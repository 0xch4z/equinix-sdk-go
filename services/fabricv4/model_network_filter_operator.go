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

// NetworkFilterOperator the model 'NetworkFilterOperator'
type NetworkFilterOperator string

// List of NetworkFilter_operator
const (
	NETWORKFILTEROPERATOR_EQUAL                    NetworkFilterOperator = "="
	NETWORKFILTEROPERATOR_NOT_EQUAL                NetworkFilterOperator = "!="
	NETWORKFILTEROPERATOR_GREATER_THAN             NetworkFilterOperator = ">"
	NETWORKFILTEROPERATOR_GREATER_THAN_OR_EQUAL_TO NetworkFilterOperator = ">="
	NETWORKFILTEROPERATOR_LESS_THAN                NetworkFilterOperator = "<"
	NETWORKFILTEROPERATOR_LESS_THAN_OR_EQUAL_TO    NetworkFilterOperator = "<="
	NETWORKFILTEROPERATOR_BETWEEN                  NetworkFilterOperator = "BETWEEN"
	NETWORKFILTEROPERATOR_NOT_BETWEEN              NetworkFilterOperator = "NOT BETWEEN"
	NETWORKFILTEROPERATOR_LIKE                     NetworkFilterOperator = "LIKE"
	NETWORKFILTEROPERATOR_NOT_LIKE                 NetworkFilterOperator = "NOT LIKE"
	NETWORKFILTEROPERATOR_ILIKE                    NetworkFilterOperator = "ILIKE"
	NETWORKFILTEROPERATOR_NOT_ILIKE                NetworkFilterOperator = "NOT ILIKE"
	NETWORKFILTEROPERATOR_IN                       NetworkFilterOperator = "IN"
	NETWORKFILTEROPERATOR_NOT_IN                   NetworkFilterOperator = "NOT IN"
)

// All allowed values of NetworkFilterOperator enum
var AllowedNetworkFilterOperatorEnumValues = []NetworkFilterOperator{
	"=",
	"!=",
	">",
	">=",
	"<",
	"<=",
	"BETWEEN",
	"NOT BETWEEN",
	"LIKE",
	"NOT LIKE",
	"ILIKE",
	"NOT ILIKE",
	"IN",
	"NOT IN",
}

func (v *NetworkFilterOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkFilterOperator(value)
	for _, existing := range AllowedNetworkFilterOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkFilterOperator", value)
}

// NewNetworkFilterOperatorFromValue returns a pointer to a valid NetworkFilterOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkFilterOperatorFromValue(v string) (*NetworkFilterOperator, error) {
	ev := NetworkFilterOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkFilterOperator: valid values are %v", v, AllowedNetworkFilterOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkFilterOperator) IsValid() bool {
	for _, existing := range AllowedNetworkFilterOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkFilter_operator value
func (v NetworkFilterOperator) Ptr() *NetworkFilterOperator {
	return &v
}

type NullableNetworkFilterOperator struct {
	value *NetworkFilterOperator
	isSet bool
}

func (v NullableNetworkFilterOperator) Get() *NetworkFilterOperator {
	return v.value
}

func (v *NullableNetworkFilterOperator) Set(val *NetworkFilterOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkFilterOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkFilterOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkFilterOperator(val *NetworkFilterOperator) *NullableNetworkFilterOperator {
	return &NullableNetworkFilterOperator{value: val, isSet: true}
}

func (v NullableNetworkFilterOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkFilterOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}