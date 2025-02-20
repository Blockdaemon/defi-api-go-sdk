# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | A timestamp represented in ISO 8601 format. | 
**Type** | [**TransactionType**](TransactionType.md) |  | 
**ChainID** | **string** | The unique identifier of a blockchain in CAIP-2 notation. | 
**BlockNumber** | **int32** | The number of a block. | 
**BlockHash** | **string** | The hash of a transaction or block. | 
**TxHash** | **string** | The hash of a transaction or block. | 
**ExplorerLink** | **string** | A link to a block explorer for a transaction or address. | 
**Status** | [**TransactionStatus**](TransactionStatus.md) |  | 
**Tenderly** | Pointer to **string** | A link to the Tenderly Explorer for a transaction. | [optional] 
**FromAccount** | [**Account**](Account.md) |  | 
**ToAccount** | [**Account**](Account.md) |  | 
**TokenAmount** | [**TokenAmount**](TokenAmount.md) |  | 
**GasDetails** | [**GasDetails**](GasDetails.md) |  | 
**Transfers** | [**[]Transfer**](Transfer.md) | The CAIP-2 identifier for a blockchain. | 

## Methods

### NewTransaction

`func NewTransaction(timestamp time.Time, type_ TransactionType, chainID string, blockNumber int32, blockHash string, txHash string, explorerLink string, status TransactionStatus, fromAccount Account, toAccount Account, tokenAmount TokenAmount, gasDetails GasDetails, transfers []Transfer, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *Transaction) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Transaction) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Transaction) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *Transaction) GetType() TransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*TransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v TransactionType)`

SetType sets Type field to given value.


### GetChainID

`func (o *Transaction) GetChainID() string`

GetChainID returns the ChainID field if non-nil, zero value otherwise.

### GetChainIDOk

`func (o *Transaction) GetChainIDOk() (*string, bool)`

GetChainIDOk returns a tuple with the ChainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainID

`func (o *Transaction) SetChainID(v string)`

SetChainID sets ChainID field to given value.


### GetBlockNumber

`func (o *Transaction) GetBlockNumber() int32`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *Transaction) GetBlockNumberOk() (*int32, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *Transaction) SetBlockNumber(v int32)`

SetBlockNumber sets BlockNumber field to given value.


### GetBlockHash

`func (o *Transaction) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *Transaction) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *Transaction) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetTxHash

`func (o *Transaction) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *Transaction) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *Transaction) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetExplorerLink

`func (o *Transaction) GetExplorerLink() string`

GetExplorerLink returns the ExplorerLink field if non-nil, zero value otherwise.

### GetExplorerLinkOk

`func (o *Transaction) GetExplorerLinkOk() (*string, bool)`

GetExplorerLinkOk returns a tuple with the ExplorerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorerLink

`func (o *Transaction) SetExplorerLink(v string)`

SetExplorerLink sets ExplorerLink field to given value.


### GetStatus

`func (o *Transaction) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.


### GetTenderly

`func (o *Transaction) GetTenderly() string`

GetTenderly returns the Tenderly field if non-nil, zero value otherwise.

### GetTenderlyOk

`func (o *Transaction) GetTenderlyOk() (*string, bool)`

GetTenderlyOk returns a tuple with the Tenderly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenderly

`func (o *Transaction) SetTenderly(v string)`

SetTenderly sets Tenderly field to given value.

### HasTenderly

`func (o *Transaction) HasTenderly() bool`

HasTenderly returns a boolean if a field has been set.

### GetFromAccount

`func (o *Transaction) GetFromAccount() Account`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *Transaction) GetFromAccountOk() (*Account, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *Transaction) SetFromAccount(v Account)`

SetFromAccount sets FromAccount field to given value.


### GetToAccount

`func (o *Transaction) GetToAccount() Account`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *Transaction) GetToAccountOk() (*Account, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *Transaction) SetToAccount(v Account)`

SetToAccount sets ToAccount field to given value.


### GetTokenAmount

`func (o *Transaction) GetTokenAmount() TokenAmount`

GetTokenAmount returns the TokenAmount field if non-nil, zero value otherwise.

### GetTokenAmountOk

`func (o *Transaction) GetTokenAmountOk() (*TokenAmount, bool)`

GetTokenAmountOk returns a tuple with the TokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAmount

`func (o *Transaction) SetTokenAmount(v TokenAmount)`

SetTokenAmount sets TokenAmount field to given value.


### GetGasDetails

`func (o *Transaction) GetGasDetails() GasDetails`

GetGasDetails returns the GasDetails field if non-nil, zero value otherwise.

### GetGasDetailsOk

`func (o *Transaction) GetGasDetailsOk() (*GasDetails, bool)`

GetGasDetailsOk returns a tuple with the GasDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasDetails

`func (o *Transaction) SetGasDetails(v GasDetails)`

SetGasDetails sets GasDetails field to given value.


### GetTransfers

`func (o *Transaction) GetTransfers() []Transfer`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *Transaction) GetTransfersOk() (*[]Transfer, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *Transaction) SetTransfers(v []Transfer)`

SetTransfers sets Transfers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


