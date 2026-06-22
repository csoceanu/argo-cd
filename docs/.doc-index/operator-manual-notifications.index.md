# OPERATOR-MANUAL/NOTIFICATIONS Documentation Index

## Overview
Argo CD Notifications is a dedicated service that monitors Application state changes and alerts users through various communication channels. It utilizes a flexible system of logic-based triggers and Go-template-based notification content, allowing operators to automate alerts for sync status, health degradation, and resource changes.

## Files Summary
*   **examples.md**: Practical walkthroughs for complex scenarios like resource change detection via webhooks and sending image lists to Slack.
*   **functions.md**: Detailed reference for built-in template functions including time parsing, string manipulation, and repository metadata retrieval.
*   **catalog.md**: A library of pre-defined, "ready-to-use" triggers and templates for common lifecycle events (created, deleted, synced, failed).
*   **templates.md**: Instructions on defining notification content using Go templates, managing shared context, and utilizing secrets within messages.
*   **subscriptions.md**: Guidance on how to link applications and projects to triggers via annotations or global configuration.
*   **troubleshooting-commands.md**: CLI reference for the `argocd admin notifications` toolset, focusing on template and trigger inspection.
*   **triggers.md**: Documentation on defining the conditions (logic) for sending notifications, including "flapping" prevention using `oncePer`.
*   **troubleshooting.md**: General debugging strategies, CLI installation guides, and cluster-based configuration validation.
*   **troubleshooting-errors.md**: A guide to resolving common errors like YAML syntax issues, multi-source application pathing, and secret permission errors.
*   **monitoring.md**: Reference for Prometheus metrics exported by the controller for tracking delivery rates and trigger evaluation.
*   **index.md**: Entry-point documentation covering installation, initial setup, and advanced namespace-based (self-service) configuration.

## Code Changes That Would Require Documentation Updates
*   **Template Engine**: Changes to the underlying Go template implementation, the introduction of new built-in functions in `functions.md`, or updates to the [Sprig](https://masterminds.github.io/sprig/) package integration.
*   **Trigger Logic**: Modifications to the expression evaluation engine ([antonmedv/expr](https://github.com/antonmedv/expr)) or changes to how `oncePer` deduplication logic is calculated.
*   **New Integration Services**: Adding support for new notification providers (e.g., Discord, Matrix, Opsgenie) or changing the configuration schema for existing ones (Slack, Webhook, Teams).
*   **CLI Tooling**: Adding new sub-commands or flags to `argocd admin notifications`, or changing the output format of existing troubleshooting commands.
*   **CRD/Schema Changes**: Updates to the Argo CD `Application` or `AppProject` specs that expose new fields which could be used in template variables (e.g., `.app.status.newField`).
*   **Multi-Source Application Logic**: Changes to how Argo CD handles applications with multiple sources, specifically how `syncResult` or `revisions` are indexed in the status block.
*   **Controller Configuration**: Alterations to environment variables (like `TZ`), default metrics ports (9001), or new startup flags for the `argocd-notifications-controller`.
*   **Self-Service Feature**: Changes to the namespace-based configuration logic, such as the flags `--application-namespaces` or `--self-service-notification-enabled`.
*   **Authentication/Security**: Changing the labels required for secret discovery (`app.kubernetes.io/part-of: argocd`) or how secrets are injected into templates.

## Key Technical Concepts
*   **Triggers**: Logic predicates that return a boolean to decide if an alert should be sent.
*   **Templates**: YAML/Go-template definitions that format the actual message sent to the user.
*   **Subscriptions**: The binding mechanism (annotations) between an Application/Project and a Trigger/Service pair.
*   **`oncePer`**: A deduplication field used to prevent redundant alerts (e.g., only notifying once per Git revision).
*   **Service Types**: Specific notification backends (Slack, Email, Webhook, Teams, PagerDutyV2).
*   **Context**: A user-defined string map in the ConfigMap used to share global variables across all templates.
*   **Built-in Functions**: Custom helper functions like `repo.GetCommitMetadata`, `time.Now`, and `sync.GetInfoItem`.
*   **Self-Service Notifications**: Decentralized configuration allowing users to define notifications within their own app namespaces.
*   **Metrics**: Prometheus endpoints (`argocd_notifications_deliveries_total`, `argocd_notifications_trigger_eval_total`).

## Related Components
*   **argocd-notifications-controller**: The primary binary/deployment responsible for evaluating triggers.
*   **argocd-notifications-cm**: The central ConfigMap containing all logic, templates, and service configurations.
*   **argocd-notifications-secret**: The Kubernetes Secret used to store sensitive integration tokens and passwords.
*   **Argo CD Application Controller**: The source of state changes that triggers the notifications system.
*   **Argo CD CLI**: Specifically the `admin notifications` sub-commands used for debugging.
*   **Argo CD Repo Server**: Accessed by the notification service to retrieve commit metadata and repository details.