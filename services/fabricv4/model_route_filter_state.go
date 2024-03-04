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

// RouteFilterState Route Filter status
type RouteFilterState string

// List of RouteFilterState
const (
	ROUTEFILTERSTATE_PROVISIONING      RouteFilterState = "PROVISIONING"
	ROUTEFILTERSTATE_REPROVISIONING    RouteFilterState = "REPROVISIONING"
	ROUTEFILTERSTATE_DEPROVISIONING    RouteFilterState = "DEPROVISIONING"
	ROUTEFILTERSTATE_PROVISIONED       RouteFilterState = "PROVISIONED"
	ROUTEFILTERSTATE_DEPROVISIONED     RouteFilterState = "DEPROVISIONED"
	ROUTEFILTERSTATE_NOT_PROVISIONED   RouteFilterState = "NOT_PROVISIONED"
	ROUTEFILTERSTATE_NOT_DEPROVISIONED RouteFilterState = "NOT_DEPROVISIONED"
)

// All allowed values of RouteFilterState enum
var AllowedRouteFilterStateEnumValues = []RouteFilterState{
	"PROVISIONING",
	"REPROVISIONING",
	"DEPROVISIONING",
	"PROVISIONED",
	"DEPROVISIONED",
	"NOT_PROVISIONED",
	"NOT_DEPROVISIONED",
}

func (v *RouteFilterState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RouteFilterState(value)
	for _, existing := range AllowedRouteFilterStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RouteFilterState", value)
}

// NewRouteFilterStateFromValue returns a pointer to a valid RouteFilterState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRouteFilterStateFromValue(v string) (*RouteFilterState, error) {
	ev := RouteFilterState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RouteFilterState: valid values are %v", v, AllowedRouteFilterStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RouteFilterState) IsValid() bool {
	for _, existing := range AllowedRouteFilterStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RouteFilterState value
func (v RouteFilterState) Ptr() *RouteFilterState {
	return &v
}

type NullableRouteFilterState struct {
	value *RouteFilterState
	isSet bool
}

func (v NullableRouteFilterState) Get() *RouteFilterState {
	return v.value
}

func (v *NullableRouteFilterState) Set(val *RouteFilterState) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFilterState) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFilterState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFilterState(val *RouteFilterState) *NullableRouteFilterState {
	return &NullableRouteFilterState{value: val, isSet: true}
}

func (v NullableRouteFilterState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFilterState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}