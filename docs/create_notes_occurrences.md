# Grafeas How To: Create Notes and Occurrences

Notes and occurrences are the basic building blocks of storing metadata in
Grafeas. A _[note](grafeas_concepts.md#notes)_ represents
high-level information about a piece of metadata, whereas an
_[occurrence](grafeas_concepts.md#occurrences)_ describes the
the occurrence of the note in the project.

This guide walks you through creating notes and occurrences on a running Grafeas
server.

## Before you begin
 * Set up a Grafeas server running on port 8080 using the instructions in
  [Running Grafeas](running_grafeas.md).

## Creating a project

All notes and occurrences are stored in a project. Create separate projects for
notes and occurrences so that setting up fine-grained access control is easier.

To create a project named `provider_example` to store the notes:

```bash
curl -X POST http://localhost:8080/v1beta1/projects \
     -H "Content-Type: application/json" \
     --data '{"name":"projects/provider_example"}'
```

To create a project named `occurrence_example` to store the occurrences:

```bash
curl -X POST http://localhost:8080/v1beta1/projects \
     -H "Content-Type: application/json" \
     --data '{"name":"projects/occurrence_example"}'
```

## Creating a note

### Using curl

First, create a `note.json` file with the information you want your note to
contain:

```json
{
    "name" : "projects/provider_example/notes/test",
    "shortDescription": "A brief description of the note",
    "longDescription": "A longer description of the note",
    "kind": "VULNERABILITY",
    "vulnerability": {
        "details": [
        {
            "package": "libexempi3",
            "cpeUri": "cpe:/o:debian:debian_linux:7",
            "minAffectedVersion": {
              "name": "2.5.7",
              "revision": "1",
              "kind": "NORMAL"
            }
        }]
    }
}
```

Now create the note with the following command:

```bash
curl http://localhost:8080/v1beta1/projects/provider_example/notes?note_id=testNote \
    -X POST -H "Content-Type: application/json" -d @note.json
```

This creates a note in the `provider_example` project. If you want to verify
that your note was created, you can list all notes in a given project with the
command:

```bash
curl http://localhost:8080/v1beta1/projects/provider_example/notes
```
#### Further examples
Examples of deployment and build notes are contained within the [json_examples](json_examples/) folder

### Using Postman
If you're familiar with [Postman](https://www.postman.com), then you can use the script contained in the [postman](postman/) folder to create projects, create and get example notes create and read example occurrences for deployment, vulnerability and build types.

### Using Java

The code example below uses Java, but the Go, Python, and Ruby client libraries
work in a similar way.

Add the following methods to the class you're using to talk to the client
library. If you followed the
[getting started guide](https://github.com/grafeas/client-java/#getting-started),
this is the `GrafeasApiExample` class.

When calling the below method, you'll need to specify the `provider_example`
project with its full name: `projects/provider_example`.

```java
/**
 * Creates and returns a new note
 *
 * @param parent The project containing the note, of the format "projects/<project_id>".
 * @param name A user-specified identifier for the note
 * @return a Note object representing the new note
 * @throws ApiException if there were errors with the API request
 */
public static ApiNote createNote(String parent, String name) throws ApiException {
  ApiNote body = new ApiNote();
  body.setName(String.format("%s/notes/%s", parent, name));
  body.setKind(ApiNoteKind.PACKAGE_VULNERABILITY);

  ApiVulnerabilityType vulnerabilityType = new ApiVulnerabilityType();
  vulnerabilityType.setCvssScore(0.3f);
  body.setVulnerabilityType(vulnerabilityType);

  return apiInstance.createNote(parent, body);
}
```

To verify that you successfully created your note, you can list all of the notes
in a given project:

```java
/**
 * Prints the notes in the specified project.
 *
 * @param parent the project to fetch notes from, of the form "projects/<project_id>".
 * @throws ApiException if there were errors with the API request
 */
public static void listNotes(String parent) throws ApiException {
  String filter = ""; // Optional; filter expression.
  int pageSize = 10; // Number of notes to return in the list.
  String pageToken = ""; // Optional; skip to a particular spot in the list.
  try {
    ApiListNotesResponse result = apiInstance.listNotes(parent, filter, pageSize, pageToken);
    System.out.println(result.getNotes());
  } catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#listNotes");
    e.printStackTrace();
  }
}
```

## Creating an occurrence

### Using curl

First, create an `occurrence.json` file describing the occurrence you want to
create. The following example creates an occurrence of the note you created
above, with some sample data about where the vulnerability occurred.

```json
{
  "name": "projects/occurrence_example/occurrences/test",
  "resource": {
    "uri": "https://gcr.io/project/image@sha256:foo"
  },
  "noteName": "projects/provider_example/notes/test",
  "kind": "VULNERABILITY",
  "vulnerability": {
    "packageIssue": [
      {
        "affectedLocation": {
          "cpeUri": "7",
          "package": "a",
          "version": {
            "name": "v1.1.1",
            "kind": "NORMAL",
            "revision": "r"
          }
        },
        "fixedLocation": {
          "cpeUri": "cpe:/o:debian:debian_linux:7",
          "package": "a",
          "version": {
            "name": "namestring",
            "kind": "NORMAL",
            "revision": "1"
          }
        }
      }
    ]
  }
}
```

Then create an occurrence with the following command:

```bash
curl http://localhost:8080/v1beta1/projects/occurrence_example/occurrences \
     -X POST -H "Content-Type: application/json" -d @occurrence.json
```

To verify that you successfully created the occurrence, you can list all
occurrences on your project with the following command:

```bash
curl http://localhost:8080/v1beta1/projects/occurrence_example/occurrences
```

You can find information about other available actions in the
[Grafeas API documentation](https://github.com/grafeas/grafeas/blob/669d9cdc0ca804bf7d29dcf6d66bb9d8e94b08b6/v1alpha1/docs/GrafeasApi.md).

#### Further examples
Examples of deployment and build occurrences are contained within the [json_examples](json_examples/) folder

### Using Java

The code sample below uses Java, but the Go, Python, and Ruby client libraries
work in a similar way. You can add the following methods to your
[GrafeasApiExample](https://github.com/nhayes/client-java/#getting-started)
class, or the equivalent class that you're using to talk to the client library.

You can create an occurrence with the following code. Note that you need to pass
in a valid note name for an existing note.

```java
/**
 * Creates and returns a new occurrence.
 *
 * @param resourceUrl the URL associated with the image. For example,
 *     "https://gcr.io/project/image@sha256:foo"
 * @param parent the project to store this occurrence in, of the form "projects/<project_id>".
 * @param noteName the full name of the note this occurrence is attached to, of the form
 *     "projects/<project_id>/notes/<note_id>"
 * @return an ApiOccurrence object representing the new occurrence
 * @throws ApiException if there were errors with the API request
 */
public static ApiOccurrence createOccurrence(String resourceUrl, String parent, String noteName)
    throws ApiException {
  ApiOccurrence body = new ApiOccurrence();
  body.setResourceUrl(resourceUrl);
  body.setNoteName(noteName);

  return apiInstance.createOccurrence(parent, body);
}
```

You can list occurrences in a similar manner:

```java
/**
 * Prints the notes in the specified project.
 *
 * @param parent the project to fetch notes from, of the form "projects/<project_id>".
 * @throws ApiException if there were errors with the API request
 */
public static void listOccurrences(String parent) throws ApiException {
  String filter = ""; // The filter expression.
  int pageSize = 56; // Number of occurrences to return in the list.
  String pageToken = "";
  try {
    ApiListOccurrencesResponse result =
        apiInstance.listOccurrences(parent, filter, pageSize, pageToken);
    System.out.println(result);
  } catch (ApiException e) {
    System.err.println("Exception when calling GrafeasApi#listOccurrences");
    e.printStackTrace();
  }
}
```

You can find information about other available actions in the
[Grafeas API documentation](https://github.com/grafeas/grafeas/blob/669d9cdc0ca804bf7d29dcf6d66bb9d8e94b08b6/v1alpha1/docs/GrafeasApi.md).
