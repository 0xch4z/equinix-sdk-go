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

// UpdateLicenseRequest struct for UpdateLicenseRequest
type UpdateLicenseRequest struct {
	Description *string  `json:"description,omitempty"`
	Size        *float32 `json:"size,omitempty"`
}

// NewUpdateLicenseRequest instantiates a new UpdateLicenseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLicenseRequest() *UpdateLicenseRequest {
	this := UpdateLicenseRequest{}
	return &this
}

// NewUpdateLicenseRequestWithDefaults instantiates a new UpdateLicenseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLicenseRequestWithDefaults() *UpdateLicenseRequest {
	this := UpdateLicenseRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateLicenseRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLicenseRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateLicenseRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateLicenseRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *UpdateLicenseRequest) GetSize() float32 {
	if o == nil || o.Size == nil {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLicenseRequest) GetSizeOk() (*float32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *UpdateLicenseRequest) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *UpdateLicenseRequest) SetSize(v float32) {
	o.Size = &v
}

func (o UpdateLicenseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateLicenseRequest struct {
	value *UpdateLicenseRequest
	isSet bool
}

func (v NullableUpdateLicenseRequest) Get() *UpdateLicenseRequest {
	return v.value
}

func (v *NullableUpdateLicenseRequest) Set(val *UpdateLicenseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLicenseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLicenseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLicenseRequest(val *UpdateLicenseRequest) *NullableUpdateLicenseRequest {
	return &NullableUpdateLicenseRequest{value: val, isSet: true}
}

func (v NullableUpdateLicenseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLicenseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
