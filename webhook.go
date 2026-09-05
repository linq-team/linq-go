// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"time"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/respjson"
	"github.com/linq-team/linq-go/shared"
	"github.com/linq-team/linq-go/shared/constant"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
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

func (r *WebhookService) Unwrap(payload []byte, headers http.Header, opts ...option.RequestOption) (*UnwrapWebhookEventUnion, error) {
	opts = slices.Concat(r.Options, opts)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	key := cfg.WebhookSecret
	if key == "" {
		return nil, errors.New("The WebhookSecret option must be set in order to verify webhook headers")
	}
	wh, err := standardwebhooks.NewWebhook(key)
	if err != nil {
		return nil, err
	}
	err = wh.Verify(payload, headers)
	if err != nil {
		return nil, err
	}
	res := &UnwrapWebhookEventUnion{}
	err = res.UnmarshalJSON(payload)
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
	// Present only when this message was recovered by reconciliation rather than
	// delivered live, and set to the time of that recovery. The field is omitted
	// entirely for normally-delivered messages, which is the overwhelming majority.
	// When present, expect `sent_at` to be substantially earlier than delivery of this
	// event: the message is genuine but is arriving late and out of real-time order,
	// so treat it as history rather than as a live inbound (for example, suppress
	// auto-replies).
	ReconciledAt time.Time `json:"reconciled_at" format:"date-time"`
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
		ReconciledAt     respjson.Field
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
	// **[BETA]** Current health for a chat. Always present — chats start at `HEALTHY`
	// and may shift based on engagement and delivery signals on the conversation. Many
	// `AT_RISK` or `CRITICAL` chats on a single line increase the risk of line
	// flagging.
	//
	// Switch on `status` to surface chat and line health in your UI — the enum is the
	// long-term contract. Each status carries a `doc_url` that deep-links to the
	// relevant section of the Chat Health guide. To gate a send, act on the response
	// rather than the status: a `403` is the authoritative answer.
	//
	// See the [Chat Health guide](/channel/imessage/guides/chats/chat-health) for what
	// each status means and how to react.
	HealthStatus MessageEventV2ChatHealthStatus `json:"health_status" api:"required"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group" api:"nullable"`
	// Your phone number's handle. Always has is_me=true.
	OwnerHandle shared.ChatHandle `json:"owner_handle" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		HealthStatus respjson.Field
		IsGroup      respjson.Field
		OwnerHandle  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEventV2Chat) RawJSON() string { return r.JSON.raw }
func (r *MessageEventV2Chat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **[BETA]** Current health for a chat. Always present — chats start at `HEALTHY`
// and may shift based on engagement and delivery signals on the conversation. Many
// `AT_RISK` or `CRITICAL` chats on a single line increase the risk of line
// flagging.
//
// Switch on `status` to surface chat and line health in your UI — the enum is the
// long-term contract. Each status carries a `doc_url` that deep-links to the
// relevant section of the Chat Health guide. To gate a send, act on the response
// rather than the status: a `403` is the authoritative answer.
//
// See the [Chat Health guide](/channel/imessage/guides/chats/chat-health) for what
// each status means and how to react.
type MessageEventV2ChatHealthStatus struct {
	// Deep-link to the relevant section of the Chat Health guide for this status.
	DocURL string `json:"doc_url" api:"required" format:"uri"`
	// Current health bucket for the chat. See the
	// [Chat Health guide](/channel/imessage/guides/chats/chat-health) for what each
	// value means and how to react. `doc_url` deep-links to the relevant section.
	//
	// `OPTED_OUT` — the recipient sent `STOP`, `UNSUBSCRIBE`, `OPTOUT`, `CANCEL`,
	// `END`, or `QUIT`. The keyword must be the whole trimmed message, never part of a
	// longer one: `STOP` counts, `please stop` does not. Most keywords must match
	// exactly, including case. `OPT OUT` is the exception — it matches in any casing,
	// with or without the space or a hyphen, so `opt out`, `Opt-Out` and `optout` all
	// count. It clears as soon as they reply again: any later message from them that
	// is not itself an opt-out keyword opts them back in immediately — a reply in any
	// conversation with you counts, the same way the block does.
	//
	// `OPTED_OUT` marks only the conversation the keyword arrived in. The block below
	// is wider than the mark, so a conversation still reading `HEALTHY` can be blocked
	// as well — gate on the `403`, not on the status. Group threads are never marked
	// and are never blocked.
	//
	// Linq enforces this: while a recipient is opted out, every send to them is
	// rejected with `403` (error code `2024`) before the message is queued, across
	// every chat and every line on your account. Nothing is delivered, including a
	// final courtesy message — to send one, set `override_optout: true` on that single
	// request.
	//
	// Any of "HEALTHY", "AT_RISK", "CRITICAL", "OPTED_OUT".
	Status string `json:"status" api:"required"`
	// When this status last changed.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocURL      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEventV2ChatHealthStatus) RawJSON() string { return r.JSON.raw }
func (r *MessageEventV2ChatHealthStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message direction - "outbound" if sent by you, "inbound" if received
type MessageEventV2Direction string

const (
	MessageEventV2DirectionInbound  MessageEventV2Direction = "inbound"
	MessageEventV2DirectionOutbound MessageEventV2Direction = "outbound"
)

// MessageEventV2PartUnion contains all possible properties and values from
// [SchemasTextPartResponse], [SchemasMediaPartResponse], [MessageEventV2PartLink],
// [MessageEventV2PartIMessageApp], [MessageEventV2PartAppClip].
//
// Use the [MessageEventV2PartUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessageEventV2PartUnion struct {
	// Any of "text", "media", "link", "imessage_app", "app_clip".
	Type  string `json:"type"`
	Value string `json:"value"`
	// This field is from variant [SchemasTextPartResponse].
	Mention string `json:"mention"`
	// This field is from variant [SchemasTextPartResponse].
	MentionRange []int64 `json:"mention_range"`
	// This field is from variant [SchemasTextPartResponse].
	Mentions []SchemasTextPartResponseMention `json:"mentions"`
	// This field is from variant [SchemasTextPartResponse].
	TextDecorations []shared.TextDecoration `json:"text_decorations"`
	// This field is from variant [SchemasMediaPartResponse].
	ID string `json:"id"`
	// This field is from variant [SchemasMediaPartResponse].
	Filename string `json:"filename"`
	// This field is from variant [SchemasMediaPartResponse].
	MimeType string `json:"mime_type"`
	// This field is from variant [SchemasMediaPartResponse].
	SizeBytes int64  `json:"size_bytes"`
	URL       string `json:"url"`
	// This field is from variant [MessageEventV2PartIMessageApp].
	App MessageEventV2PartIMessageAppApp `json:"app"`
	// This field is from variant [MessageEventV2PartIMessageApp].
	Layout MessageEventV2PartIMessageAppLayout `json:"layout"`
	// This field is from variant [MessageEventV2PartIMessageApp].
	FallbackText string `json:"fallback_text"`
	// This field is from variant [MessageEventV2PartAppClip].
	Description string `json:"description"`
	// This field is from variant [MessageEventV2PartAppClip].
	ImageURL string `json:"image_url"`
	// This field is from variant [MessageEventV2PartAppClip].
	Title string `json:"title"`
	JSON  struct {
		Type            respjson.Field
		Value           respjson.Field
		Mention         respjson.Field
		MentionRange    respjson.Field
		Mentions        respjson.Field
		TextDecorations respjson.Field
		ID              respjson.Field
		Filename        respjson.Field
		MimeType        respjson.Field
		SizeBytes       respjson.Field
		URL             respjson.Field
		App             respjson.Field
		Layout          respjson.Field
		FallbackText    respjson.Field
		Description     respjson.Field
		ImageURL        respjson.Field
		Title           respjson.Field
		raw             string
	} `json:"-"`
}

// anyMessageEventV2Part is implemented by each variant of
// [MessageEventV2PartUnion] to add type safety for the return type of
// [MessageEventV2PartUnion.AsAny]
type anyMessageEventV2Part interface {
	implMessageEventV2PartUnion()
}

func (SchemasTextPartResponse) implMessageEventV2PartUnion()       {}
func (SchemasMediaPartResponse) implMessageEventV2PartUnion()      {}
func (MessageEventV2PartLink) implMessageEventV2PartUnion()        {}
func (MessageEventV2PartIMessageApp) implMessageEventV2PartUnion() {}
func (MessageEventV2PartAppClip) implMessageEventV2PartUnion()     {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MessageEventV2PartUnion.AsAny().(type) {
//	case linqgo.SchemasTextPartResponse:
//	case linqgo.SchemasMediaPartResponse:
//	case linqgo.MessageEventV2PartLink:
//	case linqgo.MessageEventV2PartIMessageApp:
//	case linqgo.MessageEventV2PartAppClip:
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
	case "imessage_app":
		return u.AsIMessageApp()
	case "app_clip":
		return u.AsAppClip()
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

func (u MessageEventV2PartUnion) AsIMessageApp() (v MessageEventV2PartIMessageApp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageEventV2PartUnion) AsAppClip() (v MessageEventV2PartAppClip) {
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
	Type constant.Link `json:"type" default:"link"`
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

// An iMessage app card part.
type MessageEventV2PartIMessageApp struct {
	// Identifies the iMessage app (Messages app extension) that backs the card.
	App MessageEventV2PartIMessageAppApp `json:"app" api:"required"`
	// Visible layout of the card.
	Layout MessageEventV2PartIMessageAppLayout `json:"layout" api:"required"`
	// Indicates this is an iMessage app card part.
	Type constant.IMessageApp `json:"type" default:"imessage_app"`
	// The URL the recipient's app opens when the user taps the card.
	URL string `json:"url" api:"required" format:"uri"`
	// Fallback text for surfaces that cannot render the card.
	FallbackText string `json:"fallback_text" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		App          respjson.Field
		Layout       respjson.Field
		Type         respjson.Field
		URL          respjson.Field
		FallbackText respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEventV2PartIMessageApp) RawJSON() string { return r.JSON.raw }
func (r *MessageEventV2PartIMessageApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the iMessage app (Messages app extension) that backs the card.
type MessageEventV2PartIMessageAppApp struct {
	// Bundle identifier of the Messages app extension.
	BundleID string `json:"bundle_id" api:"required"`
	// Display name of the app.
	Name string `json:"name" api:"required"`
	// The app's 10-character team identifier.
	TeamID string `json:"team_id" api:"required"`
	// The owning app's App Store id, when known.
	AppStoreID int64 `json:"app_store_id" api:"nullable"`
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
func (r MessageEventV2PartIMessageAppApp) RawJSON() string { return r.JSON.raw }
func (r *MessageEventV2PartIMessageAppApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visible layout of the card.
type MessageEventV2PartIMessageAppLayout struct {
	// Primary label, top-left and bold.
	Caption string `json:"caption" api:"nullable"`
	// Secondary label, below caption on the left.
	Subcaption string `json:"subcaption" api:"nullable"`
	// Label shown top-right.
	TrailingCaption string `json:"trailing_caption" api:"nullable"`
	// Label shown below trailing_caption.
	TrailingSubcaption string `json:"trailing_subcaption" api:"nullable"`
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
func (r MessageEventV2PartIMessageAppLayout) RawJSON() string { return r.JSON.raw }
func (r *MessageEventV2PartIMessageAppLayout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Apple Pay App Clip payment card part
type MessageEventV2PartAppClip struct {
	// Indicates this is an App Clip payment card part
	Type constant.AppClip `json:"type" default:"app_clip"`
	// The checkout link the card opens
	Value string `json:"value" api:"required"`
	// The card's summary line, composed by Linq from the checkout session
	Description string `json:"description"`
	// The card's preview image
	ImageURL string `json:"image_url"`
	// The card's headline, composed by Linq from the checkout session
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		Description respjson.Field
		ImageURL    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEventV2PartAppClip) RawJSON() string { return r.JSON.raw }
func (r *MessageEventV2PartAppClip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Preferred messaging service type. Includes "auto" for default fallback behavior.
type MessageEventV2PreferredService string

const (
	MessageEventV2PreferredServiceIMessage MessageEventV2PreferredService = "iMessage"
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
	// Identifier for this reaction. Pass it to
	// `PATCH /v3/messages/{messageId}/reactions/{reactionId}` to move a sticker.
	// Stickers stack, so this is what distinguishes one sticker from another on the
	// same message.
	ReactionID string `json:"reaction_id" format:"uuid"`
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
		ReactionID   respjson.Field
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
	// DEPRECATED: Use `mentions` instead. Handle (E.164 phone number or Apple ID
	// email) of the **first** mention on this part. A part may carry several mentions;
	// this field shows only the first in `value` order, so it cannot be used to
	// determine whether a given participant was mentioned. `null` when the part
	// carries no mention.
	//
	// Deprecated: deprecated
	Mention string `json:"mention" api:"nullable"`
	// DEPRECATED: Use `mentions[].range` instead. Character range `[start, end)` in
	// `value` highlighted as the **first** mention only. `null` when the range was
	// omitted (the whole `value` is highlighted) or the part carries no mention.
	// _Characters are measured as UTF-16 code units. Most characters count as 1; some
	// emoji count as 2._
	//
	// Deprecated: deprecated
	MentionRange []int64 `json:"mention_range" api:"nullable"`
	// Every mention on this part, in the order they appear in `value`. `null` when the
	// part carries no mention. A part can carry several mentions of different people —
	// check `is_me` to tell whether this line was one of them.
	//
	// Only iMessage carries mentions. On a received message this is populated when the
	// sender was on iMessage; SMS and RCS have no way to mark a mention, so a message
	// from an SMS or RCS participant arrives as plain text with `mentions` null, even
	// in a group where other participants are on iMessage.
	Mentions []SchemasTextPartResponseMention `json:"mentions" api:"nullable"`
	// Text decorations applied to character ranges in the value
	TextDecorations []shared.TextDecoration `json:"text_decorations" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		Value           respjson.Field
		Mention         respjson.Field
		MentionRange    respjson.Field
		Mentions        respjson.Field
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

// One mention on a text part — who was mentioned, and which characters of `value`
// are the mention. A part carries one of these per mention, in the order they
// appear in the text, so a message naming two people has two entries.
type SchemasTextPartResponseMention struct {
	// Address of the mentioned participant, exactly as the device recorded it — an
	// E.164 phone number or an email address.
	Handle string `json:"handle" api:"required"`
	// Whether the mentioned participant is this line.
	IsMe bool `json:"is_me" api:"required"`
	// Character range `[start, end)` in `value` highlighted as this mention.
	// _Characters are measured as UTF-16 code units. Most characters count as 1; some
	// emoji count as 2._
	Range []int64 `json:"range" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		IsMe        respjson.Field
		Range       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchemasTextPartResponseMention) RawJSON() string { return r.JSON.raw }
func (r *SchemasTextPartResponseMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	//
	// In rare cases the message can still be delivered after this event fires — a
	// `message.delivered` webhook for the same message ID may follow.
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
//
// In rare cases the message can still be delivered after this event fires — a
// `message.delivered` webhook for the same message ID may follow.
type MessageFailedWebhookEventData struct {
	// Error codes in webhook failure events. The possible set varies by event:
	// message.failed and poll.failed can carry 3007, 4001, 4002, 4005, 4006, 4007, or
	// 4008; the group update failure events (chat.group_name_update_failed,
	// chat.group_icon_update_failed) carry 3007 or 4001; chat.background_update_failed
	// carries 1005, 2011, 4001, or 5002.
	Code int64 `json:"code" api:"required"`
	// When the failure was detected
	FailedAt time.Time `json:"failed_at" api:"required" format:"date-time"`
	// Chat identifier (UUID)
	ChatID string `json:"chat_id"`
	// Opaque diagnostic code identifying the specific failure class within `code`.
	// Values are not enumerated and may change without notice — log it and include it
	// in support requests, but do not branch on it.
	DetailCode int64 `json:"detail_code" api:"nullable"`
	// Message identifier (UUID)
	MessageID string `json:"message_id"`
	// Preferred messaging service type. Includes "auto" for default fallback behavior.
	//
	// Any of "iMessage", "SMS", "RCS", "auto".
	PreferredService string `json:"preferred_service" api:"nullable"`
	// Human-readable description of the failure
	Reason string `json:"reason"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code             respjson.Field
		FailedAt         respjson.Field
		ChatID           respjson.Field
		DetailCode       respjson.Field
		MessageID        respjson.Field
		PreferredService respjson.Field
		Reason           respjson.Field
		Service          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// **[BETA]** Current health for a chat. Always present — chats start at `HEALTHY`
	// and may shift based on engagement and delivery signals on the conversation. Many
	// `AT_RISK` or `CRITICAL` chats on a single line increase the risk of line
	// flagging.
	//
	// Switch on `status` to surface chat and line health in your UI — the enum is the
	// long-term contract. Each status carries a `doc_url` that deep-links to the
	// relevant section of the Chat Health guide. To gate a send, act on the response
	// rather than the status: a `403` is the authoritative answer.
	//
	// See the [Chat Health guide](/channel/imessage/guides/chats/chat-health) for what
	// each status means and how to react.
	HealthStatus MessageEditedWebhookEventDataChatHealthStatus `json:"health_status" api:"required"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group" api:"required"`
	// The handle that owns this chat (your phone number)
	OwnerHandle shared.ChatHandle `json:"owner_handle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		HealthStatus respjson.Field
		IsGroup      respjson.Field
		OwnerHandle  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEditedWebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *MessageEditedWebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **[BETA]** Current health for a chat. Always present — chats start at `HEALTHY`
// and may shift based on engagement and delivery signals on the conversation. Many
// `AT_RISK` or `CRITICAL` chats on a single line increase the risk of line
// flagging.
//
// Switch on `status` to surface chat and line health in your UI — the enum is the
// long-term contract. Each status carries a `doc_url` that deep-links to the
// relevant section of the Chat Health guide. To gate a send, act on the response
// rather than the status: a `403` is the authoritative answer.
//
// See the [Chat Health guide](/channel/imessage/guides/chats/chat-health) for what
// each status means and how to react.
type MessageEditedWebhookEventDataChatHealthStatus struct {
	// Deep-link to the relevant section of the Chat Health guide for this status.
	DocURL string `json:"doc_url" api:"required" format:"uri"`
	// Current health bucket for the chat. See the
	// [Chat Health guide](/channel/imessage/guides/chats/chat-health) for what each
	// value means and how to react. `doc_url` deep-links to the relevant section.
	//
	// `OPTED_OUT` — the recipient sent `STOP`, `UNSUBSCRIBE`, `OPTOUT`, `CANCEL`,
	// `END`, or `QUIT`. The keyword must be the whole trimmed message, never part of a
	// longer one: `STOP` counts, `please stop` does not. Most keywords must match
	// exactly, including case. `OPT OUT` is the exception — it matches in any casing,
	// with or without the space or a hyphen, so `opt out`, `Opt-Out` and `optout` all
	// count. It clears as soon as they reply again: any later message from them that
	// is not itself an opt-out keyword opts them back in immediately — a reply in any
	// conversation with you counts, the same way the block does.
	//
	// `OPTED_OUT` marks only the conversation the keyword arrived in. The block below
	// is wider than the mark, so a conversation still reading `HEALTHY` can be blocked
	// as well — gate on the `403`, not on the status. Group threads are never marked
	// and are never blocked.
	//
	// Linq enforces this: while a recipient is opted out, every send to them is
	// rejected with `403` (error code `2024`) before the message is queued, across
	// every chat and every line on your account. Nothing is delivered, including a
	// final courtesy message — to send one, set `override_optout: true` on that single
	// request.
	//
	// Any of "HEALTHY", "AT_RISK", "CRITICAL", "OPTED_OUT".
	Status string `json:"status" api:"required"`
	// When this status last changed.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocURL      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEditedWebhookEventDataChatHealthStatus) RawJSON() string { return r.JSON.raw }
func (r *MessageEditedWebhookEventDataChatHealthStatus) UnmarshalJSON(data []byte) error {
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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

// Complete webhook payload for poll.received events
type PollReceivedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for poll.received — a poll created by someone else and delivered to your
	// line. Carries the full poll snapshot (options, no voters yet) at receipt time.
	Data PollReceivedWebhookEventData `json:"data" api:"required"`
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r PollReceivedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PollReceivedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for poll.received — a poll created by someone else and delivered to your
// line. Carries the full poll snapshot (options, no voters yet) at receipt time.
type PollReceivedWebhookEventData struct {
	// Chat info for poll webhook events.
	Chat      PollReceivedWebhookEventDataChat `json:"chat" api:"required"`
	CreatedAt time.Time                        `json:"created_at" api:"required" format:"date-time"`
	// Any of "inbound", "outbound".
	Direction  string                           `json:"direction" api:"required"`
	MessageID  string                           `json:"message_id" api:"required" format:"uuid"`
	Poll       PollReceivedWebhookEventDataPoll `json:"poll" api:"required"`
	ReceivedAt time.Time                        `json:"received_at" api:"required" format:"date-time"`
	Service    string                           `json:"service" api:"required"`
	UpdatedAt  time.Time                        `json:"updated_at" api:"required" format:"date-time"`
	// The line that created the poll (is_me=false for an inbound poll).
	SenderHandle shared.ChatHandle `json:"sender_handle" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chat         respjson.Field
		CreatedAt    respjson.Field
		Direction    respjson.Field
		MessageID    respjson.Field
		Poll         respjson.Field
		ReceivedAt   respjson.Field
		Service      respjson.Field
		UpdatedAt    respjson.Field
		SenderHandle respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollReceivedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PollReceivedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat info for poll webhook events.
type PollReceivedWebhookEventDataChat struct {
	ID          string            `json:"id" api:"required" format:"uuid"`
	IsGroup     bool              `json:"is_group" api:"nullable"`
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
func (r PollReceivedWebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *PollReceivedWebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollReceivedWebhookEventDataPoll struct {
	Options []PollReceivedWebhookEventDataPollOption `json:"options" api:"required"`
	// Distinct participants across the whole poll.
	TotalVoters int64 `json:"total_voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		TotalVoters respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollReceivedWebhookEventDataPoll) RawJSON() string { return r.JSON.raw }
func (r *PollReceivedWebhookEventDataPoll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollReceivedWebhookEventDataPollOption struct {
	CanBeEdited bool `json:"can_be_edited" api:"required"`
	// The participant who added this option (poll creator for the initial options;
	// whoever added later ones). On a poll.updated this differs from the event's
	// `sender_handle` whenever a remote participant added the option. Null when
	// unknown.
	CreatorHandle shared.ChatHandle                             `json:"creator_handle" api:"required"`
	OptionID      string                                        `json:"option_id" api:"required" format:"uuid"`
	Text          string                                        `json:"text" api:"required"`
	Voters        []PollReceivedWebhookEventDataPollOptionVoter `json:"voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanBeEdited   respjson.Field
		CreatorHandle respjson.Field
		OptionID      respjson.Field
		Text          respjson.Field
		Voters        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollReceivedWebhookEventDataPollOption) RawJSON() string { return r.JSON.raw }
func (r *PollReceivedWebhookEventDataPollOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollReceivedWebhookEventDataPollOptionVoter struct {
	Handle  string    `json:"handle" api:"required"`
	VotedAt time.Time `json:"voted_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		VotedAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollReceivedWebhookEventDataPollOptionVoter) RawJSON() string { return r.JSON.raw }
func (r *PollReceivedWebhookEventDataPollOptionVoter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for poll.sent events
type PollSentWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for poll.sent, poll.delivered, and poll.read webhook events. Timestamps
	// indicate state (null = not yet happened): sent → sent_at; delivered →
	// +delivered_at; read → +read_at.
	Data PollSentWebhookEventData `json:"data" api:"required"`
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r PollSentWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PollSentWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for poll.sent, poll.delivered, and poll.read webhook events. Timestamps
// indicate state (null = not yet happened): sent → sent_at; delivered →
// +delivered_at; read → +read_at.
type PollSentWebhookEventData struct {
	// Chat info for poll webhook events.
	Chat      PollSentWebhookEventDataChat `json:"chat" api:"required"`
	CreatedAt time.Time                    `json:"created_at" api:"required" format:"date-time"`
	// Any of "inbound", "outbound".
	Direction   string                       `json:"direction" api:"required"`
	MessageID   string                       `json:"message_id" api:"required" format:"uuid"`
	Poll        PollSentWebhookEventDataPoll `json:"poll" api:"required"`
	Service     string                       `json:"service" api:"required"`
	UpdatedAt   time.Time                    `json:"updated_at" api:"required" format:"date-time"`
	DeliveredAt time.Time                    `json:"delivered_at" api:"nullable" format:"date-time"`
	ReadAt      time.Time                    `json:"read_at" api:"nullable" format:"date-time"`
	// The handle that sent the poll.
	SenderHandle shared.ChatHandle `json:"sender_handle" api:"nullable"`
	SentAt       time.Time         `json:"sent_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chat         respjson.Field
		CreatedAt    respjson.Field
		Direction    respjson.Field
		MessageID    respjson.Field
		Poll         respjson.Field
		Service      respjson.Field
		UpdatedAt    respjson.Field
		DeliveredAt  respjson.Field
		ReadAt       respjson.Field
		SenderHandle respjson.Field
		SentAt       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollSentWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PollSentWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat info for poll webhook events.
type PollSentWebhookEventDataChat struct {
	ID          string            `json:"id" api:"required" format:"uuid"`
	IsGroup     bool              `json:"is_group" api:"nullable"`
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
func (r PollSentWebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *PollSentWebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollSentWebhookEventDataPoll struct {
	Options []PollSentWebhookEventDataPollOption `json:"options" api:"required"`
	// Distinct participants across the whole poll.
	TotalVoters int64 `json:"total_voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		TotalVoters respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollSentWebhookEventDataPoll) RawJSON() string { return r.JSON.raw }
func (r *PollSentWebhookEventDataPoll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollSentWebhookEventDataPollOption struct {
	CanBeEdited bool `json:"can_be_edited" api:"required"`
	// The participant who added this option (poll creator for the initial options;
	// whoever added later ones). On a poll.updated this differs from the event's
	// `sender_handle` whenever a remote participant added the option. Null when
	// unknown.
	CreatorHandle shared.ChatHandle                         `json:"creator_handle" api:"required"`
	OptionID      string                                    `json:"option_id" api:"required" format:"uuid"`
	Text          string                                    `json:"text" api:"required"`
	Voters        []PollSentWebhookEventDataPollOptionVoter `json:"voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanBeEdited   respjson.Field
		CreatorHandle respjson.Field
		OptionID      respjson.Field
		Text          respjson.Field
		Voters        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollSentWebhookEventDataPollOption) RawJSON() string { return r.JSON.raw }
func (r *PollSentWebhookEventDataPollOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollSentWebhookEventDataPollOptionVoter struct {
	Handle  string    `json:"handle" api:"required"`
	VotedAt time.Time `json:"voted_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		VotedAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollSentWebhookEventDataPollOptionVoter) RawJSON() string { return r.JSON.raw }
func (r *PollSentWebhookEventDataPollOptionVoter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for poll.delivered events
type PollDeliveredWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for poll.sent, poll.delivered, and poll.read webhook events. Timestamps
	// indicate state (null = not yet happened): sent → sent_at; delivered →
	// +delivered_at; read → +read_at.
	Data PollDeliveredWebhookEventData `json:"data" api:"required"`
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r PollDeliveredWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PollDeliveredWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for poll.sent, poll.delivered, and poll.read webhook events. Timestamps
// indicate state (null = not yet happened): sent → sent_at; delivered →
// +delivered_at; read → +read_at.
type PollDeliveredWebhookEventData struct {
	// Chat info for poll webhook events.
	Chat      PollDeliveredWebhookEventDataChat `json:"chat" api:"required"`
	CreatedAt time.Time                         `json:"created_at" api:"required" format:"date-time"`
	// Any of "inbound", "outbound".
	Direction   string                            `json:"direction" api:"required"`
	MessageID   string                            `json:"message_id" api:"required" format:"uuid"`
	Poll        PollDeliveredWebhookEventDataPoll `json:"poll" api:"required"`
	Service     string                            `json:"service" api:"required"`
	UpdatedAt   time.Time                         `json:"updated_at" api:"required" format:"date-time"`
	DeliveredAt time.Time                         `json:"delivered_at" api:"nullable" format:"date-time"`
	ReadAt      time.Time                         `json:"read_at" api:"nullable" format:"date-time"`
	// The handle that sent the poll.
	SenderHandle shared.ChatHandle `json:"sender_handle" api:"nullable"`
	SentAt       time.Time         `json:"sent_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chat         respjson.Field
		CreatedAt    respjson.Field
		Direction    respjson.Field
		MessageID    respjson.Field
		Poll         respjson.Field
		Service      respjson.Field
		UpdatedAt    respjson.Field
		DeliveredAt  respjson.Field
		ReadAt       respjson.Field
		SenderHandle respjson.Field
		SentAt       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollDeliveredWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PollDeliveredWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat info for poll webhook events.
type PollDeliveredWebhookEventDataChat struct {
	ID          string            `json:"id" api:"required" format:"uuid"`
	IsGroup     bool              `json:"is_group" api:"nullable"`
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
func (r PollDeliveredWebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *PollDeliveredWebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollDeliveredWebhookEventDataPoll struct {
	Options []PollDeliveredWebhookEventDataPollOption `json:"options" api:"required"`
	// Distinct participants across the whole poll.
	TotalVoters int64 `json:"total_voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		TotalVoters respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollDeliveredWebhookEventDataPoll) RawJSON() string { return r.JSON.raw }
func (r *PollDeliveredWebhookEventDataPoll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollDeliveredWebhookEventDataPollOption struct {
	CanBeEdited bool `json:"can_be_edited" api:"required"`
	// The participant who added this option (poll creator for the initial options;
	// whoever added later ones). On a poll.updated this differs from the event's
	// `sender_handle` whenever a remote participant added the option. Null when
	// unknown.
	CreatorHandle shared.ChatHandle                              `json:"creator_handle" api:"required"`
	OptionID      string                                         `json:"option_id" api:"required" format:"uuid"`
	Text          string                                         `json:"text" api:"required"`
	Voters        []PollDeliveredWebhookEventDataPollOptionVoter `json:"voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanBeEdited   respjson.Field
		CreatorHandle respjson.Field
		OptionID      respjson.Field
		Text          respjson.Field
		Voters        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollDeliveredWebhookEventDataPollOption) RawJSON() string { return r.JSON.raw }
func (r *PollDeliveredWebhookEventDataPollOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollDeliveredWebhookEventDataPollOptionVoter struct {
	Handle  string    `json:"handle" api:"required"`
	VotedAt time.Time `json:"voted_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		VotedAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollDeliveredWebhookEventDataPollOptionVoter) RawJSON() string { return r.JSON.raw }
func (r *PollDeliveredWebhookEventDataPollOptionVoter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for poll.read events
type PollReadWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for poll.sent, poll.delivered, and poll.read webhook events. Timestamps
	// indicate state (null = not yet happened): sent → sent_at; delivered →
	// +delivered_at; read → +read_at.
	Data PollReadWebhookEventData `json:"data" api:"required"`
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r PollReadWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PollReadWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for poll.sent, poll.delivered, and poll.read webhook events. Timestamps
// indicate state (null = not yet happened): sent → sent_at; delivered →
// +delivered_at; read → +read_at.
type PollReadWebhookEventData struct {
	// Chat info for poll webhook events.
	Chat      PollReadWebhookEventDataChat `json:"chat" api:"required"`
	CreatedAt time.Time                    `json:"created_at" api:"required" format:"date-time"`
	// Any of "inbound", "outbound".
	Direction   string                       `json:"direction" api:"required"`
	MessageID   string                       `json:"message_id" api:"required" format:"uuid"`
	Poll        PollReadWebhookEventDataPoll `json:"poll" api:"required"`
	Service     string                       `json:"service" api:"required"`
	UpdatedAt   time.Time                    `json:"updated_at" api:"required" format:"date-time"`
	DeliveredAt time.Time                    `json:"delivered_at" api:"nullable" format:"date-time"`
	ReadAt      time.Time                    `json:"read_at" api:"nullable" format:"date-time"`
	// The handle that sent the poll.
	SenderHandle shared.ChatHandle `json:"sender_handle" api:"nullable"`
	SentAt       time.Time         `json:"sent_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chat         respjson.Field
		CreatedAt    respjson.Field
		Direction    respjson.Field
		MessageID    respjson.Field
		Poll         respjson.Field
		Service      respjson.Field
		UpdatedAt    respjson.Field
		DeliveredAt  respjson.Field
		ReadAt       respjson.Field
		SenderHandle respjson.Field
		SentAt       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollReadWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PollReadWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat info for poll webhook events.
type PollReadWebhookEventDataChat struct {
	ID          string            `json:"id" api:"required" format:"uuid"`
	IsGroup     bool              `json:"is_group" api:"nullable"`
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
func (r PollReadWebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *PollReadWebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollReadWebhookEventDataPoll struct {
	Options []PollReadWebhookEventDataPollOption `json:"options" api:"required"`
	// Distinct participants across the whole poll.
	TotalVoters int64 `json:"total_voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		TotalVoters respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollReadWebhookEventDataPoll) RawJSON() string { return r.JSON.raw }
func (r *PollReadWebhookEventDataPoll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollReadWebhookEventDataPollOption struct {
	CanBeEdited bool `json:"can_be_edited" api:"required"`
	// The participant who added this option (poll creator for the initial options;
	// whoever added later ones). On a poll.updated this differs from the event's
	// `sender_handle` whenever a remote participant added the option. Null when
	// unknown.
	CreatorHandle shared.ChatHandle                         `json:"creator_handle" api:"required"`
	OptionID      string                                    `json:"option_id" api:"required" format:"uuid"`
	Text          string                                    `json:"text" api:"required"`
	Voters        []PollReadWebhookEventDataPollOptionVoter `json:"voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanBeEdited   respjson.Field
		CreatorHandle respjson.Field
		OptionID      respjson.Field
		Text          respjson.Field
		Voters        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollReadWebhookEventDataPollOption) RawJSON() string { return r.JSON.raw }
func (r *PollReadWebhookEventDataPollOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollReadWebhookEventDataPollOptionVoter struct {
	Handle  string    `json:"handle" api:"required"`
	VotedAt time.Time `json:"voted_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		VotedAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollReadWebhookEventDataPollOptionVoter) RawJSON() string { return r.JSON.raw }
func (r *PollReadWebhookEventDataPollOptionVoter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for poll.updated events
type PollUpdatedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for poll.updated (option(s) added — add-only).
	Data PollUpdatedWebhookEventData `json:"data" api:"required"`
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r PollUpdatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PollUpdatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for poll.updated (option(s) added — add-only).
type PollUpdatedWebhookEventData struct {
	// Only the options this update added — never the ones the poll already had. Fetch
	// the poll to read its full option set.
	AddedOptions []PollUpdatedWebhookEventDataAddedOption `json:"added_options" api:"required"`
	// Chat info for poll webhook events.
	Chat PollUpdatedWebhookEventDataChat `json:"chat" api:"required"`
	// Any of "inbound", "outbound".
	Direction string `json:"direction" api:"required"`
	MessageID string `json:"message_id" api:"required" format:"uuid"`
	// Your line — the one that received or sent this update. Always present. On an
	// inbound update this is NOT who added the option: use
	// `added_options[].creator_handle` for that, which will be the remote participant.
	SenderHandle shared.ChatHandle `json:"sender_handle" api:"required"`
	Service      string            `json:"service" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedOptions respjson.Field
		Chat         respjson.Field
		Direction    respjson.Field
		MessageID    respjson.Field
		SenderHandle respjson.Field
		Service      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollUpdatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PollUpdatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollUpdatedWebhookEventDataAddedOption struct {
	CanBeEdited bool `json:"can_be_edited" api:"required"`
	// The participant who added this option (poll creator for the initial options;
	// whoever added later ones). On a poll.updated this differs from the event's
	// `sender_handle` whenever a remote participant added the option. Null when
	// unknown.
	CreatorHandle shared.ChatHandle                             `json:"creator_handle" api:"required"`
	OptionID      string                                        `json:"option_id" api:"required" format:"uuid"`
	Text          string                                        `json:"text" api:"required"`
	Voters        []PollUpdatedWebhookEventDataAddedOptionVoter `json:"voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanBeEdited   respjson.Field
		CreatorHandle respjson.Field
		OptionID      respjson.Field
		Text          respjson.Field
		Voters        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollUpdatedWebhookEventDataAddedOption) RawJSON() string { return r.JSON.raw }
func (r *PollUpdatedWebhookEventDataAddedOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollUpdatedWebhookEventDataAddedOptionVoter struct {
	Handle  string    `json:"handle" api:"required"`
	VotedAt time.Time `json:"voted_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		VotedAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollUpdatedWebhookEventDataAddedOptionVoter) RawJSON() string { return r.JSON.raw }
func (r *PollUpdatedWebhookEventDataAddedOptionVoter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat info for poll webhook events.
type PollUpdatedWebhookEventDataChat struct {
	ID          string            `json:"id" api:"required" format:"uuid"`
	IsGroup     bool              `json:"is_group" api:"nullable"`
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
func (r PollUpdatedWebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *PollUpdatedWebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for poll.failed events
type PollFailedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for poll.failed — an outbound poll (or poll action) that failed to send.
	// Carries the poll snapshot at failure time plus the error and when it failed.
	Data PollFailedWebhookEventData `json:"data" api:"required"`
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r PollFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PollFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for poll.failed — an outbound poll (or poll action) that failed to send.
// Carries the poll snapshot at failure time plus the error and when it failed.
type PollFailedWebhookEventData struct {
	// Chat info for poll webhook events.
	Chat PollFailedWebhookEventDataChat `json:"chat" api:"required"`
	// Any of "inbound", "outbound".
	Direction string                          `json:"direction" api:"required"`
	Error     PollFailedWebhookEventDataError `json:"error" api:"required"`
	FailedAt  time.Time                       `json:"failed_at" api:"required" format:"date-time"`
	MessageID string                          `json:"message_id" api:"required" format:"uuid"`
	Poll      PollFailedWebhookEventDataPoll  `json:"poll" api:"required"`
	Service   string                          `json:"service" api:"required"`
	// Null on failure (the send never landed).
	SenderHandle shared.ChatHandle `json:"sender_handle" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chat         respjson.Field
		Direction    respjson.Field
		Error        respjson.Field
		FailedAt     respjson.Field
		MessageID    respjson.Field
		Poll         respjson.Field
		Service      respjson.Field
		SenderHandle respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollFailedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PollFailedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat info for poll webhook events.
type PollFailedWebhookEventDataChat struct {
	ID          string            `json:"id" api:"required" format:"uuid"`
	IsGroup     bool              `json:"is_group" api:"nullable"`
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
func (r PollFailedWebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *PollFailedWebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollFailedWebhookEventDataError struct {
	// Error codes in webhook failure events. The possible set varies by event:
	// message.failed and poll.failed can carry 3007, 4001, 4002, 4005, 4006, 4007, or
	// 4008; the group update failure events (chat.group_name_update_failed,
	// chat.group_icon_update_failed) carry 3007 or 4001; chat.background_update_failed
	// carries 1005, 2011, 4001, or 5002.
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollFailedWebhookEventDataError) RawJSON() string { return r.JSON.raw }
func (r *PollFailedWebhookEventDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollFailedWebhookEventDataPoll struct {
	Options []PollFailedWebhookEventDataPollOption `json:"options" api:"required"`
	// Distinct participants across the whole poll.
	TotalVoters int64 `json:"total_voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		TotalVoters respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollFailedWebhookEventDataPoll) RawJSON() string { return r.JSON.raw }
func (r *PollFailedWebhookEventDataPoll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollFailedWebhookEventDataPollOption struct {
	CanBeEdited bool `json:"can_be_edited" api:"required"`
	// The participant who added this option (poll creator for the initial options;
	// whoever added later ones). On a poll.updated this differs from the event's
	// `sender_handle` whenever a remote participant added the option. Null when
	// unknown.
	CreatorHandle shared.ChatHandle                           `json:"creator_handle" api:"required"`
	OptionID      string                                      `json:"option_id" api:"required" format:"uuid"`
	Text          string                                      `json:"text" api:"required"`
	Voters        []PollFailedWebhookEventDataPollOptionVoter `json:"voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanBeEdited   respjson.Field
		CreatorHandle respjson.Field
		OptionID      respjson.Field
		Text          respjson.Field
		Voters        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollFailedWebhookEventDataPollOption) RawJSON() string { return r.JSON.raw }
func (r *PollFailedWebhookEventDataPollOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollFailedWebhookEventDataPollOptionVoter struct {
	Handle  string    `json:"handle" api:"required"`
	VotedAt time.Time `json:"voted_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		VotedAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollFailedWebhookEventDataPollOptionVoter) RawJSON() string { return r.JSON.raw }
func (r *PollFailedWebhookEventDataPollOptionVoter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for poll.vote.added events
type PollVoteAddedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for poll.vote.added and poll.vote.removed (one option toggled).
	Data PollVoteAddedWebhookEventData `json:"data" api:"required"`
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r PollVoteAddedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PollVoteAddedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for poll.vote.added and poll.vote.removed (one option toggled).
type PollVoteAddedWebhookEventData struct {
	// Chat info for poll webhook events.
	Chat PollVoteAddedWebhookEventDataChat `json:"chat" api:"required"`
	// Any of "inbound", "outbound".
	Direction string `json:"direction" api:"required"`
	MessageID string `json:"message_id" api:"required" format:"uuid"`
	OptionID  string `json:"option_id" api:"required" format:"uuid"`
	// The voter — always present.
	SenderHandle shared.ChatHandle `json:"sender_handle" api:"required"`
	Service      string            `json:"service" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chat         respjson.Field
		Direction    respjson.Field
		MessageID    respjson.Field
		OptionID     respjson.Field
		SenderHandle respjson.Field
		Service      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollVoteAddedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PollVoteAddedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat info for poll webhook events.
type PollVoteAddedWebhookEventDataChat struct {
	ID          string            `json:"id" api:"required" format:"uuid"`
	IsGroup     bool              `json:"is_group" api:"nullable"`
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
func (r PollVoteAddedWebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *PollVoteAddedWebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for poll.vote.removed events
type PollVoteRemovedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for poll.vote.added and poll.vote.removed (one option toggled).
	Data PollVoteRemovedWebhookEventData `json:"data" api:"required"`
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r PollVoteRemovedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PollVoteRemovedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for poll.vote.added and poll.vote.removed (one option toggled).
type PollVoteRemovedWebhookEventData struct {
	// Chat info for poll webhook events.
	Chat PollVoteRemovedWebhookEventDataChat `json:"chat" api:"required"`
	// Any of "inbound", "outbound".
	Direction string `json:"direction" api:"required"`
	MessageID string `json:"message_id" api:"required" format:"uuid"`
	OptionID  string `json:"option_id" api:"required" format:"uuid"`
	// The voter — always present.
	SenderHandle shared.ChatHandle `json:"sender_handle" api:"required"`
	Service      string            `json:"service" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chat         respjson.Field
		Direction    respjson.Field
		MessageID    respjson.Field
		OptionID     respjson.Field
		SenderHandle respjson.Field
		Service      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollVoteRemovedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PollVoteRemovedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat info for poll webhook events.
type PollVoteRemovedWebhookEventDataChat struct {
	ID          string            `json:"id" api:"required" format:"uuid"`
	IsGroup     bool              `json:"is_group" api:"nullable"`
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
func (r PollVoteRemovedWebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *PollVoteRemovedWebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for poll.reaction.added events
type PollReactionAddedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for poll.reaction.added — a reaction on a poll message. Same shape as
	// reaction.added; `message_id` is the poll-definition message's ID. Poll reactions
	// are stickers, which iMessage cannot remove, so there is no removal counterpart.
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r PollReactionAddedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PollReactionAddedWebhookEvent) UnmarshalJSON(data []byte) error {
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// **[BETA]** Current health for a chat. Always present — chats start at `HEALTHY`
	// and may shift based on engagement and delivery signals on the conversation. Many
	// `AT_RISK` or `CRITICAL` chats on a single line increase the risk of line
	// flagging.
	//
	// Switch on `status` to surface chat and line health in your UI — the enum is the
	// long-term contract. Each status carries a `doc_url` that deep-links to the
	// relevant section of the Chat Health guide. To gate a send, act on the response
	// rather than the status: a `403` is the authoritative answer.
	//
	// See the [Chat Health guide](/channel/imessage/guides/chats/chat-health) for what
	// each status means and how to react.
	HealthStatus ChatCreatedWebhookEventDataHealthStatus `json:"health_status" api:"required"`
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
		ID           respjson.Field
		CreatedAt    respjson.Field
		DisplayName  respjson.Field
		Handles      respjson.Field
		HealthStatus respjson.Field
		IsGroup      respjson.Field
		UpdatedAt    respjson.Field
		Service      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCreatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatCreatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **[BETA]** Current health for a chat. Always present — chats start at `HEALTHY`
// and may shift based on engagement and delivery signals on the conversation. Many
// `AT_RISK` or `CRITICAL` chats on a single line increase the risk of line
// flagging.
//
// Switch on `status` to surface chat and line health in your UI — the enum is the
// long-term contract. Each status carries a `doc_url` that deep-links to the
// relevant section of the Chat Health guide. To gate a send, act on the response
// rather than the status: a `403` is the authoritative answer.
//
// See the [Chat Health guide](/channel/imessage/guides/chats/chat-health) for what
// each status means and how to react.
type ChatCreatedWebhookEventDataHealthStatus struct {
	// Deep-link to the relevant section of the Chat Health guide for this status.
	DocURL string `json:"doc_url" api:"required" format:"uri"`
	// Current health bucket for the chat. See the
	// [Chat Health guide](/channel/imessage/guides/chats/chat-health) for what each
	// value means and how to react. `doc_url` deep-links to the relevant section.
	//
	// `OPTED_OUT` — the recipient sent `STOP`, `UNSUBSCRIBE`, `OPTOUT`, `CANCEL`,
	// `END`, or `QUIT`. The keyword must be the whole trimmed message, never part of a
	// longer one: `STOP` counts, `please stop` does not. Most keywords must match
	// exactly, including case. `OPT OUT` is the exception — it matches in any casing,
	// with or without the space or a hyphen, so `opt out`, `Opt-Out` and `optout` all
	// count. It clears as soon as they reply again: any later message from them that
	// is not itself an opt-out keyword opts them back in immediately — a reply in any
	// conversation with you counts, the same way the block does.
	//
	// `OPTED_OUT` marks only the conversation the keyword arrived in. The block below
	// is wider than the mark, so a conversation still reading `HEALTHY` can be blocked
	// as well — gate on the `403`, not on the status. Group threads are never marked
	// and are never blocked.
	//
	// Linq enforces this: while a recipient is opted out, every send to them is
	// rejected with `403` (error code `2024`) before the message is queued, across
	// every chat and every line on your account. Nothing is delivered, including a
	// final courtesy message — to send one, set `override_optout: true` on that single
	// request.
	//
	// Any of "HEALTHY", "AT_RISK", "CRITICAL", "OPTED_OUT".
	Status string `json:"status" api:"required"`
	// When this status last changed.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocURL      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCreatedWebhookEventDataHealthStatus) RawJSON() string { return r.JSON.raw }
func (r *ChatCreatedWebhookEventDataHealthStatus) UnmarshalJSON(data []byte) error {
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// Error codes in webhook failure events. The possible set varies by event:
	// message.failed and poll.failed can carry 3007, 4001, 4002, 4005, 4006, 4007, or
	// 4008; the group update failure events (chat.group_name_update_failed,
	// chat.group_icon_update_failed) carry 3007 or 4001; chat.background_update_failed
	// carries 1005, 2011, 4001, or 5002.
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// Error codes in webhook failure events. The possible set varies by event:
	// message.failed and poll.failed can carry 3007, 4001, 4002, 4005, 4006, 4007, or
	// 4008; the group update failure events (chat.group_name_update_failed,
	// chat.group_icon_update_failed) carry 3007 or 4001; chat.background_update_failed
	// carries 1005, 2011, 4001, or 5002.
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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

// Complete webhook payload for chat.background_updated events
type ChatBackgroundUpdatedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for chat.background_updated webhook events.
	Data ChatBackgroundUpdatedWebhookEventData `json:"data" api:"required"`
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r ChatBackgroundUpdatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatBackgroundUpdatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for chat.background_updated webhook events.
type ChatBackgroundUpdatedWebhookEventData struct {
	// Chat information
	Chat ChatBackgroundUpdatedWebhookEventDataChat `json:"chat" api:"required"`
	// Who changed it. `is_me` is true when your own number set it.
	ActorHandle shared.ChatHandle `json:"actor_handle" api:"nullable"`
	// A chat transcript background. Fields are populated per `type`.
	Background ChatBackgroundUpdatedWebhookEventDataBackground `json:"background" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chat        respjson.Field
		ActorHandle respjson.Field
		Background  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatBackgroundUpdatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatBackgroundUpdatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat information
type ChatBackgroundUpdatedWebhookEventDataChat struct {
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
func (r ChatBackgroundUpdatedWebhookEventDataChat) RawJSON() string { return r.JSON.raw }
func (r *ChatBackgroundUpdatedWebhookEventDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A chat transcript background. Fields are populated per `type`.
type ChatBackgroundUpdatedWebhookEventDataBackground struct {
	// The background family.
	//
	// Any of "color", "dynamic", "photo".
	Type string `json:"type" api:"required"`
	// Photo: a hosted URL for the background image, whether you set it or a
	// participant did. Apple stores the image, not the URL it came from, so the image
	// is re-hosted and this is our URL rather than the one you supplied. `null` only
	// if the image could not be hosted.
	ImageURL string `json:"image_url" api:"nullable"`
	// Color: the two gradient stops as hex, top then bottom.
	Shades []string `json:"shades" api:"nullable"`
	// Dynamic: the animated style.
	//
	// Any of "sky", "water", "aurora", "glitter".
	Style string `json:"style" api:"nullable"`
	// Color: `custom` (the stored two colors) or a named swatch. Dynamic: the variant
	// within the `style` (e.g. `sunrise`).
	Variant string `json:"variant" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ImageURL    respjson.Field
		Shades      respjson.Field
		Style       respjson.Field
		Variant     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatBackgroundUpdatedWebhookEventDataBackground) RawJSON() string { return r.JSON.raw }
func (r *ChatBackgroundUpdatedWebhookEventDataBackground) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for chat.background_update_failed events
type ChatBackgroundUpdateFailedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Error details for chat.background_update_failed webhook events. See
	// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
	// code reference.
	Data ChatBackgroundUpdateFailedWebhookEventData `json:"data" api:"required"`
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r ChatBackgroundUpdateFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatBackgroundUpdateFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for chat.background_update_failed webhook events. See
// [WebhookErrorCode](#/components/schemas/WebhookErrorCode) for the full error
// code reference.
type ChatBackgroundUpdateFailedWebhookEventData struct {
	// Chat identifier (UUID) whose background update failed
	ChatID string `json:"chat_id" api:"required"`
	// Error codes in webhook failure events. The possible set varies by event:
	// message.failed and poll.failed can carry 3007, 4001, 4002, 4005, 4006, 4007, or
	// 4008; the group update failure events (chat.group_name_update_failed,
	// chat.group_icon_update_failed) carry 3007 or 4001; chat.background_update_failed
	// carries 1005, 2011, 4001, or 5002.
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
func (r ChatBackgroundUpdateFailedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ChatBackgroundUpdateFailedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete webhook payload for contact_card.received events
type ContactCardReceivedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Payload for contact_card.received webhook events.
	//
	// A contact belongs to a line, not to an individual chat. You receive one event
	// per person who shares their contact, regardless of how many chats they have in
	// common with your line.
	//
	// The event fires again whenever the shared contact's name or media changes.
	Data ContactCardReceivedWebhookEventData `json:"data" api:"required"`
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
func (r ContactCardReceivedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ContactCardReceivedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for contact_card.received webhook events.
//
// A contact belongs to a line, not to an individual chat. You receive one event
// per person who shares their contact, regardless of how many chats they have in
// common with your line.
//
// The event fires again whenever the shared contact's name or media changes.
type ContactCardReceivedWebhookEventData struct {
	// First name from the shared contact card
	FirstName string `json:"first_name" api:"required"`
	// Last name from the shared contact card (may be empty)
	LastName string `json:"last_name" api:"required"`
	// Which of your lines they shared it with.
	OwnerHandle string `json:"owner_handle" api:"required"`
	// The person who shared their card — a phone number or email address.
	SenderHandle string `json:"sender_handle" api:"required"`
	// URL of the contact's media, served from `cdn.linqapp.com`. `null` when the
	// contact shared no media, and also when media was shared but could not be
	// retrieved — this field does not distinguish the two.
	//
	// Download the media and store it yourself. The URL may be signed and expire, in
	// as little as 45 minutes, and altering its query string invalidates it
	// immediately.
	MediaURL string `json:"media_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName    respjson.Field
		LastName     respjson.Field
		OwnerHandle  respjson.Field
		SenderHandle respjson.Field
		MediaURL     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactCardReceivedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ContactCardReceivedWebhookEventData) UnmarshalJSON(data []byte) error {
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
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.succeeded", "payment.canceled", "payment.expired", "payment.declined",
	// "payment.authorized", "connection.created", "connection.revoked".
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
	// The new line reputation
	//
	// Any of "HEALTHY", "AT_RISK", "CRITICAL".
	NewReputation string `json:"new_reputation" api:"required"`
	// The new service status
	//
	// Any of "ACTIVE", "FLAGGED".
	NewStatus string `json:"new_status" api:"required"`
	// Phone number in E.164 format
	PhoneNumber string `json:"phone_number" api:"required"`
	// The previous line reputation
	//
	// Any of "HEALTHY", "AT_RISK", "CRITICAL".
	PreviousReputation string `json:"previous_reputation" api:"required"`
	// The previous service status
	//
	// Any of "ACTIVE", "FLAGGED".
	PreviousStatus string `json:"previous_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedAt          respjson.Field
		NewReputation      respjson.Field
		NewStatus          respjson.Field
		PhoneNumber        respjson.Field
		PreviousReputation respjson.Field
		PreviousStatus     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
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
	PhoneNumberStatusUpdatedWebhookEventEventTypePollReceived               PhoneNumberStatusUpdatedWebhookEventEventType = "poll.received"
	PhoneNumberStatusUpdatedWebhookEventEventTypePollFailed                 PhoneNumberStatusUpdatedWebhookEventEventType = "poll.failed"
	PhoneNumberStatusUpdatedWebhookEventEventTypePollSent                   PhoneNumberStatusUpdatedWebhookEventEventType = "poll.sent"
	PhoneNumberStatusUpdatedWebhookEventEventTypePollDelivered              PhoneNumberStatusUpdatedWebhookEventEventType = "poll.delivered"
	PhoneNumberStatusUpdatedWebhookEventEventTypePollRead                   PhoneNumberStatusUpdatedWebhookEventEventType = "poll.read"
	PhoneNumberStatusUpdatedWebhookEventEventTypePollUpdated                PhoneNumberStatusUpdatedWebhookEventEventType = "poll.updated"
	PhoneNumberStatusUpdatedWebhookEventEventTypePollVoteAdded              PhoneNumberStatusUpdatedWebhookEventEventType = "poll.vote.added"
	PhoneNumberStatusUpdatedWebhookEventEventTypePollVoteRemoved            PhoneNumberStatusUpdatedWebhookEventEventType = "poll.vote.removed"
	PhoneNumberStatusUpdatedWebhookEventEventTypePollReactionAdded          PhoneNumberStatusUpdatedWebhookEventEventType = "poll.reaction.added"
	PhoneNumberStatusUpdatedWebhookEventEventTypeParticipantAdded           PhoneNumberStatusUpdatedWebhookEventEventType = "participant.added"
	PhoneNumberStatusUpdatedWebhookEventEventTypeParticipantRemoved         PhoneNumberStatusUpdatedWebhookEventEventType = "participant.removed"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatCreated                PhoneNumberStatusUpdatedWebhookEventEventType = "chat.created"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatGroupNameUpdated       PhoneNumberStatusUpdatedWebhookEventEventType = "chat.group_name_updated"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatGroupIconUpdated       PhoneNumberStatusUpdatedWebhookEventEventType = "chat.group_icon_updated"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatGroupNameUpdateFailed  PhoneNumberStatusUpdatedWebhookEventEventType = "chat.group_name_update_failed"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatGroupIconUpdateFailed  PhoneNumberStatusUpdatedWebhookEventEventType = "chat.group_icon_update_failed"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatBackgroundUpdated      PhoneNumberStatusUpdatedWebhookEventEventType = "chat.background_updated"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatBackgroundUpdateFailed PhoneNumberStatusUpdatedWebhookEventEventType = "chat.background_update_failed"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatTypingIndicatorStarted PhoneNumberStatusUpdatedWebhookEventEventType = "chat.typing_indicator.started"
	PhoneNumberStatusUpdatedWebhookEventEventTypeChatTypingIndicatorStopped PhoneNumberStatusUpdatedWebhookEventEventType = "chat.typing_indicator.stopped"
	PhoneNumberStatusUpdatedWebhookEventEventTypePhoneNumberStatusUpdated   PhoneNumberStatusUpdatedWebhookEventEventType = "phone_number.status_updated"
	PhoneNumberStatusUpdatedWebhookEventEventTypeContactCardReceived        PhoneNumberStatusUpdatedWebhookEventEventType = "contact_card.received"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallInitiated              PhoneNumberStatusUpdatedWebhookEventEventType = "call.initiated"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallRinging                PhoneNumberStatusUpdatedWebhookEventEventType = "call.ringing"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallAnswered               PhoneNumberStatusUpdatedWebhookEventEventType = "call.answered"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallEnded                  PhoneNumberStatusUpdatedWebhookEventEventType = "call.ended"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallFailed                 PhoneNumberStatusUpdatedWebhookEventEventType = "call.failed"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallDeclined               PhoneNumberStatusUpdatedWebhookEventEventType = "call.declined"
	PhoneNumberStatusUpdatedWebhookEventEventTypeCallNoAnswer               PhoneNumberStatusUpdatedWebhookEventEventType = "call.no_answer"
	PhoneNumberStatusUpdatedWebhookEventEventTypeLocationSharingStarted     PhoneNumberStatusUpdatedWebhookEventEventType = "location.sharing.started"
	PhoneNumberStatusUpdatedWebhookEventEventTypeLocationSharingStopped     PhoneNumberStatusUpdatedWebhookEventEventType = "location.sharing.stopped"
	PhoneNumberStatusUpdatedWebhookEventEventTypePaymentSucceeded           PhoneNumberStatusUpdatedWebhookEventEventType = "payment.succeeded"
	PhoneNumberStatusUpdatedWebhookEventEventTypePaymentCanceled            PhoneNumberStatusUpdatedWebhookEventEventType = "payment.canceled"
	PhoneNumberStatusUpdatedWebhookEventEventTypePaymentExpired             PhoneNumberStatusUpdatedWebhookEventEventType = "payment.expired"
	PhoneNumberStatusUpdatedWebhookEventEventTypePaymentDeclined            PhoneNumberStatusUpdatedWebhookEventEventType = "payment.declined"
	PhoneNumberStatusUpdatedWebhookEventEventTypePaymentAuthorized          PhoneNumberStatusUpdatedWebhookEventEventType = "payment.authorized"
	PhoneNumberStatusUpdatedWebhookEventEventTypeConnectionCreated          PhoneNumberStatusUpdatedWebhookEventEventType = "connection.created"
	PhoneNumberStatusUpdatedWebhookEventEventTypeConnectionRevoked          PhoneNumberStatusUpdatedWebhookEventEventType = "connection.revoked"
)

type ConnectionCreatedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The payment request, as returned by
	// `GET /v3/payment_requests/{paymentRequestId}`.
	Data ConnectionCreatedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Any of "payment.succeeded", "payment.canceled", "payment.expired",
	// "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.declined", "payment.authorized", "connection.created",
	// "connection.revoked".
	EventType ConnectionCreatedWebhookEventEventType `json:"event_type" api:"required"`
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
func (r ConnectionCreatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConnectionCreatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payment request, as returned by
// `GET /v3/payment_requests/{paymentRequestId}`.
type ConnectionCreatedWebhookEventData struct {
	// The payment request id.
	ID string `json:"id" api:"required" format:"uuid"`
	// What was charged at checkout, in the currency's minor units. In `subscription`
	// mode this is the first invoice's total — all items after any discounts are
	// applied.
	Amount int64 `json:"amount" api:"required"`
	// URL the recipient opens to pay
	// (`https://zero.linqapp.com/pay/{slug}?session=...`).
	CheckoutURL string    `json:"checkout_url" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency    string    `json:"currency" api:"required"`
	Object      string    `json:"object" api:"required"`
	// Any of "succeeded", "failed", "canceled", "expired".
	Status      string `json:"status" api:"required"`
	Description string `json:"description"`
	// Subscription mode — the discount Stripe applied, read back from the coupon.
	// Absent when none was applied.
	Discount ConnectionCreatedWebhookEventDataDiscount `json:"discount"`
	// Subscription mode — how often the subscription renews.
	//
	// Any of "day", "week", "month", "year".
	Interval string `json:"interval"`
	// Subscription mode — intervals per renewal.
	IntervalCount int64             `json:"interval_count"`
	Metadata      map[string]string `json:"metadata"`
	// Whether the request collected a one-time charge or started a subscription.
	//
	// Any of "payment", "subscription".
	Mode string `json:"mode"`
	// Natural-rail join keys, present when `rail: natural`.
	Natural ConnectionCreatedWebhookEventDataNatural `json:"natural"`
	// Subscription mode — the recurring price subscribed to.
	PriceID string `json:"price_id"`
	// Subscription mode — units of the price subscribed to.
	Quantity int64 `json:"quantity"`
	// The rail this request settled on.
	//
	// Any of "stripe", "natural".
	Rail string `json:"rail"`
	// Ids of the Stripe objects on your connected account — join keys into your own
	// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
	// `subscription_id`.
	Stripe ConnectionCreatedWebhookEventDataStripe `json:"stripe"`
	// Subscription mode — when the free trial ends and the first charge happens. On a
	// trial request, `payment.succeeded` means the payment method was collected ($0
	// moved).
	TrialEnd  time.Time `json:"trial_end" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		CheckoutURL   respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Object        respjson.Field
		Status        respjson.Field
		Description   respjson.Field
		Discount      respjson.Field
		Interval      respjson.Field
		IntervalCount respjson.Field
		Metadata      respjson.Field
		Mode          respjson.Field
		Natural       respjson.Field
		PriceID       respjson.Field
		Quantity      respjson.Field
		Rail          respjson.Field
		Stripe        respjson.Field
		TrialEnd      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionCreatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConnectionCreatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription mode — the discount Stripe applied, read back from the coupon.
// Absent when none was applied.
type ConnectionCreatedWebhookEventDataDiscount struct {
	Coupon string `json:"coupon"`
	// Name of the coupon/promo code displayed to customers.
	Label         string `json:"label"`
	PromotionCode string `json:"promotion_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coupon        respjson.Field
		Label         respjson.Field
		PromotionCode respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionCreatedWebhookEventDataDiscount) RawJSON() string { return r.JSON.raw }
func (r *ConnectionCreatedWebhookEventDataDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Natural-rail join keys, present when `rail: natural`.
type ConnectionCreatedWebhookEventDataNatural struct {
	// The Natural payment request (`prq_...`).
	PaymentRequestID string `json:"payment_request_id"`
	// The settled transaction (`txn_...`).
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentRequestID respjson.Field
		TransactionID    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionCreatedWebhookEventDataNatural) RawJSON() string { return r.JSON.raw }
func (r *ConnectionCreatedWebhookEventDataNatural) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ids of the Stripe objects on your connected account — join keys into your own
// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
// `subscription_id`.
type ConnectionCreatedWebhookEventDataStripe struct {
	// The Customer the request is attached to (`cus_...`). Always set in subscription
	// mode; set in payment mode only when the request was created with a
	// `customer_id`.
	CustomerID string `json:"customer_id"`
	// The PaymentIntent collected at checkout (`pi_...`).
	PaymentIntentID string `json:"payment_intent_id"`
	// Subscription mode — the Subscription (`sub_...`).
	SubscriptionID string `json:"subscription_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID      respjson.Field
		PaymentIntentID respjson.Field
		SubscriptionID  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionCreatedWebhookEventDataStripe) RawJSON() string { return r.JSON.raw }
func (r *ConnectionCreatedWebhookEventDataStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionCreatedWebhookEventEventType string

const (
	ConnectionCreatedWebhookEventEventTypePaymentSucceeded           ConnectionCreatedWebhookEventEventType = "payment.succeeded"
	ConnectionCreatedWebhookEventEventTypePaymentCanceled            ConnectionCreatedWebhookEventEventType = "payment.canceled"
	ConnectionCreatedWebhookEventEventTypePaymentExpired             ConnectionCreatedWebhookEventEventType = "payment.expired"
	ConnectionCreatedWebhookEventEventTypeMessageSent                ConnectionCreatedWebhookEventEventType = "message.sent"
	ConnectionCreatedWebhookEventEventTypeMessageReceived            ConnectionCreatedWebhookEventEventType = "message.received"
	ConnectionCreatedWebhookEventEventTypeMessageRead                ConnectionCreatedWebhookEventEventType = "message.read"
	ConnectionCreatedWebhookEventEventTypeMessageDelivered           ConnectionCreatedWebhookEventEventType = "message.delivered"
	ConnectionCreatedWebhookEventEventTypeMessageFailed              ConnectionCreatedWebhookEventEventType = "message.failed"
	ConnectionCreatedWebhookEventEventTypeMessageEdited              ConnectionCreatedWebhookEventEventType = "message.edited"
	ConnectionCreatedWebhookEventEventTypeReactionAdded              ConnectionCreatedWebhookEventEventType = "reaction.added"
	ConnectionCreatedWebhookEventEventTypeReactionRemoved            ConnectionCreatedWebhookEventEventType = "reaction.removed"
	ConnectionCreatedWebhookEventEventTypePollReceived               ConnectionCreatedWebhookEventEventType = "poll.received"
	ConnectionCreatedWebhookEventEventTypePollFailed                 ConnectionCreatedWebhookEventEventType = "poll.failed"
	ConnectionCreatedWebhookEventEventTypePollSent                   ConnectionCreatedWebhookEventEventType = "poll.sent"
	ConnectionCreatedWebhookEventEventTypePollDelivered              ConnectionCreatedWebhookEventEventType = "poll.delivered"
	ConnectionCreatedWebhookEventEventTypePollRead                   ConnectionCreatedWebhookEventEventType = "poll.read"
	ConnectionCreatedWebhookEventEventTypePollUpdated                ConnectionCreatedWebhookEventEventType = "poll.updated"
	ConnectionCreatedWebhookEventEventTypePollVoteAdded              ConnectionCreatedWebhookEventEventType = "poll.vote.added"
	ConnectionCreatedWebhookEventEventTypePollVoteRemoved            ConnectionCreatedWebhookEventEventType = "poll.vote.removed"
	ConnectionCreatedWebhookEventEventTypePollReactionAdded          ConnectionCreatedWebhookEventEventType = "poll.reaction.added"
	ConnectionCreatedWebhookEventEventTypeParticipantAdded           ConnectionCreatedWebhookEventEventType = "participant.added"
	ConnectionCreatedWebhookEventEventTypeParticipantRemoved         ConnectionCreatedWebhookEventEventType = "participant.removed"
	ConnectionCreatedWebhookEventEventTypeChatCreated                ConnectionCreatedWebhookEventEventType = "chat.created"
	ConnectionCreatedWebhookEventEventTypeChatGroupNameUpdated       ConnectionCreatedWebhookEventEventType = "chat.group_name_updated"
	ConnectionCreatedWebhookEventEventTypeChatGroupIconUpdated       ConnectionCreatedWebhookEventEventType = "chat.group_icon_updated"
	ConnectionCreatedWebhookEventEventTypeChatGroupNameUpdateFailed  ConnectionCreatedWebhookEventEventType = "chat.group_name_update_failed"
	ConnectionCreatedWebhookEventEventTypeChatGroupIconUpdateFailed  ConnectionCreatedWebhookEventEventType = "chat.group_icon_update_failed"
	ConnectionCreatedWebhookEventEventTypeChatBackgroundUpdated      ConnectionCreatedWebhookEventEventType = "chat.background_updated"
	ConnectionCreatedWebhookEventEventTypeChatBackgroundUpdateFailed ConnectionCreatedWebhookEventEventType = "chat.background_update_failed"
	ConnectionCreatedWebhookEventEventTypeChatTypingIndicatorStarted ConnectionCreatedWebhookEventEventType = "chat.typing_indicator.started"
	ConnectionCreatedWebhookEventEventTypeChatTypingIndicatorStopped ConnectionCreatedWebhookEventEventType = "chat.typing_indicator.stopped"
	ConnectionCreatedWebhookEventEventTypePhoneNumberStatusUpdated   ConnectionCreatedWebhookEventEventType = "phone_number.status_updated"
	ConnectionCreatedWebhookEventEventTypeContactCardReceived        ConnectionCreatedWebhookEventEventType = "contact_card.received"
	ConnectionCreatedWebhookEventEventTypeCallInitiated              ConnectionCreatedWebhookEventEventType = "call.initiated"
	ConnectionCreatedWebhookEventEventTypeCallRinging                ConnectionCreatedWebhookEventEventType = "call.ringing"
	ConnectionCreatedWebhookEventEventTypeCallAnswered               ConnectionCreatedWebhookEventEventType = "call.answered"
	ConnectionCreatedWebhookEventEventTypeCallEnded                  ConnectionCreatedWebhookEventEventType = "call.ended"
	ConnectionCreatedWebhookEventEventTypeCallFailed                 ConnectionCreatedWebhookEventEventType = "call.failed"
	ConnectionCreatedWebhookEventEventTypeCallDeclined               ConnectionCreatedWebhookEventEventType = "call.declined"
	ConnectionCreatedWebhookEventEventTypeCallNoAnswer               ConnectionCreatedWebhookEventEventType = "call.no_answer"
	ConnectionCreatedWebhookEventEventTypeLocationSharingStarted     ConnectionCreatedWebhookEventEventType = "location.sharing.started"
	ConnectionCreatedWebhookEventEventTypeLocationSharingStopped     ConnectionCreatedWebhookEventEventType = "location.sharing.stopped"
	ConnectionCreatedWebhookEventEventTypePaymentDeclined            ConnectionCreatedWebhookEventEventType = "payment.declined"
	ConnectionCreatedWebhookEventEventTypePaymentAuthorized          ConnectionCreatedWebhookEventEventType = "payment.authorized"
	ConnectionCreatedWebhookEventEventTypeConnectionCreated          ConnectionCreatedWebhookEventEventType = "connection.created"
	ConnectionCreatedWebhookEventEventTypeConnectionRevoked          ConnectionCreatedWebhookEventEventType = "connection.revoked"
)

type ConnectionRevokedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The payment request, as returned by
	// `GET /v3/payment_requests/{paymentRequestId}`.
	Data ConnectionRevokedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Any of "payment.succeeded", "payment.canceled", "payment.expired",
	// "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.declined", "payment.authorized", "connection.created",
	// "connection.revoked".
	EventType ConnectionRevokedWebhookEventEventType `json:"event_type" api:"required"`
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
func (r ConnectionRevokedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConnectionRevokedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payment request, as returned by
// `GET /v3/payment_requests/{paymentRequestId}`.
type ConnectionRevokedWebhookEventData struct {
	// The payment request id.
	ID string `json:"id" api:"required" format:"uuid"`
	// What was charged at checkout, in the currency's minor units. In `subscription`
	// mode this is the first invoice's total — all items after any discounts are
	// applied.
	Amount int64 `json:"amount" api:"required"`
	// URL the recipient opens to pay
	// (`https://zero.linqapp.com/pay/{slug}?session=...`).
	CheckoutURL string    `json:"checkout_url" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency    string    `json:"currency" api:"required"`
	Object      string    `json:"object" api:"required"`
	// Any of "succeeded", "failed", "canceled", "expired".
	Status      string `json:"status" api:"required"`
	Description string `json:"description"`
	// Subscription mode — the discount Stripe applied, read back from the coupon.
	// Absent when none was applied.
	Discount ConnectionRevokedWebhookEventDataDiscount `json:"discount"`
	// Subscription mode — how often the subscription renews.
	//
	// Any of "day", "week", "month", "year".
	Interval string `json:"interval"`
	// Subscription mode — intervals per renewal.
	IntervalCount int64             `json:"interval_count"`
	Metadata      map[string]string `json:"metadata"`
	// Whether the request collected a one-time charge or started a subscription.
	//
	// Any of "payment", "subscription".
	Mode string `json:"mode"`
	// Natural-rail join keys, present when `rail: natural`.
	Natural ConnectionRevokedWebhookEventDataNatural `json:"natural"`
	// Subscription mode — the recurring price subscribed to.
	PriceID string `json:"price_id"`
	// Subscription mode — units of the price subscribed to.
	Quantity int64 `json:"quantity"`
	// The rail this request settled on.
	//
	// Any of "stripe", "natural".
	Rail string `json:"rail"`
	// Ids of the Stripe objects on your connected account — join keys into your own
	// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
	// `subscription_id`.
	Stripe ConnectionRevokedWebhookEventDataStripe `json:"stripe"`
	// Subscription mode — when the free trial ends and the first charge happens. On a
	// trial request, `payment.succeeded` means the payment method was collected ($0
	// moved).
	TrialEnd  time.Time `json:"trial_end" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		CheckoutURL   respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Object        respjson.Field
		Status        respjson.Field
		Description   respjson.Field
		Discount      respjson.Field
		Interval      respjson.Field
		IntervalCount respjson.Field
		Metadata      respjson.Field
		Mode          respjson.Field
		Natural       respjson.Field
		PriceID       respjson.Field
		Quantity      respjson.Field
		Rail          respjson.Field
		Stripe        respjson.Field
		TrialEnd      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionRevokedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConnectionRevokedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription mode — the discount Stripe applied, read back from the coupon.
// Absent when none was applied.
type ConnectionRevokedWebhookEventDataDiscount struct {
	Coupon string `json:"coupon"`
	// Name of the coupon/promo code displayed to customers.
	Label         string `json:"label"`
	PromotionCode string `json:"promotion_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coupon        respjson.Field
		Label         respjson.Field
		PromotionCode respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionRevokedWebhookEventDataDiscount) RawJSON() string { return r.JSON.raw }
func (r *ConnectionRevokedWebhookEventDataDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Natural-rail join keys, present when `rail: natural`.
type ConnectionRevokedWebhookEventDataNatural struct {
	// The Natural payment request (`prq_...`).
	PaymentRequestID string `json:"payment_request_id"`
	// The settled transaction (`txn_...`).
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentRequestID respjson.Field
		TransactionID    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionRevokedWebhookEventDataNatural) RawJSON() string { return r.JSON.raw }
func (r *ConnectionRevokedWebhookEventDataNatural) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ids of the Stripe objects on your connected account — join keys into your own
// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
// `subscription_id`.
type ConnectionRevokedWebhookEventDataStripe struct {
	// The Customer the request is attached to (`cus_...`). Always set in subscription
	// mode; set in payment mode only when the request was created with a
	// `customer_id`.
	CustomerID string `json:"customer_id"`
	// The PaymentIntent collected at checkout (`pi_...`).
	PaymentIntentID string `json:"payment_intent_id"`
	// Subscription mode — the Subscription (`sub_...`).
	SubscriptionID string `json:"subscription_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID      respjson.Field
		PaymentIntentID respjson.Field
		SubscriptionID  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionRevokedWebhookEventDataStripe) RawJSON() string { return r.JSON.raw }
func (r *ConnectionRevokedWebhookEventDataStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionRevokedWebhookEventEventType string

const (
	ConnectionRevokedWebhookEventEventTypePaymentSucceeded           ConnectionRevokedWebhookEventEventType = "payment.succeeded"
	ConnectionRevokedWebhookEventEventTypePaymentCanceled            ConnectionRevokedWebhookEventEventType = "payment.canceled"
	ConnectionRevokedWebhookEventEventTypePaymentExpired             ConnectionRevokedWebhookEventEventType = "payment.expired"
	ConnectionRevokedWebhookEventEventTypeMessageSent                ConnectionRevokedWebhookEventEventType = "message.sent"
	ConnectionRevokedWebhookEventEventTypeMessageReceived            ConnectionRevokedWebhookEventEventType = "message.received"
	ConnectionRevokedWebhookEventEventTypeMessageRead                ConnectionRevokedWebhookEventEventType = "message.read"
	ConnectionRevokedWebhookEventEventTypeMessageDelivered           ConnectionRevokedWebhookEventEventType = "message.delivered"
	ConnectionRevokedWebhookEventEventTypeMessageFailed              ConnectionRevokedWebhookEventEventType = "message.failed"
	ConnectionRevokedWebhookEventEventTypeMessageEdited              ConnectionRevokedWebhookEventEventType = "message.edited"
	ConnectionRevokedWebhookEventEventTypeReactionAdded              ConnectionRevokedWebhookEventEventType = "reaction.added"
	ConnectionRevokedWebhookEventEventTypeReactionRemoved            ConnectionRevokedWebhookEventEventType = "reaction.removed"
	ConnectionRevokedWebhookEventEventTypePollReceived               ConnectionRevokedWebhookEventEventType = "poll.received"
	ConnectionRevokedWebhookEventEventTypePollFailed                 ConnectionRevokedWebhookEventEventType = "poll.failed"
	ConnectionRevokedWebhookEventEventTypePollSent                   ConnectionRevokedWebhookEventEventType = "poll.sent"
	ConnectionRevokedWebhookEventEventTypePollDelivered              ConnectionRevokedWebhookEventEventType = "poll.delivered"
	ConnectionRevokedWebhookEventEventTypePollRead                   ConnectionRevokedWebhookEventEventType = "poll.read"
	ConnectionRevokedWebhookEventEventTypePollUpdated                ConnectionRevokedWebhookEventEventType = "poll.updated"
	ConnectionRevokedWebhookEventEventTypePollVoteAdded              ConnectionRevokedWebhookEventEventType = "poll.vote.added"
	ConnectionRevokedWebhookEventEventTypePollVoteRemoved            ConnectionRevokedWebhookEventEventType = "poll.vote.removed"
	ConnectionRevokedWebhookEventEventTypePollReactionAdded          ConnectionRevokedWebhookEventEventType = "poll.reaction.added"
	ConnectionRevokedWebhookEventEventTypeParticipantAdded           ConnectionRevokedWebhookEventEventType = "participant.added"
	ConnectionRevokedWebhookEventEventTypeParticipantRemoved         ConnectionRevokedWebhookEventEventType = "participant.removed"
	ConnectionRevokedWebhookEventEventTypeChatCreated                ConnectionRevokedWebhookEventEventType = "chat.created"
	ConnectionRevokedWebhookEventEventTypeChatGroupNameUpdated       ConnectionRevokedWebhookEventEventType = "chat.group_name_updated"
	ConnectionRevokedWebhookEventEventTypeChatGroupIconUpdated       ConnectionRevokedWebhookEventEventType = "chat.group_icon_updated"
	ConnectionRevokedWebhookEventEventTypeChatGroupNameUpdateFailed  ConnectionRevokedWebhookEventEventType = "chat.group_name_update_failed"
	ConnectionRevokedWebhookEventEventTypeChatGroupIconUpdateFailed  ConnectionRevokedWebhookEventEventType = "chat.group_icon_update_failed"
	ConnectionRevokedWebhookEventEventTypeChatBackgroundUpdated      ConnectionRevokedWebhookEventEventType = "chat.background_updated"
	ConnectionRevokedWebhookEventEventTypeChatBackgroundUpdateFailed ConnectionRevokedWebhookEventEventType = "chat.background_update_failed"
	ConnectionRevokedWebhookEventEventTypeChatTypingIndicatorStarted ConnectionRevokedWebhookEventEventType = "chat.typing_indicator.started"
	ConnectionRevokedWebhookEventEventTypeChatTypingIndicatorStopped ConnectionRevokedWebhookEventEventType = "chat.typing_indicator.stopped"
	ConnectionRevokedWebhookEventEventTypePhoneNumberStatusUpdated   ConnectionRevokedWebhookEventEventType = "phone_number.status_updated"
	ConnectionRevokedWebhookEventEventTypeContactCardReceived        ConnectionRevokedWebhookEventEventType = "contact_card.received"
	ConnectionRevokedWebhookEventEventTypeCallInitiated              ConnectionRevokedWebhookEventEventType = "call.initiated"
	ConnectionRevokedWebhookEventEventTypeCallRinging                ConnectionRevokedWebhookEventEventType = "call.ringing"
	ConnectionRevokedWebhookEventEventTypeCallAnswered               ConnectionRevokedWebhookEventEventType = "call.answered"
	ConnectionRevokedWebhookEventEventTypeCallEnded                  ConnectionRevokedWebhookEventEventType = "call.ended"
	ConnectionRevokedWebhookEventEventTypeCallFailed                 ConnectionRevokedWebhookEventEventType = "call.failed"
	ConnectionRevokedWebhookEventEventTypeCallDeclined               ConnectionRevokedWebhookEventEventType = "call.declined"
	ConnectionRevokedWebhookEventEventTypeCallNoAnswer               ConnectionRevokedWebhookEventEventType = "call.no_answer"
	ConnectionRevokedWebhookEventEventTypeLocationSharingStarted     ConnectionRevokedWebhookEventEventType = "location.sharing.started"
	ConnectionRevokedWebhookEventEventTypeLocationSharingStopped     ConnectionRevokedWebhookEventEventType = "location.sharing.stopped"
	ConnectionRevokedWebhookEventEventTypePaymentDeclined            ConnectionRevokedWebhookEventEventType = "payment.declined"
	ConnectionRevokedWebhookEventEventTypePaymentAuthorized          ConnectionRevokedWebhookEventEventType = "payment.authorized"
	ConnectionRevokedWebhookEventEventTypeConnectionCreated          ConnectionRevokedWebhookEventEventType = "connection.created"
	ConnectionRevokedWebhookEventEventTypeConnectionRevoked          ConnectionRevokedWebhookEventEventType = "connection.revoked"
)

type LocationSharingStartedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time                              `json:"created_at" api:"required" format:"date-time"`
	Data      LocationSharingStartedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Any of "location.sharing.started", "message.sent", "message.received",
	// "message.read", "message.delivered", "message.failed", "message.edited",
	// "reaction.added", "reaction.removed", "poll.received", "poll.failed",
	// "poll.sent", "poll.delivered", "poll.read", "poll.updated", "poll.vote.added",
	// "poll.vote.removed", "poll.reaction.added", "participant.added",
	// "participant.removed", "chat.created", "chat.group_name_updated",
	// "chat.group_icon_updated", "chat.group_name_update_failed",
	// "chat.group_icon_update_failed", "chat.background_updated",
	// "chat.background_update_failed", "chat.typing_indicator.started",
	// "chat.typing_indicator.stopped", "phone_number.status_updated",
	// "contact_card.received", "call.initiated", "call.ringing", "call.answered",
	// "call.ended", "call.failed", "call.declined", "call.no_answer",
	// "location.sharing.stopped", "payment.succeeded", "payment.canceled",
	// "payment.expired", "payment.declined", "payment.authorized",
	// "connection.created", "connection.revoked".
	EventType LocationSharingStartedWebhookEventEventType `json:"event_type" api:"required"`
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
func (r LocationSharingStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *LocationSharingStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationSharingStartedWebhookEventData struct {
	// When location sharing started. Always present: falls back to when the share was
	// first observed if the device reported no start time.
	BeganAt time.Time `json:"began_at" api:"required" format:"date-time"`
	// The chat this share was first sent to. Location sharing is per-contact rather
	// than per-chat, so the location may also be visible in other chats with the same
	// handle; this identifies where the share originated and does not change if the
	// contact later shares into another chat. Null when the originating chat could not
	// be determined.
	ChatID string `json:"chat_id" api:"required" format:"uuid"`
	// When location sharing will expire. Null when sharing indefinitely.
	EndsAt time.Time `json:"ends_at" api:"required" format:"date-time"`
	// Phone number of the person sharing their location
	SharedBy string `json:"shared_by" api:"required"`
	// Your phone number receiving the location
	SharedWith string `json:"shared_with" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BeganAt     respjson.Field
		ChatID      respjson.Field
		EndsAt      respjson.Field
		SharedBy    respjson.Field
		SharedWith  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationSharingStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *LocationSharingStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationSharingStartedWebhookEventEventType string

const (
	LocationSharingStartedWebhookEventEventTypeLocationSharingStarted     LocationSharingStartedWebhookEventEventType = "location.sharing.started"
	LocationSharingStartedWebhookEventEventTypeMessageSent                LocationSharingStartedWebhookEventEventType = "message.sent"
	LocationSharingStartedWebhookEventEventTypeMessageReceived            LocationSharingStartedWebhookEventEventType = "message.received"
	LocationSharingStartedWebhookEventEventTypeMessageRead                LocationSharingStartedWebhookEventEventType = "message.read"
	LocationSharingStartedWebhookEventEventTypeMessageDelivered           LocationSharingStartedWebhookEventEventType = "message.delivered"
	LocationSharingStartedWebhookEventEventTypeMessageFailed              LocationSharingStartedWebhookEventEventType = "message.failed"
	LocationSharingStartedWebhookEventEventTypeMessageEdited              LocationSharingStartedWebhookEventEventType = "message.edited"
	LocationSharingStartedWebhookEventEventTypeReactionAdded              LocationSharingStartedWebhookEventEventType = "reaction.added"
	LocationSharingStartedWebhookEventEventTypeReactionRemoved            LocationSharingStartedWebhookEventEventType = "reaction.removed"
	LocationSharingStartedWebhookEventEventTypePollReceived               LocationSharingStartedWebhookEventEventType = "poll.received"
	LocationSharingStartedWebhookEventEventTypePollFailed                 LocationSharingStartedWebhookEventEventType = "poll.failed"
	LocationSharingStartedWebhookEventEventTypePollSent                   LocationSharingStartedWebhookEventEventType = "poll.sent"
	LocationSharingStartedWebhookEventEventTypePollDelivered              LocationSharingStartedWebhookEventEventType = "poll.delivered"
	LocationSharingStartedWebhookEventEventTypePollRead                   LocationSharingStartedWebhookEventEventType = "poll.read"
	LocationSharingStartedWebhookEventEventTypePollUpdated                LocationSharingStartedWebhookEventEventType = "poll.updated"
	LocationSharingStartedWebhookEventEventTypePollVoteAdded              LocationSharingStartedWebhookEventEventType = "poll.vote.added"
	LocationSharingStartedWebhookEventEventTypePollVoteRemoved            LocationSharingStartedWebhookEventEventType = "poll.vote.removed"
	LocationSharingStartedWebhookEventEventTypePollReactionAdded          LocationSharingStartedWebhookEventEventType = "poll.reaction.added"
	LocationSharingStartedWebhookEventEventTypeParticipantAdded           LocationSharingStartedWebhookEventEventType = "participant.added"
	LocationSharingStartedWebhookEventEventTypeParticipantRemoved         LocationSharingStartedWebhookEventEventType = "participant.removed"
	LocationSharingStartedWebhookEventEventTypeChatCreated                LocationSharingStartedWebhookEventEventType = "chat.created"
	LocationSharingStartedWebhookEventEventTypeChatGroupNameUpdated       LocationSharingStartedWebhookEventEventType = "chat.group_name_updated"
	LocationSharingStartedWebhookEventEventTypeChatGroupIconUpdated       LocationSharingStartedWebhookEventEventType = "chat.group_icon_updated"
	LocationSharingStartedWebhookEventEventTypeChatGroupNameUpdateFailed  LocationSharingStartedWebhookEventEventType = "chat.group_name_update_failed"
	LocationSharingStartedWebhookEventEventTypeChatGroupIconUpdateFailed  LocationSharingStartedWebhookEventEventType = "chat.group_icon_update_failed"
	LocationSharingStartedWebhookEventEventTypeChatBackgroundUpdated      LocationSharingStartedWebhookEventEventType = "chat.background_updated"
	LocationSharingStartedWebhookEventEventTypeChatBackgroundUpdateFailed LocationSharingStartedWebhookEventEventType = "chat.background_update_failed"
	LocationSharingStartedWebhookEventEventTypeChatTypingIndicatorStarted LocationSharingStartedWebhookEventEventType = "chat.typing_indicator.started"
	LocationSharingStartedWebhookEventEventTypeChatTypingIndicatorStopped LocationSharingStartedWebhookEventEventType = "chat.typing_indicator.stopped"
	LocationSharingStartedWebhookEventEventTypePhoneNumberStatusUpdated   LocationSharingStartedWebhookEventEventType = "phone_number.status_updated"
	LocationSharingStartedWebhookEventEventTypeContactCardReceived        LocationSharingStartedWebhookEventEventType = "contact_card.received"
	LocationSharingStartedWebhookEventEventTypeCallInitiated              LocationSharingStartedWebhookEventEventType = "call.initiated"
	LocationSharingStartedWebhookEventEventTypeCallRinging                LocationSharingStartedWebhookEventEventType = "call.ringing"
	LocationSharingStartedWebhookEventEventTypeCallAnswered               LocationSharingStartedWebhookEventEventType = "call.answered"
	LocationSharingStartedWebhookEventEventTypeCallEnded                  LocationSharingStartedWebhookEventEventType = "call.ended"
	LocationSharingStartedWebhookEventEventTypeCallFailed                 LocationSharingStartedWebhookEventEventType = "call.failed"
	LocationSharingStartedWebhookEventEventTypeCallDeclined               LocationSharingStartedWebhookEventEventType = "call.declined"
	LocationSharingStartedWebhookEventEventTypeCallNoAnswer               LocationSharingStartedWebhookEventEventType = "call.no_answer"
	LocationSharingStartedWebhookEventEventTypeLocationSharingStopped     LocationSharingStartedWebhookEventEventType = "location.sharing.stopped"
	LocationSharingStartedWebhookEventEventTypePaymentSucceeded           LocationSharingStartedWebhookEventEventType = "payment.succeeded"
	LocationSharingStartedWebhookEventEventTypePaymentCanceled            LocationSharingStartedWebhookEventEventType = "payment.canceled"
	LocationSharingStartedWebhookEventEventTypePaymentExpired             LocationSharingStartedWebhookEventEventType = "payment.expired"
	LocationSharingStartedWebhookEventEventTypePaymentDeclined            LocationSharingStartedWebhookEventEventType = "payment.declined"
	LocationSharingStartedWebhookEventEventTypePaymentAuthorized          LocationSharingStartedWebhookEventEventType = "payment.authorized"
	LocationSharingStartedWebhookEventEventTypeConnectionCreated          LocationSharingStartedWebhookEventEventType = "connection.created"
	LocationSharingStartedWebhookEventEventTypeConnectionRevoked          LocationSharingStartedWebhookEventEventType = "connection.revoked"
)

type LocationSharingStoppedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time                              `json:"created_at" api:"required" format:"date-time"`
	Data      LocationSharingStoppedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Any of "location.sharing.stopped", "message.sent", "message.received",
	// "message.read", "message.delivered", "message.failed", "message.edited",
	// "reaction.added", "reaction.removed", "poll.received", "poll.failed",
	// "poll.sent", "poll.delivered", "poll.read", "poll.updated", "poll.vote.added",
	// "poll.vote.removed", "poll.reaction.added", "participant.added",
	// "participant.removed", "chat.created", "chat.group_name_updated",
	// "chat.group_icon_updated", "chat.group_name_update_failed",
	// "chat.group_icon_update_failed", "chat.background_updated",
	// "chat.background_update_failed", "chat.typing_indicator.started",
	// "chat.typing_indicator.stopped", "phone_number.status_updated",
	// "contact_card.received", "call.initiated", "call.ringing", "call.answered",
	// "call.ended", "call.failed", "call.declined", "call.no_answer",
	// "location.sharing.started", "payment.succeeded", "payment.canceled",
	// "payment.expired", "payment.declined", "payment.authorized",
	// "connection.created", "connection.revoked".
	EventType LocationSharingStoppedWebhookEventEventType `json:"event_type" api:"required"`
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
func (r LocationSharingStoppedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *LocationSharingStoppedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationSharingStoppedWebhookEventData struct {
	// When the sharing session started, matching began_at on its started event. Always
	// present.
	BeganAt time.Time `json:"began_at" api:"required" format:"date-time"`
	// The chat the ended share was first sent to, matching the chat_id on its started
	// event. Sharing always stops for the contact as a whole, never for a single chat,
	// so this is the session's origin rather than the chat it stopped in. Null when
	// the originating chat could not be determined.
	ChatID string `json:"chat_id" api:"required" format:"uuid"`
	// When the sharing session was observed to stop.
	EndedAt time.Time `json:"ended_at" api:"required" format:"date-time"`
	// Phone number of the person who stopped sharing
	SharedBy string `json:"shared_by" api:"required"`
	// Your phone number that was receiving the location
	SharedWith string `json:"shared_with" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BeganAt     respjson.Field
		ChatID      respjson.Field
		EndedAt     respjson.Field
		SharedBy    respjson.Field
		SharedWith  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationSharingStoppedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *LocationSharingStoppedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationSharingStoppedWebhookEventEventType string

const (
	LocationSharingStoppedWebhookEventEventTypeLocationSharingStopped     LocationSharingStoppedWebhookEventEventType = "location.sharing.stopped"
	LocationSharingStoppedWebhookEventEventTypeMessageSent                LocationSharingStoppedWebhookEventEventType = "message.sent"
	LocationSharingStoppedWebhookEventEventTypeMessageReceived            LocationSharingStoppedWebhookEventEventType = "message.received"
	LocationSharingStoppedWebhookEventEventTypeMessageRead                LocationSharingStoppedWebhookEventEventType = "message.read"
	LocationSharingStoppedWebhookEventEventTypeMessageDelivered           LocationSharingStoppedWebhookEventEventType = "message.delivered"
	LocationSharingStoppedWebhookEventEventTypeMessageFailed              LocationSharingStoppedWebhookEventEventType = "message.failed"
	LocationSharingStoppedWebhookEventEventTypeMessageEdited              LocationSharingStoppedWebhookEventEventType = "message.edited"
	LocationSharingStoppedWebhookEventEventTypeReactionAdded              LocationSharingStoppedWebhookEventEventType = "reaction.added"
	LocationSharingStoppedWebhookEventEventTypeReactionRemoved            LocationSharingStoppedWebhookEventEventType = "reaction.removed"
	LocationSharingStoppedWebhookEventEventTypePollReceived               LocationSharingStoppedWebhookEventEventType = "poll.received"
	LocationSharingStoppedWebhookEventEventTypePollFailed                 LocationSharingStoppedWebhookEventEventType = "poll.failed"
	LocationSharingStoppedWebhookEventEventTypePollSent                   LocationSharingStoppedWebhookEventEventType = "poll.sent"
	LocationSharingStoppedWebhookEventEventTypePollDelivered              LocationSharingStoppedWebhookEventEventType = "poll.delivered"
	LocationSharingStoppedWebhookEventEventTypePollRead                   LocationSharingStoppedWebhookEventEventType = "poll.read"
	LocationSharingStoppedWebhookEventEventTypePollUpdated                LocationSharingStoppedWebhookEventEventType = "poll.updated"
	LocationSharingStoppedWebhookEventEventTypePollVoteAdded              LocationSharingStoppedWebhookEventEventType = "poll.vote.added"
	LocationSharingStoppedWebhookEventEventTypePollVoteRemoved            LocationSharingStoppedWebhookEventEventType = "poll.vote.removed"
	LocationSharingStoppedWebhookEventEventTypePollReactionAdded          LocationSharingStoppedWebhookEventEventType = "poll.reaction.added"
	LocationSharingStoppedWebhookEventEventTypeParticipantAdded           LocationSharingStoppedWebhookEventEventType = "participant.added"
	LocationSharingStoppedWebhookEventEventTypeParticipantRemoved         LocationSharingStoppedWebhookEventEventType = "participant.removed"
	LocationSharingStoppedWebhookEventEventTypeChatCreated                LocationSharingStoppedWebhookEventEventType = "chat.created"
	LocationSharingStoppedWebhookEventEventTypeChatGroupNameUpdated       LocationSharingStoppedWebhookEventEventType = "chat.group_name_updated"
	LocationSharingStoppedWebhookEventEventTypeChatGroupIconUpdated       LocationSharingStoppedWebhookEventEventType = "chat.group_icon_updated"
	LocationSharingStoppedWebhookEventEventTypeChatGroupNameUpdateFailed  LocationSharingStoppedWebhookEventEventType = "chat.group_name_update_failed"
	LocationSharingStoppedWebhookEventEventTypeChatGroupIconUpdateFailed  LocationSharingStoppedWebhookEventEventType = "chat.group_icon_update_failed"
	LocationSharingStoppedWebhookEventEventTypeChatBackgroundUpdated      LocationSharingStoppedWebhookEventEventType = "chat.background_updated"
	LocationSharingStoppedWebhookEventEventTypeChatBackgroundUpdateFailed LocationSharingStoppedWebhookEventEventType = "chat.background_update_failed"
	LocationSharingStoppedWebhookEventEventTypeChatTypingIndicatorStarted LocationSharingStoppedWebhookEventEventType = "chat.typing_indicator.started"
	LocationSharingStoppedWebhookEventEventTypeChatTypingIndicatorStopped LocationSharingStoppedWebhookEventEventType = "chat.typing_indicator.stopped"
	LocationSharingStoppedWebhookEventEventTypePhoneNumberStatusUpdated   LocationSharingStoppedWebhookEventEventType = "phone_number.status_updated"
	LocationSharingStoppedWebhookEventEventTypeContactCardReceived        LocationSharingStoppedWebhookEventEventType = "contact_card.received"
	LocationSharingStoppedWebhookEventEventTypeCallInitiated              LocationSharingStoppedWebhookEventEventType = "call.initiated"
	LocationSharingStoppedWebhookEventEventTypeCallRinging                LocationSharingStoppedWebhookEventEventType = "call.ringing"
	LocationSharingStoppedWebhookEventEventTypeCallAnswered               LocationSharingStoppedWebhookEventEventType = "call.answered"
	LocationSharingStoppedWebhookEventEventTypeCallEnded                  LocationSharingStoppedWebhookEventEventType = "call.ended"
	LocationSharingStoppedWebhookEventEventTypeCallFailed                 LocationSharingStoppedWebhookEventEventType = "call.failed"
	LocationSharingStoppedWebhookEventEventTypeCallDeclined               LocationSharingStoppedWebhookEventEventType = "call.declined"
	LocationSharingStoppedWebhookEventEventTypeCallNoAnswer               LocationSharingStoppedWebhookEventEventType = "call.no_answer"
	LocationSharingStoppedWebhookEventEventTypeLocationSharingStarted     LocationSharingStoppedWebhookEventEventType = "location.sharing.started"
	LocationSharingStoppedWebhookEventEventTypePaymentSucceeded           LocationSharingStoppedWebhookEventEventType = "payment.succeeded"
	LocationSharingStoppedWebhookEventEventTypePaymentCanceled            LocationSharingStoppedWebhookEventEventType = "payment.canceled"
	LocationSharingStoppedWebhookEventEventTypePaymentExpired             LocationSharingStoppedWebhookEventEventType = "payment.expired"
	LocationSharingStoppedWebhookEventEventTypePaymentDeclined            LocationSharingStoppedWebhookEventEventType = "payment.declined"
	LocationSharingStoppedWebhookEventEventTypePaymentAuthorized          LocationSharingStoppedWebhookEventEventType = "payment.authorized"
	LocationSharingStoppedWebhookEventEventTypeConnectionCreated          LocationSharingStoppedWebhookEventEventType = "connection.created"
	LocationSharingStoppedWebhookEventEventTypeConnectionRevoked          LocationSharingStoppedWebhookEventEventType = "connection.revoked"
)

type PaymentAuthorizedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The payment request, as returned by
	// `GET /v3/payment_requests/{paymentRequestId}`.
	Data PaymentAuthorizedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Any of "payment.succeeded", "payment.canceled", "payment.expired",
	// "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.declined", "payment.authorized", "connection.created",
	// "connection.revoked".
	EventType PaymentAuthorizedWebhookEventEventType `json:"event_type" api:"required"`
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
func (r PaymentAuthorizedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PaymentAuthorizedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payment request, as returned by
// `GET /v3/payment_requests/{paymentRequestId}`.
type PaymentAuthorizedWebhookEventData struct {
	// The payment request id.
	ID string `json:"id" api:"required" format:"uuid"`
	// What was charged at checkout, in the currency's minor units. In `subscription`
	// mode this is the first invoice's total — all items after any discounts are
	// applied.
	Amount int64 `json:"amount" api:"required"`
	// URL the recipient opens to pay
	// (`https://zero.linqapp.com/pay/{slug}?session=...`).
	CheckoutURL string    `json:"checkout_url" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency    string    `json:"currency" api:"required"`
	Object      string    `json:"object" api:"required"`
	// Any of "succeeded", "failed", "canceled", "expired".
	Status      string `json:"status" api:"required"`
	Description string `json:"description"`
	// Subscription mode — the discount Stripe applied, read back from the coupon.
	// Absent when none was applied.
	Discount PaymentAuthorizedWebhookEventDataDiscount `json:"discount"`
	// Subscription mode — how often the subscription renews.
	//
	// Any of "day", "week", "month", "year".
	Interval string `json:"interval"`
	// Subscription mode — intervals per renewal.
	IntervalCount int64             `json:"interval_count"`
	Metadata      map[string]string `json:"metadata"`
	// Whether the request collected a one-time charge or started a subscription.
	//
	// Any of "payment", "subscription".
	Mode string `json:"mode"`
	// Natural-rail join keys, present when `rail: natural`.
	Natural PaymentAuthorizedWebhookEventDataNatural `json:"natural"`
	// Subscription mode — the recurring price subscribed to.
	PriceID string `json:"price_id"`
	// Subscription mode — units of the price subscribed to.
	Quantity int64 `json:"quantity"`
	// The rail this request settled on.
	//
	// Any of "stripe", "natural".
	Rail string `json:"rail"`
	// Ids of the Stripe objects on your connected account — join keys into your own
	// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
	// `subscription_id`.
	Stripe PaymentAuthorizedWebhookEventDataStripe `json:"stripe"`
	// Subscription mode — when the free trial ends and the first charge happens. On a
	// trial request, `payment.succeeded` means the payment method was collected ($0
	// moved).
	TrialEnd  time.Time `json:"trial_end" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		CheckoutURL   respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Object        respjson.Field
		Status        respjson.Field
		Description   respjson.Field
		Discount      respjson.Field
		Interval      respjson.Field
		IntervalCount respjson.Field
		Metadata      respjson.Field
		Mode          respjson.Field
		Natural       respjson.Field
		PriceID       respjson.Field
		Quantity      respjson.Field
		Rail          respjson.Field
		Stripe        respjson.Field
		TrialEnd      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentAuthorizedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PaymentAuthorizedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription mode — the discount Stripe applied, read back from the coupon.
// Absent when none was applied.
type PaymentAuthorizedWebhookEventDataDiscount struct {
	Coupon string `json:"coupon"`
	// Name of the coupon/promo code displayed to customers.
	Label         string `json:"label"`
	PromotionCode string `json:"promotion_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coupon        respjson.Field
		Label         respjson.Field
		PromotionCode respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentAuthorizedWebhookEventDataDiscount) RawJSON() string { return r.JSON.raw }
func (r *PaymentAuthorizedWebhookEventDataDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Natural-rail join keys, present when `rail: natural`.
type PaymentAuthorizedWebhookEventDataNatural struct {
	// The Natural payment request (`prq_...`).
	PaymentRequestID string `json:"payment_request_id"`
	// The settled transaction (`txn_...`).
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentRequestID respjson.Field
		TransactionID    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentAuthorizedWebhookEventDataNatural) RawJSON() string { return r.JSON.raw }
func (r *PaymentAuthorizedWebhookEventDataNatural) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ids of the Stripe objects on your connected account — join keys into your own
// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
// `subscription_id`.
type PaymentAuthorizedWebhookEventDataStripe struct {
	// The Customer the request is attached to (`cus_...`). Always set in subscription
	// mode; set in payment mode only when the request was created with a
	// `customer_id`.
	CustomerID string `json:"customer_id"`
	// The PaymentIntent collected at checkout (`pi_...`).
	PaymentIntentID string `json:"payment_intent_id"`
	// Subscription mode — the Subscription (`sub_...`).
	SubscriptionID string `json:"subscription_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID      respjson.Field
		PaymentIntentID respjson.Field
		SubscriptionID  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentAuthorizedWebhookEventDataStripe) RawJSON() string { return r.JSON.raw }
func (r *PaymentAuthorizedWebhookEventDataStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentAuthorizedWebhookEventEventType string

const (
	PaymentAuthorizedWebhookEventEventTypePaymentSucceeded           PaymentAuthorizedWebhookEventEventType = "payment.succeeded"
	PaymentAuthorizedWebhookEventEventTypePaymentCanceled            PaymentAuthorizedWebhookEventEventType = "payment.canceled"
	PaymentAuthorizedWebhookEventEventTypePaymentExpired             PaymentAuthorizedWebhookEventEventType = "payment.expired"
	PaymentAuthorizedWebhookEventEventTypeMessageSent                PaymentAuthorizedWebhookEventEventType = "message.sent"
	PaymentAuthorizedWebhookEventEventTypeMessageReceived            PaymentAuthorizedWebhookEventEventType = "message.received"
	PaymentAuthorizedWebhookEventEventTypeMessageRead                PaymentAuthorizedWebhookEventEventType = "message.read"
	PaymentAuthorizedWebhookEventEventTypeMessageDelivered           PaymentAuthorizedWebhookEventEventType = "message.delivered"
	PaymentAuthorizedWebhookEventEventTypeMessageFailed              PaymentAuthorizedWebhookEventEventType = "message.failed"
	PaymentAuthorizedWebhookEventEventTypeMessageEdited              PaymentAuthorizedWebhookEventEventType = "message.edited"
	PaymentAuthorizedWebhookEventEventTypeReactionAdded              PaymentAuthorizedWebhookEventEventType = "reaction.added"
	PaymentAuthorizedWebhookEventEventTypeReactionRemoved            PaymentAuthorizedWebhookEventEventType = "reaction.removed"
	PaymentAuthorizedWebhookEventEventTypePollReceived               PaymentAuthorizedWebhookEventEventType = "poll.received"
	PaymentAuthorizedWebhookEventEventTypePollFailed                 PaymentAuthorizedWebhookEventEventType = "poll.failed"
	PaymentAuthorizedWebhookEventEventTypePollSent                   PaymentAuthorizedWebhookEventEventType = "poll.sent"
	PaymentAuthorizedWebhookEventEventTypePollDelivered              PaymentAuthorizedWebhookEventEventType = "poll.delivered"
	PaymentAuthorizedWebhookEventEventTypePollRead                   PaymentAuthorizedWebhookEventEventType = "poll.read"
	PaymentAuthorizedWebhookEventEventTypePollUpdated                PaymentAuthorizedWebhookEventEventType = "poll.updated"
	PaymentAuthorizedWebhookEventEventTypePollVoteAdded              PaymentAuthorizedWebhookEventEventType = "poll.vote.added"
	PaymentAuthorizedWebhookEventEventTypePollVoteRemoved            PaymentAuthorizedWebhookEventEventType = "poll.vote.removed"
	PaymentAuthorizedWebhookEventEventTypePollReactionAdded          PaymentAuthorizedWebhookEventEventType = "poll.reaction.added"
	PaymentAuthorizedWebhookEventEventTypeParticipantAdded           PaymentAuthorizedWebhookEventEventType = "participant.added"
	PaymentAuthorizedWebhookEventEventTypeParticipantRemoved         PaymentAuthorizedWebhookEventEventType = "participant.removed"
	PaymentAuthorizedWebhookEventEventTypeChatCreated                PaymentAuthorizedWebhookEventEventType = "chat.created"
	PaymentAuthorizedWebhookEventEventTypeChatGroupNameUpdated       PaymentAuthorizedWebhookEventEventType = "chat.group_name_updated"
	PaymentAuthorizedWebhookEventEventTypeChatGroupIconUpdated       PaymentAuthorizedWebhookEventEventType = "chat.group_icon_updated"
	PaymentAuthorizedWebhookEventEventTypeChatGroupNameUpdateFailed  PaymentAuthorizedWebhookEventEventType = "chat.group_name_update_failed"
	PaymentAuthorizedWebhookEventEventTypeChatGroupIconUpdateFailed  PaymentAuthorizedWebhookEventEventType = "chat.group_icon_update_failed"
	PaymentAuthorizedWebhookEventEventTypeChatBackgroundUpdated      PaymentAuthorizedWebhookEventEventType = "chat.background_updated"
	PaymentAuthorizedWebhookEventEventTypeChatBackgroundUpdateFailed PaymentAuthorizedWebhookEventEventType = "chat.background_update_failed"
	PaymentAuthorizedWebhookEventEventTypeChatTypingIndicatorStarted PaymentAuthorizedWebhookEventEventType = "chat.typing_indicator.started"
	PaymentAuthorizedWebhookEventEventTypeChatTypingIndicatorStopped PaymentAuthorizedWebhookEventEventType = "chat.typing_indicator.stopped"
	PaymentAuthorizedWebhookEventEventTypePhoneNumberStatusUpdated   PaymentAuthorizedWebhookEventEventType = "phone_number.status_updated"
	PaymentAuthorizedWebhookEventEventTypeContactCardReceived        PaymentAuthorizedWebhookEventEventType = "contact_card.received"
	PaymentAuthorizedWebhookEventEventTypeCallInitiated              PaymentAuthorizedWebhookEventEventType = "call.initiated"
	PaymentAuthorizedWebhookEventEventTypeCallRinging                PaymentAuthorizedWebhookEventEventType = "call.ringing"
	PaymentAuthorizedWebhookEventEventTypeCallAnswered               PaymentAuthorizedWebhookEventEventType = "call.answered"
	PaymentAuthorizedWebhookEventEventTypeCallEnded                  PaymentAuthorizedWebhookEventEventType = "call.ended"
	PaymentAuthorizedWebhookEventEventTypeCallFailed                 PaymentAuthorizedWebhookEventEventType = "call.failed"
	PaymentAuthorizedWebhookEventEventTypeCallDeclined               PaymentAuthorizedWebhookEventEventType = "call.declined"
	PaymentAuthorizedWebhookEventEventTypeCallNoAnswer               PaymentAuthorizedWebhookEventEventType = "call.no_answer"
	PaymentAuthorizedWebhookEventEventTypeLocationSharingStarted     PaymentAuthorizedWebhookEventEventType = "location.sharing.started"
	PaymentAuthorizedWebhookEventEventTypeLocationSharingStopped     PaymentAuthorizedWebhookEventEventType = "location.sharing.stopped"
	PaymentAuthorizedWebhookEventEventTypePaymentDeclined            PaymentAuthorizedWebhookEventEventType = "payment.declined"
	PaymentAuthorizedWebhookEventEventTypePaymentAuthorized          PaymentAuthorizedWebhookEventEventType = "payment.authorized"
	PaymentAuthorizedWebhookEventEventTypeConnectionCreated          PaymentAuthorizedWebhookEventEventType = "connection.created"
	PaymentAuthorizedWebhookEventEventTypeConnectionRevoked          PaymentAuthorizedWebhookEventEventType = "connection.revoked"
)

type PaymentCanceledWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The payment request, as returned by
	// `GET /v3/payment_requests/{paymentRequestId}`.
	Data PaymentCanceledWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Any of "payment.succeeded", "payment.canceled", "payment.expired",
	// "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.declined", "payment.authorized", "connection.created",
	// "connection.revoked".
	EventType PaymentCanceledWebhookEventEventType `json:"event_type" api:"required"`
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
func (r PaymentCanceledWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PaymentCanceledWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payment request, as returned by
// `GET /v3/payment_requests/{paymentRequestId}`.
type PaymentCanceledWebhookEventData struct {
	// The payment request id.
	ID string `json:"id" api:"required" format:"uuid"`
	// What was charged at checkout, in the currency's minor units. In `subscription`
	// mode this is the first invoice's total — all items after any discounts are
	// applied.
	Amount int64 `json:"amount" api:"required"`
	// URL the recipient opens to pay
	// (`https://zero.linqapp.com/pay/{slug}?session=...`).
	CheckoutURL string    `json:"checkout_url" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency    string    `json:"currency" api:"required"`
	Object      string    `json:"object" api:"required"`
	// Any of "succeeded", "failed", "canceled", "expired".
	Status      string `json:"status" api:"required"`
	Description string `json:"description"`
	// Subscription mode — the discount Stripe applied, read back from the coupon.
	// Absent when none was applied.
	Discount PaymentCanceledWebhookEventDataDiscount `json:"discount"`
	// Subscription mode — how often the subscription renews.
	//
	// Any of "day", "week", "month", "year".
	Interval string `json:"interval"`
	// Subscription mode — intervals per renewal.
	IntervalCount int64             `json:"interval_count"`
	Metadata      map[string]string `json:"metadata"`
	// Whether the request collected a one-time charge or started a subscription.
	//
	// Any of "payment", "subscription".
	Mode string `json:"mode"`
	// Natural-rail join keys, present when `rail: natural`.
	Natural PaymentCanceledWebhookEventDataNatural `json:"natural"`
	// Subscription mode — the recurring price subscribed to.
	PriceID string `json:"price_id"`
	// Subscription mode — units of the price subscribed to.
	Quantity int64 `json:"quantity"`
	// The rail this request settled on.
	//
	// Any of "stripe", "natural".
	Rail string `json:"rail"`
	// Ids of the Stripe objects on your connected account — join keys into your own
	// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
	// `subscription_id`.
	Stripe PaymentCanceledWebhookEventDataStripe `json:"stripe"`
	// Subscription mode — when the free trial ends and the first charge happens. On a
	// trial request, `payment.succeeded` means the payment method was collected ($0
	// moved).
	TrialEnd  time.Time `json:"trial_end" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		CheckoutURL   respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Object        respjson.Field
		Status        respjson.Field
		Description   respjson.Field
		Discount      respjson.Field
		Interval      respjson.Field
		IntervalCount respjson.Field
		Metadata      respjson.Field
		Mode          respjson.Field
		Natural       respjson.Field
		PriceID       respjson.Field
		Quantity      respjson.Field
		Rail          respjson.Field
		Stripe        respjson.Field
		TrialEnd      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentCanceledWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PaymentCanceledWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription mode — the discount Stripe applied, read back from the coupon.
// Absent when none was applied.
type PaymentCanceledWebhookEventDataDiscount struct {
	Coupon string `json:"coupon"`
	// Name of the coupon/promo code displayed to customers.
	Label         string `json:"label"`
	PromotionCode string `json:"promotion_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coupon        respjson.Field
		Label         respjson.Field
		PromotionCode respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentCanceledWebhookEventDataDiscount) RawJSON() string { return r.JSON.raw }
func (r *PaymentCanceledWebhookEventDataDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Natural-rail join keys, present when `rail: natural`.
type PaymentCanceledWebhookEventDataNatural struct {
	// The Natural payment request (`prq_...`).
	PaymentRequestID string `json:"payment_request_id"`
	// The settled transaction (`txn_...`).
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentRequestID respjson.Field
		TransactionID    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentCanceledWebhookEventDataNatural) RawJSON() string { return r.JSON.raw }
func (r *PaymentCanceledWebhookEventDataNatural) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ids of the Stripe objects on your connected account — join keys into your own
// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
// `subscription_id`.
type PaymentCanceledWebhookEventDataStripe struct {
	// The Customer the request is attached to (`cus_...`). Always set in subscription
	// mode; set in payment mode only when the request was created with a
	// `customer_id`.
	CustomerID string `json:"customer_id"`
	// The PaymentIntent collected at checkout (`pi_...`).
	PaymentIntentID string `json:"payment_intent_id"`
	// Subscription mode — the Subscription (`sub_...`).
	SubscriptionID string `json:"subscription_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID      respjson.Field
		PaymentIntentID respjson.Field
		SubscriptionID  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentCanceledWebhookEventDataStripe) RawJSON() string { return r.JSON.raw }
func (r *PaymentCanceledWebhookEventDataStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentCanceledWebhookEventEventType string

const (
	PaymentCanceledWebhookEventEventTypePaymentSucceeded           PaymentCanceledWebhookEventEventType = "payment.succeeded"
	PaymentCanceledWebhookEventEventTypePaymentCanceled            PaymentCanceledWebhookEventEventType = "payment.canceled"
	PaymentCanceledWebhookEventEventTypePaymentExpired             PaymentCanceledWebhookEventEventType = "payment.expired"
	PaymentCanceledWebhookEventEventTypeMessageSent                PaymentCanceledWebhookEventEventType = "message.sent"
	PaymentCanceledWebhookEventEventTypeMessageReceived            PaymentCanceledWebhookEventEventType = "message.received"
	PaymentCanceledWebhookEventEventTypeMessageRead                PaymentCanceledWebhookEventEventType = "message.read"
	PaymentCanceledWebhookEventEventTypeMessageDelivered           PaymentCanceledWebhookEventEventType = "message.delivered"
	PaymentCanceledWebhookEventEventTypeMessageFailed              PaymentCanceledWebhookEventEventType = "message.failed"
	PaymentCanceledWebhookEventEventTypeMessageEdited              PaymentCanceledWebhookEventEventType = "message.edited"
	PaymentCanceledWebhookEventEventTypeReactionAdded              PaymentCanceledWebhookEventEventType = "reaction.added"
	PaymentCanceledWebhookEventEventTypeReactionRemoved            PaymentCanceledWebhookEventEventType = "reaction.removed"
	PaymentCanceledWebhookEventEventTypePollReceived               PaymentCanceledWebhookEventEventType = "poll.received"
	PaymentCanceledWebhookEventEventTypePollFailed                 PaymentCanceledWebhookEventEventType = "poll.failed"
	PaymentCanceledWebhookEventEventTypePollSent                   PaymentCanceledWebhookEventEventType = "poll.sent"
	PaymentCanceledWebhookEventEventTypePollDelivered              PaymentCanceledWebhookEventEventType = "poll.delivered"
	PaymentCanceledWebhookEventEventTypePollRead                   PaymentCanceledWebhookEventEventType = "poll.read"
	PaymentCanceledWebhookEventEventTypePollUpdated                PaymentCanceledWebhookEventEventType = "poll.updated"
	PaymentCanceledWebhookEventEventTypePollVoteAdded              PaymentCanceledWebhookEventEventType = "poll.vote.added"
	PaymentCanceledWebhookEventEventTypePollVoteRemoved            PaymentCanceledWebhookEventEventType = "poll.vote.removed"
	PaymentCanceledWebhookEventEventTypePollReactionAdded          PaymentCanceledWebhookEventEventType = "poll.reaction.added"
	PaymentCanceledWebhookEventEventTypeParticipantAdded           PaymentCanceledWebhookEventEventType = "participant.added"
	PaymentCanceledWebhookEventEventTypeParticipantRemoved         PaymentCanceledWebhookEventEventType = "participant.removed"
	PaymentCanceledWebhookEventEventTypeChatCreated                PaymentCanceledWebhookEventEventType = "chat.created"
	PaymentCanceledWebhookEventEventTypeChatGroupNameUpdated       PaymentCanceledWebhookEventEventType = "chat.group_name_updated"
	PaymentCanceledWebhookEventEventTypeChatGroupIconUpdated       PaymentCanceledWebhookEventEventType = "chat.group_icon_updated"
	PaymentCanceledWebhookEventEventTypeChatGroupNameUpdateFailed  PaymentCanceledWebhookEventEventType = "chat.group_name_update_failed"
	PaymentCanceledWebhookEventEventTypeChatGroupIconUpdateFailed  PaymentCanceledWebhookEventEventType = "chat.group_icon_update_failed"
	PaymentCanceledWebhookEventEventTypeChatBackgroundUpdated      PaymentCanceledWebhookEventEventType = "chat.background_updated"
	PaymentCanceledWebhookEventEventTypeChatBackgroundUpdateFailed PaymentCanceledWebhookEventEventType = "chat.background_update_failed"
	PaymentCanceledWebhookEventEventTypeChatTypingIndicatorStarted PaymentCanceledWebhookEventEventType = "chat.typing_indicator.started"
	PaymentCanceledWebhookEventEventTypeChatTypingIndicatorStopped PaymentCanceledWebhookEventEventType = "chat.typing_indicator.stopped"
	PaymentCanceledWebhookEventEventTypePhoneNumberStatusUpdated   PaymentCanceledWebhookEventEventType = "phone_number.status_updated"
	PaymentCanceledWebhookEventEventTypeContactCardReceived        PaymentCanceledWebhookEventEventType = "contact_card.received"
	PaymentCanceledWebhookEventEventTypeCallInitiated              PaymentCanceledWebhookEventEventType = "call.initiated"
	PaymentCanceledWebhookEventEventTypeCallRinging                PaymentCanceledWebhookEventEventType = "call.ringing"
	PaymentCanceledWebhookEventEventTypeCallAnswered               PaymentCanceledWebhookEventEventType = "call.answered"
	PaymentCanceledWebhookEventEventTypeCallEnded                  PaymentCanceledWebhookEventEventType = "call.ended"
	PaymentCanceledWebhookEventEventTypeCallFailed                 PaymentCanceledWebhookEventEventType = "call.failed"
	PaymentCanceledWebhookEventEventTypeCallDeclined               PaymentCanceledWebhookEventEventType = "call.declined"
	PaymentCanceledWebhookEventEventTypeCallNoAnswer               PaymentCanceledWebhookEventEventType = "call.no_answer"
	PaymentCanceledWebhookEventEventTypeLocationSharingStarted     PaymentCanceledWebhookEventEventType = "location.sharing.started"
	PaymentCanceledWebhookEventEventTypeLocationSharingStopped     PaymentCanceledWebhookEventEventType = "location.sharing.stopped"
	PaymentCanceledWebhookEventEventTypePaymentDeclined            PaymentCanceledWebhookEventEventType = "payment.declined"
	PaymentCanceledWebhookEventEventTypePaymentAuthorized          PaymentCanceledWebhookEventEventType = "payment.authorized"
	PaymentCanceledWebhookEventEventTypeConnectionCreated          PaymentCanceledWebhookEventEventType = "connection.created"
	PaymentCanceledWebhookEventEventTypeConnectionRevoked          PaymentCanceledWebhookEventEventType = "connection.revoked"
)

type PaymentDeclinedWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The payment request, as returned by
	// `GET /v3/payment_requests/{paymentRequestId}`.
	Data PaymentDeclinedWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Any of "payment.succeeded", "payment.canceled", "payment.expired",
	// "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.declined", "payment.authorized", "connection.created",
	// "connection.revoked".
	EventType PaymentDeclinedWebhookEventEventType `json:"event_type" api:"required"`
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
func (r PaymentDeclinedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PaymentDeclinedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payment request, as returned by
// `GET /v3/payment_requests/{paymentRequestId}`.
type PaymentDeclinedWebhookEventData struct {
	// The payment request id.
	ID string `json:"id" api:"required" format:"uuid"`
	// What was charged at checkout, in the currency's minor units. In `subscription`
	// mode this is the first invoice's total — all items after any discounts are
	// applied.
	Amount int64 `json:"amount" api:"required"`
	// URL the recipient opens to pay
	// (`https://zero.linqapp.com/pay/{slug}?session=...`).
	CheckoutURL string    `json:"checkout_url" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency    string    `json:"currency" api:"required"`
	Object      string    `json:"object" api:"required"`
	// Any of "succeeded", "failed", "canceled", "expired".
	Status      string `json:"status" api:"required"`
	Description string `json:"description"`
	// Subscription mode — the discount Stripe applied, read back from the coupon.
	// Absent when none was applied.
	Discount PaymentDeclinedWebhookEventDataDiscount `json:"discount"`
	// Subscription mode — how often the subscription renews.
	//
	// Any of "day", "week", "month", "year".
	Interval string `json:"interval"`
	// Subscription mode — intervals per renewal.
	IntervalCount int64             `json:"interval_count"`
	Metadata      map[string]string `json:"metadata"`
	// Whether the request collected a one-time charge or started a subscription.
	//
	// Any of "payment", "subscription".
	Mode string `json:"mode"`
	// Natural-rail join keys, present when `rail: natural`.
	Natural PaymentDeclinedWebhookEventDataNatural `json:"natural"`
	// Subscription mode — the recurring price subscribed to.
	PriceID string `json:"price_id"`
	// Subscription mode — units of the price subscribed to.
	Quantity int64 `json:"quantity"`
	// The rail this request settled on.
	//
	// Any of "stripe", "natural".
	Rail string `json:"rail"`
	// Ids of the Stripe objects on your connected account — join keys into your own
	// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
	// `subscription_id`.
	Stripe PaymentDeclinedWebhookEventDataStripe `json:"stripe"`
	// Subscription mode — when the free trial ends and the first charge happens. On a
	// trial request, `payment.succeeded` means the payment method was collected ($0
	// moved).
	TrialEnd  time.Time `json:"trial_end" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		CheckoutURL   respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Object        respjson.Field
		Status        respjson.Field
		Description   respjson.Field
		Discount      respjson.Field
		Interval      respjson.Field
		IntervalCount respjson.Field
		Metadata      respjson.Field
		Mode          respjson.Field
		Natural       respjson.Field
		PriceID       respjson.Field
		Quantity      respjson.Field
		Rail          respjson.Field
		Stripe        respjson.Field
		TrialEnd      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentDeclinedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PaymentDeclinedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription mode — the discount Stripe applied, read back from the coupon.
// Absent when none was applied.
type PaymentDeclinedWebhookEventDataDiscount struct {
	Coupon string `json:"coupon"`
	// Name of the coupon/promo code displayed to customers.
	Label         string `json:"label"`
	PromotionCode string `json:"promotion_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coupon        respjson.Field
		Label         respjson.Field
		PromotionCode respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentDeclinedWebhookEventDataDiscount) RawJSON() string { return r.JSON.raw }
func (r *PaymentDeclinedWebhookEventDataDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Natural-rail join keys, present when `rail: natural`.
type PaymentDeclinedWebhookEventDataNatural struct {
	// The Natural payment request (`prq_...`).
	PaymentRequestID string `json:"payment_request_id"`
	// The settled transaction (`txn_...`).
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentRequestID respjson.Field
		TransactionID    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentDeclinedWebhookEventDataNatural) RawJSON() string { return r.JSON.raw }
func (r *PaymentDeclinedWebhookEventDataNatural) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ids of the Stripe objects on your connected account — join keys into your own
// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
// `subscription_id`.
type PaymentDeclinedWebhookEventDataStripe struct {
	// The Customer the request is attached to (`cus_...`). Always set in subscription
	// mode; set in payment mode only when the request was created with a
	// `customer_id`.
	CustomerID string `json:"customer_id"`
	// The PaymentIntent collected at checkout (`pi_...`).
	PaymentIntentID string `json:"payment_intent_id"`
	// Subscription mode — the Subscription (`sub_...`).
	SubscriptionID string `json:"subscription_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID      respjson.Field
		PaymentIntentID respjson.Field
		SubscriptionID  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentDeclinedWebhookEventDataStripe) RawJSON() string { return r.JSON.raw }
func (r *PaymentDeclinedWebhookEventDataStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentDeclinedWebhookEventEventType string

const (
	PaymentDeclinedWebhookEventEventTypePaymentSucceeded           PaymentDeclinedWebhookEventEventType = "payment.succeeded"
	PaymentDeclinedWebhookEventEventTypePaymentCanceled            PaymentDeclinedWebhookEventEventType = "payment.canceled"
	PaymentDeclinedWebhookEventEventTypePaymentExpired             PaymentDeclinedWebhookEventEventType = "payment.expired"
	PaymentDeclinedWebhookEventEventTypeMessageSent                PaymentDeclinedWebhookEventEventType = "message.sent"
	PaymentDeclinedWebhookEventEventTypeMessageReceived            PaymentDeclinedWebhookEventEventType = "message.received"
	PaymentDeclinedWebhookEventEventTypeMessageRead                PaymentDeclinedWebhookEventEventType = "message.read"
	PaymentDeclinedWebhookEventEventTypeMessageDelivered           PaymentDeclinedWebhookEventEventType = "message.delivered"
	PaymentDeclinedWebhookEventEventTypeMessageFailed              PaymentDeclinedWebhookEventEventType = "message.failed"
	PaymentDeclinedWebhookEventEventTypeMessageEdited              PaymentDeclinedWebhookEventEventType = "message.edited"
	PaymentDeclinedWebhookEventEventTypeReactionAdded              PaymentDeclinedWebhookEventEventType = "reaction.added"
	PaymentDeclinedWebhookEventEventTypeReactionRemoved            PaymentDeclinedWebhookEventEventType = "reaction.removed"
	PaymentDeclinedWebhookEventEventTypePollReceived               PaymentDeclinedWebhookEventEventType = "poll.received"
	PaymentDeclinedWebhookEventEventTypePollFailed                 PaymentDeclinedWebhookEventEventType = "poll.failed"
	PaymentDeclinedWebhookEventEventTypePollSent                   PaymentDeclinedWebhookEventEventType = "poll.sent"
	PaymentDeclinedWebhookEventEventTypePollDelivered              PaymentDeclinedWebhookEventEventType = "poll.delivered"
	PaymentDeclinedWebhookEventEventTypePollRead                   PaymentDeclinedWebhookEventEventType = "poll.read"
	PaymentDeclinedWebhookEventEventTypePollUpdated                PaymentDeclinedWebhookEventEventType = "poll.updated"
	PaymentDeclinedWebhookEventEventTypePollVoteAdded              PaymentDeclinedWebhookEventEventType = "poll.vote.added"
	PaymentDeclinedWebhookEventEventTypePollVoteRemoved            PaymentDeclinedWebhookEventEventType = "poll.vote.removed"
	PaymentDeclinedWebhookEventEventTypePollReactionAdded          PaymentDeclinedWebhookEventEventType = "poll.reaction.added"
	PaymentDeclinedWebhookEventEventTypeParticipantAdded           PaymentDeclinedWebhookEventEventType = "participant.added"
	PaymentDeclinedWebhookEventEventTypeParticipantRemoved         PaymentDeclinedWebhookEventEventType = "participant.removed"
	PaymentDeclinedWebhookEventEventTypeChatCreated                PaymentDeclinedWebhookEventEventType = "chat.created"
	PaymentDeclinedWebhookEventEventTypeChatGroupNameUpdated       PaymentDeclinedWebhookEventEventType = "chat.group_name_updated"
	PaymentDeclinedWebhookEventEventTypeChatGroupIconUpdated       PaymentDeclinedWebhookEventEventType = "chat.group_icon_updated"
	PaymentDeclinedWebhookEventEventTypeChatGroupNameUpdateFailed  PaymentDeclinedWebhookEventEventType = "chat.group_name_update_failed"
	PaymentDeclinedWebhookEventEventTypeChatGroupIconUpdateFailed  PaymentDeclinedWebhookEventEventType = "chat.group_icon_update_failed"
	PaymentDeclinedWebhookEventEventTypeChatBackgroundUpdated      PaymentDeclinedWebhookEventEventType = "chat.background_updated"
	PaymentDeclinedWebhookEventEventTypeChatBackgroundUpdateFailed PaymentDeclinedWebhookEventEventType = "chat.background_update_failed"
	PaymentDeclinedWebhookEventEventTypeChatTypingIndicatorStarted PaymentDeclinedWebhookEventEventType = "chat.typing_indicator.started"
	PaymentDeclinedWebhookEventEventTypeChatTypingIndicatorStopped PaymentDeclinedWebhookEventEventType = "chat.typing_indicator.stopped"
	PaymentDeclinedWebhookEventEventTypePhoneNumberStatusUpdated   PaymentDeclinedWebhookEventEventType = "phone_number.status_updated"
	PaymentDeclinedWebhookEventEventTypeContactCardReceived        PaymentDeclinedWebhookEventEventType = "contact_card.received"
	PaymentDeclinedWebhookEventEventTypeCallInitiated              PaymentDeclinedWebhookEventEventType = "call.initiated"
	PaymentDeclinedWebhookEventEventTypeCallRinging                PaymentDeclinedWebhookEventEventType = "call.ringing"
	PaymentDeclinedWebhookEventEventTypeCallAnswered               PaymentDeclinedWebhookEventEventType = "call.answered"
	PaymentDeclinedWebhookEventEventTypeCallEnded                  PaymentDeclinedWebhookEventEventType = "call.ended"
	PaymentDeclinedWebhookEventEventTypeCallFailed                 PaymentDeclinedWebhookEventEventType = "call.failed"
	PaymentDeclinedWebhookEventEventTypeCallDeclined               PaymentDeclinedWebhookEventEventType = "call.declined"
	PaymentDeclinedWebhookEventEventTypeCallNoAnswer               PaymentDeclinedWebhookEventEventType = "call.no_answer"
	PaymentDeclinedWebhookEventEventTypeLocationSharingStarted     PaymentDeclinedWebhookEventEventType = "location.sharing.started"
	PaymentDeclinedWebhookEventEventTypeLocationSharingStopped     PaymentDeclinedWebhookEventEventType = "location.sharing.stopped"
	PaymentDeclinedWebhookEventEventTypePaymentDeclined            PaymentDeclinedWebhookEventEventType = "payment.declined"
	PaymentDeclinedWebhookEventEventTypePaymentAuthorized          PaymentDeclinedWebhookEventEventType = "payment.authorized"
	PaymentDeclinedWebhookEventEventTypeConnectionCreated          PaymentDeclinedWebhookEventEventType = "connection.created"
	PaymentDeclinedWebhookEventEventTypeConnectionRevoked          PaymentDeclinedWebhookEventEventType = "connection.revoked"
)

type PaymentExpiredWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The payment request, as returned by
	// `GET /v3/payment_requests/{paymentRequestId}`.
	Data PaymentExpiredWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Any of "payment.succeeded", "payment.canceled", "payment.expired",
	// "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.declined", "payment.authorized", "connection.created",
	// "connection.revoked".
	EventType PaymentExpiredWebhookEventEventType `json:"event_type" api:"required"`
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
func (r PaymentExpiredWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PaymentExpiredWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payment request, as returned by
// `GET /v3/payment_requests/{paymentRequestId}`.
type PaymentExpiredWebhookEventData struct {
	// The payment request id.
	ID string `json:"id" api:"required" format:"uuid"`
	// What was charged at checkout, in the currency's minor units. In `subscription`
	// mode this is the first invoice's total — all items after any discounts are
	// applied.
	Amount int64 `json:"amount" api:"required"`
	// URL the recipient opens to pay
	// (`https://zero.linqapp.com/pay/{slug}?session=...`).
	CheckoutURL string    `json:"checkout_url" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency    string    `json:"currency" api:"required"`
	Object      string    `json:"object" api:"required"`
	// Any of "succeeded", "failed", "canceled", "expired".
	Status      string `json:"status" api:"required"`
	Description string `json:"description"`
	// Subscription mode — the discount Stripe applied, read back from the coupon.
	// Absent when none was applied.
	Discount PaymentExpiredWebhookEventDataDiscount `json:"discount"`
	// Subscription mode — how often the subscription renews.
	//
	// Any of "day", "week", "month", "year".
	Interval string `json:"interval"`
	// Subscription mode — intervals per renewal.
	IntervalCount int64             `json:"interval_count"`
	Metadata      map[string]string `json:"metadata"`
	// Whether the request collected a one-time charge or started a subscription.
	//
	// Any of "payment", "subscription".
	Mode string `json:"mode"`
	// Natural-rail join keys, present when `rail: natural`.
	Natural PaymentExpiredWebhookEventDataNatural `json:"natural"`
	// Subscription mode — the recurring price subscribed to.
	PriceID string `json:"price_id"`
	// Subscription mode — units of the price subscribed to.
	Quantity int64 `json:"quantity"`
	// The rail this request settled on.
	//
	// Any of "stripe", "natural".
	Rail string `json:"rail"`
	// Ids of the Stripe objects on your connected account — join keys into your own
	// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
	// `subscription_id`.
	Stripe PaymentExpiredWebhookEventDataStripe `json:"stripe"`
	// Subscription mode — when the free trial ends and the first charge happens. On a
	// trial request, `payment.succeeded` means the payment method was collected ($0
	// moved).
	TrialEnd  time.Time `json:"trial_end" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		CheckoutURL   respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Object        respjson.Field
		Status        respjson.Field
		Description   respjson.Field
		Discount      respjson.Field
		Interval      respjson.Field
		IntervalCount respjson.Field
		Metadata      respjson.Field
		Mode          respjson.Field
		Natural       respjson.Field
		PriceID       respjson.Field
		Quantity      respjson.Field
		Rail          respjson.Field
		Stripe        respjson.Field
		TrialEnd      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentExpiredWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PaymentExpiredWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription mode — the discount Stripe applied, read back from the coupon.
// Absent when none was applied.
type PaymentExpiredWebhookEventDataDiscount struct {
	Coupon string `json:"coupon"`
	// Name of the coupon/promo code displayed to customers.
	Label         string `json:"label"`
	PromotionCode string `json:"promotion_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coupon        respjson.Field
		Label         respjson.Field
		PromotionCode respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentExpiredWebhookEventDataDiscount) RawJSON() string { return r.JSON.raw }
func (r *PaymentExpiredWebhookEventDataDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Natural-rail join keys, present when `rail: natural`.
type PaymentExpiredWebhookEventDataNatural struct {
	// The Natural payment request (`prq_...`).
	PaymentRequestID string `json:"payment_request_id"`
	// The settled transaction (`txn_...`).
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentRequestID respjson.Field
		TransactionID    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentExpiredWebhookEventDataNatural) RawJSON() string { return r.JSON.raw }
func (r *PaymentExpiredWebhookEventDataNatural) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ids of the Stripe objects on your connected account — join keys into your own
// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
// `subscription_id`.
type PaymentExpiredWebhookEventDataStripe struct {
	// The Customer the request is attached to (`cus_...`). Always set in subscription
	// mode; set in payment mode only when the request was created with a
	// `customer_id`.
	CustomerID string `json:"customer_id"`
	// The PaymentIntent collected at checkout (`pi_...`).
	PaymentIntentID string `json:"payment_intent_id"`
	// Subscription mode — the Subscription (`sub_...`).
	SubscriptionID string `json:"subscription_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID      respjson.Field
		PaymentIntentID respjson.Field
		SubscriptionID  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentExpiredWebhookEventDataStripe) RawJSON() string { return r.JSON.raw }
func (r *PaymentExpiredWebhookEventDataStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentExpiredWebhookEventEventType string

const (
	PaymentExpiredWebhookEventEventTypePaymentSucceeded           PaymentExpiredWebhookEventEventType = "payment.succeeded"
	PaymentExpiredWebhookEventEventTypePaymentCanceled            PaymentExpiredWebhookEventEventType = "payment.canceled"
	PaymentExpiredWebhookEventEventTypePaymentExpired             PaymentExpiredWebhookEventEventType = "payment.expired"
	PaymentExpiredWebhookEventEventTypeMessageSent                PaymentExpiredWebhookEventEventType = "message.sent"
	PaymentExpiredWebhookEventEventTypeMessageReceived            PaymentExpiredWebhookEventEventType = "message.received"
	PaymentExpiredWebhookEventEventTypeMessageRead                PaymentExpiredWebhookEventEventType = "message.read"
	PaymentExpiredWebhookEventEventTypeMessageDelivered           PaymentExpiredWebhookEventEventType = "message.delivered"
	PaymentExpiredWebhookEventEventTypeMessageFailed              PaymentExpiredWebhookEventEventType = "message.failed"
	PaymentExpiredWebhookEventEventTypeMessageEdited              PaymentExpiredWebhookEventEventType = "message.edited"
	PaymentExpiredWebhookEventEventTypeReactionAdded              PaymentExpiredWebhookEventEventType = "reaction.added"
	PaymentExpiredWebhookEventEventTypeReactionRemoved            PaymentExpiredWebhookEventEventType = "reaction.removed"
	PaymentExpiredWebhookEventEventTypePollReceived               PaymentExpiredWebhookEventEventType = "poll.received"
	PaymentExpiredWebhookEventEventTypePollFailed                 PaymentExpiredWebhookEventEventType = "poll.failed"
	PaymentExpiredWebhookEventEventTypePollSent                   PaymentExpiredWebhookEventEventType = "poll.sent"
	PaymentExpiredWebhookEventEventTypePollDelivered              PaymentExpiredWebhookEventEventType = "poll.delivered"
	PaymentExpiredWebhookEventEventTypePollRead                   PaymentExpiredWebhookEventEventType = "poll.read"
	PaymentExpiredWebhookEventEventTypePollUpdated                PaymentExpiredWebhookEventEventType = "poll.updated"
	PaymentExpiredWebhookEventEventTypePollVoteAdded              PaymentExpiredWebhookEventEventType = "poll.vote.added"
	PaymentExpiredWebhookEventEventTypePollVoteRemoved            PaymentExpiredWebhookEventEventType = "poll.vote.removed"
	PaymentExpiredWebhookEventEventTypePollReactionAdded          PaymentExpiredWebhookEventEventType = "poll.reaction.added"
	PaymentExpiredWebhookEventEventTypeParticipantAdded           PaymentExpiredWebhookEventEventType = "participant.added"
	PaymentExpiredWebhookEventEventTypeParticipantRemoved         PaymentExpiredWebhookEventEventType = "participant.removed"
	PaymentExpiredWebhookEventEventTypeChatCreated                PaymentExpiredWebhookEventEventType = "chat.created"
	PaymentExpiredWebhookEventEventTypeChatGroupNameUpdated       PaymentExpiredWebhookEventEventType = "chat.group_name_updated"
	PaymentExpiredWebhookEventEventTypeChatGroupIconUpdated       PaymentExpiredWebhookEventEventType = "chat.group_icon_updated"
	PaymentExpiredWebhookEventEventTypeChatGroupNameUpdateFailed  PaymentExpiredWebhookEventEventType = "chat.group_name_update_failed"
	PaymentExpiredWebhookEventEventTypeChatGroupIconUpdateFailed  PaymentExpiredWebhookEventEventType = "chat.group_icon_update_failed"
	PaymentExpiredWebhookEventEventTypeChatBackgroundUpdated      PaymentExpiredWebhookEventEventType = "chat.background_updated"
	PaymentExpiredWebhookEventEventTypeChatBackgroundUpdateFailed PaymentExpiredWebhookEventEventType = "chat.background_update_failed"
	PaymentExpiredWebhookEventEventTypeChatTypingIndicatorStarted PaymentExpiredWebhookEventEventType = "chat.typing_indicator.started"
	PaymentExpiredWebhookEventEventTypeChatTypingIndicatorStopped PaymentExpiredWebhookEventEventType = "chat.typing_indicator.stopped"
	PaymentExpiredWebhookEventEventTypePhoneNumberStatusUpdated   PaymentExpiredWebhookEventEventType = "phone_number.status_updated"
	PaymentExpiredWebhookEventEventTypeContactCardReceived        PaymentExpiredWebhookEventEventType = "contact_card.received"
	PaymentExpiredWebhookEventEventTypeCallInitiated              PaymentExpiredWebhookEventEventType = "call.initiated"
	PaymentExpiredWebhookEventEventTypeCallRinging                PaymentExpiredWebhookEventEventType = "call.ringing"
	PaymentExpiredWebhookEventEventTypeCallAnswered               PaymentExpiredWebhookEventEventType = "call.answered"
	PaymentExpiredWebhookEventEventTypeCallEnded                  PaymentExpiredWebhookEventEventType = "call.ended"
	PaymentExpiredWebhookEventEventTypeCallFailed                 PaymentExpiredWebhookEventEventType = "call.failed"
	PaymentExpiredWebhookEventEventTypeCallDeclined               PaymentExpiredWebhookEventEventType = "call.declined"
	PaymentExpiredWebhookEventEventTypeCallNoAnswer               PaymentExpiredWebhookEventEventType = "call.no_answer"
	PaymentExpiredWebhookEventEventTypeLocationSharingStarted     PaymentExpiredWebhookEventEventType = "location.sharing.started"
	PaymentExpiredWebhookEventEventTypeLocationSharingStopped     PaymentExpiredWebhookEventEventType = "location.sharing.stopped"
	PaymentExpiredWebhookEventEventTypePaymentDeclined            PaymentExpiredWebhookEventEventType = "payment.declined"
	PaymentExpiredWebhookEventEventTypePaymentAuthorized          PaymentExpiredWebhookEventEventType = "payment.authorized"
	PaymentExpiredWebhookEventEventTypeConnectionCreated          PaymentExpiredWebhookEventEventType = "connection.created"
	PaymentExpiredWebhookEventEventTypeConnectionRevoked          PaymentExpiredWebhookEventEventType = "connection.revoked"
)

type PaymentSucceededWebhookEvent struct {
	// API version for the webhook payload format
	APIVersion string `json:"api_version" api:"required"`
	// When the event was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The payment request, as returned by
	// `GET /v3/payment_requests/{paymentRequestId}`.
	Data PaymentSucceededWebhookEventData `json:"data" api:"required"`
	// Unique identifier for this event (for deduplication)
	EventID string `json:"event_id" api:"required" format:"uuid"`
	// Any of "payment.succeeded", "payment.canceled", "payment.expired",
	// "message.sent", "message.received", "message.read", "message.delivered",
	// "message.failed", "message.edited", "reaction.added", "reaction.removed",
	// "poll.received", "poll.failed", "poll.sent", "poll.delivered", "poll.read",
	// "poll.updated", "poll.vote.added", "poll.vote.removed", "poll.reaction.added",
	// "participant.added", "participant.removed", "chat.created",
	// "chat.group_name_updated", "chat.group_icon_updated",
	// "chat.group_name_update_failed", "chat.group_icon_update_failed",
	// "chat.background_updated", "chat.background_update_failed",
	// "chat.typing_indicator.started", "chat.typing_indicator.stopped",
	// "phone_number.status_updated", "contact_card.received", "call.initiated",
	// "call.ringing", "call.answered", "call.ended", "call.failed", "call.declined",
	// "call.no_answer", "location.sharing.started", "location.sharing.stopped",
	// "payment.declined", "payment.authorized", "connection.created",
	// "connection.revoked".
	EventType PaymentSucceededWebhookEventEventType `json:"event_type" api:"required"`
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
func (r PaymentSucceededWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PaymentSucceededWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payment request, as returned by
// `GET /v3/payment_requests/{paymentRequestId}`.
type PaymentSucceededWebhookEventData struct {
	// The payment request id.
	ID string `json:"id" api:"required" format:"uuid"`
	// What was charged at checkout, in the currency's minor units. In `subscription`
	// mode this is the first invoice's total — all items after any discounts are
	// applied.
	Amount int64 `json:"amount" api:"required"`
	// URL the recipient opens to pay
	// (`https://zero.linqapp.com/pay/{slug}?session=...`).
	CheckoutURL string    `json:"checkout_url" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency    string    `json:"currency" api:"required"`
	Object      string    `json:"object" api:"required"`
	// Any of "succeeded", "failed", "canceled", "expired".
	Status      string `json:"status" api:"required"`
	Description string `json:"description"`
	// Subscription mode — the discount Stripe applied, read back from the coupon.
	// Absent when none was applied.
	Discount PaymentSucceededWebhookEventDataDiscount `json:"discount"`
	// Subscription mode — how often the subscription renews.
	//
	// Any of "day", "week", "month", "year".
	Interval string `json:"interval"`
	// Subscription mode — intervals per renewal.
	IntervalCount int64             `json:"interval_count"`
	Metadata      map[string]string `json:"metadata"`
	// Whether the request collected a one-time charge or started a subscription.
	//
	// Any of "payment", "subscription".
	Mode string `json:"mode"`
	// Natural-rail join keys, present when `rail: natural`.
	Natural PaymentSucceededWebhookEventDataNatural `json:"natural"`
	// Subscription mode — the recurring price subscribed to.
	PriceID string `json:"price_id"`
	// Subscription mode — units of the price subscribed to.
	Quantity int64 `json:"quantity"`
	// The rail this request settled on.
	//
	// Any of "stripe", "natural".
	Rail string `json:"rail"`
	// Ids of the Stripe objects on your connected account — join keys into your own
	// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
	// `subscription_id`.
	Stripe PaymentSucceededWebhookEventDataStripe `json:"stripe"`
	// Subscription mode — when the free trial ends and the first charge happens. On a
	// trial request, `payment.succeeded` means the payment method was collected ($0
	// moved).
	TrialEnd  time.Time `json:"trial_end" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		CheckoutURL   respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Object        respjson.Field
		Status        respjson.Field
		Description   respjson.Field
		Discount      respjson.Field
		Interval      respjson.Field
		IntervalCount respjson.Field
		Metadata      respjson.Field
		Mode          respjson.Field
		Natural       respjson.Field
		PriceID       respjson.Field
		Quantity      respjson.Field
		Rail          respjson.Field
		Stripe        respjson.Field
		TrialEnd      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentSucceededWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PaymentSucceededWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription mode — the discount Stripe applied, read back from the coupon.
// Absent when none was applied.
type PaymentSucceededWebhookEventDataDiscount struct {
	Coupon string `json:"coupon"`
	// Name of the coupon/promo code displayed to customers.
	Label         string `json:"label"`
	PromotionCode string `json:"promotion_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coupon        respjson.Field
		Label         respjson.Field
		PromotionCode respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentSucceededWebhookEventDataDiscount) RawJSON() string { return r.JSON.raw }
func (r *PaymentSucceededWebhookEventDataDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Natural-rail join keys, present when `rail: natural`.
type PaymentSucceededWebhookEventDataNatural struct {
	// The Natural payment request (`prq_...`).
	PaymentRequestID string `json:"payment_request_id"`
	// The settled transaction (`txn_...`).
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentRequestID respjson.Field
		TransactionID    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentSucceededWebhookEventDataNatural) RawJSON() string { return r.JSON.raw }
func (r *PaymentSucceededWebhookEventDataNatural) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ids of the Stripe objects on your connected account — join keys into your own
// Stripe Dashboard/API. Manage a subscription's post-checkout lifecycle with
// `subscription_id`.
type PaymentSucceededWebhookEventDataStripe struct {
	// The Customer the request is attached to (`cus_...`). Always set in subscription
	// mode; set in payment mode only when the request was created with a
	// `customer_id`.
	CustomerID string `json:"customer_id"`
	// The PaymentIntent collected at checkout (`pi_...`).
	PaymentIntentID string `json:"payment_intent_id"`
	// Subscription mode — the Subscription (`sub_...`).
	SubscriptionID string `json:"subscription_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID      respjson.Field
		PaymentIntentID respjson.Field
		SubscriptionID  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentSucceededWebhookEventDataStripe) RawJSON() string { return r.JSON.raw }
func (r *PaymentSucceededWebhookEventDataStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentSucceededWebhookEventEventType string

const (
	PaymentSucceededWebhookEventEventTypePaymentSucceeded           PaymentSucceededWebhookEventEventType = "payment.succeeded"
	PaymentSucceededWebhookEventEventTypePaymentCanceled            PaymentSucceededWebhookEventEventType = "payment.canceled"
	PaymentSucceededWebhookEventEventTypePaymentExpired             PaymentSucceededWebhookEventEventType = "payment.expired"
	PaymentSucceededWebhookEventEventTypeMessageSent                PaymentSucceededWebhookEventEventType = "message.sent"
	PaymentSucceededWebhookEventEventTypeMessageReceived            PaymentSucceededWebhookEventEventType = "message.received"
	PaymentSucceededWebhookEventEventTypeMessageRead                PaymentSucceededWebhookEventEventType = "message.read"
	PaymentSucceededWebhookEventEventTypeMessageDelivered           PaymentSucceededWebhookEventEventType = "message.delivered"
	PaymentSucceededWebhookEventEventTypeMessageFailed              PaymentSucceededWebhookEventEventType = "message.failed"
	PaymentSucceededWebhookEventEventTypeMessageEdited              PaymentSucceededWebhookEventEventType = "message.edited"
	PaymentSucceededWebhookEventEventTypeReactionAdded              PaymentSucceededWebhookEventEventType = "reaction.added"
	PaymentSucceededWebhookEventEventTypeReactionRemoved            PaymentSucceededWebhookEventEventType = "reaction.removed"
	PaymentSucceededWebhookEventEventTypePollReceived               PaymentSucceededWebhookEventEventType = "poll.received"
	PaymentSucceededWebhookEventEventTypePollFailed                 PaymentSucceededWebhookEventEventType = "poll.failed"
	PaymentSucceededWebhookEventEventTypePollSent                   PaymentSucceededWebhookEventEventType = "poll.sent"
	PaymentSucceededWebhookEventEventTypePollDelivered              PaymentSucceededWebhookEventEventType = "poll.delivered"
	PaymentSucceededWebhookEventEventTypePollRead                   PaymentSucceededWebhookEventEventType = "poll.read"
	PaymentSucceededWebhookEventEventTypePollUpdated                PaymentSucceededWebhookEventEventType = "poll.updated"
	PaymentSucceededWebhookEventEventTypePollVoteAdded              PaymentSucceededWebhookEventEventType = "poll.vote.added"
	PaymentSucceededWebhookEventEventTypePollVoteRemoved            PaymentSucceededWebhookEventEventType = "poll.vote.removed"
	PaymentSucceededWebhookEventEventTypePollReactionAdded          PaymentSucceededWebhookEventEventType = "poll.reaction.added"
	PaymentSucceededWebhookEventEventTypeParticipantAdded           PaymentSucceededWebhookEventEventType = "participant.added"
	PaymentSucceededWebhookEventEventTypeParticipantRemoved         PaymentSucceededWebhookEventEventType = "participant.removed"
	PaymentSucceededWebhookEventEventTypeChatCreated                PaymentSucceededWebhookEventEventType = "chat.created"
	PaymentSucceededWebhookEventEventTypeChatGroupNameUpdated       PaymentSucceededWebhookEventEventType = "chat.group_name_updated"
	PaymentSucceededWebhookEventEventTypeChatGroupIconUpdated       PaymentSucceededWebhookEventEventType = "chat.group_icon_updated"
	PaymentSucceededWebhookEventEventTypeChatGroupNameUpdateFailed  PaymentSucceededWebhookEventEventType = "chat.group_name_update_failed"
	PaymentSucceededWebhookEventEventTypeChatGroupIconUpdateFailed  PaymentSucceededWebhookEventEventType = "chat.group_icon_update_failed"
	PaymentSucceededWebhookEventEventTypeChatBackgroundUpdated      PaymentSucceededWebhookEventEventType = "chat.background_updated"
	PaymentSucceededWebhookEventEventTypeChatBackgroundUpdateFailed PaymentSucceededWebhookEventEventType = "chat.background_update_failed"
	PaymentSucceededWebhookEventEventTypeChatTypingIndicatorStarted PaymentSucceededWebhookEventEventType = "chat.typing_indicator.started"
	PaymentSucceededWebhookEventEventTypeChatTypingIndicatorStopped PaymentSucceededWebhookEventEventType = "chat.typing_indicator.stopped"
	PaymentSucceededWebhookEventEventTypePhoneNumberStatusUpdated   PaymentSucceededWebhookEventEventType = "phone_number.status_updated"
	PaymentSucceededWebhookEventEventTypeContactCardReceived        PaymentSucceededWebhookEventEventType = "contact_card.received"
	PaymentSucceededWebhookEventEventTypeCallInitiated              PaymentSucceededWebhookEventEventType = "call.initiated"
	PaymentSucceededWebhookEventEventTypeCallRinging                PaymentSucceededWebhookEventEventType = "call.ringing"
	PaymentSucceededWebhookEventEventTypeCallAnswered               PaymentSucceededWebhookEventEventType = "call.answered"
	PaymentSucceededWebhookEventEventTypeCallEnded                  PaymentSucceededWebhookEventEventType = "call.ended"
	PaymentSucceededWebhookEventEventTypeCallFailed                 PaymentSucceededWebhookEventEventType = "call.failed"
	PaymentSucceededWebhookEventEventTypeCallDeclined               PaymentSucceededWebhookEventEventType = "call.declined"
	PaymentSucceededWebhookEventEventTypeCallNoAnswer               PaymentSucceededWebhookEventEventType = "call.no_answer"
	PaymentSucceededWebhookEventEventTypeLocationSharingStarted     PaymentSucceededWebhookEventEventType = "location.sharing.started"
	PaymentSucceededWebhookEventEventTypeLocationSharingStopped     PaymentSucceededWebhookEventEventType = "location.sharing.stopped"
	PaymentSucceededWebhookEventEventTypePaymentDeclined            PaymentSucceededWebhookEventEventType = "payment.declined"
	PaymentSucceededWebhookEventEventTypePaymentAuthorized          PaymentSucceededWebhookEventEventType = "payment.authorized"
	PaymentSucceededWebhookEventEventTypeConnectionCreated          PaymentSucceededWebhookEventEventType = "connection.created"
	PaymentSucceededWebhookEventEventTypeConnectionRevoked          PaymentSucceededWebhookEventEventType = "connection.revoked"
)

// UnwrapWebhookEventUnion contains all possible properties and values from
// [MessageSentWebhookEvent], [MessageReceivedWebhookEvent],
// [MessageReadWebhookEvent], [MessageDeliveredWebhookEvent],
// [MessageFailedWebhookEvent], [MessageEditedWebhookEvent],
// [ReactionAddedWebhookEvent], [ReactionRemovedWebhookEvent],
// [PollReceivedWebhookEvent], [PollSentWebhookEvent], [PollDeliveredWebhookEvent],
// [PollReadWebhookEvent], [PollUpdatedWebhookEvent], [PollFailedWebhookEvent],
// [PollVoteAddedWebhookEvent], [PollVoteRemovedWebhookEvent],
// [PollReactionAddedWebhookEvent], [ParticipantAddedWebhookEvent],
// [ParticipantRemovedWebhookEvent], [ChatCreatedWebhookEvent],
// [ChatGroupNameUpdatedWebhookEvent], [ChatGroupIconUpdatedWebhookEvent],
// [ChatGroupNameUpdateFailedWebhookEvent],
// [ChatGroupIconUpdateFailedWebhookEvent],
// [ChatTypingIndicatorStartedWebhookEvent],
// [ChatTypingIndicatorStoppedWebhookEvent], [ChatBackgroundUpdatedWebhookEvent],
// [ChatBackgroundUpdateFailedWebhookEvent], [ContactCardReceivedWebhookEvent],
// [PhoneNumberStatusUpdatedWebhookEvent], [ConnectionCreatedWebhookEvent],
// [ConnectionRevokedWebhookEvent], [LocationSharingStartedWebhookEvent],
// [LocationSharingStoppedWebhookEvent], [PaymentAuthorizedWebhookEvent],
// [PaymentCanceledWebhookEvent], [PaymentDeclinedWebhookEvent],
// [PaymentExpiredWebhookEvent], [PaymentSucceededWebhookEvent].
//
// Use the [UnwrapWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	APIVersion string    `json:"api_version"`
	CreatedAt  time.Time `json:"created_at"`
	// This field is a union of [MessageEventV2], [MessageFailedWebhookEventData],
	// [MessageEditedWebhookEventData], [ReactionEventBase],
	// [PollReceivedWebhookEventData], [PollSentWebhookEventData],
	// [PollDeliveredWebhookEventData], [PollReadWebhookEventData],
	// [PollUpdatedWebhookEventData], [PollFailedWebhookEventData],
	// [PollVoteAddedWebhookEventData], [PollVoteRemovedWebhookEventData],
	// [ParticipantAddedWebhookEventData], [ParticipantRemovedWebhookEventData],
	// [ChatCreatedWebhookEventData], [ChatGroupNameUpdatedWebhookEventData],
	// [ChatGroupIconUpdatedWebhookEventData],
	// [ChatGroupNameUpdateFailedWebhookEventData],
	// [ChatGroupIconUpdateFailedWebhookEventData],
	// [ChatTypingIndicatorStartedWebhookEventData],
	// [ChatTypingIndicatorStoppedWebhookEventData],
	// [ChatBackgroundUpdatedWebhookEventData],
	// [ChatBackgroundUpdateFailedWebhookEventData],
	// [ContactCardReceivedWebhookEventData],
	// [PhoneNumberStatusUpdatedWebhookEventData], [ConnectionCreatedWebhookEventData],
	// [ConnectionRevokedWebhookEventData], [LocationSharingStartedWebhookEventData],
	// [LocationSharingStoppedWebhookEventData], [PaymentAuthorizedWebhookEventData],
	// [PaymentCanceledWebhookEventData], [PaymentDeclinedWebhookEventData],
	// [PaymentExpiredWebhookEventData], [PaymentSucceededWebhookEventData]
	Data    UnwrapWebhookEventUnionData `json:"data"`
	EventID string                      `json:"event_id"`
	// Any of nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
	// nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
	// nil, nil, nil, nil, nil, nil, nil, nil, nil.
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

func (u UnwrapWebhookEventUnion) AsMessageSentWebhookEvent() (v MessageSentWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsMessageReceivedWebhookEvent() (v MessageReceivedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsMessageReadWebhookEvent() (v MessageReadWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsMessageDeliveredWebhookEvent() (v MessageDeliveredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsMessageFailedWebhookEvent() (v MessageFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsMessageEditedWebhookEvent() (v MessageEditedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsReactionAddedWebhookEvent() (v ReactionAddedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsReactionRemovedWebhookEvent() (v ReactionRemovedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPollReceivedWebhookEvent() (v PollReceivedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPollSentWebhookEvent() (v PollSentWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPollDeliveredWebhookEvent() (v PollDeliveredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPollReadWebhookEvent() (v PollReadWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPollUpdatedWebhookEvent() (v PollUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPollFailedWebhookEvent() (v PollFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPollVoteAddedWebhookEvent() (v PollVoteAddedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPollVoteRemovedWebhookEvent() (v PollVoteRemovedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPollReactionAddedWebhookEvent() (v PollReactionAddedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsParticipantAddedWebhookEvent() (v ParticipantAddedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsParticipantRemovedWebhookEvent() (v ParticipantRemovedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsChatCreatedWebhookEvent() (v ChatCreatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsChatGroupNameUpdatedWebhookEvent() (v ChatGroupNameUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsChatGroupIconUpdatedWebhookEvent() (v ChatGroupIconUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsChatGroupNameUpdateFailedWebhookEvent() (v ChatGroupNameUpdateFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsChatGroupIconUpdateFailedWebhookEvent() (v ChatGroupIconUpdateFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsChatTypingIndicatorStartedWebhookEvent() (v ChatTypingIndicatorStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsChatTypingIndicatorStoppedWebhookEvent() (v ChatTypingIndicatorStoppedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsChatBackgroundUpdatedWebhookEvent() (v ChatBackgroundUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsChatBackgroundUpdateFailedWebhookEvent() (v ChatBackgroundUpdateFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsContactCardReceivedWebhookEvent() (v ContactCardReceivedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPhoneNumberStatusUpdatedWebhookEvent() (v PhoneNumberStatusUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConnectionCreatedWebhookEvent() (v ConnectionCreatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConnectionRevokedWebhookEvent() (v ConnectionRevokedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsLocationSharingStartedWebhookEvent() (v LocationSharingStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsLocationSharingStoppedWebhookEvent() (v LocationSharingStoppedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPaymentAuthorizedWebhookEvent() (v PaymentAuthorizedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPaymentCanceledWebhookEvent() (v PaymentCanceledWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPaymentDeclinedWebhookEvent() (v PaymentDeclinedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPaymentExpiredWebhookEvent() (v PaymentExpiredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPaymentSucceededWebhookEvent() (v PaymentSucceededWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionData is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionData struct {
	ID string `json:"id"`
	// This field is a union of [MessageEventV2Chat],
	// [MessageEditedWebhookEventDataChat], [PollReceivedWebhookEventDataChat],
	// [PollSentWebhookEventDataChat], [PollDeliveredWebhookEventDataChat],
	// [PollReadWebhookEventDataChat], [PollUpdatedWebhookEventDataChat],
	// [PollFailedWebhookEventDataChat], [PollVoteAddedWebhookEventDataChat],
	// [PollVoteRemovedWebhookEventDataChat],
	// [ChatBackgroundUpdatedWebhookEventDataChat]
	Chat      UnwrapWebhookEventUnionDataChat `json:"chat"`
	Direction string                          `json:"direction"`
	// This field is from variant [MessageEventV2].
	Parts []MessageEventV2PartUnion `json:"parts"`
	// This field is a union of [shared.ChatHandle], [string]
	SenderHandle UnwrapWebhookEventUnionDataSenderHandle `json:"sender_handle"`
	Service      string                                  `json:"service"`
	DeliveredAt  time.Time                               `json:"delivered_at"`
	// This field is from variant [MessageEventV2].
	Effect SchemasMessageEffect `json:"effect"`
	// This field is from variant [MessageEventV2].
	IdempotencyKey   string    `json:"idempotency_key"`
	PreferredService string    `json:"preferred_service"`
	ReadAt           time.Time `json:"read_at"`
	// This field is from variant [MessageEventV2].
	ReconciledAt time.Time `json:"reconciled_at"`
	// This field is from variant [MessageEventV2].
	ReplyTo MessageEventV2ReplyTo `json:"reply_to"`
	SentAt  time.Time             `json:"sent_at"`
	// This field is from variant [MessageFailedWebhookEventData].
	Code     int64     `json:"code"`
	FailedAt time.Time `json:"failed_at"`
	ChatID   string    `json:"chat_id"`
	// This field is from variant [MessageFailedWebhookEventData].
	DetailCode int64  `json:"detail_code"`
	MessageID  string `json:"message_id"`
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
	ReactionID string `json:"reaction_id"`
	// This field is from variant [ReactionEventBase].
	Sticker   ReactionEventBaseSticker `json:"sticker"`
	CreatedAt time.Time                `json:"created_at"`
	// This field is a union of [PollReceivedWebhookEventDataPoll],
	// [PollSentWebhookEventDataPoll], [PollDeliveredWebhookEventDataPoll],
	// [PollReadWebhookEventDataPoll], [PollFailedWebhookEventDataPoll]
	Poll UnwrapWebhookEventUnionDataPoll `json:"poll"`
	// This field is from variant [PollReceivedWebhookEventData].
	ReceivedAt time.Time `json:"received_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	// This field is from variant [PollUpdatedWebhookEventData].
	AddedOptions []PollUpdatedWebhookEventDataAddedOption `json:"added_options"`
	// This field is from variant [PollFailedWebhookEventData].
	Error    PollFailedWebhookEventDataError `json:"error"`
	OptionID string                          `json:"option_id"`
	Handle   string                          `json:"handle"`
	// This field is from variant [ParticipantAddedWebhookEventData].
	AddedAt time.Time `json:"added_at"`
	// This field is from variant [ParticipantAddedWebhookEventData].
	Participant shared.ChatHandle `json:"participant"`
	// This field is from variant [ParticipantRemovedWebhookEventData].
	RemovedAt time.Time `json:"removed_at"`
	// This field is from variant [ChatCreatedWebhookEventData].
	DisplayName string `json:"display_name"`
	// This field is from variant [ChatCreatedWebhookEventData].
	Handles []shared.ChatHandle `json:"handles"`
	// This field is from variant [ChatCreatedWebhookEventData].
	HealthStatus ChatCreatedWebhookEventDataHealthStatus `json:"health_status"`
	// This field is from variant [ChatCreatedWebhookEventData].
	IsGroup bool `json:"is_group"`
	// This field is from variant [ChatGroupNameUpdatedWebhookEventData].
	ChangedByHandle shared.ChatHandle `json:"changed_by_handle"`
	NewValue        string            `json:"new_value"`
	OldValue        string            `json:"old_value"`
	ErrorCode       int64             `json:"error_code"`
	// This field is from variant [ChatBackgroundUpdatedWebhookEventData].
	ActorHandle shared.ChatHandle `json:"actor_handle"`
	// This field is from variant [ChatBackgroundUpdatedWebhookEventData].
	Background ChatBackgroundUpdatedWebhookEventDataBackground `json:"background"`
	// This field is from variant [ContactCardReceivedWebhookEventData].
	FirstName string `json:"first_name"`
	// This field is from variant [ContactCardReceivedWebhookEventData].
	LastName string `json:"last_name"`
	// This field is from variant [ContactCardReceivedWebhookEventData].
	OwnerHandle string `json:"owner_handle"`
	// This field is from variant [ContactCardReceivedWebhookEventData].
	MediaURL string `json:"media_url"`
	// This field is from variant [PhoneNumberStatusUpdatedWebhookEventData].
	ChangedAt time.Time `json:"changed_at"`
	// This field is from variant [PhoneNumberStatusUpdatedWebhookEventData].
	NewReputation string `json:"new_reputation"`
	// This field is from variant [PhoneNumberStatusUpdatedWebhookEventData].
	NewStatus string `json:"new_status"`
	// This field is from variant [PhoneNumberStatusUpdatedWebhookEventData].
	PhoneNumber string `json:"phone_number"`
	// This field is from variant [PhoneNumberStatusUpdatedWebhookEventData].
	PreviousReputation string `json:"previous_reputation"`
	// This field is from variant [PhoneNumberStatusUpdatedWebhookEventData].
	PreviousStatus string `json:"previous_status"`
	Amount         int64  `json:"amount"`
	CheckoutURL    string `json:"checkout_url"`
	Currency       string `json:"currency"`
	Object         string `json:"object"`
	Status         string `json:"status"`
	Description    string `json:"description"`
	// This field is a union of [ConnectionCreatedWebhookEventDataDiscount],
	// [ConnectionRevokedWebhookEventDataDiscount],
	// [PaymentAuthorizedWebhookEventDataDiscount],
	// [PaymentCanceledWebhookEventDataDiscount],
	// [PaymentDeclinedWebhookEventDataDiscount],
	// [PaymentExpiredWebhookEventDataDiscount],
	// [PaymentSucceededWebhookEventDataDiscount]
	Discount      UnwrapWebhookEventUnionDataDiscount `json:"discount"`
	Interval      string                              `json:"interval"`
	IntervalCount int64                               `json:"interval_count"`
	Metadata      string                              `json:"metadata"`
	Mode          string                              `json:"mode"`
	// This field is a union of [ConnectionCreatedWebhookEventDataNatural],
	// [ConnectionRevokedWebhookEventDataNatural],
	// [PaymentAuthorizedWebhookEventDataNatural],
	// [PaymentCanceledWebhookEventDataNatural],
	// [PaymentDeclinedWebhookEventDataNatural],
	// [PaymentExpiredWebhookEventDataNatural],
	// [PaymentSucceededWebhookEventDataNatural]
	Natural  UnwrapWebhookEventUnionDataNatural `json:"natural"`
	PriceID  string                             `json:"price_id"`
	Quantity int64                              `json:"quantity"`
	Rail     string                             `json:"rail"`
	// This field is a union of [ConnectionCreatedWebhookEventDataStripe],
	// [ConnectionRevokedWebhookEventDataStripe],
	// [PaymentAuthorizedWebhookEventDataStripe],
	// [PaymentCanceledWebhookEventDataStripe],
	// [PaymentDeclinedWebhookEventDataStripe], [PaymentExpiredWebhookEventDataStripe],
	// [PaymentSucceededWebhookEventDataStripe]
	Stripe   UnwrapWebhookEventUnionDataStripe `json:"stripe"`
	TrialEnd time.Time                         `json:"trial_end"`
	BeganAt  time.Time                         `json:"began_at"`
	// This field is from variant [LocationSharingStartedWebhookEventData].
	EndsAt     time.Time `json:"ends_at"`
	SharedBy   string    `json:"shared_by"`
	SharedWith string    `json:"shared_with"`
	// This field is from variant [LocationSharingStoppedWebhookEventData].
	EndedAt time.Time `json:"ended_at"`
	JSON    struct {
		ID                 respjson.Field
		Chat               respjson.Field
		Direction          respjson.Field
		Parts              respjson.Field
		SenderHandle       respjson.Field
		Service            respjson.Field
		DeliveredAt        respjson.Field
		Effect             respjson.Field
		IdempotencyKey     respjson.Field
		PreferredService   respjson.Field
		ReadAt             respjson.Field
		ReconciledAt       respjson.Field
		ReplyTo            respjson.Field
		SentAt             respjson.Field
		Code               respjson.Field
		FailedAt           respjson.Field
		ChatID             respjson.Field
		DetailCode         respjson.Field
		MessageID          respjson.Field
		Reason             respjson.Field
		EditedAt           respjson.Field
		Part               respjson.Field
		IsFromMe           respjson.Field
		ReactionType       respjson.Field
		CustomEmoji        respjson.Field
		From               respjson.Field
		FromHandle         respjson.Field
		PartIndex          respjson.Field
		ReactedAt          respjson.Field
		ReactionID         respjson.Field
		Sticker            respjson.Field
		CreatedAt          respjson.Field
		Poll               respjson.Field
		ReceivedAt         respjson.Field
		UpdatedAt          respjson.Field
		AddedOptions       respjson.Field
		Error              respjson.Field
		OptionID           respjson.Field
		Handle             respjson.Field
		AddedAt            respjson.Field
		Participant        respjson.Field
		RemovedAt          respjson.Field
		DisplayName        respjson.Field
		Handles            respjson.Field
		HealthStatus       respjson.Field
		IsGroup            respjson.Field
		ChangedByHandle    respjson.Field
		NewValue           respjson.Field
		OldValue           respjson.Field
		ErrorCode          respjson.Field
		ActorHandle        respjson.Field
		Background         respjson.Field
		FirstName          respjson.Field
		LastName           respjson.Field
		OwnerHandle        respjson.Field
		MediaURL           respjson.Field
		ChangedAt          respjson.Field
		NewReputation      respjson.Field
		NewStatus          respjson.Field
		PhoneNumber        respjson.Field
		PreviousReputation respjson.Field
		PreviousStatus     respjson.Field
		Amount             respjson.Field
		CheckoutURL        respjson.Field
		Currency           respjson.Field
		Object             respjson.Field
		Status             respjson.Field
		Description        respjson.Field
		Discount           respjson.Field
		Interval           respjson.Field
		IntervalCount      respjson.Field
		Metadata           respjson.Field
		Mode               respjson.Field
		Natural            respjson.Field
		PriceID            respjson.Field
		Quantity           respjson.Field
		Rail               respjson.Field
		Stripe             respjson.Field
		TrialEnd           respjson.Field
		BeganAt            respjson.Field
		EndsAt             respjson.Field
		SharedBy           respjson.Field
		SharedWith         respjson.Field
		EndedAt            respjson.Field
		raw                string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataChat is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataChat provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataChat struct {
	ID string `json:"id"`
	// This field is a union of [MessageEventV2ChatHealthStatus],
	// [MessageEditedWebhookEventDataChatHealthStatus]
	HealthStatus UnwrapWebhookEventUnionDataChatHealthStatus `json:"health_status"`
	IsGroup      bool                                        `json:"is_group"`
	// This field is from variant [MessageEventV2Chat].
	OwnerHandle shared.ChatHandle `json:"owner_handle"`
	JSON        struct {
		ID           respjson.Field
		HealthStatus respjson.Field
		IsGroup      respjson.Field
		OwnerHandle  respjson.Field
		raw          string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataChatHealthStatus is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataChatHealthStatus provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataChatHealthStatus struct {
	DocURL    string    `json:"doc_url"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	JSON      struct {
		DocURL    respjson.Field
		Status    respjson.Field
		UpdatedAt respjson.Field
		raw       string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataChatHealthStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataSenderHandle is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataSenderHandle provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type UnwrapWebhookEventUnionDataSenderHandle struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [shared.ChatHandle].
	ID string `json:"id"`
	// This field is from variant [shared.ChatHandle].
	Handle string `json:"handle"`
	// This field is from variant [shared.ChatHandle].
	JoinedAt time.Time `json:"joined_at"`
	// This field is from variant [shared.ChatHandle].
	Service shared.ServiceType `json:"service"`
	// This field is from variant [shared.ChatHandle].
	IsMe bool `json:"is_me"`
	// This field is from variant [shared.ChatHandle].
	LeftAt time.Time `json:"left_at"`
	// This field is from variant [shared.ChatHandle].
	Status shared.ChatHandleStatus `json:"status"`
	JSON   struct {
		OfString respjson.Field
		ID       respjson.Field
		Handle   respjson.Field
		JoinedAt respjson.Field
		Service  respjson.Field
		IsMe     respjson.Field
		LeftAt   respjson.Field
		Status   respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataSenderHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPoll is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPoll provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataPoll struct {
	// This field is a union of [[]PollReceivedWebhookEventDataPollOption],
	// [[]PollSentWebhookEventDataPollOption],
	// [[]PollDeliveredWebhookEventDataPollOption],
	// [[]PollReadWebhookEventDataPollOption], [[]PollFailedWebhookEventDataPollOption]
	Options     UnwrapWebhookEventUnionDataPollOptions `json:"options"`
	TotalVoters int64                                  `json:"total_voters"`
	JSON        struct {
		Options     respjson.Field
		TotalVoters respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPoll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPollOptions is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPollOptions provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPollReceivedWebhookEventDataPollOptions
// OfPollSentWebhookEventDataPollOptions OfPollDeliveredWebhookEventDataPollOptions
// OfPollReadWebhookEventDataPollOptions OfPollFailedWebhookEventDataPollOptions]
type UnwrapWebhookEventUnionDataPollOptions struct {
	// This field will be present if the value is a
	// [[]PollReceivedWebhookEventDataPollOption] instead of an object.
	OfPollReceivedWebhookEventDataPollOptions []PollReceivedWebhookEventDataPollOption `json:",inline"`
	// This field will be present if the value is a
	// [[]PollSentWebhookEventDataPollOption] instead of an object.
	OfPollSentWebhookEventDataPollOptions []PollSentWebhookEventDataPollOption `json:",inline"`
	// This field will be present if the value is a
	// [[]PollDeliveredWebhookEventDataPollOption] instead of an object.
	OfPollDeliveredWebhookEventDataPollOptions []PollDeliveredWebhookEventDataPollOption `json:",inline"`
	// This field will be present if the value is a
	// [[]PollReadWebhookEventDataPollOption] instead of an object.
	OfPollReadWebhookEventDataPollOptions []PollReadWebhookEventDataPollOption `json:",inline"`
	// This field will be present if the value is a
	// [[]PollFailedWebhookEventDataPollOption] instead of an object.
	OfPollFailedWebhookEventDataPollOptions []PollFailedWebhookEventDataPollOption `json:",inline"`
	JSON                                    struct {
		OfPollReceivedWebhookEventDataPollOptions  respjson.Field
		OfPollSentWebhookEventDataPollOptions      respjson.Field
		OfPollDeliveredWebhookEventDataPollOptions respjson.Field
		OfPollReadWebhookEventDataPollOptions      respjson.Field
		OfPollFailedWebhookEventDataPollOptions    respjson.Field
		raw                                        string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPollOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataDiscount is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataDiscount provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataDiscount struct {
	Coupon        string `json:"coupon"`
	Label         string `json:"label"`
	PromotionCode string `json:"promotion_code"`
	JSON          struct {
		Coupon        respjson.Field
		Label         respjson.Field
		PromotionCode respjson.Field
		raw           string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataNatural is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataNatural provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataNatural struct {
	PaymentRequestID string `json:"payment_request_id"`
	TransactionID    string `json:"transaction_id"`
	JSON             struct {
		PaymentRequestID respjson.Field
		TransactionID    respjson.Field
		raw              string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataNatural) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataStripe is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataStripe provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataStripe struct {
	CustomerID      string `json:"customer_id"`
	PaymentIntentID string `json:"payment_intent_id"`
	SubscriptionID  string `json:"subscription_id"`
	JSON            struct {
		CustomerID      respjson.Field
		PaymentIntentID respjson.Field
		SubscriptionID  respjson.Field
		raw             string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
