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

// checks if the Tags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tags{}

// Tags A list of tags which can be associated with a token.
type Tags struct {
}

// NewTags instantiates a new Tags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTags() *Tags {
	this := Tags{}
	return &this
}

// NewTagsWithDefaults instantiates a new Tags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagsWithDefaults() *Tags {
	this := Tags{}
	return &this
}

func (o Tags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tags) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return toSerialize, nil
}

func (o *Tags) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTags struct {
	value Tags
	isSet bool
}

func (v NullableTags) Get() Tags {
	return v.value
}

func (v *NullableTags) Set(val Tags) {
	v.value = val
	v.isSet = true
}

func (v NullableTags) IsSet() bool {
	return v.isSet
}

func (v *NullableTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTags(val Tags) *NullableTags {
	return &NullableTags{value: val, isSet: true}
}

func (v NullableTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


