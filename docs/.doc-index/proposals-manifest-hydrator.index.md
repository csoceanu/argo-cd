# PROPOSALS/MANIFEST-HYDRATOR Documentation Index

## Overview
This documentation area describes the "Manifest Hydrator" feature for Argo CD, which automates the process of rendering (hydrating) Helm charts or Kustomize templates into plain Kubernetes manifests and committing them back to a Git repository. Its primary purpose is to provide a fully auditable history of an application's final state and improve visibility for developers by ensuring the rendered output is stored in version control.

## Files Summary
*   **proposals/manifest-hydrator/README.md**: Outlines the rationale for manifest hydration, provides security guidelines for Git push access, and details the configuration of specific "repository-push" secrets required for the feature.

## Code Changes That Would Require Documentation Updates
*   **Application CRD Modifications**: Any changes to the `sourceHydrator` field or structural changes to the `source`/`sources` fields in the Argo CD Application specification.
*   **Credential Management**: Changes to how Argo CD identifies or consumes push secrets, such as changing the required label `argocd.argoproj.io/secret-type: repository-push`.
*   **Namespace Requirements**: Modifications to the hardcoded or default namespace (currently `argocd-push`) where push secrets must be stored.
*   **Hydration Logic**: Updates to the core rendering engine (Helm or Kustomize integrations) that change how manifests are transformed before being pushed.
*   **Git Integration Logic**: Changes to the way Argo CD interacts with SCM providers for push operations, including branch handling, commit signing, or commit message formatting.
*   **Security Architecture**: Alterations to the security model that separates push and pull credentials or changes to the permission model between the repo-server and the push secrets.
*   **Migration Path**: Updates to the logic or tooling used to migrate existing Applications from standard sources to the hydrated source model.

## Key Technical Concepts
*   **Manifest Hydration**: The process of rendering templates (Helm/Kustomize) into static Kubernetes YAML.
*   **Source Hydrator**: The specific Application field used to configure the hydration and push back mechanism.
*   **repository-push**: A specialized secret type (`argocd.argoproj.io/secret-type`) distinct from standard pull-only repository credentials.
*   **argocd-push Namespace**: The dedicated namespace used to isolate push-access credentials from other Argo CD components.
*   **Credential Separation**: The security practice of using different secrets for Git pull (read) and Git push (write) operations to minimize blast radius.
*   **Rendered Manifests**: The final Kubernetes objects produced after template expansion.
*   **SCM Security Mechanisms**: Branch protection and repository-level permissions used to control Argo CD's write access.

## Related Components
*   **Argo CD Application Controller**: Responsible for reconciling the Application state and triggering the hydration workflow.
*   **Argo CD Repo Server**: The component likely responsible for the heavy lifting of rendering templates and executing Git push commands.
*   **Argo CD Secret Manager**: Handles the discovery and injection of `repository-push` credentials.
*   **Git SCM Providers**: External systems (GitHub, GitLab, Bitbucket) that receive the hydrated manifest commits.