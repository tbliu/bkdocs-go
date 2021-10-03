# BankData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**BankCode** | **string** |  | 
**BankCodeType** | **string** |  | 
**Country** | Pointer to **string** |  | [optional] 
**StateProvince** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**StreetAddress** | Pointer to **string** |  | [optional] 
**AccountNumber** | **string** |  | 

## Methods

### NewBankData

`func NewBankData(name string, bankCode string, bankCodeType string, accountNumber string, ) *BankData`

NewBankData instantiates a new BankData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankDataWithDefaults

`func NewBankDataWithDefaults() *BankData`

NewBankDataWithDefaults instantiates a new BankData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BankData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankData) SetName(v string)`

SetName sets Name field to given value.


### GetBankCode

`func (o *BankData) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *BankData) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *BankData) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.


### GetBankCodeType

`func (o *BankData) GetBankCodeType() string`

GetBankCodeType returns the BankCodeType field if non-nil, zero value otherwise.

### GetBankCodeTypeOk

`func (o *BankData) GetBankCodeTypeOk() (*string, bool)`

GetBankCodeTypeOk returns a tuple with the BankCodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCodeType

`func (o *BankData) SetBankCodeType(v string)`

SetBankCodeType sets BankCodeType field to given value.


### GetCountry

`func (o *BankData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BankData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BankData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BankData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetStateProvince

`func (o *BankData) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *BankData) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *BankData) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *BankData) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.

### GetPostalCode

`func (o *BankData) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BankData) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BankData) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *BankData) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCity

`func (o *BankData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BankData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BankData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BankData) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStreetAddress

`func (o *BankData) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *BankData) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *BankData) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *BankData) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetAccountNumber

`func (o *BankData) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BankData) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BankData) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


