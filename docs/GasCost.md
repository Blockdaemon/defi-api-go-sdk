# GasCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **string** | The gas limit for the (internal-)transaction, specified as a string to maintain precision. | 
**Token** | [**TokenAmount**](TokenAmount.md) |  | 

## Methods

### NewGasCost

`func NewGasCost(limit string, token TokenAmount, ) *GasCost`

NewGasCost instantiates a new GasCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGasCostWithDefaults

`func NewGasCostWithDefaults() *GasCost`

NewGasCostWithDefaults instantiates a new GasCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GasCost) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GasCost) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GasCost) SetLimit(v string)`

SetLimit sets Limit field to given value.


### GetToken

`func (o *GasCost) GetToken() TokenAmount`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GasCost) GetTokenOk() (*TokenAmount, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GasCost) SetToken(v TokenAmount)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


