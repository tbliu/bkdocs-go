# Identity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | **string** |  | 
**FamilyName** | **string** |  | 
**DateOfBirth** | **string** |  | 
**TaxId** | Pointer to **string** |  | [optional] 
**TaxIdType** | Pointer to **string** |  | [optional] 
**CountryOfCitizenship** | Pointer to **string** | [ISO 3166-1 alpha-3](https://www.iso.org/iso-3166-country-codes.html).  | [optional] 
**CountryOfBirth** | Pointer to **string** | [ISO 3166-1 alpha-3](https://www.iso.org/iso-3166-country-codes.html).  | [optional] 
**CountryOfTaxResidence** | **string** | [ISO 3166-1 alpha-3](https://www.iso.org/iso-3166-country-codes.html).  | 
**FundingSource** | **[]string** |  | 
**AnnualIncomeMin** | Pointer to **float32** |  | [optional] 
**AnnualIncomeMax** | Pointer to **float32** |  | [optional] 
**LiquidNetWorthMin** | Pointer to **float32** |  | [optional] 
**LiquidNetWorthMax** | Pointer to **float32** |  | [optional] 
**TotalNetWorthMin** | Pointer to **float32** |  | [optional] 
**TotalNetWorthMax** | Pointer to **float32** |  | [optional] 
**Extra** | Pointer to **map[string]interface{}** | any extra information used for KYC purposes  | [optional] 

## Methods

### NewIdentity

`func NewIdentity(givenName string, familyName string, dateOfBirth string, countryOfTaxResidence string, fundingSource []string, ) *Identity`

NewIdentity instantiates a new Identity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityWithDefaults

`func NewIdentityWithDefaults() *Identity`

NewIdentityWithDefaults instantiates a new Identity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *Identity) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *Identity) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *Identity) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.


### GetFamilyName

`func (o *Identity) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *Identity) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *Identity) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.


### GetDateOfBirth

`func (o *Identity) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Identity) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Identity) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetTaxId

`func (o *Identity) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *Identity) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *Identity) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *Identity) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxIdType

`func (o *Identity) GetTaxIdType() string`

GetTaxIdType returns the TaxIdType field if non-nil, zero value otherwise.

### GetTaxIdTypeOk

`func (o *Identity) GetTaxIdTypeOk() (*string, bool)`

GetTaxIdTypeOk returns a tuple with the TaxIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdType

`func (o *Identity) SetTaxIdType(v string)`

SetTaxIdType sets TaxIdType field to given value.

### HasTaxIdType

`func (o *Identity) HasTaxIdType() bool`

HasTaxIdType returns a boolean if a field has been set.

### GetCountryOfCitizenship

`func (o *Identity) GetCountryOfCitizenship() string`

GetCountryOfCitizenship returns the CountryOfCitizenship field if non-nil, zero value otherwise.

### GetCountryOfCitizenshipOk

`func (o *Identity) GetCountryOfCitizenshipOk() (*string, bool)`

GetCountryOfCitizenshipOk returns a tuple with the CountryOfCitizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfCitizenship

`func (o *Identity) SetCountryOfCitizenship(v string)`

SetCountryOfCitizenship sets CountryOfCitizenship field to given value.

### HasCountryOfCitizenship

`func (o *Identity) HasCountryOfCitizenship() bool`

HasCountryOfCitizenship returns a boolean if a field has been set.

### GetCountryOfBirth

`func (o *Identity) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *Identity) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *Identity) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.

### HasCountryOfBirth

`func (o *Identity) HasCountryOfBirth() bool`

HasCountryOfBirth returns a boolean if a field has been set.

### GetCountryOfTaxResidence

`func (o *Identity) GetCountryOfTaxResidence() string`

GetCountryOfTaxResidence returns the CountryOfTaxResidence field if non-nil, zero value otherwise.

### GetCountryOfTaxResidenceOk

`func (o *Identity) GetCountryOfTaxResidenceOk() (*string, bool)`

GetCountryOfTaxResidenceOk returns a tuple with the CountryOfTaxResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfTaxResidence

`func (o *Identity) SetCountryOfTaxResidence(v string)`

SetCountryOfTaxResidence sets CountryOfTaxResidence field to given value.


### GetFundingSource

`func (o *Identity) GetFundingSource() []string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *Identity) GetFundingSourceOk() (*[]string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *Identity) SetFundingSource(v []string)`

SetFundingSource sets FundingSource field to given value.


### GetAnnualIncomeMin

`func (o *Identity) GetAnnualIncomeMin() float32`

GetAnnualIncomeMin returns the AnnualIncomeMin field if non-nil, zero value otherwise.

### GetAnnualIncomeMinOk

`func (o *Identity) GetAnnualIncomeMinOk() (*float32, bool)`

GetAnnualIncomeMinOk returns a tuple with the AnnualIncomeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualIncomeMin

`func (o *Identity) SetAnnualIncomeMin(v float32)`

SetAnnualIncomeMin sets AnnualIncomeMin field to given value.

### HasAnnualIncomeMin

`func (o *Identity) HasAnnualIncomeMin() bool`

HasAnnualIncomeMin returns a boolean if a field has been set.

### GetAnnualIncomeMax

`func (o *Identity) GetAnnualIncomeMax() float32`

GetAnnualIncomeMax returns the AnnualIncomeMax field if non-nil, zero value otherwise.

### GetAnnualIncomeMaxOk

`func (o *Identity) GetAnnualIncomeMaxOk() (*float32, bool)`

GetAnnualIncomeMaxOk returns a tuple with the AnnualIncomeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualIncomeMax

`func (o *Identity) SetAnnualIncomeMax(v float32)`

SetAnnualIncomeMax sets AnnualIncomeMax field to given value.

### HasAnnualIncomeMax

`func (o *Identity) HasAnnualIncomeMax() bool`

HasAnnualIncomeMax returns a boolean if a field has been set.

### GetLiquidNetWorthMin

`func (o *Identity) GetLiquidNetWorthMin() float32`

GetLiquidNetWorthMin returns the LiquidNetWorthMin field if non-nil, zero value otherwise.

### GetLiquidNetWorthMinOk

`func (o *Identity) GetLiquidNetWorthMinOk() (*float32, bool)`

GetLiquidNetWorthMinOk returns a tuple with the LiquidNetWorthMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidNetWorthMin

`func (o *Identity) SetLiquidNetWorthMin(v float32)`

SetLiquidNetWorthMin sets LiquidNetWorthMin field to given value.

### HasLiquidNetWorthMin

`func (o *Identity) HasLiquidNetWorthMin() bool`

HasLiquidNetWorthMin returns a boolean if a field has been set.

### GetLiquidNetWorthMax

`func (o *Identity) GetLiquidNetWorthMax() float32`

GetLiquidNetWorthMax returns the LiquidNetWorthMax field if non-nil, zero value otherwise.

### GetLiquidNetWorthMaxOk

`func (o *Identity) GetLiquidNetWorthMaxOk() (*float32, bool)`

GetLiquidNetWorthMaxOk returns a tuple with the LiquidNetWorthMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidNetWorthMax

`func (o *Identity) SetLiquidNetWorthMax(v float32)`

SetLiquidNetWorthMax sets LiquidNetWorthMax field to given value.

### HasLiquidNetWorthMax

`func (o *Identity) HasLiquidNetWorthMax() bool`

HasLiquidNetWorthMax returns a boolean if a field has been set.

### GetTotalNetWorthMin

`func (o *Identity) GetTotalNetWorthMin() float32`

GetTotalNetWorthMin returns the TotalNetWorthMin field if non-nil, zero value otherwise.

### GetTotalNetWorthMinOk

`func (o *Identity) GetTotalNetWorthMinOk() (*float32, bool)`

GetTotalNetWorthMinOk returns a tuple with the TotalNetWorthMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetWorthMin

`func (o *Identity) SetTotalNetWorthMin(v float32)`

SetTotalNetWorthMin sets TotalNetWorthMin field to given value.

### HasTotalNetWorthMin

`func (o *Identity) HasTotalNetWorthMin() bool`

HasTotalNetWorthMin returns a boolean if a field has been set.

### GetTotalNetWorthMax

`func (o *Identity) GetTotalNetWorthMax() float32`

GetTotalNetWorthMax returns the TotalNetWorthMax field if non-nil, zero value otherwise.

### GetTotalNetWorthMaxOk

`func (o *Identity) GetTotalNetWorthMaxOk() (*float32, bool)`

GetTotalNetWorthMaxOk returns a tuple with the TotalNetWorthMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetWorthMax

`func (o *Identity) SetTotalNetWorthMax(v float32)`

SetTotalNetWorthMax sets TotalNetWorthMax field to given value.

### HasTotalNetWorthMax

`func (o *Identity) HasTotalNetWorthMax() bool`

HasTotalNetWorthMax returns a boolean if a field has been set.

### GetExtra

`func (o *Identity) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *Identity) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *Identity) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *Identity) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


