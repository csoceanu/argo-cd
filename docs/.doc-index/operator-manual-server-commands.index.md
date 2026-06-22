# OPERATOR-MANUAL/SERVER-COMMANDS Documentation Index

## Overview
This documentation area provides a comprehensive command-line reference for the core backend components of Argo CD. It details the execution flags, configuration options, and operational parameters for the primary services that manage application state, manifest generation, identity federation, and API exposure. These documents serve as the definitive guide for operators configuring the Argo CD server-side stack via CLI arguments or environment-mapped ConfigMaps.

## Files Summary

*   **argocd-application-controller.md**: Detailed reference for the controller responsible for monitoring live cluster state, comparing it against the desired state in Git, and performing reconciliation/self-healing.
*   **argocd-applicationset-controller.md**: Documentation for the controller that automates the generation of Argo CD Applications using various generators (SCM, Git, etc.).
*   **argocd-repo-server.md**: Reference for the internal service that clones repositories, caches manifests, and handles the heavy lifting of manifest generation (Helm, Kustomize, etc.).
*   **argocd-server.md**: Reference for the main API server that provides the gRPC/REST backend for the Web UI and CLI, including authentication and RBAC.
*   **argocd-server_version.md**: Specific documentation for the versioning subcommand of the API server.
*   **argocd-dex.md / argocd-dex_rundex.md / argocd-dex_gendexcfg.md**: A suite of documents covering the Dex utility tools used to generate identity provider configurations and run the integrated Dex server.
*   **additional-configuration-method.md**: Explains how to use the `argocd-cmd-params-cm.yaml` ConfigMap to persistently configure server flags without modifying deployment manifests directly.

## Code Changes That Would Require Documentation Updates

*   **CLI Flag Modifications**: Adding, removing, or renaming any `flags` in the Golang source code for any of the server components (using libraries like Cobra).
*   **Default Value Changes**: Modifying the default values of existing flags (e.g., changing the default `repo-cache-expiration` or `metrics-port`).
*   **Component Logic Updates**:
    *   Changes to **Sharding** logic (e.g., adding new `sharding-method` options).
    *   Updates to **Manifest Generation** limits (e.g., Helm/OCI max extracted size).
    *   Modifications to **Sync Policies** or self-healing backoff algorithms.
*   **Configuration Mapping**: Changes to how the `argocd-cmd-params-cm` ConfigMap keys map to internal component variables or prefixes (e.g., adding a new prefix for a new service).
*   **Integration Updates**: Changes to how components interact with **Redis** (caching, TLS, Sentinel), **Dex** (OIDC config generation), or **OpenTelemetry** (tracing attributes/headers).
*   **Security/TLS Policy**: Updating supported TLS versions, ciphers, or certificate validation logic for internal gRPC or external HTTPS traffic.
*   **Feature Flag Additions**: Introducing new functionality controlled by flags (e.g., `--hydrator-enabled`, `--server-side-diff-enabled`).

## Key Technical Concepts

*   **Reconciliation & Self-Healing**: Parameters governing how often the controller checks state and how it reacts to drift (e.g., `--app-resync`, `--self-heal-timeout-seconds`).
*   **Manifest Generation**: Controls for Helm, OCI, and Git directory processing (e.g., `--helm-manifest-max-extracted-size`, `--oci-layer-media-types`).
*   **Caching & Redis**: Configuration for application state and repository metadata caching (e.g., `--repo-cache-expiration`, `--redis-compress`).
*   **Sharding**: Methods for distributing applications across multiple controller replicas (e.g., `consistent-hashing`, `round-robin`).
*   **Identity & Auth**: Configuration for OIDC, Dex, and impersonation (e.g., `--dex-server`, `--oidc-cache-expiration`, `--as-group`).
*   **Observability**: Settings for metrics endpoints and OpenTelemetry tracing (e.g., `--metrics-port`, `--otlp-address`).
*   **ConfigMap-based Overrides**: The use of `argocd-cmd-params-cm` for centralized configuration management using component prefixes (`server.`, `reposerver.`, `controller.`).

## Related Components

*   **Argo CD API Server**: The primary interface for users and the CLI.
*   **Argo CD Repository Server**: The worker service for Git and manifest operations.
*   **Argo CD Application Controller**: The state reconciliation engine.
*   **Argo CD ApplicationSet Controller**: The automation engine for multi-app management.
*   **Dex**: The embedded identity federation service.
*   **Redis**: Used for state caching across all major server components.
*   **Kubernetes API Server**: The target for resource orchestration and impersonation.