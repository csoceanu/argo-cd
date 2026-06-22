# DEVELOPER-GUIDE/ARCHITECTURE Documentation Index

## Overview
This documentation area provides a high-level blueprint of Argo CD’s system design, emphasizing its modular, component-based architecture and its security implementation. It defines the structural boundaries of the system, the interaction between logical layers (UI, Application, Core, Infra), and the specific request-handling pipeline for authentication and authorization within the API server.

## Files Summary
*   **components.md**: Outlines the modular design of Argo CD, defining the responsibilities of core deployable units (API Server, Controllers, Repo Server, etc.) and their hierarchical dependencies.
*   **authz-authn.md**: Provides a deep dive into the API Server's internal logic, detailing how connection multiplexing (cmux), gRPC/REST translation, and security enforcement (AuthN/AuthZ via Casbin) are layered.

## Code Changes That Would Require Documentation Updates
*   **Architectural Shifts**: Adding a new microservice/deployable component or merging existing ones (e.g., merging ApplicationSet Controller into Application Controller).
*   **Dependency Graph Changes**: Introducing a new dependency between existing components (e.g., if the Repo Server suddenly required a direct connection to the Application Controller).
*   **Communication Protocol Updates**: Changing the interface between components, such as moving from gRPC to a different RPC framework or changing the port-sharing logic in `cmux`.
*   **Authentication Logic Refactoring**: Modifying the `Session Manager` logic, changing how OIDC/Dex integrations are handled, or updating the gRPC interceptors used for AuthN.
*   **Authorization/RBAC Framework Updates**: Changing the underlying RBAC engine (Casbin), modifying how RBAC rules are parsed from Project CRDs, or changing the enforcement points in `Service Methods`.
*   **API Gateway Modifications**: Changes to `grpc-gateway` configurations, REST API pathing, or how HTTP middleware is applied to non-gRPC endpoints.
*   **Infrastructure Requirements**: Changing the role of Redis (e.g., moving from cache to primary persistence) or adding support for new repository types beyond Git/Helm/OCI.

## Key Technical Concepts
*   **Component Layers**: UI Layer, Application Layer, Core Layer (GitOps functionality), Infrastructure Layer.
*   **Core Controllers**: Application Controller (Sync/Reconciliation), ApplicationSet Controller.
*   **Service Components**: Repo Server (Manifest generation), API Server, Redis (Caching), Dex (OIDC Provider).
*   **Request Handling**: Connection Multiplexing (cmux), gRPC-gateway (Protobuf to REST translation), HTTP Multiplexer.
*   **Security Enforcement**: AuthN (Authentication), AuthZ (Authorization), RBAC, Casbin (Policy Engine), gRPC Interceptors, Session Management.
*   **Kubernetes Integration**: Custom Resource Definitions (CRDs), Reconciliation Loops, Kube API interaction.

## Related Components
*   **Argo CD API Server**: The central hub for UI/CLI interactions and security enforcement.
*   **Argo CD Repository Server**: Responsible for cloning repositories and rendering manifests.
*   **Argo CD Application Controller**: The primary engine for state reconciliation.
*   **Argo CD CLI & Webapp**: The primary clients consuming the documented architectural interfaces.
*   **External Providers**: Redis (Cache), Dex (Auth), Git/Helm/OCI (Source Control).