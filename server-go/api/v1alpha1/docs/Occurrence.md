# Occurrence

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the occurrence in the form \&quot;projects/{project_id}/occurrences/{occurrence_id}\&quot; @OutputOnly | [optional] [default to null]
**ResourceUrl** | **string** | The unique url of the image or container for which the occurrence applies. Example: https://gcr.io/project/image@sha256:foo This field can be used as a filter in list requests. | [optional] [default to null]
**NoteName** | **string** | An analysis note associated with this image, in the form \&quot;projects/{project_id}/notes/{note_id}\&quot; This field can be used as a filter in list requests. | [optional] [default to null]
**Kind** | **string** | This explicitly denotes which of the occurrence details is specified. This field can be used as a filter in list requests. @OutputOnly | [optional] [default to null]
**CustomDetails** | [**CustomDetails**](CustomDetails.md) | Details of the custom note. | [optional] [default to null]
**VulnerabilityDetails** | [**VulnerabilityDetails**](VulnerabilityDetails.md) | Details of a security vulnerability note. | [optional] [default to null]
**BuildDetails** | [**BuildDetails**](BuildDetails.md) | Build details for a verifiable build. | [optional] [default to null]
**DerivedImage** | [**Derived**](Derived.md) | Describes how this resource derives from the basis in the associated note. | [optional] [default to null]
**Installation** | [**Installation**](Installation.md) | Describes the installation of a package on the linked resource. | [optional] [default to null]
**Deployment** | [**Deployment**](Deployment.md) | Describes the deployment of an artifact on a runtime. | [optional] [default to null]
**Discovered** | [**Discovered**](Discovered.md) | Describes the initial scan status for this resource. | [optional] [default to null]
**Attestation** | [**Attestation**](Attestation.md) | Describes an attestation of an artifact. | [optional] [default to null]
**Remediation** | **string** | A description of actions that can be taken to remedy the note | [optional] [default to null]
**CreateTime** | **string** | The time this occurrence was created. @OutputOnly | [optional] [default to null]
**UpdateTime** | **string** | The time this occurrence was last updated. @OutputOnly | [optional] [default to null]
**OperationName** | **string** | The name of the operation that created this note. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


