# DocumentUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentType** | [**DocumentType**](DocumentType.md) |  | 
**DocumentSubType** | Pointer to **string** |  | [optional] 
**Content** | **string** |  | 
**MimeType** | **string** |  | 

## Methods

### NewDocumentUploadRequest

`func NewDocumentUploadRequest(documentType DocumentType, content string, mimeType string, ) *DocumentUploadRequest`

NewDocumentUploadRequest instantiates a new DocumentUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentUploadRequestWithDefaults

`func NewDocumentUploadRequestWithDefaults() *DocumentUploadRequest`

NewDocumentUploadRequestWithDefaults instantiates a new DocumentUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentType

`func (o *DocumentUploadRequest) GetDocumentType() DocumentType`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *DocumentUploadRequest) GetDocumentTypeOk() (*DocumentType, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *DocumentUploadRequest) SetDocumentType(v DocumentType)`

SetDocumentType sets DocumentType field to given value.


### GetDocumentSubType

`func (o *DocumentUploadRequest) GetDocumentSubType() string`

GetDocumentSubType returns the DocumentSubType field if non-nil, zero value otherwise.

### GetDocumentSubTypeOk

`func (o *DocumentUploadRequest) GetDocumentSubTypeOk() (*string, bool)`

GetDocumentSubTypeOk returns a tuple with the DocumentSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSubType

`func (o *DocumentUploadRequest) SetDocumentSubType(v string)`

SetDocumentSubType sets DocumentSubType field to given value.

### HasDocumentSubType

`func (o *DocumentUploadRequest) HasDocumentSubType() bool`

HasDocumentSubType returns a boolean if a field has been set.

### GetContent

`func (o *DocumentUploadRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DocumentUploadRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DocumentUploadRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetMimeType

`func (o *DocumentUploadRequest) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *DocumentUploadRequest) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *DocumentUploadRequest) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


