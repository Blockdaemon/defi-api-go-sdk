# TokenApprovalData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** | The unique identifier of a blockchain in CAIP-2 notation. | 
**AccountAddress** | **string** | The address of an account. | 
**SpenderAddress** | **string** | The address of an account. | 
**Token** | [**Token**](Token.md) |  | 
**Amount** | **string** | The amount of a token represented as a string. | 
**AmountUSD** | Pointer to **string** | The amount of a token represented as a string. | [optional] 
**ValueAtRisk** | **string** | The amount of a token represented as a string. | 
**ValueAtRiskUSD** | Pointer to **string** | The amount of a token represented as a string. | [optional] 
**SpenderName** | Pointer to **string** | The address of an account. | [optional] 

## Methods

### NewTokenApprovalData

`func NewTokenApprovalData(chainId string, accountAddress string, spenderAddress string, token Token, amount string, valueAtRisk string, ) *TokenApprovalData`

NewTokenApprovalData instantiates a new TokenApprovalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenApprovalDataWithDefaults

`func NewTokenApprovalDataWithDefaults() *TokenApprovalData`

NewTokenApprovalDataWithDefaults instantiates a new TokenApprovalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *TokenApprovalData) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenApprovalData) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenApprovalData) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetAccountAddress

`func (o *TokenApprovalData) GetAccountAddress() string`

GetAccountAddress returns the AccountAddress field if non-nil, zero value otherwise.

### GetAccountAddressOk

`func (o *TokenApprovalData) GetAccountAddressOk() (*string, bool)`

GetAccountAddressOk returns a tuple with the AccountAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAddress

`func (o *TokenApprovalData) SetAccountAddress(v string)`

SetAccountAddress sets AccountAddress field to given value.


### GetSpenderAddress

`func (o *TokenApprovalData) GetSpenderAddress() string`

GetSpenderAddress returns the SpenderAddress field if non-nil, zero value otherwise.

### GetSpenderAddressOk

`func (o *TokenApprovalData) GetSpenderAddressOk() (*string, bool)`

GetSpenderAddressOk returns a tuple with the SpenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpenderAddress

`func (o *TokenApprovalData) SetSpenderAddress(v string)`

SetSpenderAddress sets SpenderAddress field to given value.


### GetToken

`func (o *TokenApprovalData) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenApprovalData) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenApprovalData) SetToken(v Token)`

SetToken sets Token field to given value.


### GetAmount

`func (o *TokenApprovalData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenApprovalData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenApprovalData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAmountUSD

`func (o *TokenApprovalData) GetAmountUSD() string`

GetAmountUSD returns the AmountUSD field if non-nil, zero value otherwise.

### GetAmountUSDOk

`func (o *TokenApprovalData) GetAmountUSDOk() (*string, bool)`

GetAmountUSDOk returns a tuple with the AmountUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountUSD

`func (o *TokenApprovalData) SetAmountUSD(v string)`

SetAmountUSD sets AmountUSD field to given value.

### HasAmountUSD

`func (o *TokenApprovalData) HasAmountUSD() bool`

HasAmountUSD returns a boolean if a field has been set.

### GetValueAtRisk

`func (o *TokenApprovalData) GetValueAtRisk() string`

GetValueAtRisk returns the ValueAtRisk field if non-nil, zero value otherwise.

### GetValueAtRiskOk

`func (o *TokenApprovalData) GetValueAtRiskOk() (*string, bool)`

GetValueAtRiskOk returns a tuple with the ValueAtRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAtRisk

`func (o *TokenApprovalData) SetValueAtRisk(v string)`

SetValueAtRisk sets ValueAtRisk field to given value.


### GetValueAtRiskUSD

`func (o *TokenApprovalData) GetValueAtRiskUSD() string`

GetValueAtRiskUSD returns the ValueAtRiskUSD field if non-nil, zero value otherwise.

### GetValueAtRiskUSDOk

`func (o *TokenApprovalData) GetValueAtRiskUSDOk() (*string, bool)`

GetValueAtRiskUSDOk returns a tuple with the ValueAtRiskUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAtRiskUSD

`func (o *TokenApprovalData) SetValueAtRiskUSD(v string)`

SetValueAtRiskUSD sets ValueAtRiskUSD field to given value.

### HasValueAtRiskUSD

`func (o *TokenApprovalData) HasValueAtRiskUSD() bool`

HasValueAtRiskUSD returns a boolean if a field has been set.

### GetSpenderName

`func (o *TokenApprovalData) GetSpenderName() string`

GetSpenderName returns the SpenderName field if non-nil, zero value otherwise.

### GetSpenderNameOk

`func (o *TokenApprovalData) GetSpenderNameOk() (*string, bool)`

GetSpenderNameOk returns a tuple with the SpenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpenderName

`func (o *TokenApprovalData) SetSpenderName(v string)`

SetSpenderName sets SpenderName field to given value.

### HasSpenderName

`func (o *TokenApprovalData) HasSpenderName() bool`

HasSpenderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


