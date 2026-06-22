# OPERATOR-MANUAL Documentation Index

## Overview
The Operator Manual documentation provides comprehensive guidance for administrators and DevOps engineers responsible for installing, configuring, securing, and scaling Argo CD. It covers the architectural internals, security hardening (AuthN/AuthZ/TLS), high availability tuning, and the declarative management of clusters, repositories, and applications.

## Files Summary
*   **argocd-cmd-params-cm-yaml.md**: Provides a template and example for the `argocd-cmd-params-cm` ConfigMap used to set environment-level parameters.
*   **argocd-secret-yaml.md**: Contains examples of the `argocd-secret` manifest used to store sensitive data like passwords and OIDC client secrets.
*   **security.md**: Detailed guide on Argo CD security architecture, including authentication (JWT/SSO), authorization, TLS, and Git repository security considerations.
*   **metrics.md**: Lists all Prometheus metrics exposed by the Application Controller, API Server, Repo Server, and ApplicationSet Controller.
*   **tls.md**: Technical instructions for configuring inbound TLS for components and inter-component communication (e.g., Server to Repo-Server).
*   **argocd-cm-yaml.md**: Provides an example of the main `argocd-cm` ConfigMap for general system settings and customizations.
*   **argocd-tls-certs-cm-yaml.md**: Example for the `argocd-tls-certs-cm` used to manage custom CA certificates for Git/Helm repositories.
*   **argocd-ssh-known-hosts-cm-yaml.md**: Example manifest for managing SSH public keys of Git servers to prevent man-in-the-middle attacks.
*   **feature-maturity.md**: Tracks the current lifecycle status (Alpha, Beta, Stable) of various Argo CD features and configurations.
*   **argocd-repo-creds-yaml.md**: Example for creating repository credential templates to share credentials across multiple repositories.
*   **installation.md**: Covers the different installation flavors (Multi-tenant vs. Core) and methods (Kustomize, Helm, HA vs. non-HA).
*   **resource_actions.md**: Explains how to define custom Lua scripts to allow users to perform actions (like "restart" or "scale") directly on Kubernetes resources.
*   **secret-management.md**: Compares different secret management strategies, specifically destination-cluster operators versus manifest-generation plugins.
*   **reconcile.md**: Documentation on optimizing reconciliation performance by ignoring resource updates on specific fields using JSON pointers or JQ.
*   **rbac.md**: Comprehensive guide on the RBAC model, policy syntax (Casbin), SSO group mapping, and fine-grained resource permissions.
*   **config-management-plugins.md**: Detailed instructions for creating and installing sidecar-based Config Management Plugins (CMP) to support custom templating tools.
*   **cluster-bootstrapping.md**: Describes the "App of Apps" pattern for declaratively managing groups of applications across multiple clusters.
*   **web_based_terminal.md**: Instructions for enabling and securing the UI-based terminal feature (`exec`) for troubleshooting pods.
*   **ingress.md**: Provides configuration examples for various Ingress controllers (Nginx, ALB, Istio, Traefik, etc.) to expose the Argo CD UI and gRPC API.
*   **cluster-management.md**: Guide for managing target clusters using the `argocd cluster` CLI commands.
*   **troubleshooting.md**: Introduction to `argocd admin` subcommands for validating settings, resource overrides, and connectivity issues.
*   **app-sync-using-impersonation.md**: Explains the alpha feature for syncing applications using specific Kubernetes Service Accounts via impersonation.
*   **project-specification.md**: Provides the full YAML schema and field descriptions for the `AppProject` Custom Resource.
*   **ui-customization.md**: Guides on customizing the UI with banners, custom CSS, and specific default views or node labels.
*   **high_availability.md**: Advanced scaling guide for high-volume environments, covering sharding algorithms, processor tuning, and monorepo optimizations.
*   **health.md**: Documentation for built-in and custom Lua-based health checks used to determine the "Healthy" or "Degraded" status of resources.
*   **deep_links.md**: Instructions for configuring external deep links (e.g., to Splunk or Datadog) based on application or resource data.
*   **declarative-setup.md**: The foundational guide for managing Applications, Projects, Repositories, and Clusters using Kubernetes manifests.
*   **custom_tools.md**: Instructions on how to add custom binaries (like specific Helm or Kustomize versions) via volume mounts or custom images.
*   **index.md**: The landing page for the operator manual, directing users to installation and configuration entry points.
*   **disaster_recovery.md**: Brief guide on using `argocd admin export/import` for backing up and restoring Argo CD state.
*   **core.md**: Describes the "Argo CD Core" (headless) mode for lightweight GitOps without the API server or Web UI.
*   **architecture.md**: High-level overview of the internal components (API Server, Repository Server, Application Controller) and their interactions.
*   **signed-release-assets.md**: Procedures for verifying the integrity of Argo CD container images and CLI binaries using Cosign and SLSA.
*   **webhook.md**: Guide for configuring Git provider webhooks (GitHub, GitLab, etc.) to trigger instant application refreshes.
*   **argocd-repositories-yaml.md**: Example manifest for defining repository connection secrets for Git and Helm.

## Code Changes That Would Require Documentation Updates
*   **Adding/Renaming Metrics**: Changes to the Prometheus metrics emitted by any controller must be reflected in `metrics.md`.
*   **New RBAC Resources/Actions**: Adding a new resource type (e.g., `extensions`) or a new action (e.g., `invoke`) requires updates to `rbac.md`.
*   **Configuration Logic**: Adding new fields to `argocd-cm` or `argocd-cmd-params-cm` requires updates to the respective YAML examples and potentially `high_availability.md` or `reconcile.md`.
*   **API/Architecture Changes**: Modifying how the Server, Repo-Server, or Controller interact requires updates to `architecture.md` and `tls.md`.
*   **Feature Graduation**: Moving a feature from Alpha to Beta or Beta to Stable must be updated in `feature-maturity.md`.
*   **Custom Resource Changes**: Any schema changes to the `Application` or `AppProject` CRDs require updates to `project-specification.md` and `declarative-setup.md`.
*   **Lua Environment**: Changes to the global variables or libraries available to Lua scripts require updates to `health.md`, `resource_actions.md`, and `deep_links.md`.
*   **CLI Admin Commands**: New subcommands for `argocd admin` or changes to existing troubleshooting workflows require updates to `troubleshooting.md` and `disaster_recovery.md`.
*   **AuthN/AuthZ Workflows**: Changes to the session management or OIDC handling code require updates to `security.md`.

## Key Technical Concepts
*   **Resource Tracking**: Label-based vs. Annotation-based methods for identifying managed resources.
*   **Sharding**: The distribution of cluster management across multiple controller replicas (Legacy, Round-Robin, Consistent-Hashing).
*   **Lua Customization**: Scripting for health checks, resource actions, and deep link evaluation.
*   **Casbin**: The policy engine used for RBAC.
*   **App of Apps**: A hierarchical pattern for cluster bootstrapping.
*   **Config Management Plugins (CMP)**: The sidecar architecture for extending templating capabilities.
*   **Keyless Signing**: Using Cosign and OIDC for artifact verification.
*   **Reconciliation Loop**: The process of comparing Git desired state vs. Live cluster state.
*   **Impersonation**: Using Kubernetes headers to sync resources with reduced privileges.

## Related Components
*   **argocd-server**: The API and UI gateway.
*   **argocd-application-controller**: The state reconciliation engine.
*   **argocd-repo-server**: The manifest generation engine.
*   **argocd-dex-server**: The identity federation service.
*   **argocd-applicationset-controller**: The automated application generator.
*   **argocd-redis**: The caching layer.
*   **argocd-notifications-controller**: The event notification system.