# OPERATOR-MANUAL/NOTIFICATIONS Documentation Index

## Overview
Argo CD Notifications is a specialized controller that monitors Argo CD applications and alerts users to state changes (e.g., sync failures, health degradation, or successful deployments). The system utilizes a flexible architecture of triggers (conditions), templates (content), and subscriptions (recipients) to integrate with various communication services like Slack, Email, and Webhooks.

## Files Summary
*   **operator-manual/notifications/index.md**: Provides a high-level introduction, quick-start installation guide for the catalog, and instructions for namespace-based self-service configuration.
*   **operator-manual/notifications/catalog.md**: A reference list of pre-defined triggers and templates available in the community-maintained catalog for standard application events.
*   **operator-manual/notifications/triggers.md**: Explains how to define conditions using the `expr` language and how to use the `oncePer` field to prevent notification fatigue.
*   **operator-manual/notifications/templates.md**: Details the creation of notification messages using Golang templates, including the use of shared context, secrets, and service-specific fields.
*   **operator-manual/notifications/functions.md**: A comprehensive reference of built-in helper functions for time manipulation, string formatting, and retrieving repository/commit metadata.
*   **operator-manual/notifications/subscriptions.md**: Describes how to subscribe applications or projects to specific triggers using annotations or global configuration selectors.
*   **operator-manual/notifications/examples.md**: Offers practical, end-to-end examples for advanced use cases like image change tracking, Slack formatting, and webhook resource analysis.
*   **operator-manual/notifications/troubleshooting.md**: An overview of diagnostic strategies, including global CLI flags, Kustomize integration, and cluster-access methods.
*   **operator-manual/notifications/troubleshooting-commands.md**: A detailed technical reference for the `argocd admin notifications` CLI tool, covering template testing and trigger evaluation.
*   **operator-manual/notifications/troubleshooting-errors.md**: A guide to resolving common configuration errors, YAML syntax issues, and challenges associated with multi-source applications.
*   **operator-manual/notifications/monitoring.md**: Documents the Prometheus metrics exported by the controller for tracking delivery success rates and trigger evaluations.

## Code Changes That Would Require Documentation Updates
*   **Notification Controller Logic**: Changes to how triggers are evaluated, the introduction of new built-in variables in the template `context`, or modifications to the `oncePer` deduplication logic.
*   **New Integration Services**: Adding support for new notification providers (e.g., Discord, PagerDuty, Matrix) would require updates to service registration and template fields.
*   **Template Function Additions**: Introducing new helper functions in the `functions` package (e.g., new `repo`, `strings`, or `sync` helpers).
*   **CLI Tooling**: Adding new sub-commands to `argocd admin notifications` or changing existing flags and output formats.
*   **Metrics & Observability**: Adding or renaming Prometheus metrics in the controller or changing the default metrics port.
*   **Multi-Source Application Support**: Any changes in how Argo CD handles status for applications with multiple sources, as this affects template syntax (e.g., `revisions` vs `revision`).
*   **Security & Permissions**: Changes to how secrets are accessed from templates or modifications to the mandatory labels (e.g., `app.kubernetes.io/part-of: argocd`) required for secret discovery.
*   **Catalog Updates**: Adding new standard triggers or templates to the official Argo CD notifications catalog.

## Key Technical Concepts
*   **Triggers**: Conditional expressions powered by `antonmedv/expr` that determine if a notification should be sent.
*   **Templates**: Golang `html/template` definitions that generate the notification payload.
*   **Subscriptions**: The link between an Application/Project and a Trigger/Service destination, usually managed via the `notifications.argoproj.io/subscribe` annotation.
*   **oncePer**: A mechanism to ensure notifications are only sent once per unique value (e.g., per Git revision or per annotation version).
*   **argocd-notifications-cm**: The central ConfigMap for defining services, triggers, templates, and global subscriptions.
*   **argocd-notifications-secret**: The core secret for storing sensitive integration tokens and credentials.
*   **Template Functions**: Specific helpers like `repo.GetCommitMetadata`, `sync.GetInfoItem`, and `time.Parse`.
*   **Self-Service Notifications**: A feature allowing configuration of notifications within individual application namespaces using `--self-service-notification-enabled`.
*   **CLI Commands**: `template notify` (for testing delivery) and `trigger run` (for testing conditions).

## Related Components
*   **argocd-notifications-controller**: The primary microservice executing the logic documented here.
*   **argocd-repo-server**: Used by notification functions to fetch commit and manifest metadata.
*   **argocd-application-controller**: Manages the Application CRDs that triggers monitor.
*   **Argo CD API Server**: Handles the CLI requests for notification troubleshooting.
*   **External Notification Services**: Slack, Microsoft Teams, SMTP (Email), GitHub (Commit Status), and generic Webhook endpoints.