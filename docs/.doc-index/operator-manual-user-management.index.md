# OPERATOR-MANUAL/USER-MANAGEMENT Documentation Index

## Overview
This documentation area provides comprehensive guides for managing user access to Argo CD. It covers the configuration of local accounts, administrative security, and Single Sign-On (SSO) integrations using either the bundled Dex OIDC provider (for SAML, LDAP, etc.) or direct OIDC provider connections.

## Files Summary

*   **index.md**: Serves as the primary landing page, detailing local account management, login rate limiting (throttling), and the foundational architecture for SSO via Dex or direct OIDC.
*   **okta.md**: Provides configuration steps for Okta integration using SAML (via Dex) or OIDC (direct), including group mapping and private deployment strategies.
*   **auth0.md**: Describes how to register Argo CD with Auth0 and configure RBAC using non-standard FQDN claims for group membership.
*   **zitadel.md**: A step-by-step guide for Zitadel integration using OIDC, covering Zitadel "Actions" to inject group claims into tokens.
*   **microsoft.md**: Detailed instructions for Microsoft Entra ID (Azure AD), supporting OIDC, SAML via Dex, Workload Identity Federation, and CLI-specific authentication.
*   **onelogin.md**: Explains the integration process for OneLogin OIDC, specifically how to map OneLogin "User Roles" to Argo CD groups.
*   **openunison.md**: Details integration with OpenUnison, focusing on portal "badges" and handling LDAP/AD Distinguished Names (DNs) in RBAC policies.
*   **identity-center.md**: Covers AWS IAM Identity Center (formerly AWS SSO) integration using SAML with Dex, including group attribute mapping workarounds.
*   **keycloak.md**: Comprehensive guide for Keycloak, including standard client authentication and PKCE (Proof Key for Code Exchange) for CLI-based login.
*   **google.md**: Outlines three methods for Google Workspace integration: standard OIDC, SAML, and the "Google Groups" method requiring a Service Account and Directory API access.

## Code Changes That Would Require Documentation Updates

*   **ConfigMap Schema Changes**: Any modifications to the data keys in `argocd-cm` (e.g., new `oidc.config` fields) or `argocd-rbac-cm` (e.g., changes to `policy.csv` syntax).
*   **Auth Provider Updates**: Adding support for a new Identity Provider or updating the versions/features of existing connectors.
*   **Dex Version Upgrades**: Updates to the bundled Dex server that change supported connectors, configuration syntax, or security defaults (e.g., `insecureEnableGroups`).
*   **Authentication Logic**: Changes to the Argo CD API server regarding how it handles callback URIs (`/auth/callback` vs `/api/dex/callback`) or token validation logic.
*   **CLI Authentication**: Changes to the `argocd login --sso` flow, such as port selection logic or the addition of new PKCE-related flags.
*   **Secret Handling**: Modifications to how Argo CD references sensitive data (e.g., the `$` prefix logic for Kubernetes secrets) or the introduction of new workload identity mechanisms.
*   **Rate Limiting/Security**: Changes to environment variables governing brute-force protection (e.g., `ARGOCD_SESSION_FAILURE_MAX_FAIL_COUNT`).
*   **RBAC Evaluation**: Changes to the RBAC engine that affect how group claims are parsed from JWT tokens or SAML assertions.

## Key Technical Concepts

*   **Core ConfigMaps**: `argocd-cm` (General config), `argocd-rbac-cm` (Role-based access control).
*   **Security Objects**: `argocd-secret` (Stores client secrets and admin passwords).
*   **Authentication Protocols**: OIDC (OpenID Connect), SAML 2.0, OAuth2, PKCE.
*   **Dex Connectors**: GitHub, Microsoft, Google, SAML, LDAP.
*   **OIDC Parameters**: `issuer`, `clientID`, `clientSecret`, `requestedScopes`, `requestedIDTokenClaims`, `logoutURL`.
*   **Dex-Specific Fields**: `ssoURL`, `caData`, `entityIssuer`, `redirectURI`, `groupsAttr`.
*   **Identity Provider Concepts**: Entra ID App Registrations, Okta Group Attribute Statements, Zitadel Actions, Google Directory API, Keycloak Client Scopes.
*   **CLI Commands**: `argocd login --sso`, `argocd account update-password`, `argocd account generate-token`.
*   **Throttling Variables**: `ARGOCD_SESSION_FAILURE_WINDOW_SECONDS`, `ARGOCD_MAX_CONCURRENT_LOGIN_REQUESTS_COUNT`.

## Related Components

*   **argocd-server**: The API server that handles direct OIDC authentication and RBAC enforcement.
*   **argocd-dex-server**: The bundled Dex instance that acts as an intermediary for non-OIDC identity providers.
*   **Argo CD CLI**: The command-line tool that performs SSO login via local callback listeners.
*   **RBAC Controller**: The system that matches claims (groups/email) from identity tokens against the `policy.csv` rules.
*   **Kubernetes Ingress**: Often critical for routing authentication callbacks to the correct internal service.