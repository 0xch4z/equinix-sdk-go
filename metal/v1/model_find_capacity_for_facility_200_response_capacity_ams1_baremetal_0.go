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

// FindCapacityForFacility200ResponseCapacityAms1Baremetal0 struct for FindCapacityForFacility200ResponseCapacityAms1Baremetal0
type FindCapacityForFacility200ResponseCapacityAms1Baremetal0 struct {
	Level *string `json:"level,omitempty"`
}

// NewFindCapacityForFacility200ResponseCapacityAms1Baremetal0 instantiates a new FindCapacityForFacility200ResponseCapacityAms1Baremetal0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindCapacityForFacility200ResponseCapacityAms1Baremetal0() *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	this := FindCapacityForFacility200ResponseCapacityAms1Baremetal0{}
	return &this
}

// NewFindCapacityForFacility200ResponseCapacityAms1Baremetal0WithDefaults instantiates a new FindCapacityForFacility200ResponseCapacityAms1Baremetal0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindCapacityForFacility200ResponseCapacityAms1Baremetal0WithDefaults() *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	this := FindCapacityForFacility200ResponseCapacityAms1Baremetal0{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *FindCapacityForFacility200ResponseCapacityAms1Baremetal0) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForFacility200ResponseCapacityAms1Baremetal0) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *FindCapacityForFacility200ResponseCapacityAms1Baremetal0) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *FindCapacityForFacility200ResponseCapacityAms1Baremetal0) SetLevel(v string) {
	o.Level = &v
}

func (o FindCapacityForFacility200ResponseCapacityAms1Baremetal0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableFindCapacityForFacility200ResponseCapacityAms1Baremetal0 struct {
	value *FindCapacityForFacility200ResponseCapacityAms1Baremetal0
	isSet bool
}

func (v NullableFindCapacityForFacility200ResponseCapacityAms1Baremetal0) Get() *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	return v.value
}

func (v *NullableFindCapacityForFacility200ResponseCapacityAms1Baremetal0) Set(val *FindCapacityForFacility200ResponseCapacityAms1Baremetal0) {
	v.value = val
	v.isSet = true
}

func (v NullableFindCapacityForFacility200ResponseCapacityAms1Baremetal0) IsSet() bool {
	return v.isSet
}

func (v *NullableFindCapacityForFacility200ResponseCapacityAms1Baremetal0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindCapacityForFacility200ResponseCapacityAms1Baremetal0(val *FindCapacityForFacility200ResponseCapacityAms1Baremetal0) *NullableFindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	return &NullableFindCapacityForFacility200ResponseCapacityAms1Baremetal0{value: val, isSet: true}
}

func (v NullableFindCapacityForFacility200ResponseCapacityAms1Baremetal0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindCapacityForFacility200ResponseCapacityAms1Baremetal0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
