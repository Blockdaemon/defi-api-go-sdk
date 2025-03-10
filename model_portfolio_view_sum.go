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

// checks if the PortfolioViewSum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioViewSum{}

// PortfolioViewSum A wrapper over account sum and chain sum objects.
type PortfolioViewSum struct {
	Sums *PortfolioViewSumSums `json:"sums,omitempty"`
}

// NewPortfolioViewSum instantiates a new PortfolioViewSum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioViewSum() *PortfolioViewSum {
	this := PortfolioViewSum{}
	return &this
}

// NewPortfolioViewSumWithDefaults instantiates a new PortfolioViewSum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioViewSumWithDefaults() *PortfolioViewSum {
	this := PortfolioViewSum{}
	return &this
}

// GetSums returns the Sums field value if set, zero value otherwise.
func (o *PortfolioViewSum) GetSums() PortfolioViewSumSums {
	if o == nil || IsNil(o.Sums) {
		var ret PortfolioViewSumSums
		return ret
	}
	return *o.Sums
}

// GetSumsOk returns a tuple with the Sums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioViewSum) GetSumsOk() (*PortfolioViewSumSums, bool) {
	if o == nil || IsNil(o.Sums) {
		return nil, false
	}
	return o.Sums, true
}

// HasSums returns a boolean if a field has been set.
func (o *PortfolioViewSum) HasSums() bool {
	if o != nil && !IsNil(o.Sums) {
		return true
	}

	return false
}

// SetSums gets a reference to the given PortfolioViewSumSums and assigns it to the Sums field.
func (o *PortfolioViewSum) SetSums(v PortfolioViewSumSums) {
	o.Sums = &v
}

func (o PortfolioViewSum) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioViewSum) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sums) {
		toSerialize["sums"] = o.Sums
	}
	return toSerialize, nil
}

type NullablePortfolioViewSum struct {
	value *PortfolioViewSum
	isSet bool
}

func (v NullablePortfolioViewSum) Get() *PortfolioViewSum {
	return v.value
}

func (v *NullablePortfolioViewSum) Set(val *PortfolioViewSum) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioViewSum) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioViewSum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioViewSum(val *PortfolioViewSum) *NullablePortfolioViewSum {
	return &NullablePortfolioViewSum{value: val, isSet: true}
}

func (v NullablePortfolioViewSum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioViewSum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


