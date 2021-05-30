/*
 * CrossID Public API
 *
 * CrossID API Reference 
 *
 * API version: 1.0.0
 * Contact: developer@crossid.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cidclient

import (
	"encoding/json"
)

// ResourceType struct for ResourceType
type ResourceType struct {
	AppId *string `json:"appId,omitempty"`
	Description *string `json:"description,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Id *string `json:"id,omitempty"`
	Meta *ApiTokenMeta `json:"meta,omitempty"`
	Mode *string `json:"mode,omitempty"`
	Name *string `json:"name,omitempty"`
	Schema *string `json:"schema,omitempty"`
	SchemaExtensions *[]ResourceTypeSchemaExtensions `json:"schemaExtensions,omitempty"`
	SchemaInterfaces *[]ResourceTypeSchemaExtensions `json:"schemaInterfaces,omitempty"`
	ToApp *string `json:"toApp,omitempty"`
	ToStore *string `json:"toStore,omitempty"`
	Ui *ResourceTypeUi `json:"ui,omitempty"`
}

// NewResourceType instantiates a new ResourceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceType() *ResourceType {
	this := ResourceType{}
	return &this
}

// NewResourceTypeWithDefaults instantiates a new ResourceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTypeWithDefaults() *ResourceType {
	this := ResourceType{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ResourceType) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ResourceType) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *ResourceType) SetAppId(v string) {
	o.AppId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceType) SetDescription(v string) {
	o.Description = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ResourceType) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ResourceType) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *ResourceType) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceType) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceType) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ResourceType) GetMeta() ApiTokenMeta {
	if o == nil || o.Meta == nil {
		var ret ApiTokenMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetMetaOk() (*ApiTokenMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ResourceType) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiTokenMeta and assigns it to the Meta field.
func (o *ResourceType) SetMeta(v ApiTokenMeta) {
	o.Meta = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ResourceType) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ResourceType) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ResourceType) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceType) SetName(v string) {
	o.Name = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ResourceType) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ResourceType) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *ResourceType) SetSchema(v string) {
	o.Schema = &v
}

// GetSchemaExtensions returns the SchemaExtensions field value if set, zero value otherwise.
func (o *ResourceType) GetSchemaExtensions() []ResourceTypeSchemaExtensions {
	if o == nil || o.SchemaExtensions == nil {
		var ret []ResourceTypeSchemaExtensions
		return ret
	}
	return *o.SchemaExtensions
}

// GetSchemaExtensionsOk returns a tuple with the SchemaExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetSchemaExtensionsOk() (*[]ResourceTypeSchemaExtensions, bool) {
	if o == nil || o.SchemaExtensions == nil {
		return nil, false
	}
	return o.SchemaExtensions, true
}

// HasSchemaExtensions returns a boolean if a field has been set.
func (o *ResourceType) HasSchemaExtensions() bool {
	if o != nil && o.SchemaExtensions != nil {
		return true
	}

	return false
}

// SetSchemaExtensions gets a reference to the given []ResourceTypeSchemaExtensions and assigns it to the SchemaExtensions field.
func (o *ResourceType) SetSchemaExtensions(v []ResourceTypeSchemaExtensions) {
	o.SchemaExtensions = &v
}

// GetSchemaInterfaces returns the SchemaInterfaces field value if set, zero value otherwise.
func (o *ResourceType) GetSchemaInterfaces() []ResourceTypeSchemaExtensions {
	if o == nil || o.SchemaInterfaces == nil {
		var ret []ResourceTypeSchemaExtensions
		return ret
	}
	return *o.SchemaInterfaces
}

// GetSchemaInterfacesOk returns a tuple with the SchemaInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetSchemaInterfacesOk() (*[]ResourceTypeSchemaExtensions, bool) {
	if o == nil || o.SchemaInterfaces == nil {
		return nil, false
	}
	return o.SchemaInterfaces, true
}

// HasSchemaInterfaces returns a boolean if a field has been set.
func (o *ResourceType) HasSchemaInterfaces() bool {
	if o != nil && o.SchemaInterfaces != nil {
		return true
	}

	return false
}

// SetSchemaInterfaces gets a reference to the given []ResourceTypeSchemaExtensions and assigns it to the SchemaInterfaces field.
func (o *ResourceType) SetSchemaInterfaces(v []ResourceTypeSchemaExtensions) {
	o.SchemaInterfaces = &v
}

// GetToApp returns the ToApp field value if set, zero value otherwise.
func (o *ResourceType) GetToApp() string {
	if o == nil || o.ToApp == nil {
		var ret string
		return ret
	}
	return *o.ToApp
}

// GetToAppOk returns a tuple with the ToApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetToAppOk() (*string, bool) {
	if o == nil || o.ToApp == nil {
		return nil, false
	}
	return o.ToApp, true
}

// HasToApp returns a boolean if a field has been set.
func (o *ResourceType) HasToApp() bool {
	if o != nil && o.ToApp != nil {
		return true
	}

	return false
}

// SetToApp gets a reference to the given string and assigns it to the ToApp field.
func (o *ResourceType) SetToApp(v string) {
	o.ToApp = &v
}

// GetToStore returns the ToStore field value if set, zero value otherwise.
func (o *ResourceType) GetToStore() string {
	if o == nil || o.ToStore == nil {
		var ret string
		return ret
	}
	return *o.ToStore
}

// GetToStoreOk returns a tuple with the ToStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetToStoreOk() (*string, bool) {
	if o == nil || o.ToStore == nil {
		return nil, false
	}
	return o.ToStore, true
}

// HasToStore returns a boolean if a field has been set.
func (o *ResourceType) HasToStore() bool {
	if o != nil && o.ToStore != nil {
		return true
	}

	return false
}

// SetToStore gets a reference to the given string and assigns it to the ToStore field.
func (o *ResourceType) SetToStore(v string) {
	o.ToStore = &v
}

// GetUi returns the Ui field value if set, zero value otherwise.
func (o *ResourceType) GetUi() ResourceTypeUi {
	if o == nil || o.Ui == nil {
		var ret ResourceTypeUi
		return ret
	}
	return *o.Ui
}

// GetUiOk returns a tuple with the Ui field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetUiOk() (*ResourceTypeUi, bool) {
	if o == nil || o.Ui == nil {
		return nil, false
	}
	return o.Ui, true
}

// HasUi returns a boolean if a field has been set.
func (o *ResourceType) HasUi() bool {
	if o != nil && o.Ui != nil {
		return true
	}

	return false
}

// SetUi gets a reference to the given ResourceTypeUi and assigns it to the Ui field.
func (o *ResourceType) SetUi(v ResourceTypeUi) {
	o.Ui = &v
}

func (o ResourceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.SchemaExtensions != nil {
		toSerialize["schemaExtensions"] = o.SchemaExtensions
	}
	if o.SchemaInterfaces != nil {
		toSerialize["schemaInterfaces"] = o.SchemaInterfaces
	}
	if o.ToApp != nil {
		toSerialize["toApp"] = o.ToApp
	}
	if o.ToStore != nil {
		toSerialize["toStore"] = o.ToStore
	}
	if o.Ui != nil {
		toSerialize["ui"] = o.Ui
	}
	return json.Marshal(toSerialize)
}

type NullableResourceType struct {
	value *ResourceType
	isSet bool
}

func (v NullableResourceType) Get() *ResourceType {
	return v.value
}

func (v *NullableResourceType) Set(val *ResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceType(val *ResourceType) *NullableResourceType {
	return &NullableResourceType{value: val, isSet: true}
}

func (v NullableResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

