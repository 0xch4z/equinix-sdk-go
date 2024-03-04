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
	"fmt"
	"time"
)

// checks if the ServiceToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceToken{}

// ServiceToken Create Service Tokens (v4) generates Equinix Fabric service tokens. These tokens authorize users to access protected resources and services. The tokens remove sensitive content from the environment, rather than just masking it, making the protected data impossible to unencrypt or decrypt. Resource owners can distribute the tokens to trusted partners and vendors, allowing selected third parties to work directly with Equinix network assets.
type ServiceToken struct {
	Type *ServiceTokenType `json:"type,omitempty"`
	// An absolute URL that is the subject of the link's context.
	Href *string `json:"href,omitempty"`
	// Equinix-assigned service token identifier
	Uuid string `json:"uuid"`
	// Customer-provided service token name
	Name *string `json:"name,omitempty"`
	// Customer-provided service token description
	Description *string `json:"description,omitempty"`
	// Expiration date and time of the service token.
	ExpirationDateTime *time.Time              `json:"expirationDateTime,omitempty"`
	Connection         *ServiceTokenConnection `json:"connection,omitempty"`
	State              *ServiceTokenState      `json:"state,omitempty"`
	// Service token related notifications
	Notifications        []SimplifiedNotification `json:"notifications,omitempty"`
	Account              *SimplifiedAccount       `json:"account,omitempty"`
	Changelog            *Changelog               `json:"changelog,omitempty"`
	Project              *Project                 `json:"project,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceToken ServiceToken

// NewServiceToken instantiates a new ServiceToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceToken(uuid string) *ServiceToken {
	this := ServiceToken{}
	this.Uuid = uuid
	return &this
}

// NewServiceTokenWithDefaults instantiates a new ServiceToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTokenWithDefaults() *ServiceToken {
	this := ServiceToken{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceToken) GetType() ServiceTokenType {
	if o == nil || IsNil(o.Type) {
		var ret ServiceTokenType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetTypeOk() (*ServiceTokenType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceToken) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ServiceTokenType and assigns it to the Type field.
func (o *ServiceToken) SetType(v ServiceTokenType) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ServiceToken) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ServiceToken) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ServiceToken) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value
func (o *ServiceToken) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ServiceToken) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceToken) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceToken) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceToken) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceToken) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceToken) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceToken) SetDescription(v string) {
	o.Description = &v
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise.
func (o *ServiceToken) GetExpirationDateTime() time.Time {
	if o == nil || IsNil(o.ExpirationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDateTime) {
		return nil, false
	}
	return o.ExpirationDateTime, true
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *ServiceToken) HasExpirationDateTime() bool {
	if o != nil && !IsNil(o.ExpirationDateTime) {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.
func (o *ServiceToken) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *ServiceToken) GetConnection() ServiceTokenConnection {
	if o == nil || IsNil(o.Connection) {
		var ret ServiceTokenConnection
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetConnectionOk() (*ServiceTokenConnection, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *ServiceToken) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ServiceTokenConnection and assigns it to the Connection field.
func (o *ServiceToken) SetConnection(v ServiceTokenConnection) {
	o.Connection = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ServiceToken) GetState() ServiceTokenState {
	if o == nil || IsNil(o.State) {
		var ret ServiceTokenState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetStateOk() (*ServiceTokenState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ServiceToken) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ServiceTokenState and assigns it to the State field.
func (o *ServiceToken) SetState(v ServiceTokenState) {
	o.State = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *ServiceToken) GetNotifications() []SimplifiedNotification {
	if o == nil || IsNil(o.Notifications) {
		var ret []SimplifiedNotification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetNotificationsOk() ([]SimplifiedNotification, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *ServiceToken) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []SimplifiedNotification and assigns it to the Notifications field.
func (o *ServiceToken) SetNotifications(v []SimplifiedNotification) {
	o.Notifications = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ServiceToken) GetAccount() SimplifiedAccount {
	if o == nil || IsNil(o.Account) {
		var ret SimplifiedAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetAccountOk() (*SimplifiedAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ServiceToken) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given SimplifiedAccount and assigns it to the Account field.
func (o *ServiceToken) SetAccount(v SimplifiedAccount) {
	o.Account = &v
}

// GetChangelog returns the Changelog field value if set, zero value otherwise.
func (o *ServiceToken) GetChangelog() Changelog {
	if o == nil || IsNil(o.Changelog) {
		var ret Changelog
		return ret
	}
	return *o.Changelog
}

// GetChangelogOk returns a tuple with the Changelog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetChangelogOk() (*Changelog, bool) {
	if o == nil || IsNil(o.Changelog) {
		return nil, false
	}
	return o.Changelog, true
}

// HasChangelog returns a boolean if a field has been set.
func (o *ServiceToken) HasChangelog() bool {
	if o != nil && !IsNil(o.Changelog) {
		return true
	}

	return false
}

// SetChangelog gets a reference to the given Changelog and assigns it to the Changelog field.
func (o *ServiceToken) SetChangelog(v Changelog) {
	o.Changelog = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ServiceToken) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToken) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ServiceToken) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *ServiceToken) SetProject(v Project) {
	o.Project = &v
}

func (o ServiceToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpirationDateTime) {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime
	}
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Changelog) {
		toSerialize["changelog"] = o.Changelog
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServiceToken := _ServiceToken{}

	err = json.Unmarshal(data, &varServiceToken)

	if err != nil {
		return err
	}

	*o = ServiceToken(varServiceToken)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "expirationDateTime")
		delete(additionalProperties, "connection")
		delete(additionalProperties, "state")
		delete(additionalProperties, "notifications")
		delete(additionalProperties, "account")
		delete(additionalProperties, "changelog")
		delete(additionalProperties, "project")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceToken struct {
	value *ServiceToken
	isSet bool
}

func (v NullableServiceToken) Get() *ServiceToken {
	return v.value
}

func (v *NullableServiceToken) Set(val *ServiceToken) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceToken) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceToken(val *ServiceToken) *NullableServiceToken {
	return &NullableServiceToken{value: val, isSet: true}
}

func (v NullableServiceToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}