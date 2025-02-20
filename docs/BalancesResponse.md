# BalancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | The address of an account. | [optional] 
**Balances** | [**[]ChainBalance**](ChainBalance.md) |  | 
**HasMissingBalances** | Pointer to **[]string** | Array of chain IDs for which balances are missing | [optional] 

## Methods

### NewBalancesResponse

`func NewBalancesResponse(balances []ChainBalance, ) *BalancesResponse`

NewBalancesResponse instantiates a new BalancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalancesResponseWithDefaults

`func NewBalancesResponseWithDefaults() *BalancesResponse`

NewBalancesResponseWithDefaults instantiates a new BalancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *BalancesResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BalancesResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BalancesResponse) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BalancesResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBalances

`func (o *BalancesResponse) GetBalances() []ChainBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *BalancesResponse) GetBalancesOk() (*[]ChainBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *BalancesResponse) SetBalances(v []ChainBalance)`

SetBalances sets Balances field to given value.


### GetHasMissingBalances

`func (o *BalancesResponse) GetHasMissingBalances() []string`

GetHasMissingBalances returns the HasMissingBalances field if non-nil, zero value otherwise.

### GetHasMissingBalancesOk

`func (o *BalancesResponse) GetHasMissingBalancesOk() (*[]string, bool)`

GetHasMissingBalancesOk returns a tuple with the HasMissingBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMissingBalances

`func (o *BalancesResponse) SetHasMissingBalances(v []string)`

SetHasMissingBalances sets HasMissingBalances field to given value.

### HasHasMissingBalances

`func (o *BalancesResponse) HasHasMissingBalances() bool`

HasHasMissingBalances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


