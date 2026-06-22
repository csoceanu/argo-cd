This analysis provides a comprehensive summary of the `faq.md` documentation for Argo CD.

### 1. Primary Purpose
The file serves as a **troubleshooting guide and knowledge base** for Argo CD. It addresses common pitfalls, explains "by-design" behaviors that may be counter-intuitive to new users (especially regarding Helm and Kubernetes synchronization), and provides manual recovery procedures for broken states.

### 2. Key Topics Covered
*   **Application Management & Recovery**: Handling deleted/corrupted repositories, terminating stuck syncs, and deleting applications manually via `--cascade=false`.
*   **Sync & Diffing Logic**: Explanations for persistent `OutOfSync` states, unit normalization (e.g., `1000m` vs `1`), and conflicts with the `app.kubernetes.io/instance` label.
*   **Health Checks**: Troubleshooting the `Progressing` state for specific resources like `Ingress`, `StatefulSet`, and `SealedSecret`, including custom health check Lua scripts.
*   **Authentication & User Management**: Resetting the admin password, using `bcrypt`, disabling the admin account, and troubleshooting SSO (Dex) and JWT cookie size limits.
*   **Helm Integration**: Why Argo CD apps don't show up in `helm list` (use of `helm template`) and managing Helm dependencies in air-gapped environments.
*   **Cluster & Infrastructure**: Troubleshooting cluster connectivity, rotating Redis secrets, and configuring Redis authentication.
*   **CLI & Technical Errors**: Handling gRPC/proxy issues, X.509 certificate errors, and schema validation errors related to Server-Side Apply.

### 3. Technical Keywords
*   **Configuration Maps**: `argocd-cm` (main config), `argocd-secret` (passwords/keys), `argocd-redis` (auth).
*   **Config Keys**: `timeout.reconciliation`, `application.instanceLabelKey`, `admin.enabled`, `dex.config`, `resource.customizations.health`.
*   **CLI Commands/Flags**: `argocd account bcrypt`, `argocd admin cluster kubeconfig`, `--cascade=false`, `--grpc-web`, `--insecure`, `kubectl patch`.
*   **Labels**: `argocd.argoproj.io/secret-type: cluster`, `app.kubernetes.io/instance`.
*   **Technical Concepts**: JWT tokens, bcrypt hashing, Server-Side Apply (SSA), Server-Side Diff, Lua health scripts, Manifest generation caching.

### 4. Target Audience
*   **DevOps/Platform Engineers**: Responsible for installing, configuring, and maintaining Argo CD.
*   **Kubernetes Administrators**: Troubleshooting cluster connectivity and resource health.
*   **Application Developers**: Using Argo CD to deploy manifests and needing to understand why their apps are "Out of Sync" or "Progressing."

### 5. Related Concepts
*   **GitOps**: The foundational pattern this tool implements.
*   **Helm**: Specifically the distinction between `helm template` and `helm install`.
*   **Kubernetes Controllers**: Specifically how `StatefulSet` and `Ingress` controllers update status fields.
*   **Redis**: Used by Argo CD for caching and state management.
*   **Dex**: The identity provider connector used for SSO.

---

### AI Update Trigger Summary
An AI system should monitor and update this file when the following code or environmental changes occur:

1.  **CLI Flag Changes**: If arguments like `--cascade`, `--grpc-web`, or `--insecure` are deprecated or changed in the Argo CD CLI source code.
2.  **Default Secret Logic**: If the logic for generating the initial admin password or storing Redis credentials changes (e.g., moving from pod name to `argocd-initial-admin-secret`).
3.  **ConfigMap Schema Changes**: If new keys are added to `argocd-cm` or if existing keys (like `timeout.reconciliation`) are renamed.
4.  **Dependency Upgrades**:
    *   When the **Kubernetes client-go** version is updated (affects the "field not declared in schema" section).
    *   When **Redis** installation manifests are modified (affects secret rotation steps).
    *   When **Helm** integration logic changes how charts are rendered.
5.  **Health Check Defaults**: If Argo CD introduces built-in health checks for new resource types or fixes known issues with `StatefulSet` or `Ingress` health monitoring.
6.  **Labeling Logic**: If the default `instanceLabelKey` changes or if the way Argo CD tracks resources via labels is modified.