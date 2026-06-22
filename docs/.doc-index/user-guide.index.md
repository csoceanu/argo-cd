# USER-GUIDE Documentation Index

## Overview
This documentation area provides comprehensive guidance for developers and users on managing applications within Argo CD. It covers application lifecycle management, supported manifest tools (Helm, Kustomize, OCI, etc.), security configurations like GPG verification and private repository access, and advanced deployment strategies such as sync waves and automated synchronization.

## Files Summary
*   **user-guide/private-repositories.md**: Explains how to configure credentials (HTTPS, SSH, GitHub Apps, GCP, Azure) and TLS certificates for accessing private Git and Helm repositories.
*   **user-guide/annotations-and-labels.md**: A comprehensive reference of the specific Kubernetes annotations and labels used by Argo CD to control application behavior and tracking.
*   **user-guide/ci_automation.md**: Describes recommended workflows for integrating CI pipelines with Argo CD, including image updates and triggering syncs via the CLI.
*   **user-guide/compare-options.md**: Details how to configure Argo CD to ignore extraneous resources during the comparison process to maintain a "Synced" status.
*   **user-guide/oci.md**: Provides instructions on using OCI-compliant registries as application sources, including ORAS usage and metadata mapping.
*   **user-guide/jsonnet.md**: Explains how Argo CD processes Jsonnet files, including support for external variables, Top-Level Arguments (TLAs), and shared libraries.
*   **user-guide/gpg-verification.md**: Guides users through enforcing GnuPG signature verification for Git commits at the project level to ensure manifest integrity.
*   **user-guide/multiple_sources.md**: Describes the "Multiple Sources" feature that allows combining resources from different repositories into a single Argo CD application.
*   **user-guide/external-url.md**: Shows how to add clickable external links (e.g., to logs or metrics) to the Argo CD UI using resource annotations.
*   **user-guide/directory.md**: Covers "Directory" type applications, including recursive file detection and include/exclude patterns for plain YAML/JSON manifests.
*   **user-guide/orphaned-resources.md**: Explains how to monitor and manage resources within a namespace that are not defined in the corresponding GitOps manifests.
*   **user-guide/auto_sync.md**: Details the automated synchronization policy, including self-healing, automatic pruning, and sync retry logic.
*   **user-guide/subscriptions.md**: Describes how to subscribe to application events and notifications using annotations and global configuration.
*   **user-guide/kustomize.md**: Exhaustive guide on Kustomize support, covering patches, components, private remote bases, and version management.
*   **user-guide/selective_sync.md**: Explains how to manually synchronize only a subset of resources within an application.
*   **user-guide/config-management-plugins.md**: A placeholder file indicating that Config Management Plugin documentation has moved to the Operator Manual.
*   **user-guide/best_practices.md**: Provides architectural recommendations such as separating configuration from source code and ensuring manifest immutability.
*   **user-guide/projects.md**: Details the `AppProject` resource, covering logical grouping, RBAC roles, JWT tokens, and project-scoped cluster/repo restrictions.
*   **user-guide/application-set.md**: Introduces the ApplicationSet controller and its generators for automating the creation of multiple Argo CD applications.
*   **user-guide/tool_detection.md**: Explains the logic Argo CD uses to automatically detect which manifest tool (Helm, Kustomize, etc.) to use for a repository.
*   **user-guide/status-badge.md**: Explains how to enable and customize public-facing status badges that display an application's health and sync state.
*   **user-guide/application-specification.md**: Points to the full YAML specification and schema for the Argo CD `Application` Custom Resource.
*   **user-guide/sync-kubectl.md**: Shows how to trigger Argo CD sync operations directly using `kubectl apply` or `kubectl patch` on the Application resource.
*   **user-guide/resource_hooks.md**: A redirect page for resource hooks, pointing users toward the unified sync-waves documentation.
*   **user-guide/resource_tracking.md**: Compares different resource tracking methods, including the newer `tracking-id` annotation versus the traditional instance label.
*   **user-guide/parameters.md**: Explains how to use parameter overrides to dynamically change manifest values without modifying Git.
*   **user-guide/import.md**: Provides solutions for Go developers importing Argo CD packages who encounter "unknown revision" errors in dependency management.
*   **user-guide/sync_windows.md**: Details how to define "Sync Windows" to allow or block synchronization during specific time periods using cron syntax.
*   **user-guide/helm.md**: Extensive guide on Helm support, including values precedence, OCI charts, plugins, and mapping Helm hooks to Argo CD hooks.
*   **user-guide/source-hydrator.md**: Describes an alpha feature for Hydrated Manifests, allowing Argo CD to push rendered manifests back to a "staging" or "sync" branch in Git.
*   **user-guide/diff-strategies.md**: Compares different diffing algorithms, specifically highlighting the benefits of Server-Side Diff over Legacy and Structured-Merge Diff.
*   **user-guide/sync-options.md**: A detailed reference for sync-time behaviors like `CreateNamespace`, `ServerSideApply`, `PruneLast`, and `Replace`.
*   **user-guide/scale_application_resources.md**: Documents the UI feature allowing users to manually scale Deployments and StatefulSets directly from the dashboard.
*   **user-guide/sync-waves.md**: Explains the execution order of deployments using sync phases (PreSync, PostSync, etc.) and numerical sync waves.
*   **user-guide/tracking_strategies.md**: Discusses strategies for tracking versions in Git (branches, tags, SHAs) and Helm (SemVer ranges).
*   **user-guide/application_sources.md**: Summarizes the various tools supported for application manifest generation in both production and development environments.
*   **user-guide/extra_info.md**: Explains how to add custom key-value metadata to the Application details page in the UI via the `info` field.
*   **user-guide/build-environment.md**: Lists the standard environment variables (like `ARGOCD_APP_NAME`) available to manifest tools during the build process.
*   **user-guide/index.md**: The entry point for the user guide, providing a brief overview and link to the getting started guide.
*   **user-guide/diffing.md**: Details how to customize the diffing engine to ignore specific JSON paths or managed fields that cause "false" out-of-sync states.
*   **user-guide/skip_reconcile.md**: Documents the alpha feature that allows users to pause all reconciliation for a specific application via an annotation.
*   **user-guide/environment-variables.md**: Lists the environment variables supported by the `argocd` CLI for configuration and authentication.
*   **user-guide/plugins.md**: Explains how to extend the `argocd` CLI functionality by creating and distributing custom binary plugins.
*   **user-guide/app_deletion.md**: Describes the differences between cascade and non-cascade application deletion and the use of the deletion finalizer.

## Code Changes That Would Require Documentation Updates
*   **CRD Schema Changes**: Any modifications to `applications.argoproj.io` or `appprojects.argoproj.io` fields (e.g., adding a new `syncOption` or `source` parameter).
*   **Manifest Tool Integration**: Updates to the versions of Helm, Kustomize, or Jsonnet bundled with the repo-server, or changes in how they are invoked.
*   **Sync Logic**: Changes to the application controller’s reconciliation loop, specifically regarding sync phases, waves, or the order of operations.
*   **Tracking Logic**: Modifications to how Argo CD identifies owned resources (e.g., changes to the `tracking-id` generation or label truncation).
*   **RBAC and Projects**: Adding new resources to the RBAC engine (e.g., `applicationsets`, `logs`) or changing how project roles/JWTs are validated.
*   **CLI Development**: Adding new subcommands to the `argocd` binary or changing the behavior of global flags like `--server` or `--auth-token`.
*   **Diffing Engine**: Changes to the Server-Side Diff implementation or the addition of new system-level resource customizations.
*   **Credential Handling**: Updates to supported Git hosters, credential types (e.g., new Workload Identity providers), or TLS/SSH scanning logic.
*   **UI Features**: Dashboard modifications that change how "Info" sections, "External Links," or "Scale" actions are rendered or configured.

## Key Technical Concepts
*   **Application & AppProject**: The core CRDs managing GitOps state and logical boundaries.
*   **Sync Policy**: Automated, manual, self-healing, and pruning behaviors.
*   **Sync Waves & Phases**: `PreSync`, `Sync`, `PostSync`, `SyncFail`, and numerical wave ordering.
*   **Sync Options**: `ServerSideApply`, `CreateNamespace`, `Replace`, `PruneLast`, `ApplyOutOfSyncOnly`.
*   **Resource Tracking**: `argocd.argoproj.io/tracking-id`, `app.kubernetes.io/instance`, and installation IDs.
*   **Diffing Strategies**: Legacy 3-way diff, Server-Side Diff, `ignoreDifferences`, and JQ path expressions.
*   **Manifest Tools**: Helm (OCI/Classic), Kustomize (Patches/Components), Jsonnet, OCI, and Directory.
*   **Security**: GPG verification, SSH Known Hosts, TLS Client Certs, and RBAC Roles.
*   **ApplicationSet**: Generators (Git, List, Cluster, Matrix) for automated app factories.

## Related Components
*   **argocd-application-controller**: Responsible for reconciliation, sync operations, and health monitoring.
*   **argocd-repo-server**: Handles manifest generation, repository cloning, and tool detection.
*   **argocd-server**: The API server and UI backend; manages RBAC and external authentication.
*   **argocd-applicationset-controller**: Automates the lifecycle of Applications based on templates.
*   **argocd-notifications**: Manages subscriptions and notification delivery based on application state.
*   **argocd-redis**: Caches generated manifests and cluster state to optimize performance.