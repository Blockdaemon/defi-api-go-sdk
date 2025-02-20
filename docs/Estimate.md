# Estimate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationDetails** | [**IntegrationDetails**](IntegrationDetails.md) |  | 
**ApprovalAddress** | **string** | The address of a token, account, or smart contract. | 
**ToAmountMin** | Pointer to **string** | The minimum amount to be received after the transaction, specified as a string to maintain precision. | [optional] 
**ToAmount** | **string** | The estimated amount to be received after the transaction, specified as a string to maintain precision. | 
**FromAmount** | **string** | The amount of tokens to be sent, including decimals. | 
**FeeCosts** | [**[]FeeCost**](FeeCost.md) | Transaction fees, if any. This can be empty if there are no extra fees. | 
**GasCosts** | [**[]GasCost**](GasCost.md) | The CAIP-2 identifier for a blockchain. | 
**ExecutionDuration** | Pointer to **int32** | The estimated time for the transaction to be executed, in seconds. | [optional] 
**FromAmountUSD** | Pointer to **string** | The value of the amount to be sent in USD. | [optional] 
**ToAmountUSD** | Pointer to **string** | The estimated value of the amount to be received in USD. | [optional] 

## Methods

### NewEstimate

`func NewEstimate(integrationDetails IntegrationDetails, approvalAddress string, toAmount string, fromAmount string, feeCosts []FeeCost, gasCosts []GasCost, ) *Estimate`

NewEstimate instantiates a new Estimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateWithDefaults

`func NewEstimateWithDefaults() *Estimate`

NewEstimateWithDefaults instantiates a new Estimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationDetails

`func (o *Estimate) GetIntegrationDetails() IntegrationDetails`

GetIntegrationDetails returns the IntegrationDetails field if non-nil, zero value otherwise.

### GetIntegrationDetailsOk

`func (o *Estimate) GetIntegrationDetailsOk() (*IntegrationDetails, bool)`

GetIntegrationDetailsOk returns a tuple with the IntegrationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationDetails

`func (o *Estimate) SetIntegrationDetails(v IntegrationDetails)`

SetIntegrationDetails sets IntegrationDetails field to given value.


### GetApprovalAddress

`func (o *Estimate) GetApprovalAddress() string`

GetApprovalAddress returns the ApprovalAddress field if non-nil, zero value otherwise.

### GetApprovalAddressOk

`func (o *Estimate) GetApprovalAddressOk() (*string, bool)`

GetApprovalAddressOk returns a tuple with the ApprovalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalAddress

`func (o *Estimate) SetApprovalAddress(v string)`

SetApprovalAddress sets ApprovalAddress field to given value.


### GetToAmountMin

`func (o *Estimate) GetToAmountMin() string`

GetToAmountMin returns the ToAmountMin field if non-nil, zero value otherwise.

### GetToAmountMinOk

`func (o *Estimate) GetToAmountMinOk() (*string, bool)`

GetToAmountMinOk returns a tuple with the ToAmountMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAmountMin

`func (o *Estimate) SetToAmountMin(v string)`

SetToAmountMin sets ToAmountMin field to given value.

### HasToAmountMin

`func (o *Estimate) HasToAmountMin() bool`

HasToAmountMin returns a boolean if a field has been set.

### GetToAmount

`func (o *Estimate) GetToAmount() string`

GetToAmount returns the ToAmount field if non-nil, zero value otherwise.

### GetToAmountOk

`func (o *Estimate) GetToAmountOk() (*string, bool)`

GetToAmountOk returns a tuple with the ToAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAmount

`func (o *Estimate) SetToAmount(v string)`

SetToAmount sets ToAmount field to given value.


### GetFromAmount

`func (o *Estimate) GetFromAmount() string`

GetFromAmount returns the FromAmount field if non-nil, zero value otherwise.

### GetFromAmountOk

`func (o *Estimate) GetFromAmountOk() (*string, bool)`

GetFromAmountOk returns a tuple with the FromAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAmount

`func (o *Estimate) SetFromAmount(v string)`

SetFromAmount sets FromAmount field to given value.


### GetFeeCosts

`func (o *Estimate) GetFeeCosts() []FeeCost`

GetFeeCosts returns the FeeCosts field if non-nil, zero value otherwise.

### GetFeeCostsOk

`func (o *Estimate) GetFeeCostsOk() (*[]FeeCost, bool)`

GetFeeCostsOk returns a tuple with the FeeCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCosts

`func (o *Estimate) SetFeeCosts(v []FeeCost)`

SetFeeCosts sets FeeCosts field to given value.


### GetGasCosts

`func (o *Estimate) GetGasCosts() []GasCost`

GetGasCosts returns the GasCosts field if non-nil, zero value otherwise.

### GetGasCostsOk

`func (o *Estimate) GetGasCostsOk() (*[]GasCost, bool)`

GetGasCostsOk returns a tuple with the GasCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasCosts

`func (o *Estimate) SetGasCosts(v []GasCost)`

SetGasCosts sets GasCosts field to given value.


### GetExecutionDuration

`func (o *Estimate) GetExecutionDuration() int32`

GetExecutionDuration returns the ExecutionDuration field if non-nil, zero value otherwise.

### GetExecutionDurationOk

`func (o *Estimate) GetExecutionDurationOk() (*int32, bool)`

GetExecutionDurationOk returns a tuple with the ExecutionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDuration

`func (o *Estimate) SetExecutionDuration(v int32)`

SetExecutionDuration sets ExecutionDuration field to given value.

### HasExecutionDuration

`func (o *Estimate) HasExecutionDuration() bool`

HasExecutionDuration returns a boolean if a field has been set.

### GetFromAmountUSD

`func (o *Estimate) GetFromAmountUSD() string`

GetFromAmountUSD returns the FromAmountUSD field if non-nil, zero value otherwise.

### GetFromAmountUSDOk

`func (o *Estimate) GetFromAmountUSDOk() (*string, bool)`

GetFromAmountUSDOk returns a tuple with the FromAmountUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAmountUSD

`func (o *Estimate) SetFromAmountUSD(v string)`

SetFromAmountUSD sets FromAmountUSD field to given value.

### HasFromAmountUSD

`func (o *Estimate) HasFromAmountUSD() bool`

HasFromAmountUSD returns a boolean if a field has been set.

### GetToAmountUSD

`func (o *Estimate) GetToAmountUSD() string`

GetToAmountUSD returns the ToAmountUSD field if non-nil, zero value otherwise.

### GetToAmountUSDOk

`func (o *Estimate) GetToAmountUSDOk() (*string, bool)`

GetToAmountUSDOk returns a tuple with the ToAmountUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAmountUSD

`func (o *Estimate) SetToAmountUSD(v string)`

SetToAmountUSD sets ToAmountUSD field to given value.

### HasToAmountUSD

`func (o *Estimate) HasToAmountUSD() bool`

HasToAmountUSD returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


