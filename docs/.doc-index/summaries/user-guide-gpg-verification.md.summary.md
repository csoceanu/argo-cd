This documentation serves as the authoritative guide for configuring and enforcing GnuPG (GPG) signature verification within ArgoCD.

### 1. Primary Purpose
The file documents the end-to-end process of ensuring that ArgoCD only synchronizes with Git commits or tags that have been cryptographically signed by trusted GnuPG keys. It explains how to transition from a "no-verification" state to a secure, enforced state at the project level.

### 2. Key Topics Covered
*   **Enforcement Logic**: How ArgoCD handles signed vs. unsigned commits and the resulting `ResourceComparison` errors.
*   **Signature Targets**: Differences in verification behavior between lightweight tags (verifies commit) and annotated tags (verifies the tag object itself).
*   **Key Management**: Three methods for importing public keys: CLI, Web UI, and Declarative (ConfigMaps).
*   **Project Configuration**: Associating specific GPG key IDs with an `AppProject` to enforce verification.
*   **RBAC**: Configuring permissions specifically for the `gpgkeys` resource.
*   **Architecture & Troubleshooting**: Understanding how the `argocd-repo-server` maintains a transient keyring and how to disable the feature globally via environment variables.

### 3. Technical Keywords
*   **Environment Variables**: `ARGOCD_GPG_ENABLED`
*   **Configuration Resources**: `argocd-gpg-keys-cm` (ConfigMap), `AppProject` (CRD)
*   **CRD Fields**: `signatureKeys`, `keyID`
*   **CLI Commands**: `argocd gpg add/list/get/rm`, `argocd proj add-signature-key`, `argocd proj set --signature-keys`
*   **RBAC Resource**: `gpgkeys`
*   **Internal Paths**: `/app/config/gpg/keys` (GnuPG home in repo-server)
*   **Components**: `argocd-server`, `argocd-repo-server`, `argocd-application-controller`, `argocd-applicationset-controller`

### 4. Target Audience
*   **Security Administrators**: Responsible for defining trust models and ensuring code integrity.
*   **DevOps/Platform Engineers**: Responsible for configuring ArgoCD projects and managing the infrastructure (ConfigMaps, Deployment manifests).
*   **Developers**: To understand why a sync might fail due to signature requirements.

### 5. Related Concepts
*   **Git Security**: Specifically GPG commit/tag signing (`git commit -S`, `git tag -s`).
*   **ArgoCD Multi-tenancy**: Using `AppProjects` to isolate security requirements between different teams.
*   **GitOps Workflow Integrity**: Ensuring that the "Source of Truth" has not been tampered with by unauthorized actors.

---

### Maintenance Guide: When to update this file
An AI system or developer should update this documentation if any of the following code changes occur:

1.  **Logic Changes in Signature Verification**: If the strategy for verifying tags vs. commits changes, or if support is added for Helm repository signatures (currently unsupported).
2.  **Environment Variable Modifications**: If the `ARGOCD_GPG_ENABLED` flag is renamed, deprecated, or if its requirement spreads to new components.
3.  **CLI/API Updates**: If new sub-commands are added to `argocd gpg` or if the `argocd proj` flags for signature management are modified.
4.  **Schema Changes in CRDs**: If the `AppProject` spec is updated to include more properties under `signatureKeys` beyond the current `keyID`.
5.  **RBAC Scope Changes**: If the resource name for GPG keys is changed from `gpgkeys` to something else in the policy enforcer.
6.  **Internal Filesystem Changes**: If the `argocd-repo-server` changes the location of the GPG keyring or the method by which it consumes the `argocd-gpg-keys-cm`.
7.  **Trust Model Updates**: If ArgoCD moves away from its "simple trust" model (trusting any imported key) to a more complex Web of Trust or CA-based model.