
# Detail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpeUri** | **String** | The cpe_uri in [cpe format] (https://cpe.mitre.org/specification/) in which the vulnerability manifests.  Examples include distro or storage location for vulnerable jar. This field can be used as a filter in list requests. |  [optional]
**_package** | **String** | The name of the package where the vulnerability was found. This field can be used as a filter in list requests. |  [optional]
**minAffectedVersion** | [**Version**](Version.md) | The min version of the package in which the vulnerability exists. |  [optional]
**maxAffectedVersion** | [**Version**](Version.md) | The max version of the package in which the vulnerability exists. This field can be used as a filter in list requests. |  [optional]
**severityName** | **String** | The severity (eg: distro assigned severity) for this vulnerability. |  [optional]
**description** | **String** | A vendor-specific description of this note. |  [optional]
**fixedLocation** | [**VulnerabilityLocation**](VulnerabilityLocation.md) | The fix for this specific package version. |  [optional]



