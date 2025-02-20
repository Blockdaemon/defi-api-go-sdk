# GasDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | [**Token**](Token.md) |  | 
**Gas** | **int32** | The gas limit provided by the sender. | 
**GasUsed** | **int32** | The amount of gas used by the transaction. | 
**GasPrice** | **string** | The amount of a token represented as a string. | 
**GasUSD** | Pointer to **float64** | The price of an asset, coin, or token. | [optional] 
**PriorityFee** | **string** | The priority fee (tip) paid to miners. | 
**BaseFee** | **string** | The base fee per gas unit, applicable in EIP-1559 transactions. | 

## Methods

### NewGasDetails

`func NewGasDetails(token Token, gas int32, gasUsed int32, gasPrice string, priorityFee string, baseFee string, ) *GasDetails`

NewGasDetails instantiates a new GasDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGasDetailsWithDefaults

`func NewGasDetailsWithDefaults() *GasDetails`

NewGasDetailsWithDefaults instantiates a new GasDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *GasDetails) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GasDetails) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GasDetails) SetToken(v Token)`

SetToken sets Token field to given value.


### GetGas

`func (o *GasDetails) GetGas() int32`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *GasDetails) GetGasOk() (*int32, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *GasDetails) SetGas(v int32)`

SetGas sets Gas field to given value.


### GetGasUsed

`func (o *GasDetails) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GasDetails) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GasDetails) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetGasPrice

`func (o *GasDetails) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GasDetails) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GasDetails) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetGasUSD

`func (o *GasDetails) GetGasUSD() float64`

GetGasUSD returns the GasUSD field if non-nil, zero value otherwise.

### GetGasUSDOk

`func (o *GasDetails) GetGasUSDOk() (*float64, bool)`

GetGasUSDOk returns a tuple with the GasUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUSD

`func (o *GasDetails) SetGasUSD(v float64)`

SetGasUSD sets GasUSD field to given value.

### HasGasUSD

`func (o *GasDetails) HasGasUSD() bool`

HasGasUSD returns a boolean if a field has been set.

### GetPriorityFee

`func (o *GasDetails) GetPriorityFee() string`

GetPriorityFee returns the PriorityFee field if non-nil, zero value otherwise.

### GetPriorityFeeOk

`func (o *GasDetails) GetPriorityFeeOk() (*string, bool)`

GetPriorityFeeOk returns a tuple with the PriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityFee

`func (o *GasDetails) SetPriorityFee(v string)`

SetPriorityFee sets PriorityFee field to given value.


### GetBaseFee

`func (o *GasDetails) GetBaseFee() string`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *GasDetails) GetBaseFeeOk() (*string, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *GasDetails) SetBaseFee(v string)`

SetBaseFee sets BaseFee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


