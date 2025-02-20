# AccountSum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address of an account. | 
**ApprovedSum** | **string** | The amount of a token represented as a string. | 
**ValueAtRiskSum** | **string** | The amount of a token represented as a string. | 

## Methods

### NewAccountSum

`func NewAccountSum(address string, approvedSum string, valueAtRiskSum string, ) *AccountSum`

NewAccountSum instantiates a new AccountSum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSumWithDefaults

`func NewAccountSumWithDefaults() *AccountSum`

NewAccountSumWithDefaults instantiates a new AccountSum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AccountSum) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccountSum) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccountSum) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetApprovedSum

`func (o *AccountSum) GetApprovedSum() string`

GetApprovedSum returns the ApprovedSum field if non-nil, zero value otherwise.

### GetApprovedSumOk

`func (o *AccountSum) GetApprovedSumOk() (*string, bool)`

GetApprovedSumOk returns a tuple with the ApprovedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedSum

`func (o *AccountSum) SetApprovedSum(v string)`

SetApprovedSum sets ApprovedSum field to given value.


### GetValueAtRiskSum

`func (o *AccountSum) GetValueAtRiskSum() string`

GetValueAtRiskSum returns the ValueAtRiskSum field if non-nil, zero value otherwise.

### GetValueAtRiskSumOk

`func (o *AccountSum) GetValueAtRiskSumOk() (*string, bool)`

GetValueAtRiskSumOk returns a tuple with the ValueAtRiskSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAtRiskSum

`func (o *AccountSum) SetValueAtRiskSum(v string)`

SetValueAtRiskSum sets ValueAtRiskSum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


