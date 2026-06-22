# USER-GUIDE/COMMANDS Documentation Index

## Overview
This documentation provides a comprehensive reference for the Argo CD Command Line Interface (CLI), covering syntax, options, and usage for managing core entities such as Applications, ApplicationSets, Projects, Clusters, Repositories, and Accounts. It serves as both a functional guide for developers performing lifecycle operations and a configuration reference for administrators managing system security, performance, data integrity, and infrastructure troubleshooting.

## Files Summary
*   **argocd.md**: Global CLI options and high-level command list.
*   **argocd_account.md**: Base command for managing account settings, listing accounts, and checking permissions.
*   **argocd_account_bcrypt.md**: Utility to generate bcrypt hashes for passwords.
*   **argocd_account_can-i.md**: Checks if the current account has permissions for specific actions on resources.
*   **argocd_account_delete-token.md**: Revokes API tokens for a specified account or the current user.
*   **argocd_account_generate-token.md**: Generates authentication tokens for local accounts.
*   **argocd_account_get.md**: Retrieves specific account details.
*   **argocd_account_get-user-info.md**: Retrieves information about the currently logged-in user.
*   **argocd_account_list.md**: Lists all configured accounts.
*   **argocd_account_update-password.md**: Changes the password for the current or a specified account.
*   **argocd_admin.md**: Entry point for administrative commands requiring direct Kubernetes access.
*   **argocd_admin_app.md**: Advanced application management including diffing reconciliation results and spec generation.
*   **argocd_admin_app_diff-reconcile-results.md**: Admin utility to compare the output of two reconciliation cycles for debugging.
*   **argocd_admin_app_generate-spec.md**: Instructions for generating declarative application configurations.
*   **argocd_admin_app_get-reconcile-results.md**: Details on reconciling applications and exporting summaries to files.
*   **argocd_admin_cluster.md**: Base reference for administrative cluster configuration commands.
*   **argocd_admin_cluster_generate-spec.md**: Guides the generation of declarative configurations for clusters.
*   **argocd_admin_cluster_kubeconfig.md**: Generates, lists, or deletes kubeconfig entries for managed clusters.
*   **argocd_admin_cluster_namespaces.md**: Manages namespaced-mode for cluster management.
*   **argocd_admin_cluster_namespaces_enable-namespaced-mode.md**: Enables the namespaced-mode for specific clusters.
*   **argocd_admin_cluster_shards.md**: Explains how to view controller sharding distribution and resource responsibility.
*   **argocd_admin_cluster_stats.md**: Provides performance metrics and sharding information for clusters.
*   **argocd_admin_dashboard.md**: Launches the Argo CD Web UI locally via port-forwarding.
*   **argocd_admin_export.md**: Reference for backing up all Argo CD data to stdout or a file.
*   **argocd_admin_import.md**: Restores Argo CD state from a backup file.
*   **argocd_admin_initial-password.md**: Retrieves or resets the default admin password.
*   **argocd_admin_notifications.md**: Root command reference for managing notification settings.
*   **argocd_admin_notifications_template.md**: Manages notification templates.
*   **argocd_admin_notifications_template_get.md**: Retrieves details of a specific notification template.
*   **argocd_admin_notifications_template_notify.md**: Instructions for testing and sending manual notifications.
*   **argocd_admin_notifications_trigger.md**: Overview of notification trigger management commands.
*   **argocd_admin_notifications_trigger_get.md**: Retrieves details for a notification trigger.
*   **argocd_admin_notifications_trigger_run.md**: Debugs and executes notification triggers.
*   **argocd_admin_proj.md**: Parent command for administrative project management.
*   **argocd_admin_proj_generate-allow-list.md**: Utility to generate project allow-lists from Kubernetes ClusterRole files.
*   **argocd_admin_proj_generate-spec.md**: Instructions for generating declarative project configurations.
*   **argocd_admin_proj_update-role-policy.md**: Performs bulk updates on project role policies.
*   **argocd_admin_redis-initial-password.md**: Administrative tool to manage the initial Redis password.
*   **argocd_admin_repo_generate-spec.md**: Generates declarative YAML for repository configurations.
*   **argocd_admin_settings.md**: Root command for administrative settings validation and troubleshooting.
*   **argocd_admin_settings_rbac.md**: Tools for validating and testing complex RBAC configurations.
*   **argocd_admin_settings_rbac_can.md**: Troubleshooting tool to check specific RBAC permissions for a role or subject.
*   **argocd_admin_settings_rbac_validate.md**: Command to syntactically validate RBAC policy files or ConfigMaps.
*   **argocd_admin_settings_resource-overrides.md**: Root command for troubleshooting resource customizations.
*   **argocd_admin_settings_resource-overrides_health.md**: Assesses resource health using custom Lua scripts.
*   **argocd_admin_settings_resource-overrides_ignore-differences.md**: Verifies fields excluded from diffing via configuration.
*   **argocd_admin_settings_resource-overrides_ignore-resource-updates.md**: Renders fields excluded from updates based on customization settings.
*   **argocd_admin_settings_resource-overrides_list-actions.md**: Lists available Lua-based resource actions.
*   **argocd_admin_settings_resource-overrides_run-action.md**: Executes custom resource actions (e.g., restarts) via Lua scripts.
*   **argocd_admin_settings_validate.md**: Checks configuration files for errors.
*   **argocd_app.md**: Primary command for application lifecycle management.
*   **argocd_app_actions.md**: Root command for managing resource-level actions within applications.
*   **argocd_app_actions_list.md**: Lists available custom actions for a specific resource.
*   **argocd_app_actions_run.md**: Executes a specific action on application resources.
*   **argocd_app_add-source.md**: Adds a source repository to a multi-source Argo CD application.
*   **argocd_app_confirm-deletion.md**: Confirms the manual pruning or deletion of resources.
*   **argocd_app_create.md**: Creates a new application from various sources (Git, Helm, Kustomize, etc.).
*   **argocd_app_delete.md**: Reference for deleting applications with support for cascading and selectors.
*   **argocd_app_delete-resource.md**: Removes a specific Kubernetes resource from an application.
*   **argocd_app_diff.md**: Details on performing diffs between target state (Git) and live state (K8s).
*   **argocd_app_edit.md**: Opens the application configuration in a text editor.
*   **argocd_app_get.md**: Provides comprehensive details, status, and parameters of an application.
*   **argocd_app_get-resource.md**: Retrieves live Kubernetes manifests for specific resources.
*   **argocd_app_history.md**: Displays the deployment history and previous revisions.
*   **argocd_app_list.md**: Filters and lists applications by project, repo, or labels.
*   **argocd_app_logs.md**: Streams or retrieves logs from pods associated with an application.
*   **argocd_app_manifests.md**: Prints rendered Kubernetes manifests for a specific revision.
*   **argocd_app_patch.md**: Updates application specifications using JSON or Merge patches.
*   **argocd_app_patch-resource.md**: Applies a patch to a specific resource within an app.
*   **argocd_app_remove-source.md**: Removes a specific source from a multi-source application.
*   **argocd_app_resources.md**: Lists all resources managed by an application, including orphaned resources.
*   **argocd_app_rollback.md**: Reverts an application to a previous state using its history ID.
*   **argocd_app_set.md**: Detailed reference for modifying application parameters and sync policies.
*   **argocd_app_sync.md**: Manually triggers a reconciliation/synchronization of an application.
*   **argocd_app_terminate-op.md**: Cancels a currently running application operation.
*   **argocd_app_unset.md**: Removes specific parameter overrides (Kustomize, Helm, etc.).
*   **argocd_app_wait.md**: Blocks until an application reaches a specific state (e.g., Healthy, Synced).
*   **argocd_appset.md**: Base command for managing ApplicationSet controllers.
*   **argocd_appset_create.md**: Reference for creating ApplicationSets from files or URLs.
*   **argocd_appset_delete.md**: Deletes one or more ApplicationSets.
*   **argocd_appset_generate.md**: Previews applications generated by an ApplicationSet template.
*   **argocd_appset_get.md**: Retrieves detailed configuration and status for an ApplicationSet.
*   **argocd_appset_list.md**: Lists managed ApplicationSets with filtering options.
*   **argocd_cert.md**: Root command for managing repository security credentials.
*   **argocd_cert_add-ssh.md**: Manages SSH known host entries for private repository access.
*   **argocd_cert_add-tls.md**: Adds TLS certificates for secure repository connections.
*   **argocd_cert_list.md**: Lists all configured repository certificates and SSH hosts.
*   **argocd_cert_rm.md**: Deletes repository certificates from the Argo CD repo server.
*   **argocd_cluster.md**: General management of target cluster credentials.
*   **argocd_cluster_add.md**: Guides adding new Kubernetes clusters to Argo CD management.
*   **argocd_cluster_get.md**: Shows details for a registered cluster.
*   **argocd_cluster_list.md**: Lists all clusters currently configured in Argo CD.
*   **argocd_cluster_rm.md**: Removes cluster credentials and stops management of that target.
*   **argocd_cluster_rotate-auth.md**: Updates credentials used for cluster communication.
*   **argocd_cluster_set.md**: Updates metadata and permitted namespaces for clusters.
*   **argocd_completion.md**: Generates shell completion scripts (Bash, Zsh, Fish).
*   **argocd_configure.md**: Manages local CLI settings and interactive prompts.
*   **argocd_context.md**: Manages CLI local contexts for switching between multiple servers.
*   **argocd_gpg.md**: Reference for managing GPG keys used for commit signature verification.
*   **argocd_gpg_add.md**: Adds a GPG public key.
*   **argocd_gpg_get.md**: Retrieves a specific GPG key detail.
*   **argocd_gpg_list.md**: Lists all GPG keys.
*   **argocd_gpg_rm.md**: Removes a GPG key.
*   **argocd_login.md**: Details the authentication process (User/Pass, SSO, or Core).
*   **argocd_logout.md**: Terminates the session for the active server context.
*   **argocd_proj.md**: Root command for general AppProject management.
*   **argocd_proj_add-destination.md**: Adds permitted target clusters and namespaces to a project.
*   **argocd_proj_add-destination-service-account.md**: Configures default service accounts for project destinations.
*   **argocd_proj_add-orphaned-ignore.md**: Excludes specific resources from orphaned resource monitoring.
*   **argocd_proj_add-signature-key.md**: Links GPG keys to projects for enforced signature verification.
*   **argocd_proj_add-source.md**: Authorizes a new repository URL as a valid source for a project.
*   **argocd_proj_add-source-namespace.md**: Manages namespaces that can host Argo CD applications in a project.
*   **argocd_proj_allow-cluster-resource.md**: Manages the allow-list for cluster-scoped resources.
*   **argocd_proj_allow-namespace-resource.md**: Manages resource-level allow-lists for namespaces.
*   **argocd_proj_create.md**: Reference for creating new AppProjects.
*   **argocd_proj_delete.md**: Reference for deleting AppProjects.
*   **argocd_proj_deny-cluster-resource.md**: Manages cluster-scoped resource restrictions.
*   **argocd_proj_deny-namespace-resource.md**: Manages resource-level deny-lists for namespaces.
*   **argocd_proj_edit.md**: Opens project configuration for interactive editing.
*   **argocd_proj_get.md**: Displays detailed configuration and status for a project.
*   **argocd_proj_list.md**: Lists all AppProjects.
*   **argocd_proj_remove-destination.md**: Removes a target cluster and namespace from a project.
*   **argocd_proj_remove-destination-service-account.md**: Removes specific destination-scoped service accounts.
*   **argocd_proj_remove-orphaned-ignore.md**: Manages the ignore-list for orphaned resource monitoring.
*   **argocd_proj_remove-signature-key.md**: Removes GPG keys from a project.
*   **argocd_proj_remove-source.md**: Removes a repository URL from the allowed sources list.
*   **argocd_proj_remove-source-namespace.md**: Removes authorized namespaces for hosting apps.
*   **argocd_proj_role.md**: Base command for project role management.
*   **argocd_proj_role_add-group.md**: Maps OIDC/SAML groups to specific project roles.
*   **argocd_proj_role_add-policy.md**: Manages RBAC policy rules within a project role.
*   **argocd_proj_role_create.md**: Defines new roles within an AppProject.
*   **argocd_proj_role_create-token.md**: Generates JWT tokens for project-specific roles.
*   **argocd_proj_role_delete.md**: Deletes a project role.
*   **argocd_proj_role_delete-token.md**: Revokes a JWT token associated with a project role.
*   **argocd_proj_role_get.md**: Shows details of a project role, including policies and active tokens.
*   **argocd_proj_role_list.md**: Lists all roles associated with a specific project.
*   **argocd_proj_role_list-tokens.md**: Lists active JWT tokens for a project role.
*   **argocd_proj_role_remove-group.md**: Removes OIDC group claims from a role.
*   **argocd_proj_role_remove-policy.md**: Removes specific permission policies from project roles.
*   **argocd_proj_set.md**: Configures project parameters like allowed resources and destinations.
*   **argocd_proj_windows.md**: Manages maintenance windows for project synchronization.
*   **argocd_proj_windows_add.md**: Defines sync windows (Allow/Deny) for deployment timing.
*   **argocd_proj_windows_delete.md**: Deletes scheduled sync windows from a project.
*   **argocd_proj_windows_disable-manual-sync.md**: Removes manual override capability from a sync window.
*   **argocd_proj_windows_enable-manual-sync.md**: Overrides sync windows for a one-time manual synchronization.
*   **argocd_proj_windows_list.md**: Displays configured sync windows for a project.
*   **argocd_proj_windows_update.md**: Modifies existing sync window schedules or constraints.
*   **argocd_relogin.md**: Refreshes expired authentication tokens.
*   **argocd_repo.md**: Root command for repository connection management.
*   **argocd_repo_add.md**: Reference for connecting Git, Helm, or OCI repositories.
*   **argocd_repo_get.md**: Retrieves configuration for a specific repository.
*   **argocd_repo_list.md**: Lists and refreshes connection status of configured repositories.
*   **argocd_repo_rm.md**: Removes credentials and metadata for source repositories.
*   **argocd_repocreds.md**: Manages credential templates for bulk repository authentication.
*   **argocd_repocreds_add.md**: Adds repository credential templates (SSH, HTTPS, etc.).
*   **argocd_repocreds_list.md**: Lists all repository credential templates.
*   **argocd_repocreds_rm.md**: Removes repository credential templates.
*   **argocd_version.md**: Prints client and server version information.

## Code Changes That Would Require Documentation Updates
*   **CLI Structure**: Adding, renaming, or removing any subcommand or changing the global `argocd` options.
*   **Flag Modifications**: Adding new flags, changing default values (e.g., timeout or retry limits), or deprecating existing flags.
*   **API & Schema**: Changes to Go structs returned by the API or updates to CRDs (Application, AppProject, ApplicationSet) affecting declarative configuration.
*   **Authentication & Security**: Modifying the login flow, adding support for new authentication providers (OIDC/SSO), or changing GPG verification logic.
*   **RBAC & Scoping**: Refactoring Casbin policy formats, introducing new resource types/actions, or changing namespace scoping features.
*   **Sync & Reconciliation**: Implementing new sync policies, self-healing behaviors, retry strategies, or sharding algorithms.
*   **Configuration Maps**: Updates to the schema or processing of `argocd-cm`, `argocd-secret`, or `argocd-cmd-params-cm`.
*   **Manifest Generation**: Changes to how manifest generation tools (Helm, Kustomize, Jsonnet) are invoked or the addition of new tool versions.
*   **System Notifications**: Updates to the notification engine, including how templates and triggers are evaluated or rendered.

## Key Technical Concepts
*   **AppProject**: The primary security boundary used to group applications and restrict sources/destinations.
*   **RBAC (Role-Based Access Control)**: CSV-based policy management for granular access control.
*   **Sync Policy**: Automated vs. Manual reconciliation, including auto-prune, self-heal, and retry strategies.
*   **Sync Windows**: Time-based constraints (cron-based) that allow or block deployment activities.
*   **Sharding**: Mechanism for distributing managed clusters across multiple controller replicas for scalability.
*   **Core Mode**: Bypassing the Argo CD API server to communicate directly with the Kubernetes API.
*   **JWT Tokens**: Role-based tokens used for programmatic and automated access.
*   **Declarative Config**: The generation and management of Argo CD entities via YAML/JSON manifests.
*   **Resource Overrides**: Custom Lua scripts used to modify health assessments or execute resource actions.
*   **GPG Signature Verification**: Ensuring commit integrity by verifying keys before synchronization.
*   **Orphaned Resources**: Monitoring and managing live resources not defined in the source repository.
*   **Context Management**: Handling multiple server connections and user identities via the local CLI configuration.

## Related Components
*   **Argo CD API Server**: The central gateway for all CLI and UI requests.
*   **Application Controller**: The component responsible for monitoring state, sharding, and enforcing policies.
*   **Repo Server**: Handles manifest generation (hydration) and repository credential management.
*   **Notifications Controller**: Processes triggers and sends alerts based on system events.
*   **Redis**: High-performance cache for repository state, sharding data, and manifest generation results.
*   **Kubernetes API Server**: The final destination for all synchronized resource manifests.
*   **DEX / OIDC**: External identity providers used for user authentication and group mapping.
*   **ConfigMaps & Secrets**: Persistent storage for system settings (`argocd-cm`, `argocd-secret`).