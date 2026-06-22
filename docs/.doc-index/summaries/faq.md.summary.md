This documentation file, `faq.md`, serves as the primary troubleshooting guide and "Frequently Asked Questions" repository for **Argo CD**. It addresses common operational hurdles, configuration nuances, and known limitations of the system.

### 1. Primary Purpose
The file provides actionable solutions for common issues encountered by administrators and users during the lifecycle of an Argo CD application—from initial deployment and authentication to complex sync errors and manifest generation failures.

### 2. Key Topics Covered
*   **Application Lifecycle Management**: Handling corrupted repositories, deleting applications using `--cascade=false`, and terminating stuck synchronizations.
*   **Health and Sync Status**: Troubleshooting why apps stay `OutOfSync` or stuck in `Progressing` (specifically for `Ingress`, `StatefulSet`, and `SealedSecret`).
*   **Authentication and User Management**: Resetting the admin password (v1.8 vs v1.9+), disabling the admin account, Dex/SSO configuration, and handling JWT cookie size limits.
*   **Helm Integration**: Explaining why Argo-deployed apps don't appear in `helm ls` and managing Helm dependencies in air-gapped environments.
*   **Cluster & Connectivity**: Troubleshooting cluster connection issues via CLI, verifying cluster secrets, and handling gRPC/TLS errors.
*   **Configuration & Performance**: Adjusting Git/Helm polling intervals (`timeout.reconciliation`) and managing the `app.kubernetes.io/instance` label.
*   **Backend Infrastructure**: Rotating Redis secrets and configuring/disabling Redis authentication.
*   **Manifest Generation**: Handling "field not declared in schema" errors (version mismatches) and cached generation errors.

### 3. Technical Keywords
*   **ConfigMaps/Secrets**: `argocd-cm`, `argocd-secret`, `argocd-initial-admin-secret`, `argocd-redis`, `argocd-redis-ha-haproxy`.
*   **CLI Commands**: `argocd account bcrypt`, `argocd admin cluster kubeconfig`, `kubectl patch secret`, `argocd cluster list`.
*   **Configuration Keys**: `timeout.reconciliation`, `admin.enabled`, `application.instanceLabelKey`, `resource.customizations.health`, `admin.passwordMtime`.
*   **Technical Specs**: `bcrypt`, `JWT`, `gRPC-web`, `x509`, `HTTP 2`, `Server-side apply`.
*   **Kubernetes Fields**: `status.loadBalancer.ingress`, `status.updatedReplicas`, `app.kubernetes.io/instance`, `$setElementOrder`.

### 4. Target Audience
*   **Kubernetes Administrators/Platform Engineers**: Responsible for installing, securing, and maintaining Argo CD.
*   **DevOps Engineers**: Who manage application deployments and need to understand sync behaviors and manifest diffs.
*   **Security Teams**: Interested in password rotation, SSO/Dex configuration, and Redis hardening.

### 5. Related Concepts
*   **GitOps**: The underlying methodology for Argo CD.
*   **Helm & Kustomize**: Supported manifest generation tools.
*   **Kubernetes Controller Pattern**: How Argo CD monitors resources and determines health.
*   **RBAC & SSO**: Integration with external identity providers via Dex.
*   **Redis**: Used as the caching layer for manifests and cluster state.

---

### AI Update Triggers
An AI system should recommend updating this file if code changes occur in the following areas:
*   **Security Defaults**: If the default admin password logic or secret naming conventions change.
*   **Health Assessment Logic**: If the hardcoded health checks for `Ingress`, `StatefulSet`, or `DaemonSet` are modified in the controller code.
*   **Helm/Kustomize Integration**: If the way Argo CD invokes manifest generators changes (e.g., moving away from `helm template`).
*   **Configuration Schema**: If new fields are added to `argocd-cm` or if the polling/reconciliation logic is refactored.
*   **CLI Flags**: If global flags like `--grpc-web` or `--insecure` are deprecated or their behavior changes.
*   **Dependency Upgrades**: When the Kubernetes library version (`go.mod`) is updated, as this affects the "field not declared in schema" section.
*   **Redis Integration**: If the internal Redis architecture or authentication requirements are modified.