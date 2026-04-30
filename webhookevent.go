// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/respjson"
	"github.com/linq-team/linq-go/shared/constant"
)

// Webhook Subscriptions allow you to receive real-time notifications when events
// occur on your account.
//
// Configure webhook endpoints to receive events such as messages sent/received,
// delivery status changes, reactions, typing indicators, and more.
//
// Failed deliveries (5xx, 429, network errors) are retried up to 10 times over ~25
// minutes with exponential backoff. Each event includes a unique ID for
// deduplication.
//
// ## Webhook Headers
//
// Each webhook request includes the following headers:
//
// | Header                      | Description                                               |
// | --------------------------- | --------------------------------------------------------- |
// | `X-Webhook-Event`           | The event type (e.g., `message.sent`, `message.received`) |
// | `X-Webhook-Subscription-ID` | Your webhook subscription ID                              |
// | `X-Webhook-Timestamp`       | Unix timestamp (seconds) when the webhook was sent        |
// | `X-Webhook-Signature`       | HMAC-SHA256 signature for verification                    |
//
// ## Verifying Webhook Signatures
//
// All webhooks are signed using HMAC-SHA256. You should always verify the
// signature to ensure the webhook originated from Linq and hasn't been tampered
// with.
//
// **Signature Construction:**
//
// The signature is computed over a concatenation of the timestamp and payload:
//
// ```
// {timestamp}.{payload}
// ```
//
// Where:
//
// - `timestamp` is the value from the `X-Webhook-Timestamp` header
// - `payload` is the raw JSON request body (exact bytes, not re-serialized)
//
// **Verification Steps:**
//
// 1. Extract the `X-Webhook-Timestamp` and `X-Webhook-Signature` headers
// 2. Get the raw request body bytes (do not parse and re-serialize)
// 3. Concatenate: `"{timestamp}.{payload}"`
// 4. Compute HMAC-SHA256 using your signing secret as the key
// 5. Hex-encode the result and compare with `X-Webhook-Signature`
// 6. Use constant-time comparison to prevent timing attacks
//
// **Example (Python):**
//
// ```python
// import hmac
// import hashlib
//
// def verify_webhook(signing_secret, payload, timestamp, signature):
//
//	message = f"{timestamp}.{payload.decode('utf-8')}"
//	expected = hmac.new(
//	    signing_secret.encode('utf-8'),
//	    message.encode('utf-8'),
//	    hashlib.sha256
//	).hexdigest()
//	return hmac.compare_digest(expected, signature)
//
// ```
//
// **Example (Node.js):**
//
// ```javascript
// const crypto = require("crypto");
//
//	function verifyWebhook(signingSecret, payload, timestamp, signature) {
//	  const message = `${timestamp}.${payload}`;
//	  const expected = crypto
//	    .createHmac("sha256", signingSecret)
//	    .update(message)
//	    .digest("hex");
//	  return crypto.timingSafeEqual(Buffer.from(expected), Buffer.from(signature));
//	}
//
// ```
//
// **Security Best Practices:**
//
//   - Reject webhooks with timestamps older than 5 minutes to prevent replay attacks
//   - Always use constant-time comparison for signature verification
//   - Store your signing secret securely (e.g., environment variable, secrets
//     manager)
//   - Return a 2xx status code quickly, then process the webhook asynchronously
//
// WebhookEventService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookEventService] method instead.
type WebhookEventService struct {
	Options []option.RequestOption
}

// NewWebhookEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookEventService(opts ...option.RequestOption) (r WebhookEventService) {
	r = WebhookEventService{}
	r.Options = opts
	return
}

// Returns all available webhook event types that can be subscribed to. Use this
// endpoint to discover valid values for the `subscribed_events` field when
// creating or updating webhook subscriptions.
func (r *WebhookEventService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/webhook-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Valid webhook event types that can be subscribed to.
//
// **Note:** `message.edited` is only delivered to subscriptions using
// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
// subscription will not produce any deliveries.
type WebhookEventType string

const (
	WebhookEventTypeMessageSent                WebhookEventType = "message.sent"
	WebhookEventTypeMessageReceived            WebhookEventType = "message.received"
	WebhookEventTypeMessageRead                WebhookEventType = "message.read"
	WebhookEventTypeMessageDelivered           WebhookEventType = "message.delivered"
	WebhookEventTypeMessageFailed              WebhookEventType = "message.failed"
	WebhookEventTypeMessageEdited              WebhookEventType = "message.edited"
	WebhookEventTypeReactionAdded              WebhookEventType = "reaction.added"
	WebhookEventTypeReactionRemoved            WebhookEventType = "reaction.removed"
	WebhookEventTypeParticipantAdded           WebhookEventType = "participant.added"
	WebhookEventTypeParticipantRemoved         WebhookEventType = "participant.removed"
	WebhookEventTypeChatCreated                WebhookEventType = "chat.created"
	WebhookEventTypeChatGroupNameUpdated       WebhookEventType = "chat.group_name_updated"
	WebhookEventTypeChatGroupIconUpdated       WebhookEventType = "chat.group_icon_updated"
	WebhookEventTypeChatGroupNameUpdateFailed  WebhookEventType = "chat.group_name_update_failed"
	WebhookEventTypeChatGroupIconUpdateFailed  WebhookEventType = "chat.group_icon_update_failed"
	WebhookEventTypeChatTypingIndicatorStarted WebhookEventType = "chat.typing_indicator.started"
	WebhookEventTypeChatTypingIndicatorStopped WebhookEventType = "chat.typing_indicator.stopped"
	WebhookEventTypePhoneNumberStatusUpdated   WebhookEventType = "phone_number.status_updated"
	WebhookEventTypeCallInitiated              WebhookEventType = "call.initiated"
	WebhookEventTypeCallRinging                WebhookEventType = "call.ringing"
	WebhookEventTypeCallAnswered               WebhookEventType = "call.answered"
	WebhookEventTypeCallEnded                  WebhookEventType = "call.ended"
	WebhookEventTypeCallFailed                 WebhookEventType = "call.failed"
	WebhookEventTypeCallDeclined               WebhookEventType = "call.declined"
	WebhookEventTypeCallNoAnswer               WebhookEventType = "call.no_answer"
)

type WebhookEventListResponse struct {
	// URL to the webhook events documentation
	DocURL constant.HTTPSApidocsLinqappComDocumentationWebhookEvents `json:"doc_url" default:"https://apidocs.linqapp.com/documentation/webhook-events"`
	// List of all available webhook event types
	Events []WebhookEventType `json:"events" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocURL      respjson.Field
		Events      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
