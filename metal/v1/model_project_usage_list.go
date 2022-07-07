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

// ProjectUsageList struct for ProjectUsageList
type ProjectUsageList struct {
	Usages []FindProjectUsage200ResponseUsagesInner `json:"usages,omitempty"`
}

// NewProjectUsageList instantiates a new ProjectUsageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectUsageList() *ProjectUsageList {
	this := ProjectUsageList{}
	return &this
}

// NewProjectUsageListWithDefaults instantiates a new ProjectUsageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectUsageListWithDefaults() *ProjectUsageList {
	this := ProjectUsageList{}
	return &this
}

// GetUsages returns the Usages field value if set, zero value otherwise.
func (o *ProjectUsageList) GetUsages() []FindProjectUsage200ResponseUsagesInner {
	if o == nil || o.Usages == nil {
		var ret []FindProjectUsage200ResponseUsagesInner
		return ret
	}
	return o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsageList) GetUsagesOk() ([]FindProjectUsage200ResponseUsagesInner, bool) {
	if o == nil || o.Usages == nil {
		return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *ProjectUsageList) HasUsages() bool {
	if o != nil && o.Usages != nil {
		return true
	}

	return false
}

// SetUsages gets a reference to the given []FindProjectUsage200ResponseUsagesInner and assigns it to the Usages field.
func (o *ProjectUsageList) SetUsages(v []FindProjectUsage200ResponseUsagesInner) {
	o.Usages = v
}

func (o ProjectUsageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usages != nil {
		toSerialize["usages"] = o.Usages
	}
	return json.Marshal(toSerialize)
}

type NullableProjectUsageList struct {
	value *ProjectUsageList
	isSet bool
}

func (v NullableProjectUsageList) Get() *ProjectUsageList {
	return v.value
}

func (v *NullableProjectUsageList) Set(val *ProjectUsageList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectUsageList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectUsageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectUsageList(val *ProjectUsageList) *NullableProjectUsageList {
	return &NullableProjectUsageList{value: val, isSet: true}
}

func (v NullableProjectUsageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectUsageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
