# OPERATOR-MANUAL/UPGRADING Documentation Index

## Overview
This documentation provides a comprehensive guide for upgrading Argo CD across various versions. It outlines the project's adherence to semantic versioning (SemVer), provides standard installation commands for both High Availability (HA) and non-HA environments, and details specific breaking changes, API deprecations, and configuration migrations required for each release to ensure system stability and security.

## Files Summary
*   **overview.md**: The entry point for upgrades, defining versioning rules (Patch/Minor/Major) and providing the base `kubectl apply` commands for upgrading Argo CD.
*   **3.1-3.2.md**: Details the move to non-root hydration paths, support for Kustomize versions in git-overrides, and changes to CronJob health assessment.
*   **3.0-3.1.md**: Covers symlink protection in the API server, the deprecation of the V1 Actions API, and changes to OIDC PKCE flow handling.
*   **2.14-3.0.md**: A major release guide highlighting fine-grained RBAC for sub-resources, mandatory logs RBAC, default resource exclusions, and the switch to annotation-based resource tracking.
*   **2.13-2.14.md**: Briefly notes a manifest fix for 2.14.0/2.14.1 and Helm version updates.
*   **2.12-2.13.md**: Covers Flux resource actions/health checks, LDAP connector strictness in Dex, and naming changes for manual CronJob triggers.
*   **2.11-2.12.md**: Details cluster secret scoping changes, Redis/HAProxy image registry shifts, and atomic field requirements for Server-Side Apply.
*   **2.10-2.11.md**: Documents the addition of `initiatedBy` in the Application CRD and Redis Egress policy changes.
*   **2.9-2.10.md**: Focuses on `managedNamespaceMetadata` behavior changes due to kubectl upgrades and Helm version bumps.
*   **2.8-2.9.md**: Covers Kustomize 5.2.1 upgrade and Egress NetworkPolicy modifications for Redis.
*   **2.7-2.8.md**: Documents the removal of `argocd-cm` plugin support, the switch to `tini` as entrypoint, and new RBAC actions for Jobs/Workflows.
*   **2.6-2.7.md**: Details the new `extensions` RBAC resource, Proxy Extensions, and changes to Deep Links templating syntax.
*   **2.5-2.6.md**: Covers the `project` vs `projects` filter bug, Sprig v3/Semver v3 upgrades, and the requirement for `aud` claims in API tokens.
*   **2.4-2.5.md**: Significant update covering `applicationsets` RBAC, deprecation of `argocd-cm` plugins, Dex TLS defaults, and out-of-bounds symlink blocking.
*   **2.3-2.4.md**: Documents the removal of KSonnet and Helm 2, new `exec` RBAC resources, and the requirement for `ARGOCD_ENV_` prefixes in plugin variables.
*   **2.2-2.3.md**: Highlights the bundling of Notifications and ApplicationSet into the core, removal of non-Linux binaries from images, and removal of Python from the base image.
*   **2.1-2.2.md**: Focuses on Helm 3.7+ (credential passing changes) and the removal of SHA-1 SSH signature support (OpenSSH 8.8 upgrade).
*   **2.0-2.1.md**: Covers the deprecation of `--app-resync`, the move from `argocd-cm` to Secrets for repo config, and the merge of `argocd-util` into `argocd admin`.
*   **1.8-2.0.md**: A major transition file covering the move to Ubuntu base images, Quay.io registry shift, and making Helm v3 the default renderer.
*   **1.7-1.8.md**: Documents the conversion of the controller to a StatefulSet and removal of default health assessments for Application CRDs.
*   **1.6-1.7.md**: Covers AppProject token migration, YAML library upgrades affecting multiline strings, and GnuPG RBAC rules.
*   **1.5-1.6.md**: Notes the removal of the deprecated `diff` field in the managed-resources API.
*   **1.4-1.5.md**: Covers Prometheus metric deprecations and the introduction of Redis HA Proxy.
*   **1.3-1.4.md**: Details changes to Sync Hook termination behavior and public API breaking changes.
*   **1.2-1.3.md**: Briefly notes API breaking changes.
*   **1.1-1.2.md**: Covers the removal of Kustomize v1 and new ConfigMap label requirements.
*   **1.0-1.1.md**: Notes the deprecation of Kustomize v1.0.

## Code Changes That Would Require Documentation Updates
*   **Bundled Tool Upgrades**: Updating versions of Helm, Kustomize, Redis, Dex, HAProxy, or OpenSSH.
*   **CRD Schema Modifications**: Adding, removing, or renaming fields in `Application`, `AppProject`, or `ApplicationSet` (e.g., the `project` to `projects` filter change).
*   **RBAC Policy Updates**: Introducing new resources (e.g., `applicationsets`, `exec`, `extensions`) or new actions (e.g., `action/batch/CronJob/create-job`).
*   **API Deprecations/Removals**: Changing REST or gRPC endpoints, or altering the structure of API responses (e.g., sanitizing sensitive fields).
*   **Default Behavior Shifts**: Changing system defaults like resource tracking methods (Labels vs. Annotations), health check logic, or sync hook assessment.
*   **Security Hardening**: Implementation of new validations (e.g., out-of-bounds symlink checks, `aud` claim requirements, or NetworkPolicy restrictions).
*   **Metric Changes**: Renaming, removing, or adding labels to Prometheus metrics (e.g., `argocd_app_info`).
*   **Container/Manifest Changes**: Changing the base OS image (e.g., Debian to Ubuntu), changing the entrypoint (e.g., `tini`), or moving components between controllers (e.g., Deployment to StatefulSet).
*   **Plugin Architecture**: Modifying how Config Management Plugins (CMPs) are loaded, environment variable prefixing, or filesystem access/isolation.

## Key Technical Concepts
*   **Semantic Versioning (SemVer)**: Rules governing patch, minor, and major release breaking changes.
*   **Config Management Plugins (CMP)**: Transition from `argocd-cm` to sidecar containers.
*   **Resource Tracking**: The mechanism (Label or Annotation) used to link live cluster resources to an Argo CD Application.
*   **RBAC Resources/Actions**: Specific permissions required for new features like UI Terminal (`exec`) or Proxy Extensions.
*   **OIDC/Dex**: Configuration for authentication providers and specific claims like `aud` and `federated_claims`.
*   **Sync Hooks & Waves**: Logic for resource application ordering and cleanup.
*   **Server-Side Apply (SSA)**: Impact on field management and atomic maps in ApplicationSets.
*   **Out-of-bounds Symlinks**: Security checks preventing path traversal in Git/Helm sources.
*   **Hydration**: The process of rendering manifests into a dedicated path for GitOps.

## Related Components
*   **argocd-server**: API handling, RBAC enforcement, and UI serving.
*   **argocd-repo-server**: Manifest generation (Helm/Kustomize/CMP) and repository access.
*   **argocd-application-controller**: Reconciler responsible for sync and health status.
*   **argocd-applicationset-controller**: Generator for multi-cluster/multi-app patterns.
*   **argocd-notifications**: System for alerting on application state changes.
*   **argocd-dex-server**: The bundled identity service for OIDC/LDAP.
*   **argocd-redis**: The caching layer (including HA Proxy configurations).