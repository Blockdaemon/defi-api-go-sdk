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

// checks if the TransactionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionsResponse{}

// TransactionsResponse A response containing transaction details for an account.
type TransactionsResponse struct {
	// The next page of the paginated response.
	NextPage string `json:"nextPage"`
	// List of transactions.
	Items []Transaction `json:"items"`
}

type _TransactionsResponse TransactionsResponse

// NewTransactionsResponse instantiates a new TransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsResponse(nextPage string, items []Transaction) *TransactionsResponse {
	this := TransactionsResponse{}
	this.NextPage = nextPage
	this.Items = items
	return &this
}

// NewTransactionsResponseWithDefaults instantiates a new TransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsResponseWithDefaults() *TransactionsResponse {
	this := TransactionsResponse{}
	return &this
}

// GetNextPage returns the NextPage field value
func (o *TransactionsResponse) GetNextPage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value
// and a boolean to check if the value has been set.
func (o *TransactionsResponse) GetNextPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPage, true
}

// SetNextPage sets field value
func (o *TransactionsResponse) SetNextPage(v string) {
	o.NextPage = v
}

// GetItems returns the Items field value
func (o *TransactionsResponse) GetItems() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *TransactionsResponse) GetItemsOk() ([]Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *TransactionsResponse) SetItems(v []Transaction) {
	o.Items = v
}

func (o TransactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nextPage"] = o.NextPage
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *TransactionsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nextPage",
		"items",
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

	varTransactionsResponse := _TransactionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionsResponse)

	if err != nil {
		return err
	}

	*o = TransactionsResponse(varTransactionsResponse)

	return err
}

type NullableTransactionsResponse struct {
	value *TransactionsResponse
	isSet bool
}

func (v NullableTransactionsResponse) Get() *TransactionsResponse {
	return v.value
}

func (v *NullableTransactionsResponse) Set(val *TransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsResponse(val *TransactionsResponse) *NullableTransactionsResponse {
	return &NullableTransactionsResponse{value: val, isSet: true}
}

func (v NullableTransactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


