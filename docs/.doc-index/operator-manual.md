# Index: operator-manual/

## What this folder documents
The operator-manual covers everything a platform team or cluster administrator needs to install, configure, secure, and operate Argo CD at scale. It addresses declarative configuration, RBAC, SSO, TLS, ingress, high availability, multi-cluster management, ApplicationSets, notifications, config management plugins, upgrading, and Prometheus metrics. The audience is operators maintaining shared Argo CD instances for multiple teams.

## Key topics
- Installation: `install.yaml`, `namespace-install.yaml`, `ha/install.yaml`, `ha/namespace-install.yaml`, Helm chart, core (headless) mode
- Declarative setup: `argocd-cm` ConfigMap, `argocd-secret` Secret, `argocd-rbac-cm`, `argocd-cmd-params-cm`, `argocd-tls-certs-cm`, `argocd-ssh-known-hosts-cm`, Application and AppProject CRDs
- RBAC: Casbin policy syntax (`p, <role>, <resource>, <action>, <object>, <effect>`), `role:readonly`, `role:admin`, `policy.default`, anonymous access via `users.anonymous.enabled`
- User management: local users in `argocd-cm` (`accounts.<name>: apiKey, login`), disabling admin, SSO integration overview
- SSO integrations: Dex-based OIDC (Auth0, Google, Okta, Keycloak, Microsoft, identity-center, OneLogin, OpenUnison, Zitadel)
- TLS: `argocd-server-tls` secret, `argocd-repo-server-tls`, `argocd-dex-server-tls`, `--tlsminversion`, cert-manager integration
- Ingress: Ambassador, Nginx, Traefik, AWS ALB, Istio, Contour configurations for gRPC+HTTPS on port 443
- Git webhooks: `/api/webhook` endpoint, GitHub/GitLab/Bitbucket/Bitbucket Server/Azure DevOps/Gogs, `webhook.maxPayloadSizeMB`
- High availability: `argocd-repo-server` parallelism (`--parallelismlimit`), monorepo scaling, `ARGOCD_GIT_ATTEMPTS_COUNT`, `ARGOCD_EXEC_TIMEOUT`, Redis HA
- Cluster management: `argocd cluster add`, dynamic cluster distribution, cluster sharding via `ARGOCD_CONTROLLER_REPLICAS`
- Config Management Plugins (CMP): sidecar plugin pattern, `ConfigManagementPlugin` CRD, `init`/`generate`/`discover` commands
- Health checks: custom health check Lua scripts in `argocd-cm`, `resource.customizations.health`
- Resource actions: built-in and custom actions via Lua scripts in `argocd-cm`, `resource_actions.md`
- Secret management: integrations with Vault, Sealed Secrets, External Secrets; no native secret encryption
- Metrics: Prometheus endpoints (`argocd-metrics:8082/metrics`, `argocd-server-metrics:8083/metrics`, `argocd-repo-server:8084/metrics`), key metrics `argocd_app_info`, `argocd_app_reconcile`, `argocd_app_sync_total`, `argocd_cluster_connection_status`
- Disaster recovery: `argocd admin export/import`
- Web-based terminal: pod exec via UI (`web_based_terminal.md`)
- Deep links: custom external links in `argocd-cm` (`deepLinks`)
- Custom styles: injecting CSS via `argocd-cm` `ui.cssurl`
- Feature maturity levels: Alpha/Beta/Stable
- Signed release assets and supply chain security
- ApplicationSet: generators (Git directory, Git file, Cluster, List, Matrix, Merge, SCM Provider, Pull Request, Cluster Decision Resource, Plugin, Post-Selector), Go Template support, Progressive Syncs, Security considerations for templated `project` field
- Notifications: triggers, templates, subscriptions via `notifications.argoproj.io/subscribe.*` annotations, services (Slack, email, PagerDuty, GitHub, Grafana, Alertmanager, OpsGenie, Teams, Telegram, Webex, Rocket.Chat, Mattermost, New Relic, Pushover, AWS SQS)
- Upgrading: per-version breaking changes guides from v1.0 through v3.2

## Code areas that affect this folder
- `manifests/` — install manifests referenced by `installation.md`
- `manifests/crds/` — CRD specs; changes affect `declarative-setup.md`, `project-specification.md`, `applicationset/applicationset-specification.md`
- `server/` — API server flags; changes affect `argocd-cmd-params-cm-yaml.md`, `tls.md`, `ingress.md`, `security.md`
- `controller/` — Application Controller flags and reconciliation; affects `high_availability.md`, `metrics.md`, `reconcile.md`
- `applicationset/` — ApplicationSet controller; affects all files under `applicationset/`
- `reposerver/` — Repo Server behavior; affects `config-management-plugins.md`, `high_availability.md`
- `util/rbac/`, `util/session/` — RBAC and auth; affects `rbac.md`, `security.md`
- `util/notification/` — Notifications engine; affects all files under `notifications/`
- `assets/builtin-policy.csv` — Default RBAC policy; affects `rbac.md`
- `hack/` — release tooling; affects `signed-release-assets.md`

## Subfolders
- `applicationset/` — Full reference for the ApplicationSet CRD: all generator types, templates, Go templates, Progressive Syncs, security, and the applicationset specification
- `notifications/` — Notification system configuration: triggers, templates, subscriptions, per-service configuration, troubleshooting, and the built-in catalog
- `server-commands/` — Auto-generated CLI flag reference for `argocd-server`, `argocd-application-controller`, `argocd-applicationset-controller`, `argocd-repo-server`, and `argocd-dex`
- `upgrading/` — Per-version migration guides covering breaking changes and required manual steps for each minor release from v1.0 to v3.2
- `user-management/` — SSO integration guides for specific identity providers: Okta, Keycloak, Auth0, Google, Microsoft, AWS Identity Center, OneLogin, OpenUnison, Zitadel
