This analysis provides a comprehensive overview of the `operator-manual/ingress.md` documentation for Argo CD.

### 1. Primary Purpose
The file serves as a technical guide for exposing the Argo CD API server (which handles both UI traffic via HTTP/HTTPS and CLI traffic via gRPC) to external networks. It provides specific configuration patterns for various Kubernetes Ingress controllers and cloud-provider load balancers to ensure both protocols function correctly.

### 2. Key Topics Covered
*   **Dual-Protocol Requirements**: Explaining how Argo CD uses port 443 for both gRPC and HTTPS.
*   **Ingress Controller Implementations**: Detailed guides and YAML manifests for:
    *   **Ambassador**: Host-based and path-based routing.
    *   **Contour**: Handling multiple Ingress objects for different protocols and private/public access splits.
    *   **NGINX**: Comparing SSL-Passthrough (easiest for gRPC) vs. SSL Termination (requiring multiple hostnames).
    *   **Traefik (v3.0)**: Utilizing `IngressRoute` CRDs to terminate both protocols on one port.
    *   **Istio**: Configuring Gateways and VirtualServices for subpath routing.
*   **Cloud Load Balancers**:
    *   **AWS ALB**: Setting up specific target groups for gRPC/HTTP2 and health check configurations.
    *   **Google Cloud LB (GKE)**: Integrating with BackendConfigs, FrontendConfigs, and NEGs.
*   **Subpath Configuration**: Instructions for running Argo CD under a non-root URL (e.g., `/argo-cd`) using `--rootpath` and `--basehref`.
*   **Proxy Authentication**: Using the CLI to pass custom headers through authenticating reverse proxies.

### 3. Technical Keywords
*   **APIs/Services**: `argocd-server`, `argocd-cmd-params-cm` (ConfigMap).
*   **CLI Flags**: `--insecure`, `--rootpath`, `--basehref`, `--grpc-web-root-path`, `--header`.
*   **Protocols**: gRPC, HTTPS, h2c (HTTP/2 without TLS), GRPC-Web.
*   **Annotations**:
    *   `nginx.ingress.kubernetes.io/ssl-passthrough`
    *   `alb.ingress.kubernetes.io/backend-protocol-version: GRPC`
    *   `projectcontour.io/upstream-protocol.h2c`
    *   `cloud.google.com/neg`
*   **Kubernetes Resources**: `Ingress`, `Service`, `HTTPProxy` (Contour), `IngressRoute` (Traefik), `BackendConfig` (GCP).

### 4. Target Audience
*   **Platform Engineers/DevOps**: Responsible for the initial setup and networking of Argo CD.
*   **Kubernetes Administrators**: Managing ingress controllers, SSL certificates, and cluster security.
*   **Cloud Architects**: Designing external access via AWS or Google Cloud infrastructure.

### 5. Related Concepts
*   **TLS/SSL Management**: Frequently references `cert-manager` and Let's Encrypt for certificate automation.
*   **Authentication (Dex/SSO)**: Mentioned in the context of external callback URLs for Single Sign-On.
*   **Network Security**: Concepts of private vs. public subnets and internal/external load balancing.

---

### Update Triggers: When to modify this file
An AI system should flag this file for updates if code changes occur in the following areas:

1.  **API Server Flags**: If the `argocd-server` command-line arguments related to networking (like `--insecure`, `--rootpath`, or protocol handling) are added, renamed, or deprecated.
2.  **Service Definitions**: If the default ports (80/443) or names for the `argocd-server` service are modified in the core manifests.
3.  **gRPC Implementation**: If the way Argo CD handles gRPC changes (e.g., moving to a different port or requiring specific HTTP/2 settings).
4.  **CRD Versioning**: If supported third-party tools (Traefik, Contour, Ambassador) release major versions that change their CRD schemas (e.g., moving from `v1alpha1` to `v1`).
5.  **New Cloud Integrations**: If official support or recommended patterns for other cloud providers (like Azure Application Gateway) are added to the codebase.
6.  **CLI Updates**: If new authentication or connection flags are added to the `argocd login` command.