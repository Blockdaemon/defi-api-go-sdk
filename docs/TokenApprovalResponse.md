# TokenApprovalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainID** | **string** | The unique identifier of a blockchain in CAIP-2 notation. | 
**AccountAddress** | **string** | The address of an account. | 
**SpenderAddress** | **string** | The address of an account. | 
**Token** | [**Token**](Token.md) |  | 
**FromApprovedAmount** | **string** | The amount of a token represented as a string. | 
**ToApprovedAmount** | **string** | The amount of a token represented as a string. | 
**TransactionRequest** | [**TransactionRequest**](TransactionRequest.md) |  | 

## Methods

### NewTokenApprovalResponse

`func NewTokenApprovalResponse(chainID string, accountAddress string, spenderAddress string, token Token, fromApprovedAmount string, toApprovedAmount string, transactionRequest TransactionRequest, ) *TokenApprovalResponse`

NewTokenApprovalResponse instantiates a new TokenApprovalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenApprovalResponseWithDefaults

`func NewTokenApprovalResponseWithDefaults() *TokenApprovalResponse`

NewTokenApprovalResponseWithDefaults instantiates a new TokenApprovalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainID

`func (o *TokenApprovalResponse) GetChainID() string`

GetChainID returns the ChainID field if non-nil, zero value otherwise.

### GetChainIDOk

`func (o *TokenApprovalResponse) GetChainIDOk() (*string, bool)`

GetChainIDOk returns a tuple with the ChainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainID

`func (o *TokenApprovalResponse) SetChainID(v string)`

SetChainID sets ChainID field to given value.


### GetAccountAddress

`func (o *TokenApprovalResponse) GetAccountAddress() string`

GetAccountAddress returns the AccountAddress field if non-nil, zero value otherwise.

### GetAccountAddressOk

`func (o *TokenApprovalResponse) GetAccountAddressOk() (*string, bool)`

GetAccountAddressOk returns a tuple with the AccountAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAddress

`func (o *TokenApprovalResponse) SetAccountAddress(v string)`

SetAccountAddress sets AccountAddress field to given value.


### GetSpenderAddress

`func (o *TokenApprovalResponse) GetSpenderAddress() string`

GetSpenderAddress returns the SpenderAddress field if non-nil, zero value otherwise.

### GetSpenderAddressOk

`func (o *TokenApprovalResponse) GetSpenderAddressOk() (*string, bool)`

GetSpenderAddressOk returns a tuple with the SpenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpenderAddress

`func (o *TokenApprovalResponse) SetSpenderAddress(v string)`

SetSpenderAddress sets SpenderAddress field to given value.


### GetToken

`func (o *TokenApprovalResponse) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenApprovalResponse) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenApprovalResponse) SetToken(v Token)`

SetToken sets Token field to given value.


### GetFromApprovedAmount

`func (o *TokenApprovalResponse) GetFromApprovedAmount() string`

GetFromApprovedAmount returns the FromApprovedAmount field if non-nil, zero value otherwise.

### GetFromApprovedAmountOk

`func (o *TokenApprovalResponse) GetFromApprovedAmountOk() (*string, bool)`

GetFromApprovedAmountOk returns a tuple with the FromApprovedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromApprovedAmount

`func (o *TokenApprovalResponse) SetFromApprovedAmount(v string)`

SetFromApprovedAmount sets FromApprovedAmount field to given value.


### GetToApprovedAmount

`func (o *TokenApprovalResponse) GetToApprovedAmount() string`

GetToApprovedAmount returns the ToApprovedAmount field if non-nil, zero value otherwise.

### GetToApprovedAmountOk

`func (o *TokenApprovalResponse) GetToApprovedAmountOk() (*string, bool)`

GetToApprovedAmountOk returns a tuple with the ToApprovedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToApprovedAmount

`func (o *TokenApprovalResponse) SetToApprovedAmount(v string)`

SetToApprovedAmount sets ToApprovedAmount field to given value.


### GetTransactionRequest

`func (o *TokenApprovalResponse) GetTransactionRequest() TransactionRequest`

GetTransactionRequest returns the TransactionRequest field if non-nil, zero value otherwise.

### GetTransactionRequestOk

`func (o *TokenApprovalResponse) GetTransactionRequestOk() (*TransactionRequest, bool)`

GetTransactionRequestOk returns a tuple with the TransactionRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRequest

`func (o *TokenApprovalResponse) SetTransactionRequest(v TransactionRequest)`

SetTransactionRequest sets TransactionRequest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


