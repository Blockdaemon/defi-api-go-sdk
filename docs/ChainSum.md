# ChainSum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** | The unique identifier of a blockchain in CAIP-2 notation. | 
**ApprovedSum** | **string** | The amount of a token represented as a string. | 
**ValueAtRiskSum** | **string** | The amount of a token represented as a string. | 

## Methods

### NewChainSum

`func NewChainSum(chainId string, approvedSum string, valueAtRiskSum string, ) *ChainSum`

NewChainSum instantiates a new ChainSum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainSumWithDefaults

`func NewChainSumWithDefaults() *ChainSum`

NewChainSumWithDefaults instantiates a new ChainSum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *ChainSum) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ChainSum) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ChainSum) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetApprovedSum

`func (o *ChainSum) GetApprovedSum() string`

GetApprovedSum returns the ApprovedSum field if non-nil, zero value otherwise.

### GetApprovedSumOk

`func (o *ChainSum) GetApprovedSumOk() (*string, bool)`

GetApprovedSumOk returns a tuple with the ApprovedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedSum

`func (o *ChainSum) SetApprovedSum(v string)`

SetApprovedSum sets ApprovedSum field to given value.


### GetValueAtRiskSum

`func (o *ChainSum) GetValueAtRiskSum() string`

GetValueAtRiskSum returns the ValueAtRiskSum field if non-nil, zero value otherwise.

### GetValueAtRiskSumOk

`func (o *ChainSum) GetValueAtRiskSumOk() (*string, bool)`

GetValueAtRiskSumOk returns a tuple with the ValueAtRiskSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAtRiskSum

`func (o *ChainSum) SetValueAtRiskSum(v string)`

SetValueAtRiskSum sets ValueAtRiskSum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


