# OPERATOR-MANUAL Documentation Index

## Overview
The `operator-manual` provides the definitive guide for Argo CD administrators to install, configure, secure, and maintain the system. It covers core architecture, multi-tenancy via RBAC and SSO, automated application management through the ApplicationSet controller, event-driven notifications, and detailed upgrade paths across major and minor versions.

## Files Summary

### Core Architecture & Installation
*   **architecture.md**: High-level architectural overview of the API Server, Repository Server, and Application Controller.
*   **installation.md**: Multi-tenant installation modes (HA and non-HA) using Kustomize and Helm.
*   **core.md**: Describes "Argo CD Core," a lightweight, headless installation for GitOps without the API or UI.
*   **feature-maturity.md**: Status table tracking stability (Alpha/Beta/Stable) of features.
*   **declarative-setup.md**: Primary reference for managing Argo CD resources using Kubernetes manifests.
*   **cluster-bootstrapping.md**: Details the "App of Apps" pattern for managing groups of applications.
*   **cluster-management.md**: CLI-based instructions for adding and removing Kubernetes clusters.

### Configuration Templates (YAML)
*   **argocd-cm-yaml.md / argocd-cmd-params-cm-yaml.md / argocd-rbac-cm-yaml.md / argocd-secret-yaml.md**: Templates for core settings, command parameters, access control, and sensitive data.
*   **argocd-repo-creds-yaml.md / argocd-repositories-yaml.md**: Templates for Git/Helm repository credentials.
*   **argocd-ssh-known-hosts-cm-yaml.md / argocd-tls-certs-cm-yaml.md**: Managing SSH public keys and custom CA/TLS certificates.

### User Management & Security
*   **rbac.md**: Detailed guide on the Casbin-based RBAC policy engine and group binding.
*   **security.md**: Master security document covering auth, TLS, and safe Git/Helm handling.
*   **user-management/index.md**: Overview of local accounts and SSO paths (Dex vs. OIDC).
*   **user-management/[microsoft.md/google.md/keycloak.md/onelogin.md/identity-center.md]**: Specific integration guides for Azure AD, Google Workspace, Keycloak, OneLogin, and AWS IAM Identity Center.
*   **secret-management.md**: Comparison of secret management strategies (Sealed Secrets, External Secrets, etc.).
*   **signed-release-assets.md**: Verifying container images and binaries using Cosign and SLSA.

### ApplicationSet Controller
*   **applicationset/index.md**: Introduction to the automated "factory" pattern for Application creation.
*   **applicationset/Generators.md**: Summary of all available generator types.
*   **applicationset/Generators-[Git.md/Cluster.md/SCM-Provider.md/Pull-Request.md]**: Deep dives into specific generators for monorepos, cluster targeting, and discovery.
*   **applicationset/Generators-Matrix.md / Merge.md / Plugin.md**: Advanced logic for combining generators or using custom third-party RPC generators.
*   **applicationset/GoTemplate.md**: Guide to the enhanced Go Text Template engine and Sprig functions.
*   **applicationset/Progressive-Syncs.md**: Experimental feature for staging updates via RollingSync.
*   **applicationset/[Security.md/Controlling-Resource-Modification.md/Appset-Any-Namespace.md]**: Security policies, dry-runs, and multi-namespace configuration.

### Notifications
*   **notifications/index.md**: Introduction to the notification controller and non-control-plane self-service.
*   **notifications/[triggers.md/templates.md/subscriptions.md/functions.md]**: Reference for trigger conditions, Go templating, and subscription annotations.
*   **notifications/[catalog.md/examples.md/monitoring.md/troubleshooting.md]**: Pre-configured triggers, practical recipes, Prometheus metrics, and debugging.
*   **notifications/services/overview.md**: Configuration of notification services in `argocd-notifications-cm`.
*   **notifications/services/[slack.md/webhook.md/github.md/aws/email/pagerduty/etc]**: Service-specific guides for integrations like Slack, GitHub Status, PagerDuty, and generic webhooks.

### Server Commands & Operations
*   **server-commands/[argocd-server.md/argocd-application-controller.md/argocd-repo-server.md/argocd-applicationset-controller.md/argocd-dex.md]**: CLI references for all core components.
*   **server-commands/additional-configuration-method.md**: Using `argocd-cmd-params-cm` to set global flags.
*   **high_availability.md**: Scaling strategies, monorepo optimizations, and sharding algorithms.
*   **dynamic-cluster-distribution.md**: Alpha feature for heartbeat-based shard management.
*   **metrics.md**: Reference of Prometheus metrics for all components.
*   **health.md / resource_actions.md**: Built-in and custom Lua-based health assessments and actions.
*   **reconcile.md**: Optimization techniques for controller load and resource ignoring.
*   **disaster_recovery.md**: Exporting and importing Argo CD data.
*   **troubleshooting.md**: Using `argocd admin` to validate settings and connectivity.

### Upgrading Path
*   **upgrading/overview.md**: Semantic versioning rules and standard upgrade commands.
*   **upgrading/[3.x.md/2.x.md/1.x.md]**: Version-specific breaking changes, manifest fixes, and migration steps (e.g., CMP v2 transition, resource tracking changes, API deprecations).

### UI & Extensibility
*   **ingress.md / tls.md**: Configuration for Ingress controllers (Nginx, ALB, Istio, etc.) and inbound/inter-component TLS.
*   **config-management-plugins.md / custom_tools.md**: Extending Argo CD with CMP sidecars and custom binaries.
*   **ui-customization.md / custom-styles.md / deep_links.md**: Branding, custom CSS, and conditional external links to monitoring tools.
*   **web_based_terminal.md**: Enabling browser-based `exec` into pods.
*   **webhook.md**: Configuring Git provider webhooks for instant refreshes.

## Code Changes That Would Require Documentation Updates
*   **Metrics**: Adding new metrics, changing labels, or modifying default ports.
*   **CLI/Server Parameters**: Adding or deprecating flags in any server component or ConfigMap parameters.
*   **CRD Schema**: Updates to `Application`, `AppProject`, `ApplicationSet`, or `ConfigManagementPlugin`.
*   **RBAC & Multi-tenancy**: Introducing new resources/actions (e.g., `extensions`, `exec`) or namespace validation logic.
*   **Auth/SSO**: Changes to Dex connectors, OIDC flows (e.g., PKCE), or session handling.
*   **Resource Tracking & Diffing**: Modifying reconciliation logic, Server-Side Apply (SSA) behavior, or diffing algorithms.
*   **Lua Environment**: Updating Lua versions or available library whitelists for health/actions.
*   **ApplicationSet Generators**: New generator types, field additions (SCM/PR filters), or templating logic changes.
*   **Notification Engine**: New service providers, template functions, or trigger evaluation logic.
*   **Tooling/Environment**: Upgrading bundled binaries (Helm, Kustomize), base image shifts (e.g., Ubuntu/Tini), or security defaults (e.g., symlink blocking).

## Key Technical Concepts
*   **Controllers**: Application Controller (reconciliation), ApplicationSet Controller (automation), Notifications Controller (alerting).
*   **Servers**: API Server (argocd-server), Repo Server (manifest generation), Dex (identity).
*   **Generators**: The engine for ApplicationSets (List, Cluster, Git, SCM, PR, Matrix, Merge, Plugin).
*   **Authentication**: OIDC, SAML, JWT, Dex, PKCE, Workload Identity (Azure/AWS).
*   **Authorization**: Casbin RBAC, Scopes, AppProjects.
*   **Configuration Management**: Helm, Kustomize, CMP (Config Management Plugins), Sidecars.
*   **Sharding**: Cluster distribution via Round-Robin, Consistent Hashing, and Heartbeats.
*   **Networking**: gRPC/gRPC-Web, TLS/mTLS, SSL-Passthrough, Webhooks, Egress NetworkPolicies.
*   **State Management**: Health Checks (Lua), Resource Actions (Lua), Reconciliation, Resource Tracking (Annotation vs. Label).
*   **Notifications Logic**: Triggers, Templates, Subscriptions, and `oncePer` deduplication.

## Related Components
*   **Kubernetes API Server**: Target for resource deployment and status.
*   **Infrastructure Services**: Redis (caching), HA-Proxy, Prometheus (metrics).
*   **Git Providers**: GitHub, GitLab, Bitbucket, Azure DevOps, Gitea.
*   **Ingress Controllers**: Nginx, Traefik, AWS ALB, Ambassador, Istio, Contour.
*   **Secret Management**: Sealed Secrets, External Secrets, Secrets Store CSI.
*   **Upstream Libraries**: Dex (Auth), Sprig (Templating), Cosign (Signing).