
# Derived

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fingerprint** | [**Fingerprint**](Fingerprint.md) | The fingerprint of the derived image |  [optional]
**distance** | **Integer** | The number of layers by which this image differs from the associated image basis. @OutputOnly |  [optional]
**layerInfo** | [**List&lt;Layer&gt;**](Layer.md) | This contains layer-specific metadata, if populated it has length “distance” and is ordered with [distance] being the layer immediately following the base image and [1] being the final layer. |  [optional]
**baseResourceUrl** | **String** | This contains the base image url for the derived image Occurrence @OutputOnly |  [optional]



