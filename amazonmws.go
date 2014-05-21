// amazonmws provides methods for interacting with the Amazon Marketplace Services API.
package amazonmws

import (
	"fmt"
	"strconv"
)

/*
GetLowestOfferListingsForASIN takes a list of ASINs and returns the result.
*/
func (api AmazonMWSAPI) GetLowestOfferListingsForASIN(items []string) (string, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceId)

	return api.genSignAndFetch2("GetLowestOfferListingsForASIN", "/Products/2011-10-01", params)
}

/*
GetCompetitivePricingForAsin takes a list of ASINs and returns the result.
*/
func (api AmazonMWSAPI) GetCompetitivePricingForASIN(items []string) (string, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceId)

	return api.genSignAndFetch2("GetCompetitivePricingForASIN", "/Products/2011-10-01", params)
}

func (api AmazonMWSAPI) GetMatchingProductForId(idType string, idList []string) (string, error) {
	params := make(map[string]string)

	for k, v := range idList {
		key := fmt.Sprintf("IdList.Id.%d", (k + 1))
		params[key] = string(v)
	}

	params["IdType"] = idType
	params["MarketplaceId"] = string(api.MarketplaceId)

	return api.genSignAndFetch2("GetMatchingProductForId", "/Products/2011-10-01", params)
}

/*
List orders
Takes a map of parameters, instead of specifying exactly what do to
e.g. ["CreatedAfter"], ["UpdatedBefore"]
*/
func (api AmazonMWSAPI) ListOrders(inputParameters map[string]string) ([]byte, error) {
	params := make(map[string]string)

	for k, v := range inputParameters {
		params[k] = v
	}

	params["MarketplaceId.Id.1"] = string(api.MarketplaceId)
	//params["CreatedAfter"] = string(createdAfter)
	//params["CreatedBefore"] = string(createdBefore)

	return api.genSignAndFetch("ListOrders", "/Orders/2013-09-01", params)
}

func (api AmazonMWSAPI) ListOrdersByNextToken(token string) ([]byte, error) {
	params := make(map[string]string)

	//params["MarketplaceId.Id.1"] = string(api.MarketplaceId)
	//params["CreatedAfter"] = string(createdAfter)
	params["NextToken"] = string(token)

	return api.genSignAndFetch("ListOrdersByNextToken", "/Orders/2013-09-01", params)
}

func (api AmazonMWSAPI) GetOrder(orderId string) ([]byte, error) {
	params := make(map[string]string)

	//params["MarketplaceId.Id.1"] = string(api.MarketplaceId)
	//params["CreatedAfter"] = string(createdAfter)
	params["AmazonOrderId.Id.1"] = string(orderId)

	return api.genSignAndFetch("GetOrder", "/Orders/2013-09-01", params)
}

func (api AmazonMWSAPI) ListOrderItems(orderId string) ([]byte, error) {
	params := make(map[string]string)

	//params["MarketplaceId.Id.1"] = string(api.MarketplaceId)
	//params["CreatedAfter"] = string(createdAfter)
	params["AmazonOrderId"] = string(orderId)

	return api.genSignAndFetch("ListOrderItems", "/Orders/2013-09-01", params)
}

func (api AmazonMWSAPI) GetReportList(reportType string) ([]byte, error) {
	params := make(map[string]string)

	//params["MarketplaceId.Id.1"] = string(api.MarketplaceId)
	//params["CreatedAfter"] = string(createdAfter)
	params["ReportTypeList.Type.1"] = string(reportType)
	// /Reports/2009-01-01

	return api.genSignAndFetch("GetReportList", "/", params)
}

func (api AmazonMWSAPI) GetReportListByNextToken(token string) ([]byte, error) {
	params := make(map[string]string)

	//params["MarketplaceId.Id.1"] = string(api.MarketplaceId)
	//params["CreatedAfter"] = string(createdAfter)
	params["NextToken"] = string(token)
	// /Reports/2009-01-01

	return api.genSignAndFetch("GetReportListByNextToken", "/", params)
}

func (api AmazonMWSAPI) GetReport(reportID string) ([]byte, error) {
	params := make(map[string]string)

	//params["MarketplaceId.Id.1"] = string(api.MarketplaceId)
	//params["CreatedAfter"] = string(createdAfter)
	params["ReportId"] = string(reportID)
	// /Reports/2009-01-01

	return api.genSignAndFetch("GetReport", "/", params)
}

/*
ListInventorySupply
Takes an string array of SKUs
Returns a byte array in response
*/
func (api AmazonMWSAPI) ListInventorySupply(SKUs []string, ResponseGroup string) ([]byte, error) {
	params := make(map[string]string)

	//params["MarketplaceId.Id.1"] = string(api.MarketplaceId)
	//params["CreatedAfter"] = string(createdAfter)
	//params["ReportId"] = string(reportID)
	// /Reports/2009-01-01
	params["ResponseGroup"] = ResponseGroup

	for index, sku := range SKUs {
		params["SellerSkus.member."+strconv.Itoa(index+1)] = string(sku)
	}

	return api.genSignAndFetch("ListInventorySupply", "/FulfillmentInventory/2010-10-01", params)
}

/*
ListInventorySupplyByNextToken
Takes a string token for next page of supply listings
Returns a byte array in response
*/
func (api AmazonMWSAPI) ListInventorySupplyByNextToken(token string) ([]byte, error) {
	params := make(map[string]string)

	params["NextToken"] = token

	return api.genSignAndFetch("ListInventorySupplyByNextToken", "/FulfillmentInventory/2010-10-01", params)
}
