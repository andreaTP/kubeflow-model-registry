// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package generated

import (
	"fmt"
	converter "github.com/kubeflow/model-registry/internal/converter"
	proto "github.com/kubeflow/model-registry/internal/ml_metadata/proto"
	openapi "github.com/kubeflow/model-registry/pkg/openapi"
)

type MLMDToOpenAPIConverterImpl struct{}

func (c *MLMDToOpenAPIConverterImpl) ConvertDocArtifact(source *proto.Artifact) (*openapi.DocArtifact, error) {
	var pOpenapiDocArtifact *openapi.DocArtifact
	if source != nil {
		var openapiDocArtifact openapi.DocArtifact
		xstring, err := converter.MapArtifactType(source)
		if err != nil {
			return nil, fmt.Errorf("error setting field ArtifactType: %w", err)
		}
		openapiDocArtifact.ArtifactType = xstring
		var pString *string
		if (*source).Uri != nil {
			xstring2 := *(*source).Uri
			pString = &xstring2
		}
		openapiDocArtifact.Uri = pString
		openapiDocArtifact.State = converter.MapMLMDArtifactState((*source).State)
		openapiDocArtifact.Id = converter.Int64ToString((*source).Id)
		openapiDocArtifact.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiDocArtifact.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		openapiDocArtifact.Name = converter.MapNameFromOwned((*source).Name)
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, err
		}
		openapiDocArtifact.CustomProperties = &mapStringOpenapiMetadataValue
		openapiDocArtifact.Description = converter.MapDescription((*source).Properties)
		var pString2 *string
		if (*source).ExternalId != nil {
			xstring3 := *(*source).ExternalId
			pString2 = &xstring3
		}
		openapiDocArtifact.ExternalId = pString2
		pOpenapiDocArtifact = &openapiDocArtifact
	}
	return pOpenapiDocArtifact, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertInferenceService(source *proto.Context) (*openapi.InferenceService, error) {
	var pOpenapiInferenceService *openapi.InferenceService
	if source != nil {
		var openapiInferenceService openapi.InferenceService
		openapiInferenceService.Id = converter.Int64ToString((*source).Id)
		openapiInferenceService.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiInferenceService.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		openapiInferenceService.Name = converter.MapNameFromOwned((*source).Name)
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, err
		}
		openapiInferenceService.CustomProperties = &mapStringOpenapiMetadataValue
		openapiInferenceService.Description = converter.MapDescription((*source).Properties)
		var pString *string
		if (*source).ExternalId != nil {
			xstring := *(*source).ExternalId
			pString = &xstring
		}
		openapiInferenceService.ExternalId = pString
		openapiInferenceService.ModelVersionId = converter.MapPropertyModelVersionId((*source).Properties)
		openapiInferenceService.Runtime = converter.MapPropertyRuntime((*source).Properties)
		openapiInferenceService.DesiredState = converter.MapInferenceServiceDesiredState((*source).Properties)
		openapiInferenceService.RegisteredModelId = converter.MapPropertyRegisteredModelId((*source).Properties)
		openapiInferenceService.ServingEnvironmentId = converter.MapPropertyServingEnvironmentId((*source).Properties)
		pOpenapiInferenceService = &openapiInferenceService
	}
	return pOpenapiInferenceService, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertModelArtifact(source *proto.Artifact) (*openapi.ModelArtifact, error) {
	var pOpenapiModelArtifact *openapi.ModelArtifact
	if source != nil {
		var openapiModelArtifact openapi.ModelArtifact
		xstring, err := converter.MapArtifactType(source)
		if err != nil {
			return nil, fmt.Errorf("error setting field ArtifactType: %w", err)
		}
		openapiModelArtifact.ArtifactType = xstring
		var pString *string
		if (*source).Uri != nil {
			xstring2 := *(*source).Uri
			pString = &xstring2
		}
		openapiModelArtifact.Uri = pString
		openapiModelArtifact.State = converter.MapMLMDArtifactState((*source).State)
		openapiModelArtifact.Id = converter.Int64ToString((*source).Id)
		openapiModelArtifact.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiModelArtifact.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		openapiModelArtifact.Name = converter.MapNameFromOwned((*source).Name)
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, err
		}
		openapiModelArtifact.CustomProperties = &mapStringOpenapiMetadataValue
		openapiModelArtifact.Description = converter.MapDescription((*source).Properties)
		var pString2 *string
		if (*source).ExternalId != nil {
			xstring3 := *(*source).ExternalId
			pString2 = &xstring3
		}
		openapiModelArtifact.ExternalId = pString2
		openapiModelArtifact.ModelFormatName = converter.MapModelArtifactFormatName((*source).Properties)
		openapiModelArtifact.StorageKey = converter.MapModelArtifactStorageKey((*source).Properties)
		openapiModelArtifact.StoragePath = converter.MapModelArtifactStoragePath((*source).Properties)
		openapiModelArtifact.ModelFormatVersion = converter.MapModelArtifactFormatVersion((*source).Properties)
		openapiModelArtifact.ServiceAccountName = converter.MapModelArtifactServiceAccountName((*source).Properties)
		pOpenapiModelArtifact = &openapiModelArtifact
	}
	return pOpenapiModelArtifact, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertModelVersion(source *proto.Context) (*openapi.ModelVersion, error) {
	var pOpenapiModelVersion *openapi.ModelVersion
	if source != nil {
		var openapiModelVersion openapi.ModelVersion
		openapiModelVersion.Id = converter.Int64ToString((*source).Id)
		openapiModelVersion.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiModelVersion.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		openapiModelVersion.Name = converter.MapNameFromOwned((*source).Name)
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, err
		}
		openapiModelVersion.CustomProperties = &mapStringOpenapiMetadataValue
		openapiModelVersion.Description = converter.MapDescription((*source).Properties)
		var pString *string
		if (*source).ExternalId != nil {
			xstring := *(*source).ExternalId
			pString = &xstring
		}
		openapiModelVersion.ExternalId = pString
		openapiModelVersion.State = converter.MapModelVersionState((*source).Properties)
		openapiModelVersion.Author = converter.MapPropertyAuthor((*source).Properties)
		pOpenapiModelVersion = &openapiModelVersion
	}
	return pOpenapiModelVersion, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertRegisteredModel(source *proto.Context) (*openapi.RegisteredModel, error) {
	var pOpenapiRegisteredModel *openapi.RegisteredModel
	if source != nil {
		var openapiRegisteredModel openapi.RegisteredModel
		openapiRegisteredModel.Id = converter.Int64ToString((*source).Id)
		openapiRegisteredModel.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiRegisteredModel.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		var pString *string
		if (*source).Name != nil {
			xstring := *(*source).Name
			pString = &xstring
		}
		openapiRegisteredModel.Name = pString
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, err
		}
		openapiRegisteredModel.CustomProperties = &mapStringOpenapiMetadataValue
		openapiRegisteredModel.Description = converter.MapDescription((*source).Properties)
		var pString2 *string
		if (*source).ExternalId != nil {
			xstring2 := *(*source).ExternalId
			pString2 = &xstring2
		}
		openapiRegisteredModel.ExternalId = pString2
		openapiRegisteredModel.State = converter.MapRegisteredModelState((*source).Properties)
		pOpenapiRegisteredModel = &openapiRegisteredModel
	}
	return pOpenapiRegisteredModel, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertServeModel(source *proto.Execution) (*openapi.ServeModel, error) {
	var pOpenapiServeModel *openapi.ServeModel
	if source != nil {
		var openapiServeModel openapi.ServeModel
		openapiServeModel.LastKnownState = converter.MapMLMDServeModelLastKnownState((*source).LastKnownState)
		openapiServeModel.Id = converter.Int64ToString((*source).Id)
		openapiServeModel.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiServeModel.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		openapiServeModel.Name = converter.MapNameFromOwned((*source).Name)
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, err
		}
		openapiServeModel.CustomProperties = &mapStringOpenapiMetadataValue
		openapiServeModel.Description = converter.MapDescription((*source).Properties)
		var pString *string
		if (*source).ExternalId != nil {
			xstring := *(*source).ExternalId
			pString = &xstring
		}
		openapiServeModel.ExternalId = pString
		openapiServeModel.ModelVersionId = converter.MapPropertyModelVersionIdAsValue((*source).Properties)
		pOpenapiServeModel = &openapiServeModel
	}
	return pOpenapiServeModel, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertServingEnvironment(source *proto.Context) (*openapi.ServingEnvironment, error) {
	var pOpenapiServingEnvironment *openapi.ServingEnvironment
	if source != nil {
		var openapiServingEnvironment openapi.ServingEnvironment
		openapiServingEnvironment.Id = converter.Int64ToString((*source).Id)
		openapiServingEnvironment.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiServingEnvironment.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		openapiServingEnvironment.Name = converter.MapNameFromOwned((*source).Name)
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, err
		}
		openapiServingEnvironment.CustomProperties = &mapStringOpenapiMetadataValue
		openapiServingEnvironment.Description = converter.MapDescription((*source).Properties)
		var pString *string
		if (*source).ExternalId != nil {
			xstring := *(*source).ExternalId
			pString = &xstring
		}
		openapiServingEnvironment.ExternalId = pString
		pOpenapiServingEnvironment = &openapiServingEnvironment
	}
	return pOpenapiServingEnvironment, nil
}
