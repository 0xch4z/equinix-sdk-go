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
)

// checks if the RouteFiltersSearchBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteFiltersSearchBase{}

// RouteFiltersSearchBase struct for RouteFiltersSearchBase
type RouteFiltersSearchBase struct {
	Filter               *RouteFiltersSearchBaseFilter `json:"filter,omitempty"`
	Pagination           *Pagination                   `json:"pagination,omitempty"`
	Sort                 []SortItem                    `json:"sort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteFiltersSearchBase RouteFiltersSearchBase

// NewRouteFiltersSearchBase instantiates a new RouteFiltersSearchBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteFiltersSearchBase() *RouteFiltersSearchBase {
	this := RouteFiltersSearchBase{}
	return &this
}

// NewRouteFiltersSearchBaseWithDefaults instantiates a new RouteFiltersSearchBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteFiltersSearchBaseWithDefaults() *RouteFiltersSearchBase {
	this := RouteFiltersSearchBase{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RouteFiltersSearchBase) GetFilter() RouteFiltersSearchBaseFilter {
	if o == nil || IsNil(o.Filter) {
		var ret RouteFiltersSearchBaseFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFiltersSearchBase) GetFilterOk() (*RouteFiltersSearchBaseFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RouteFiltersSearchBase) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given RouteFiltersSearchBaseFilter and assigns it to the Filter field.
func (o *RouteFiltersSearchBase) SetFilter(v RouteFiltersSearchBaseFilter) {
	o.Filter = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *RouteFiltersSearchBase) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFiltersSearchBase) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *RouteFiltersSearchBase) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *RouteFiltersSearchBase) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *RouteFiltersSearchBase) GetSort() []SortItem {
	if o == nil || IsNil(o.Sort) {
		var ret []SortItem
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFiltersSearchBase) GetSortOk() ([]SortItem, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *RouteFiltersSearchBase) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []SortItem and assigns it to the Sort field.
func (o *RouteFiltersSearchBase) SetSort(v []SortItem) {
	o.Sort = v
}

func (o RouteFiltersSearchBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteFiltersSearchBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteFiltersSearchBase) UnmarshalJSON(data []byte) (err error) {
	varRouteFiltersSearchBase := _RouteFiltersSearchBase{}

	err = json.Unmarshal(data, &varRouteFiltersSearchBase)

	if err != nil {
		return err
	}

	*o = RouteFiltersSearchBase(varRouteFiltersSearchBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "sort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteFiltersSearchBase struct {
	value *RouteFiltersSearchBase
	isSet bool
}

func (v NullableRouteFiltersSearchBase) Get() *RouteFiltersSearchBase {
	return v.value
}

func (v *NullableRouteFiltersSearchBase) Set(val *RouteFiltersSearchBase) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFiltersSearchBase) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFiltersSearchBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFiltersSearchBase(val *RouteFiltersSearchBase) *NullableRouteFiltersSearchBase {
	return &NullableRouteFiltersSearchBase{value: val, isSet: true}
}

func (v NullableRouteFiltersSearchBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFiltersSearchBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
