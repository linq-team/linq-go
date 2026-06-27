// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/apiquery"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/pagination"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
	"github.com/linq-team/linq-go/shared"
)

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
//
// ## Ephemeral Messages (Privacy Tier)
//
// For regulated or sensitive conversations, opt in to the **ephemeral messages**
// tier by contacting your Linq support contact. When enabled, every message on the
// covered phone numbers is automatically given a fixed **24-hour retention
// window** — after that window the platform permanently deletes the message from
// Linq storage. There is no per-message flag; ephemerality is applied
// automatically based on your configuration.
//
// You can request it at two scopes:
//
// | Scope                | Effect                                                                                                                    |
// | -------------------- | ------------------------------------------------------------------------------------------------------------------------- |
// | **Partner-wide**     | Every outbound and inbound message on every phone number under your account is retained for 24 hours, then deleted.       |
// | **Per phone number** | Only the specified phone numbers have their messages auto-deleted. The rest follow the standard message-retention policy. |
//
// **Behavioral differences vs the standard default:**
//
// | Aspect                  | Standard                                           | Ephemeral                                                                                                                                   |
// | ----------------------- | -------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
// | Retention               | Retained per the standard message-retention policy | **Hard backstop: 24 hours** from when the message is created                                                                                |
// | After expiry            | Message stays retrievable                          | Message is permanently deleted — `GET /v3/messages/{messageId}` returns `404` and it no longer appears in `GET /v3/chats/{chatId}/messages` |
// | Content on expiry       | N/A                                                | Text, formatting, and attachment references are scrubbed; the message is gone, not blanked out                                              |
// | Cross-partner isolation | Enforced                                           | Enforced                                                                                                                                    |
//
// **How the 24-hour window works:**
//
//   - The window is fixed at **24 hours from message creation** (`created_at`) and
//     cannot be configured per message.
//   - It mirrors the ephemeral _attachments_ 1-day backstop, so a message and any
//     media it carries expire together.
//   - Expiry is delivery-independent — the clock starts when the message is created,
//     not when it is delivered or read.
//
// **What you observe:**
//
//   - **No expiry timestamp is exposed.** API responses and webhook payloads do not
//     include the deletion time. If you need it, compute `created_at + 24h`
//     yourself.
//   - **No deletion webhook is sent.** There is no `message.deleted` event — a
//     message simply stops being retrievable once its window passes.
//   - **Delivery is unaffected.** Ephemeral messages send, deliver, and fire the
//     usual `message.sent` / `message.received` and status webhooks exactly like
//     standard messages. Only retention changes.
//
// **When to choose ephemeral:**
//
//   - You have a compliance requirement that the platform must not retain message
//     content beyond a short window.
//   - The conversation is high-sensitivity (PHI, financial, identity verification)
//     and you do not want it sitting in storage long-term.
//   - Your application is the system of record — you capture what you need from the
//     delivery webhook in real time and do not rely on reading message history back
//     from Linq later.
//
// **Important:** ephemeral applies in _both directions_ — messages you send
// **and** messages received by the phone numbers in that scope. Because Linq can
// no longer return the message after 24 hours, persist anything you need to keep
// from the webhook payload at the time it is delivered.
//
// ChatMessageService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatMessageService] method instead.
type ChatMessageService struct {
	Options []option.RequestOption
}

// NewChatMessageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatMessageService(opts ...option.RequestOption) (r ChatMessageService) {
	r = ChatMessageService{}
	r.Options = opts
	return
}

// Retrieve messages from a specific chat with pagination support.
func (r *ChatMessageService) List(ctx context.Context, chatID string, query ChatMessageListParams, opts ...option.RequestOption) (res *pagination.ListMessagesPagination[Message], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/chats/%s/messages", chatID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve messages from a specific chat with pagination support.
func (r *ChatMessageService) ListAutoPaging(ctx context.Context, chatID string, query ChatMessageListParams, opts ...option.RequestOption) *pagination.ListMessagesPaginationAutoPager[Message] {
	return pagination.NewListMessagesPaginationAutoPager(r.List(ctx, chatID, query, opts...))
}

// Send a message to an existing chat. Use this endpoint when you already have a
// chat ID and want to send additional messages to it.
//
// ## Message Effects
//
// You can add iMessage effects to make your messages more expressive. Effects are
// optional and can be either screen effects (full-screen animations) or bubble
// effects (message bubble animations).
//
// **Screen Effects:** `confetti`, `fireworks`, `lasers`, `sparkles`,
// `celebration`, `hearts`, `love`, `balloons`, `happy_birthday`, `echo`,
// `spotlight`
//
// **Bubble Effects:** `slam`, `loud`, `gentle`, `invisible`
//
// Only one effect type can be applied per message.
//
// ## Inline Text Decorations (iMessage only)
//
// Use the `text_decorations` array on a text part to apply styling and animations
// to character ranges.
//
// Each decoration specifies a `range: [start, end)` and exactly one of `style` or
// `animation`.
//
// **Styles:** `bold`, `italic`, `strikethrough`, `underline` **Animations:**
// `big`, `small`, `shake`, `nod`, `explode`, `ripple`, `bloom`, `jitter`
//
// ```json
//
//	{
//	  "type": "text",
//	  "value": "Hello world",
//	  "text_decorations": [
//	    { "range": [0, 5], "style": "bold" },
//	    { "range": [6, 11], "animation": "shake" }
//	  ]
//	}
//
// ```
//
// **Note:** Style ranges (bold, italic, etc.) may overlap, but animation ranges
// must not overlap with other animations or styles. Text decorations only render
// for iMessage recipients. For SMS/RCS, text decorations are not applied.
func (r *ChatMessageService) Send(ctx context.Context, chatID string, body ChatMessageSendParams, opts ...option.RequestOption) (res *ChatMessageSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/chats/%s/messages", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A message that was sent (used in CreateChat and SendMessage responses)
type SentMessage struct {
	// Message identifier (UUID)
	ID string `json:"id" api:"required" format:"uuid"`
	// When the message was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Current delivery status of a message
	//
	// Any of "pending", "queued", "sent", "delivered", "received", "read", "failed".
	DeliveryStatus SentMessageDeliveryStatus `json:"delivery_status" api:"required"`
	// DEPRECATED: Use `delivery_status == "read"` instead. Whether the message has
	// been read.
	//
	// Deprecated: deprecated
	IsRead bool `json:"is_read" api:"required"`
	// Message parts in order (text, media, and link)
	Parts []SentMessagePartUnion `json:"parts" api:"required"`
	// When the message was actually sent (null if still queued)
	SentAt time.Time `json:"sent_at" api:"required" format:"date-time"`
	// When the message was delivered
	DeliveredAt time.Time `json:"delivered_at" api:"nullable" format:"date-time"`
	// iMessage effect applied to a message (screen or bubble effect)
	Effect MessageEffect `json:"effect" api:"nullable"`
	// The sender of this message as a full handle object
	FromHandle shared.ChatHandle `json:"from_handle" api:"nullable"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	PreferredService shared.ServiceType `json:"preferred_service" api:"nullable"`
	// Indicates this message is a threaded reply to another message
	ReplyTo ReplyTo `json:"reply_to" api:"nullable"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		DeliveryStatus   respjson.Field
		IsRead           respjson.Field
		Parts            respjson.Field
		SentAt           respjson.Field
		DeliveredAt      respjson.Field
		Effect           respjson.Field
		FromHandle       respjson.Field
		PreferredService respjson.Field
		ReplyTo          respjson.Field
		Service          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SentMessage) RawJSON() string { return r.JSON.raw }
func (r *SentMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current delivery status of a message
type SentMessageDeliveryStatus string

const (
	SentMessageDeliveryStatusPending   SentMessageDeliveryStatus = "pending"
	SentMessageDeliveryStatusQueued    SentMessageDeliveryStatus = "queued"
	SentMessageDeliveryStatusSent      SentMessageDeliveryStatus = "sent"
	SentMessageDeliveryStatusDelivered SentMessageDeliveryStatus = "delivered"
	SentMessageDeliveryStatusReceived  SentMessageDeliveryStatus = "received"
	SentMessageDeliveryStatusRead      SentMessageDeliveryStatus = "read"
	SentMessageDeliveryStatusFailed    SentMessageDeliveryStatus = "failed"
)

// SentMessagePartUnion contains all possible properties and values from
// [shared.TextPartResponse], [shared.MediaPartResponse],
// [shared.LinkPartResponse], [SentMessagePartIMessageAppPartResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SentMessagePartUnion struct {
	Reactions []shared.Reaction `json:"reactions"`
	Type      string            `json:"type"`
	Value     string            `json:"value"`
	// This field is from variant [shared.TextPartResponse].
	TextDecorations []shared.TextDecoration `json:"text_decorations"`
	// This field is from variant [shared.MediaPartResponse].
	ID string `json:"id"`
	// This field is from variant [shared.MediaPartResponse].
	Filename string `json:"filename"`
	// This field is from variant [shared.MediaPartResponse].
	MimeType string `json:"mime_type"`
	// This field is from variant [shared.MediaPartResponse].
	SizeBytes int64  `json:"size_bytes"`
	URL       string `json:"url"`
	// This field is from variant [SentMessagePartIMessageAppPartResponse].
	App SentMessagePartIMessageAppPartResponseApp `json:"app"`
	// This field is from variant [SentMessagePartIMessageAppPartResponse].
	Layout SentMessagePartIMessageAppPartResponseLayout `json:"layout"`
	// This field is from variant [SentMessagePartIMessageAppPartResponse].
	FallbackText string `json:"fallback_text"`
	JSON         struct {
		Reactions       respjson.Field
		Type            respjson.Field
		Value           respjson.Field
		TextDecorations respjson.Field
		ID              respjson.Field
		Filename        respjson.Field
		MimeType        respjson.Field
		SizeBytes       respjson.Field
		URL             respjson.Field
		App             respjson.Field
		Layout          respjson.Field
		FallbackText    respjson.Field
		raw             string
	} `json:"-"`
}

func (u SentMessagePartUnion) AsTextPartResponse() (v shared.TextPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SentMessagePartUnion) AsMediaPartResponse() (v shared.MediaPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SentMessagePartUnion) AsLinkPartResponse() (v shared.LinkPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SentMessagePartUnion) AsSentMessagePartIMessageAppPartResponse() (v SentMessagePartIMessageAppPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SentMessagePartUnion) RawJSON() string { return u.JSON.raw }

func (r *SentMessagePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An iMessage app card part.
type SentMessagePartIMessageAppPartResponse struct {
	// Identifies the iMessage app (Messages app extension) that backs the card.
	App SentMessagePartIMessageAppPartResponseApp `json:"app" api:"required"`
	// Visible layout of the card. At least one of `caption`, `subcaption`,
	// `trailing_caption`, or `trailing_subcaption` must be set, otherwise the card
	// renders as an empty bubble. Any image on the card is drawn by the recipient's
	// installed app extension; it cannot be supplied here.
	Layout SentMessagePartIMessageAppPartResponseLayout `json:"layout" api:"required"`
	// Reactions on this message part
	Reactions []shared.Reaction `json:"reactions" api:"required"`
	// Indicates this is an iMessage app card part.
	//
	// Any of "imessage_app".
	Type string `json:"type" api:"required"`
	// The URL delivered to the iMessage app on tap.
	URL string `json:"url" api:"required" format:"uri"`
	// Fallback text for surfaces that cannot render the card.
	FallbackText string `json:"fallback_text" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		App          respjson.Field
		Layout       respjson.Field
		Reactions    respjson.Field
		Type         respjson.Field
		URL          respjson.Field
		FallbackText respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SentMessagePartIMessageAppPartResponse) RawJSON() string { return r.JSON.raw }
func (r *SentMessagePartIMessageAppPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the iMessage app (Messages app extension) that backs the card.
type SentMessagePartIMessageAppPartResponseApp struct {
	// Bundle identifier of the Messages app extension. Must not contain `:`.
	BundleID string `json:"bundle_id" api:"required"`
	// Display name of the app, shown by Messages' fallback UI.
	Name string `json:"name" api:"required"`
	// The app's 10-character uppercase alphanumeric team identifier.
	TeamID string `json:"team_id" api:"required"`
	// The owning app's App Store id (optional). When set, recipients without the
	// iMessage app installed see a "Get the app" affordance.
	AppStoreID int64 `json:"app_store_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BundleID    respjson.Field
		Name        respjson.Field
		TeamID      respjson.Field
		AppStoreID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SentMessagePartIMessageAppPartResponseApp) RawJSON() string { return r.JSON.raw }
func (r *SentMessagePartIMessageAppPartResponseApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visible layout of the card. At least one of `caption`, `subcaption`,
// `trailing_caption`, or `trailing_subcaption` must be set, otherwise the card
// renders as an empty bubble. Any image on the card is drawn by the recipient's
// installed app extension; it cannot be supplied here.
type SentMessagePartIMessageAppPartResponseLayout struct {
	// Primary label, top-left and bold.
	Caption string `json:"caption"`
	// Secondary label, below `caption` on the left.
	Subcaption string `json:"subcaption"`
	// Label shown top-right.
	TrailingCaption string `json:"trailing_caption"`
	// Label shown below `trailing_caption`, on the right.
	TrailingSubcaption string `json:"trailing_subcaption"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption            respjson.Field
		Subcaption         respjson.Field
		TrailingCaption    respjson.Field
		TrailingSubcaption respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SentMessagePartIMessageAppPartResponseLayout) RawJSON() string { return r.JSON.raw }
func (r *SentMessagePartIMessageAppPartResponseLayout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for sending a message to a chat
type ChatMessageSendResponse struct {
	// Unique identifier of the chat this message was sent to
	ChatID string `json:"chat_id" api:"required" format:"uuid"`
	// A message that was sent (used in CreateChat and SendMessage responses)
	Message SentMessage `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageSendResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageListParams struct {
	// Pagination cursor from previous next_cursor response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of messages to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChatMessageListParams]'s query parameters as `url.Values`.
func (r ChatMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChatMessageSendParams struct {
	// Message content container. Groups all message-related fields together,
	// separating the "what" (message content) from the "where" (routing fields like
	// from/to).
	Message MessageContentParam `json:"message,omitzero" api:"required"`
	paramObj
}

func (r ChatMessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
