# DEVELOPER-GUIDE/EXTENSIONS Documentation Index

## Overview
This documentation area provides instructions for extending Argo CD through custom user interface elements and backend proxy services. It defines the architecture for injecting JavaScript-based UI components and configuring the Argo CD API server to act as a secure reverse-proxy for external backend services.

## Files Summary
*   **ui-extensions.md**: Explains how to deliver and register custom React components to add resource tabs, sidebar items, status panel elements, and top bar actions within the Argo CD web interface.
*   **proxy-extensions.md**: Describes the configuration and security protocols for the API server's reverse-proxy functionality, allowing UI extensions to communicate with external backend services using Argo CD's authentication and RBAC.

## Code Changes That Would Require Documentation Updates
*   **UI API Surface**: Modifications to the global `extensionsAPI` object or the method signatures for `registerResourceExtension`, `registerSystemLevelExtension`, `registerStatusPanelExtension`, or `registerTopBarActionMenuExt`.
*   **Extension Loading Logic**: Changes to how `argocd-server` scans the filesystem (e.g., changing the `/tmp/extensions` path) or the regex used to identify extension files (`^extension(.*)\.js$`).
*   **UI Properties/Models**: Updates to the `Application`, `State` (resource), or `ApplicationTree` interfaces in `models.ts` that are passed as props to UI extensions.
*   **Proxy Configuration Schema**: Changes to the YAML structure in `argocd-cm` for `extension.config`, including timeout settings, header management, or cluster matching logic.
*   **Feature Flag Management**: Promoting the proxy extension feature from Beta to GA or changing the command-line parameter key `server.enable.proxy.extension`.
*   **Header Protocol**: Altering the mandatory incoming headers (`Argocd-Application-Name`, `Argocd-Project-Name`) or the identity headers injected into outgoing proxy requests (`Argocd-User-Id`, `Argocd-User-Groups`).
*   **RBAC & Security**: Modifications to how the API server validates permissions for the `extensions` resource or how it sanitizes sensitive headers (Cookie/Authorization) before proxying.
*   **UI Rendering Lifecycle**: Changes to how Argo CD loads external scripts or handles React global variables/externals during initial page rendering.

## Key Technical Concepts
*   **extensionsAPI**: The global JavaScript object used by extensions to register themselves with the Argo CD UI.
*   **Resource Tab Extension**: Custom tabs added to the resource sliding panel, filtered by Group and Kind.
*   **System Level Extension**: Custom pages accessible via new icons in the main Argo CD sidebar.
*   **Flyout Widget**: A sliding panel UI component used to display detailed information for status panel or action menu extensions.
*   **Proxy Extension**: A backend configuration allowing the API server to forward requests to external services under the `/extensions/<name>` endpoint.
*   **argocd-cm (extension.config)**: The primary configuration point for defining backend service URLs, timeouts, and custom headers.
*   **Application-Destination Mapping**: The logic used by the proxy to route requests to specific backends based on an application's target cluster name or server URL.
*   **Header Injection**: The process of adding Argo CD user identity (ID, Groups, Username) to requests forwarded to backend services.

## Related Components
*   **argocd-server**: The primary component responsible for serving the UI, loading JS extensions, and executing the proxy logic.
*   **Argo CD Web UI**: The React-based frontend that executes the extension scripts and renders the custom components.
*   **argocd-cm / argocd-cmd-params-cm**: ConfigMaps used to enable and configure extension behaviors.
*   **Argo CD RBAC Subsystem**: Validates user permissions against the `extensions` resource before allowing proxy access.
*   **Application Controller**: Manages the application resources that provide context (labels, destination) to extensions.