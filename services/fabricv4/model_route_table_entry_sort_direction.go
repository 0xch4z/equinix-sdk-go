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

// RouteTableEntrySortDirection Sorting direction
type RouteTableEntrySortDirection string

// List of RouteTableEntrySortDirection
const (
	ROUTETABLEENTRYSORTDIRECTION_DESC RouteTableEntrySortDirection = "DESC"
	ROUTETABLEENTRYSORTDIRECTION_ASC  RouteTableEntrySortDirection = "ASC"
)

// All allowed values of RouteTableEntrySortDirection enum
var AllowedRouteTableEntrySortDirectionEnumValues = []RouteTableEntrySortDirection{
	"DESC",
	"ASC",
}

func (v *RouteTableEntrySortDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RouteTableEntrySortDirection(value)
	for _, existing := range AllowedRouteTableEntrySortDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RouteTableEntrySortDirection", value)
}

// NewRouteTableEntrySortDirectionFromValue returns a pointer to a valid RouteTableEntrySortDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRouteTableEntrySortDirectionFromValue(v string) (*RouteTableEntrySortDirection, error) {
	ev := RouteTableEntrySortDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RouteTableEntrySortDirection: valid values are %v", v, AllowedRouteTableEntrySortDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RouteTableEntrySortDirection) IsValid() bool {
	for _, existing := range AllowedRouteTableEntrySortDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RouteTableEntrySortDirection value
func (v RouteTableEntrySortDirection) Ptr() *RouteTableEntrySortDirection {
	return &v
}

type NullableRouteTableEntrySortDirection struct {
	value *RouteTableEntrySortDirection
	isSet bool
}

func (v NullableRouteTableEntrySortDirection) Get() *RouteTableEntrySortDirection {
	return v.value
}

func (v *NullableRouteTableEntrySortDirection) Set(val *RouteTableEntrySortDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteTableEntrySortDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteTableEntrySortDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteTableEntrySortDirection(val *RouteTableEntrySortDirection) *NullableRouteTableEntrySortDirection {
	return &NullableRouteTableEntrySortDirection{value: val, isSet: true}
}

func (v NullableRouteTableEntrySortDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteTableEntrySortDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
