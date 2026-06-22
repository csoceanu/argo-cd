# USER-GUIDE/COMMANDS Documentation Index

## Overview
This documentation area provides a comprehensive reference for the Argo CD Command Line Interface (CLI), covering application lifecycle management, project security configuration, and cluster administration. It serves as the primary resource for understanding command syntax, available flags, and practical usage examples for both standard user operations and advanced administrative tasks required to maintain an Argo CD instance.

## Files Summary

* **argocd.md**: Root command reference detailing global flags like `--server`, `--auth-token`, and `--core`.
* **argocd_account.md**: Main entry point for managing local account settings, password updates, and permission checks.
* **argocd_account_bcrypt.md**: Utility for generating bcrypt-hashed passwords for use in Argo CD configuration maps.
* **argocd_account_can-i.md**: Validates whether the current account has RBAC permissions to perform a specific action on a resource.
* **argocd_account_delete-token.md**: Command reference for revoking API tokens associated with a specific Argo CD account.
* **argocd_account_generate-token.md**: Issues a new JWT authentication token for a local account, optionally with an expiration.
* **argocd_account_get.md**: Retrieves detailed configuration and status for a specific Argo CD account.
* **argocd_account_get-user-info.md**: Retrieves and displays metadata about the currently authenticated user session.
* **argocd_account_list.md**: Displays a list of all local and SSO-synced accounts known to the Argo CD instance.
* **argocd_account_update-password.md**: Changes the authentication password for the current user or a specified local account.
* **argocd_admin.md**: Parent command for system-level administrative utilities that require direct Kubernetes access.
* **argocd_admin_app.md**: Acts as the parent command for administrative application tasks, including reconciliation diffs and spec generation.
* **argocd_admin_app_diff-reconcile-results.md**: Administrative command for comparing two reconciliation outputs to debug state discrepancies.
* **argocd_admin_app_generate-spec.md**: Provides instructions for generating declarative YAML or JSON configurations for Applications from existing sources.
* **argocd_admin_app_get-reconcile-results.md**: Explains how to reconcile all applications and export the results to a summary file for troubleshooting.
* **argocd_admin_cluster.md**: Serves as the parent reference for administrative cluster configuration and management tasks.
* **argocd_admin_cluster_generate-spec.md**: Details the process for generating declarative configuration manifests for registering new clusters.
* **argocd_admin_cluster_kubeconfig.md**: Reference for generating, listing, or deleting kubeconfig files for clusters managed by Argo CD.
* **argocd_admin_cluster_namespaces.md**: Displays the mapping of managed namespaces across all registered clusters.
* **argocd_admin_cluster_namespaces_disable-namespaced-mode.md**: Disables restricted namespace management for a set of clusters.
* **argocd_admin_cluster_namespaces_enable-namespaced-mode.md**: Enables "Namespaced Mode" for specific clusters to restrict Argo CD to specific namespaces.
* **argocd_admin_cluster_shards.md**: Describes how to view information about controller sharding and the distribution of resources across shards.
* **argocd_admin_cluster_stats.md**: Provides detailed metrics regarding cluster resource usage and controller sharding distribution.
* **argocd_admin_dashboard.md**: Starts a local proxy and opens the Argo CD Web UI in the default browser.
* **argocd_admin_export.md**: Guide for exporting all Argo CD configuration data to stdout or a file for backup and migration.
* **argocd_admin_import.md**: Facilitates the restoration of Argo CD state by importing data from a backup file or stdin.
* **argocd_admin_initial-password.md**: Manages the retrieval and reset of the default administrative password for a new installation.
* **argocd_admin_notifications.md**: Parent reference for commands used to manage and configure the Argo CD notifications system.
* **argocd_admin_notifications_template.md**: Parent command for notification template management, used to define how alerts are formatted.
* **argocd_admin_notifications_template_get.md**: Retrieves the specific definition and contents of a notification template.
* **argocd_admin_notifications_template_notify.md**: Explains how to test and send notifications using specific templates to verify recipient configuration.
* **argocd_admin_notifications_trigger.md**: Parent reference for managing notification triggers that determine when alerts are sent.
* **argocd_admin_notifications_trigger_get.md**: Displays the configuration details for specific notification triggers.
* **argocd_admin_notifications_trigger_run.md**: Allows administrators to manually test and evaluate a notification trigger against a resource.
* **argocd_admin_proj.md**: High-level administrative entry point for project configuration and allow-list generation.
* **argocd_admin_proj_generate-allow-list.md**: Guides the generation of project resource allow-lists based on existing Kubernetes ClusterRole files.
* **argocd_admin_proj_generate-spec.md**: Explains how to generate declarative manifests for AppProject resources.
* **argocd_admin_proj_update-role-policy.md**: Enables bulk updates to project role policies across multiple projects using glob patterns.
* **argocd_admin_redis-initial-password.md**: Provides the command to ensure a Redis password exists for secure internal communication.
* **argocd_admin_repo.md**: Parent reference for administrative repository configuration and spec generation.
* **argocd_admin_repo_generate-spec.md**: Generates declarative configuration YAML for adding repositories with various auth methods.
* **argocd_admin_settings.md**: Parent reference for system-wide settings validation and troubleshooting commands.
* **argocd_admin_settings_rbac.md**: Provides tools for validating and testing complex RBAC configurations.
* **argocd_admin_settings_rbac_can.md**: Details how to test RBAC permissions for a specific role or subject against a policy.
* **argocd_admin_settings_rbac_validate.md**: Explains the process for checking the syntax and structural correctness of RBAC policy files.
* **argocd_admin_settings_resource-overrides.md**: Parent reference for troubleshooting and testing resource-specific customizations.
* **argocd_admin_settings_resource-overrides_health.md**: Describes how to test custom Lua health scripts against specific resource manifests.
* **argocd_admin_settings_resource-overrides_ignore-differences.md**: Administrative tool for testing fields excluded from diffing based on customization settings.
* **argocd_admin_settings_resource-overrides_ignore-resource-updates.md**: Renders fields that are currently excluded from resource updates based on system settings.
* **argocd_admin_settings_resource-overrides_list-actions.md**: Lists available custom Lua-based resource actions defined in the global configuration.
* **argocd_admin_settings_resource-overrides_run-action.md**: Guides the execution and testing of custom resource actions using Lua scripts.
* **argocd_admin_settings_validate.md**: Performs structural and logical validation of the Argo CD global configuration files (`argocd-cm.yaml`).
* **argocd_app.md**: Primary command for all application-level operations and lifecycle management.
* **argocd_app_actions.md**: Parent reference for managing resource-level actions within an application.
* **argocd_app_actions_list.md**: Lists the available actions that can be performed on specific resources within an application.
* **argocd_app_actions_run.md**: Explains how to execute a specific action on resources matching defined filters.
* **argocd_app_add-source.md**: Details the process of adding multiple source repositories to an existing Argo CD application.
* **argocd_app_confirm-deletion.md**: Manual confirmation step for deleting resources during a cascaded application deletion.
* **argocd_app_create.md**: Provisions a new application with parameters for Helm, Kustomize, Jsonnet, or directory sources.
* **argocd_app_delete.md**: Explains how to remove applications and their associated resources.
* **argocd_app_delete-resource.md**: Removes a specific Kubernetes resource that was deployed as part of an application.
* **argocd_app_diff.md**: Provides instructions for comparing the live state of an application against its desired state in Git.
* **argocd_app_edit.md**: Provides an interactive way to modify an application's YAML configuration using a local text editor.
* **argocd_app_get.md**: Provides a comprehensive view of an application's status, configuration, and health.
* **argocd_app_get-resource.md**: Explains how to retrieve the live manifest of a specific resource managed by an application.
* **argocd_app_history.md**: Displays the deployment history of an application, including previous sync operations.
* **argocd_app_list.md**: Details how to list and filter applications based on projects, clusters, or labels.
* **argocd_app_logs.md**: Provides the command for streaming or viewing logs from pods associated with an application.
* **argocd_app_manifests.md**: Guide for viewing the rendered Kubernetes manifests for an application from live state or git.
* **argocd_app_patch.md**: Explains how to apply JSON or Merge patches to an application's specification.
* **argocd_app_patch-resource.md**: Applies a Kubernetes patch to a specific resource managed by an application.
* **argocd_app_remove-source.md**: Removes a specific source definition from an application that utilizes multiple sources.
* **argocd_app_resources.md**: Lists all Kubernetes resources associated with an application, including identifying orphaned resources.
* **argocd_app_rollback.md**: Reverts an application to a previous state recorded in its deployment history.
* **argocd_app_set.md**: Detailed reference for modifying application parameters, sync policies, and tool-specific configurations.
* **argocd_app_sync.md**: Triggers a synchronization operation to move an application to the target state defined in Git.
* **argocd_app_terminate-op.md**: Cancels a currently running background operation, such as a hung sync.
* **argocd_app_unset.md**: Details how to remove specific parameters or overrides from an application's configuration.
* **argocd_app_wait.md**: Blocking CLI command that waits until an application reaches a desired health or sync state.
* **argocd_appset.md**: Parent command for managing ApplicationSet resources and their lifecycle.
* **argocd_appset_create.md**: Explains how to create or update ApplicationSets from files or URLs.
* **argocd_appset_delete.md**: Deletes ApplicationSet resources, optionally triggering a cascaded deletion of generated apps.
* **argocd_appset_generate.md**: Previews the application manifests that would be produced by an ApplicationSet template.
* **argocd_appset_get.md**: Provides detailed information and parameters for a specific ApplicationSet.
* **argocd_appset_list.md**: Lists all configured ApplicationSets with filtering options.
* **argocd_cert.md**: Parent reference for managing TLS certificates and SSH known host entries.
* **argocd_cert_add-ssh.md**: Explains how to add SSH public host keys for secure repository access.
* **argocd_cert_add-tls.md**: Guides the addition of TLS certificates for repository servers.
* **argocd_cert_list.md**: Displays all configured certificates and known hosts in the Argo CD store.
* **argocd_cert_rm.md**: Reference for removing repository certificates or SSH known hosts entries.
* **argocd_cluster.md**: General management reference for clusters, covering adding, listing, and updating credentials.
* **argocd_cluster_add.md**: Details the process for registering a new Kubernetes cluster with Argo CD.
* **argocd_cluster_get.md**: Displays detailed connection and status information for a registered cluster.
* **argocd_cluster_list.md**: Lists all clusters currently managed by the Argo CD instance.
* **argocd_cluster_rm.md**: Command for deleting cluster credentials and removing a cluster from management.
* **argocd_cluster_rotate-auth.md**: Provides a mechanism to refresh or rotate the authentication credentials for a specific remote cluster.
* **argocd_cluster_set.md**: Explains how to update metadata, labels, and allowed namespaces for an existing cluster.
* **argocd_completion.md**: Outputs auto-completion scripts for various shell environments (Bash, Zsh, Fish).
* **argocd_configure.md**: Manual for managing local CLI settings, such as enabling/disabling interactive prompts.
* **argocd_context.md**: Provides commands for managing and switching between different Argo CD server contexts.
* **argocd_gpg.md**: Parent reference for managing GnuPG keys used for commit signature verification.
* **argocd_gpg_add.md**: Uploads a new public GPG key to the Argo CD server.
* **argocd_gpg_get.md**: Retrieves the detailed content of a specific public GPG key by its ID.
* **argocd_gpg_list.md**: Lists all public GPG keys registered on the server.
* **argocd_gpg_rm.md**: Deletes a public GPG key from the server's keyring.
* **argocd_login.md**: Details the authentication process for the CLI via password, token, or SSO.
* **argocd_logout.md**: Simple reference for terminating the active Argo CD session.
* **argocd_proj.md**: Parent reference for managing Argo CD AppProjects.
* **argocd_proj_add-destination.md**: Explains how to add permitted destination clusters and namespaces to a project.
* **argocd_proj_add-destination-service-account.md**: Details adding default service accounts for specific project destinations.
* **argocd_proj_add-orphaned-ignore.md**: Adds specific resource types to a project's ignore list for orphaned resource monitoring.
* **argocd_proj_add-signature-key.md**: Guides the addition of GPG keys to a project for mandatory signature verification.
* **argocd_proj_add-source.md**: Instructions for whitelisting a new repository URL as a valid source for a project.
* **argocd_proj_add-source-namespace.md**: Permitting specific Kubernetes namespaces to host source manifests for a project.
* **argocd_proj_allow-cluster-resource.md**: Explains how to allow specific cluster-scoped resources within a project.
* **argocd_proj_allow-namespace-resource.md**: Modifies a project's allow list for namespaced Kubernetes API resources.
* **argocd_proj_create.md**: Provisions a new AppProject with defined security boundaries and resource constraints.
* **argocd_proj_delete.md**: Removes an existing AppProject resource from the Argo CD system.
* **argocd_proj_deny-cluster-resource.md**: Restricts the creation of specific cluster-level resources within a project's scope.
* **argocd_proj_deny-namespace-resource.md**: Adds a namespaced resource type to a project's deny list.
* **argocd_proj_edit.md**: Provides an interactive method for editing project configurations.
* **argocd_proj_get.md**: Reference for retrieving detailed information about a project, including roles, policies, and sync windows.
* **argocd_proj_list.md**: Displays a summary of all existing AppProject resources.
* **argocd_proj_remove-destination.md**: Instructions for removing a destination cluster and namespace pairing from a project.
* **argocd_proj_remove-destination-service-account.md**: Deletes a specific destination service account mapping from a project.
* **argocd_proj_remove-orphaned-ignore.md**: Details how to remove resources from the orphaned ignore list in a project.
* **argocd_proj_remove-signature-key.md**: Removes a GPG key ID from a project's list of required commit signatures.
* **argocd_proj_remove-source.md**: Removes a repository URL from the allowed list of sources in a project.
* **argocd_proj_remove-source-namespace.md**: Revokes a namespace's permission to serve as a source for applications in a project.
* **argocd_proj_role.md**: Parent command for managing project-level roles and security tokens.
* **argocd_proj_role_add-group.md**: Explains how to bind OIDC group claims to project-specific roles.
* **argocd_proj_role_add-policy.md**: Appends a new RBAC policy rule to a project role.
* **argocd_proj_role_create.md**: Details the creation of new roles within a project for fine-grained access control.
* **argocd_proj_role_create-token.md**: Guides the generation of JWT tokens for automation against project roles.
* **argocd_proj_role_delete.md**: Deletes a specific RBAC role and its associated policies from an AppProject.
* **argocd_proj_role_delete-token.md**: Invalidates a specific JWT token associated with a project role.
* **argocd_proj_role_get.md**: Displays the details, policies, and tokens associated with a project role.
* **argocd_proj_role_list.md**: Lists all roles defined within a specific project.
* **argocd_proj_role_list-tokens.md**: Lists the metadata of all active JWT tokens issued for a specific project role.
* **argocd_proj_role_remove-group.md**: Removes an OIDC group mapping from a project's RBAC role.
* **argocd_proj_role_remove-policy.md**: Instructions for deleting specific RBAC policy entries from a project role.
* **argocd_proj_set.md**: Updates the parameters of an AppProject, such as descriptions or destination clusters.
* **argocd_proj_windows.md**: Parent command for managing project synchronization windows.
* **argocd_proj_windows_add.md**: Explains how to define allow or deny sync windows for a project.
* **argocd_proj_windows_delete.md**: Guide for removing defined sync windows from a project using their unique IDs.
* **argocd_proj_windows_disable-manual-sync.md**: Details how to restrict manual overrides during defined sync windows.
* **argocd_proj_windows_enable-manual-sync.md**: Overrides a deny sync window to allow manual synchronization of applications.
* **argocd_proj_windows_list.md**: Lists all configured synchronization windows for a specific AppProject.
* **argocd_proj_windows_update.md**: Guides the modification of existing sync window schedules or durations.
* **argocd_relogin.md**: Forces a re-authentication flow to refresh expired session tokens.
* **argocd_repo.md**: Parent reference for managing repository connections.
* **argocd_repo_add.md**: Provides instructions for adding Git, Helm, or OCI repositories with various auth methods.
* **argocd_repo_get.md**: Retrieves the connection settings and status for a registered repository.
* **argocd_repo_list.md**: Lists all repositories connected to Argo CD and their connection status.
* **argocd_repo_rm.md**: Reference for removing connection credentials for one or more source repositories.
* **argocd_repocreds.md**: Parent reference for managing repository credential templates.
* **argocd_repocreds_add.md**: Registers new repository credentials supporting SSH, HTTPS, GitHub Apps, and GCP.
* **argocd_repocreds_list.md**: Lists all repository credential templates used for bulk authentication.
* **argocd_repocreds_rm.md**: Details how to remove stored repository credential templates.
* **argocd_version.md**: Displays the version information for both the Argo CD client and server.

## Code Changes That Would Require Documentation Updates

* **CLI Syntax & Framework**: Modifications to `cobra` command definitions, help templates, global flags, subcommands, or default values in `cmd/argocd`.
* **API & Schema Changes**: New gRPC/REST endpoints or changes to underlying CRD Go structs for `Application`, `AppProject`, or `ApplicationSet`.
* **Authentication & Auth Providers**: Adding support for new repository types (OCI), authentication protocols (SSO/OIDC), or changes to JWT token lifecycle/account management.
* **Resource Customization & Engines**: Updates to Lua integration for health checks and resource actions, or alterations to the reconciliation/diffing algorithms.
* **Infrastructure & Sharding**: Changes to multi-cluster logic, namespace management modes, or controller sharding distribution algorithms.
* **Notification Engine**: Changes to how triggers and templates are processed, validated, or delivered in the `notifications-controller`.
* **Manifest Generation**: Changes in how the `repo-server` processes Helm values, Kustomize parameters, or Jsonnet variables.
* **Output Formats**: Updates to the serialization logic for `json`, `yaml`, `wide`, `tree`, or administrative export data models.

## Key Technical Concepts

* **Application Lifecycle**: Sync, Rollback, History, Pruning, Self-Healing, Cascaded Deletion, and Multi-Source Application support.
* **Security & Multi-tenancy**: AppProject boundaries, RBAC Policies (`p, sub, res, act, obj`), OIDC Group Claims, JWT Tokens, and GPG Signature Verification.
* **Infrastructure Management**: Cluster Registration, Controller Sharding, Repository Credentials, and Namespaced Mode.
* **Synchronization Control**: Sync Windows (allow/deny schedules), Sync Policies (automated/manual), and Sync Options.
* **Reconciliation & Health**: Server-Side Diff, `ignoreDifferences` logic, Lua-based Health Checks, and Custom Resource Actions.
* **Automation**: ApplicationSet templates and Notification Triggers/Templates.
* **Administration**: ConfigMap/Secret Validation, Data Export/Import, and RBAC testing.

## Related Components

* **argocd-server**: The API server handling CLI requests, authentication, and RBAC.
* **argocd-application-controller**: The engine managing state reconciliation, health checks, and sync windows.
* **argocd-repo-server**: Component responsible for manifest generation (Helm/Kustomize) and Git/OCI interactions.
* **argocd-notifications-controller**: Processes triggers and templates to send alerts to external recipients.
* **argocd-applicationset-controller**: Manages the generation of applications from defined templates.
* **Redis**: Used for caching manifest data, repository status, and reconciliation results.
* **Kubernetes API Server**: The target for resource deployment and the primary endpoint for `--core` mode operations.
* **Argo CD ConfigMaps**: `argocd-cm` (settings), `argocd-rbac-cm` (policy), and `argocd-cmd-params-cm` (system parameters).