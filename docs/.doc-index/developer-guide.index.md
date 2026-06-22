# DEVELOPER-GUIDE Documentation Index

## Overview
This documentation area provides a comprehensive roadmap for contributors to the Argo CD project. It covers everything from initial environment setup and local development toolchains to architectural deep-dives, release processes, and the submission of pull requests. Its purpose is to ensure that developers can efficiently build, test, and debug Argo CD components while adhering to the project's quality standards and governance models.

## Files Summary
*   **index.md**: The high-level entry point directing contributors to specific guides based on their goals (docs, frontend, backend).
*   **architecture/components.md**: Describes the modular component-based architecture, logical layers (UI, Application, Core, Infra), and specific responsibilities of each service.
*   **architecture/authz-authn.md**: Explains the technical implementation of authentication and authorization, including the use of Cmux, gRPC-gateway, and Casbin.
*   **development-environment.md**: Lists mandatory tools (Go, Docker, K8s clusters) and steps for forking/cloning the repository.
*   **toolchain-guide.md**: Detailed comparison between virtualized (Docker-based) and local development toolchains.
*   **development-cycle.md**: Step-by-step workflow for a contribution, including dependency management, code generation, linting, and testing.
*   **running-locally.md**: Instructions for running Argo CD services outside of a Kubernetes cluster to speed up development.
*   **debugging-locally.md**: Guide for configuring IDEs (VS Code/GoLand) to debug individual services while others run via `goreman`.
*   **debugging-remote-environment.md**: Instructions for using Telepresence to debug local code against a live remote Kubernetes cluster.
*   **tilt.md**: Explains how to use Tilt for real-time log visibility and iterative container-based development.
*   **test-e2e.md**: Configuration and troubleshooting guide for running the end-to-end test suite.
*   **submit-your-pr.md**: Outlines PR quality gates, title naming conventions, CI checks, and code coverage requirements.
*   **ci.md**: Troubleshooting guide for failing GitHub Actions checks and builder image maintenance.
*   **code-contributions.md**: Details the enhancement proposal process, triage states, and design document requirements.
*   **releasing.md**: Technical steps for triggering automated releases, updating versions, and generating manifests.
*   **release-process-and-cadence.md**: Overview of the 3-month minor release cycle, security patch policies, and feature acceptance criteria.
*   **dependencies.md**: Procedures for updating shared repositories like `gitops-engine`, `notifications-engine`, and `argo-ui`.
*   **api-docs.md**: Instructions for accessing Swagger UI and technical details on the Applications API behavior.
*   **extensions/ui-extensions.md**: Documentation on building custom JavaScript plugins for Resource Tabs and the Status Panel.
*   **extensions/proxy-extensions.md**: Details the beta feature for reverse-proxying requests to backend services via the API Server.
*   **docs-site.md**: Instructions for building and testing the documentation website locally using MkDocs.
*   **faq.md**: Common questions regarding PR status, labeling, and a list of all auto-generated code files.
*   **static-code-analysis.md**: Summarizes the security and quality scanning tools (Snyk, SonarCloud, Codecov).
*   **mac-users.md**: Troubleshooting specific to macOS, including port conflicts and execution format errors.

## Code Changes That Would Require Documentation Updates
*   **Adding/Removing CLI Commands or Flags**: Requires updating `faq.md` (list of generated docs) and running `make codegen` to update CLI documentation.
*   **Modifying API Protobufs (`.proto` files)**: Requires updating `api-docs.md` if auth or behavior changes, and regenerating glue code via `make codegen`.
*   **Adding New Microservices**: Must update `architecture/components.md`, `running-locally.md`, `debugging-locally.md`, and `tilt.md` with new port/service definitions.
*   **Changes to Build/Test Infrastructure**: If `Makefile` targets or `Dockerfile` environments change, `development-cycle.md` and `ci.md` must be updated.
*   **Dependency Version Bumps**: Updating `gitops-engine` or `argo-ui` requires updating the SHA references in `dependencies.md`.
*   **New Extension Types**: Any change to how UI or Proxy extensions are registered or authorized requires updates to the files in the `extensions/` folder.
*   **RBAC/Auth Logic Changes**: Changes to how Casbin or session management works require updates to `architecture/authz-authn.md`.
*   **Installation Manifest Changes**: Modifying `manifests/install.yaml` logic requires updating `releasing.md` and `running-locally.md`.

## Key Technical Concepts
*   **Toolchains**: Virtualized (Docker-based) vs. Local (Native OS).
*   **Code Generation**: `make codegen`, Swagger/OpenAPI, Protobuf (`.pb.go`).
*   **Local Execution**: `goreman`, `Procfile`, `make start`, `make start-local`.
*   **CI/CD**: GitHub Actions, `Init ArgoCD Release`, `Publish ArgoCD Release`, Snyk, Codecov.
*   **Testing**: Unit tests (`make test`), End-to-End tests (`make test-e2e`), Testdata repositories.
*   **Debugging**: Delve, Telepresence (Interceptors), IDE Launch Configurations (`launch.json`).
*   **Extensions**: `extensionsAPI`, Resource Tab Extensions, Proxy Extensions (Reverse Proxy).
*   **RBAC**: Casbin rules, Project-level vs. Global-level permissions.
*   **Network Protocols**: Cmux (connection multiplexing), gRPC, gRPC-gateway (REST translation).

## Related Components
*   **argocd-server**: The primary API and UI server.
*   **argocd-repo-server**: Handles Git/Helm interaction and manifest generation.
*   **argocd-application-controller**: Manages the reconciliation loop and live-state monitoring.
*   **argocd-applicationset-controller**: Manages ApplicationSet resources.
*   **argocd-dex-server**: Handles OIDC authentication.
*   **argocd-redis**: Cache layer for Git and K8s data.
*   **argocd-notifications-controller**: Handles outbound alerts/notifications.
*   **gitops-engine**: Shared library for Kubernetes state synchronization.
*   **argo-ui**: Shared React component library.