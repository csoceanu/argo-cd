# PROPOSALS Documentation Index

## Overview
This documentation area contains architectural blueprints and feature proposals for the Argo CD ecosystem, including core Application behavior, ApplicationSet enhancements, UI/CLI extensions, and scalability improvements. Its purpose is to record the design decisions, technical constraints, and implementation strategies for major changes before they are merged into the codebase, serving as a historical record and a roadmap for developers.

## Files Summary

*   **001-proposal-template.md**: A standardized template for submitting new enhancement proposals to the Argo CD project.
*   **002-ui-extensions.md**: Proposes a mechanism to load external Javascript/Lua scripts from Git to provide custom visualizations and actions for CRDs in the UI.
*   **003-applications-outside-argocd-namespace.md**: Outlines the design for allowing Application CRs to exist in namespaces other than the control plane's namespace.
*   **004-scalability-benchmarking.md**: Describes repeatable procedures and metrics for measuring the performance limits of Argo CD at scale.
*   **2022-07-13-appset-progressive-rollout-strategy.md**: Introduces `RollingUpdate` and `RollingSync` strategies for the ApplicationSet controller.
*   **application-name-identifier.md**: Proposes shifting resource tracking from labels to annotations to support longer application names and avoid label-copying conflicts.
*   **applicationset-plugin-generator.md**: Defines a new ApplicationSet generator that fetches parameters from external services via RPC calls.
*   **argocd-cli-pluin.md**: Details the implementation of a plugin system for the `argocd` CLI, similar to `kubectl` plugins.
*   **backend-support-appset.md**: Proposes API server endpoints and CLI/UI support for creating and managing ApplicationSets with proper RBAC.
*   **config-management-plugin-v2.md**: Enhances Config Management Plugins (CMPs) by running them as sidecars in the repo-server pod for better isolation.
*   **decouple-application-sync-user-using-impersonation.md**: Design for performing application syncs using Kubernetes user/group impersonation to enforce least-privilege security.
*   **deep-links.md**: Provides a configuration mechanism to add external links to the Argo CD UI for projects, applications, and resources.
*   **deletion-strategy-progressive-sync.md**: Adds configurable deletion orders (e.g., `Reverse`) for applications managed by ApplicationSet Progressive Sync.
*   **feature-bounties.md**: Establishes an experimental program for offering monetary rewards for specific community-contributed features.
*   **feature-bounties/hide-annotations.md**: A specific bounty proposal for adding a configuration to hide certain annotations from the Web UI.
*   **headless-argocd.md**: Proposes a "Headless" distribution of Argo CD that operates without an API server or UI, relying on direct Kubernetes access.
*   **manifest-hydrator.md**: Introduces first-class support for manifest hydration (rendering Helm/Kustomize to plain YAML) and pushing results back to Git.
*   **manifest-hydrator/README.md**: User-facing documentation for configuring and using the manifest hydration feature.
*   **manifest-hydrator/commit-server/README.md**: Technical specification for the internal gRPC service that handles pushing hydrated manifests to Git.
*   **multiple-sources-for-applications.md**: The core proposal for adding a `sources` field to Application CRs to support multiple repositories in one app.
*   **multiple-sources-for-applications-ui.md**: Specific UI design changes to support the viewing and editing of multi-source applications.
*   **native-oci-support.md**: Outlines the retrieval of manifests stored as artifacts in OCI registries as a first-class source type.
*   **notifications-API.md**: Adds API endpoints to list available notification triggers and services to improve the UI subscription experience.
*   **parameterized-config-management-plugins.md**: Enables CMPs to "announce" parameters to the UI and CLI, providing an experience similar to Helm or Kustomize.
*   **project-repos-and-clusters.md**: Design for project-scoped repositories and clusters to allow self-service onboarding for developers.
*   **project-scoped-repository-enhancements.md**: Allows multiple teams to use identical repository URLs with different credentials by scoping them to AppProjects.
*   **proxy-extensions.md**: Introduces a reverse-proxy in the API server to allow UI extensions to communicate with dedicated backend services.
*   **rebalancing-clusters-across-shards-dynamically.md**: Proposes moving the controller to a stateless Deployment and using ConfigMap heartbeats for dynamic sharding.
*   **resource-deletion-with-approval.md**: Adds a safety mechanism requiring manual user approval in the UI/CLI before pruning or deleting specific resources.
*   **respect-rbac-for-resource-exclusions.md**: Configures the controller to monitor only resources for which its Service Account has explicit RBAC read permissions.
*   **server-side-apply.md**: Proposes leveraging Kubernetes Server-Side Apply (SSA) for syncing resources to resolve large CRD and admission controller issues.
*   **server-side-pagination.md**: Introduces pagination and filtering to the Application List and Watch APIs to improve UI responsiveness.
*   **sync-timeout.md**: Adds configurable timeouts and termination settings to synchronization operations.

## Code Changes That Would Require Documentation Updates

*   **CRD Schema Changes**: Any modifications to `Application`, `ApplicationSet`, `AppProject`, or `ArgoCDExtension` specs (e.g., adding `sources`, `sourceHydrator`, `strategy`, or `impersonation` fields).
*   **API/Protobuf Modifications**: Updates to the gRPC service definitions for Application, Project, or Notification services (especially new endpoints for pagination or links).
*   **Controller Logic Updates**: Changes to how sharding is calculated, how sync operations are timed out, or how RBAC is respected during resource discovery.
*   **UI Component Overhauls**: Significant changes to the Application Details page, the Summary tab, or the addition of new "Sources" or "Extensions" tabs.
*   **CLI Subcommand Additions**: New commands under `argocd appset`, `argocd admin`, or the introduction of a new plugin architecture for the CLI.
*   **Environment Variable Changes**: New global configuration flags for the Application Controller (e.g., enabling SSA or Impersonation).
*   **Sidecar Configurations**: Changes to the `argocd-repo-server` deployment regarding CMP v2 socket communication or shared volumes.

## Key Technical Concepts

*   **Tracking Methods**: `label`, `annotation`, `annotation+label` (used to identify managed resources).
*   **Progressive Sync**: `RollingUpdate`, `RollingSync`, `Reverse` deletion, `AllAtOnce`.
*   **CMP v2**: Sidecar architecture, `ConfigManagementPlugin` manifest, `argocd-cmp-server`.
*   **Sharding Logic**: Dynamic rebalancing, ConfigMap heartbeats, heartbeat interval.
*   **Manifest Hydration**: `sourceHydrator`, `drySource`, `syncSource`, `hydrateTo`.
*   **Security/Auth**: `impersonation`, `SelfSubjectAccessReview`, project-scoped credentials.
*   **API Optimization**: Server-side pagination, `offset`, `limit`, `minName`, `maxName`, Application Stats.
*   **Extensions**: `ArgoCDExtension` CRD, UI extension reverse proxy, Lua-based actions/health checks.

## Related Components

*   **Application Controller**: The primary logic for reconciling and syncing applications.
*   **ApplicationSet Controller**: The generator-based system for managing groups of applications.
*   **Argo CD API Server**: The central gateway for UI, CLI, and external integrations.
*   **Argo CD Repo Server**: The component responsible for manifest generation (Helm, Kustomize, CMPs).
*   **GitOps Engine**: The underlying library handling Kubernetes resource diffing and application.
*   **Argo CD CLI**: The command-line interface for administrative and developer operations.