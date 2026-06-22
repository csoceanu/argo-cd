# OPERATOR-MANUAL/NOTIFICATIONS/SERVICES Documentation Index

## Overview
This documentation area provides comprehensive configuration guides for integrating Argo CD Notifications with various third-party communication platforms, alerting tools, and CI/CD services. It details how to configure the `argocd-notifications-cm` ConfigMap and `argocd-notifications-secret` Secret to enable automated messaging based on application lifecycle events.

## Files Summary
*   **operator-manual/notifications/services/overview.md**: Provides the foundational concepts of notification services, including global configuration syntax, handling sensitive data via secrets, and implementing custom service names.
*   **operator-manual/notifications/services/alertmanager.md**: Details integration with Prometheus Alertmanager, covering authentication methods (Basic and Bearer), custom API paths, and template-based labels/annotations.
*   **operator-manual/notifications/services/awssqs.md**: Explains how to send notification messages to AWS SQS queues, including credential management, region settings, and support for FIFO queues.
*   **operator-manual/notifications/services/email.md**: Covers SMTP-based email notification setup, including server parameters, authentication, and HTML message support.
*   **operator-manual/notifications/services/github.md**: Describes integration with GitHub via GitHub Apps to update commit statuses, create deployments, or post pull request comments.
*   **operator-manual/notifications/services/googlechat.md**: Instructions for sending text and CardV2 messages to Google Chat spaces using webhooks and thread keys.
*   **operator-manual/notifications/services/grafana.md**: Explains how to create Grafana annotations via the Grafana HTTP API to mark deployment events on dashboards.
*   **operator-manual/notifications/services/mattermost.md**: Guides users on setting up Mattermost bot integrations, highlighting its compatibility with Slack-style attachments.
*   **operator-manual/notifications/services/newrelic.md**: Details how to send deployment markers to New Relic, including configuration for API keys and application IDs.
*   **operator-manual/notifications/services/opsgenie.md**: Covers integration with Opsgenie for incident management, including responder settings, alert priorities, and custom tags.
*   **operator-manual/notifications/services/pagerduty.md**: Provides configuration for the PagerDuty service using the integration service ID to trigger incidents.
*   **operator-manual/notifications/services/pagerduty_v2.md**: Details the advanced PagerDuty integration using the Events API v2, supporting severity levels and service keys.
*   **operator-manual/notifications/services/pushover.md**: A concise guide for integrating mobile push notifications via the Pushover API.
*   **operator-manual/notifications/services/rocketchat.md**: Explains how to configure a Rocket.Chat bot user to deliver notifications to channels or teams.
*   **operator-manual/notifications/services/slack.md**: A deep dive into Slack integration via OAuth, covering message blocks, attachments, threading, and delivery policies.
*   **operator-manual/notifications/services/teams.md**: Instructions for using Microsoft Teams incoming webhooks to send messages featuring sections, facts, and potential actions.
*   **operator-manual/notifications/services/telegram.md**: Guides setup for Telegram bots, including public channel usernames and private chat IDs with thread support.
*   **operator-manual/notifications/services/webex.md**: Covers integration with Webex Teams using bot tokens for delivery to personal emails or room IDs.
*   **operator-manual/notifications/services/webhook.md**: Provides a generic framework for sending HTTP requests (GET, POST, PUT, PATCH) to custom endpoints with configurable headers and retry logic.

## Code Changes That Would Require Documentation Updates
*   **New Service Implementation**: Adding support for a new third-party platform (e.g., Discord, Matrix).
*   **Schema Changes in ConfigMap**: Modifying the expected keys under `service.<type>` in `argocd-notifications-cm`.
*   **Secret Handling Logic**: Changing the way sensitive data is referenced (the `$` syntax) or adding support for external secret stores.
*   **Template Variable Changes**: Adding, removing, or renaming the metadata fields available to templates (e.g., `.app.status`, `.repo.GetCommitMetadata`).
*   **Annotation Namespace Updates**: Changing the prefix `notifications.argoproj.io/subscribe`.
*   **Upstream API Version Shifts**: Updating integrations to match breaking changes in external APIs (e.g., GitHub Deployment API v3 to v4).
*   **Retry and Timeout Logic**: Introducing new global or service-specific parameters for request handling (e.g., `retryMax`, `timeout`).
*   **Authentication Method Expansion**: Adding support for OAuth2, mTLS, or specific cloud IAM roles for existing services.

## Key Technical Concepts
*   **ConfigMap Configuration**: Defining services within `argocd-notifications-cm` using the `service.<type>.<name>` format.
*   **Secret Referencing**: Injecting sensitive credentials using `$secret-key` syntax referencing `argocd-notifications-secret`.
*   **Subscriptions**: Linking specific triggers to service destinations via Application/Rollout annotations.
*   **Templating**: Using Golang templates to construct payloads (JSON, Markdown, or HTML) for different notification providers.
*   **Triggers**: Conditional logic (`when` expressions) that determine when a notification should be dispatched.
*   **Payload Customization**: Platform-specific UI features like Slack "Blocks," Google Chat "CardsV2," and Teams "Facts."
*   **Threading/Grouping**: Concepts for aggregating multiple notifications into a single conversation thread (Slack `groupingKey`, Google Chat `threadKey`).
*   **Webhook Methods**: HTTP verbs (POST, PUT, PATCH, GET) used for generic integrations.

## Related Components
*   **Argo CD Application Controller**: The source of application state and events.
*   **Argo Rollouts**: Used in conjunction with specific notification integrations (like SQS and PagerDuty) for deployment lifecycle events.
*   **Argo Notifications Engine**: The underlying library responsible for parsing templates and executing service calls.
*   **Kubernetes Secrets/ConfigMaps**: The primary storage mechanism for notifications configuration.
*   **External APIs**: Third-party endpoints provided by Slack, GitHub, PagerDuty, etc.