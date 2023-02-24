/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. The fields available for search differ by resource, as does the search strategy.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// checks if the Project type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Project{}

// Project struct for Project
type Project struct {
	BgpConfig     *Href                  `json:"bgp_config,omitempty"`
	CreatedAt     *time.Time             `json:"created_at,omitempty"`
	Customdata    map[string]interface{} `json:"customdata,omitempty"`
	Devices       []Href                 `json:"devices,omitempty"`
	Href          *string                `json:"href,omitempty"`
	Id            *string                `json:"id,omitempty"`
	Invitations   []Href                 `json:"invitations,omitempty"`
	MaxDevices    map[string]interface{} `json:"max_devices,omitempty"`
	Members       []Href                 `json:"members,omitempty"`
	Memberships   []Href                 `json:"memberships,omitempty"`
	Name          *string                `json:"name,omitempty"`
	NetworkStatus map[string]interface{} `json:"network_status,omitempty"`
	PaymentMethod *Href                  `json:"payment_method,omitempty"`
	SshKeys       []Href                 `json:"ssh_keys,omitempty"`
	UpdatedAt     *time.Time             `json:"updated_at,omitempty"`
	Volumes       []Href                 `json:"volumes,omitempty"`
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

// GetBgpConfig returns the BgpConfig field value if set, zero value otherwise.
func (o *Project) GetBgpConfig() Href {
	if o == nil || isNil(o.BgpConfig) {
		var ret Href
		return ret
	}
	return *o.BgpConfig
}

// GetBgpConfigOk returns a tuple with the BgpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetBgpConfigOk() (*Href, bool) {
	if o == nil || isNil(o.BgpConfig) {
		return nil, false
	}
	return o.BgpConfig, true
}

// HasBgpConfig returns a boolean if a field has been set.
func (o *Project) HasBgpConfig() bool {
	if o != nil && !isNil(o.BgpConfig) {
		return true
	}

	return false
}

// SetBgpConfig gets a reference to the given Href and assigns it to the BgpConfig field.
func (o *Project) SetBgpConfig(v Href) {
	o.BgpConfig = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Project) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Project) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Project) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *Project) GetCustomdata() map[string]interface{} {
	if o == nil || isNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *Project) HasCustomdata() bool {
	if o != nil && !isNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *Project) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *Project) GetDevices() []Href {
	if o == nil || isNil(o.Devices) {
		var ret []Href
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetDevicesOk() ([]Href, bool) {
	if o == nil || isNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *Project) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []Href and assigns it to the Devices field.
func (o *Project) SetDevices(v []Href) {
	o.Devices = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Project) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Project) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Project) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Project) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Project) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Project) SetId(v string) {
	o.Id = &v
}

// GetInvitations returns the Invitations field value if set, zero value otherwise.
func (o *Project) GetInvitations() []Href {
	if o == nil || isNil(o.Invitations) {
		var ret []Href
		return ret
	}
	return o.Invitations
}

// GetInvitationsOk returns a tuple with the Invitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetInvitationsOk() ([]Href, bool) {
	if o == nil || isNil(o.Invitations) {
		return nil, false
	}
	return o.Invitations, true
}

// HasInvitations returns a boolean if a field has been set.
func (o *Project) HasInvitations() bool {
	if o != nil && !isNil(o.Invitations) {
		return true
	}

	return false
}

// SetInvitations gets a reference to the given []Href and assigns it to the Invitations field.
func (o *Project) SetInvitations(v []Href) {
	o.Invitations = v
}

// GetMaxDevices returns the MaxDevices field value if set, zero value otherwise.
func (o *Project) GetMaxDevices() map[string]interface{} {
	if o == nil || isNil(o.MaxDevices) {
		var ret map[string]interface{}
		return ret
	}
	return o.MaxDevices
}

// GetMaxDevicesOk returns a tuple with the MaxDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetMaxDevicesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.MaxDevices) {
		return map[string]interface{}{}, false
	}
	return o.MaxDevices, true
}

// HasMaxDevices returns a boolean if a field has been set.
func (o *Project) HasMaxDevices() bool {
	if o != nil && !isNil(o.MaxDevices) {
		return true
	}

	return false
}

// SetMaxDevices gets a reference to the given map[string]interface{} and assigns it to the MaxDevices field.
func (o *Project) SetMaxDevices(v map[string]interface{}) {
	o.MaxDevices = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *Project) GetMembers() []Href {
	if o == nil || isNil(o.Members) {
		var ret []Href
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetMembersOk() ([]Href, bool) {
	if o == nil || isNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *Project) HasMembers() bool {
	if o != nil && !isNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []Href and assigns it to the Members field.
func (o *Project) SetMembers(v []Href) {
	o.Members = v
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *Project) GetMemberships() []Href {
	if o == nil || isNil(o.Memberships) {
		var ret []Href
		return ret
	}
	return o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetMembershipsOk() ([]Href, bool) {
	if o == nil || isNil(o.Memberships) {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *Project) HasMemberships() bool {
	if o != nil && !isNil(o.Memberships) {
		return true
	}

	return false
}

// SetMemberships gets a reference to the given []Href and assigns it to the Memberships field.
func (o *Project) SetMemberships(v []Href) {
	o.Memberships = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Project) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Project) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Project) SetName(v string) {
	o.Name = &v
}

// GetNetworkStatus returns the NetworkStatus field value if set, zero value otherwise.
func (o *Project) GetNetworkStatus() map[string]interface{} {
	if o == nil || isNil(o.NetworkStatus) {
		var ret map[string]interface{}
		return ret
	}
	return o.NetworkStatus
}

// GetNetworkStatusOk returns a tuple with the NetworkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetNetworkStatusOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.NetworkStatus) {
		return map[string]interface{}{}, false
	}
	return o.NetworkStatus, true
}

// HasNetworkStatus returns a boolean if a field has been set.
func (o *Project) HasNetworkStatus() bool {
	if o != nil && !isNil(o.NetworkStatus) {
		return true
	}

	return false
}

// SetNetworkStatus gets a reference to the given map[string]interface{} and assigns it to the NetworkStatus field.
func (o *Project) SetNetworkStatus(v map[string]interface{}) {
	o.NetworkStatus = v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *Project) GetPaymentMethod() Href {
	if o == nil || isNil(o.PaymentMethod) {
		var ret Href
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetPaymentMethodOk() (*Href, bool) {
	if o == nil || isNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *Project) HasPaymentMethod() bool {
	if o != nil && !isNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given Href and assigns it to the PaymentMethod field.
func (o *Project) SetPaymentMethod(v Href) {
	o.PaymentMethod = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *Project) GetSshKeys() []Href {
	if o == nil || isNil(o.SshKeys) {
		var ret []Href
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSshKeysOk() ([]Href, bool) {
	if o == nil || isNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *Project) HasSshKeys() bool {
	if o != nil && !isNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []Href and assigns it to the SshKeys field.
func (o *Project) SetSshKeys(v []Href) {
	o.SshKeys = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Project) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Project) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Project) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *Project) GetVolumes() []Href {
	if o == nil || isNil(o.Volumes) {
		var ret []Href
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetVolumesOk() ([]Href, bool) {
	if o == nil || isNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *Project) HasVolumes() bool {
	if o != nil && !isNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Href and assigns it to the Volumes field.
func (o *Project) SetVolumes(v []Href) {
	o.Volumes = v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Project) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BgpConfig) {
		toSerialize["bgp_config"] = o.BgpConfig
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Invitations) {
		toSerialize["invitations"] = o.Invitations
	}
	if !isNil(o.MaxDevices) {
		toSerialize["max_devices"] = o.MaxDevices
	}
	if !isNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !isNil(o.Memberships) {
		toSerialize["memberships"] = o.Memberships
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.NetworkStatus) {
		toSerialize["network_status"] = o.NetworkStatus
	}
	if !isNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !isNil(o.SshKeys) {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return toSerialize, nil
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
