This documentation file, `declarative-setup.md`, serves as the authoritative guide for managing **Argo CD as Code**. It explains how to define, configure, and manage Argo CD applications, projects, clusters, and global settings using Kubernetes manifests and `kubectl` instead of the command-line interface (CLI) or user interface (UI).

### 1. Primary Purpose
The file documents the **declarative management** of the Argo CD ecosystem. It provides the schema and examples for Custom Resource Definitions (CRDs) and specialized Kubernetes Secrets/ConfigMaps that control how Argo CD functions and what it manages.

### 2. Key Topics Covered
*   **Core CRDs**: Detailed specifications for `Application` and `AppProject` resources, including source/destination definitions and lifecycle management (finalizers).
*   **Infrastructure Configuration**: How to connect Git/Helm repositories and external Kubernetes clusters using Kubernetes Secrets.
*   **Authentication & Security**: 
    *   Repository credentials (HTTPS, SSH, GitHub Apps, GCP, Azure).
    *   Managing TLS certificates and SSH known hosts via ConfigMaps.
    *   Cloud-specific cluster authentication (AWS EKS/IRSA, Google GKE Workload Identity, Azure AKS/kubelogin).
*   **Global Settings**: Using `argocd-cm` to exclude/include specific Kubernetes resources from being tracked, masking sensitive data, and configuring RBAC respect.
*   **Operational Patterns**: The "App of Apps" pattern and "Managing Argo CD using Argo CD" (Self-management).

### 3. Technical Keywords
*   **CRDs**: `Application`, `AppProject`.
*   **ConfigMaps**: `argocd-cm`, `argocd-rbac-cm`, `argocd-tls-certs-cm`, `argocd-ssh-known-hosts-cm`, `argocd-cmd-params-cm`.
*   **Secrets Labels**: `argocd.argoproj.io/secret-type` (values: `repository`, `cluster`, `repo-creds`).
*   **Auth/K8s APIs**: `execProviderConfig`, `argocd-k8s-auth`, `client.authentication.k8s.io/v1beta1`, `resources-finalizer.argocd.argoproj.io`.
*   **Cloud Integrations**: IRSA (AWS), Workload Identity (GCP/Azure), `awsAuthConfig`, `AmazonEKSClusterAdminPolicy`.
*   **Configuration Keys**: `resource.exclusions`, `resource.inclusions`, `resource.respectRBAC`, `enableOCI`.

### 4. Target Audience
*   **Platform Engineers/SREs**: Responsible for installing and maintaining the Argo CD infrastructure.
*   **DevOps Engineers**: Implementing GitOps pipelines and automating the creation of applications and projects.
*   **Security Administrators**: Configuring RBAC, secret masking, and cluster-access policies.

### 5. Related Concepts
*   **GitOps**: The underlying philosophy of using Git as the single source of truth for infrastructure.
*   **Kubernetes RBAC**: Argo CD uses native K8s RBAC and its own internal RBAC (documented in `rbac.md`).
*   **Cloud Provider IAM**: Deeply integrated with AWS IAM, GCP IAM, and Azure Active Directory for cluster authentication.
*   **Helm & Kustomize**: Sourcing manifests from these template engines.

---

### 💡 Update Triggers for AI Systems
This file should be updated whenever code changes occur in the following areas:
1.  **CRD Schema Changes**: If new fields are added to or removed from the `Application` or `AppProject` Go structs (e.g., additions to `ApplicationSpec`).
2.  **Secret/ConfigMap Logic**: If the controller's logic for parsing `argocd-cm` or secret labels (like `secret-type: cluster`) is modified.
3.  **Authentication Providers**: If new authentication methods for Git providers (e.g., a new Git host) or Cloud providers (e.g., a new EKS auth flag) are introduced.
4.  **Resource Handling Logic**: If the way the controller filters resources (`resource.exclusions`) or respects RBAC (`resource.respectRBAC`) changes.
5.  **New Global Environment Variables**: If variables managed via `argocd-cmd-params-cm` are added or deprecated.