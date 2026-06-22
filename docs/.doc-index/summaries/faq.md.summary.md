This analysis provides a comprehensive overview of the `faq.md` file from the Argo CD documentation to help users and AI systems understand its content and maintenance requirements.

---

### 1. Primary Purpose
The `faq.md` file serves as a **troubleshooting and configuration guide** for common issues encountered by Argo CD users. It provides practical workarounds, CLI commands, and configuration snippets to resolve state mismatches, authentication hurdles, and integration challenges with Kubernetes resources and Helm.

### 2. Key Topics Covered
*   **Application Lifecycle Management**: Handling deleted/corrupted repositories, terminating stuck syncs, and managing the `OutOfSync` status.
*   **Health Assessments**: Troubleshooting why specific resources (Ingress, StatefulSet, SealedSecrets) stay in a `Progressing` state.
*   **Identity and Access Management**: Resetting admin passwords (v1.8 vs. v1.9+), disabling the admin user, SSO/Dex configuration, and handling JWT cookie size limits.
*   **Helm Integration**: Explaining why Argo CD doesn't use `helm install`, how to deploy Helm charts in air-gapped environments, and managing Helm dependencies.
*   **Cluster & Networking**: Connectivity troubleshooting, gRPC-web proxies, TLS certificate issues, and cluster secret labels.
*   **Internal Mechanics**: Adjusting polling intervals (reconciliation timeouts), understanding Kubernetes resource normalization (diffing), and managing Redis (authentication, secret rotation).
*   **Technical Errors**: Solving "Manifest generation error," "field not declared in schema," and JSON patch order conflicts.

### 3. Technical Keywords
*   **ConfigMaps/Secrets**: `argocd-cm`, `argocd-secret`, `argocd-initial-admin-secret`, `argocd-redis`.
*   **CLI Commands**: `argocd account bcrypt`, `argocd admin cluster kubeconfig`, `argocd cluster list`, `kubectl patch`, `kubectl rollout restart`.
*   **Configuration Keys**: `timeout.reconciliation`, `admin.enabled`, `application.instanceLabelKey`, `resource.customizations.health`, `status.loadBalancer.ingress`.
*   **Flags**: `--cascade=false`, `--grpc-web`, `--insecure`, `--update-status`.
*   **Core Concepts**: Bcrypt hashing, JWT, gRPC, Server-side apply, Diffing strategies, Health checks.

### 4. Target Audience
*   **DevOps/SREs**: Responsible for installing, configuring, and maintaining the Argo CD instance.
*   **Kubernetes Administrators**: Managing the clusters Argo CD connects to and troubleshooting resource-level issues (e.g., Ingress status).
*   **App Developers**: Using Argo CD for GitOps who need to understand why their applications are "Out of Sync" or stuck "Progressing."

### 5. Related Concepts
*   **GitOps**: The underlying philosophy of using Git as the source of truth for cluster state.
*   **Kubernetes Controllers**: How Argo CD monitors resource states and calculates health.
*   **Helm**: The specific manifest generation tool behavior within Argo CD.
*   **Redis**: Used by Argo CD for caching and state management.
*   **RBAC & SSO**: User management and authentication flows via Dex.

---

### Maintenance Guide: When to Update this File
This file should be updated whenever the following code or architectural changes occur:

1.  **Authentication Changes**: If the default admin password logic changes or a new secret is introduced (as seen in the v1.8 to v1.9 transition).
2.  **Health Check Logic**: If the built-in health assessment logic for standard K8s types (like `StatefulSet` or `Ingress`) is modified or if new "known issues" arise with newer Kubernetes versions.
3.  **ConfigMap Schema Updates**: If new parameters are added to `argocd-cm` or if existing keys (like `timeout.reconciliation`) are deprecated.
4.  **CLI Flag Updates**: If flags like `--grpc-web` or `--insecure` are changed or if new troubleshooting commands are added to the `argocd` CLI.
5.  **Dependency Versioning**: If Argo CD is upgraded to a newer Kubernetes library version, the "field not declared in schema" section may need updated version references (e.g., updating the `go.mod` search instructions).
6.  **Component Architecture**: If the internal usage of Redis or Helm is significantly altered (e.g., moving away from `helm template` or changing the default Redis auth implementation).
7.  **Resource Handling**: If the way Argo CD labels resources (like `app.kubernetes.io/instance`) changes.