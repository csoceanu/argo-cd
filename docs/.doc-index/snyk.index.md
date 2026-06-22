# SNYK Documentation Index

## Overview
This documentation area tracks and displays the results of weekly Snyk security scans performed on the Argo CD project. It provides a transparency report on vulnerabilities (Critical, High, Medium, and Low) across the project's dependencies, container images, and Infrastructure as Code (IaC) manifests for the `master` branch and the three most recent stable release branches.

## Files Summary
*   **snyk/index.md**: Serves as the primary dashboard for security scan results, aggregating vulnerability counts for Go modules, Yarn dependencies, core container images (Argo CD, Dex, Redis, HAProxy), and installation manifests.

## Code Changes That Would Require Documentation Updates
*   **Release Lifecycle Events**: Adding a new minor release branch or removing an end-of-life (EOL) release version from the supported scan list.
*   **Dependency Management**: Updates to backend dependencies in `go.mod` or frontend dependencies in `ui/yarn.lock` that significantly alter the vulnerability profile.
*   **Base Image Version Upgrades**: Changing the versions of bundled third-party images such as `dex`, `redis`, or `haproxy` (e.g., upgrading Redis from 7.0.x to 7.2.x).
*   **Registry Changes**: Moving project images to different container registries (e.g., switching from `quay.io` to `ghcr.io` or `public.ecr.aws`).
*   **Infrastructure Changes**: Adding or renaming core installation manifests (e.g., creating a new `ha-install.yaml` or modifying `namespace-install.yaml`) that should be tracked for IaC security.
*   **Scan Tooling Configuration**: Changes to the Snyk CLI configuration or the automation frequency (currently weekly on Sundays) that would change how results are reported.
*   **Packaging Changes**: Modifications to the Argo CD Dockerfile that result in new image tags or different base image layers.

## Key Technical Concepts
*   **Snyk**: The security platform used for Software Composition Analysis (SCA) and Static Analysis.
*   **Vulnerability Severity**: Classification of security risks into Critical, High, Medium, and Low tiers.
*   **SCA (Software Composition Analysis)**: Scanning of `go.mod` and `yarn.lock` for known vulnerabilities in open-source libraries.
*   **Container Scanning**: Analysis of OCI/Docker images (`argocd`, `dex`, `redis`, `haproxy`) for operating system and application layer vulnerabilities.
*   **IaC (Infrastructure as Code) Scanning**: Evaluation of Kubernetes manifests (`install.yaml`) for security misconfigurations.
*   **Master vs. Patch Releases**: The distinction between the bleeding-edge development branch and stable, supported versions.

## Related Components
*   **Argo CD Core**: The main application scanned as `argocd:latest` or versioned tags.
*   **Argo CD UI**: The frontend component whose dependencies are tracked via `yarn.lock`.
*   **Redis**: The distributed caching layer used by Argo CD.
*   **Dex**: The identity service used for OIDC/external authentication.
*   **HAProxy**: The load balancer component used in High Availability (HA) configurations.
*   **Installation Manifests**: The official Kubernetes YAML files used to deploy Argo CD.