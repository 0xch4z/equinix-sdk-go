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

// SnapshotPolicyInput struct for SnapshotPolicyInput
type SnapshotPolicyInput struct {
	SnapshotCount *int32 `json:"snapshot_count,omitempty"`
	SnapshotFrequency *string `json:"snapshot_frequency,omitempty"`
}

// NewSnapshotPolicyInput instantiates a new SnapshotPolicyInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotPolicyInput() *SnapshotPolicyInput {
	this := SnapshotPolicyInput{}
	return &this
}

// NewSnapshotPolicyInputWithDefaults instantiates a new SnapshotPolicyInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotPolicyInputWithDefaults() *SnapshotPolicyInput {
	this := SnapshotPolicyInput{}
	return &this
}

// GetSnapshotCount returns the SnapshotCount field value if set, zero value otherwise.
func (o *SnapshotPolicyInput) GetSnapshotCount() int32 {
	if o == nil || o.SnapshotCount == nil {
		var ret int32
		return ret
	}
	return *o.SnapshotCount
}

// GetSnapshotCountOk returns a tuple with the SnapshotCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPolicyInput) GetSnapshotCountOk() (*int32, bool) {
	if o == nil || o.SnapshotCount == nil {
		return nil, false
	}
	return o.SnapshotCount, true
}

// HasSnapshotCount returns a boolean if a field has been set.
func (o *SnapshotPolicyInput) HasSnapshotCount() bool {
	if o != nil && o.SnapshotCount != nil {
		return true
	}

	return false
}

// SetSnapshotCount gets a reference to the given int32 and assigns it to the SnapshotCount field.
func (o *SnapshotPolicyInput) SetSnapshotCount(v int32) {
	o.SnapshotCount = &v
}

// GetSnapshotFrequency returns the SnapshotFrequency field value if set, zero value otherwise.
func (o *SnapshotPolicyInput) GetSnapshotFrequency() string {
	if o == nil || o.SnapshotFrequency == nil {
		var ret string
		return ret
	}
	return *o.SnapshotFrequency
}

// GetSnapshotFrequencyOk returns a tuple with the SnapshotFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotPolicyInput) GetSnapshotFrequencyOk() (*string, bool) {
	if o == nil || o.SnapshotFrequency == nil {
		return nil, false
	}
	return o.SnapshotFrequency, true
}

// HasSnapshotFrequency returns a boolean if a field has been set.
func (o *SnapshotPolicyInput) HasSnapshotFrequency() bool {
	if o != nil && o.SnapshotFrequency != nil {
		return true
	}

	return false
}

// SetSnapshotFrequency gets a reference to the given string and assigns it to the SnapshotFrequency field.
func (o *SnapshotPolicyInput) SetSnapshotFrequency(v string) {
	o.SnapshotFrequency = &v
}

func (o SnapshotPolicyInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SnapshotCount != nil {
		toSerialize["snapshot_count"] = o.SnapshotCount
	}
	if o.SnapshotFrequency != nil {
		toSerialize["snapshot_frequency"] = o.SnapshotFrequency
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotPolicyInput struct {
	value *SnapshotPolicyInput
	isSet bool
}

func (v NullableSnapshotPolicyInput) Get() *SnapshotPolicyInput {
	return v.value
}

func (v *NullableSnapshotPolicyInput) Set(val *SnapshotPolicyInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotPolicyInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotPolicyInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotPolicyInput(val *SnapshotPolicyInput) *NullableSnapshotPolicyInput {
	return &NullableSnapshotPolicyInput{value: val, isSet: true}
}

func (v NullableSnapshotPolicyInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotPolicyInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


