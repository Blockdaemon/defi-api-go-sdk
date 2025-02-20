# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetID** | **string** | A unique identifier for a plugin or sub-component. | 
**From** | [**TokenAmount**](TokenAmount.md) |  | 
**To** | [**TokenAmount**](TokenAmount.md) |  | 
**Slippage** | Pointer to **float32** | The amount of a token represented as a decimal number. | [optional] 
**Steps** | [**IncludedSteps**](IncludedSteps.md) |  | 
**TransactionRequest** | [**TransactionRequest**](TransactionRequest.md) |  | 
**Extensions** | Pointer to **map[string]string** | Additional metadata about the route. | [optional] 

## Methods

### NewRoute

`func NewRoute(targetID string, from TokenAmount, to TokenAmount, steps IncludedSteps, transactionRequest TransactionRequest, ) *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetID

`func (o *Route) GetTargetID() string`

GetTargetID returns the TargetID field if non-nil, zero value otherwise.

### GetTargetIDOk

`func (o *Route) GetTargetIDOk() (*string, bool)`

GetTargetIDOk returns a tuple with the TargetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetID

`func (o *Route) SetTargetID(v string)`

SetTargetID sets TargetID field to given value.


### GetFrom

`func (o *Route) GetFrom() TokenAmount`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Route) GetFromOk() (*TokenAmount, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Route) SetFrom(v TokenAmount)`

SetFrom sets From field to given value.


### GetTo

`func (o *Route) GetTo() TokenAmount`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Route) GetToOk() (*TokenAmount, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Route) SetTo(v TokenAmount)`

SetTo sets To field to given value.


### GetSlippage

`func (o *Route) GetSlippage() float32`

GetSlippage returns the Slippage field if non-nil, zero value otherwise.

### GetSlippageOk

`func (o *Route) GetSlippageOk() (*float32, bool)`

GetSlippageOk returns a tuple with the Slippage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippage

`func (o *Route) SetSlippage(v float32)`

SetSlippage sets Slippage field to given value.

### HasSlippage

`func (o *Route) HasSlippage() bool`

HasSlippage returns a boolean if a field has been set.

### GetSteps

`func (o *Route) GetSteps() IncludedSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Route) GetStepsOk() (*IncludedSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Route) SetSteps(v IncludedSteps)`

SetSteps sets Steps field to given value.


### GetTransactionRequest

`func (o *Route) GetTransactionRequest() TransactionRequest`

GetTransactionRequest returns the TransactionRequest field if non-nil, zero value otherwise.

### GetTransactionRequestOk

`func (o *Route) GetTransactionRequestOk() (*TransactionRequest, bool)`

GetTransactionRequestOk returns a tuple with the TransactionRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRequest

`func (o *Route) SetTransactionRequest(v TransactionRequest)`

SetTransactionRequest sets TransactionRequest field to given value.


### GetExtensions

`func (o *Route) GetExtensions() map[string]string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Route) GetExtensionsOk() (*map[string]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Route) SetExtensions(v map[string]string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Route) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


