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

// PaymentMethod struct for PaymentMethod
type PaymentMethod struct {
	BillingAddress  *FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress `json:"billing_address,omitempty"`
	CardType        *string                                                                     `json:"card_type,omitempty"`
	CardholderName  *string                                                                     `json:"cardholder_name,omitempty"`
	CreatedAt       *time.Time                                                                  `json:"created_at,omitempty"`
	CreatedByUser   *FindBatchById200ResponseDevicesInner                                       `json:"created_by_user,omitempty"`
	Default         *bool                                                                       `json:"default,omitempty"`
	Email           *string                                                                     `json:"email,omitempty"`
	ExpirationMonth *string                                                                     `json:"expiration_month,omitempty"`
	ExpirationYear  *string                                                                     `json:"expiration_year,omitempty"`
	Id              *string                                                                     `json:"id,omitempty"`
	Name            *string                                                                     `json:"name,omitempty"`
	Organization    *FindBatchById200ResponseDevicesInner                                       `json:"organization,omitempty"`
	Projects        []FindBatchById200ResponseDevicesInner                                      `json:"projects,omitempty"`
	Type            *string                                                                     `json:"type,omitempty"`
	UpdatedAt       *time.Time                                                                  `json:"updated_at,omitempty"`
}

// NewPaymentMethod instantiates a new PaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethod() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodWithDefaults() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *PaymentMethod) GetBillingAddress() FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress {
	if o == nil || o.BillingAddress == nil {
		var ret FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetBillingAddressOk() (*FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress, bool) {
	if o == nil || o.BillingAddress == nil {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *PaymentMethod) HasBillingAddress() bool {
	if o != nil && o.BillingAddress != nil {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress and assigns it to the BillingAddress field.
func (o *PaymentMethod) SetBillingAddress(v FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress) {
	o.BillingAddress = &v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *PaymentMethod) GetCardType() string {
	if o == nil || o.CardType == nil {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardTypeOk() (*string, bool) {
	if o == nil || o.CardType == nil {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *PaymentMethod) HasCardType() bool {
	if o != nil && o.CardType != nil {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *PaymentMethod) SetCardType(v string) {
	o.CardType = &v
}

// GetCardholderName returns the CardholderName field value if set, zero value otherwise.
func (o *PaymentMethod) GetCardholderName() string {
	if o == nil || o.CardholderName == nil {
		var ret string
		return ret
	}
	return *o.CardholderName
}

// GetCardholderNameOk returns a tuple with the CardholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardholderNameOk() (*string, bool) {
	if o == nil || o.CardholderName == nil {
		return nil, false
	}
	return o.CardholderName, true
}

// HasCardholderName returns a boolean if a field has been set.
func (o *PaymentMethod) HasCardholderName() bool {
	if o != nil && o.CardholderName != nil {
		return true
	}

	return false
}

// SetCardholderName gets a reference to the given string and assigns it to the CardholderName field.
func (o *PaymentMethod) SetCardholderName(v string) {
	o.CardholderName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PaymentMethod) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PaymentMethod) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *PaymentMethod) GetCreatedByUser() FindBatchById200ResponseDevicesInner {
	if o == nil || o.CreatedByUser == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCreatedByUserOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *PaymentMethod) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the CreatedByUser field.
func (o *PaymentMethod) SetCreatedByUser(v FindBatchById200ResponseDevicesInner) {
	o.CreatedByUser = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PaymentMethod) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PaymentMethod) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *PaymentMethod) SetDefault(v bool) {
	o.Default = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PaymentMethod) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PaymentMethod) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PaymentMethod) SetEmail(v string) {
	o.Email = &v
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise.
func (o *PaymentMethod) GetExpirationMonth() string {
	if o == nil || o.ExpirationMonth == nil {
		var ret string
		return ret
	}
	return *o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetExpirationMonthOk() (*string, bool) {
	if o == nil || o.ExpirationMonth == nil {
		return nil, false
	}
	return o.ExpirationMonth, true
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *PaymentMethod) HasExpirationMonth() bool {
	if o != nil && o.ExpirationMonth != nil {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given string and assigns it to the ExpirationMonth field.
func (o *PaymentMethod) SetExpirationMonth(v string) {
	o.ExpirationMonth = &v
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise.
func (o *PaymentMethod) GetExpirationYear() string {
	if o == nil || o.ExpirationYear == nil {
		var ret string
		return ret
	}
	return *o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetExpirationYearOk() (*string, bool) {
	if o == nil || o.ExpirationYear == nil {
		return nil, false
	}
	return o.ExpirationYear, true
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *PaymentMethod) HasExpirationYear() bool {
	if o != nil && o.ExpirationYear != nil {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given string and assigns it to the ExpirationYear field.
func (o *PaymentMethod) SetExpirationYear(v string) {
	o.ExpirationYear = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentMethod) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentMethod) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentMethod) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentMethod) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentMethod) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentMethod) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *PaymentMethod) GetOrganization() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Organization == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetOrganizationOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PaymentMethod) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Organization field.
func (o *PaymentMethod) SetOrganization(v FindBatchById200ResponseDevicesInner) {
	o.Organization = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *PaymentMethod) GetProjects() []FindBatchById200ResponseDevicesInner {
	if o == nil || o.Projects == nil {
		var ret []FindBatchById200ResponseDevicesInner
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetProjectsOk() ([]FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PaymentMethod) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []FindBatchById200ResponseDevicesInner and assigns it to the Projects field.
func (o *PaymentMethod) SetProjects(v []FindBatchById200ResponseDevicesInner) {
	o.Projects = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethod) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethod) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethod) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PaymentMethod) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PaymentMethod) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingAddress != nil {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if o.CardType != nil {
		toSerialize["card_type"] = o.CardType
	}
	if o.CardholderName != nil {
		toSerialize["cardholder_name"] = o.CardholderName
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedByUser != nil {
		toSerialize["created_by_user"] = o.CreatedByUser
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.ExpirationMonth != nil {
		toSerialize["expiration_month"] = o.ExpirationMonth
	}
	if o.ExpirationYear != nil {
		toSerialize["expiration_year"] = o.ExpirationYear
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethod struct {
	value *PaymentMethod
	isSet bool
}

func (v NullablePaymentMethod) Get() *PaymentMethod {
	return v.value
}

func (v *NullablePaymentMethod) Set(val *PaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethod(val *PaymentMethod) *NullablePaymentMethod {
	return &NullablePaymentMethod{value: val, isSet: true}
}

func (v NullablePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
