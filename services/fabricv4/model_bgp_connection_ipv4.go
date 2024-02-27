/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

API version: 4.12
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the BGPConnectionIpv4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BGPConnectionIpv4{}

// BGPConnectionIpv4 struct for BGPConnectionIpv4
type BGPConnectionIpv4 struct {
	// Customer side peering ip
	CustomerPeerIp string `json:"customerPeerIp"`
	// Equinix side peering ip
	EquinixPeerIp *string `json:"equinixPeerIp,omitempty"`
	// Admin status for the BGP session
	Enabled              bool `json:"enabled"`
	AdditionalProperties map[string]interface{}
}

type _BGPConnectionIpv4 BGPConnectionIpv4

// NewBGPConnectionIpv4 instantiates a new BGPConnectionIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBGPConnectionIpv4(customerPeerIp string, enabled bool) *BGPConnectionIpv4 {
	this := BGPConnectionIpv4{}
	this.CustomerPeerIp = customerPeerIp
	this.Enabled = enabled
	return &this
}

// NewBGPConnectionIpv4WithDefaults instantiates a new BGPConnectionIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBGPConnectionIpv4WithDefaults() *BGPConnectionIpv4 {
	this := BGPConnectionIpv4{}
	return &this
}

// GetCustomerPeerIp returns the CustomerPeerIp field value
func (o *BGPConnectionIpv4) GetCustomerPeerIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerPeerIp
}

// GetCustomerPeerIpOk returns a tuple with the CustomerPeerIp field value
// and a boolean to check if the value has been set.
func (o *BGPConnectionIpv4) GetCustomerPeerIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerPeerIp, true
}

// SetCustomerPeerIp sets field value
func (o *BGPConnectionIpv4) SetCustomerPeerIp(v string) {
	o.CustomerPeerIp = v
}

// GetEquinixPeerIp returns the EquinixPeerIp field value if set, zero value otherwise.
func (o *BGPConnectionIpv4) GetEquinixPeerIp() string {
	if o == nil || IsNil(o.EquinixPeerIp) {
		var ret string
		return ret
	}
	return *o.EquinixPeerIp
}

// GetEquinixPeerIpOk returns a tuple with the EquinixPeerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPConnectionIpv4) GetEquinixPeerIpOk() (*string, bool) {
	if o == nil || IsNil(o.EquinixPeerIp) {
		return nil, false
	}
	return o.EquinixPeerIp, true
}

// HasEquinixPeerIp returns a boolean if a field has been set.
func (o *BGPConnectionIpv4) HasEquinixPeerIp() bool {
	if o != nil && !IsNil(o.EquinixPeerIp) {
		return true
	}

	return false
}

// SetEquinixPeerIp gets a reference to the given string and assigns it to the EquinixPeerIp field.
func (o *BGPConnectionIpv4) SetEquinixPeerIp(v string) {
	o.EquinixPeerIp = &v
}

// GetEnabled returns the Enabled field value
func (o *BGPConnectionIpv4) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *BGPConnectionIpv4) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *BGPConnectionIpv4) SetEnabled(v bool) {
	o.Enabled = v
}

func (o BGPConnectionIpv4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BGPConnectionIpv4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerPeerIp"] = o.CustomerPeerIp
	if !IsNil(o.EquinixPeerIp) {
		toSerialize["equinixPeerIp"] = o.EquinixPeerIp
	}
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BGPConnectionIpv4) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerPeerIp",
		"enabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBGPConnectionIpv4 := _BGPConnectionIpv4{}

	err = json.Unmarshal(data, &varBGPConnectionIpv4)

	if err != nil {
		return err
	}

	*o = BGPConnectionIpv4(varBGPConnectionIpv4)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customerPeerIp")
		delete(additionalProperties, "equinixPeerIp")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBGPConnectionIpv4 struct {
	value *BGPConnectionIpv4
	isSet bool
}

func (v NullableBGPConnectionIpv4) Get() *BGPConnectionIpv4 {
	return v.value
}

func (v *NullableBGPConnectionIpv4) Set(val *BGPConnectionIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableBGPConnectionIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableBGPConnectionIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBGPConnectionIpv4(val *BGPConnectionIpv4) *NullableBGPConnectionIpv4 {
	return &NullableBGPConnectionIpv4{value: val, isSet: true}
}

func (v NullableBGPConnectionIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBGPConnectionIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
