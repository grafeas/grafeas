
# Note

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | The name of the note in the form \&quot;projects/{project_id}/notes/{note_id}\&quot; |  [optional]
**shortDescription** | **String** | A one sentence description of this note |  [optional]
**longDescription** | **String** | A detailed description of this note |  [optional]
**kind** | [**KindEnum**](#KindEnum) | This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests. @OutputOnly |  [optional]
**vulnerabilityType** | [**VulnerabilityType**](VulnerabilityType.md) | A package vulnerability type of note. |  [optional]
**buildType** | [**BuildType**](BuildType.md) | Build provenance type for a verifiable build. |  [optional]
**baseImage** | [**Basis**](Basis.md) | A note describing a base image. |  [optional]
**_package** | [**ModelPackage**](ModelPackage.md) | A note describing a package hosted by various package managers. |  [optional]
**deployable** | [**Deployable**](Deployable.md) | A note describing something that can be deployed. |  [optional]
**discovery** | [**Discovery**](Discovery.md) | A note describing a project/analysis type. |  [optional]
**relatedUrl** | [**List&lt;RelatedUrl&gt;**](RelatedUrl.md) | Urls associated with this note |  [optional]
**expirationTime** | **String** | Time of expiration for this Note, null if Note currently does not expire. |  [optional]
**createTime** | **String** | The time this note was created. This field can be used as a filter in list requests. @OutputOnly |  [optional]
**updateTime** | **String** | The time this note was last updated. This field can be used as a filter in list requests. @OutputOnly |  [optional]
**operationName** | **String** | The name of the operation that created this note. |  [optional]
**relatedNoteNames** | **List&lt;String&gt;** | Other notes related to this note. |  [optional]


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



