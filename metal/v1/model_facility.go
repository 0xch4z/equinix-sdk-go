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

// Facility struct for Facility
type Facility struct {
	Address  *FindDeviceById200ResponseFacilityAddress `json:"address,omitempty"`
	Code     *string                                   `json:"code,omitempty"`
	Features []string                                  `json:"features,omitempty"`
	Id       *string                                   `json:"id,omitempty"`
	// IP ranges registered in facility. Can be used for GeoIP location
	IpRanges []string                                `json:"ip_ranges,omitempty"`
	Metro    *FindDeviceById200ResponseFacilityMetro `json:"metro,omitempty"`
	Name     *string                                 `json:"name,omitempty"`
}

// NewFacility instantiates a new Facility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacility() *Facility {
	this := Facility{}
	return &this
}

// NewFacilityWithDefaults instantiates a new Facility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacilityWithDefaults() *Facility {
	this := Facility{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Facility) GetAddress() FindDeviceById200ResponseFacilityAddress {
	if o == nil || o.Address == nil {
		var ret FindDeviceById200ResponseFacilityAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetAddressOk() (*FindDeviceById200ResponseFacilityAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Facility) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given FindDeviceById200ResponseFacilityAddress and assigns it to the Address field.
func (o *Facility) SetAddress(v FindDeviceById200ResponseFacilityAddress) {
	o.Address = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Facility) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Facility) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Facility) SetCode(v string) {
	o.Code = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Facility) GetFeatures() []string {
	if o == nil || o.Features == nil {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetFeaturesOk() ([]string, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Facility) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *Facility) SetFeatures(v []string) {
	o.Features = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Facility) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Facility) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Facility) SetId(v string) {
	o.Id = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *Facility) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetIpRangesOk() ([]string, bool) {
	if o == nil || o.IpRanges == nil {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *Facility) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *Facility) SetIpRanges(v []string) {
	o.IpRanges = v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *Facility) GetMetro() FindDeviceById200ResponseFacilityMetro {
	if o == nil || o.Metro == nil {
		var ret FindDeviceById200ResponseFacilityMetro
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetMetroOk() (*FindDeviceById200ResponseFacilityMetro, bool) {
	if o == nil || o.Metro == nil {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *Facility) HasMetro() bool {
	if o != nil && o.Metro != nil {
		return true
	}

	return false
}

// SetMetro gets a reference to the given FindDeviceById200ResponseFacilityMetro and assigns it to the Metro field.
func (o *Facility) SetMetro(v FindDeviceById200ResponseFacilityMetro) {
	o.Metro = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Facility) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facility) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Facility) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Facility) SetName(v string) {
	o.Name = &v
}

func (o Facility) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IpRanges != nil {
		toSerialize["ip_ranges"] = o.IpRanges
	}
	if o.Metro != nil {
		toSerialize["metro"] = o.Metro
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableFacility struct {
	value *Facility
	isSet bool
}

func (v NullableFacility) Get() *Facility {
	return v.value
}

func (v *NullableFacility) Set(val *Facility) {
	v.value = val
	v.isSet = true
}

func (v NullableFacility) IsSet() bool {
	return v.isSet
}

func (v *NullableFacility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacility(val *Facility) *NullableFacility {
	return &NullableFacility{value: val, isSet: true}
}

func (v NullableFacility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
