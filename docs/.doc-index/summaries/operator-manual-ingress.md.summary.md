This documentation provides a technical guide for configuring external access to an Argo CD instance using various Kubernetes Ingress controllers and cloud-provider load balancers. It specifically addresses the complexity of Argo CD’s dual-protocol nature (gRPC for CLI and HTTPS for the UI).

### 1. Primary Purpose
The file documents the configuration required to expose the **Argo CD API Server** to external networks. It provides specific recipes for routing both web traffic (UI) and gRPC traffic (CLI) through different ingress providers while handling TLS termination and path-based routing.

### 2. Key Topics Covered
*   **Protocol Handling**: Managing the coexistence of gRPC and HTTP/HTTPS on the same or different ports.
*   **TLS Strategies**: Configuring SSL Passthrough (terminating at the pod) versus SSL Termination (terminating at the Ingress/Load Balancer).
*   **Controller-Specific Configurations**: Tailored guides for Ambassador, Contour, Nginx, Traefik, Istio, AWS ALB, and GKE Ingress.
*   **Subpath/Rootpath Configuration**: Running Argo CD under a non-root URL (e.g., `example.com/argo-cd`).
*   **Authentication Multi-layering**: Using CLI headers to bypass external authentication proxies.

### 3. Technical Keywords
*   **Core Flags**: `--insecure`, `--rootpath`, `--basehref`, `--grpc-web-root-path`.
*   **ConfigMaps**: `argocd-cmd-params-cm`.
*   **Service/Ingress Annotations**: 
    *   Nginx: `ssl-passthrough`, `backend-protocol`.
    *   AWS ALB: `alb.ingress.kubernetes.io/backend-protocol-version: GRPC`.
    *   GKE: `cloud.google.com/neg`, `networking.gke.io/v1beta1.FrontendConfig`.
*   **Custom Resource Definitions (CRDs)**: `Mapping` (Ambassador), `HTTPProxy` (Contour), `IngressRoute` (Traefik), `VirtualService`/`Gateway` (Istio).
*   **Protocols/Headers**: `gRPC`, `h2c` (HTTP/2 cleartext), `application/grpc` Content-Type.

### 4. Target Audience
*   **Kubernetes Cluster Administrators**: Responsible for networking and security.
*   **DevOps Engineers**: Deploying and maintaining Argo CD infrastructure.
*   **SREs**: Troubleshooting connectivity issues between the Argo CD CLI/UI and the server.

### 5. Related Concepts
*   **TLS Termination & Certificates**: Integration with `cert-manager` and Let's Encrypt.
*   **Load Balancing (L4 vs L7)**: Differences between Network Load Balancers and Application Load Balancers.
*   **SSO/OIDC**: Configuration for external callback URLs (specifically discussed in the Contour section).
*   **Reverse Proxies**: General behavior of Nginx and Envoy-based controllers.

---

### Triggering Updates: When to Edit This File
An AI system should suggest updates to this file if code changes occur in the following areas:
1.  **Server Startup Arguments**: If new flags are added to `argocd-server` that affect networking, ports, or protocol handling (similar to the existing `--insecure` or `--rootpath`).
2.  **ConfigMap Schema**: If the `argocd-cmd-params-cm` entry keys for `server.*` properties are modified or expanded.
3.  **API Protocol Changes**: If the API server introduces new protocols or changes the existing gRPC/HTTP implementation (e.g., moving away from port 8080/443).
4.  **Health Check Endpoints**: If the internal health check paths (currently `/healthz` or `/grpc.health.v1.Health/Check`) are renamed or changed.
5.  **CLI Authentication**: If the `argocd login` command changes how it handles headers or root paths.
6.  **Dependency Versions**: If supported ingress providers (like Traefik or Ambassador) release major versions requiring different CRD versions or annotation syntax.