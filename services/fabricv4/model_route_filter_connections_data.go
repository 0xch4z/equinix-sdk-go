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

// checks if the RouteFilterConnectionsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteFilterConnectionsData{}

// RouteFilterConnectionsData struct for RouteFilterConnectionsData
type RouteFilterConnectionsData struct {
	// Connection URI
	Href *string         `json:"href,omitempty"`
	Type *ConnectionType `json:"type,omitempty"`
	// Route Filter identifier
	Uuid                 *string `json:"uuid,omitempty"`
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteFilterConnectionsData RouteFilterConnectionsData

// NewRouteFilterConnectionsData instantiates a new RouteFilterConnectionsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteFilterConnectionsData() *RouteFilterConnectionsData {
	this := RouteFilterConnectionsData{}
	return &this
}

// NewRouteFilterConnectionsDataWithDefaults instantiates a new RouteFilterConnectionsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteFilterConnectionsDataWithDefaults() *RouteFilterConnectionsData {
	this := RouteFilterConnectionsData{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RouteFilterConnectionsData) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterConnectionsData) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RouteFilterConnectionsData) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RouteFilterConnectionsData) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RouteFilterConnectionsData) GetType() ConnectionType {
	if o == nil || IsNil(o.Type) {
		var ret ConnectionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterConnectionsData) GetTypeOk() (*ConnectionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RouteFilterConnectionsData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ConnectionType and assigns it to the Type field.
func (o *RouteFilterConnectionsData) SetType(v ConnectionType) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RouteFilterConnectionsData) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterConnectionsData) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RouteFilterConnectionsData) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RouteFilterConnectionsData) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RouteFilterConnectionsData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterConnectionsData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RouteFilterConnectionsData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RouteFilterConnectionsData) SetName(v string) {
	o.Name = &v
}

func (o RouteFilterConnectionsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteFilterConnectionsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteFilterConnectionsData) UnmarshalJSON(data []byte) (err error) {
	varRouteFilterConnectionsData := _RouteFilterConnectionsData{}

	err = json.Unmarshal(data, &varRouteFilterConnectionsData)

	if err != nil {
		return err
	}

	*o = RouteFilterConnectionsData(varRouteFilterConnectionsData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteFilterConnectionsData struct {
	value *RouteFilterConnectionsData
	isSet bool
}

func (v NullableRouteFilterConnectionsData) Get() *RouteFilterConnectionsData {
	return v.value
}

func (v *NullableRouteFilterConnectionsData) Set(val *RouteFilterConnectionsData) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFilterConnectionsData) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFilterConnectionsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFilterConnectionsData(val *RouteFilterConnectionsData) *NullableRouteFilterConnectionsData {
	return &NullableRouteFilterConnectionsData{value: val, isSet: true}
}

func (v NullableRouteFilterConnectionsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFilterConnectionsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}