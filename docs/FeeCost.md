# FeeCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the fee cost. | 
**Description** | **string** | Description of the fee cost. | 
**Token** | [**TokenAmount**](TokenAmount.md) |  | 
**Included** | **bool** | Flag indicating if the fee cost is included in the fromAmount. If true, the fee will be deducted from the fromAmount to calculate the final toAmount. If false, the fee needs to be payed on top of the fromAmount. | 

## Methods

### NewFeeCost

`func NewFeeCost(name string, description string, token TokenAmount, included bool, ) *FeeCost`

NewFeeCost instantiates a new FeeCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeCostWithDefaults

`func NewFeeCostWithDefaults() *FeeCost`

NewFeeCostWithDefaults instantiates a new FeeCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeeCost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeeCost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeeCost) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FeeCost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeeCost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeeCost) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetToken

`func (o *FeeCost) GetToken() TokenAmount`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FeeCost) GetTokenOk() (*TokenAmount, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FeeCost) SetToken(v TokenAmount)`

SetToken sets Token field to given value.


### GetIncluded

`func (o *FeeCost) GetIncluded() bool`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *FeeCost) GetIncludedOk() (*bool, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *FeeCost) SetIncluded(v bool)`

SetIncluded sets Included field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


