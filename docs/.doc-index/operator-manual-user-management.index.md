# OPERATOR-MANUAL/USER-MANAGEMENT Documentation Index

## Overview
This documentation area provides comprehensive guides for managing user authentication and authorization within Argo CD. It covers the configuration of local accounts, security hardening (rate limiting, disabling admin), and deep-dive integration steps for various external Identity Providers (IdPs) using OIDC and SAML protocols, primarily through the bundled Dex server or direct OIDC connections.

## Files Summary
*   **index.md**: The foundational overview covering local user management, account capabilities (apiKey/login), failed login rate limiting, and general SSO architecture (Dex vs. Direct OIDC).
*   **okta.md**: Instructions for Okta integration via SAML (with Dex) or OIDC (Direct), including private deployment patterns and group-to-role mapping.
*   **auth0.md**: Configuration for Auth0 OIDC, focusing on FQDN-based group claims and secure client secret storage.
*   **zitadel.md**: A detailed walkthrough for Zitadel integration, including the use of "Zitadel Actions" to inject roles into ID tokens for Argo CD RBAC.
*   **microsoft.md**: Comprehensive guide for Microsoft Entra ID (formerly Azure AD), covering OIDC, SAML, Workload Identity Federation, and CLI-specific redirect settings.
*   **onelogin.md**: Setup for OneLogin OIDC, specifically detailing how to map OneLogin "User Roles" to the OIDC groups claim.
*   **keycloak.md**: Instructions for Keycloak integration, contrasting Client Secret authentication with PKCE flows (required for CLI login) and group mapper configuration.
*   **google.md**: Covers three paths for Google Workspace: OIDC (Dex), SAML (Dex - deprecated), and a specialized OIDC + Google Groups method using Google Directory API and Service Accounts.
*   **openunison.md**: Integration guide for OpenUnison, focusing on portal "badges" for single-point access and handling LDAP Distinguished Names (DNs) in RBAC.
*   **identity-center.md**: Guide for AWS IAM Identity Center (AWS SSO) integration using SAML via Dex, including workarounds for group attribute mapping.

## Code Changes That Would Require Documentation Updates
*   **Authentication Flow Logic**: Any changes to the OIDC/SAML handshake logic in `argocd-server` or the bundled `dex` configuration.
*   **CLI Login Mechanics**: Modifications to the local callback server (default port 8085), PKCE implementation, or the `--sso` flag behavior.
*   **ConfigMap Schema Updates**: Changes to the structure or supported keys within `argocd-cm` (specifically `oidc.config` and `dex.config`) or `argocd-rbac-cm`.
*   **RBAC Engine Alterations**: Changes to how scopes are parsed, how group claims are matched, or how `policy.csv` interprets subjects (e.g., changes to handling special characters or DNs).
*   **Secret Management**: Updates to how Argo CD references sensitive data via the `$` prefix or changes to the `argocd-secret` lookup logic.
*   **Security Defaults**: Altering default session timeouts, rate-limiting environment variables (`ARGOCD_SESSION_FAILURE_MAX_FAIL_COUNT`), or password hashing algorithms.
*   **Resource Labels**: Changing the required labels (e.g., `app.kubernetes.io/part-of: argocd`) for custom secrets used in SSO.
*   **New Protocol Support**: Implementation of new authentication protocols or major version upgrades of the bundled Dex server that introduce new connectors.

## Key Technical Concepts
*   **OIDC (OpenID Connect)**: The primary protocol for modern identity integration.
*   **SAML (Security Assertion Markup Language)**: XML-based authentication, often requiring Dex as a bridge.
*   **Dex**: The bundled identity service that acts as a connector/proxy to various IdPs.
*   **RBAC (Role-Based Access Control)**: Mapping external groups/claims to internal Argo CD roles (`admin`, `readonly`).
*   **PKCE (Proof Key for Code Exchange)**: Security extension for OIDC, critical for CLI/Desktop application logins.
*   **Workload Identity Federation**: Azure-specific method for passwordless authentication between Argo CD and Entra ID.
*   **Claims Mapping**: The process of extracting `groups`, `email`, or `roles` from an ID token.
*   **Domain-Wide Delegation**: Required for Google Workspace to allow Argo CD to fetch group memberships via the Directory API.
*   **Rate Limiting/Throttling**: Protecting the `admin` and local accounts from brute-force attacks via environment variables.

## Related Components
*   **argocd-server**: Handles the API and UI authentication logic and session management.
*   **argocd-dex-server**: The component running the Dex OIDC provider.
*   **Argo CD CLI**: Client-side tool that performs SSO handshakes via local browser redirects.
*   **ConfigMaps (`argocd-cm`, `argocd-rbac-cm`)**: Primary configuration storage for auth settings.
*   **Kubernetes Secrets (`argocd-secret`)**: Storage for IdP client secrets and local account passwords.
*   **Ingress Controller**: Manages the public endpoints and callback URLs required for SSO redirects.