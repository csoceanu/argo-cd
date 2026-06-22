# _ROOT Documentation Index

## Overview
This documentation serves as the entry point for the Argo CD project, covering high-level architecture, core GitOps concepts, and the end-to-end "Day 0" experience. It provides the necessary instructions for installing the platform, configuring the CLI, and troubleshooting common initial deployment issues.

## Files Summary
*   **getting_started.md**: A step-by-step guide for installing Argo CD on Kubernetes, accessing the API server, and deploying a first "Guestbook" application.
*   **SUPPORT.md**: Provides a hierarchy of resources for getting help, including the FAQ, Slack channels, and the GitHub issue tracker.
*   **security_considerations.md**: An archived record of historical CVEs and security-related workarounds for older versions of Argo CD.
*   **try_argo_cd_locally.md**: A focused guide for developers to set up Argo CD in a local environment using Kind (Kubernetes in Docker).
*   **core_concepts.md**: Defines specific terminology used within the platform, such as Application, Target State, Live State, and Refresh.
*   **cli_installation.md**: Detailed instructions for downloading and installing the `argocd` CLI binary on Linux, macOS (Intel/Silicon), and Windows.
*   **CONTRIBUTING.md**: A pointer to the official developer-oriented contribution guidelines for the project.
*   **faq.md**: A comprehensive collection of troubleshooting tips for common errors related to sync status, Redis, Helm, and resource health.
*   **bug_triage.md**: Internal documentation describing the process for core maintainers to label, prioritize, and classify incoming GitHub issues.
*   **understand_the_basics.md**: Curates external learning resources for prerequisite technologies like Docker, Kubernetes, Helm, and Kustomize.
*   **index.md**: The landing page for the documentation, offering an architectural overview, feature list, and high-level "How it works" explanation.
*   **roadmap.md**: Directs users to the live GitHub Project board for the current release plans and development milestones.

## Code Changes That Would Require Documentation Updates
*   **Installation Manifests**: Any changes to the `install.yaml` (e.g., new services, changed resource names, or modified `ClusterRoleBinding` logic) must be reflected in `getting_started.md` and `index.md`.
*   **CLI Command Structure**: Adding new subcommands, changing flag names (e.g., `--insecure`, `--grpc-web`), or modifying the output of `argocd app get` requires updates to `getting_started.md` and `cli_installation.md`.
*   **Initial Admin Credentials**: If the logic for generating the default admin password changes (currently using `argocd-initial-admin-secret`), both the `getting_started.md` and `faq.md` must be updated.
*   **Health Assessment Logic**: Any modification to how Argo CD determines the "Healthy" or "Progressing" state for standard K8s resources (like Ingress or StatefulSets) requires updating the `faq.md`.
*   **Authentication/Security Defaults**: Changes to default Redis password requirements, JWT cookie length handling, or SSO configuration patterns (Dex/OIDC) must be updated in `faq.md` and `security_considerations.md`.
*   **Application Source Types**: Adding native support for new config management tools (beyond Helm, Kustomize, Jsonnet) requires updates to `index.md` and `core_concepts.md`.
*   **Cluster Management**: Changes to how external clusters are registered (e.g., the `argocd-manager` service account or role permissions) must be updated in `getting_started.md`.
*   **Labeling & Triage**: Changes to the project's GitHub label schema (priority/severity) require updates to `bug_triage.md`.

## Key Technical Concepts
*   **CRDs**: `Application` is the primary Custom Resource Definition.
*   **State Management**: `Target State` (Git) vs. `Live State` (Cluster).
*   **Syncing**: `Sync Policy` (Manual/Automatic), `OutOfSync` status, and `Sync Hooks` (Pre/Post).
*   **CLI Commands**: `argocd login`, `argocd app create`, `argocd cluster add`, `argocd admin initial-password`, `argocd account update-password`.
*   **Connectivity**: `Port Forwarding`, `LoadBalancer` service type, and `gRPC-web`.
*   **Security/Auth**: `Bcrypt` hashes for passwords, `JWT` tokens, `OIDC`, `SSO`, and `RBAC`.
*   **Internal Components**: `argocd-server`, `argocd-repo-server`, `argocd-application-controller`, `argocd-redis`.
*   **Normalization**: Differences in how Kubernetes normalizes resource units (e.g., `1000m` to `1`) and how Argo CD diffs them.

## Related Components
*   **Kubernetes API Server**: The destination for deployed manifests.
*   **Git Providers**: Source of truth (GitHub, GitLab, Bitbucket).
*   **Helm / Kustomize / Jsonnet**: Template engines and manifest generators.
*   **Redis**: Used for caching and authentication state.
*   **Dex**: The identity provider connector used for SSO.
*   **SealedSecrets / Bitnami**: Referenced specifically regarding health status issues.
*   **Kind / MicroK8s**: Supported local development/testing environments.