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

// OperatingSystem struct for OperatingSystem
type OperatingSystem struct {
	Distro *string `json:"distro,omitempty"`
	Id     *string `json:"id,omitempty"`
	// Licenced OS is priced according to pricing property
	Licensed *bool   `json:"licensed,omitempty"`
	Name     *string `json:"name,omitempty"`
	// Servers can be already preinstalled with OS in order to shorten provision time.
	Preinstallable *bool `json:"preinstallable,omitempty"`
	// This object contains price per time unit and optional multiplier value if licence price depends on hardware plan or components (e.g. number of cores)
	Pricing         map[string]interface{} `json:"pricing,omitempty"`
	ProvisionableOn []string               `json:"provisionable_on,omitempty"`
	Slug            *string                `json:"slug,omitempty"`
	Version         *string                `json:"version,omitempty"`
}

// NewOperatingSystem instantiates a new OperatingSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatingSystem() *OperatingSystem {
	this := OperatingSystem{}
	return &this
}

// NewOperatingSystemWithDefaults instantiates a new OperatingSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatingSystemWithDefaults() *OperatingSystem {
	this := OperatingSystem{}
	return &this
}

// GetDistro returns the Distro field value if set, zero value otherwise.
func (o *OperatingSystem) GetDistro() string {
	if o == nil || o.Distro == nil {
		var ret string
		return ret
	}
	return *o.Distro
}

// GetDistroOk returns a tuple with the Distro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetDistroOk() (*string, bool) {
	if o == nil || o.Distro == nil {
		return nil, false
	}
	return o.Distro, true
}

// HasDistro returns a boolean if a field has been set.
func (o *OperatingSystem) HasDistro() bool {
	if o != nil && o.Distro != nil {
		return true
	}

	return false
}

// SetDistro gets a reference to the given string and assigns it to the Distro field.
func (o *OperatingSystem) SetDistro(v string) {
	o.Distro = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OperatingSystem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OperatingSystem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OperatingSystem) SetId(v string) {
	o.Id = &v
}

// GetLicensed returns the Licensed field value if set, zero value otherwise.
func (o *OperatingSystem) GetLicensed() bool {
	if o == nil || o.Licensed == nil {
		var ret bool
		return ret
	}
	return *o.Licensed
}

// GetLicensedOk returns a tuple with the Licensed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetLicensedOk() (*bool, bool) {
	if o == nil || o.Licensed == nil {
		return nil, false
	}
	return o.Licensed, true
}

// HasLicensed returns a boolean if a field has been set.
func (o *OperatingSystem) HasLicensed() bool {
	if o != nil && o.Licensed != nil {
		return true
	}

	return false
}

// SetLicensed gets a reference to the given bool and assigns it to the Licensed field.
func (o *OperatingSystem) SetLicensed(v bool) {
	o.Licensed = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OperatingSystem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OperatingSystem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OperatingSystem) SetName(v string) {
	o.Name = &v
}

// GetPreinstallable returns the Preinstallable field value if set, zero value otherwise.
func (o *OperatingSystem) GetPreinstallable() bool {
	if o == nil || o.Preinstallable == nil {
		var ret bool
		return ret
	}
	return *o.Preinstallable
}

// GetPreinstallableOk returns a tuple with the Preinstallable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetPreinstallableOk() (*bool, bool) {
	if o == nil || o.Preinstallable == nil {
		return nil, false
	}
	return o.Preinstallable, true
}

// HasPreinstallable returns a boolean if a field has been set.
func (o *OperatingSystem) HasPreinstallable() bool {
	if o != nil && o.Preinstallable != nil {
		return true
	}

	return false
}

// SetPreinstallable gets a reference to the given bool and assigns it to the Preinstallable field.
func (o *OperatingSystem) SetPreinstallable(v bool) {
	o.Preinstallable = &v
}

// GetPricing returns the Pricing field value if set, zero value otherwise.
func (o *OperatingSystem) GetPricing() map[string]interface{} {
	if o == nil || o.Pricing == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetPricingOk() (map[string]interface{}, bool) {
	if o == nil || o.Pricing == nil {
		return nil, false
	}
	return o.Pricing, true
}

// HasPricing returns a boolean if a field has been set.
func (o *OperatingSystem) HasPricing() bool {
	if o != nil && o.Pricing != nil {
		return true
	}

	return false
}

// SetPricing gets a reference to the given map[string]interface{} and assigns it to the Pricing field.
func (o *OperatingSystem) SetPricing(v map[string]interface{}) {
	o.Pricing = v
}

// GetProvisionableOn returns the ProvisionableOn field value if set, zero value otherwise.
func (o *OperatingSystem) GetProvisionableOn() []string {
	if o == nil || o.ProvisionableOn == nil {
		var ret []string
		return ret
	}
	return o.ProvisionableOn
}

// GetProvisionableOnOk returns a tuple with the ProvisionableOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetProvisionableOnOk() ([]string, bool) {
	if o == nil || o.ProvisionableOn == nil {
		return nil, false
	}
	return o.ProvisionableOn, true
}

// HasProvisionableOn returns a boolean if a field has been set.
func (o *OperatingSystem) HasProvisionableOn() bool {
	if o != nil && o.ProvisionableOn != nil {
		return true
	}

	return false
}

// SetProvisionableOn gets a reference to the given []string and assigns it to the ProvisionableOn field.
func (o *OperatingSystem) SetProvisionableOn(v []string) {
	o.ProvisionableOn = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *OperatingSystem) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *OperatingSystem) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *OperatingSystem) SetSlug(v string) {
	o.Slug = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OperatingSystem) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OperatingSystem) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *OperatingSystem) SetVersion(v string) {
	o.Version = &v
}

func (o OperatingSystem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Distro != nil {
		toSerialize["distro"] = o.Distro
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Licensed != nil {
		toSerialize["licensed"] = o.Licensed
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Preinstallable != nil {
		toSerialize["preinstallable"] = o.Preinstallable
	}
	if o.Pricing != nil {
		toSerialize["pricing"] = o.Pricing
	}
	if o.ProvisionableOn != nil {
		toSerialize["provisionable_on"] = o.ProvisionableOn
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableOperatingSystem struct {
	value *OperatingSystem
	isSet bool
}

func (v NullableOperatingSystem) Get() *OperatingSystem {
	return v.value
}

func (v *NullableOperatingSystem) Set(val *OperatingSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatingSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatingSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatingSystem(val *OperatingSystem) *NullableOperatingSystem {
	return &NullableOperatingSystem{value: val, isSet: true}
}

func (v NullableOperatingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatingSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
