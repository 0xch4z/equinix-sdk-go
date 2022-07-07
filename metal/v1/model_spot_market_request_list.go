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

// SpotMarketRequestList struct for SpotMarketRequestList
type SpotMarketRequestList struct {
	SpotMarketRequests []ListSpotMarketRequests200ResponseSpotMarketRequestsInner `json:"spot_market_requests,omitempty"`
}

// NewSpotMarketRequestList instantiates a new SpotMarketRequestList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotMarketRequestList() *SpotMarketRequestList {
	this := SpotMarketRequestList{}
	return &this
}

// NewSpotMarketRequestListWithDefaults instantiates a new SpotMarketRequestList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotMarketRequestListWithDefaults() *SpotMarketRequestList {
	this := SpotMarketRequestList{}
	return &this
}

// GetSpotMarketRequests returns the SpotMarketRequests field value if set, zero value otherwise.
func (o *SpotMarketRequestList) GetSpotMarketRequests() []ListSpotMarketRequests200ResponseSpotMarketRequestsInner {
	if o == nil || o.SpotMarketRequests == nil {
		var ret []ListSpotMarketRequests200ResponseSpotMarketRequestsInner
		return ret
	}
	return o.SpotMarketRequests
}

// GetSpotMarketRequestsOk returns a tuple with the SpotMarketRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestList) GetSpotMarketRequestsOk() ([]ListSpotMarketRequests200ResponseSpotMarketRequestsInner, bool) {
	if o == nil || o.SpotMarketRequests == nil {
		return nil, false
	}
	return o.SpotMarketRequests, true
}

// HasSpotMarketRequests returns a boolean if a field has been set.
func (o *SpotMarketRequestList) HasSpotMarketRequests() bool {
	if o != nil && o.SpotMarketRequests != nil {
		return true
	}

	return false
}

// SetSpotMarketRequests gets a reference to the given []ListSpotMarketRequests200ResponseSpotMarketRequestsInner and assigns it to the SpotMarketRequests field.
func (o *SpotMarketRequestList) SetSpotMarketRequests(v []ListSpotMarketRequests200ResponseSpotMarketRequestsInner) {
	o.SpotMarketRequests = v
}

func (o SpotMarketRequestList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SpotMarketRequests != nil {
		toSerialize["spot_market_requests"] = o.SpotMarketRequests
	}
	return json.Marshal(toSerialize)
}

type NullableSpotMarketRequestList struct {
	value *SpotMarketRequestList
	isSet bool
}

func (v NullableSpotMarketRequestList) Get() *SpotMarketRequestList {
	return v.value
}

func (v *NullableSpotMarketRequestList) Set(val *SpotMarketRequestList) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotMarketRequestList) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotMarketRequestList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotMarketRequestList(val *SpotMarketRequestList) *NullableSpotMarketRequestList {
	return &NullableSpotMarketRequestList{value: val, isSet: true}
}

func (v NullableSpotMarketRequestList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotMarketRequestList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
