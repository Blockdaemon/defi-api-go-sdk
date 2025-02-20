# ChainBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainID** | **string** | The unique identifier of a blockchain in CAIP-2 notation. | 
**TokenBalances** | [**[]TokenAmount**](TokenAmount.md) |  | 

## Methods

### NewChainBalance

`func NewChainBalance(chainID string, tokenBalances []TokenAmount, ) *ChainBalance`

NewChainBalance instantiates a new ChainBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainBalanceWithDefaults

`func NewChainBalanceWithDefaults() *ChainBalance`

NewChainBalanceWithDefaults instantiates a new ChainBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainID

`func (o *ChainBalance) GetChainID() string`

GetChainID returns the ChainID field if non-nil, zero value otherwise.

### GetChainIDOk

`func (o *ChainBalance) GetChainIDOk() (*string, bool)`

GetChainIDOk returns a tuple with the ChainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainID

`func (o *ChainBalance) SetChainID(v string)`

SetChainID sets ChainID field to given value.


### GetTokenBalances

`func (o *ChainBalance) GetTokenBalances() []TokenAmount`

GetTokenBalances returns the TokenBalances field if non-nil, zero value otherwise.

### GetTokenBalancesOk

`func (o *ChainBalance) GetTokenBalancesOk() (*[]TokenAmount, bool)`

GetTokenBalancesOk returns a tuple with the TokenBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBalances

`func (o *ChainBalance) SetTokenBalances(v []TokenAmount)`

SetTokenBalances sets TokenBalances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


