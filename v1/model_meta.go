/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// Meta struct for Meta
type Meta struct {
	First *Href `json:"first,omitempty"`
	Previous *Href `json:"previous,omitempty"`
	Self *Href `json:"self,omitempty"`
	Next *Href `json:"next,omitempty"`
	Last *Href `json:"last,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewMeta instantiates a new Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeta() *Meta {
	this := Meta{}
	return &this
}

// NewMetaWithDefaults instantiates a new Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaWithDefaults() *Meta {
	this := Meta{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *Meta) GetFirst() Href {
	if o == nil || o.First == nil {
		var ret Href
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetFirstOk() (*Href, bool) {
	if o == nil || o.First == nil {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *Meta) HasFirst() bool {
	if o != nil && o.First != nil {
		return true
	}

	return false
}

// SetFirst gets a reference to the given Href and assigns it to the First field.
func (o *Meta) SetFirst(v Href) {
	o.First = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *Meta) GetPrevious() Href {
	if o == nil || o.Previous == nil {
		var ret Href
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetPreviousOk() (*Href, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *Meta) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given Href and assigns it to the Previous field.
func (o *Meta) SetPrevious(v Href) {
	o.Previous = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *Meta) GetSelf() Href {
	if o == nil || o.Self == nil {
		var ret Href
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetSelfOk() (*Href, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *Meta) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Href and assigns it to the Self field.
func (o *Meta) SetSelf(v Href) {
	o.Self = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *Meta) GetNext() Href {
	if o == nil || o.Next == nil {
		var ret Href
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetNextOk() (*Href, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *Meta) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given Href and assigns it to the Next field.
func (o *Meta) SetNext(v Href) {
	o.Next = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *Meta) GetLast() Href {
	if o == nil || o.Last == nil {
		var ret Href
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetLastOk() (*Href, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *Meta) HasLast() bool {
	if o != nil && o.Last != nil {
		return true
	}

	return false
}

// SetLast gets a reference to the given Href and assigns it to the Last field.
func (o *Meta) SetLast(v Href) {
	o.Last = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Meta) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Meta) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *Meta) SetTotal(v int32) {
	o.Total = &v
}

func (o Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.First != nil {
		toSerialize["first"] = o.First
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableMeta struct {
	value *Meta
	isSet bool
}

func (v NullableMeta) Get() *Meta {
	return v.value
}

func (v *NullableMeta) Set(val *Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeta(val *Meta) *NullableMeta {
	return &NullableMeta{value: val, isSet: true}
}

func (v NullableMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


