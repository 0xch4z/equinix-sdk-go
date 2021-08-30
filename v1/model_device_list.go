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

// DeviceList struct for DeviceList
type DeviceList struct {
	Devices *[]Device `json:"devices,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
}

// NewDeviceList instantiates a new DeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceList() *DeviceList {
	this := DeviceList{}
	return &this
}

// NewDeviceListWithDefaults instantiates a new DeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceListWithDefaults() *DeviceList {
	this := DeviceList{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *DeviceList) GetDevices() []Device {
	if o == nil || o.Devices == nil {
		var ret []Device
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetDevicesOk() (*[]Device, bool) {
	if o == nil || o.Devices == nil {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *DeviceList) HasDevices() bool {
	if o != nil && o.Devices != nil {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []Device and assigns it to the Devices field.
func (o *DeviceList) SetDevices(v []Device) {
	o.Devices = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DeviceList) GetMeta() Meta {
	if o == nil || o.Meta == nil {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetMetaOk() (*Meta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DeviceList) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *DeviceList) SetMeta(v Meta) {
	o.Meta = &v
}

func (o DeviceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Devices != nil {
		toSerialize["devices"] = o.Devices
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceList struct {
	value *DeviceList
	isSet bool
}

func (v NullableDeviceList) Get() *DeviceList {
	return v.value
}

func (v *NullableDeviceList) Set(val *DeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceList(val *DeviceList) *NullableDeviceList {
	return &NullableDeviceList{value: val, isSet: true}
}

func (v NullableDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


