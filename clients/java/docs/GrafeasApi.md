# GrafeasApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createNote**](GrafeasApi.md#createNote) | **POST** /v1alpha1/projects/{projectsId}/notes | 
[**createOccurrence**](GrafeasApi.md#createOccurrence) | **POST** /v1alpha1/projects/{projectsId}/occurrences | 
[**deleteNote**](GrafeasApi.md#deleteNote) | **DELETE** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
[**deleteOccurrence**](GrafeasApi.md#deleteOccurrence) | **DELETE** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
[**getNote**](GrafeasApi.md#getNote) | **GET** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
[**getOccurrence**](GrafeasApi.md#getOccurrence) | **GET** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
[**getOccurrenceNote**](GrafeasApi.md#getOccurrenceNote) | **GET** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}/notes | 
[**getOperation**](GrafeasApi.md#getOperation) | **GET** /v1alpha1/projects/{projectsId}/operations/{operationsId} | 
[**listNoteOccurrences**](GrafeasApi.md#listNoteOccurrences) | **GET** /v1alpha1/projects/{projectsId}/notes/{notesId}/occurrences | 
[**listNotes**](GrafeasApi.md#listNotes) | **GET** /v1alpha1/projects/{projectsId}/notes | 
[**listOccurrences**](GrafeasApi.md#listOccurrences) | **GET** /v1alpha1/projects/{projectsId}/occurrences | 
[**listOperations**](GrafeasApi.md#listOperations) | **GET** /v1alpha1/projects/{projectsId}/operations | 
[**updateNote**](GrafeasApi.md#updateNote) | **PUT** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
[**updateOccurrence**](GrafeasApi.md#updateOccurrence) | **PUT** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
[**updateOperation**](GrafeasApi.md#updateOperation) | **PUT** /v1alpha1/projects/{projectsId}/operations/{operationsId} | 


<a name="createNote"></a>
# **createNote**
> Note createNote(projectsId, noteId, note)



Creates a new note.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `parent`. This field contains the projectId for example: \"project/{project_id}
String noteId = "noteId_example"; // String | The ID to use for this note.
Note note = new Note(); // Note | The Note to be inserted
try {
    Note result = apiInstance.createNote(projectsId, noteId, note);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#createNote");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;parent&#x60;. This field contains the projectId for example: \&quot;project/{project_id} |
 **noteId** | **String**| The ID to use for this note. | [optional]
 **note** | [**Note**](Note.md)| The Note to be inserted | [optional]

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="createOccurrence"></a>
# **createOccurrence**
> Occurrence createOccurrence(projectsId, occurrence)



Creates a new occurrence.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `parent`. This field contains the projectId for example: \"projects/{project_id}\"
Occurrence occurrence = new Occurrence(); // Occurrence | The occurrence to be inserted
try {
    Occurrence result = apiInstance.createOccurrence(projectsId, occurrence);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#createOccurrence");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;parent&#x60;. This field contains the projectId for example: \&quot;projects/{project_id}\&quot; |
 **occurrence** | [**Occurrence**](Occurrence.md)| The occurrence to be inserted | [optional]

### Return type

[**Occurrence**](Occurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="deleteNote"></a>
# **deleteNote**
> Empty deleteNote(projectsId, notesId)



Deletes the given note from the system.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `name`. The name of the note in the form \"projects/{project_id}/notes/{note_id}\"
String notesId = "notesId_example"; // String | Part of `name`. See documentation of `projectsId`.
try {
    Empty result = apiInstance.deleteNote(projectsId, notesId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#deleteNote");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;name&#x60;. The name of the note in the form \&quot;projects/{project_id}/notes/{note_id}\&quot; |
 **notesId** | **String**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. |

### Return type

[**Empty**](Empty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="deleteOccurrence"></a>
# **deleteOccurrence**
> Empty deleteOccurrence(projectsId, occurrencesId)



Deletes the given occurrence from the system.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `name`. The name of the occurrence in the form \"projects/{project_id}/occurrences/{occurrence_id}\"
String occurrencesId = "occurrencesId_example"; // String | Part of `name`. See documentation of `projectsId`.
try {
    Empty result = apiInstance.deleteOccurrence(projectsId, occurrencesId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#deleteOccurrence");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;name&#x60;. The name of the occurrence in the form \&quot;projects/{project_id}/occurrences/{occurrence_id}\&quot; |
 **occurrencesId** | **String**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. |

### Return type

[**Empty**](Empty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="getNote"></a>
# **getNote**
> Note getNote(projectsId, notesId)



Returns the requested occurrence

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `name`. The name of the note in the form \"projects/{project_id}/notes/{note_id}\"
String notesId = "notesId_example"; // String | Part of `name`. See documentation of `projectsId`.
try {
    Note result = apiInstance.getNote(projectsId, notesId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#getNote");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;name&#x60;. The name of the note in the form \&quot;projects/{project_id}/notes/{note_id}\&quot; |
 **notesId** | **String**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. |

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="getOccurrence"></a>
# **getOccurrence**
> Occurrence getOccurrence(projectsId, occurrencesId)



Returns the requested occurrence

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `name`. The name of the occurrence in the form \"projects/{project_id}/occurrences/{occurrence_id}\"
String occurrencesId = "occurrencesId_example"; // String | Part of `name`. See documentation of `projectsId`.
try {
    Occurrence result = apiInstance.getOccurrence(projectsId, occurrencesId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#getOccurrence");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;name&#x60;. The name of the occurrence in the form \&quot;projects/{project_id}/occurrences/{occurrence_id}\&quot; |
 **occurrencesId** | **String**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. |

### Return type

[**Occurrence**](Occurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="getOccurrenceNote"></a>
# **getOccurrenceNote**
> Note getOccurrenceNote(projectsId, occurrencesId)



Gets the note that this occurrence is attached to.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `name`. The name of the occurrence in the form \"projects/{project_id}/occurrences/{occurrence_id}\"
String occurrencesId = "occurrencesId_example"; // String | Part of `name`. See documentation of `projectsId`.
try {
    Note result = apiInstance.getOccurrenceNote(projectsId, occurrencesId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#getOccurrenceNote");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;name&#x60;. The name of the occurrence in the form \&quot;projects/{project_id}/occurrences/{occurrence_id}\&quot; |
 **occurrencesId** | **String**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. |

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="getOperation"></a>
# **getOperation**
> Operation getOperation(projectsId, operationsId)



Returns the requested occurrence

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `name`. The name of the operation in the form \"projects/{project_id}/operations/{operation_id}\"
String operationsId = "operationsId_example"; // String | Part of `name`. See documentation of `projectsId`.
try {
    Operation result = apiInstance.getOperation(projectsId, operationsId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#getOperation");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;name&#x60;. The name of the operation in the form \&quot;projects/{project_id}/operations/{operation_id}\&quot; |
 **operationsId** | **String**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. |

### Return type

[**Operation**](Operation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="listNoteOccurrences"></a>
# **listNoteOccurrences**
> ListNoteOccurrencesResponse listNoteOccurrences(projectsId, notesId, filter, pageSize, pageToken)



Lists the names of Occurrences linked to a particular Note.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `name`. The name field will contain the note name for example:   \"project/{project_id}/notes/{note_id}\"
String notesId = "notesId_example"; // String | Part of `name`. See documentation of `projectsId`.
String filter = "filter_example"; // String | The filter expression.
Integer pageSize = 56; // Integer | Number of notes to return in the list.
String pageToken = "pageToken_example"; // String | Token to provide to skip to a particular spot in the list.
try {
    ListNoteOccurrencesResponse result = apiInstance.listNoteOccurrences(projectsId, notesId, filter, pageSize, pageToken);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#listNoteOccurrences");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;name&#x60;. The name field will contain the note name for example:   \&quot;project/{project_id}/notes/{note_id}\&quot; |
 **notesId** | **String**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. |
 **filter** | **String**| The filter expression. | [optional]
 **pageSize** | **Integer**| Number of notes to return in the list. | [optional]
 **pageToken** | **String**| Token to provide to skip to a particular spot in the list. | [optional]

### Return type

[**ListNoteOccurrencesResponse**](ListNoteOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="listNotes"></a>
# **listNotes**
> ListNotesResponse listNotes(projectsId, filter, pageSize, pageToken)



Lists all notes for a given project.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `parent`. This field contains the projectId for example: \"project/{project_id}
String filter = "filter_example"; // String | The filter expression.
Integer pageSize = 56; // Integer | Number of notes to return in the list.
String pageToken = "pageToken_example"; // String | Token to provide to skip to a particular spot in the list.
try {
    ListNotesResponse result = apiInstance.listNotes(projectsId, filter, pageSize, pageToken);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#listNotes");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;parent&#x60;. This field contains the projectId for example: \&quot;project/{project_id} |
 **filter** | **String**| The filter expression. | [optional]
 **pageSize** | **Integer**| Number of notes to return in the list. | [optional]
 **pageToken** | **String**| Token to provide to skip to a particular spot in the list. | [optional]

### Return type

[**ListNotesResponse**](ListNotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="listOccurrences"></a>
# **listOccurrences**
> ListOccurrencesResponse listOccurrences(projectsId, filter, pageSize, pageToken)



Lists active occurrences for a given project/Digest.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `parent`. This contains the projectId for example: projects/{project_id}
String filter = "filter_example"; // String | The filter expression.
Integer pageSize = 56; // Integer | Number of occurrences to return in the list.
String pageToken = "pageToken_example"; // String | Token to provide to skip to a particular spot in the list.
try {
    ListOccurrencesResponse result = apiInstance.listOccurrences(projectsId, filter, pageSize, pageToken);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#listOccurrences");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;parent&#x60;. This contains the projectId for example: projects/{project_id} |
 **filter** | **String**| The filter expression. | [optional]
 **pageSize** | **Integer**| Number of occurrences to return in the list. | [optional]
 **pageToken** | **String**| Token to provide to skip to a particular spot in the list. | [optional]

### Return type

[**ListOccurrencesResponse**](ListOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="listOperations"></a>
# **listOperations**
> ListOperationsResponse listOperations(projectsId, filter, pageSize, pageToken)



Lists all operations for a given project.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `parent`. This field contains the projectId for example: \"project/{project_id}
String filter = "filter_example"; // String | The filter expression.
Integer pageSize = 56; // Integer | Number of operations to return in the list.
String pageToken = "pageToken_example"; // String | Token to provide to skip to a particular spot in the list.
try {
    ListOperationsResponse result = apiInstance.listOperations(projectsId, filter, pageSize, pageToken);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#listOperations");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;parent&#x60;. This field contains the projectId for example: \&quot;project/{project_id} |
 **filter** | **String**| The filter expression. | [optional]
 **pageSize** | **Integer**| Number of operations to return in the list. | [optional]
 **pageToken** | **String**| Token to provide to skip to a particular spot in the list. | [optional]

### Return type

[**ListOperationsResponse**](ListOperationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="updateNote"></a>
# **updateNote**
> Note updateNote(projectsId, notesId, note)



Updates an existing note.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `name`. The name of the note. Should be of the form \"projects/{project_id}/notes/{note_id}\".
String notesId = "notesId_example"; // String | Part of `name`. See documentation of `projectsId`.
Note note = new Note(); // Note | The updated note.
try {
    Note result = apiInstance.updateNote(projectsId, notesId, note);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#updateNote");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;name&#x60;. The name of the note. Should be of the form \&quot;projects/{project_id}/notes/{note_id}\&quot;. |
 **notesId** | **String**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. |
 **note** | [**Note**](Note.md)| The updated note. | [optional]

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="updateOccurrence"></a>
# **updateOccurrence**
> Occurrence updateOccurrence(projectsId, occurrencesId, occurrence)



Updates an existing occurrence.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `name`. The name of the occurrence. Should be of the form \"projects/{project_id}/occurrences/{occurrence_id}\".
String occurrencesId = "occurrencesId_example"; // String | Part of `name`. See documentation of `projectsId`.
Occurrence occurrence = new Occurrence(); // Occurrence | The updated occurrence.
try {
    Occurrence result = apiInstance.updateOccurrence(projectsId, occurrencesId, occurrence);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#updateOccurrence");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;name&#x60;. The name of the occurrence. Should be of the form \&quot;projects/{project_id}/occurrences/{occurrence_id}\&quot;. |
 **occurrencesId** | **String**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. |
 **occurrence** | [**Occurrence**](Occurrence.md)| The updated occurrence. | [optional]

### Return type

[**Occurrence**](Occurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="updateOperation"></a>
# **updateOperation**
> Operation updateOperation(projectsId, operationsId, body)



Updates an existing operation returns an error if operation  does not exist. The only valid operations are to update mark the done bit change the result.

### Example
```java
// Import classes:
//import java.io.grafeas.ApiException;
//import io.swagger.client.api.GrafeasApi;


GrafeasApi apiInstance = new GrafeasApi();
String projectsId = "projectsId_example"; // String | Part of `name`. The name of the Operation. Should be of the form \"projects/{project_id}/operations/{operation_id}\".
String operationsId = "operationsId_example"; // String | Part of `name`. See documentation of `projectsId`.
UpdateOperationRequest body = new UpdateOperationRequest(); // UpdateOperationRequest | The request body.
try {
    Operation result = apiInstance.updateOperation(projectsId, operationsId, body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#updateOperation");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsId** | **String**| Part of &#x60;name&#x60;. The name of the Operation. Should be of the form \&quot;projects/{project_id}/operations/{operation_id}\&quot;. |
 **operationsId** | **String**| Part of &#x60;name&#x60;. See documentation of &#x60;projectsId&#x60;. |
 **body** | [**UpdateOperationRequest**](UpdateOperationRequest.md)| The request body. | [optional]

### Return type

[**Operation**](Operation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

