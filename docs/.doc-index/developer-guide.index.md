# DEVELOPER-GUIDE Documentation Index

## Overview
This documentation area provides a comprehensive roadmap for contributors to set up, develop, test, and release Argo CD. It covers environment configuration for local and virtualized toolchains, the end-to-end development lifecycle, architectural debugging strategies, and the formal processes for submitting code and managing releases.

## Files Summary
*   **running-locally.md**: Instructions for running Argo CD components outside of a Kubernetes cluster using Docker/Podman or local binaries to speed up the development loop.
*   **tilt.md**: A guide to using Tilt for real-time service updates and integrated debugging within a local Kubernetes cluster.
*   **development-environment.md**: Lists mandatory tools (Go, Docker, Kind/Minikube) and provides the initial setup for forking and cloning the repository.
*   **development-cycle.md**: Details the standard workflow for contributors, including dependency management, code generation, linting, and running unit/E2E tests.
*   **releasing.md**: Outlines the automated two-step release process via GitHub Actions, including manifest updates and CLI binary generation.
*   **ui-extensions.md**: A placeholder file indicating that UI extension documentation has moved to the extensions guide.
*   **debugging-locally.md**: Explains how to isolate and debug a single Argo CD component in an IDE while running other services via Goreman.
*   **mac-users.md**: Documents known issues and fixes specific to macOS, such as port conflicts (port 5000) and binary execution errors.
*   **toolchain-guide.md**: Compares the virtualized (Docker-based) toolchain with the local toolchain, providing setup instructions for each.
*   **submit-your-pr.md**: Defines the quality gates, PR naming conventions (fix, feat, docs), and CI requirements for merging code.
*   **contributing.md**: A redirect file pointing users toward the comprehensive Toolchain guide for contribution details.
*   **use-gitpod.md**: States the current status of Gitpod availability for the project (currently not available).
*   **release-process-and-cadence.md**: Documents the minor and patch release schedules, feature acceptance criteria, and security patch policies.
*   **ci.md**: Provides troubleshooting steps for failed GitHub Actions checks and instructions for updating the builder image.
*   **code-contributions.md**: Outlines the enhancement proposal process (triage, design documents) and regular contributor meetings.
*   **test-e2e.md**: Describes the architecture and execution of the end-to-end test suite, including namespace isolation and troubleshooting.
*   **faq.md**: Answers common questions regarding PR review wait times, generated code locations, and PR labeling.
*   **static-code-analysis.md**: Lists the static analysis tools used by the project, including golangci-lint, Codecov, and Snyk.
*   **dependencies.md**: Explains the process for pulling in changes from shared upstream repositories like gitops-engine, notifications-engine, and argo-ui.
*   **api-docs.md**: Provides instructions for accessing Swagger documentation and interacting with the Applications API via bearer tokens.
*   **debugging-remote-environment.md**: A guide to using Telepresence to swap a remote cluster component with a locally running debug instance.
*   **docs-site.md**: Instructions for building, serving, and testing the mkdocs-based documentation site locally.
*   **index.md**: The high-level entry point and navigation guide for different types of contributors (docs, backend, frontend).

## Code Changes That Would Require Documentation Updates
*   **Adding a New Microservice/Component**: Updates required in `Procfile` (in `debugging-locally.md`), `tilt.md` (port forwarding/debug ports), and `running-locally.md`.
*   **Changing API Definitions**: Updates needed for `api-docs.md` and triggering the code generation workflow described in `faq.md` and `development-cycle.md`.
*   **Modifying Build/Test Infrastructure**: Updates to `toolchain-guide.md`, `ci.md`, and `Makefile` target descriptions in `development-cycle.md`.
*   **Altering the Release Workflow**: Updates to `releasing.md` if GitHub Actions or tagging logic changes, and `release-process-and-cadence.md` for schedule shifts.
*   **Updating External Dependencies**: Changes to the sync process in `dependencies.md` (e.g., if moving to a new version of `gitops-engine`).
*   **Adding New Linter or Security Tool**: Updates to `static-code-analysis.md` and the CI troubleshooting section in `ci.md`.
*   **Changes to E2E Test Framework**: Updates to `test-e2e.md` regarding namespace management, port usage, or setup logic.
*   **New Environment Prerequisites**: Changes to `development-environment.md` if the minimum Go, Node, or Kubernetes versions are incremented.

## Key Technical Concepts
*   **Toolchains**: Virtualized (Docker-based) vs. Local (Native binary) environments.
*   **Goreman**: Process manager using `Procfile` to orchestrate multiple local microservices.
*   **Codegen**: Automated generation of Protobuf stubs, Swagger specs, and Kubernetes manifests.
*   **E2E (End-to-End)**: Testing the full system in a dedicated Kubernetes namespace (`argocd-e2e`).
*   **Telepresence**: Network intercept tool used for "swapping" a cluster pod with a local process.
*   **Cherry-picking**: The process of backporting bug fixes to older release branches using labels.
*   **Enhancement Proposals**: The formal triage and design process for new features.
*   **Image Namespace/Tag**: Environment variables (`IMAGE_NAMESPACE`, `IMAGE_TAG`) used to customize build outputs.

## Related Components
*   **Argo CD API Server**: The primary interface for the CLI and UI.
*   **Argo CD Repo Server**: Responsible for managing Git repositories and generating manifests.
*   **Argo CD Application Controller**: The operator that synchronizes state with the Kubernetes cluster.
*   **Argo CD ApplicationSet Controller**: Handles multi-cluster/template-based application generation.
*   **Argo CD Notifications Controller**: Manages outbound alerts based on application state changes.
*   **Dex**: The OIDC provider for identity management.
*   **Redis**: Used for caching repository and OIDC state.
*   **GitOps Engine**: The core library shared across Argo projects for reconciliation logic.
*   **Argo UI**: Shared React components used in the frontend.