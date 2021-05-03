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

// ParentBlock struct for ParentBlock
type ParentBlock struct {
	Network *string `json:"network,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	Cidr *int32 `json:"cidr,omitempty"`
	Href *string `json:"href,omitempty"`
}

// NewParentBlock instantiates a new ParentBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParentBlock() *ParentBlock {
	this := ParentBlock{}
	return &this
}

// NewParentBlockWithDefaults instantiates a new ParentBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParentBlockWithDefaults() *ParentBlock {
	this := ParentBlock{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ParentBlock) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentBlock) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ParentBlock) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *ParentBlock) SetNetwork(v string) {
	o.Network = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *ParentBlock) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentBlock) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *ParentBlock) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *ParentBlock) SetNetmask(v string) {
	o.Netmask = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *ParentBlock) GetCidr() int32 {
	if o == nil || o.Cidr == nil {
		var ret int32
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentBlock) GetCidrOk() (*int32, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *ParentBlock) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given int32 and assigns it to the Cidr field.
func (o *ParentBlock) SetCidr(v int32) {
	o.Cidr = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ParentBlock) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentBlock) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ParentBlock) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ParentBlock) SetHref(v string) {
	o.Href = &v
}

func (o ParentBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Netmask != nil {
		toSerialize["netmask"] = o.Netmask
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullableParentBlock struct {
	value *ParentBlock
	isSet bool
}

func (v NullableParentBlock) Get() *ParentBlock {
	return v.value
}

func (v *NullableParentBlock) Set(val *ParentBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableParentBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableParentBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentBlock(val *ParentBlock) *NullableParentBlock {
	return &NullableParentBlock{value: val, isSet: true}
}

func (v NullableParentBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


