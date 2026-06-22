# PROPOSALS Documentation Index

## Overview
This documentation area contains architectural blueprints and feature specifications for the Argo CD ecosystem. It covers fundamental shifts in the project's design, including multi-tenancy models, UI extensibility, enhanced configuration management, and scalability improvements for both the core Application controller and ApplicationSets.

## Files Summary
*   **multiple-sources-for-applications-ui.md**: Details the UI/UX changes required to support applications with multiple Git/Helm sources (tabs, history, and creation dialog).
*   **deletion-strategy-progressive-sync.md**: Proposes deletion orders (`AllAtOnce`, `Reverse`) for ApplicationSet progressive syncs.
*   **config-management-plugin-v2.md**: Defines the "CMP v2" architecture using sidecar containers and a dedicated `argocd-cmp-server`.
*   **server-side-pagination.md**: Outlines the implementation of server-side pagination for Application List/Watch APIs to improve UI responsiveness.
*   **004-scalability-benchmarking.md**: Establishes a framework and tooling for objective Argo CD scalability testing.
*   **argocd-cli-pluin.md**: Introduces `kubectl`-like plugin support for the `argocd` CLI.
*   **respect-rbac-for-resource-exclusions.md**: Logic for the controller to respect Kubernetes RBAC permissions when deciding which resources to monitor.
*   **applicationset-plugin-generator.md**: Introduces a plugin-based generator for ApplicationSets that utilizes RPC calls for data.
*   **rebalancing-clusters-across-shards-dynamically.md**: Migration from StatefulSet to Deployment for the controller and heartbeat-based shard assignment.
*   **feature-bounties.md**: Guidelines for an experimental monetary reward program for community contributions.
*   **notifications-API.md**: Adds API endpoints and UI elements to allow users to subscribe applications to notification triggers.
*   **2022-07-13-appset-progressive-rollout-strategy.md**: Defines `RollingUpdate` and `RollingSync` strategies for ApplicationSets.
*   **headless-argocd.md**: A mode for Argo CD without an API/UI server, relying on local CLI-driven proxying.
*   **manifest-hydrator.md**: Proposes a "rendered manifests" pattern where Argo CD hydrates manifests and pushes them back to a Git branch.
*   **002-ui-extensions.md**: Defines the `ArgoCDExtension` CRD for injecting custom React components and Lua scripts into the UI.
*   **parameterized-config-management-plugins.md**: Enhances CMPs to "announce" and consume UI-driven parameters similar to Helm/Kustomize.
*   **server-side-apply.md**: Implementation details for using Kubernetes Server-Side Apply (SSA) during synchronization.
*   **native-oci-support.md**: Support for pulling manifests from OCI registries (independent of Helm).
*   **project-scoped-repository-enhancements.md**: Allows multiple repository credentials for the same URL, isolated by `AppProject`.
*   **project-repos-and-clusters.md**: Enables self-service onboarding by allowing repos and clusters to be scoped to specific projects.
*   **decouple-application-sync-user-using-impersonation.md**: Uses Kubernetes Impersonation to allow app syncs to run with specific ServiceAccount privileges.
*   **multiple-sources-for-applications.md**: The core proposal for the `sources` field (plural) in the Application CRD.
*   **003-applications-outside-argocd-namespace.md**: Allows Application CRs to exist in any namespace, utilizing `AppProject` source namespace whitelists.
*   **application-name-identifier.md**: Proposes moving from labels to annotations for resource tracking to allow long application names.
*   **backend-support-appset.md**: Adds API/CLI/UI support for managing ApplicationSet resources with proper RBAC.
*   **deep-links.md**: Configurable external links in the UI for Projects, Applications, and Resources.
*   **resource-deletion-with-approval.md**: Adds a safety net requiring manual approval before pruning or deleting specific resources.
*   **sync-timeout.md**: Introduces sync operation timeouts and health-based termination settings.
*   **proxy-extensions.md**: Backend service support for UI extensions via a reverse proxy in the API server.
*   **hide-annotations.md**: A feature to hide sensitive or internal annotations (e.g., secrets) from the Web UI.

## Code Changes That Would Require Documentation Updates
*   **CRD Schema Modifications**: Any changes to `Application`, `AppProject`, `ApplicationSet`, or `ArgoCDExtension` fields (e.g., adding `sources`, `sourceHydrator`, `sourceNamespaces`, or `destinationServiceAccounts`).
*   **Sync Engine Logic**: Changes to how resources are applied (Client-side vs. Server-side Apply) or how they are tracked (`app.kubernetes.io/instance` label vs. annotation).
*   **Multi-tenancy & RBAC**: Updates to the `AppProject` validation logic or the introduction of new RBAC resources (e.g., `extensions`, `applicationsets`).
*   **Controller Architecture**: Changing the Application Controller from a StatefulSet to a Deployment, or modifying the sharding/rebalancing algorithm.
*   **Plugin Systems**: Alterations to the CMP v2 sidecar protocol, the `ConfigManagementPlugin` spec, or UI extension loading mechanisms.
*   **API/CLI Framework**: Adding new subcommands to the `argocd` CLI or introducing new gRPC service endpoints in the API server (e.g., the Notifications API or Server-side pagination).
*   **Sync Policy Behaviors**: Modifying sync waves, hooks, or introducing new termination/timeout logic in the synchronization loop.
*   **Git/OCI Integration**: Changes to how Argo CD connects to repositories, handles OCI media types, or performs Git push operations for hydration.

## Key Technical Concepts
*   **CMP v2**: Sidecar-based Config Management Plugins using `argocd-cmp-server`.
*   **Multiple Sources**: Using the `sources` field to combine multiple Git/Helm repos in one Application.
*   **Manifest Hydration**: The process of rendering "dry" manifests into a "hydrated" branch.
*   **Progressive Sync**: Advanced ApplicationSet rollout/deletion strategies (`RollingUpdate`, `RollingSync`).
*   **Server-Side Apply (SSA)**: Delegating resource merging/patching to the Kubernetes API server.
*   **Tracking Methods**: The mechanism (label vs. annotation) used to map live resources to Argo CD Applications.
*   **Impersonation**: Using `kubectl` impersonation headers to sync resources as a specific ServiceAccount.
*   **Sharding**: The distribution of cluster management across multiple controller replicas.
*   **Deep Links**: Configurable UI shortcuts to external observability or auditing tools.
*   **Headless Mode**: A CLI-first operational mode that bypasses the centralized API server.

## Related Components
*   **argocd-application-controller**: Core logic for reconciliation and sharding.
*   **argocd-server**: The API and UI backend, including reverse-proxy and extension logic.
*   **argocd-repo-server**: Component responsible for manifest generation and hydration.
*   **argocd-applicationset-controller**: Specialized controller for generating and managing groups of applications.
*   **Argo CD CLI**: The `argocd` command-line tool and its plugin system.
*   **Argo CD UI**: The React-based frontend and its extension framework.
*   **GitOps Engine**: The underlying library used for diffing and applying resources.