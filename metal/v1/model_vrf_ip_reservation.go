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

// VrfIpReservation struct for VrfIpReservation
type VrfIpReservation struct {
	Tags          []string                                                        `json:"tags,omitempty"`
	AddressFamily *int32                                                          `json:"address_family,omitempty"`
	Cidr          *int32                                                          `json:"cidr,omitempty"`
	CreatedAt     *time.Time                                                      `json:"created_at,omitempty"`
	CreatedBy     *FindBatchById200ResponseDevicesInner                           `json:"created_by,omitempty"`
	Details       *string                                                         `json:"details,omitempty"`
	Href          *string                                                         `json:"href,omitempty"`
	Id            *string                                                         `json:"id,omitempty"`
	MetalGateway  *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway `json:"metal_gateway,omitempty"`
	Netmask       *string                                                         `json:"netmask,omitempty"`
	Network       *string                                                         `json:"network,omitempty"`
	Project       *MoveHardwareReservation201ResponseProject                      `json:"project,omitempty"`
	State         *string                                                         `json:"state,omitempty"`
	Type          *string                                                         `json:"type,omitempty"`
	Vrf           *FindVrfs200ResponseVrfsInner                                   `json:"vrf,omitempty"`
}

// NewVrfIpReservation instantiates a new VrfIpReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfIpReservation() *VrfIpReservation {
	this := VrfIpReservation{}
	return &this
}

// NewVrfIpReservationWithDefaults instantiates a new VrfIpReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfIpReservationWithDefaults() *VrfIpReservation {
	this := VrfIpReservation{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VrfIpReservation) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VrfIpReservation) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VrfIpReservation) SetTags(v []string) {
	o.Tags = v
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *VrfIpReservation) GetAddressFamily() int32 {
	if o == nil || o.AddressFamily == nil {
		var ret int32
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetAddressFamilyOk() (*int32, bool) {
	if o == nil || o.AddressFamily == nil {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *VrfIpReservation) HasAddressFamily() bool {
	if o != nil && o.AddressFamily != nil {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given int32 and assigns it to the AddressFamily field.
func (o *VrfIpReservation) SetAddressFamily(v int32) {
	o.AddressFamily = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *VrfIpReservation) GetCidr() int32 {
	if o == nil || o.Cidr == nil {
		var ret int32
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetCidrOk() (*int32, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *VrfIpReservation) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given int32 and assigns it to the Cidr field.
func (o *VrfIpReservation) SetCidr(v int32) {
	o.Cidr = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VrfIpReservation) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VrfIpReservation) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VrfIpReservation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *VrfIpReservation) GetCreatedBy() FindBatchById200ResponseDevicesInner {
	if o == nil || o.CreatedBy == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetCreatedByOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *VrfIpReservation) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the CreatedBy field.
func (o *VrfIpReservation) SetCreatedBy(v FindBatchById200ResponseDevicesInner) {
	o.CreatedBy = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *VrfIpReservation) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *VrfIpReservation) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *VrfIpReservation) SetDetails(v string) {
	o.Details = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *VrfIpReservation) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *VrfIpReservation) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *VrfIpReservation) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VrfIpReservation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VrfIpReservation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VrfIpReservation) SetId(v string) {
	o.Id = &v
}

// GetMetalGateway returns the MetalGateway field value if set, zero value otherwise.
func (o *VrfIpReservation) GetMetalGateway() FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway {
	if o == nil || o.MetalGateway == nil {
		var ret FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway
		return ret
	}
	return *o.MetalGateway
}

// GetMetalGatewayOk returns a tuple with the MetalGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetMetalGatewayOk() (*FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway, bool) {
	if o == nil || o.MetalGateway == nil {
		return nil, false
	}
	return o.MetalGateway, true
}

// HasMetalGateway returns a boolean if a field has been set.
func (o *VrfIpReservation) HasMetalGateway() bool {
	if o != nil && o.MetalGateway != nil {
		return true
	}

	return false
}

// SetMetalGateway gets a reference to the given FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway and assigns it to the MetalGateway field.
func (o *VrfIpReservation) SetMetalGateway(v FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) {
	o.MetalGateway = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *VrfIpReservation) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *VrfIpReservation) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *VrfIpReservation) SetNetmask(v string) {
	o.Netmask = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *VrfIpReservation) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *VrfIpReservation) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *VrfIpReservation) SetNetwork(v string) {
	o.Network = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *VrfIpReservation) GetProject() MoveHardwareReservation201ResponseProject {
	if o == nil || o.Project == nil {
		var ret MoveHardwareReservation201ResponseProject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetProjectOk() (*MoveHardwareReservation201ResponseProject, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *VrfIpReservation) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given MoveHardwareReservation201ResponseProject and assigns it to the Project field.
func (o *VrfIpReservation) SetProject(v MoveHardwareReservation201ResponseProject) {
	o.Project = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VrfIpReservation) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VrfIpReservation) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VrfIpReservation) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VrfIpReservation) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VrfIpReservation) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VrfIpReservation) SetType(v string) {
	o.Type = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *VrfIpReservation) GetVrf() FindVrfs200ResponseVrfsInner {
	if o == nil || o.Vrf == nil {
		var ret FindVrfs200ResponseVrfsInner
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservation) GetVrfOk() (*FindVrfs200ResponseVrfsInner, bool) {
	if o == nil || o.Vrf == nil {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *VrfIpReservation) HasVrf() bool {
	if o != nil && o.Vrf != nil {
		return true
	}

	return false
}

// SetVrf gets a reference to the given FindVrfs200ResponseVrfsInner and assigns it to the Vrf field.
func (o *VrfIpReservation) SetVrf(v FindVrfs200ResponseVrfsInner) {
	o.Vrf = &v
}

func (o VrfIpReservation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.AddressFamily != nil {
		toSerialize["address_family"] = o.AddressFamily
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MetalGateway != nil {
		toSerialize["metal_gateway"] = o.MetalGateway
	}
	if o.Netmask != nil {
		toSerialize["netmask"] = o.Netmask
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Vrf != nil {
		toSerialize["vrf"] = o.Vrf
	}
	return json.Marshal(toSerialize)
}

type NullableVrfIpReservation struct {
	value *VrfIpReservation
	isSet bool
}

func (v NullableVrfIpReservation) Get() *VrfIpReservation {
	return v.value
}

func (v *NullableVrfIpReservation) Set(val *VrfIpReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfIpReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfIpReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfIpReservation(val *VrfIpReservation) *NullableVrfIpReservation {
	return &NullableVrfIpReservation{value: val, isSet: true}
}

func (v NullableVrfIpReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfIpReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
