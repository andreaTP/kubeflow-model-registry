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

// checks if the BaseArtifactUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseArtifactUpdate{}

// BaseArtifactUpdate struct for BaseArtifactUpdate
type BaseArtifactUpdate struct {
	// The uniform resource identifier of the physical artifact. May be empty if there is no physical artifact.
	Uri   *string        `json:"uri,omitempty"`
	State *ArtifactState `json:"state,omitempty"`
	// User provided custom properties which are not defined by its type.
	CustomProperties *map[string]MetadataValue `json:"customProperties,omitempty"`
	// An optional description about the resource.
	Description *string `json:"description,omitempty"`
	// The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance.
	ExternalID *string `json:"externalID,omitempty"`
}

// NewBaseArtifactUpdate instantiates a new BaseArtifactUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseArtifactUpdate() *BaseArtifactUpdate {
	this := BaseArtifactUpdate{}
	return &this
}

// NewBaseArtifactUpdateWithDefaults instantiates a new BaseArtifactUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseArtifactUpdateWithDefaults() *BaseArtifactUpdate {
	this := BaseArtifactUpdate{}
	var state ArtifactState = ARTIFACTSTATE_UNKNOWN
	this.State = &state
	return &this
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BaseArtifactUpdate) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseArtifactUpdate) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BaseArtifactUpdate) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BaseArtifactUpdate) SetUri(v string) {
	o.Uri = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BaseArtifactUpdate) GetState() ArtifactState {
	if o == nil || IsNil(o.State) {
		var ret ArtifactState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseArtifactUpdate) GetStateOk() (*ArtifactState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BaseArtifactUpdate) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ArtifactState and assigns it to the State field.
func (o *BaseArtifactUpdate) SetState(v ArtifactState) {
	o.State = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *BaseArtifactUpdate) GetCustomProperties() map[string]MetadataValue {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]MetadataValue
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseArtifactUpdate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *BaseArtifactUpdate) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]MetadataValue and assigns it to the CustomProperties field.
func (o *BaseArtifactUpdate) SetCustomProperties(v map[string]MetadataValue) {
	o.CustomProperties = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseArtifactUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseArtifactUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseArtifactUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseArtifactUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *BaseArtifactUpdate) GetExternalID() string {
	if o == nil || IsNil(o.ExternalID) {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseArtifactUpdate) GetExternalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalID) {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *BaseArtifactUpdate) HasExternalID() bool {
	if o != nil && !IsNil(o.ExternalID) {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *BaseArtifactUpdate) SetExternalID(v string) {
	o.ExternalID = &v
}

func (o BaseArtifactUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseArtifactUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
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

type NullableBaseArtifactUpdate struct {
	value *BaseArtifactUpdate
	isSet bool
}

func (v NullableBaseArtifactUpdate) Get() *BaseArtifactUpdate {
	return v.value
}

func (v *NullableBaseArtifactUpdate) Set(val *BaseArtifactUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseArtifactUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseArtifactUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseArtifactUpdate(val *BaseArtifactUpdate) *NullableBaseArtifactUpdate {
	return &NullableBaseArtifactUpdate{value: val, isSet: true}
}

func (v NullableBaseArtifactUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseArtifactUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
