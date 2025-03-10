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

// checks if the AllApprovalsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllApprovalsResponse{}

// AllApprovalsResponse A response containing token approvals and combined sums for accounts and chains.
type AllApprovalsResponse struct {
	// List of token approval data for the requested chains and accounts.
	Approvals []TokenApprovalData `json:"approvals"`
	PortfolioViewSum PortfolioViewSum `json:"portfolioViewSum"`
	// Array of chain IDs for which balances are missing
	HasMissingApprovals []string `json:"hasMissingApprovals,omitempty"`
}

type _AllApprovalsResponse AllApprovalsResponse

// NewAllApprovalsResponse instantiates a new AllApprovalsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllApprovalsResponse(approvals []TokenApprovalData, portfolioViewSum PortfolioViewSum) *AllApprovalsResponse {
	this := AllApprovalsResponse{}
	this.Approvals = approvals
	this.PortfolioViewSum = portfolioViewSum
	return &this
}

// NewAllApprovalsResponseWithDefaults instantiates a new AllApprovalsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllApprovalsResponseWithDefaults() *AllApprovalsResponse {
	this := AllApprovalsResponse{}
	return &this
}

// GetApprovals returns the Approvals field value
func (o *AllApprovalsResponse) GetApprovals() []TokenApprovalData {
	if o == nil {
		var ret []TokenApprovalData
		return ret
	}

	return o.Approvals
}

// GetApprovalsOk returns a tuple with the Approvals field value
// and a boolean to check if the value has been set.
func (o *AllApprovalsResponse) GetApprovalsOk() ([]TokenApprovalData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approvals, true
}

// SetApprovals sets field value
func (o *AllApprovalsResponse) SetApprovals(v []TokenApprovalData) {
	o.Approvals = v
}

// GetPortfolioViewSum returns the PortfolioViewSum field value
func (o *AllApprovalsResponse) GetPortfolioViewSum() PortfolioViewSum {
	if o == nil {
		var ret PortfolioViewSum
		return ret
	}

	return o.PortfolioViewSum
}

// GetPortfolioViewSumOk returns a tuple with the PortfolioViewSum field value
// and a boolean to check if the value has been set.
func (o *AllApprovalsResponse) GetPortfolioViewSumOk() (*PortfolioViewSum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortfolioViewSum, true
}

// SetPortfolioViewSum sets field value
func (o *AllApprovalsResponse) SetPortfolioViewSum(v PortfolioViewSum) {
	o.PortfolioViewSum = v
}

// GetHasMissingApprovals returns the HasMissingApprovals field value if set, zero value otherwise.
func (o *AllApprovalsResponse) GetHasMissingApprovals() []string {
	if o == nil || IsNil(o.HasMissingApprovals) {
		var ret []string
		return ret
	}
	return o.HasMissingApprovals
}

// GetHasMissingApprovalsOk returns a tuple with the HasMissingApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllApprovalsResponse) GetHasMissingApprovalsOk() ([]string, bool) {
	if o == nil || IsNil(o.HasMissingApprovals) {
		return nil, false
	}
	return o.HasMissingApprovals, true
}

// HasHasMissingApprovals returns a boolean if a field has been set.
func (o *AllApprovalsResponse) HasHasMissingApprovals() bool {
	if o != nil && !IsNil(o.HasMissingApprovals) {
		return true
	}

	return false
}

// SetHasMissingApprovals gets a reference to the given []string and assigns it to the HasMissingApprovals field.
func (o *AllApprovalsResponse) SetHasMissingApprovals(v []string) {
	o.HasMissingApprovals = v
}

func (o AllApprovalsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllApprovalsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["approvals"] = o.Approvals
	toSerialize["portfolioViewSum"] = o.PortfolioViewSum
	if !IsNil(o.HasMissingApprovals) {
		toSerialize["hasMissingApprovals"] = o.HasMissingApprovals
	}
	return toSerialize, nil
}

func (o *AllApprovalsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approvals",
		"portfolioViewSum",
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

	varAllApprovalsResponse := _AllApprovalsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAllApprovalsResponse)

	if err != nil {
		return err
	}

	*o = AllApprovalsResponse(varAllApprovalsResponse)

	return err
}

type NullableAllApprovalsResponse struct {
	value *AllApprovalsResponse
	isSet bool
}

func (v NullableAllApprovalsResponse) Get() *AllApprovalsResponse {
	return v.value
}

func (v *NullableAllApprovalsResponse) Set(val *AllApprovalsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllApprovalsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllApprovalsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllApprovalsResponse(val *AllApprovalsResponse) *NullableAllApprovalsResponse {
	return &NullableAllApprovalsResponse{value: val, isSet: true}
}

func (v NullableAllApprovalsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllApprovalsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


