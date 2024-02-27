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

// checks if the Expression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Expression{}

// Expression struct for Expression
type Expression struct {
	And                  []Expression        `json:"and,omitempty"`
	Or                   []Expression        `json:"or,omitempty"`
	Property             *SearchFieldName    `json:"property,omitempty"`
	Operator             *ExpressionOperator `json:"operator,omitempty"`
	Values               []string            `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Expression Expression

// NewExpression instantiates a new Expression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpression() *Expression {
	this := Expression{}
	return &this
}

// NewExpressionWithDefaults instantiates a new Expression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpressionWithDefaults() *Expression {
	this := Expression{}
	return &this
}

// GetAnd returns the And field value if set, zero value otherwise.
func (o *Expression) GetAnd() []Expression {
	if o == nil || IsNil(o.And) {
		var ret []Expression
		return ret
	}
	return o.And
}

// GetAndOk returns a tuple with the And field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expression) GetAndOk() ([]Expression, bool) {
	if o == nil || IsNil(o.And) {
		return nil, false
	}
	return o.And, true
}

// HasAnd returns a boolean if a field has been set.
func (o *Expression) HasAnd() bool {
	if o != nil && !IsNil(o.And) {
		return true
	}

	return false
}

// SetAnd gets a reference to the given []Expression and assigns it to the And field.
func (o *Expression) SetAnd(v []Expression) {
	o.And = v
}

// GetOr returns the Or field value if set, zero value otherwise.
func (o *Expression) GetOr() []Expression {
	if o == nil || IsNil(o.Or) {
		var ret []Expression
		return ret
	}
	return o.Or
}

// GetOrOk returns a tuple with the Or field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expression) GetOrOk() ([]Expression, bool) {
	if o == nil || IsNil(o.Or) {
		return nil, false
	}
	return o.Or, true
}

// HasOr returns a boolean if a field has been set.
func (o *Expression) HasOr() bool {
	if o != nil && !IsNil(o.Or) {
		return true
	}

	return false
}

// SetOr gets a reference to the given []Expression and assigns it to the Or field.
func (o *Expression) SetOr(v []Expression) {
	o.Or = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *Expression) GetProperty() SearchFieldName {
	if o == nil || IsNil(o.Property) {
		var ret SearchFieldName
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expression) GetPropertyOk() (*SearchFieldName, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *Expression) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given SearchFieldName and assigns it to the Property field.
func (o *Expression) SetProperty(v SearchFieldName) {
	o.Property = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *Expression) GetOperator() ExpressionOperator {
	if o == nil || IsNil(o.Operator) {
		var ret ExpressionOperator
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expression) GetOperatorOk() (*ExpressionOperator, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *Expression) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given ExpressionOperator and assigns it to the Operator field.
func (o *Expression) SetOperator(v ExpressionOperator) {
	o.Operator = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Expression) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expression) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Expression) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *Expression) SetValues(v []string) {
	o.Values = v
}

func (o Expression) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Expression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.And) {
		toSerialize["and"] = o.And
	}
	if !IsNil(o.Or) {
		toSerialize["or"] = o.Or
	}
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

func (o *Expression) UnmarshalJSON(data []byte) (err error) {
	varExpression := _Expression{}

	err = json.Unmarshal(data, &varExpression)

	if err != nil {
		return err
	}

	*o = Expression(varExpression)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "and")
		delete(additionalProperties, "or")
		delete(additionalProperties, "property")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExpression struct {
	value *Expression
	isSet bool
}

func (v NullableExpression) Get() *Expression {
	return v.value
}

func (v *NullableExpression) Set(val *Expression) {
	v.value = val
	v.isSet = true
}

func (v NullableExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpression(val *Expression) *NullableExpression {
	return &NullableExpression{value: val, isSet: true}
}

func (v NullableExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
