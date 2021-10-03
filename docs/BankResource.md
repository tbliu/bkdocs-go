# BankResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Name** | **string** |  | 
**BankCode** | **string** |  | 
**BankCodeType** | **string** |  | 
**Country** | Pointer to **string** |  | [optional] 
**StateProvince** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**StreetAddress** | Pointer to **string** |  | [optional] 
**AccountNumber** | **string** |  | 
**AccountId** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewBankResource

`func NewBankResource(id string, createdAt time.Time, updatedAt time.Time, name string, bankCode string, bankCodeType string, accountNumber string, accountId string, status string, ) *BankResource`

NewBankResource instantiates a new BankResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankResourceWithDefaults

`func NewBankResourceWithDefaults() *BankResource`

NewBankResourceWithDefaults instantiates a new BankResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankResource) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *BankResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BankResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BankResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BankResource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BankResource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BankResource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *BankResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankResource) SetName(v string)`

SetName sets Name field to given value.


### GetBankCode

`func (o *BankResource) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *BankResource) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *BankResource) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.


### GetBankCodeType

`func (o *BankResource) GetBankCodeType() string`

GetBankCodeType returns the BankCodeType field if non-nil, zero value otherwise.

### GetBankCodeTypeOk

`func (o *BankResource) GetBankCodeTypeOk() (*string, bool)`

GetBankCodeTypeOk returns a tuple with the BankCodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCodeType

`func (o *BankResource) SetBankCodeType(v string)`

SetBankCodeType sets BankCodeType field to given value.


### GetCountry

`func (o *BankResource) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BankResource) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BankResource) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BankResource) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetStateProvince

`func (o *BankResource) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *BankResource) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *BankResource) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *BankResource) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.

### GetPostalCode

`func (o *BankResource) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BankResource) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BankResource) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *BankResource) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCity

`func (o *BankResource) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BankResource) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BankResource) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BankResource) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStreetAddress

`func (o *BankResource) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *BankResource) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *BankResource) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *BankResource) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetAccountNumber

`func (o *BankResource) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BankResource) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BankResource) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountId

`func (o *BankResource) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BankResource) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BankResource) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetStatus

`func (o *BankResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BankResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BankResource) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


