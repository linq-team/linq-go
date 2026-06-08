// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
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
// WebhookSubscriptionService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookSubscriptionService] method instead.
type WebhookSubscriptionService struct {
	Options []option.RequestOption
}

// NewWebhookSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWebhookSubscriptionService(opts ...option.RequestOption) (r WebhookSubscriptionService) {
	r = WebhookSubscriptionService{}
	r.Options = opts
	return
}

// Create a new webhook subscription to receive events at a target URL. Upon
// creation, a signing secret is generated for verifying webhook authenticity.
// **Store this secret securely — it cannot be retrieved later.**
//
// **Phone Number Filtering:**
//
//   - Optionally specify `phone_numbers` to only receive events for specific lines
//   - If omitted, events from all phone numbers are delivered (default behavior)
//   - Use multiple subscriptions with different `phone_numbers` to route different
//     lines to different endpoints
//   - Each `target_url` can only be used once per account. To route different lines
//     to different destinations, use a unique URL per subscription (e.g., append a
//     query parameter: `https://example.com/webhook?line=1`)
//
// **Webhook Delivery:**
//
//   - Events are sent via HTTP POST to the target URL
//   - Each request includes
//     [Standard Webhooks](https://github.com/standard-webhooks/standard-webhooks)
//     headers (`webhook-id`, `webhook-timestamp`, `webhook-signature`) for signature
//     verification
//   - Legacy `X-Webhook-*` headers are also sent for backwards compatibility
//     (deprecated)
//   - See
//     [Verifying Webhook Signatures](https://docs.linqapp.com/guides/webhooks#verifying-webhook-signatures)
//     for verification details
//   - Failed deliveries (5xx, 429, network errors) are retried up to 10 times over
//     ~25 minutes with exponential backoff
//   - Client errors (4xx except 429) are not retried
func (r *WebhookSubscriptionService) New(ctx context.Context, body WebhookSubscriptionNewParams, opts ...option.RequestOption) (res *WebhookSubscriptionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/webhook-subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve details for a specific webhook subscription including its target URL,
// subscribed events, and current status.
func (r *WebhookSubscriptionService) Get(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *WebhookSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/webhook-subscriptions/%s", url.PathEscape(subscriptionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing webhook subscription. You can modify the target URL,
// subscribed events, or activate/deactivate the subscription.
//
// **Note:** The signing secret cannot be changed via this endpoint.
func (r *WebhookSubscriptionService) Update(ctx context.Context, subscriptionID string, body WebhookSubscriptionUpdateParams, opts ...option.RequestOption) (res *WebhookSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/webhook-subscriptions/%s", url.PathEscape(subscriptionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Retrieve all webhook subscriptions for the authenticated partner. Returns a list
// of active and inactive subscriptions with their configuration and status.
func (r *WebhookSubscriptionService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookSubscriptionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/webhook-subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a webhook subscription.
func (r *WebhookSubscriptionService) Delete(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return err
	}
	path := fmt.Sprintf("v3/webhook-subscriptions/%s", url.PathEscape(subscriptionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type WebhookSubscription struct {
	// Unique identifier for the webhook subscription
	ID string `json:"id" api:"required"`
	// When the subscription was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this subscription is currently active
	IsActive bool `json:"is_active" api:"required"`
	// List of event types this subscription receives
	SubscribedEvents []WebhookEventType `json:"subscribed_events" api:"required"`
	// URL where webhook events will be sent
	TargetURL string `json:"target_url" api:"required" format:"uri"`
	// When the subscription was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Phone numbers this subscription filters for. If null or empty, events from all
	// phone numbers are delivered.
	PhoneNumbers []string `json:"phone_numbers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		IsActive         respjson.Field
		SubscribedEvents respjson.Field
		TargetURL        respjson.Field
		UpdatedAt        respjson.Field
		PhoneNumbers     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookSubscription) RawJSON() string { return r.JSON.raw }
func (r *WebhookSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response returned when creating a webhook subscription. Includes the signing
// secret which is only shown once.
type WebhookSubscriptionNewResponse struct {
	// Unique identifier for the webhook subscription
	ID string `json:"id" api:"required"`
	// When the subscription was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this subscription is currently active
	IsActive bool `json:"is_active" api:"required"`
	// Secret for verifying webhook signatures. Store this securely - it cannot be
	// retrieved again.
	SigningSecret string `json:"signing_secret" api:"required"`
	// List of event types this subscription receives
	SubscribedEvents []WebhookEventType `json:"subscribed_events" api:"required"`
	// URL where webhook events will be sent
	TargetURL string `json:"target_url" api:"required" format:"uri"`
	// When the subscription was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Phone numbers this subscription filters for. If null or empty, events from all
	// phone numbers are delivered.
	PhoneNumbers []string `json:"phone_numbers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		IsActive         respjson.Field
		SigningSecret    respjson.Field
		SubscribedEvents respjson.Field
		TargetURL        respjson.Field
		UpdatedAt        respjson.Field
		PhoneNumbers     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookSubscriptionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookSubscriptionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookSubscriptionListResponse struct {
	// List of webhook subscriptions
	Subscriptions []WebhookSubscription `json:"subscriptions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Subscriptions respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookSubscriptionListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookSubscriptionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookSubscriptionNewParams struct {
	// List of event types to subscribe to
	SubscribedEvents []WebhookEventType `json:"subscribed_events,omitzero" api:"required"`
	// URL where webhook events will be sent. Must be HTTPS.
	TargetURL string `json:"target_url" api:"required" format:"uri"`
	// Optional list of phone numbers to filter events for. Only events originating
	// from these phone numbers will be delivered to this subscription. If omitted or
	// empty, events from all phone numbers are delivered. Phone numbers must be in
	// E.164 format.
	PhoneNumbers []string `json:"phone_numbers,omitzero"`
	paramObj
}

func (r WebhookSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookSubscriptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookSubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookSubscriptionUpdateParams struct {
	// Activate or deactivate the subscription
	IsActive param.Opt[bool] `json:"is_active,omitzero"`
	// New target URL for webhook events
	TargetURL param.Opt[string] `json:"target_url,omitzero" format:"uri"`
	// Updated list of phone numbers to filter events for. Set to a non-empty array to
	// filter events to specific phone numbers. Set to an empty array or null to remove
	// the filter and receive events from all phone numbers. Phone numbers must be in
	// E.164 format.
	PhoneNumbers []string `json:"phone_numbers,omitzero"`
	// Updated list of event types to subscribe to
	SubscribedEvents []WebhookEventType `json:"subscribed_events,omitzero"`
	paramObj
}

func (r WebhookSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookSubscriptionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookSubscriptionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
