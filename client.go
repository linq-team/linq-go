// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the linq-api-v3 API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options []option.RequestOption
	Chats   ChatService
	// Messages are individual text or multimedia communications within a chat thread.
	//
	// Messages can include text, attachments, special effects (like confetti or
	// fireworks), and reactions. All messages are associated with a specific chat and
	// sent from a phone number you own.
	//
	// Messages support delivery status tracking, read receipts, and editing
	// capabilities.
	Messages MessageService
	// Send files (images, videos, documents, audio) with messages by providing a URL
	// in a media part. Pre-uploading via `POST /v3/attachments` is **optional** and
	// only needed for specific optimization scenarios.
	//
	// ## Sending Media via URL (up to 10MB)
	//
	// Provide a publicly accessible HTTPS URL with a
	// [supported media type](#supported-file-types) in the `url` field of a media
	// part.
	//
	// ```json
	//
	//	{
	//	  "parts": [{ "type": "media", "url": "https://your-cdn.com/images/photo.jpg" }]
	//	}
	//
	// ```
	//
	// This works with any URL you already host — no pre-upload step required.
	// **Maximum file size: 10MB.**
	//
	// ## Pre-Upload (required for files over 10MB)
	//
	// Use `POST /v3/attachments` when you want to:
	//
	//   - **Send files larger than 10MB** (up to 100MB) — URL-based downloads are
	//     limited to 10MB
	//   - **Send the same file to many recipients** — upload once, reuse the
	//     `attachment_id` without re-downloading each time
	//   - **Reduce message send latency** — the file is already stored, so sending is
	//     faster
	//
	// **How it works:**
	//
	//  1. `POST /v3/attachments` with file metadata → returns a presigned `upload_url`
	//     (valid for **15 minutes**) and a permanent `attachment_id`
	//  2. PUT the raw file bytes to the `upload_url` with the `required_headers` (no
	//     JSON or multipart — just the binary content)
	//  3. Reference the `attachment_id` in your media part when sending messages (no
	//     expiration)
	//
	// **Key difference:** When you provide an external `url`, we download and process
	// the file on every send. When you use a pre-uploaded `attachment_id`, the file is
	// already stored — so repeated sends skip the download step entirely.
	//
	// ## Domain Allowlisting
	//
	// Attachment URLs in API responses are served from `cdn.linqapp.com`. This
	// includes:
	//
	// - `url` fields in media and voice memo message parts
	// - `download_url` fields in attachment and upload response objects
	//
	// If your application enforces domain allowlists (e.g., for SSRF protection), add:
	//
	// ```
	// cdn.linqapp.com
	// ```
	//
	// ## Supported File Types
	//
	// - **Images:** JPEG, PNG, GIF, HEIC, HEIF, TIFF, BMP
	// - **Videos:** MP4, MOV, M4V
	// - **Audio:** M4A, AAC, MP3, WAV, AIFF, CAF, AMR
	// - **Documents:** PDF, TXT, RTF, CSV, Office formats, ZIP
	// - **Contact & Calendar:** VCF, ICS
	//
	// ## Audio: Attachment vs Voice Memo
	//
	// Audio files sent as media parts appear as **downloadable file attachments** in
	// iMessage. To send audio as an **iMessage voice memo bubble** (with native inline
	// playback UI), use the dedicated `POST /v3/chats/{chatId}/voicememo` endpoint
	// instead.
	//
	// ## File Size Limits
	//
	// - **URL-based (`url` field):** 10MB maximum
	// - **Pre-upload (`attachment_id`):** 100MB maximum
	Attachments AttachmentService
	// Phone Numbers represent the phone numbers assigned to your partner account.
	//
	// Use the list phone numbers endpoint to discover which phone numbers are
	// available for sending messages.
	//
	// When creating chats, listing chats, or sending a voice memo, use one of your
	// assigned phone numbers in the `from` field.
	Phonenumbers PhonenumberService
	// Phone Numbers represent the phone numbers assigned to your partner account.
	//
	// Use the list phone numbers endpoint to discover which phone numbers are
	// available for sending messages.
	//
	// When creating chats, listing chats, or sending a voice memo, use one of your
	// assigned phone numbers in the `from` field.
	PhoneNumbers PhoneNumberService
	// Webhook Subscriptions allow you to receive real-time notifications when events
	// occur on your account.
	//
	// Configure webhook endpoints to receive events such as messages sent/received,
	// delivery status changes, reactions, typing indicators, and more.
	//
	// Failed deliveries (5xx, 429, network errors) are retried up to 10 times over ~2
	// hours with exponential backoff. Each event includes a unique ID for
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
	WebhookEvents WebhookEventService
	// Webhook Subscriptions allow you to receive real-time notifications when events
	// occur on your account.
	//
	// Configure webhook endpoints to receive events such as messages sent/received,
	// delivery status changes, reactions, typing indicators, and more.
	//
	// Failed deliveries (5xx, 429, network errors) are retried up to 10 times over ~2
	// hours with exponential backoff. Each event includes a unique ID for
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
	WebhookSubscriptions WebhookSubscriptionService
	// Check whether a recipient address supports iMessage or RCS before sending a
	// message.
	Capability CapabilityService
	Webhooks   WebhookService
}

// DefaultClientOptions read from the environment (LINQ_API_V3_API_KEY,
// LINQ_API_V3_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("LINQ_API_V3_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("LINQ_API_V3_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (LINQ_API_V3_API_KEY, LINQ_API_V3_BASE_URL). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.Chats = NewChatService(opts...)
	r.Messages = NewMessageService(opts...)
	r.Attachments = NewAttachmentService(opts...)
	r.Phonenumbers = NewPhonenumberService(opts...)
	r.PhoneNumbers = NewPhoneNumberService(opts...)
	r.WebhookEvents = NewWebhookEventService(opts...)
	r.WebhookSubscriptions = NewWebhookSubscriptionService(opts...)
	r.Capability = NewCapabilityService(opts...)
	r.Webhooks = NewWebhookService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
