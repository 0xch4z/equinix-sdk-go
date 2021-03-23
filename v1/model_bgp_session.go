/*
 * Metal API
 *
 * This is the API for Equinix Metal Product. Interact with your devices, user account, and projects.
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// BgpSession struct for BgpSession
type BgpSession struct {
	Id *string `json:"id,omitempty"`
	// The status of the BGP Session will start \"unknown\" and progress to \"up\" or \"down\" depending on the devices.
	Status *string `json:"status,omitempty"`
	LearnedRoutes *[]string `json:"learned_routes,omitempty"`
	AddressFamily *string `json:"address_family,omitempty"`
	Device *Href `json:"device,omitempty"`
	Href *string `json:"href,omitempty"`
	DefaultRoute *bool `json:"default_route,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewBgpSession instantiates a new BgpSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpSession() *BgpSession {
	this := BgpSession{}
	return &this
}

// NewBgpSessionWithDefaults instantiates a new BgpSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpSessionWithDefaults() *BgpSession {
	this := BgpSession{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BgpSession) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BgpSession) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BgpSession) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BgpSession) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BgpSession) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BgpSession) SetStatus(v string) {
	o.Status = &v
}

// GetLearnedRoutes returns the LearnedRoutes field value if set, zero value otherwise.
func (o *BgpSession) GetLearnedRoutes() []string {
	if o == nil || o.LearnedRoutes == nil {
		var ret []string
		return ret
	}
	return *o.LearnedRoutes
}

// GetLearnedRoutesOk returns a tuple with the LearnedRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetLearnedRoutesOk() (*[]string, bool) {
	if o == nil || o.LearnedRoutes == nil {
		return nil, false
	}
	return o.LearnedRoutes, true
}

// HasLearnedRoutes returns a boolean if a field has been set.
func (o *BgpSession) HasLearnedRoutes() bool {
	if o != nil && o.LearnedRoutes != nil {
		return true
	}

	return false
}

// SetLearnedRoutes gets a reference to the given []string and assigns it to the LearnedRoutes field.
func (o *BgpSession) SetLearnedRoutes(v []string) {
	o.LearnedRoutes = &v
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *BgpSession) GetAddressFamily() string {
	if o == nil || o.AddressFamily == nil {
		var ret string
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetAddressFamilyOk() (*string, bool) {
	if o == nil || o.AddressFamily == nil {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *BgpSession) HasAddressFamily() bool {
	if o != nil && o.AddressFamily != nil {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given string and assigns it to the AddressFamily field.
func (o *BgpSession) SetAddressFamily(v string) {
	o.AddressFamily = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *BgpSession) GetDevice() Href {
	if o == nil || o.Device == nil {
		var ret Href
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetDeviceOk() (*Href, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *BgpSession) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Href and assigns it to the Device field.
func (o *BgpSession) SetDevice(v Href) {
	o.Device = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BgpSession) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BgpSession) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BgpSession) SetHref(v string) {
	o.Href = &v
}

// GetDefaultRoute returns the DefaultRoute field value if set, zero value otherwise.
func (o *BgpSession) GetDefaultRoute() bool {
	if o == nil || o.DefaultRoute == nil {
		var ret bool
		return ret
	}
	return *o.DefaultRoute
}

// GetDefaultRouteOk returns a tuple with the DefaultRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetDefaultRouteOk() (*bool, bool) {
	if o == nil || o.DefaultRoute == nil {
		return nil, false
	}
	return o.DefaultRoute, true
}

// HasDefaultRoute returns a boolean if a field has been set.
func (o *BgpSession) HasDefaultRoute() bool {
	if o != nil && o.DefaultRoute != nil {
		return true
	}

	return false
}

// SetDefaultRoute gets a reference to the given bool and assigns it to the DefaultRoute field.
func (o *BgpSession) SetDefaultRoute(v bool) {
	o.DefaultRoute = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BgpSession) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BgpSession) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BgpSession) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BgpSession) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSession) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BgpSession) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *BgpSession) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o BgpSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.LearnedRoutes != nil {
		toSerialize["learned_routes"] = o.LearnedRoutes
	}
	if o.AddressFamily != nil {
		toSerialize["address_family"] = o.AddressFamily
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.DefaultRoute != nil {
		toSerialize["default_route"] = o.DefaultRoute
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableBgpSession struct {
	value *BgpSession
	isSet bool
}

func (v NullableBgpSession) Get() *BgpSession {
	return v.value
}

func (v *NullableBgpSession) Set(val *BgpSession) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpSession) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpSession(val *BgpSession) *NullableBgpSession {
	return &NullableBgpSession{value: val, isSet: true}
}

func (v NullableBgpSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


