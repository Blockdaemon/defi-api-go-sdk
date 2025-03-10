/*
Blockdaemon DeFi API

The Blockdaemon DeFi API provides a single interface to a multitude of DeFi projects and blockchains

API version: 1.0.0
Contact: info@blockdaemon.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Token type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Token{}

// Token The metadata for a supported token.
type Token struct {
	// The unique identifier of a blockchain in CAIP-2 notation.
	ChainID string `json:"chainID" validate:"regexp=^[-a-z0-9]{3,8}:[-_a-zA-Z0-9]{1,32}$"`
	// The address of an account.
	Address string `json:"address"`
	// The CAIP-2 identifier for a blockchain.
	Name string `json:"name"`
	// The symbol of the token
	Symbol string `json:"symbol"`
	// How many decimals the token supports
	Decimals int32 `json:"decimals"`
	// The URI for the logo or icon of a token, chain, or DeFi project.
	LogoURI string `json:"logoURI"`
	Tags Tags `json:"tags"`
	// The price of the asset or token in US Dollars.
	PriceUSD *float64 `json:"priceUSD,omitempty"`
	Extensions Extensions `json:"extensions"`
}

type _Token Token

// NewToken instantiates a new Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToken(chainID string, address string, name string, symbol string, decimals int32, logoURI string, tags Tags, extensions Extensions) *Token {
	this := Token{}
	this.ChainID = chainID
	this.Address = address
	this.Name = name
	this.Symbol = symbol
	this.Decimals = decimals
	this.LogoURI = logoURI
	this.Tags = tags
	this.Extensions = extensions
	return &this
}

// NewTokenWithDefaults instantiates a new Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWithDefaults() *Token {
	this := Token{}
	return &this
}

// GetChainID returns the ChainID field value
func (o *Token) GetChainID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainID
}

// GetChainIDOk returns a tuple with the ChainID field value
// and a boolean to check if the value has been set.
func (o *Token) GetChainIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainID, true
}

// SetChainID sets field value
func (o *Token) SetChainID(v string) {
	o.ChainID = v
}

// GetAddress returns the Address field value
func (o *Token) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Token) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *Token) SetAddress(v string) {
	o.Address = v
}

// GetName returns the Name field value
func (o *Token) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Token) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Token) SetName(v string) {
	o.Name = v
}

// GetSymbol returns the Symbol field value
func (o *Token) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *Token) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *Token) SetSymbol(v string) {
	o.Symbol = v
}

// GetDecimals returns the Decimals field value
func (o *Token) GetDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value
// and a boolean to check if the value has been set.
func (o *Token) GetDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// SetDecimals sets field value
func (o *Token) SetDecimals(v int32) {
	o.Decimals = v
}

// GetLogoURI returns the LogoURI field value
func (o *Token) GetLogoURI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoURI
}

// GetLogoURIOk returns a tuple with the LogoURI field value
// and a boolean to check if the value has been set.
func (o *Token) GetLogoURIOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogoURI, true
}

// SetLogoURI sets field value
func (o *Token) SetLogoURI(v string) {
	o.LogoURI = v
}

// GetTags returns the Tags field value
func (o *Token) GetTags() Tags {
	if o == nil {
		var ret Tags
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Token) GetTagsOk() (*Tags, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *Token) SetTags(v Tags) {
	o.Tags = v
}

// GetPriceUSD returns the PriceUSD field value if set, zero value otherwise.
func (o *Token) GetPriceUSD() float64 {
	if o == nil || IsNil(o.PriceUSD) {
		var ret float64
		return ret
	}
	return *o.PriceUSD
}

// GetPriceUSDOk returns a tuple with the PriceUSD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetPriceUSDOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceUSD) {
		return nil, false
	}
	return o.PriceUSD, true
}

// HasPriceUSD returns a boolean if a field has been set.
func (o *Token) HasPriceUSD() bool {
	if o != nil && !IsNil(o.PriceUSD) {
		return true
	}

	return false
}

// SetPriceUSD gets a reference to the given float64 and assigns it to the PriceUSD field.
func (o *Token) SetPriceUSD(v float64) {
	o.PriceUSD = &v
}

// GetExtensions returns the Extensions field value
func (o *Token) GetExtensions() Extensions {
	if o == nil {
		var ret Extensions
		return ret
	}

	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value
// and a boolean to check if the value has been set.
func (o *Token) GetExtensionsOk() (*Extensions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Extensions, true
}

// SetExtensions sets field value
func (o *Token) SetExtensions(v Extensions) {
	o.Extensions = v
}

func (o Token) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Token) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chainID"] = o.ChainID
	toSerialize["address"] = o.Address
	toSerialize["name"] = o.Name
	toSerialize["symbol"] = o.Symbol
	toSerialize["decimals"] = o.Decimals
	toSerialize["logoURI"] = o.LogoURI
	toSerialize["tags"] = o.Tags
	if !IsNil(o.PriceUSD) {
		toSerialize["priceUSD"] = o.PriceUSD
	}
	toSerialize["extensions"] = o.Extensions
	return toSerialize, nil
}

func (o *Token) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chainID",
		"address",
		"name",
		"symbol",
		"decimals",
		"logoURI",
		"tags",
		"extensions",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varToken := _Token{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToken)

	if err != nil {
		return err
	}

	*o = Token(varToken)

	return err
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


