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

// checks if the TokenApprovalData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenApprovalData{}

// TokenApprovalData Details about a token approval for an account and spender.
type TokenApprovalData struct {
	// The unique identifier of a blockchain in CAIP-2 notation.
	ChainId string `json:"chainId" validate:"regexp=^[-a-z0-9]{3,8}:[-_a-zA-Z0-9]{1,32}$"`
	// The address of an account.
	AccountAddress string `json:"accountAddress"`
	// The address of an account.
	SpenderAddress string `json:"spenderAddress"`
	Token Token `json:"token"`
	// The amount of a token represented as a string.
	Amount string `json:"amount" validate:"regexp=^([1-9][0-9]*|0)(\\\\.[0-9]+)?$"`
	// The amount of a token represented as a string.
	AmountUSD *string `json:"amountUSD,omitempty" validate:"regexp=^([1-9][0-9]*|0)(\\\\.[0-9]+)?$"`
	// The amount of a token represented as a string.
	ValueAtRisk string `json:"valueAtRisk" validate:"regexp=^([1-9][0-9]*|0)(\\\\.[0-9]+)?$"`
	// The amount of a token represented as a string.
	ValueAtRiskUSD *string `json:"valueAtRiskUSD,omitempty" validate:"regexp=^([1-9][0-9]*|0)(\\\\.[0-9]+)?$"`
	// The address of an account.
	SpenderName *string `json:"spenderName,omitempty"`
}

type _TokenApprovalData TokenApprovalData

// NewTokenApprovalData instantiates a new TokenApprovalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenApprovalData(chainId string, accountAddress string, spenderAddress string, token Token, amount string, valueAtRisk string) *TokenApprovalData {
	this := TokenApprovalData{}
	this.ChainId = chainId
	this.AccountAddress = accountAddress
	this.SpenderAddress = spenderAddress
	this.Token = token
	this.Amount = amount
	this.ValueAtRisk = valueAtRisk
	return &this
}

// NewTokenApprovalDataWithDefaults instantiates a new TokenApprovalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenApprovalDataWithDefaults() *TokenApprovalData {
	this := TokenApprovalData{}
	return &this
}

// GetChainId returns the ChainId field value
func (o *TokenApprovalData) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *TokenApprovalData) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *TokenApprovalData) SetChainId(v string) {
	o.ChainId = v
}

// GetAccountAddress returns the AccountAddress field value
func (o *TokenApprovalData) GetAccountAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountAddress
}

// GetAccountAddressOk returns a tuple with the AccountAddress field value
// and a boolean to check if the value has been set.
func (o *TokenApprovalData) GetAccountAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountAddress, true
}

// SetAccountAddress sets field value
func (o *TokenApprovalData) SetAccountAddress(v string) {
	o.AccountAddress = v
}

// GetSpenderAddress returns the SpenderAddress field value
func (o *TokenApprovalData) GetSpenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpenderAddress
}

// GetSpenderAddressOk returns a tuple with the SpenderAddress field value
// and a boolean to check if the value has been set.
func (o *TokenApprovalData) GetSpenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpenderAddress, true
}

// SetSpenderAddress sets field value
func (o *TokenApprovalData) SetSpenderAddress(v string) {
	o.SpenderAddress = v
}

// GetToken returns the Token field value
func (o *TokenApprovalData) GetToken() Token {
	if o == nil {
		var ret Token
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *TokenApprovalData) GetTokenOk() (*Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *TokenApprovalData) SetToken(v Token) {
	o.Token = v
}

// GetAmount returns the Amount field value
func (o *TokenApprovalData) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TokenApprovalData) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TokenApprovalData) SetAmount(v string) {
	o.Amount = v
}

// GetAmountUSD returns the AmountUSD field value if set, zero value otherwise.
func (o *TokenApprovalData) GetAmountUSD() string {
	if o == nil || IsNil(o.AmountUSD) {
		var ret string
		return ret
	}
	return *o.AmountUSD
}

// GetAmountUSDOk returns a tuple with the AmountUSD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenApprovalData) GetAmountUSDOk() (*string, bool) {
	if o == nil || IsNil(o.AmountUSD) {
		return nil, false
	}
	return o.AmountUSD, true
}

// HasAmountUSD returns a boolean if a field has been set.
func (o *TokenApprovalData) HasAmountUSD() bool {
	if o != nil && !IsNil(o.AmountUSD) {
		return true
	}

	return false
}

// SetAmountUSD gets a reference to the given string and assigns it to the AmountUSD field.
func (o *TokenApprovalData) SetAmountUSD(v string) {
	o.AmountUSD = &v
}

// GetValueAtRisk returns the ValueAtRisk field value
func (o *TokenApprovalData) GetValueAtRisk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueAtRisk
}

// GetValueAtRiskOk returns a tuple with the ValueAtRisk field value
// and a boolean to check if the value has been set.
func (o *TokenApprovalData) GetValueAtRiskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueAtRisk, true
}

// SetValueAtRisk sets field value
func (o *TokenApprovalData) SetValueAtRisk(v string) {
	o.ValueAtRisk = v
}

// GetValueAtRiskUSD returns the ValueAtRiskUSD field value if set, zero value otherwise.
func (o *TokenApprovalData) GetValueAtRiskUSD() string {
	if o == nil || IsNil(o.ValueAtRiskUSD) {
		var ret string
		return ret
	}
	return *o.ValueAtRiskUSD
}

// GetValueAtRiskUSDOk returns a tuple with the ValueAtRiskUSD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenApprovalData) GetValueAtRiskUSDOk() (*string, bool) {
	if o == nil || IsNil(o.ValueAtRiskUSD) {
		return nil, false
	}
	return o.ValueAtRiskUSD, true
}

// HasValueAtRiskUSD returns a boolean if a field has been set.
func (o *TokenApprovalData) HasValueAtRiskUSD() bool {
	if o != nil && !IsNil(o.ValueAtRiskUSD) {
		return true
	}

	return false
}

// SetValueAtRiskUSD gets a reference to the given string and assigns it to the ValueAtRiskUSD field.
func (o *TokenApprovalData) SetValueAtRiskUSD(v string) {
	o.ValueAtRiskUSD = &v
}

// GetSpenderName returns the SpenderName field value if set, zero value otherwise.
func (o *TokenApprovalData) GetSpenderName() string {
	if o == nil || IsNil(o.SpenderName) {
		var ret string
		return ret
	}
	return *o.SpenderName
}

// GetSpenderNameOk returns a tuple with the SpenderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenApprovalData) GetSpenderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SpenderName) {
		return nil, false
	}
	return o.SpenderName, true
}

// HasSpenderName returns a boolean if a field has been set.
func (o *TokenApprovalData) HasSpenderName() bool {
	if o != nil && !IsNil(o.SpenderName) {
		return true
	}

	return false
}

// SetSpenderName gets a reference to the given string and assigns it to the SpenderName field.
func (o *TokenApprovalData) SetSpenderName(v string) {
	o.SpenderName = &v
}

func (o TokenApprovalData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenApprovalData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chainId"] = o.ChainId
	toSerialize["accountAddress"] = o.AccountAddress
	toSerialize["spenderAddress"] = o.SpenderAddress
	toSerialize["token"] = o.Token
	toSerialize["amount"] = o.Amount
	if !IsNil(o.AmountUSD) {
		toSerialize["amountUSD"] = o.AmountUSD
	}
	toSerialize["valueAtRisk"] = o.ValueAtRisk
	if !IsNil(o.ValueAtRiskUSD) {
		toSerialize["valueAtRiskUSD"] = o.ValueAtRiskUSD
	}
	if !IsNil(o.SpenderName) {
		toSerialize["spenderName"] = o.SpenderName
	}
	return toSerialize, nil
}

func (o *TokenApprovalData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chainId",
		"accountAddress",
		"spenderAddress",
		"token",
		"amount",
		"valueAtRisk",
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

	varTokenApprovalData := _TokenApprovalData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenApprovalData)

	if err != nil {
		return err
	}

	*o = TokenApprovalData(varTokenApprovalData)

	return err
}

type NullableTokenApprovalData struct {
	value *TokenApprovalData
	isSet bool
}

func (v NullableTokenApprovalData) Get() *TokenApprovalData {
	return v.value
}

func (v *NullableTokenApprovalData) Set(val *TokenApprovalData) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenApprovalData) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenApprovalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenApprovalData(val *TokenApprovalData) *NullableTokenApprovalData {
	return &NullableTokenApprovalData{value: val, isSet: true}
}

func (v NullableTokenApprovalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenApprovalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


