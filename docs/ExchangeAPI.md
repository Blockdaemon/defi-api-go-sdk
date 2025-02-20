# \ExchangeAPI

All URIs are relative to *https://svc.blockdaemon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRoutes**](ExchangeAPI.md#GetRoutes) | **Get** /exchange/routes | Get a list of routes for swapping assets
[**GetStatus**](ExchangeAPI.md#GetStatus) | **Get** /exchange/status | Get swap status



## GetRoutes

> RoutesResponse GetRoutes(ctx).FromChain(fromChain).FromAmount(fromAmount).FromToken(fromToken).ToChain(toChain).ToToken(toToken).FromAddress(fromAddress).ToAddress(toAddress).Slippage(slippage).AllowBridges(allowBridges).AllowExchanges(allowExchanges).DenyBridges(denyBridges).DenyExchanges(denyExchanges).Execute()

Get a list of routes for swapping assets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Blockdaemon/defi-api-go-sdk"
)

func main() {
	fromChain := "eip155:1" // string | The sending blockchain in CAIP-2 notation.
	fromAmount := "3000000000000000" // string | The amount of tokens to be sent, including decimals.
	fromToken := "0x0000000000000000000000000000000000000000" // string | The address or symbol of the token to be transferred.
	toChain := "eip155:10" // string | The receiving blockchain in CAIP-2 notation.
	toToken := "0x0b2c639c533813f4aa9d7837caf62653d097ff85" // string | The address or symbol of the token to be received.
	fromAddress := "0xb8DC8a293fb4c6677dE0A29E505d33314eBB9161" // string | The address of the wallet sending the tokens.
	toAddress := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045" // string | The address of the wallet receiving the tokens. This can be the same as the `fromAddress`. `fromAddress` will be used.
	slippage := float32(0.001) // float32 | The maximum allowed price slippage for the transaction, as a decimal fraction. (default to 0.01)
	allowBridges := []string{"Inner_example"} // []string | A list of bridges that are allowed for the transaction. (optional)
	allowExchanges := []string{"Inner_example"} // []string | A list of exchanges that are allowed for the transaction. (optional)
	denyBridges := []string{"Inner_example"} // []string | A list of bridges that are not allowed for the transaction. (optional)
	denyExchanges := []string{"Inner_example"} // []string | A list of exchanges that are not allowed for the transaction. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExchangeAPI.GetRoutes(context.Background()).FromChain(fromChain).FromAmount(fromAmount).FromToken(fromToken).ToChain(toChain).ToToken(toToken).FromAddress(fromAddress).ToAddress(toAddress).Slippage(slippage).AllowBridges(allowBridges).AllowExchanges(allowExchanges).DenyBridges(denyBridges).DenyExchanges(denyExchanges).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeAPI.GetRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoutes`: RoutesResponse
	fmt.Fprintf(os.Stdout, "Response from `ExchangeAPI.GetRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromChain** | **string** | The sending blockchain in CAIP-2 notation. | 
 **fromAmount** | **string** | The amount of tokens to be sent, including decimals. | 
 **fromToken** | **string** | The address or symbol of the token to be transferred. | 
 **toChain** | **string** | The receiving blockchain in CAIP-2 notation. | 
 **toToken** | **string** | The address or symbol of the token to be received. | 
 **fromAddress** | **string** | The address of the wallet sending the tokens. | 
 **toAddress** | **string** | The address of the wallet receiving the tokens. This can be the same as the &#x60;fromAddress&#x60;. &#x60;fromAddress&#x60; will be used. | 
 **slippage** | **float32** | The maximum allowed price slippage for the transaction, as a decimal fraction. | [default to 0.01]
 **allowBridges** | **[]string** | A list of bridges that are allowed for the transaction. | 
 **allowExchanges** | **[]string** | A list of exchanges that are allowed for the transaction. | 
 **denyBridges** | **[]string** | A list of bridges that are not allowed for the transaction. | 
 **denyExchanges** | **[]string** | A list of exchanges that are not allowed for the transaction. | 

### Return type

[**RoutesResponse**](RoutesResponse.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> StatusResponse GetStatus(ctx).TransactionID(transactionID).TargetID(targetID).FromChain(fromChain).ToChain(toChain).Execute()

Get swap status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Blockdaemon/defi-api-go-sdk"
)

func main() {
	transactionID := "0xced78f9a9a3ad2816b6ba573aaf93e5c8685f10e1ef827c0b31090ed23746987" // string | The ID of the transaction.
	targetID := "906180c0-7ade-5710-905d-2583c3ab71eb" // string | A reference to the transaction target used. This value can be found in the routes response.
	fromChain := "eip155:10" // string | The sending Chain ID or symbol.
	toChain := "eip155:10" // string | The receiver of the transaction (Chain ID). This can be the same as the `fromChain` for same-chain swaps.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExchangeAPI.GetStatus(context.Background()).TransactionID(transactionID).TargetID(targetID).FromChain(fromChain).ToChain(toChain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeAPI.GetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatus`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ExchangeAPI.GetStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionID** | **string** | The ID of the transaction. | 
 **targetID** | **string** | A reference to the transaction target used. This value can be found in the routes response. | 
 **fromChain** | **string** | The sending Chain ID or symbol. | 
 **toChain** | **string** | The receiver of the transaction (Chain ID). This can be the same as the &#x60;fromChain&#x60; for same-chain swaps. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

