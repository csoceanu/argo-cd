# OPERATOR-MANUAL/NOTIFICATIONS/SERVICES Documentation Index

## Overview
This documentation area provides comprehensive guides for integrating Argo CD Notifications with various third-party services, including chat platforms, incident management tools, and generic webhooks. It details the configuration of the `argocd-notifications-cm` ConfigMap and `argocd-notifications-secret` Secret, explaining how to define service types, manage sensitive authentication data, and customize notification templates for specific delivery channels.

## Files Summary
*   **overview.md**: Core concepts of notification services, including ConfigMap/Secret structure, sensitive data referencing via `$key` syntax, and support for multiple instances of the same service type using custom names.
*   **alertmanager.md**: Configuration for pushing alerts to Prometheus Alertmanager, supporting HA clusters, basic/bearer auth, and custom labels/annotations.
*   **awssqs.md**: Integration with AWS SQS for queuing messages, including support for FIFO queues, MessageGroupIds, and credential management via environment variables or Secrets.
*   **email.md**: Setup for SMTP-based email notifications, covering host/port configuration, authentication, and HTML support.
*   **github.md**: Guide for using GitHub Apps to update commit statuses, create deployments, and post pull request comments.
*   **googlechat.md**: Configuration for Google Chat webhooks, supporting simple text, CardV2 formats, and message threading.
*   **grafana.md**: Instructions for creating Grafana annotations to visualize deployment events on dashboards using API keys.
*   **mattermost.md**: Integration with Mattermost using bot tokens, featuring compatibility with Slack-style message attachments.
*   **newrelic.md**: Configuration for recording NewRelic deployments, including custom fields for changelogs and commit metadata.
*   **opsgenie.md**: Comprehensive setup for Opsgenie alerts, detailing priorities, responders, tags, and alert de-duplication (aliases).
*   **pagerduty.md**: Setup for PagerDuty incident creation (V1 style) using service IDs and integration tokens.
*   **pagerduty_v2.md**: Integration for PagerDuty Events API v2, supporting severity levels, source identification, and multiple service keys.
*   **pushover.md**: Simple configuration for Pushover mobile/desktop notifications using API and user keys.
*   **rocketchat.md**: Bot-based integration for Rocket.Chat, utilizing Slack-compatible attachment structures for rich messaging.
*   **slack.md**: Extensive guide for Slack bot integration, covering OAuth scopes, message blocks, attachments, threading, and delivery policies.
*   **teams.md**: Webhook-based integration for Microsoft Teams, utilizing "facts," sections, and potential actions (O365 Connectors).
*   **telegram.md**: Setup for Telegram bots using chat IDs or usernames, with support for private chat threads.
*   **webex.md**: Simple bot-based integration for Webex Teams (Cisco Webex) using room IDs or emails.
*   **webhook.md**: Documentation for generic HTTP webhooks, including custom headers, basic auth, retry logic, and templatized request bodies.

## Code Changes That Would Require Documentation Updates
*   **Service Provider Schema Changes**: Adding, removing, or renaming parameters in the notification engine for a specific service (e.g., adding `insecureSkipVerify` to a service that didn't have it).
*   **Authentication Logic Updates**: Changes to how secrets are retrieved or injected into service configurations (e.g., moving from plain tokens to support for workload identity/IAM).
*   **Payload/Template Field Additions**: Introducing new specialized fields within templates for specific services (e.g., adding a new `card` type for Google Chat or a new `potentialAction` type for Teams).
*   **New Service Integration**: Implementing a new notification backend (e.g., Discord, Zoom, or Jira) requires a new service-specific markdown file and an update to `overview.md`.
*   **Retry and Timeout Logic**: Modifying default retry behaviors, backoff strategies, or timeout parameters across the notifications engine or for specific services (like `webhook.md`).
*   **API Version Migration**: Updating a service integration to use a newer version of a third-party API (e.g., migrating from PagerDuty V1 to V2).
*   **Notification Controller Annotations**: Changing the prefix or structure of the subscription annotations (e.g., `notifications.argoproj.io/subscribe`).
*   **Template Metadata/Context Changes**: Adding new variables to the global template context (like `argocdUrl`) that users should be aware of for their message formatting.

## Key Technical Concepts
*   **ConfigMap `argocd-notifications-cm`**: The primary configuration hub for defining service instances and templates.
*   **Secret `argocd-notifications-secret`**: The secure store for tokens, passwords, and private keys.
*   **Sensitive Data Referencing**: The use of `$key` syntax to pull values from Secrets into the ConfigMap.
*   **Service Custom Naming**: The `service.<type>.<name>` format allowing multiple instances of one service (e.g., two different Slack workspaces).
*   **Go Templating**: The use of `{{ ... }}` syntax and helper functions (like `call .repo.GetCommitMetadata`) to generate dynamic content.
*   **Message Aggregation/Threading**: Concepts like `groupingKey` (Slack) or `threadKey` (Google Chat) used to collapse multiple notifications into a single conversation.
*   **Subscription Annotations**: The mechanism for attaching specific triggers and services to Argo CD Application resources.
*   **Slack-Compatible Attachments**: A common format used by Mattermost and Rocket.Chat for rich message layout.
*   **Retry Policy**: Parameters like `retryMax`, `retryWaitMin`, and `retryWaitMax` specifically used in webhook configurations.

## Related Components
*   **Argo CD Notifications Controller**: The service responsible for watching applications and dispatching notifications.
*   **Argo CD Notifications Engine**: The underlying library that handles template parsing and service integrations.
*   **Argo CD Repository Server**: Source of commit metadata used in templates (author, message, revision).
*   **Argo CD Application Controller**: Provides the application state and status used as triggers.
*   **External Service APIs**: (Slack API, GitHub API, PagerDuty Events API, etc.) The targets for the notification payloads.