# NativeCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the native token. | 
**Symbol** | **string** | Symbol of the native token. | 
**Decimals** | **int32** | Number of decimal places used by the native token. | 
**Icon** | Pointer to **string** | URL of the native token icon. | [optional] 

## Methods

### NewNativeCurrency

`func NewNativeCurrency(name string, symbol string, decimals int32, ) *NativeCurrency`

NewNativeCurrency instantiates a new NativeCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNativeCurrencyWithDefaults

`func NewNativeCurrencyWithDefaults() *NativeCurrency`

NewNativeCurrencyWithDefaults instantiates a new NativeCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NativeCurrency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NativeCurrency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NativeCurrency) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *NativeCurrency) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NativeCurrency) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NativeCurrency) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *NativeCurrency) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *NativeCurrency) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *NativeCurrency) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetIcon

`func (o *NativeCurrency) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *NativeCurrency) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *NativeCurrency) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *NativeCurrency) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


