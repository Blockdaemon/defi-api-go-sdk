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

// checks if the Chains type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Chains{}

// Chains A list of supported blockchains and their metadata.
type Chains struct {
}

// NewChains instantiates a new Chains object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChains() *Chains {
	this := Chains{}
	return &this
}

// NewChainsWithDefaults instantiates a new Chains object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChainsWithDefaults() *Chains {
	this := Chains{}
	return &this
}

func (o Chains) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Chains) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return toSerialize, nil
}

func (o *Chains) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableChains struct {
	value Chains
	isSet bool
}

func (v NullableChains) Get() Chains {
	return v.value
}

func (v *NullableChains) Set(val Chains) {
	v.value = val
	v.isSet = true
}

func (v NullableChains) IsSet() bool {
	return v.isSet
}

func (v *NullableChains) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChains(val Chains) *NullableChains {
	return &NullableChains{value: val, isSet: true}
}

func (v NullableChains) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChains) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


