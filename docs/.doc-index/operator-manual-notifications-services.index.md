# OPERATOR-MANUAL/NOTIFICATIONS/SERVICES Documentation Index

## Overview
This documentation area provides detailed configuration guides for the various notification service integrations supported by Argo CD Notifications. It explains how to define service providers in the `argocd-notifications-cm` ConfigMap, manage sensitive credentials in Kubernetes Secrets, and customize notification delivery using templates and application annotations.

## Files Summary
*   **operator-manual/notifications/services/overview.md**: Provides the foundational concepts for notification services, including global configuration syntax, sensitive data handling using secret references, and the use of custom names for multiple instances of the same service type.
*   **operator-manual/notifications/services/alertmanager.md**: Details the integration with Prometheus Alertmanager, covering target addresses, authentication (Basic/Bearer), and the use of custom labels and annotations in alert payloads.
*   **operator-manual/notifications/services/awssqs.md**: Explains how to send notification messages to AWS SQS queues, including credential management via IAM/OIDC or Secrets and support for FIFO queues.
*   **operator-manual/notifications/services/email.md**: Covers the configuration of SMTP-based email notifications, including server settings, authentication, and HTML content support.
*   **operator-manual/notifications/services/github.md**: Describes how to use GitHub Apps to update commit statuses, create deployment markers, and post pull request comments based on Argo CD events.
*   **operator-manual/notifications/services/googlechat.md**: Instructions for sending text or card-based messages to Google Chat spaces via incoming webhooks, including support for chat threading.
*   **operator-manual/notifications/services/grafana.md**: Explains how to create annotations on Grafana dashboards to visualize deployment events alongside metrics.
*   **operator-manual/notifications/services/mattermost.md**: Details integration with Mattermost via bot accounts, utilizing Slack-compatible attachment formats for rich messaging.
*   **operator-manual/notifications/services/newrelic.md**: Guides users on creating NewRelic deployment markers, including the mapping of commit metadata to NewRelic attributes.
*   **operator-manual/notifications/services/opsgenie.md**: Covers the creation and routing of alerts in Opsgenie, including detailed metadata mapping for priorities, aliases, and responders.
*   **operator-manual/notifications/services/pagerduty.md**: Explains the integration with PagerDuty for incident creation using the service ID and integration tokens.
*   **operator-manual/notifications/services/pagerduty_v2.md**: Focuses on the PagerDuty Events API v2, allowing for more granular event telemetry including severity levels and component groupings.
*   **operator-manual/notifications/services/pushover.md**: A brief guide on sending mobile push notifications through the Pushover API.
*   **operator-manual/notifications/services/rocketchat.md**: Explains how to configure a bot user to post messages and Slack-compatible attachments to Rocket.Chat channels.
*   **operator-manual/notifications/services/slack.md**: A comprehensive guide for Slack integration using OAuth apps, covering blocks, attachments, message threading (grouping keys), and delivery policies.
*   **operator-manual/notifications/services/teams.md**: Details how to send notifications to Microsoft Teams via webhooks using "Facts" and "Potential Actions" for interactive messages.
*   **operator-manual/notifications/services/telegram.md**: Instructions for setting up a Telegram bot to deliver notifications to public or private channels.
*   **operator-manual/notifications/services/webex.md**: Covers the configuration for sending notifications to Cisco Webex Teams rooms using bot tokens.
*   **operator-manual/notifications/services/webhook.md**: Explains how to configure generic HTTP webhooks for custom integrations, including retry logic, custom headers, and templatized request bodies.

## Code Changes That Would Require Documentation Updates
*   **Addition of New Service Providers**: Any new integration added to the notifications engine must have a corresponding file created here and be added to the `overview.md` list.
*   **Changes to ConfigMap Schema**: Modifying the expected YAML keys for a service in `argocd-notifications-cm` (e.g., changing `apiURL` to `serverUrl`).
*   **Template Field Additions**: Introducing new specialized fields within the `template` definition for a specific service (e.g., adding `priorityID` to PagerDuty or `blocks` to Slack).
*   **Authentication Method Updates**: Support for new auth types, such as adding OIDC/Workload Identity support to AWS SQS or GitHub.
*   **Logic in Delivery Behavior**: Changes to how messages are grouped (e.g., `groupingKey`), retried (e.g., `retryMax` in webhooks), or updated (e.g., `deliveryPolicy` in Slack).
*   **Deprecation of Third-Party API Features**: If a service provider changes their API (like Google Chat moving from `cards` to `cardsV2`), the documentation must reflect the recommended usage.
*   **Changes to Subscription Annotations**: Altering the prefix or format of the `notifications.argoproj.io/subscribe` annotations.

## Key Technical Concepts
*   **Service Definition**: The `service.<type>.<name>` syntax used in the ConfigMap to register a notification provider.
*   **Sensitive Data Referencing**: The `$key` syntax used to pull values from `argocd-notifications-secret`.
*   **Subscriptions**: The use of annotations on Argo CD `Application` or `Rollout` resources to map triggers to specific service destinations.
*   **Rich Messaging Elements**: Service-specific UI components like Slack `blocks`, MS Teams `facts`, Google Chat `cardsV2`, and GitHub `checkRun`.
*   **Message Aggregation**: Concepts like `groupingKey` (Slack), `threadKey` (Google Chat), or `messageGroupId` (SQS FIFO) used to organize notifications.
*   **HTTP Configuration**: Advanced webhook settings including `insecureSkipVerify`, `retryWaitMax`, and custom `headers`.
*   **Service Compatibility**: The ability for services like Mattermost and Rocket.Chat to consume Slack-formatted JSON attachments.

## Related Components
*   **argocd-notifications-controller**: The primary controller that watches for triggers and executes these service integrations.
*   **argocd-notifications-cm**: The central ConfigMap for all service and template definitions.
*   **argocd-notifications-secret**: The default location for storing API tokens and passwords.
*   **Notifications Engine**: The underlying library (argoproj/notifications-engine) that provides the core logic for these integrations.
*   **Argo CD Application Controller**: The source of the application state changes that trigger notifications.
*   **Argo Rollouts**: Mentioned as a resource type capable of producing notification triggers (specifically in PagerDuty and SQS docs).