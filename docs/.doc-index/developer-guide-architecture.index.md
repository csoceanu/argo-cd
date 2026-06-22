# DEVELOPER-GUIDE/ARCHITECTURE Documentation Index

## Overview
This documentation area provides a high-level blueprint of the Argo CD system design, detailing its modular component-based architecture and its internal security enforcement mechanisms. It defines the responsibilities of various microservices and explains the logical flow of requests through the API server for authentication and authorization.

## Files Summary
- **developer-guide/architecture/components.md**: Outlines the modular structure of Argo CD, categorizing services into four logical layers (UI, Application, Core, and Infra) and defining the specific roles of components like the API Server, Controllers, and Repo Server.
- **developer-guide/architecture/authz-authn.md**: Provides a deep dive into the API server's internal request processing, explaining how connection multiplexing (cmux), gRPC-gateway, and Casbin-based RBAC are used to secure the system.

## Code Changes That Would Require Documentation Updates
- **Introduction of New Services**: Adding a new deployable unit or Kubernetes controller (e.g., a notification service or a new resource generator).
- **Changes to Inter-Component Communication**: Modifying how components interact (e.g., changing the protocol between the API Server and Repo Server).
- **Authentication Provider Updates**: Changes to how Argo CD integrates with Dex or other OIDC providers, or modifications to the internal `Session Manager` logic.
- **Authorization Logic Refactors**: Replacing the RBAC engine (Casbin) or changing the enforcement points within the `Service Method` or `HTTP Handler`.
- **Infrastructure Dependency Shifts**: Replacing Redis with a different caching mechanism or introducing new mandatory infrastructure tools.
- **API Server Entry-point Modifications**: Changes to the `cmux` logic, port configurations (default 8080), or the way gRPC and HTTP/1.x traffic is multiplexed.
- **Logical Layering Adjustments**: Reorganizing the hierarchy of dependencies or moving functionality between the "Core" and "Application" layers.
- **CLI/UI Communication Changes**: Altering how the CLI or Webapp interacts with the API server, such as bypassing the gRPC-gateway or changing REST API definitions.

## Key Technical Concepts
- **Core Components**: API Server, Application Controller, ApplicationSet Controller, Repo Server, Redis, Dex.
- **Architectural Layers**: UI Layer, Application Layer, Core Layer (GitOps functionality), Infra Layer.
- **Request Multiplexing**: Cmux, HTTP Mux, gRPC Server, Connection Multiplexer.
- **API Translation**: gRPC-gateway, REST API, Protocol Buffers (implied), HTTP/2 vs HTTP/1.x.
- **Security Enforcement**: AuthN (Authentication), AuthZ (Authorization), RBAC (Role-Based Access Control), Casbin, Session Management, gRPC Interceptors, HTTP Middleware.
- **State Management**: Git/Helm/OCI repositories, Kube API, Reconciliation Loop, Project CRDs.

## Related Components
- **Argo CD API Server** (argocd-server)
- **Argo CD Application Controller** (argocd-application-controller)
- **Argo CD ApplicationSet Controller** (argocd-applicationset-controller)
- **Argo CD Repo Server** (argocd-repo-server)
- **Argo CD CLI**
- **Argo CD Web UI** (Webapp)
- **Redis**
- **Dex** (Identity Service)
- **Kubernetes API Server**