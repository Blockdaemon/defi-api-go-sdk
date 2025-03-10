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

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account Information about an account, including its address and metadata.
type Account struct {
	// Human-readable name for the account.
	Name *string `json:"name,omitempty"`
	// URI for the account's logo or icon.
	LogoURI *string `json:"logoURI,omitempty"`
	// The address of an account.
	Address string `json:"address"`
	// A link to a block explorer for a transaction or address.
	ExplorerLink *string `json:"explorerLink,omitempty"`
	// Any additional metadata sourced from platforms like OpenLabelInitiative, ENS, etc.
	Metadata map[string]string `json:"metadata"`
}

type _Account Account

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(address string, metadata map[string]string) *Account {
	this := Account{}
	this.Address = address
	this.Metadata = metadata
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Account) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Account) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Account) SetName(v string) {
	o.Name = &v
}

// GetLogoURI returns the LogoURI field value if set, zero value otherwise.
func (o *Account) GetLogoURI() string {
	if o == nil || IsNil(o.LogoURI) {
		var ret string
		return ret
	}
	return *o.LogoURI
}

// GetLogoURIOk returns a tuple with the LogoURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetLogoURIOk() (*string, bool) {
	if o == nil || IsNil(o.LogoURI) {
		return nil, false
	}
	return o.LogoURI, true
}

// HasLogoURI returns a boolean if a field has been set.
func (o *Account) HasLogoURI() bool {
	if o != nil && !IsNil(o.LogoURI) {
		return true
	}

	return false
}

// SetLogoURI gets a reference to the given string and assigns it to the LogoURI field.
func (o *Account) SetLogoURI(v string) {
	o.LogoURI = &v
}

// GetAddress returns the Address field value
func (o *Account) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Account) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *Account) SetAddress(v string) {
	o.Address = v
}

// GetExplorerLink returns the ExplorerLink field value if set, zero value otherwise.
func (o *Account) GetExplorerLink() string {
	if o == nil || IsNil(o.ExplorerLink) {
		var ret string
		return ret
	}
	return *o.ExplorerLink
}

// GetExplorerLinkOk returns a tuple with the ExplorerLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetExplorerLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ExplorerLink) {
		return nil, false
	}
	return o.ExplorerLink, true
}

// HasExplorerLink returns a boolean if a field has been set.
func (o *Account) HasExplorerLink() bool {
	if o != nil && !IsNil(o.ExplorerLink) {
		return true
	}

	return false
}

// SetExplorerLink gets a reference to the given string and assigns it to the ExplorerLink field.
func (o *Account) SetExplorerLink(v string) {
	o.ExplorerLink = &v
}

// GetMetadata returns the Metadata field value
func (o *Account) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Account) GetMetadataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Account) SetMetadata(v map[string]string) {
	o.Metadata = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LogoURI) {
		toSerialize["logoURI"] = o.LogoURI
	}
	toSerialize["address"] = o.Address
	if !IsNil(o.ExplorerLink) {
		toSerialize["explorerLink"] = o.ExplorerLink
	}
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *Account) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"metadata",
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

	varAccount := _Account{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccount)

	if err != nil {
		return err
	}

	*o = Account(varAccount)

	return err
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


