# DocumentUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentType** | [**DocumentType**](DocumentType.md) |  | 
**DocumentSubType** | Pointer to **string** |  | [optional] 
**Content** | **string** |  | 
**MimeType** | **string** |  | 

## Methods

### NewDocumentUpload

`func NewDocumentUpload(documentType DocumentType, content string, mimeType string, ) *DocumentUpload`

NewDocumentUpload instantiates a new DocumentUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentUploadWithDefaults

`func NewDocumentUploadWithDefaults() *DocumentUpload`

NewDocumentUploadWithDefaults instantiates a new DocumentUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentType

`func (o *DocumentUpload) GetDocumentType() DocumentType`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *DocumentUpload) GetDocumentTypeOk() (*DocumentType, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *DocumentUpload) SetDocumentType(v DocumentType)`

SetDocumentType sets DocumentType field to given value.


### GetDocumentSubType

`func (o *DocumentUpload) GetDocumentSubType() string`

GetDocumentSubType returns the DocumentSubType field if non-nil, zero value otherwise.

### GetDocumentSubTypeOk

`func (o *DocumentUpload) GetDocumentSubTypeOk() (*string, bool)`

GetDocumentSubTypeOk returns a tuple with the DocumentSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSubType

`func (o *DocumentUpload) SetDocumentSubType(v string)`

SetDocumentSubType sets DocumentSubType field to given value.

### HasDocumentSubType

`func (o *DocumentUpload) HasDocumentSubType() bool`

HasDocumentSubType returns a boolean if a field has been set.

### GetContent

`func (o *DocumentUpload) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DocumentUpload) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DocumentUpload) SetContent(v string)`

SetContent sets Content field to given value.


### GetMimeType

`func (o *DocumentUpload) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *DocumentUpload) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *DocumentUpload) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


