# OPERATOR-MANUAL/APPLICATIONSET Documentation Index

## Overview
This documentation area covers the **ApplicationSet controller**, a sub-component of Argo CD designed to automate the generation and management of multiple `Application` resources. It details the architecture of "Generators" which provide parameter substitution for templates, lifecycle management (creation, updates, and deletion), security protocols, and advanced sync strategies like Progressive Syncs.

## Files Summary
*   **Generators-Git-File-Globbing.md**: Explains the transition from greedy default globbing to the `doublestar` implementation and how to enable this optional behavior via controller flags or ConfigMaps.
*   **Application-Deletion.md**: Describes the relationship between ApplicationSet deletion and its child Applications, focusing on owner references, finalizers, and the `preserveResourcesOnDeletion` setting.
*   **Controlling-Resource-Modification.md**: Details methods to restrict the controller's behavior, including dry-run mode, sync policies (`create-only`, `create-update`, etc.), and ignoring specific Application field differences.
*   **Generators-Matrix.md**: Explains how to combine two generators to produce a Cartesian product of parameters, including parameter overriding and path prefixing.
*   **Generators-List.md**: Covers the List generator for targeting clusters based on literal key/value pairs and dynamic generation via `elementsYaml`.
*   **Argo-CD-Integration.md**: Provides a high-level overview of how the ApplicationSet controller acts as an "Application factory" within the Argo CD ecosystem.
*   **GoTemplate.md**: A comprehensive guide on using Go Text Templates, including Sprig functions, custom functions like `normalize`/`slugify`, and migration steps from the legacy engine.
*   **Generators-SCM-Provider.md**: Details repository discovery across various providers (GitHub, GitLab, Gitea, Bitbucket, Azure DevOps, AWS CodeCommit) and filtering mechanisms.
*   **Security.md**: Outlines critical security requirements, such as restricting ApplicationSet creation to admins and the risks associated with templated `project` fields.
*   **Generators-Cluster.md**: Explains how to target clusters defined in Argo CD secrets using label selectors and how to pass additional metadata via the `values` field.
*   **Appset-Any-Namespace.md**: Documents the beta feature allowing ApplicationSets to exist outside the `argocd` namespace, including the necessary security allow-lists for SCM providers.
*   **Generators-Cluster-Decision-Resource.md**: Describes using Kubernetes duck-typing to interface with external cluster managers like Open Cluster Management (OCM).
*   **Generators-Post-Selector.md**: Explains how to apply a standard Kubernetes label selector to filter the output of any generator.
*   **Generators-Plugin.md**: Provides instructions for creating custom generators via an HTTP RPC interface (sidecar or standalone).
*   **Generators-Pull-Request.md**: Details how to discover open pull/merge requests to create ephemeral test environments, including webhook configuration for instant updates.
*   **Generators-Git.md**: Covers the Git Directory and File generators, including directory exclusion rules, polling intervals, and webhook triggers.
*   **Generators.md**: Serves as a high-level directory and introduction to the nine types of available generators.
*   **Generators-Merge.md**: Explains how to merge parameter sets from multiple generators using specific `mergeKeys`.
*   **Progressive-Syncs.md**: Documents the alpha feature for staged rollouts (`RollingSync`) and ordered deletion of Applications based on labels.
*   **Template.md**: Describes the structure of the ApplicationSet template, generator-level overrides, and the `templatePatch` feature for advanced JSON/YAML patching.
*   **applicationset-specification.md**: Provides the full YAML specification and schema for the ApplicationSet Custom Resource Definition.
*   **index.md**: The entry point for the ApplicationSet controller, explaining basic concepts and parameter substitution.
*   **Use-Cases.md**: Outlines practical applications including cluster add-ons, monorepo management, and developer self-service patterns.
*   **Getting-Started.md**: Provides installation instructions, high-availability configuration, and post-upgrade safety tips.

## Code Changes That Would Require Documentation Updates
*   **New Generator Implementation**: Adding a new generator type (e.g., a "Cloud-Provider" generator) would require new files in the `Generators-*.md` series.
*   **Changes to Parameter Mapping**: Modifying the default parameters produced by existing generators (e.g., adding `head_short_sha_7` to the PR generator).
*   **Templating Engine Updates**: Changes to the Go Template implementation, adding new custom functions (like `slugify`), or deprecating the `fasttemplate` engine.
*   **Controller CLI/Environment Variables**: Adding new startup flags (e.g., `--enable-new-git-file-globbing` or `--enable-progressive-syncs`).
*   **RBAC and Multi-Namespace Logic**: Modifications to how the controller validates namespaces or projects, especially concerning the "AppSet in any namespace" feature.
*   **Sync Strategy Extensions**: Enhancements to Progressive Syncs, such as new `matchExpressions` operators or new rollout types.
*   **CRD Schema Changes**: Any update to the `ApplicationSet` spec in the Go types, which must be reflected in `applicationset-specification.md`.
*   **External Integration Changes**: Updates to how the controller interacts with SCM APIs (GitHub/GitLab/Bitbucket) or the Argo CD Repo Server.

## Key Technical Concepts
*   **Generators**: List, Cluster, Git (File/Directory), Matrix, Merge, SCM Provider, PR, Cluster Decision Resource, Plugin.
*   **Templating**: `goTemplate`, `templatePatch`, `fasttemplate`, Sprig functions, `normalize`, `slugify`.
*   **Parameters**: `{{.path.path}}`, `{{.branch}}`, `{{.url}}`, `{{.nameNormalized}}`, `{{.values}}`.
*   **Policies**: `sync`, `create-only`, `create-update`, `create-delete`, `dry-run`.
*   **Lifecycle**: `ownerReferences`, `resources-finalizer`, `preserveResourcesOnDeletion`.
*   **Filtering**: `labelSelector`, `matchExpressions`, `repositoryMatch`, `pathsExist`.
*   **Connectivity**: `requeueAfterSeconds`, `webhook` (GitHub/GitLab), `tokenRef`, `appNamespace`.
*   **Sync Strategies**: `RollingSync`, `AllAtOnce`, `maxUpdate`, `deletionOrder`.

## Related Components
*   **Argo CD Application Controller**: Responsible for the actual deployment of resources generated by ApplicationSets.
*   **Argo CD Repo Server**: Used by Git and PR generators to fetch repository metadata and files.
*   **Argo CD API Server**: Used for managing ApplicationSets via CLI/UI.
*   **SCM Providers**: GitHub, GitLab, Gitea, Bitbucket, Azure DevOps, AWS CodeCommit.
*   **Kubernetes API**: The source of truth for Cluster Secrets and Cluster Decision Resources.
*   **Argocd-cmd-params-cm**: The central ConfigMap for tuning controller behavior.