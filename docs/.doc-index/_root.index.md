# _ROOT Documentation Index

## Overview
This documentation area serves as the high-level entry point for Argo CD, covering core architectural concepts, installation procedures across different environments, and fundamental operational workflows. It provides the necessary context for new users to understand GitOps principles and guides administrators through initial setup, security considerations, and troubleshooting common issues.

## Files Summary
*   **getting_started.md**: Provides a step-by-step walkthrough for installing Argo CD on Kubernetes, accessing the API server, and deploying a first application.
*   **SUPPORT.md**: Directs users to official help channels, including the FAQ, Slack community, and GitHub issue reporting.
*   **security_considerations.md**: An archived record of historical CVEs and security audits, offering context on past vulnerabilities and general security best practices.
*   **try_argo_cd_locally.md**: A specific guide for developers to set up Argo CD on a local machine using `Kind` (Kubernetes in Docker).
*   **core_concepts.md**: Defines the specialized terminology used within Argo CD, such as Applications, Target State, Syncing, and Health status.
*   **cli_installation.md**: Detailed instructions for downloading and installing the `argocd` command-line interface on Linux, macOS (Apple Silicon), and Windows.
*   **CONTRIBUTING.md**: A brief pointer to the external developer guide for users looking to contribute code or documentation to the project.
*   **faq.md**: A comprehensive troubleshooting repository covering common technical hurdles, configuration errors, and specific resource behavior (e.g., Helm, SealedSecrets).
*   **bug_triage.md**: Outlines the internal process for classifying, labeling, and prioritizing GitHub issues based on severity and urgency.
*   **understand_the_basics.md**: Lists prerequisite knowledge and external resources for tools that Argo CD relies on, such as Docker, Kubernetes, and Helm.
*   **index.md**: The landing page for the documentation, providing an architectural overview, feature list, and an explanation of the GitOps pattern.
*   **roadmap.md**: Points to the live GitHub Project and release cadence documentation to track the future direction of the platform.

## Code Changes That Would Require Documentation Updates
*   **Manifest Path Changes**: Any change to the raw URL or location of `install.yaml` (affects `index.md`, `getting_started.md`, `try_argo_cd_locally.md`).
*   **CLI Command Refactoring**: Changes to flags, subcommands, or syntax for `argocd login`, `app create`, `app sync`, or `admin initial-password`.
*   **Default Password Logic**: Modifications to how `argocd-initial-admin-secret` is generated or how the default admin password is set (affects `getting_started.md`, `faq.md`).
*   **New Manifest Generators**: Adding support for a new tool (e.g., a new config management plugin) requires updates to `index.md` and `core_concepts.md`.
*   **Health Assessment Logic**: Updates to how Argo CD calculates the "Healthy" or "Progressing" status for K8s resources (affects `faq.md` and `core_concepts.md`).
*   **ConfigMap Key Changes**: Changes to the default keys in `argocd-cm` (e.g., `timeout.reconciliation`, `admin.enabled`, `application.instanceLabelKey`).
*   **Architecture Evolution**: Introduction of new microservices or changes to the interaction between `argocd-server`, `repo-server`, and `controller` (affects `index.md`).
*   **Redis/Auth Logic**: Changes to how Argo CD authenticates with Redis or rotates secrets (affects `faq.md`).
*   **New Platform Support**: Changes in CLI binary names or supported architectures/OS types (affects `cli_installation.md`).
*   **Security Patches**: Discovery of new vulnerabilities requiring an entry in the CVE table (affects `security_considerations.md`).

## Key Technical Concepts
*   **Application (CRD)**: The primary object representing a group of K8s resources.
*   **GitOps**: The operational pattern using Git as the "source of truth."
*   **Sync (Synchronize)**: The process of making the Live State match the Target State.
*   **Live State vs. Target State**: The current cluster reality vs. the desired state in Git.
*   **OutOfSync**: The state when Git and the Cluster differ.
*   **Health Status**: `Healthy`, `Progressing`, `Missing`, `Degraded`.
*   **Refresh**: The action of pulling the latest Git manifests to compare against the cluster.
*   **Tool / Config Management Tool**: Kustomize, Helm, Jsonnet, or Plugins.
*   **Self-healing**: Automatic syncing when drift is detected.
*   **RBAC & SSO**: Authorization and Authentication (OIDC, SAML, etc.).
*   **Port-forwarding**: The primary method recommended for local API access.

## Related Components
*   **argocd-server**: The API server and Web UI provider.
*   **argocd-repo-server**: The service that clones Git repos and generates manifests.
*   **argocd-application-controller**: The K8s controller monitoring applications and clusters.
*   **argocd-dex-server**: Handles SSO integration.
*   **argocd-redis**: Caching layer for manifest generation and session data.
*   **argocd-initial-admin-secret**: The secret containing the bootstrap password.
*   **argocd-cm**: The central ConfigMap for system-wide configuration.
*   **Kind (Kubernetes in Docker)**: Recommended local development cluster tool.