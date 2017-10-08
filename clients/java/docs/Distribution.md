
# Distribution

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpeUri** | **String** | The cpe_uri in [cpe format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package. |  [optional]
**architecture** | [**ArchitectureEnum**](#ArchitectureEnum) | The CPU architecture for which packages in this distribution channel were built |  [optional]
**latestVersion** | [**Version**](Version.md) | The latest available version of this package in this distribution channel. |  [optional]
**maintainer** | **String** | A freeform string denoting the maintainer of this package. |  [optional]
**url** | **String** | The distribution channel-specific homepage for this package. |  [optional]
**description** | **String** | The distribution channel-specific description of this package. |  [optional]


<a name="ArchitectureEnum"></a>
## Enum: ArchitectureEnum
Name | Value
---- | -----
UNKNOWN | &quot;UNKNOWN&quot;
X86 | &quot;X86&quot;
X64 | &quot;X64&quot;



