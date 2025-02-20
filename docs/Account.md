# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human-readable name for the account. | [optional] 
**LogoURI** | Pointer to **string** | URI for the account&#39;s logo or icon. | [optional] 
**Address** | **string** | The address of an account. | 
**ExplorerLink** | Pointer to **string** | A link to a block explorer for a transaction or address. | [optional] 
**Metadata** | **map[string]string** | Any additional metadata sourced from platforms like OpenLabelInitiative, ENS, etc. | 

## Methods

### NewAccount

`func NewAccount(address string, metadata map[string]string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Account) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoURI

`func (o *Account) GetLogoURI() string`

GetLogoURI returns the LogoURI field if non-nil, zero value otherwise.

### GetLogoURIOk

`func (o *Account) GetLogoURIOk() (*string, bool)`

GetLogoURIOk returns a tuple with the LogoURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoURI

`func (o *Account) SetLogoURI(v string)`

SetLogoURI sets LogoURI field to given value.

### HasLogoURI

`func (o *Account) HasLogoURI() bool`

HasLogoURI returns a boolean if a field has been set.

### GetAddress

`func (o *Account) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Account) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Account) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetExplorerLink

`func (o *Account) GetExplorerLink() string`

GetExplorerLink returns the ExplorerLink field if non-nil, zero value otherwise.

### GetExplorerLinkOk

`func (o *Account) GetExplorerLinkOk() (*string, bool)`

GetExplorerLinkOk returns a tuple with the ExplorerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorerLink

`func (o *Account) SetExplorerLink(v string)`

SetExplorerLink sets ExplorerLink field to given value.

### HasExplorerLink

`func (o *Account) HasExplorerLink() bool`

HasExplorerLink returns a boolean if a field has been set.

### GetMetadata

`func (o *Account) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Account) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Account) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


