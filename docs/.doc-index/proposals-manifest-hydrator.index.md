# PROPOSALS/MANIFEST-HYDRATOR Documentation Index

## Overview
This documentation describes the proposal and implementation details for the Argo CD Manifest Hydrator, a tool designed to render high-level configurations (Helm, Kustomize) into plain Kubernetes manifests and commit them back to Git. Its primary purpose is to provide a fully auditable history of an application's state and improve visibility into what is actually running in a cluster.

## Files Summary
*   **proposals/manifest-hydrator/README.md**: Provides a high-level conceptual overview of manifest hydration, the security model for Git push access, and instructions for configuring push credentials via Kubernetes secrets.
*   **proposals/manifest-hydrator/commit-server/README.md**: Defines the technical specifications for the Commit Server, specifically documenting the gRPC service interface and message structures used to transmit hydrated manifests to Git.

## Code Changes That Would Require Documentation Updates
*   **Application CRD Schema**: Any changes to the `ApplicationSpec`, specifically the addition or modification of the `sourceHydrator`, `source`, or `sources` fields.
*   **Credential Management**: Modifications to how Argo CD handles Git credentials, particularly changes to the `argocd.argoproj.io/secret-type: repository-push` label or the discovery mechanism in the `argocd-push` namespace.
*   **gRPC API Definitions**: Updates to the `CommitManifests` or `CommitPathDetails` protobuf messages in the `commit-server`, including adding new fields for metadata or changing the structure of the manifest payload.
*   **Security Policies**: Changes to the security isolation between push and pull credentials or the introduction of new RBAC/permission models for the Commit Server.
*   **Hydration Logic**: If the engine used to render Helm/Kustomize changes in a way that affects the resulting commit structure (e.g., changes to how `README.md` files are generated or where manifests are placed in the directory tree).
*   **Migration Paths**: Implementation of the logic to transition from standard `source`/`sources` fields to the `sourceHydrator` field.

## Key Technical Concepts
*   **Manifest Hydration (Rendering)**: The process of transforming "dry" sources (templates/Kustomizations) into "hydrated" (plain YAML) manifests.
*   **Repository-Push Secret**: A specialized secret type (`repository-push`) stored in the `argocd-push` namespace used exclusively for write access to Git.
*   **Commit Server**: The backend component that exposes a gRPC service to handle the physical pushing of manifests to remote repositories.
*   **Dry SHA**: The SHA256 hash of the source ("dry") commit used as a reference point for the hydrated commit.
*   **CommitManifests gRPC Service**: The interface defining the communication between the hydration controller and the commit backend.
*   **Target Branch**: The destination branch in the Git repository where hydrated manifests are stored.
*   **Hydrated State Visibility**: The concept of ensuring state changes are auditable and discoverable in Git rather than existing only in live cluster state.

## Related Components
*   **Argo CD Application Controller**: Responsible for identifying when hydration is required based on Application specs.
*   **Argo CD Repo Server**: The component likely performing the rendering/hydration of Helm and Kustomize.
*   **Commit Server**: A dedicated service for executing Git push operations.
*   **Kubernetes Secret Store**: Specifically the `argocd-push` namespace used for securing push-capable credentials.
*   **SCM Providers**: External Git hosts (GitHub, GitLab, Bitbucket) that receive the pushed hydrated manifests.