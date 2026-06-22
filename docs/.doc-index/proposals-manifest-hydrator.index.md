# PROPOSALS/MANIFEST-HYDRATOR Documentation Index

## Overview
This documentation area describes the "Manifest Hydrator" feature in Argo CD, which automates the process of rendering Helm charts or Kustomize templates into plain Kubernetes manifests and pushing them back to a Git repository. Its primary purpose is to provide a transparent, auditable history of the final application state, enhancing visibility and security within the GitOps workflow.

## Files Summary
*   **proposals/manifest-hydrator/README.md**: Serves as the primary guide for the feature, covering the rationale for manifest hydration, security considerations for Git push access, credential configuration, and the transition to the `sourceHydrator` field.

## Code Changes That Would Require Documentation Updates
*   **CRD Schema Modifications**: Any changes to the `Application` CRD, specifically adding, removing, or renaming fields within `spec.sourceHydrator`, `spec.source`, or `spec.sources`.
*   **Credential Management Logic**: Changes to how Argo CD identifies push secrets, such as changing the required label (`argocd.argoproj.io/secret-type: repository-push`) or the expected secret format.
*   **Namespace Defaults**: Alterations to the hardcoded or default namespace for push credentials (currently `argocd-push`).
*   **Hydration Logic**: Updates to how Argo CD handles the transformation of Helm/Kustomize sources before pushing (e.g., support for new rendering engines or parameters).
*   **Security Model Changes**: Any shifts in the architecture regarding the separation of push and pull credentials or how SCM authentication is handled.
*   **Migration Path**: Changes in the internal logic that handles the transition from the legacy `source`/`sources` fields to the new `sourceHydrator` model.
*   **Git Provider Support**: Adding specific requirements or limitations for different SCM providers (GitHub, GitLab, Bitbucket) regarding push access.

## Key Technical Concepts
*   **Manifest Hydration/Rendering**: The process of converting high-level templates (Helm/Kustomize) into final Kubernetes YAML.
*   **sourceHydrator**: The specific field in the Application spec used to configure the hydration and push back mechanism.
*   **repository-push**: A specialized Argo CD secret type used exclusively for write access to Git repositories.
*   **argocd-push Namespace**: The designated secure namespace for storing sensitive push credentials.
*   **Credential Isolation**: The security practice of separating Git pull (read) and Git push (write) secrets to minimize the blast radius of a credential leak.
*   **Live Cluster State vs. Git State**: The comparison between what is currently running and what is stored in the repository.

## Related Components
*   **Argo CD Application Controller**: Responsible for managing the lifecycle and synchronization of the Application.
*   **Argo CD Repo Server**: The component likely tasked with performing the actual manifest generation/rendering.
*   **Kubernetes API Server**: For managing the Secrets and Application CRDs.
*   **SCM Providers (Git)**: External systems where hydrated manifests are stored and audited.
*   **Rendering Tools**: External binaries like `helm` and `kustomize` utilized during the hydration process.