# Chain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainName** | **string** | The name of a blockchain. | 
**ChainID** | **string** | The unique identifier of a blockchain in CAIP-2 notation. | 
**NetworkType** | Pointer to [**NetworkType**](NetworkType.md) |  | [optional] 
**ChainIconURI** | Pointer to **string** | The type of network (mainnet, testnet, or devnet). | [optional] 
**BlockTime** | Pointer to **int32** | estimated block confirmation time in seconds | [optional] 
**BlockConfirmations** | Pointer to **int32** | number of additional blocks required to ensure the transaction is not reversed | [optional] 
**Rpc** | **[]string** | List of RPC gateway URLs. | 
**NativeCurrency** | Pointer to [**NativeCurrency**](NativeCurrency.md) |  | [optional] 
**BlockExplorerUrls** | Pointer to **[]string** | List of block explorer URLs. | [optional] 

## Methods

### NewChain

`func NewChain(chainName string, chainID string, rpc []string, ) *Chain`

NewChain instantiates a new Chain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainWithDefaults

`func NewChainWithDefaults() *Chain`

NewChainWithDefaults instantiates a new Chain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainName

`func (o *Chain) GetChainName() string`

GetChainName returns the ChainName field if non-nil, zero value otherwise.

### GetChainNameOk

`func (o *Chain) GetChainNameOk() (*string, bool)`

GetChainNameOk returns a tuple with the ChainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainName

`func (o *Chain) SetChainName(v string)`

SetChainName sets ChainName field to given value.


### GetChainID

`func (o *Chain) GetChainID() string`

GetChainID returns the ChainID field if non-nil, zero value otherwise.

### GetChainIDOk

`func (o *Chain) GetChainIDOk() (*string, bool)`

GetChainIDOk returns a tuple with the ChainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainID

`func (o *Chain) SetChainID(v string)`

SetChainID sets ChainID field to given value.


### GetNetworkType

`func (o *Chain) GetNetworkType() NetworkType`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *Chain) GetNetworkTypeOk() (*NetworkType, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *Chain) SetNetworkType(v NetworkType)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *Chain) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetChainIconURI

`func (o *Chain) GetChainIconURI() string`

GetChainIconURI returns the ChainIconURI field if non-nil, zero value otherwise.

### GetChainIconURIOk

`func (o *Chain) GetChainIconURIOk() (*string, bool)`

GetChainIconURIOk returns a tuple with the ChainIconURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainIconURI

`func (o *Chain) SetChainIconURI(v string)`

SetChainIconURI sets ChainIconURI field to given value.

### HasChainIconURI

`func (o *Chain) HasChainIconURI() bool`

HasChainIconURI returns a boolean if a field has been set.

### GetBlockTime

`func (o *Chain) GetBlockTime() int32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *Chain) GetBlockTimeOk() (*int32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *Chain) SetBlockTime(v int32)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *Chain) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetBlockConfirmations

`func (o *Chain) GetBlockConfirmations() int32`

GetBlockConfirmations returns the BlockConfirmations field if non-nil, zero value otherwise.

### GetBlockConfirmationsOk

`func (o *Chain) GetBlockConfirmationsOk() (*int32, bool)`

GetBlockConfirmationsOk returns a tuple with the BlockConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockConfirmations

`func (o *Chain) SetBlockConfirmations(v int32)`

SetBlockConfirmations sets BlockConfirmations field to given value.

### HasBlockConfirmations

`func (o *Chain) HasBlockConfirmations() bool`

HasBlockConfirmations returns a boolean if a field has been set.

### GetRpc

`func (o *Chain) GetRpc() []string`

GetRpc returns the Rpc field if non-nil, zero value otherwise.

### GetRpcOk

`func (o *Chain) GetRpcOk() (*[]string, bool)`

GetRpcOk returns a tuple with the Rpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpc

`func (o *Chain) SetRpc(v []string)`

SetRpc sets Rpc field to given value.


### GetNativeCurrency

`func (o *Chain) GetNativeCurrency() NativeCurrency`

GetNativeCurrency returns the NativeCurrency field if non-nil, zero value otherwise.

### GetNativeCurrencyOk

`func (o *Chain) GetNativeCurrencyOk() (*NativeCurrency, bool)`

GetNativeCurrencyOk returns a tuple with the NativeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeCurrency

`func (o *Chain) SetNativeCurrency(v NativeCurrency)`

SetNativeCurrency sets NativeCurrency field to given value.

### HasNativeCurrency

`func (o *Chain) HasNativeCurrency() bool`

HasNativeCurrency returns a boolean if a field has been set.

### GetBlockExplorerUrls

`func (o *Chain) GetBlockExplorerUrls() []string`

GetBlockExplorerUrls returns the BlockExplorerUrls field if non-nil, zero value otherwise.

### GetBlockExplorerUrlsOk

`func (o *Chain) GetBlockExplorerUrlsOk() (*[]string, bool)`

GetBlockExplorerUrlsOk returns a tuple with the BlockExplorerUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExplorerUrls

`func (o *Chain) SetBlockExplorerUrls(v []string)`

SetBlockExplorerUrls sets BlockExplorerUrls field to given value.

### HasBlockExplorerUrls

`func (o *Chain) HasBlockExplorerUrls() bool`

HasBlockExplorerUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


