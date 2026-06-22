# PROPOSALS Documentation Index

## Overview
This documentation area contains design proposals and architectural blueprints for the Argo CD ecosystem. It covers major feature enhancements, core system refactors, API changes, and user experience improvements for the Application Controller, API Server, Repository Server, and ApplicationSet Controller.

## Files Summary
*   **proposals/multiple-sources-for-applications-ui.md**: Details the UI modifications required to support and visualize applications that utilize multiple git/helm sources.
*   **proposals/deletion-strategy-progressive-sync.md**: Proposes configurable deletion orders (AllAtOnce vs. Reverse) for applications managed by ApplicationSet Progressive Sync.
*   **proposals/config-management-plugin-v2.md**: Describes the "v2" architecture for Config Management Plugins using sidecar containers and a dedicated CMP API server.
*   **proposals/server-side-pagination.md**: Outlines the implementation of pagination for Application List and Watch APIs to improve UI/CLI performance at scale.
*   **proposals/004-scalability-benchmarking.md**: Establishes a framework and methodology for objectively measuring Argo CD performance limits across different topologies.
*   **proposals/argocd-cli-pluin.md**: Proposes a plugin mechanism for the `argocd` CLI modeled after `kubectl` plugins to extend functionality without bloating the core binary.
*   **proposals/respect-rbac-for-resource-exclusions.md**: Details a feature to make the controller respect Kubernetes RBAC permissions when deciding which resources to monitor.
*   **proposals/applicationset-plugin-generator.md**: Introduces a "plugin" generator for ApplicationSets that retrieves parameters via RPC calls to external services.
*   **proposals/rebalancing-clusters-across-shards-dynamically.md**: Proposes moving the application controller to a Deployment model to allow for dynamic cluster-to-shard rebalancing.
*   **proposals/feature-bounties.md**: Outlines an experimental program for offering monetary rewards for the implementation of specific Argo CD features.
*   **proposals/notifications-API.md**: Describes new API endpoints to allow users to subscribe to notifications directly from the Application Details UI.
*   **proposals/2022-07-13-appset-progressive-rollout-strategy.md**: Proposes `RollingUpdate` and `RollingSync` strategies for controlled Application deployment in ApplicationSets.
*   **proposals/headless-argocd.md**: Outlines a "headless" mode that disables the API server and UI in favor of direct Kubernetes RBAC and local CLI management.
*   **proposals/manifest-hydrator.md**: Proposes a first-class "Rendered Manifests" pattern where Argo CD hydrates dry manifests and pushes them to a separate git branch.
*   **proposals/002-ui-extensions.md**: Defines a mechanism for resource-specific UI visualizations, health checks, and actions using dynamically loaded JS and Lua.
*   **proposals/parameterized-config-management-plugins.md**: Enhances CMPs to allow them to "announce" parameters to the UI, providing a user experience similar to native Helm/Kustomize.
*   **proposals/server-side-apply.md**: Describes integration with Kubernetes Server-Side Apply (SSA) to handle large CRDs and better interoperate with admission controllers.
*   **proposals/native-oci-support.md**: Proposes native support for sourcing manifests directly from OCI registries as a first-class repository type.
*   **proposals/project-scoped-repository-enhancements.md**: Allows multiple repository credentials for the same URL, isolated by AppProject scope for multi-tenant environments.
*   **proposals/project-repos-and-clusters.md**: Enables self-service onboarding by allowing developers to register repositories and clusters within a specific Project scope.
*   **proposals/decouple-application-sync-user-using-impersonation.md**: Proposes using Kubernetes user impersonation to perform syncs with the privileges of a specific ServiceAccount.
*   **proposals/multiple-sources-for-applications.md**: The core proposal for supporting the `sources` field in the Application CRD to combine multiple git/helm inputs.
*   **proposals/003-applications-outside-argocd-namespace.md**: Outlines the architectural changes needed to allow Application CRs to exist in namespaces other than the Argo CD install namespace.
*   **proposals/application-name-identifier.md**: Proposes changing the resource tracking mechanism from labels to annotations to support longer application names.
*   **proposals/backend-support-appset.md**: Details the API and CLI changes needed to manage ApplicationSets through the standard Argo CD interfaces.
*   **proposals/001-proposal-template.md**: Provides the standard structure and required sections for all new enhancement proposals.
*   **proposals/deep-links.md**: Proposes configurable UI links to third-party systems (monitoring, logging) based on application and resource metadata.
*   **proposals/resource-deletion-with-approval.md**: Introduces a manual approval step for pruning or deleting specific resources during synchronization.
*   **proposals/sync-timeout.md**: Proposes configurable timeouts and termination settings for synchronization operations that get stuck.
*   **proposals/proxy-extensions.md**: Extends the UI extension framework to include a reverse-proxy for communicating with custom backend services.

## Code Changes That Would Require Documentation Updates
*   **API/Protobuf Definitions**: Any modification to `app.proto`, `applicationset.proto`, or `project.proto` that adds fields or services.
*   **CRD Schema Updates**: Changes to `Application`, `ApplicationSet`, or `AppProject` spec/status structures.
*   **Controller Logic**: Changes to the reconciliation loop in `argocd-application-controller` or `applicationset-controller` (e.g., sharding, sync strategies).
*   **Tracking & Diffing**: Alterations to how resources are identified (labels vs. annotations) or how diffs are calculated (SSA vs. 3-way merge).
*   **Plugin Systems**: Changes to the `argocd-cmp-server` or the ways plugins are discovered and executed in sidecars.
*   **RBAC Policy**: New resource types or actions added to the Argo CD RBAC evaluator.
*   **Repository Clients**: Updates to how Argo CD interacts with Git, Helm, or OCI registries.
*   **UI Architecture**: Changes to the extension loading mechanism, dashboard layouts, or how application parameters are rendered.

## Key Technical Concepts
*   **Multiple Sources**: Combining multiple Git repos or Helm charts into a single Application.
*   **Progressive Sync/Rollout**: Ordered deployment and deletion of applications within an ApplicationSet.
*   **Tracking Method**: The mechanism (Label vs. Annotation) used to link live resources to an Argo Application.
*   **CMP v2**: Sidecar-based Config Management Plugins.
*   **Server-Side Apply (SSA)**: Leveraging Kubernetes native field ownership and patching.
*   **Cluster Sharding**: Distributing the management of multiple clusters across multiple controller replicas.
*   **Impersonation**: Executing sync operations as a specific Kubernetes ServiceAccount.
*   **Manifest Hydration**: The process of converting "dry" templates (Helm/Kustomize) into "hydrated" plain YAML.
*   **Resource Actions**: Lua-based custom scripts to perform operations on Kubernetes resources.

## Related Components
*   **argocd-application-controller**: The core state reconciliation engine.
*   **argocd-server**: The API gateway for UI and CLI.
*   **argocd-repo-server**: Handles manifest generation and repository interactions.
*   **applicationset-controller**: Manages the lifecycle of multiple applications from a template.
*   **gitops-engine**: The underlying library for resource synchronization and diffing.
*   **argocd-notifications**: Handles external alerts based on application state changes.