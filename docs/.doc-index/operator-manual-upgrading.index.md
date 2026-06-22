# OPERATOR-MANUAL/UPGRADING Documentation Index

## Overview
This documentation area provides comprehensive guidance for upgrading Argo CD between various versions. It details the semver-based versioning logic, provides standard installation commands for both High Availability (HA) and Non-HA environments, and lists critical breaking changes, deprecations, and mandatory manual interventions required during specific version migrations from v1.0 through v3.2.

## Files Summary
*   **operator-manual/upgrading/overview.md**: Serves as the landing page for upgrades, explaining versioning rules (patch, minor, major) and providing core `kubectl` upgrade commands and a master list of all version-specific guides.
*   **operator-manual/upgrading/3.1-3.2.md**: Details breaking changes regarding non-root hydration paths, Kustomize overrides in `.argocd-source.yaml`, CronJob health status transitions, and ApplicationSet status limits.
*   **operator-manual/upgrading/3.0-3.1.md**: Covers symlink protection for static assets, the deprecation of the v1 Actions API, and the transition of OIDC PKCE handling from the UI to the server.
*   **operator-manual/upgrading/2.14-3.0.md**: A major version guide covering fine-grained RBAC for sub-resources, mandatory logs RBAC enforcement, default resource exclusions for performance, Dex user ID changes, and the switch to annotation-based resource tracking.
*   **operator-manual/upgrading/2.13-2.14.md**: Highlights a manifest issue in v2.14.0, Helm 3.16 upgrades, and security-related sanitization of project API responses.
*   **operator-manual/upgrading/2.12-2.13.md**: Documents new Custom Resource Actions and Health checks for Flux resources, LDAP connector strictness in Dex, and changes to CronJob job naming conventions.
*   **operator-manual/upgrading/2.11-2.12.md**: Explains cluster secret scoping changes based on projects, the migration of Redis/HAProxy images to ECR, and atomic field requirements for Server-Side Apply in ApplicationSets.
*   **operator-manual/upgrading/2.10-2.11.md**: Introduces the `initiatedBy` field in the Application CRD and documents egress NetworkPolicy modifications for Redis.
*   **operator-manual/upgrading/2.9-2.10.md**: Details how `managedNamespaceMetadata` no longer preserves client-side labels due to a kubectl upgrade and repeats Redis egress NetworkPolicy instructions.
*   **operator-manual/upgrading/2.8-2.9.md**: Covers the Kustomize 5.2.1 upgrade and provides specific `diff` blocks for adjusting Redis and HAProxy NetworkPolicy egress rules.
*   **operator-manual/upgrading/2.7-2.8.md**: Documents the final removal of `argocd-cm` plugin support, the switch to `tini` as the container entrypoint, and new RBAC actions for Jobs and Workflows.
*   **operator-manual/upgrading/2.6-2.7.md**: Introduces the `extensions` RBAC resource, deep link template syntax updates, support for `helm.sh/resource-policy`, and the upgrade to Kustomize 5.0.
*   **operator-manual/upgrading/2.5-2.6.md**: Covers breaking API changes in project filters, Sprig semver function behavior changes, and the mandatory requirement for `aud` claims in OIDC tokens.
*   **operator-manual/upgrading/2.4-2.5.md**: Details the deprecation of `argocd-cm` plugins, the introduction of the `applicationsets` RBAC resource, Dex TLS enablement, and server-side manifest diffing.
*   **operator-manual/upgrading/2.3-2.4.md**: Documents the removal of KSonnet and Helm 2, the introduction of the `exec` RBAC resource, and mandatory prefixing (`ARGOCD_ENV_`) for plugin environment variables.
*   **operator-manual/upgrading/2.2-2.3.md**: Explains the bundling of Notifications and ApplicationSet into core Argo CD and the removal of non-Linux binaries from the default image.
*   **operator-manual/upgrading/2.1-2.2.md**: Focuses on Helm 3.7 breaking changes regarding repository credentials and the removal of SHA-1 SSH key signature support.
*   **operator-manual/upgrading/2.0-2.1.md**: Details the replacement of `--app-resync` with ConfigMap settings, the migration of repository configs to Secrets, and the merging of `argocd-util` into `argocd admin`.
*   **operator-manual/upgrading/1.8-2.0.md**: A major release guide covering the migration to Ubuntu-based images, registry change to Quay.io, Redis 6.2 upgrade, and the transition of CRDs to `apiextensions/v1`.
*   **operator-manual/upgrading/1.7-1.8.md**: Documents the conversion of the application controller to a StatefulSet and the removal of default health assessment for the Application CRD.
*   **operator-manual/upgrading/1.6-1.7.md**: Covers the migration of AppProject tokens to the status field and new RBAC requirements for GnuPG features.
*   **operator-manual/upgrading/1.5-1.6.md**: Details the removal of the deprecated `diff` field from the managed-resources API.
*   **operator-manual/upgrading/1.4-1.5.md**: Documents the introduction of Redis HA Proxy and the deprecation of legacy Prometheus metrics.
*   **operator-manual/upgrading/1.3-1.4.md**: Explains changes in sync hook state assessment and backward incompatible API changes.
*   **operator-manual/upgrading/1.2-1.3.md**: Briefly notes backward incompatible API changes and CLI requirements.
*   **operator-manual/upgrading/1.1-1.2.md**: Documents the removal of Kustomize v1 support and the requirement for specific labels on Argo CD ConfigMaps.
*   **operator-manual/upgrading/1.0-1.1.md**: Notes the deprecation of Kustomize v1.0.

## Code Changes That Would Require Documentation Updates
*   **Bundled Tool Upgrades**: Any update to the versions of Helm, Kustomize, Redis, Dex, or OpenSSH bundled in the manifests or base image.
*   **RBAC Schema Changes**: Adding new RBAC resources (e.g., `extensions`, `exec`, `applicationsets`) or actions that change existing permission inheritance.
*   **CRD Modifications**: Changes to `Application`, `AppProject`, or `ApplicationSet` schemas, including moving fields between `spec` and `status` or changing field types.
*   **ConfigMap Key Changes**: Deprecating, renaming, or adding new keys to `argocd-cm`, `argocd-rbac-cm`, or `argocd-cmd-params-cm`.
*   **API Breaking Changes**: Modifying gRPC or REST API responses, especially sanitizing fields or changing field names in JSON payloads (e.g., `project` to `projects`).
*   **Architecture/Deployment Changes**: Converting Deployments to StatefulSets, changing default container users (BYOI), or modifying entrypoint scripts (e.g., `tini`).
*   **Default Behavior Shifts**: Changes to resource tracking methods (Label vs. Annotation), health check logic, or sync hook assessment.
*   **Security Policy Updates**: New restrictions on symlinks, Egress NetworkPolicies, or OIDC claim requirements (e.g., mandatory `aud` or `federated:id` scopes).

## Key Technical Concepts
*   **Semantic Versioning (SemVer)**: Rules governing patch, minor, and major release behavior.
*   **Config Management Plugins (CMP)**: The transition from `argocd-cm` based plugins to sidecar-based plugins.
*   **Resource Tracking**: Methods used to link live resources to Argo Applications (`label`, `annotation`).
*   **RBAC Resources**: `applications`, `clusters`, `repositories`, `logs`, `exec`, `extensions`, `applicationsets`, `gpgkeys`.
*   **Sync Hooks & Waves**: Logic determining the order and health assessment of resource application.
*   **Server-Side Apply (SSA)**: Atomic field management and merge strategies for CRDs.
*   **OIDC/Dex Authentication**: Claims handling (`aud`, `sub`, `federated:id`), PKCE, and connector configurations (LDAP).
*   **HA vs. Non-HA Manifests**: Installation patterns and the use of Redis HA Proxy.

## Related Components
*   **argocd-server**: API server handling REST/gRPC, RBAC enforcement, and UI assets.
*   **argocd-repo-server**: Component responsible for manifest generation (Helm, Kustomize, Plugins).
*   **argocd-application-controller**: Controller managing application state, health, and reconciliation.
*   **argocd-dex-server**: Embedded Dex instance for identity provider integration.
*   **argocd-redis / argocd-redis-ha**: Caching layer for repository and application state.
*   **argocd-applicationset-controller**: Controller for generating applications from templates and generators.
*   **Argo CD CLI**: Command-line interface required to stay in sync with server versions.