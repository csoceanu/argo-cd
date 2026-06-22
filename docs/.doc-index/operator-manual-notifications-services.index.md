# OPERATOR-MANUAL/NOTIFICATIONS/SERVICES Documentation Index

## Overview
This documentation area provides comprehensive configuration guides for integrating Argo CD notifications with various external communication platforms and services. It details how to define notification services in the `argocd-notifications-cm` ConfigMap, manage sensitive credentials in Kubernetes Secrets, and customize notification delivery through templates and subscriptions.

## Files Summary
*   **overview.md**: Provides the foundational concepts for service configuration, including sensitive data handling, custom naming for multiple instances, and a directory of supported service types.
*   **alertmanager.md**: Details integration with Prometheus Alertmanager, covering authentication, target clusters, and routing via labels and annotations.
*   **awssqs.md**: Explains how to send notification messages to AWS SQS queues, including FIFO queue support and credential management via environment variables or secrets.
*   **email.md**: Configures SMTP-based email notifications, supporting custom subjects and HTML formatting.
*   **github.md**: Covers integration with GitHub Apps to update commit statuses, deployment environments, and pull request comments.
*   **googlechat.md**: Describes webhook-based integration with Google Chat, supporting text messages, CardV2 layouts, and chat threads.
*   **grafana.md**: Explains how to create Grafana annotations via API keys to track deployments and sync events on dashboards.
*   **mattermost.md**: Provides setup for Mattermost bot integrations, utilizing Slack-compatible attachment formats.
*   **newrelic.md**: Details how to record deployment events in New Relic, including changelog and user metadata mapping.
*   **opsgenie.md**: Configures incident alerting in Opsgenie, with support for team-based API keys and detailed alert property mapping (priority, alias, responders).
*   **pagerduty.md**: Focuses on PagerDuty incident creation for Argo CD Applications and Rollouts using the classic integration method.
*   **pagerduty_v2.md**: Details the modern Events API v2 integration for PagerDuty, supporting advanced severity levels and service key dictionaries.
*   **pushover.md**: A concise guide for mobile push notifications via the Pushover API.
*   **rocketchat.md**: Covers Rocket.Chat bot configuration, supporting Slack-compatible attachments and channel/room ID targeting.
*   **slack.md**: A deep dive into Slack App integration, featuring message blocks, attachments, threading (grouping keys), and delivery policies.
*   **teams.md**: Explains Microsoft Teams integration via webhooks, utilizing "facts" and "potentialAction" for interactive notifications.
*   **telegram.md**: Guides the setup of Telegram bots for public channels or private group chats with thread support.
*   **webex.md**: Brief instructions for integrating Cisco Webex Teams via bot tokens.
*   **webhook.md**: A flexible guide for sending generic HTTP requests (POST, PUT, etc.) to any external API, including retry logic and custom header support.

## Code Changes That Would Require Documentation Updates
*   **New Service Integration**: Adding a new notification provider (e.g., Discord, Matrix) requires a new service-specific Markdown file and an entry in `overview.md`.
*   **Parameter Schema Changes**: Adding, renaming, or removing configuration fields in the Go structs representing service settings (e.g., adding `timeout` to the Email service).
*   **Authentication Logic Updates**: Changes to how secrets are referenced or how authentication is performed (e.g., moving from API keys to OIDC or IAM Roles).
*   **Template Variable Changes**: Modifying the data structure passed to templates (the `app` or `context` objects) which would break existing examples.
*   **API Version Upgrades**: Updating the implementation to support newer versions of external APIs (e.g., migrating from Slack attachments to Blocks or GitHub API v3 to v4).
*   **Retry/Network Behavior**: Changing the default retry logic, backoff strategies, or timeout handling globally or for a specific service like `webhook`.
*   **Annotation Syntax**: Altering the prefix or structure of the `notifications.argoproj.io` annotations used in Application or Rollout resources.
*   **CRD Support Expansion**: If notifications are extended to support new resource types beyond `Application` and `Rollout`.

## Key Technical Concepts
*   **Service Definition**: Configuration mapping under `service.<type>.<name>` in the ConfigMap.
*   **Sensitive Data Referencing**: Using the `$` prefix to pull values from `argocd-notifications-secret`.
*   **Subscriptions**: Annotations on resources that link triggers to specific notification services and recipients.
*   **Dynamic Templating**: Use of Go templates to inject application metadata (`{{.app.metadata.name}}`) into notification payloads.
*   **Service Custom Names**: Logic allowing multiple instances of the same service type (e.g., `service.slack.internal` and `service.slack.external`).
*   **Webhook Method/Body**: Customizable HTTP verbs and payload formats for generic integrations.
*   **Slack Blocks/Attachments**: Specialized JSON structures for rich messaging in Slack and compatible platforms (Mattermost, Rocket.Chat).
*   **Grouping Keys**: Logic for aggregating multiple notifications into a single thread (Slack, Google Chat).
*   **Delivery Policy**: Rules for whether a notification creates a new message or updates an existing one (e.g., `PostAndUpdate`).

## Related Components
*   **Argo CD Application Controller**: The core component monitoring application state.
*   **Argo Rollouts**: Used in examples for `on-rollout-aborted` and related triggers.
*   **Kubernetes ConfigMaps & Secrets**: The primary storage for configuration and credentials.
*   **Notification Engine**: The underlying library that processes triggers and dispatches messages.
*   **External API Gateways**: GitHub API, Slack API, PagerDuty Events API, etc.