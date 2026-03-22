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

// Message content nested within webhook events
type MessagePayload struct {
	// Message identifier
	ID string `json:"id" format:"uuid"`
	// When the message record was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// When the message was delivered
	DeliveredAt time.Time `json:"delivered_at" api:"nullable" format:"date-time"`
	// iMessage effect applied to a message (screen or bubble animation)
	Effect SchemasMessageEffect `json:"effect"`
	// Whether the message has been delivered
	IsDelivered bool `json:"is_delivered"`
	// Whether the message has been read
	IsRead bool `json:"is_read"`
	// Message content parts (text and/or media)
	Parts []MessagePayloadPartUnion `json:"parts"`
	// When the message was read
	ReadAt time.Time `json:"read_at" api:"nullable" format:"date-time"`
	// Reference to the message this is replying to
	ReplyTo MessagePayloadReplyTo `json:"reply_to"`
	// When the message was sent
	SentAt time.Time `json:"sent_at" api:"nullable" format:"date-time"`
	// When the message record was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DeliveredAt respjson.Field
		Effect      respjson.Field
		IsDelivered respjson.Field
		IsRead      respjson.Field
		Parts       respjson.Field
		ReadAt      respjson.Field
		ReplyTo     respjson.Field
		SentAt      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagePayload) RawJSON() string { return r.JSON.raw }
func (r *MessagePayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessagePayloadPartUnion contains all possible properties and values from
// [SchemasTextPartResponse], [SchemasMediaPartResponse],
// [MessagePayloadPartSchemasLinkPartResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessagePayloadPartUnion struct {
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

func (u MessagePayloadPartUnion) AsSchemasTextPartResponse() (v SchemasTextPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessagePayloadPartUnion) AsSchemasMediaPartResponse() (v SchemasMediaPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessagePayloadPartUnion) AsMessagePayloadPartSchemasLinkPartResponse() (v MessagePayloadPartSchemasLinkPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessagePayloadPartUnion) RawJSON() string { return u.JSON.raw }

func (r *MessagePayloadPartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A rich link preview part
type MessagePayloadPartSchemasLinkPartResponse struct {
	// Indicates this is a rich link preview part
	//
	// Any of "link".
	Type string `json:"type" api:"required"`
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
func (r MessagePayloadPartSchemasLinkPartResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagePayloadPartSchemasLinkPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reference to the message this is replying to
type MessagePayloadReplyTo struct {
	// The ID of the message being replied to
	MessageID string `json:"message_id" format:"uuid"`
	// Index of the message part being replied to (0-based)
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
func (r MessagePayloadReplyTo) RawJSON() string { return r.JSON.raw }
func (r *MessagePayloadReplyTo) UnmarshalJSON(data []byte) error {
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
type MessageSentV2026WebhookEvent struct {
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
func (r MessageSentV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageSentV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.received events (2026-02-03 format)
type MessageReceivedV2026WebhookEvent struct {
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
func (r MessageReceivedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageReceivedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.read events (2026-02-03 format)
type MessageReadV2026WebhookEvent struct {
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
func (r MessageReadV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageReadV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.delivered events (2026-02-03 format)
type MessageDeliveredV2026WebhookEvent struct {
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
func (r MessageDeliveredV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageDeliveredV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.failed events
type MessageFailedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error details for message.failed webhook events. See
	// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
	// code reference.
	Data MessageFailedV2026WebhookEventData `json:"data" api:"required"`
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
func (r MessageFailedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageFailedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for message.failed webhook events. See
// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
// code reference.
type MessageFailedV2026WebhookEventData struct {
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
func (r MessageFailedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *MessageFailedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for reaction.added events
type ReactionAddedV2026WebhookEvent struct {
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
func (r ReactionAddedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ReactionAddedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for reaction.removed events
type ReactionRemovedV2026WebhookEvent struct {
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
func (r ReactionRemovedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ReactionRemovedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for participant.added events
type ParticipantAddedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for participant.added webhook events
	Data ParticipantAddedV2026WebhookEventData `json:"data" api:"required"`
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
func (r ParticipantAddedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ParticipantAddedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for participant.added webhook events
type ParticipantAddedV2026WebhookEventData struct {
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
func (r ParticipantAddedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ParticipantAddedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for participant.removed events
type ParticipantRemovedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for participant.removed webhook events
	Data ParticipantRemovedV2026WebhookEventData `json:"data" api:"required"`
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
func (r ParticipantRemovedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ParticipantRemovedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for participant.removed webhook events
type ParticipantRemovedV2026WebhookEventData struct {
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
func (r ParticipantRemovedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ParticipantRemovedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_name_updated events
type ChatGroupNameUpdatedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.group_name_updated webhook events
	Data ChatGroupNameUpdatedV2026WebhookEventData `json:"data" api:"required"`
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
func (r ChatGroupNameUpdatedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdatedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.group_name_updated webhook events
type ChatGroupNameUpdatedV2026WebhookEventData struct {
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
func (r ChatGroupNameUpdatedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdatedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_icon_updated events
type ChatGroupIconUpdatedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.group_icon_updated webhook events
	Data ChatGroupIconUpdatedV2026WebhookEventData `json:"data" api:"required"`
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
func (r ChatGroupIconUpdatedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdatedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.group_icon_updated webhook events
type ChatGroupIconUpdatedV2026WebhookEventData struct {
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
func (r ChatGroupIconUpdatedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdatedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_name_update_failed events
type ChatGroupNameUpdateFailedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error details for chat.group_name_update_failed webhook events. See
	// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
	// code reference.
	Data ChatGroupNameUpdateFailedV2026WebhookEventData `json:"data" api:"required"`
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
func (r ChatGroupNameUpdateFailedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdateFailedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for chat.group_name_update_failed webhook events. See
// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
// code reference.
type ChatGroupNameUpdateFailedV2026WebhookEventData struct {
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
func (r ChatGroupNameUpdateFailedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdateFailedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_icon_update_failed events
type ChatGroupIconUpdateFailedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error details for chat.group_icon_update_failed webhook events. See
	// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
	// code reference.
	Data ChatGroupIconUpdateFailedV2026WebhookEventData `json:"data" api:"required"`
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
func (r ChatGroupIconUpdateFailedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdateFailedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for chat.group_icon_update_failed webhook events. See
// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
// code reference.
type ChatGroupIconUpdateFailedV2026WebhookEventData struct {
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
func (r ChatGroupIconUpdateFailedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdateFailedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.created events
type ChatCreatedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.created webhook events. Matches GET /v3/chats/{chatId}
	// response.
	Data ChatCreatedV2026WebhookEventData `json:"data" api:"required"`
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
func (r ChatCreatedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatCreatedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.created webhook events. Matches GET /v3/chats/{chatId}
// response.
type ChatCreatedV2026WebhookEventData struct {
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
func (r ChatCreatedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatCreatedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.typing_indicator.started events
type ChatTypingIndicatorStartedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.typing_indicator.started webhook events
	Data ChatTypingIndicatorStartedV2026WebhookEventData `json:"data" api:"required"`
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
func (r ChatTypingIndicatorStartedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStartedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.typing_indicator.started webhook events
type ChatTypingIndicatorStartedV2026WebhookEventData struct {
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
func (r ChatTypingIndicatorStartedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStartedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.typing_indicator.stopped events
type ChatTypingIndicatorStoppedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.typing_indicator.stopped webhook events
	Data ChatTypingIndicatorStoppedV2026WebhookEventData `json:"data" api:"required"`
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
func (r ChatTypingIndicatorStoppedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStoppedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.typing_indicator.stopped webhook events
type ChatTypingIndicatorStoppedV2026WebhookEventData struct {
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
func (r ChatTypingIndicatorStoppedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStoppedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.edited events (2026-02-03 format only)
type MessageEditedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for `message.edited` events (2026-02-03 format).
	//
	// Describes which part of a message was edited and when. Only text parts can be
	// edited. Only available for subscriptions using `webhook_version: "2026-02-03"`.
	Data MessageEditedV2026WebhookEventData `json:"data" api:"required"`
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
func (r MessageEditedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageEditedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for `message.edited` events (2026-02-03 format).
//
// Describes which part of a message was edited and when. Only text parts can be
// edited. Only available for subscriptions using `webhook_version: "2026-02-03"`.
type MessageEditedV2026WebhookEventData struct {
	// Message identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Chat context
	Chat MessageEditedV2026WebhookEventDataChat `json:"chat" api:"required"`
	// "outbound" if you sent the original message, "inbound" if you received it
	//
	// Any of "outbound", "inbound".
	Direction string `json:"direction" api:"required"`
	// When the edit occurred
	EditedAt time.Time `json:"edited_at" api:"required" format:"date-time"`
	// The edited part
	Part MessageEditedV2026WebhookEventDataPart `json:"part" api:"required"`
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
func (r MessageEditedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *MessageEditedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat context
type MessageEditedV2026WebhookEventDataChat struct {
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
func (r MessageEditedV2026WebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *MessageEditedV2026WebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The edited part
type MessageEditedV2026WebhookEventDataPart struct {
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
func (r MessageEditedV2026WebhookEventDataPart) RawJSON() string { return r.JSON.raw }
func (r *MessageEditedV2026WebhookEventDataPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for phone_number.status_updated events
type PhoneNumberStatusUpdatedV2026WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for phone_number.status_updated webhook events
	Data PhoneNumberStatusUpdatedV2026WebhookEventData `json:"data" api:"required"`
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
	EventType PhoneNumberStatusUpdatedV2026WebhookEventEventType `json:"event_type" api:"required"`
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
func (r PhoneNumberStatusUpdatedV2026WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberStatusUpdatedV2026WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for phone_number.status_updated webhook events
type PhoneNumberStatusUpdatedV2026WebhookEventData struct {
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
func (r PhoneNumberStatusUpdatedV2026WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberStatusUpdatedV2026WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event
type PhoneNumberStatusUpdatedV2026WebhookEventEventType string

const (
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeMessageSent                PhoneNumberStatusUpdatedV2026WebhookEventEventType = "message.sent"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeMessageReceived            PhoneNumberStatusUpdatedV2026WebhookEventEventType = "message.received"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeMessageRead                PhoneNumberStatusUpdatedV2026WebhookEventEventType = "message.read"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeMessageDelivered           PhoneNumberStatusUpdatedV2026WebhookEventEventType = "message.delivered"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeMessageFailed              PhoneNumberStatusUpdatedV2026WebhookEventEventType = "message.failed"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeMessageEdited              PhoneNumberStatusUpdatedV2026WebhookEventEventType = "message.edited"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeReactionAdded              PhoneNumberStatusUpdatedV2026WebhookEventEventType = "reaction.added"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeReactionRemoved            PhoneNumberStatusUpdatedV2026WebhookEventEventType = "reaction.removed"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeParticipantAdded           PhoneNumberStatusUpdatedV2026WebhookEventEventType = "participant.added"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeParticipantRemoved         PhoneNumberStatusUpdatedV2026WebhookEventEventType = "participant.removed"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeChatCreated                PhoneNumberStatusUpdatedV2026WebhookEventEventType = "chat.created"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeChatGroupNameUpdated       PhoneNumberStatusUpdatedV2026WebhookEventEventType = "chat.group_name_updated"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeChatGroupIconUpdated       PhoneNumberStatusUpdatedV2026WebhookEventEventType = "chat.group_icon_updated"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeChatGroupNameUpdateFailed  PhoneNumberStatusUpdatedV2026WebhookEventEventType = "chat.group_name_update_failed"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeChatGroupIconUpdateFailed  PhoneNumberStatusUpdatedV2026WebhookEventEventType = "chat.group_icon_update_failed"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeChatTypingIndicatorStarted PhoneNumberStatusUpdatedV2026WebhookEventEventType = "chat.typing_indicator.started"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeChatTypingIndicatorStopped PhoneNumberStatusUpdatedV2026WebhookEventEventType = "chat.typing_indicator.stopped"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypePhoneNumberStatusUpdated   PhoneNumberStatusUpdatedV2026WebhookEventEventType = "phone_number.status_updated"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeCallInitiated              PhoneNumberStatusUpdatedV2026WebhookEventEventType = "call.initiated"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeCallRinging                PhoneNumberStatusUpdatedV2026WebhookEventEventType = "call.ringing"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeCallAnswered               PhoneNumberStatusUpdatedV2026WebhookEventEventType = "call.answered"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeCallEnded                  PhoneNumberStatusUpdatedV2026WebhookEventEventType = "call.ended"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeCallFailed                 PhoneNumberStatusUpdatedV2026WebhookEventEventType = "call.failed"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeCallDeclined               PhoneNumberStatusUpdatedV2026WebhookEventEventType = "call.declined"
	PhoneNumberStatusUpdatedV2026WebhookEventEventTypeCallNoAnswer               PhoneNumberStatusUpdatedV2026WebhookEventEventType = "call.no_answer"
)

// Complete webhook payload for message.sent events (2025-01-01 format)
type MessageSentV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unified payload for message.sent and message.received webhook events (2025-01-01
	// format)
	Data MessageSentV2025WebhookEventData `json:"data" api:"required"`
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
func (r MessageSentV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageSentV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified payload for message.sent and message.received webhook events (2025-01-01
// format)
type MessageSentV2025WebhookEventData struct {
	// Chat identifier
	ChatID string `json:"chat_id" format:"uuid"`
	// DEPRECATED: Use from_handle instead. Phone number or email address of the
	// message sender.
	//
	// Deprecated: deprecated
	From string `json:"from"`
	// The sender of this message as a full handle object
	FromHandle shared.ChatHandle `json:"from_handle"`
	// Idempotency key for the message. Used for deduplication of outbound messages.
	IdempotencyKey string `json:"idempotency_key" api:"nullable"`
	// Whether the message was sent by us (true for sent events, false for received
	// events)
	IsFromMe bool `json:"is_from_me"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group"`
	// Message content nested within webhook events
	Message MessagePayload `json:"message"`
	// Preferred messaging service type. Includes "auto" for default fallback behavior.
	//
	// Any of "iMessage", "SMS", "RCS", "auto".
	PreferredService string `json:"preferred_service" api:"nullable"`
	// When the message was received. Null for sent events.
	ReceivedAt time.Time `json:"received_at" api:"nullable" format:"date-time"`
	// Our phone number that received the message as a full handle object. Null for
	// sent events.
	RecipientHandle shared.ChatHandle `json:"recipient_handle" api:"nullable"`
	// DEPRECATED: Use recipient_handle instead. Our phone number that received the
	// message. Null for sent events.
	//
	// Deprecated: deprecated
	RecipientPhone string `json:"recipient_phone" api:"nullable"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID           respjson.Field
		From             respjson.Field
		FromHandle       respjson.Field
		IdempotencyKey   respjson.Field
		IsFromMe         respjson.Field
		IsGroup          respjson.Field
		Message          respjson.Field
		PreferredService respjson.Field
		ReceivedAt       respjson.Field
		RecipientHandle  respjson.Field
		RecipientPhone   respjson.Field
		Service          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSentV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *MessageSentV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.received events (2025-01-01 format)
type MessageReceivedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unified payload for message.sent and message.received webhook events (2025-01-01
	// format)
	Data MessageReceivedV2025WebhookEventData `json:"data" api:"required"`
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
func (r MessageReceivedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageReceivedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified payload for message.sent and message.received webhook events (2025-01-01
// format)
type MessageReceivedV2025WebhookEventData struct {
	// Chat identifier
	ChatID string `json:"chat_id" format:"uuid"`
	// DEPRECATED: Use from_handle instead. Phone number or email address of the
	// message sender.
	//
	// Deprecated: deprecated
	From string `json:"from"`
	// The sender of this message as a full handle object
	FromHandle shared.ChatHandle `json:"from_handle"`
	// Idempotency key for the message. Used for deduplication of outbound messages.
	IdempotencyKey string `json:"idempotency_key" api:"nullable"`
	// Whether the message was sent by us (true for sent events, false for received
	// events)
	IsFromMe bool `json:"is_from_me"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group"`
	// Message content nested within webhook events
	Message MessagePayload `json:"message"`
	// Preferred messaging service type. Includes "auto" for default fallback behavior.
	//
	// Any of "iMessage", "SMS", "RCS", "auto".
	PreferredService string `json:"preferred_service" api:"nullable"`
	// When the message was received. Null for sent events.
	ReceivedAt time.Time `json:"received_at" api:"nullable" format:"date-time"`
	// Our phone number that received the message as a full handle object. Null for
	// sent events.
	RecipientHandle shared.ChatHandle `json:"recipient_handle" api:"nullable"`
	// DEPRECATED: Use recipient_handle instead. Our phone number that received the
	// message. Null for sent events.
	//
	// Deprecated: deprecated
	RecipientPhone string `json:"recipient_phone" api:"nullable"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID           respjson.Field
		From             respjson.Field
		FromHandle       respjson.Field
		IdempotencyKey   respjson.Field
		IsFromMe         respjson.Field
		IsGroup          respjson.Field
		Message          respjson.Field
		PreferredService respjson.Field
		ReceivedAt       respjson.Field
		RecipientHandle  respjson.Field
		RecipientPhone   respjson.Field
		Service          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageReceivedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *MessageReceivedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.read events (2025-01-01 format)
type MessageReadV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for message.read webhook events (2025-01-01 format). Extends
	// MessageEvent with read_at and message_id.
	Data MessageReadV2025WebhookEventData `json:"data" api:"required"`
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
func (r MessageReadV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageReadV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for message.read webhook events (2025-01-01 format). Extends
// MessageEvent with read_at and message_id.
type MessageReadV2025WebhookEventData struct {
	// When the message was read
	ReadAt time.Time `json:"read_at" api:"required" format:"date-time"`
	// Chat identifier
	ChatID string `json:"chat_id" format:"uuid"`
	// DEPRECATED: Use from_handle instead. Phone number or email address of the
	// message sender.
	//
	// Deprecated: deprecated
	From string `json:"from"`
	// The sender of this message as a full handle object
	FromHandle shared.ChatHandle `json:"from_handle"`
	// Idempotency key for the message. Used for deduplication of outbound messages.
	IdempotencyKey string `json:"idempotency_key" api:"nullable"`
	// Whether the message was sent by us (true for sent events, false for received
	// events)
	IsFromMe bool `json:"is_from_me"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group"`
	// Message content nested within webhook events
	Message MessagePayload `json:"message"`
	// Message identifier (UUID)
	MessageID string `json:"message_id"`
	// Preferred messaging service type. Includes "auto" for default fallback behavior.
	//
	// Any of "iMessage", "SMS", "RCS", "auto".
	PreferredService string `json:"preferred_service" api:"nullable"`
	// When the message was received. Null for sent events.
	ReceivedAt time.Time `json:"received_at" api:"nullable" format:"date-time"`
	// Our phone number that received the message as a full handle object. Null for
	// sent events.
	RecipientHandle shared.ChatHandle `json:"recipient_handle" api:"nullable"`
	// DEPRECATED: Use recipient_handle instead. Our phone number that received the
	// message. Null for sent events.
	//
	// Deprecated: deprecated
	RecipientPhone string `json:"recipient_phone" api:"nullable"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReadAt           respjson.Field
		ChatID           respjson.Field
		From             respjson.Field
		FromHandle       respjson.Field
		IdempotencyKey   respjson.Field
		IsFromMe         respjson.Field
		IsGroup          respjson.Field
		Message          respjson.Field
		MessageID        respjson.Field
		PreferredService respjson.Field
		ReceivedAt       respjson.Field
		RecipientHandle  respjson.Field
		RecipientPhone   respjson.Field
		Service          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageReadV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *MessageReadV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.delivered events (2025-01-01 format)
type MessageDeliveredV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for message.delivered webhook events (2025-01-01 format). Extends
	// MessageEvent with delivered_at and message_id.
	Data MessageDeliveredV2025WebhookEventData `json:"data" api:"required"`
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
func (r MessageDeliveredV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageDeliveredV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for message.delivered webhook events (2025-01-01 format). Extends
// MessageEvent with delivered_at and message_id.
type MessageDeliveredV2025WebhookEventData struct {
	// When the message was delivered to the recipient's device
	DeliveredAt time.Time `json:"delivered_at" api:"required" format:"date-time"`
	// Chat identifier
	ChatID string `json:"chat_id" format:"uuid"`
	// DEPRECATED: Use from_handle instead. Phone number or email address of the
	// message sender.
	//
	// Deprecated: deprecated
	From string `json:"from"`
	// The sender of this message as a full handle object
	FromHandle shared.ChatHandle `json:"from_handle"`
	// Idempotency key for the message. Used for deduplication of outbound messages.
	IdempotencyKey string `json:"idempotency_key" api:"nullable"`
	// Whether the message was sent by us (true for sent events, false for received
	// events)
	IsFromMe bool `json:"is_from_me"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group"`
	// Message content nested within webhook events
	Message MessagePayload `json:"message"`
	// Message identifier (UUID)
	MessageID string `json:"message_id"`
	// Preferred messaging service type. Includes "auto" for default fallback behavior.
	//
	// Any of "iMessage", "SMS", "RCS", "auto".
	PreferredService string `json:"preferred_service" api:"nullable"`
	// When the message was received. Null for sent events.
	ReceivedAt time.Time `json:"received_at" api:"nullable" format:"date-time"`
	// Our phone number that received the message as a full handle object. Null for
	// sent events.
	RecipientHandle shared.ChatHandle `json:"recipient_handle" api:"nullable"`
	// DEPRECATED: Use recipient_handle instead. Our phone number that received the
	// message. Null for sent events.
	//
	// Deprecated: deprecated
	RecipientPhone string `json:"recipient_phone" api:"nullable"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeliveredAt      respjson.Field
		ChatID           respjson.Field
		From             respjson.Field
		FromHandle       respjson.Field
		IdempotencyKey   respjson.Field
		IsFromMe         respjson.Field
		IsGroup          respjson.Field
		Message          respjson.Field
		MessageID        respjson.Field
		PreferredService respjson.Field
		ReceivedAt       respjson.Field
		RecipientHandle  respjson.Field
		RecipientPhone   respjson.Field
		Service          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageDeliveredV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *MessageDeliveredV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for message.failed events
type MessageFailedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error details for message.failed webhook events. See
	// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
	// code reference.
	Data MessageFailedV2025WebhookEventData `json:"data" api:"required"`
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
func (r MessageFailedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageFailedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for message.failed webhook events. See
// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
// code reference.
type MessageFailedV2025WebhookEventData struct {
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
func (r MessageFailedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *MessageFailedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for reaction.added events
type ReactionAddedV2025WebhookEvent struct {
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
func (r ReactionAddedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ReactionAddedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for reaction.removed events
type ReactionRemovedV2025WebhookEvent struct {
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
func (r ReactionRemovedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ReactionRemovedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for participant.added events
type ParticipantAddedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for participant.added webhook events
	Data ParticipantAddedV2025WebhookEventData `json:"data" api:"required"`
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
func (r ParticipantAddedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ParticipantAddedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for participant.added webhook events
type ParticipantAddedV2025WebhookEventData struct {
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
func (r ParticipantAddedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ParticipantAddedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for participant.removed events
type ParticipantRemovedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for participant.removed webhook events
	Data ParticipantRemovedV2025WebhookEventData `json:"data" api:"required"`
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
func (r ParticipantRemovedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ParticipantRemovedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for participant.removed webhook events
type ParticipantRemovedV2025WebhookEventData struct {
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
func (r ParticipantRemovedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ParticipantRemovedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_name_updated events
type ChatGroupNameUpdatedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.group_name_updated webhook events
	Data ChatGroupNameUpdatedV2025WebhookEventData `json:"data" api:"required"`
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
func (r ChatGroupNameUpdatedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdatedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.group_name_updated webhook events
type ChatGroupNameUpdatedV2025WebhookEventData struct {
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
func (r ChatGroupNameUpdatedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdatedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_icon_updated events
type ChatGroupIconUpdatedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.group_icon_updated webhook events
	Data ChatGroupIconUpdatedV2025WebhookEventData `json:"data" api:"required"`
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
func (r ChatGroupIconUpdatedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdatedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.group_icon_updated webhook events
type ChatGroupIconUpdatedV2025WebhookEventData struct {
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
func (r ChatGroupIconUpdatedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdatedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_name_update_failed events
type ChatGroupNameUpdateFailedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error details for chat.group_name_update_failed webhook events. See
	// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
	// code reference.
	Data ChatGroupNameUpdateFailedV2025WebhookEventData `json:"data" api:"required"`
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
func (r ChatGroupNameUpdateFailedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdateFailedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for chat.group_name_update_failed webhook events. See
// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
// code reference.
type ChatGroupNameUpdateFailedV2025WebhookEventData struct {
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
func (r ChatGroupNameUpdateFailedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupNameUpdateFailedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.group_icon_update_failed events
type ChatGroupIconUpdateFailedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error details for chat.group_icon_update_failed webhook events. See
	// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
	// code reference.
	Data ChatGroupIconUpdateFailedV2025WebhookEventData `json:"data" api:"required"`
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
func (r ChatGroupIconUpdateFailedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdateFailedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for chat.group_icon_update_failed webhook events. See
// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
// code reference.
type ChatGroupIconUpdateFailedV2025WebhookEventData struct {
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
func (r ChatGroupIconUpdateFailedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatGroupIconUpdateFailedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.created events
type ChatCreatedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.created webhook events. Matches GET /v3/chats/{chatId}
	// response.
	Data ChatCreatedV2025WebhookEventData `json:"data" api:"required"`
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
func (r ChatCreatedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatCreatedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.created webhook events. Matches GET /v3/chats/{chatId}
// response.
type ChatCreatedV2025WebhookEventData struct {
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
func (r ChatCreatedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatCreatedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.typing_indicator.started events
type ChatTypingIndicatorStartedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.typing_indicator.started webhook events
	Data ChatTypingIndicatorStartedV2025WebhookEventData `json:"data" api:"required"`
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
func (r ChatTypingIndicatorStartedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStartedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.typing_indicator.started webhook events
type ChatTypingIndicatorStartedV2025WebhookEventData struct {
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
func (r ChatTypingIndicatorStartedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStartedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.typing_indicator.stopped events
type ChatTypingIndicatorStoppedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.typing_indicator.stopped webhook events
	Data ChatTypingIndicatorStoppedV2025WebhookEventData `json:"data" api:"required"`
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
func (r ChatTypingIndicatorStoppedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStoppedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.typing_indicator.stopped webhook events
type ChatTypingIndicatorStoppedV2025WebhookEventData struct {
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
func (r ChatTypingIndicatorStoppedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatTypingIndicatorStoppedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for phone_number.status_updated events
type PhoneNumberStatusUpdatedV2025WebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for phone_number.status_updated webhook events
	Data PhoneNumberStatusUpdatedV2025WebhookEventData `json:"data" api:"required"`
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
	EventType PhoneNumberStatusUpdatedV2025WebhookEventEventType `json:"event_type" api:"required"`
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
func (r PhoneNumberStatusUpdatedV2025WebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberStatusUpdatedV2025WebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for phone_number.status_updated webhook events
type PhoneNumberStatusUpdatedV2025WebhookEventData struct {
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
func (r PhoneNumberStatusUpdatedV2025WebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberStatusUpdatedV2025WebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event
type PhoneNumberStatusUpdatedV2025WebhookEventEventType string

const (
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeMessageSent                PhoneNumberStatusUpdatedV2025WebhookEventEventType = "message.sent"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeMessageReceived            PhoneNumberStatusUpdatedV2025WebhookEventEventType = "message.received"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeMessageRead                PhoneNumberStatusUpdatedV2025WebhookEventEventType = "message.read"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeMessageDelivered           PhoneNumberStatusUpdatedV2025WebhookEventEventType = "message.delivered"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeMessageFailed              PhoneNumberStatusUpdatedV2025WebhookEventEventType = "message.failed"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeMessageEdited              PhoneNumberStatusUpdatedV2025WebhookEventEventType = "message.edited"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeReactionAdded              PhoneNumberStatusUpdatedV2025WebhookEventEventType = "reaction.added"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeReactionRemoved            PhoneNumberStatusUpdatedV2025WebhookEventEventType = "reaction.removed"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeParticipantAdded           PhoneNumberStatusUpdatedV2025WebhookEventEventType = "participant.added"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeParticipantRemoved         PhoneNumberStatusUpdatedV2025WebhookEventEventType = "participant.removed"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeChatCreated                PhoneNumberStatusUpdatedV2025WebhookEventEventType = "chat.created"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeChatGroupNameUpdated       PhoneNumberStatusUpdatedV2025WebhookEventEventType = "chat.group_name_updated"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeChatGroupIconUpdated       PhoneNumberStatusUpdatedV2025WebhookEventEventType = "chat.group_icon_updated"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeChatGroupNameUpdateFailed  PhoneNumberStatusUpdatedV2025WebhookEventEventType = "chat.group_name_update_failed"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeChatGroupIconUpdateFailed  PhoneNumberStatusUpdatedV2025WebhookEventEventType = "chat.group_icon_update_failed"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeChatTypingIndicatorStarted PhoneNumberStatusUpdatedV2025WebhookEventEventType = "chat.typing_indicator.started"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeChatTypingIndicatorStopped PhoneNumberStatusUpdatedV2025WebhookEventEventType = "chat.typing_indicator.stopped"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypePhoneNumberStatusUpdated   PhoneNumberStatusUpdatedV2025WebhookEventEventType = "phone_number.status_updated"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeCallInitiated              PhoneNumberStatusUpdatedV2025WebhookEventEventType = "call.initiated"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeCallRinging                PhoneNumberStatusUpdatedV2025WebhookEventEventType = "call.ringing"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeCallAnswered               PhoneNumberStatusUpdatedV2025WebhookEventEventType = "call.answered"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeCallEnded                  PhoneNumberStatusUpdatedV2025WebhookEventEventType = "call.ended"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeCallFailed                 PhoneNumberStatusUpdatedV2025WebhookEventEventType = "call.failed"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeCallDeclined               PhoneNumberStatusUpdatedV2025WebhookEventEventType = "call.declined"
	PhoneNumberStatusUpdatedV2025WebhookEventEventTypeCallNoAnswer               PhoneNumberStatusUpdatedV2025WebhookEventEventType = "call.no_answer"
)

// EventsWebhookEventUnion contains all possible properties and values from
// [MessageSentV2026WebhookEvent], [MessageReceivedV2026WebhookEvent],
// [MessageReadV2026WebhookEvent], [MessageDeliveredV2026WebhookEvent],
// [MessageFailedV2026WebhookEvent], [ReactionAddedV2026WebhookEvent],
// [ReactionRemovedV2026WebhookEvent], [ParticipantAddedV2026WebhookEvent],
// [ParticipantRemovedV2026WebhookEvent], [ChatGroupNameUpdatedV2026WebhookEvent],
// [ChatGroupIconUpdatedV2026WebhookEvent],
// [ChatGroupNameUpdateFailedV2026WebhookEvent],
// [ChatGroupIconUpdateFailedV2026WebhookEvent], [ChatCreatedV2026WebhookEvent],
// [ChatTypingIndicatorStartedV2026WebhookEvent],
// [ChatTypingIndicatorStoppedV2026WebhookEvent], [MessageEditedV2026WebhookEvent],
// [PhoneNumberStatusUpdatedV2026WebhookEvent], [MessageSentV2025WebhookEvent],
// [MessageReceivedV2025WebhookEvent], [MessageReadV2025WebhookEvent],
// [MessageDeliveredV2025WebhookEvent], [MessageFailedV2025WebhookEvent],
// [ReactionAddedV2025WebhookEvent], [ReactionRemovedV2025WebhookEvent],
// [ParticipantAddedV2025WebhookEvent], [ParticipantRemovedV2025WebhookEvent],
// [ChatGroupNameUpdatedV2025WebhookEvent],
// [ChatGroupIconUpdatedV2025WebhookEvent],
// [ChatGroupNameUpdateFailedV2025WebhookEvent],
// [ChatGroupIconUpdateFailedV2025WebhookEvent], [ChatCreatedV2025WebhookEvent],
// [ChatTypingIndicatorStartedV2025WebhookEvent],
// [ChatTypingIndicatorStoppedV2025WebhookEvent],
// [PhoneNumberStatusUpdatedV2025WebhookEvent].
//
// Use the [EventsWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EventsWebhookEventUnion struct {
	APIVersion string    `json:"api_version"`
	CreatedAt  time.Time `json:"created_at"`
	// This field is a union of [MessageEventV2], [MessageFailedV2026WebhookEventData],
	// [ReactionEventBase], [ParticipantAddedV2026WebhookEventData],
	// [ParticipantRemovedV2026WebhookEventData],
	// [ChatGroupNameUpdatedV2026WebhookEventData],
	// [ChatGroupIconUpdatedV2026WebhookEventData],
	// [ChatGroupNameUpdateFailedV2026WebhookEventData],
	// [ChatGroupIconUpdateFailedV2026WebhookEventData],
	// [ChatCreatedV2026WebhookEventData],
	// [ChatTypingIndicatorStartedV2026WebhookEventData],
	// [ChatTypingIndicatorStoppedV2026WebhookEventData],
	// [MessageEditedV2026WebhookEventData],
	// [PhoneNumberStatusUpdatedV2026WebhookEventData],
	// [MessageSentV2025WebhookEventData], [MessageReceivedV2025WebhookEventData],
	// [MessageReadV2025WebhookEventData], [MessageDeliveredV2025WebhookEventData],
	// [MessageFailedV2025WebhookEventData], [ParticipantAddedV2025WebhookEventData],
	// [ParticipantRemovedV2025WebhookEventData],
	// [ChatGroupNameUpdatedV2025WebhookEventData],
	// [ChatGroupIconUpdatedV2025WebhookEventData],
	// [ChatGroupNameUpdateFailedV2025WebhookEventData],
	// [ChatGroupIconUpdateFailedV2025WebhookEventData],
	// [ChatCreatedV2025WebhookEventData],
	// [ChatTypingIndicatorStartedV2025WebhookEventData],
	// [ChatTypingIndicatorStoppedV2025WebhookEventData],
	// [PhoneNumberStatusUpdatedV2025WebhookEventData]
	Data    EventsWebhookEventUnionData `json:"data"`
	EventID string                      `json:"event_id"`
	// Any of nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
	// nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
	// nil, nil, nil, nil, nil.
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

func (u EventsWebhookEventUnion) AsMessageSentV2026WebhookEvent() (v MessageSentV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageReceivedV2026WebhookEvent() (v MessageReceivedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageReadV2026WebhookEvent() (v MessageReadV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageDeliveredV2026WebhookEvent() (v MessageDeliveredV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageFailedV2026WebhookEvent() (v MessageFailedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsReactionAddedV2026WebhookEvent() (v ReactionAddedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsReactionRemovedV2026WebhookEvent() (v ReactionRemovedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsParticipantAddedV2026WebhookEvent() (v ParticipantAddedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsParticipantRemovedV2026WebhookEvent() (v ParticipantRemovedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupNameUpdatedV2026WebhookEvent() (v ChatGroupNameUpdatedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupIconUpdatedV2026WebhookEvent() (v ChatGroupIconUpdatedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupNameUpdateFailedV2026WebhookEvent() (v ChatGroupNameUpdateFailedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupIconUpdateFailedV2026WebhookEvent() (v ChatGroupIconUpdateFailedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatCreatedV2026WebhookEvent() (v ChatCreatedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatTypingIndicatorStartedV2026WebhookEvent() (v ChatTypingIndicatorStartedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatTypingIndicatorStoppedV2026WebhookEvent() (v ChatTypingIndicatorStoppedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageEditedV2026WebhookEvent() (v MessageEditedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsPhoneNumberStatusUpdatedV2026WebhookEvent() (v PhoneNumberStatusUpdatedV2026WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageSentV2025WebhookEvent() (v MessageSentV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageReceivedV2025WebhookEvent() (v MessageReceivedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageReadV2025WebhookEvent() (v MessageReadV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageDeliveredV2025WebhookEvent() (v MessageDeliveredV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsMessageFailedV2025WebhookEvent() (v MessageFailedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsReactionAddedV2025WebhookEvent() (v ReactionAddedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsReactionRemovedV2025WebhookEvent() (v ReactionRemovedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsParticipantAddedV2025WebhookEvent() (v ParticipantAddedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsParticipantRemovedV2025WebhookEvent() (v ParticipantRemovedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupNameUpdatedV2025WebhookEvent() (v ChatGroupNameUpdatedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupIconUpdatedV2025WebhookEvent() (v ChatGroupIconUpdatedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupNameUpdateFailedV2025WebhookEvent() (v ChatGroupNameUpdateFailedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatGroupIconUpdateFailedV2025WebhookEvent() (v ChatGroupIconUpdateFailedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatCreatedV2025WebhookEvent() (v ChatCreatedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatTypingIndicatorStartedV2025WebhookEvent() (v ChatTypingIndicatorStartedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsChatTypingIndicatorStoppedV2025WebhookEvent() (v ChatTypingIndicatorStoppedV2025WebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsWebhookEventUnion) AsPhoneNumberStatusUpdatedV2025WebhookEvent() (v PhoneNumberStatusUpdatedV2025WebhookEvent) {
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
	// [MessageEditedV2026WebhookEventDataChat]
	Chat      EventsWebhookEventUnionDataChat `json:"chat"`
	Direction string                          `json:"direction"`
	// This field is from variant [MessageEventV2].
	Parts []MessageEventV2PartUnion `json:"parts"`
	// This field is from variant [MessageEventV2].
	SenderHandle shared.ChatHandle `json:"sender_handle"`
	// This field is from variant [MessageEventV2].
	Service     shared.ServiceType `json:"service"`
	DeliveredAt time.Time          `json:"delivered_at"`
	// This field is from variant [MessageEventV2].
	Effect           SchemasMessageEffect `json:"effect"`
	IdempotencyKey   string               `json:"idempotency_key"`
	PreferredService string               `json:"preferred_service"`
	ReadAt           time.Time            `json:"read_at"`
	// This field is from variant [MessageEventV2].
	ReplyTo MessageEventV2ReplyTo `json:"reply_to"`
	// This field is from variant [MessageEventV2].
	SentAt    time.Time `json:"sent_at"`
	Code      int64     `json:"code"`
	FailedAt  time.Time `json:"failed_at"`
	ChatID    string    `json:"chat_id"`
	MessageID string    `json:"message_id"`
	Reason    string    `json:"reason"`
	IsFromMe  bool      `json:"is_from_me"`
	// This field is from variant [ReactionEventBase].
	ReactionType shared.ReactionType `json:"reaction_type"`
	// This field is from variant [ReactionEventBase].
	CustomEmoji string `json:"custom_emoji"`
	From        string `json:"from"`
	// This field is from variant [ReactionEventBase].
	FromHandle shared.ChatHandle `json:"from_handle"`
	// This field is from variant [ReactionEventBase].
	PartIndex int64 `json:"part_index"`
	// This field is from variant [ReactionEventBase].
	ReactedAt time.Time `json:"reacted_at"`
	// This field is from variant [ReactionEventBase].
	Sticker ReactionEventBaseSticker `json:"sticker"`
	Handle  string                   `json:"handle"`
	AddedAt time.Time                `json:"added_at"`
	// This field is from variant [ParticipantAddedV2026WebhookEventData].
	Participant shared.ChatHandle `json:"participant"`
	RemovedAt   time.Time         `json:"removed_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	// This field is from variant [ChatGroupNameUpdatedV2026WebhookEventData].
	ChangedByHandle shared.ChatHandle   `json:"changed_by_handle"`
	NewValue        string              `json:"new_value"`
	OldValue        string              `json:"old_value"`
	ErrorCode       int64               `json:"error_code"`
	CreatedAt       time.Time           `json:"created_at"`
	DisplayName     string              `json:"display_name"`
	Handles         []shared.ChatHandle `json:"handles"`
	IsGroup         bool                `json:"is_group"`
	// This field is from variant [MessageEditedV2026WebhookEventData].
	EditedAt time.Time `json:"edited_at"`
	// This field is from variant [MessageEditedV2026WebhookEventData].
	Part           MessageEditedV2026WebhookEventDataPart `json:"part"`
	ChangedAt      time.Time                              `json:"changed_at"`
	NewStatus      string                                 `json:"new_status"`
	PhoneNumber    string                                 `json:"phone_number"`
	PreviousStatus string                                 `json:"previous_status"`
	// This field is from variant [MessageSentV2025WebhookEventData].
	Message    MessagePayload `json:"message"`
	ReceivedAt time.Time      `json:"received_at"`
	// This field is from variant [MessageSentV2025WebhookEventData].
	RecipientHandle shared.ChatHandle `json:"recipient_handle"`
	RecipientPhone  string            `json:"recipient_phone"`
	JSON            struct {
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
		UpdatedAt        respjson.Field
		ChangedByHandle  respjson.Field
		NewValue         respjson.Field
		OldValue         respjson.Field
		ErrorCode        respjson.Field
		CreatedAt        respjson.Field
		DisplayName      respjson.Field
		Handles          respjson.Field
		IsGroup          respjson.Field
		EditedAt         respjson.Field
		Part             respjson.Field
		ChangedAt        respjson.Field
		NewStatus        respjson.Field
		PhoneNumber      respjson.Field
		PreviousStatus   respjson.Field
		Message          respjson.Field
		ReceivedAt       respjson.Field
		RecipientHandle  respjson.Field
		RecipientPhone   respjson.Field
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
