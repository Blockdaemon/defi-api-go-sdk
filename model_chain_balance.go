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

// checks if the ChainBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChainBalance{}

// ChainBalance The token balances for an account across multiple chains.
type ChainBalance struct {
	// The unique identifier of a blockchain in CAIP-2 notation.
	ChainID string `json:"chainID" validate:"regexp=^[-a-z0-9]{3,8}:[-_a-zA-Z0-9]{1,32}$"`
	TokenBalances []TokenAmount `json:"tokenBalances"`
}

type _ChainBalance ChainBalance

// NewChainBalance instantiates a new ChainBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChainBalance(chainID string, tokenBalances []TokenAmount) *ChainBalance {
	this := ChainBalance{}
	this.ChainID = chainID
	this.TokenBalances = tokenBalances
	return &this
}

// NewChainBalanceWithDefaults instantiates a new ChainBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChainBalanceWithDefaults() *ChainBalance {
	this := ChainBalance{}
	return &this
}

// GetChainID returns the ChainID field value
func (o *ChainBalance) GetChainID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainID
}

// GetChainIDOk returns a tuple with the ChainID field value
// and a boolean to check if the value has been set.
func (o *ChainBalance) GetChainIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainID, true
}

// SetChainID sets field value
func (o *ChainBalance) SetChainID(v string) {
	o.ChainID = v
}

// GetTokenBalances returns the TokenBalances field value
func (o *ChainBalance) GetTokenBalances() []TokenAmount {
	if o == nil {
		var ret []TokenAmount
		return ret
	}

	return o.TokenBalances
}

// GetTokenBalancesOk returns a tuple with the TokenBalances field value
// and a boolean to check if the value has been set.
func (o *ChainBalance) GetTokenBalancesOk() ([]TokenAmount, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenBalances, true
}

// SetTokenBalances sets field value
func (o *ChainBalance) SetTokenBalances(v []TokenAmount) {
	o.TokenBalances = v
}

func (o ChainBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChainBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chainID"] = o.ChainID
	toSerialize["tokenBalances"] = o.TokenBalances
	return toSerialize, nil
}

func (o *ChainBalance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chainID",
		"tokenBalances",
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

	varChainBalance := _ChainBalance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChainBalance)

	if err != nil {
		return err
	}

	*o = ChainBalance(varChainBalance)

	return err
}

type NullableChainBalance struct {
	value *ChainBalance
	isSet bool
}

func (v NullableChainBalance) Get() *ChainBalance {
	return v.value
}

func (v *NullableChainBalance) Set(val *ChainBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableChainBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableChainBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChainBalance(val *ChainBalance) *NullableChainBalance {
	return &NullableChainBalance{value: val, isSet: true}
}

func (v NullableChainBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChainBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


