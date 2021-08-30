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

// Project struct for Project
type Project struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	MaxDevices *map[string]interface{} `json:"max_devices,omitempty"`
	Members *[]Href `json:"members,omitempty"`
	Memberships *[]Href `json:"memberships,omitempty"`
	NetworkStatus *map[string]interface{} `json:"network_status,omitempty"`
	Invitations *[]Href `json:"invitations,omitempty"`
	PaymentMethod *Href `json:"payment_method,omitempty"`
	Devices *[]Href `json:"devices,omitempty"`
	SshKeys *[]Href `json:"ssh_keys,omitempty"`
	Volumes *[]Href `json:"volumes,omitempty"`
	BgpConfig *Href `json:"bgp_config,omitempty"`
	Customdata *map[string]interface{} `json:"customdata,omitempty"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject() *Project {
	this := Project{}
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Project) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Project) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Project) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Project) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Project) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Project) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Project) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Project) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Project) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Project) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Project) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Project) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetMaxDevices returns the MaxDevices field value if set, zero value otherwise.
func (o *Project) GetMaxDevices() map[string]interface{} {
	if o == nil || o.MaxDevices == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.MaxDevices
}

// GetMaxDevicesOk returns a tuple with the MaxDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetMaxDevicesOk() (*map[string]interface{}, bool) {
	if o == nil || o.MaxDevices == nil {
		return nil, false
	}
	return o.MaxDevices, true
}

// HasMaxDevices returns a boolean if a field has been set.
func (o *Project) HasMaxDevices() bool {
	if o != nil && o.MaxDevices != nil {
		return true
	}

	return false
}

// SetMaxDevices gets a reference to the given map[string]interface{} and assigns it to the MaxDevices field.
func (o *Project) SetMaxDevices(v map[string]interface{}) {
	o.MaxDevices = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *Project) GetMembers() []Href {
	if o == nil || o.Members == nil {
		var ret []Href
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetMembersOk() (*[]Href, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *Project) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []Href and assigns it to the Members field.
func (o *Project) SetMembers(v []Href) {
	o.Members = &v
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *Project) GetMemberships() []Href {
	if o == nil || o.Memberships == nil {
		var ret []Href
		return ret
	}
	return *o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetMembershipsOk() (*[]Href, bool) {
	if o == nil || o.Memberships == nil {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *Project) HasMemberships() bool {
	if o != nil && o.Memberships != nil {
		return true
	}

	return false
}

// SetMemberships gets a reference to the given []Href and assigns it to the Memberships field.
func (o *Project) SetMemberships(v []Href) {
	o.Memberships = &v
}

// GetNetworkStatus returns the NetworkStatus field value if set, zero value otherwise.
func (o *Project) GetNetworkStatus() map[string]interface{} {
	if o == nil || o.NetworkStatus == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.NetworkStatus
}

// GetNetworkStatusOk returns a tuple with the NetworkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetNetworkStatusOk() (*map[string]interface{}, bool) {
	if o == nil || o.NetworkStatus == nil {
		return nil, false
	}
	return o.NetworkStatus, true
}

// HasNetworkStatus returns a boolean if a field has been set.
func (o *Project) HasNetworkStatus() bool {
	if o != nil && o.NetworkStatus != nil {
		return true
	}

	return false
}

// SetNetworkStatus gets a reference to the given map[string]interface{} and assigns it to the NetworkStatus field.
func (o *Project) SetNetworkStatus(v map[string]interface{}) {
	o.NetworkStatus = &v
}

// GetInvitations returns the Invitations field value if set, zero value otherwise.
func (o *Project) GetInvitations() []Href {
	if o == nil || o.Invitations == nil {
		var ret []Href
		return ret
	}
	return *o.Invitations
}

// GetInvitationsOk returns a tuple with the Invitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetInvitationsOk() (*[]Href, bool) {
	if o == nil || o.Invitations == nil {
		return nil, false
	}
	return o.Invitations, true
}

// HasInvitations returns a boolean if a field has been set.
func (o *Project) HasInvitations() bool {
	if o != nil && o.Invitations != nil {
		return true
	}

	return false
}

// SetInvitations gets a reference to the given []Href and assigns it to the Invitations field.
func (o *Project) SetInvitations(v []Href) {
	o.Invitations = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *Project) GetPaymentMethod() Href {
	if o == nil || o.PaymentMethod == nil {
		var ret Href
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetPaymentMethodOk() (*Href, bool) {
	if o == nil || o.PaymentMethod == nil {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *Project) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod != nil {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given Href and assigns it to the PaymentMethod field.
func (o *Project) SetPaymentMethod(v Href) {
	o.PaymentMethod = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *Project) GetDevices() []Href {
	if o == nil || o.Devices == nil {
		var ret []Href
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetDevicesOk() (*[]Href, bool) {
	if o == nil || o.Devices == nil {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *Project) HasDevices() bool {
	if o != nil && o.Devices != nil {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []Href and assigns it to the Devices field.
func (o *Project) SetDevices(v []Href) {
	o.Devices = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *Project) GetSshKeys() []Href {
	if o == nil || o.SshKeys == nil {
		var ret []Href
		return ret
	}
	return *o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSshKeysOk() (*[]Href, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *Project) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []Href and assigns it to the SshKeys field.
func (o *Project) SetSshKeys(v []Href) {
	o.SshKeys = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *Project) GetVolumes() []Href {
	if o == nil || o.Volumes == nil {
		var ret []Href
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetVolumesOk() (*[]Href, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *Project) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Href and assigns it to the Volumes field.
func (o *Project) SetVolumes(v []Href) {
	o.Volumes = &v
}

// GetBgpConfig returns the BgpConfig field value if set, zero value otherwise.
func (o *Project) GetBgpConfig() Href {
	if o == nil || o.BgpConfig == nil {
		var ret Href
		return ret
	}
	return *o.BgpConfig
}

// GetBgpConfigOk returns a tuple with the BgpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetBgpConfigOk() (*Href, bool) {
	if o == nil || o.BgpConfig == nil {
		return nil, false
	}
	return o.BgpConfig, true
}

// HasBgpConfig returns a boolean if a field has been set.
func (o *Project) HasBgpConfig() bool {
	if o != nil && o.BgpConfig != nil {
		return true
	}

	return false
}

// SetBgpConfig gets a reference to the given Href and assigns it to the BgpConfig field.
func (o *Project) SetBgpConfig(v Href) {
	o.BgpConfig = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *Project) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCustomdataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *Project) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *Project) SetCustomdata(v map[string]interface{}) {
	o.Customdata = &v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.MaxDevices != nil {
		toSerialize["max_devices"] = o.MaxDevices
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Memberships != nil {
		toSerialize["memberships"] = o.Memberships
	}
	if o.NetworkStatus != nil {
		toSerialize["network_status"] = o.NetworkStatus
	}
	if o.Invitations != nil {
		toSerialize["invitations"] = o.Invitations
	}
	if o.PaymentMethod != nil {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if o.Devices != nil {
		toSerialize["devices"] = o.Devices
	}
	if o.SshKeys != nil {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	if o.BgpConfig != nil {
		toSerialize["bgp_config"] = o.BgpConfig
	}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	return json.Marshal(toSerialize)
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


