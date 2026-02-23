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
	"github.com/stainless-sdks/linq-api-v3-go/shared/constant"
)

// ChatMessageService contains methods and other services that help with
// interacting with the linq API.
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

type ChatMessageListResponse struct {
	// List of messages
	Messages []ChatMessageListResponseMessage `json:"messages,required"`
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

type ChatMessageListResponseMessage struct {
	// Unique identifier for the message
	ID string `json:"id,required" format:"uuid"`
	// ID of the chat this message belongs to
	ChatID string `json:"chat_id,required" format:"uuid"`
	// When the message was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the message has been delivered
	IsDelivered bool `json:"is_delivered,required"`
	// Whether this message was sent by the authenticated user
	IsFromMe bool `json:"is_from_me,required"`
	// Whether the message has been read
	IsRead bool `json:"is_read,required"`
	// When the message was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// When the message was delivered
	DeliveredAt time.Time `json:"delivered_at,nullable" format:"date-time"`
	// iMessage effect applied to a message (screen or bubble effect)
	Effect ChatMessageListResponseMessageEffect `json:"effect,nullable"`
	// DEPRECATED: Use from_handle instead. Phone number of the message sender.
	//
	// Deprecated: deprecated
	From string `json:"from,nullable"`
	// The sender of this message as a full handle object
	FromHandle ChatMessageListResponseMessageFromHandle `json:"from_handle,nullable"`
	// Message parts in order (text and media)
	Parts []ChatMessageListResponseMessagePartUnion `json:"parts,nullable"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	PreferredService string `json:"preferred_service,nullable"`
	// When the message was read
	ReadAt time.Time `json:"read_at,nullable" format:"date-time"`
	// Indicates this message is a threaded reply to another message
	ReplyTo ChatMessageListResponseMessageReplyTo `json:"reply_to,nullable"`
	// When the message was sent
	SentAt time.Time `json:"sent_at,nullable" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ChatID           respjson.Field
		CreatedAt        respjson.Field
		IsDelivered      respjson.Field
		IsFromMe         respjson.Field
		IsRead           respjson.Field
		UpdatedAt        respjson.Field
		DeliveredAt      respjson.Field
		Effect           respjson.Field
		From             respjson.Field
		FromHandle       respjson.Field
		Parts            respjson.Field
		PreferredService respjson.Field
		ReadAt           respjson.Field
		ReplyTo          respjson.Field
		SentAt           respjson.Field
		Service          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageListResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// iMessage effect applied to a message (screen or bubble effect)
type ChatMessageListResponseMessageEffect struct {
	// Name of the effect. Common values:
	//
	//   - Screen effects: confetti, fireworks, lasers, sparkles, celebration, hearts,
	//     love, balloons, happy_birthday, echo, spotlight
	//   - Bubble effects: slam, loud, gentle, invisible
	Name string `json:"name"`
	// Type of effect
	//
	// Any of "screen", "bubble".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponseMessageEffect) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageListResponseMessageEffect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The sender of this message as a full handle object
type ChatMessageListResponseMessageFromHandle struct {
	// Unique identifier for this handle
	ID string `json:"id,required" format:"uuid"`
	// Phone number (E.164) or email address of the participant
	Handle string `json:"handle,required"`
	// When this participant joined the chat
	JoinedAt time.Time `json:"joined_at,required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,required"`
	// Whether this handle belongs to the sender (your phone number)
	IsMe bool `json:"is_me,nullable"`
	// When they left (if applicable)
	LeftAt time.Time `json:"left_at,nullable" format:"date-time"`
	// Participant status
	//
	// Any of "active", "left", "removed".
	Status string `json:"status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Handle      respjson.Field
		JoinedAt    respjson.Field
		Service     respjson.Field
		IsMe        respjson.Field
		LeftAt      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponseMessageFromHandle) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageListResponseMessageFromHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatMessageListResponseMessagePartUnion contains all possible properties and
// values from [ChatMessageListResponseMessagePartTextPartResponse],
// [ChatMessageListResponseMessagePartMediaPartResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatMessageListResponseMessagePartUnion struct {
	// This field is a union of
	// [[]ChatMessageListResponseMessagePartTextPartResponseReaction],
	// [[]ChatMessageListResponseMessagePartMediaPartResponseReaction]
	Reactions ChatMessageListResponseMessagePartUnionReactions `json:"reactions"`
	Type      string                                           `json:"type"`
	// This field is from variant [ChatMessageListResponseMessagePartTextPartResponse].
	Value string `json:"value"`
	// This field is from variant
	// [ChatMessageListResponseMessagePartMediaPartResponse].
	ID string `json:"id"`
	// This field is from variant
	// [ChatMessageListResponseMessagePartMediaPartResponse].
	Filename string `json:"filename"`
	// This field is from variant
	// [ChatMessageListResponseMessagePartMediaPartResponse].
	MimeType string `json:"mime_type"`
	// This field is from variant
	// [ChatMessageListResponseMessagePartMediaPartResponse].
	SizeBytes int64 `json:"size_bytes"`
	// This field is from variant
	// [ChatMessageListResponseMessagePartMediaPartResponse].
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

func (u ChatMessageListResponseMessagePartUnion) AsChatMessageListResponseMessagePartTextPartResponse() (v ChatMessageListResponseMessagePartTextPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatMessageListResponseMessagePartUnion) AsChatMessageListResponseMessagePartMediaPartResponse() (v ChatMessageListResponseMessagePartMediaPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatMessageListResponseMessagePartUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatMessageListResponseMessagePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatMessageListResponseMessagePartUnionReactions is an implicit subunion of
// [ChatMessageListResponseMessagePartUnion].
// ChatMessageListResponseMessagePartUnionReactions provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatMessageListResponseMessagePartUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChatMessageListResponseMessagePartTextPartResponseReactions
// OfChatMessageListResponseMessagePartMediaPartResponseReactions]
type ChatMessageListResponseMessagePartUnionReactions struct {
	// This field will be present if the value is a
	// [[]ChatMessageListResponseMessagePartTextPartResponseReaction] instead of an
	// object.
	OfChatMessageListResponseMessagePartTextPartResponseReactions []ChatMessageListResponseMessagePartTextPartResponseReaction `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatMessageListResponseMessagePartMediaPartResponseReaction] instead of an
	// object.
	OfChatMessageListResponseMessagePartMediaPartResponseReactions []ChatMessageListResponseMessagePartMediaPartResponseReaction `json:",inline"`
	JSON                                                           struct {
		OfChatMessageListResponseMessagePartTextPartResponseReactions  respjson.Field
		OfChatMessageListResponseMessagePartMediaPartResponseReactions respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (r *ChatMessageListResponseMessagePartUnionReactions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text message part
type ChatMessageListResponseMessagePartTextPartResponse struct {
	// Reactions on this message part
	Reactions []ChatMessageListResponseMessagePartTextPartResponseReaction `json:"reactions,required"`
	// Indicates this is a text message part
	//
	// Any of "text".
	Type string `json:"type,required"`
	// The text content
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reactions   respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponseMessagePartTextPartResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageListResponseMessagePartTextPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageListResponseMessagePartTextPartResponseReaction struct {
	Handle ChatMessageListResponseMessagePartTextPartResponseReactionHandle `json:"handle,required"`
	// Whether this reaction is from the current user
	IsMe bool `json:"is_me,required"`
	// Type of reaction. Standard iMessage tapbacks are love, like, dislike, laugh,
	// emphasize, question. Custom emoji reactions have type "custom" with the actual
	// emoji in the custom_emoji field.
	//
	// Any of "love", "like", "dislike", "laugh", "emphasize", "question", "custom".
	Type string `json:"type,required"`
	// Custom emoji if type is "custom", null otherwise
	CustomEmoji string `json:"custom_emoji,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		IsMe        respjson.Field
		Type        respjson.Field
		CustomEmoji respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponseMessagePartTextPartResponseReaction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatMessageListResponseMessagePartTextPartResponseReaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageListResponseMessagePartTextPartResponseReactionHandle struct {
	// Unique identifier for this handle
	ID string `json:"id,required" format:"uuid"`
	// Phone number (E.164) or email address of the participant
	Handle string `json:"handle,required"`
	// When this participant joined the chat
	JoinedAt time.Time `json:"joined_at,required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,required"`
	// Whether this handle belongs to the sender (your phone number)
	IsMe bool `json:"is_me,nullable"`
	// When they left (if applicable)
	LeftAt time.Time `json:"left_at,nullable" format:"date-time"`
	// Participant status
	//
	// Any of "active", "left", "removed".
	Status string `json:"status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Handle      respjson.Field
		JoinedAt    respjson.Field
		Service     respjson.Field
		IsMe        respjson.Field
		LeftAt      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponseMessagePartTextPartResponseReactionHandle) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatMessageListResponseMessagePartTextPartResponseReactionHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A media attachment part
type ChatMessageListResponseMessagePartMediaPartResponse struct {
	// Unique attachment identifier
	ID string `json:"id,required" format:"uuid"`
	// Original filename
	Filename string `json:"filename,required"`
	// MIME type of the file
	MimeType string `json:"mime_type,required"`
	// Reactions on this message part
	Reactions []ChatMessageListResponseMessagePartMediaPartResponseReaction `json:"reactions,required"`
	// File size in bytes
	SizeBytes int64 `json:"size_bytes,required"`
	// Indicates this is a media attachment part
	//
	// Any of "media".
	Type string `json:"type,required"`
	// Presigned URL for downloading the attachment (expires in 1 hour).
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Filename    respjson.Field
		MimeType    respjson.Field
		Reactions   respjson.Field
		SizeBytes   respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponseMessagePartMediaPartResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageListResponseMessagePartMediaPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageListResponseMessagePartMediaPartResponseReaction struct {
	Handle ChatMessageListResponseMessagePartMediaPartResponseReactionHandle `json:"handle,required"`
	// Whether this reaction is from the current user
	IsMe bool `json:"is_me,required"`
	// Type of reaction. Standard iMessage tapbacks are love, like, dislike, laugh,
	// emphasize, question. Custom emoji reactions have type "custom" with the actual
	// emoji in the custom_emoji field.
	//
	// Any of "love", "like", "dislike", "laugh", "emphasize", "question", "custom".
	Type string `json:"type,required"`
	// Custom emoji if type is "custom", null otherwise
	CustomEmoji string `json:"custom_emoji,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		IsMe        respjson.Field
		Type        respjson.Field
		CustomEmoji respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponseMessagePartMediaPartResponseReaction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatMessageListResponseMessagePartMediaPartResponseReaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageListResponseMessagePartMediaPartResponseReactionHandle struct {
	// Unique identifier for this handle
	ID string `json:"id,required" format:"uuid"`
	// Phone number (E.164) or email address of the participant
	Handle string `json:"handle,required"`
	// When this participant joined the chat
	JoinedAt time.Time `json:"joined_at,required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,required"`
	// Whether this handle belongs to the sender (your phone number)
	IsMe bool `json:"is_me,nullable"`
	// When they left (if applicable)
	LeftAt time.Time `json:"left_at,nullable" format:"date-time"`
	// Participant status
	//
	// Any of "active", "left", "removed".
	Status string `json:"status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Handle      respjson.Field
		JoinedAt    respjson.Field
		Service     respjson.Field
		IsMe        respjson.Field
		LeftAt      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageListResponseMessagePartMediaPartResponseReactionHandle) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatMessageListResponseMessagePartMediaPartResponseReactionHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates this message is a threaded reply to another message
type ChatMessageListResponseMessageReplyTo struct {
	// The ID of the message to reply to
	MessageID string `json:"message_id,required" format:"uuid"`
	// The specific message part to reply to (0-based index). Defaults to 0 (first
	// part) if not provided. Use this when replying to a specific part of a multipart
	// message.
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
func (r ChatMessageListResponseMessageReplyTo) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageListResponseMessageReplyTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for sending a message to a chat
type ChatMessageSendResponse struct {
	// Unique identifier of the chat this message was sent to
	ChatID string `json:"chat_id,required" format:"uuid"`
	// A message that was sent (used in CreateChat and SendMessage responses)
	Message ChatMessageSendResponseMessage `json:"message,required"`
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

// A message that was sent (used in CreateChat and SendMessage responses)
type ChatMessageSendResponseMessage struct {
	// Message identifier (UUID)
	ID string `json:"id,required" format:"uuid"`
	// Current delivery status of a message
	//
	// Any of "pending", "queued", "sent", "delivered", "failed".
	DeliveryStatus string `json:"delivery_status,required"`
	// Whether the message has been read
	IsRead bool `json:"is_read,required"`
	// Message parts in order (text and media)
	Parts []ChatMessageSendResponseMessagePartUnion `json:"parts,required"`
	// When the message was sent
	SentAt time.Time `json:"sent_at,required" format:"date-time"`
	// When the message was delivered
	DeliveredAt time.Time `json:"delivered_at,nullable" format:"date-time"`
	// iMessage effect applied to a message (screen or bubble effect)
	Effect ChatMessageSendResponseMessageEffect `json:"effect,nullable"`
	// The sender of this message as a full handle object
	FromHandle ChatMessageSendResponseMessageFromHandle `json:"from_handle,nullable"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	PreferredService string `json:"preferred_service,nullable"`
	// Indicates this message is a threaded reply to another message
	ReplyTo ChatMessageSendResponseMessageReplyTo `json:"reply_to,nullable"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,nullable"`
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
func (r ChatMessageSendResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageSendResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatMessageSendResponseMessagePartUnion contains all possible properties and
// values from [ChatMessageSendResponseMessagePartTextPartResponse],
// [ChatMessageSendResponseMessagePartMediaPartResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatMessageSendResponseMessagePartUnion struct {
	// This field is a union of
	// [[]ChatMessageSendResponseMessagePartTextPartResponseReaction],
	// [[]ChatMessageSendResponseMessagePartMediaPartResponseReaction]
	Reactions ChatMessageSendResponseMessagePartUnionReactions `json:"reactions"`
	Type      string                                           `json:"type"`
	// This field is from variant [ChatMessageSendResponseMessagePartTextPartResponse].
	Value string `json:"value"`
	// This field is from variant
	// [ChatMessageSendResponseMessagePartMediaPartResponse].
	ID string `json:"id"`
	// This field is from variant
	// [ChatMessageSendResponseMessagePartMediaPartResponse].
	Filename string `json:"filename"`
	// This field is from variant
	// [ChatMessageSendResponseMessagePartMediaPartResponse].
	MimeType string `json:"mime_type"`
	// This field is from variant
	// [ChatMessageSendResponseMessagePartMediaPartResponse].
	SizeBytes int64 `json:"size_bytes"`
	// This field is from variant
	// [ChatMessageSendResponseMessagePartMediaPartResponse].
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

func (u ChatMessageSendResponseMessagePartUnion) AsChatMessageSendResponseMessagePartTextPartResponse() (v ChatMessageSendResponseMessagePartTextPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatMessageSendResponseMessagePartUnion) AsChatMessageSendResponseMessagePartMediaPartResponse() (v ChatMessageSendResponseMessagePartMediaPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatMessageSendResponseMessagePartUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatMessageSendResponseMessagePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatMessageSendResponseMessagePartUnionReactions is an implicit subunion of
// [ChatMessageSendResponseMessagePartUnion].
// ChatMessageSendResponseMessagePartUnionReactions provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatMessageSendResponseMessagePartUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChatMessageSendResponseMessagePartTextPartResponseReactions
// OfChatMessageSendResponseMessagePartMediaPartResponseReactions]
type ChatMessageSendResponseMessagePartUnionReactions struct {
	// This field will be present if the value is a
	// [[]ChatMessageSendResponseMessagePartTextPartResponseReaction] instead of an
	// object.
	OfChatMessageSendResponseMessagePartTextPartResponseReactions []ChatMessageSendResponseMessagePartTextPartResponseReaction `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatMessageSendResponseMessagePartMediaPartResponseReaction] instead of an
	// object.
	OfChatMessageSendResponseMessagePartMediaPartResponseReactions []ChatMessageSendResponseMessagePartMediaPartResponseReaction `json:",inline"`
	JSON                                                           struct {
		OfChatMessageSendResponseMessagePartTextPartResponseReactions  respjson.Field
		OfChatMessageSendResponseMessagePartMediaPartResponseReactions respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (r *ChatMessageSendResponseMessagePartUnionReactions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text message part
type ChatMessageSendResponseMessagePartTextPartResponse struct {
	// Reactions on this message part
	Reactions []ChatMessageSendResponseMessagePartTextPartResponseReaction `json:"reactions,required"`
	// Indicates this is a text message part
	//
	// Any of "text".
	Type string `json:"type,required"`
	// The text content
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reactions   respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageSendResponseMessagePartTextPartResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageSendResponseMessagePartTextPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageSendResponseMessagePartTextPartResponseReaction struct {
	Handle ChatMessageSendResponseMessagePartTextPartResponseReactionHandle `json:"handle,required"`
	// Whether this reaction is from the current user
	IsMe bool `json:"is_me,required"`
	// Type of reaction. Standard iMessage tapbacks are love, like, dislike, laugh,
	// emphasize, question. Custom emoji reactions have type "custom" with the actual
	// emoji in the custom_emoji field.
	//
	// Any of "love", "like", "dislike", "laugh", "emphasize", "question", "custom".
	Type string `json:"type,required"`
	// Custom emoji if type is "custom", null otherwise
	CustomEmoji string `json:"custom_emoji,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		IsMe        respjson.Field
		Type        respjson.Field
		CustomEmoji respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageSendResponseMessagePartTextPartResponseReaction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatMessageSendResponseMessagePartTextPartResponseReaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageSendResponseMessagePartTextPartResponseReactionHandle struct {
	// Unique identifier for this handle
	ID string `json:"id,required" format:"uuid"`
	// Phone number (E.164) or email address of the participant
	Handle string `json:"handle,required"`
	// When this participant joined the chat
	JoinedAt time.Time `json:"joined_at,required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,required"`
	// Whether this handle belongs to the sender (your phone number)
	IsMe bool `json:"is_me,nullable"`
	// When they left (if applicable)
	LeftAt time.Time `json:"left_at,nullable" format:"date-time"`
	// Participant status
	//
	// Any of "active", "left", "removed".
	Status string `json:"status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Handle      respjson.Field
		JoinedAt    respjson.Field
		Service     respjson.Field
		IsMe        respjson.Field
		LeftAt      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageSendResponseMessagePartTextPartResponseReactionHandle) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatMessageSendResponseMessagePartTextPartResponseReactionHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A media attachment part
type ChatMessageSendResponseMessagePartMediaPartResponse struct {
	// Unique attachment identifier
	ID string `json:"id,required" format:"uuid"`
	// Original filename
	Filename string `json:"filename,required"`
	// MIME type of the file
	MimeType string `json:"mime_type,required"`
	// Reactions on this message part
	Reactions []ChatMessageSendResponseMessagePartMediaPartResponseReaction `json:"reactions,required"`
	// File size in bytes
	SizeBytes int64 `json:"size_bytes,required"`
	// Indicates this is a media attachment part
	//
	// Any of "media".
	Type string `json:"type,required"`
	// Presigned URL for downloading the attachment (expires in 1 hour).
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Filename    respjson.Field
		MimeType    respjson.Field
		Reactions   respjson.Field
		SizeBytes   respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageSendResponseMessagePartMediaPartResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageSendResponseMessagePartMediaPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageSendResponseMessagePartMediaPartResponseReaction struct {
	Handle ChatMessageSendResponseMessagePartMediaPartResponseReactionHandle `json:"handle,required"`
	// Whether this reaction is from the current user
	IsMe bool `json:"is_me,required"`
	// Type of reaction. Standard iMessage tapbacks are love, like, dislike, laugh,
	// emphasize, question. Custom emoji reactions have type "custom" with the actual
	// emoji in the custom_emoji field.
	//
	// Any of "love", "like", "dislike", "laugh", "emphasize", "question", "custom".
	Type string `json:"type,required"`
	// Custom emoji if type is "custom", null otherwise
	CustomEmoji string `json:"custom_emoji,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		IsMe        respjson.Field
		Type        respjson.Field
		CustomEmoji respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageSendResponseMessagePartMediaPartResponseReaction) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatMessageSendResponseMessagePartMediaPartResponseReaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageSendResponseMessagePartMediaPartResponseReactionHandle struct {
	// Unique identifier for this handle
	ID string `json:"id,required" format:"uuid"`
	// Phone number (E.164) or email address of the participant
	Handle string `json:"handle,required"`
	// When this participant joined the chat
	JoinedAt time.Time `json:"joined_at,required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,required"`
	// Whether this handle belongs to the sender (your phone number)
	IsMe bool `json:"is_me,nullable"`
	// When they left (if applicable)
	LeftAt time.Time `json:"left_at,nullable" format:"date-time"`
	// Participant status
	//
	// Any of "active", "left", "removed".
	Status string `json:"status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Handle      respjson.Field
		JoinedAt    respjson.Field
		Service     respjson.Field
		IsMe        respjson.Field
		LeftAt      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageSendResponseMessagePartMediaPartResponseReactionHandle) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatMessageSendResponseMessagePartMediaPartResponseReactionHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// iMessage effect applied to a message (screen or bubble effect)
type ChatMessageSendResponseMessageEffect struct {
	// Name of the effect. Common values:
	//
	//   - Screen effects: confetti, fireworks, lasers, sparkles, celebration, hearts,
	//     love, balloons, happy_birthday, echo, spotlight
	//   - Bubble effects: slam, loud, gentle, invisible
	Name string `json:"name"`
	// Type of effect
	//
	// Any of "screen", "bubble".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageSendResponseMessageEffect) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageSendResponseMessageEffect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The sender of this message as a full handle object
type ChatMessageSendResponseMessageFromHandle struct {
	// Unique identifier for this handle
	ID string `json:"id,required" format:"uuid"`
	// Phone number (E.164) or email address of the participant
	Handle string `json:"handle,required"`
	// When this participant joined the chat
	JoinedAt time.Time `json:"joined_at,required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,required"`
	// Whether this handle belongs to the sender (your phone number)
	IsMe bool `json:"is_me,nullable"`
	// When they left (if applicable)
	LeftAt time.Time `json:"left_at,nullable" format:"date-time"`
	// Participant status
	//
	// Any of "active", "left", "removed".
	Status string `json:"status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Handle      respjson.Field
		JoinedAt    respjson.Field
		Service     respjson.Field
		IsMe        respjson.Field
		LeftAt      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageSendResponseMessageFromHandle) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageSendResponseMessageFromHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates this message is a threaded reply to another message
type ChatMessageSendResponseMessageReplyTo struct {
	// The ID of the message to reply to
	MessageID string `json:"message_id,required" format:"uuid"`
	// The specific message part to reply to (0-based index). Defaults to 0 (first
	// part) if not provided. Use this when replying to a specific part of a multipart
	// message.
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
func (r ChatMessageSendResponseMessageReplyTo) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageSendResponseMessageReplyTo) UnmarshalJSON(data []byte) error {
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
	Message ChatMessageSendParamsMessage `json:"message,omitzero,required"`
	paramObj
}

func (r ChatMessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message content container. Groups all message-related fields together,
// separating the "what" (message content) from the "where" (routing fields like
// from/to).
//
// The property Parts is required.
type ChatMessageSendParamsMessage struct {
	// Array of message parts. Each part can be either text or media. Parts are
	// displayed in order. Text and media can be mixed.
	//
	// **Supported Media:**
	//
	//   - Images: .jpg, .jpeg, .png, .gif, .heic, .heif, .tif, .tiff, .bmp
	//   - Videos: .mp4, .mov, .m4v, .mpeg, .mpg, .3gp
	//   - Audio: .m4a, .mp3, .aac, .caf, .wav, .aiff, .amr
	//   - Documents: .pdf, .txt, .rtf, .csv, .doc, .docx, .xls, .xlsx, .ppt, .pptx,
	//     .pages, .numbers, .key, .epub, .zip, .html, .htm
	//   - Contact & Calendar: .vcf, .ics
	//
	// **Audio:**
	//
	//   - Audio files (.m4a, .mp3, .aac, .caf, .wav, .aiff, .amr) are fully supported as
	//     media parts
	//   - To send audio as an **iMessage voice memo bubble** (inline playback UI), use
	//     the dedicated `/v3/chats/{chatId}/voicememo` endpoint instead
	//
	// **Validation Rule:** Consecutive text parts are not allowed. Text parts must be
	// separated by media parts. For example, [text, text] is invalid, but [text,
	// media, text] is valid.
	Parts []ChatMessageSendParamsMessagePartUnion `json:"parts,omitzero,required"`
	// Optional idempotency key for this message. Use this to prevent duplicate sends
	// of the same message.
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// iMessage effect to apply to this message (screen or bubble effect)
	Effect ChatMessageSendParamsMessageEffect `json:"effect,omitzero"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	PreferredService string `json:"preferred_service,omitzero"`
	// Reply to another message to create a threaded conversation
	ReplyTo ChatMessageSendParamsMessageReplyTo `json:"reply_to,omitzero"`
	paramObj
}

func (r ChatMessageSendParamsMessage) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParamsMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParamsMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatMessageSendParamsMessage](
		"preferred_service", "iMessage", "SMS", "RCS",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatMessageSendParamsMessagePartUnion struct {
	OfText  *ChatMessageSendParamsMessagePartText  `json:",omitzero,inline"`
	OfMedia *ChatMessageSendParamsMessagePartMedia `json:",omitzero,inline"`
	paramUnion
}

func (u ChatMessageSendParamsMessagePartUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfMedia)
}
func (u *ChatMessageSendParamsMessagePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ChatMessageSendParamsMessagePartUnion](
		"type",
		apijson.Discriminator[ChatMessageSendParamsMessagePartText]("text"),
		apijson.Discriminator[ChatMessageSendParamsMessagePartMedia]("media"),
	)
}

// The properties Type, Value are required.
type ChatMessageSendParamsMessagePartText struct {
	// The text content
	Value string `json:"value,required"`
	// Optional idempotency key for this specific message part. Use this to prevent
	// duplicate sends of the same part.
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Indicates this is a text message part
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ChatMessageSendParamsMessagePartText) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParamsMessagePartText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParamsMessagePartText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ChatMessageSendParamsMessagePartMedia struct {
	// Reference to a file pre-uploaded via `POST /v3/attachments` (optional). The file
	// is already stored, so sends using this ID skip the download step — useful when
	// sending the same file to many recipients.
	//
	// Either `url` or `attachment_id` must be provided, but not both.
	AttachmentID param.Opt[string] `json:"attachment_id,omitzero" format:"uuid"`
	// Optional idempotency key for this specific message part. Use this to prevent
	// duplicate sends of the same part.
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Any publicly accessible HTTPS URL to the media file. The server downloads and
	// sends the file automatically — no pre-upload step required.
	//
	// **Size limit:** 10MB maximum for URL-based downloads. For larger files (up to
	// 100MB), use the pre-upload flow: `POST /v3/attachments` to get a presigned URL,
	// upload directly, then reference by `attachment_id`.
	//
	// **Requirements:**
	//
	//   - URL must use HTTPS
	//   - File content must be a supported format (the server validates the actual file
	//     content)
	//
	// **Supported formats:**
	//
	//   - Images: .jpg, .jpeg, .png, .gif, .heic, .heif, .tif, .tiff, .bmp
	//   - Videos: .mp4, .mov, .m4v, .mpeg, .mpg, .3gp
	//   - Audio: .m4a, .mp3, .aac, .caf, .wav, .aiff, .amr
	//   - Documents: .pdf, .txt, .rtf, .csv, .doc, .docx, .xls, .xlsx, .ppt, .pptx,
	//     .pages, .numbers, .key, .epub, .zip, .html, .htm
	//   - Contact & Calendar: .vcf, .ics
	//
	// **Tip:** Audio sent here appears as a regular file attachment. To send audio as
	// an iMessage voice memo bubble (with inline playback), use
	// `/v3/chats/{chatId}/voicememo`. For repeated sends of the same file, use
	// `attachment_id` to avoid redundant downloads.
	//
	// Either `url` or `attachment_id` must be provided, but not both.
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Indicates this is a media attachment part
	//
	// This field can be elided, and will marshal its zero value as "media".
	Type constant.Media `json:"type,required"`
	paramObj
}

func (r ChatMessageSendParamsMessagePartMedia) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParamsMessagePartMedia
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParamsMessagePartMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// iMessage effect to apply to this message (screen or bubble effect)
type ChatMessageSendParamsMessageEffect struct {
	// Name of the effect. Common values:
	//
	//   - Screen effects: confetti, fireworks, lasers, sparkles, celebration, hearts,
	//     love, balloons, happy_birthday, echo, spotlight
	//   - Bubble effects: slam, loud, gentle, invisible
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of effect
	//
	// Any of "screen", "bubble".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatMessageSendParamsMessageEffect) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParamsMessageEffect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParamsMessageEffect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatMessageSendParamsMessageEffect](
		"type", "screen", "bubble",
	)
}

// Reply to another message to create a threaded conversation
//
// The property MessageID is required.
type ChatMessageSendParamsMessageReplyTo struct {
	// The ID of the message to reply to
	MessageID string `json:"message_id,required" format:"uuid"`
	// The specific message part to reply to (0-based index). Defaults to 0 (first
	// part) if not provided. Use this when replying to a specific part of a multipart
	// message.
	PartIndex param.Opt[int64] `json:"part_index,omitzero"`
	paramObj
}

func (r ChatMessageSendParamsMessageReplyTo) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageSendParamsMessageReplyTo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageSendParamsMessageReplyTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
