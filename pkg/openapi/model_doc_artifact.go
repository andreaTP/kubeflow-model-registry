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

// checks if the DocArtifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocArtifact{}

// DocArtifact A document.
type DocArtifact struct {
	ArtifactType *string `json:"artifactType,omitempty"`
	// The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set.
	Name *string `json:"name,omitempty"`
	// User provided custom properties which are not defined by its type.
	CustomProperties *map[string]MetadataValue `json:"customProperties,omitempty"`
	// An optional description about the resource.
	Description *string `json:"description,omitempty"`
	// The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance.
	ExternalID *string `json:"externalID,omitempty"`
}

// NewDocArtifact instantiates a new DocArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocArtifact() *DocArtifact {
	this := DocArtifact{}
	return &this
}

// NewDocArtifactWithDefaults instantiates a new DocArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocArtifactWithDefaults() *DocArtifact {
	this := DocArtifact{}
	var artifactType string = "doc-artifact"
	this.ArtifactType = &artifactType
	return &this
}

// GetArtifactType returns the ArtifactType field value if set, zero value otherwise.
func (o *DocArtifact) GetArtifactType() string {
	if o == nil || IsNil(o.ArtifactType) {
		var ret string
		return ret
	}
	return *o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocArtifact) GetArtifactTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ArtifactType) {
		return nil, false
	}
	return o.ArtifactType, true
}

// HasArtifactType returns a boolean if a field has been set.
func (o *DocArtifact) HasArtifactType() bool {
	if o != nil && !IsNil(o.ArtifactType) {
		return true
	}

	return false
}

// SetArtifactType gets a reference to the given string and assigns it to the ArtifactType field.
func (o *DocArtifact) SetArtifactType(v string) {
	o.ArtifactType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DocArtifact) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocArtifact) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DocArtifact) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DocArtifact) SetName(v string) {
	o.Name = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *DocArtifact) GetCustomProperties() map[string]MetadataValue {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]MetadataValue
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocArtifact) GetCustomPropertiesOk() (*map[string]MetadataValue, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *DocArtifact) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]MetadataValue and assigns it to the CustomProperties field.
func (o *DocArtifact) SetCustomProperties(v map[string]MetadataValue) {
	o.CustomProperties = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DocArtifact) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocArtifact) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DocArtifact) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DocArtifact) SetDescription(v string) {
	o.Description = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *DocArtifact) GetExternalID() string {
	if o == nil || IsNil(o.ExternalID) {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocArtifact) GetExternalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalID) {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *DocArtifact) HasExternalID() bool {
	if o != nil && !IsNil(o.ExternalID) {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *DocArtifact) SetExternalID(v string) {
	o.ExternalID = &v
}

func (o DocArtifact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocArtifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArtifactType) {
		toSerialize["artifactType"] = o.ArtifactType
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

type NullableDocArtifact struct {
	value *DocArtifact
	isSet bool
}

func (v NullableDocArtifact) Get() *DocArtifact {
	return v.value
}

func (v *NullableDocArtifact) Set(val *DocArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableDocArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableDocArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocArtifact(val *DocArtifact) *NullableDocArtifact {
	return &NullableDocArtifact{value: val, isSet: true}
}

func (v NullableDocArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
