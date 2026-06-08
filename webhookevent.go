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
// All webhook requests include two sets of headers. **If you have an existing
// integration using the `X-Webhook-*` headers, nothing changes** — those headers
// are still sent on every delivery and work exactly as before. The new `webhook-*`
// headers follow the
// [Standard Webhooks](https://github.com/standard-webhooks/standard-webhooks)
// specification. You can safely ignore them if your current verification code
// works and you don't want to use this convention.
//
// ### Standard Webhooks Headers (Recommended)
//
// Used by [our SDK](https://github.com/linq-team/linq-node) and any
// [Standard Webhooks library](https://github.com/standard-webhooks/standard-webhooks).
//
// | Header              | Description                                        |
// | ------------------- | -------------------------------------------------- |
// | `webhook-id`        | Unique event identifier (use as idempotency key)   |
// | `webhook-timestamp` | Unix timestamp (seconds) when the webhook was sent |
// | `webhook-signature` | Standard Webhooks signature (`v1,{base64}` format) |
//
// ### Legacy Headers (Deprecated)
//
// Still sent on every delivery for backwards compatibility. Existing verification
// code using these headers continues to work — no changes required.
//
// | Header                      | Description                                        |
// | --------------------------- | -------------------------------------------------- |
// | `X-Webhook-Event`           | _(deprecated)_ Event type (e.g., `message.sent`)   |
// | `X-Webhook-Subscription-ID` | _(deprecated)_ Webhook subscription ID             |
// | `X-Webhook-Timestamp`       | _(deprecated)_ Unix timestamp (seconds)            |
// | `X-Webhook-Signature`       | _(deprecated)_ HMAC-SHA256 signature (hex-encoded) |
//
// ## Signing Secrets
//
// Signing secrets use the Standard Webhooks format: a `whsec_` prefix followed by
// base64-encoded random bytes (e.g.,
// `whsec_MfKQ9r8GKYqrTwjUPD8ILPZIo2LaLaSw7Jxx2Oll+OE=`).
//
// Strip the `whsec_` prefix and base64-decode the remainder to get the raw key
// bytes.
//
// ## Verifying Webhook Signatures
//
// Webhooks are signed following the
// [Standard Webhooks specification](https://github.com/standard-webhooks/standard-webhooks).
// You can use any
// [Standard Webhooks library](https://github.com/standard-webhooks/standard-webhooks)
// to verify signatures, or implement verification manually:
//
// **Signed content:** `{webhook-id}.{webhook-timestamp}.{body}`
//
// **Verification Steps:**
//
//  1. Extract the `webhook-id`, `webhook-timestamp`, and `webhook-signature`
//     headers
//  2. Reject if the timestamp is more than 5 minutes old (replay protection)
//  3. Get the raw request body bytes (do not parse and re-serialize)
//  4. Construct signed content: `"{webhook-id}.{webhook-timestamp}.{body}"`
//  5. Strip the `whsec_` prefix from your secret and base64-decode to get key bytes
//  6. Compute HMAC-SHA256 using the key bytes over the signed content
//  7. Base64-encode the result and compare with the value after `v1,` in
//     `webhook-signature`
//  8. Use constant-time comparison to prevent timing attacks
//
// **Example (Python):**
//
// ```python
// import base64, hmac, hashlib
//
// def verify_webhook(secret, body, headers):
//
//	msg_id = headers['webhook-id']
//	timestamp = headers['webhook-timestamp']
//	signature = headers['webhook-signature']
//
//	secret_str = secret.removeprefix('whsec_')
//	key = base64.b64decode(secret_str)
//
//	signed_content = f"{msg_id}.{timestamp}.{body}"
//	expected = base64.b64encode(
//	    hmac.new(key, signed_content.encode(), hashlib.sha256).digest()
//	).decode()
//
//	for sig in signature.split(' '):
//	    if sig.startswith('v1,') and hmac.compare_digest(expected, sig[3:]):
//	        return True
//	return False
//
// ```
//
// **Example (Node.js):**
//
// ```javascript
// const crypto = require("crypto");
//
//	function verifyWebhook(secret, rawBody, headers) {
//	  const msgId = headers["webhook-id"];
//	  const timestamp = headers["webhook-timestamp"];
//	  const signature = headers["webhook-signature"];
//
//	  const secretStr = secret.startsWith("whsec_") ? secret.slice(6) : secret;
//	  const keyBytes = Buffer.from(secretStr, "base64");
//	  const signedContent = `${msgId}.${timestamp}.${rawBody}`;
//	  const expected = crypto
//	    .createHmac("sha256", keyBytes)
//	    .update(signedContent)
//	    .digest("base64");
//
//	  return signature.split(" ").some((sig) => {
//	    if (!sig.startsWith("v1,")) return false;
//	    try {
//	      return crypto.timingSafeEqual(
//	        Buffer.from(expected, "base64"),
//	        Buffer.from(sig.slice(3), "base64")
//	      );
//	    } catch {
//	      return false;
//	    }
//	  });
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
	WebhookEventTypeLocationSharingStarted     WebhookEventType = "location.sharing.started"
	WebhookEventTypeLocationSharingStopped     WebhookEventType = "location.sharing.stopped"
)

type WebhookEventListResponse struct {
	// URL to the webhook events documentation
	DocURL constant.HTTPSDocsLinqappComGuidesWebhooksEvents `json:"doc_url" default:"https://docs.linqapp.com/guides/webhooks/events"`
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
