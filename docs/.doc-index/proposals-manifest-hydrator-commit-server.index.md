# PROPOSALS/MANIFEST-HYDRATOR/COMMIT-SERVER Documentation Index

## Overview
This documentation describes the Argo CD Commit Server, a component of the manifest hydrator proposal designed to provide push access to Git repositories. It specifically defines the gRPC interface used to transmit hydrated Kubernetes manifests and their associated metadata from the hydration process to a target Git branch.

## Files Summary
*   **README.md**: Provides a high-level description of the Commit Server's purpose and defines the gRPC message structures (`CommitManifests` and `CommitPathDetails`) required for its operation.

## Code Changes That Would Require Documentation Updates
*   **Protobuf Interface Modifications**: Any changes to the `CommitManifests` or `CommitPathDetails` message definitions (e.g., adding fields for GPG signatures, changing field types, or renaming existing fields).
*   **Authentication Logic**: Changes to how the server handles Git credentials or the types of URLs supported (currently HTTPS and SSH).
*   **Manifest Serialization**: If the server shifts from accepting JSON strings to other formats (e.g., YAML, binary blobs, or specific Kubernetes object types).
*   **Repository Structure Logic**: Changes to how the `path` or `readme` fields are interpreted or where they are placed within the target repository.
*   **Traceability Requirements**: Altering the metadata required for provenance, such as the `drySHA`, `commitAuthor`, or `commitTime`.
*   **API Protocol Changes**: Moving away from gRPC or changing the service definition (e.g., implementing a streaming interface for large manifest sets).

## Key Technical Concepts
*   **Commit Server**: The service responsible for the "write" phase of the manifest hydration lifecycle.
*   **Hydrated Manifests**: The final, rendered Kubernetes YAML/JSON output that is ready for cluster application.
*   **Dry SHA**: The unique identifier (SHA256) of the source/input commit used to generate the hydrated manifests, ensuring traceability.
*   **gRPC Service**: The communication protocol used by callers to interact with the Commit Server.
*   **CommitPathDetails**: The data structure mapping specific directory paths to their corresponding manifests and documentation.
*   **Target Branch**: The destination branch in Git where the hydrated manifests will be committed.
*   **JSON Manifests**: The specific format currently expected for the list of Kubernetes documents.

## Related Components
*   **Argo CD**: The primary system this server integrates with.
*   **Manifest Hydrator**: The broader architectural proposal that uses this server to store rendered outputs.
*   **Git Repository**: The storage backend (HTTPS/SSH) where manifests are persisted.
*   **Hydration Controller/Engine**: The upstream component that calls this server after processing source manifests.