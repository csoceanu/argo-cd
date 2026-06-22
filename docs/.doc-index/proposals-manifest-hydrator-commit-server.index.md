# PROPOSALS/MANIFEST-HYDRATOR/COMMIT-SERVER Documentation Index

## Overview
This documentation describes the Argo CD Commit Server, a component designed to provide push access to git repositories for hydrated manifests. It specifically details the gRPC service interface used to accept and commit Kubernetes manifests to a target repository and branch.

## Files Summary
*   **proposals/manifest-hydrator/commit-server/README.md**: Provides a high-level description of the Commit Server's purpose and defines the protobuf message structures for the gRPC service interface.

## Code Changes That Would Require Documentation Updates
*   **gRPC API Modifications**: Any changes to the `CommitManifests` or `CommitPathDetails` message structures in the `.proto` definitions (e.g., adding, renaming, or removing fields).
*   **Authentication Requirements**: Changes to how the server handles authentication for HTTPS or SSH git URLs.
*   **Manifest Format Support**: If the server begins accepting formats other than JSON (e.g., YAML strings) within the `manifests` repeated field.
*   **Metadata Changes**: Adding or modifying required metadata fields like `drySHA`, `commitAuthor`, or `commitTime`.
*   **Response Handling**: Updates to the `CommitManifestsResponse` to include error codes, success messages, or status details.
*   **Path Logic**: Changes to how the server interprets or writes to the directory structure defined in `CommitPathDetails.path`.

## Key Technical Concepts
*   **CommitManifests**: The primary gRPC request message containing repository targets and manifest metadata.
*   **CommitPathDetails**: A sub-message defining specific paths, manifest content (JSON), and README documentation for the repository.
*   **Hydrated Manifests**: Final Kubernetes resource definitions generated from "dry" sources (like Helm or Kustomize).
*   **drySHA**: The full SHA256 hash of the source ("dry") commit used for traceability between source and hydrated state.
*   **Target Branch**: The specific git branch where the Commit Server will push the hydrated changes.
*   **gRPC Service**: The communication protocol used by callers to interact with the Commit Server.

## Related Components
*   **Argo CD Manifest Hydrator**: The architectural pattern/component that generates the manifests sent to this server.
*   **Git Repository**: The external storage system where the Commit Server pushes manifests.
*   **Kubernetes Manifests**: The workload definitions (Service, Deployment, etc.) being managed.
*   **Protobuf/gRPC Framework**: The underlying technology for the service definition.