# \GrafeasApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNote**](GrafeasApi.md#CreateNote) | **Post** /v1alpha1/projects/{projectsId}/notes | 
[**CreateOccurrence**](GrafeasApi.md#CreateOccurrence) | **Post** /v1alpha1/projects/{projectsId}/occurrences | 
[**DeleteNote**](GrafeasApi.md#DeleteNote) | **Delete** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
[**DeleteOccurrence**](GrafeasApi.md#DeleteOccurrence) | **Delete** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
[**GetNote**](GrafeasApi.md#GetNote) | **Get** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
[**GetOccurrence**](GrafeasApi.md#GetOccurrence) | **Get** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
[**GetOccurrenceNote**](GrafeasApi.md#GetOccurrenceNote) | **Get** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}/notes | 
[**GetOperation**](GrafeasApi.md#GetOperation) | **Get** /v1alpha1/projects/{projectsId}/operations/{operationsId} | 
[**ListNoteOccurrences**](GrafeasApi.md#ListNoteOccurrences) | **Get** /v1alpha1/projects/{projectsId}/notes/{notesId}/occurrences | 
[**ListNotes**](GrafeasApi.md#ListNotes) | **Get** /v1alpha1/projects/{projectsId}/notes | 
[**ListOccurrences**](GrafeasApi.md#ListOccurrences) | **Get** /v1alpha1/projects/{projectsId}/occurrences | 
[**ListOperations**](GrafeasApi.md#ListOperations) | **Get** /v1alpha1/projects/{projectsId}/operations | 
[**UpdateNote**](GrafeasApi.md#UpdateNote) | **Put** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
[**UpdateOccurrence**](GrafeasApi.md#UpdateOccurrence) | **Put** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
[**UpdateOperation**](GrafeasApi.md#UpdateOperation) | **Put** /v1alpha1/projects/{projectsId}/operations/{operationsId} | 


# **CreateNote**
> Note CreateNote($projectsId, $noteId, $note)



Creates a new note.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;parent&#x60;. This field contains the projectId for example: \&quot;project/{project_id} | 
 **noteId** | **string**| The ID to use for this note. | [optional] 
 **note** | [**Note**](Note.md)| The Note to be inserted | [optional] 

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOccurrence**
> Occurrence CreateOccurrence($projectsId, $occurrence)



Creates a new occurrence.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;parent&#x60;. This field contains the projectId for example: \&quot;projects/{project_id}\&quot; | 
 **occurrence** | [**Occurrence**](Occurrence.md)| The occurrence to be inserted | [optional] 

### Return type

[**Occurrence**](Occurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNote**
> Empty DeleteNote($projectsId, $notesId)



Deletes the given note from the system.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;name&#x60;. The name of the note in the form \&quot;projects/{project_id}/notes/{note_id}\&quot; | 
 **notesId** | **string**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. | 

### Return type

[**Empty**](Empty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOccurrence**
> Empty DeleteOccurrence($projectsId, $occurrencesId)



Deletes the given occurrence from the system.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;name&#x60;. The name of the occurrence in the form \&quot;projects/{project_id}/occurrences/{occurrence_id}\&quot; | 
 **occurrencesId** | **string**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. | 

### Return type

[**Empty**](Empty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNote**
> Note GetNote($projectsId, $notesId)



Returns the requested occurrence


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;name&#x60;. The name of the note in the form \&quot;projects/{project_id}/notes/{note_id}\&quot; | 
 **notesId** | **string**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. | 

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOccurrence**
> Occurrence GetOccurrence($projectsId, $occurrencesId)



Returns the requested occurrence


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;name&#x60;. The name of the occurrence in the form \&quot;projects/{project_id}/occurrences/{occurrence_id}\&quot; | 
 **occurrencesId** | **string**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. | 

### Return type

[**Occurrence**](Occurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOccurrenceNote**
> Note GetOccurrenceNote($projectsId, $occurrencesId)



Gets the note that this occurrence is attached to.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;name&#x60;. The name of the occurrence in the form \&quot;projects/{project_id}/occurrences/{occurrence_id}\&quot; | 
 **occurrencesId** | **string**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. | 

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOperation**
> Operation GetOperation($projectsId, $operationsId)



Returns the requested occurrence


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;name&#x60;. The name of the operation in the form \&quot;projects/{project_id}/operations/{operation_id}\&quot; | 
 **operationsId** | **string**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. | 

### Return type

[**Operation**](Operation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNoteOccurrences**
> ListNoteOccurrencesResponse ListNoteOccurrences($projectsId, $notesId, $filter, $pageSize, $pageToken)



Lists the names of Occurrences linked to a particular Note.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;name&#x60;. The name field will contain the note name for example:   \&quot;project/{project_id}/notes/{note_id}\&quot; | 
 **notesId** | **string**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. | 
 **filter** | **string**| The filter expression. | [optional] 
 **pageSize** | **int32**| Number of notes to return in the list. | [optional] 
 **pageToken** | **string**| Token to provide to skip to a particular spot in the list. | [optional] 

### Return type

[**ListNoteOccurrencesResponse**](ListNoteOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotes**
> ListNotesResponse ListNotes($projectsId, $filter, $pageSize, $pageToken)



Lists all notes for a given project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;parent&#x60;. This field contains the projectId for example: \&quot;project/{project_id} | 
 **filter** | **string**| The filter expression. | [optional] 
 **pageSize** | **int32**| Number of notes to return in the list. | [optional] 
 **pageToken** | **string**| Token to provide to skip to a particular spot in the list. | [optional] 

### Return type

[**ListNotesResponse**](ListNotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOccurrences**
> ListOccurrencesResponse ListOccurrences($projectsId, $filter, $pageSize, $pageToken)



Lists active occurrences for a given project/Digest.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;parent&#x60;. This contains the projectId for example: projects/{project_id} | 
 **filter** | **string**| The filter expression. | [optional] 
 **pageSize** | **int32**| Number of occurrences to return in the list. | [optional] 
 **pageToken** | **string**| Token to provide to skip to a particular spot in the list. | [optional] 

### Return type

[**ListOccurrencesResponse**](ListOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOperations**
> ListOperationsResponse ListOperations($projectsId, $filter, $pageSize, $pageToken)



Lists all operations for a given project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;parent&#x60;. This field contains the projectId for example: \&quot;project/{project_id} | 
 **filter** | **string**| The filter expression. | [optional] 
 **pageSize** | **int32**| Number of operations to return in the list. | [optional] 
 **pageToken** | **string**| Token to provide to skip to a particular spot in the list. | [optional] 

### Return type

[**ListOperationsResponse**](ListOperationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNote**
> Note UpdateNote($projectsId, $notesId, $note)



Updates an existing note.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;name&#x60;. The name of the note. Should be of the form \&quot;projects/{project_id}/notes/{note_id}\&quot;. | 
 **notesId** | **string**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. | 
 **note** | [**Note**](Note.md)| The updated note. | [optional] 

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOccurrence**
> Occurrence UpdateOccurrence($projectsId, $occurrencesId, $occurrence)



Updates an existing occurrence.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;name&#x60;. The name of the occurrence. Should be of the form \&quot;projects/{project_id}/occurrences/{occurrence_id}\&quot;. | 
 **occurrencesId** | **string**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. | 
 **occurrence** | [**Occurrence**](Occurrence.md)| The updated occurrence. | [optional] 

### Return type

[**Occurrence**](Occurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOperation**
> Operation UpdateOperation($projectsId, $operationsId, $body)



Updates an existing operation returns an error if operation  does not exist. The only valid operations are to update mark the done bit change the result.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **string**| Part of &#x60;name&#x60;. The name of the Operation. Should be of the form \&quot;projects/{project_id}/operations/{operation_id}\&quot;. | 
 **operationsId** | **string**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. | 
 **body** | [**UpdateOperationRequest**](UpdateOperationRequest.md)| The request body. | [optional] 

### Return type

[**Operation**](Operation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

