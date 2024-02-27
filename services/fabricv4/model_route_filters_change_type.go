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

// RouteFiltersChangeType Type of change
type RouteFiltersChangeType string

// List of RouteFiltersChange_type
const (
	ROUTEFILTERSCHANGETYPE_IPV4_PREFIX_FILTER_UPDATE   RouteFiltersChangeType = "BGP_IPv4_PREFIX_FILTER_UPDATE"
	ROUTEFILTERSCHANGETYPE_IPV4_PREFIX_FILTER_CREATION RouteFiltersChangeType = "BGP_IPv4_PREFIX_FILTER_CREATION"
	ROUTEFILTERSCHANGETYPE_IPV4_PREFIX_FILTER_DELETION RouteFiltersChangeType = "BGP_IPv4_PREFIX_FILTER_DELETION"
	ROUTEFILTERSCHANGETYPE_IPV6_PREFIX_FILTER_UPDATE   RouteFiltersChangeType = "BGP_IPv6_PREFIX_FILTER_UPDATE"
	ROUTEFILTERSCHANGETYPE_IPV6_PREFIX_FILTER_CREATION RouteFiltersChangeType = "BGP_IPv6_PREFIX_FILTER_CREATION"
	ROUTEFILTERSCHANGETYPE_IPV6_PREFIX_FILTER_DELETION RouteFiltersChangeType = "BGP_IPv6_PREFIX_FILTER_DELETION"
)

// All allowed values of RouteFiltersChangeType enum
var AllowedRouteFiltersChangeTypeEnumValues = []RouteFiltersChangeType{
	"BGP_IPv4_PREFIX_FILTER_UPDATE",
	"BGP_IPv4_PREFIX_FILTER_CREATION",
	"BGP_IPv4_PREFIX_FILTER_DELETION",
	"BGP_IPv6_PREFIX_FILTER_UPDATE",
	"BGP_IPv6_PREFIX_FILTER_CREATION",
	"BGP_IPv6_PREFIX_FILTER_DELETION",
}

func (v *RouteFiltersChangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RouteFiltersChangeType(value)
	for _, existing := range AllowedRouteFiltersChangeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RouteFiltersChangeType", value)
}

// NewRouteFiltersChangeTypeFromValue returns a pointer to a valid RouteFiltersChangeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRouteFiltersChangeTypeFromValue(v string) (*RouteFiltersChangeType, error) {
	ev := RouteFiltersChangeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RouteFiltersChangeType: valid values are %v", v, AllowedRouteFiltersChangeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RouteFiltersChangeType) IsValid() bool {
	for _, existing := range AllowedRouteFiltersChangeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RouteFiltersChange_type value
func (v RouteFiltersChangeType) Ptr() *RouteFiltersChangeType {
	return &v
}

type NullableRouteFiltersChangeType struct {
	value *RouteFiltersChangeType
	isSet bool
}

func (v NullableRouteFiltersChangeType) Get() *RouteFiltersChangeType {
	return v.value
}

func (v *NullableRouteFiltersChangeType) Set(val *RouteFiltersChangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFiltersChangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFiltersChangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFiltersChangeType(val *RouteFiltersChangeType) *NullableRouteFiltersChangeType {
	return &NullableRouteFiltersChangeType{value: val, isSet: true}
}

func (v NullableRouteFiltersChangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFiltersChangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
