This Amazon MWS API client is based heavily on both @DDRBoxman's [go-amazon-product-api](https://github.com/DDRBoxman/go-amazon-product-api) and @ezkl's [go-amazon-mws-api](https://github.com/ezkl/go-amazon-mws-api).

## Documentation

You can view the auto-generated documentation at http://godoc.org/github.com/nickrobison/go-amazon-mws-api

The primary changes are better error handling, expanded API support, selectable API versions, and moving many of the declared arguments into maps, to provide better flexibility.

## Supported APIs

+ Orders
  + ListOrders
  + ListOrdersByNextToken
  + ListOrderItems
  + GetOrder
+ Reports
  + GetReportList
  + GetReportListByNextToken
  + GetReport
+ Fulfillment
  + ListInventorySupply
+ Products (not fully updated to new URL call, still returns string)
  + GetCompetitivePricingForASIN
  + GetMatchingProductForId
+ Fulfillment
  + ListInventorySupply
  + ListInventorySupplyByNextToken

## TODO
- [ ] Move APIs to Map based parameters
- [ ] Update all APIs to new URL call
- [ ] Update generated godoc
- [ ] Cleanup Comments

## Usage

```go
package main

import (
       "fmt"
       "github.com/nickrobison/go-amazon-mws-api"
)

func main() {
       var api amazonmws.AmazonMWSAPI

       api.AccessKey = ""
       api.SecretKey = ""
       api.Host = "mws.amazonservices.com"
       api.MarketplaceId = "ATVPDKIKX0DER"
       api.SellerId = ""

       asins := []string{"0195019199"}

       result,err := api.GetLowestOfferListingsForASIN(asins)

       if (err != nil) {
           fmt.Println(err)
       }

       fmt.Println(result)
}
```

