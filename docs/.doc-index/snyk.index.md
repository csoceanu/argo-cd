# SNYK Documentation Index

## Overview
This documentation area tracks the security posture of the Argo CD project by reporting weekly Snyk vulnerability scans. It provides a high-level summary of security vulnerabilities (categorized by severity) for the project's dependencies, container images, and installation manifests across the master branch and supported stable releases.

## Files Summary
*   **snyk/index.md**: Serves as the central dashboard for Snyk scan results, providing summary tables that link to detailed vulnerability reports for Go dependencies, UI packages, container images (Dex, Redis, HAProxy, Argo CD), and IaC manifests for current and legacy versions.

## Code Changes That Would Require Documentation Updates
*   **Dependency Updates**: Modifying `go.mod` or `ui/yarn.lock` which will result in different vulnerability counts in subsequent scans.
*   **Release Lifecycle**: Releasing a new minor or patch version of Argo CD, requiring the addition of new tables and the removal of older, unsupported versions.
*   **Base Image Upgrades**: Updating the versions of bundled components such as Dex, HAProxy, or Redis in the Dockerfiles or build scripts.
*   **Manifest Changes**: Significant modifications to `install.yaml` or `namespace-install.yaml` that would impact Infrastructure as Code (IaC) security results.
*   **Scanning Configuration**: Changes to the Snyk scanning frequency (currently weekly) or the specific branches targeted for scans.
*   **Image Registry Migration**: Changing the registries used for Argo CD components (e.g., moving from Quay to GHCR or ECR), which would change the report URLs.

## Key Technical Concepts
*   **Snyk Scans**: Automated security audits for code, dependencies, and containers.
*   **Vulnerability Severity Levels**: Classification of risks into Critical, High, Medium, and Low.
*   **go.mod**: The Go modules file defining backend dependencies.
*   **ui/yarn.lock**: The lockfile defining frontend JavaScript dependencies for the Argo CD UI.
*   **IaC (Infrastructure as Code) Scanning**: Security analysis of Kubernetes manifests (`install.yaml`).
*   **Container Image Security**: Scanning of specialized images including `argocd`, `dex`, `redis`, and `haproxy`.
*   **Release Versioning**: Documentation is partitioned by `master` and specific semver tags (e.g., `v3.1.5`).

## Related Components
*   **Argo CD Server & UI**: Primary components being scanned for application-level vulnerabilities.
*   **Identity Management (Dex)**: The bundled authentication component.
*   **Caching Layer (Redis)**: The bundled data store component.
*   **Load Balancing (HAProxy)**: The bundled proxy component used in specific deployment configurations.
*   **Deployment Manifests**: The `install.yaml` and `namespace-install.yaml` files used for cluster setup.
*   **CI/CD Pipeline**: The automation that triggers these scans every Sunday.