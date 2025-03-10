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

// checks if the PortfolioViewSumSums type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioViewSumSums{}

// PortfolioViewSumSums struct for PortfolioViewSumSums
type PortfolioViewSumSums struct {
	// Combined approval dollar amounts by account Address.
	Accounts []AccountSum `json:"accounts,omitempty"`
	// Combined approval dollar amounts by chain ID.
	Chains []ChainSum `json:"chains,omitempty"`
}

// NewPortfolioViewSumSums instantiates a new PortfolioViewSumSums object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioViewSumSums() *PortfolioViewSumSums {
	this := PortfolioViewSumSums{}
	return &this
}

// NewPortfolioViewSumSumsWithDefaults instantiates a new PortfolioViewSumSums object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioViewSumSumsWithDefaults() *PortfolioViewSumSums {
	this := PortfolioViewSumSums{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *PortfolioViewSumSums) GetAccounts() []AccountSum {
	if o == nil || IsNil(o.Accounts) {
		var ret []AccountSum
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioViewSumSums) GetAccountsOk() ([]AccountSum, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *PortfolioViewSumSums) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []AccountSum and assigns it to the Accounts field.
func (o *PortfolioViewSumSums) SetAccounts(v []AccountSum) {
	o.Accounts = v
}

// GetChains returns the Chains field value if set, zero value otherwise.
func (o *PortfolioViewSumSums) GetChains() []ChainSum {
	if o == nil || IsNil(o.Chains) {
		var ret []ChainSum
		return ret
	}
	return o.Chains
}

// GetChainsOk returns a tuple with the Chains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioViewSumSums) GetChainsOk() ([]ChainSum, bool) {
	if o == nil || IsNil(o.Chains) {
		return nil, false
	}
	return o.Chains, true
}

// HasChains returns a boolean if a field has been set.
func (o *PortfolioViewSumSums) HasChains() bool {
	if o != nil && !IsNil(o.Chains) {
		return true
	}

	return false
}

// SetChains gets a reference to the given []ChainSum and assigns it to the Chains field.
func (o *PortfolioViewSumSums) SetChains(v []ChainSum) {
	o.Chains = v
}

func (o PortfolioViewSumSums) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioViewSumSums) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Chains) {
		toSerialize["chains"] = o.Chains
	}
	return toSerialize, nil
}

type NullablePortfolioViewSumSums struct {
	value *PortfolioViewSumSums
	isSet bool
}

func (v NullablePortfolioViewSumSums) Get() *PortfolioViewSumSums {
	return v.value
}

func (v *NullablePortfolioViewSumSums) Set(val *PortfolioViewSumSums) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioViewSumSums) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioViewSumSums) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioViewSumSums(val *PortfolioViewSumSums) *NullablePortfolioViewSumSums {
	return &NullablePortfolioViewSumSums{value: val, isSet: true}
}

func (v NullablePortfolioViewSumSums) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioViewSumSums) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


