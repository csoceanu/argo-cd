# DEVELOPER-GUIDE/EXTENSIONS Documentation Index

## Overview
This documentation area provides instructions for extending the Argo CD platform through both frontend UI components and backend proxy services. It outlines how developers can inject custom tabs, sidebar items, and action buttons into the web interface, and how to securely proxy requests through the Argo CD API server to external backend services.

## Files Summary
*   **ui-extensions.md**: Describes the JavaScript-based extensibility framework for the Argo CD UI, detailing the `extensionsAPI` for registering resource tabs, system-level sidebar items, status panel widgets, and top-bar action buttons.
*   **proxy-extensions.md**: Details the backend reverse-proxy feature that allows UI extensions to communicate with external services, covering configuration in `argocd-cm`, security/RBAC enforcement, and header management.

## Code Changes That Would Require Documentation Updates
Documentation updates are necessary if any of the following areas of the Argo CD codebase are modified:

### UI & Frontend Framework
*   **Extensions Loading Logic**: Changes to the directory (`/tmp/extensions`), file naming regex (`^extension(.*)\.js$`), or the method by which the UI server fetches and injects extension scripts.
*   **Global Variable Exposure**: Renaming or restructuring the `extensionsAPI` or `React` global variables provided to extensions.
*   **API Interface Changes**: Adding, removing, or changing arguments for registration methods:
    *   `registerResourceExtension`
    *   `registerSystemLevelExtension`
    *   `registerStatusPanelExtension`
    *   `registerTopBarActionMenuExt`
*   **Data Model Updates**: Modifications to the TypeScript interfaces for `Application`, `ApplicationTree`, or Kubernetes resource states (referenced in `models.ts`) that are passed as props to extensions.
*   **UI Component Lifecycle**: Changes to how "Flyout" widgets are triggered (`openFlyout`) or rendered.

### Backend & Proxy Server
*   **Feature Flag Management**: Changes to the toggle key `server.enable.proxy.extension` in `argocd-cmd-params-cm`.
*   **Configuration Schema**: Alterations to the `extension.config` structure within `argocd-cm` (e.g., changing field names like `connectionTimeout`, `keepAlive`, or `maxIdleConnections`).
*   **Routing Logic**: Modifications to how the API server handles the `<argocd-host>/extensions/` base path or how it performs cluster-based routing for multi-backend services.
*   **Security & RBAC**: 
    *   Changes to the "Extensions" RBAC resource definition.
    *   Modifications to header sanitization (removing `Cookie` or `Authorization` before proxying).
*   **Request/Response Headers**:
    *   Changes to mandatory incoming headers: `Argocd-Application-Name`, `Argocd-Project-Name`.
    *   Changes to outgoing headers sent to backends: `Argocd-Target-Cluster-Name`, `Argocd-User-Id`, `Argocd-User-Groups`, etc.
*   **Secret Resolution**: Changes to how the proxy resolves `$secret.key` references in the configuration.

## Key Technical Concepts
*   **extensionsAPI**: The global JS object used to interface with the Argo CD UI.
*   **Resource Tab Extension**: Custom UI components added to the sliding panel of specific K8s resource types.
*   **System Level Extension**: Global sidebar entries that link to a full-page custom component.
*   **Status Panel Extension**: Widgets added to the top sync-status bar of an application.
*   **Top Bar Action Menu**: Custom buttons added next to "Sync", "Refresh", and "Details" buttons.
*   **Flyout Widget**: A sliding panel UI pattern used for complex extension interactions.
*   **extension.config**: The YAML configuration block in `argocd-cm` defining proxy routes.
*   **Header Injection**: The process of adding Argo CD user/context metadata (like `Argocd-Username`) to outgoing proxy requests.
*   **Cluster-Aware Routing**: Forwarding proxy requests to specific backend URLs based on the target cluster of the Argo CD Application.

## Related Components
*   **Argo CD UI (argocd-ui)**: The React-based frontend that executes the extension JS.
*   **Argo CD Server (argocd-server)**: The backend API service responsible for serving extension files and handling proxy requests.
*   **ConfigMaps**: `argocd-cm` and `argocd-cmd-params-cm` for feature enablement and proxy definitions.
*   **RBAC Subsystem**: Validates if the user has `invoke` permissions on the `extensions` resource.
*   **Kubernetes API**: Extensions often interact with resource data (`ApplicationTree`) derived from the K8s API.