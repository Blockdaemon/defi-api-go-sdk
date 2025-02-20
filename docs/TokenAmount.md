# TokenAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | The amount of a token represented as a string. | 
**AmountUSD** | Pointer to **float32** | The amount of a token represented as a decimal number. | [optional] 
**Token** | [**Token**](Token.md) |  | 

## Methods

### NewTokenAmount

`func NewTokenAmount(amount string, token Token, ) *TokenAmount`

NewTokenAmount instantiates a new TokenAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenAmountWithDefaults

`func NewTokenAmountWithDefaults() *TokenAmount`

NewTokenAmountWithDefaults instantiates a new TokenAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TokenAmount) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenAmount) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenAmount) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAmountUSD

`func (o *TokenAmount) GetAmountUSD() float32`

GetAmountUSD returns the AmountUSD field if non-nil, zero value otherwise.

### GetAmountUSDOk

`func (o *TokenAmount) GetAmountUSDOk() (*float32, bool)`

GetAmountUSDOk returns a tuple with the AmountUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountUSD

`func (o *TokenAmount) SetAmountUSD(v float32)`

SetAmountUSD sets AmountUSD field to given value.

### HasAmountUSD

`func (o *TokenAmount) HasAmountUSD() bool`

HasAmountUSD returns a boolean if a field has been set.

### GetToken

`func (o *TokenAmount) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenAmount) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenAmount) SetToken(v Token)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


