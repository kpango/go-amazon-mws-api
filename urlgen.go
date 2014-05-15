package amazonmws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type AmazonMWSAPI struct {
	AccessKey     string
	SecretKey     string
	Host          string
	MarketplaceId string
	SellerId      string
}

func (api AmazonMWSAPI) genSignAndFetch2(Action string, ActionPath string, Parameters map[string]string) (string, error) {
	genUrl, err := GenerateAmazonUrl(api, Action, ActionPath, Parameters)
	if err != nil {
		return "", err
	}

	SetTimestamp(genUrl)

	signedurl, err := SignAmazonUrl(genUrl, api)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(signedurl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (api AmazonMWSAPI) genSignAndFetch(Action string, ActionPath string, Parameters map[string]string) ([]byte, error) {
	genUrl, err := GenerateAmazonUrl(api, Action, ActionPath, Parameters)
	if err != nil {
		return nil, err
	}

	SetTimestamp(genUrl)

	signedurl, err := SignAmazonUrl(genUrl, api)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(signedurl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//api.genSignAndFetch("ListOrders", "/Orders/2011-01-01", params)
func GenerateAmazonUrl(api AmazonMWSAPI, Action string, ActionPath string, Parameters map[string]string) (finalUrl *url.URL, err error) {
	result, err := url.Parse(api.Host)
	if err != nil {
		return nil, err
	}

	result.Host = api.Host
	result.Scheme = "https"
	result.Path = ActionPath

	// Generate version from ActionPath

	actionString := strings.Split(ActionPath, "/")

	values := url.Values{}
	values.Add("Action", Action)
	values.Add("AWSAccessKeyId", api.AccessKey)
	// The Report API uses Merchant instead of Seller. So we need to adjust for that.
	if len(ActionPath) > 1 {
		values.Add("SellerId", api.SellerId)
	} else {
		values.Add("Merchant", api.SellerId)
	}
	values.Add("SignatureVersion", "2")
	values.Add("SignatureMethod", "HmacSHA256")
	//values.Add("Version", "2013-09-01")
	// The Report API doesn't have an ActionPath, so we need to simply append the version number
	if len(ActionPath) > 1 {
		values.Add("Version", string(actionString[len(actionString)-1]))
	} else {
		values.Add("Version", "2009-01-01")
	}

	for k, v := range Parameters {
		values.Set(k, v)
	}

	params := values.Encode()
	result.RawQuery = params

	return result, nil
}

func SetTimestamp(origUrl *url.URL) (err error) {
	values, err := url.ParseQuery(origUrl.RawQuery)
	if err != nil {
		return err
	}
	values.Set("Timestamp", time.Now().UTC().Format(time.RFC3339))
	origUrl.RawQuery = values.Encode()

	return nil
}

func SignAmazonUrl(origUrl *url.URL, api AmazonMWSAPI) (signedUrl string, err error) {
	escapeUrl := strings.Replace(origUrl.RawQuery, ",", "%2C", -1)
	escapeUrl = strings.Replace(escapeUrl, ":", "%3A", -1)

	params := strings.Split(escapeUrl, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")

	toSign := fmt.Sprintf("GET\n%s\n%s\n%s", origUrl.Host, origUrl.Path, sortedParams)

	hasher := hmac.New(sha256.New, []byte(api.SecretKey))
	_, err = hasher.Write([]byte(toSign))
	if err != nil {
		return "", nil
	}

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	newParams := fmt.Sprintf("%s&Signature=%s", sortedParams, hash)

	origUrl.RawQuery = newParams

	return origUrl.String(), nil
}
