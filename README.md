# Grafeas: A Component Metadata API

![Grafeas logo](logo/grafeas-logo-128.png)

Grafeas ("scribe" in Greek) is an open-source artifact metadata API that provides a uniform way to audit and govern your software supply chain. Grafeas defines an API spec for managing metadata about software resources, such
as container images, Virtual Machine (VM) images, JAR files, and scripts. You can use Grafeas to define and aggregate information about your project's components. Grafeas provides organizations with a central source of truth for tracking and enforcing policies across an ever growing set of software development teams and pipelines. Build, auditing, and compliance tools can use the Grafeas API to store, query, and retrieve comprehensive metadata on software components of all kinds.

Grafeas's benefits include:

- **Universal coverage**: Grafeas stores structured metadata against the software component’s unique identifier (e.g., container image digest), so you don’t have to co-locate it with the component’s registry, and so it can store metadata about components from many different repositories.
- **Hybrid cloud-friendly**: Just as you can use JFrog Artifactory as the central, universal component repository across hybrid cloud deployments, you can use the Grafeas API as a central, universal metadata store.
- **Pluggable**: Grafeas makes it easy to add new metadata producers and consumers (for example, if you decide to add or change security scanners, add new build systems, etc.)
- **Structured**: Structured metadata schemas for common metadata types (e.g., vulnerability, build, attestation, and package index metadata) let you add new metadata types and providers, and the tools that depend on Grafeas can immediately understand those new sources.
- **Strong access controls**: Grafeas allows you to carefully control access for multiple metadata producers and consumers.
- **Rich query-ability**: With Grafeas, you can easily query all metadata across all of your components so you don’t have to parse monolithic reports on each component.

Grafeas divides the metadata information into [_notes_](docs/grafeas_concepts.md#notes) and
[_occurrences_](docs/grafeas_concepts.md#occurrences). Notes are high-level descriptions of particular
types of metadata. Occurrences are instantiations of notes, which describe how
and when a given note occurs on the resource associated with the occurrence.
This division allows third-party metadata providers to create and manage
metadata on behalf of many customers. It also allows for fine-grained access
control of different types of metadata.

## Getting Started

* Learn the [Grafeas concepts](docs/grafeas_concepts.md)
* Run Grafeas locally following [these
instructions](docs/running_grafeas.md)
* Once you have a running server, you can
use the [client libraries](https://github.com/grafeas) to experiment with
creating notes and occurrences in Grafeas. There are client libraries available in Java, Go, Ruby, and Python.
* The authoritative API for grafeas is the [protobuf
files](https://github.com/Grafeas/Grafeas/tree/master/proto/v1beta1).

## Support

If you have questions, reach out to us on
[grafeas-users](https://groups.google.com/forum/#!forum/grafeas-users). For
questions about contributing, please see the [section](#contributing) below or
use [grafeas-dev](https://groups.google.com/forum/#!forum/grafeas-dev).

Grafeas announcements will be posted to its
[@grafeasio](https://twitter.com/Grafeasio) Twitter account and to
[grafeas-users](https://groups.google.com/forum/#!forum/grafeas-users).

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details on how you can contribute.

See [DEVELOPMENT](DEVELOPMENT.md) for details on the  development and testing workflow.

## License

Grafeas is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
