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

// FindDeviceById200ResponseFacilityAddressCoordinates struct for FindDeviceById200ResponseFacilityAddressCoordinates
type FindDeviceById200ResponseFacilityAddressCoordinates struct {
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
}

// NewFindDeviceById200ResponseFacilityAddressCoordinates instantiates a new FindDeviceById200ResponseFacilityAddressCoordinates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceById200ResponseFacilityAddressCoordinates() *FindDeviceById200ResponseFacilityAddressCoordinates {
	this := FindDeviceById200ResponseFacilityAddressCoordinates{}
	return &this
}

// NewFindDeviceById200ResponseFacilityAddressCoordinatesWithDefaults instantiates a new FindDeviceById200ResponseFacilityAddressCoordinates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceById200ResponseFacilityAddressCoordinatesWithDefaults() *FindDeviceById200ResponseFacilityAddressCoordinates {
	this := FindDeviceById200ResponseFacilityAddressCoordinates{}
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseFacilityAddressCoordinates) GetLatitude() string {
	if o == nil || o.Latitude == nil {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseFacilityAddressCoordinates) GetLatitudeOk() (*string, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseFacilityAddressCoordinates) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *FindDeviceById200ResponseFacilityAddressCoordinates) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseFacilityAddressCoordinates) GetLongitude() string {
	if o == nil || o.Longitude == nil {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseFacilityAddressCoordinates) GetLongitudeOk() (*string, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseFacilityAddressCoordinates) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *FindDeviceById200ResponseFacilityAddressCoordinates) SetLongitude(v string) {
	o.Longitude = &v
}

func (o FindDeviceById200ResponseFacilityAddressCoordinates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceById200ResponseFacilityAddressCoordinates struct {
	value *FindDeviceById200ResponseFacilityAddressCoordinates
	isSet bool
}

func (v NullableFindDeviceById200ResponseFacilityAddressCoordinates) Get() *FindDeviceById200ResponseFacilityAddressCoordinates {
	return v.value
}

func (v *NullableFindDeviceById200ResponseFacilityAddressCoordinates) Set(val *FindDeviceById200ResponseFacilityAddressCoordinates) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceById200ResponseFacilityAddressCoordinates) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceById200ResponseFacilityAddressCoordinates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceById200ResponseFacilityAddressCoordinates(val *FindDeviceById200ResponseFacilityAddressCoordinates) *NullableFindDeviceById200ResponseFacilityAddressCoordinates {
	return &NullableFindDeviceById200ResponseFacilityAddressCoordinates{value: val, isSet: true}
}

func (v NullableFindDeviceById200ResponseFacilityAddressCoordinates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceById200ResponseFacilityAddressCoordinates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
