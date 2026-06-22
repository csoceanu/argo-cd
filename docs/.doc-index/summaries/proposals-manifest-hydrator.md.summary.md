This documentation analysis provides a comprehensive overview of the **Manifest Hydrator** proposal for Argo CD.

### 1. Primary Purpose
The file documents a proposal to make **manifest hydration** (the "rendered manifests pattern") a native, first-class feature of Argo CD. Its goal is to automate the process of taking "dry" sources (templates like Helm or Kustomize) and committing the resulting "hydrated" (plain Kubernetes) manifests back into a Git repository. This ensures that every change to an application's state is recorded in Git history, facilitating better auditing, debugging, and environment promotion.

### 2. Key Topics Covered
*   **The Rendered Manifest Pattern**: Shifting from hydration-at-runtime to hydration-at-commit-time within the GitOps lifecycle.
*   **Operational Modes**:
    *   **Push-to-deploy**: Hydrated manifests are pushed to the deployment branch.
    *   **Push-to-stage**: Hydrated manifests are pushed to a staging branch for external promotion systems to handle.
*   **Determinism and Reproducibility**: Strict opinions on avoiding non-deterministic configurations (e.g., unpinned Helm versions or runtime parameter overrides) to ensure the Git history is a reliable source of truth.
*   **Developer Experience**: Providing local CLI commands to exactly reproduce how Argo CD hydrates manifests.
*   **Workflow Automation**: How the application controller detects "dry" commits, performs hydration, and delegates pushes to a new "commit server."
*   **Metadata Generation**: Automatic creation of `README.md`, `manifest.yaml`, and `hydrator.metadata` files in the hydrated branch.

### 3. Technical Keywords
*   **CRD Fields**: `spec.sourceHydrator`, `drySource`, `syncSource`, `hydrateTo`.
*   **Configuration Files**: `.argocd-source.yaml` (replaces `spec.source` fields to keep config in Git).
*   **Output Files**: `manifest.yaml`, `hydrator.metadata`, `README.md`.
*   **APIs/Protobuf**: `CommitManifests` service, `CommitPathDetails`.
*   **Environment Variables**: `ARGOCD_APP_REVISION`, `ARGOCD_APP_SOURCE_REPO_URL` (and restricted/ignored variables like `KUBE_VERSION`).
*   **Security/Auth**: `repository-write` secret type, `githubAppID`, `githubInstallationID`.
*   **Tools**: Helm (Umbrella charts), Kustomize, Argo CD Application Controller.

### 4. Target Audience
*   **Argo CD Maintainers/Contributors**: To understand the architectural impact and new component requirements (Commit Server).
*   **Platform Engineers**: Designing GitOps workflows and automated environment promotion systems.
*   **DevOps Engineers**: Seeking to improve auditability and reduce "hidden" state changes in their CI/CD pipelines.

### 5. Related Concepts
*   **GitOps**: Specifically the "Rendered Manifests" vs. "Dry Source" debate.
*   **Environment Promotion**: Using Git branches as a promotion mechanism rather than just deployment targets.
*   **Argo CD Image Updater**: Shared logic for automated Git pushes.
*   **Mutating Webhooks**: Suggested as the alternative for cluster-specific runtime configuration that shouldn't be in Git.
*   **Helm/Kustomize**: The primary tools being "hydrated."

---

### Update Triggers for AI Systems
This proposal file should be referenced or updated if any of the following occur:
1.  **Application CRD Changes**: Any modification to `v1alpha1.Application` that impacts how sources or sync options are defined.
2.  **Repo-Server Logic Updates**: Changes to how Argo CD invokes Helm or Kustomize (as the Hydrator must mirror this logic).
3.  **New Secret Types**: Introduction of new repository credential types for write access.
4.  **Promotion Workflow Features**: Development of native Argo CD promotion tools that would interact with the `hydrateTo` field.
5.  **Environment Variable Changes**: Updates to the standard Argo CD build environment variables that might break the "Determinism" requirement.