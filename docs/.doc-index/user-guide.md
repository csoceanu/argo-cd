# Index: user-guide/

## What this folder documents
The user-guide targets application developers and GitOps practitioners who use Argo CD to deploy and manage applications. It covers application sources (Helm, Kustomize, Jsonnet, plain YAML, OCI, directory), sync behavior, projects, private repositories, CI automation, and the full `argocd` CLI command reference. This is the primary reference for day-to-day Argo CD usage.

## Key topics
- Application sources: Helm charts (values files, `--values`, `valuesObject`, `helm.releaseName`, OCI charts), Kustomize overlays (`namePrefix`, `nameSuffix`, `images`, `replicas`, `patches`, `components`), Jsonnet, plain-YAML directory, OCI artifacts (`oci://` scheme)
- Multiple sources: `spec.sources` array, combining manifests from multiple repos
- Source Hydrator (rendered manifest pattern, Alpha): `hydrator.enabled` in `argocd-cmd-params-cm`, commit server, `install-with-hydrator.yaml`
- Tracking strategies: branch/HEAD tracking, tag tracking (semver ranges via `>=`/`*`), commit pinning
- Auto-sync: `spec.syncPolicy.automated`, `prune`, `selfHeal`, `allowEmpty`
- Sync options: `Prune=false`, `Prune=confirm`, `Validate=false`, `SkipDryRunOnMissingResource=true`, `ApplyOutOfSyncOnly=true`, `ServerSideApply=true`, `Replace=true`, `RespectIgnoreDifferences=true`, `CreateNamespace=true`, `PrunePropagationPolicy`
- Sync waves: `argocd.argoproj.io/sync-wave` annotation ordering resource creation/deletion
- Resource hooks: `argocd.argoproj.io/hook` annotation (PreSync, Sync, PostSync, SyncFail, PostDelete), `argocd.argoproj.io/hook-delete-policy`
- Selective sync: syncing only OutOfSync resources, `--selective-sync` flag
- Diffing customization: `ignoreDifferences` with `jsonPointers`, `jqPathExpressions`, `managedFieldsManagers`; system-level `resource.customizations.ignoreDifferences`
- Compare options: `IgnoreExtraneous`, `ServerSideDiff` in `spec.ignoreDifferences`
- Diff strategies: `last-applied-configuration` vs server-side apply field manager
- Projects (AppProject): source repo restrictions, destination cluster/namespace restrictions, resource whitelist/blacklist, project roles, sync windows, `orphanedResources` monitoring
- Private repositories: HTTPS credentials, SSH keys, GitHub App auth, TLS client certificates, `argocd repo add`, credential templates (`argocd repocreds add`)
- GnuPG signature verification: `ARGOCD_GPG_ENABLED`, per-project enforcement, `argocd gpg add`
- Resource tracking: label vs annotation vs `managedFields` tracking strategies (`resource.compareoptions.trackingMethod`)
- Orphaned resources: `spec.orphanedResources.warn`, `argocd.argoproj.io/managed-by` label
- Annotations and labels: `argocd.argoproj.io/app-name`, `notifications.argoproj.io/subscribe.*`, deep links
- CI/CD automation: `argocd app sync`, `argocd app wait`, `ARGOCD_AUTH_TOKEN` env var, `argocd login --core`
- Environment variables: `ARGOCD_OPTS`, `ARGOCD_SERVER`, `ARGOCD_AUTH_TOKEN`, `ARGOCD_GRPC_KEEPALIVE_*`
- Status badge: `argocd.argoproj.io/badge` endpoint
- Extra info: `info` field in Application spec for custom key-value display
- Skip reconcile: `argocd.argoproj.io/skip-reconcile` annotation
- Scale application resources: `argocd app actions run <app> restart --kind Deployment`
- Import/export: `argocd admin export`, `argocd admin import`
- Config management plugins (user perspective): specifying `spec.source.plugin.name`, passing env vars via `spec.source.plugin.env`
- Build environment variables available to CMP: `ARGOCD_APP_NAME`, `ARGOCD_APP_NAMESPACE`, `ARGOCD_ENV_*`, `KUBE_VERSION`, `KUBE_API_VERSIONS`
- Tool detection: how Argo CD auto-detects Helm, Kustomize, Jsonnet, or plain-YAML based on repository contents
- `argocd` CLI: full command tree including `app`, `appset`, `proj`, `repo`, `repocreds`, `cluster`, `cert`, `gpg`, `account`, `admin`, `context`, `login`, `logout`, `configure`

## Code areas that affect this folder
- `cmd/argocd/` — CLI implementation; changes affect all files under `commands/`
- `util/helm/` — Helm rendering; affects `helm.md`
- `util/kustomize/` — Kustomize rendering; affects `kustomize.md`
- `util/diff/` — Diff engine; affects `diffing.md`, `diff-strategies.md`, `compare-options.md`
- `controller/sync.go`, `controller/sync_namespace.go` — sync logic; affects `sync-options.md`, `sync-waves.md`
- `reposerver/repository/` — manifest generation, CMP, source hydrator; affects `config-management-plugins.md`, `source-hydrator.md`
- `pkg/apis/application/v1alpha1/` — Application and AppProject CRD types; affects `application-specification.md`, `projects.md`
- `util/gpg/` — GPG verification; affects `gpg-verification.md`
- `util/resource/` — resource tracking logic; affects `resource_tracking.md`
- `server/application/` — app actions, resource operations; affects sync-related docs and `scale_application_resources.md`

## Subfolders
- `commands/` — Auto-generated reference pages for every `argocd` CLI subcommand (app, appset, proj, repo, repocreds, cluster, cert, gpg, account, admin, context, login, logout, configure, completion, version, relogin)
