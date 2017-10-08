# swagger-java-client

## Requirements

Building the API client library requires [Maven](https://maven.apache.org/) to be installed.

## Installation

To install the API client library to your local Maven repository, simply execute:

```shell
mvn install
```

To deploy it to a remote Maven repository instead, configure the settings of the repository and execute:

```shell
mvn deploy
```

Refer to the [official documentation](https://maven.apache.org/plugins/maven-deploy-plugin/usage.html) for more information.

### Maven users

Add this dependency to your project's POM:

```xml
<dependency>
    <groupId>io.swagger</groupId>
    <artifactId>swagger-java-client</artifactId>
    <version>1.0.0</version>
    <scope>compile</scope>
</dependency>
```

### Gradle users

Add this dependency to your project's build file:

```groovy
compile "io.swagger:swagger-java-client:1.0.0"
```

### Others

At first generate the JAR by executing:

    mvn package

Then manually install the following JARs:

* target/swagger-java-client-1.0.0.jar
* target/lib/*.jar

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Java code:

```java

import java.io.grafeas.*;
import java.io.grafeas.auth.*;
import java.io.grafeas.model.*;
import io.swagger.client.api.GrafeasApi;

import java.io.File;
import java.util.*;

public class GrafeasApiExample {

    public static void main(String[] args) {
        
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
    }
}

```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*GrafeasApi* | [**createNote**](docs/GrafeasApi.md#createNote) | **POST** /v1alpha1/projects/{projectsId}/notes | 
*GrafeasApi* | [**createOccurrence**](docs/GrafeasApi.md#createOccurrence) | **POST** /v1alpha1/projects/{projectsId}/occurrences | 
*GrafeasApi* | [**deleteNote**](docs/GrafeasApi.md#deleteNote) | **DELETE** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
*GrafeasApi* | [**deleteOccurrence**](docs/GrafeasApi.md#deleteOccurrence) | **DELETE** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
*GrafeasApi* | [**getNote**](docs/GrafeasApi.md#getNote) | **GET** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
*GrafeasApi* | [**getOccurrence**](docs/GrafeasApi.md#getOccurrence) | **GET** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
*GrafeasApi* | [**getOccurrenceNote**](docs/GrafeasApi.md#getOccurrenceNote) | **GET** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}/notes | 
*GrafeasApi* | [**getOperation**](docs/GrafeasApi.md#getOperation) | **GET** /v1alpha1/projects/{projectsId}/operations/{operationsId} | 
*GrafeasApi* | [**listNoteOccurrences**](docs/GrafeasApi.md#listNoteOccurrences) | **GET** /v1alpha1/projects/{projectsId}/notes/{notesId}/occurrences | 
*GrafeasApi* | [**listNotes**](docs/GrafeasApi.md#listNotes) | **GET** /v1alpha1/projects/{projectsId}/notes | 
*GrafeasApi* | [**listOccurrences**](docs/GrafeasApi.md#listOccurrences) | **GET** /v1alpha1/projects/{projectsId}/occurrences | 
*GrafeasApi* | [**listOperations**](docs/GrafeasApi.md#listOperations) | **GET** /v1alpha1/projects/{projectsId}/operations | 
*GrafeasApi* | [**updateNote**](docs/GrafeasApi.md#updateNote) | **PUT** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
*GrafeasApi* | [**updateOccurrence**](docs/GrafeasApi.md#updateOccurrence) | **PUT** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
*GrafeasApi* | [**updateOperation**](docs/GrafeasApi.md#updateOperation) | **PUT** /v1alpha1/projects/{projectsId}/operations/{operationsId} | 


## Documentation for Models

 - [AliasContext](docs/AliasContext.md)
 - [Artifact](docs/Artifact.md)
 - [Basis](docs/Basis.md)
 - [BuildDetails](docs/BuildDetails.md)
 - [BuildProvenance](docs/BuildProvenance.md)
 - [BuildSignature](docs/BuildSignature.md)
 - [BuildType](docs/BuildType.md)
 - [CloudRepoSourceContext](docs/CloudRepoSourceContext.md)
 - [CloudWorkspaceId](docs/CloudWorkspaceId.md)
 - [CloudWorkspaceSourceContext](docs/CloudWorkspaceSourceContext.md)
 - [Command](docs/Command.md)
 - [CreateOperationRequest](docs/CreateOperationRequest.md)
 - [CustomDetails](docs/CustomDetails.md)
 - [Deployable](docs/Deployable.md)
 - [Deployment](docs/Deployment.md)
 - [Derived](docs/Derived.md)
 - [Detail](docs/Detail.md)
 - [Discovered](docs/Discovered.md)
 - [Discovery](docs/Discovery.md)
 - [Distribution](docs/Distribution.md)
 - [Empty](docs/Empty.md)
 - [ExtendedSourceContext](docs/ExtendedSourceContext.md)
 - [FileHashes](docs/FileHashes.md)
 - [Fingerprint](docs/Fingerprint.md)
 - [GerritSourceContext](docs/GerritSourceContext.md)
 - [GitSourceContext](docs/GitSourceContext.md)
 - [Hash](docs/Hash.md)
 - [Installation](docs/Installation.md)
 - [Layer](docs/Layer.md)
 - [ListNoteOccurrencesResponse](docs/ListNoteOccurrencesResponse.md)
 - [ListNotesResponse](docs/ListNotesResponse.md)
 - [ListOccurrencesResponse](docs/ListOccurrencesResponse.md)
 - [ListOperationsResponse](docs/ListOperationsResponse.md)
 - [Location](docs/Location.md)
 - [ModelPackage](docs/ModelPackage.md)
 - [Note](docs/Note.md)
 - [Occurrence](docs/Occurrence.md)
 - [Operation](docs/Operation.md)
 - [PackageIssue](docs/PackageIssue.md)
 - [ProjectRepoId](docs/ProjectRepoId.md)
 - [RelatedUrl](docs/RelatedUrl.md)
 - [RepoId](docs/RepoId.md)
 - [RepoSource](docs/RepoSource.md)
 - [Source](docs/Source.md)
 - [SourceContext](docs/SourceContext.md)
 - [Status](docs/Status.md)
 - [StorageSource](docs/StorageSource.md)
 - [UpdateOperationRequest](docs/UpdateOperationRequest.md)
 - [Version](docs/Version.md)
 - [VulnerabilityDetails](docs/VulnerabilityDetails.md)
 - [VulnerabilityLocation](docs/VulnerabilityLocation.md)
 - [VulnerabilityType](docs/VulnerabilityType.md)


## Documentation for Authorization

All endpoints do not require authorization.
Authentication schemes defined for the API:

## Recommendation

It's recommended to create an instance of `ApiClient` per thread in a multithreaded environment to avoid any potential issue.

## Author



