This documentation provides a comprehensive upgrade guide for migrating from **Argo CD version 2.14 to 3.0**. It outlines breaking changes, performance optimizations, and security enhancements introduced in the 3.0 release.

### 1. Primary Purpose
The file serves as the **official upgrade manual** for Argo CD 3.0. Its goal is to provide operators with a "low-risk" path to version 3.0 by detailing breaking changes, offering detection methods to see if an environment is impacted, and providing remediation steps or "opt-out" configurations to restore legacy (v2.x) behavior.

### 2. Key Topics Covered
*   **RBAC Enhancements**: Shift toward more restrictive defaults, specifically regarding sub-resource updates and log access.
*   **Performance & Scale**: Introduction of default resource exclusions (e.g., hiding high-churn objects like `Endpoints` or `CiliumIdentity`) and changes to how resource health is persisted to reduce load on the K8s API and Application controller.
*   **Authentication (Dex/SSO)**: Changes to how subject claims (`sub`) are handled in Dex to use more stable federated IDs.
*   **Resource Tracking**: The transition from label-based tracking to annotation-based tracking as the default mechanism for managing resources.
*   **ApplicationSet Logic**: Standardizing the behavior of nested selectors.
*   **Deprecations & Removals**: Final removal of legacy `argocd-cm` repository configurations and specific legacy Prometheus metrics.
*   **Dependency Updates**: Implications of upgrading to Helm 3.17.1.

### 3. Technical Keywords
*   **Configuration Objects**: `argocd-cm`, `argocd-rbac-cm`, `argocd-cmd-params-cm.yaml`.
*   **RBAC Actions**: `update/*`, `delete/*`, `logs, get`.
*   **Configuration Keys**: 
    *   `server.rbac.disableApplicationFineGrainedRBACInheritance`
    *   `application.resourceTrackingMethod`
    *   `resource.exclusions`
    *   `resource.customizations.ignoreResourceUpdates`
    *   `controller.resource.health.persist`
*   **Metrics**: `argocd_app_info` (replacement), `argocd_app_sync_status` (removed).
*   **Authentication**: `federated:id`, `federated_claims.user_id`, Dex SSO.
*   **Tools/CLI**: `kubectl`, `jq`, `helm template`, `argocd app sync`.

### 4. Target Audience
*   **Kubernetes Platform Engineers**: Responsible for maintaining the Argo CD installation and its performance.
*   **DevOps/SREs**: Users who manage RBAC policies, monitoring dashboards, and CI/CD pipelines.
*   **Security Administrators**: Those auditing access controls and SSO integrations.

### 5. Related Concepts
*   **GitOps Reconciliation**: How Argo CD detects and manages differences between Git and the cluster.
*   **Least Privilege (RBAC)**: The shift toward explicit log permissions and sub-resource isolation.
*   **Kubernetes Controller Performance**: Reducing API churn by ignoring high-frequency status updates.
*   **OIDC/Identity Management**: Integration with Dex for external authentication.

---

### Update Triggers for AI Systems
This file should be updated or referenced when:
*   **Code Changes in RBAC Logic**: If the underlying enforcement of application actions or sub-resources is modified.
*   **Default Configuration Changes**: If the default list of `resource.exclusions` or `ignoreResourceUpdates` in the source code is expanded or reduced.
*   **Metric Schema Changes**: If the labels or names of controller metrics are altered in the Go source code.
*   **Dependency Bumps**: When the bundled version of **Helm** or **Dex** is updated, especially if those updates contain breaking upstream changes.
*   **Resource Tracking Logic**: If the format of the `argocd.argoproj.io/tracking-id` annotation is changed.
*   **CRD Sanitization**: If the API response for Projects or Clusters is further restricted for security.