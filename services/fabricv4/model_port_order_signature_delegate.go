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

// checks if the PortOrderSignatureDelegate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortOrderSignatureDelegate{}

// PortOrderSignatureDelegate delegate oder details
type PortOrderSignatureDelegate struct {
	// name of delegate
	FirstName *string `json:"firstName,omitempty"`
	// last Name of delegate
	LastName *string `json:"lastName,omitempty"`
	// email of delegate
	Email                *string `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortOrderSignatureDelegate PortOrderSignatureDelegate

// NewPortOrderSignatureDelegate instantiates a new PortOrderSignatureDelegate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortOrderSignatureDelegate() *PortOrderSignatureDelegate {
	this := PortOrderSignatureDelegate{}
	return &this
}

// NewPortOrderSignatureDelegateWithDefaults instantiates a new PortOrderSignatureDelegate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortOrderSignatureDelegateWithDefaults() *PortOrderSignatureDelegate {
	this := PortOrderSignatureDelegate{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PortOrderSignatureDelegate) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortOrderSignatureDelegate) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PortOrderSignatureDelegate) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PortOrderSignatureDelegate) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PortOrderSignatureDelegate) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortOrderSignatureDelegate) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PortOrderSignatureDelegate) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PortOrderSignatureDelegate) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PortOrderSignatureDelegate) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortOrderSignatureDelegate) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PortOrderSignatureDelegate) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PortOrderSignatureDelegate) SetEmail(v string) {
	o.Email = &v
}

func (o PortOrderSignatureDelegate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortOrderSignatureDelegate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortOrderSignatureDelegate) UnmarshalJSON(data []byte) (err error) {
	varPortOrderSignatureDelegate := _PortOrderSignatureDelegate{}

	err = json.Unmarshal(data, &varPortOrderSignatureDelegate)

	if err != nil {
		return err
	}

	*o = PortOrderSignatureDelegate(varPortOrderSignatureDelegate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortOrderSignatureDelegate struct {
	value *PortOrderSignatureDelegate
	isSet bool
}

func (v NullablePortOrderSignatureDelegate) Get() *PortOrderSignatureDelegate {
	return v.value
}

func (v *NullablePortOrderSignatureDelegate) Set(val *PortOrderSignatureDelegate) {
	v.value = val
	v.isSet = true
}

func (v NullablePortOrderSignatureDelegate) IsSet() bool {
	return v.isSet
}

func (v *NullablePortOrderSignatureDelegate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortOrderSignatureDelegate(val *PortOrderSignatureDelegate) *NullablePortOrderSignatureDelegate {
	return &NullablePortOrderSignatureDelegate{value: val, isSet: true}
}

func (v NullablePortOrderSignatureDelegate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortOrderSignatureDelegate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}