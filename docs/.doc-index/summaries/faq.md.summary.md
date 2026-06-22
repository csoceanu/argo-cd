This analysis provides a comprehensive summary of the `faq.md` documentation for Argo CD.

### 1. Primary Purpose
The file serves as a **troubleshooting and configuration guide** for Argo CD. It addresses common pain points, architectural misunderstandings (specifically regarding Helm), and provides recovery steps for common failure scenarios (lost passwords, corrupted repos, and stuck syncs).

### 2. Key Topics Covered
*   **Application Lifecycle & Syncing**: Troubleshooting "stuck" applications in `Progressing` or `OutOfSync` states, terminating syncs, and handling deleted/corrupted repositories.
*   **Authentication & User Management**: Admin password resets (v1.8 vs. v1.9+ logic), disabling the admin user, and fixing SSO/Dex issues or cookie length errors.
*   **Helm Integration**: Explaining why Argo CD-deployed Helm apps don't appear in `helm ls` and how to handle air-gapped Helm repositories.
*   **Cluster & Infrastructure**: Troubleshooting cluster connectivity, manual `kubeconfig` generation, and cluster secret labels.
*   **Internal Component Management**: Rotating Redis secrets, configuring Redis authentication, and understanding repo-server manifest caching.
*   **Advanced Diffing & Validation**: Handling unit normalization (e.g., `1000m` vs `1`), "field not declared in schema" errors, and Kubernetes-specific resource health (Ingress, StatefulSets, SealedSecrets).

### 3. Technical Keywords
*   **ConfigMaps/Secrets**: `argocd-cm`, `argocd-secret`, `argocd-redis`, `argocd-initial-admin-secret`, `argocd-repo-server`.
*   **Configuration Keys**: `timeout.reconciliation`, `application.instanceLabelKey`, `admin.enabled`, `resource.customizations.health`, `dex.config`.
*   **Argo CD CLI Commands**: `argocd account bcrypt`, `argocd admin cluster kubeconfig`, `argocd cluster list`.
*   **CLI Flags**: `--cascade=false`, `--grpc-web`, `--insecure`.
*   **Kubernetes Concepts**: `CustomResourceDefinition (CRD)`, `Server-Side Apply`, `LoadBalancer status`, `bcrypt hash`, `JWT`.
*   **Labels**: `argocd.argoproj.io/secret-type: cluster`, `app.kubernetes.io/instance`.

### 4. Target Audience
*   **Argo CD Administrators/Operators**: Users responsible for installing, configuring, and maintaining the Argo CD instance.
*   **DevOps/SRE Engineers**: Users troubleshooting application deployment failures or sync discrepancies.
*   **Kubernetes Developers**: Users needing to understand why their manifests (Helm/Kustomize) behave differently in Argo CD than in local environments.

### 5. Related Concepts
*   **GitOps**: The underlying philosophy of using Git as the source of truth.
*   **Declarative Setup**: Automating the configuration of Argo CD itself.
*   **Health Checks**: The Lua-based customization engine used to determine resource status.
*   **Manifest Generation**: The process of converting templates (Helm/Kustomize) into raw YAML via the `repo-server`.
*   **RBAC & User Management**: Integrating with OIDC/Dex for enterprise identity.

---

### AI Update Triggers: When to update this file
An AI system should suggest updates to `faq.md` if code changes occur in the following areas:

1.  **Security/Auth Logic**: If the default admin password generation logic changes or the name of the initial secret (`argocd-initial-admin-secret`) is modified.
2.  **ConfigMap Schema**: If new fields are added to `argocd-cm` that affect sync behavior, polling intervals, or resource tracking.
3.  **Health Assessment Logic**: If the hard-coded health checks for standard types (Ingress, StatefulSet, etc.) are rewritten or if the default health status for a common controller (like `SealedSecret`) changes.
4.  **Component Architecture**: If a new internal component is introduced or if the Redis dependency is replaced/significantly altered (e.g., changing how Redis auth is handled).
5.  **CLI Interface**: If existing flags like `--grpc-web` are deprecated or if new troubleshooting flags are added.
6.  **Dependency Upgrades**: When the Go modules (`go.mod`) update the Kubernetes library versions, as this dictates the "field not declared in schema" troubleshooting section.
7.  **Resource Tracking**: If the default label used for tracking (currently `app.kubernetes.io/instance`) is changed in the source code.