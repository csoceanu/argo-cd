# OPERATOR-MANUAL/NOTIFICATIONS Documentation Index

## Overview
This documentation area provides a comprehensive guide to the Argo CD Notifications service, a system that monitors Argo CD applications and alerts users about state changes. It details the configuration of triggers, templates, and subscriptions, as well as troubleshooting techniques and monitoring practices for administrators.

## Files Summary
*   **operator-manual/notifications/index.md**: Serves as the primary entry point, covering installation from the catalog, basic email setup, and the self-service "applications in any namespace" configuration.
*   **operator-manual/notifications/catalog.md**: Provides a reference list of the pre-defined triggers (e.g., `on-deployed`, `on-sync-failed`) and templates (e.g., `app-sync-succeeded`) included in the official Argo CD catalog.
*   **operator-manual/notifications/triggers.md**: Explains how to define the conditions that initiate a notification, including the use of expression language, condition bundles, and the `oncePer` field for deduplication.
*   **operator-manual/notifications/templates.md**: Details how to construct notification content using Golang templates, including how to reference application fields, secrets, and custom context.
*   **operator-manual/notifications/subscriptions.md**: Describes the various ways to subscribe to notifications via Application/AppProject annotations or global configuration in the `argocd-notifications-cm`.
*   **operator-manual/notifications/functions.md**: A technical reference of built-in Golang-based functions available for use within triggers and templates, covering time, strings, and repository metadata.
*   **operator-manual/notifications/examples.md**: Offers practical, advanced configuration examples such as notifying on sync changes via webhooks and sending image lists to Slack.
*   **operator-manual/notifications/monitoring.md**: Lists the Prometheus metrics exported by the notifications controller and provides information for Grafana dashboarding.
*   **operator-manual/notifications/troubleshooting.md**: Provides a high-level overview of diagnostic tools, including Kustomize integration and cluster-based debugging steps.
*   **operator-manual/notifications/troubleshooting-commands.md**: A detailed reference for the `argocd admin notifications` CLI toolset, including commands for testing templates and evaluating triggers.
*   **operator-manual/notifications/troubleshooting-errors.md**: Catalogs common configuration and runtime errors, with a specific focus on issues arising from multi-source applications and YAML syntax.

## Code Changes That Would Require Documentation Updates
*   **Application/AppProject CRD Changes**: Any modification to the `ApplicationSpec`, `ApplicationStatus`, or `AppProject` structures, as notification templates and triggers directly reference these fields.
*   **New Notification Drivers**: Implementation of new notification services (e.g., a new integration for Discord, Matrix, or generic SMS).
*   **Expression Engine Updates**: Upgrading or changing the behavior of the `antonmedv/expr` library used for trigger conditions.
*   **Custom Function Additions**: Adding new helper functions to the `functions` package (e.g., new `repo`, `sync`, or `strings` helpers).
*   **CLI Flag/Command Modifications**: Changes to the `argocd admin notifications` command structure, arguments, or output formats.
*   **Multi-Source Logic**: Changes to how Argo CD handles multiple sources in an application, specifically how `syncResult` or `revisions` are indexed.
*   **Namespace-Based Controller Logic**: Altering how the controller discovers `ConfigMaps` and `Secrets` in non-default namespaces or changes to the `--self-service-notification-enabled` logic.
*   **Metrics Collection**: Adding or renaming Prometheus metrics in the notifications controller.
*   **Default Catalog Updates**: Adding new triggers or templates to the official `notifications_catalog/install.yaml`.

## Key Technical Concepts
*   **Configuration Objects**: `argocd-notifications-cm` (ConfigMap), `argocd-notifications-secret` (Secret).
*   **Annotations**: `notifications.argoproj.io/subscribe.<trigger>.<service>`.
*   **Trigger Fields**: `when` (condition logic), `oncePer` (deduplication key), `send` (template reference).
*   **Template Context**: `.app` (Application object), `.context` (shared variables), `.secrets` (sensitive data), `.serviceType` (slack, email, etc.).
*   **Built-in Functions**: `repo.GetCommitMetadata`, `repo.RepoURLToHTTPS`, `time.Now`, `strings.ReplaceAll`, `sync.GetInfoItem`.
*   **CLI Subcommands**: `template get`, `template notify`, `trigger get`, `trigger run`.
*   **Metrics**: `argocd_notifications_deliveries_total`, `argocd_notifications_trigger_eval_total`.
*   **Delivery Policy**: Post-sync behaviors, grouping keys, and broadcast settings within service-specific template fields.

## Related Components
*   **Argo CD Application Controller**: The source of state changes that the notifications service monitors.
*   **Argo CD Repo Server**: Accessed via `repo` functions to retrieve commit metadata and repository details.
*   **Argo CD API Server**: Used by the CLI to interact with the cluster configuration.
*   **External Integration Services**: Slack, Microsoft Teams, PagerDuty, Email (SMTP), and Webhooks.
*   **Prometheus Stack**: For consuming the metrics exported on port 9001.