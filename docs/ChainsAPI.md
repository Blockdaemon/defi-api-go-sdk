# \ChainsAPI

All URIs are relative to *https://svc.blockdaemon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChains**](ChainsAPI.md#GetChains) | **Get** /chains | Get supported blockchain networks with metadata



## GetChains

> Chains GetChains(ctx).ChainID(chainID).ChainName(chainName).NetworkType(networkType).ChainType(chainType).Execute()

Get supported blockchain networks with metadata



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
	chainType := openapiclient.ChainType("evm") // ChainType | The type of chain (e.g evm, or aptos). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.GetChains(context.Background()).ChainID(chainID).ChainName(chainName).NetworkType(networkType).ChainType(chainType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.GetChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChains`: Chains
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.GetChains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainID** | **string** | The CAIP-2 identifier for a blockchain. | 
 **chainName** | **string** | The name of a blockchain. | 
 **networkType** | [**NetworkType**](NetworkType.md) | The type of network (mainnet, testnet, or devnet). | 
 **chainType** | [**ChainType**](ChainType.md) | The type of chain (e.g evm, or aptos). | 

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

