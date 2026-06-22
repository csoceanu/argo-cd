# OPERATOR-MANUAL Documentation Index

## Overview
This documentation area provides comprehensive guidance for administrators and platform engineers responsible for installing, configuring, and maintaining Argo CD. It covers core architectural components, security hardening, high availability scaling, resource customization through Lua, and integration with external systems like OIDC providers and Ingress controllers.

## Files Summary
*   **operator-manual/app-any-namespace.md**: Explains how to enable and configure Argo CD to manage Application resources located in namespaces other than the control plane namespace.
*   **operator-manual/app-sync-using-impersonation.md**: Details the alpha feature for syncing applications using specific Kubernetes Service Accounts via impersonation to limit control plane privileges.
*   **operator-manual/architecture.md**: Provides a high-level overview of Argo CD's internal components, including the API Server, Repository Server, and Application Controller.
*   **operator-manual/argocd-cm-yaml.md**: Provides a reference example for the primary `argocd-cm.yaml` configuration file.
*   **operator-manual/argocd-cmd-params-cm-yaml.md**: Provides an example manifest for `argocd-cmd-params-cm.yaml`, used for setting environment-level parameters.
*   **operator-manual/argocd-rbac-cm-yaml.md**: Provides a reference example for configuring Role-Based Access Control via the `argocd-rbac-cm.yaml` file.
*   **operator-manual/argocd-repo-creds-yaml.md**: Displays an example of the `argocd-repo-creds.yaml` file used for repository credential templates.
*   **operator-manual/argocd-repositories-yaml.md**: Provides an example for defining Git and Helm repository connections in `argocd-repositories.yaml`.
*   **operator-manual/argocd-secret-yaml.md**: Provides an example of the `argocd-secret.yaml` file which stores sensitive data like passwords and OIDC client secrets.
*   **operator-manual/argocd-ssh-known-hosts-cm-yaml.md**: Provides an example for the `argocd-ssh-known-hosts-cm.yaml` file used for SSH public key verification.
*   **operator-manual/argocd-tls-certs-cm-yaml.md**: Provides an example for the `argocd-tls-certs-cm.yaml` file used for custom CA and self-signed certificates.
*   **operator-manual/cluster-bootstrapping.md**: Describes the "App of Apps" pattern for declaratively managing groups of applications in a new cluster.
*   **operator-manual/cluster-management.md**: Covers CLI-based management of clusters, including adding and removing destination clusters.
*   **operator-manual/config-management-plugins.md**: Detailed guide on creating and installing Config Management Plugins (CMP) using sidecar containers.
*   **operator-manual/core.md**: Introduces Argo CD Core, a lightweight, headless installation intended for cluster administrators.
*   **operator-manual/custom-styles.md**: Explains how to customize the Argo CD UI appearance using CSS injection and informational banners.
*   **operator-manual/custom_tools.md**: Guidance on including custom binaries or toolchains (like specific Helm/Kustomize versions) via volume mounts or custom images.
*   **operator-manual/declarative-setup.md**: The primary guide for defining applications, projects, repositories, and clusters using Kubernetes manifests.
*   **operator-manual/deep_links.md**: Explains how to configure external deep links in the UI to redirect users to third-party monitoring or logging systems.
*   **operator-manual/disaster_recovery.md**: Provides instructions for exporting and importing all Argo CD data for backup and recovery purposes.
*   **operator-manual/dynamic-cluster-distribution.md**: Covers the alpha feature for dynamic sharding and cluster distribution based on controller heartbeats.
*   **operator-manual/feature-maturity.md**: Tracks the stability status (Alpha, Beta, Stable) of various Argo CD features and configurations.
*   **operator-manual/health.md**: Detailed instructions on built-in health checks and writing custom health assessment logic using Lua.
*   **operator-manual/high_availability.md**: Advanced guide for scaling Argo CD components, including sharding, monorepo optimizations, and cache tuning.
*   **operator-manual/index.md**: The entry point for the operator manual, providing a high-level roadmap for administrators.
*   **operator-manual/ingress.md**: Comprehensive configuration examples for various Ingress controllers including NGINX, Traefik, AWS ALB, and Istio.
*   **operator-manual/installation.md**: Covers various installation methods, including Multi-Tenant, HA, Core, Helm, and Kustomize.
*   **operator-manual/metrics.md**: A complete reference of Prometheus metrics exposed by the controller, API server, and repo server.
*   **operator-manual/project-specification.md**: Provides the full YAML schema and field descriptions for the `AppProject` Custom Resource.
*   **operator-manual/rbac.md**: In-depth guide to the RBAC model, policy CSV syntax, and mapping SSO groups to internal roles.
*   **operator-manual/reconcile.md**: Documentation on optimizing the reconciliation process by ignoring specific resource updates via JSON pointers or JQ.
*   **operator-manual/resource_actions.md**: Explains how to define custom Lua scripts to perform operational actions on resources from the UI.
*   **operator-manual/resource_actions_builtin.md**: A list of links to the Lua definitions for built-in resource actions like restarting deployments.
*   **operator-manual/secret-management.md**: Compares different secret management strategies, including External Secrets and manifest-generation plugins.
*   **operator-manual/security.md**: Outlines security implementation details including authentication, authorization, TLS, and auditing.
*   **operator-manual/signed-release-assets.md**: Instructions for verifying the authenticity of Argo CD container images and binaries using Cosign and SLSA provenance.
*   **operator-manual/tested-kubernetes-versions.md**: Placeholder/guide for checking compatibility between Argo CD versions and Kubernetes versions.
*   **operator-manual/tls.md**: Detailed configuration of inbound and inter-component TLS for all Argo CD services.
*   **operator-manual/troubleshooting.md**: Describes admin CLI tools for validating settings, resource overrides, and cluster connectivity.
*   **operator-manual/ui-customization.md**: Covers UI-specific settings like default application views and node label propagation in pod views.
*   **operator-manual/web_based_terminal.md**: Instructions for enabling the browser-based terminal (`exec`) feature and securing it via RBAC.
*   **operator-manual/webhook.md**: Explains how to configure Git provider webhooks to trigger immediate application refreshes on code changes.

## Code Changes That Would Require Documentation Updates
*   **CRD Schema Changes**: Any additions or modifications to the `Application`, `AppProject`, or `ApplicationSet` spec/status fields.
*   **New Configuration Options**: Adding new keys to `argocd-cm`, `argocd-cmd-params-cm`, or `argocd-rbac-cm`.
*   **Metrics Updates**: Adding, renaming, or changing the labels of any Prometheus metrics in any component.
*   **CLI Modifications**: Updates to `argocd admin`, `argocd cluster`, or `argocd app` commands, specifically those used for operator tasks.
*   **Sharding/Reconciliation Logic**: Changes to how clusters are distributed across shards or how the controller caches resources.
*   **Security & Auth**: Changes to JWT handling, SSO integration logic, or RBAC resource/action definitions.
*   **Lua Environment**: Updates to the Lua libraries available for custom health checks or resource actions.
*   **Installation Manifests**: Changes to the base manifests, HA configurations, or Core installation logic.
*   **CMP Logic**: Modifications to how Config Management Plugins are discovered, initialized, or executed.
*   **Terminal/Exec Feature**: Changes to how the web terminal connects to pods or how shells are detected.

## Key Technical Concepts
*   **Custom Resources**: `Application`, `AppProject`, `ApplicationSet`.
*   **ConfigMaps**: `argocd-cm`, `argocd-cmd-params-cm`, `argocd-rbac-cm`, `argocd-tls-certs-cm`, `argocd-ssh-known-hosts-cm`.
*   **Sharding**: `legacy`, `round-robin`, `consistent-hashing`, `heartbeat`, `dynamic distribution`.
*   **Security**: RBAC, OIDC, JWT, TLS termination, SSL passthrough, Signed assets, SLSA.
*   **Extensibility**: Lua Health Checks, Custom Resource Actions, Config Management Plugins (CMP).
*   **Reconciliation**: Resource tracking, sync waves, ignoreDifferences, manifest-generate-paths.
*   **Protocols**: gRPC, gRPC-web, HTTPS.

## Related Components
*   **argocd-server**: The API and UI gateway.
*   **argocd-application-controller**: The main reconciliation engine.
*   **argocd-repo-server**: The manifest generation service.
*   **argocd-dex-server**: The bundled authentication provider.
*   **argocd-applicationset-controller**: The automation engine for generating multiple applications.
*   **argocd-notifications-controller**: The notification delivery service.
*   **Redis**: The distributed cache for manifests and cluster state.