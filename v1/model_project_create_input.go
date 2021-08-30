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

// ProjectCreateInput struct for ProjectCreateInput
type ProjectCreateInput struct {
	Name string `json:"name"`
	PaymentMethodId *string `json:"payment_method_id,omitempty"`
	Customdata *map[string]interface{} `json:"customdata,omitempty"`
}

// NewProjectCreateInput instantiates a new ProjectCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCreateInput(name string) *ProjectCreateInput {
	this := ProjectCreateInput{}
	this.Name = name
	return &this
}

// NewProjectCreateInputWithDefaults instantiates a new ProjectCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCreateInputWithDefaults() *ProjectCreateInput {
	this := ProjectCreateInput{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectCreateInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectCreateInput) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectCreateInput) SetName(v string) {
	o.Name = v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *ProjectCreateInput) GetPaymentMethodId() string {
	if o == nil || o.PaymentMethodId == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateInput) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || o.PaymentMethodId == nil {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *ProjectCreateInput) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId != nil {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *ProjectCreateInput) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *ProjectCreateInput) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateInput) GetCustomdataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *ProjectCreateInput) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *ProjectCreateInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = &v
}

func (o ProjectCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.PaymentMethodId != nil {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	return json.Marshal(toSerialize)
}

type NullableProjectCreateInput struct {
	value *ProjectCreateInput
	isSet bool
}

func (v NullableProjectCreateInput) Get() *ProjectCreateInput {
	return v.value
}

func (v *NullableProjectCreateInput) Set(val *ProjectCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCreateInput(val *ProjectCreateInput) *NullableProjectCreateInput {
	return &NullableProjectCreateInput{value: val, isSet: true}
}

func (v NullableProjectCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


