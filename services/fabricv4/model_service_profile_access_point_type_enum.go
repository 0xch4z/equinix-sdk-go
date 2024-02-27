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

// ServiceProfileAccessPointTypeEnum Access point type
type ServiceProfileAccessPointTypeEnum string

// List of ServiceProfileAccessPointTypeEnum
const (
	SERVICEPROFILEACCESSPOINTTYPEENUM_VD   ServiceProfileAccessPointTypeEnum = "VD"
	SERVICEPROFILEACCESSPOINTTYPEENUM_COLO ServiceProfileAccessPointTypeEnum = "COLO"
)

// All allowed values of ServiceProfileAccessPointTypeEnum enum
var AllowedServiceProfileAccessPointTypeEnumEnumValues = []ServiceProfileAccessPointTypeEnum{
	"VD",
	"COLO",
}

func (v *ServiceProfileAccessPointTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceProfileAccessPointTypeEnum(value)
	for _, existing := range AllowedServiceProfileAccessPointTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceProfileAccessPointTypeEnum", value)
}

// NewServiceProfileAccessPointTypeEnumFromValue returns a pointer to a valid ServiceProfileAccessPointTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceProfileAccessPointTypeEnumFromValue(v string) (*ServiceProfileAccessPointTypeEnum, error) {
	ev := ServiceProfileAccessPointTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceProfileAccessPointTypeEnum: valid values are %v", v, AllowedServiceProfileAccessPointTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceProfileAccessPointTypeEnum) IsValid() bool {
	for _, existing := range AllowedServiceProfileAccessPointTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceProfileAccessPointTypeEnum value
func (v ServiceProfileAccessPointTypeEnum) Ptr() *ServiceProfileAccessPointTypeEnum {
	return &v
}

type NullableServiceProfileAccessPointTypeEnum struct {
	value *ServiceProfileAccessPointTypeEnum
	isSet bool
}

func (v NullableServiceProfileAccessPointTypeEnum) Get() *ServiceProfileAccessPointTypeEnum {
	return v.value
}

func (v *NullableServiceProfileAccessPointTypeEnum) Set(val *ServiceProfileAccessPointTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileAccessPointTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileAccessPointTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileAccessPointTypeEnum(val *ServiceProfileAccessPointTypeEnum) *NullableServiceProfileAccessPointTypeEnum {
	return &NullableServiceProfileAccessPointTypeEnum{value: val, isSet: true}
}

func (v NullableServiceProfileAccessPointTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileAccessPointTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
