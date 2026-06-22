# DEVELOPER-GUIDE/EXTENSIONS Documentation Index

## Overview
This documentation area provides instructions for extending Argo CD's functionality through custom UI elements and backend proxy services. It details how developers can inject React-based components into specific areas of the web interface and how to configure the Argo CD API server as a secure reverse-proxy to integrate third-party backend data.

## Files Summary
*   **developer-guide/extensions/ui-extensions.md**: Describes how to build and register UI extensions for resource tabs, system-level pages, status panels, and top bar action menus using the `extensionsAPI`.
*   **developer-guide/extensions/proxy-extensions.md**: Details the configuration and usage of the backend proxy feature, including feature flags, `argocd-cm` settings, header propagation, and security/RBAC enforcement.

## Code Changes That Would Require Documentation Updates
*   **UI Registration API**: Any changes to the `extensionsAPI` global variable or the signature of registration methods (e.g., `registerResourceExtension`, `registerStatusPanelExtension`).
*   **Extension Loading Logic**: Modifying the directory path (`/tmp/extensions`), the file naming convention (`^extension(.*)\.js$`), or the way the Argo CD server serves these static assets.
*   **Component Property Interfaces**: Changes to the data structures passed to React components, such as `application`, `resource`, or `tree` models.
*   **New UI Extension Points**: Adding new areas in the Argo CD interface where extensions can be rendered (e.g., new menus, footer, or login page extensions).
*   **Proxy Configuration Schema**: Updates to the `extension.config` YAML structure in `argocd-cm`, such as adding new backend settings or changing timeout defaults.
*   **Proxy Feature Flags**: Changes to the command-line parameters or ConfigMap keys used to enable proxying (`server.enable.proxy.extension`).
*   **Header Propagation**: Modifying the set of mandatory or optional headers sent by the API server to backend services (e.g., `Argocd-User-Id`, `Argocd-Project-Name`).
*   **RBAC Policy Changes**: Alterations to how the `extensions` resource is handled within the Argo CD RBAC engine for authorization.
*   **React Integration**: Changing the way Argo CD handles external dependencies for extensions, such as moving away from the `React` global variable.

## Key Technical Concepts
*   **extensionsAPI**: The global JavaScript object used by extensions to register themselves with the Argo CD UI.
*   **Resource Tab Extension**: Custom tabs added to the Kubernetes resource sliding panel.
*   **System Level Extension**: Custom sidebar items that render full-page components.
*   **Flyout Widget**: A sliding panel triggered from an extension component via the `openFlyout` function.
*   **Top Bar Action Menu**: The primary action area (containing Details, Sync, etc.) where custom buttons can be injected.
*   **Proxy Backend**: Configuration for the API server to forward requests to external URLs based on cluster context.
*   **Header Injection**: The process of adding authentication context (User ID, Groups, Application name) to proxied requests.
*   **ConfigMap `argocd-cm`**: The primary configuration source for defining proxy extension backends.
*   **ConfigMap `argocd-cmd-params-cm`**: Used to enable the beta proxy extension feature flag.

## Related Components
*   **Argo CD API Server (argocd-server)**: Responsible for serving extension JS files and acting as the reverse-proxy.
*   **Argo CD Web UI**: The React-based frontend that loads and renders the UI extensions.
*   **Argo CD RBAC System**: Controls access permissions for both the UI elements and the proxy endpoints.
*   **Kubernetes API**: The source of resource data (`Application`, `ApplicationTree`) passed to extension components.
*   **Backend Services**: Third-party services that provide additional data/functionality via the proxy extension mechanism.