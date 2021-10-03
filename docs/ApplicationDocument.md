# ApplicationDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DocumentType** | [**DocumentType**](DocumentType.md) |  | 
**DocumentSubType** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewApplicationDocument

`func NewApplicationDocument(id string, documentType DocumentType, createdAt time.Time, ) *ApplicationDocument`

NewApplicationDocument instantiates a new ApplicationDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDocumentWithDefaults

`func NewApplicationDocumentWithDefaults() *ApplicationDocument`

NewApplicationDocumentWithDefaults instantiates a new ApplicationDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationDocument) SetId(v string)`

SetId sets Id field to given value.


### GetDocumentType

`func (o *ApplicationDocument) GetDocumentType() DocumentType`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ApplicationDocument) GetDocumentTypeOk() (*DocumentType, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ApplicationDocument) SetDocumentType(v DocumentType)`

SetDocumentType sets DocumentType field to given value.


### GetDocumentSubType

`func (o *ApplicationDocument) GetDocumentSubType() string`

GetDocumentSubType returns the DocumentSubType field if non-nil, zero value otherwise.

### GetDocumentSubTypeOk

`func (o *ApplicationDocument) GetDocumentSubTypeOk() (*string, bool)`

GetDocumentSubTypeOk returns a tuple with the DocumentSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSubType

`func (o *ApplicationDocument) SetDocumentSubType(v string)`

SetDocumentSubType sets DocumentSubType field to given value.

### HasDocumentSubType

`func (o *ApplicationDocument) HasDocumentSubType() bool`

HasDocumentSubType returns a boolean if a field has been set.

### GetMimeType

`func (o *ApplicationDocument) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ApplicationDocument) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ApplicationDocument) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ApplicationDocument) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationDocument) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationDocument) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationDocument) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


