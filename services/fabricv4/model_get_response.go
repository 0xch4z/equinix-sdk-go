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

// checks if the GetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetResponse{}

// GetResponse struct for GetResponse
type GetResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// List of Routing Protocols
	Data                 []RoutingProtocolData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetResponse GetResponse

// NewGetResponse instantiates a new GetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetResponse() *GetResponse {
	this := GetResponse{}
	return &this
}

// NewGetResponseWithDefaults instantiates a new GetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetResponseWithDefaults() *GetResponse {
	this := GetResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *GetResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetResponse) GetData() []RoutingProtocolData {
	if o == nil || IsNil(o.Data) {
		var ret []RoutingProtocolData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResponse) GetDataOk() ([]RoutingProtocolData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RoutingProtocolData and assigns it to the Data field.
func (o *GetResponse) SetData(v []RoutingProtocolData) {
	o.Data = v
}

func (o GetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetResponse) UnmarshalJSON(data []byte) (err error) {
	varGetResponse := _GetResponse{}

	err = json.Unmarshal(data, &varGetResponse)

	if err != nil {
		return err
	}

	*o = GetResponse(varGetResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetResponse struct {
	value *GetResponse
	isSet bool
}

func (v NullableGetResponse) Get() *GetResponse {
	return v.value
}

func (v *NullableGetResponse) Set(val *GetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetResponse(val *GetResponse) *NullableGetResponse {
	return &NullableGetResponse{value: val, isSet: true}
}

func (v NullableGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
