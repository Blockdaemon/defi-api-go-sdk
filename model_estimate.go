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

// checks if the Estimate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Estimate{}

// Estimate An estimate of the costs and outcome of a step in a route.
type Estimate struct {
	IntegrationDetails IntegrationDetails `json:"integrationDetails"`
	// The address of a token, account, or smart contract.
	ApprovalAddress string `json:"approvalAddress" validate:"regexp=^0x[a-fA-F0-9]{40}$"`
	// The minimum amount to be received after the transaction, specified as a string to maintain precision.
	ToAmountMin *string `json:"toAmountMin,omitempty" validate:"regexp=^([1-9][0-9]*|0)(\\\\.[0-9]+)?$"`
	// The estimated amount to be received after the transaction, specified as a string to maintain precision.
	ToAmount string `json:"toAmount" validate:"regexp=^([1-9][0-9]*|0)(\\\\.[0-9]+)?$"`
	// The amount of tokens to be sent, including decimals.
	FromAmount string `json:"fromAmount" validate:"regexp=^([1-9][0-9]*|0)(\\\\.[0-9]+)?$"`
	// Transaction fees, if any. This can be empty if there are no extra fees.
	FeeCosts []FeeCost `json:"feeCosts"`
	// The CAIP-2 identifier for a blockchain.
	GasCosts []GasCost `json:"gasCosts"`
	// The estimated time for the transaction to be executed, in seconds.
	ExecutionDuration *int32 `json:"executionDuration,omitempty"`
	// The value of the amount to be sent in USD.
	FromAmountUSD *string `json:"fromAmountUSD,omitempty" validate:"regexp=^(0(\\\\.[0-9]+)?|[1-9][0-9]*(\\\\.[0-9]+)?)$"`
	// The estimated value of the amount to be received in USD.
	ToAmountUSD *string `json:"toAmountUSD,omitempty" validate:"regexp=^(0(\\\\.[0-9]+)?|[1-9][0-9]*(\\\\.[0-9]+)?)$"`
}

type _Estimate Estimate

// NewEstimate instantiates a new Estimate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimate(integrationDetails IntegrationDetails, approvalAddress string, toAmount string, fromAmount string, feeCosts []FeeCost, gasCosts []GasCost) *Estimate {
	this := Estimate{}
	this.IntegrationDetails = integrationDetails
	this.ApprovalAddress = approvalAddress
	this.ToAmount = toAmount
	this.FromAmount = fromAmount
	this.FeeCosts = feeCosts
	this.GasCosts = gasCosts
	return &this
}

// NewEstimateWithDefaults instantiates a new Estimate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimateWithDefaults() *Estimate {
	this := Estimate{}
	return &this
}

// GetIntegrationDetails returns the IntegrationDetails field value
func (o *Estimate) GetIntegrationDetails() IntegrationDetails {
	if o == nil {
		var ret IntegrationDetails
		return ret
	}

	return o.IntegrationDetails
}

// GetIntegrationDetailsOk returns a tuple with the IntegrationDetails field value
// and a boolean to check if the value has been set.
func (o *Estimate) GetIntegrationDetailsOk() (*IntegrationDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationDetails, true
}

// SetIntegrationDetails sets field value
func (o *Estimate) SetIntegrationDetails(v IntegrationDetails) {
	o.IntegrationDetails = v
}

// GetApprovalAddress returns the ApprovalAddress field value
func (o *Estimate) GetApprovalAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApprovalAddress
}

// GetApprovalAddressOk returns a tuple with the ApprovalAddress field value
// and a boolean to check if the value has been set.
func (o *Estimate) GetApprovalAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalAddress, true
}

// SetApprovalAddress sets field value
func (o *Estimate) SetApprovalAddress(v string) {
	o.ApprovalAddress = v
}

// GetToAmountMin returns the ToAmountMin field value if set, zero value otherwise.
func (o *Estimate) GetToAmountMin() string {
	if o == nil || IsNil(o.ToAmountMin) {
		var ret string
		return ret
	}
	return *o.ToAmountMin
}

// GetToAmountMinOk returns a tuple with the ToAmountMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimate) GetToAmountMinOk() (*string, bool) {
	if o == nil || IsNil(o.ToAmountMin) {
		return nil, false
	}
	return o.ToAmountMin, true
}

// HasToAmountMin returns a boolean if a field has been set.
func (o *Estimate) HasToAmountMin() bool {
	if o != nil && !IsNil(o.ToAmountMin) {
		return true
	}

	return false
}

// SetToAmountMin gets a reference to the given string and assigns it to the ToAmountMin field.
func (o *Estimate) SetToAmountMin(v string) {
	o.ToAmountMin = &v
}

// GetToAmount returns the ToAmount field value
func (o *Estimate) GetToAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAmount
}

// GetToAmountOk returns a tuple with the ToAmount field value
// and a boolean to check if the value has been set.
func (o *Estimate) GetToAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAmount, true
}

// SetToAmount sets field value
func (o *Estimate) SetToAmount(v string) {
	o.ToAmount = v
}

// GetFromAmount returns the FromAmount field value
func (o *Estimate) GetFromAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAmount
}

// GetFromAmountOk returns a tuple with the FromAmount field value
// and a boolean to check if the value has been set.
func (o *Estimate) GetFromAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAmount, true
}

// SetFromAmount sets field value
func (o *Estimate) SetFromAmount(v string) {
	o.FromAmount = v
}

// GetFeeCosts returns the FeeCosts field value
func (o *Estimate) GetFeeCosts() []FeeCost {
	if o == nil {
		var ret []FeeCost
		return ret
	}

	return o.FeeCosts
}

// GetFeeCostsOk returns a tuple with the FeeCosts field value
// and a boolean to check if the value has been set.
func (o *Estimate) GetFeeCostsOk() ([]FeeCost, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeeCosts, true
}

// SetFeeCosts sets field value
func (o *Estimate) SetFeeCosts(v []FeeCost) {
	o.FeeCosts = v
}

// GetGasCosts returns the GasCosts field value
func (o *Estimate) GetGasCosts() []GasCost {
	if o == nil {
		var ret []GasCost
		return ret
	}

	return o.GasCosts
}

// GetGasCostsOk returns a tuple with the GasCosts field value
// and a boolean to check if the value has been set.
func (o *Estimate) GetGasCostsOk() ([]GasCost, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasCosts, true
}

// SetGasCosts sets field value
func (o *Estimate) SetGasCosts(v []GasCost) {
	o.GasCosts = v
}

// GetExecutionDuration returns the ExecutionDuration field value if set, zero value otherwise.
func (o *Estimate) GetExecutionDuration() int32 {
	if o == nil || IsNil(o.ExecutionDuration) {
		var ret int32
		return ret
	}
	return *o.ExecutionDuration
}

// GetExecutionDurationOk returns a tuple with the ExecutionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimate) GetExecutionDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.ExecutionDuration) {
		return nil, false
	}
	return o.ExecutionDuration, true
}

// HasExecutionDuration returns a boolean if a field has been set.
func (o *Estimate) HasExecutionDuration() bool {
	if o != nil && !IsNil(o.ExecutionDuration) {
		return true
	}

	return false
}

// SetExecutionDuration gets a reference to the given int32 and assigns it to the ExecutionDuration field.
func (o *Estimate) SetExecutionDuration(v int32) {
	o.ExecutionDuration = &v
}

// GetFromAmountUSD returns the FromAmountUSD field value if set, zero value otherwise.
func (o *Estimate) GetFromAmountUSD() string {
	if o == nil || IsNil(o.FromAmountUSD) {
		var ret string
		return ret
	}
	return *o.FromAmountUSD
}

// GetFromAmountUSDOk returns a tuple with the FromAmountUSD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimate) GetFromAmountUSDOk() (*string, bool) {
	if o == nil || IsNil(o.FromAmountUSD) {
		return nil, false
	}
	return o.FromAmountUSD, true
}

// HasFromAmountUSD returns a boolean if a field has been set.
func (o *Estimate) HasFromAmountUSD() bool {
	if o != nil && !IsNil(o.FromAmountUSD) {
		return true
	}

	return false
}

// SetFromAmountUSD gets a reference to the given string and assigns it to the FromAmountUSD field.
func (o *Estimate) SetFromAmountUSD(v string) {
	o.FromAmountUSD = &v
}

// GetToAmountUSD returns the ToAmountUSD field value if set, zero value otherwise.
func (o *Estimate) GetToAmountUSD() string {
	if o == nil || IsNil(o.ToAmountUSD) {
		var ret string
		return ret
	}
	return *o.ToAmountUSD
}

// GetToAmountUSDOk returns a tuple with the ToAmountUSD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimate) GetToAmountUSDOk() (*string, bool) {
	if o == nil || IsNil(o.ToAmountUSD) {
		return nil, false
	}
	return o.ToAmountUSD, true
}

// HasToAmountUSD returns a boolean if a field has been set.
func (o *Estimate) HasToAmountUSD() bool {
	if o != nil && !IsNil(o.ToAmountUSD) {
		return true
	}

	return false
}

// SetToAmountUSD gets a reference to the given string and assigns it to the ToAmountUSD field.
func (o *Estimate) SetToAmountUSD(v string) {
	o.ToAmountUSD = &v
}

func (o Estimate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Estimate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["integrationDetails"] = o.IntegrationDetails
	toSerialize["approvalAddress"] = o.ApprovalAddress
	if !IsNil(o.ToAmountMin) {
		toSerialize["toAmountMin"] = o.ToAmountMin
	}
	toSerialize["toAmount"] = o.ToAmount
	toSerialize["fromAmount"] = o.FromAmount
	toSerialize["feeCosts"] = o.FeeCosts
	toSerialize["gasCosts"] = o.GasCosts
	if !IsNil(o.ExecutionDuration) {
		toSerialize["executionDuration"] = o.ExecutionDuration
	}
	if !IsNil(o.FromAmountUSD) {
		toSerialize["fromAmountUSD"] = o.FromAmountUSD
	}
	if !IsNil(o.ToAmountUSD) {
		toSerialize["toAmountUSD"] = o.ToAmountUSD
	}
	return toSerialize, nil
}

func (o *Estimate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"integrationDetails",
		"approvalAddress",
		"toAmount",
		"fromAmount",
		"feeCosts",
		"gasCosts",
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

	varEstimate := _Estimate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEstimate)

	if err != nil {
		return err
	}

	*o = Estimate(varEstimate)

	return err
}

type NullableEstimate struct {
	value *Estimate
	isSet bool
}

func (v NullableEstimate) Get() *Estimate {
	return v.value
}

func (v *NullableEstimate) Set(val *Estimate) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimate) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimate(val *Estimate) *NullableEstimate {
	return &NullableEstimate{value: val, isSet: true}
}

func (v NullableEstimate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


