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

// CreateProjectAPIKeyRequest struct for CreateProjectAPIKeyRequest
type CreateProjectAPIKeyRequest struct {
	Description *string `json:"description,omitempty"`
	ReadOnly    *bool   `json:"read_only,omitempty"`
}

// NewCreateProjectAPIKeyRequest instantiates a new CreateProjectAPIKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectAPIKeyRequest() *CreateProjectAPIKeyRequest {
	this := CreateProjectAPIKeyRequest{}
	return &this
}

// NewCreateProjectAPIKeyRequestWithDefaults instantiates a new CreateProjectAPIKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectAPIKeyRequestWithDefaults() *CreateProjectAPIKeyRequest {
	this := CreateProjectAPIKeyRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateProjectAPIKeyRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectAPIKeyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateProjectAPIKeyRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateProjectAPIKeyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *CreateProjectAPIKeyRequest) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectAPIKeyRequest) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *CreateProjectAPIKeyRequest) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *CreateProjectAPIKeyRequest) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o CreateProjectAPIKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ReadOnly != nil {
		toSerialize["read_only"] = o.ReadOnly
	}
	return json.Marshal(toSerialize)
}

type NullableCreateProjectAPIKeyRequest struct {
	value *CreateProjectAPIKeyRequest
	isSet bool
}

func (v NullableCreateProjectAPIKeyRequest) Get() *CreateProjectAPIKeyRequest {
	return v.value
}

func (v *NullableCreateProjectAPIKeyRequest) Set(val *CreateProjectAPIKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectAPIKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectAPIKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectAPIKeyRequest(val *CreateProjectAPIKeyRequest) *NullableCreateProjectAPIKeyRequest {
	return &NullableCreateProjectAPIKeyRequest{value: val, isSet: true}
}

func (v NullableCreateProjectAPIKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectAPIKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
