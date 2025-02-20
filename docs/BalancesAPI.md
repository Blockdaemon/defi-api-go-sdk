# \BalancesAPI

All URIs are relative to *https://svc.blockdaemon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalances**](BalancesAPI.md#GetBalances) | **Get** /balances | Get token balances and market data for an account
[**GetSupportedChainsForBalances**](BalancesAPI.md#GetSupportedChainsForBalances) | **Get** /balances/supported-chains | Get supported chains for token balance queries



## GetBalances

> BalancesResponse GetBalances(ctx).AccountAddress(accountAddress).ChainIDs(chainIDs).TokenAddress(tokenAddress).Execute()

Get token balances and market data for an account



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
	accountAddress := "0xb8DC8a293fb4c6677dE0A29E505d33314eBB9161" // string | The address of an account.
	chainIDs := []string{"Inner_example"} // []string | A list of CAIP-2 identifiers for blockchains.
	tokenAddress := "0x7F5c764cBc14f9669B88837ca1490cCa17c31607" // string | The address of an token contract to query. If omitted, all supported tokens will be queried. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BalancesAPI.GetBalances(context.Background()).AccountAddress(accountAddress).ChainIDs(chainIDs).TokenAddress(tokenAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalancesAPI.GetBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalances`: BalancesResponse
	fmt.Fprintf(os.Stdout, "Response from `BalancesAPI.GetBalances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountAddress** | **string** | The address of an account. | 
 **chainIDs** | **[]string** | A list of CAIP-2 identifiers for blockchains. | 
 **tokenAddress** | **string** | The address of an token contract to query. If omitted, all supported tokens will be queried. | 

### Return type

[**BalancesResponse**](BalancesResponse.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedChainsForBalances

> Chains GetSupportedChainsForBalances(ctx).ChainID(chainID).ChainName(chainName).NetworkType(networkType).Execute()

Get supported chains for token balance queries



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
	chainID := "eip155:1" // string | The CAIP-2 identifier for a blockchain. (optional)
	chainName := "Ethereum" // string | The name of a blockchain. (optional)
	networkType := openapiclient.NetworkType("mainnet") // NetworkType | The type of network (mainnet, testnet, or devnet). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BalancesAPI.GetSupportedChainsForBalances(context.Background()).ChainID(chainID).ChainName(chainName).NetworkType(networkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalancesAPI.GetSupportedChainsForBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportedChainsForBalances`: Chains
	fmt.Fprintf(os.Stdout, "Response from `BalancesAPI.GetSupportedChainsForBalances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedChainsForBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainID** | **string** | The CAIP-2 identifier for a blockchain. | 
 **chainName** | **string** | The name of a blockchain. | 
 **networkType** | [**NetworkType**](NetworkType.md) | The type of network (mainnet, testnet, or devnet). | 

### Return type

[**Chains**](Chains.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

