# USER-GUIDE/COMMANDS Documentation Index

## Overview
This documentation provides a comprehensive command-line interface (CLI) reference for Argo CD, detailing the syntax and options for managing applications, projects, clusters, and administrative settings. It serves as the primary technical guide for developers and platform engineers to interact with the Argo CD API and underlying Kubernetes resources via the `argocd` command-line tool.

## Files Summary
*   **argocd.md**: The root command reference listing global flags and all top-level sub-commands for controlling the Argo CD server.
*   **argocd_account.md**: Provides an overview of commands used to manage local account settings, including listing accounts and updating passwords.
*   **argocd_account_bcrypt.md**: Documentation for generating bcrypt hashes for passwords, useful for manual configuration of local users.
*   **argocd_account_can-i.md**: A utility to check specific RBAC permissions (actions on resources) for the currently logged-in user.
*   **argocd_account_delete-token.md**: Describes how to invalidate and delete authentication tokens for specific Argo CD accounts.
*   **argocd_account_generate-token.md**: Explains how to generate authentication tokens for local accounts with optional expiration times.
*   **argocd_account_get.md**: Describes how to retrieve specific details for a given local account or the current user.
*   **argocd_account_get-user-info.md**: Instructions for retrieving information about the currently authenticated user session.
*   **argocd_account_list.md**: Lists all accounts configured within the Argo CD instance.
*   **argocd_account_update-password.md**: Details the process for changing the password of the current user or a specified account.
*   **argocd_admin.md**: Parent reference for administrative commands that typically require direct Kubernetes cluster access.
*   **argocd_admin_app.md**: High-level command for managing application configurations from an administrative level.
*   **argocd_admin_app_diff-reconcile-results.md**: A troubleshooting command used to compare the results of two different reconciliation cycles.
*   **argocd_admin_app_generate-spec.md**: Generates declarative Application specifications for various configuration management tools like Helm, Kustomize, and Jsonnet.
*   **argocd_admin_app_get-reconcile-results.md**: Triggers a global reconciliation and saves a summary of the results for all applications to a local file.
*   **argocd_admin_cluster.md**: Root command for administrative management and diagnostic reporting for cluster connections.
*   **argocd_admin_cluster_generate-spec.md**: Generates the declarative configuration required to add a cluster to Argo CD.
*   **argocd_admin_cluster_kubeconfig.md**: Explains how to generate or manage local kubeconfig files for clusters managed by Argo CD.
*   **argocd_admin_cluster_namespaces.md**: Displays information about which namespaces Argo CD is currently managing across various clusters.
*   **argocd_admin_cluster_namespaces_disable-namespaced-mode.md**: Explains how to disable the "namespaced mode" for specific clusters.
*   **argocd_admin_cluster_namespaces_enable-namespaced-mode.md**: Explains how to enable "namespaced mode" to restrict Argo CD to specific namespaces on a cluster.
*   **argocd_admin_cluster_shards.md**: Provides technical insights into how clusters are distributed across different controller shards.
*   **argocd_admin_cluster_stats.md**: Provides cluster-level statistics and helps determine/inferred sharding distributions.
*   **argocd_admin_dashboard.md**: Instructions for launching the Argo CD Web UI on a local port using port-forwarding.
*   **argocd_admin_export.md**: Covers the administrative utility for exporting all Argo CD data (applications, settings, etc.) to stdout or a file for backup or migration.
*   **argocd_admin_import.md**: Guidance on importing Argo CD configuration data (apps, projects, settings) from a backup file or stdin.
*   **argocd_admin_initial-password.md**: Explains how to retrieve the auto-generated initial admin password from Kubernetes secrets.
*   **argocd_admin_notifications.md**: Root command for notification management, allowing for template and trigger configuration troubleshooting.
*   **argocd_admin_notifications_template.md**: Overview of commands for managing notification templates used by the notifications controller.
*   **argocd_admin_notifications_template_get.md**: Shows how to retrieve and view specific notification template definitions.
*   **argocd_admin_notifications_template_notify.md**: A testing utility to trigger and send notifications using specific templates to verify integration with services like Slack.
*   **argocd_admin_notifications_trigger.md**: Provides sub-commands to inspect and test the conditions that trigger notifications.
*   **argocd_admin_notifications_trigger_get.md**: Lists configured notification triggers and their conditions.
*   **argocd_admin_notifications_trigger_run.md**: Allows administrators to test/evaluate a notification trigger against a specific resource.
*   **argocd_admin_proj.md**: Provides a high-level entry point for administrative project management commands, including policy updates and spec generation.
*   **argocd_admin_proj_generate-allow-list.md**: A utility to generate a project resource allow-list by parsing an existing Kubernetes ClusterRole YAML file.
*   **argocd_admin_proj_generate-spec.md**: Generates the declarative YAML/JSON configuration for an AppProject, including source and destination restrictions.
*   **argocd_admin_proj_update-role-policy.md**: Enables bulk updates of project role policies across multiple projects using globs.
*   **argocd_admin_redis-initial-password.md**: Ensures a Redis password exists in the cluster, creating a new secret if it is missing.
*   **argocd_admin_repo.md**: Administrative root command for managing the global repository configuration of the Argo CD instance.
*   **argocd_admin_repo_generate-spec.md**: Generates declarative Kubernetes manifests for repository credentials.
*   **argocd_admin_settings.md**: Root command for administrative troubleshooting of global settings, including RBAC and resource customizations.
*   **argocd_admin_settings_rbac.md**: Parent command for validating and testing the global RBAC configuration.
*   **argocd_admin_settings_rbac_can.md**: A diagnostic tool to verify if a specific subject or role has permission to perform an action on a resource under the current RBAC policy.
*   **argocd_admin_settings_rbac_validate.md**: Validates the syntax and structural correctness of an RBAC policy file or ConfigMap.
*   **argocd_admin_settings_resource-overrides.md**: Root command for troubleshooting resource-specific customizations like health checks and actions.
*   **argocd_admin_settings_resource-overrides_health.md**: Tests custom Lua health scripts against resource manifests to troubleshoot status reporting.
*   **argocd_admin_settings_resource-overrides_ignore-differences.md**: Explains how to render and troubleshoot fields excluded from diffing based on the `ignoreDifferences` configuration in the `argocd-cm` ConfigMap.
*   **argocd_admin_settings_resource-overrides_ignore-resource-updates.md**: Renders fields that are currently ignored during resource updates based on global customizations.
*   **argocd_admin_settings_resource-overrides_list-actions.md**: Lists available custom Lua actions for specific resources based on system settings.
*   **argocd_admin_settings_resource-overrides_run-action.md**: Manually executes a custom resource action using the internal Lua engine for testing purposes.
*   **argocd_admin_settings_validate.md**: Validates the contents of `argocd-cm` and `argocd-secret` for configuration errors.
*   **argocd_app.md**: Central command for application lifecycle management, listing all available sub-commands.
*   **argocd_app_actions.md**: The root command for discovering and invoking administrative actions on application resources.
*   **argocd_app_actions_list.md**: Discovers which specific actions (e.g., restart, resume) are available for a given resource in an application.
*   **argocd_app_actions_run.md**: Executes resource-specific actions (such as restarting a Deployment) on resources managed within an application.
*   **argocd_app_add-source.md**: Adds an additional source repository to an application, supporting the multi-source applications feature.
*   **argocd_app_confirm-deletion.md**: Used to confirm the deletion of resources when a prune operation is pending.
*   **argocd_app_create.md**: Comprehensive guide for creating new applications from various sources (Helm, Git, Kustomize, etc.).
*   **argocd_app_delete.md**: Removes applications from Argo CD, with options for cascaded deletion of cluster resources and label-based selection.
*   **argocd_app_delete-resource.md**: Explains how to delete a specific Kubernetes resource from a managed application.
*   **argocd_app_diff.md**: Performs a comparison between the target state in Git and the live state in the cluster, highlighting configuration drifts.
*   **argocd_app_edit.md**: Opens the application's YAML definition in the default editor for live modification.
*   **argocd_app_get.md**: Provides a comprehensive overview of an application's health, sync status, and configuration parameters.
*   **argocd_app_get-resource.md**: Fetches the live manifests of specific Kubernetes resources within an application, allowing for field-level filtering.
*   **argocd_app_history.md**: Displays the deployment history of an application, providing IDs for rollback operations.
*   **argocd_app_list.md**: Lists managed applications with extensive filtering options by project, repository, cluster, or labels.
*   **argocd_app_logs.md**: Streams or retrieves logs from pods associated with a specific application, supporting container and time-based filtering.
*   **argocd_app_manifests.md**: Covers retrieving and printing the rendered manifests of an application from live state, Git, or local paths.
*   **argocd_app_patch.md**: Updates an application specification using JSON or Merge patch types directly from the CLI.
*   **argocd_app_patch-resource.md**: Documentation for applying JSON or merge patches to specific resources within an application.
*   **argocd_app_remove-source.md**: Explains how to remove a single source from a multiple-sources application.
*   **argocd_app_resources.md**: Lists all resources belonging to an application, with an optional tree view and orphaned resource filtering.
*   **argocd_app_rollback.md**: Details the process of reverting an application to a previous state found in its history.
*   **argocd_app_set.md**: Describes how to modify application parameters, including sync policies, source details, and tool-specific settings.
*   **argocd_app_sync.md**: The primary command for triggering the reconciliation of an application with its target state.
*   **argocd_app_terminate-op.md**: Allows users to stop a currently running application operation, such as a long-running sync.
*   **argocd_app_unset.md**: Removes specific parameter overrides (like Kustomize images or Helm values) from an application's configuration.
*   **argocd_app_wait.md**: Details how to block the CLI until an application reaches a specified state, such as Synced, Healthy, or Deleted.
*   **argocd_appset.md**: Overview for managing ApplicationSet resources which automate application creation.
*   **argocd_appset_create.md**: Facilitates the creation of ApplicationSet resources from files or URLs, including support for dry-runs and upserts.
*   **argocd_appset_delete.md**: Instructions for removing one or more ApplicationSets and their generated applications.
*   **argocd_appset_generate.md**: A utility to preview the applications that an ApplicationSet would generate without applying them.
*   **argocd_appset_get.md**: Retrieves the detailed status and specification of a specific ApplicationSet.
*   **argocd_appset_list.md**: Displays all ApplicationSets currently managed by the controller.
*   **argocd_cert.md**: The primary command for managing TLS certificates and SSH known hosts entries required for secure repository communication.
*   **argocd_cert_add-ssh.md**: Adds SSH known host entries for Git servers to the internal store via batch processing.
*   **argocd_cert_add-tls.md**: Adds a TLS certificate for a specific repository server to ensure secure HTTPS connections.
*   **argocd_cert_list.md**: Lists all stored repository certificates, categorized by type (HTTPS or SSH) and hostname.
*   **argocd_cert_rm.md**: Explains how to remove repository certificates (HTTPS) or SSH known hosts entries from the Argo CD configuration.
*   **argocd_cluster.md**: Serves as the primary reference for cluster management, including adding, listing, and updating target cluster configurations.
*   **argocd_cluster_add.md**: Adds a new cluster to Argo CD by importing settings from a local kubeconfig context.
*   **argocd_cluster_get.md**: Retrieves detailed connection status and metadata for a registered cluster.
*   **argocd_cluster_list.md**: Lists all Kubernetes clusters configured in Argo CD, providing details on their connection status and server URLs.
*   **argocd_cluster_rm.md**: Describes how to remove cluster credentials and stop Argo CD from managing a specific target cluster.
*   **argocd_cluster_rotate-auth.md**: Instructions for updating/rotating the authentication credentials used to access a managed cluster.
*   **argocd_cluster_set.md**: Updates the metadata, labels, and managed namespaces for an existing cluster connection.
*   **argocd_completion.md**: Documentation for generating shell completion scripts for Bash, Zsh, and Fish.
*   **argocd_configure.md**: Details commands for managing local CLI configuration, such as enabling or disabling interactive prompts.
*   **argocd_context.md**: Manages the local CLI configuration, allowing users to switch between different Argo CD server contexts.
*   **argocd_gpg.md**: The root command for managing GPG public keys used by the server to verify the signatures of Git commits.
*   **argocd_gpg_add.md**: Guide for adding GPG public keys to the server for commit signature verification.
*   **argocd_gpg_get.md**: Retrieves a specific GPG public key's details from the server's keyring.
*   **argocd_gpg_list.md**: Lists all GPG public keys currently used for signature verification.
*   **argocd_gpg_rm.md**: Explains how to remove a GPG key from the Argo CD keyring.
*   **argocd_login.md**: Handles user authentication with the Argo CD server, supporting passwords, SSO, and direct Kubernetes access.
*   **argocd_logout.md**: Provides instructions for logging out of the current Argo CD server context and clearing local session data.
*   **argocd_proj.md**: The root command for project lifecycle management, including creation, listing, and deletion.
*   **argocd_proj_add-destination.md**: Adds a permitted destination cluster and namespace to an AppProject's allow-list.
*   **argocd_proj_add-destination-service-account.md**: Configures a specific service account to be used by default for a project's destination cluster and namespace.
*   **argocd_proj_add-orphaned-ignore.md**: Adds specific resource types to a project's "ignore" list for orphaned resource monitoring.
*   **argocd_proj_add-signature-key.md**: Adds a GPG key ID to a project's configuration to enforce commit signature verification.
*   **argocd_proj_add-source.md**: Describes how to authorize new source repository URLs for use within a specific project.
*   **argocd_proj_add-source-namespace.md**: Authorizes specific namespaces to host applications belonging to a project.
*   **argocd_proj_allow-cluster-resource.md**: Permits a specific cluster-scoped API resource (Group/Kind) to be managed within a project.
*   **argocd_proj_allow-namespace-resource.md**: Manages the allow-list for Kubernetes resource kinds that can be deployed within a project.
*   **argocd_proj_create.md**: Details for creating new AppProject resources with defined restrictions.
*   **argocd_proj_delete.md**: Instructions for removing a project and its associated metadata.
*   **argocd_proj_deny-cluster-resource.md**: Adds cluster-scoped resource types to the deny-list for a project.
*   **argocd_proj_deny-namespace-resource.md**: Adds namespace-scoped resource types to the deny-list for a project.
*   **argocd_proj_edit.md**: Launches an interactive editor to modify the configuration of an existing AppProject.
*   **argocd_proj_get.md**: Explains how to retrieve and display detailed configuration and status information for a specific project.
*   **argocd_proj_list.md**: Lists all defined AppProjects.
*   **argocd_proj_remove-destination.md**: Details how to remove target cluster and namespace destinations from a project's permitted list.
*   **argocd_proj_remove-destination-service-account.md**: Removes specific service account mappings for a project destination.
*   **argocd_proj_remove-orphaned-ignore.md**: Removes a previously ignored resource from the orphaned resources monitoring exclusion list of a project.
*   **argocd_proj_remove-signature-key.md**: Removes a GPG key ID from a project's signature verification list.
*   **argocd_proj_remove-source.md**: Provides the command reference for removing authorized source repository URLs from a specific Argo CD project.
*   **argocd_proj_remove-source-namespace.md**: Revokes a namespace's authorization to host a project's applications.
*   **argocd_proj_role.md**: Parent command for managing roles and permissions within a specific project.
*   **argocd_proj_role_add-group.md**: Maps an external OIDC or SAML group claim to an internal project role for RBAC.
*   **argocd_proj_role_add-policy.md**: Explains how to add fine-grained RBAC policies to a project role.
*   **argocd_proj_role_create.md**: Creates a new RBAC role within a project to facilitate granular access control.
*   **argocd_proj_role_create-token.md**: Generates a new JWT authentication token for a project role with optional expiration.
*   **argocd_proj_role_delete.md**: Deletes a specific role from a project.
*   **argocd_proj_role_delete-token.md**: Deletes a specific JWT token associated with a project role.
*   **argocd_proj_role_get.md**: Retrieves the detailed configuration of a project role, including policies and tokens.
*   **argocd_proj_role_list.md**: Lists all roles defined within a specific project.
*   **argocd_proj_role_list-tokens.md**: Lists all active JWT tokens for a specific project role.
*   **argocd_proj_role_remove-group.md**: Removes an OIDC group claim mapping from a project role.
*   **argocd_proj_role_remove-policy.md**: Provides instructions for removing specific RBAC policy entries from a project.
*   **argocd_proj_set.md**: General command for updating various project parameters like descriptions, sources, and destinations.
*   **argocd_proj_windows.md**: Parent command for managing Sync Windows (schedules that allow or block syncs).
*   **argocd_proj_windows_add.md**: Defines sync windows for a project to either allow or deny resource synchronization.
*   **argocd_proj_windows_delete.md**: Explains how to delete scheduled sync windows within a project.
*   **argocd_proj_windows_disable-manual-sync.md**: Disables the ability to perform manual syncs when a deny-sync window is active.
*   **argocd_proj_windows_enable-manual-sync.md**: Allows manual sync overrides for a scheduled sync window.
*   **argocd_proj_windows_list.md**: Lists all configured sync windows for a specific project.
*   **argocd_proj_windows_update.md**: Modifies the properties of an existing sync window, such as its schedule or duration.
*   **argocd_relogin.md**: Explains how to refresh an expired session token using existing credentials or SSO.
*   **argocd_repo.md**: The root command for managing the list of connected Git, Helm, and OCI repositories.
*   **argocd_repo_add.md**: Configures connection parameters for repositories using SSH, HTTPS, or GitHub Apps.
*   **argocd_repo_get.md**: Retrieves connection details and status for a registered Git or Helm repository.
*   **argocd_repo_list.md**: Lists all repository connections, with an option to force a hard refresh.
*   **argocd_repo_rm.md**: Details the process for removing configured repository connection parameters and credentials.
*   **argocd_repocreds.md**: Manages repository credential templates that provide shared authentication for multiple repositories.
*   **argocd_repocreds_add.md**: Guide for adding credential templates that apply to multiple repositories.
*   **argocd_repocreds_list.md**: Lists all configured repository credential templates.
*   **argocd_repocreds_rm.md**: Deletes repository credential templates from the Argo CD configuration.
*   **argocd_version.md**: Prints the version information for both the Argo CD client and server.

## Code Changes That Would Require Documentation Updates
*   **CLI Framework & Structure**: Any changes to the underlying Cobra/Pflag implementation, global flag inheritance, renaming sub-commands, or changing default flag values.
*   **API/Struct & CRD Extensions**: Adding new fields to `ApplicationSpec`, `AppProjectSpec`, or `Cluster` objects requires corresponding flag updates in `create`, `set`, and `generate-spec` commands.
*   **Auth & Security Logic**: New repository authentication methods (e.g., cloud IAM roles), modifications to session management, or changes to JWT token generation and invalidation.
*   **RBAC Policy & Logic**: Changes to internal Casbin policy evaluation, the list of supported resources (e.g., extensions), or available actions in the `can-i` logic.
*   **Resource Customization & Lua Engine**: Updates to the Lua engine for health checks or new override types (e.g., new field-ignore categories) in the `resource-overrides` commands.
*   **Reconciliation & Diffing Engine**: Modifying how the controller calculates resource differences, reconciliation results, or `ignoreDifferences` logic.
*   **Configuration Management**: Changes to how Argo CD handles Helm, Kustomize, or Jsonnet parameters, including Config Management Plugin (CMP) logic.
*   **Infrastructure & Sharding**: Modifying cluster distribution algorithms or introducing new sharding methods.
*   **Administrative Tooling**: Updates to the export/import logic, namespaced mode features, or how `kubeconfig` files are generated.
*   **Notification System**: Introducing new trigger conditions, template variables, or notification controller behaviors.

## Key Technical Concepts
*   **Resources**: Application, ApplicationSet, AppProject, Cluster, Repository, Role, Sync Window.
*   **Auth/Security**: JWT, GPG Signature Verification, SSO, RBAC, OIDC/SAML Group Claims, TLS/SSH Certificates, Contexts.
*   **Config Management**: Helm, Kustomize, Jsonnet, OCI, Config Management Plugins (CMP), Application Parameter Overrides.
*   **States & Sync**: Target State, Live State, Diff, Reconciliation, Health Status (Healthy, Degraded, etc.), Sync Policy (Automated vs. Manual, Self-healing, Auto-pruning).
*   **Infrastructure**: Controller Shards, Redis Caching, Repo Server, API Server, Kube-Context.
*   **Advanced Monitoring**: Orphaned Resources, IgnoreDifferences, Multi-Source Applications.

## Related Components
*   **Argo CD API Server**: Primary backend interface for the CLI.
*   **Argo CD Application Controller**: Responsible for sharding, health checks, reconciliation, and sync operations.
*   **Argo CD Repo Server**: Handles manifest generation and Git/Helm/OCI interactions.
*   **Argo CD Notifications Controller**: Processes notification triggers and templates.
*   **Redis**: Used for caching connection status, manifest data, and session information.
*   **Dex**: Handles external authentication (SSO/OIDC).
*   **Kubernetes API Server**: Direct target for cluster state and administrative `--core` commands.
*   **Argo CD ConfigMap (argocd-cm)**: Global configuration source for settings and resource overrides.