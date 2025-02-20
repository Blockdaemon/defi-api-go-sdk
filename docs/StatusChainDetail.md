# StatusChainDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**StatusEnum**](StatusEnum.md) |  | 
**TxHash** | **string** | The hash of a transaction or block. | 
**ExplorerLink** | **string** | A link to a block explorer for a transaction or address. | 
**ChainID** | **string** | The unique identifier of a blockchain in CAIP-2 notation. | 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewStatusChainDetail

`func NewStatusChainDetail(status StatusEnum, txHash string, explorerLink string, chainID string, ) *StatusChainDetail`

NewStatusChainDetail instantiates a new StatusChainDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusChainDetailWithDefaults

`func NewStatusChainDetailWithDefaults() *StatusChainDetail`

NewStatusChainDetailWithDefaults instantiates a new StatusChainDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StatusChainDetail) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatusChainDetail) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatusChainDetail) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetTxHash

`func (o *StatusChainDetail) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *StatusChainDetail) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *StatusChainDetail) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetExplorerLink

`func (o *StatusChainDetail) GetExplorerLink() string`

GetExplorerLink returns the ExplorerLink field if non-nil, zero value otherwise.

### GetExplorerLinkOk

`func (o *StatusChainDetail) GetExplorerLinkOk() (*string, bool)`

GetExplorerLinkOk returns a tuple with the ExplorerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorerLink

`func (o *StatusChainDetail) SetExplorerLink(v string)`

SetExplorerLink sets ExplorerLink field to given value.


### GetChainID

`func (o *StatusChainDetail) GetChainID() string`

GetChainID returns the ChainID field if non-nil, zero value otherwise.

### GetChainIDOk

`func (o *StatusChainDetail) GetChainIDOk() (*string, bool)`

GetChainIDOk returns a tuple with the ChainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainID

`func (o *StatusChainDetail) SetChainID(v string)`

SetChainID sets ChainID field to given value.


### GetExtensions

`func (o *StatusChainDetail) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *StatusChainDetail) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *StatusChainDetail) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *StatusChainDetail) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


