/*
Metal API

This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// GetInterconnection200ResponseServiceTokensInner struct for GetInterconnection200ResponseServiceTokensInner
type GetInterconnection200ResponseServiceTokensInner struct {
	// The expiration date and time of the Fabric service token. Once a service token is expired, it is no longer redeemable.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// The service token UUID that can be used on the Fabric Portal to create an connection from Metal to another Fabric service provider.
	Id *string `json:"id,omitempty"`
	// The maximum speed that can be selected on the Fabric Portal when configuring a connection with the service token. The speed is recorded in bps, but can be set by using any of the following units: 'bps', 'mbps', or 'gbps'. This speed is automatically capped depending on the tier of the organization. If you would like to upgrade to another tier, please contact our Support team.
	MaxAllowedSpeed *string `json:"max_allowed_speed,omitempty"`
	// Either primary or redundant, depending on the role of the connection port the token is associated with.
	Role *string `json:"role,omitempty"`
	// The type of service token that has been created. Currently, only A-side service tokens are available.
	ServiceTokenType *string `json:"service_token_type,omitempty"`
	// The state of the service token that corresponds with the service token state on Fabric. An inactive state refers to a token that has not been redeemed yet on the Fabric side, an active state refers to a token that has already been redeemed, and an expired state refers to a token that has reached its expiry time.
	State *string `json:"state,omitempty"`
}

// NewGetInterconnection200ResponseServiceTokensInner instantiates a new GetInterconnection200ResponseServiceTokensInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInterconnection200ResponseServiceTokensInner() *GetInterconnection200ResponseServiceTokensInner {
	this := GetInterconnection200ResponseServiceTokensInner{}
	return &this
}

// NewGetInterconnection200ResponseServiceTokensInnerWithDefaults instantiates a new GetInterconnection200ResponseServiceTokensInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInterconnection200ResponseServiceTokensInnerWithDefaults() *GetInterconnection200ResponseServiceTokensInner {
	this := GetInterconnection200ResponseServiceTokensInner{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *GetInterconnection200ResponseServiceTokensInner) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *GetInterconnection200ResponseServiceTokensInner) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetInterconnection200ResponseServiceTokensInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetInterconnection200ResponseServiceTokensInner) SetId(v string) {
	o.Id = &v
}

// GetMaxAllowedSpeed returns the MaxAllowedSpeed field value if set, zero value otherwise.
func (o *GetInterconnection200ResponseServiceTokensInner) GetMaxAllowedSpeed() string {
	if o == nil || o.MaxAllowedSpeed == nil {
		var ret string
		return ret
	}
	return *o.MaxAllowedSpeed
}

// GetMaxAllowedSpeedOk returns a tuple with the MaxAllowedSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) GetMaxAllowedSpeedOk() (*string, bool) {
	if o == nil || o.MaxAllowedSpeed == nil {
		return nil, false
	}
	return o.MaxAllowedSpeed, true
}

// HasMaxAllowedSpeed returns a boolean if a field has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) HasMaxAllowedSpeed() bool {
	if o != nil && o.MaxAllowedSpeed != nil {
		return true
	}

	return false
}

// SetMaxAllowedSpeed gets a reference to the given string and assigns it to the MaxAllowedSpeed field.
func (o *GetInterconnection200ResponseServiceTokensInner) SetMaxAllowedSpeed(v string) {
	o.MaxAllowedSpeed = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *GetInterconnection200ResponseServiceTokensInner) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *GetInterconnection200ResponseServiceTokensInner) SetRole(v string) {
	o.Role = &v
}

// GetServiceTokenType returns the ServiceTokenType field value if set, zero value otherwise.
func (o *GetInterconnection200ResponseServiceTokensInner) GetServiceTokenType() string {
	if o == nil || o.ServiceTokenType == nil {
		var ret string
		return ret
	}
	return *o.ServiceTokenType
}

// GetServiceTokenTypeOk returns a tuple with the ServiceTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) GetServiceTokenTypeOk() (*string, bool) {
	if o == nil || o.ServiceTokenType == nil {
		return nil, false
	}
	return o.ServiceTokenType, true
}

// HasServiceTokenType returns a boolean if a field has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) HasServiceTokenType() bool {
	if o != nil && o.ServiceTokenType != nil {
		return true
	}

	return false
}

// SetServiceTokenType gets a reference to the given string and assigns it to the ServiceTokenType field.
func (o *GetInterconnection200ResponseServiceTokensInner) SetServiceTokenType(v string) {
	o.ServiceTokenType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetInterconnection200ResponseServiceTokensInner) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetInterconnection200ResponseServiceTokensInner) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetInterconnection200ResponseServiceTokensInner) SetState(v string) {
	o.State = &v
}

func (o GetInterconnection200ResponseServiceTokensInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MaxAllowedSpeed != nil {
		toSerialize["max_allowed_speed"] = o.MaxAllowedSpeed
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.ServiceTokenType != nil {
		toSerialize["service_token_type"] = o.ServiceTokenType
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableGetInterconnection200ResponseServiceTokensInner struct {
	value *GetInterconnection200ResponseServiceTokensInner
	isSet bool
}

func (v NullableGetInterconnection200ResponseServiceTokensInner) Get() *GetInterconnection200ResponseServiceTokensInner {
	return v.value
}

func (v *NullableGetInterconnection200ResponseServiceTokensInner) Set(val *GetInterconnection200ResponseServiceTokensInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInterconnection200ResponseServiceTokensInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInterconnection200ResponseServiceTokensInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInterconnection200ResponseServiceTokensInner(val *GetInterconnection200ResponseServiceTokensInner) *NullableGetInterconnection200ResponseServiceTokensInner {
	return &NullableGetInterconnection200ResponseServiceTokensInner{value: val, isSet: true}
}

func (v NullableGetInterconnection200ResponseServiceTokensInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInterconnection200ResponseServiceTokensInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
