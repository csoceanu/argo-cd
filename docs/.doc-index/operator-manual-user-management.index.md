# OPERATOR-MANUAL/USER-MANAGEMENT Documentation Index

## Overview
This documentation area covers the configuration and management of user authentication and authorization within Argo CD. It provides detailed instructions for managing local accounts, setting up Single Sign-On (SSO) using the bundled Dex server or native OIDC, and mapping external group memberships to Argo CD RBAC roles across various Identity Providers (IdPs).

## Files Summary
*   **auth0.md**: Provides configuration steps for integrating Auth0 as an OIDC provider, specifically detailing custom FQDN claims for group-based RBAC.
*   **google.md**: Outlines three integration methods for Google Workspace: OIDC via Dex, SAML via Dex, and a specialized Google connector for Dex that uses service accounts to fetch Google Groups.
*   **identity-center.md**: Details how to configure AWS IAM Identity Center (formerly AWS SSO) using SAML via Dex, including workarounds for group attribute mapping.
*   **index.md**: The foundational overview for user management, covering local account creation, password management, rate limiting, secret referencing, and general SSO architecture (Dex vs. native OIDC).
*   **keycloak.md**: Explains Keycloak integration via OIDC, contrasting standard client authentication with PKCE flows required for CLI SSO support.
*   **microsoft.md**: A comprehensive guide for Microsoft Entra ID (Azure AD), covering OIDC native integration, SAML via Dex, the Microsoft Dex connector, and Azure Workload Identity.
*   **okta.md**: Describes Okta integration using SAML (with Dex) or native OIDC, including specific instructions for private deployments where only the callback is public.
*   **onelogin.md**: Instructions for connecting OneLogin via OIDC, focusing on mapping OneLogin UserRoles to the "groups" claim for authorization.
*   **openunison.md**: Details integration with the OpenUnison portal to provide unified access to both Kubernetes and Argo CD with automated "badge" logins.
*   **zitadel.md**: A step-by-step guide for Zitadel integration, covering project/application setup, custom actions for group claims, and role mapping.

## Code Changes That Would Require Documentation Updates
*   **Authentication Logic**: Changes to how Argo CD handles OIDC tokens, SAML assertions, or the internal handshake with the bundled Dex server.
*   **ConfigMap Schema**: Modifications to the structure or supported keys within `argocd-cm` (e.g., `oidc.config`, `dex.config`, `accounts.*`) or `argocd-rbac-cm`.
*   **Local Account Management**: Updates to the `argocd account` CLI commands, password hashing mechanisms, or account capability flags (`apiKey`, `login`).
*   **Secret Handling**: Changes to the logic that dereferences Kubernetes secrets using the `$` syntax in ConfigMaps.
*   **CLI Login Flow**: Updates to the SSO login process, such as changing the default callback port (`8085`), PKCE implementation, or browser redirection logic.
*   **RBAC & Scopes**: Changes to how Argo CD parses "scopes" or "claims" from JWT tokens to determine user permissions.
*   **Security & Throttling**: Updates to login rate limiting, session failure windows, or concurrent login request limits.
*   **Dependency Updates**: Significant version upgrades to the bundled Dex server that introduce new connectors or deprecate existing configuration fields.

## Key Technical Concepts
*   **ConfigMaps**: `argocd-cm` (main config), `argocd-rbac-cm` (permissions).
*   **Protocols**: OIDC (OpenID Connect), SAML 2.0, OAuth2, PKCE (Proof Key for Code Exchange).
*   **Dex Connectors**: GitHub, SAML, OIDC, Microsoft, Google, LDAP.
*   **Authentication Terms**: Issuer URI, Client ID, Client Secret, Redirect URI (Callback), Scopes, Claims, ID Token, UserInfo Endpoint.
*   **Local Management**: `apiKey`, `login` capabilities, `argocd account` commands, password bcrypt hashes.
*   **RBAC Mapping**: `policy.csv`, `g` (group) entries, `role:admin`, `role:readonly`, `scopes` configuration.
*   **Secret Referencing**: `$<secret_name>:<key>` syntax for sensitive SSO credentials.
*   **Environment Variables**: `ARGOCD_SESSION_FAILURE_MAX_FAIL_COUNT`, `ARGOCD_MAX_CONCURRENT_LOGIN_REQUESTS_COUNT`.

## Related Components
*   **argocd-server**: The primary API and Web UI component that handles authentication requests.
*   **argocd-dex-server**: The bundled instance of Dex used for multi-protocol identity brokerage.
*   **Argo CD CLI**: The command-line tool used for `argocd login --sso` and account management.
*   **Kubernetes RBAC/Secrets**: The underlying infrastructure used for storing configuration and sensitive credentials.
*   **Identity Providers (IdPs)**: External systems like Okta, Entra ID, Keycloak, Google, and Auth0.