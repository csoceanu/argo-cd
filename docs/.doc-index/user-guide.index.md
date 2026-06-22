# USER-GUIDE Documentation Index

## Overview
This documentation serves as the primary resource for developers and operators managing the lifecycle of Kubernetes applications using Argo CD. It provides a comprehensive guide to GitOps workflows—including repository connectivity, manifest generation, and synchronization policies—alongside a detailed command-line reference for the Argo CD CLI (`argocd`) to manage applications, projects, cluster administration, and security settings.

## Files Summary
*   **annotations-and-labels.md**: A reference table of all Argo CD-specific annotations and labels used to control resource behavior.
*   **app_deletion.md**: Understanding cascading vs. non-cascading deletion and the deletion finalizer.
*   **application-set.md**: High-level overview of the ApplicationSet controller for automating multi-cluster application generation.
*   **application-specification.md**: Reference to the full Application CRD schema.
*   **application_sources.md**: Overview of supported manifest tools and local development manifest syncing.
*   **argocd.md**: The root command reference listing global options and all primary subcommands.
*   **argocd_account.md**: User-centric commands for password management and account auditing.
*   **argocd_account_bcrypt.md**: Tool to generate bcrypt hashes for account passwords.
*   **argocd_account_can-i.md**: A policy testing tool to check RBAC permissions for specific actions and resources.
*   **argocd_account_delete-token.md**: Command to revoke specific authentication tokens for an account.
*   **argocd_account_generate-token.md**: Command to create new authentication tokens for accounts.
*   **argocd_account_get.md**: Retrieves detailed information about the current or specified account.
*   **argocd_account_list.md**: Lists all configured accounts.
*   **argocd_account_update-password.md**: Changes the password for the current or specified user.
*   **argocd_admin.md**: Root for administrative commands requiring direct Kubernetes access.
*   **argocd_admin_app.md**: Advanced tools for auditing and managing applications at the system level.
*   **argocd_admin_app_diff-reconcile-results.md**: Troubleshooting tool to compare reconciliation outputs across applications.
*   **argocd_admin_app_get-reconcile-results.md**: Retrieves internal reconciliation data for debugging.
*   **argocd_admin_cluster.md**: Tools for generating declarative Kubernetes manifests for cluster secrets.
*   **argocd_admin_cluster_generate-spec.md**: Generates the secret specification for a cluster.
*   **argocd_admin_cluster_kubeconfig.md**: Generates a kubeconfig for clusters managed by Argo CD.
*   **argocd_admin_cluster_namespaces.md**: Root for managing namespaced mode for clusters.
*   **argocd_admin_cluster_namespaces_disable-namespaced-mode.md**: Disables namespace-specific management for clusters.
*   **argocd_admin_cluster_namespaces_enable-namespaced-mode.md**: Enables namespace-specific management for clusters.
*   **argocd_admin_cluster_stats.md**: Displays resource statistics and sharding distribution data.
*   **argocd_admin_dashboard.md**: Commands to launch the Argo CD Web UI locally.
*   **argocd_admin_export.md**: Exports all Argo CD configuration data for backup or migration.
*   **argocd_admin_import.md**: Facilitates disaster recovery by importing Argo CD state from a backup file.
*   **argocd_admin_initial-password.md**: Retrieves the default admin password generated during installation.
*   **argocd_admin_notifications.md**: Commands for testing and configuring the notifications engine.
*   **argocd_admin_notifications_template.md**: Management of notification message templates.
*   **argocd_admin_notifications_trigger.md**: Configuration for event-based notification triggers.
*   **argocd_admin_proj.md**: Root for administrative project configuration and bulk updates.
*   **argocd_admin_proj_update-role-policy.md**: Bulk update tool for project role policies.
*   **argocd_admin_redis-initial-password.md**: Utility to manage the initial Redis password secret.
*   **argocd_admin_repo.md**: Administrative management of repository configurations and secret generation.
*   **argocd_admin_repo_generate-spec.md**: Generates declarative YAML/JSON for repository secrets.
*   **argocd_admin_settings.md**: Tools to validate and troubleshoot system-wide configurations.
*   **argocd_admin_settings_rbac.md**: Root for validating and testing RBAC configurations.
*   **argocd_admin_settings_resource-overrides.md**: Troubleshooting Lua-based resource customizations (health, actions).
*   **argocd_admin_settings_resource-overrides_ignore-differences.md**: Debugs field exclusion logic for diffing.
*   **argocd_admin_settings_validate.md**: Validates `argocd-cm` and `argocd-secret` configurations.
*   **argocd_app.md**: Primary entry point for managing the Application lifecycle.
*   **argocd_app_actions.md**: Manage and list custom resource actions (e.g., restarting a deployment).
*   **argocd_app_confirm-deletion.md**: Finalizes the removal of application-related resources.
*   **argocd_app_create.md**: Extensive guide for creating apps (Helm, Kustomize, Directory, Jsonnet, etc.).
*   **argocd_app_delete-resource.md**: Removes specific Kubernetes resources from an application.
*   **argocd_app_edit.md**: Opens the application manifest in an editor for manual updates.
*   **argocd_app_get.md**: Retrieves detailed status, health, and configuration parameters for an application.
*   **argocd_app_history.md**: Displays deployment history, including previous revisions and sync times.
*   **argocd_app_logs.md**: Provides real-time streaming or tailing of logs from pods in an application.
*   **argocd_app_manifests.md**: Displays generated or live manifests for an application.
*   **argocd_app_patch-resource.md**: Applies patches (JSON/Merge/Strategic) to specific app resources.
*   **argocd_app_remove-source.md**: Removes a specific source from a multi-source application.
*   **argocd_app_resources.md**: Lists resources managed by an application, including orphaned resource filtering.
*   **argocd_app_rollback.md**: Reverts an application to a previous state in its history.
*   **argocd_app_set.md**: Modifies application parameters like sync policies or source details.
*   **argocd_app_sync.md**: Triggers synchronization to the target state in Git.
*   **argocd_app_terminate-op.md**: Stops currently running operations (like a stuck sync).
*   **argocd_app_unset.md**: Removes specific parameter overrides (Kustomize images, Helm values, etc.).
*   **argocd_app_wait.md**: Blocks execution until an application reaches a desired health/sync status.
*   **argocd_appset.md**: Root command for managing ApplicationSet controllers.
*   **argocd_appset_get.md / argocd_appset_list.md**: Commands for retrieving ApplicationSet details and status.
*   **argocd_appset_generate.md**: Renders templates to preview ApplicationSet output.
*   **argocd_appset_delete.md**: Safely removes ApplicationSets and their generated apps.
*   **argocd_cert.md**: Manage repository certificates and SSH known hosts (`add-tls`, `add-ssh`, `list`, `rm`).
*   **argocd_cluster.md**: Entry point for managing credentials and metadata for target Kubernetes clusters.
*   **argocd_cluster_add.md / argocd_cluster_rm.md**: Adding and removing cluster credentials.
*   **argocd_cluster_rotate-auth.md**: Rotates authentication credentials for external clusters.
*   **argocd_completion.md**: Generates shell autocompletion scripts (Bash, Zsh, Fish).
*   **argocd_configure.md**: Manages local CLI settings like interactive prompts.
*   **argocd_context.md**: Manages local CLI configuration for switching between multiple servers.
*   **argocd_gpg.md**: System-wide GPG public key management (`add`, `list`, `get`, `rm`).
*   **argocd_login.md**: Authentication entry point supporting local users, SSO, and EKS flows.
*   **argocd_logout.md**: Invalidates local sessions for specific contexts.
*   **argocd_proj.md**: Core commands for defining project boundaries and restrictions.
*   **argocd_proj_role.md**: Root for managing project-specific RBAC roles and policies.
*   **argocd_proj_windows.md**: Management of Sync Windows (calendared maintenance windows).
*   **argocd_relogin.md**: Refreshes expired authentication tokens via CLI.
*   **argocd_repo.md**: Manage Git, Helm, or OCI repository connections and configurations.
*   **argocd_repocreds.md**: Manage credential templates that apply to multiple repositories.
*   **auto_sync.md**: Deep dive into automated synchronization policies, pruning, self-healing, and retry logic.
*   **best_practices.md**: Architectural advice on separating config from source code and ensuring manifest immutability.
*   **build-environment.md**: List of standard environment variables available to config management tools.
*   **ci_automation.md**: Guidelines for integrating CI pipelines with Argo CD.
*   **compare-options.md**: Configuration for ignoring extraneous resources during application comparison.
*   **diff-strategies.md**: Comparison of Legacy, Structured-Merge, and Server-Side Diff strategies.
*   **diffing.md**: System and Application level customizations for ignoring specific diff paths.
*   **directory.md**: Management of plain YAML/JSON manifest directories.
*   **environment-variables.md**: Global environment variables for configuring the `argocd` CLI.
*   **external-url.md**: Methods for adding clickable external links to the Argo CD UI.
*   **extra_info.md**: Adding custom key-value metadata to the application dashboard.
*   **gpg-verification.md**: Steps to enforce GnuPG signature verification on Git commits.
*   **helm.md**: Comprehensive guide on Helm integration, OCI charts, and values precedence.
*   **import.md**: Guidance for developers importing Argo CD Go packages into their own projects.
*   **jsonnet.md**: Documentation on using Jsonnet for manifest generation.
*   **kustomize.md**: Declarative setup for Kustomize-based applications.
*   **multiple_sources.md**: Combining multiple Git/Helm repositories into a single logical application.
*   **oci.md**: Detailed usage of OCI images and Helm charts as application sources.
*   **orphaned-resources.md**: Monitoring resources in a namespace not managed by Argo CD.
*   **parameters.md**: Mechanisms for overriding application parameters via CLI, UI, or config files.
*   **plugins.md**: Guide for writing and distributing custom sub-commands for the CLI.
*   **private-repositories.md**: Instructions for connecting private repositories using HTTPS, SSH, or GitHub Apps.
*   **projects.md**: Managing `AppProject` resources, RBAC roles, and inheritance.
*   **resource_tracking.md**: Technical details on how Argo CD tracks live resources using labels or annotations.
*   **scale_application_resources.md**: Instructions for scaling workloads directly from the UI.
*   **selective_sync.md**: Instructions for syncing a subset of resources within an application.
*   **skip_reconcile.md**: Pausing the reconciliation loop for specific applications via annotations.
*   **source-hydrator.md**: Documentation for the "Source Hydrator" (Rendered Manifest Pattern).
*   **status-badge.md**: Configuration for application health and sync status badges.
*   **subscriptions.md**: Configuration for event-based notifications to Slack, email, or webhooks.
*   **sync-kubectl.md**: Using standard `kubectl` commands to trigger Argo CD operations.
*   **sync-options.md**: Reference for advanced sync flags (Server-Side Apply, CreateNamespace, etc.).
*   **sync-waves.md**: Logic for Sync Phases (PreSync, PostSync) and Sync Waves.
*   **sync_windows.md**: Defining time-based windows to allow or block synchronization.
*   **tool_detection.md**: Logic used to automatically identify Helm, Kustomize, or Directory tools.
*   **tracking_strategies.md**: Differences between tracking Git HEAD, branches, tags, and versions.

## Code Changes That Would Require Documentation Updates
*   **CRD Schema Changes**: Modifications to `Application`, `AppProject`, or `ApplicationSet` specifications (e.g., new fields in `syncPolicy`).
*   **CLI Structure & Flags**: Adding/renaming subcommands or flags in the Cobra-based CLI; changing default flag values.
*   **New Annotations/Labels**: Introduction of new `argocd.argoproj.io/` or `notifications.argoproj.io/` metadata keys.
*   **Sync & Reconciliation Logic**: Changes to the application controller’s loop, sync waves, phases, retry logic, or sharding algorithms.
*   **Diffing Engine**: Modifications to how Server-Side Diffing or Structured-Merge Diffing is calculated or customized.
*   **Config Management Tooling**: Updating supported versions or default build flags for Helm, Kustomize, or Jsonnet.
*   **UI Features**: New dashboard elements, scaling modals, or repository connection wizards.
*   **Security & RBAC**: Changes to GitHub App support, OIDC handling, Casbin policy logic, or `can-i` command resources.
*   **Resource Tracking**: Changes to the default label (`app.kubernetes.io/instance`) or tracking annotation logic.
*   **Notification Engine**: Adding new default triggers or services in `argocd-notifications-cm`.

## Key Technical Concepts
*   **Sync Waves & Phases**: Orchestrating the order of resource application during a sync.
*   **Self-Healing & Pruning**: Automatic alignment of cluster state to Git and removal of extraneous resources.
*   **Server-Side Apply**: Utilizing Kubernetes native field management for large or shared resources.
*   **Resource Tracking**: The mechanism (Annotation vs. Label) used to map cluster objects to Applications.
*   **Multi-Source Apps**: Combining diverse repositories into a single logical deployment.
*   **Sync Windows**: Cron-based time constraints for automated or manual synchronization.
*   **App of Apps / ApplicationSets**: Patterns for bootstrapping clusters and automating multi-cluster app generation.
*   **Rendered Manifest Pattern (Hydration)**: Committing hydrated manifests back to Git for audit and visibility.
*   **Casbin RBAC**: The policy engine used for project-level permissions and isolation.
*   **Sharding**: The distribution of cluster management across multiple controller replicas.
*   **Resource Actions**: User-defined Lua scripts for performing operations like "restart" on custom resources.

## Related Components
*   **Argo CD Application Controller**: Manages reconciliation, sync operations, and sync windows.
*   **Argo CD Repo Server**: Handles Git/Helm cloning and manifest generation (Helm template, Kustomize build).
*   **Argo CD API Server**: Exposes the GRC/REST API for the CLI and UI.
*   **Argo CD Notifications Controller**: Processes triggers and sends alerts to external services.
*   **ApplicationSet Controller**: Automates Application CRD generation based on generators.
*   **argocd CLI**: The primary terminal interface for imperative and administrative tasks.
*   **Redis**: Used for caching manifest generation results, connection statuses, and data compression.
*   **Kubernetes API Server**: The ultimate target for resource application and state source for reconciliation.