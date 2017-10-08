
# Occurrence

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | The name of the occurrence in the form \&quot;projects/{project_id}/occurrences/{occurrence_id}\&quot; @OutputOnly |  [optional]
**resourceUrl** | **String** | The unique url of the image or container for which the occurrence applies. Example: https://gcr.io/project/image@sha256:foo This field can be used as a filter in list requests. |  [optional]
**noteName** | **String** | An analysis note associated with this image, in the form \&quot;projects/{project_id}/notes/{note_id}\&quot; This field can be used as a filter in list requests. |  [optional]
**kind** | [**KindEnum**](#KindEnum) | This explicitly denotes which of the occurrence details is specified. This field can be used as a filter in list requests. @OutputOnly |  [optional]
**customDetails** | [**CustomDetails**](CustomDetails.md) | Details of the custom note. |  [optional]
**vulnerabilityDetails** | [**VulnerabilityDetails**](VulnerabilityDetails.md) | Details of a security vulnerability note. |  [optional]
**buildDetails** | [**BuildDetails**](BuildDetails.md) | Build details for a verifiable build. |  [optional]
**derivedImage** | [**Derived**](Derived.md) | Describes how this resource derives from the basis in the associated note. |  [optional]
**installation** | [**Installation**](Installation.md) | Describes the installation of a package on the linked resource. |  [optional]
**deployment** | [**Deployment**](Deployment.md) | Describes the deployment of an artifact on a runtime. |  [optional]
**discovered** | [**Discovered**](Discovered.md) | Describes the initial scan status for this resource. |  [optional]
**remediation** | **String** | A description of actions that can be taken to remedy the note |  [optional]
**createTime** | **String** | The time this occurrence was created. @OutputOnly |  [optional]
**updateTime** | **String** | The time this occurrence was last updated. @OutputOnly |  [optional]
**operationName** | **String** | The name of the operation that created this note. |  [optional]


<a name="KindEnum"></a>
## Enum: KindEnum
Name | Value
---- | -----
CUSTOM | &quot;CUSTOM&quot;
PACKAGE_VULNERABILITY | &quot;PACKAGE_VULNERABILITY&quot;
BUILD_DETAILS | &quot;BUILD_DETAILS&quot;
IMAGE_BASIS | &quot;IMAGE_BASIS&quot;
PACKAGE_MANAGER | &quot;PACKAGE_MANAGER&quot;
DEPLOYABLE | &quot;DEPLOYABLE&quot;
DISCOVERY | &quot;DISCOVERY&quot;



