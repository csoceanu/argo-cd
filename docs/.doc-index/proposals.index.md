# PROPOSALS Documentation Index

## Overview
This documentation area contains the design specifications, architectural blueprints, and feature proposals for the Argo CD ecosystem. It covers fundamental changes to the core controller, multi-tenancy models, UI/CLI enhancements, and integrations with external tools like OCI registries and configuration management plugins.

## Files Summary

*   **proposals/multiple-sources-for-applications-ui.md**: Outlines the UI/UX changes required to support applications that pull from multiple repositories, including new "Sources" tabs and edit capabilities.
*   **proposals/deletion-strategy-progressive-sync.md**: Proposes a new functionality for ApplicationSet ProgressiveSync to allow configurable deletion orders, such as `AllAtOnce` or `Reverse`.
*   **proposals/config-management-plugin-v2.md**: Details the transition of Config Management Plugins to a sidecar architecture to improve security, installation, and tool discovery.
*   **proposals/server-side-pagination.md**: Describes the implementation of server-side pagination for the Application List and Watch APIs to improve UI performance with large datasets.
*   **proposals/004-scalability-benchmarking.md**: Establishes a repeatable procedure and set of metrics for benchmarking Argo CD performance across various cluster topologies.
*   **proposals/argocd-cli-pluin.md**: Proposes a plugin mechanism for the `argocd` CLI, similar to `kubectl` plugins, to keep the core binary lean.
*   **proposals/respect-rbac-for-resource-exclusions.md**: Details how the Argo CD controller can be configured to respect Kubernetes RBAC when deciding which resources to monitor.
*   **proposals/applicationset-plugin-generator.md**: Introduces a "plugin" generator for ApplicationSets that retrieves parameters via external RPC calls.
*   **proposals/rebalancing-clusters-across-shards-dynamically.md**: Proposes moving the application controller from a StatefulSet to a Deployment to enable dynamic horizontal scaling and cluster rebalancing.
*   **proposals/feature-bounties.md**: Defines an experimental program for offering monetary rewards for significant community-contributed features.
*   **proposals/notifications-API.md**: Describes new API endpoints to expose notification triggers and services directly to the Argo CD UI.
*   **proposals/2022-07-13-appset-progressive-rollout-strategy.md**: Proposes `RollingUpdate` and `RollingSync` strategies for controlled application updates within an ApplicationSet.
*   **proposals/headless-argocd.md**: Introduces a distribution of Argo CD that operates without a central API server or UI, relying on local CLI/UI components.
*   **proposals/manifest-hydrator.md**: Proposes a first-class "rendered manifest" pattern where Argo CD hydrates dry manifests and pushes them back to a Git branch.
*   **proposals/002-ui-extensions.md**: Defines a framework for loading external JavaScript/Lua components to provide custom resource visualizations and actions in the UI.
*   **proposals/parameterized-config-management-plugins.md**: Enhances CMPs to "announce" parameters to the UI, providing a native-like experience for custom tools.
*   **proposals/server-side-apply.md**: Outlines the integration of Kubernetes Server-Side Apply (SSA) into the sync process to improve conflict resolution and CRD handling.
*   **proposals/native-oci-support.md**: Proposes adding native support for OCI registries as a source for manifests (Helm and non-Helm content).
*   **proposals/project-scoped-repository-enhancements.md**: Allows multiple repository credentials for the same URL by scoping them to specific AppProjects.
*   **proposals/project-repos-and-clusters.md**: Enables self-service onboarding by allowing developers to register clusters and repositories within a specific project scope.
*   **proposals/decouple-application-sync-user-using-impersonation.md**: Details the use of Kubernetes User Impersonation to execute sync operations with restricted ServiceAccount privileges.
*   **proposals/multiple-sources-for-applications.md**: The core proposal for the `sources` field in the Application CRD to support combining multiple Git/Helm sources.
*   **proposals/003-applications-outside-argocd-namespace.md**: Proposes allowing Application CRDs to be reconciled from namespaces other than the default `argocd` namespace.
*   **proposals/application-name-identifier.md**: Proposes a tracking method using annotations instead of labels to support application names exceeding 63 characters.
*   **proposals/backend-support-appset.md**: Details the API server changes required to support CRUD operations for ApplicationSets via the CLI and Web UI.
*   **proposals/001-proposal-template.md**: The standardized template used for creating all Argo CD enhancement proposals.
*   **proposals/deep-links.md**: Introduces configurable links in the UI that allow users to jump to external monitoring tools like Splunk or Datadog.
*   **proposals/resource-deletion-with-approval.md**: Proposes a manual approval step for resource pruning or deletion during the sync process.
*   **proposals/sync-timeout.md**: Adds configuration for operation timeouts and termination settings based on resource health states.
*   **proposals/proxy-extensions.md**: Extends the UI extension framework to support reverse-proxying requests to external backend services.

## Code Changes That Would Require Documentation Updates

*   **CRD Modifications**: Any changes to `Application`, `ApplicationSet`, or `AppProject` schemas (e.g., adding `sources`, `sourceHydrator`, or `destinationServiceAccounts`).
*   **API Protocol Changes**: Updates to Protobuf definitions for the Application, Project, or Notification services (e.g., adding pagination fields or deep link info).
*   **Controller Reconciliation Logic**: Changes to how the Application Controller handles sync phases, sharding logic, or resource tracking (label vs. annotation).
*   **RBAC and Security**: Modifications to the Argo CD RBAC evaluator or the introduction of new resource types (e.g., `extensions`).
*   **CLI Structure**: Addition of new subcommands (e.g., `argocd appset`, `argocd admin dashboard`) or global flags (e.g., `--headless`, `--as`).
*   **Git/OCI Interaction**: Changes to the repo-server's ability to fetch content, handle credentials, or parse OCI media types.
*   **UI Architecture**: Updates to how JavaScript extensions are loaded, how tabs are rendered, or how pagination is handled in the frontend.
*   **Sharding and Scaling**: Altering the deployment model from StatefulSet to Deployment or changing how `ARGOCD_CONTROLLER_REPLICAS` is consumed.

## Key Technical Concepts

*   **Multi-Source Applications**: Using a list of sources in a single Application CRD.
*   **Progressive Rollout**: Controlled updates to applications in an ApplicationSet using steps and match expressions.
*   **Config Management Plugin (CMP) v2**: Sidecar-based manifest generation.
*   **Server-Side Apply (SSA)**: Delegating merge logic to the Kubernetes API server.
*   **User Impersonation**: Syncing resources as a specific ServiceAccount.
*   **Manifest Hydration**: The process of rendering dry manifests (Helm/Kustomize) and committing the output back to Git.
*   **Resource Tracking**: The mechanism (label or annotation) used to link live cluster resources to an Application.
*   **Deep Links**: Templated external URLs displayed in the UI based on resource metadata.
*   **Headless Mode**: Running Argo CD components without the central API server.

## Related Components

*   **Application Controller**: The core logic responsible for reconciliation and sync.
*   **ApplicationSet Controller**: The generator-based automation for creating multiple applications.
*   **Argo CD API Server**: The GRPC/REST gateway for the UI and CLI.
*   **Repo-Server**: The component responsible for cloning Git/OCI repos and rendering manifests.
*   **GitOps Engine**: The underlying library shared with other projects for cluster synchronization.
*   **Dex/SSO**: The authentication integration for multi-tenant access.
*   **Redis**: The caching layer for manifest generation and cluster state.