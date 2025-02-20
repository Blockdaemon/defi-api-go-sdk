# PortfolioViewSumSums

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]AccountSum**](AccountSum.md) | Combined approval dollar amounts by account Address. | [optional] 
**Chains** | Pointer to [**[]ChainSum**](ChainSum.md) | Combined approval dollar amounts by chain ID. | [optional] 

## Methods

### NewPortfolioViewSumSums

`func NewPortfolioViewSumSums() *PortfolioViewSumSums`

NewPortfolioViewSumSums instantiates a new PortfolioViewSumSums object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortfolioViewSumSumsWithDefaults

`func NewPortfolioViewSumSumsWithDefaults() *PortfolioViewSumSums`

NewPortfolioViewSumSumsWithDefaults instantiates a new PortfolioViewSumSums object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *PortfolioViewSumSums) GetAccounts() []AccountSum`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *PortfolioViewSumSums) GetAccountsOk() (*[]AccountSum, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *PortfolioViewSumSums) SetAccounts(v []AccountSum)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *PortfolioViewSumSums) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetChains

`func (o *PortfolioViewSumSums) GetChains() []ChainSum`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *PortfolioViewSumSums) GetChainsOk() (*[]ChainSum, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *PortfolioViewSumSums) SetChains(v []ChainSum)`

SetChains sets Chains field to given value.

### HasChains

`func (o *PortfolioViewSumSums) HasChains() bool`

HasChains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


