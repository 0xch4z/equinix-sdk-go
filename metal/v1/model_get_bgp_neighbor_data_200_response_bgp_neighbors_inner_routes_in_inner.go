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

// GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner struct for GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner
type GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner struct {
	Exact *bool `json:"exact,omitempty"`
	// A project network
	Route *string `json:"route,omitempty"`
}

// NewGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner instantiates a new GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner() *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner {
	this := GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner{}
	return &this
}

// NewGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInnerWithDefaults instantiates a new GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInnerWithDefaults() *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner {
	this := GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner{}
	return &this
}

// GetExact returns the Exact field value if set, zero value otherwise.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) GetExact() bool {
	if o == nil || o.Exact == nil {
		var ret bool
		return ret
	}
	return *o.Exact
}

// GetExactOk returns a tuple with the Exact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) GetExactOk() (*bool, bool) {
	if o == nil || o.Exact == nil {
		return nil, false
	}
	return o.Exact, true
}

// HasExact returns a boolean if a field has been set.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) HasExact() bool {
	if o != nil && o.Exact != nil {
		return true
	}

	return false
}

// SetExact gets a reference to the given bool and assigns it to the Exact field.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) SetExact(v bool) {
	o.Exact = &v
}

// GetRoute returns the Route field value if set, zero value otherwise.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) GetRoute() string {
	if o == nil || o.Route == nil {
		var ret string
		return ret
	}
	return *o.Route
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) GetRouteOk() (*string, bool) {
	if o == nil || o.Route == nil {
		return nil, false
	}
	return o.Route, true
}

// HasRoute returns a boolean if a field has been set.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) HasRoute() bool {
	if o != nil && o.Route != nil {
		return true
	}

	return false
}

// SetRoute gets a reference to the given string and assigns it to the Route field.
func (o *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) SetRoute(v string) {
	o.Route = &v
}

func (o GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exact != nil {
		toSerialize["exact"] = o.Exact
	}
	if o.Route != nil {
		toSerialize["route"] = o.Route
	}
	return json.Marshal(toSerialize)
}

type NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner struct {
	value *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner
	isSet bool
}

func (v NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) Get() *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner {
	return v.value
}

func (v *NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) Set(val *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner(val *GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) *NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner {
	return &NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner{value: val, isSet: true}
}

func (v NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
