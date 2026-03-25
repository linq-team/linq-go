// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"encoding/json"
	"time"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/respjson"
	"github.com/linq-team/linq-go/shared"
	"github.com/linq-team/linq-go/shared/constant"
)

// WebhookService contains methods and other services that help with interacting
// with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

func (r *WebhookService) Events(payload []byte, opts ...option.RequestOption) (*EventsWebhookEventUnion, error) {
	res := &EventsWebhookEventUnion{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Unified payload for message webhooks when using `webhook_version: "2026-02-03"`.
//
// This schema is used for message.sent, message.received, message.delivered, and
// message.read events when the subscription URL includes `?version=2026-02-03`.
//
// Key differences from V1 (2025-01-01):
//
//   - `direction`: "inbound" or "outbound" instead of `is_from_me` boolean
//   - `sender_handle`: Full handle object for the sender
//   - `chat`: Nested object with `id`, `is_group`, and `owner_handle`
//   - Message fields (`id`, `parts`, `effect`, etc.) are at the top level, not
//     nested in `message`
//
// Timestamps indicate the message state:
//
// - `message.sent`: sent_at set, delivered_at=null, read_at=null
// - `message.received`: sent_at set, delivered_at=null, read_at=null
// - `message.delivered`: sent_at set, delivered_at set, read_at=null
// - `message.read`: sent_at set, delivered_at set, read_at set
type MessageEventV2 struct {
	// Message identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Chat information
	Chat MessageEventV2Chat `json:"chat" api:"required"`
	// Message direction - "outbound" if sent by you, "inbound" if received
	//
	// Any of "inbound", "outbound".
	Direction MessageEventV2Direction `json:"direction" api:"required"`
	// Message parts (text and/or media)
	Parts []MessageEventV2PartUnion `json:"parts" api:"required"`
	// The handle that sent this message
	SenderHandle shared.ChatHandle `json:"sender_handle" api:"required"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service" api:"required"`
	// When the message was delivered. Null if not yet delivered.
	DeliveredAt time.Time `json:"delivered_at" api:"nullable" format:"date-time"`
	// iMessage effect applied to a message (screen or bubble animation)
	Effect SchemasMessageEffect `json:"effect" api:"nullable"`
	// Idempotency key for deduplication of outbound messages.
	IdempotencyKey string `json:"idempotency_key" api:"nullable"`
	// Preferred messaging service type. Includes "auto" for default fallback behavior.
	//
	// Any of "iMessage", "SMS", "RCS", "auto".
	PreferredService MessageEventV2PreferredService `json:"preferred_service" api:"nullable"`
	// When the message was read. Null if not yet read.
	ReadAt time.Time `json:"read_at" api:"nullable" format:"date-time"`
	// Reference to the message this is replying to (for threaded replies)
	ReplyTo MessageEventV2ReplyTo `json:"reply_to" api:"nullable"`
	// When the message was sent. Null if not yet sent.
	SentAt time.Time `json:"sent_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Chat             respjson.Field
		Direction        respjson.Field
		Parts            respjson.Field
		SenderHandle     respjson.Field
		Service          respjson.Field
		DeliveredAt      respjson.Field
		Effect           respjson.Field
		IdempotencyKey   respjson.Field
		PreferredService respjson.Field
		ReadAt           respjson.Field
		ReplyTo          respjson.Field
		SentAt           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEventV2) RawJSON() string { return r.JSON.raw }
func (r *MessageEventV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat information
type MessageEventV2Chat struct {
	// Chat identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group" api:"nullable"`
	// Your phone number's handle. Always has is_me=true.
	OwnerHandle shared.ChatHandle `json:"owner_handle" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsGroup     respjson.Field
		OwnerHandle respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEventV2Chat) RawJSON() string { return r.JSON.raw }
func (r *MessageEventV2Chat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message direction - "outbound" if sent by you, "inbound" if received
type MessageEventV2Direction string

const (
	MessageEventV2DirectionInbound  MessageEventV2Direction = "inbound"
	MessageEventV2DirectionOutbound MessageEventV2Direction = "outbound"
)

// MessageEventV2PartUnion contains all possible properties and values from
// [SchemasTextPartResponse], [SchemasMediaPartResponse], [MessageEventV2PartLink].
//
// Use the [MessageEventV2PartUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessageEventV2PartUnion struct {
	// Any of "text", "media", "link".
	Type  string `json:"type"`
	Value string `json:"value"`
	// This field is from variant [SchemasTextPartResponse].
	TextDecorations []shared.TextDecoration `json:"text_decorations"`
	// This field is from variant [SchemasMediaPartResponse].
	ID string `json:"id"`
	// This field is from variant [SchemasMediaPartResponse].
	Filename string `json:"filename"`
	// This field is from variant [SchemasMediaPartResponse].
	MimeType string `json:"mime_type"`
	// This field is from variant [SchemasMediaPartResponse].
	SizeBytes int64 `json:"size_bytes"`
	// This field is from variant [SchemasMediaPartResponse].
	URL  string `json:"url"`
	JSON struct {
		Type            respjson.Field
		Value           respjson.Field
		TextDecorations respjson.Field
		ID              respjson.Field
		Filename        respjson.Field
		MimeType        respjson.Field
		SizeBytes       respjson.Field
		URL             respjson.Field
		raw             string
	} `json:"-"`
}

// anyMessageEventV2Part is implemented by each variant of
// [MessageEventV2PartUnion] to add type safety for the return type of
// [MessageEventV2PartUnion.AsAny]
type anyMessageEventV2Part interface {
	implMessageEventV2PartUnion()
}

func (SchemasTextPartResponse) implMessageEventV2PartUnion()  {}
func (SchemasMediaPartResponse) implMessageEventV2PartUnion() {}
func (MessageEventV2PartLink) implMessageEventV2PartUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MessageEventV2PartUnion.AsAny().(type) {
//	case linqgo.SchemasTextPartResponse:
//	case linqgo.SchemasMediaPartResponse:
//	case linqgo.MessageEventV2PartLink:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MessageEventV2PartUnion) AsAny() anyMessageEventV2Part {
	switch u.Type {
	case "text":
		return u.AsText()
	case "media":
		return u.AsMedia()
	case "link":
		return u.AsLink()
	}
	return nil
}

func (u MessageEventV2PartUnion) AsText() (v SchemasTextPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageEventV2PartUnion) AsMedia() (v SchemasMediaPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageEventV2PartUnion) AsLink() (v MessageEventV2PartLink) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageEventV2PartUnion) RawJSON() string { return u.JSON.raw }

func (r *MessageEventV2PartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A rich link preview part
type MessageEventV2PartLink struct {
	// Indicates this is a rich link preview part
	Type constant.Link `json:"type" api:"required"`
	// The URL
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEventV2PartLink) RawJSON() string { return r.JSON.raw }
func (r *MessageEventV2PartLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Preferred messaging service type. Includes "auto" for default fallback behavior.
type MessageEventV2PreferredService string

const (
	MessageEventV2PreferredServiceiMessage MessageEventV2PreferredService = "iMessage"
	MessageEventV2PreferredServiceSMS      MessageEventV2PreferredService = "SMS"
	MessageEventV2PreferredServiceRCS      MessageEventV2PreferredService = "RCS"
	MessageEventV2PreferredServiceAuto     MessageEventV2PreferredService = "auto"
)

// Reference to the message this is replying to (for threaded replies)
type MessageEventV2ReplyTo struct {
	// ID of the message being replied to
	MessageID string `json:"message_id" format:"uuid"`
	// Index of the part being replied to
	PartIndex int64 `json:"part_index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		PartIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEventV2ReplyTo) RawJSON() string { return r.JSON.raw }
func (r *MessageEventV2ReplyTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReactionEventBase struct {
	// Whether this reaction was from the owner of the phone number (true) or from
	// someone else (false)
	IsFromMe bool `json:"is_from_me" api:"required"`
	// Type of reaction. Standard iMessage tapbacks are love, like, dislike, laugh,
	// emphasize, question. Custom emoji reactions have type "custom" with the actual
	// emoji in the custom_emoji field. Sticker reactions have type "sticker" with
	// sticker attachment details in the sticker field.
	//
	// Any of "love", "like", "dislike", "laugh", "emphasize", "question", "custom",
	// "sticker".
	ReactionType shared.ReactionType `json:"reaction_type" api:"required"`
	// Chat identifier (UUID)
	ChatID string `json:"chat_id"`
	// The actual emoji when reaction_type is "custom". Null for standard tapbacks.
	CustomEmoji string `json:"custom_emoji" api:"nullable"`
	// DEPRECATED: Use from_handle instead. Phone number or email address of the person
	// who added/removed the reaction.
	//
	// Deprecated: deprecated
	From string `json:"from"`
	// The person who added/removed the reaction as a full handle object
	FromHandle shared.ChatHandle `json:"from_handle"`
	// Message identifier (UUID) that the reaction was added to or removed from
	MessageID string `json:"message_id"`
	// Index of the message part that was reacted to (0-based)
	PartIndex int64 `json:"part_index"`
	// When the reaction was added or removed
	ReactedAt time.Time `json:"reacted_at" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service"`
	// Sticker attachment details when reaction_type is "sticker". Null for non-sticker
	// reactions.
	Sticker ReactionEventBaseSticker `json:"sticker" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsFromMe     respjson.Field
		ReactionType respjson.Field
		ChatID       respjson.Field
		CustomEmoji  respjson.Field
		From         respjson.Field
		FromHandle   respjson.Field
		MessageID    respjson.Field
		PartIndex    respjson.Field
		ReactedAt    respjson.Field
		Service      respjson.Field
		Sticker      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReactionEventBase) RawJSON() string { return r.JSON.raw }
func (r *ReactionEventBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sticker attachment details when reaction_type is "sticker". Null for non-sticker
// reactions.
type ReactionEventBaseSticker struct {
	// Filename of the sticker
	FileName string `json:"file_name"`
	// Sticker image height in pixels
	Height int64 `json:"height"`
	// MIME type of the sticker image
	MimeType string `json:"mime_type"`
	// Presigned URL for downloading the sticker image (expires in 1 hour).
	URL string `json:"url" format:"uri"`
	// Sticker image width in pixels
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileName    respjson.Field
		Height      respjson.Field
		MimeType    respjson.Field
		URL         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReactionEventBaseSticker) RawJSON() string { return r.JSON.raw }
func (r *ReactionEventBaseSticker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A media attachment part
type SchemasMediaPartResponse struct {
	// Unique attachment identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Original filename
	Filename string `json:"filename" api:"required"`
	// MIME type of the file
	MimeType string `json:"mime_type" api:"required"`
	// File size in bytes
	SizeBytes int64 `json:"size_bytes" api:"required"`
	// Indicates this is a media attachment part
	//
	// Any of "media".
	Type SchemasMediaPartResponseType `json:"type" api:"required"`
	// Presigned URL for downloading the attachment (expires in 1 hour).
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Filename    respjson.Field
		MimeType    respjson.Field
		SizeBytes   respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchemasMediaPartResponse) RawJSON() string { return r.JSON.raw }
func (r *SchemasMediaPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates this is a media attachment part
type SchemasMediaPartResponseType string

const (
	SchemasMediaPartResponseTypeMedia SchemasMediaPartResponseType = "media"
)

// iMessage effect applied to a message (screen or bubble animation)
type SchemasMessageEffect struct {
	// Effect name (confetti, fireworks, slam, gentle, etc.)
	Name string `json:"name"`
	// Effect category
	//
	// Any of "screen", "bubble".
	Type SchemasMessageEffectType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchemasMessageEffect) RawJSON() string { return r.JSON.raw }
func (r *SchemasMessageEffect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Effect category
type SchemasMessageEffectType string

const (
	SchemasMessageEffectTypeScreen SchemasMessageEffectType = "screen"
	SchemasMessageEffectTypeBubble SchemasMessageEffectType = "bubble"
)

// A text message part
type SchemasTextPartResponse struct {
	// Indicates this is a text message part
	//
	// Any of "text".
	Type SchemasTextPartResponseType `json:"type" api:"required"`
	// The text content
	Value string `json:"value" api:"required"`
	// Text decorations applied to character ranges in the value
	TextDecorations []shared.TextDecoration `json:"text_decorations" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		Value           respjson.Field
		TextDecorations respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchemasTextPartResponse) RawJSON() string { return r.JSON.raw }
func (r *SchemasTextPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates this is a text message part
type SchemasTextPartResponseType string

const (
	SchemasTextPartResponseTypeText SchemasTextPartResponseType = "text"
)

// Complete webhook payload for message.sent events (2026-02-03 format)
type MessageSentWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unified payload for message webhooks when using `webhook_version: "2026-02-03"`.
	//
	// This schema is used for message.sent, message.received, message.delivered, and
	// message.read events when the subscription URL includes `?version=2026-02-03`.
	//
	// Key differences from V1 (2025-01-01):
	//
	//   - `direction`: "inbound" or "outbound" instead of `is_from_me` boolean
	//   - `sender_handle`: Full handle object for the sender
	//   - `chat`: Nested object with `id`, `is_group`, and `owner_handle`
	//   - Message fields (`id`, `parts`, `effect`, etc.) are at the top level, not
	//     nested in `message`
	//
	// Timestamps indicate the message state:
	//
	// - `message.sent`: sent_at set, delivered_at=null, read_at=null
	// - `message.received`: sent_at set, delivered_at=null, read_at=null
	// - `message.delivered`: sent_at set, delivered_at set, read_at=null
	// - `message.read`: sent_at set, delivered_at set, read_at set
	Data MessageEventV2 `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSentWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageSentWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.received events (2026-02-03 format)
type MessageReceivedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unified payload for message webhooks when using `webhook_version: "2026-02-03"`.
	//
	// This schema is used for message.sent, message.received, message.delivered, and
	// message.read events when the subscription URL includes `?version=2026-02-03`.
	//
	// Key differences from V1 (2025-01-01):
	//
	//   - `direction`: "inbound" or "outbound" instead of `is_from_me` boolean
	//   - `sender_handle`: Full handle object for the sender
	//   - `chat`: Nested object with `id`, `is_group`, and `owner_handle`
	//   - Message fields (`id`, `parts`, `effect`, etc.) are at the top level, not
	//     nested in `message`
	//
	// Timestamps indicate the message state:
	//
	// - `message.sent`: sent_at set, delivered_at=null, read_at=null
	// - `message.received`: sent_at set, delivered_at=null, read_at=null
	// - `message.delivered`: sent_at set, delivered_at set, read_at=null
	// - `message.read`: sent_at set, delivered_at set, read_at set
	Data MessageEventV2 `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageReceivedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageReceivedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.read events (2026-02-03 format)
type MessageReadWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unified payload for message webhooks when using `webhook_version: "2026-02-03"`.
	//
	// This schema is used for message.sent, message.received, message.delivered, and
	// message.read events when the subscription URL includes `?version=2026-02-03`.
	//
	// Key differences from V1 (2025-01-01):
	//
	//   - `direction`: "inbound" or "outbound" instead of `is_from_me` boolean
	//   - `sender_handle`: Full handle object for the sender
	//   - `chat`: Nested object with `id`, `is_group`, and `owner_handle`
	//   - Message fields (`id`, `parts`, `effect`, etc.) are at the top level, not
	//     nested in `message`
	//
	// Timestamps indicate the message state:
	//
	// - `message.sent`: sent_at set, delivered_at=null, read_at=null
	// - `message.received`: sent_at set, delivered_at=null, read_at=null
	// - `message.delivered`: sent_at set, delivered_at set, read_at=null
	// - `message.read`: sent_at set, delivered_at set, read_at set
	Data MessageEventV2 `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageReadWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageReadWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.delivered events (2026-02-03 format)
type MessageDeliveredWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unified payload for message webhooks when using `webhook_version: "2026-02-03"`.
	//
	// This schema is used for message.sent, message.received, message.delivered, and
	// message.read events when the subscription URL includes `?version=2026-02-03`.
	//
	// Key differences from V1 (2025-01-01):
	//
	//   - `direction`: "inbound" or "outbound" instead of `is_from_me` boolean
	//   - `sender_handle`: Full handle object for the sender
	//   - `chat`: Nested object with `id`, `is_group`, and `owner_handle`
	//   - Message fields (`id`, `parts`, `effect`, etc.) are at the top level, not
	//     nested in `message`
	//
	// Timestamps indicate the message state:
	//
	// - `message.sent`: sent_at set, delivered_at=null, read_at=null
	// - `message.received`: sent_at set, delivered_at=null, read_at=null
	// - `message.delivered`: sent_at set, delivered_at set, read_at=null
	// - `message.read`: sent_at set, delivered_at set, read_at set
	Data MessageEventV2 `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageDeliveredWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageDeliveredWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.failed events
type MessageFailedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error details for message.failed webhook events. See
	// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
	// code reference.
	Data MessageFailedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for message.failed webhook events. See
// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
// code reference.
type MessageFailedWebhookEventData struct {
	// Error codes in webhook failure events (3007, 4001).
	Code int64 `json:"code" api:"required"`
	// When the failure was detected
	FailedAt time.Time `json:"failed_at" api:"required" format:"date-time"`
	// Chat identifier (UUID)
	ChatID string `json:"chat_id"`
	// Message identifier (UUID)
	MessageID string `json:"message_id"`
	// Human-readable description of the failure
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		FailedAt    respjson.Field
		ChatID      respjson.Field
		MessageID   respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageFailedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *MessageFailedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.edited events (2026-02-03 format only)
type MessageEditedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for `message.edited` events (2026-02-03 format).
	//
	// Describes which part of a message was edited and when. Only text parts can be
	// edited. Only available for subscriptions using `webhook_version: "2026-02-03"`.
	Data MessageEditedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEditedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageEditedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for `message.edited` events (2026-02-03 format).
//
// Describes which part of a message was edited and when. Only text parts can be
// edited. Only available for subscriptions using `webhook_version: "2026-02-03"`.
type MessageEditedWebhookEventData struct {
	// Message identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Chat context
	Chat MessageEditedWebhookEventDataChat `json:"chat" api:"required"`
	// "outbound" if you sent the original message, "inbound" if you received it
	//
	// Any of "outbound", "inbound".
	Direction string `json:"direction" api:"required"`
	// When the edit occurred
	EditedAt time.Time `json:"edited_at" api:"required" format:"date-time"`
	// The edited part
	Part MessageEditedWebhookEventDataPart `json:"part" api:"required"`
	// The handle that sent (and edited) this message
	SenderHandle shared.ChatHandle `json:"sender_handle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Chat         respjson.Field
		Direction    respjson.Field
		EditedAt     respjson.Field
		Part         respjson.Field
		SenderHandle respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEditedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *MessageEditedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat context
type MessageEditedWebhookEventDataChat struct {
	// Chat identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group" api:"required"`
	// The handle that owns this chat (your phone number)
	OwnerHandle shared.ChatHandle `json:"owner_handle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsGroup     respjson.Field
		OwnerHandle respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEditedWebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *MessageEditedWebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The edited part
type MessageEditedWebhookEventDataPart struct {
	// Zero-based index of the edited part within the message
	Index int64 `json:"index" api:"required"`
	// New text content of the part
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEditedWebhookEventDataPart) RawJSON() string { return r.JSON.raw }
func (r *MessageEditedWebhookEventDataPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for reaction.added events
type ReactionAddedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for reaction.added webhook events
	Data ReactionEventBase `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReactionAddedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ReactionAddedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for reaction.removed events
type ReactionRemovedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for reaction.removed webhook events
	Data ReactionEventBase `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReactionRemovedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ReactionRemovedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for participant.added events
type ParticipantAddedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for participant.added webhook events
	Data ParticipantAddedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParticipantAddedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ParticipantAddedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for participant.added webhook events
type ParticipantAddedWebhookEventData struct {
	// DEPRECATED: Use participant instead. Handle (phone number or email address) of
	// the added participant.
	//
	// Deprecated: deprecated
	Handle string `json:"handle" api:"required"`
	// When the participant was added
	AddedAt time.Time `json:"added_at" format:"date-time"`
	// Chat identifier (UUID) of the group chat
	ChatID string `json:"chat_id"`
	// The added participant as a full handle object
	Participant shared.ChatHandle `json:"participant"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		AddedAt     respjson.Field
		ChatID      respjson.Field
		Participant respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParticipantAddedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ParticipantAddedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for participant.removed events
type ParticipantRemovedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for participant.removed webhook events
	Data ParticipantRemovedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParticipantRemovedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ParticipantRemovedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for participant.removed webhook events
type ParticipantRemovedWebhookEventData struct {
	// DEPRECATED: Use participant instead. Handle (phone number or email address) of
	// the removed participant.
	//
	// Deprecated: deprecated
	Handle string `json:"handle" api:"required"`
	// Chat identifier (UUID) of the group chat
	ChatID string `json:"chat_id"`
	// The removed participant as a full handle object
	Participant shared.ChatHandle `json:"participant"`
	// When the participant was removed
	RemovedAt time.Time `json:"removed_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		ChatID      respjson.Field
		Participant respjson.Field
		RemovedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParticipantRemovedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ParticipantRemovedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.created events
type ChatCreatedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.created webhook events. Matches GET /v3/chats/{chatId}
	// response.
	Data ChatCreatedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCreatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatCreatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.created webhook events. Matches GET /v3/chats/{chatId}
// response.
type ChatCreatedWebhookEventData struct {
	// Unique identifier for the chat
	ID string `json:"id" api:"required" format:"uuid"`
	// When the chat was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name for the chat. Defaults to a comma-separated list of recipient
	// handles. Can be updated for group chats.
	DisplayName string `json:"display_name" api:"required"`
	// List of chat participants with full handle details. Always contains at least two
	// handles (your phone number and the other participant).
	Handles []shared.ChatHandle `json:"handles" api:"required"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group" api:"required"`
	// When the chat was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DisplayName respjson.Field
		Handles     respjson.Field
		IsGroup     respjson.Field
		UpdatedAt   respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCreatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatCreatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_name_updated events
type ChatGroupNameUpdatedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.group_name_updated webhook events
	Data ChatGroupNameUpdatedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGroupNameUpdatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.group_name_updated webhook events
type ChatGroupNameUpdatedWebhookEventData struct {
	// Chat identifier (UUID) of the group chat
	ChatID string `json:"chat_id" api:"required"`
	// When the update occurred
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The handle who made the change.
	ChangedByHandle shared.ChatHandle `json:"changed_by_handle" api:"nullable"`
	// New group name (null if the name was removed)
	NewValue string `json:"new_value" api:"nullable"`
	// Previous group name (null if no previous name)
	OldValue string `json:"old_value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID          respjson.Field
		UpdatedAt       respjson.Field
		ChangedByHandle respjson.Field
		NewValue        respjson.Field
		OldValue        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGroupNameUpdatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_icon_updated events
type ChatGroupIconUpdatedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.group_icon_updated webhook events
	Data ChatGroupIconUpdatedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGroupIconUpdatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.group_icon_updated webhook events
type ChatGroupIconUpdatedWebhookEventData struct {
	// Chat identifier (UUID) of the group chat
	ChatID string `json:"chat_id" api:"required"`
	// When the update occurred
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The handle who made the change.
	ChangedByHandle shared.ChatHandle `json:"changed_by_handle" api:"nullable"`
	// New icon URL (null if the icon was removed)
	NewValue string `json:"new_value" api:"nullable"`
	// Previous icon URL (null if no previous icon)
	OldValue string `json:"old_value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID          respjson.Field
		UpdatedAt       respjson.Field
		ChangedByHandle respjson.Field
		NewValue        respjson.Field
		OldValue        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGroupIconUpdatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_name_update_failed events
type ChatGroupNameUpdateFailedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error details for chat.group_name_update_failed webhook events. See
	// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
	// code reference.
	Data ChatGroupNameUpdateFailedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGroupNameUpdateFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdateFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for chat.group_name_update_failed webhook events. See
// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
// code reference.
type ChatGroupNameUpdateFailedWebhookEventData struct {
	// Chat identifier (UUID) of the group chat
	ChatID string `json:"chat_id" api:"required"`
	// Error codes in webhook failure events (3007, 4001).
	ErrorCode int64 `json:"error_code" api:"required"`
	// When the failure was detected
	FailedAt time.Time `json:"failed_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		ErrorCode   respjson.Field
		FailedAt    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGroupNameUpdateFailedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdateFailedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_icon_update_failed events
type ChatGroupIconUpdateFailedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error details for chat.group_icon_update_failed webhook events. See
	// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
	// code reference.
	Data ChatGroupIconUpdateFailedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGroupIconUpdateFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdateFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for chat.group_icon_update_failed webhook events. See
// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
// code reference.
type ChatGroupIconUpdateFailedWebhookEventData struct {
	// Chat identifier (UUID) of the group chat
	ChatID string `json:"chat_id" api:"required"`
	// Error codes in webhook failure events (3007, 4001).
	ErrorCode int64 `json:"error_code" api:"required"`
	// When the failure was detected
	FailedAt time.Time `json:"failed_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		ErrorCode   respjson.Field
		FailedAt    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGroupIconUpdateFailedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdateFailedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.typing_indicator.started events
type ChatTypingIndicatorStartedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.typing_indicator.started webhook events
	Data ChatTypingIndicatorStartedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatTypingIndicatorStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.typing_indicator.started webhook events
type ChatTypingIndicatorStartedWebhookEventData struct {
	// Chat identifier
	ChatID string `json:"chat_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatTypingIndicatorStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.typing_indicator.stopped events
type ChatTypingIndicatorStoppedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.typing_indicator.stopped webhook events
	Data ChatTypingIndicatorStoppedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Valid webhook event types that can be subscribed to.
	//
	// **Note:** `message.edited` is only delivered to subscriptions using
	// `webhook_version: "2026-02-03"`. Subscribing to this event on a v2025
	// subscription will not produce any deliveries.
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType WebhookEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatTypingIndicatorStoppedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStoppedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.typing_indicator.stopped webhook events
type ChatTypingIndicatorStoppedWebhookEventData struct {
	// Chat identifier
	ChatID string `json:"chat_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatTypingIndicatorStoppedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStoppedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for phone_number.status_updated events
type PhoneNumberStatusUpdatedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for phone_number.status_updated webhook events
	Data PhoneNumberStatusUpdatedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// The type of event
	//
	// Any of "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "call.initiated", "call.ringing",
	// "call.answered", "call.ended", "call.failed", "call.declined", "call.no_answer".
	EventType PhoneNumberStatusUpdatedWebhookEventEventType `json:"event_type" api:"required"`
	// Partner identifier. Present on all webhooks for cross-referencing.
	PartnerID string `json:"partner_id" api:"required"`
	// Trace ID for debugging and correlation across systems.
	TraceID string `json:"trace_id" api:"required"`
	// Date-based webhook payload version. Determined by the `?version=` query
	// parameter in your webhook subscription URL. If no version parameter is
	// specified, defaults based on subscription creation date.
	WebhookVersion string `json:"webhook_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberStatusUpdatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberStatusUpdatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for phone_number.status_updated webhook events
type PhoneNumberStatusUpdatedWebhookEventData struct {
	// When the status change occurred
	ChangedAt time.Time `json:"changed_at" api:"required" format:"date-time"`
	// The new service status
	//
	// Any of "ACTIVE", "FLAGGED".
	NewStatus string `json:"new_status" api:"required"`
	// Phone number in E.164 format
	PhoneNumber string `json:"phone_number" api:"required"`
	// The previous service status
	//
	// Any of "ACTIVE", "FLAGGED".
	PreviousStatus string `json:"previous_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedAt      respjson.Field
		NewStatus      respjson.Field
		PhoneNumber    respjson.Field
		PreviousStatus respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberStatusUpdatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberStatusUpdatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event
type PhoneNumberStatusUpdatedWebhookEventEventType string

const (
	PhoneNumberStatusUpdatedWebhookEventEventTypeMessageSent                PhoneNumberStatusUpdatedWebhookEventEventType = "message.sent"
	PhoneNumberStatusUpdatedWebhookEventEventTypeMessageReceived            PhoneNumberStatusUpdatedWebhookEventEventType = "message.received"
	PhoneNumberStatusUpdatedWebhookEventEventTypeMessageRead                PhoneNumberStatusUpdatedWebhookEventEventType = "message.read"
	PhoneNumberStatusUpdatedWebhookEventEventTypeMessageDelivered           PhoneNumberStatusUpdatedWebhookEventEventType = "message.delivered"
	PhoneNumberStatusUpdatedWebhookEventEventTypeMessageFailed              PhoneNumberStatusUpdatedWebhookEventEventType = "message.failed"
	PhoneNumberStatusUpdatedWebhookEventEventTypeMessageEdited              PhoneNumberStatusUpdatedWebhookEventEventType = "message.edited"
	PhoneNumberStatusUpdatedWebhookEventEventTypeReactionAdded              PhoneNumberStatusUpdatedWebhookEventEventType = "reaction.added"
	PhoneNumberStatusUpdatedWebhookEventEventTypeReactionRemoved            PhoneNumberStatusUpdatedWebhookEventEventType = "reaction.removed"
	PhoneNumberStatusUpdatedWebhookEventEventTypeParticipantAdded           PhoneNumberStatusUpdatedWebhookEventEventType = "participant.added"
	PhoneNumberStatusUpdatedWebhookEventEventTypeParticipantRemoved         PhoneNumberStatusUpdatedWebhookEventEventType = "participant.removed"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatCreated                PhoneNumberStatusUpdatedWebhookEventEventType = "chat.created"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatGroupNameUpdated       PhoneNumberStatusUpdatedWebhookEventEventType = "chat.group_name_updated"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatGroupIconUpdated       PhoneNumberStatusUpdatedWebhookEventEventType = "chat.group_icon_updated"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatGroupNameUpdateFailed  PhoneNumberStatusUpdatedWebhookEventEventType = "chat.group_name_update_failed"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatGroupIconUpdateFailed  PhoneNumberStatusUpdatedWebhookEventEventType = "chat.group_icon_update_failed"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatTypingIndicatorStarted PhoneNumberStatusUpdatedWebhookEventEventType = "chat.typing_indicator.started"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatTypingIndicatorStopped PhoneNumberStatusUpdatedWebhookEventEventType = "chat.typing_indicator.stopped"
	PhoneNumberStatusUpdatedWebhookEventEventTypePhoneNumberStatusUpdated   PhoneNumberStatusUpdatedWebhookEventEventType = "phone_number.status_updated"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallInitiated              PhoneNumberStatusUpdatedWebhookEventEventType = "call.initiated"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallRinging                PhoneNumberStatusUpdatedWebhookEventEventType = "call.ringing"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallAnswered               PhoneNumberStatusUpdatedWebhookEventEventType = "call.answered"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallEnded                  PhoneNumberStatusUpdatedWebhookEventEventType = "call.ended"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallFailed                 PhoneNumberStatusUpdatedWebhookEventEventType = "call.failed"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallDeclined               PhoneNumberStatusUpdatedWebhookEventEventType = "call.declined"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallNoAnswer               PhoneNumberStatusUpdatedWebhookEventEventType = "call.no_answer"
)

// EventsWebhookEventUnion contains all possible properties and values from
// [MessageSentWebhookEvent], [MessageReceivedWebhookEvent],
// [MessageReadWebhookEvent], [MessageDeliveredWebhookEvent],
// [MessageFailedWebhookEvent], [MessageEditedWebhookEvent],
// [ReactionAddedWebhookEvent], [ReactionRemovedWebhookEvent],
// [ParticipantAddedWebhookEvent], [ParticipantRemovedWebhookEvent],
// [ChatCreatedWebhookEvent], [ChatGroupNameUpdatedWebhookEvent],
// [ChatGroupIconUpdatedWebhookEvent], [ChatGroupNameUpdateFailedWebhookEvent],
// [ChatGroupIconUpdateFailedWebhookEvent],
// [ChatTypingIndicatorStartedWebhookEvent],
// [ChatTypingIndicatorStoppedWebhookEvent],
// [PhoneNumberStatusUpdatedWebhookEvent].
//
// Use the [EventsWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EventsWebhookEventUnion struct {
	APIVersion string    `json:"api_version"`
	CreatedAt  time.Time `json:"created_at"`
	// This field is a union of [MessageEventV2], [MessageFailedWebhookEventData],
	// [MessageEditedWebhookEventData], [ReactionEventBase],
	// [ParticipantAddedWebhookEventData], [ParticipantRemovedWebhookEventData],
	// [ChatCreatedWebhookEventData], [ChatGroupNameUpdatedWebhookEventData],
	// [ChatGroupIconUpdatedWebhookEventData],
	// [ChatGroupNameUpdateFailedWebhookEventData],
	// [ChatGroupIconUpdateFailedWebhookEventData],
	// [ChatTypingIndicatorStartedWebhookEventData],
	// [ChatTypingIndicatorStoppedWebhookEventData],
	// [PhoneNumberStatusUpdatedWebhookEventData]
	Data    EventsWebhookEventUnionData `json:"data"`
	EventID string                      `json:"event_id"`
	// Any of nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
	// nil, nil, nil, nil.
	EventType      string `json:"event_type"`
	PartnerID      string `json:"partner_id"`
	TraceID        string `json:"trace_id"`
	WebhookVersion string `json:"webhook_version"`
	JSON           struct {
		APIVersion     respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		PartnerID      respjson.Field
		TraceID        respjson.Field
		WebhookVersion respjson.Field
		raw            string
	} `json:"-"`
}

func (u EventsWebhookEventUnion) AsMessageSentWebhookEvent() (v MessageSentWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageReceivedWebhookEvent() (v MessageReceivedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageReadWebhookEvent() (v MessageReadWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageDeliveredWebhookEvent() (v MessageDeliveredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageFailedWebhookEvent() (v MessageFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageEditedWebhookEvent() (v MessageEditedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsReactionAddedWebhookEvent() (v ReactionAddedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsReactionRemovedWebhookEvent() (v ReactionRemovedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsParticipantAddedWebhookEvent() (v ParticipantAddedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsParticipantRemovedWebhookEvent() (v ParticipantRemovedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatCreatedWebhookEvent() (v ChatCreatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupNameUpdatedWebhookEvent() (v ChatGroupNameUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupIconUpdatedWebhookEvent() (v ChatGroupIconUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupNameUpdateFailedWebhookEvent() (v ChatGroupNameUpdateFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupIconUpdateFailedWebhookEvent() (v ChatGroupIconUpdateFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatTypingIndicatorStartedWebhookEvent() (v ChatTypingIndicatorStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatTypingIndicatorStoppedWebhookEvent() (v ChatTypingIndicatorStoppedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsPhoneNumberStatusUpdatedWebhookEvent() (v PhoneNumberStatusUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventsWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *EventsWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventsWebhookEventUnionData is an implicit subunion of
// [EventsWebhookEventUnion]. EventsWebhookEventUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [EventsWebhookEventUnion].
type EventsWebhookEventUnionData struct {
	ID string `json:"id"`
	// This field is a union of [MessageEventV2Chat],
	// [MessageEditedWebhookEventDataChat]
	Chat      EventsWebhookEventUnionDataChat `json:"chat"`
	Direction string                          `json:"direction"`
	// This field is from variant [MessageEventV2].
	Parts []MessageEventV2PartUnion `json:"parts"`
	// This field is from variant [MessageEventV2].
	SenderHandle shared.ChatHandle `json:"sender_handle"`
	// This field is from variant [MessageEventV2].
	Service shared.ServiceType `json:"service"`
	// This field is from variant [MessageEventV2].
	DeliveredAt time.Time `json:"delivered_at"`
	// This field is from variant [MessageEventV2].
	Effect SchemasMessageEffect `json:"effect"`
	// This field is from variant [MessageEventV2].
	IdempotencyKey string `json:"idempotency_key"`
	// This field is from variant [MessageEventV2].
	PreferredService MessageEventV2PreferredService `json:"preferred_service"`
	// This field is from variant [MessageEventV2].
	ReadAt time.Time `json:"read_at"`
	// This field is from variant [MessageEventV2].
	ReplyTo MessageEventV2ReplyTo `json:"reply_to"`
	// This field is from variant [MessageEventV2].
	SentAt time.Time `json:"sent_at"`
	// This field is from variant [MessageFailedWebhookEventData].
	Code      int64     `json:"code"`
	FailedAt  time.Time `json:"failed_at"`
	ChatID    string    `json:"chat_id"`
	MessageID string    `json:"message_id"`
	// This field is from variant [MessageFailedWebhookEventData].
	Reason string `json:"reason"`
	// This field is from variant [MessageEditedWebhookEventData].
	EditedAt time.Time `json:"edited_at"`
	// This field is from variant [MessageEditedWebhookEventData].
	Part MessageEditedWebhookEventDataPart `json:"part"`
	// This field is from variant [ReactionEventBase].
	IsFromMe bool `json:"is_from_me"`
	// This field is from variant [ReactionEventBase].
	ReactionType shared.ReactionType `json:"reaction_type"`
	// This field is from variant [ReactionEventBase].
	CustomEmoji string `json:"custom_emoji"`
	// This field is from variant [ReactionEventBase].
	From string `json:"from"`
	// This field is from variant [ReactionEventBase].
	FromHandle shared.ChatHandle `json:"from_handle"`
	// This field is from variant [ReactionEventBase].
	PartIndex int64 `json:"part_index"`
	// This field is from variant [ReactionEventBase].
	ReactedAt time.Time `json:"reacted_at"`
	// This field is from variant [ReactionEventBase].
	Sticker ReactionEventBaseSticker `json:"sticker"`
	Handle  string                   `json:"handle"`
	// This field is from variant [ParticipantAddedWebhookEventData].
	AddedAt time.Time `json:"added_at"`
	// This field is from variant [ParticipantAddedWebhookEventData].
	Participant shared.ChatHandle `json:"participant"`
	// This field is from variant [ParticipantRemovedWebhookEventData].
	RemovedAt time.Time `json:"removed_at"`
	// This field is from variant [ChatCreatedWebhookEventData].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [ChatCreatedWebhookEventData].
	DisplayName string `json:"display_name"`
	// This field is from variant [ChatCreatedWebhookEventData].
	Handles []shared.ChatHandle `json:"handles"`
	// This field is from variant [ChatCreatedWebhookEventData].
	IsGroup   bool      `json:"is_group"`
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant [ChatGroupNameUpdatedWebhookEventData].
	ChangedByHandle shared.ChatHandle `json:"changed_by_handle"`
	NewValue        string            `json:"new_value"`
	OldValue        string            `json:"old_value"`
	ErrorCode       int64             `json:"error_code"`
	// This field is from variant [PhoneNumberStatusUpdatedWebhookEventData].
	ChangedAt time.Time `json:"changed_at"`
	// This field is from variant [PhoneNumberStatusUpdatedWebhookEventData].
	NewStatus string `json:"new_status"`
	// This field is from variant [PhoneNumberStatusUpdatedWebhookEventData].
	PhoneNumber string `json:"phone_number"`
	// This field is from variant [PhoneNumberStatusUpdatedWebhookEventData].
	PreviousStatus string `json:"previous_status"`
	JSON           struct {
		ID               respjson.Field
		Chat             respjson.Field
		Direction        respjson.Field
		Parts            respjson.Field
		SenderHandle     respjson.Field
		Service          respjson.Field
		DeliveredAt      respjson.Field
		Effect           respjson.Field
		IdempotencyKey   respjson.Field
		PreferredService respjson.Field
		ReadAt           respjson.Field
		ReplyTo          respjson.Field
		SentAt           respjson.Field
		Code             respjson.Field
		FailedAt         respjson.Field
		ChatID           respjson.Field
		MessageID        respjson.Field
		Reason           respjson.Field
		EditedAt         respjson.Field
		Part             respjson.Field
		IsFromMe         respjson.Field
		ReactionType     respjson.Field
		CustomEmoji      respjson.Field
		From             respjson.Field
		FromHandle       respjson.Field
		PartIndex        respjson.Field
		ReactedAt        respjson.Field
		Sticker          respjson.Field
		Handle           respjson.Field
		AddedAt          respjson.Field
		Participant      respjson.Field
		RemovedAt        respjson.Field
		CreatedAt        respjson.Field
		DisplayName      respjson.Field
		Handles          respjson.Field
		IsGroup          respjson.Field
		UpdatedAt        respjson.Field
		ChangedByHandle  respjson.Field
		NewValue         respjson.Field
		OldValue         respjson.Field
		ErrorCode        respjson.Field
		ChangedAt        respjson.Field
		NewStatus        respjson.Field
		PhoneNumber      respjson.Field
		PreviousStatus   respjson.Field
		raw              string
	} `json:"-"`
}

func (r *EventsWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventsWebhookEventUnionDataChat is an implicit subunion of
// [EventsWebhookEventUnion]. EventsWebhookEventUnionDataChat provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [EventsWebhookEventUnion].
type EventsWebhookEventUnionDataChat struct {
	ID      string `json:"id"`
	IsGroup bool   `json:"is_group"`
	// This field is from variant [MessageEventV2Chat].
	OwnerHandle shared.ChatHandle `json:"owner_handle"`
	JSON        struct {
		ID          respjson.Field
		IsGroup     respjson.Field
		OwnerHandle respjson.Field
		raw         string
	} `json:"-"`
}

func (r *EventsWebhookEventUnionDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
