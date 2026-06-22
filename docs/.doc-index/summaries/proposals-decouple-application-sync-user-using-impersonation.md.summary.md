This document is a technical proposal for Argo CD that outlines the design and implementation for decoupling the Argo CD control plane's privileges from the privileges used to synchronize applications. It introduces **Kubernetes Impersonation** as the mechanism to achieve a more secure, least-privilege model in multi-tenant environments.

### 1. Primary Purpose
The primary purpose of this file is to document a feature that allows Argo CD to perform Application sync operations as a specific Kubernetes **ServiceAccount** rather than using the high-privilege identity of the Argo CD Application Controller. This reduces the security risk ("blast radius") if the Argo CD control plane is compromised, as sync operations are restricted by native Kubernetes RBAC rather than just Argo CD's internal logic.

### 2. Key Topics Covered
*   **The Impersonation Mechanism**: Using Kubernetes impersonation headers (`kubectl --as`) to execute syncs.
*   **AppProject CRD Extensions**: Introduction of the `destinationServiceAccounts` field within the `AppProject` specification to map clusters and namespaces to specific ServiceAccounts.
*   **Resolution Logic**: How Argo CD determines which ServiceAccount to use based on the Application’s destination (server and namespace), including glob pattern matching and "first-match" priority.
*   **Configuration & Enabling**: Methods to enable the feature via ConfigMaps, environment variables, or CLI flags.
*   **Component Impact**: Detailed implementation requirements for the GitOps Engine, Argo CD API, Application Controller, UI, and CLI.
*   **Security & Multi-tenancy**: Handling audit trails, preventing privilege escalation, and managing cross-namespace ServiceAccount references.

### 3. Technical Keywords
*   **APIs/CRDs**: `AppProject`, `Application`, `DestinationServiceAccount`, `SyncContext`.
*   **Configuration Options**: 
    *   `applicationcontroller.enable.impersonation` (in `argocd-cm`)
    *   `ARGOCD_APPLICATION_CONTROLLER_ENABLE_IMPERSONATION` (Env var)
    *   `--enable-impersonation` (CLI flag)
*   **Fields**: `destinationServiceAccounts`, `defaultServiceAccount`, `server`, `namespace`.
*   **Kubernetes Concepts**: Impersonation headers, ServiceAccounts, RBAC (Roles/RoleBindings), Cluster-admin, Audit logs.
*   **Argo Components**: Application Controller, GitOps Engine, `argocd-cm`.

### 4. Target Audience
*   **Argo CD Maintainers/Contributors**: To understand the architectural changes required across the codebase.
*   **Cluster Administrators**: To learn how to configure granular permissions for different teams or projects.
*   **Security Engineers**: To evaluate the security posture and auditability of the Argo CD installation.

### 5. Related Concepts
*   **Kubernetes Native RBAC**: This feature shifts authorization logic from Argo CD's `AppProject` (which acts as a secondary layer) to the cluster's native RBAC.
*   **Multi-tenancy**: Directly relates to how Argo CD isolates different teams (tenants) within a shared cluster.
*   **Principle of Least Privilege**: The underlying security philosophy driving the proposal.
*   **Kubernetes Audit Logs**: Impersonation events are recorded in K8s audit trails, providing visibility into which user/system triggered an action.

---

### AI Update Trigger Summary
This documentation should be updated or referenced if code changes occur in the following areas:
1.  **CRD Changes**: Any modification to `AppProject` or `Application` schemas, specifically regarding destinations or authentication.
2.  **Controller Logic**: Updates to how the Application Controller handles the `SyncContext` or how it invokes the GitOps Engine for `kubectl` operations.
3.  **GitOps Engine**: Changes to the underlying library that executes Kubernetes commands (ensuring impersonation headers are still passed).
4.  **Configuration Defaults**: If the feature flag status changes from Beta/Opt-in to enabled-by-default.
5.  **CLI/UI**: Changes to how `AppProjects` are created or viewed, specifically involving the `destinationServiceAccounts` mapping.