# TokenApprovalModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainID** | **string** | The unique identifier of a blockchain in CAIP-2 notation. | 
**AccountAddress** | **string** | The address of an account. | 
**SpenderAddress** | **string** | The address of an account. | 
**TokenAddress** | **string** | The address of an account. | 
**ToApprovedAmount** | **string** | The amount of a token represented as a string. | 

## Methods

### NewTokenApprovalModification

`func NewTokenApprovalModification(chainID string, accountAddress string, spenderAddress string, tokenAddress string, toApprovedAmount string, ) *TokenApprovalModification`

NewTokenApprovalModification instantiates a new TokenApprovalModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenApprovalModificationWithDefaults

`func NewTokenApprovalModificationWithDefaults() *TokenApprovalModification`

NewTokenApprovalModificationWithDefaults instantiates a new TokenApprovalModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainID

`func (o *TokenApprovalModification) GetChainID() string`

GetChainID returns the ChainID field if non-nil, zero value otherwise.

### GetChainIDOk

`func (o *TokenApprovalModification) GetChainIDOk() (*string, bool)`

GetChainIDOk returns a tuple with the ChainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainID

`func (o *TokenApprovalModification) SetChainID(v string)`

SetChainID sets ChainID field to given value.


### GetAccountAddress

`func (o *TokenApprovalModification) GetAccountAddress() string`

GetAccountAddress returns the AccountAddress field if non-nil, zero value otherwise.

### GetAccountAddressOk

`func (o *TokenApprovalModification) GetAccountAddressOk() (*string, bool)`

GetAccountAddressOk returns a tuple with the AccountAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAddress

`func (o *TokenApprovalModification) SetAccountAddress(v string)`

SetAccountAddress sets AccountAddress field to given value.


### GetSpenderAddress

`func (o *TokenApprovalModification) GetSpenderAddress() string`

GetSpenderAddress returns the SpenderAddress field if non-nil, zero value otherwise.

### GetSpenderAddressOk

`func (o *TokenApprovalModification) GetSpenderAddressOk() (*string, bool)`

GetSpenderAddressOk returns a tuple with the SpenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpenderAddress

`func (o *TokenApprovalModification) SetSpenderAddress(v string)`

SetSpenderAddress sets SpenderAddress field to given value.


### GetTokenAddress

`func (o *TokenApprovalModification) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TokenApprovalModification) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TokenApprovalModification) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetToApprovedAmount

`func (o *TokenApprovalModification) GetToApprovedAmount() string`

GetToApprovedAmount returns the ToApprovedAmount field if non-nil, zero value otherwise.

### GetToApprovedAmountOk

`func (o *TokenApprovalModification) GetToApprovedAmountOk() (*string, bool)`

GetToApprovedAmountOk returns a tuple with the ToApprovedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToApprovedAmount

`func (o *TokenApprovalModification) SetToApprovedAmount(v string)`

SetToApprovedAmount sets ToApprovedAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


