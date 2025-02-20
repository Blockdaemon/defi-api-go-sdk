# IncludedStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the step, such as &#39;swap&#39; or &#39;cross&#39;. | 
**Action** | [**Action**](Action.md) |  | 
**Estimate** | [**Estimate**](Estimate.md) |  | 
**IntegrationDetails** | [**IntegrationDetails**](IntegrationDetails.md) |  | 

## Methods

### NewIncludedStep

`func NewIncludedStep(type_ string, action Action, estimate Estimate, integrationDetails IntegrationDetails, ) *IncludedStep`

NewIncludedStep instantiates a new IncludedStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncludedStepWithDefaults

`func NewIncludedStepWithDefaults() *IncludedStep`

NewIncludedStepWithDefaults instantiates a new IncludedStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IncludedStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncludedStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncludedStep) SetType(v string)`

SetType sets Type field to given value.


### GetAction

`func (o *IncludedStep) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IncludedStep) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IncludedStep) SetAction(v Action)`

SetAction sets Action field to given value.


### GetEstimate

`func (o *IncludedStep) GetEstimate() Estimate`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *IncludedStep) GetEstimateOk() (*Estimate, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *IncludedStep) SetEstimate(v Estimate)`

SetEstimate sets Estimate field to given value.


### GetIntegrationDetails

`func (o *IncludedStep) GetIntegrationDetails() IntegrationDetails`

GetIntegrationDetails returns the IntegrationDetails field if non-nil, zero value otherwise.

### GetIntegrationDetailsOk

`func (o *IncludedStep) GetIntegrationDetailsOk() (*IntegrationDetails, bool)`

GetIntegrationDetailsOk returns a tuple with the IntegrationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationDetails

`func (o *IncludedStep) SetIntegrationDetails(v IntegrationDetails)`

SetIntegrationDetails sets IntegrationDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


