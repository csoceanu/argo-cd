# OPERATOR-MANUAL/APPLICATIONSET Documentation Index

## Overview
This documentation covers the configuration, management, and architecture of the Argo CD ApplicationSet controller. It describes how to use various generators (List, Cluster, Git, SCM, etc.) to automate the creation and lifecycle management of Argo CD Applications across multiple clusters and repositories, including advanced features like progressive syncs, Go templating, and multi-namespace support.

## Files Summary

*   **Appset-Any-Namespace.md**: Explains how to enable and configure ApplicationSets to be managed in namespaces other than the default `argocd` namespace, including security considerations for SCM providers.
*   **Application-Deletion.md**: Describes the resource lifecycle, owner references, and the use of finalizers to control how child Applications and their resources are deleted.
*   **applicationset-specification.md**: Provides the full YAML specification and field definitions for the ApplicationSet Custom Resource Definition (CRD).
*   **Argo-CD-Integration.md**: Outlines the architectural relationship between the ApplicationSet controller and the core Argo CD components, framing the controller as an "Application factory."
*   **Controlling-Resource-Modification.md**: Details settings for limiting controller actions, including dry-run mode, sync policies (create-only, create-update), and ignoring specific Application field differences.
*   **Generators-Cluster-Decision-Resource.md**: Documentation on using duck-typing to interface with external resources (like Open Cluster Management) to decide target clusters.
*   **Generators-Cluster.md**: Explains how to generate parameters based on clusters defined within Argo CD, including label selectors and the `flatList` option.
*   **Generators-Git-File-Globbing.md**: Discusses the transition from greedy globbing to the `doublestar` library and how to enable the new behavior.
*   **Generators-Git.md**: Covers the Git Directory and Git File generators, including path parameterization, exclusion rules, and webhook configurations.
*   **Generators-List.md**: Describes the simplest generator which uses a literal list of key/value pairs to generate Applications.
*   **Generators-Matrix.md**: Explains how to combine the output of two generators to create a Cartesian product of parameters.
*   **Generators-Merge.md**: Details how to merge parameters from multiple generators using specific merge keys to override base configurations.
*   **Generators-Plugin.md**: Documentation on extending ApplicationSet functionality by creating custom generators that communicate via RPC/HTTP.
*   **Generators-Post-Selector.md**: Explains how to use standard Kubernetes label selectors to filter the results of any generator.
*   **Generators-Pull-Request.md**: Covers the discovery of open pull/merge requests across various SCM providers to automate preview environment deployments.
*   **Generators-SCM-Provider.md**: Describes how to automatically discover repositories within an SCM organization (GitHub, GitLab, etc.) to drive Application generation.
*   **Generators.md**: Provides a high-level overview and categorization of all available generators within the ApplicationSet controller.
*   **Getting-Started.md**: Provides installation instructions, requirements, and steps for enabling High Availability (HA) mode.
*   **GoTemplate.md**: Detailed guide on using Go Text Templates, Sprig functions, and custom ApplicationSet functions for advanced parameter substitution.
*   **index.md**: The landing page and introduction to the ApplicationSet controller, explaining its purpose and core mechanics.
*   **Progressive-Syncs.md**: Documentation on Alpha-stage features for staged rollouts, including `RollingSync` strategies and `Reverse` deletion orders.
*   **Security.md**: Outlines critical security practices, RBAC requirements, and the risks associated with templated project fields.
*   **Template.md**: Explains the `template` and `templatePatch` fields, including how to patch non-string values and handle Helm-in-Helm templating conflicts.
*   **Use-Cases.md**: Illustrates practical applications such as managing cluster add-ons, monorepo support, and self-service Application provisioning.

## Code Changes That Would Require Documentation Updates

*   **Generator Logic**: Adding a new generator type or modifying the internal logic/parameters of existing ones (e.g., adding a new SCM provider to the SCM generator).
*   **Templating Engine**: Changes to `fasttemplate` implementation or additions/removals of Sprig/custom functions in the GoTemplate engine.
*   **CRD Schema Changes**: Any modification to the `argoproj.io/v1alpha1` ApplicationSet spec, including new fields in `syncPolicy`, `strategy`, or `template`.
*   **Controller Arguments**: Adding or changing CLI flags for the `argocd-applicationset-controller` (e.g., enabling/disabling features like progressive syncs or new globbing).
*   **Environment Variables**: Changes to the `ARGOCD_APPLICATIONSET_CONTROLLER_*` variables that govern global behavior like log levels, namespace scoping, or SCM allow-lists.
*   **Sync & Deletion Policies**: Modifications to how `ownerReferences`, finalizers, or the `--policy` flag influence Application lifecycle and resource pruning.
*   **RBAC & Multi-tenancy**: Changes to how the controller validates namespaces or how it interacts with `AppProject` resources.
*   **Webhook Handling**: Updates to the `/api/webhook` endpoint logic or changes in supported Git provider event payloads.
*   **Progressive Sync Algorithms**: Updates to health-checking logic, `maxUpdate` calculation, or rollout group sorting.

## Key Technical Concepts

*   **Generators**: The core mechanism for data sourcing (List, Cluster, Git, SCM, PR, Matrix, Merge, Plugin).
*   **Parameters**: Key-value pairs generated by generators for use in templates (e.g., `{{.cluster}}`, `{{.path.basename}}`).
*   **GoTemplate**: The advanced templating engine supporting Sprig and logic.
*   **templatePatch**: A field allowing YAML/JSON-based patches to the generated Application spec.
*   **RollingSync**: A progressive sync strategy for staged application updates.
*   **MergeKeys**: Specific identifiers used by the Merge generator to correlate data from different sources.
*   **Duck-typing**: Used by the Cluster Decision Resource generator to interpret custom resources.
*   **resources-finalizer**: The Argo CD finalizer used to ensure clean deletion of cluster resources.
*   **PreserveResourcesOnDeletion**: A sync policy setting to prevent cascading deletion of manifests.
*   **scm-creds**: A specific secret type label required for secure SCM token referencing.

## Related Components

*   **Argo CD Application Controller**: The component that actually reconciles the Applications generated by the ApplicationSet controller.
*   **Argo CD Repo Server**: Used by the Git generator to fetch and parse files/directories.
*   **Argo CD API Server**: Interfaces with the controller for cluster discovery and RBAC.
*   **SCM Providers**: External systems (GitHub, GitLab, Bitbucket, Azure DevOps, Gitea) providing API data for discovery.
*   **Kubernetes API Server**: The source of truth for the ApplicationSet CRDs and cluster secrets.