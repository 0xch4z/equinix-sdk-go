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

// SSHKeyList struct for SSHKeyList
type SSHKeyList struct {
	SshKeys *[]SSHKey `json:"ssh_keys,omitempty"`
}

// NewSSHKeyList instantiates a new SSHKeyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSHKeyList() *SSHKeyList {
	this := SSHKeyList{}
	return &this
}

// NewSSHKeyListWithDefaults instantiates a new SSHKeyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHKeyListWithDefaults() *SSHKeyList {
	this := SSHKeyList{}
	return &this
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *SSHKeyList) GetSshKeys() []SSHKey {
	if o == nil || o.SshKeys == nil {
		var ret []SSHKey
		return ret
	}
	return *o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHKeyList) GetSshKeysOk() (*[]SSHKey, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *SSHKeyList) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []SSHKey and assigns it to the SshKeys field.
func (o *SSHKeyList) SetSshKeys(v []SSHKey) {
	o.SshKeys = &v
}

func (o SSHKeyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SshKeys != nil {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	return json.Marshal(toSerialize)
}

type NullableSSHKeyList struct {
	value *SSHKeyList
	isSet bool
}

func (v NullableSSHKeyList) Get() *SSHKeyList {
	return v.value
}

func (v *NullableSSHKeyList) Set(val *SSHKeyList) {
	v.value = val
	v.isSet = true
}

func (v NullableSSHKeyList) IsSet() bool {
	return v.isSet
}

func (v *NullableSSHKeyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSHKeyList(val *SSHKeyList) *NullableSSHKeyList {
	return &NullableSSHKeyList{value: val, isSet: true}
}

func (v NullableSSHKeyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSHKeyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


