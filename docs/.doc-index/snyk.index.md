# SNYK Documentation Index

## Overview
This documentation area provides transparency regarding the security posture of Argo CD by displaying the results of automated Snyk security scans. These scans are performed weekly on the project's primary branch and the most recent patches of the supported minor releases to track and manage vulnerabilities in dependencies, container images, and infrastructure-as-code manifests.

## Files Summary
*   **snyk/index.md**: Acts as the central dashboard for security scan results, providing a tabular breakdown of critical, high, medium, and low vulnerabilities across different versions of Argo CD and its sub-components.

## Code Changes That Would Require Documentation Updates
*   **Release Lifecycle Events**: When a new minor version of Argo CD is released, this page must be updated to include the new version and remove the oldest supported version to maintain the "three most recent minor releases" policy.
*   **Dependency Management Changes**: Modifications to the backend dependency management (e.g., migrating from `go.mod` to another tool) or frontend package management (e.g., changing from `yarn.lock` to `package-lock.json`).
*   **Image Version Upgrades**: Updating the versions of bundled or recommended sidecar components such as Dex, HAProxy, or Redis, as the specific version tags (e.g., `redis:7.2.7-alpine`) are hardcoded in the report tables.
*   **Deployment Manifest Refactoring**: Adding new core installation files or renaming existing ones (e.g., `install.yaml` or `namespace-install.yaml`) which are tracked under the Infrastructure-as-Code (IaC) scan section.
*   **Scan Tooling/Policy Changes**: Changing the frequency of scans (currently weekly on Sundays) or switching from Snyk to a different vulnerability scanning provider.
*   **Repository Restructuring**: Moving the location of `go.mod` or the `ui/` directory, which would invalidate the paths used for the scan reports.

## Key Technical Concepts
*   **Snyk Scans**: Automated security audits for code, dependencies, and containers.
*   **Vulnerability Severity Levels**: Classification of security risks into Critical, High, Medium, and Low categories.
*   **go.mod**: The Go modules file defining backend dependencies for the Argo CD project.
*   **yarn.lock**: The lockfile for Node.js dependencies used by the Argo CD user interface.
*   **Container Image Scanning**: Security analysis of specific image tags (e.g., `quay.io/argoproj/argocd`).
*   **Infrastructure-as-Code (IaC) Scanning**: Security analysis of Kubernetes manifests (`install.yaml`) to detect misconfigurations.
*   **Patch Releases**: Targeted updates to specific minor versions (e.g., v3.1.5) to fix bugs or security issues.

## Related Components
*   **Argo CD Core**: The primary application components scanned for Go-based vulnerabilities.
*   **Argo CD UI**: The frontend component scanned for JavaScript/Node.js vulnerabilities.
*   **Dex**: The identity service used by Argo CD for authentication.
*   **Redis**: The caching layer used by Argo CD for performance optimization.
*   **HAProxy**: The load balancer/ingress component used in certain Argo CD deployment patterns.
*   **Installation Manifests**: The official YAML files used to deploy Argo CD into Kubernetes clusters.