# USER-GUIDE Documentation Index

## Overview
This documentation area provides a comprehensive guide for developers and administrators to manage the lifecycle of Kubernetes applications using Argo CD. It covers conceptual guidance on deployment strategies and manifest rendering tools (Helm, Kustomize, Jsonnet), as well as a complete technical reference for the `argocd` CLI used to manage applications, multi-tenancy via projects, repository connectivity, and administrative diagnostics.

## Files Summary
*   **annotations-and-labels.md**: Reference table of standard annotations and labels for sync options, tracking, and UI customization.
*   **app_deletion.md**: Explanation of cascading vs. non-cascading deletion and the use of finalizers.
*   **application-set.md**: Overview of the ApplicationSet controller for automating application generation.
*   **application-specification.md**: Reference for the full `Application` CRD schema.
*   **application_sources.md**: High-level overview of supported tools and local manifest uploading.
*   **argocd_account.md (and sub-pages)**: Management of local accounts, passwords (bcrypt), tokens, and RBAC permission testing (`can-i`).
*   **argocd_admin.md (and sub-pages)**: Administrative tools for troubleshooting reconciliation, cluster sharding stats, export/import (disaster recovery), and notification configuration.
*   **argocd_app.md (and sub-pages)**: CLI reference for application lifecycle: create, sync, logs, history, rollback, manifests, and resource patching.
*   **argocd_appset.md (and sub-pages)**: Commands for managing and previewing (generate) ApplicationSet resources.
*   **argocd_cert.md (and sub-pages)**: Management of TLS certificates and SSH known hosts for repository connections.
*   **argocd_cluster.md (and sub-pages)**: Lifecycle management of external clusters, including sharding, namespaces, and auth rotation.
*   **argocd_completion.md**: Generation of shell completion scripts for Bash, Zsh, and Fish.
*   **argocd_configure.md**: Manages local CLI settings and interactive prompts.
*   **argocd_context.md / argocd_login.md / argocd_logout.md**: Authentication and session management for the CLI.
*   **argocd_gpg.md (and sub-pages)**: Server-wide management of GPG public keys for signature verification.
*   **argocd_proj.md (and sub-pages)**: Management of `AppProject` resources, including destination allow-lists, RBAC roles, and sync windows.
*   **argocd_repo.md / argocd_repocreds.md**: Management of Git/Helm/OCI repositories and credential templates.
*   **auto_sync.md**: Details on automated synchronization policies, pruning, and self-healing.
*   **best_practices.md**: Guidance on repository separation, manifest immutability, and state management.
*   **build-environment.md**: List of environment variables available to tools during manifest generation.
*   **ci_automation.md**: Workflows for integrating Argo CD with CI pipelines and image updates.
*   **compare-options.md**: Documentation on excluding specific resources from sync status.
*   **diff-strategies.md**: Comparison of Legacy, Structured-Merge, and Server-Side Diff strategies.
*   **diffing.md**: Configuration to ignore specific fields during the diffing process.
*   **directory.md**: Usage of plain YAML/JSON directory applications and file patterns.
*   **environment-variables.md**: Global environment variables affecting CLI behavior.
*   **external-url.md**: Adding custom clickable links to resources within the UI.
*   **extra_info.md**: Adding custom key-value pairs to the application dashboard.
*   **gpg-verification.md**: Setup for GnuPG signature verification of Git commits.
*   **helm.md**: Comprehensive Helm guide covering values, OCI charts, and authentication.
*   **import.md**: Guidance on importing Argo CD Go packages and dependency management.
*   **jsonnet.md**: Configuration for Jsonnet applications, arguments, and variables.
*   **kustomize.md**: Advanced Kustomize configurations, patches, and components.
*   **multiple_sources.md**: Combining multiple Git/Helm sources into a single application.
*   **oci.md**: Guidelines for using OCI-compliant registries as application sources.
*   **orphaned-resources.md**: Monitoring and managing resources not managed by any application.
*   **parameters.md**: Overriding application parameters via CLI, Git, or UI.
*   **plugins.md**: Guide for developing and distributing custom config management plugins.
*   **private-repositories.md**: Configuring credentials for private Git/Helm repos (HTTPS, SSH, GitHub Apps).
*   **projects.md**: Managing multi-tenancy and resource restrictions via `AppProject`.
*   **resource_tracking.md**: Deep dive into tracking methods (annotation vs. label).
*   **scale_application_resources.md**: Details on UI-driven scaling for Deployments and StatefulSets.
*   **selective_sync.md**: Instructions for synchronizing specific resources.
*   **skip_reconcile.md**: Using annotations to temporarily pause application reconciliation.
*   **source-hydrator.md**: Documentation for the Source Hydrator (Alpha) feature.
*   **status-badge.md**: Configuration for public-facing application health badges.
*   **subscriptions.md**: Configuring notification triggers and recipients via annotations.
*   **sync-kubectl.md**: Triggering sync operations directly via `kubectl`.
*   **sync-options.md**: List of sync-time behaviors (ServerSideApply, Replace, CreateNamespace).
*   **sync-waves.md**: Logic for Sync Phases (Pre/Post) and Sync Waves (ordering).
*   **sync_windows.md**: Defining time-based windows to allow or block sync activities.
*   **tool_detection.md**: Logic used to automatically identify manifest tools.
*   **tracking_strategies.md**: Comparison of tracking Git branches, tags, SHAs, and Helm versions.

## Code Changes That Would Require Documentation Updates
*   **CRD Schema Changes**: Any modification to `Application`, `AppProject`, or `ApplicationSet` specifications.
*   **CLI Command Updates**: Changes to `argocd` CLI sub-commands, flags, or output formats.
*   **API Changes**: Changes to the underlying Protobuf/Swagger definitions for the Argo CD API.
*   **New Annotations/Labels**: Introduction of new `argocd.argoproj.io/` or `link.argocd.argoproj.io/` strings.
*   **Rendering Tool Integrations**: Upgrading bundled Helm/Kustomize versions or changing default execution flags.
*   **Authentication & Security**: Modifications to OIDC handling, SSO, local account security (bcrypt), or Casbin RBAC policies.
*   **Controller Logic**: Changes to sharding algorithms, reconciliation loop behaviors, or health assessment logic.
*   **Diffing & Sync**: Modifications to Server-Side Diff, 3-way merge, sync options, hook execution order, or wave delays.
*   **Infrastructure & Defaults**: Changes to default component names, configuration paths, or global CLI environment variables.
*   **UI/Dashboard**: New visual features like the "Scale" action or badge customizations.

## Key Technical Concepts
*   **Sync Waves & Phases**: Sequential ordering and logic for resource application.
*   **Self-Healing & Pruning**: Automatic drift correction and cleanup of resources.
*   **Server-Side Apply (SSA)**: Kubernetes-native field management for large manifests.
*   **Multi-Source Applications**: Pulling manifests and values from disparate repositories.
*   **AppProject Scoping**: Restricting sources, destinations, and GVKs for multi-tenancy.
*   **GnuPG (GPG)**: Signature verification to ensure only authorized commits are synced.
*   **Resource Actions & Health Checks**: Lua-based scripts for assessing and interacting with live resources.
*   **Cluster Sharding**: Distributing managed clusters across multiple controller replicas.
*   **Config Management Plugins (CMP)**: Sidecar-based extensions for custom manifest generation.
*   **Sync Windows**: Cron-based gates for synchronization operations.
*   **Orphaned Resource Monitoring**: Detection of unmanaged resources in target namespaces.
*   **Source Hydration**: Process of committing rendered manifests back to Git.

## Related Components
*   **argocd-server**: The API server and UI host.
*   **argocd-repo-server**: Handles manifest generation and repository interaction.
*   **argocd-application-controller**: The primary engine for reconciliation, health, and sync.
*   **argocd-notifications-controller**: Handles notification triggers and templates.
*   **argocd-applicationset-controller**: Manages ApplicationSet CRDs.
*   **argocd-redis**: Caching layer for manifest generation, sessions, and cluster state.
*   **Dex**: Used for OIDC and SSO integration.