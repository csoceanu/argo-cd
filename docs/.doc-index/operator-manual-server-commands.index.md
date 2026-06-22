# OPERATOR-MANUAL/SERVER-COMMANDS Documentation Index

## Overview
This documentation area provides a comprehensive reference for the server-side command-line interfaces (CLIs) and services that constitute the Argo CD backend. It details the operational parameters, flags, and configuration methods for the core controllers, API servers, and utility tools required to run and maintain an Argo CD instance.

## Files Summary
*   **argocd-dex_rundex.md**: Provides the reference for the command that executes the Dex identity service using configurations derived from Argo CD’s ConfigMap and Secret.
*   **argocd-repo-server.md**: Detailed reference for the internal repository server responsible for manifest generation, Git repository caching, and OCI/Helm interactions.
*   **argocd-dex.md**: The parent command reference for `argocd-dex` utility tools used for identity management within Argo CD.
*   **argocd-dex_gendexcfg.md**: Reference for the utility that generates a Dex-compatible configuration file from Argo CD's environment settings.
*   **argocd-server_version.md**: Reference for the command used to retrieve the version and build information of the Argo CD API server.
*   **argocd-server.md**: Comprehensive reference for the primary API server (gRPC/REST) which facilitates communication between the Web UI, CLI, and external CI/CD systems.
*   **argocd-applicationset-controller.md**: Documentation for the controller responsible for managing ApplicationSet resources and automating large-scale application deployments.
*   **argocd-application-controller.md**: Detailed reference for the core Kubernetes controller that reconciles the live state of applications against the desired state defined in Git.
*   **additional-configuration-method.md**: Explains how to use the `argocd-cmd-params-cm.yaml` ConfigMap as an alternative to command-line flags for global and component-specific configuration.

## Code Changes That Would Require Documentation Updates
*   **New CLI Flags**: Adding, renaming, or deprecating any command-line options in the Golang source code for any of the listed binaries.
*   **Default Value Updates**: Changing the default value of an existing flag (e.g., increasing the default cache expiration or changing a default port).
*   **Controller Logic Changes**: Modifications to the reconciliation loop, sharding methods (legacy, round-robin, consistent-hashing), or self-heal backoff strategies.
*   **Security & Auth**: Changes to TLS version support, cipher suites, or identity provider integration logic (Dex/OIDC).
*   **Metric Additions**: Registering new Prometheus metrics or changing existing metric names/labels that are exposed via the `--metrics-port`.
*   **Cache Provider Changes**: Updates to how Argo CD interacts with Redis, including new compression algorithms (e.g., gzip) or TLS connection settings.
*   **Manifest Generation Evolution**: Changes to how the Repo Server handles Helm, Kustomize, OCI, or plugin-based manifest generation.
*   **ConfigMap Schema Changes**: Altering the prefix or key structure in `argocd-cmd-params-cm.yaml` used to override server parameters.
*   **API Protocol Changes**: Updates to allowed content types, OpenTelemetry (OTLP) headers, or gRPC/REST endpoint behavior.

## Key Technical Concepts
*   **Reconciliation Loop**: The process by which the `application-controller` detects and corrects out-of-sync resources.
*   **Manifest Generation**: The conversion of source code (Git/Helm) into Kubernetes manifests performed by the `repo-server`.
*   **ApplicationSet Generators**: Logic used by the `applicationset-controller` to dynamically create applications (e.g., SCM, PR, or Git file generators).
*   **Sharding**: Distributing the management of clusters across multiple controller replicas.
*   **Dex/OIDC**: The authentication layers for mapping external identity providers to Argo CD.
*   **Hydrator**: Feature flag mentioned in server/controller commands for managing large resource states.
*   **argocd-cmd-params-cm**: The central configuration mechanism for overriding hardcoded defaults or flags.
*   **Workqueue Rate Limiting**: Parameters (wq-bucket-qps, wq-maxdelay-ns) used to tune the performance of the controller's processing queue.

## Related Components
*   **Argo CD API Server**: The primary interface for users and automation.
*   **Argo CD Repository Server**: The backend manifest engine.
*   **Argo CD Application Controller**: The state enforcement engine.
*   **Argo CD ApplicationSet Controller**: The automation engine for Application resources.
*   **Dex**: The bundled identity federation service.
*   **Redis**: The distributed caching layer used by the server, controller, and repo-server.
*   **Kubernetes API Server**: The target environment for resource reconciliation and source for Argo CD configuration (Secrets/ConfigMaps).