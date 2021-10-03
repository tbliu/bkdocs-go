# Agreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agreement** | **string** | - margin_agreement: Alpaca Margin Agreement - account_agreement: Alpaca Account Agreement - customer_agreement: Alpaca Customer Agreement  | 
**SignedAt** | **string** |  | 
**IpAddress** | **string** |  | 

## Methods

### NewAgreement

`func NewAgreement(agreement string, signedAt string, ipAddress string, ) *Agreement`

NewAgreement instantiates a new Agreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementWithDefaults

`func NewAgreementWithDefaults() *Agreement`

NewAgreementWithDefaults instantiates a new Agreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreement

`func (o *Agreement) GetAgreement() string`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *Agreement) GetAgreementOk() (*string, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *Agreement) SetAgreement(v string)`

SetAgreement sets Agreement field to given value.


### GetSignedAt

`func (o *Agreement) GetSignedAt() string`

GetSignedAt returns the SignedAt field if non-nil, zero value otherwise.

### GetSignedAtOk

`func (o *Agreement) GetSignedAtOk() (*string, bool)`

GetSignedAtOk returns a tuple with the SignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedAt

`func (o *Agreement) SetSignedAt(v string)`

SetSignedAt sets SignedAt field to given value.


### GetIpAddress

`func (o *Agreement) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Agreement) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Agreement) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


