# DEVELOPER-GUIDE/EXTENSIONS Documentation Index

## Overview
This documentation provides a comprehensive guide for developers to extend Argo CD’s functionality through UI and Proxy extensions. It covers the registration of custom React components within the web interface and the configuration of the Argo CD API server as a reverse-proxy to integrate external backend services.

## Files Summary
*   **ui-extensions.md**: Explains how to load and register custom JavaScript/React components to add tabs, sidebar items, status panels, and action menu buttons to the Argo CD UI.
*   **proxy-extensions.md**: Details the configuration and security protocols for the Beta proxy feature, which allows UI extensions to communicate with backend services through the Argo CD API server.

## Code Changes That Would Require Documentation Updates
*   **Extension Discovery Logic**: Any change to the file path (`/tmp/extensions`), filename regex (`^extension(.*)\.js$`), or the way the `argocd-server` scans for extension files.
*   **Global API Modifications**: Changes to the `extensionsAPI` global variable or its registration methods (`registerResourceExtension`, `registerSystemLevelExtension`, etc.).
*   **React Integration**: Updating the bundled version of React or changing the requirement for extensions to use `externals` in Webpack.
*   **UI Component Props**: Modifying the interfaces for `Application`, `State` (resource), or `ApplicationTree` passed to extension components.
*   **Feature Flag Management**: Renaming or removing the `server.enable.proxy.extension` flag in `argocd-cmd-params-cm`.
*   **Proxy Configuration Schema**: Altering the structure of `extension.config` in `argocd-cm`, including timeout settings, header management, or cluster-matching logic.
*   **Mandatory Headers**: Adding, renaming, or changing the format requirements for proxy headers such as `Argocd-Application-Name`, `Argocd-Project-Name`, or `Argocd-User-Id`.
*   **RBAC Policy Changes**: Modifications to how the `extensions` resource is handled within the Argo CD RBAC engine.
*   **UI Injection Points**: Adding new extension points in the UI (e.g., a new "Settings" extension point) or changing the layout of existing ones (Resource Tabs, Top Bar, Status Panel).

## Key Technical Concepts
*   **`extensionsAPI`**: The global JavaScript object used to register all types of UI extensions.
*   **Registration Methods**: `registerResourceExtension`, `registerSystemLevelExtension`, `registerStatusPanelExtension`, `registerTopBarActionMenuExt`.
*   **Flyout Widget**: A sliding panel component triggered by UI extensions for detailed views.
*   **Reverse-Proxy Extension**: The mechanism allowing the API server to forward authenticated requests to external backends.
*   **`extension.config`**: The ConfigMap key used to define backend service mappings, timeouts, and headers.
*   **Secret Interpolation**: The use of `$` prefix in header values to reference `argocd-secret` keys.
*   **Header-based Authorization**: The requirement for specific `Argocd-*` headers to validate user permissions before proxying.
*   **Cluster Matching**: Logic that routes proxy requests based on `Application.Spec.Destination` (name or server).

## Related Components
*   **argocd-server**: The primary backend component that serves the UI, discovers extension files, and acts as the proxy.
*   **Argo CD UI**: The React-based frontend that executes the extension JS and renders custom components.
*   **argocd-cm / argocd-cmd-params-cm**: Configuration Maps used to enable and configure the extension ecosystem.
*   **RBAC Subsystem**: Controls access to both the UI visibility and the proxy endpoints.
*   **Secret Manager**: Handles the retrieval of sensitive credentials used in proxy headers.