# USER-GUIDE Documentation Index

## Overview
This documentation area provides the essential guidance for developers and end-users managing applications with Argo CD. It covers the end-to-end lifecycle of Kubernetes resources—from configuring private repository access and defining application projects to advanced synchronization strategies, manifest hydration, and CLI extensibility.

## Files Summary
*   **private-repositories.md**: Detailed instructions for configuring HTTPS, SSH, GitHub Apps, and Cloud provider credentials for private Git/Helm repositories.
*   **annotations-and-labels.md**: A comprehensive reference of internal metadata used by Argo CD for resource tracking, synchronization behavior, and UI customization.
*   **ci_automation.md**: Patterns for integrating Argo CD into CI pipelines, including image updates and CLI-based synchronization.
*   **compare-options.md**: Documentation on the `IgnoreExtraneous` option to prevent specific resources from affecting sync status.
*   **oci.md**: Instructions for using OCI-compliant registries (DockerHub, ECR, etc.) as application sources for manifests or Helm charts.
*   **jsonnet.md**: Usage of Jsonnet for manifest generation, including build environments, variables, and shared libraries.
*   **gpg-verification.md**: Configuring GnuPG signature verification to ensure only signed commits are synchronized to clusters.
*   **multiple_sources.md**: Combining multiple Git repositories or Helm charts into a single Application, including external value file references.
*   **external-url.md**: Methods for adding clickable monitoring or documentation links to resources in the Argo CD UI.
*   **directory.md**: Management of plain YAML/JSON manifest directories, including file inclusion/exclusion and recursive detection.
*   **orphaned-resources.md**: Configuration for detecting and alerting on top-level resources in a namespace not managed by Argo CD.
*   **auto_sync.md**: Explains the automated sync policy, including self-healing, automatic pruning, and retry mechanisms.
*   **subscriptions.md**: Defining notification subscriptions for application events via annotations or global configuration.
*   **kustomize.md**: Deep dive into Kustomize support, covering patches, components, private remote bases, and version management.
*   **selective_sync.md**: Manual synchronization of specific resources within an application rather than the whole set.
*   **best_practices.md**: Guidance on repository separation, manifest immutability, and handling imperative cluster elements like HPA.
*   **projects.md**: Documentation of the `AppProject` CRD for multi-tenancy, RBAC, and scoping of clusters/repositories.
*   **application-set.md**: High-level overview of the ApplicationSet controller for automating application generation at scale.
*   **tool_detection.md**: Logic used by Argo CD to automatically identify Helm, Kustomize, or plain directory-type sources.
*   **status-badge.md**: Embedding health and sync status images into external sites or README files.
*   **application-specification.md**: Points to the full schema definition for the `Application` Custom Resource.
*   **sync-kubectl.md**: Using standard `kubectl` commands to trigger and configure Argo CD operations via the API.
*   **resource_tracking.md**: Technical methods for how Argo CD identifies managed resources (Annotations vs. Labels).
*   **parameters.md**: Overriding configuration values (Helm/Kustomize) via CLI, UI, or `.argocd-source.yaml`.
*   **import.md**: Solutions for Go developers importing Argo CD packages who encounter dependency versioning issues.
*   **sync_windows.md**: Defining time-based schedules to allow or deny synchronization operations for specific apps or clusters.
*   **helm.md**: Extensive guide on Helm support, covering values, release names, hooks, OCI, and custom plugins.
*   **source-hydrator.md**: Documentation of the "Rendered Manifest Pattern" (Alpha) for pushing hydrated manifests back to Git.
*   **diff-strategies.md**: Comparison of Legacy, Structured-Merge, and Server-Side Diff engines.
*   **sync-options.md**: Reference for specific sync behaviors like `ServerSideApply`, `Replace`, `CreateNamespace`, and `FailOnSharedResource`.
*   **scale_application_resources.md**: Guidance on using the UI-based scaling feature for Deployments and StatefulSets.
*   **sync-waves.md**: Logic for Sync Phases (Pre/Post) and Waves (numeric ordering) to control deployment sequences.
*   **tracking_strategies.md**: Comparison of version tracking methods including branch tracking, tag SemVer ranges, and commit pinning.
*   **application_sources.md**: Summary of supported manifest tools and instructions for local/development manifest uploads.
*   **extra_info.md**: Adding custom key-value metadata to the Application dashboard for better visibility.
*   **build-environment.md**: List of standard environment variables (e.g., `ARGOCD_APP_NAME`) available during manifest generation.
*   **diffing.md**: Customizing how Argo CD ignores differences at the application or system level to avoid "false" out-of-sync states.
*   **skip_reconcile.md**: Documentation on pausing all controller processing for an application via annotation.
*   **environment-variables.md**: Global environment variables for configuring the `argocd` CLI tool.
*   **plugins.md**: Guide for creating custom sub-commands and extensions for the `argocd` CLI.
*   **app_deletion.md**: Explanation of cascading vs. non-cascading deletion and the use of finalizers.

## Code Changes That Would Require Documentation Updates
*   **CRD Schema Changes**: Any additions or modifications to the `Application` or `AppProject` specs (affects `application-specification.md`, `projects.md`).
*   **New Annotations/Labels**: Introduction of new `argocd.argoproj.io/` metadata for sync, diff, or UI logic (affects `annotations-and-labels.md`).
*   **Manifest Tool Updates**: Changes to bundled versions of Helm, Kustomize, or Jsonnet, or changes in how they are invoked (affects `helm.md`, `kustomize.md`, `jsonnet.md`).
*   **Sync Engine Logic**: Modifications to waves, phases, hooks, or the auto-sync state machine (affects `sync-waves.md`, `auto_sync.md`).
*   **Diffing Logic**: Changes to `ServerSideDiff` behavior or how `ignoreDifferences` is evaluated (affects `diff-strategies.md`, `diffing.md`).
*   **CLI Command Updates**: Adding new flags or sub-commands to the `argocd` CLI (affects `environment-variables.md`, `plugins.md`, and tool-specific docs).
*   **Credential Handling**: Changes in how Git/Helm/OCI credentials are encrypted, stored, or retrieved (affects `private-repositories.md`).
*   **UI Features**: New dashboard elements, scaling modals, or extra info fields (affects `scale_application_resources.md`, `extra_info.md`).
*   **Tracking Logic**: Changing the default resource tracking method or implementation (affects `resource_tracking.md`).

## Key Technical Concepts
*   **Resources**: `Application`, `AppProject`, `ApplicationSet`, `Secret` (repo/cluster type).
*   **Sync Behavior**: `Sync Waves`, `Sync Phases`, `Hooks`, `Self-Heal`, `Pruning`, `Automated Sync`, `Sync Windows`.
*   **Diffing**: `Server-Side Diff`, `IgnoreDifferences`, `JSON Pointers`, `JQ Path Expressions`, `Legacy Diff`.
*   **Tracking**: `tracking-id`, `instanceLabelKey`, `managedFields`.
*   **Manifest Tools**: `Helm (v2/v3)`, `Kustomize`, `Jsonnet`, `OCI`, `Config Management Plugins`.
*   **Security**: `GnuPG Signature Verification`, `RBAC`, `JWT Tokens`, `Credential Templates`, `TLS/SSH Certificates`.
*   **Automation**: `Source Hydrator`, `Webhook`, `CI Pipeline Integration`, `Local Manifest Sync`.

## Related Components
*   **argocd-application-controller**: The primary state machine for sync and health.
*   **argocd-repo-server**: Responsible for manifest generation and hydration.
*   **argocd-server**: The API and UI gateway.
*   **argocd-notifications**: The subsystem handling trigger-based alerting.
*   **argocd CLI**: The primary interface for imperative management.
*   **argocd-applicationset-controller**: Automation for massive-scale application deployments.