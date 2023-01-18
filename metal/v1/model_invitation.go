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

// checks if the Invitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invitation{}

// Invitation struct for Invitation
type Invitation struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	Href         *string    `json:"href,omitempty"`
	Id           *string    `json:"id,omitempty"`
	Invitation   *Href      `json:"invitation,omitempty"`
	InvitedBy    *Href      `json:"invited_by,omitempty"`
	Invitee      *string    `json:"invitee,omitempty"`
	Nonce        *string    `json:"nonce,omitempty"`
	Organization *Href      `json:"organization,omitempty"`
	Projects     []Href     `json:"projects,omitempty"`
	Roles        []string   `json:"roles,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

// NewInvitation instantiates a new Invitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitation() *Invitation {
	this := Invitation{}
	return &this
}

// NewInvitationWithDefaults instantiates a new Invitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationWithDefaults() *Invitation {
	this := Invitation{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Invitation) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Invitation) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Invitation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Invitation) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Invitation) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Invitation) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Invitation) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Invitation) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Invitation) SetId(v string) {
	o.Id = &v
}

// GetInvitation returns the Invitation field value if set, zero value otherwise.
func (o *Invitation) GetInvitation() Href {
	if o == nil || isNil(o.Invitation) {
		var ret Href
		return ret
	}
	return *o.Invitation
}

// GetInvitationOk returns a tuple with the Invitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetInvitationOk() (*Href, bool) {
	if o == nil || isNil(o.Invitation) {
		return nil, false
	}
	return o.Invitation, true
}

// HasInvitation returns a boolean if a field has been set.
func (o *Invitation) HasInvitation() bool {
	if o != nil && !isNil(o.Invitation) {
		return true
	}

	return false
}

// SetInvitation gets a reference to the given Href and assigns it to the Invitation field.
func (o *Invitation) SetInvitation(v Href) {
	o.Invitation = &v
}

// GetInvitedBy returns the InvitedBy field value if set, zero value otherwise.
func (o *Invitation) GetInvitedBy() Href {
	if o == nil || isNil(o.InvitedBy) {
		var ret Href
		return ret
	}
	return *o.InvitedBy
}

// GetInvitedByOk returns a tuple with the InvitedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetInvitedByOk() (*Href, bool) {
	if o == nil || isNil(o.InvitedBy) {
		return nil, false
	}
	return o.InvitedBy, true
}

// HasInvitedBy returns a boolean if a field has been set.
func (o *Invitation) HasInvitedBy() bool {
	if o != nil && !isNil(o.InvitedBy) {
		return true
	}

	return false
}

// SetInvitedBy gets a reference to the given Href and assigns it to the InvitedBy field.
func (o *Invitation) SetInvitedBy(v Href) {
	o.InvitedBy = &v
}

// GetInvitee returns the Invitee field value if set, zero value otherwise.
func (o *Invitation) GetInvitee() string {
	if o == nil || isNil(o.Invitee) {
		var ret string
		return ret
	}
	return *o.Invitee
}

// GetInviteeOk returns a tuple with the Invitee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetInviteeOk() (*string, bool) {
	if o == nil || isNil(o.Invitee) {
		return nil, false
	}
	return o.Invitee, true
}

// HasInvitee returns a boolean if a field has been set.
func (o *Invitation) HasInvitee() bool {
	if o != nil && !isNil(o.Invitee) {
		return true
	}

	return false
}

// SetInvitee gets a reference to the given string and assigns it to the Invitee field.
func (o *Invitation) SetInvitee(v string) {
	o.Invitee = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *Invitation) GetNonce() string {
	if o == nil || isNil(o.Nonce) {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetNonceOk() (*string, bool) {
	if o == nil || isNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *Invitation) HasNonce() bool {
	if o != nil && !isNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *Invitation) SetNonce(v string) {
	o.Nonce = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Invitation) GetOrganization() Href {
	if o == nil || isNil(o.Organization) {
		var ret Href
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetOrganizationOk() (*Href, bool) {
	if o == nil || isNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Invitation) HasOrganization() bool {
	if o != nil && !isNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Href and assigns it to the Organization field.
func (o *Invitation) SetOrganization(v Href) {
	o.Organization = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *Invitation) GetProjects() []Href {
	if o == nil || isNil(o.Projects) {
		var ret []Href
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetProjectsOk() ([]Href, bool) {
	if o == nil || isNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *Invitation) HasProjects() bool {
	if o != nil && !isNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []Href and assigns it to the Projects field.
func (o *Invitation) SetProjects(v []Href) {
	o.Projects = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Invitation) GetRoles() []string {
	if o == nil || isNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetRolesOk() ([]string, bool) {
	if o == nil || isNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Invitation) HasRoles() bool {
	if o != nil && !isNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *Invitation) SetRoles(v []string) {
	o.Roles = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Invitation) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Invitation) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Invitation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Invitation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Invitation) {
		toSerialize["invitation"] = o.Invitation
	}
	if !isNil(o.InvitedBy) {
		toSerialize["invited_by"] = o.InvitedBy
	}
	if !isNil(o.Invitee) {
		toSerialize["invitee"] = o.Invitee
	}
	if !isNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !isNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !isNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !isNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableInvitation struct {
	value *Invitation
	isSet bool
}

func (v NullableInvitation) Get() *Invitation {
	return v.value
}

func (v *NullableInvitation) Set(val *Invitation) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitation(val *Invitation) *NullableInvitation {
	return &NullableInvitation{value: val, isSet: true}
}

func (v NullableInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
