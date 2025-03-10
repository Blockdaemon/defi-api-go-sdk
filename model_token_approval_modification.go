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

// checks if the TokenApprovalModification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenApprovalModification{}

// TokenApprovalModification The address of an account.
type TokenApprovalModification struct {
	// The unique identifier of a blockchain in CAIP-2 notation.
	ChainID string `json:"chainID" validate:"regexp=^[-a-z0-9]{3,8}:[-_a-zA-Z0-9]{1,32}$"`
	// The address of an account.
	AccountAddress string `json:"accountAddress"`
	// The address of an account.
	SpenderAddress string `json:"spenderAddress"`
	// The address of an account.
	TokenAddress string `json:"tokenAddress"`
	// The amount of a token represented as a string.
	ToApprovedAmount string `json:"toApprovedAmount" validate:"regexp=^([1-9][0-9]*|0)(\\\\.[0-9]+)?$"`
}

type _TokenApprovalModification TokenApprovalModification

// NewTokenApprovalModification instantiates a new TokenApprovalModification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenApprovalModification(chainID string, accountAddress string, spenderAddress string, tokenAddress string, toApprovedAmount string) *TokenApprovalModification {
	this := TokenApprovalModification{}
	this.ChainID = chainID
	this.AccountAddress = accountAddress
	this.SpenderAddress = spenderAddress
	this.TokenAddress = tokenAddress
	this.ToApprovedAmount = toApprovedAmount
	return &this
}

// NewTokenApprovalModificationWithDefaults instantiates a new TokenApprovalModification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenApprovalModificationWithDefaults() *TokenApprovalModification {
	this := TokenApprovalModification{}
	return &this
}

// GetChainID returns the ChainID field value
func (o *TokenApprovalModification) GetChainID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainID
}

// GetChainIDOk returns a tuple with the ChainID field value
// and a boolean to check if the value has been set.
func (o *TokenApprovalModification) GetChainIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainID, true
}

// SetChainID sets field value
func (o *TokenApprovalModification) SetChainID(v string) {
	o.ChainID = v
}

// GetAccountAddress returns the AccountAddress field value
func (o *TokenApprovalModification) GetAccountAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountAddress
}

// GetAccountAddressOk returns a tuple with the AccountAddress field value
// and a boolean to check if the value has been set.
func (o *TokenApprovalModification) GetAccountAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountAddress, true
}

// SetAccountAddress sets field value
func (o *TokenApprovalModification) SetAccountAddress(v string) {
	o.AccountAddress = v
}

// GetSpenderAddress returns the SpenderAddress field value
func (o *TokenApprovalModification) GetSpenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpenderAddress
}

// GetSpenderAddressOk returns a tuple with the SpenderAddress field value
// and a boolean to check if the value has been set.
func (o *TokenApprovalModification) GetSpenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpenderAddress, true
}

// SetSpenderAddress sets field value
func (o *TokenApprovalModification) SetSpenderAddress(v string) {
	o.SpenderAddress = v
}

// GetTokenAddress returns the TokenAddress field value
func (o *TokenApprovalModification) GetTokenAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value
// and a boolean to check if the value has been set.
func (o *TokenApprovalModification) GetTokenAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenAddress, true
}

// SetTokenAddress sets field value
func (o *TokenApprovalModification) SetTokenAddress(v string) {
	o.TokenAddress = v
}

// GetToApprovedAmount returns the ToApprovedAmount field value
func (o *TokenApprovalModification) GetToApprovedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToApprovedAmount
}

// GetToApprovedAmountOk returns a tuple with the ToApprovedAmount field value
// and a boolean to check if the value has been set.
func (o *TokenApprovalModification) GetToApprovedAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToApprovedAmount, true
}

// SetToApprovedAmount sets field value
func (o *TokenApprovalModification) SetToApprovedAmount(v string) {
	o.ToApprovedAmount = v
}

func (o TokenApprovalModification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenApprovalModification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chainID"] = o.ChainID
	toSerialize["accountAddress"] = o.AccountAddress
	toSerialize["spenderAddress"] = o.SpenderAddress
	toSerialize["tokenAddress"] = o.TokenAddress
	toSerialize["toApprovedAmount"] = o.ToApprovedAmount
	return toSerialize, nil
}

func (o *TokenApprovalModification) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chainID",
		"accountAddress",
		"spenderAddress",
		"tokenAddress",
		"toApprovedAmount",
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

	varTokenApprovalModification := _TokenApprovalModification{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenApprovalModification)

	if err != nil {
		return err
	}

	*o = TokenApprovalModification(varTokenApprovalModification)

	return err
}

type NullableTokenApprovalModification struct {
	value *TokenApprovalModification
	isSet bool
}

func (v NullableTokenApprovalModification) Get() *TokenApprovalModification {
	return v.value
}

func (v *NullableTokenApprovalModification) Set(val *TokenApprovalModification) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenApprovalModification) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenApprovalModification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenApprovalModification(val *TokenApprovalModification) *NullableTokenApprovalModification {
	return &NullableTokenApprovalModification{value: val, isSet: true}
}

func (v NullableTokenApprovalModification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenApprovalModification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


