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

// checks if the ValidateRequestFilterAnd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateRequestFilterAnd{}

// ValidateRequestFilterAnd struct for ValidateRequestFilterAnd
type ValidateRequestFilterAnd struct {
	// Path to property
	Property *string `json:"property,omitempty"`
	// Type of operation
	Operator *string `json:"operator,omitempty"`
	// Values for the given property
	Values               []string `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ValidateRequestFilterAnd ValidateRequestFilterAnd

// NewValidateRequestFilterAnd instantiates a new ValidateRequestFilterAnd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateRequestFilterAnd() *ValidateRequestFilterAnd {
	this := ValidateRequestFilterAnd{}
	return &this
}

// NewValidateRequestFilterAndWithDefaults instantiates a new ValidateRequestFilterAnd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateRequestFilterAndWithDefaults() *ValidateRequestFilterAnd {
	this := ValidateRequestFilterAnd{}
	return &this
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *ValidateRequestFilterAnd) GetProperty() string {
	if o == nil || IsNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateRequestFilterAnd) GetPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *ValidateRequestFilterAnd) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *ValidateRequestFilterAnd) SetProperty(v string) {
	o.Property = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ValidateRequestFilterAnd) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateRequestFilterAnd) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ValidateRequestFilterAnd) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ValidateRequestFilterAnd) SetOperator(v string) {
	o.Operator = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ValidateRequestFilterAnd) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateRequestFilterAnd) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ValidateRequestFilterAnd) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ValidateRequestFilterAnd) SetValues(v []string) {
	o.Values = v
}

func (o ValidateRequestFilterAnd) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateRequestFilterAnd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidateRequestFilterAnd) UnmarshalJSON(data []byte) (err error) {
	varValidateRequestFilterAnd := _ValidateRequestFilterAnd{}

	err = json.Unmarshal(data, &varValidateRequestFilterAnd)

	if err != nil {
		return err
	}

	*o = ValidateRequestFilterAnd(varValidateRequestFilterAnd)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "property")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidateRequestFilterAnd struct {
	value *ValidateRequestFilterAnd
	isSet bool
}

func (v NullableValidateRequestFilterAnd) Get() *ValidateRequestFilterAnd {
	return v.value
}

func (v *NullableValidateRequestFilterAnd) Set(val *ValidateRequestFilterAnd) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateRequestFilterAnd) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateRequestFilterAnd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateRequestFilterAnd(val *ValidateRequestFilterAnd) *NullableValidateRequestFilterAnd {
	return &NullableValidateRequestFilterAnd{value: val, isSet: true}
}

func (v NullableValidateRequestFilterAnd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateRequestFilterAnd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}