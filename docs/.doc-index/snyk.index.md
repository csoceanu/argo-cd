# SNYK Documentation Index

## Overview
This documentation area tracks and summarizes the security vulnerability status of Argo CD across its active release branches. It provides a high-level dashboard of Snyk scan results, categorized by severity, for application dependencies, container images, and infrastructure-as-code (IaC) manifests.

## Files Summary
*   **snyk/index.md**: Acts as the central dashboard for security monitoring, displaying vulnerability counts (Critical, High, Medium, Low) for the `master` branch and the three most recent minor releases, with links to detailed HTML reports for each component.

## Code Changes That Would Require Documentation Updates
*   **Release Lifecycle Events**: Adding a new minor release or a new patch version requires creating a new table entry and updating the list of supported versions.
*   **Dependency Management Changes**: If new dependency files are added (e.g., a new `package.json` in a different subdirectory or a change from `yarn` to `npm`), the scan targets must be updated.
*   **Base Image Updates**: Changing the version or provider of third-party images (e.g., updating Redis from `7.0` to `7.2` or switching from Alpine to a different base image) requires updating the labels and report links in the tables.
*   **Infrastructure Changes**: Adding or renaming core installation manifests (e.g., adding a new `ha-install.yaml`) would require adding new rows for IaC scanning.
*   **Security Policy Updates**: If the scanning frequency or the criteria for "supported releases" changes (currently the three most recent minor releases), the introductory text must be revised.
*   **Component Architecture**: Adding a new architectural component that requires a dedicated container image (e.g., a new sidecar or utility) would necessitate a new entry in the scan tables.

## Key Technical Concepts
*   **Snyk Scans**: Automated security analysis for vulnerabilities in code, dependencies, and containers.
*   **Vulnerability Severities**: Classification of security risks into Critical, High, Medium, and Low categories.
*   **Software Composition Analysis (SCA)**: Monitoring of `go.mod` (Golang) and `yarn.lock` (JavaScript/UI) files for known vulnerable libraries.
*   **Container Image Scanning**: Security auditing of Docker/OCI images including `dex`, `haproxy`, `redis`, and the core `argocd` image.
*   **Infrastructure as Code (IaC) Scanning**: Static analysis of Kubernetes manifests (`install.yaml`, `namespace-install.yaml`) for misconfigurations.
*   **Patch Versioning**: Tracking security posture across different maintenance releases (e.g., v3.1.5, v3.0.16).

## Related Components
*   **Argo CD Core**: The main application logic whose dependencies and final image are scanned.
*   **Identity Provider (Dex)**: The bundled authentication component.
*   **Data Store (Redis)**: The caching layer used by Argo CD.
*   **Load Balancer (HAProxy)**: Used in high-availability configurations.
*   **Argo CD UI**: The frontend component whose `yarn.lock` is specifically monitored.
*   **Installation Manifests**: The raw YAML files used for deploying Argo CD in Kubernetes.