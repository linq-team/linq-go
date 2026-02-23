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

// ChatService contains methods and other services that help with interacting with
// the linq API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	Options      []option.RequestOption
	Participants ChatParticipantService
	Typing       ChatTypingService
	Messages     ChatMessageService
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r ChatService) {
	r = ChatService{}
	r.Options = opts
	r.Participants = NewChatParticipantService(opts...)
	r.Typing = NewChatTypingService(opts...)
	r.Messages = NewChatMessageService(opts...)
	return
}

// Create a new chat with specified participants and send an initial message. The
// initial message is required when creating a chat.
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
func (r *ChatService) New(ctx context.Context, body ChatNewParams, opts ...option.RequestOption) (res *ChatNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/chats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a chat by its unique identifier.
func (r *ChatService) Get(ctx context.Context, chatID string, opts ...option.RequestOption) (res *ChatGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return
	}
	path := fmt.Sprintf("v3/chats/%s", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update chat properties such as display name and group chat icon.
func (r *ChatService) Update(ctx context.Context, chatID string, body ChatUpdateParams, opts ...option.RequestOption) (res *ChatUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return
	}
	path := fmt.Sprintf("v3/chats/%s", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of chats for the authenticated partner filtered by
// phone number. Returns all chats involving the specified phone number with their
// participants and recent activity.
//
// **Pagination:**
//
// - Use `limit` to control page size (default: 20, max: 100)
// - The response includes `next_cursor` for fetching the next page
// - When `next_cursor` is `null`, there are no more results to fetch
// - Pass the `next_cursor` value as the `cursor` parameter for the next request
//
// **Example pagination flow:**
//
// 1. First request: `GET /v3/chats?from=%2B12223334444&limit=20`
// 2. Response includes `next_cursor: "20"` (more results exist)
// 3. Next request: `GET /v3/chats?from=%2B12223334444&limit=20&cursor=20`
// 4. Response includes `next_cursor: null` (no more results)
func (r *ChatService) List(ctx context.Context, query ChatListParams, opts ...option.RequestOption) (res *ChatListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/chats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Mark all messages in a chat as read.
func (r *ChatService) MarkAsRead(ctx context.Context, chatID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return
	}
	path := fmt.Sprintf("v3/chats/%s/read", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Send an audio file as an **iMessage voice memo bubble** to all participants in a
// chat. Voice memos appear with iMessage's native inline playback UI, unlike
// regular audio attachments sent via media parts which appear as downloadable
// files.
//
// **Supported audio formats:**
//
// - MP3 (audio/mpeg)
// - M4A (audio/x-m4a, audio/mp4)
// - AAC (audio/aac)
// - CAF (audio/x-caf) - Core Audio Format
// - WAV (audio/wav)
// - AIFF (audio/aiff, audio/x-aiff)
// - AMR (audio/amr)
func (r *ChatService) SendVoicememo(ctx context.Context, chatID string, body ChatSendVoicememoParams, opts ...option.RequestOption) (res *ChatSendVoicememoResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return
	}
	path := fmt.Sprintf("v3/chats/%s/voicememo", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Share your contact information (Name and Photo Sharing) with a chat.
//
// **Note:** A contact card must be configured before sharing. You can set up your
// contact card on the
// [Linq dashboard](https://dashboard.linqapp.com/contact-cards).
func (r *ChatService) ShareContactCard(ctx context.Context, chatID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return
	}
	path := fmt.Sprintf("v3/chats/%s/share_contact_card", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Response for creating a new chat with an initial message
type ChatNewResponse struct {
	Chat ChatNewResponseChat `json:"chat,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chat        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatNewResponseChat struct {
	// Unique identifier for the created chat (UUID)
	ID string `json:"id,required" format:"uuid"`
	// Display name for the chat. Defaults to a comma-separated list of recipient
	// handles. Can be updated for group chats.
	DisplayName string `json:"display_name,required"`
	// List of participants in the chat. Always contains at least two handles (your
	// phone number and the other participant).
	Handles []ChatNewResponseChatHandle `json:"handles,required"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group,required"`
	// A message that was sent (used in CreateChat and SendMessage responses)
	Message ChatNewResponseChatMessage `json:"message,required"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DisplayName respjson.Field
		Handles     respjson.Field
		IsGroup     respjson.Field
		Message     respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatNewResponseChat) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponseChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatNewResponseChatHandle struct {
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
func (r ChatNewResponseChatHandle) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponseChatHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message that was sent (used in CreateChat and SendMessage responses)
type ChatNewResponseChatMessage struct {
	// Message identifier (UUID)
	ID string `json:"id,required" format:"uuid"`
	// Current delivery status of a message
	//
	// Any of "pending", "queued", "sent", "delivered", "failed".
	DeliveryStatus string `json:"delivery_status,required"`
	// Whether the message has been read
	IsRead bool `json:"is_read,required"`
	// Message parts in order (text and media)
	Parts []ChatNewResponseChatMessagePartUnion `json:"parts,required"`
	// When the message was sent
	SentAt time.Time `json:"sent_at,required" format:"date-time"`
	// When the message was delivered
	DeliveredAt time.Time `json:"delivered_at,nullable" format:"date-time"`
	// iMessage effect applied to a message (screen or bubble effect)
	Effect ChatNewResponseChatMessageEffect `json:"effect,nullable"`
	// The sender of this message as a full handle object
	FromHandle ChatNewResponseChatMessageFromHandle `json:"from_handle,nullable"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	PreferredService string `json:"preferred_service,nullable"`
	// Indicates this message is a threaded reply to another message
	ReplyTo ChatNewResponseChatMessageReplyTo `json:"reply_to,nullable"`
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
func (r ChatNewResponseChatMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponseChatMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatNewResponseChatMessagePartUnion contains all possible properties and values
// from [ChatNewResponseChatMessagePartTextPartResponse],
// [ChatNewResponseChatMessagePartMediaPartResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatNewResponseChatMessagePartUnion struct {
	// This field is a union of
	// [[]ChatNewResponseChatMessagePartTextPartResponseReaction],
	// [[]ChatNewResponseChatMessagePartMediaPartResponseReaction]
	Reactions ChatNewResponseChatMessagePartUnionReactions `json:"reactions"`
	Type      string                                       `json:"type"`
	// This field is from variant [ChatNewResponseChatMessagePartTextPartResponse].
	Value string `json:"value"`
	// This field is from variant [ChatNewResponseChatMessagePartMediaPartResponse].
	ID string `json:"id"`
	// This field is from variant [ChatNewResponseChatMessagePartMediaPartResponse].
	Filename string `json:"filename"`
	// This field is from variant [ChatNewResponseChatMessagePartMediaPartResponse].
	MimeType string `json:"mime_type"`
	// This field is from variant [ChatNewResponseChatMessagePartMediaPartResponse].
	SizeBytes int64 `json:"size_bytes"`
	// This field is from variant [ChatNewResponseChatMessagePartMediaPartResponse].
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

func (u ChatNewResponseChatMessagePartUnion) AsChatNewResponseChatMessagePartTextPartResponse() (v ChatNewResponseChatMessagePartTextPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatNewResponseChatMessagePartUnion) AsChatNewResponseChatMessagePartMediaPartResponse() (v ChatNewResponseChatMessagePartMediaPartResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatNewResponseChatMessagePartUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatNewResponseChatMessagePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatNewResponseChatMessagePartUnionReactions is an implicit subunion of
// [ChatNewResponseChatMessagePartUnion].
// ChatNewResponseChatMessagePartUnionReactions provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChatNewResponseChatMessagePartUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChatNewResponseChatMessagePartTextPartResponseReactions
// OfChatNewResponseChatMessagePartMediaPartResponseReactions]
type ChatNewResponseChatMessagePartUnionReactions struct {
	// This field will be present if the value is a
	// [[]ChatNewResponseChatMessagePartTextPartResponseReaction] instead of an object.
	OfChatNewResponseChatMessagePartTextPartResponseReactions []ChatNewResponseChatMessagePartTextPartResponseReaction `json:",inline"`
	// This field will be present if the value is a
	// [[]ChatNewResponseChatMessagePartMediaPartResponseReaction] instead of an
	// object.
	OfChatNewResponseChatMessagePartMediaPartResponseReactions []ChatNewResponseChatMessagePartMediaPartResponseReaction `json:",inline"`
	JSON                                                       struct {
		OfChatNewResponseChatMessagePartTextPartResponseReactions  respjson.Field
		OfChatNewResponseChatMessagePartMediaPartResponseReactions respjson.Field
		raw                                                        string
	} `json:"-"`
}

func (r *ChatNewResponseChatMessagePartUnionReactions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text message part
type ChatNewResponseChatMessagePartTextPartResponse struct {
	// Reactions on this message part
	Reactions []ChatNewResponseChatMessagePartTextPartResponseReaction `json:"reactions,required"`
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
func (r ChatNewResponseChatMessagePartTextPartResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponseChatMessagePartTextPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatNewResponseChatMessagePartTextPartResponseReaction struct {
	Handle ChatNewResponseChatMessagePartTextPartResponseReactionHandle `json:"handle,required"`
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
func (r ChatNewResponseChatMessagePartTextPartResponseReaction) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponseChatMessagePartTextPartResponseReaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatNewResponseChatMessagePartTextPartResponseReactionHandle struct {
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
func (r ChatNewResponseChatMessagePartTextPartResponseReactionHandle) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatNewResponseChatMessagePartTextPartResponseReactionHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A media attachment part
type ChatNewResponseChatMessagePartMediaPartResponse struct {
	// Unique attachment identifier
	ID string `json:"id,required" format:"uuid"`
	// Original filename
	Filename string `json:"filename,required"`
	// MIME type of the file
	MimeType string `json:"mime_type,required"`
	// Reactions on this message part
	Reactions []ChatNewResponseChatMessagePartMediaPartResponseReaction `json:"reactions,required"`
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
func (r ChatNewResponseChatMessagePartMediaPartResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponseChatMessagePartMediaPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatNewResponseChatMessagePartMediaPartResponseReaction struct {
	Handle ChatNewResponseChatMessagePartMediaPartResponseReactionHandle `json:"handle,required"`
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
func (r ChatNewResponseChatMessagePartMediaPartResponseReaction) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponseChatMessagePartMediaPartResponseReaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatNewResponseChatMessagePartMediaPartResponseReactionHandle struct {
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
func (r ChatNewResponseChatMessagePartMediaPartResponseReactionHandle) RawJSON() string {
	return r.JSON.raw
}
func (r *ChatNewResponseChatMessagePartMediaPartResponseReactionHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// iMessage effect applied to a message (screen or bubble effect)
type ChatNewResponseChatMessageEffect struct {
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
func (r ChatNewResponseChatMessageEffect) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponseChatMessageEffect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The sender of this message as a full handle object
type ChatNewResponseChatMessageFromHandle struct {
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
func (r ChatNewResponseChatMessageFromHandle) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponseChatMessageFromHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates this message is a threaded reply to another message
type ChatNewResponseChatMessageReplyTo struct {
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
func (r ChatNewResponseChatMessageReplyTo) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponseChatMessageReplyTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponse struct {
	// Unique identifier for the chat
	ID string `json:"id,required" format:"uuid"`
	// When the chat was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Display name for the chat. Defaults to a comma-separated list of recipient
	// handles. Can be updated for group chats.
	DisplayName string `json:"display_name,required"`
	// List of chat participants with full handle details. Always contains at least two
	// handles (your phone number and the other participant).
	Handles []ChatGetResponseHandle `json:"handles,required"`
	// Whether the chat is archived
	IsArchived bool `json:"is_archived,required"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group,required"`
	// When the chat was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service ChatGetResponseService `json:"service,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DisplayName respjson.Field
		Handles     respjson.Field
		IsArchived  respjson.Field
		IsGroup     respjson.Field
		UpdatedAt   respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponseHandle struct {
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
func (r ChatGetResponseHandle) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging service type
type ChatGetResponseService string

const (
	ChatGetResponseServiceIMessage ChatGetResponseService = "iMessage"
	ChatGetResponseServiceSMS      ChatGetResponseService = "SMS"
	ChatGetResponseServiceRcs      ChatGetResponseService = "RCS"
)

type ChatUpdateResponse struct {
	// Unique identifier for the chat
	ID string `json:"id,required" format:"uuid"`
	// When the chat was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Display name for the chat. Defaults to a comma-separated list of recipient
	// handles. Can be updated for group chats.
	DisplayName string `json:"display_name,required"`
	// List of chat participants with full handle details. Always contains at least two
	// handles (your phone number and the other participant).
	Handles []ChatUpdateResponseHandle `json:"handles,required"`
	// Whether the chat is archived
	IsArchived bool `json:"is_archived,required"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group,required"`
	// When the chat was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service ChatUpdateResponseService `json:"service,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DisplayName respjson.Field
		Handles     respjson.Field
		IsArchived  respjson.Field
		IsGroup     respjson.Field
		UpdatedAt   respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatUpdateResponseHandle struct {
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
func (r ChatUpdateResponseHandle) RawJSON() string { return r.JSON.raw }
func (r *ChatUpdateResponseHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging service type
type ChatUpdateResponseService string

const (
	ChatUpdateResponseServiceIMessage ChatUpdateResponseService = "iMessage"
	ChatUpdateResponseServiceSMS      ChatUpdateResponseService = "SMS"
	ChatUpdateResponseServiceRcs      ChatUpdateResponseService = "RCS"
)

type ChatListResponse struct {
	// List of chats
	Chats []ChatListResponseChat `json:"chats,required"`
	// Cursor for fetching the next page of results. Null if there are no more results
	// to fetch. Pass this value as the `cursor` parameter in the next request.
	NextCursor string `json:"next_cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chats       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatListResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatListResponseChat struct {
	// Unique identifier for the chat
	ID string `json:"id,required" format:"uuid"`
	// When the chat was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Display name for the chat. Defaults to a comma-separated list of recipient
	// handles. Can be updated for group chats.
	DisplayName string `json:"display_name,required"`
	// List of chat participants with full handle details. Always contains at least two
	// handles (your phone number and the other participant).
	Handles []ChatListResponseChatHandle `json:"handles,required"`
	// Whether the chat is archived
	IsArchived bool `json:"is_archived,required"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group,required"`
	// When the chat was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DisplayName respjson.Field
		Handles     respjson.Field
		IsArchived  respjson.Field
		IsGroup     respjson.Field
		UpdatedAt   respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatListResponseChat) RawJSON() string { return r.JSON.raw }
func (r *ChatListResponseChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatListResponseChatHandle struct {
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
func (r ChatListResponseChatHandle) RawJSON() string { return r.JSON.raw }
func (r *ChatListResponseChatHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for sending a voice memo to a chat
type ChatSendVoicememoResponse struct {
	VoiceMemo ChatSendVoicememoResponseVoiceMemo `json:"voice_memo,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VoiceMemo   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatSendVoicememoResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatSendVoicememoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatSendVoicememoResponseVoiceMemo struct {
	// Message identifier
	ID   string                                 `json:"id,required" format:"uuid"`
	Chat ChatSendVoicememoResponseVoiceMemoChat `json:"chat,required"`
	// When the voice memo was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Sender phone number
	From string `json:"from,required"`
	// Current delivery status
	Status string `json:"status,required"`
	// Recipient handles (phone numbers or email addresses)
	To        []string                                    `json:"to,required"`
	VoiceMemo ChatSendVoicememoResponseVoiceMemoVoiceMemo `json:"voice_memo,required"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Chat        respjson.Field
		CreatedAt   respjson.Field
		From        respjson.Field
		Status      respjson.Field
		To          respjson.Field
		VoiceMemo   respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatSendVoicememoResponseVoiceMemo) RawJSON() string { return r.JSON.raw }
func (r *ChatSendVoicememoResponseVoiceMemo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatSendVoicememoResponseVoiceMemoChat struct {
	// Chat identifier
	ID string `json:"id,required" format:"uuid"`
	// Chat participants
	Handles []ChatSendVoicememoResponseVoiceMemoChatHandle `json:"handles,required"`
	// Whether the chat is active
	IsActive bool `json:"is_active,required"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group,required"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service string `json:"service,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Handles     respjson.Field
		IsActive    respjson.Field
		IsGroup     respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatSendVoicememoResponseVoiceMemoChat) RawJSON() string { return r.JSON.raw }
func (r *ChatSendVoicememoResponseVoiceMemoChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatSendVoicememoResponseVoiceMemoChatHandle struct {
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
func (r ChatSendVoicememoResponseVoiceMemoChatHandle) RawJSON() string { return r.JSON.raw }
func (r *ChatSendVoicememoResponseVoiceMemoChatHandle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatSendVoicememoResponseVoiceMemoVoiceMemo struct {
	// Attachment identifier
	ID string `json:"id,required" format:"uuid"`
	// Original filename
	Filename string `json:"filename,required"`
	// Audio MIME type
	MimeType string `json:"mime_type,required"`
	// File size in bytes
	SizeBytes int64 `json:"size_bytes,required"`
	// CDN URL for downloading the voice memo
	URL string `json:"url,required" format:"uri"`
	// Duration in milliseconds
	DurationMs int64 `json:"duration_ms,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Filename    respjson.Field
		MimeType    respjson.Field
		SizeBytes   respjson.Field
		URL         respjson.Field
		DurationMs  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatSendVoicememoResponseVoiceMemoVoiceMemo) RawJSON() string { return r.JSON.raw }
func (r *ChatSendVoicememoResponseVoiceMemoVoiceMemo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatNewParams struct {
	// Sender phone number in E.164 format. Must be a phone number that the
	// authenticated partner has permission to send from.
	From string `json:"from,required"`
	// Message content container. Groups all message-related fields together,
	// separating the "what" (message content) from the "where" (routing fields like
	// from/to).
	Message ChatNewParamsMessage `json:"message,omitzero,required"`
	// Array of recipient handles (phone numbers in E.164 format or email addresses).
	// For individual chats, provide one recipient. For group chats, provide multiple.
	To []string `json:"to,omitzero,required"`
	paramObj
}

func (r ChatNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message content container. Groups all message-related fields together,
// separating the "what" (message content) from the "where" (routing fields like
// from/to).
//
// The property Parts is required.
type ChatNewParamsMessage struct {
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
	Parts []ChatNewParamsMessagePartUnion `json:"parts,omitzero,required"`
	// Optional idempotency key for this message. Use this to prevent duplicate sends
	// of the same message.
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// iMessage effect to apply to this message (screen or bubble effect)
	Effect ChatNewParamsMessageEffect `json:"effect,omitzero"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	PreferredService string `json:"preferred_service,omitzero"`
	// Reply to another message to create a threaded conversation
	ReplyTo ChatNewParamsMessageReplyTo `json:"reply_to,omitzero"`
	paramObj
}

func (r ChatNewParamsMessage) MarshalJSON() (data []byte, err error) {
	type shadow ChatNewParamsMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatNewParamsMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatNewParamsMessage](
		"preferred_service", "iMessage", "SMS", "RCS",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatNewParamsMessagePartUnion struct {
	OfText  *ChatNewParamsMessagePartText  `json:",omitzero,inline"`
	OfMedia *ChatNewParamsMessagePartMedia `json:",omitzero,inline"`
	paramUnion
}

func (u ChatNewParamsMessagePartUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfMedia)
}
func (u *ChatNewParamsMessagePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ChatNewParamsMessagePartUnion](
		"type",
		apijson.Discriminator[ChatNewParamsMessagePartText]("text"),
		apijson.Discriminator[ChatNewParamsMessagePartMedia]("media"),
	)
}

// The properties Type, Value are required.
type ChatNewParamsMessagePartText struct {
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

func (r ChatNewParamsMessagePartText) MarshalJSON() (data []byte, err error) {
	type shadow ChatNewParamsMessagePartText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatNewParamsMessagePartText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ChatNewParamsMessagePartMedia struct {
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

func (r ChatNewParamsMessagePartMedia) MarshalJSON() (data []byte, err error) {
	type shadow ChatNewParamsMessagePartMedia
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatNewParamsMessagePartMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// iMessage effect to apply to this message (screen or bubble effect)
type ChatNewParamsMessageEffect struct {
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

func (r ChatNewParamsMessageEffect) MarshalJSON() (data []byte, err error) {
	type shadow ChatNewParamsMessageEffect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatNewParamsMessageEffect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatNewParamsMessageEffect](
		"type", "screen", "bubble",
	)
}

// Reply to another message to create a threaded conversation
//
// The property MessageID is required.
type ChatNewParamsMessageReplyTo struct {
	// The ID of the message to reply to
	MessageID string `json:"message_id,required" format:"uuid"`
	// The specific message part to reply to (0-based index). Defaults to 0 (first
	// part) if not provided. Use this when replying to a specific part of a multipart
	// message.
	PartIndex param.Opt[int64] `json:"part_index,omitzero"`
	paramObj
}

func (r ChatNewParamsMessageReplyTo) MarshalJSON() (data []byte, err error) {
	type shadow ChatNewParamsMessageReplyTo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatNewParamsMessageReplyTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatUpdateParams struct {
	// New display name for the chat (group chats only)
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// URL of an image to set as the group chat icon (group chats only)
	GroupChatIcon param.Opt[string] `json:"group_chat_icon,omitzero" format:"uri"`
	paramObj
}

func (r ChatUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatListParams struct {
	// Phone number to filter chats by. Returns all chats made from this phone number.
	// Must be in E.164 format (e.g., `+13343284472`). The `+` is automatically
	// URL-encoded by HTTP clients.
	From string `query:"from,required" json:"-"`
	// Pagination cursor from the previous response's `next_cursor` field. Omit this
	// parameter for the first page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of chats to return per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChatListParams]'s query parameters as `url.Values`.
func (r ChatListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChatSendVoicememoParams struct {
	// Sender phone number in E.164 format
	From string `json:"from,required"`
	// URL of the voice memo audio file. Must be a publicly accessible HTTPS URL.
	VoiceMemoURL string `json:"voice_memo_url,required" format:"uri"`
	paramObj
}

func (r ChatSendVoicememoParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatSendVoicememoParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatSendVoicememoParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
