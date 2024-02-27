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
)

// checks if the ServiceTokenSide type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTokenSide{}

// ServiceTokenSide Connection link protocol configuration
type ServiceTokenSide struct {
	// List of AccessPointSelectors
	AccessPointSelectors []AccessPointSelector `json:"accessPointSelectors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceTokenSide ServiceTokenSide

// NewServiceTokenSide instantiates a new ServiceTokenSide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTokenSide() *ServiceTokenSide {
	this := ServiceTokenSide{}
	return &this
}

// NewServiceTokenSideWithDefaults instantiates a new ServiceTokenSide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTokenSideWithDefaults() *ServiceTokenSide {
	this := ServiceTokenSide{}
	return &this
}

// GetAccessPointSelectors returns the AccessPointSelectors field value if set, zero value otherwise.
func (o *ServiceTokenSide) GetAccessPointSelectors() []AccessPointSelector {
	if o == nil || IsNil(o.AccessPointSelectors) {
		var ret []AccessPointSelector
		return ret
	}
	return o.AccessPointSelectors
}

// GetAccessPointSelectorsOk returns a tuple with the AccessPointSelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTokenSide) GetAccessPointSelectorsOk() ([]AccessPointSelector, bool) {
	if o == nil || IsNil(o.AccessPointSelectors) {
		return nil, false
	}
	return o.AccessPointSelectors, true
}

// HasAccessPointSelectors returns a boolean if a field has been set.
func (o *ServiceTokenSide) HasAccessPointSelectors() bool {
	if o != nil && !IsNil(o.AccessPointSelectors) {
		return true
	}

	return false
}

// SetAccessPointSelectors gets a reference to the given []AccessPointSelector and assigns it to the AccessPointSelectors field.
func (o *ServiceTokenSide) SetAccessPointSelectors(v []AccessPointSelector) {
	o.AccessPointSelectors = v
}

func (o ServiceTokenSide) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTokenSide) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessPointSelectors) {
		toSerialize["accessPointSelectors"] = o.AccessPointSelectors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceTokenSide) UnmarshalJSON(data []byte) (err error) {
	varServiceTokenSide := _ServiceTokenSide{}

	err = json.Unmarshal(data, &varServiceTokenSide)

	if err != nil {
		return err
	}

	*o = ServiceTokenSide(varServiceTokenSide)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessPointSelectors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceTokenSide struct {
	value *ServiceTokenSide
	isSet bool
}

func (v NullableServiceTokenSide) Get() *ServiceTokenSide {
	return v.value
}

func (v *NullableServiceTokenSide) Set(val *ServiceTokenSide) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTokenSide) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTokenSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTokenSide(val *ServiceTokenSide) *NullableServiceTokenSide {
	return &NullableServiceTokenSide{value: val, isSet: true}
}

func (v NullableServiceTokenSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTokenSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
