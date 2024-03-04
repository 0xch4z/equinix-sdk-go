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

// checks if the ConnectionSide type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionSide{}

// ConnectionSide Connection configuration object for each side of multi-segment connection
type ConnectionSide struct {
	ServiceToken   *ServiceToken             `json:"serviceToken,omitempty"`
	AccessPoint    *AccessPoint              `json:"accessPoint,omitempty"`
	CompanyProfile *ConnectionCompanyProfile `json:"companyProfile,omitempty"`
	Invitation     *ConnectionInvitation     `json:"invitation,omitempty"`
	// Any additional information, which is not part of connection metadata or configuration
	AdditionalInfo       []ConnectionSideAdditionalInfo `json:"additionalInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionSide ConnectionSide

// NewConnectionSide instantiates a new ConnectionSide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionSide() *ConnectionSide {
	this := ConnectionSide{}
	return &this
}

// NewConnectionSideWithDefaults instantiates a new ConnectionSide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionSideWithDefaults() *ConnectionSide {
	this := ConnectionSide{}
	return &this
}

// GetServiceToken returns the ServiceToken field value if set, zero value otherwise.
func (o *ConnectionSide) GetServiceToken() ServiceToken {
	if o == nil || IsNil(o.ServiceToken) {
		var ret ServiceToken
		return ret
	}
	return *o.ServiceToken
}

// GetServiceTokenOk returns a tuple with the ServiceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSide) GetServiceTokenOk() (*ServiceToken, bool) {
	if o == nil || IsNil(o.ServiceToken) {
		return nil, false
	}
	return o.ServiceToken, true
}

// HasServiceToken returns a boolean if a field has been set.
func (o *ConnectionSide) HasServiceToken() bool {
	if o != nil && !IsNil(o.ServiceToken) {
		return true
	}

	return false
}

// SetServiceToken gets a reference to the given ServiceToken and assigns it to the ServiceToken field.
func (o *ConnectionSide) SetServiceToken(v ServiceToken) {
	o.ServiceToken = &v
}

// GetAccessPoint returns the AccessPoint field value if set, zero value otherwise.
func (o *ConnectionSide) GetAccessPoint() AccessPoint {
	if o == nil || IsNil(o.AccessPoint) {
		var ret AccessPoint
		return ret
	}
	return *o.AccessPoint
}

// GetAccessPointOk returns a tuple with the AccessPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSide) GetAccessPointOk() (*AccessPoint, bool) {
	if o == nil || IsNil(o.AccessPoint) {
		return nil, false
	}
	return o.AccessPoint, true
}

// HasAccessPoint returns a boolean if a field has been set.
func (o *ConnectionSide) HasAccessPoint() bool {
	if o != nil && !IsNil(o.AccessPoint) {
		return true
	}

	return false
}

// SetAccessPoint gets a reference to the given AccessPoint and assigns it to the AccessPoint field.
func (o *ConnectionSide) SetAccessPoint(v AccessPoint) {
	o.AccessPoint = &v
}

// GetCompanyProfile returns the CompanyProfile field value if set, zero value otherwise.
func (o *ConnectionSide) GetCompanyProfile() ConnectionCompanyProfile {
	if o == nil || IsNil(o.CompanyProfile) {
		var ret ConnectionCompanyProfile
		return ret
	}
	return *o.CompanyProfile
}

// GetCompanyProfileOk returns a tuple with the CompanyProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSide) GetCompanyProfileOk() (*ConnectionCompanyProfile, bool) {
	if o == nil || IsNil(o.CompanyProfile) {
		return nil, false
	}
	return o.CompanyProfile, true
}

// HasCompanyProfile returns a boolean if a field has been set.
func (o *ConnectionSide) HasCompanyProfile() bool {
	if o != nil && !IsNil(o.CompanyProfile) {
		return true
	}

	return false
}

// SetCompanyProfile gets a reference to the given ConnectionCompanyProfile and assigns it to the CompanyProfile field.
func (o *ConnectionSide) SetCompanyProfile(v ConnectionCompanyProfile) {
	o.CompanyProfile = &v
}

// GetInvitation returns the Invitation field value if set, zero value otherwise.
func (o *ConnectionSide) GetInvitation() ConnectionInvitation {
	if o == nil || IsNil(o.Invitation) {
		var ret ConnectionInvitation
		return ret
	}
	return *o.Invitation
}

// GetInvitationOk returns a tuple with the Invitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSide) GetInvitationOk() (*ConnectionInvitation, bool) {
	if o == nil || IsNil(o.Invitation) {
		return nil, false
	}
	return o.Invitation, true
}

// HasInvitation returns a boolean if a field has been set.
func (o *ConnectionSide) HasInvitation() bool {
	if o != nil && !IsNil(o.Invitation) {
		return true
	}

	return false
}

// SetInvitation gets a reference to the given ConnectionInvitation and assigns it to the Invitation field.
func (o *ConnectionSide) SetInvitation(v ConnectionInvitation) {
	o.Invitation = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *ConnectionSide) GetAdditionalInfo() []ConnectionSideAdditionalInfo {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret []ConnectionSideAdditionalInfo
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSide) GetAdditionalInfoOk() ([]ConnectionSideAdditionalInfo, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *ConnectionSide) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given []ConnectionSideAdditionalInfo and assigns it to the AdditionalInfo field.
func (o *ConnectionSide) SetAdditionalInfo(v []ConnectionSideAdditionalInfo) {
	o.AdditionalInfo = v
}

func (o ConnectionSide) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionSide) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceToken) {
		toSerialize["serviceToken"] = o.ServiceToken
	}
	if !IsNil(o.AccessPoint) {
		toSerialize["accessPoint"] = o.AccessPoint
	}
	if !IsNil(o.CompanyProfile) {
		toSerialize["companyProfile"] = o.CompanyProfile
	}
	if !IsNil(o.Invitation) {
		toSerialize["invitation"] = o.Invitation
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionSide) UnmarshalJSON(data []byte) (err error) {
	varConnectionSide := _ConnectionSide{}

	err = json.Unmarshal(data, &varConnectionSide)

	if err != nil {
		return err
	}

	*o = ConnectionSide(varConnectionSide)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serviceToken")
		delete(additionalProperties, "accessPoint")
		delete(additionalProperties, "companyProfile")
		delete(additionalProperties, "invitation")
		delete(additionalProperties, "additionalInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionSide struct {
	value *ConnectionSide
	isSet bool
}

func (v NullableConnectionSide) Get() *ConnectionSide {
	return v.value
}

func (v *NullableConnectionSide) Set(val *ConnectionSide) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionSide) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionSide(val *ConnectionSide) *NullableConnectionSide {
	return &NullableConnectionSide{value: val, isSet: true}
}

func (v NullableConnectionSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}