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

// VirtualConnectionBridgePackageCode Virtual Connection bridge package code
type VirtualConnectionBridgePackageCode string

// List of VirtualConnectionBridgePackageCode
const (
	VIRTUALCONNECTIONBRIDGEPACKAGECODE_REGIONAL VirtualConnectionBridgePackageCode = "REGIONAL"
	VIRTUALCONNECTIONBRIDGEPACKAGECODE_GLOBAL   VirtualConnectionBridgePackageCode = "GLOBAL"
)

// All allowed values of VirtualConnectionBridgePackageCode enum
var AllowedVirtualConnectionBridgePackageCodeEnumValues = []VirtualConnectionBridgePackageCode{
	"REGIONAL",
	"GLOBAL",
}

func (v *VirtualConnectionBridgePackageCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualConnectionBridgePackageCode(value)
	for _, existing := range AllowedVirtualConnectionBridgePackageCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualConnectionBridgePackageCode", value)
}

// NewVirtualConnectionBridgePackageCodeFromValue returns a pointer to a valid VirtualConnectionBridgePackageCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualConnectionBridgePackageCodeFromValue(v string) (*VirtualConnectionBridgePackageCode, error) {
	ev := VirtualConnectionBridgePackageCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualConnectionBridgePackageCode: valid values are %v", v, AllowedVirtualConnectionBridgePackageCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualConnectionBridgePackageCode) IsValid() bool {
	for _, existing := range AllowedVirtualConnectionBridgePackageCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VirtualConnectionBridgePackageCode value
func (v VirtualConnectionBridgePackageCode) Ptr() *VirtualConnectionBridgePackageCode {
	return &v
}

type NullableVirtualConnectionBridgePackageCode struct {
	value *VirtualConnectionBridgePackageCode
	isSet bool
}

func (v NullableVirtualConnectionBridgePackageCode) Get() *VirtualConnectionBridgePackageCode {
	return v.value
}

func (v *NullableVirtualConnectionBridgePackageCode) Set(val *VirtualConnectionBridgePackageCode) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualConnectionBridgePackageCode) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualConnectionBridgePackageCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualConnectionBridgePackageCode(val *VirtualConnectionBridgePackageCode) *NullableVirtualConnectionBridgePackageCode {
	return &NullableVirtualConnectionBridgePackageCode{value: val, isSet: true}
}

func (v NullableVirtualConnectionBridgePackageCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualConnectionBridgePackageCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
