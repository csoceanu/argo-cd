This documentation defines the **SCM (Source Control Management) Provider Generator** for Argo CD ApplicationSets. It serves as the reference guide for automatically discovering repositories across various Git platforms to dynamically create Argo CD Applications.

### 1. Primary Purpose
The SCM Provider generator automates the "discovery" phase of a GitOps pipeline. Instead of manually defining an Application for every repository, this generator queries the API of an SCM provider (like GitHub or GitLab) to find repositories matching specific criteria and automatically generates Applications for them. It is specifically designed for organizations following a microservices pattern where services are spread across many repositories.

### 2. Key Topics Covered
*   **Provider-Specific Configurations**: Detailed setup for GitHub, GitLab, Gitea, Bitbucket Server (API 1.0), Bitbucket Cloud (API 2.0), Azure DevOps, and AWS CodeCommit.
*   **Authentication Mechanisms**: Usage of Kubernetes Secrets for PATs (Personal Access Tokens), OAuth tokens, GitHub Apps, and Basic Auth.
*   **Repository Filtering**: Logic for including/excluding repositories based on name regex, existence of specific files, or repository labels/topics.
*   **Security Constraints**: Requirements for admin-level permissions to prevent secret leaking or unauthorized resource management.
*   **Advanced Networking**: Handling self-signed TLS certificates and custom Root CAs for on-premises SCM instances.
*   **Templating**: How to use the discovered metadata (repo URL, branch name, etc.) to populate the Application spec.

### 3. Technical Keywords
*   **CRD Fields**: `scmProvider`, `cloneProtocol`, `tokenRef`, `appSecretName`, `allBranches`, `repositoryMatch`, `pathsExist`, `labelMatch`.
*   **Supported Platforms**: `github`, `gitlab`, `gitea`, `bitbucketServer`, `azureDevOps`, `bitbucket`, `awsCodeCommit`.
*   **Template Parameters**: `{{.organization}}`, `{{.repository}}`, `{{.url}}`, `{{.branch}}`, `{{.sha}}`, `{{.labels}}`, `{{.branchNormalized}}`.
*   **Configuration**: `argocd-cmd-params-cm`, `ARGOCD_APPLICATIONSET_CONTROLLER_SCM_ROOT_CA_PATH`, `goTemplate`.

### 4. Target Audience
*   **Platform Engineers**: Responsible for setting up automated onboarding of new services into Argo CD.
*   **DevOps Engineers**: Managing large-scale Kubernetes deployments across multiple repositories.
*   **Security Auditors**: Reviewing how SCM tokens are handled and how repo discovery is restricted.

### 5. Related Concepts
*   **Argo CD ApplicationSets**: The parent controller that evaluates these generators.
*   **Pull Request Generator**: A sibling generator often used alongside SCM discovery to handle ephemeral environments.
*   **GitOps Layout Patterns**: Specifically the "App-of-Apps" or "Generator" patterns.
*   **Kubernetes RBAC**: Used to control who can create ApplicationSets (due to the security implications of generators).

---

### Triggering Documentation Updates (AI Analysis Guide)
Update this file if code changes occur in the following areas:

1.  **Provider Support**: A new SCM provider is added to the `applicationset` controller source code.
2.  **API Schema Changes**: New fields are added to the `SCMProviderGenerator` Go struct (e.g., a new filtering logic like `topicMatch` or a new auth field).
3.  **Template Variables**: If the controller begins exporting new metadata for use in templates (e.g., adding `{{.repository_description}}`).
4.  **Security Defaults**: Changes in how the ApplicationSet controller handles tokens, certificate validation, or RBAC requirements for generators.
5.  **Environment Variables**: Changes to the global configuration variables like those managing SCM Root CA paths or rate limiting.
6.  **AWS/Cloud Integrations**: Changes in how IAM roles or regional discovery are handled for cloud-native Git services like CodeCommit.