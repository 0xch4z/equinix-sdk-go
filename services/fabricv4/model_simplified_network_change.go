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

// checks if the SimplifiedNetworkChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedNetworkChange{}

// SimplifiedNetworkChange Current state of latest network change
type SimplifiedNetworkChange struct {
	// Network URI
	Href *string `json:"href,omitempty"`
	// Uniquely identifies a change
	Uuid                 *string            `json:"uuid,omitempty"`
	Type                 *NetworkChangeType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimplifiedNetworkChange SimplifiedNetworkChange

// NewSimplifiedNetworkChange instantiates a new SimplifiedNetworkChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedNetworkChange() *SimplifiedNetworkChange {
	this := SimplifiedNetworkChange{}
	return &this
}

// NewSimplifiedNetworkChangeWithDefaults instantiates a new SimplifiedNetworkChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedNetworkChangeWithDefaults() *SimplifiedNetworkChange {
	this := SimplifiedNetworkChange{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SimplifiedNetworkChange) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetworkChange) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SimplifiedNetworkChange) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SimplifiedNetworkChange) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SimplifiedNetworkChange) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetworkChange) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SimplifiedNetworkChange) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SimplifiedNetworkChange) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimplifiedNetworkChange) GetType() NetworkChangeType {
	if o == nil || IsNil(o.Type) {
		var ret NetworkChangeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetworkChange) GetTypeOk() (*NetworkChangeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimplifiedNetworkChange) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given NetworkChangeType and assigns it to the Type field.
func (o *SimplifiedNetworkChange) SetType(v NetworkChangeType) {
	o.Type = &v
}

func (o SimplifiedNetworkChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedNetworkChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimplifiedNetworkChange) UnmarshalJSON(data []byte) (err error) {
	varSimplifiedNetworkChange := _SimplifiedNetworkChange{}

	err = json.Unmarshal(data, &varSimplifiedNetworkChange)

	if err != nil {
		return err
	}

	*o = SimplifiedNetworkChange(varSimplifiedNetworkChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimplifiedNetworkChange struct {
	value *SimplifiedNetworkChange
	isSet bool
}

func (v NullableSimplifiedNetworkChange) Get() *SimplifiedNetworkChange {
	return v.value
}

func (v *NullableSimplifiedNetworkChange) Set(val *SimplifiedNetworkChange) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedNetworkChange) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedNetworkChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedNetworkChange(val *SimplifiedNetworkChange) *NullableSimplifiedNetworkChange {
	return &NullableSimplifiedNetworkChange{value: val, isSet: true}
}

func (v NullableSimplifiedNetworkChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedNetworkChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
