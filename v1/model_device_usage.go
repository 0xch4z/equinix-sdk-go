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

// DeviceUsage struct for DeviceUsage
type DeviceUsage struct {
	Quantity *string `json:"quantity,omitempty"`
	Unit *string `json:"unit,omitempty"`
	Total *string `json:"total,omitempty"`
}

// NewDeviceUsage instantiates a new DeviceUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceUsage() *DeviceUsage {
	this := DeviceUsage{}
	return &this
}

// NewDeviceUsageWithDefaults instantiates a new DeviceUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceUsageWithDefaults() *DeviceUsage {
	this := DeviceUsage{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *DeviceUsage) GetQuantity() string {
	if o == nil || o.Quantity == nil {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUsage) GetQuantityOk() (*string, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *DeviceUsage) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *DeviceUsage) SetQuantity(v string) {
	o.Quantity = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *DeviceUsage) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUsage) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *DeviceUsage) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *DeviceUsage) SetUnit(v string) {
	o.Unit = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DeviceUsage) GetTotal() string {
	if o == nil || o.Total == nil {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUsage) GetTotalOk() (*string, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DeviceUsage) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *DeviceUsage) SetTotal(v string) {
	o.Total = &v
}

func (o DeviceUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceUsage struct {
	value *DeviceUsage
	isSet bool
}

func (v NullableDeviceUsage) Get() *DeviceUsage {
	return v.value
}

func (v *NullableDeviceUsage) Set(val *DeviceUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceUsage(val *DeviceUsage) *NullableDeviceUsage {
	return &NullableDeviceUsage{value: val, isSet: true}
}

func (v NullableDeviceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


