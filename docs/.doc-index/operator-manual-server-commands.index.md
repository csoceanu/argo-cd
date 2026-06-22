# OPERATOR-MANUAL/SERVER-COMMANDS Documentation Index

## Overview
This documentation area provides a comprehensive reference for the server-side command-line interfaces and background processes that power Argo CD. It details the configuration options, execution flags, and architectural roles for core components including the API server, the application and applicationset controllers, the repository server, and the Dex identity service.

## Files Summary
*   **argocd-application-controller.md**: Detailed reference for the core controller that monitors live cluster state against Git and orchestrates synchronization, including sharding and self-healing settings.
*   **argocd-applicationset-controller.md**: Documentation for the controller managing ApplicationSet resources, covering SCM provider configurations, progressive syncs, and reconciliation policies.
*   **argocd-repo-server.md**: Reference for the internal service responsible for cloning Git repositories, caching manifests, and invoking rendering tools like Helm or Kustomize.
*   **argocd-server.md**: Documentation for the Argo CD API server, which provides the gRPC/REST interface for the Web UI and CLI, including authentication and Redis configuration.
*   **argocd-server_version.md**: Command reference for retrieving the version information of the Argo CD API server.
*   **argocd-dex.md**: High-level reference for the internal Dex utility tools used for identity management.
*   **argocd-dex_gendexcfg.md**: Reference for generating Dex configuration files based on Argo CD settings.
*   **argocd-dex_rundex.md**: Instructions for running the Dex service using configurations derived from Argo CD ConfigMaps and Secrets.
*   **additional-configuration-method.md**: Explains how to use the `argocd-cmd-params-cm` ConfigMap as an alternative to command-line flags for global configuration.

## Code Changes That Would Require Documentation Updates
*   **New CLI Flags**: Adding any new `flag` to the `cobra` commands in the server-side binaries (e.g., adding a new timeout or feature toggle).
*   **Default Value Modifications**: Changing the default value of any existing flag (e.g., increasing the default `repo-cache-expiration` or changing the default `tlsminversion`).
*   **Service Port Changes**: Modifying default ports for gRPC, HTTP, or metrics endpoints (e.g., changing the default `8081` for the repo-server).
*   **Caching Logic Updates**: Introducing new caching layers or changing how Argo CD interacts with Redis (e.g., new compression algorithms or Redis Sentinel configurations).
*   **Controller Logic Changes**: Altering the reconciliation loop behavior, sharding methods (`legacy`, `round-robin`, `consistent-hashing`), or self-healing backoff algorithms.
*   **Security Header/Policy Updates**: Changes to default Content Security Policy (CSP), X-Frame-Options, or TLS cipher suites.
*   **SCM Provider Additions**: Adding support for new SCM providers or generators within the ApplicationSet controller.
*   **OpenTelemetry/Observability Changes**: Adding new OTLP attributes, changing trace propagation, or adding new Prometheus metric labels.
*   **ConfigMap Mapping**: Changes to the `argocd-cmd-params-cm.yaml` logic or the prefixes used to map ConfigMap keys to command-line flags.
*   **Resource Limits**: Modifying hardcoded or configurable limits for manifest extraction sizes (Helm/OCI) or concurrent request limits.

## Key Technical Concepts
*   **Reconciliation & Sync**: `app-resync`, `self-heal-backoff`, `progressive-syncs`, `hard-resync`.
*   **Caching**: `redis`, `repo-cache`, `app-state-cache`, `revision-cache`, `redis-compress`.
*   **Identity & Auth**: `dex`, `oidc-cache`, `disable-auth`, `impersonation` (`--as`), `bearer token`.
*   **Scaling & Performance**: `sharding-method`, `parallelismlimit`, `operation-processors`, `status-processors`, `webhook-parallelism-limit`.
*   **Security**: `TLS`, `tlsciphers`, `insecure-skip-tls-verify`, `CSP` (Content-Security-Policy), `X-Frame-Options`.
*   **Observability**: `OTLP` (OpenTelemetry), `metrics-port`, `logformat` (json/text), `loglevel`.
*   **Manifest Generation**: `helm-manifest-max-extracted-size`, `oci-layer-media-types`, `plugin-tar-exclude`.

## Related Components
*   **Argo CD API Server (`argocd-server`)**: The frontend API gateway.
*   **Argo CD Application Controller**: The state reconciliation engine.
*   **Argo CD Repository Server (`argocd-repo-server`)**: The manifest generation and Git caching service.
*   **Argo CD ApplicationSet Controller**: The generator for multi-cluster/multi-app automation.
*   **Dex**: The embedded identity provider for SSO.
*   **Redis**: The mandatory caching backend for state and manifest data.
*   **argocd-cmd-params-cm**: The centralized configuration ConfigMap for server components.