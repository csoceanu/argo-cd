# DEVELOPER-GUIDE Documentation Index

## Overview
The Argo CD Developer Guide provides a comprehensive framework for contributors to set up, develop, test, and release the Argo CD ecosystem. It covers the entire technical lifecycle, including architectural deep-dives into authentication/authorization, local and virtualized development toolchains, and detailed procedures for submitting high-quality pull requests and managing releases.

## Files Summary
*   **index.md**: The entry point for contributors, categorizing guides based on contribution type (Docs, Backend, Frontend).
*   **development-environment.md**: Lists mandatory tools (Go, Docker, K8s clusters) and initial repository setup steps.
*   **toolchain-guide.md**: Compares local vs. virtualized development environments and provides setup instructions for Docker, Podman, and K8s providers.
*   **running-locally.md**: Detailed instructions for running Argo CD components outside of Kubernetes using Goreman or Docker for faster iteration.
*   **debugging-locally.md**: Explains how to debug specific backend components (like `api-server`) within IDEs while running the rest of the stack locally.
*   **development-cycle.md**: Defines the standard workflow for code changes, including dependency management, code generation (`codegen`), and linting.
*   **tilt.md**: Guide for using Tilt for real-time K8s-based development, offering better log visibility and hot-reloading.
*   **submit-your-pr.md**: Outlines PR requirements, title naming conventions, CI gates, and code coverage standards.
*   **code-contributions.md**: Details the "Enhancement Proposal" process and triage workflow for new features.
*   **releasing.md**: Step-by-step instructions for the automated release process using GitHub Actions.
*   **release-process-and-cadence.md**: Defines the release schedule (Minor vs. Patch), feature freeze periods, and support policies.
*   **ci.md**: Troubleshooting guide for Continuous Integration failures and updates to the builder image.
*   **test-e2e.md**: Instructions for running and troubleshooting the end-to-end test suite in isolated namespaces.
*   **architecture/components.md**: Visualizes the 4-layer architecture (UI, Application, Core, Infra) and describes component responsibilities.
*   **architecture/authz-authn.md**: Deep-dive into the API server logic, covering gRPC-gateway, Cmux, and RBAC enforcement.
*   **extensions/ui-extensions.md**: Documentation for extending the Web UI with custom resource tabs, system-level pages, and status panels.
*   **extensions/proxy-extensions.md**: Explains how to configure the API server as a reverse-proxy for backend-integrated UI extensions.
*   **dependencies.md**: Guidance for updating core shared libraries like `gitops-engine`, `notifications-engine`, and `argo-ui`.
*   **api-docs.md**: Instructions for accessing Swagger UI and utilizing the Application API via bearer tokens.
*   **debugging-remote-environment.md**: Utilizing Telepresence to intercept remote cluster traffic for local debugging.
*   **docs-site.md**: How to build and preview the MkDocs-based documentation site locally.
*   **faq.md**: Answers common development questions and lists all auto-generated code files.
*   **mac-users.md**: Troubleshooting specific to macOS, such as Port 5000 conflicts and architecture-related binary errors.
*   **static-code-analysis.md**: Overview of security and quality tools like Snyk, SonarCloud, and Codecov.

## Code Changes That Would Require Documentation Updates
*   **API Modifications**: Changing Protobuf definitions, adding new gRPC services, or altering REST endpoints requires updates to `api-docs.md` and potentially `architecture/authz-authn.md`.
*   **Infrastructure/Tooling Updates**: Upgrading Go versions, changing minimum Docker/K8s versions, or adding new `Makefile` targets requires updates to `development-environment.md`, `toolchain-guide.md`, and `development-cycle.md`.
*   **RBAC/Auth Logic Changes**: Modifying how Casbin or OIDC tokens are processed necessitates updates to `architecture/authz-authn.md`.
*   **New Components**: Adding a new controller or microservice requires updating `architecture/components.md`, the `Procfile` references in `running-locally.md`, and `tilt.md`.
*   **UI Extension Hooks**: Adding new areas for UI customization (e.g., new menu locations) requires updates to `extensions/ui-extensions.md`.
*   **Dependency Shifts**: Changing how the project interacts with `gitops-engine` or `argo-ui` needs to be reflected in `dependencies.md`.
*   **Release Process Changes**: Modifying GitHub Action workflows or tagging logic requires updating `releasing.md`.
*   **CI/Test Pipeline Changes**: Adding new linting rules or E2E test requirements requires updates to `ci.md`, `static-code-analysis.md`, or `test-e2e.md`.

## Key Technical Concepts
*   **Codegen**: The process of generating `*.pb.go`, Swagger specs, and manifests via `make codegen`.
*   **Toolchains**: The choice between "Virtualized" (Docker-based) and "Local" (native OS-based) development.
*   **Goreman**: The process manager used to run the multi-component Argo CD stack locally.
*   **Cmux**: The connection multiplexer used to serve gRPC and HTTP on the same port.
*   **gRPC-gateway**: The proxy that translates REST calls into gRPC for the backend.
*   **RBAC/Casbin**: The authorization engine used to validate user permissions.
*   **Telepresence**: The tool used for debugging local code against a live remote Kubernetes cluster.
*   **ApplicationSet Controller**: The specific controller for generating multiple applications from templates.
*   **Repo Server**: The component responsible for cloning Git repos and generating manifests (Kustomize/Helm).

## Related Components
*   **argocd-server**: The primary API and UI host.
*   **argocd-repo-server**: Manifest generation and Git interaction.
*   **argocd-application-controller**: The reconciliation loop for live state vs. Git.
*   **argocd-applicationset-controller**: Automation for app creation.
*   **argocd-notifications-controller**: Handling outbound alerts.
*   **dex**: The OIDC identity provider.
*   **redis**: The caching layer for API and manifest data.
*   **gitops-engine**: The shared core library for Kubernetes synchronization.