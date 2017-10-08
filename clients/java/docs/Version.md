
# Version

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**epoch** | **Integer** | Used to correct mistakes in the version numbering scheme. |  [optional]
**name** | **String** | The main part of the version name. |  [optional]
**revision** | **String** | The iteration of the package build from the above version. |  [optional]
**kind** | [**KindEnum**](#KindEnum) | Distinguish between sentinel MIN/MAX versions and normal versions. If kind is not NORMAL, then the other fields are ignored. |  [optional]


<a name="KindEnum"></a>
## Enum: KindEnum
Name | Value
---- | -----
NORMAL | &quot;NORMAL&quot;
MINIMUM | &quot;MINIMUM&quot;
MAXIMUM | &quot;MAXIMUM&quot;



