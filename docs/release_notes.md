# Release Notes

## Grafeas v1 beta launch

The v1beta1 release of the Grafeas API includes the following improvements:
*   Names of the different metadata types are more consistent.
*   Structure of the metadata types are more consistent.
*   Deprecated methods and fields removed.
*   Better conformance to the [Google API style guide](https://cloud.google.com/apis/design/).

### Method changes

*   [Long running operations](https://github.com/googleapis/googleapis/blob/master/google/longrunning/operations.proto)
    removed.
*   The `name` field is no longer part of any create or list methods. You can
    only specify the `parent` to create or list resources in.

v1alpha1                    | v1beta1
--------------------------- | ------------------------------------
GetVulnzOccurrencesSummary  | GetVulnerabilityOccurrencesSummary
CreateOperation             | _[Removed]_
BatchCreateOperations       | _[Removed]_
UpdateOperation             | _[Removed]_


### Metadata schema changes

Proto changes, as described below.

#### General changes

*   Each metadata type now has its own file and package.
*   A common file now houses all common messages.
*   `CustomType` metadata type was removed.

#### Note changes

*   `RelatedUrl` now lives in the common file.
*   Redundant suffixes were removed from some `oneof` field names.

v1alpha1 Note                  | v1beta1 Note
------------------------------ | ------------------------------------
note\_type                     | type
_[N/A]_                        | related\_note\_names
note\_type.vulnerability\_type | type.vulnerability
note\_type.build\_type         | type.build

#### Occurrence changes

*   `resource_url` was replaced with a resource object.
*   Redundant suffixes were removed from some `oneof` field names.
*   Every details `oneof` type has either been renamed "Details" or wrapped in a
    `Details` message for consistency.

v1alpha1 Occurrence            | v1beta1 Occurrence
------------------------------ | ------------------------------------
resource\_url                  | resource.uri
details.custom\_details        | _[Removed]_
details.vulnerability\_details | details.vulnerability
details.build\_details         | details.build

#### Note kind changes

*   `NoteKind` enum values were simplified.
*   The `NoteKind` enum now lives in the common file.

v1alpha1 NoteKind            | v1beta1 NoteKind
---------------------------- | ------------------------
KIND\_UNSPECIFIED            | NOTE\_KIND\_UNSPECIFIED
CUSTOM                       | _[Removed]_
PACKAGE\_VULNERABILITY       | VULNERABILITY
BUILD\_DETAILS               | BUILD
IMAGE\_BASIS                 | IMAGE
PACKAGE\_MANAGER             | PACKAGE
DEPLOYABLE                   | DEPLOYMENT
ATTESTATION\_AUTHORITY       | ATTESTATION

#### Attestation changes

*   `AttestationAuthority` renamed to `Authority`.
*   `Attestation` wrapped in a `Details` message and pulled out to top-level.
*   `AttestationAuthorityHint` renamed to `Hint`.

#### Build changes

*   `BuildType` renamed to `Build`.
*   `BuildDetails` renamed to `Details`.
*   `BuildSignature.signature` has changed from a string to a bytes type.

#### Deployment changes

*   `Deployment` wrapped in a `Details` message and pulled out to top-level.

#### Discovery changes

*   `Discovered` wrapped in a `Details` message and pulled out to top-level.

v1alpha1 Discovered         | v1beta1 Discovered
--------------------------- | ---------------------
operation                   | _[Removed]_

#### Image changes

*   `DockerImage` wrapper message removed and all messages pulled out to
    top-level.
*   `Derived` wrapped in a Details message.
*   `Derived.distance` has changed from a `uint32` to an `int32` type.

#### Package changes

*   `PackageManager` wrapper message removed and all messages pulled out to
    top-level.
*   `Version` has been moved into this proto.
*   `Installation` wrapped in a `Details` message.

#### Provenance changes

*   `StorageSource` has been replaced with a URI to make it more generic.

v1alpha1 Source              | v1beta1 Source
---------------------------- | ------------------------------------
source.storage\_source       | _[Removed]_
source.repo\_source          | _[Removed]_
artifact\_storage\_source    | artifact\_storage\_source\_uri

<br>

v1alpha1 Artifact    | v1beta1 Artifact
-------------------- | ---------------------
name                 | _[Removed]_

### Vulnerability changes

*   `VulnerabilityType` renamed to `Vulnerability`.
*   `Version` was moved into the package proto.
*   `VulnerabilityDetails` renamed to Details and pulled out to top-level.
*   `PackageIssue` pulled out to top-level.
*   `VulnerabilityLocation` pulled out to top-level.

## Grafeas alpha launch

The alpha launch of Grafeas included:

*   Grafeas API spec
*   Client libraries for Java, Python, Go
*   Reference implementation
*   Kritis overview
