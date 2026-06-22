This documentation file, `declarative-setup.md`, serves as the foundational guide for managing **Argo CD** itself through Kubernetes manifests. It promotes the GitOps "self-hosting" philosophy by documenting how to configure Argo CD resources without relying on the command-line interface (CLI).

### 1. Primary Purpose
The file documents the **declarative configuration** of Argo CD. It provides the schema and examples for the Custom Resource Definitions (CRDs) and Kubernetes native resources (ConfigMaps/Secrets) required to manage applications, projects, clusters, and repository credentials within an Argo CD instance.

### 2. Key Topics Covered
*   **Atomic Configuration:** A reference of the specific ConfigMaps and Secrets used for core settings (RBAC, TLS, SSH hosts, global params).
*   **Application CRD:** How to define a deployment unit, including source (Git/Helm), destination (Cluster/Namespace), and cascading deletion via finalizers.
*   **AppProject CRD:** Logical grouping of applications, multi-tenancy, and security boundaries (whitelisting/blacklisting resources).
*   **Repository Management:** Connecting private Git/Helm repositories via HTTPS, SSH, GitHub Apps, or OCI registries.
*   **Cluster Management:** Declaratively adding target Kubernetes clusters, with deep dives into cloud-specific authentication (EKS, GKE, AKS).
*   **Resource Filtering:** Methods for excluding or including specific Kubernetes resource types from Argo CD’s management using `resource.exclusions` and `resource.inclusions`.
*   **Advanced Controller Settings:** Features like `resource.respectRBAC` (to align controller discovery with K8s RBAC) and masking sensitive annotations.

### 3. Technical Keywords
*   **CRDs:** `Application`, `AppProject` (`argoproj.io/v1alpha1`).
*   **ConfigMaps:** `argocd-cm`, `argocd-rbac-cm`, `argocd-tls-certs-cm`, `argocd-ssh-known-hosts-cm`, `argocd-cmd-params-cm`.
*   **Secrets Labels:** `argocd.argoproj.io/secret-type: repository`, `argocd.argoproj.io/secret-type: cluster`.
*   **Auth Mechanisms:** IRSA (IAM Roles for Service Accounts), Workload Identity (GCP/Azure), `argocd-k8s-auth`, `kubelogin`, JWT Tokens.
*   **Finalizers:** `resources-finalizer.argocd.argoproj.io`.
*   **Configuration Keys:** `resource.exclusions`, `resource.inclusions`, `resource.respectRBAC`, `enableOCI`.

### 4. Target Audience
*   **Platform Engineers & SREs:** Responsible for installing, hardening, and maintaining the Argo CD infrastructure.
*   **DevOps Engineers:** Automating the onboarding of new applications and clusters.
*   **Security Teams:** Auditing the RBAC, project boundaries, and credential management strategies.

### 5. Related Concepts
*   **GitOps:** The overarching methodology of using Git as the source of truth for infrastructure.
*   **App of Apps Pattern:** Using one Argo CD Application to manage a manifest containing other Applications.
*   **RBAC (Role-Based Access Control):** Specifically how `argocd-rbac-cm` and `AppProject` roles interact.
*   **Helm & Kustomize:** Supported manifest generation tools within the `Application` spec.

---

### AI Update Trigger Analysis
This file is a critical reference for the **Argo CD API surface and configuration schema**. An AI system should monitor for code changes in the following areas to determine if this documentation requires updates:

*   **CRD Golang Structs:** Any changes to `pkg/apis/application/v1alpha1/types.go` (e.g., adding fields to `ApplicationSpec` or `AppProjectSpec`) must be reflected here.
*   **Controller Logic:** If the `argocd-application-controller` introduces new behavior for resource tracking (like the `respectRBAC` feature), the "Resource Exclusion/Inclusion" section needs an update.
*   **Authentication Plugins:** Updates to `argocd-k8s-auth` or changes in how Argo CD interacts with cloud provider APIs (EKS Access Entries, Azure Workload Identity) require syncing the "Clusters" section.
*   **Default ConfigMap Keys:** Any new keys added to the core configuration logic (processed by `argocd-cm`) should be added to the "Quick Reference" table.
*   **Security/Finalizer Changes:** Modifications to how cascading deletes or namespace-scoped restrictions are handled by the controller logic.