# StatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**StatusEnum**](StatusEnum.md) |  | 
**Source** | [**StatusChainDetail**](StatusChainDetail.md) |  | 
**Target** | Pointer to [**StatusChainDetail**](StatusChainDetail.md) |  | [optional] 

## Methods

### NewStatusResponse

`func NewStatusResponse(status StatusEnum, source StatusChainDetail, ) *StatusResponse`

NewStatusResponse instantiates a new StatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseWithDefaults

`func NewStatusResponseWithDefaults() *StatusResponse`

NewStatusResponseWithDefaults instantiates a new StatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StatusResponse) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatusResponse) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatusResponse) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetSource

`func (o *StatusResponse) GetSource() StatusChainDetail`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StatusResponse) GetSourceOk() (*StatusChainDetail, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StatusResponse) SetSource(v StatusChainDetail)`

SetSource sets Source field to given value.


### GetTarget

`func (o *StatusResponse) GetTarget() StatusChainDetail`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *StatusResponse) GetTargetOk() (*StatusChainDetail, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *StatusResponse) SetTarget(v StatusChainDetail)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *StatusResponse) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


