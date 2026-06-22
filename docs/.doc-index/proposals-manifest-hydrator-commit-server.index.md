# PROPOSALS/MANIFEST-HYDRATOR/COMMIT-SERVER Documentation Index

## Overview
This documentation describes the Argo CD Commit Server, a specialized gRPC service designed to handle push access to git repositories for hydrated manifests. It defines the communication protocol and data structures required to automate the commitment of processed Kubernetes manifests back into version control.

## Files Summary
*   **proposals/manifest-hydrator/commit-server/README.md**: Defines the purpose of the Commit Server and provides the full Protobuf interface specification for the gRPC service, including request and response message structures.

## Code Changes That Would Require Documentation Updates
*   **gRPC Service Definition**: Any modifications to the RPC method names or the addition of new service methods in the Commit Server.
*   **Protobuf Message Schema**: Changes to the fields within `CommitManifests` or `CommitPathDetails`, such as adding support for new metadata (e.g., tags, labels, or commit signatures).
*   **Data Format Requirements**: Changes to how Kubernetes manifests are represented (e.g., moving from a list of JSON strings to a different format like YAML or binary blobs).
*   **Authentication/Authorization**: Updates to how the server handles credentials for the `repoURL`, which might require new fields in the message specification.
*   **Git Integration Logic**: Changes to how the server interprets `targetBranch` or handles branch creation/updates.
*   **Path Handling**: Modifications to how directories are structured or how the `README.md` file is generated/written within the repository.

## Key Technical Concepts
*   **Commit Server**: The gRPC-based component responsible for git write operations.
*   **Hydrated Manifests**: Fully processed/rendered Kubernetes manifests ready for deployment.
*   **Dry SHA**: The SHA256 hash of the original "dry" source commit used to track the origin of the hydration.
*   **gRPC Interface**: The specific Protobuf service definition used for communication.
*   **CommitManifests**: The primary message structure containing repository URL, branch, and metadata.
*   **CommitPathDetails**: A sub-structure defining specific file paths, manifest content, and documentation for the target repository.
*   **JSON Manifests**: The current expected format for Kubernetes resources in the gRPC request.

## Related Components
*   **Argo CD**: The parent project and orchestration framework.
*   **Manifest Hydrator**: The broader system responsible for transforming "dry" manifests into "hydrated" manifests.
*   **Git Repository**: The destination storage for the hydrated manifests.
*   **Protocol Buffers (Protobuf)**: The underlying technology for service definition and serialization.