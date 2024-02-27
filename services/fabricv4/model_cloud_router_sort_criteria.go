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

// checks if the CloudRouterSortCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouterSortCriteria{}

// CloudRouterSortCriteria struct for CloudRouterSortCriteria
type CloudRouterSortCriteria struct {
	Direction            *CloudRouterSortDirection `json:"direction,omitempty"`
	Property             *CloudRouterSortBy        `json:"property,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudRouterSortCriteria CloudRouterSortCriteria

// NewCloudRouterSortCriteria instantiates a new CloudRouterSortCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouterSortCriteria() *CloudRouterSortCriteria {
	this := CloudRouterSortCriteria{}
	var direction CloudRouterSortDirection = CLOUDROUTERSORTDIRECTION_DESC
	this.Direction = &direction
	var property CloudRouterSortBy = CLOUDROUTERSORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// NewCloudRouterSortCriteriaWithDefaults instantiates a new CloudRouterSortCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterSortCriteriaWithDefaults() *CloudRouterSortCriteria {
	this := CloudRouterSortCriteria{}
	var direction CloudRouterSortDirection = CLOUDROUTERSORTDIRECTION_DESC
	this.Direction = &direction
	var property CloudRouterSortBy = CLOUDROUTERSORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *CloudRouterSortCriteria) GetDirection() CloudRouterSortDirection {
	if o == nil || IsNil(o.Direction) {
		var ret CloudRouterSortDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterSortCriteria) GetDirectionOk() (*CloudRouterSortDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *CloudRouterSortCriteria) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given CloudRouterSortDirection and assigns it to the Direction field.
func (o *CloudRouterSortCriteria) SetDirection(v CloudRouterSortDirection) {
	o.Direction = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *CloudRouterSortCriteria) GetProperty() CloudRouterSortBy {
	if o == nil || IsNil(o.Property) {
		var ret CloudRouterSortBy
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterSortCriteria) GetPropertyOk() (*CloudRouterSortBy, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *CloudRouterSortCriteria) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given CloudRouterSortBy and assigns it to the Property field.
func (o *CloudRouterSortCriteria) SetProperty(v CloudRouterSortBy) {
	o.Property = &v
}

func (o CloudRouterSortCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouterSortCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudRouterSortCriteria) UnmarshalJSON(data []byte) (err error) {
	varCloudRouterSortCriteria := _CloudRouterSortCriteria{}

	err = json.Unmarshal(data, &varCloudRouterSortCriteria)

	if err != nil {
		return err
	}

	*o = CloudRouterSortCriteria(varCloudRouterSortCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		delete(additionalProperties, "property")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouterSortCriteria struct {
	value *CloudRouterSortCriteria
	isSet bool
}

func (v NullableCloudRouterSortCriteria) Get() *CloudRouterSortCriteria {
	return v.value
}

func (v *NullableCloudRouterSortCriteria) Set(val *CloudRouterSortCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterSortCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterSortCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterSortCriteria(val *CloudRouterSortCriteria) *NullableCloudRouterSortCriteria {
	return &NullableCloudRouterSortCriteria{value: val, isSet: true}
}

func (v NullableCloudRouterSortCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterSortCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
