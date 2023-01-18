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
)

// checks if the PlanSpecs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanSpecs{}

// PlanSpecs struct for PlanSpecs
type PlanSpecs struct {
	Cpus     []PlanSpecsCpusInner   `json:"cpus,omitempty"`
	Drives   []PlanSpecsDrivesInner `json:"drives,omitempty"`
	Nics     []PlanSpecsNicsInner   `json:"nics,omitempty"`
	Features *PlanSpecsFeatures     `json:"features,omitempty"`
}

// NewPlanSpecs instantiates a new PlanSpecs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanSpecs() *PlanSpecs {
	this := PlanSpecs{}
	return &this
}

// NewPlanSpecsWithDefaults instantiates a new PlanSpecs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanSpecsWithDefaults() *PlanSpecs {
	this := PlanSpecs{}
	return &this
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *PlanSpecs) GetCpus() []PlanSpecsCpusInner {
	if o == nil || isNil(o.Cpus) {
		var ret []PlanSpecsCpusInner
		return ret
	}
	return o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanSpecs) GetCpusOk() ([]PlanSpecsCpusInner, bool) {
	if o == nil || isNil(o.Cpus) {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *PlanSpecs) HasCpus() bool {
	if o != nil && !isNil(o.Cpus) {
		return true
	}

	return false
}

// SetCpus gets a reference to the given []PlanSpecsCpusInner and assigns it to the Cpus field.
func (o *PlanSpecs) SetCpus(v []PlanSpecsCpusInner) {
	o.Cpus = v
}

// GetDrives returns the Drives field value if set, zero value otherwise.
func (o *PlanSpecs) GetDrives() []PlanSpecsDrivesInner {
	if o == nil || isNil(o.Drives) {
		var ret []PlanSpecsDrivesInner
		return ret
	}
	return o.Drives
}

// GetDrivesOk returns a tuple with the Drives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanSpecs) GetDrivesOk() ([]PlanSpecsDrivesInner, bool) {
	if o == nil || isNil(o.Drives) {
		return nil, false
	}
	return o.Drives, true
}

// HasDrives returns a boolean if a field has been set.
func (o *PlanSpecs) HasDrives() bool {
	if o != nil && !isNil(o.Drives) {
		return true
	}

	return false
}

// SetDrives gets a reference to the given []PlanSpecsDrivesInner and assigns it to the Drives field.
func (o *PlanSpecs) SetDrives(v []PlanSpecsDrivesInner) {
	o.Drives = v
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *PlanSpecs) GetNics() []PlanSpecsNicsInner {
	if o == nil || isNil(o.Nics) {
		var ret []PlanSpecsNicsInner
		return ret
	}
	return o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanSpecs) GetNicsOk() ([]PlanSpecsNicsInner, bool) {
	if o == nil || isNil(o.Nics) {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *PlanSpecs) HasNics() bool {
	if o != nil && !isNil(o.Nics) {
		return true
	}

	return false
}

// SetNics gets a reference to the given []PlanSpecsNicsInner and assigns it to the Nics field.
func (o *PlanSpecs) SetNics(v []PlanSpecsNicsInner) {
	o.Nics = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *PlanSpecs) GetFeatures() PlanSpecsFeatures {
	if o == nil || isNil(o.Features) {
		var ret PlanSpecsFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanSpecs) GetFeaturesOk() (*PlanSpecsFeatures, bool) {
	if o == nil || isNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *PlanSpecs) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given PlanSpecsFeatures and assigns it to the Features field.
func (o *PlanSpecs) SetFeatures(v PlanSpecsFeatures) {
	o.Features = &v
}

func (o PlanSpecs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanSpecs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Cpus) {
		toSerialize["cpus"] = o.Cpus
	}
	if !isNil(o.Drives) {
		toSerialize["drives"] = o.Drives
	}
	if !isNil(o.Nics) {
		toSerialize["nics"] = o.Nics
	}
	if !isNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	return toSerialize, nil
}

type NullablePlanSpecs struct {
	value *PlanSpecs
	isSet bool
}

func (v NullablePlanSpecs) Get() *PlanSpecs {
	return v.value
}

func (v *NullablePlanSpecs) Set(val *PlanSpecs) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanSpecs) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanSpecs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanSpecs(val *PlanSpecs) *NullablePlanSpecs {
	return &NullablePlanSpecs{value: val, isSet: true}
}

func (v NullablePlanSpecs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanSpecs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
