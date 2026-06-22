# PROPOSALS/FEATURE-BOUNTIES Documentation Index

## Overview
This documentation area contains specific feature proposals for the Argo CD project that have been designated as "bounties." These documents serve as technical specifications for contributors, outlining desired functionality, implementation hints, and financial rewards for successful pull requests.

## Files Summary
*   **hide-annotations.md**: Proposes a new configuration option in `argocd-cm` to allow users to hide specific, often sensitive or redundant, annotations (such as OpenShift token secrets) from the Argo CD Web UI.

## Code Changes That Would Require Documentation Updates
*   **ConfigMap Schema Changes**: Any modifications to `argocd-cm` that change the key name `hide.secret.annotations` or its data structure (e.g., changing from a list to a map).
*   **UI Rendering Logic**: Changes to how the Argo CD Web UI processes and displays resource metadata, specifically annotations.
*   **Diffing Engine Updates**: Modifications to how `gitops-engine` or the Argo CD controller handles "ignored" or "hidden" fields during resource reconciliation and visualization.
*   **Feature Expansion**: If the ability to hide fields is expanded beyond annotations (e.g., hiding specific labels or entire status fields) or beyond Secret resources.
*   **Bounty Lifecycle**: Updates to the status of the proposal (e.g., marking it as "In Progress," "Completed," or "Closed") or changes to the award amount.
*   **Environment-Specific Defaults**: If the project decides to include certain annotations (like OpenShift-specific ones) in a default "hidden" list without requiring user configuration.

## Key Technical Concepts
*   **argocd-cm**: The central ConfigMap used for global configuration of Argo CD.
*   **hide.secret.annotations**: The specific configuration key proposed to manage annotation visibility.
*   **gitops-engine**: The core library used by Argo CD for diffing and resource management; specifically the `pkg/diff` package.
*   **last-applied-configuration**: The existing internal mechanism used as a reference point for hiding specific metadata fields.
*   **Resource Annotations**: Kubernetes metadata key-value pairs that are the target of this filtering feature.
*   **Argo CD Web UI**: The frontend interface where the filtering/hiding of these annotations is visually realized.

## Related Components
*   **Argo CD API Server**: Responsible for serving the configuration and resource data to the frontend.
*   **Argo CD Web UI (Frontend)**: The component responsible for filtering out the specified annotations before rendering the resource view.
*   **Argo CD Repo Server / Controller**: Components that interact with the `gitops-engine` for resource diffing.
*   **gitops-engine**: The external dependency that handles the logic of which fields are relevant for diffs and display.