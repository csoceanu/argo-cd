# DEVELOPER-GUIDE/ARCHITECTURE Documentation Index

## Overview
This documentation area provides a high-level blueprint of Argo CD's architectural design and security implementation. It outlines the modular component structure, logical layering of services, and the internal mechanics of the API server's authentication and authorization workflows.

## Files Summary
*   **components.md**: Describes the component-based architecture of Argo CD, detailing the four logical layers (UI, Application, Core, Infra) and the specific responsibilities of key services like the API Server, Controllers, and Repo Server.
*   **authz-authn.md**: Provides a deep dive into the API server's internal request handling, explaining how it uses multiplexing to handle gRPC and HTTP, and how authentication (AuthN) and authorization (AuthZ) are enforced via interceptors and Casbin.

## Code Changes That Would Require Documentation Updates
*   **Introduction of New Components**: Adding a new microservice or deployable unit to the Argo CD ecosystem.
*   **Controller Modifications**: Changes to the core reconciliation logic or the introduction of new Custom Resource Definitions (CRDs) beyond Application and ApplicationSet.
*   **API Protocol Changes**: Modifying how the API server handles requests, such as switching from gRPC/REST to a different protocol or changing the multiplexing logic (e.g., replacing `cmux`).
*   **Authentication Flow Updates**: Changes to how tokens are validated, how the Session Manager interacts with external providers (like Dex), or updates to the login/OIDC handshake.
*   **Authorization Engine Changes**: Modifications to the RBAC implementation, updating the Casbin library version/configuration, or changing how permissions are defined in Projects.
*   **Dependency Shifts**: Swapping out core infrastructure dependencies, such as replacing Redis for caching or Dex for authentication.
*   **API Gateway Logic**: Changes to `grpc-gateway` configurations or the way gRPC services are exposed as REST endpoints.
*   **Layering Violations/Adjustments**: Any change that alters the top-down dependency rule (e.g., making a Core component depend on an Application layer component).

## Key Technical Concepts
*   **Component Modularity**: Separation of concerns into independent deployable units.
*   **Logical Layering**: UI, Application, Core, and Infra tiers.
*   **Reconciliation Loop**: The process by which controllers sync Git state to Kubernetes state.
*   **Connection Multiplexing (Cmux)**: Routing HTTP/1.1 and gRPC (HTTP/2) traffic on a single port (8080).
*   **gRPC-gateway**: Translation layer between RESTful HTTP and internal gRPC services.
*   **AuthN Interceptors**: Middleware-like functions that trigger authentication for every gRPC request.
*   **RBAC Enforcement**: The use of Casbin and internal functions to validate user permissions against Project or global rules.
*   **Session Management**: The lifecycle and verification of authentication tokens.
*   **Desired vs. Live State**: The core GitOps concept managed by the Repo Server and Application Controller.

## Related Components
*   **API Server**: The primary entry point for Web UI and CLI traffic.
*   **Application Controller**: The engine for resource reconciliation.
*   **ApplicationSet Controller**: The generator for multiple Application resources.
*   **Repo Server**: The service responsible for manifest generation from Git/Helm/OCI.
*   **Redis**: The shared cache for Git and Kube API data.
*   **Dex**: The default OIDC authentication provider.
*   **Argo CD CLI / Webapp**: The primary consumers of the architectural layers described.
*   **Kubernetes API**: The target for state reconciliation.
*   **Casbin**: The library powering the RBAC engine.