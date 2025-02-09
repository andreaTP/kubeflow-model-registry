// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package generated

import (
	converter "github.com/kubeflow/model-registry/internal/converter"
	openapi "github.com/kubeflow/model-registry/pkg/openapi"
)

type OpenAPIConverterImpl struct{}

func (c *OpenAPIConverterImpl) ConvertInferenceServiceCreate(source *openapi.InferenceServiceCreate) (*openapi.InferenceService, error) {
	var pOpenapiInferenceService *openapi.InferenceService
	if source != nil {
		var openapiInferenceService openapi.InferenceService
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiInferenceService.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiInferenceService.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiInferenceService.ExternalID = pString2
		var pString3 *string
		if (*source).Name != nil {
			xstring3 := *(*source).Name
			pString3 = &xstring3
		}
		openapiInferenceService.Name = pString3
		var pString4 *string
		if (*source).ModelVersionId != nil {
			xstring4 := *(*source).ModelVersionId
			pString4 = &xstring4
		}
		openapiInferenceService.ModelVersionId = pString4
		var pString5 *string
		if (*source).Runtime != nil {
			xstring5 := *(*source).Runtime
			pString5 = &xstring5
		}
		openapiInferenceService.Runtime = pString5
		var pOpenapiInferenceServiceState *openapi.InferenceServiceState
		if (*source).DesiredState != nil {
			openapiInferenceServiceState := openapi.InferenceServiceState(*(*source).DesiredState)
			pOpenapiInferenceServiceState = &openapiInferenceServiceState
		}
		openapiInferenceService.DesiredState = pOpenapiInferenceServiceState
		openapiInferenceService.RegisteredModelId = (*source).RegisteredModelId
		openapiInferenceService.ServingEnvironmentId = (*source).ServingEnvironmentId
		pOpenapiInferenceService = &openapiInferenceService
	}
	return pOpenapiInferenceService, nil
}
func (c *OpenAPIConverterImpl) ConvertInferenceServiceUpdate(source *openapi.InferenceServiceUpdate) (*openapi.InferenceService, error) {
	var pOpenapiInferenceService *openapi.InferenceService
	if source != nil {
		var openapiInferenceService openapi.InferenceService
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiInferenceService.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiInferenceService.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiInferenceService.ExternalID = pString2
		var pString3 *string
		if (*source).ModelVersionId != nil {
			xstring3 := *(*source).ModelVersionId
			pString3 = &xstring3
		}
		openapiInferenceService.ModelVersionId = pString3
		var pString4 *string
		if (*source).Runtime != nil {
			xstring4 := *(*source).Runtime
			pString4 = &xstring4
		}
		openapiInferenceService.Runtime = pString4
		var pOpenapiInferenceServiceState *openapi.InferenceServiceState
		if (*source).DesiredState != nil {
			openapiInferenceServiceState := openapi.InferenceServiceState(*(*source).DesiredState)
			pOpenapiInferenceServiceState = &openapiInferenceServiceState
		}
		openapiInferenceService.DesiredState = pOpenapiInferenceServiceState
		pOpenapiInferenceService = &openapiInferenceService
	}
	return pOpenapiInferenceService, nil
}
func (c *OpenAPIConverterImpl) ConvertModelArtifactCreate(source *openapi.ModelArtifactCreate) (*openapi.ModelArtifact, error) {
	var pOpenapiModelArtifact *openapi.ModelArtifact
	if source != nil {
		var openapiModelArtifact openapi.ModelArtifact
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiModelArtifact.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiModelArtifact.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiModelArtifact.ExternalID = pString2
		var pString3 *string
		if (*source).Uri != nil {
			xstring3 := *(*source).Uri
			pString3 = &xstring3
		}
		openapiModelArtifact.Uri = pString3
		var pOpenapiArtifactState *openapi.ArtifactState
		if (*source).State != nil {
			openapiArtifactState := openapi.ArtifactState(*(*source).State)
			pOpenapiArtifactState = &openapiArtifactState
		}
		openapiModelArtifact.State = pOpenapiArtifactState
		var pString4 *string
		if (*source).Name != nil {
			xstring4 := *(*source).Name
			pString4 = &xstring4
		}
		openapiModelArtifact.Name = pString4
		var pString5 *string
		if (*source).ModelFormatName != nil {
			xstring5 := *(*source).ModelFormatName
			pString5 = &xstring5
		}
		openapiModelArtifact.ModelFormatName = pString5
		var pString6 *string
		if (*source).StorageKey != nil {
			xstring6 := *(*source).StorageKey
			pString6 = &xstring6
		}
		openapiModelArtifact.StorageKey = pString6
		var pString7 *string
		if (*source).StoragePath != nil {
			xstring7 := *(*source).StoragePath
			pString7 = &xstring7
		}
		openapiModelArtifact.StoragePath = pString7
		var pString8 *string
		if (*source).ModelFormatVersion != nil {
			xstring8 := *(*source).ModelFormatVersion
			pString8 = &xstring8
		}
		openapiModelArtifact.ModelFormatVersion = pString8
		var pString9 *string
		if (*source).ServiceAccountName != nil {
			xstring9 := *(*source).ServiceAccountName
			pString9 = &xstring9
		}
		openapiModelArtifact.ServiceAccountName = pString9
		pOpenapiModelArtifact = &openapiModelArtifact
	}
	return pOpenapiModelArtifact, nil
}
func (c *OpenAPIConverterImpl) ConvertModelArtifactUpdate(source *openapi.ModelArtifactUpdate) (*openapi.ModelArtifact, error) {
	var pOpenapiModelArtifact *openapi.ModelArtifact
	if source != nil {
		var openapiModelArtifact openapi.ModelArtifact
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiModelArtifact.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiModelArtifact.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiModelArtifact.ExternalID = pString2
		var pString3 *string
		if (*source).Uri != nil {
			xstring3 := *(*source).Uri
			pString3 = &xstring3
		}
		openapiModelArtifact.Uri = pString3
		var pOpenapiArtifactState *openapi.ArtifactState
		if (*source).State != nil {
			openapiArtifactState := openapi.ArtifactState(*(*source).State)
			pOpenapiArtifactState = &openapiArtifactState
		}
		openapiModelArtifact.State = pOpenapiArtifactState
		var pString4 *string
		if (*source).ModelFormatName != nil {
			xstring4 := *(*source).ModelFormatName
			pString4 = &xstring4
		}
		openapiModelArtifact.ModelFormatName = pString4
		var pString5 *string
		if (*source).StorageKey != nil {
			xstring5 := *(*source).StorageKey
			pString5 = &xstring5
		}
		openapiModelArtifact.StorageKey = pString5
		var pString6 *string
		if (*source).StoragePath != nil {
			xstring6 := *(*source).StoragePath
			pString6 = &xstring6
		}
		openapiModelArtifact.StoragePath = pString6
		var pString7 *string
		if (*source).ModelFormatVersion != nil {
			xstring7 := *(*source).ModelFormatVersion
			pString7 = &xstring7
		}
		openapiModelArtifact.ModelFormatVersion = pString7
		var pString8 *string
		if (*source).ServiceAccountName != nil {
			xstring8 := *(*source).ServiceAccountName
			pString8 = &xstring8
		}
		openapiModelArtifact.ServiceAccountName = pString8
		pOpenapiModelArtifact = &openapiModelArtifact
	}
	return pOpenapiModelArtifact, nil
}
func (c *OpenAPIConverterImpl) ConvertModelVersionCreate(source *openapi.ModelVersionCreate) (*openapi.ModelVersion, error) {
	var pOpenapiModelVersion *openapi.ModelVersion
	if source != nil {
		var openapiModelVersion openapi.ModelVersion
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiModelVersion.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiModelVersion.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiModelVersion.ExternalID = pString2
		var pString3 *string
		if (*source).Name != nil {
			xstring3 := *(*source).Name
			pString3 = &xstring3
		}
		openapiModelVersion.Name = pString3
		var pOpenapiModelVersionState *openapi.ModelVersionState
		if (*source).State != nil {
			openapiModelVersionState := openapi.ModelVersionState(*(*source).State)
			pOpenapiModelVersionState = &openapiModelVersionState
		}
		openapiModelVersion.State = pOpenapiModelVersionState
		var pString4 *string
		if (*source).Author != nil {
			xstring4 := *(*source).Author
			pString4 = &xstring4
		}
		openapiModelVersion.Author = pString4
		pOpenapiModelVersion = &openapiModelVersion
	}
	return pOpenapiModelVersion, nil
}
func (c *OpenAPIConverterImpl) ConvertModelVersionUpdate(source *openapi.ModelVersionUpdate) (*openapi.ModelVersion, error) {
	var pOpenapiModelVersion *openapi.ModelVersion
	if source != nil {
		var openapiModelVersion openapi.ModelVersion
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiModelVersion.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiModelVersion.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiModelVersion.ExternalID = pString2
		var pOpenapiModelVersionState *openapi.ModelVersionState
		if (*source).State != nil {
			openapiModelVersionState := openapi.ModelVersionState(*(*source).State)
			pOpenapiModelVersionState = &openapiModelVersionState
		}
		openapiModelVersion.State = pOpenapiModelVersionState
		var pString3 *string
		if (*source).Author != nil {
			xstring3 := *(*source).Author
			pString3 = &xstring3
		}
		openapiModelVersion.Author = pString3
		pOpenapiModelVersion = &openapiModelVersion
	}
	return pOpenapiModelVersion, nil
}
func (c *OpenAPIConverterImpl) ConvertRegisteredModelCreate(source *openapi.RegisteredModelCreate) (*openapi.RegisteredModel, error) {
	var pOpenapiRegisteredModel *openapi.RegisteredModel
	if source != nil {
		var openapiRegisteredModel openapi.RegisteredModel
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiRegisteredModel.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiRegisteredModel.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiRegisteredModel.ExternalID = pString2
		var pString3 *string
		if (*source).Name != nil {
			xstring3 := *(*source).Name
			pString3 = &xstring3
		}
		openapiRegisteredModel.Name = pString3
		var pOpenapiRegisteredModelState *openapi.RegisteredModelState
		if (*source).State != nil {
			openapiRegisteredModelState := openapi.RegisteredModelState(*(*source).State)
			pOpenapiRegisteredModelState = &openapiRegisteredModelState
		}
		openapiRegisteredModel.State = pOpenapiRegisteredModelState
		pOpenapiRegisteredModel = &openapiRegisteredModel
	}
	return pOpenapiRegisteredModel, nil
}
func (c *OpenAPIConverterImpl) ConvertRegisteredModelUpdate(source *openapi.RegisteredModelUpdate) (*openapi.RegisteredModel, error) {
	var pOpenapiRegisteredModel *openapi.RegisteredModel
	if source != nil {
		var openapiRegisteredModel openapi.RegisteredModel
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiRegisteredModel.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiRegisteredModel.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiRegisteredModel.ExternalID = pString2
		var pOpenapiRegisteredModelState *openapi.RegisteredModelState
		if (*source).State != nil {
			openapiRegisteredModelState := openapi.RegisteredModelState(*(*source).State)
			pOpenapiRegisteredModelState = &openapiRegisteredModelState
		}
		openapiRegisteredModel.State = pOpenapiRegisteredModelState
		pOpenapiRegisteredModel = &openapiRegisteredModel
	}
	return pOpenapiRegisteredModel, nil
}
func (c *OpenAPIConverterImpl) ConvertServeModelCreate(source *openapi.ServeModelCreate) (*openapi.ServeModel, error) {
	var pOpenapiServeModel *openapi.ServeModel
	if source != nil {
		var openapiServeModel openapi.ServeModel
		var pOpenapiExecutionState *openapi.ExecutionState
		if (*source).LastKnownState != nil {
			openapiExecutionState := openapi.ExecutionState(*(*source).LastKnownState)
			pOpenapiExecutionState = &openapiExecutionState
		}
		openapiServeModel.LastKnownState = pOpenapiExecutionState
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiServeModel.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiServeModel.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiServeModel.ExternalID = pString2
		var pString3 *string
		if (*source).Name != nil {
			xstring3 := *(*source).Name
			pString3 = &xstring3
		}
		openapiServeModel.Name = pString3
		openapiServeModel.ModelVersionId = (*source).ModelVersionId
		pOpenapiServeModel = &openapiServeModel
	}
	return pOpenapiServeModel, nil
}
func (c *OpenAPIConverterImpl) ConvertServeModelUpdate(source *openapi.ServeModelUpdate) (*openapi.ServeModel, error) {
	var pOpenapiServeModel *openapi.ServeModel
	if source != nil {
		var openapiServeModel openapi.ServeModel
		var pOpenapiExecutionState *openapi.ExecutionState
		if (*source).LastKnownState != nil {
			openapiExecutionState := openapi.ExecutionState(*(*source).LastKnownState)
			pOpenapiExecutionState = &openapiExecutionState
		}
		openapiServeModel.LastKnownState = pOpenapiExecutionState
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiServeModel.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiServeModel.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiServeModel.ExternalID = pString2
		pOpenapiServeModel = &openapiServeModel
	}
	return pOpenapiServeModel, nil
}
func (c *OpenAPIConverterImpl) ConvertServingEnvironmentCreate(source *openapi.ServingEnvironmentCreate) (*openapi.ServingEnvironment, error) {
	var pOpenapiServingEnvironment *openapi.ServingEnvironment
	if source != nil {
		var openapiServingEnvironment openapi.ServingEnvironment
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiServingEnvironment.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiServingEnvironment.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiServingEnvironment.ExternalID = pString2
		var pString3 *string
		if (*source).Name != nil {
			xstring3 := *(*source).Name
			pString3 = &xstring3
		}
		openapiServingEnvironment.Name = pString3
		pOpenapiServingEnvironment = &openapiServingEnvironment
	}
	return pOpenapiServingEnvironment, nil
}
func (c *OpenAPIConverterImpl) ConvertServingEnvironmentUpdate(source *openapi.ServingEnvironmentUpdate) (*openapi.ServingEnvironment, error) {
	var pOpenapiServingEnvironment *openapi.ServingEnvironment
	if source != nil {
		var openapiServingEnvironment openapi.ServingEnvironment
		var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
		if (*source).CustomProperties != nil {
			mapStringOpenapiMetadataValue := make(map[string]openapi.MetadataValue, len((*(*source).CustomProperties)))
			for key, value := range *(*source).CustomProperties {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
			pMapStringOpenapiMetadataValue = &mapStringOpenapiMetadataValue
		}
		openapiServingEnvironment.CustomProperties = pMapStringOpenapiMetadataValue
		var pString *string
		if (*source).Description != nil {
			xstring := *(*source).Description
			pString = &xstring
		}
		openapiServingEnvironment.Description = pString
		var pString2 *string
		if (*source).ExternalID != nil {
			xstring2 := *(*source).ExternalID
			pString2 = &xstring2
		}
		openapiServingEnvironment.ExternalID = pString2
		pOpenapiServingEnvironment = &openapiServingEnvironment
	}
	return pOpenapiServingEnvironment, nil
}
func (c *OpenAPIConverterImpl) OverrideNotEditableForDocArtifact(source converter.OpenapiUpdateWrapper[openapi.DocArtifact]) (openapi.DocArtifact, error) {
	openapiDocArtifact := converter.InitDocArtifactWithUpdate(source)
	_ = source
	return openapiDocArtifact, nil
}
func (c *OpenAPIConverterImpl) OverrideNotEditableForInferenceService(source converter.OpenapiUpdateWrapper[openapi.InferenceService]) (openapi.InferenceService, error) {
	openapiInferenceService := converter.InitInferenceServiceWithUpdate(source)
	var pString *string
	if source.Existing != nil {
		pString = source.Existing.Name
	}
	var pString2 *string
	if pString != nil {
		xstring := *pString
		pString2 = &xstring
	}
	openapiInferenceService.Name = pString2
	var pString3 *string
	if source.Existing != nil {
		pString3 = &source.Existing.RegisteredModelId
	}
	var xstring2 string
	if pString3 != nil {
		xstring2 = *pString3
	}
	openapiInferenceService.RegisteredModelId = xstring2
	var pString4 *string
	if source.Existing != nil {
		pString4 = &source.Existing.ServingEnvironmentId
	}
	var xstring3 string
	if pString4 != nil {
		xstring3 = *pString4
	}
	openapiInferenceService.ServingEnvironmentId = xstring3
	return openapiInferenceService, nil
}
func (c *OpenAPIConverterImpl) OverrideNotEditableForModelArtifact(source converter.OpenapiUpdateWrapper[openapi.ModelArtifact]) (openapi.ModelArtifact, error) {
	openapiModelArtifact := converter.InitModelArtifactWithUpdate(source)
	var pString *string
	if source.Existing != nil {
		pString = &source.Existing.ArtifactType
	}
	var xstring string
	if pString != nil {
		xstring = *pString
	}
	openapiModelArtifact.ArtifactType = xstring
	var pString2 *string
	if source.Existing != nil {
		pString2 = source.Existing.Name
	}
	var pString3 *string
	if pString2 != nil {
		xstring2 := *pString2
		pString3 = &xstring2
	}
	openapiModelArtifact.Name = pString3
	return openapiModelArtifact, nil
}
func (c *OpenAPIConverterImpl) OverrideNotEditableForModelVersion(source converter.OpenapiUpdateWrapper[openapi.ModelVersion]) (openapi.ModelVersion, error) {
	openapiModelVersion := converter.InitModelVersionWithUpdate(source)
	var pString *string
	if source.Existing != nil {
		pString = source.Existing.Name
	}
	var pString2 *string
	if pString != nil {
		xstring := *pString
		pString2 = &xstring
	}
	openapiModelVersion.Name = pString2
	return openapiModelVersion, nil
}
func (c *OpenAPIConverterImpl) OverrideNotEditableForRegisteredModel(source converter.OpenapiUpdateWrapper[openapi.RegisteredModel]) (openapi.RegisteredModel, error) {
	openapiRegisteredModel := converter.InitRegisteredModelWithUpdate(source)
	var pString *string
	if source.Existing != nil {
		pString = source.Existing.Name
	}
	var pString2 *string
	if pString != nil {
		xstring := *pString
		pString2 = &xstring
	}
	openapiRegisteredModel.Name = pString2
	return openapiRegisteredModel, nil
}
func (c *OpenAPIConverterImpl) OverrideNotEditableForServeModel(source converter.OpenapiUpdateWrapper[openapi.ServeModel]) (openapi.ServeModel, error) {
	openapiServeModel := converter.InitServeModelWithUpdate(source)
	var pString *string
	if source.Existing != nil {
		pString = source.Existing.Name
	}
	var pString2 *string
	if pString != nil {
		xstring := *pString
		pString2 = &xstring
	}
	openapiServeModel.Name = pString2
	var pString3 *string
	if source.Existing != nil {
		pString3 = &source.Existing.ModelVersionId
	}
	var xstring2 string
	if pString3 != nil {
		xstring2 = *pString3
	}
	openapiServeModel.ModelVersionId = xstring2
	return openapiServeModel, nil
}
func (c *OpenAPIConverterImpl) OverrideNotEditableForServingEnvironment(source converter.OpenapiUpdateWrapper[openapi.ServingEnvironment]) (openapi.ServingEnvironment, error) {
	openapiServingEnvironment := converter.InitServingEnvironmentWithUpdate(source)
	var pString *string
	if source.Existing != nil {
		pString = source.Existing.Name
	}
	var pString2 *string
	if pString != nil {
		xstring := *pString
		pString2 = &xstring
	}
	openapiServingEnvironment.Name = pString2
	return openapiServingEnvironment, nil
}
func (c *OpenAPIConverterImpl) openapiMetadataValueToOpenapiMetadataValue(source openapi.MetadataValue) openapi.MetadataValue {
	var openapiMetadataValue openapi.MetadataValue
	openapiMetadataValue.MetadataBoolValue = c.pOpenapiMetadataBoolValueToPOpenapiMetadataBoolValue(source.MetadataBoolValue)
	openapiMetadataValue.MetadataDoubleValue = c.pOpenapiMetadataDoubleValueToPOpenapiMetadataDoubleValue(source.MetadataDoubleValue)
	openapiMetadataValue.MetadataIntValue = c.pOpenapiMetadataIntValueToPOpenapiMetadataIntValue(source.MetadataIntValue)
	openapiMetadataValue.MetadataProtoValue = c.pOpenapiMetadataProtoValueToPOpenapiMetadataProtoValue(source.MetadataProtoValue)
	openapiMetadataValue.MetadataStringValue = c.pOpenapiMetadataStringValueToPOpenapiMetadataStringValue(source.MetadataStringValue)
	openapiMetadataValue.MetadataStructValue = c.pOpenapiMetadataStructValueToPOpenapiMetadataStructValue(source.MetadataStructValue)
	return openapiMetadataValue
}
func (c *OpenAPIConverterImpl) pOpenapiMetadataBoolValueToPOpenapiMetadataBoolValue(source *openapi.MetadataBoolValue) *openapi.MetadataBoolValue {
	var pOpenapiMetadataBoolValue *openapi.MetadataBoolValue
	if source != nil {
		var openapiMetadataBoolValue openapi.MetadataBoolValue
		openapiMetadataBoolValue.BoolValue = (*source).BoolValue
		openapiMetadataBoolValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataBoolValue = &openapiMetadataBoolValue
	}
	return pOpenapiMetadataBoolValue
}
func (c *OpenAPIConverterImpl) pOpenapiMetadataDoubleValueToPOpenapiMetadataDoubleValue(source *openapi.MetadataDoubleValue) *openapi.MetadataDoubleValue {
	var pOpenapiMetadataDoubleValue *openapi.MetadataDoubleValue
	if source != nil {
		var openapiMetadataDoubleValue openapi.MetadataDoubleValue
		openapiMetadataDoubleValue.DoubleValue = (*source).DoubleValue
		openapiMetadataDoubleValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataDoubleValue = &openapiMetadataDoubleValue
	}
	return pOpenapiMetadataDoubleValue
}
func (c *OpenAPIConverterImpl) pOpenapiMetadataIntValueToPOpenapiMetadataIntValue(source *openapi.MetadataIntValue) *openapi.MetadataIntValue {
	var pOpenapiMetadataIntValue *openapi.MetadataIntValue
	if source != nil {
		var openapiMetadataIntValue openapi.MetadataIntValue
		openapiMetadataIntValue.IntValue = (*source).IntValue
		openapiMetadataIntValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataIntValue = &openapiMetadataIntValue
	}
	return pOpenapiMetadataIntValue
}
func (c *OpenAPIConverterImpl) pOpenapiMetadataProtoValueToPOpenapiMetadataProtoValue(source *openapi.MetadataProtoValue) *openapi.MetadataProtoValue {
	var pOpenapiMetadataProtoValue *openapi.MetadataProtoValue
	if source != nil {
		var openapiMetadataProtoValue openapi.MetadataProtoValue
		openapiMetadataProtoValue.Type = (*source).Type
		openapiMetadataProtoValue.ProtoValue = (*source).ProtoValue
		openapiMetadataProtoValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataProtoValue = &openapiMetadataProtoValue
	}
	return pOpenapiMetadataProtoValue
}
func (c *OpenAPIConverterImpl) pOpenapiMetadataStringValueToPOpenapiMetadataStringValue(source *openapi.MetadataStringValue) *openapi.MetadataStringValue {
	var pOpenapiMetadataStringValue *openapi.MetadataStringValue
	if source != nil {
		var openapiMetadataStringValue openapi.MetadataStringValue
		openapiMetadataStringValue.StringValue = (*source).StringValue
		openapiMetadataStringValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataStringValue = &openapiMetadataStringValue
	}
	return pOpenapiMetadataStringValue
}
func (c *OpenAPIConverterImpl) pOpenapiMetadataStructValueToPOpenapiMetadataStructValue(source *openapi.MetadataStructValue) *openapi.MetadataStructValue {
	var pOpenapiMetadataStructValue *openapi.MetadataStructValue
	if source != nil {
		var openapiMetadataStructValue openapi.MetadataStructValue
		openapiMetadataStructValue.StructValue = (*source).StructValue
		openapiMetadataStructValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataStructValue = &openapiMetadataStructValue
	}
	return pOpenapiMetadataStructValue
}
