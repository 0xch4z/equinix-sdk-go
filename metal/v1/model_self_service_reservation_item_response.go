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

// SelfServiceReservationItemResponse struct for SelfServiceReservationItemResponse
type SelfServiceReservationItemResponse struct {
	Amount    *float32 `json:"amount,omitempty"`
	Id        *string  `json:"id,omitempty"`
	MetroCode *string  `json:"metro_code,omitempty"`
	MetroId   *string  `json:"metro_id,omitempty"`
	MetroName *string  `json:"metro_name,omitempty"`
	PlanId    *string  `json:"plan_id,omitempty"`
	PlanName  *string  `json:"plan_name,omitempty"`
	PlanSlug  *string  `json:"plan_slug,omitempty"`
	Quantity  *int32   `json:"quantity,omitempty"`
	Term      *string  `json:"term,omitempty"`
}

// NewSelfServiceReservationItemResponse instantiates a new SelfServiceReservationItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceReservationItemResponse() *SelfServiceReservationItemResponse {
	this := SelfServiceReservationItemResponse{}
	return &this
}

// NewSelfServiceReservationItemResponseWithDefaults instantiates a new SelfServiceReservationItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceReservationItemResponseWithDefaults() *SelfServiceReservationItemResponse {
	this := SelfServiceReservationItemResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *SelfServiceReservationItemResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SelfServiceReservationItemResponse) SetId(v string) {
	o.Id = &v
}

// GetMetroCode returns the MetroCode field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetMetroCode() string {
	if o == nil || o.MetroCode == nil {
		var ret string
		return ret
	}
	return *o.MetroCode
}

// GetMetroCodeOk returns a tuple with the MetroCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetMetroCodeOk() (*string, bool) {
	if o == nil || o.MetroCode == nil {
		return nil, false
	}
	return o.MetroCode, true
}

// HasMetroCode returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasMetroCode() bool {
	if o != nil && o.MetroCode != nil {
		return true
	}

	return false
}

// SetMetroCode gets a reference to the given string and assigns it to the MetroCode field.
func (o *SelfServiceReservationItemResponse) SetMetroCode(v string) {
	o.MetroCode = &v
}

// GetMetroId returns the MetroId field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetMetroId() string {
	if o == nil || o.MetroId == nil {
		var ret string
		return ret
	}
	return *o.MetroId
}

// GetMetroIdOk returns a tuple with the MetroId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetMetroIdOk() (*string, bool) {
	if o == nil || o.MetroId == nil {
		return nil, false
	}
	return o.MetroId, true
}

// HasMetroId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasMetroId() bool {
	if o != nil && o.MetroId != nil {
		return true
	}

	return false
}

// SetMetroId gets a reference to the given string and assigns it to the MetroId field.
func (o *SelfServiceReservationItemResponse) SetMetroId(v string) {
	o.MetroId = &v
}

// GetMetroName returns the MetroName field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetMetroName() string {
	if o == nil || o.MetroName == nil {
		var ret string
		return ret
	}
	return *o.MetroName
}

// GetMetroNameOk returns a tuple with the MetroName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetMetroNameOk() (*string, bool) {
	if o == nil || o.MetroName == nil {
		return nil, false
	}
	return o.MetroName, true
}

// HasMetroName returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasMetroName() bool {
	if o != nil && o.MetroName != nil {
		return true
	}

	return false
}

// SetMetroName gets a reference to the given string and assigns it to the MetroName field.
func (o *SelfServiceReservationItemResponse) SetMetroName(v string) {
	o.MetroName = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetPlanId() string {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetPlanIdOk() (*string, bool) {
	if o == nil || o.PlanId == nil {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasPlanId() bool {
	if o != nil && o.PlanId != nil {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SelfServiceReservationItemResponse) SetPlanId(v string) {
	o.PlanId = &v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetPlanName() string {
	if o == nil || o.PlanName == nil {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetPlanNameOk() (*string, bool) {
	if o == nil || o.PlanName == nil {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasPlanName() bool {
	if o != nil && o.PlanName != nil {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *SelfServiceReservationItemResponse) SetPlanName(v string) {
	o.PlanName = &v
}

// GetPlanSlug returns the PlanSlug field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetPlanSlug() string {
	if o == nil || o.PlanSlug == nil {
		var ret string
		return ret
	}
	return *o.PlanSlug
}

// GetPlanSlugOk returns a tuple with the PlanSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetPlanSlugOk() (*string, bool) {
	if o == nil || o.PlanSlug == nil {
		return nil, false
	}
	return o.PlanSlug, true
}

// HasPlanSlug returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasPlanSlug() bool {
	if o != nil && o.PlanSlug != nil {
		return true
	}

	return false
}

// SetPlanSlug gets a reference to the given string and assigns it to the PlanSlug field.
func (o *SelfServiceReservationItemResponse) SetPlanSlug(v string) {
	o.PlanSlug = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *SelfServiceReservationItemResponse) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *SelfServiceReservationItemResponse) GetTerm() string {
	if o == nil || o.Term == nil {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemResponse) GetTermOk() (*string, bool) {
	if o == nil || o.Term == nil {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *SelfServiceReservationItemResponse) HasTerm() bool {
	if o != nil && o.Term != nil {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *SelfServiceReservationItemResponse) SetTerm(v string) {
	o.Term = &v
}

func (o SelfServiceReservationItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MetroCode != nil {
		toSerialize["metro_code"] = o.MetroCode
	}
	if o.MetroId != nil {
		toSerialize["metro_id"] = o.MetroId
	}
	if o.MetroName != nil {
		toSerialize["metro_name"] = o.MetroName
	}
	if o.PlanId != nil {
		toSerialize["plan_id"] = o.PlanId
	}
	if o.PlanName != nil {
		toSerialize["plan_name"] = o.PlanName
	}
	if o.PlanSlug != nil {
		toSerialize["plan_slug"] = o.PlanSlug
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Term != nil {
		toSerialize["term"] = o.Term
	}
	return json.Marshal(toSerialize)
}

type NullableSelfServiceReservationItemResponse struct {
	value *SelfServiceReservationItemResponse
	isSet bool
}

func (v NullableSelfServiceReservationItemResponse) Get() *SelfServiceReservationItemResponse {
	return v.value
}

func (v *NullableSelfServiceReservationItemResponse) Set(val *SelfServiceReservationItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceReservationItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceReservationItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceReservationItemResponse(val *SelfServiceReservationItemResponse) *NullableSelfServiceReservationItemResponse {
	return &NullableSelfServiceReservationItemResponse{value: val, isSet: true}
}

func (v NullableSelfServiceReservationItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceReservationItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
