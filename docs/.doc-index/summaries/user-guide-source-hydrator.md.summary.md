This documentation outlines the **Source Hydrator** feature in Argo CD, which implements the "rendered manifest pattern." It allows Argo CD to render (hydrate) concise, "DRY" (Don't Repeat Yourself) manifests—such as Helm charts or Kustomize configurations—into flat, "WET" Kubernetes manifests and push them back to a Git repository before they are synced to a cluster.

### 1. Primary Purpose
The file documents how to configure and use the Argo CD Source Hydrator to improve visibility into what is actually being deployed. By pushing hydrated manifests to Git, users can use Git-native workflows (like Pull Requests) for environment promotion and auditing, while keeping the original templates separate from the final rendered state.

### 2. Key Topics Covered
*   **Enabling the Feature**: Instructions for activating the "commit server" and modifying the `argocd-cmd-params-cm` ConfigMap.
*   **Authentication**: Setting up separate Push and Pull secrets using the `argocd.argoproj.io/secret-type: repository-write` label.
*   **Application Configuration**: Configuring the `spec.sourceHydrator` field, including `drySource` (templates) and `syncSource` (where hydrated files go).
*   **Staging & Promotion**: Using the `hydrateTo` field to push to a temporary branch for PR-based promotion.
*   **Commit Tracing & Metadata**: Using custom Git trailers (`Argocd-reference-commit-*`) to link hydrated manifests back to the original source code commits.
*   **Customization**: Using Go text templates and Sprig functions to customize Git commit messages.
*   **Limitations & Best Practices**: Warnings regarding signature verification, project scoping, and the necessity of deterministic hydration.

### 3. Technical Keywords
*   **Components**: `commit-server`, `argocd-repo-server`, `Application controller`.
*   **Configuration Fields**: `hydrator.enabled`, `sourceHydrator.commitMessageTemplate`, `spec.sourceHydrator`, `drySource`, `syncSource`, `hydrateTo`.
*   **Kubernetes Resources**: `ConfigMap` (`argocd-cmd-params-cm`, `argocd-cm`), `Secret`, `Application` (CRD).
*   **Labels/Annotations**: `argocd.argoproj.io/secret-type: repository-write`, `manifest-generate-paths`.
*   **Commit Trailers**: `Argocd-reference-commit-sha`, `Argocd-reference-commit-author`, `Argocd-reference-commit-repourl`.
*   **Tools**: Helm, Kustomize, GitHub Apps, jq, Sprig.

### 4. Target Audience
*   **DevOps/Platform Engineers**: Responsible for setting up Argo CD and defining deployment patterns.
*   **CI/CD Developers**: Those writing scripts to automate image bumps and manifest updates.
*   **Security Teams**: Interested in the auditability of manifests pushed to Git.

### 5. Related Concepts
*   **GitOps**: The core philosophy of using Git as the single source of truth.
*   **Rendered Manifest Pattern**: A specific GitOps architectural pattern where WET manifests are stored in Git.
*   **Environment Promotion**: The process of moving changes from dev to staging to production.
*   **Secret Management**: Specifically how to handle secrets on the destination cluster (e.g., External Secrets Operator) rather than during hydration.

---

### AI Update Trigger Analysis
This file should be updated if any of the following code changes occur:
*   **CRD Changes**: Modifications to the `Application` spec regarding `sourceHydrator` or its sub-fields.
*   **Component Evolution**: If the feature moves from **Alpha** to **Beta/GA** (changing default settings or removing specific `-with-hydrator.yaml` manifests).
*   **Feature Expansion**: Support for signature verification, multi-commit references, or `manifest-generate-paths` for the hydrator.
*   **ConfigMap Updates**: Changes to key names in `argocd-cmd-params-cm` or `argocd-cm`.
*   **Security Logic**: Changes to how `repository-write` secrets are selected or how project-scoping is enforced.
*   **Metadata Changes**: Adding or renaming Git trailers used for tracing.