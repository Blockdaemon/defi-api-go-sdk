# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**TokenAmount**](TokenAmount.md) |  | 
**To** | [**TokenAmount**](TokenAmount.md) |  | 
**Slippage** | **float32** | The maximum allowed price slippage for a transaction, expressed as a fraction of 100. | [default to 0.01]
**FromAccount** | [**Account**](Account.md) |  | 
**ToAccount** | [**Account**](Account.md) |  | 

## Methods

### NewAction

`func NewAction(from TokenAmount, to TokenAmount, slippage float32, fromAccount Account, toAccount Account, ) *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *Action) GetFrom() TokenAmount`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Action) GetFromOk() (*TokenAmount, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Action) SetFrom(v TokenAmount)`

SetFrom sets From field to given value.


### GetTo

`func (o *Action) GetTo() TokenAmount`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Action) GetToOk() (*TokenAmount, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Action) SetTo(v TokenAmount)`

SetTo sets To field to given value.


### GetSlippage

`func (o *Action) GetSlippage() float32`

GetSlippage returns the Slippage field if non-nil, zero value otherwise.

### GetSlippageOk

`func (o *Action) GetSlippageOk() (*float32, bool)`

GetSlippageOk returns a tuple with the Slippage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippage

`func (o *Action) SetSlippage(v float32)`

SetSlippage sets Slippage field to given value.


### GetFromAccount

`func (o *Action) GetFromAccount() Account`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *Action) GetFromAccountOk() (*Account, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *Action) SetFromAccount(v Account)`

SetFromAccount sets FromAccount field to given value.


### GetToAccount

`func (o *Action) GetToAccount() Account`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *Action) GetToAccountOk() (*Account, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *Action) SetToAccount(v Account)`

SetToAccount sets ToAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


