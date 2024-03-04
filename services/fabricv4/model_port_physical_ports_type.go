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

// PortPhysicalPortsType Physical Ports Type
type PortPhysicalPortsType string

// List of Port_physicalPortsType
const (
	PORTPHYSICALPORTSTYPE__1000_BASE_LX  PortPhysicalPortsType = "1000BASE_LX"
	PORTPHYSICALPORTSTYPE__10_GBASE_LR   PortPhysicalPortsType = "10GBASE_LR"
	PORTPHYSICALPORTSTYPE__100_GBASE_LR4 PortPhysicalPortsType = "100GBASE_LR4"
	PORTPHYSICALPORTSTYPE__10_GBASE_ER   PortPhysicalPortsType = "10GBASE_ER"
	PORTPHYSICALPORTSTYPE__1000_BASE_SX  PortPhysicalPortsType = "1000BASE_SX"
)

// All allowed values of PortPhysicalPortsType enum
var AllowedPortPhysicalPortsTypeEnumValues = []PortPhysicalPortsType{
	"1000BASE_LX",
	"10GBASE_LR",
	"100GBASE_LR4",
	"10GBASE_ER",
	"1000BASE_SX",
}

func (v *PortPhysicalPortsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortPhysicalPortsType(value)
	for _, existing := range AllowedPortPhysicalPortsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortPhysicalPortsType", value)
}

// NewPortPhysicalPortsTypeFromValue returns a pointer to a valid PortPhysicalPortsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortPhysicalPortsTypeFromValue(v string) (*PortPhysicalPortsType, error) {
	ev := PortPhysicalPortsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortPhysicalPortsType: valid values are %v", v, AllowedPortPhysicalPortsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortPhysicalPortsType) IsValid() bool {
	for _, existing := range AllowedPortPhysicalPortsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Port_physicalPortsType value
func (v PortPhysicalPortsType) Ptr() *PortPhysicalPortsType {
	return &v
}

type NullablePortPhysicalPortsType struct {
	value *PortPhysicalPortsType
	isSet bool
}

func (v NullablePortPhysicalPortsType) Get() *PortPhysicalPortsType {
	return v.value
}

func (v *NullablePortPhysicalPortsType) Set(val *PortPhysicalPortsType) {
	v.value = val
	v.isSet = true
}

func (v NullablePortPhysicalPortsType) IsSet() bool {
	return v.isSet
}

func (v *NullablePortPhysicalPortsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortPhysicalPortsType(val *PortPhysicalPortsType) *NullablePortPhysicalPortsType {
	return &NullablePortPhysicalPortsType{value: val, isSet: true}
}

func (v NullablePortPhysicalPortsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortPhysicalPortsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}