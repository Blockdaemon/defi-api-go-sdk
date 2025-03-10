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

// checks if the Chain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Chain{}

// Chain Details about a supported blockchain.
type Chain struct {
	// The name of a blockchain.
	ChainName string `json:"chainName"`
	// The unique identifier of a blockchain in CAIP-2 notation.
	ChainID string `json:"chainID" validate:"regexp=^[-a-z0-9]{3,8}:[-_a-zA-Z0-9]{1,32}$"`
	NetworkType *NetworkType `json:"networkType,omitempty"`
	ChainType *ChainType `json:"chainType,omitempty"`
	// The type of network (mainnet, testnet, or devnet).
	ChainIconURI *string `json:"chainIconURI,omitempty"`
	// estimated block confirmation time in seconds
	BlockTime *int32 `json:"blockTime,omitempty"`
	// number of additional blocks required to ensure the transaction is not reversed
	BlockConfirmations *int32 `json:"blockConfirmations,omitempty"`
	// List of RPC gateway URLs.
	Rpc []string `json:"rpc"`
	NativeCurrency *NativeCurrency `json:"nativeCurrency,omitempty"`
	// List of block explorer URLs.
	BlockExplorerUrls []string `json:"blockExplorerUrls,omitempty"`
}

type _Chain Chain

// NewChain instantiates a new Chain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChain(chainName string, chainID string, rpc []string) *Chain {
	this := Chain{}
	this.ChainName = chainName
	this.ChainID = chainID
	this.Rpc = rpc
	return &this
}

// NewChainWithDefaults instantiates a new Chain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChainWithDefaults() *Chain {
	this := Chain{}
	return &this
}

// GetChainName returns the ChainName field value
func (o *Chain) GetChainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainName
}

// GetChainNameOk returns a tuple with the ChainName field value
// and a boolean to check if the value has been set.
func (o *Chain) GetChainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainName, true
}

// SetChainName sets field value
func (o *Chain) SetChainName(v string) {
	o.ChainName = v
}

// GetChainID returns the ChainID field value
func (o *Chain) GetChainID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainID
}

// GetChainIDOk returns a tuple with the ChainID field value
// and a boolean to check if the value has been set.
func (o *Chain) GetChainIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainID, true
}

// SetChainID sets field value
func (o *Chain) SetChainID(v string) {
	o.ChainID = v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *Chain) GetNetworkType() NetworkType {
	if o == nil || IsNil(o.NetworkType) {
		var ret NetworkType
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chain) GetNetworkTypeOk() (*NetworkType, bool) {
	if o == nil || IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *Chain) HasNetworkType() bool {
	if o != nil && !IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given NetworkType and assigns it to the NetworkType field.
func (o *Chain) SetNetworkType(v NetworkType) {
	o.NetworkType = &v
}

// GetChainType returns the ChainType field value if set, zero value otherwise.
func (o *Chain) GetChainType() ChainType {
	if o == nil || IsNil(o.ChainType) {
		var ret ChainType
		return ret
	}
	return *o.ChainType
}

// GetChainTypeOk returns a tuple with the ChainType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chain) GetChainTypeOk() (*ChainType, bool) {
	if o == nil || IsNil(o.ChainType) {
		return nil, false
	}
	return o.ChainType, true
}

// HasChainType returns a boolean if a field has been set.
func (o *Chain) HasChainType() bool {
	if o != nil && !IsNil(o.ChainType) {
		return true
	}

	return false
}

// SetChainType gets a reference to the given ChainType and assigns it to the ChainType field.
func (o *Chain) SetChainType(v ChainType) {
	o.ChainType = &v
}

// GetChainIconURI returns the ChainIconURI field value if set, zero value otherwise.
func (o *Chain) GetChainIconURI() string {
	if o == nil || IsNil(o.ChainIconURI) {
		var ret string
		return ret
	}
	return *o.ChainIconURI
}

// GetChainIconURIOk returns a tuple with the ChainIconURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chain) GetChainIconURIOk() (*string, bool) {
	if o == nil || IsNil(o.ChainIconURI) {
		return nil, false
	}
	return o.ChainIconURI, true
}

// HasChainIconURI returns a boolean if a field has been set.
func (o *Chain) HasChainIconURI() bool {
	if o != nil && !IsNil(o.ChainIconURI) {
		return true
	}

	return false
}

// SetChainIconURI gets a reference to the given string and assigns it to the ChainIconURI field.
func (o *Chain) SetChainIconURI(v string) {
	o.ChainIconURI = &v
}

// GetBlockTime returns the BlockTime field value if set, zero value otherwise.
func (o *Chain) GetBlockTime() int32 {
	if o == nil || IsNil(o.BlockTime) {
		var ret int32
		return ret
	}
	return *o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chain) GetBlockTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockTime) {
		return nil, false
	}
	return o.BlockTime, true
}

// HasBlockTime returns a boolean if a field has been set.
func (o *Chain) HasBlockTime() bool {
	if o != nil && !IsNil(o.BlockTime) {
		return true
	}

	return false
}

// SetBlockTime gets a reference to the given int32 and assigns it to the BlockTime field.
func (o *Chain) SetBlockTime(v int32) {
	o.BlockTime = &v
}

// GetBlockConfirmations returns the BlockConfirmations field value if set, zero value otherwise.
func (o *Chain) GetBlockConfirmations() int32 {
	if o == nil || IsNil(o.BlockConfirmations) {
		var ret int32
		return ret
	}
	return *o.BlockConfirmations
}

// GetBlockConfirmationsOk returns a tuple with the BlockConfirmations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chain) GetBlockConfirmationsOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockConfirmations) {
		return nil, false
	}
	return o.BlockConfirmations, true
}

// HasBlockConfirmations returns a boolean if a field has been set.
func (o *Chain) HasBlockConfirmations() bool {
	if o != nil && !IsNil(o.BlockConfirmations) {
		return true
	}

	return false
}

// SetBlockConfirmations gets a reference to the given int32 and assigns it to the BlockConfirmations field.
func (o *Chain) SetBlockConfirmations(v int32) {
	o.BlockConfirmations = &v
}

// GetRpc returns the Rpc field value
func (o *Chain) GetRpc() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Rpc
}

// GetRpcOk returns a tuple with the Rpc field value
// and a boolean to check if the value has been set.
func (o *Chain) GetRpcOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rpc, true
}

// SetRpc sets field value
func (o *Chain) SetRpc(v []string) {
	o.Rpc = v
}

// GetNativeCurrency returns the NativeCurrency field value if set, zero value otherwise.
func (o *Chain) GetNativeCurrency() NativeCurrency {
	if o == nil || IsNil(o.NativeCurrency) {
		var ret NativeCurrency
		return ret
	}
	return *o.NativeCurrency
}

// GetNativeCurrencyOk returns a tuple with the NativeCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chain) GetNativeCurrencyOk() (*NativeCurrency, bool) {
	if o == nil || IsNil(o.NativeCurrency) {
		return nil, false
	}
	return o.NativeCurrency, true
}

// HasNativeCurrency returns a boolean if a field has been set.
func (o *Chain) HasNativeCurrency() bool {
	if o != nil && !IsNil(o.NativeCurrency) {
		return true
	}

	return false
}

// SetNativeCurrency gets a reference to the given NativeCurrency and assigns it to the NativeCurrency field.
func (o *Chain) SetNativeCurrency(v NativeCurrency) {
	o.NativeCurrency = &v
}

// GetBlockExplorerUrls returns the BlockExplorerUrls field value if set, zero value otherwise.
func (o *Chain) GetBlockExplorerUrls() []string {
	if o == nil || IsNil(o.BlockExplorerUrls) {
		var ret []string
		return ret
	}
	return o.BlockExplorerUrls
}

// GetBlockExplorerUrlsOk returns a tuple with the BlockExplorerUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chain) GetBlockExplorerUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockExplorerUrls) {
		return nil, false
	}
	return o.BlockExplorerUrls, true
}

// HasBlockExplorerUrls returns a boolean if a field has been set.
func (o *Chain) HasBlockExplorerUrls() bool {
	if o != nil && !IsNil(o.BlockExplorerUrls) {
		return true
	}

	return false
}

// SetBlockExplorerUrls gets a reference to the given []string and assigns it to the BlockExplorerUrls field.
func (o *Chain) SetBlockExplorerUrls(v []string) {
	o.BlockExplorerUrls = v
}

func (o Chain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Chain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chainName"] = o.ChainName
	toSerialize["chainID"] = o.ChainID
	if !IsNil(o.NetworkType) {
		toSerialize["networkType"] = o.NetworkType
	}
	if !IsNil(o.ChainType) {
		toSerialize["chainType"] = o.ChainType
	}
	if !IsNil(o.ChainIconURI) {
		toSerialize["chainIconURI"] = o.ChainIconURI
	}
	if !IsNil(o.BlockTime) {
		toSerialize["blockTime"] = o.BlockTime
	}
	if !IsNil(o.BlockConfirmations) {
		toSerialize["blockConfirmations"] = o.BlockConfirmations
	}
	toSerialize["rpc"] = o.Rpc
	if !IsNil(o.NativeCurrency) {
		toSerialize["nativeCurrency"] = o.NativeCurrency
	}
	if !IsNil(o.BlockExplorerUrls) {
		toSerialize["blockExplorerUrls"] = o.BlockExplorerUrls
	}
	return toSerialize, nil
}

func (o *Chain) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chainName",
		"chainID",
		"rpc",
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

	varChain := _Chain{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChain)

	if err != nil {
		return err
	}

	*o = Chain(varChain)

	return err
}

type NullableChain struct {
	value *Chain
	isSet bool
}

func (v NullableChain) Get() *Chain {
	return v.value
}

func (v *NullableChain) Set(val *Chain) {
	v.value = val
	v.isSet = true
}

func (v NullableChain) IsSet() bool {
	return v.isSet
}

func (v *NullableChain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChain(val *Chain) *NullableChain {
	return &NullableChain{value: val, isSet: true}
}

func (v NullableChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


