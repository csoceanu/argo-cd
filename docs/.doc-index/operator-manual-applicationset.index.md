# OPERATOR-MANUAL/APPLICATIONSET Documentation Index

## Overview
This documentation covers the architecture, configuration, and operational management of the Argo CD ApplicationSet controller. It describes how to use "Generators" to automatically create, update, and delete Argo CD Applications across multiple clusters and repositories, enabling multi-cluster management, monorepo support, and self-service GitOps.

## Files Summary
*   **index.md**: High-level introduction to the ApplicationSet controller, parameter substitution, and its relationship with Argo CD Applications.
*   **Argo-CD-Integration.md**: Explains the "factory" role of the controller and its strict operational boundaries within the Argo CD namespace.
*   **Getting-Started.md**: Installation instructions (bundled vs. standalone), high-availability configuration, and post-upgrade safeguards.
*   **Security.md**: Critical security guidelines regarding RBAC, admin-only access, and risks associated with templated `project` fields.
*   **Application-Deletion.md**: Details on resource lifecycle, owner references, finalizers, and the `preserveResourcesOnDeletion` setting.
*   **Controlling-Resource-Modification.md**: Guide on limiting controller actions via policies (create-only, etc.), dry-run mode, and ignoring specific Application field differences.
*   **GoTemplate.md**: Instructions for using Go Text Templates, Sprig functions, and custom functions like `normalize` and `slugify`.
*   **Template.md**: Deep dive into the `template` and `templatePatch` fields for generating Application manifests.
*   **Progressive-Syncs.md**: Documentation for the Alpha feature enabling staged rollouts (`RollingSync`) and ordered deletions.
*   **Appset-Any-Namespace.md**: Detailed configuration for the Beta feature allowing ApplicationSets to exist outside the `argocd` namespace.
*   **Generators.md**: A central menu and summary of all available generator types.
*   **Generators-Cluster.md**: How to discover target clusters using Argo CD cluster secrets, including label selectors and the `flatList` feature.
*   **Generators-Git.md**: Details on the Git Directory and Git File generators for discovering app structures in repositories.
*   **Generators-Git-File-Globbing.md**: Specifically addresses the new `doublestar` globbing implementation and how to enable it.
*   **Generators-List.md**: Using literal lists or dynamic `elementsYaml` to provide parameters.
*   **Generators-Matrix.md**: Combining two generators to produce every possible combination of parameters.
*   **Generators-Merge.md**: Merging parameter sets from multiple generators using specific `mergeKeys`.
*   **Generators-SCM-Provider.md**: Automating repository discovery across GitHub, GitLab, Bitbucket, Azure DevOps, Gitea, and AWS CodeCommit.
*   **Generators-Pull-Request.md**: Automating ephemeral environment creation by discovering open PRs/MRs.
*   **Generators-Plugin.md**: Architecture for external generators using RPC HTTP requests.
*   **Generators-Cluster-Decision-Resource.md**: Interface for duck-typed Kubernetes resources (like Open Cluster Management placements).
*   **Generators-Post-Selector.md**: Standardized filtering of generator outputs using Kubernetes label selectors.
*   **Use-Cases.md**: Practical scenarios including cluster add-ons, monorepo management, and developer self-service.
*   **applicationset-specification.md**: Reference for the ApplicationSet CRD structure.

## Code Changes That Would Require Documentation Updates
*   **New Generator Types**: Any PR adding a new generator (e.g., a "Cloud-Provider" generator) would require a new `Generators-X.md` file and an update to `Generators.md`.
*   **Template Engine Changes**: Modifications to the templating logic, adding new Sprig-like functions, or changing the default behavior of `fasttemplate` vs. `GoTemplate` must be reflected in `GoTemplate.md`.
*   **Controller CLI Flags**: Adding new startup parameters (like `--enable-new-git-file-globbing` or `--enable-progressive-syncs`) requires updates to the relevant feature page and `Getting-Started.md`.
*   **CRD Schema Changes**: Any change to `ApplicationSet` fields (e.g., adding fields to `syncPolicy` or `strategy`) requires updates to `applicationset-specification.md` and `Template.md`.
*   **SCM API Integrations**: Updates to how the controller interacts with GitHub/GitLab APIs, or adding support for new SCM providers (e.g., Woodpecker, Forgejo), requires updates to `Generators-SCM-Provider.md`.
*   **Security Policies**: Changes to how `ownerReferences` are handled or new RBAC requirements for multi-namespace support require updates to `Security.md` and `Appset-Any-Namespace.md`.
*   **Finalizer Logic**: Altering how `resources-finalizer.argocd.argoproj.io` is injected or removed must be documented in `Application-Deletion.md`.
*   **Feature Maturity**: Moving "Progressive Syncs" from Alpha to Beta or "App in any namespace" to GA requires updating the warning banners in their respective files.

## Key Technical Concepts
*   **Generators**: Modules that produce a list of key-value pairs (parameters).
*   **Parameters**: Variables (e.g., `{{.cluster}}`, `{{.path}}`) substituted into the Application template.
*   **Duck-typing**: Used by the Cluster Decision Resource generator to read external CRD statuses without knowing the full schema.
*   **RollingSync**: A strategy for staging updates across groups of applications based on health.
*   **OwnerReferences**: The Kubernetes mechanism linking generated Applications to the parent ApplicationSet for garbage collection.
*   **resources-finalizer**: The Argo CD finalizer that ensures child resources (Deployments, etc.) are deleted when an Application is removed.
*   **MergeKeys**: The identifiers used to join parameter sets in the Merge generator.
*   **pathParamPrefix**: A setting to avoid namespace collisions when using multiple Git generators in a Matrix.
*   **Slugify/Normalize**: Custom functions to ensure string parameters are safe for Kubernetes resource naming.

## Related Components
*   **Argo CD Application Controller**: Responsible for syncing the Applications created by the AppSet controller.
*   **Argo CD Repo Server**: Used by Git generators to fetch repository contents and files.
*   **Argo CD API Server**: Handles the RBAC and API requests for ApplicationSets.
*   **SCM Providers**: (GitHub, GitLab, Bitbucket, etc.) External APIs polled by SCM/PR generators.
*   **Kubernetes Secrets**: Used to store cluster credentials and SCM access tokens.
*   **ConfigMaps**: Used for plugin configurations and global controller settings (`argocd-cmd-params-cm`).