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

// MessageService contains methods and other services that help with interacting
// with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.Options = opts
	return
}

// Retrieve a specific message by its ID. This endpoint returns the full message
// details including text, attachments, reactions, and metadata.
func (r *MessageService) Get(ctx context.Context, messageID string, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("v3/messages/%s", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a message from the Linq API only. This does NOT unsend or remove the
// message from the actual chat - recipients will still see the message.
//
// Use this endpoint to remove messages from your records and prevent them from
// appearing in API responses.
func (r *MessageService) Delete(ctx context.Context, messageID string, body MessageDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("v3/messages/%s", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Add or remove emoji reactions to messages. Reactions let users express their
// response to a message without sending a new message.
//
// **Supported Reactions:**
//
// - love ❤️
// - like 👍
// - dislike 👎
// - laugh 😂
// - emphasize ‼️
// - question ❓
// - custom - any emoji (use `custom_emoji` field to specify)
func (r *MessageService) AddReaction(ctx context.Context, messageID string, body MessageAddReactionParams, opts ...option.RequestOption) (res *Reaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("v3/messages/%s/reactions", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve all messages in a conversation thread. Given any message ID in the
// thread, returns the originator message and all replies in chronological order.
//
// If the message is not part of a thread, returns just that single message.
//
// Supports pagination and configurable ordering.
func (r *MessageService) GetThread(ctx context.Context, messageID string, query MessageGetThreadParams, opts ...option.RequestOption) (res *MessageGetThreadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("v3/messages/%s/thread", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ChatHandle struct {
	// Unique identifier for this handle
	ID string `json:"id,required" format:"uuid"`
	// Phone number (E.164) or email address of the participant
	Handle string `json:"handle,required"`
	// When this participant joined the chat
	JoinedAt time.Time `json:"joined_at,required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service ChatHandleService `json:"service,required"`
	// Whether this handle belongs to the sender (your phone number)
	IsMe bool `json:"is_me,nullable"`
	// When they left (if applicable)
	LeftAt time.Time `json:"left_at,nullable" format:"date-time"`
	// Participant status
	//
	// Any of "active", "left", "removed".
	Status ChatHandleStatus `json:"status,nullable"`
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
func (r ChatHandle) RawJSON() string { return r.JSON.raw }
func (r *ChatHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging service type
type ChatHandleService string

const (
	ChatHandleServiceIMessage ChatHandleService = "iMessage"
	ChatHandleServiceSMS      ChatHandleService = "SMS"
	ChatHandleServiceRcs      ChatHandleService = "RCS"
)

// Participant status
type ChatHandleStatus string

const (
	ChatHandleStatusActive  ChatHandleStatus = "active"
	ChatHandleStatusLeft    ChatHandleStatus = "left"
	ChatHandleStatusRemoved ChatHandleStatus = "removed"
)

// A media attachment part
type MediaPart struct {
	// Unique attachment identifier
	ID string `json:"id,required" format:"uuid"`
	// Original filename
	Filename string `json:"filename,required"`
	// MIME type of the file
	MimeType string `json:"mime_type,required"`
	// Reactions on this message part
	Reactions []Reaction `json:"reactions,required"`
	// File size in bytes
	SizeBytes int64 `json:"size_bytes,required"`
	// Indicates this is a media attachment part
	//
	// Any of "media".
	Type MediaPartType `json:"type,required"`
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
func (r MediaPart) RawJSON() string { return r.JSON.raw }
func (r *MediaPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates this is a media attachment part
type MediaPartType string

const (
	MediaPartTypeMedia MediaPartType = "media"
)

type Message struct {
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
	Effect MessageEffect `json:"effect,nullable"`
	// DEPRECATED: Use from_handle instead. Phone number of the message sender.
	//
	// Deprecated: deprecated
	From string `json:"from,nullable"`
	// The sender of this message as a full handle object
	FromHandle ChatHandle `json:"from_handle,nullable"`
	// Message parts in order (text and media)
	Parts []MessagePartUnion `json:"parts,nullable"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	PreferredService MessagePreferredService `json:"preferred_service,nullable"`
	// When the message was read
	ReadAt time.Time `json:"read_at,nullable" format:"date-time"`
	// Indicates this message is a threaded reply to another message
	ReplyTo ReplyTo `json:"reply_to,nullable"`
	// When the message was sent
	SentAt time.Time `json:"sent_at,nullable" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service MessageService `json:"service,nullable"`
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
func (r Message) RawJSON() string { return r.JSON.raw }
func (r *Message) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessagePartUnion contains all possible properties and values from [TextPart],
// [MediaPart].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessagePartUnion struct {
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

func (u MessagePartUnion) AsTextPart() (v TextPart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessagePartUnion) AsMediaPart() (v MediaPart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessagePartUnion) RawJSON() string { return u.JSON.raw }

func (r *MessagePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging service type
type MessagePreferredService string

const (
	MessagePreferredServiceIMessage MessagePreferredService = "iMessage"
	MessagePreferredServiceSMS      MessagePreferredService = "SMS"
	MessagePreferredServiceRcs      MessagePreferredService = "RCS"
)

// Messaging service type
type MessageService string

const (
	MessageServiceIMessage MessageService = "iMessage"
	MessageServiceSMS      MessageService = "SMS"
	MessageServiceRcs      MessageService = "RCS"
)

// iMessage effect applied to a message (screen or bubble effect)
type MessageEffect struct {
	// Name of the effect. Common values:
	//
	//   - Screen effects: confetti, fireworks, lasers, sparkles, celebration, hearts,
	//     love, balloons, happy_birthday, echo, spotlight
	//   - Bubble effects: slam, loud, gentle, invisible
	Name string `json:"name"`
	// Type of effect
	//
	// Any of "screen", "bubble".
	Type MessageEffectType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEffect) RawJSON() string { return r.JSON.raw }
func (r *MessageEffect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageEffect to a MessageEffectParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageEffectParam.Overrides()
func (r MessageEffect) ToParam() MessageEffectParam {
	return param.Override[MessageEffectParam](json.RawMessage(r.RawJSON()))
}

// Type of effect
type MessageEffectType string

const (
	MessageEffectTypeScreen MessageEffectType = "screen"
	MessageEffectTypeBubble MessageEffectType = "bubble"
)

// iMessage effect applied to a message (screen or bubble effect)
type MessageEffectParam struct {
	// Name of the effect. Common values:
	//
	//   - Screen effects: confetti, fireworks, lasers, sparkles, celebration, hearts,
	//     love, balloons, happy_birthday, echo, spotlight
	//   - Bubble effects: slam, loud, gentle, invisible
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of effect
	//
	// Any of "screen", "bubble".
	Type MessageEffectType `json:"type,omitzero"`
	paramObj
}

func (r MessageEffectParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageEffectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageEffectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Reaction struct {
	Handle ChatHandle `json:"handle,required"`
	// Whether this reaction is from the current user
	IsMe bool `json:"is_me,required"`
	// Type of reaction. Standard iMessage tapbacks are love, like, dislike, laugh,
	// emphasize, question. Custom emoji reactions have type "custom" with the actual
	// emoji in the custom_emoji field. Sticker reactions have type "sticker" with
	// sticker attachment details in the sticker field.
	//
	// Any of "love", "like", "dislike", "laugh", "emphasize", "question", "custom",
	// "sticker".
	Type ReactionType `json:"type,required"`
	// Custom emoji if type is "custom", null otherwise
	CustomEmoji string `json:"custom_emoji,nullable"`
	// Sticker attachment details when reaction_type is "sticker". Null for non-sticker
	// reactions.
	Sticker ReactionSticker `json:"sticker,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		IsMe        respjson.Field
		Type        respjson.Field
		CustomEmoji respjson.Field
		Sticker     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Reaction) RawJSON() string { return r.JSON.raw }
func (r *Reaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sticker attachment details when reaction_type is "sticker". Null for non-sticker
// reactions.
type ReactionSticker struct {
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
func (r ReactionSticker) RawJSON() string { return r.JSON.raw }
func (r *ReactionSticker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of reaction. Standard iMessage tapbacks are love, like, dislike, laugh,
// emphasize, question. Custom emoji reactions have type "custom" with the actual
// emoji in the custom_emoji field. Sticker reactions have type "sticker" with
// sticker attachment details in the sticker field.
type ReactionType string

const (
	ReactionTypeLove      ReactionType = "love"
	ReactionTypeLike      ReactionType = "like"
	ReactionTypeDislike   ReactionType = "dislike"
	ReactionTypeLaugh     ReactionType = "laugh"
	ReactionTypeEmphasize ReactionType = "emphasize"
	ReactionTypeQuestion  ReactionType = "question"
	ReactionTypeCustom    ReactionType = "custom"
	ReactionTypeSticker   ReactionType = "sticker"
)

// Indicates this message is a threaded reply to another message
type ReplyTo struct {
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
func (r ReplyTo) RawJSON() string { return r.JSON.raw }
func (r *ReplyTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ReplyTo to a ReplyToParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ReplyToParam.Overrides()
func (r ReplyTo) ToParam() ReplyToParam {
	return param.Override[ReplyToParam](json.RawMessage(r.RawJSON()))
}

// Indicates this message is a threaded reply to another message
//
// The property MessageID is required.
type ReplyToParam struct {
	// The ID of the message to reply to
	MessageID string `json:"message_id,required" format:"uuid"`
	// The specific message part to reply to (0-based index). Defaults to 0 (first
	// part) if not provided. Use this when replying to a specific part of a multipart
	// message.
	PartIndex param.Opt[int64] `json:"part_index,omitzero"`
	paramObj
}

func (r ReplyToParam) MarshalJSON() (data []byte, err error) {
	type shadow ReplyToParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReplyToParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text message part
type TextPart struct {
	// Reactions on this message part
	Reactions []Reaction `json:"reactions,required"`
	// Indicates this is a text message part
	//
	// Any of "text".
	Type TextPartType `json:"type,required"`
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
func (r TextPart) RawJSON() string { return r.JSON.raw }
func (r *TextPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates this is a text message part
type TextPartType string

const (
	TextPartTypeText TextPartType = "text"
)

// Response containing messages in a thread with pagination
type MessageGetThreadResponse struct {
	// Messages in the thread, ordered by the specified order parameter
	Messages []Message `json:"messages,required"`
	// Cursor for fetching the next page of results (null if no more results)
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
func (r MessageGetThreadResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageGetThreadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageDeleteParams struct {
	// ID of the chat containing the message to delete
	ChatID string `json:"chat_id,required" format:"uuid"`
	paramObj
}

func (r MessageDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageAddReactionParams struct {
	// Whether to add or remove the reaction
	//
	// Any of "add", "remove".
	Operation MessageAddReactionParamsOperation `json:"operation,omitzero,required"`
	// Type of reaction. Standard iMessage tapbacks are love, like, dislike, laugh,
	// emphasize, question. Custom emoji reactions have type "custom" with the actual
	// emoji in the custom_emoji field. Sticker reactions have type "sticker" with
	// sticker attachment details in the sticker field.
	//
	// Any of "love", "like", "dislike", "laugh", "emphasize", "question", "custom",
	// "sticker".
	Type ReactionType `json:"type,omitzero,required"`
	// Custom emoji string. Required when type is "custom".
	CustomEmoji param.Opt[string] `json:"custom_emoji,omitzero"`
	// Optional index of the message part to react to. If not provided, reacts to the
	// entire message (part 0).
	PartIndex param.Opt[int64] `json:"part_index,omitzero"`
	paramObj
}

func (r MessageAddReactionParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageAddReactionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageAddReactionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether to add or remove the reaction
type MessageAddReactionParamsOperation string

const (
	MessageAddReactionParamsOperationAdd    MessageAddReactionParamsOperation = "add"
	MessageAddReactionParamsOperationRemove MessageAddReactionParamsOperation = "remove"
)

type MessageGetThreadParams struct {
	// Pagination cursor from previous next_cursor response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of messages to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order for messages (asc = oldest first, desc = newest first)
	//
	// Any of "asc", "desc".
	Order MessageGetThreadParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageGetThreadParams]'s query parameters as `url.Values`.
func (r MessageGetThreadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order for messages (asc = oldest first, desc = newest first)
type MessageGetThreadParamsOrder string

const (
	MessageGetThreadParamsOrderAsc  MessageGetThreadParamsOrder = "asc"
	MessageGetThreadParamsOrderDesc MessageGetThreadParamsOrder = "desc"
)
