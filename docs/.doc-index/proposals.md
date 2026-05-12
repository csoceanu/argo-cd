# Index: proposals/

## What this folder documents
The proposals folder contains formal design documents (KEP-style) for new Argo CD features and significant architectural changes. Each proposal describes motivation, goals, non-goals, and the detailed design. Proposals may be in various states (draft, implemented, or archived). This folder is mainly relevant to core maintainers and contributors evaluating or implementing new features.

## Key topics
- Proposal template (`001-proposal-template.md`) — standard front-matter (title, authors, sponsors, reviewers, approvers, creation-date) and section structure
- UI Extensions (`002-ui-extensions.md`) — adding custom React tabs/panels to the Argo CD web UI
- Applications outside argocd namespace (`003-applications-outside-argocd-namespace.md`) — app-in-any-namespace feature
- Scalability benchmarking (`004-scalability-benchmarking.md`) — methodology for measuring Argo CD at scale
- ApplicationSet Progressive Rollout Strategy (`2022-07-13-appset-progressive-rollout-strategy.md`) — declarative ordered rollout of ApplicationSet-managed apps
- Application name identifier (`application-name-identifier.md`) — globally unique application identifiers
- ApplicationSet Plugin Generator (`applicationset-plugin-generator.md`) — custom generator plugins for ApplicationSet
- Argo CD CLI plugin system (`argocd-cli-pluin.md`) — extensible CLI via plugins
- Backend support for ApplicationSet (`backend-support-appset.md`) — server-side support improvements
- Config Management Plugin v2 (`config-management-plugin-v2.md`) — sidecar CMP model replacing configmap-based plugins (cdk8s, Tanka, jkcfg, QBEC, Dhall, Pulumi)
- Decouple application sync using impersonation (`decouple-application-sync-user-using-impersonation.md`) — per-app service account impersonation to reduce control-plane privilege blast radius
- Deep links (`deep-links.md`) — custom external links in the UI for resources
- ApplicationSet deletion strategy for Progressive Sync (`deletion-strategy-progressive-sync.md`)
- Feature bounties index (`feature-bounties.md`) and `feature-bounties/hide-annotations.md`
- Headless Argo CD (`headless-argocd.md`) — core mode without API server/UI
- Manifest Hydrator (`manifest-hydrator.md`, `manifest-hydrator/README.md`, `manifest-hydrator/commit-server/README.md`) — rendered manifest pattern as first-class feature, push-to-deploy vs push-to-stage, commit server component
- Multiple sources for applications UI (`multiple-sources-for-applications-ui.md`) and core proposal (`multiple-sources-for-applications.md`) — `spec.sources` array
- Native OCI support (`native-oci-support.md`) — storing and retrieving manifests from OCI registries
- Notifications API (`notifications-API.md`) — trigger/template/subscription notification system
- Parameterized Config Management Plugins (`parameterized-config-management-plugins.md`) — parameter passing to CMPs
- Project-scoped repositories and clusters (`project-repos-and-clusters.md`, `project-scoped-repository-enhancements.md`)
- Proxy extensions (`proxy-extensions.md`) — reverse-proxy from UI extensions to backend services via API server
- Rebalancing clusters across shards dynamically (`rebalancing-clusters-across-shards-dynamically.md`) — converting controller from StatefulSet to Deployment, dynamic shard assignment
- Resource deletion with approval (`resource-deletion-with-approval.md`) — Prune=confirm pattern
- Respect RBAC for resource exclusions (`respect-rbac-for-resource-exclusions.md`)
- Server-Side Apply (`server-side-apply.md`) — using Kubernetes SSA for syncs to reduce diff noise and manage field ownership
- Server-side pagination (`server-side-pagination.md`) — paginated app/resource listing API
- Sync timeout (`sync-timeout.md`) — configurable timeout for sync operations

## Code areas that affect this folder
- `pkg/apis/application/v1alpha1/` — CRD type changes that correspond to implemented proposals
- `applicationset/` — ApplicationSet controller; relevant when proposals about generators or progressive sync are implemented
- `reposerver/` — manifest generation; relevant for CMP v2 and manifest hydrator proposals
- `server/` — API server; relevant for proxy extensions, pagination, and notifications proposals
- `controller/` — Application Controller; relevant for impersonation, shard rebalancing proposals
- `util/notification/` — relevant when notification API proposals are implemented

## Subfolders
- `feature-bounties/` — Individual bounty proposals for specific features (e.g., hiding annotations in the UI)
- `manifest-hydrator/` — Detailed sub-design for the manifest hydrator feature including the commit server component design
