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

// FabricCloudRouterCode Cloud Router code
type FabricCloudRouterCode string

// List of FabricCloudRouterCode
const (
	FABRICCLOUDROUTERCODE_LAB      FabricCloudRouterCode = "LAB"
	FABRICCLOUDROUTERCODE_ADVANCED FabricCloudRouterCode = "ADVANCED"
	FABRICCLOUDROUTERCODE_STANDARD FabricCloudRouterCode = "STANDARD"
	FABRICCLOUDROUTERCODE_PREMIUM  FabricCloudRouterCode = "PREMIUM"
)

// All allowed values of FabricCloudRouterCode enum
var AllowedFabricCloudRouterCodeEnumValues = []FabricCloudRouterCode{
	"LAB",
	"ADVANCED",
	"STANDARD",
	"PREMIUM",
}

func (v *FabricCloudRouterCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FabricCloudRouterCode(value)
	for _, existing := range AllowedFabricCloudRouterCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FabricCloudRouterCode", value)
}

// NewFabricCloudRouterCodeFromValue returns a pointer to a valid FabricCloudRouterCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFabricCloudRouterCodeFromValue(v string) (*FabricCloudRouterCode, error) {
	ev := FabricCloudRouterCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FabricCloudRouterCode: valid values are %v", v, AllowedFabricCloudRouterCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FabricCloudRouterCode) IsValid() bool {
	for _, existing := range AllowedFabricCloudRouterCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FabricCloudRouterCode value
func (v FabricCloudRouterCode) Ptr() *FabricCloudRouterCode {
	return &v
}

type NullableFabricCloudRouterCode struct {
	value *FabricCloudRouterCode
	isSet bool
}

func (v NullableFabricCloudRouterCode) Get() *FabricCloudRouterCode {
	return v.value
}

func (v *NullableFabricCloudRouterCode) Set(val *FabricCloudRouterCode) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricCloudRouterCode) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricCloudRouterCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricCloudRouterCode(val *FabricCloudRouterCode) *NullableFabricCloudRouterCode {
	return &NullableFabricCloudRouterCode{value: val, isSet: true}
}

func (v NullableFabricCloudRouterCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricCloudRouterCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}