# OPERATOR-MANUAL/APPLICATIONSET Documentation Index

## Overview
This documentation covers the **ApplicationSet controller**, a sub-component of Argo CD that acts as an "Application factory." It explains how to use **Generators** to automate the creation, update, and deletion of multiple Argo CD Applications across diverse clusters and repositories using a single declarative manifest.

## Files Summary
*   **index.md**: General introduction to the controller, the `ApplicationSet` CRD, and the basic workflow of parameter substitution.
*   **Getting-Started.md**: Installation instructions (bundled vs. standalone), High Availability (HA) configuration, and post-upgrade safeguards.
*   **Argo-CD-Integration.md**: High-level architectural overview of how the ApplicationSet controller interacts with the core Argo CD components.
*   **Use-Cases.md**: Detailed scenarios including cluster add-ons, monorepos, and self-service multitenancy.
*   **Generators.md**: An overview and directory of the nine supported generator types.
*   **Generators-Git.md**: Deep dive into the **Git Directory** and **Git File** generators for discovering apps via repository structure.
*   **Generators-Git-File-Globbing.md**: Technical explanation of the greedy globbing issue and how to enable/use the new `doublestar` globbing implementation.
*   **Generators-List.md**: Instructions for the **List** generator, covering both literal lists and dynamic `elementsYaml`.
*   **Generators-Cluster.md**: Explains the **Cluster** generator for targeting clusters registered in Argo CD, including label selectors and `flatList` mode.
*   **Generators-SCM-Provider.md**: Comprehensive guide for discovering repositories across GitHub, GitLab, Gitea, Bitbucket, Azure DevOps, and AWS CodeCommit.
*   **Generators-Pull-Request.md**: Details the **Pull Request** generator for creating ephemeral environments based on open PRs/MRs, including webhook setup.
*   **Generators-Matrix.md**: Explains how to combine two generators to create a Cartesian product of parameters.
*   **Generators-Merge.md**: Explains how to merge parameters from multiple generators using specific merge keys.
*   **Generators-Plugin.md**: Documentation for creating custom HTTP-based generators via the **Plugin** system.
*   **Generators-Cluster-Decision-Resource.md**: Covers the **Cluster Decision Resource** generator using duck-typing to interface with external cluster managers like OCM.
*   **Generators-Post-Selector.md**: Explains how to apply a common Kubernetes `labelSelector` to filter results from any generator.
*   **Template.md**: Detailed guide on Application templates, generator-level overrides, and the advanced `templatePatch` feature.
*   **GoTemplate.md**: Manual for enabling Go Text Templates, using Sprig functions, and migrating from the legacy fasttemplate engine.
*   **Progressive-Syncs.md**: Documentation for alpha-stage features like `RollingSync` strategies and `Reverse` deletion orders.
*   **Controlling-Resource-Modification.md**: Covers safety features like dry-run mode, sync policies (`create-only`, `sync`), and `ignoreApplicationDifferences`.
*   **Application-Deletion.md**: Explains the lifecycle relationship between AppSets, Applications, and child resources, including finalizer behavior.
*   **Appset-Any-Namespace.md**: Guide for the beta feature allowing ApplicationSets to be managed outside the `argocd` namespace.
*   **Security.md**: Critical security considerations regarding RBAC, templated `project` fields, and secret exfiltration risks.
*   **applicationset-specification.md**: A placeholder/reference for the full YAML specification of the `ApplicationSet` resource.

## Code Changes That Would Require Documentation Updates
*   **CRD Schema Changes**: Any modification to the `ApplicationSet` spec or status (e.g., adding fields to `generators`, `template`, or `strategy`).
*   **Generator Logic**: Changes to how parameters are generated, especially directory scanning, file parsing, or SCM API interactions.
*   **New Generators**: Adding a new generator type requires a new dedicated file and updates to `Generators.md`.
*   **Templating Engine**: Changes to the rendering logic for Go Templates, Sprig function updates, or modifications to the `templatePatch` JSON/YAML merging logic.
*   **Controller CLI/Environment Flags**: Adding or removing flags (like `--enable-progressive-syncs` or `--enable-new-git-file-globbing`) or corresponding environment variables.
*   **RBAC & Multi-namespace Logic**: Changes to how the controller validates namespaces, checks `sourceNamespaces`, or handles cross-namespace secrets.
*   **SCM Provider Integrations**: Updates to supported SCM APIs (GitHub, GitLab, etc.), new authentication methods (App Tokens, IAM roles), or filter logic.
*   **Sync & Deletion Strategies**: Modifications to the `RollingSync` algorithm, health check monitoring, or the implementation of deletion finalizers.
*   **Webhook Handling**: Changes to the `/api/webhook` endpoint or the payload formats expected from Git providers.

## Key Technical Concepts
*   **Generators**: (List, Cluster, Git, Matrix, Merge, SCM, PR, Plugin, ClusterDecision).
*   **Parameter Substitution**: The `{{.variable}}` syntax used to inject data into templates.
*   **Go Templates & Sprig**: Advanced logic functions like `normalize`, `slugify`, `dig`, and `toJson`.
*   **templatePatch**: Advanced YAML/JSON patching for fields that cannot be easily templated as strings.
*   **RollingSync**: Sequential deployment strategy based on application labels and health status.
*   **Finalizers**: Specifically `resources-finalizer.argocd.argoproj.io` and its role in cascading deletions.
*   **Duck-typing**: Used by the Cluster Decision Resource generator to find cluster names in external CRDs.
*   **Greedy Globbing**: The historical behavior of Git file discovery and the `doublestar` fix.
*   **Strict Mode (tokenRef)**: Security setting requiring specific labels on secrets used for SCM authentication.

## Related Components
*   **Argo CD Application Controller**: Responsible for the actual sync/reconciliation of the generated Applications.
*   **Argo CD Repo Server**: Used by Git generators to fetch repository contents and files.
*   **Argo CD API Server**: Handles RBAC and API requests for ApplicationSet resources.
*   **SCM Providers**: External systems (GitHub, GitLab, Bitbucket, Azure DevOps, AWS) that provide data to SCM and PR generators.
*   **Kubernetes API**: The source for Cluster secrets and the target for Application resource creation.