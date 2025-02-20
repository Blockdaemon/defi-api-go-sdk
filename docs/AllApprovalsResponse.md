# AllApprovalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | [**[]TokenApprovalData**](TokenApprovalData.md) | List of token approval data for the requested chains and accounts. | 
**PortfolioViewSum** | [**PortfolioViewSum**](PortfolioViewSum.md) |  | 
**HasMissingApprovals** | Pointer to **[]string** | Array of chain IDs for which balances are missing | [optional] 

## Methods

### NewAllApprovalsResponse

`func NewAllApprovalsResponse(approvals []TokenApprovalData, portfolioViewSum PortfolioViewSum, ) *AllApprovalsResponse`

NewAllApprovalsResponse instantiates a new AllApprovalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllApprovalsResponseWithDefaults

`func NewAllApprovalsResponseWithDefaults() *AllApprovalsResponse`

NewAllApprovalsResponseWithDefaults instantiates a new AllApprovalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *AllApprovalsResponse) GetApprovals() []TokenApprovalData`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *AllApprovalsResponse) GetApprovalsOk() (*[]TokenApprovalData, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *AllApprovalsResponse) SetApprovals(v []TokenApprovalData)`

SetApprovals sets Approvals field to given value.


### GetPortfolioViewSum

`func (o *AllApprovalsResponse) GetPortfolioViewSum() PortfolioViewSum`

GetPortfolioViewSum returns the PortfolioViewSum field if non-nil, zero value otherwise.

### GetPortfolioViewSumOk

`func (o *AllApprovalsResponse) GetPortfolioViewSumOk() (*PortfolioViewSum, bool)`

GetPortfolioViewSumOk returns a tuple with the PortfolioViewSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolioViewSum

`func (o *AllApprovalsResponse) SetPortfolioViewSum(v PortfolioViewSum)`

SetPortfolioViewSum sets PortfolioViewSum field to given value.


### GetHasMissingApprovals

`func (o *AllApprovalsResponse) GetHasMissingApprovals() []string`

GetHasMissingApprovals returns the HasMissingApprovals field if non-nil, zero value otherwise.

### GetHasMissingApprovalsOk

`func (o *AllApprovalsResponse) GetHasMissingApprovalsOk() (*[]string, bool)`

GetHasMissingApprovalsOk returns a tuple with the HasMissingApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMissingApprovals

`func (o *AllApprovalsResponse) SetHasMissingApprovals(v []string)`

SetHasMissingApprovals sets HasMissingApprovals field to given value.

### HasHasMissingApprovals

`func (o *AllApprovalsResponse) HasHasMissingApprovals() bool`

HasHasMissingApprovals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


