
# GerritSourceContext

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**hostUri** | **String** | The URI of a running Gerrit instance. |  [optional]
**gerritProject** | **String** | The full project name within the host. Projects may be nested, so \&quot;project/subproject\&quot; is a valid project name. The \&quot;repo name\&quot; is hostURI/project. |  [optional]
**revisionId** | **String** | A revision (commit) ID. |  [optional]
**aliasName** | **String** | The name of an alias (branch, tag, etc.). |  [optional]
**aliasContext** | [**AliasContext**](AliasContext.md) | An alias, which may be a branch or tag. |  [optional]



