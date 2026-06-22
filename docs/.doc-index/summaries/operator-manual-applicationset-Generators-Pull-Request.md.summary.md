This documentation provides a comprehensive guide to the **Pull Request (PR) Generator** for Argo CD ApplicationSets. This generator enables "Preview Environments" by automatically creating Argo CD Applications when a pull request is opened in a supported Source Control Management (SCM) system and deleting them when the PR is closed.

### 1. Primary Purpose
The file documents how to configure the ApplicationSet controller to poll or receive webhooks from SCM providers (GitHub, GitLab, Gitea, Bitbucket, Azure DevOps) to dynamically discover open pull requests and generate Kubernetes applications based on those PRs.

### 2. Key Topics Covered
*   **Provider-Specific Configuration**: Detailed setup for GitHub, GitLab, Gitea, Bitbucket Server, Bitbucket Cloud, and Azure DevOps.
*   **Authentication**: Methods for connecting to SCM APIs using Secrets, Personal Access Tokens (PATs), GitHub Apps, and Basic/Bearer Auth.
*   **Filtering Logic**: How to narrow down PRs using labels (GitHub/GitLab/Azure), branch name regex (`branchMatch`), target branch regex (`targetBranchMatch`), and PR titles.
*   **Template Variables**: List of parameters (like `{{.branch}}`, `{{.head_sha}}`, and `{{.number}}`) available for use in the Application template.
*   **Webhook Integration**: Instructions on bypassing the default 30-minute polling interval by configuring SCM webhooks to trigger immediate refreshes.
*   **Security Precautions**: Warnings regarding the risks of allowing non-admins to create ApplicationSets with PR generators, specifically concerning Secret leaking and resource management.

### 3. Technical Keywords
*   **Core Resources**: `ApplicationSet`, `generator`, `pullRequest`, `template`.
*   **Configuration Keys**: `requeueAfterSeconds`, `tokenRef`, `appSecretName`, `api`, `insecure`, `caRef`, `labels`, `filters`.
*   **Filtering/Regex**: `branchMatch`, `targetBranchMatch`, `titleMatch`.
*   **Template Parameters**: `{{.number}}`, `{{.branch}}`, `{{.branch_slug}}`, `{{.head_sha}}`, `{{.head_short_sha_7}}`, `{{.author}}`, `{{.values.key}}`.
*   **Protocols/Standards**: RFC 1123 (DNS label standard), TLS, Webhooks (JSON content-type), Basic Auth, Bearer Token.

### 4. Target Audience
*   **DevOps Engineers**: Setting up automated preview environments for developers.
*   **Argo CD Administrators**: Managing the security and performance of the ApplicationSet controller.
*   **CI/CD Pipeline Developers**: Integrating git-flow workflows with Kubernetes deployments.

### 5. Related Concepts
*   **Argo CD ApplicationSets**: The parent controller that manages the lifecycle of multiple Applications.
*   **SCM Provider Generator**: A related generator that discovers repositories rather than PRs.
*   **Git Generator**: A generator that discovers files or folders within a Git repo.
*   **Kubernetes RBAC**: Related via the security warning about who can create/update ApplicationSets.
*   **Ephemeral Environments**: The architectural pattern this generator facilitates.

---

### AI Update Triggers
An AI system should suggest updates to this file if code changes occur in the following areas:
1.  **New SCM Support**: If a new provider (e.g., AWS CodeCommit or Woodpecker CI) is added to the ApplicationSet controller.
2.  **Filter Logic Changes**: If new filtering capabilities are added (e.g., filtering by PR author, date, or comment content).
3.  **Template Variable Expansion**: If the controller begins providing new metadata to the generator (e.g., `{{.pr_description}}` or `{{.base_ref_sha}}`).
4.  **Auth Refactoring**: If the way Argo CD handles SCM credentials changes (e.g., moving from individual Secrets to a centralized credential store).
5.  **Webhook Payload Changes**: If the ApplicationSet webhook receiver is updated to handle new events or different payload structures.
6.  **Slugification Logic**: If the logic for `branch_slug` (RFC 1123 compliance) is modified (e.g., changing the character limit or sanitization rules).