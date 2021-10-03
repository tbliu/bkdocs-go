# TrustedContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | **string** |  | 
**FamilyName** | **string** |  | 
**EmailAddress** | Pointer to **string** | at least one of &#x60;email_address&#x60;, &#x60;phone_number&#x60; or &#x60;street_address&#x60; is required  | [optional] 
**PhoneNumber** | Pointer to **string** | at least one of &#x60;email_address&#x60;, &#x60;phone_number&#x60; or &#x60;street_address&#x60; is required  | [optional] 
**StreetAddress** | Pointer to **[]string** | at least one of &#x60;email_address&#x60;, &#x60;phone_number&#x60; or &#x60;street_address&#x60; is required  | [optional] 
**City** | Pointer to **string** | required if &#x60;street_address&#x60; is set  | [optional] 
**State** | Pointer to **string** | required if &#x60;street_address&#x60; is set  | [optional] 
**PostalCode** | Pointer to **string** | required if &#x60;street_address&#x60; is set  | [optional] 
**Country** | Pointer to **string** | [ISO 3166-1 alpha-3](https://www.iso.org/iso-3166-country-codes.html). required if &#x60;street_address&#x60; is set  | [optional] 

## Methods

### NewTrustedContact

`func NewTrustedContact(givenName string, familyName string, ) *TrustedContact`

NewTrustedContact instantiates a new TrustedContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedContactWithDefaults

`func NewTrustedContactWithDefaults() *TrustedContact`

NewTrustedContactWithDefaults instantiates a new TrustedContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *TrustedContact) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *TrustedContact) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *TrustedContact) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.


### GetFamilyName

`func (o *TrustedContact) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *TrustedContact) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *TrustedContact) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.


### GetEmailAddress

`func (o *TrustedContact) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *TrustedContact) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *TrustedContact) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *TrustedContact) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *TrustedContact) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *TrustedContact) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *TrustedContact) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *TrustedContact) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStreetAddress

`func (o *TrustedContact) GetStreetAddress() []string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *TrustedContact) GetStreetAddressOk() (*[]string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *TrustedContact) SetStreetAddress(v []string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *TrustedContact) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetCity

`func (o *TrustedContact) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *TrustedContact) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *TrustedContact) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *TrustedContact) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *TrustedContact) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TrustedContact) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TrustedContact) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TrustedContact) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *TrustedContact) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *TrustedContact) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *TrustedContact) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *TrustedContact) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *TrustedContact) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TrustedContact) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TrustedContact) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TrustedContact) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


