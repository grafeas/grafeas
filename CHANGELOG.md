This is the changelog of Grafeas server releases. For more information on
versioning, see [versioning](docs/versioning.md) document.

v0.2.2
  * Enhanced support for Vulnerability Notes and Occurrences, with the following changes:
    * Indicate the location at which an affected package was found in the container image.
    * CVSS proto compatible with v2 and v3.
    * More detailed status of language package scans in the Discovery occurrence.
    * Added cvss_version field to indicate which version was used to populate fields: cvss_score and severity.
    * Added support for Vex Assessments.
  * Added support for language packages in Package Notes and Occurrences.
  * Enhanced SLSA support.
    * Added support for SLSA v0.2 to the intoto statement.
    * Added SLSA v0.2 converter.
  * Added SBOM support.
  * Updated versions of frameworks and libraries
    * Use ANTLR v4.
    * Golang 1.20.

v0.2.1:
  * Update grpc-gateway v2.7.3 to generate compatible openapi spec.

v0.2.0:
  * Support for SBOM, using SPDX spec.
  * Enhanced support for Occurrences, with the following additions (via v1beta1 API):
    * [DSSE Envelope](https://github.com/secure-systems-lab/dsse),
    * archival timestamp.
  * Enhanced support for Vulnerability Notes and Occurrences, with the following additions (via v1beta1 API):
    * per-source CVE reporting,
    * CWEs tracking,
    * CVSS v2 and v3 scores,
    * improved package types support,
    * vendor information,
    * inclusive version ranges.
  * Added support for ARM builds.
  * Improved validation and permissions checks.
  * Improved development support on Windows.
  * Documentation and code cleanups and minor fixes.

v0.1.6:
  * Added support for in-toto.
  * Added JWT support to attestation notes.

v0.1.5:
  * Upgraded to golang 1.14.0
  * Added last_scan_time to discovery occurrences
  * Added support for Windows updates

v0.1.4:
  * Support for use of existing secret and certs in Helm chart, in addition to generating them.
  * Fix for handling http requests.
  * Support for multi-platform protobuf compiler download.
  * Checked in `v1beta1` go generated protos, to simplify integration downstream.

v0.1.3:
  * same as `v0.1.2`, but with generated protos uploaded with the release.

v0.1.2:
  * Support for multiple storage implementations.

v0.1.1:
  * Grafeas helm chart is compliant with stable requirements.
  * Code cleanup.
  * Added back `max_affected_version` to Vulnerability.

v0.1.0:
  * Grafeas server implements v1beta1 Grafeas API.
  * Grafeas server can run:
    * as standalone server,
    * as k8s pod.
