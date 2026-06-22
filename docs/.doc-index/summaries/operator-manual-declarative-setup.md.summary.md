This analysis provides a comprehensive summary of `operator-manual/declarative-setup.md` for Argo CD.

### 1. Primary Purpose
The document serves as the definitive guide for managing Argo CD's entire state—including applications, projects, settings, and credentials—using **declarative Kubernetes manifests**. It explains how to bypass the `argocd` CLI in favor of `kubectl apply` or GitOps workflows, enabling "Argo CD to manage Argo CD."

### 2. Key Topics Covered
*   **Atomic Configuration Reference**: A catalog of the specific ConfigMaps and Secrets that control Argo CD (e.g., `argocd-cm`, `argocd-rbac-cm`).
*   **Application & Project Specifications**: Detailed breakdowns of the `Application` and `AppProject` Custom Resource Definitions (CRDs).
*   **Credential Management**: How to declaratively define Git repositories, Helm charts (including OCI), and SSH/TLS trust configurations using Kubernetes Secrets and ConfigMaps.
*   **Multi-Cluster Connectivity**: Procedures for registering external clusters, with specific deep dives into cloud-native authentication for **AWS EKS, Google GKE, and Azure AKS**.
*   **Controller Behavior Tuning**: Settings for excluding/including specific Kubernetes resources from being tracked and configuring how the controller respects RBAC.
*   **UI/UX Customization**: Masking sensitive annotations and defining custom labels for the web interface.

### 3. Technical Keywords
*   **CRDs**: `Application`, `AppProject`.
*   **Core ConfigMaps**: `argocd-cm`, `argocd-cmd-params-cm`, `argocd-rbac-cm`, `argocd-tls-certs-cm`, `argocd-ssh-known-hosts-cm`.
*   **Kubernetes Labels/Annotations**: `argocd.argoproj.io/secret-type` (cluster/repository), `app.kubernetes.io/part-of: argocd`, `resources-finalizer.argocd.argoproj.io`.
*   **Cloud Auth Tools**: `argocd-k8s-auth`, IRSA (AWS), Workload Identity (GCP/Azure), `kubelogin`.
*   **Configuration Keys**: `resource.exclusions`, `resource.inclusions`, `resource.respectRBAC`, `execProviderConfig`, `awsAuthConfig`.
*   **Repo Types**: `git`, `helm`, `oci`.

### 4. Target Audience
*   **DevOps/Platform Engineers**: Who need to automate the deployment and scaling of Argo CD.
*   **Kubernetes Administrators**: Responsible for security, RBAC, and multi-cluster networking.
*   **Security Architects**: Focusing on secret management, credential templates, and resource masking.

### 5. Related Concepts
*   **GitOps**: The underlying philosophy of the entire document.
*   **App of Apps Pattern**: Mentioned as the method for bootstrapping multiple applications.
*   **RBAC & SSO**: Directly links to user management and permissioning systems.
*   **Cluster Bootstrapping**: Relates to how new clusters are provisioned and instantly managed by Argo CD.

---

### 6. Update Triggers for AI Systems
An AI should update or flag this file for review if code changes occur in the following areas:

1.  **CRD Schema Changes**: Any modification to the `argoproj.io` API group, specifically fields within `ApplicationSpec` or `AppProjectSpec` (e.g., adding a new field to `source` or `destination`).
2.  **ConfigMap Key Addition**: When new functional toggles are added to `argocd-cm` or `argocd-cmd-params-cm` (e.g., new performance tuning or resource filtering options).
3.  **Authentication Logic**: Changes to how Argo CD handles Git/Helm credentials or cluster authentication (e.g., supporting a new cloud provider or a new `argocd-k8s-auth` command).
4.  **Secret Labeling Logic**: If the controller changes the labels it looks for to identify repository or cluster secrets.
5.  **Default Resource Handling**: If the list of "always excluded" Kubernetes resources (like `events.k8s.io`) is modified in the source code.
6.  **Dependency Updates**: Significant changes in how integrated tools like `helm`, `kustomize`, or `kubelogin` are invoked (e.g., changes to proxy support or environment variable handling).