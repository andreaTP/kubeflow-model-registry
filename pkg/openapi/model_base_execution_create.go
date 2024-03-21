/*
Model Registry REST API

REST API for Model Registry to create and manage ML model metadata

API version: v1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BaseExecutionCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseExecutionCreate{}

// BaseExecutionCreate struct for BaseExecutionCreate
type BaseExecutionCreate struct {
	LastKnownState *ExecutionState `json:"lastKnownState,omitempty"`
	// The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set.
	Name *string `json:"name,omitempty"`
	// User provided custom properties which are not defined by its type.
	CustomProperties *map[string]MetadataValue `json:"customProperties,omitempty"`
	// An optional description about the resource.
	Description *string `json:"description,omitempty"`
	// The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance.
	ExternalID *string `json:"externalID,omitempty"`
}

// NewBaseExecutionCreate instantiates a new BaseExecutionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseExecutionCreate() *BaseExecutionCreate {
	this := BaseExecutionCreate{}
	var lastKnownState ExecutionState = EXECUTIONSTATE_UNKNOWN
	this.LastKnownState = &lastKnownState
	return &this
}

// NewBaseExecutionCreateWithDefaults instantiates a new BaseExecutionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseExecutionCreateWithDefaults() *BaseExecutionCreate {
	this := BaseExecutionCreate{}
	var lastKnownState ExecutionState = EXECUTIONSTATE_UNKNOWN
	this.LastKnownState = &lastKnownState
	return &this
}

// GetLastKnownState returns the LastKnownState field value if set, zero value otherwise.
func (o *BaseExecutionCreate) GetLastKnownState() ExecutionState {
	if o == nil || IsNil(o.LastKnownState) {
		var ret ExecutionState
		return ret
	}
	return *o.LastKnownState
}

// GetLastKnownStateOk returns a tuple with the LastKnownState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecutionCreate) GetLastKnownStateOk() (*ExecutionState, bool) {
	if o == nil || IsNil(o.LastKnownState) {
		return nil, false
	}
	return o.LastKnownState, true
}

// HasLastKnownState returns a boolean if a field has been set.
func (o *BaseExecutionCreate) HasLastKnownState() bool {
	if o != nil && !IsNil(o.LastKnownState) {
		return true
	}

	return false
}

// SetLastKnownState gets a reference to the given ExecutionState and assigns it to the LastKnownState field.
func (o *BaseExecutionCreate) SetLastKnownState(v ExecutionState) {
	o.LastKnownState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseExecutionCreate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecutionCreate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseExecutionCreate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseExecutionCreate) SetName(v string) {
	o.Name = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *BaseExecutionCreate) GetCustomProperties() map[string]MetadataValue {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]MetadataValue
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecutionCreate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *BaseExecutionCreate) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]MetadataValue and assigns it to the CustomProperties field.
func (o *BaseExecutionCreate) SetCustomProperties(v map[string]MetadataValue) {
	o.CustomProperties = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseExecutionCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecutionCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseExecutionCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseExecutionCreate) SetDescription(v string) {
	o.Description = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *BaseExecutionCreate) GetExternalID() string {
	if o == nil || IsNil(o.ExternalID) {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecutionCreate) GetExternalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalID) {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *BaseExecutionCreate) HasExternalID() bool {
	if o != nil && !IsNil(o.ExternalID) {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *BaseExecutionCreate) SetExternalID(v string) {
	o.ExternalID = &v
}

func (o BaseExecutionCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseExecutionCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastKnownState) {
		toSerialize["lastKnownState"] = o.LastKnownState
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalID) {
		toSerialize["externalID"] = o.ExternalID
	}
	return toSerialize, nil
}

type NullableBaseExecutionCreate struct {
	value *BaseExecutionCreate
	isSet bool
}

func (v NullableBaseExecutionCreate) Get() *BaseExecutionCreate {
	return v.value
}

func (v *NullableBaseExecutionCreate) Set(val *BaseExecutionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseExecutionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseExecutionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseExecutionCreate(val *BaseExecutionCreate) *NullableBaseExecutionCreate {
	return &NullableBaseExecutionCreate{value: val, isSet: true}
}

func (v NullableBaseExecutionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseExecutionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
