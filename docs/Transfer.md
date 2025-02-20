# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenAmount** | [**TokenAmount**](TokenAmount.md) |  | 
**FromAccount** | [**Account**](Account.md) |  | 
**ToAccount** | [**Account**](Account.md) |  | 
**Type** | [**TransferType**](TransferType.md) |  | 

## Methods

### NewTransfer

`func NewTransfer(tokenAmount TokenAmount, fromAccount Account, toAccount Account, type_ TransferType, ) *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenAmount

`func (o *Transfer) GetTokenAmount() TokenAmount`

GetTokenAmount returns the TokenAmount field if non-nil, zero value otherwise.

### GetTokenAmountOk

`func (o *Transfer) GetTokenAmountOk() (*TokenAmount, bool)`

GetTokenAmountOk returns a tuple with the TokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAmount

`func (o *Transfer) SetTokenAmount(v TokenAmount)`

SetTokenAmount sets TokenAmount field to given value.


### GetFromAccount

`func (o *Transfer) GetFromAccount() Account`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *Transfer) GetFromAccountOk() (*Account, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *Transfer) SetFromAccount(v Account)`

SetFromAccount sets FromAccount field to given value.


### GetToAccount

`func (o *Transfer) GetToAccount() Account`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *Transfer) GetToAccountOk() (*Account, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *Transfer) SetToAccount(v Account)`

SetToAccount sets ToAccount field to given value.


### GetType

`func (o *Transfer) GetType() TransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transfer) GetTypeOk() (*TransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transfer) SetType(v TransferType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


