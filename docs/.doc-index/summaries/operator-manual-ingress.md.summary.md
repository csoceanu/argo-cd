This documentation provides a technical guide for exposing the Argo CD API server to external traffic using various Kubernetes Ingress controllers and cloud-provider load balancers.

### 1. Primary Purpose
The file documents the network configuration required to make the Argo CD API server accessible for both the **Web UI (HTTPS)** and the **CLI (gRPC)**. It provides specific implementation patterns for various Ingress controllers, handling the complexity of multiplexing or separating gRPC and standard HTTP traffic.

### 2. Key Topics Covered
*   **Protocol Handling**: Managing the dual nature of Argo CD’s API server (gRPC for CLI and HTTP for UI).
*   **Ingress Controller Implementations**:
    *   **Ambassador**: Using Mappings for host and path-based routing.
    *   **Contour**: Using `HTTPProxy` and handling private/public ingress paths.
    *   **NGINX**: Strategies for SSL-Passthrough vs. SSL Termination at the controller.
    *   **Traefik (v3.0)**: Utilizing `IngressRoute` for same-port TCP/HTTP termination.
    *   **Istio**: Configuration via Gateways and VirtualServices, including subpath overrides.
*   **Cloud Provider Integrations**:
    *   **AWS ALB**: Configuring target groups for gRPC and using specific health check paths (`/grpc.health.v1.Health/Check`).
    *   **GKE Ingress**: Leveraging Google Cloud-specific objects like `BackendConfig`, `FrontendConfig`, and Network Endpoint Groups (NEGs).
*   **Path Customization**: Detailed instructions for running Argo CD under a subpath (e.g., `/argo-cd`) using `--rootpath` and `--basehref`.
*   **Authentication**: Handling multiple layers of reverse proxies using custom CLI headers.

### 3. Technical Keywords
*   **Argo CD Server Flags**: `--insecure`, `--rootpath`, `--basehref`, `--grpc-web-root-path`, `--staticassets`.
*   **Configuration Objects**: `argocd-cmd-params-cm` (ConfigMap), `argocd-server` (Deployment/Service).
*   **Kubernetes Resources**: `Ingress`, `Service`, `Secret`, `ConfigMap`.
*   **Controller-Specific CRDs**: `Mapping` (Ambassador), `HTTPProxy` (Contour), `IngressRoute` (Traefik), `VirtualService`/`Gateway` (Istio), `BackendConfig`/`FrontendConfig` (GKE).
*   **Annotations**: 
    *   `nginx.ingress.kubernetes.io/ssl-passthrough`
    *   `alb.ingress.kubernetes.io/backend-protocol-version: GRPC`
    *   `projectcontour.io/upstream-protocol.h2c`
*   **Protocols**: gRPC, HTTPS, h2c (HTTP/2 without TLS).

### 4. Target Audience
*   **Kubernetes Operators/SREs**: Responsible for deploying and securing Argo CD infrastructure.
*   **DevOps Engineers**: Setting up CI/CD pipelines that interact with the Argo CD CLI from outside the cluster.
*   **Network Engineers**: Configuring load balancers, DNS, and TLS termination for internal tools.

### 5. Related Concepts
*   **TLS Termination**: Deciding whether to terminate SSL at the Ingress, Load Balancer, or the Argo CD Pod.
*   **gRPC Multiplexing**: The technical challenge of routing gRPC and HTTP traffic which often share the same port (443).
*   **Identity Providers/SSO**: Mentions of Dex and SSO callbacks (`/api/dex/callback`) which are critical for external access.
*   **Health Checking**: Protocol-specific health checks (standard HTTP vs. gRPC health probes).

---

### AI Update Trigger Guide
This file should be updated if any of the following changes occur in the codebase:
*   **CLI Changes**: If new flags are added to `argocd login` related to connectivity or headers.
*   **Server Logic**: If the API server changes how it handles gRPC vs. HTTP (e.g., moving them to separate ports or changing the internal health check path).
*   **Configuration Schema**: If the `argocd-cmd-params-cm` ConfigMap introduces new keys for server pathing or security.
*   **Dependency Updates**: If supported versions of Ingress controllers (like Traefik or NGINX) introduce breaking changes in their CRDs or annotation syntax.
*   **Default Behavior**: If the default TLS behavior of `argocd-server` changes (e.g., switching from secure by default to a different protocol).