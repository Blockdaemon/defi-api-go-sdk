# TransactionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPage** | **string** | The next page of the paginated response. | 
**Items** | [**[]Transaction**](Transaction.md) | List of transactions. | 

## Methods

### NewTransactionsResponse

`func NewTransactionsResponse(nextPage string, items []Transaction, ) *TransactionsResponse`

NewTransactionsResponse instantiates a new TransactionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsResponseWithDefaults

`func NewTransactionsResponseWithDefaults() *TransactionsResponse`

NewTransactionsResponseWithDefaults instantiates a new TransactionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPage

`func (o *TransactionsResponse) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *TransactionsResponse) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *TransactionsResponse) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.


### GetItems

`func (o *TransactionsResponse) GetItems() []Transaction`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TransactionsResponse) GetItemsOk() (*[]Transaction, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TransactionsResponse) SetItems(v []Transaction)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


