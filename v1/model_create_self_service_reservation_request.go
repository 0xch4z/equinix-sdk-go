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

// CreateSelfServiceReservationRequest struct for CreateSelfServiceReservationRequest
type CreateSelfServiceReservationRequest struct {
	StartDate *time.Time `json:"start_date,omitempty"`
	Period *CreateSelfServiceReservationRequestPeriod `json:"period,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Item *[]SelfServiceReservationItemRequest `json:"item,omitempty"`
}

// NewCreateSelfServiceReservationRequest instantiates a new CreateSelfServiceReservationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSelfServiceReservationRequest() *CreateSelfServiceReservationRequest {
	this := CreateSelfServiceReservationRequest{}
	return &this
}

// NewCreateSelfServiceReservationRequestWithDefaults instantiates a new CreateSelfServiceReservationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSelfServiceReservationRequestWithDefaults() *CreateSelfServiceReservationRequest {
	this := CreateSelfServiceReservationRequest{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CreateSelfServiceReservationRequest) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSelfServiceReservationRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreateSelfServiceReservationRequest) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *CreateSelfServiceReservationRequest) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *CreateSelfServiceReservationRequest) GetPeriod() CreateSelfServiceReservationRequestPeriod {
	if o == nil || o.Period == nil {
		var ret CreateSelfServiceReservationRequestPeriod
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSelfServiceReservationRequest) GetPeriodOk() (*CreateSelfServiceReservationRequestPeriod, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *CreateSelfServiceReservationRequest) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given CreateSelfServiceReservationRequestPeriod and assigns it to the Period field.
func (o *CreateSelfServiceReservationRequest) SetPeriod(v CreateSelfServiceReservationRequestPeriod) {
	o.Period = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CreateSelfServiceReservationRequest) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSelfServiceReservationRequest) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CreateSelfServiceReservationRequest) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CreateSelfServiceReservationRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *CreateSelfServiceReservationRequest) GetItem() []SelfServiceReservationItemRequest {
	if o == nil || o.Item == nil {
		var ret []SelfServiceReservationItemRequest
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSelfServiceReservationRequest) GetItemOk() (*[]SelfServiceReservationItemRequest, bool) {
	if o == nil || o.Item == nil {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *CreateSelfServiceReservationRequest) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given []SelfServiceReservationItemRequest and assigns it to the Item field.
func (o *CreateSelfServiceReservationRequest) SetItem(v []SelfServiceReservationItemRequest) {
	o.Item = &v
}

func (o CreateSelfServiceReservationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Item != nil {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSelfServiceReservationRequest struct {
	value *CreateSelfServiceReservationRequest
	isSet bool
}

func (v NullableCreateSelfServiceReservationRequest) Get() *CreateSelfServiceReservationRequest {
	return v.value
}

func (v *NullableCreateSelfServiceReservationRequest) Set(val *CreateSelfServiceReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSelfServiceReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSelfServiceReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSelfServiceReservationRequest(val *CreateSelfServiceReservationRequest) *NullableCreateSelfServiceReservationRequest {
	return &NullableCreateSelfServiceReservationRequest{value: val, isSet: true}
}

func (v NullableCreateSelfServiceReservationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSelfServiceReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


