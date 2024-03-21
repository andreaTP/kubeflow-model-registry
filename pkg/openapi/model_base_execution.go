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

// checks if the BaseExecution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseExecution{}

// BaseExecution struct for BaseExecution
type BaseExecution struct {
	LastKnownState *ExecutionState `json:"lastKnownState,omitempty"`
	// Output only. The unique server generated id of the resource.
	Id *string `json:"id,omitempty"`
	// Output only. Create time of the resource in millisecond since epoch.
	CreateTimeSinceEpoch *string `json:"createTimeSinceEpoch,omitempty"`
	// Output only. Last update time of the resource since epoch in millisecond since epoch.
	LastUpdateTimeSinceEpoch *string `json:"lastUpdateTimeSinceEpoch,omitempty"`
	// The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set.
	Name *string `json:"name,omitempty"`
	// User provided custom properties which are not defined by its type.
	CustomProperties *map[string]MetadataValue `json:"customProperties,omitempty"`
	// An optional description about the resource.
	Description *string `json:"description,omitempty"`
	// The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance.
	ExternalID *string `json:"externalID,omitempty"`
}

// NewBaseExecution instantiates a new BaseExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseExecution() *BaseExecution {
	this := BaseExecution{}
	var lastKnownState ExecutionState = EXECUTIONSTATE_UNKNOWN
	this.LastKnownState = &lastKnownState
	return &this
}

// NewBaseExecutionWithDefaults instantiates a new BaseExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseExecutionWithDefaults() *BaseExecution {
	this := BaseExecution{}
	var lastKnownState ExecutionState = EXECUTIONSTATE_UNKNOWN
	this.LastKnownState = &lastKnownState
	return &this
}

// GetLastKnownState returns the LastKnownState field value if set, zero value otherwise.
func (o *BaseExecution) GetLastKnownState() ExecutionState {
	if o == nil || IsNil(o.LastKnownState) {
		var ret ExecutionState
		return ret
	}
	return *o.LastKnownState
}

// GetLastKnownStateOk returns a tuple with the LastKnownState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecution) GetLastKnownStateOk() (*ExecutionState, bool) {
	if o == nil || IsNil(o.LastKnownState) {
		return nil, false
	}
	return o.LastKnownState, true
}

// HasLastKnownState returns a boolean if a field has been set.
func (o *BaseExecution) HasLastKnownState() bool {
	if o != nil && !IsNil(o.LastKnownState) {
		return true
	}

	return false
}

// SetLastKnownState gets a reference to the given ExecutionState and assigns it to the LastKnownState field.
func (o *BaseExecution) SetLastKnownState(v ExecutionState) {
	o.LastKnownState = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseExecution) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecution) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseExecution) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseExecution) SetId(v string) {
	o.Id = &v
}

// GetCreateTimeSinceEpoch returns the CreateTimeSinceEpoch field value if set, zero value otherwise.
func (o *BaseExecution) GetCreateTimeSinceEpoch() string {
	if o == nil || IsNil(o.CreateTimeSinceEpoch) {
		var ret string
		return ret
	}
	return *o.CreateTimeSinceEpoch
}

// GetCreateTimeSinceEpochOk returns a tuple with the CreateTimeSinceEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecution) GetCreateTimeSinceEpochOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTimeSinceEpoch) {
		return nil, false
	}
	return o.CreateTimeSinceEpoch, true
}

// HasCreateTimeSinceEpoch returns a boolean if a field has been set.
func (o *BaseExecution) HasCreateTimeSinceEpoch() bool {
	if o != nil && !IsNil(o.CreateTimeSinceEpoch) {
		return true
	}

	return false
}

// SetCreateTimeSinceEpoch gets a reference to the given string and assigns it to the CreateTimeSinceEpoch field.
func (o *BaseExecution) SetCreateTimeSinceEpoch(v string) {
	o.CreateTimeSinceEpoch = &v
}

// GetLastUpdateTimeSinceEpoch returns the LastUpdateTimeSinceEpoch field value if set, zero value otherwise.
func (o *BaseExecution) GetLastUpdateTimeSinceEpoch() string {
	if o == nil || IsNil(o.LastUpdateTimeSinceEpoch) {
		var ret string
		return ret
	}
	return *o.LastUpdateTimeSinceEpoch
}

// GetLastUpdateTimeSinceEpochOk returns a tuple with the LastUpdateTimeSinceEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecution) GetLastUpdateTimeSinceEpochOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdateTimeSinceEpoch) {
		return nil, false
	}
	return o.LastUpdateTimeSinceEpoch, true
}

// HasLastUpdateTimeSinceEpoch returns a boolean if a field has been set.
func (o *BaseExecution) HasLastUpdateTimeSinceEpoch() bool {
	if o != nil && !IsNil(o.LastUpdateTimeSinceEpoch) {
		return true
	}

	return false
}

// SetLastUpdateTimeSinceEpoch gets a reference to the given string and assigns it to the LastUpdateTimeSinceEpoch field.
func (o *BaseExecution) SetLastUpdateTimeSinceEpoch(v string) {
	o.LastUpdateTimeSinceEpoch = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseExecution) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecution) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseExecution) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseExecution) SetName(v string) {
	o.Name = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *BaseExecution) GetCustomProperties() map[string]MetadataValue {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]MetadataValue
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecution) GetCustomPropertiesOk() (*map[string]MetadataValue, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *BaseExecution) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]MetadataValue and assigns it to the CustomProperties field.
func (o *BaseExecution) SetCustomProperties(v map[string]MetadataValue) {
	o.CustomProperties = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseExecution) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecution) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseExecution) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseExecution) SetDescription(v string) {
	o.Description = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *BaseExecution) GetExternalID() string {
	if o == nil || IsNil(o.ExternalID) {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseExecution) GetExternalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalID) {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *BaseExecution) HasExternalID() bool {
	if o != nil && !IsNil(o.ExternalID) {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *BaseExecution) SetExternalID(v string) {
	o.ExternalID = &v
}

func (o BaseExecution) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseExecution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastKnownState) {
		toSerialize["lastKnownState"] = o.LastKnownState
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreateTimeSinceEpoch) {
		toSerialize["createTimeSinceEpoch"] = o.CreateTimeSinceEpoch
	}
	if !IsNil(o.LastUpdateTimeSinceEpoch) {
		toSerialize["lastUpdateTimeSinceEpoch"] = o.LastUpdateTimeSinceEpoch
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

type NullableBaseExecution struct {
	value *BaseExecution
	isSet bool
}

func (v NullableBaseExecution) Get() *BaseExecution {
	return v.value
}

func (v *NullableBaseExecution) Set(val *BaseExecution) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseExecution(val *BaseExecution) *NullableBaseExecution {
	return &NullableBaseExecution{value: val, isSet: true}
}

func (v NullableBaseExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
