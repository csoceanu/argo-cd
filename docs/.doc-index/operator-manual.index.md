# OPERATOR-MANUAL Documentation Index

## Overview
The "operator-manual" documentation provides comprehensive instructions for cluster administrators and platform engineers to install, configure, secure, and maintain Argo CD. It covers the internal architecture, security hardening (including RBAC and TLS), performance optimization for large-scale GitOps (sharding and reconciliation), and extensibility through custom Lua scripts and sidecar plugins.

## Files Summary
- **architecture.md**: High-level overview of system components (API Server, Repo Server, Application Controller).
- **security.md**: Details on authentication (JWT/SSO), authorization, Git/Helm security, and auditing.
- **rbac.md**: Guide to configuring Casbin-based access control policies for users and groups.
- **tls.md**: Configuration for inbound TLS, inter-component communication, and certificate management.
- **installation.md**: Instructions for Multi-tenant (HA/non-HA) and Core (headless) installation methods.
- **high_availability.md**: Advanced scaling techniques, including sharding algorithms and monorepo optimizations.
- **config-management-plugins.md**: Documentation for CMP v2, detailing sidecar configuration and manifest generation.
- **metrics.md**: Inventory of Prometheus metrics for all Argo CD components.
- **health.md**: Guide to built-in and custom Lua-based health assessment for Kubernetes resources.
- **resource_actions.md**: Instructions for defining custom Lua scripts to perform UI-driven operations on resources.
- **reconcile.md**: Optimization settings to reduce controller load by ignoring specific resource field updates.
- **declarative-setup.md**: Reference for defining Applications, Projects, and Repositories via Kubernetes manifests.
- **ingress.md**: Recipes for configuring Ingress controllers (NGINX, Traefik, Istio, ALB) with gRPC support.
- **secret-management.md**: Comparison of secret handling strategies (Operators vs. Manifest Generation).
- **web_based_terminal.md**: Setup guide for the browser-based terminal (`exec`) feature.
- **deep_links.md**: Configuration for templated external links to third-party monitoring or logging tools.
- **core.md**: Specifics on the "Argo CD Core" lightweight installation for cluster admins.
- **app-sync-using-impersonation.md**: Alpha feature guide for syncing applications using specific Service Accounts.
- **signed-release-assets.md**: Procedures for verifying container images and binaries using Cosign and SLSA.
- **webhook.md**: Configuration for Git provider webhooks to trigger immediate application refreshes.
- **cluster-management.md / cluster-bootstrapping.md**: CLI-based cluster operations and the "App of Apps" pattern.
- **custom_tools.md**: Instructions for overriding bundled binaries or using custom container images.
- **feature-maturity.md**: Status tracker (Alpha/Beta/Stable) for experimental and new features.
- **troubleshooting.md / disaster_recovery.md**: Admin tools for settings validation and backup/restore procedures.
- **ui-customization.md / custom-styles.md**: Branding, CSS overrides, and notification banners for the Web UI.
- **project-specification.md**: Full schema reference for the `AppProject` Custom Resource.
- **argocd-*-yaml.md**: Various reference examples for core ConfigMaps and Secrets.

## Code Changes That Would Require Documentation Updates
- **CRD Schema Changes**: Any modification to the `Application`, `AppProject`, or `ApplicationSet` spec or status fields.
- **CLI Command Updates**: Adding, removing, or changing flags/arguments for `argocd admin` or `argocd cluster` commands.
- **Metric Definitions**: Adding new Prometheus metrics or changing existing label names in any component.
- **Auth/Security Logic**: Changes to JWT handling, RBAC evaluation logic (Casbin), or OIDC integration.
- **Default Manifests**: Updates to the base installation YAMLs (HA vs. non-HA) or ServiceAccount permissions.
- **Templating Tool Versions**: Upgrading or changing the bundled versions of Helm, Kustomize, or Jsonnet.
- **Controller Logic**: Changes to the reconciliation loop, sharding algorithms, or resource tracking (labels/annotations).
- **Environment Variables**: Introducing new `ARGOCD_` environment variables in `argocd-cmd-params-cm`.
- **Lua Environment**: Modifications to the Lua sandbox or standard libraries available for health checks and actions.
- **Web UI Features**: Changes to default views, terminal behavior, or customization options (CSS/Banners).

## Key Technical Concepts
- **Resource Tracking**: The mechanism (labels vs. annotations) used to link live resources to Argo CD Applications.
- **Sharding**: Distributing clusters across multiple controller replicas using legacy, round-robin, or consistent-hashing methods.
- **Casbin**: The policy engine used for RBAC (Syntax: `p, sub, res, act, obj, eft`).
- **CMP v2**: Config Management Plugins using sidecar containers for manifest generation.
- **App of Apps**: A recursive deployment pattern where one Application manages multiple child Applications.
- **Headless Mode (Core)**: A minimalist installation without the API Server or Web UI.
- **Health Status**: The lifecycle state of a resource (Healthy, Progressing, Degraded, Suspended).
- **SLSA/Provenance**: Security attestations used to verify the supply chain of release artifacts.
- **gRPC/HTTP2**: The protocols required for CLI/UI communication, often requiring specific Ingress configuration.
- **Impersonation**: The ability for the controller to act as a specific Service Account during syncs.

## Related Components
- **API Server (`argocd-server`)**: Handles UI/CLI traffic and RBAC.
- **Application Controller (`argocd-application-controller`)**: Executes reconciliation and syncs.
- **Repository Server (`argocd-repo-server`)**: Clones repos and generates manifests.
- **Dex (`argocd-dex-server`)**: Manages OIDC authentication.
- **Redis**: Caching layer for manifest generation and cluster state.
- **Prometheus**: Monitoring target for system health and application status.
- **Git Providers**: GitHub, GitLab, Bitbucket, Azure DevOps (Webhook/Repo sources).