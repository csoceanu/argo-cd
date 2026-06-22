# PROPOSALS/FEATURE-BOUNTIES Documentation Index

## Overview
This documentation area outlines proposed features for Argo CD that have been designated for community bounties. Each proposal serves as a technical specification and a call-to-action for contributors, detailing specific problems, proposed configuration changes, and implementation references to guide the development of rewarded features.

## Files Summary
*   **hide-annotations.md**: Proposes a new configuration option in `argocd-cm` to allow administrators to hide specific sensitive or high-noise annotations (such as OpenShift token secrets) from being displayed in the Argo CD Web UI.

## Code Changes That Would Require Documentation Updates
*   **Configuration Schema Changes**: Any modification to the `argocd-cm` structure, specifically renaming or restructuring the `hide.secret.annotations` key.
*   **UI Rendering Logic**: Changes to how the Argo CD Web UI fetches or displays resource metadata and annotations.
*   **Diffing Engine Updates**: Modifications to the `gitops-engine` or the Argo CD controller logic that filters metadata during the reconciliation or diffing process.
*   **Scope Expansion**: If the feature is updated to support hiding metadata beyond just annotations (e.g., labels) or resources beyond just Secrets.
*   **Bounty Program Adjustments**: Changes to the award amount or the criteria for "accepted" pull requests related to this feature.
*   **Alternative Implementation**: If the community decides to implement this via a different mechanism, such as server-side resource customization or RBAC-based field masking.

## Key Technical Concepts
*   **argocd-cm**: The primary ConfigMap used for global Argo CD settings.
*   **hide.secret.annotations**: The specific configuration key proposed for defining which annotations to mask.
*   **Resource Annotations**: Kubernetes metadata used to store non-identifying information, often targeted for hiding if they contain sensitive tokens or redundant data.
*   **last-applied-configuration**: The standard Kubernetes annotation that serves as a precedent for field hiding in Argo CD.
*   **gitops-engine**: The underlying library used by Argo CD for diffing and resource state management.
*   **openshift.io/token-secret.value**: A specific example of a sensitive annotation that prompted this proposal.

## Related Components
*   **Argo CD Web UI**: The frontend component responsible for displaying resource details to users.
*   **Argo CD API Server**: The backend component that reads `argocd-cm` and serves resource data to the UI.
*   **Argo CD Controller**: Manages the state of applications and interacts with the `gitops-engine`.
*   **GitOps Engine**: Provides the core logic for comparing live state versus desired state and handling metadata filtering.