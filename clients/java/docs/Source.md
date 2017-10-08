
# Source

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**storageSource** | [**StorageSource**](StorageSource.md) | If provided, get the source from this location in in Google Cloud Storage. |  [optional]
**repoSource** | [**RepoSource**](RepoSource.md) | If provided, get source from this location in a Cloud Repo. |  [optional]
**artifactStorageSource** | [**StorageSource**](StorageSource.md) | If provided, the input binary artifacts for the build came from this location. |  [optional]
**sourceContext** | [**ExtendedSourceContext**](ExtendedSourceContext.md) | If provided, the source code used for the build came from this location. |  [optional]
**additionalSourceContexts** | [**List&lt;ExtendedSourceContext&gt;**](ExtendedSourceContext.md) | If provided, some of the source code used for the build may be found in these locations, in the case where the source repository had multiple remotes or submodules. This list will not include the context specified in the source_context field. |  [optional]
**fileHashes** | [**Map&lt;String, FileHashes&gt;**](FileHashes.md) | Hash(es) of the build source, which can be used to verify that the original source integrity was maintained in the build.  The keys to this map are file paths used as build source and the values contain the hash values for those files.  If the build source came in a single package such as a gzipped tarfile (.tar.gz), the FileHash will be for the single path to that file. |  [optional]



