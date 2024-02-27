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

// PrecisionTimeServiceCreateResponseType the model 'PrecisionTimeServiceCreateResponseType'
type PrecisionTimeServiceCreateResponseType string

// List of precisionTimeServiceCreateResponse_type
const (
	PRECISIONTIMESERVICECREATERESPONSETYPE_NTP PrecisionTimeServiceCreateResponseType = "NTP"
	PRECISIONTIMESERVICECREATERESPONSETYPE_PTP PrecisionTimeServiceCreateResponseType = "PTP"
)

// All allowed values of PrecisionTimeServiceCreateResponseType enum
var AllowedPrecisionTimeServiceCreateResponseTypeEnumValues = []PrecisionTimeServiceCreateResponseType{
	"NTP",
	"PTP",
}

func (v *PrecisionTimeServiceCreateResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrecisionTimeServiceCreateResponseType(value)
	for _, existing := range AllowedPrecisionTimeServiceCreateResponseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrecisionTimeServiceCreateResponseType", value)
}

// NewPrecisionTimeServiceCreateResponseTypeFromValue returns a pointer to a valid PrecisionTimeServiceCreateResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrecisionTimeServiceCreateResponseTypeFromValue(v string) (*PrecisionTimeServiceCreateResponseType, error) {
	ev := PrecisionTimeServiceCreateResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrecisionTimeServiceCreateResponseType: valid values are %v", v, AllowedPrecisionTimeServiceCreateResponseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrecisionTimeServiceCreateResponseType) IsValid() bool {
	for _, existing := range AllowedPrecisionTimeServiceCreateResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to precisionTimeServiceCreateResponse_type value
func (v PrecisionTimeServiceCreateResponseType) Ptr() *PrecisionTimeServiceCreateResponseType {
	return &v
}

type NullablePrecisionTimeServiceCreateResponseType struct {
	value *PrecisionTimeServiceCreateResponseType
	isSet bool
}

func (v NullablePrecisionTimeServiceCreateResponseType) Get() *PrecisionTimeServiceCreateResponseType {
	return v.value
}

func (v *NullablePrecisionTimeServiceCreateResponseType) Set(val *PrecisionTimeServiceCreateResponseType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrecisionTimeServiceCreateResponseType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrecisionTimeServiceCreateResponseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrecisionTimeServiceCreateResponseType(val *PrecisionTimeServiceCreateResponseType) *NullablePrecisionTimeServiceCreateResponseType {
	return &NullablePrecisionTimeServiceCreateResponseType{value: val, isSet: true}
}

func (v NullablePrecisionTimeServiceCreateResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrecisionTimeServiceCreateResponseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
