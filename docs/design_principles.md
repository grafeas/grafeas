# Grafeas Core Design Principles

This page outlines the core principles that guide Grafeas's design:

- **Universal coverage**: Grafeas stores structured metadata against the software component’s unique identifier (e.g., container image digest), so you don’t have to co-locate it with the component’s registry, and so it can store metadata about components from many different repositories.
- **Hybrid cloud-friendly**: Just as you can use JFrog Artifactory as the central, universal component repository across hybrid cloud deployments, you can use the Grafeas API as a central, universal metadata store.
- **Pluggable**: Grafeas makes it easy to add new metadata producers and consumers (for example, if you decide to add or change security scanners, add new build systems, etc.)
- **Structured**: Structured metadata schemas for common metadata types (e.g., vulnerability, build, attestation, and package index metadata) let you add new metadata types and providers, and the tools that depend on Grafeas can immediately understand those new sources.
- **Strong access controls**: Grafeas allows you to carefully control access for multiple metadata producers and consumers.
- **Rich query-ability**: With Grafeas, you can easily query all metadata across all of your components so you don’t have to parse monolithic reports on each component.
