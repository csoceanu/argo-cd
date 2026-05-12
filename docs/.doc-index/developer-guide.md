# Index: developer-guide/

## What this folder documents
The developer-guide covers everything needed for contributors and third-party integrators: setting up a local development environment, building and testing Argo CD from source, understanding the internal component architecture, authoring and submitting contributions, and extending the UI or API via the extensions system. It targets Argo CD contributors, release engineers, and developers building integrations or plugins.

## Key topics
- Development environment setup: required tools (Go, Git, Docker/Podman, Kind/Minikube/K3d), `kind create cluster`, Tilt for live reload (`tiltfile`)
- Running Argo CD locally: `make start`, `make test`, `make start-e2e`, `make test-e2e`
- E2E tests: `argocd-e2e` namespace, `test/e2e/testdata` as local git repo, `ARGOCD_E2E_APISERVER_PORT` / `ARGOCD_E2E_REPOSERVER_PORT` env vars
- Component architecture: API Server, Application Controller, ApplicationSet Controller, Repo Server, Dex, Redis, Webapp, CLI — layered UI / Application / Core / Infra model
- Authentication/authorization internals: cmux, gRPC-gateway, gRPC interceptors for AuthN, Casbin for RBAC, Session Manager
- API docs: Swagger UI at `/swagger-ui`, bearer token auth via `/api/v1/session`, Applications API `project` parameter for 403/404 disambiguation
- UI extensions: `extensionsAPI.registerResourceExtension`, `^extension(.*)\.js$` files in `/tmp/extensions`, React externals pattern
- Proxy extensions: `server.enable.proxy.extension` feature flag in `argocd-cmd-params-cm`, `extension.config` in `argocd-cm`, reverse-proxy with AuthN/AuthZ
- Release process: quarterly minor releases (Feb/May/Aug/Nov), patch releases on-demand for 3 most recent minors, Release Champion role
- Static code analysis, Mac-specific setup, Gitpod environment, submitting PRs, code contribution guidelines
- Dependency management and toolchain guide

## Code areas that affect this folder
- `server/server.go`, `server/application/` — API server behavior; affects `api-docs.md` and `architecture/authz-authn.md`
- `controller/` — Application Controller reconciliation; affects `architecture/components.md`
- `applicationset/` — ApplicationSet Controller; affects `architecture/components.md`
- `reposerver/` — Repo Server; affects `architecture/components.md`
- `ui/src/` — Web UI source; affects `extensions/ui-extensions.md` and `docs-site.md`
- `test/e2e/` — E2E test suite; affects `test-e2e.md`
- `Makefile` — build targets; affects `running-locally.md` and `development-environment.md`
- `go.mod` — Go version requirements; affects `development-environment.md`

## Subfolders
- `architecture/` — Deep-dive diagrams and explanations of Argo CD component responsibilities, AuthN/AuthZ request flow (cmux, gRPC-gateway, Casbin)
- `extensions/` — Guide for building UI extensions (resource tab, system-level, application-level) and proxy extensions (reverse-proxy backend services)
