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

// checks if the UserLite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserLite{}

// UserLite struct for UserLite
type UserLite struct {
	// Avatar thumbnail URL of the User
	AvatarThumbUrl *string `json:"avatar_thumb_url,omitempty"`
	// When the user was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Primary email address of the User
	Email *string `json:"email,omitempty"`
	// First name of the User
	FirstName *string `json:"first_name,omitempty"`
	// Full name of the User
	FullName *string `json:"full_name,omitempty"`
	// API URL uniquely representing the User
	Href *string `json:"href,omitempty"`
	// ID of the User
	Id string `json:"id"`
	// Last name of the User
	LastName *string `json:"last_name,omitempty"`
	// Short ID of the User
	ShortId string `json:"short_id"`
	// When the user details were last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewUserLite instantiates a new UserLite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLite(id string, shortId string) *UserLite {
	this := UserLite{}
	this.Id = id
	this.ShortId = shortId
	return &this
}

// NewUserLiteWithDefaults instantiates a new UserLite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLiteWithDefaults() *UserLite {
	this := UserLite{}
	return &this
}

// GetAvatarThumbUrl returns the AvatarThumbUrl field value if set, zero value otherwise.
func (o *UserLite) GetAvatarThumbUrl() string {
	if o == nil || isNil(o.AvatarThumbUrl) {
		var ret string
		return ret
	}
	return *o.AvatarThumbUrl
}

// GetAvatarThumbUrlOk returns a tuple with the AvatarThumbUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetAvatarThumbUrlOk() (*string, bool) {
	if o == nil || isNil(o.AvatarThumbUrl) {
		return nil, false
	}
	return o.AvatarThumbUrl, true
}

// HasAvatarThumbUrl returns a boolean if a field has been set.
func (o *UserLite) HasAvatarThumbUrl() bool {
	if o != nil && !isNil(o.AvatarThumbUrl) {
		return true
	}

	return false
}

// SetAvatarThumbUrl gets a reference to the given string and assigns it to the AvatarThumbUrl field.
func (o *UserLite) SetAvatarThumbUrl(v string) {
	o.AvatarThumbUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserLite) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserLite) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UserLite) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserLite) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserLite) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserLite) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserLite) GetFirstName() string {
	if o == nil || isNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetFirstNameOk() (*string, bool) {
	if o == nil || isNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserLite) HasFirstName() bool {
	if o != nil && !isNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserLite) SetFirstName(v string) {
	o.FirstName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *UserLite) GetFullName() string {
	if o == nil || isNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetFullNameOk() (*string, bool) {
	if o == nil || isNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *UserLite) HasFullName() bool {
	if o != nil && !isNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *UserLite) SetFullName(v string) {
	o.FullName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *UserLite) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *UserLite) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *UserLite) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value
func (o *UserLite) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserLite) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserLite) SetId(v string) {
	o.Id = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserLite) GetLastName() string {
	if o == nil || isNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetLastNameOk() (*string, bool) {
	if o == nil || isNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserLite) HasLastName() bool {
	if o != nil && !isNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserLite) SetLastName(v string) {
	o.LastName = &v
}

// GetShortId returns the ShortId field value
func (o *UserLite) GetShortId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortId
}

// GetShortIdOk returns a tuple with the ShortId field value
// and a boolean to check if the value has been set.
func (o *UserLite) GetShortIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortId, true
}

// SetShortId sets field value
func (o *UserLite) SetShortId(v string) {
	o.ShortId = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserLite) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserLite) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *UserLite) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o UserLite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AvatarThumbUrl) {
		toSerialize["avatar_thumb_url"] = o.AvatarThumbUrl
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !isNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	toSerialize["id"] = o.Id
	if !isNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	toSerialize["short_id"] = o.ShortId
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableUserLite struct {
	value *UserLite
	isSet bool
}

func (v NullableUserLite) Get() *UserLite {
	return v.value
}

func (v *NullableUserLite) Set(val *UserLite) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLite) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLite(val *UserLite) *NullableUserLite {
	return &NullableUserLite{value: val, isSet: true}
}

func (v NullableUserLite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
