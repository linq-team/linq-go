// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqapiv3

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/linq-api-v3-go/internal/apijson"
	"github.com/stainless-sdks/linq-api-v3-go/internal/apiquery"
	"github.com/stainless-sdks/linq-api-v3-go/internal/requestconfig"
	"github.com/stainless-sdks/linq-api-v3-go/option"
	"github.com/stainless-sdks/linq-api-v3-go/packages/param"
	"github.com/stainless-sdks/linq-api-v3-go/packages/respjson"
)

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
func (r *ChatMessageService) List(ctx context.Context, chatID string, query ChatMessageListParams, opts ...option.RequestOption) (res *ChatMessageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return
	}
	path := fmt.Sprintf("v3/chats/%s/messages", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
func (r *ChatMessageService) Send(ctx context.Context, chatID string, body ChatMessageSendParams, opts ...option.RequestOption) (res *ChatMessageSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return
	}
	path := fmt.Sprintf("v3/chats/%s/messages", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A message that was sent (used in CreateChat and SendMessage responses)
type SentMessage struct {
	// Message identifier (UUID)
	ID string `json:"id,required" format:"uuid"`
	// Current delivery status of a message
	//
	// Any of "pending", "queued", "sent", "delivered", "failed".
	DeliveryStatus SentMessageDeliveryStatus `json:"delivery_status,required"`
	// Whether the message has been read
	IsRead bool `json:"is_read,required"`
	// Message parts in order (text and media)
	Parts []SentMessagePartUnion `json:"parts,required"`
	// When the message was sent
	SentAt time.Time `json:"sent_at,required" format:"date-time"`
	// When the message was delivered
	DeliveredAt time.Time `json:"delivered_at,nullable" format:"date-time"`
	// iMessage effect applied to a message (screen or bubble effect)
	Effect MessageEffect `json:"effect,nullable"`
	// The sender of this message as a full handle object
	FromHandle ChatHandle `json:"from_handle,nullable"`
	// Preferred service for sending this message
	//
	// Any of "iMessage", "SMS", "RCS".
	PreferredService SentMessagePreferredService `json:"preferred_service,nullable"`
	// Indicates this message is a threaded reply to another message
	ReplyTo ReplyTo `json:"reply_to,nullable"`
	// Service used to send this message
	//
	// Any of "iMessage", "SMS", "RCS".
	Service SentMessageService `json:"service,nullable"`
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
// [TextPart], [MediaPart].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SentMessagePartUnion struct {
	Reactions []Reaction `json:"reactions"`
	Type      string     `json:"type"`
	// This field is from variant [TextPart].
	Value string `json:"value"`
	// This field is from variant [MediaPart].
	ID string `json:"id"`
	// This field is from variant [MediaPart].
	Filename string `json:"filename"`
	// This field is from variant [MediaPart].
	MimeType string `json:"mime_type"`
	// This field is from variant [MediaPart].
	SizeBytes int64 `json:"size_bytes"`
	// This field is from variant [MediaPart].
	URL  string `json:"url"`
	JSON struct {
		Reactions respjson.Field
		Type      respjson.Field
		Value     respjson.Field
		ID        respjson.Field
		Filename  respjson.Field
		MimeType  respjson.Field
		SizeBytes respjson.Field
		URL       respjson.Field
		raw       string
	} `json:"-"`
}

func (u SentMessagePartUnion) AsTextPart() (v TextPart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SentMessagePartUnion) AsMediaPart() (v MediaPart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SentMessagePartUnion) RawJSON() string { return u.JSON.raw }

func (r *SentMessagePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Preferred service for sending this message
type SentMessagePreferredService string

const (
	SentMessagePreferredServiceIMessage SentMessagePreferredService = "iMessage"
	SentMessagePreferredServiceSMS      SentMessagePreferredService = "SMS"
	SentMessagePreferredServiceRcs      SentMessagePreferredService = "RCS"
)

// Service used to send this message
type SentMessageService string

const (
	SentMessageServiceIMessage SentMessageService = "iMessage"
	SentMessageServiceSMS      SentMessageService = "SMS"
	SentMessageServiceRcs      SentMessageService = "RCS"
)

type ChatMessageListResponse struct {
	// List of messages
	Messages []Message `json:"messages,required"`
	// Cursor for fetching the next page of results. Null if there are no more results
	// to fetch. Pass this value as the `cursor` parameter in the next request.
	NextCursor string `json:"next_cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages    respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for sending a message to a chat
type ChatMessageSendResponse struct {
	// Unique identifier of the chat this message was sent to
	ChatID string `json:"chat_id,required" format:"uuid"`
	// A message that was sent (used in CreateChat and SendMessage responses)
	Message SentMessage `json:"message,required"`
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
	Message MessageContentParam `json:"message,omitzero,required"`
	paramObj
}

func (r ChatMessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
