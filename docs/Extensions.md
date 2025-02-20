# Extensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verified** | Pointer to **bool** | Describes whether the token data has been verified by Blockdaemon | [optional] 
**Homepage** | Pointer to **string** | The homepage of the token issuer | [optional] 
**Unknown** | Pointer to **bool** |  | [optional] 

## Methods

### NewExtensions

`func NewExtensions() *Extensions`

NewExtensions instantiates a new Extensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionsWithDefaults

`func NewExtensionsWithDefaults() *Extensions`

NewExtensionsWithDefaults instantiates a new Extensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerified

`func (o *Extensions) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Extensions) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Extensions) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *Extensions) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetHomepage

`func (o *Extensions) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *Extensions) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *Extensions) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *Extensions) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetUnknown

`func (o *Extensions) GetUnknown() bool`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *Extensions) GetUnknownOk() (*bool, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *Extensions) SetUnknown(v bool)`

SetUnknown sets Unknown field to given value.

### HasUnknown

`func (o *Extensions) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


