This documentation file, `declarative-setup.md`, serves as the definitive guide for managing Argo CD itself through Kubernetes manifests. It promotes a "GitOps for Argo CD" approach, where settings, applications, and cluster connections are defined as code rather than via CLI or UI.

### 1. Primary Purpose
The file documents the **declarative configuration** of Argo CD. It explains how to define and manage Argo CD Applications, Projects, Settings, Repositories, and Clusters using Kubernetes Custom Resource Definitions (CRDs), Secrets, and ConfigMaps.

### 2. Key Topics Covered
*   **Core Resource Management**: Usage of `Application` and `AppProject` CRDs.
*   **System Configuration**: Overview of standard ConfigMaps (e.g., `argocd-cm`, `argocd-rbac-cm`, `argocd-cmd-params-cm`) and Secrets (e.g., `argocd-secret`).
*   **Repository Connectivity**: Configuring Git and Helm repositories via Secrets using HTTPS, SSH, GitHub Apps, or OCI registries.
*   **Credential Templates**: Setting up shared credentials for multiple repositories based on URL prefixes.
*   **Multi-Cluster Management**: Defining remote clusters as Secrets, including cloud-specific authentication.
*   **Cloud Provider Integration**: Deep-dives into EKS (IRSA, Access Entries), GKE (Workload Identity), and AKS (kubelogin/federated identity).
*   **Resource Control**: Inclusion/Exclusion of specific Kubernetes resources from Argo CD tracking and management.
*   **Self-Management**: Patterns for using Argo CD to manage its own installation and configuration.

### 3. Technical Keywords
*   **CRDs**: `Application`, `AppProject`.
*   **Standard ConfigMaps**: `argocd-cm`, `argocd-rbac-cm`, `argocd-tls-certs-cm`, `argocd-ssh-known-hosts-cm`, `argocd-cmd-params-cm`.
*   **Auth/Identity**: `IRSA` (AWS), `Workload Identity` (GCP/Azure), `argocd-k8s-auth`, `execProviderConfig`, `AAD_LOGIN_METHOD`, `JWT`.
*   **Labels/Annotations**: `argocd.argoproj.io/secret-type`, `app.kubernetes.io/part-of: argocd`, `resources-finalizer.argocd.argoproj.io`.
*   **Features**: `App of Apps`, `Resource Exclusion/Inclusion`, `resource.respectRBAC`, `OCI support`.

### 4. Target Audience
*   **DevOps/Platform Engineers**: Responsible for the initial setup and maintenance of Argo CD.
*   **SREs**: Looking to automate infrastructure management and cluster bootstrapping.
*   **Security Architects**: Defining RBAC policies and secure credential storage patterns.

### 5. Related Concepts
*   **GitOps**: The underlying philosophy of using Git as the source of truth for infrastructure.
*   **Cluster Bootstrapping**: Using the "App of Apps" pattern to provision multiple services on new clusters.
*   **Kubernetes RBAC**: How Argo CD interacts with cluster-level security.
*   **Secret Management**: Integration with tools like Bitnami Sealed Secrets for encrypted manifest storage.

---

### AI Update Trigger Analysis
This file should be updated by an AI system if any of the following code changes occur:

1.  **CRD Schema Changes**: If new fields are added to the `Application` or `AppProject` spec (e.g., new source types, destination fields, or project-level restrictions).
2.  **Controller Logic Updates**: If the behavior of resource inclusion/exclusion or "Respect RBAC" logic is modified in the source code.
3.  **Authentication Providers**: If `argocd-k8s-auth` adds support for new cloud providers or if the configuration schema for existing ones (AWS, GCP, Azure) changes.
4.  **New Configuration Keys**: If new keys are added to `argocd-cm` (e.g., new masking options, custom label settings, or performance tuning variables).
5.  **Standard Resource Names**: If the required names for core ConfigMaps or Secrets are changed or if new mandatory labels/annotations are introduced.
6.  **Dependency Changes**: If a change in a supported tool (like Helm or Kustomize) requires a change in how Argo CD declaratively handles those sources (e.g., the OCI syntax changes).