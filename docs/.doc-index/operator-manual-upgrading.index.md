# OPERATOR-MANUAL/UPGRADING Documentation Index

## Overview
This documentation provides a comprehensive guide for upgrading Argo CD, outlining the versioning strategy and specific migration steps required for various releases. It details breaking changes, deprecated features, and required manual interventions to ensure cluster stability during minor and major version transitions.

## Files Summary
*   **operator-manual/upgrading/overview.md**: Provides the fundamental rules for Argo CD versioning (SemVer), standard upgrade commands for HA and non-HA installs, and a master list of version-to-version upgrade guides.
*   **operator-manual/upgrading/1.0-1.1.md**: Highlights that v1.1 introduced no breaking changes but marked Kustomize v1.0 as deprecated.
*   **operator-manual/upgrading/1.1-1.2.md**: Documents the final removal of Kustomize v1 support and the requirement to annotate ConfigMaps with the `part-of: argocd` label.
*   **operator-manual/upgrading/1.2-1.3.md**: Alerts users to backward incompatible changes in the public Argo CD APIs requiring a CLI upgrade.
*   **operator-manual/upgrading/1.3-1.4.md**: Details changes to Sync Hook behavior, specifically regarding the deletion of in-flight hooks during sync termination.
*   **operator-manual/upgrading/1.4-1.5.md**: Covers Prometheus metric deprecations, the introduction of Redis HA Proxy, and the Kustomize v3.6.1 upgrade.
*   **operator-manual/upgrading/1.5-1.6.md**: Documents the removal of the deprecated `diff` field from the managed-resources API.
*   **operator-manual/upgrading/1.6-1.7.md**: Explains the migration of AppProject tokens to the status field and YAML library changes affecting multi-line strings.
*   **operator-manual/upgrading/1.7-1.8.md**: Covers the conversion of the application controller to a StatefulSet and the disabling of gRPC metrics by default.
*   **operator-manual/upgrading/1.8-2.0.md**: A major release guide covering the move to Helm v3 as default, Ubuntu base image migration, and the move to the quay.io registry.
*   **operator-manual/upgrading/2.0-2.1.md**: Details the replacement of the `--app-resync` flag, moving repository configs to Secrets, and the merger of `argocd-util` into `argocd admin`.
*   **operator-manual/upgrading/2.1-2.2.md**: Focuses on Helm 3.7 breaking changes and the removal of support for SHA-1 SSH key signatures.
*   **operator-manual/upgrading/2.2-2.3.md**: Documents the bundling of Notifications and ApplicationSets into the core Argo CD installation and the removal of Python from the base image.
*   **operator-manual/upgrading/2.3-2.4.md**: Covers the removal of KSonnet and Helm 2, new RBAC resources for `exec`, and the enforcement of logs RBAC.
*   **operator-manual/upgrading/2.4-2.5.md**: Details the deprecation of `argocd-cm` plugins (CMPs), Dex TLS configuration, and the blocking of out-of-bounds symlinks.
*   **operator-manual/upgrading/2.5-2.6.md**: Addresses ApplicationSet Go templating changes, OIDC `aud` claim requirements, and job health status updates.
*   **operator-manual/upgrading/2.6-2.7.md**: Introduces Proxy Extensions RBAC, the transition to Tini as entrypoint, and Deep Links template syntax updates.
*   **operator-manual/upgrading/2.7-2.8.md**: Documents the total removal of `argocd-cm` plugins and new RBAC actions for Job/Workflow creation.
*   **operator-manual/upgrading/2.8-2.9.md**: Details Kustomize upgrades and critical Egress NetworkPolicy changes for Redis.
*   **operator-manual/upgrading/2.9-2.10.md**: Covers the impact of kubectl upgrades on `managedNamespaceMetadata` and Helm version bumps.
*   **operator-manual/upgrading/2.10-2.11.md**: Notes the addition of the `initiatedBy` field in the Application CRD.
*   **operator-manual/upgrading/2.11-2.12.md**: Explains cluster secret scoping changes for projects and atomic field management for ApplicationSet selectors.
*   **operator-manual/upgrading/2.12-2.13.md**: Introduces health checks and actions for Flux resources and standardizes log file extensions to `.log`.
*   **operator-manual/upgrading/2.13-2.14.md**: Advises on avoiding v2.14.0 manifests due to image errors and covers Helm schema validation flags.
*   **operator-manual/upgrading/2.14-3.0.md**: A major upgrade guide covering fine-grained RBAC, Dex claim changes, annotation-based tracking, and resource exclusion defaults.
*   **operator-manual/upgrading/3.0-3.1.md**: Documents the deprecation of the v1 Actions API and the server-side enforcement of OIDC PKCE flows.
*   **operator-manual/upgrading/3.1-3.2.md**: Details the requirement for non-root hydration paths and support for Kustomize versions via `.argocd-source.yaml`.

## Code Changes That Would Require Documentation Updates
*   **Binary Dependency Updates**: Upgrading bundled versions of Helm, Kustomize, Redis, or HAProxy.
*   **CRD Schema Modifications**: Changes to `Application`, `AppProject`, or `ApplicationSet` specifications or status fields.
*   **RBAC Changes**: Adding new resources (e.g., `extensions`, `exec`, `logs`) or new actions (e.g., `action/batch/CronJob/create-job`) to the RBAC system.
*   **API Deprecations**: Modifying, versioning (v1 to v2), or removing GRPC or REST API endpoints.
*   **Default Configuration Shifts**: Changing default values in `argocd-cm`, `argocd-cmd-params-cm`, or `argocd-rbac-cm` (e.g., switching from label to annotation tracking).
*   **Security Policy Updates**: Changes to OIDC claim handling, TLS defaults, or NetworkPolicy restrictions for internal components.
*   **Container Architecture**: Altering the base Docker image, switching registries (e.g., Docker Hub to Quay), or changing entrypoint logic (e.g., `entrypoint.sh` to `tini`).
*   **Manifest Structural Changes**: Converting Deployments to StatefulSets or changing service account assignments.
*   **Third-party Integrations**: Adding or modifying support for external tools like Flux, Dex, or GnuPG.

## Key Technical Concepts
*   **SemVer (Semantic Versioning)**: The logic governing breaking changes vs. patch updates.
*   **RBAC Resources/Actions**: Specific strings like `applications`, `clusters`, `repositories`, `exec`, and `logs`.
*   **Config Management Plugins (CMP)**: The transition from `argocd-cm` plugins to Sidecar containers.
*   **Resource Tracking**: Methods including `label`, `annotation`, and `ApplyOutOfSyncOnly`.
*   **OIDC Claims**: `sub`, `aud`, `federated:id`, and `federated_claims.user_id`.
*   **Sync Hooks & Waves**: Logic for resource synchronization order and termination.
*   **Hydration**: The process of rendering manifests, specifically relating to path isolation.
*   **Server-Side Apply (SSA)**: Impact on field management and atomic map types.

## Related Components
*   **Argo CD Application Controller**: Responsible for state reconciliation and health assessments.
*   **Argo CD Repo Server**: Handles manifest generation (Helm, Kustomize, CMPs).
*   **Argo CD API Server**: Manages user authentication, RBAC enforcement, and API requests.
*   **Dex Server**: Handles identity provider integration and OIDC flows.
*   **Redis/Redis-HA**: Provides caching for the API and Repo servers.
*   **ApplicationSet Controller**: Manages multi-cluster/multi-app generation logic.
*   **Argo CD CLI**: The command-line interface used for administrative and user actions.