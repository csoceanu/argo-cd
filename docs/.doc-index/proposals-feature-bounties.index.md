# PROPOSALS/FEATURE-BOUNTIES Documentation Index

## Overview
This documentation area outlines proposed feature enhancements for Argo CD that are eligible for community bounties. It provides technical specifications, implementation strategies, and configuration examples to guide developers in contributing specific requested features to the project.

## Files Summary
* **hide-annotations.md**: Proposes a feature to allow administrators to hide specific sensitive or redundant annotations (such as OpenShift tokens) from the Argo CD Web UI via settings in the `argocd-cm` ConfigMap.

## Code Changes That Would Require Documentation Updates
* **Configuration Schema Changes**: Any modifications to the `argocd-cm` ConfigMap structure or the addition of new keys for resource visibility.
* **UI Rendering Logic**: Changes to how the Argo CD frontend (Web UI) processes and displays resource metadata, specifically annotations.
* **API Filtering**: Updates to the backend API responsible for sanitizing or filtering resource manifests before they are transmitted to the client.
* **Diffing Engine Modifications**: Changes to the underlying `gitops-engine` or internal diffing logic that controls which metadata fields are ignored or masked during resource comparison.
* **Secret Handling Policies**: Changes to the default security posture regarding how sensitive metadata in Kubernetes Secrets is exposed in the interface.
* **Platform-Specific Integrations**: Updates to how Argo CD handles specific third-party annotations, such as those from OpenShift or other cloud providers.

## Key Technical Concepts
* **argocd-cm**: The central ConfigMap used for managing Argo CD global configurations.
* **hide.secret.annotations**: The proposed configuration key for defining a list of annotations to be suppressed in the UI.
* **last-applied-configuration**: The standard Kubernetes annotation used as a reference for existing hiding/masking logic.
* **gitops-engine**: The core library used by Argo CD for manifest comparison and state synchronization.
* **Resource Annotations**: Kubernetes metadata key-value pairs that this feature aims to filter.
* **openshift.io/token-secret.value**: A specific example of a sensitive annotation that necessitates this feature.
* **Feature Bounty**: The programmatic framework for rewarding external contributors for implementing specific enhancements.

## Related Components
* **Argo CD Web UI**: The frontend component where the hidden annotations would be filtered from view.
* **Argo CD API Server**: The backend component that would likely implement the filtering logic before serving resource data.
* **Argo CD Repo Server**: The component responsible for manifest generation and initial metadata handling.
* **GitOps Engine**: The external dependency that handles low-level resource diffing and metadata exclusion.
* **Config Management**: The subsystem responsible for watching and applying changes from `argocd-cm`.