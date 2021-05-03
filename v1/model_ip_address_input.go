/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// IpAddressInput struct for IpAddressInput
type IpAddressInput struct {
	// Address Family for IP Address
	AddressFamily *float32 `json:"address_family,omitempty"`
}

// NewIpAddressInput instantiates a new IpAddressInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpAddressInput() *IpAddressInput {
	this := IpAddressInput{}
	return &this
}

// NewIpAddressInputWithDefaults instantiates a new IpAddressInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpAddressInputWithDefaults() *IpAddressInput {
	this := IpAddressInput{}
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *IpAddressInput) GetAddressFamily() float32 {
	if o == nil || o.AddressFamily == nil {
		var ret float32
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAddressInput) GetAddressFamilyOk() (*float32, bool) {
	if o == nil || o.AddressFamily == nil {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *IpAddressInput) HasAddressFamily() bool {
	if o != nil && o.AddressFamily != nil {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given float32 and assigns it to the AddressFamily field.
func (o *IpAddressInput) SetAddressFamily(v float32) {
	o.AddressFamily = &v
}

func (o IpAddressInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressFamily != nil {
		toSerialize["address_family"] = o.AddressFamily
	}
	return json.Marshal(toSerialize)
}

type NullableIpAddressInput struct {
	value *IpAddressInput
	isSet bool
}

func (v NullableIpAddressInput) Get() *IpAddressInput {
	return v.value
}

func (v *NullableIpAddressInput) Set(val *IpAddressInput) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAddressInput) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAddressInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAddressInput(val *IpAddressInput) *NullableIpAddressInput {
	return &NullableIpAddressInput{value: val, isSet: true}
}

func (v NullableIpAddressInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAddressInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


