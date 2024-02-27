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

// checks if the PortOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortOrder{}

// PortOrder struct for PortOrder
type PortOrder struct {
	PurchaseOrder *PortOrderPurchaseOrder `json:"purchaseOrder,omitempty"`
	// Order Identification
	OrderId *string `json:"orderId,omitempty"`
	// Order Reference Number
	OrderNumber *string `json:"orderNumber,omitempty"`
	// Equinix-assigned order identifier
	Uuid                 *string             `json:"uuid,omitempty"`
	Signature            *PortOrderSignature `json:"signature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortOrder PortOrder

// NewPortOrder instantiates a new PortOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortOrder() *PortOrder {
	this := PortOrder{}
	return &this
}

// NewPortOrderWithDefaults instantiates a new PortOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortOrderWithDefaults() *PortOrder {
	this := PortOrder{}
	return &this
}

// GetPurchaseOrder returns the PurchaseOrder field value if set, zero value otherwise.
func (o *PortOrder) GetPurchaseOrder() PortOrderPurchaseOrder {
	if o == nil || IsNil(o.PurchaseOrder) {
		var ret PortOrderPurchaseOrder
		return ret
	}
	return *o.PurchaseOrder
}

// GetPurchaseOrderOk returns a tuple with the PurchaseOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortOrder) GetPurchaseOrderOk() (*PortOrderPurchaseOrder, bool) {
	if o == nil || IsNil(o.PurchaseOrder) {
		return nil, false
	}
	return o.PurchaseOrder, true
}

// HasPurchaseOrder returns a boolean if a field has been set.
func (o *PortOrder) HasPurchaseOrder() bool {
	if o != nil && !IsNil(o.PurchaseOrder) {
		return true
	}

	return false
}

// SetPurchaseOrder gets a reference to the given PortOrderPurchaseOrder and assigns it to the PurchaseOrder field.
func (o *PortOrder) SetPurchaseOrder(v PortOrderPurchaseOrder) {
	o.PurchaseOrder = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *PortOrder) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortOrder) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *PortOrder) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *PortOrder) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *PortOrder) GetOrderNumber() string {
	if o == nil || IsNil(o.OrderNumber) {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortOrder) GetOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNumber) {
		return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *PortOrder) HasOrderNumber() bool {
	if o != nil && !IsNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *PortOrder) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *PortOrder) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortOrder) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *PortOrder) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *PortOrder) SetUuid(v string) {
	o.Uuid = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *PortOrder) GetSignature() PortOrderSignature {
	if o == nil || IsNil(o.Signature) {
		var ret PortOrderSignature
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortOrder) GetSignatureOk() (*PortOrderSignature, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *PortOrder) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given PortOrderSignature and assigns it to the Signature field.
func (o *PortOrder) SetSignature(v PortOrderSignature) {
	o.Signature = &v
}

func (o PortOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PurchaseOrder) {
		toSerialize["purchaseOrder"] = o.PurchaseOrder
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.OrderNumber) {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortOrder) UnmarshalJSON(data []byte) (err error) {
	varPortOrder := _PortOrder{}

	err = json.Unmarshal(data, &varPortOrder)

	if err != nil {
		return err
	}

	*o = PortOrder(varPortOrder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "purchaseOrder")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderNumber")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "signature")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortOrder struct {
	value *PortOrder
	isSet bool
}

func (v NullablePortOrder) Get() *PortOrder {
	return v.value
}

func (v *NullablePortOrder) Set(val *PortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullablePortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullablePortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortOrder(val *PortOrder) *NullablePortOrder {
	return &NullablePortOrder{value: val, isSet: true}
}

func (v NullablePortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
