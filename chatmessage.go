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

// Messages are individual text or multimedia communications within a chat thread.
//
// Messages can include text, attachments, special effects (like confetti or
// fireworks), and reactions. All messages are associated with a specific chat and
// sent from a phone number you own.
//
// Messages support delivery status tracking, read receipts, and editing
// capabilities.
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
	// Current delivery status of a message
	//
	// Any of "pending", "queued", "sent", "delivered", "failed".
	DeliveryStatus SentMessageDeliveryStatus `json:"delivery_status" api:"required"`
	// Whether the message has been read
	IsRead bool `json:"is_read" api:"required"`
	// Message parts in order (text and media)
	Parts []SentMessagePartUnion `json:"parts" api:"required"`
	// When the message was sent
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
	SentMessageDeliveryStatusFailed    SentMessageDeliveryStatus = "failed"
)

// SentMessagePartUnion contains all possible properties and values from
// [shared.TextPartResponse], [shared.MediaPartResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SentMessagePartUnion struct {
	Reactions []shared.Reaction `json:"reactions"`
	Type      string            `json:"type"`
	// This field is from variant [shared.TextPartResponse].
	Value string `json:"value"`
	// This field is from variant [shared.TextPartResponse].
	TextDecorations []TextDecoration `json:"text_decorations"`
	// This field is from variant [shared.MediaPartResponse].
	ID string `json:"id"`
	// This field is from variant [shared.MediaPartResponse].
	Filename string `json:"filename"`
	// This field is from variant [shared.MediaPartResponse].
	MimeType string `json:"mime_type"`
	// This field is from variant [shared.MediaPartResponse].
	SizeBytes int64 `json:"size_bytes"`
	// This field is from variant [shared.MediaPartResponse].
	URL  string `json:"url"`
	JSON struct {
		Reactions       respjson.Field
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

func (u SentMessagePartUnion) AsTextPartResponse() (v shared.TextPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SentMessagePartUnion) AsMediaPartResponse() (v shared.MediaPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SentMessagePartUnion) RawJSON() string { return u.JSON.raw }

func (r *SentMessagePartUnion) UnmarshalJSON(data []byte) error {
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
