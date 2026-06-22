This comprehensive analysis of `operator-manual/ingress.md` provides a technical overview of how Argo CD is exposed to external traffic.

### 1. Primary Purpose
The file documents the configuration of external access to the Argo CD API server. It provides specific implementation patterns for exposing both the **Argo CD UI (HTTPS)** and the **Argo CD CLI (gRPC)** through various Kubernetes Ingress controllers and Cloud Load Balancers.

### 2. Key Topics Covered
*   **Protocol Multiplexing**: Handling the unique requirement of `argocd-server` which serves both gRPC and HTTPS/HTTP on the same ports (443/80).
*   **Ingress Controller Implementations**: Detailed guides for Ambassador, Contour, Nginx, Traefik, and Istio.
*   **Cloud-Specific Routing**: Configurations for AWS Application Load Balancers (ALB) and Google Kubernetes Engine (GKE) Ingress.
*   **Subpath Configuration**: How to host Argo CD on a non-root URL (e.g., `example.com/argo-cd`) using `--rootpath` and `--basehref`.
*   **TLS Strategies**: Documentation of both "SSL Passthrough" (terminating at the pod) and "SSL Termination" (terminating at the Ingress).
*   **Authentication Proxies**: Passing headers through multiple layers of reverse proxies.

### 3. Technical Keywords
*   **Core Components**: `argocd-server`, `argocd-cmd-params-cm` (ConfigMap).
*   **Configuration Flags**: `--insecure`, `--rootpath`, `--basehref`, `--grpc-web-root-path`, `--staticassets`.
*   **Protocols**: gRPC, HTTPS, HTTP/2 (h2c), WebSockets (streaming).
*   **Kubernetes Resources**: `Ingress`, `Service`, `ConfigMap`, `Secret` (TLS).
*   **Ingress-Specific Annotations**:
    *   *Nginx*: `ssl-passthrough`, `backend-protocol: "GRPC"`.
    *   *AWS ALB*: `backend-protocol-version: GRPC`, `conditions`.
    *   *Contour*: `upstream-protocol.h2c`.
    *   *GKE*: `cloud.google.com/neg`, `BackendConfig`, `FrontendConfig`.
*   **CRDs**: `Mapping` (Ambassador), `HTTPProxy` (Contour), `IngressRoute` (Traefik), `VirtualService`/`Gateway` (Istio).

### 4. Target Audience
*   **DevOps Engineers/SREs**: Responsible for the initial setup and exposure of Argo CD.
*   **Cluster Administrators**: Managing networking, DNS, and SSL/TLS certificates within Kubernetes.
*   **Security Engineers**: Configuring secure access, TLS termination policies, and header-based authentication.

### 5. Related Concepts
*   **SSO/Dex**: The documentation mentions configuring external callbacks for Dex (SSO) which is critical for enterprise identity management.
*   **Cert-Manager**: Frequently referenced for automating TLS certificate issuance (e.g., Let's Encrypt).
*   **High Availability**: Configuring Load Balancers is a prerequisite for a highly available Argo CD installation.
*   **gRPC-Web**: An alternative protocol mentioned as a way to simplify ingress requirements by avoiding pure gRPC.

---

### AI Update Trigger Analysis
This file should be updated if any of the following changes occur in the codebase:
1.  **Command-Line Arguments**: If `argocd-server` introduces new flags or deprecates current ones (like `--insecure`, `--rootpath`, or `--basehref`).
2.  **Service Port Changes**: If the default ports for the API server (80/443/8080) are modified.
3.  **Protocol Shifts**: If the API server switches from pure gRPC to only gRPC-Web, or introduces a new communication protocol.
4.  **Health Check Paths**: If the internal health check endpoints (`/healthz` or `/grpc.health.v1.Health/Check`) are renamed.
5.  **Dependency Versions**: If the supported versions of Ingress controllers (e.g., Traefik v3.0) undergo breaking changes in their CRD schemas or annotation requirements.
6.  **ConfigMap Schema**: Changes to the keys supported within `argocd-cmd-params-cm`.