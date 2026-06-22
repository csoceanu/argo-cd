# PROPOSALS/MANIFEST-HYDRATOR/COMMIT-SERVER Documentation Index

## Overview
This documentation describes the Argo CD Commit Server, a component of the manifest hydration proposal designed to provide push access to Git repositories. Its primary purpose is to receive hydrated Kubernetes manifests via a gRPC interface and commit them to specific target branches, bridging the gap between manifest generation and Git storage.

## Files Summary
*   **README.md**: Defines the core gRPC service interface, message structures, and the functional responsibility of the Commit Server in the hydration workflow.

## Code Changes That Would Require Documentation Updates
*   **gRPC API Modifications**: Any changes to the `.proto` definitions, including adding, renaming, or removing fields in the `CommitManifests` or `CommitPathDetails` messages.
*   **Protocol Transitions**: Changing the communication method from gRPC to another protocol (e.g., REST/OpenAPI).
*   **Git Authentication/Transport Changes**: Changes to how the server handles repository URLs (HTTPS/SSH) or updates to supported Git authentication methods.
*   **Manifest Packaging Logic**: Changes to how manifests are bundled or structured within the Git repository (e.g., changing the default filename from `manifest.yaml` to something else).
*   **Metadata Requirements**: Updates to the required metadata for commits, such as adding new fields for traceability (e.g., environment tags, signature requirements, or build IDs).
*   **README Generation Logic**: Alterations to how the `README.md` content is processed or where it is placed relative to the manifests.

## Key Technical Concepts
*   **Hydrated Manifests**: The final, rendered Kubernetes YAML/JSON documents produced after a hydration process.
*   **Dry Commit/SHA**: The reference to the source commit (unhydrated) that triggered the hydration process, used for traceability.
*   **gRPC Service**: The interface through which callers (like a Hydrator) request manifest commits.
*   **CommitManifests Message**: The primary data structure containing repo URLs, target branches, and commit metadata.
*   **CommitPathDetails**: The structure defining the specific filesystem path, manifest content, and documentation for a hydration output.
*   **Target Branch**: The specific Git branch designated to hold the hydrated manifests.

## Related Components
*   **Argo CD Manifest Hydrator**: The broader system that generates the manifests passed to this server.
*   **Git Repositories**: The destination targets for the hydrated manifests.
*   **Argo CD Repo Server**: The component likely responsible for interacting with the Commit Server during the hydration lifecycle.