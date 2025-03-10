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
)

// checks if the Tokens type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tokens{}

// Tokens A list of supported tokens and their metadata.
type Tokens struct {
}

// NewTokens instantiates a new Tokens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokens() *Tokens {
	this := Tokens{}
	return &this
}

// NewTokensWithDefaults instantiates a new Tokens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokensWithDefaults() *Tokens {
	this := Tokens{}
	return &this
}

func (o Tokens) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tokens) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return toSerialize, nil
}

func (o *Tokens) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTokens struct {
	value Tokens
	isSet bool
}

func (v NullableTokens) Get() Tokens {
	return v.value
}

func (v *NullableTokens) Set(val Tokens) {
	v.value = val
	v.isSet = true
}

func (v NullableTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokens(val Tokens) *NullableTokens {
	return &NullableTokens{value: val, isSet: true}
}

func (v NullableTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


