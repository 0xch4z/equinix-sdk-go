/*
Metal API

This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// PaymentMethodBillingAddress struct for PaymentMethodBillingAddress
type PaymentMethodBillingAddress struct {
	StreetAddress *string `json:"street_address,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	CountryCodeAlpha2 *string `json:"country_code_alpha2,omitempty"`
}

// NewPaymentMethodBillingAddress instantiates a new PaymentMethodBillingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodBillingAddress() *PaymentMethodBillingAddress {
	this := PaymentMethodBillingAddress{}
	return &this
}

// NewPaymentMethodBillingAddressWithDefaults instantiates a new PaymentMethodBillingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodBillingAddressWithDefaults() *PaymentMethodBillingAddress {
	this := PaymentMethodBillingAddress{}
	return &this
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *PaymentMethodBillingAddress) GetStreetAddress() string {
	if o == nil || o.StreetAddress == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBillingAddress) GetStreetAddressOk() (*string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *PaymentMethodBillingAddress) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *PaymentMethodBillingAddress) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *PaymentMethodBillingAddress) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBillingAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *PaymentMethodBillingAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *PaymentMethodBillingAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCodeAlpha2 returns the CountryCodeAlpha2 field value if set, zero value otherwise.
func (o *PaymentMethodBillingAddress) GetCountryCodeAlpha2() string {
	if o == nil || o.CountryCodeAlpha2 == nil {
		var ret string
		return ret
	}
	return *o.CountryCodeAlpha2
}

// GetCountryCodeAlpha2Ok returns a tuple with the CountryCodeAlpha2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBillingAddress) GetCountryCodeAlpha2Ok() (*string, bool) {
	if o == nil || o.CountryCodeAlpha2 == nil {
		return nil, false
	}
	return o.CountryCodeAlpha2, true
}

// HasCountryCodeAlpha2 returns a boolean if a field has been set.
func (o *PaymentMethodBillingAddress) HasCountryCodeAlpha2() bool {
	if o != nil && o.CountryCodeAlpha2 != nil {
		return true
	}

	return false
}

// SetCountryCodeAlpha2 gets a reference to the given string and assigns it to the CountryCodeAlpha2 field.
func (o *PaymentMethodBillingAddress) SetCountryCodeAlpha2(v string) {
	o.CountryCodeAlpha2 = &v
}

func (o PaymentMethodBillingAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StreetAddress != nil {
		toSerialize["street_address"] = o.StreetAddress
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.CountryCodeAlpha2 != nil {
		toSerialize["country_code_alpha2"] = o.CountryCodeAlpha2
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodBillingAddress struct {
	value *PaymentMethodBillingAddress
	isSet bool
}

func (v NullablePaymentMethodBillingAddress) Get() *PaymentMethodBillingAddress {
	return v.value
}

func (v *NullablePaymentMethodBillingAddress) Set(val *PaymentMethodBillingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodBillingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodBillingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodBillingAddress(val *PaymentMethodBillingAddress) *NullablePaymentMethodBillingAddress {
	return &NullablePaymentMethodBillingAddress{value: val, isSet: true}
}

func (v NullablePaymentMethodBillingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodBillingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


