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

// UpdateVrfRequest struct for UpdateVrfRequest
type UpdateVrfRequest struct {
	Description *string `json:"description,omitempty"`
	// A list of CIDR network addresses. Like [\"10.0.0.0/16\", \"2001:d78::/56\"]. IPv4 blocks must be between /8 and /29 in size. IPv6 blocks must be between /56 and /64. A VRF\\'s IP ranges must be defined in order to create VRF IP Reservations, which can then be used for Metal Gateways or Virtual Circuits. Adding a new CIDR address to the list will result in the creation of a new IP Range for this VRF. Removal of an existing CIDR address from the list will result in the deletion of an existing IP Range for this VRF. Deleting an IP Range will result in the deletion of any VRF IP Reservations contained within the IP Range, as well as the VRF IP Reservation\\'s associated Metal Gateways or Virtual Circuits. If you do not wish to add or remove IP Ranges, either include the full existing list of IP Ranges in the update request, or do not specify the `ip_ranges` field in the update request. Specifying a value of `[]` will remove all existing IP Ranges from the VRF.
	IpRanges []string `json:"ip_ranges,omitempty"`
	// The new `local_asn` value for the VRF. This field cannot be updated when there are active Interconnection Virtual Circuits associated to the VRF.
	LocalAsn *int32  `json:"local_asn,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// NewUpdateVrfRequest instantiates a new UpdateVrfRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVrfRequest() *UpdateVrfRequest {
	this := UpdateVrfRequest{}
	return &this
}

// NewUpdateVrfRequestWithDefaults instantiates a new UpdateVrfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVrfRequestWithDefaults() *UpdateVrfRequest {
	this := UpdateVrfRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateVrfRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVrfRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateVrfRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateVrfRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *UpdateVrfRequest) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVrfRequest) GetIpRangesOk() ([]string, bool) {
	if o == nil || o.IpRanges == nil {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *UpdateVrfRequest) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *UpdateVrfRequest) SetIpRanges(v []string) {
	o.IpRanges = v
}

// GetLocalAsn returns the LocalAsn field value if set, zero value otherwise.
func (o *UpdateVrfRequest) GetLocalAsn() int32 {
	if o == nil || o.LocalAsn == nil {
		var ret int32
		return ret
	}
	return *o.LocalAsn
}

// GetLocalAsnOk returns a tuple with the LocalAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVrfRequest) GetLocalAsnOk() (*int32, bool) {
	if o == nil || o.LocalAsn == nil {
		return nil, false
	}
	return o.LocalAsn, true
}

// HasLocalAsn returns a boolean if a field has been set.
func (o *UpdateVrfRequest) HasLocalAsn() bool {
	if o != nil && o.LocalAsn != nil {
		return true
	}

	return false
}

// SetLocalAsn gets a reference to the given int32 and assigns it to the LocalAsn field.
func (o *UpdateVrfRequest) SetLocalAsn(v int32) {
	o.LocalAsn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateVrfRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVrfRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVrfRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateVrfRequest) SetName(v string) {
	o.Name = &v
}

func (o UpdateVrfRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IpRanges != nil {
		toSerialize["ip_ranges"] = o.IpRanges
	}
	if o.LocalAsn != nil {
		toSerialize["local_asn"] = o.LocalAsn
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateVrfRequest struct {
	value *UpdateVrfRequest
	isSet bool
}

func (v NullableUpdateVrfRequest) Get() *UpdateVrfRequest {
	return v.value
}

func (v *NullableUpdateVrfRequest) Set(val *UpdateVrfRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVrfRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVrfRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVrfRequest(val *UpdateVrfRequest) *NullableUpdateVrfRequest {
	return &NullableUpdateVrfRequest{value: val, isSet: true}
}

func (v NullableUpdateVrfRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVrfRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
