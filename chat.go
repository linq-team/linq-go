// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
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
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
	"github.com/linq-team/linq-go/shared"
	"github.com/linq-team/linq-go/shared/constant"
)

// ChatService contains methods and other services that help with interacting with
// the linq-api-v3 API.
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
func (r *ChatService) Get(ctx context.Context, chatID string, opts ...option.RequestOption) (res *Chat, err error) {
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
func (r *ChatService) Update(ctx context.Context, chatID string, body ChatUpdateParams, opts ...option.RequestOption) (res *Chat, err error) {
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

type Chat struct {
	// Unique identifier for the chat
	ID string `json:"id" api:"required" format:"uuid"`
	// When the chat was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name for the chat. Defaults to a comma-separated list of recipient
	// handles. Can be updated for group chats.
	DisplayName string `json:"display_name" api:"required"`
	// List of chat participants with full handle details. Always contains at least two
	// handles (your phone number and the other participant).
	Handles []ChatHandle `json:"handles" api:"required"`
	// Whether the chat is archived
	IsArchived bool `json:"is_archived" api:"required"`
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
		IsArchived  respjson.Field
		IsGroup     respjson.Field
		UpdatedAt   respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Chat) RawJSON() string { return r.JSON.raw }
func (r *Chat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message content container. Groups all message-related fields together,
// separating the "what" (message content) from the "where" (routing fields like
// from/to).
//
// The property Parts is required.
type MessageContentParam struct {
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
	Parts []MessageContentPartUnionParam `json:"parts,omitzero" api:"required"`
	// Optional idempotency key for this message. Use this to prevent duplicate sends
	// of the same message.
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// iMessage effect to apply to this message (screen or bubble effect)
	Effect MessageEffectParam `json:"effect,omitzero"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	PreferredService shared.ServiceType `json:"preferred_service,omitzero"`
	// Reply to another message to create a threaded conversation
	ReplyTo ReplyToParam `json:"reply_to,omitzero"`
	paramObj
}

func (r MessageContentParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageContentPartUnionParam struct {
	OfText  *MessageContentPartTextParam  `json:",omitzero,inline"`
	OfMedia *MessageContentPartMediaParam `json:",omitzero,inline"`
	paramUnion
}

func (u MessageContentPartUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfMedia)
}
func (u *MessageContentPartUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[MessageContentPartUnionParam](
		"type",
		apijson.Discriminator[MessageContentPartTextParam]("text"),
		apijson.Discriminator[MessageContentPartMediaParam]("media"),
	)
}

// The properties Type, Value are required.
type MessageContentPartTextParam struct {
	// The text content
	Value string `json:"value" api:"required"`
	// Indicates this is a text message part
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type" api:"required"`
	paramObj
}

func (r MessageContentPartTextParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContentPartTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContentPartTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type MessageContentPartMediaParam struct {
	// Reference to a file pre-uploaded via `POST /v3/attachments` (optional). The file
	// is already stored, so sends using this ID skip the download step — useful when
	// sending the same file to many recipients.
	//
	// Either `url` or `attachment_id` must be provided, but not both.
	AttachmentID param.Opt[string] `json:"attachment_id,omitzero" format:"uuid"`
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
	Type constant.Media `json:"type" api:"required"`
	paramObj
}

func (r MessageContentPartMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContentPartMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContentPartMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for creating a new chat with an initial message
type ChatNewResponse struct {
	Chat ChatNewResponseChat `json:"chat" api:"required"`
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
	ID string `json:"id" api:"required" format:"uuid"`
	// Display name for the chat. Defaults to a comma-separated list of recipient
	// handles. Can be updated for group chats.
	DisplayName string `json:"display_name" api:"required"`
	// List of participants in the chat. Always contains at least two handles (your
	// phone number and the other participant).
	Handles []ChatHandle `json:"handles" api:"required"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group" api:"required"`
	// A message that was sent (used in CreateChat and SendMessage responses)
	Message SentMessage `json:"message" api:"required"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service" api:"required"`
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

type ChatListResponse struct {
	// List of chats
	Chats []Chat `json:"chats" api:"required"`
	// Cursor for fetching the next page of results. Null if there are no more results
	// to fetch. Pass this value as the `cursor` parameter in the next request.
	NextCursor string `json:"next_cursor" api:"nullable"`
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

// Response for sending a voice memo to a chat
type ChatSendVoicememoResponse struct {
	VoiceMemo ChatSendVoicememoResponseVoiceMemo `json:"voice_memo" api:"required"`
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
	ID   string                                 `json:"id" api:"required" format:"uuid"`
	Chat ChatSendVoicememoResponseVoiceMemoChat `json:"chat" api:"required"`
	// When the voice memo was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Sender phone number
	From string `json:"from" api:"required"`
	// Current delivery status
	Status string `json:"status" api:"required"`
	// Recipient handles (phone numbers or email addresses)
	To        []string                                    `json:"to" api:"required"`
	VoiceMemo ChatSendVoicememoResponseVoiceMemoVoiceMemo `json:"voice_memo" api:"required"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service" api:"nullable"`
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
	ID string `json:"id" api:"required" format:"uuid"`
	// Chat participants
	Handles []ChatHandle `json:"handles" api:"required"`
	// Whether the chat is active
	IsActive bool `json:"is_active" api:"required"`
	// Whether this is a group chat
	IsGroup bool `json:"is_group" api:"required"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service shared.ServiceType `json:"service" api:"required"`
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

type ChatSendVoicememoResponseVoiceMemoVoiceMemo struct {
	// Attachment identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Original filename
	Filename string `json:"filename" api:"required"`
	// Audio MIME type
	MimeType string `json:"mime_type" api:"required"`
	// File size in bytes
	SizeBytes int64 `json:"size_bytes" api:"required"`
	// CDN URL for downloading the voice memo
	URL string `json:"url" api:"required" format:"uri"`
	// Duration in milliseconds
	DurationMs int64 `json:"duration_ms" api:"nullable"`
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
	From string `json:"from" api:"required"`
	// Message content container. Groups all message-related fields together,
	// separating the "what" (message content) from the "where" (routing fields like
	// from/to).
	Message MessageContentParam `json:"message,omitzero" api:"required"`
	// Array of recipient handles (phone numbers in E.164 format or email addresses).
	// For individual chats, provide one recipient. For group chats, provide multiple.
	To []string `json:"to,omitzero" api:"required"`
	paramObj
}

func (r ChatNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatNewParams) UnmarshalJSON(data []byte) error {
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
	From string `query:"from" api:"required" json:"-"`
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
	From string `json:"from" api:"required"`
	// URL of the voice memo audio file. Must be a publicly accessible HTTPS URL.
	VoiceMemoURL string `json:"voice_memo_url" api:"required" format:"uri"`
	paramObj
}

func (r ChatSendVoicememoParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatSendVoicememoParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatSendVoicememoParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
