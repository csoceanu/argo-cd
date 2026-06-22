This documentation provides a technical guide for implementing and managing **Config Management Plugins (CMP)** in Argo CD. CMPs allow users to extend Argo CD’s manifest generation capabilities beyond its built-in support for Helm, Kustomize, and Jsonnet.

### 1. Primary Purpose
The file documents the **Sidecar Plugin architecture**, which is the modern method for integrating custom configuration management tools into Argo CD. It provides instructions on how to define plugin configurations, install them as sidecars to the `argocd-repo-server`, and migrate from the deprecated ConfigMap-based plugin system.

### 2. Key Topics Covered
*   **Plugin Configuration**: Structure of the `ConfigManagementPlugin` manifest (a pseudo-CRD).
*   **Lifecycle Phases**: Detailed explanation of the `init` (preparation), `generate` (manifest creation), and `discover` (automatic matching) phases.
*   **Installation & Registration**: How to patch the `argocd-repo-server` to include a plugin sidecar and where to place configuration files.
*   **Variable/Parameter Passing**: How environment variables, application specs, and UI parameters are injected into the plugin environment.
*   **Performance & Optimization**: Using tar exclusions and manifest-generate-paths to speed up processing in monorepos.
*   **Security & Permissions**: Handling file modes (`preserveFileMode`) and sharing Git credentials (`provideGitCreds` via ASKPASS).
*   **Migration**: Steps to move from the legacy `argocd-cm` plugin system (removed in v2.8) to the sidecar model.

### 3. Technical Keywords
*   **APIs/Resources**: `ConfigManagementPlugin` (v1alpha1), `Application`, `argocd-repo-server`, `argocd-cmp-server`.
*   **Configuration Fields**: `init`, `generate`, `discover`, `parameters`, `static`, `dynamic`, `fileName`, `find.glob`, `find.command`.
*   **Environment Variables**: `ARGOCD_ENV_*`, `ARGOCD_APP_NAME`, `ARGOCD_APP_NAMESPACE`, `ARGOCD_APP_PARAMETERS`, `KUBE_VERSION`, `ARGOCD_EXEC_TIMEOUT`.
*   **CLI/Settings**: `reposerver.plugin.tar.exclusions`, `server.repo.server.timeout.seconds`, `argocd.argoproj.io/manifest-generate-paths`.
*   **Security**: `runAsUser: 999`, `ASKPASS`, `preserveFileMode`.

### 4. Target Audience
*   **Argo CD Administrators**: Responsible for installing and securing plugins.
*   **DevOps/Platform Engineers**: Creating custom manifest generation workflows for their teams.
*   **Tooling Developers**: Building integrations between specialized config tools and Argo CD.

### 5. Related Concepts
*   **Native Tools**: Helm, Kustomize, Jsonnet (the tools CMPs are designed to supplement or replace).
*   **Sidecar Pattern**: Kubernetes architectural pattern used to extend the repo-server.
*   **Manifest Generation**: The core process of converting source code (Git/OCI) into Kubernetes YAML.
*   **RBAC/Security**: The trust model regarding how external scripts interact with Argo CD’s internal environment.

---

### AI Update Triggers: When to update this file
An AI system should flag this documentation for updates if code changes occur in the following areas:

1.  **CMP Server Logic**: Changes to `argocd-cmp-server` (the Go binary that facilitates the plugin protocol).
2.  **Manifest Schema**: Updates to the `ConfigManagementPlugin` struct in the `argoproj/argo-cd` repository (e.g., adding new fields to `spec`).
3.  **Env Var Prefixing**: Changes to the `ARGOCD_ENV_` prefixing logic or how parameters are serialized into `ARGOCD_APP_PARAMETERS`.
4.  **Timeout Handling**: Changes to default timeouts or the introduction of new timeout environment variables.
5.  **Repo-Server Architecture**: Modifications to how `argocd-repo-server` communicates with sidecars (e.g., changes to the gRPC socket or mount paths).
6.  **Discovery Engine**: Changes to the globbing library or the execution flow of the `discover` command.
7.  **Credential Management**: New methods for passing secrets or Git credentials to sidecars.