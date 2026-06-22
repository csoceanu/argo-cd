# OPERATOR-MANUAL Documentation Index

## Overview
This documentation area provides comprehensive guidance for Argo CD administrators and platform engineers responsible for the full operational lifecycle of the platform. It covers everything from initial installation (Helm/Kustomize) and security hardening (RBAC, SSO, TLS) to high-availability scaling, automated application lifecycles via ApplicationSets, outbound notification orchestration, and version-specific upgrade procedures.

## Files Summary

### Core Installation & Infrastructure
*   **architecture.md**: High-level overview of Argo CD components and their interactions.
*   **installation.md**: Detailed installation methods (Multi-tenant vs. Core, HA vs. Non-HA).
*   **core.md**: Guidance on the lightweight, headless "Argo CD Core" mode.
*   **high_availability.md**: Scaling components, cluster sharding, and monorepo performance.
*   **dynamic-cluster-distribution.md**: (Alpha) Documentation for runtime cluster sharding.
*   **ingress.md**: Configuration examples for Ingress controllers (Nginx, Traefik, ALB, Istio, etc.).
*   **tls.md**: Configuring inbound TLS and securing inter-component communication.
*   **metrics.md**: Prometheus metrics reference for all core components.
*   **signed-release-assets.md**: Verifying authenticity of binaries and images using Cosign/SLSA.
*   **feature-maturity.md**: Tracks the stability status (Alpha/Beta/Stable) of features.

### Configuration & Management
*   **declarative-setup.md**: Managing Argo CD resources via Kubernetes manifests.
*   **cluster-management.md**: CLI-based management for adding/removing target clusters.
*   **cluster-bootstrapping.md**: Implementing the "App of Apps" pattern.
*   **project-specification.md**: Technical schema for the `AppProject` Custom Resource.
*   **reconcile.md**: Optimizing reconciliation by ignoring specific resource updates.
*   **health.md**: Built-in health assessments and custom Lua health checks.
*   **resource_actions.md**: Creating custom Lua-based UI buttons for Kubernetes resources.
*   **config-management-plugins.md**: Creating and installing CMP sidecars for manifest generation.
*   **custom_tools.md**: Overriding bundled binaries (Helm/Kustomize) or building custom images.
*   **webhook.md**: Configuration for Git provider webhooks to trigger instant refreshes.
*   **troubleshooting.md / disaster_recovery.md**: Validating settings and backup/restore procedures.

### User Management, Security & UI
*   **security.md**: Comprehensive guide on authentication, authorization, and repo security.
*   **rbac.md**: Policy syntax, resource types, and mapping SSO groups to roles.
*   **user-management/index.md**: Local accounts, rate-limiting, and Dex architecture.
*   **Identity Provider Guides**: Specific OIDC/SAML integration for Okta, Auth0, Zitadel, Microsoft, OneLogin, OpenUnison, Identity Center, Keycloak, and Google.
*   **secret-management.md**: Strategies for secrets (CSI, External Secrets, Plugins).
*   **app-sync-using-impersonation.md**: (Alpha) Running syncs as specific Service Accounts.
*   **web_based_terminal.md**: Setup and security for the UI-based pod terminal.
*   **ui-customization.md**: Customizing the UI with banners, CSS, and view preferences.
*   **deep_links.md**: Templating external links (Splunk, Datadog) into the UI.

### ApplicationSets
*   **applicationset/index.md / applicationset-specification.md**: Introduction and YAML schema.
*   **applicationset/Generators.md**: High-level overview of all parameter sources.
*   **applicationset/Generators-*.md**: Deep-dives into specific generators (List, Cluster, Git, SCM, PR, Matrix, Merge, Plugin).
*   **applicationset/Security.md**: Security constraints, secret exfiltration risks, and project fields.
*   **applicationset/GoTemplate.md**: Using Go Text Templates and Sprig functions.
*   **applicationset/Progressive-Syncs.md**: Documentation for rolling updates and sequenced lifecycles.
*   **applicationset/Template.md**: Instructions on using templates and `templatePatch`.
*   **applicationset/Appset-Any-Namespace.md**: Enabling ApplicationSets in non-control-plane namespaces.

### Notifications
*   **notifications/index.md**: Entry point and namespace-based configuration.
*   **notifications/templates.md / triggers.md / subscriptions.md**: Defining content, conditions, and routing.
*   **notifications/catalog.md / examples.md**: Library of pre-defined triggers and Slack/Webhook use cases.
*   **notifications/functions.md**: Built-in Go template functions for metadata manipulation.
*   **notifications/monitoring.md**: Prometheus metrics for notification delivery.
*   **notifications/services/overview.md**: High-level service configuration and secret management.
*   **notifications/services/[service-name].md**: Provider-specific setup (Slack, Webhook, PagerDuty, AWS SQS, etc.).
*   **notifications/troubleshooting.md / commands.md / errors.md**: Debugging via CLI.

### Reference & Upgrading
*   **server-commands/*.md**: CLI reference for all core binaries (server, repo-server, etc.).
*   **upgrading/overview.md**: Semantic versioning rules and standard upgrade commands.
*   **upgrading/[version-to-version].md**: Detailed migration guides for specific version jumps.
*   **Placeholder Configs**: Example `.yaml.md` files for core ConfigMaps and Secrets.

## Code Changes That Would Require Documentation Updates
*   **API & CRD Modifications**: Changes to the schema of `Application`, `AppProject`, or `ApplicationSet` (updates to specifications and declarative setup).
*   **Authentication & Authz**: Adding auth protocols, changing Dex configurations, modifying group claim parsing, or introducing new RBAC resource types.
*   **ApplicationSet Logic**: Adding/removing generator fields, modifying templating engines (Sprig/Go), or updating Progressive Sync logic.
*   **Notification Engine**: Adding new service integrations, updating trigger evaluation (expr), or adding helper functions to the notification template context.
*   **Infrastructure & CLI**: Adding/renaming binary flags, changing `argocd-cm` ConfigMap keys, or adding new subcommands to `argocd admin`.
*   **System Metrics**: Adding new Prometheus counters, gauges, or histograms to any component.
*   **Security & TLS**: Changing default security headers, TLS versions, or password hashing algorithms.
*   **Bundled Tools & Lifecycle**: Upgrading bundled versions of Helm, Kustomize, or Redis; shifting base container images or registry paths.
*   **Resource Management**: Altering resource tracking logic (labels vs. annotations) or changing default health assessment Lua scripts.

## Key Technical Concepts
*   **Resource Tracking**: Label, annotation, and `ServerSideApply` (SSA) methods.
*   **Sync Strategies**: Prune, IgnoreExtraneous, Sync Hooks/Waves, and Progressive Syncs.
*   **RBAC Syntax**: Policy structure: `p, <role>, <resource>, <action>, <object>, <effect>`.
*   **Generators**: Parameter sources for ApplicationSets (Git, SCM, Cluster, Matrix, etc.).
*   **Notification Elements**: Triggers (when), Templates (what), Subscriptions (who), and Services (how).
*   **Sharding**: Methods for distributing cluster load (legacy, round-robin, consistent-hashing).
*   **Health Statuses**: Healthy, Progressing, Degraded, Suspended, Missing, Unknown.
*   **CMP Lifecycle**: Plugin phases (discover, init, generate).
*   **Hydration**: Rendering manifests into dedicated repository subdirectories.
*   **Versioning**: Semantic versioning rules regarding breaking changes in minor/major releases.

## Related Components
*   **argocd-server**: API gateway, RBAC enforcer, and UI host.
*   **argocd-repo-server**: Manifest generation service (Helm, Kustomize).
*   **argocd-application-controller**: Primary state reconciliation engine.
*   **argocd-applicationset-controller**: Automation factory for generating Applications.
*   **argocd-notifications-controller**: Outbound alert and event processing engine.
*   **argocd-dex-server**: Identity federation service for OIDC/SAML.
*   **Redis / Sentinel**: Caching layer for manifest state and cluster coordination.