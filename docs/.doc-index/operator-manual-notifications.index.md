# OPERATOR-MANUAL/NOTIFICATIONS Documentation Index

## Overview
This documentation area describes the Argo CD Notifications service, a system that monitors Argo CD applications and alerts users about state changes (e.g., sync failures, health degradation). It covers the configuration of the notification lifecycle: defining **Triggers** (the "when"), **Templates** (the "what"), **Subscriptions** (the "who"), and **Services** (the "how" or delivery mechanism).

## Files Summary
*   **index.md**: The entry point providing a quick-start guide and details on namespace-based/self-service configuration.
*   **catalog.md**: A reference list of pre-configured triggers and templates available in the stable distribution.
*   **triggers.md**: Documentation on defining conditions using the `expr` language, including deduplication logic via `oncePer`.
*   **templates.md**: Instructions for crafting messages using Golang templates, custom context, and access to secrets.
*   **functions.md**: A reference of specialized functions available within templates (time, strings, repo metadata, sync info).
*   **subscriptions.md**: Explains how to subscribe applications or projects to notifications using annotations or global selectors.
*   **monitoring.md**: Details on Prometheus metrics exposed by the notifications controller for tracking deliveries and evaluations.
*   **examples.md**: Practical walkthroughs for complex scenarios like monitoring resource changes and Slack image listing.
*   **troubleshooting.md**: General debugging strategies and how to use the CLI tools.
*   **troubleshooting-commands.md**: Reference for `argocd admin notifications` CLI commands (get, notify, run).
*   **troubleshooting-errors.md**: A guide to common configuration errors, including multi-source application pitfalls.
*   **services/overview.md**: Introduction to notification service types and sensitive data handling.
*   **services/[service-name].md**: Individual deep-dives for specific integrations (Slack, GitHub, Webhook, PagerDuty, Alertmanager, etc.).

## Code Changes That Would Require Documentation Updates
*   **Notification Engine & Services**: 
    *   Adding support for a new notification provider (e.g., Discord, Matrix).
    *   Updating existing service parameters (e.g., adding `priority` to PagerDuty or `blocks` to Slack).
    *   Changing default retry logic or timeout settings for webhooks.
*   **Template Engine**:
    *   Adding new custom functions to the `repo`, `sync`, `strings`, or `time` packages.
    *   Changing the schema of the objects passed to templates (the `app`, `context`, or `secrets` variables).
    *   Upgrading the Golang template version or the Sprig library integration.
*   **Trigger Logic**:
    *   Modifying the expression evaluation engine (`antonmedv/expr`).
    *   Adding new default triggers to the built-in catalog.
    *   Changing the behavior of `oncePer` (e.g., how it handles mono-repos).
*   **Controller Infrastructure**:
    *   Changing the controller's CLI flags or environment variables (e.g., `--metrics-port`, `TZ`).
    *   Updating the self-service/namespace-based configuration logic (`--application-namespaces`).
    *   Modifying how the controller interacts with the `argocd-notifications-cm` or secrets.
*   **CLI Tools**:
    *   Adding new sub-commands to `argocd admin notifications`.
    *   Changing the output format or flags of existing troubleshooting commands.
*   **Observability**:
    *   Adding, renaming, or removing Prometheus metrics.
    *   Changing the labels associated with `argocd_notifications_deliveries_total` or `argocd_notifications_trigger_eval_total`.

## Key Technical Concepts
*   **CRDs & Config**: `argocd-notifications-cm` (ConfigMap), `argocd-notifications-secret` (Secret).
*   **Annotations**: `notifications.argoproj.io/subscribe.<trigger>.<service>`.
*   **Expression Language**: `antonmedv/expr` used in `when` conditions.
*   **Deduplication**: `oncePer` logic based on app status fields or annotations.
*   **Template Variables**: `.app` (Application object), `.context` (shared string map), `.secrets`, `.serviceType`, `.recipient`.
*   **Template Functions**: `repo.GetCommitMetadata`, `repo.RepoURLToHTTPS`, `sync.GetInfoItem`, `time.Now`.
*   **CLI Operations**: `template notify` (test rendering/sending), `trigger run` (test evaluation).
*   **Service Configurations**: OAuth tokens, Webhook URLs, Basic Auth, SMTP settings.
*   **Multi-source Apps**: Handling arrays in `app.spec.sources` and `app.status.operationState.syncResult.revisions`.

## Related Components
*   **Argo CD Application Controller**: The source of truth for application status and health.
*   **Argo CD Repo Server**: Accessed by the notifications controller to retrieve commit metadata (authors, messages).
*   **Argo CD API Server**: Used by the CLI for authentication and resource management.
*   **External Integration Providers**: Slack API, GitHub Apps, PagerDuty Events API v2, SMTP Servers, Alertmanager.
*   **Prometheus Stack**: For consuming metrics and alerting on notification delivery failures.