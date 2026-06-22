# DEVELOPER-GUIDE Documentation Index

## Overview
The Argo CD Developer Guide provides the foundational technical instructions and workflow requirements for contributors to build, test, and maintain the Argo CD ecosystem. It covers everything from initial environment setup and local debugging strategies to the internal release processes, continuous integration standards, and dependency management for core engines.

## Files Summary
*   **running-locally.md**: Instructions for running Argo CD components outside of a Kubernetes cluster to speed up development cycles.
*   **tilt.md**: Guide for using Tilt to manage container-based local development with real-time logs and automated restarts.
*   **development-environment.md**: Lists mandatory tools (Go, Docker, K8s clusters) and steps to clone the repository and initialize the environment.
*   **development-cycle.md**: Defines the standard workflow for code changes, including dependency sync, code generation, linting, and testing.
*   **releasing.md**: Outlines the two-step automated release process using GitHub Actions and the necessary manual triggers.
*   **debugging-locally.md**: Detailed instructions for configuring IDEs (VS Code/GoLand) to debug specific Argo CD components using environment variables and `Procfile` configs.
*   **mac-users.md**: Troubleshooting guide for macOS-specific issues like port 5000 conflicts and architecture-related execution errors.
*   **toolchain-guide.md**: Compares local vs. virtualized (Docker-based) development toolchains and provides troubleshooting for cluster connectivity.
*   **submit-your-pr.md**: Quality standards for Pull Requests, including naming conventions, checklist requirements, and cherry-picking policies.
*   **release-process-and-cadence.md**: Details the schedule for minor/patch releases, feature acceptance criteria, and security patch policies.
*   **ci.md**: Troubleshooting guide for failing Continuous Integration checks and instructions for updating the builder image.
*   **code-contributions.md**: Explains the triage process for new features, including enhancement proposals, design documents, and contributor meetings.
*   **test-e2e.md**: Guidance on running and troubleshooting the end-to-end test suite in isolated namespaces.
*   **faq.md**: Frequently asked questions regarding PR labels, generated code, and contribution philosophy.
*   **static-code-analysis.md**: List of security and quality scanning tools (Snyk, SonarCloud, Codecov) used in the project.
*   **dependencies.md**: Protocols for updating upstream dependencies like `gitops-engine`, `notifications-engine`, and `argo-ui`.
*   **api-docs.md**: Instructions for accessing and authorizing the Swagger/OpenAPI documentation for the Argo CD API.
*   **debugging-remote-environment.md**: Guide for using Telepresence to intercept cluster traffic and debug services running in a remote Kubernetes environment.
*   **docs-site.md**: Instructions for building and serving the MkDocs-based documentation website locally.
*   **index.md**: The landing page and high-level roadmap for different types of contributors (Docs, Backend, Frontend).

## Code Changes That Would Require Documentation Updates
*   **Adding/Renaming Microservices**: Any change to the architectural components (e.g., a new controller) requires updates to `running-locally.md`, `tilt.md`, `debugging-locally.md` (Procfile section), and port-forwarding tables.
*   **API Modifications**: Changes to Protobuf (`.proto`) files or Swagger definitions require updates to `api-docs.md` and `faq.md` (generated code section).
*   **Toolchain/Build Logic**: Updates to the `Makefile`, Go version, or required system utilities (Docker/Podman/Node) necessitate updates to `development-environment.md` and `toolchain-guide.md`.
*   **CI/CD Pipeline Alterations**: Changes to GitHub Action workflows or the introduction of new linting/testing tools must be reflected in `ci.md`, `static-code-analysis.md`, and `development-cycle.md`.
*   **Dependency Upgrades**: Pulling new versions of `gitops-engine` or `argo-ui` should follow and potentially update the procedures in `dependencies.md`.
*   **Release Policy Changes**: Modifications to the release frequency, support windows, or "Release Champion" duties require updates to `release-process-and-cadence.md` and `releasing.md`.
*   **New Test Frameworks**: Introduction of new E2E testing patterns or requirements for test isolation must be documented in `test-e2e.md`.

## Key Technical Concepts
*   **Virtualized Toolchain**: Running builds and tests inside a Docker container to ensure environment parity across developer machines.
*   **Codegen**: The automated process (`make codegen`) of generating API glue code, Swagger specs, and installation manifests.
*   **Procfile**: A configuration file used by `goreman` to manage multiple local processes (API server, Repo server, etc.) simultaneously.
*   **Interception (Telepresence)**: Routing traffic from a remote K8s cluster to a local process for live debugging.
*   **Cherry-picking**: The process of backporting bug fixes from the `master` branch to specific release branches (e.g., `release-2.x`).
*   **Release Candidate (RC)**: Pre-GA versions used for feature freeze and final testing during the 7-week cycle.
*   **Protobuf/gRPC**: The underlying technology for Argo CD's internal and external API communication.

## Related Components
*   **argocd-server**: The API and UI hosting component.
*   **argocd-repo-server**: Internal service managing Git/Helm repository interactions.
*   **argocd-application-controller**: The K8s controller responsible for monitoring and syncing applications.
*   **argocd-applicationset-controller**: Manages automation for generating multiple applications.
*   **argocd-notifications-controller**: Handles outbound alerts and status updates.
*   **gitops-engine**: The core logic library shared between Argo CD and other GitOps tools.
*   **argo-ui**: Shared React components library for the frontend.