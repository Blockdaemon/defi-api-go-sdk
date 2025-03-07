# \ApprovalsAPI

All URIs are relative to *https://svc.blockdaemon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTokenApproval**](ApprovalsAPI.md#DeleteTokenApproval) | **Delete** /approval | Delete an ERC20 token approval
[**GetAllApprovals**](ApprovalsAPI.md#GetAllApprovals) | **Get** /approvals | Get a list of ERC20 token approvals
[**GetSupportedChains**](ApprovalsAPI.md#GetSupportedChains) | **Get** /approvals/supported-chains | Get supported chains for token approvals
[**GetTokenApproval**](ApprovalsAPI.md#GetTokenApproval) | **Get** /approval | List ERC20 token approval
[**ModifyTokenApproval**](ApprovalsAPI.md#ModifyTokenApproval) | **Post** /approval | Modify an ERC20 token approval



## DeleteTokenApproval

> TokenApprovalResponse DeleteTokenApproval(ctx).TokenApprovalDeletion(tokenApprovalDeletion).Execute()

Delete an ERC20 token approval



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
	tokenApprovalDeletion := *openapiclient.NewTokenApprovalDeletion("ChainID_example", "0xf271AAFC62634e6Dc9A276ac0f6145C4fDbE2Ced", "0xf271AAFC62634e6Dc9A276ac0f6145C4fDbE2Ced", "0xf271AAFC62634e6Dc9A276ac0f6145C4fDbE2Ced") // TokenApprovalDeletion | The request body for a delete approval

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApprovalsAPI.DeleteTokenApproval(context.Background()).TokenApprovalDeletion(tokenApprovalDeletion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsAPI.DeleteTokenApproval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTokenApproval`: TokenApprovalResponse
	fmt.Fprintf(os.Stdout, "Response from `ApprovalsAPI.DeleteTokenApproval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenApprovalDeletion** | [**TokenApprovalDeletion**](TokenApprovalDeletion.md) | The request body for a delete approval | 

### Return type

[**TokenApprovalResponse**](TokenApprovalResponse.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllApprovals

> AllApprovalsResponse GetAllApprovals(ctx).ChainIDs(chainIDs).AccountAddresses(accountAddresses).SpenderAddresses(spenderAddresses).TokenAddresses(tokenAddresses).Execute()

Get a list of ERC20 token approvals



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
	chainIDs := []string{"Inner_example"} // []string | A list of CAIP-2 identifiers for blockchains.
	accountAddresses := []string{"0xf271AAFC62634e6Dc9A276ac0f6145C4fDbE2Ced"} // []string | A list of account addresses to query.
	spenderAddresses := []string{"0xf271AAFC62634e6Dc9A276ac0f6145C4fDbE2Ced"} // []string | A list of spender addresses to query. If omitted, all supported spenders will be queried. (optional)
	tokenAddresses := []string{"0xf271AAFC62634e6Dc9A276ac0f6145C4fDbE2Ced"} // []string | A list of token contract addresses to query. If omitted, all supported tokens will be queried. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApprovalsAPI.GetAllApprovals(context.Background()).ChainIDs(chainIDs).AccountAddresses(accountAddresses).SpenderAddresses(spenderAddresses).TokenAddresses(tokenAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsAPI.GetAllApprovals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllApprovals`: AllApprovalsResponse
	fmt.Fprintf(os.Stdout, "Response from `ApprovalsAPI.GetAllApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainIDs** | **[]string** | A list of CAIP-2 identifiers for blockchains. | 
 **accountAddresses** | **[]string** | A list of account addresses to query. | 
 **spenderAddresses** | **[]string** | A list of spender addresses to query. If omitted, all supported spenders will be queried. | 
 **tokenAddresses** | **[]string** | A list of token contract addresses to query. If omitted, all supported tokens will be queried. | 

### Return type

[**AllApprovalsResponse**](AllApprovalsResponse.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedChains

> Chains GetSupportedChains(ctx).ChainID(chainID).ChainName(chainName).NetworkType(networkType).Execute()

Get supported chains for token approvals



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
	resp, r, err := apiClient.ApprovalsAPI.GetSupportedChains(context.Background()).ChainID(chainID).ChainName(chainName).NetworkType(networkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsAPI.GetSupportedChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportedChains`: Chains
	fmt.Fprintf(os.Stdout, "Response from `ApprovalsAPI.GetSupportedChains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedChainsRequest struct via the builder pattern


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


## GetTokenApproval

> TokenApprovalData GetTokenApproval(ctx).ChainID(chainID).AccountAddress(accountAddress).TokenAddress(tokenAddress).SpenderAddress(spenderAddress).Execute()

List ERC20 token approval



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
	chainID := "eip155:10" // string | The CAIP-2 identifier for a blockchain.
	accountAddress := "0xb8DC8a293fb4c6677dE0A29E505d33314eBB9161" // string | The address of an account.
	tokenAddress := "0x7F5c764cBc14f9669B88837ca1490cCa17c31607" // string | The address of an token contract.
	spenderAddress := "0xB0D502E938ed5f4df2E681fE6E419ff29631d62b" // string | The address of a spender.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApprovalsAPI.GetTokenApproval(context.Background()).ChainID(chainID).AccountAddress(accountAddress).TokenAddress(tokenAddress).SpenderAddress(spenderAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsAPI.GetTokenApproval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenApproval`: TokenApprovalData
	fmt.Fprintf(os.Stdout, "Response from `ApprovalsAPI.GetTokenApproval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainID** | **string** | The CAIP-2 identifier for a blockchain. | 
 **accountAddress** | **string** | The address of an account. | 
 **tokenAddress** | **string** | The address of an token contract. | 
 **spenderAddress** | **string** | The address of a spender. | 

### Return type

[**TokenApprovalData**](TokenApprovalData.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTokenApproval

> TokenApprovalResponse ModifyTokenApproval(ctx).TokenApprovalModification(tokenApprovalModification).Execute()

Modify an ERC20 token approval



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
	tokenApprovalModification := *openapiclient.NewTokenApprovalModification("ChainID_example", "0xf271AAFC62634e6Dc9A276ac0f6145C4fDbE2Ced", "0xf271AAFC62634e6Dc9A276ac0f6145C4fDbE2Ced", "0xf271AAFC62634e6Dc9A276ac0f6145C4fDbE2Ced", "ToApprovedAmount_example") // TokenApprovalModification | The request body for a post approval

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApprovalsAPI.ModifyTokenApproval(context.Background()).TokenApprovalModification(tokenApprovalModification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsAPI.ModifyTokenApproval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTokenApproval`: TokenApprovalResponse
	fmt.Fprintf(os.Stdout, "Response from `ApprovalsAPI.ModifyTokenApproval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyTokenApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenApprovalModification** | [**TokenApprovalModification**](TokenApprovalModification.md) | The request body for a post approval | 

### Return type

[**TokenApprovalResponse**](TokenApprovalResponse.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

