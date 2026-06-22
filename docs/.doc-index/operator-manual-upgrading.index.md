# OPERATOR-MANUAL/UPGRADING Documentation Index

## Overview
This documentation area provides a chronological and version-specific roadmap for upgrading Argo CD. It details breaking changes, deprecated features, required manual migrations, security hardening measures, and updates to bundled third-party tools (like Helm and Kustomize) to ensure a safe transition between versions.

## Files Summary
*   **overview.md**: Defines Argo CD's semantic versioning rules (patch, minor, major) and provides the base installation commands for HA and non-HA environments.
*   **3.1-3.2.md**: Enforces non-root hydration paths, introduces `.argocd-source.yaml` support for Kustomize versions, and updates CronJob health status logic.
*   **3.0-3.1.md**: Hardens API security with symlink protection in static assets and migrates OIDC PKCE flow handling from the UI to the server.
*   **2.14-3.0.md**: A major release guide covering fine-grained RBAC for sub-resources, mandatory log RBAC, the switch to annotation-based resource tracking, and default resource exclusions for performance.
*   **2.13-2.14.md**: Addresses manifest tagging issues and documents Helm 3.16 upgrades along with project API sanitization.
*   **2.12-2.13.md**: Adds native support for Flux resource actions/health and documents breaking changes in the Dex LDAP connector.
*   **2.11-2.12.md**: Details strict cluster secret scoping by project and the migration of ApplicationSet field management to atomic server-side apply.
*   **2.10-2.11.md**: Adds the `initiatedBy` field to the Application CRD and continues Redis NetworkPolicy egress updates.
*   **2.9-2.10.md**: Covers Helm 3.14 upgrades and the impact of kubectl 1.26 on `managedNamespaceMetadata` labels/annotations.
*   **2.8-2.9.md**: Documents Kustomize 5.2 upgrades and security adjustments to Redis Egress NetworkPolicies for Kubernetes API access.
*   **2.7-2.8.md**: Marks the final removal of `argocd-cm` based plugins and introduces new RBAC actions for Job and Workflow creation.
*   **2.6-2.7.md**: Introduces Proxy Extensions RBAC, the transition to `tini` as a container entrypoint, and deep link template variable updates.
*   **2.5-2.6.md**: Documents breaking changes in ApplicationSet semver functions, OIDC `aud` claim requirements, and the "Suspended" status for Jobs.
*   **2.4-2.5.md**: Deprecates `argocd-cm` plugins in favor of sidecars, introduces the `applicationsets` RBAC resource, and implements out-of-bounds symlink blocking.
*   **2.3-2.4.md**: Final removal of KSonnet and Helm 2, introduction of the `exec` RBAC resource, and prefixing of plugin environment variables (`ARGOCD_ENV_`).
*   **2.2-2.3.md**: Bundles Notifications and ApplicationSet into the core project and removes Python from the base image.
*   **2.1-2.2.md**: Upgrades Helm to 3.7 (rewritten OCI support) and documents the removal of SHA-1 SSH key signature support.
*   **2.0-2.1.md**: Deprecates `--app-resync` in favor of `timeout.reconciliation` and migrates repository configurations from `argocd-cm` to Secrets.
*   **1.8-2.0.md**: Major update switching to Ubuntu base images, Quay.io as the primary registry, and Helm 3 as the default rendering engine.
*   **1.7-1.8.md**: Converts the application controller to a StatefulSet and removes default health assessment for `Application` CRDs.
*   **1.6-1.7.md**: Migrates AppProject tokens to the `status` field and updates RBAC for GnuPG key management.
*   **1.5-1.6.md**: Removes the deprecated `diff` field from the managed-resources API.
*   **1.4-1.5.md**: Introduces Redis HA Proxy manifests and migrates specific Prometheus metrics to labels on `argocd_app_info`.
*   **1.3-1.4.md**: Implements stricter sync hook termination behavior and public API breaking changes.
*   **1.2-1.3.md**: Notes public API breaking changes and mandates a CLI upgrade to v1.3.
*   **1.1-1.2.md**: Removes Kustomize v1 support and requires specific `part-of` labels on ConfigMaps.
*   **1.0-1.1.md**: Notes Kustomize v1 deprecation and stable v1.1 upgrade path.

## Code Changes That Would Require Documentation Updates
*   **Tool Version Bumps**: Updating bundled versions of Helm, Kustomize, Redis, HAProxy, Dex, or OpenSSH.
*   **CRD Schema Modifications**: Adding, renaming, or removing fields in `Application`, `ApplicationSet`, or `AppProject` manifests.
*   **RBAC Resource/Action Additions**: Defining new resources (like `extensions` or `exec`) or new actions (like `action/batch/CronJob/create-job`).
*   **API Response Changes**: Sanitizing fields in API responses, changing error codes (e.g., `NotFound` to `PermissionDenied`), or deprecating endpoints.
*   **Configuration Defaults**: Changing default values in `argocd-cm`, `argocd-rbac-cm`, or `argocd-cmd-params-cm`.
*   **Controller Architecture**: Changing deployments to StatefulSets or modifying Service Accounts.
*   **Security Logic**: Implementing new symlink checks, path traversal protections, or OIDC claim requirements.
*   **Resource Tracking/Diffing**: Modifying how Argo CD identifies managed resources (e.g., switching from labels to annotations) or how it calculates diffs.
*   **Base Image/Entrypoint**: Changing the Docker base image (e.g., Debian to Ubuntu) or the binary entrypoint logic (e.g., `entrypoint.sh` to `tini`).

## Key Technical Concepts
*   **Config Management**: `argocd-cm`, `argocd-cmd-params-cm`, Config Management Plugins (CMPs), Sidecar containers.
*   **RBAC Resources**: `applicationsets`, `extensions`, `exec`, `logs`, `gpgkeys`, `clusters`, `projects`.
*   **Resource Tracking**: `application.resourceTrackingMethod`, `argocd.argoproj.io/tracking-id`, `ApplyOutOfSyncOnly`.
*   **Sync Logic**: Sync Waves, Sync Hooks, Server-side Apply, `managedNamespaceMetadata`.
*   **Security**: OIDC `aud` claims, PKCE, Egress NetworkPolicies, Out-of-bounds symlinks, SSH SHA-1 deprecation.
*   **CLI Commands**: `argocd admin`, `argocd app diff --server-side-generate`, `argocd-util` (deprecated).
*   **External Generators**: Matrix, Merge, List, and Git generators in ApplicationSets.

## Related Components
*   **argocd-application-controller**: Manages application state and health.
*   **argocd-repo-server**: Handles manifest generation and repository communication.
*   **argocd-server**: Provides the API and UI.
*   **argocd-applicationset-controller**: Manages the generation of multiple applications.
*   **argocd-dex-server**: Handles OIDC and LDAP authentication.
*   **argocd-redis**: Provides caching for the repo-server and API.
*   **argocd-notifications**: (Bundled) Handles alerts and status updates.