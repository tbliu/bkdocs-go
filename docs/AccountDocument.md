# AccountDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DocumentType** | [**DocumentType**](DocumentType.md) |  | 
**DocumentSubType** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewAccountDocument

`func NewAccountDocument(id string, documentType DocumentType, createdAt time.Time, ) *AccountDocument`

NewAccountDocument instantiates a new AccountDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDocumentWithDefaults

`func NewAccountDocumentWithDefaults() *AccountDocument`

NewAccountDocumentWithDefaults instantiates a new AccountDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountDocument) SetId(v string)`

SetId sets Id field to given value.


### GetDocumentType

`func (o *AccountDocument) GetDocumentType() DocumentType`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *AccountDocument) GetDocumentTypeOk() (*DocumentType, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *AccountDocument) SetDocumentType(v DocumentType)`

SetDocumentType sets DocumentType field to given value.


### GetDocumentSubType

`func (o *AccountDocument) GetDocumentSubType() string`

GetDocumentSubType returns the DocumentSubType field if non-nil, zero value otherwise.

### GetDocumentSubTypeOk

`func (o *AccountDocument) GetDocumentSubTypeOk() (*string, bool)`

GetDocumentSubTypeOk returns a tuple with the DocumentSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSubType

`func (o *AccountDocument) SetDocumentSubType(v string)`

SetDocumentSubType sets DocumentSubType field to given value.

### HasDocumentSubType

`func (o *AccountDocument) HasDocumentSubType() bool`

HasDocumentSubType returns a boolean if a field has been set.

### GetMimeType

`func (o *AccountDocument) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AccountDocument) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AccountDocument) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *AccountDocument) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountDocument) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountDocument) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountDocument) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


