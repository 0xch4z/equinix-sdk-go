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

// FindSpotMarketPricesHistory200ResponsePricesHistory struct for FindSpotMarketPricesHistory200ResponsePricesHistory
type FindSpotMarketPricesHistory200ResponsePricesHistory struct {
	Datapoints [][]float32 `json:"datapoints,omitempty"`
}

// NewFindSpotMarketPricesHistory200ResponsePricesHistory instantiates a new FindSpotMarketPricesHistory200ResponsePricesHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindSpotMarketPricesHistory200ResponsePricesHistory() *FindSpotMarketPricesHistory200ResponsePricesHistory {
	this := FindSpotMarketPricesHistory200ResponsePricesHistory{}
	return &this
}

// NewFindSpotMarketPricesHistory200ResponsePricesHistoryWithDefaults instantiates a new FindSpotMarketPricesHistory200ResponsePricesHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindSpotMarketPricesHistory200ResponsePricesHistoryWithDefaults() *FindSpotMarketPricesHistory200ResponsePricesHistory {
	this := FindSpotMarketPricesHistory200ResponsePricesHistory{}
	return &this
}

// GetDatapoints returns the Datapoints field value if set, zero value otherwise.
func (o *FindSpotMarketPricesHistory200ResponsePricesHistory) GetDatapoints() [][]float32 {
	if o == nil || o.Datapoints == nil {
		var ret [][]float32
		return ret
	}
	return o.Datapoints
}

// GetDatapointsOk returns a tuple with the Datapoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPricesHistory200ResponsePricesHistory) GetDatapointsOk() ([][]float32, bool) {
	if o == nil || o.Datapoints == nil {
		return nil, false
	}
	return o.Datapoints, true
}

// HasDatapoints returns a boolean if a field has been set.
func (o *FindSpotMarketPricesHistory200ResponsePricesHistory) HasDatapoints() bool {
	if o != nil && o.Datapoints != nil {
		return true
	}

	return false
}

// SetDatapoints gets a reference to the given [][]float32 and assigns it to the Datapoints field.
func (o *FindSpotMarketPricesHistory200ResponsePricesHistory) SetDatapoints(v [][]float32) {
	o.Datapoints = v
}

func (o FindSpotMarketPricesHistory200ResponsePricesHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datapoints != nil {
		toSerialize["datapoints"] = o.Datapoints
	}
	return json.Marshal(toSerialize)
}

type NullableFindSpotMarketPricesHistory200ResponsePricesHistory struct {
	value *FindSpotMarketPricesHistory200ResponsePricesHistory
	isSet bool
}

func (v NullableFindSpotMarketPricesHistory200ResponsePricesHistory) Get() *FindSpotMarketPricesHistory200ResponsePricesHistory {
	return v.value
}

func (v *NullableFindSpotMarketPricesHistory200ResponsePricesHistory) Set(val *FindSpotMarketPricesHistory200ResponsePricesHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableFindSpotMarketPricesHistory200ResponsePricesHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableFindSpotMarketPricesHistory200ResponsePricesHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindSpotMarketPricesHistory200ResponsePricesHistory(val *FindSpotMarketPricesHistory200ResponsePricesHistory) *NullableFindSpotMarketPricesHistory200ResponsePricesHistory {
	return &NullableFindSpotMarketPricesHistory200ResponsePricesHistory{value: val, isSet: true}
}

func (v NullableFindSpotMarketPricesHistory200ResponsePricesHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindSpotMarketPricesHistory200ResponsePricesHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
