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

// checks if the RouteFilterRulesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteFilterRulesData{}

// RouteFilterRulesData struct for RouteFilterRulesData
type RouteFilterRulesData struct {
	// Route Filter Rules URI
	Href *string                   `json:"href,omitempty"`
	Type *RouteFilterRulesDataType `json:"type,omitempty"`
	// Route Filter Rule identifier
	Uuid *string `json:"uuid,omitempty"`
	Name *string `json:"name,omitempty"`
	// Customer-provided Route Filter Rule description
	Description *string               `json:"description,omitempty"`
	State       *RouteFilterRuleState `json:"state,omitempty"`
	// prefix matching operator
	PrefixMatch          *string                     `json:"prefixMatch,omitempty"`
	Change               *RouteFilterRulesChange     `json:"change,omitempty"`
	Action               *RouteFilterRulesDataAction `json:"action,omitempty"`
	Prefix               *string                     `json:"prefix,omitempty"`
	Changelog            *Changelog                  `json:"changelog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteFilterRulesData RouteFilterRulesData

// NewRouteFilterRulesData instantiates a new RouteFilterRulesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteFilterRulesData() *RouteFilterRulesData {
	this := RouteFilterRulesData{}
	var prefixMatch string = "orlonger"
	this.PrefixMatch = &prefixMatch
	return &this
}

// NewRouteFilterRulesDataWithDefaults instantiates a new RouteFilterRulesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteFilterRulesDataWithDefaults() *RouteFilterRulesData {
	this := RouteFilterRulesData{}
	var prefixMatch string = "orlonger"
	this.PrefixMatch = &prefixMatch
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RouteFilterRulesData) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesData) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RouteFilterRulesData) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RouteFilterRulesData) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RouteFilterRulesData) GetType() RouteFilterRulesDataType {
	if o == nil || IsNil(o.Type) {
		var ret RouteFilterRulesDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesData) GetTypeOk() (*RouteFilterRulesDataType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RouteFilterRulesData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RouteFilterRulesDataType and assigns it to the Type field.
func (o *RouteFilterRulesData) SetType(v RouteFilterRulesDataType) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RouteFilterRulesData) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesData) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RouteFilterRulesData) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RouteFilterRulesData) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RouteFilterRulesData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RouteFilterRulesData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RouteFilterRulesData) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RouteFilterRulesData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RouteFilterRulesData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RouteFilterRulesData) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RouteFilterRulesData) GetState() RouteFilterRuleState {
	if o == nil || IsNil(o.State) {
		var ret RouteFilterRuleState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesData) GetStateOk() (*RouteFilterRuleState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RouteFilterRulesData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given RouteFilterRuleState and assigns it to the State field.
func (o *RouteFilterRulesData) SetState(v RouteFilterRuleState) {
	o.State = &v
}

// GetPrefixMatch returns the PrefixMatch field value if set, zero value otherwise.
func (o *RouteFilterRulesData) GetPrefixMatch() string {
	if o == nil || IsNil(o.PrefixMatch) {
		var ret string
		return ret
	}
	return *o.PrefixMatch
}

// GetPrefixMatchOk returns a tuple with the PrefixMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesData) GetPrefixMatchOk() (*string, bool) {
	if o == nil || IsNil(o.PrefixMatch) {
		return nil, false
	}
	return o.PrefixMatch, true
}

// HasPrefixMatch returns a boolean if a field has been set.
func (o *RouteFilterRulesData) HasPrefixMatch() bool {
	if o != nil && !IsNil(o.PrefixMatch) {
		return true
	}

	return false
}

// SetPrefixMatch gets a reference to the given string and assigns it to the PrefixMatch field.
func (o *RouteFilterRulesData) SetPrefixMatch(v string) {
	o.PrefixMatch = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *RouteFilterRulesData) GetChange() RouteFilterRulesChange {
	if o == nil || IsNil(o.Change) {
		var ret RouteFilterRulesChange
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesData) GetChangeOk() (*RouteFilterRulesChange, bool) {
	if o == nil || IsNil(o.Change) {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *RouteFilterRulesData) HasChange() bool {
	if o != nil && !IsNil(o.Change) {
		return true
	}

	return false
}

// SetChange gets a reference to the given RouteFilterRulesChange and assigns it to the Change field.
func (o *RouteFilterRulesData) SetChange(v RouteFilterRulesChange) {
	o.Change = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RouteFilterRulesData) GetAction() RouteFilterRulesDataAction {
	if o == nil || IsNil(o.Action) {
		var ret RouteFilterRulesDataAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesData) GetActionOk() (*RouteFilterRulesDataAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RouteFilterRulesData) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given RouteFilterRulesDataAction and assigns it to the Action field.
func (o *RouteFilterRulesData) SetAction(v RouteFilterRulesDataAction) {
	o.Action = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *RouteFilterRulesData) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesData) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *RouteFilterRulesData) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *RouteFilterRulesData) SetPrefix(v string) {
	o.Prefix = &v
}

// GetChangelog returns the Changelog field value if set, zero value otherwise.
func (o *RouteFilterRulesData) GetChangelog() Changelog {
	if o == nil || IsNil(o.Changelog) {
		var ret Changelog
		return ret
	}
	return *o.Changelog
}

// GetChangelogOk returns a tuple with the Changelog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesData) GetChangelogOk() (*Changelog, bool) {
	if o == nil || IsNil(o.Changelog) {
		return nil, false
	}
	return o.Changelog, true
}

// HasChangelog returns a boolean if a field has been set.
func (o *RouteFilterRulesData) HasChangelog() bool {
	if o != nil && !IsNil(o.Changelog) {
		return true
	}

	return false
}

// SetChangelog gets a reference to the given Changelog and assigns it to the Changelog field.
func (o *RouteFilterRulesData) SetChangelog(v Changelog) {
	o.Changelog = &v
}

func (o RouteFilterRulesData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteFilterRulesData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PrefixMatch) {
		toSerialize["prefixMatch"] = o.PrefixMatch
	}
	if !IsNil(o.Change) {
		toSerialize["change"] = o.Change
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Changelog) {
		toSerialize["changelog"] = o.Changelog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteFilterRulesData) UnmarshalJSON(data []byte) (err error) {
	varRouteFilterRulesData := _RouteFilterRulesData{}

	err = json.Unmarshal(data, &varRouteFilterRulesData)

	if err != nil {
		return err
	}

	*o = RouteFilterRulesData(varRouteFilterRulesData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "state")
		delete(additionalProperties, "prefixMatch")
		delete(additionalProperties, "change")
		delete(additionalProperties, "action")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "changelog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteFilterRulesData struct {
	value *RouteFilterRulesData
	isSet bool
}

func (v NullableRouteFilterRulesData) Get() *RouteFilterRulesData {
	return v.value
}

func (v *NullableRouteFilterRulesData) Set(val *RouteFilterRulesData) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFilterRulesData) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFilterRulesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFilterRulesData(val *RouteFilterRulesData) *NullableRouteFilterRulesData {
	return &NullableRouteFilterRulesData{value: val, isSet: true}
}

func (v NullableRouteFilterRulesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFilterRulesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
