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
)

// InvitationList struct for InvitationList
type InvitationList struct {
	Invitations []FindInvitations200ResponseInvitationsInner `json:"invitations,omitempty"`
}

// NewInvitationList instantiates a new InvitationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationList() *InvitationList {
	this := InvitationList{}
	return &this
}

// NewInvitationListWithDefaults instantiates a new InvitationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationListWithDefaults() *InvitationList {
	this := InvitationList{}
	return &this
}

// GetInvitations returns the Invitations field value if set, zero value otherwise.
func (o *InvitationList) GetInvitations() []FindInvitations200ResponseInvitationsInner {
	if o == nil || o.Invitations == nil {
		var ret []FindInvitations200ResponseInvitationsInner
		return ret
	}
	return o.Invitations
}

// GetInvitationsOk returns a tuple with the Invitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationList) GetInvitationsOk() ([]FindInvitations200ResponseInvitationsInner, bool) {
	if o == nil || o.Invitations == nil {
		return nil, false
	}
	return o.Invitations, true
}

// HasInvitations returns a boolean if a field has been set.
func (o *InvitationList) HasInvitations() bool {
	if o != nil && o.Invitations != nil {
		return true
	}

	return false
}

// SetInvitations gets a reference to the given []FindInvitations200ResponseInvitationsInner and assigns it to the Invitations field.
func (o *InvitationList) SetInvitations(v []FindInvitations200ResponseInvitationsInner) {
	o.Invitations = v
}

func (o InvitationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Invitations != nil {
		toSerialize["invitations"] = o.Invitations
	}
	return json.Marshal(toSerialize)
}

type NullableInvitationList struct {
	value *InvitationList
	isSet bool
}

func (v NullableInvitationList) Get() *InvitationList {
	return v.value
}

func (v *NullableInvitationList) Set(val *InvitationList) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationList) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationList(val *InvitationList) *NullableInvitationList {
	return &NullableInvitationList{value: val, isSet: true}
}

func (v NullableInvitationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
