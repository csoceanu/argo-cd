# PROPOSALS Documentation Index

## Overview
This documentation area contains architectural and feature proposals for the Argo CD ecosystem, encompassing core Application and ApplicationSet controllers, the API server, and user interfaces. Its purpose is to define the design specifications for multi-tenancy, scalability enhancements, extensibility through plugins, and advanced GitOps workflow strategies like manifest hydration and progressive rollouts.

## Files Summary
*   **multiple-sources-for-applications-ui.md**: Details UI enhancements to support the "Multiple Sources" feature, including a new "Sources" tab and parameter editing.
*   **deletion-strategy-progressive-sync.md**: Proposes `AllAtOnce` and `Reverse` deletion orders for Applications managed by ApplicationSet progressive sync.
*   **config-management-plugin-v2.md**: Outlines the transition to sidecar-based Config Management Plugins (CMP v2) for improved isolation and discovery.
*   **server-side-pagination.md**: Proposes adding `offset`, `limit`, and name-based cursors to Application List and Watch APIs to improve UI/CLI performance.
*   **004-scalability-benchmarking.md**: Defines a framework and metrics (e.g., sync time, queue depth) for benchmarking Argo CD performance in large-scale environments.
*   **argocd-cli-pluin.md**: Introduces a `kubectl`-style plugin system for the `argocd` CLI based on `$PATH` binary discovery.
*   **respect-rbac-for-resource-exclusions.md**: Proposes a controller setting (`resource.respectRBAC`) to automatically exclude resources the service account cannot read.
*   **applicationset-plugin-generator.md**: Introduces a `plugin` generator for ApplicationSets that fetches parameters via external RPC calls.
*   **rebalancing-clusters-across-shards-dynamically.md**: Proposes moving the Application Controller to a Deployment and using a ConfigMap-based heartbeat for dynamic sharding.
*   **feature-bounties.md**: Estapes a process for the Argo Project to offer and claim monetary bounties for feature implementations.
*   **notifications-API.md**: Proposes API endpoints to list notification triggers and services to enable UI-based subscriptions.
*   **2022-07-13-appset-progressive-rollout-strategy.md**: Defines `RollingUpdate` and `RollingSync` strategies for ApplicationSets to control update waves.
*   **headless-argocd.md**: Outlines "Headless" mode to run Argo CD without a centralized API/UI, relying instead on local CLI-driven dashboards.
*   **manifest-hydrator.md**: Proposes a first-class "rendered manifest" pattern where Argo CD hydrates "dry" sources and pushes them back to git.
*   **002-ui-extensions.md**: Introduces the `ArgoCDExtension` CRD and a framework for custom JS visualizations and Lua health/actions.
*   **parameterized-config-management-plugins.md**: Enhances CMPs to "announce" and consume structured parameters, mirroring the native Helm/Kustomize UI experience.
*   **server-side-apply.md**: Details support for Kubernetes Server-Side Apply (SSA) to improve conflict management and support large CRDs.
*   **native-oci-support.md**: Proposes sourcing any manifest type (not just Helm) directly from OCI registries as a first-class repository type.
*   **project-scoped-repository-enhancements.md**: Allows multiple AppProjects to share the same Repository URL using distinct credentials.
*   **project-repos-and-clusters.md**: Enables self-service onboarding by allowing project-scoped repositories and clusters stored as secrets.
*   **decouple-application-sync-user-using-impersonation.md**: Uses Kubernetes Impersonation to run syncs as a specific ServiceAccount defined in the AppProject.
*   **multiple-sources-for-applications.md**: The core proposal for the `sources` field in the Application spec, allowing apps to combine multiple git/helm sources.
*   **003-applications-outside-argocd-namespace.md**: Proposes allowing Application CRs to exist in namespaces other than the control plane namespace.
*   **application-name-identifier.md**: Introduces an annotation-based tracking method to bypass the 63-character Kubernetes label limit for app names.
*   **backend-support-appset.md**: Adds API Server endpoints and RBAC logic to support CRUD operations for ApplicationSets via CLI and UI.
*   **001-proposal-template.md**: The standard boilerplate for creating new enhancement proposals.
*   **deep-links.md**: Proposes a configuration in `argocd-cm` to display templated links to external tools (Splunk, Datadog) in the UI.
*   **resource-deletion-with-approval.md**: Introduces `Prune=confirm` and `Delete=confirm` sync options requiring manual UI/CLI approval.
*   **sync-timeout.md**: Proposes `syncPolicy.terminate` settings to enforce global sync timeouts or per-resource health timeouts.
*   **proxy-extensions.md**: Enables UI extensions to proxy requests to external backend services through the Argo CD API Server.

## Code Changes That Would Require Documentation Updates

### API and CRD Changes
*   Changes to `ApplicationSpec` or `ApplicationSetSpec` (specifically adding/removing fields like `sources`, `sourceHydrator`, `strategy`, or `syncPolicy`).
*   Modifications to the `AppProject` CRD, particularly regarding `destinations`, `sourceNamespaces`, or impersonation settings.
*   Introduction of new CRDs (e.g., `ArgoCDExtension`).
*   Updates to Protobuf definitions for `ApplicationService`, `ProjectService`, or the new `NotificationService`.

### Controller Logic
*   Changes to the reconciliation loop for Applications or ApplicationSets (e.g., sharding logic, rollout wave calculation, or deletion order).
*   Modifications to how Argo CD tracks resources (label vs. annotation tracking methods).
*   Updates to the Git/Helm/OCI fetch logic in the Repo Server.
*   Changes to the RBAC enforcement logic within the controller (especially regarding resource exclusions or impersonation).

### CLI and UI
*   Adding new subcommands to the `argocd` CLI (e.g., `argocd admin dashboard`, `argocd appset ...`, or CLI plugin discovery).
*   Changes to UI navigation or the introduction of new tabs (e.g., "Sources" tab, Extension tabs, or Notification subscription panels).
*   Updates to the `argocd-cm` or `argocd-rbac-cm` configuration schemas.

### Sync and Hydration
*   Modifications to the manifest generation pipeline (CMPs, hydration, or OCI artifact extraction).
*   Changes to sync options or synchronization hooks (e.g., adding `confirm` options or `timeout` logic).
*   Changes to how Server-Side Apply is invoked or how conflicts are reported.

## Key Technical Concepts
*   **Tracking Methods**: `label`, `annotation`, `annotation+label`.
*   **Sync Options**: `Prune=confirm`, `Delete=confirm`, `ServerSideApply=true`.
*   **Sharding**: `ARGOCD_CONTROLLER_REPLICAS`, dynamic rebalancing, heartbeat via ConfigMap.
*   **CMPs**: Sidecar architecture, `argocd-cmp-server`, `plugin.yaml`, dynamic parameter announcement.
*   **ApplicationSet Strategies**: `RollingUpdate`, `RollingSync`, `AllAtOnce`, `Reverse` deletion.
*   **Impersonation**: `DestinationServiceAccounts`, `as-user`, `as-group`.
*   **Manifest Hydration**: `drySource`, `syncSource`, `hydrateTo`, push-to-deploy/stage.
*   **OCI Support**: `vnd.cncf.argoproj.argocd.content.v1.tar+gzip`.

## Related Components
*   **Application Controller**: Primary resource management and reconciliation.
*   **ApplicationSet Controller**: Automated application generation and rollout orchestration.
*   **Argo CD API Server**: Gateway for UI/CLI requests and extension proxying.
*   **Argo CD Repo Server**: Manifest generation and repository (Git/Helm/OCI) interfacing.
*   **Argo CD UI**: The React-based dashboard.
*   **Argo CD CLI**: The command-line interface for administrative and user actions.