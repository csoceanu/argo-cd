# DEVELOPER-GUIDE Documentation Index

## Overview
This documentation area provides a comprehensive roadmap for contributors to Argo CD, covering the entire development lifecycle from initial environment setup to the final release process. It defines the standards for code quality, testing procedures (unit, lint, and e2e), debugging strategies in local and remote environments, and the administrative workflows for managing pull requests and repository dependencies.

## Files Summary

*   **api-docs.md**: Provides instructions on accessing Swagger UI for API exploration and explains how to authorize requests using bearer tokens.
*   **ci.md**: Explains the Continuous Integration pipeline, including troubleshooting failed checks, retriggering builds, and the public CD image publishing process.
*   **code-contributions.md**: Outlines the philosophy for contributing features, the enhancement proposal triage process, and requirements for design documents.
*   **contributing.md**: A placeholder file redirecting contributors to the toolchain guide.
*   **debugging-locally.md**: A detailed guide on configuring IDEs (VS Code/GoLand) to debug individual Argo CD components while others run via the local toolchain.
*   **debugging-remote-environment.md**: Instructions for using Telepresence to bridge local development with a remote Kubernetes cluster for real-world environment testing.
*   **dependencies.md**: Technical procedures for pulling changes from and contributing to shared libraries like `gitops-engine`, `notifications-engine`, and `argo-ui`.
*   **development-cycle.md**: Defines the standard sequence of operations for development, including dependency syncing, code generation, linting, and running test suites.
*   **development-environment.md**: Lists mandatory tools (Go, Docker, Kubectl) and provides setup guides for local Kubernetes clusters like Kind, Minikube, and K3d.
*   **docs-site.md**: Instructions for building and serving the documentation website locally using `mkdocs` to preview changes.
*   **faq.md**: Answers common contributor questions regarding PR labels, review processes, and the nature of auto-generated code.
*   **index.md**: The landing page and navigational overview for different contribution paths (documentation, frontend, or backend).
*   **mac-users.md**: Addresses macOS-specific development hurdles, such as port 5000 conflicts with AirPlay and virtualized toolchain execution errors.
*   **release-process-and-cadence.md**: Sets the schedule for minor and patch releases, defines the role of the "Release Champion," and establishes feature acceptance criteria.
*   **releasing.md**: Detailed technical workflow for project maintainers to trigger automated releases using GitHub Actions and the release script.
*   **running-locally.md**: Comprehensive guide for running Argo CD services outside a cluster using Docker (virtualized) or Goreman (local) for rapid iteration.
*   **static-code-analysis.md**: Lists the automated tools used for code linting, coverage reporting, and security/vulnerability scanning.
*   **submit-your-pr.md**: Documentation on PR quality gates, naming conventions (feat, fix, docs, chore), and automated test requirements for merging.
*   **test-e2e.md**: Guide for executing and troubleshooting end-to-end tests that validate the interaction between Argo CD components and Kubernetes.
*   **tilt.md**: Explains how to use Tilt for a container-based development workflow with live reloading and integrated log visibility.
*   **toolchain-guide.md**: Compares local and virtualized development environments, providing setup instructions for Docker, Podman, and networking.
*   **ui-extensions.md**: A placeholder file redirecting to the specific guide for developing UI extensions.
*   **use-gitpod.md**: States the current status of Gitpod availability for the project (currently unavailable).

## Code Changes That Would Require Documentation Updates

*   **API Specification Changes**: Any modification to `.proto` files or the addition of new API endpoints requires updates to `api-docs.md` and `development-cycle.md` (regarding `make codegen`).
*   **Build System/Tooling Updates**: Changing Go versions, updating the `Makefile`, or introducing new mandatory development tools requires updating `development-environment.md` and `toolchain-guide.md`.
*   **New Microservices**: Adding a new component (e.g., a new controller) requires updates to `running-locally.md`, `debugging-locally.md`, `tilt.md`, and `test-e2e.md` to include the new service in start/debug configs.
*   **Dependency Management Changes**: Altering the way `gitops-engine` or `argo-ui` are integrated or changing the vendor process requires updating `dependencies.md`.
*   **Release Policy Shifts**: Changing the frequency of minor releases or the criteria for feature freeze requires updates to `release-process-and-cadence.md`.
*   **CLI Structure Changes**: Adding or renaming CLI commands requires updates to `faq.md` (which references command generation) and `submit-your-pr.md`.
*   **CI Pipeline Modifications**: Introducing new linting rules or changing GitHub Action workflows requires updates to `ci.md`, `static-code-analysis.md`, and `submit-your-pr.md`.
*   **Installation Manifest Changes**: Modifying how Argo CD is installed (e.g., changing `install.yaml` or CRDs) requires updates to the "Deploy Argo CD resources" sections in `running-locally.md` and `development-environment.md`.

## Key Technical Concepts

*   **Development Tools**: `make`, `goreman`, `tilt`, `kubectl`, `kind`, `k3d`, `minikube`, `telepresence`.
*   **Build/Codegen Commands**: `make codegen`, `make build`, `make test`, `make test-e2e`, `make install-tools-local`.
*   **Configuration Files**: `Procfile` (process management), `go.mod` (Go dependencies), `yarn.lock` (UI dependencies), `goreleaser.yaml` (release config), `launch.json` (IDE debug).
*   **Environment Variables**: `ARGOCD_SERVER`, `ARGOCD_OPTS`, `IMAGE_NAMESPACE`, `IMAGE_TAG`, `DOCKER=podman`.
*   **Testing/Linting**: `golangci-lint`, `eslint`, `codecov`, `unit tests`, `e2e tests`.
*   **PR/Release Workflow**: `cherry-pick/x.y` labels, `feat/fix/docs/chore` prefixes, `Release Champion`, `Release Candidate (RC)`.

## Related Components

*   **Argo CD Server (`argocd-server`)**: The API and UI server.
*   **Repo Server (`argocd-repo-server`)**: Handles manifest generation and Git interactions.
*   **Application Controller (`argocd-application-controller`)**: Monitors and reconciles application state.
*   **ApplicationSet Controller (`argocd-applicationset-controller`)**: Automates application creation.
*   **Notifications Controller (`argocd-notifications-controller`)**: Handles outbound notifications.
*   **Dex Server (`argocd-dex-server`)**: Manages identity and authentication.
*   **Redis**: Provides caching for the system.
*   **Shared Engines**: `gitops-engine` and `notifications-engine`.
*   **UI Components**: `argo-ui`.