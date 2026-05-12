# Index: snyk/

## What this folder documents
The snyk folder contains a single markdown index page that presents weekly automated Snyk vulnerability scan results for Argo CD's `master` branch and the three most recently patched minor releases. It links to generated HTML scan reports for Go modules, the UI yarn.lock, Docker images (argocd, dex, haproxy, redis), and Kubernetes install manifests (IAC scans of `install.yaml` and `namespace-install.yaml`). This folder is maintained automatically and is not manually authored documentation.

## Key topics
- Weekly Snyk scans for `master` branch and the three most recent minor release patches (e.g., v3.1.5, v3.0.16, v2.14.17)
- Vulnerability severity columns: Critical, High, Medium, Low
- Scanned targets per release: `go.mod`, `ui/yarn.lock`, `dex` image, `haproxy` image, `redis` image, `argocd` container image, `install.yaml` (IAC), `namespace-install.yaml` (IAC)
- HTML report files are stored under version-named subdirectories (`master/`, `v3.1.5/`, `v3.0.16/`, `v2.14.17/`) — these are not markdown files

## Code areas that affect this folder
- `go.mod`, `go.sum` — Go dependency updates change scan results for `argocd-test.html`
- `ui/yarn.lock` — UI dependency updates change scan results
- `Dockerfile`, base image pins (haproxy, redis, dex versions in manifests or Dockerfile) — container image updates change image scan results
- `manifests/install.yaml`, `manifests/namespace-install.yaml` — IAC scan targets; RBAC or privilege changes affect IAC findings
- `.github/workflows/` — the CI workflow that runs Snyk and publishes results updates this folder's HTML files
