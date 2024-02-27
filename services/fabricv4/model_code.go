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

// Code Cloud Router package code
type Code string

// List of code
const (
	CODE_LAB      Code = "LAB"
	CODE_ADVANCED Code = "ADVANCED"
	CODE_STANDARD Code = "STANDARD"
	CODE_PREMIUM  Code = "PREMIUM"
)

// All allowed values of Code enum
var AllowedCodeEnumValues = []Code{
	"LAB",
	"ADVANCED",
	"STANDARD",
	"PREMIUM",
}

func (v *Code) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Code(value)
	for _, existing := range AllowedCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Code", value)
}

// NewCodeFromValue returns a pointer to a valid Code
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCodeFromValue(v string) (*Code, error) {
	ev := Code(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Code: valid values are %v", v, AllowedCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Code) IsValid() bool {
	for _, existing := range AllowedCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to code value
func (v Code) Ptr() *Code {
	return &v
}

type NullableCode struct {
	value *Code
	isSet bool
}

func (v NullableCode) Get() *Code {
	return v.value
}

func (v *NullableCode) Set(val *Code) {
	v.value = val
	v.isSet = true
}

func (v NullableCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCode(val *Code) *NullableCode {
	return &NullableCode{value: val, isSet: true}
}

func (v NullableCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
