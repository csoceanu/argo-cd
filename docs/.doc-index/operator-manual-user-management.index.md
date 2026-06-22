# OPERATOR-MANUAL/USER-MANAGEMENT Documentation Index

## Overview
This documentation area provides comprehensive guides and reference configurations for managing user identity, authentication, and authorization within Argo CD. It details the setup of local accounts for small teams and automation, alongside complex Single Sign-On (SSO) integrations using OIDC and SAML via the bundled Dex server or native OIDC providers.

## Files Summary
*   **index.md**: The foundational guide covering local user management (passwords, tokens, rate limiting), architectural differences between Dex and native OIDC, and secure handling of SSO secrets.
*   **okta.md**: Provides configuration steps for Okta integration using either SAML (via Dex) or native OIDC, including group attribute mapping and private deployment scenarios.
*   **microsoft.md**: A detailed guide for Microsoft Entra ID (Azure AD), covering native OIDC, SAML, and Microsoft-specific Dex connectors, including Azure Workload Identity and CLI login.
*   **google.md**: Outlines three methods for Google Workspace integration, with specific emphasis on using Google Service Accounts to sync Google Group memberships for RBAC.
*   **keycloak.md**: Instructions for Keycloak integration using standard Client Secret authentication or PKCE (Proof Key for Code Exchange) to support CLI-based SSO.
*   **zitadel.md**: A specialized guide for Zitadel integration involving custom "Actions" (JavaScript) to inject user roles into ID tokens for Argo CD consumption.
*   **auth0.md**: Covers Auth0 OIDC registration, focusing on the requirement for FQDN-based group claims and RBAC mapping.
*   **onelogin.md**: Setup for OneLogin OIDC, detailing how to map OneLogin User Roles to the `groups` claim in the JWT.
*   **identity-center.md**: Guidance for AWS IAM Identity Center (formerly AWS SSO) using SAML via Dex, including workarounds for group attribute mapping.
*   **openunison.md**: Integration with the OpenUnison portal for unified access, focusing on LDAP/Active Directory DN (Distinguished Name) group handling.

## Code Changes That Would Require Documentation Updates
*   **ConfigMap Schema Changes**: Modifications to the structure or supported keys within `argocd-cm` (especially `oidc.config` and `dex.config`) or `argocd-rbac-cm`.
*   **Dex Version Upgrades**: Upgrading the bundled Dex server, which may introduce new connectors, deprecate SAML features, or change configuration syntax.
*   **CLI Authentication Logic**: Changes to `argocd login` behavior, port handling for SSO (e.g., `--sso-port`), or the introduction of new authentication flags.
*   **RBAC Logic & Scopes**: Altering how Argo CD parses scopes (e.g., `groups`, `email`) or changes to the default RBAC roles (`admin`, `readonly`).
*   **Secret Referencing**: Changes to how Argo CD retrieves sensitive data from Kubernetes Secrets using the `$` prefix syntax.
*   **Local Account Management**: Updates to the `argocd account` command suite or changes to the local user capability settings (`apiKey`, `login`).
*   **Security/Throttling Policies**: Modifications to environment variables governing login rate limiting (`ARGOCD_SESSION_FAILURE_MAX_FAIL_COUNT`, etc.).
*   **New Identity Provider Support**: Adding native support for new IdPs that bypass Dex.

## Key Technical Concepts
*   **Authentication Protocols**: OIDC (OpenID Connect), SAML (Security Assertion Markup Language), OAuth2, PKCE.
*   **Argo CD Components**: `argocd-server`, `argocd-dex-server`, `argocd-cm`, `argocd-rbac-cm`, `argocd-secret`.
*   **JWT Claims**: `sub`, `email`, `groups`, `aud` (audience), `iss` (issuer).
*   **Configuration Fields**: `issuer`, `clientID`, `clientSecret`, `requestedScopes`, `redirectURI`, `caData`, `policy.csv`.
*   **Security Concepts**: Domain-Wide Delegation (Google), Workload Identity (Azure), Federated Credentials, Token Mappers.
*   **CLI Commands**: `argocd login --sso`, `argocd account update-password`, `argocd account generate-token`.

## Related Components
*   **Dex (Bundled)**: The multi-connector identity provider used by Argo CD to bridge various IdPs.
*   **Argo CD RBAC System**: The policy engine that consumes group claims from this documentation to enforce access control.
*   **Kubernetes RBAC**: Necessary for managing the Secrets and ConfigMaps described in these guides.
*   **Ingress Controllers**: Often involved in exposing the `/api/dex/callback` or `/auth/callback` endpoints for SSO.
*   **External IdPs**: Okta, Microsoft Entra ID, Google Workspace, Keycloak, Auth0, OneLogin, Zitadel, AWS Identity Center.