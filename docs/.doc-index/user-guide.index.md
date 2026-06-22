# USER-GUIDE Documentation Index

## Overview
This documentation area provides a comprehensive guide for developers and users to manage the lifecycle of applications within Argo CD. It covers application manifest generation (Helm, Kustomize, Jsonnet, OCI), deployment strategies, security configurations (GPG, private repositories, RBAC), and advanced synchronization techniques like sync waves, windows, and server-side diffs.

## Files Summary

*   **private-repositories.md**: Explains how to configure access to private Git and Helm repositories using HTTPS, SSH, GitHub Apps, and cloud-provider-specific identities.
*   **annotations-and-labels.md**: A reference guide for all Argo CD-specific Kubernetes annotations and labels used to control application behavior.
*   **ci_automation.md**: Describes best practices and workflows for integrating Argo CD into CI/CD pipelines, including image updates and CLI synchronization.
*   **compare-options.md**: Details how to use the `IgnoreExtraneous` option to prevent generated resources from affecting application sync status.
*   **oci.md**: Instructions for using OCI-compliant registries as a source for Kubernetes manifests and Helm charts.
*   **jsonnet.md**: Covers the usage of Jsonnet for manifest generation, including external variables and shared library support.
*   **gpg-verification.md**: Documentation on enforcing GnuPG signature verification for Git commits to ensure only trusted code is deployed.
*   **multiple_sources.md**: Explains how to combine resources from multiple repositories or paths into a single Argo CD application.
*   **external-url.md**: Shows how to add custom external links to the Argo CD UI for specific resources via annotations.
*   **directory.md**: Describes how Argo CD handles plain YAML/JSON directories, including recursive detection and file include/exclude patterns.
*   **orphaned-resources.md**: Details the feature for monitoring and managing namespaced resources that are not part of any Argo CD application.
*   **auto_sync.md**: Explains automated synchronization policies, including self-healing, automated pruning, and retry logic.
*   **subscriptions.md**: Covers how to subscribe to application event notifications using the Argo CD Notifications controller.
*   **kustomize.md**: A deep dive into Kustomize support, including image overrides, patches, components, and custom Kustomize versions.
*   **selective_sync.md**: Instructions for synchronizing only specific resources within an application rather than the whole set.
*   **config-management-plugins.md**: A redirection page pointing to the operator manual for configuring custom config management tools.
*   **best_practices.md**: Recommendations for repository structure, manifest immutability, and handling imperative Kubernetes features like HPA.
*   **projects.md**: Documentation on `AppProject` resources, used for logical grouping, RBAC, and restricting deployment sources/destinations.
*   **application-set.md**: Overview of the ApplicationSet controller for automating the creation of multiple applications across clusters.
*   **tool_detection.md**: Explains the logic Argo CD uses to automatically detect if a repository contains Helm, Kustomize, or plain manifests.
*   **status-badge.md**: How to enable and customize the health and sync status badges for embedding in external sites/READMEs.
*   **application-specification.md**: A reference placeholder for the full Kubernetes Application CRD specification.
*   **sync-kubectl.md**: Demonstrates how to trigger Argo CD sync operations directly using `kubectl` by modifying the `operation` field.
*   **resource_hooks.md**: A redirection page pointing to the documentation on sync waves and resource phases.
*   **resource_tracking.md**: Explains the various methods Argo CD uses to track managed resources (annotations vs. labels).
*   **parameters.md**: Details how to override application parameters via the CLI, UI, or `.argocd-source.yaml` files.
*   **import.md**: Provides technical solutions for developers importing Argo CD as a Go library, focusing on dependency management.
*   **sync_windows.md**: Documentation on time-based windows that allow or deny synchronization based on schedules.
*   **helm.md**: Extensive guide on Helm support, covering value precedence, release names, hooks, plugins, and OCI charts.
*   **source-hydrator.md**: Explains the "Source Hydrator" (rendered manifest pattern) for pushing hydrated manifests back to Git.
*   **diff-strategies.md**: Compares diffing strategies, including Legacy, Structured-Merge, and the newer Server-Side Diff.
*   **sync-options.md**: A comprehensive list of synchronization behavior overrides, such as `ServerSideApply` and `CreateNamespace`.
*   **scale_application_resources.md**: Explains the UI feature allowing users to scale Deployments and StatefulSets manually.
*   **sync-waves.md**: Details the execution order of resources using phases (PreSync, PostSync) and numerical waves.
*   **tracking_strategies.md**: Covers strategies for tracking application versions via Git tags, SHAs, branches, and Helm SemVer ranges.
*   **application_sources.md**: High-level overview of supported manifest tools and the local manifest override feature.
*   **extra_info.md**: Explains how to add custom metadata fields to the Application details panel in the UI.
*   **build-environment.md**: Lists standard environment variables available to all manifest generation tools during the build process.
*   **index.md**: The landing page and introduction for the Argo CD user guide.
*   **diffing.md**: Instructions for customizing the diff engine to ignore specific fields or managed managers to avoid perma-drift.
*   **skip_reconcile.md**: Documentation on the annotation used to pause all reconciliation activity for a specific application.
*   **environment-variables.md**: Reference for environment variables that modify the behavior of the `argocd` CLI.
*   **plugins.md**: Guide for creating and distributing custom sub-commands for the Argo CD CLI.
*   **app_deletion.md**: Explains the difference between cascading and non-cascading deletion and the role of finalizers.

## Code Changes That Would Require Documentation Updates

*   **CRD Changes**: Updates to `applications.argoproj.io` or `appprojects.argoproj.io` (e.g., new fields in `spec` or `status`).
*   **Manifest Tool Upgrades**: Changes in the default versions or supported flags for Helm, Kustomize, or Jsonnet.
*   **Controller Logic**: Modifications to the `argocd-application-controller` regarding sync phases, waves, or reconciliation loops.
*   **Diffing Engine**: Updates to how Argo CD calculates differences between Git and the live cluster state (e.g., Server-Side Diff logic).
*   **Security & Auth**: Changes to credential handling, GPG verification logic, or RBAC resource names (like `gpgkeys`).
*   **UI/UX Enhancements**: New interactive features in the dashboard, such as the resource scaling modal or link annotations.
*   **CLI Development**: Changes to `argocd` command syntax, environment variables, or the CLI plugin architecture.
*   **Resource Tracking**: Adjustments to how Argo CD labels or annotates resources for ownership tracking.
*   **Notifications**: Adding new triggers or services to the `argocd-notifications` module.

## Key Technical Concepts

*   **Application & AppProject**: The primary Custom Resource Definitions for managing deployments and multi-tenancy.
*   **Sync Phases & Waves**: The mechanism for ordering resource application (PreSync, Sync, PostSync).
*   **Resource Hooks**: Jobs or pods triggered at specific points in the deployment lifecycle.
*   **Server-Side Apply (SSA)**: Kubernetes-native application strategy supported as a sync option.
*   **Self-Healing & Pruning**: Automated behaviors to ensure cluster state matches Git.
*   **OCI Artifacts**: Using container registries to store Helm charts and Kubernetes manifests.
*   **JSON Pointers & JQ Path**: Methods used to ignore specific fields during the diffing process.
*   **Credential Templates**: Shared authentication configuration for multiple repositories.
*   **Sync Windows**: Cron-based schedules for controlling when syncs can occur.

## Related Components

*   **argocd-application-controller**: Manages application state and executes syncs.
*   **argocd-repo-server**: Generates Kubernetes manifests from Git/Helm/OCI sources.
*   **argocd-server**: Provides the API and UI for user interaction.
*   **ApplicationSet Controller**: Automates large-scale application generation.
*   **Argo CD CLI**: Command-line tool for management and automation.
*   **Argo CD Notifications**: Handles event-driven alerts and subscriptions.
*   **Source Hydrator**: Component responsible for the rendered manifest GitOps pattern.