This analysis provides a comprehensive overview of the `operator-manual/rbac.md` documentation for Argo CD.

### 1. Primary Purpose
The file serves as the definitive guide for configuring **Role-Based Access Control (RBAC)** in Argo CD. It explains how to restrict access to resources (like applications, clusters, and projects) for both local users and those authenticated via External Single Sign-On (SSO). It defines the syntax, hierarchy, and evaluation logic of the RBAC engine.

### 2. Key Topics Covered
*   **Built-in Roles:** Definitions for `role:readonly` and `role:admin`.
*   **RBAC Architecture:** Configuration via the global `argocd-rbac-cm` ConfigMap and `AppProject` CRDs.
*   **Casbin Policy Model:** Explanation of the `p` (policy) and `g` (group/role assignment) syntax.
*   **Permission Matrix:** A detailed table of resources (applications, clusters, certificates, logs, etc.) and their allowed actions (get, create, update, delete, sync, etc.).
*   **Application-Specific Granularity:** How to scope permissions to specific projects, namespaces, or sub-resources (e.g., allowing a user to delete Pods but not the Application itself).
*   **Special Actions:** Details on `action` (custom resource actions), `override` (syncing local manifests), and `exec` (web terminal access).
*   **SSO and Local User Integration:** Mapping OIDC scopes (groups/email) to Argo CD roles and handling local account security.
*   **Policy Composition & Validation:** Methods for splitting policies into multiple files and using the CLI to validate or test rules.

### 3. Technical Keywords
*   **Core Components:** `argocd-rbac-cm`, `AppProject`, `argocd-cm`.
*   **Policy Syntax:** `p (policy)`, `g (group)`, `allow`, `deny`.
*   **Matching Engine:** `Casbin`, `glob`, `regex`, `policy.matchMode`.
*   **Configuration Keys:** `policy.default`, `policy.csv`, `scopes`, `users.anonymous.enabled`, `server.rbac.disableApplicationFineGrainedRBACInheritance`.
*   **CLI Commands:** `argocd admin settings rbac validate`, `argocd admin settings rbac can`.
*   **Resource Identifiers:** `<app-project>/<app-name>`, `<app-project>/<app-ns>/<app-name>`.

### 4. Target Audience
*   **Cluster Administrators:** Responsible for securing the Argo CD instance.
*   **DevSecOps Engineers:** Designing least-privilege access models for development teams.
*   **Security Auditors:** Understanding the inheritance and priority (e.g., `deny` overriding `allow`) of the system.

### 5. Related Concepts
*   **SSO/OIDC Integration:** RBAC is useless without an identity provider (documented in `user-management/index.md`).
*   **AppProjects:** RBAC policies are often embedded within AppProject resources to provide multi-tenancy.
*   **ApplicationSets:** RBAC controls who can create ApplicationSets to prevent privilege escalation via templated projects.
*   **Web-based Terminal:** Specifically tied to the `exec` resource permissions.
*   **Proxy Extensions:** Controlled via the `extensions` resource permission.

---

### AI Maintenance Trigger: When to update this file
An AI system should flag this file for updates if any of the following code changes occur:

1.  **API Changes:** If a new resource type (e.g., a new CRD) is added to Argo CD, it must be added to the Resource/Action table.
2.  **RBAC Logic Modifications:** If the evaluation engine changes (e.g., how `glob` patterns handle separators or changes in `deny`/`allow` priority).
3.  **New Global Settings:** If new fields are added to `argocd-rbac-cm` (like the v3.0 fine-grained inheritance toggle).
4.  **CLI Updates:** If the `argocd admin` subcommands for RBAC testing are renamed or their flags are modified.
5.  **Built-in Role Changes:** If the default permissions in `builtin-policy.csv` are modified in the source code.