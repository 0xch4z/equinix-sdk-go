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

// Port struct for Port
type Port struct {
	Data map[string]interface{} `json:"data,omitempty"`
	// Indicates whether or not the bond can be broken on the port (when applicable).
	DisbondOperationSupported *bool                                  `json:"disbond_operation_supported,omitempty"`
	Href                      *string                                `json:"href,omitempty"`
	Id                        *string                                `json:"id,omitempty"`
	Name                      *string                                `json:"name,omitempty"`
	Type                      *string                                `json:"type,omitempty"`
	VirtualNetworks           []FindBatchById200ResponseDevicesInner `json:"virtual_networks,omitempty"`
}

// NewPort instantiates a new Port object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPort() *Port {
	this := Port{}
	return &this
}

// NewPortWithDefaults instantiates a new Port object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortWithDefaults() *Port {
	this := Port{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Port) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Port) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *Port) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetDisbondOperationSupported returns the DisbondOperationSupported field value if set, zero value otherwise.
func (o *Port) GetDisbondOperationSupported() bool {
	if o == nil || o.DisbondOperationSupported == nil {
		var ret bool
		return ret
	}
	return *o.DisbondOperationSupported
}

// GetDisbondOperationSupportedOk returns a tuple with the DisbondOperationSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetDisbondOperationSupportedOk() (*bool, bool) {
	if o == nil || o.DisbondOperationSupported == nil {
		return nil, false
	}
	return o.DisbondOperationSupported, true
}

// HasDisbondOperationSupported returns a boolean if a field has been set.
func (o *Port) HasDisbondOperationSupported() bool {
	if o != nil && o.DisbondOperationSupported != nil {
		return true
	}

	return false
}

// SetDisbondOperationSupported gets a reference to the given bool and assigns it to the DisbondOperationSupported field.
func (o *Port) SetDisbondOperationSupported(v bool) {
	o.DisbondOperationSupported = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Port) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Port) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Port) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Port) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Port) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Port) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Port) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Port) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Port) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Port) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Port) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Port) SetType(v string) {
	o.Type = &v
}

// GetVirtualNetworks returns the VirtualNetworks field value if set, zero value otherwise.
func (o *Port) GetVirtualNetworks() []FindBatchById200ResponseDevicesInner {
	if o == nil || o.VirtualNetworks == nil {
		var ret []FindBatchById200ResponseDevicesInner
		return ret
	}
	return o.VirtualNetworks
}

// GetVirtualNetworksOk returns a tuple with the VirtualNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetVirtualNetworksOk() ([]FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.VirtualNetworks == nil {
		return nil, false
	}
	return o.VirtualNetworks, true
}

// HasVirtualNetworks returns a boolean if a field has been set.
func (o *Port) HasVirtualNetworks() bool {
	if o != nil && o.VirtualNetworks != nil {
		return true
	}

	return false
}

// SetVirtualNetworks gets a reference to the given []FindBatchById200ResponseDevicesInner and assigns it to the VirtualNetworks field.
func (o *Port) SetVirtualNetworks(v []FindBatchById200ResponseDevicesInner) {
	o.VirtualNetworks = v
}

func (o Port) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.DisbondOperationSupported != nil {
		toSerialize["disbond_operation_supported"] = o.DisbondOperationSupported
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.VirtualNetworks != nil {
		toSerialize["virtual_networks"] = o.VirtualNetworks
	}
	return json.Marshal(toSerialize)
}

type NullablePort struct {
	value *Port
	isSet bool
}

func (v NullablePort) Get() *Port {
	return v.value
}

func (v *NullablePort) Set(val *Port) {
	v.value = val
	v.isSet = true
}

func (v NullablePort) IsSet() bool {
	return v.isSet
}

func (v *NullablePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePort(val *Port) *NullablePort {
	return &NullablePort{value: val, isSet: true}
}

func (v NullablePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
