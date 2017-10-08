
# BuildProvenance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **String** | Unique identifier of the build. |  [optional]
**projectId** | **String** | ID of the project. |  [optional]
**projectNum** | **String** | Numerical ID of the project. |  [optional]
**commands** | [**List&lt;Command&gt;**](Command.md) | Commands requested by the build. |  [optional]
**builtArtifacts** | [**List&lt;Artifact&gt;**](Artifact.md) | Output of the build. |  [optional]
**createTime** | **String** | Time at which the build was created. |  [optional]
**startTime** | **String** | Time at which execution of the build was started. |  [optional]
**finishTime** | **String** | Time at whihc execution of the build was finished. |  [optional]
**userId** | **String** | GAIA ID of end user who initiated this build; at the time that the BuildProvenance is uploaded to Analysis, this will be resolved to the primary e-mail address of the user and stored in the Creator field. |  [optional]
**creator** | **String** | E-mail address of the user who initiated this build. Note that this was the user&#39;s e-mail address at the time the build was initiated; this address may not represent the same end-user for all time. |  [optional]
**logsBucket** | **String** | Google Cloud Storage bucket where logs were written. |  [optional]
**sourceProvenance** | [**Source**](Source.md) | Details of the Source input to the build. |  [optional]
**triggerId** | **String** | Trigger identifier if the build was triggered automatically; empty if not. |  [optional]
**buildOptions** | **Map&lt;String, String&gt;** | Special options applied to this build. This is a catch-all field where build providers can enter any desired additional details. |  [optional]
**builderVersion** | **String** | Version string of the builder at the time this build was executed. |  [optional]



