This documentation provides a technical guide for implementing and using **Plugin Generators** within Argo CD ApplicationSets. The Plugin Generator allows users to extend Argo CD's capabilities by calling external services (via HTTP/RPC) to generate parameters for Application creation, bypassing the limitations of built-in generators.

### 1. Primary Purpose
The file documents how to build, configure, and deploy custom generators for Argo CD ApplicationSets. It specifically addresses how to use external logic (written in any programming language) to fetch data from non-Git sources (like CI systems, custom APIs, or databases) and feed that data into the ApplicationSet controller.

### 2. Key Topics Covered
*   **Plugin Architecture**: Explains that plugins are simple HTTP services responding to RPC-style requests (`/api/v1/getparams.execute`).
*   **Configuration**: How to define the plugin in an `ApplicationSet` spec using `configMapRef` and `input.parameters`.
*   **Security & Authentication**: Mechanisms for securing the communication between Argo CD and the plugin using Bearer tokens stored in Kubernetes Secrets.
*   **Polling Behavior**: Usage of `requeueAfterSeconds` to control how often the controller polls the plugin for changes.
*   **Advanced Composition**: Using the Plugin Generator as a child of `Matrix` or `Merge` generators to create complex workflows (e.g., combining a Pull Request generator with a custom Image Digest lookup plugin).
*   **Implementation Example**: A reference Python implementation of the required HTTP server logic.

### 3. Technical Keywords
*   **CRD/Schema**: `ApplicationSet`, `generators`, `plugin`, `configMapRef`, `input.parameters`, `requeueAfterSeconds`.
*   **API/Endpoints**: `/api/v1/getparams.execute`.
*   **Templating**: `goTemplate: true`, `{{ .generator.input.parameters }}`, `{{ .values }}`.
*   **Configuration Keys**: `baseUrl`, `token`, `requestTimeout`.
*   **Environment/Secrets**: `argocd-secret`, `app.kubernetes.io/part-of: argocd`, `Authorization: Bearer`.
*   **Outputs**: `output.parameters`.

### 4. Target Audience
*   **Platform Engineers**: Who need to integrate Argo CD with internal developer portals or bespoke metadata services.
*   **DevOps Engineers**: Who need to automate application deployments based on external triggers (like specific CI build outputs) not natively supported by Argo CD.
*   **Plugin Developers**: Software engineers writing the actual HTTP services that interface with Argo CD.

### 5. Related Concepts
*   **Argo CD ApplicationSets**: The parent controller that manages multiple Applications.
*   **Matrix & Merge Generators**: Built-in generators that can wrap the Plugin generator to iterate over multiple data sets.
*   **GitOps**: The document emphasizes that while plugins fetch external data, the goal is to remain complementary to GitOps principles.
*   **Sidecar vs. Standalone Deployment**: Deployment strategies for the plugin service within Kubernetes.

---

### AI Update Trigger Analysis
An AI system should flag this file for updates if any of the following code changes occur in the Argo CD repository:

1.  **CRD Changes**: Any modification to the `ApplicationSet` CustomResourceDefinition, specifically under the `generators.plugin` or `generators.matrix/merge` fields.
2.  **Controller Logic**: Changes in the ApplicationSet controller related to how it handles HTTP requests, timeout logic, or the specific RPC path (`/api/v1/getparams.execute`).
3.  **Secret Handling**: Changes to the way Argo CD parses secrets for plugins (e.g., changing the required labels or the `$<secret>:<key>` syntax).
4.  **Templating Engine**: Updates to the Go Templating integration within ApplicationSets that might change how plugin output is accessed (e.g., changes to `goTemplateOptions`).
5.  **Reserved Keys**: If the controller adds or changes "reserved keys" (currently `generator.input.parameters` and `values`) that the plugin is not allowed to overwrite.