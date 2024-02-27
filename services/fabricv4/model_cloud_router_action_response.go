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

// checks if the CloudRouterActionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouterActionResponse{}

// CloudRouterActionResponse Cloud Router actions response object
type CloudRouterActionResponse struct {
	Type                 CloudRouterActionType  `json:"type"`
	Uuid                 string                 `json:"uuid"`
	Description          *string                `json:"description,omitempty"`
	State                CloudRouterActionState `json:"state"`
	ChangeLog            Changelog              `json:"changeLog"`
	AdditionalProperties map[string]interface{}
}

type _CloudRouterActionResponse CloudRouterActionResponse

// NewCloudRouterActionResponse instantiates a new CloudRouterActionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouterActionResponse(type_ CloudRouterActionType, uuid string, state CloudRouterActionState, changeLog Changelog) *CloudRouterActionResponse {
	this := CloudRouterActionResponse{}
	this.Type = type_
	this.Uuid = uuid
	this.State = state
	this.ChangeLog = changeLog
	return &this
}

// NewCloudRouterActionResponseWithDefaults instantiates a new CloudRouterActionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterActionResponseWithDefaults() *CloudRouterActionResponse {
	this := CloudRouterActionResponse{}
	return &this
}

// GetType returns the Type field value
func (o *CloudRouterActionResponse) GetType() CloudRouterActionType {
	if o == nil {
		var ret CloudRouterActionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudRouterActionResponse) GetTypeOk() (*CloudRouterActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CloudRouterActionResponse) SetType(v CloudRouterActionType) {
	o.Type = v
}

// GetUuid returns the Uuid field value
func (o *CloudRouterActionResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *CloudRouterActionResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *CloudRouterActionResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudRouterActionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterActionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudRouterActionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudRouterActionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value
func (o *CloudRouterActionResponse) GetState() CloudRouterActionState {
	if o == nil {
		var ret CloudRouterActionState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CloudRouterActionResponse) GetStateOk() (*CloudRouterActionState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CloudRouterActionResponse) SetState(v CloudRouterActionState) {
	o.State = v
}

// GetChangeLog returns the ChangeLog field value
func (o *CloudRouterActionResponse) GetChangeLog() Changelog {
	if o == nil {
		var ret Changelog
		return ret
	}

	return o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value
// and a boolean to check if the value has been set.
func (o *CloudRouterActionResponse) GetChangeLogOk() (*Changelog, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeLog, true
}

// SetChangeLog sets field value
func (o *CloudRouterActionResponse) SetChangeLog(v Changelog) {
	o.ChangeLog = v
}

func (o CloudRouterActionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouterActionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["state"] = o.State
	toSerialize["changeLog"] = o.ChangeLog

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudRouterActionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"uuid",
		"state",
		"changeLog",
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

	varCloudRouterActionResponse := _CloudRouterActionResponse{}

	err = json.Unmarshal(data, &varCloudRouterActionResponse)

	if err != nil {
		return err
	}

	*o = CloudRouterActionResponse(varCloudRouterActionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "description")
		delete(additionalProperties, "state")
		delete(additionalProperties, "changeLog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouterActionResponse struct {
	value *CloudRouterActionResponse
	isSet bool
}

func (v NullableCloudRouterActionResponse) Get() *CloudRouterActionResponse {
	return v.value
}

func (v *NullableCloudRouterActionResponse) Set(val *CloudRouterActionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterActionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterActionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterActionResponse(val *CloudRouterActionResponse) *NullableCloudRouterActionResponse {
	return &NullableCloudRouterActionResponse{value: val, isSet: true}
}

func (v NullableCloudRouterActionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterActionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
