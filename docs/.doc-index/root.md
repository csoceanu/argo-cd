# Index: docs/ (root)

## What this folder documents
The docs root covers introductory and top-level Argo CD content aimed at all audiences: new users getting started, general product overview, core concepts, CLI installation, and high-level security considerations. It serves as the entry point to the full documentation site and links out to the operator, user, and developer sections.

## Key topics
- GitOps pattern and why Argo CD uses Git repositories as the source of truth
- Application, sync status, live state, target state, health, refresh — core Argo CD terminology
- Getting started: installing Argo CD (`install.yaml`), downloading the `argocd` CLI, accessing the API server via LoadBalancer / Ingress / port-forward, logging in, creating and syncing a first application
- CLI installation methods: Homebrew, curl, ArchLinux, WSL, macOS, Windows (manual and package manager)
- `try_argo_cd_locally.md` — running Argo CD locally without a cloud cluster
- FAQ: troubleshooting OutOfSync, Progressing state for Ingress/StatefulSet/SealedSecret, resetting admin password via `argocd-secret` bcrypt hash
- Security considerations (archived CVE table): CVE-2020-5260 git credential leak, CVE-2020-8828 insecure default admin password, CVE-2020-8827 brute-force, CVE-2020-8826 session fixation, CVE-2018-21034 info disclosure
- Roadmap and project adoption (USERS.md reference)
- CONTRIBUTING.md and SUPPORT.md for community engagement
- Bug triage process (`bug_triage.md`)

## Code areas that affect this folder
- `manifests/install.yaml`, `manifests/ha/install.yaml`, `manifests/namespace-install.yaml` — installation manifests referenced in getting_started.md
- `cmd/argocd/` — CLI entry point; changes to CLI commands affect `cli_installation.md` and `getting_started.md`
- `server/` — API server; security or auth changes affect `security_considerations.md`
- `util/session/` — session management; token expiry changes affect security docs
- Release scripts / `CHANGELOG` — roadmap.md and version references
