# Note

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the note in the form \&quot;projects/{project_id}/notes/{note_id}\&quot; | [optional] [default to null]
**ShortDescription** | **string** | A one sentence description of this note | [optional] [default to null]
**LongDescription** | **string** | A detailed description of this note | [optional] [default to null]
**Kind** | **string** | This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests. @OutputOnly | [optional] [default to null]
**VulnerabilityType** | [**VulnerabilityType**](VulnerabilityType.md) | A package vulnerability type of note. | [optional] [default to null]
**BuildType** | [**BuildType**](BuildType.md) | Build provenance type for a verifiable build. | [optional] [default to null]
**BaseImage** | [**Basis**](Basis.md) | A note describing a base image. | [optional] [default to null]
**Package_** | [**ModelPackage**](Package.md) | A note describing a package hosted by various package managers. | [optional] [default to null]
**Deployable** | [**Deployable**](Deployable.md) | A note describing something that can be deployed. | [optional] [default to null]
**Discovery** | [**Discovery**](Discovery.md) | A note describing a project/analysis type. | [optional] [default to null]
**AttestationAuthority** | [**AttestationAuthority**](AttestationAuthority.md) | A note describing an attestation role. | [optional] [default to null]
**RelatedUrl** | [**[]RelatedUrl**](RelatedUrl.md) | Urls associated with this note | [optional] [default to null]
**ExpirationTime** | **string** | Time of expiration for this Note, null if Note currently does not expire. | [optional] [default to null]
**CreateTime** | **string** | The time this note was created. This field can be used as a filter in list requests. @OutputOnly | [optional] [default to null]
**UpdateTime** | **string** | The time this note was last updated. This field can be used as a filter in list requests. @OutputOnly | [optional] [default to null]
**OperationName** | **string** | The name of the operation that created this note. | [optional] [default to null]
**RelatedNoteNames** | **[]string** | Other notes related to this note. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


