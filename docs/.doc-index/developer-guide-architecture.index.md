# DEVELOPER-GUIDE/ARCHITECTURE Documentation Index

## Overview
This documentation provides a high-level blueprint of the Argo CD system design, focusing on its modular component architecture and the internal mechanics of its security framework. It explains how different deployable units interact across logical layers and details the specific path a request takes through the API server for authentication and authorization.

## Files Summary
* **developer-guide/architecture/components.md**: Outlines the component-based architecture of Argo CD, detailing the responsibilities of controllers, servers, and external dependencies, as well as their hierarchical relationships.
* **developer-guide/architecture/authz-authn.md**: Explains the internal implementation of the API server's authentication and authorization layers, including connection multiplexing, gRPC interceptors, and RBAC enforcement via Casbin.

## Code Changes That Would Require Documentation Updates
* **Architectural Refactoring**: Adding, removing, or renaming core components (e.g., introducing a new controller or splitting the Repo Server).
* **Dependency Changes**: Changing the version or role of foundational infrastructure like Redis, Dex, or the Kubernetes API.
* **API Protocol Shifts**: Modifying how the API server handles requests, such as replacing `cmux` or `grpc-gateway`, or adding support for new protocols beyond gRPC and HTTP.
* **AuthN/AuthZ Logic**: Updates to the `Session Manager`, changes to the authentication interceptor logic, or modifying the `AuthN Provider` interface.
* **RBAC Engine Updates**: Replacing `Casbin` with a different library or changing how RBAC rules are evaluated within the `Project` CRD or the API server.
* **Component Layering**: Altering the dependency flow between the UI, Application, Core, and Infra layers (e.g., making a Core component depend on a UI component).
* **New CRD Integration**: Introducing new Custom Resource Definitions that require specialized controllers or reconciliation logic.

## Key Technical Concepts
* **Modular Architecture**: Separation of concerns into independent deployable units.
* **Logical Layers**: The hierarchy of UI, Application, Core, and Infra.
* **Reconciliation Loop**: The process used by the Application and ApplicationSet controllers to sync state.
* **Repo Server**: The component responsible for manifest generation from Git/Helm/OCI.
* **Connection Multiplexing (cmux)**: Handling gRPC and HTTP requests on a single port (8080).
* **gRPC-gateway**: The translation layer that exposes gRPC services as a REST API.
* **AuthN (Authentication)**: Token verification and session management logic.
* **AuthZ (Authorization)**: Permission validation using RBAC and Casbin.
* **gRPC Interceptor**: Middleware used to automatically trigger AuthN logic for every gRPC request.
* **Casbin**: The library used for policy-based access control.
* **Project CRD**: Kubernetes resource used to define fine-grained RBAC rules.

## Related Components
* **API Server**: The central gateway for UI and CLI interactions.
* **Application Controller**: Manages the lifecycle and sync of Application resources.
* **ApplicationSet Controller**: Automates the generation of multiple Applications.
* **Repo Server**: Interfaces with Git, Helm, and OCI repositories.
* **Redis**: Provides the caching layer for Kube API and Git requests.
* **Dex**: Handles OIDC-based external authentication.
* **Argo CD CLI & Webapp**: The primary clients for the API server.
* **Casbin & gRPC-gateway**: Key libraries facilitating security and API access.