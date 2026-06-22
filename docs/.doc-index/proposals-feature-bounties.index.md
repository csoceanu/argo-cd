# PROPOSALS/FEATURE-BOUNTIES Documentation Index

## Overview
This documentation area outlines proposed feature enhancements for Argo CD that are eligible for community bounties. Its purpose is to define specific technical requirements, proposed solutions, and implementation paths for contributors to earn awards while improving the project.

## Files Summary
*   **hide-annotations.md**: Proposes a feature to allow administrators to hide specific, sensitive, or noisy annotations (such as OpenShift tokens) from appearing in the Argo CD Web UI via configuration in the `argocd-cm` ConfigMap.

## Code Changes That Would Require Documentation Updates
*   **ConfigMap Schema Changes**: Any modification to the `argocd-cm` structure, specifically regarding how resource metadata filtering is configured.
*   **Configuration Key Renaming**: If the proposed key `hide.secret.annotations` is renamed or replaced by a more generic filtering mechanism.
*   **Feature Scope Expansion**: If the ability to hide metadata is extended beyond annotations to include labels or other resource fields.
*   **UI Rendering Logic**: Changes to how the Argo CD Web UI processes and displays resource metadata that would render this specific filtering logic obsolete or different.
*   **API/Diffing Logic**: Updates to the backend diffing engine (e.g., in `gitops-engine` or Argo CD's controller) that change how annotations are stripped before being sent to the frontend.
*   **Bounty Status/Amount**: Changes to the award amount or the acceptance status of the proposal.

## Key Technical Concepts
*   **argocd-cm**: The primary ConfigMap used to manage Argo CD settings.
*   **hide.secret.annotations**: The specific configuration key proposed to list annotations for exclusion.
*   **Annotations**: Kubernetes metadata fields used to store non-identifying information.
*   **GitOps Engine**: The underlying library Argo CD uses for resource diffing and state comparison.
*   **Resource Diffing**: The process of comparing the desired state in Git with the live state in the cluster.
*   **openshift.io/token-secret.value**: A specific example of a sensitive annotation that motivated this proposal.
*   **last-applied-configuration**: An existing annotation that is already hidden, serving as a technical precedent for this feature.

## Related Components
*   **Argo CD Web UI**: The frontend component where the annotations are currently displayed and should be hidden.
*   **Argo CD API Server**: The component responsible for reading `argocd-cm` and serving resource data to the UI.
*   **GitOps Engine (pkg/diff)**: The external dependency where the core logic for filtering metadata during diffing resides.
*   **Application Controller**: The component that manages resource state and interacts with the Kubernetes API.