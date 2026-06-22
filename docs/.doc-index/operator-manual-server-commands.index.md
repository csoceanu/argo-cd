# OPERATOR-MANUAL/SERVER-COMMANDS Documentation Index

## Overview
This documentation area provides comprehensive command references and configuration guides for the core server-side components of Argo CD. It details the operational flags, environment settings, and configuration methods required to run the API server, repository server, application controllers, and identity management (Dex) utilities.

## Files Summary
* **argocd-dex_rundex.md**: Provides the command reference for running the Dex server using settings derived from Argo CD’s internal ConfigMaps and Secrets.
* **argocd-repo-server.md**: Details the Repository Server, which handles Git repository local caching and the generation of Kubernetes manifests from source code.
* **argocd-dex.md**: Serves as the parent entry point for internal Dex-related utility tools used within the Argo CD ecosystem.
* **argocd-dex_gendexcfg.md**: Explains the command used to manually generate a Dex configuration file based on current Argo CD settings.
* **argocd-server_version.md**: Document the specific flags for the version subcommand of the API server, used to retrieve build and version metadata.
* **argocd-server.md**: The primary reference for the Argo CD API server, covering gRPC/REST endpoints, authentication, OIDC, and global cache settings.
* **argocd-applicationset-controller.md**: Describes the controller responsible for managing ApplicationSet resources and automating multi-application generation via SCM/PR generators.
* **argocd-application-controller.md**: Detailed reference for the core application controller that reconciles the live state of the cluster with the desired state in Git.
* **additional-configuration-method.md**: Explains how to use the `argocd-cmd-params-cm.yaml` ConfigMap to configure server components globally as an alternative to CLI flags.

## Code Changes That Would Require Documentation Updates
* **Flag Modifications**: Adding, renaming, or removing CLI flags in the Go source code for any server component (e.g., adding a new `--enable-feature-x` flag).
* **Default Value Changes**: Changing the default value of any configuration parameter (e.g., increasing `default-cache-expiration` or changing a default port).
* **New Subcommands**: Adding new sub-functional commands to `argocd-server`, `argocd-dex`, or other binaries.
* **Cache Logic Updates**: Modifications to how Redis, repository state, or OIDC state is cached or expired.
* **Integration Extensions**: Adding support for new SCM providers, OCI media types, or OpenTelemetry (OTLP) attributes.
* **Scaling & Performance Tuning**: Changes to parallelism limits (webhook, manifest generation, or kubectl execution) and sharding methods.
* **Security Enhancements**: Updates to TLS cipher suites, minimum/maximum TLS versions, or RBAC impersonation logic.
* **ConfigMap Schema Changes**: Updating the `argocd-cmd-params-cm` mapping logic or adding new prefixes for component configuration.
* **Experimental Features**: Transitioning experimental features (like Hydrator or Progressive Syncs) to GA or changing their activation flags.

## Key Technical Concepts
* **Server Components**: `argocd-server`, `argocd-repo-server`, `argocd-application-controller`, `argocd-applicationset-controller`, `argocd-dex`.
* **Caching Infrastructure**: Redis, Redis Sentinel, `repo-cache-expiration`, `app-state-cache-expiration`, compression algorithms (gzip/none).
* **Authentication & Identity**: Dex, OIDC, JWT tokens, impersonation (`--as`, `--as-group`), Basic Auth.
* **Manifest Generation**: Helm, OCI, Git hidden directories, `allow-oob-symlinks`, manifest size limits.
* **Observability**: OpenTelemetry (OTLP), Prometheus metrics, log formats (JSON/Text), log levels.
* **Reconciliation & Sync**: Self-healing backoff, sharding methods (legacy, round-robin, consistent-hashing), hard resync, workqueue tuning.
* **Global Configuration**: `argocd-cmd-params-cm`, `kubeconfig` contexts, namespace scoping.

## Related Components
* **Argo CD API Server**: The central gateway for UI and CLI interaction.
* **Repository Server**: The internal manifest generation engine.
* **Application/ApplicationSet Controllers**: The logic engines for state reconciliation.
* **Dex Server**: The integrated identity provider for SSO.
* **Redis**: The state and manifest cache provider.
* **Kubernetes API Server**: The underlying platform where Argo CD reconciles resources.
* **OpenTelemetry Collector**: External system for receiving traces and metrics.