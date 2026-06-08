// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the linq-api-v3 API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options []option.RequestOption
	Chats   ChatService
	// Messages are individual communications within a chat thread.
	//
	// Messages can include text, media attachments, rich link previews, special
	// effects (like confetti or fireworks), and reactions. All messages are associated
	// with a specific chat and sent from a phone number you own.
	//
	// Messages support delivery status tracking, read receipts, and editing
	// capabilities.
	//
	// ## Rich Link Previews
	//
	// Send a URL as a `link` part to deliver it with a rich preview card showing the
	// page's title, description, and image (when available). A `link` part must be the
	// **only** part in the message — it cannot be combined with text or media parts.
	// To send a URL without a preview card, include it in a `text` part instead.
	//
	// **Limitations:**
	//
	// - A `link` part cannot be combined with other parts in the same message.
	// - Maximum URL length: 2,048 characters.
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
	//
	// ## Security & Ownership
	//
	// Every attachment is bound to the partner account that created or received it.
	// The API enforces ownership on every operation that touches an attachment —
	// sending, retrieving, deleting.
	//
	// **What this means for you:**
	//
	//   - An attachment created under your API key can only be referenced by your API
	//     key.
	//   - Submitting another partner's `attachment_id` returns `404 Not Found`. We do
	//     not disclose whether the id exists or belongs to someone else.
	//   - Submitting a CDN URL that resolves to another partner's attachment is rejected
	//     before the send is attempted.
	//   - Ownership enforcement applies uniformly across send, create-chat, voice memo,
	//     retrieve, and delete operations.
	//
	// Every attachment-affecting endpoint requires a valid partner API key.
	// Unauthenticated calls return `401 Unauthorized`.
	//
	// ## Attachment URL Patterns
	//
	// Attachment URLs in API responses and webhook payloads use one of two layouts,
	// depending on the attachment's tier:
	//
	// | Tier                 | URL pattern                                                                            | TTL                                                              |
	// | -------------------- | -------------------------------------------------------------------------------------- | ---------------------------------------------------------------- |
	// | Persistent (default) | `https://cdn.linqapp.com/attachments/partners/{partner_id}/{attachment_id}/{filename}` | Long-lived                                                       |
	// | Ephemeral            | Pre-signed URL pointing at the ephemeral prefix on `cdn.linqapp.com`                   | 15 minutes per signed URL — re-fetch via the API for a fresh URL |
	//
	// Inbound media you receive over webhooks uses the same layout your outbound sends
	// produce, so the URL you store and the URL you build look identical — no special
	// casing in your client.
	//
	// ## Ephemeral Attachments (Privacy Tier)
	//
	// For regulated or sensitive content, opt in to the **ephemeral attachments** tier
	// by contacting your Linq support contact. You can request it at two scopes:
	//
	// | Scope                | Effect                                                                                                                     |
	// | -------------------- | -------------------------------------------------------------------------------------------------------------------------- |
	// | **Partner-wide**     | Every outbound and inbound attachment on every phone number under your account is routed through the ephemeral tier.       |
	// | **Per phone number** | Only the specified phone numbers route their attachments through the ephemeral tier. The rest stay on the persistent tier. |
	//
	// **Behavioral differences vs the persistent default:**
	//
	// | Aspect                  | Persistent                           | Ephemeral                                                                                                              |
	// | ----------------------- | ------------------------------------ | ---------------------------------------------------------------------------------------------------------------------- |
	// | Download URL form       | Long-lived CDN URL                   | Pre-signed URL with short TTL                                                                                          |
	// | Retention floor         | Indefinite (until you call `DELETE`) | **Hard backstop: 1 day** — even without an explicit `DELETE`, the platform removes the underlying bytes after 24 hours |
	// | URL re-fetch            | Not required                         | Fetch via `GET /v3/attachments/{attachmentId}` for a fresh signed URL after TTL expiry                                 |
	// | Cross-partner isolation | Enforced                             | Enforced                                                                                                               |
	//
	// **When to choose ephemeral:**
	//
	//   - Your downstream system processes the file immediately on receipt and does not
	//     need to re-read it later.
	//   - You have a compliance requirement that the platform must not retain
	//     attachments beyond a short window.
	//   - The content is high-sensitivity (PHI, financial documents, identity
	//     verification) and you do not want it sitting behind a long-lived URL.
	//
	// **Important:** ephemeral applies in _both directions_ — outbound files you
	// upload **and** inbound media received by the phone numbers in that scope.
	// Download bytes you need to keep promptly, or fetch a fresh signed URL via the
	// API when needed.
	//
	// ## Deleting an Attachment
	//
	// To permanently remove an attachment you own, use:
	//
	// ```http
	// DELETE /v3/attachments/{attachmentId}
	// Authorization: Bearer <your_api_key>
	// ```
	//
	// **What this does:**
	//
	// 1. Verifies the attachment is owned by your account. Returns `404` otherwise.
	// 2. Removes the underlying file from Linq storage.
	// 3. Records an audit entry (timestamp, partner, attachment id).
	//
	// **Response codes:**
	//
	// | Status                      | Meaning                                                          |
	// | --------------------------- | ---------------------------------------------------------------- |
	// | `204 No Content`            | Deletion succeeded. The attachment is removed from Linq storage. |
	// | `400 Bad Request`           | `attachmentId` is not a valid UUID.                              |
	// | `401 Unauthorized`          | Missing or invalid API key.                                      |
	// | `404 Not Found`             | Attachment does not exist or is not owned by your account.       |
	// | `500 Internal Server Error` | Transient infrastructure issue — safe to retry.                  |
	//
	// **Effect on message history:**
	//
	//   - Messages that referenced the deleted attachment remain visible.
	//   - The message part that pointed at the attachment is preserved with no
	//     attachment reference.
	//   - Webhook payloads previously delivered to you retain the original URL string,
	//     but downloads from that URL return `404` going forward.
	//
	// Deletion is **irreversible**. Once `204` is returned, the bytes are gone — there
	// is no undelete.
	//
	// ## Inbound Media Flow
	//
	// When one of your phone numbers receives a message with media (image, video,
	// audio, document), the platform:
	//
	//  1. Stores the file under your partner account.
	//  2. Records metadata linked to the inbound message.
	//  3. Delivers a webhook whose `parts[]` array includes a `media` part with a `url`
	//     pointing at `cdn.linqapp.com`.
	//  4. If the receiving phone is opted in to ephemeral, the `url` is a short-TTL
	//     signed URL.
	//
	// You can acknowledge the webhook without fetching the file inline, and lazy-load
	// via `GET /v3/attachments/{attachmentId}` later. For ephemeral attachments,
	// retrieving via the API always returns a freshly-signed URL.
	//
	// ## Data Lifecycle Summary
	//
	// | Data                                                | Persistent tier                        | Ephemeral tier                                            |
	// | --------------------------------------------------- | -------------------------------------- | --------------------------------------------------------- |
	// | Attachment bytes                                    | Retained until you `DELETE`            | **Auto-removed after 1 day**, also removable via `DELETE` |
	// | Attachment metadata (id, filename, mime type, size) | Retained until you `DELETE`            | Removed alongside the bytes                               |
	// | Message body & parts                                | Retained per message-retention policy  | Retained per message-retention policy                     |
	// | Audit log of deletions                              | Retained per platform retention policy | Retained per platform retention policy                    |
	//
	// **In transit:** TLS 1.2+ everywhere. **At rest:** AES-256 (server-side
	// encryption).
	//
	// ## Compliance Checklist
	//
	// If you're integrating Linq under a security or privacy review, here is the short
	// list:
	//
	//   - Allowlist exactly one outbound domain: `cdn.linqapp.com`.
	//   - Decide whether you need ephemeral attachments (high-sensitivity content) —
	//     request enablement through your Linq support contact.
	//   - Implement `DELETE /v3/attachments/{attachmentId}` calls in your deletion
	//     workflow.
	//   - Persist any attachments your application needs long-term — Linq is the
	//     authoritative source until you delete, but the ephemeral tier auto-purges
	//     after 1 day.
	//   - For audit: every deletion is logged on Linq's side. Surface a confirmation in
	//     your application UI based on the `204` response.
	//   - For end-user "right to delete" requests: enumerate attachment ids and `DELETE`
	//     each. The platform does not provide a partner-wide wipe endpoint — deletion is
	//     per-attachment by design.
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
	WebhookEvents WebhookEventService
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
	WebhookSubscriptions WebhookSubscriptionService
	// Check whether a recipient address supports iMessage or RCS before sending a
	// message.
	Capability CapabilityService
	Webhooks   WebhookService
	// Contact Card lets you set and share your contact information (name and profile
	// photo) with chat participants via iMessage Name and Photo Sharing.
	//
	// Use `POST /v3/contact_card` to create or update a card for a phone number. Use
	// `PATCH /v3/contact_card` to update an existing active card. Use
	// `GET /v3/contact_card` to retrieve the active card(s) for your partner account.
	//
	// **Sharing behavior:** Sharing may not take effect in every chat due to
	// limitations outside our control. We recommend calling the share endpoint once
	// per day, after the first outbound activity.
	ContactCard ContactCardService
}

// DefaultClientOptions read from the environment (LINQ_API_V3_API_KEY,
// LINQ_API_V3_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithHTTPClient(defaultHTTPClient()), option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("LINQ_API_V3_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("LINQ_API_V3_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("LINQ_API_V3_CUSTOM_HEADERS"); ok {
		for _, line := range strings.Split(o, "\n") {
			colon := strings.Index(line, ":")
			if colon >= 0 {
				defaults = append(defaults, option.WithHeader(strings.TrimSpace(line[:colon]), strings.TrimSpace(line[colon+1:])))
			}
		}
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
	r.ContactCard = NewContactCardService(opts...)

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
