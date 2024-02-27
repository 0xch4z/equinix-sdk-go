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

// VirtualConnectionPriceAccessPointType Virtual Connection access point type
type VirtualConnectionPriceAccessPointType string

// List of VirtualConnectionPriceAccessPointType
const (
	VIRTUALCONNECTIONPRICEACCESSPOINTTYPE_VD           VirtualConnectionPriceAccessPointType = "VD"
	VIRTUALCONNECTIONPRICEACCESSPOINTTYPE_SP           VirtualConnectionPriceAccessPointType = "SP"
	VIRTUALCONNECTIONPRICEACCESSPOINTTYPE_COLO         VirtualConnectionPriceAccessPointType = "COLO"
	VIRTUALCONNECTIONPRICEACCESSPOINTTYPE_CLOUD_ROUTER VirtualConnectionPriceAccessPointType = "CLOUD_ROUTER"
	VIRTUALCONNECTIONPRICEACCESSPOINTTYPE_CHAINGROUP   VirtualConnectionPriceAccessPointType = "CHAINGROUP"
	VIRTUALCONNECTIONPRICEACCESSPOINTTYPE_NETWORK      VirtualConnectionPriceAccessPointType = "NETWORK"
)

// All allowed values of VirtualConnectionPriceAccessPointType enum
var AllowedVirtualConnectionPriceAccessPointTypeEnumValues = []VirtualConnectionPriceAccessPointType{
	"VD",
	"SP",
	"COLO",
	"CLOUD_ROUTER",
	"CHAINGROUP",
	"NETWORK",
}

func (v *VirtualConnectionPriceAccessPointType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualConnectionPriceAccessPointType(value)
	for _, existing := range AllowedVirtualConnectionPriceAccessPointTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualConnectionPriceAccessPointType", value)
}

// NewVirtualConnectionPriceAccessPointTypeFromValue returns a pointer to a valid VirtualConnectionPriceAccessPointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualConnectionPriceAccessPointTypeFromValue(v string) (*VirtualConnectionPriceAccessPointType, error) {
	ev := VirtualConnectionPriceAccessPointType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualConnectionPriceAccessPointType: valid values are %v", v, AllowedVirtualConnectionPriceAccessPointTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualConnectionPriceAccessPointType) IsValid() bool {
	for _, existing := range AllowedVirtualConnectionPriceAccessPointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VirtualConnectionPriceAccessPointType value
func (v VirtualConnectionPriceAccessPointType) Ptr() *VirtualConnectionPriceAccessPointType {
	return &v
}

type NullableVirtualConnectionPriceAccessPointType struct {
	value *VirtualConnectionPriceAccessPointType
	isSet bool
}

func (v NullableVirtualConnectionPriceAccessPointType) Get() *VirtualConnectionPriceAccessPointType {
	return v.value
}

func (v *NullableVirtualConnectionPriceAccessPointType) Set(val *VirtualConnectionPriceAccessPointType) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualConnectionPriceAccessPointType) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualConnectionPriceAccessPointType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualConnectionPriceAccessPointType(val *VirtualConnectionPriceAccessPointType) *NullableVirtualConnectionPriceAccessPointType {
	return &NullableVirtualConnectionPriceAccessPointType{value: val, isSet: true}
}

func (v NullableVirtualConnectionPriceAccessPointType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualConnectionPriceAccessPointType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
