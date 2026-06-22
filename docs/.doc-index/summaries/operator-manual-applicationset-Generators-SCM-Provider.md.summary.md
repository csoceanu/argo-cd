This documentation provides a technical guide for the **SCM Provider Generator** within Argo CD ApplicationSets. This generator enables the automatic discovery of Git repositories across various Source Control Management (SCM) platforms to automate the creation of Argo CD Applications at scale.

### 1. Primary Purpose
The file documents how to use the **SCM Provider Generator** to scan SCM organizations/groups and automatically generate Application resources for discovered repositories. It is designed to support GitOps patterns where microservices are distributed across many individual repositories.

### 2. Key Topics Covered
*   **Provider Configurations**: Detailed setup for GitHub, GitLab, Gitea, Bitbucket Server (API 1.0), Bitbucket Cloud (API 2.0), Azure DevOps, and AWS CodeCommit.
*   **Authentication**: Methods for connecting to providers, including Personal Access Tokens (PAT), Basic Auth, Bearer Tokens, GitHub Apps, and AWS IAM roles.
*   **Filtering Logic**: Mechanisms to include/exclude repositories based on name matches, branch matches, labels/topics, or the presence (or absence) of specific files/paths.
*   **Template Parameters**: The variables (e.g., repository name, URL, SHA) available for use in the ApplicationSet template.
*   **Security & TLS**: Guidance on RBAC implications and handling self-signed certificates (via `insecure` flags or `caRef`).

### 3. Technical Keywords
*   **CRD/Kind**: `ApplicationSet` (`argoproj.io/v1alpha1`)
*   **Generator Field**: `scmProvider`
*   **Provider Specifics**: `github`, `gitlab`, `gitea`, `bitbucketServer`, `bitbucket`, `azureDevOps`, `awsCodeCommit`.
*   **Configuration Options**: `cloneProtocol` (ssh/https), `allBranches`, `api` (for self-hosted), `tokenRef`, `appSecretName`, `caRef`, `insecure`.
*   **Filtering**: `repositoryMatch`, `pathsExist`, `pathsDoNotExist`, `labelMatch`, `branchMatch`.
*   **Template Variables**: `{{.organization}}`, `{{.repository}}`, `{{.url}}`, `{{.branch}}`, `{{.sha}}`, `{{.labels}}`, `{{.branchNormalized}}`.
*   **Advanced**: `goTemplate: true`, `values` field interpolation.

### 4. Target Audience
*   **Platform Engineers**: Responsible for designing automated application onboarding workflows.
*   **Argo CD Administrators**: Setting up the infrastructure and security for ApplicationSets.
*   **DevOps Practitioners**: Implementing GitOps at scale across large organizations with many repositories.

### 5. Related Concepts
*   **GitOps Layout Patterns**: Strategies for managing microservices across multiple repos.
*   **Argo CD ApplicationSets**: The overarching controller that manages these generators.
*   **RBAC & Security**: Crucial for ensuring that SCM scanning doesn't leak secrets or allow unauthorized repo management.
*   **Repository Credentials**: Centralized credential management in Argo CD.

---

### Update Triggers for AI Systems
This documentation should be updated if any of the following occur in the codebase:
*   **New SCM Provider Support**: If a new provider (e.g., Woodpecker, Codeberg) is added.
*   **Schema Changes**: If new fields are added to the `scmProvider` spec (e.g., adding pagination settings or new filtering logic like `commitAge`).
*   **Template Parameter Changes**: If new variables are exposed to the template (e.g., exposing repository descriptions or default reviewer lists).
*   **Authentication Updates**: If new auth methods are implemented (e.g., OIDC for certain providers).
*   **Breaking Changes**: If default behaviors change (e.g., changing the default `cloneProtocol` or `allBranches` logic).