// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type ChatHandle struct {
	// Unique identifier for this handle
	ID string `json:"id" api:"required" format:"uuid"`
	// Phone number (E.164) or email address of the participant
	Handle string `json:"handle" api:"required"`
	// When this participant joined the chat
	JoinedAt time.Time `json:"joined_at" api:"required" format:"date-time"`
	// Messaging service type
	//
	// Any of "iMessage", "SMS", "RCS".
	Service ServiceType `json:"service" api:"required"`
	// Whether this handle belongs to the sender (your phone number)
	IsMe bool `json:"is_me" api:"nullable"`
	// When they left (if applicable)
	LeftAt time.Time `json:"left_at" api:"nullable" format:"date-time"`
	// Participant status
	//
	// Any of "active", "left", "removed".
	Status ChatHandleStatus `json:"status" api:"nullable"`
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

// Participant status
type ChatHandleStatus string

const (
	ChatHandleStatusActive  ChatHandleStatus = "active"
	ChatHandleStatusLeft    ChatHandleStatus = "left"
	ChatHandleStatusRemoved ChatHandleStatus = "removed"
)

// A media attachment part
type MediaPartResponse struct {
	// Unique attachment identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Original filename
	Filename string `json:"filename" api:"required"`
	// MIME type of the file
	MimeType string `json:"mime_type" api:"required"`
	// Reactions on this message part
	Reactions []Reaction `json:"reactions" api:"required"`
	// File size in bytes
	SizeBytes int64 `json:"size_bytes" api:"required"`
	// Indicates this is a media attachment part
	//
	// Any of "media".
	Type MediaPartResponseType `json:"type" api:"required"`
	// Presigned URL for downloading the attachment (expires in 1 hour).
	URL string `json:"url" api:"required" format:"uri"`
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
func (r MediaPartResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates this is a media attachment part
type MediaPartResponseType string

const (
	MediaPartResponseTypeMedia MediaPartResponseType = "media"
)

type Reaction struct {
	Handle ChatHandle `json:"handle" api:"required"`
	// Whether this reaction is from the current user
	IsMe bool `json:"is_me" api:"required"`
	// Type of reaction. Standard iMessage tapbacks are love, like, dislike, laugh,
	// emphasize, question. Custom emoji reactions have type "custom" with the actual
	// emoji in the custom_emoji field. Sticker reactions have type "sticker" with
	// sticker attachment details in the sticker field.
	//
	// Any of "love", "like", "dislike", "laugh", "emphasize", "question", "custom",
	// "sticker".
	Type ReactionType `json:"type" api:"required"`
	// Custom emoji if type is "custom", null otherwise
	CustomEmoji string `json:"custom_emoji" api:"nullable"`
	// Sticker attachment details when reaction_type is "sticker". Null for non-sticker
	// reactions.
	Sticker ReactionSticker `json:"sticker" api:"nullable"`
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

// Messaging service type
type ServiceType string

const (
	ServiceTypeiMessage ServiceType = "iMessage"
	ServiceTypeSMS      ServiceType = "SMS"
	ServiceTypeRCS      ServiceType = "RCS"
)

type TextDecoration struct {
	// Character range `[start, end)` in the `value` string where the decoration
	// applies. `start` is inclusive, `end` is exclusive. _Characters are measured as
	// UTF-16 code units. Most characters count as 1; some emoji count as 2._
	Range []int64 `json:"range" api:"required"`
	// Animated text effect to apply. Mutually exclusive with `style`.
	//
	// Any of "big", "small", "shake", "nod", "explode", "ripple", "bloom", "jitter".
	Animation TextDecorationAnimation `json:"animation"`
	// Text style to apply. Mutually exclusive with `animation`.
	//
	// Any of "bold", "italic", "strikethrough", "underline".
	Style TextDecorationStyle `json:"style"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Range       respjson.Field
		Animation   respjson.Field
		Style       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextDecoration) RawJSON() string { return r.JSON.raw }
func (r *TextDecoration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TextDecoration to a TextDecorationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TextDecorationParam.Overrides()
func (r TextDecoration) ToParam() TextDecorationParam {
	return param.Override[TextDecorationParam](json.RawMessage(r.RawJSON()))
}

// Animated text effect to apply. Mutually exclusive with `style`.
type TextDecorationAnimation string

const (
	TextDecorationAnimationBig     TextDecorationAnimation = "big"
	TextDecorationAnimationSmall   TextDecorationAnimation = "small"
	TextDecorationAnimationShake   TextDecorationAnimation = "shake"
	TextDecorationAnimationNod     TextDecorationAnimation = "nod"
	TextDecorationAnimationExplode TextDecorationAnimation = "explode"
	TextDecorationAnimationRipple  TextDecorationAnimation = "ripple"
	TextDecorationAnimationBloom   TextDecorationAnimation = "bloom"
	TextDecorationAnimationJitter  TextDecorationAnimation = "jitter"
)

// Text style to apply. Mutually exclusive with `animation`.
type TextDecorationStyle string

const (
	TextDecorationStyleBold          TextDecorationStyle = "bold"
	TextDecorationStyleItalic        TextDecorationStyle = "italic"
	TextDecorationStyleStrikethrough TextDecorationStyle = "strikethrough"
	TextDecorationStyleUnderline     TextDecorationStyle = "underline"
)

// The property Range is required.
type TextDecorationParam struct {
	// Character range `[start, end)` in the `value` string where the decoration
	// applies. `start` is inclusive, `end` is exclusive. _Characters are measured as
	// UTF-16 code units. Most characters count as 1; some emoji count as 2._
	Range []int64 `json:"range,omitzero" api:"required"`
	// Animated text effect to apply. Mutually exclusive with `style`.
	//
	// Any of "big", "small", "shake", "nod", "explode", "ripple", "bloom", "jitter".
	Animation TextDecorationAnimation `json:"animation,omitzero"`
	// Text style to apply. Mutually exclusive with `animation`.
	//
	// Any of "bold", "italic", "strikethrough", "underline".
	Style TextDecorationStyle `json:"style,omitzero"`
	paramObj
}

func (r TextDecorationParam) MarshalJSON() (data []byte, err error) {
	type shadow TextDecorationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextDecorationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text message part
type TextPartResponse struct {
	// Reactions on this message part
	Reactions []Reaction `json:"reactions" api:"required"`
	// Indicates this is a text message part
	//
	// Any of "text".
	Type TextPartResponseType `json:"type" api:"required"`
	// The text content
	Value string `json:"value" api:"required"`
	// Text decorations applied to character ranges in the value
	TextDecorations []TextDecoration `json:"text_decorations" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reactions       respjson.Field
		Type            respjson.Field
		Value           respjson.Field
		TextDecorations respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextPartResponse) RawJSON() string { return r.JSON.raw }
func (r *TextPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates this is a text message part
type TextPartResponseType string

const (
	TextPartResponseTypeText TextPartResponseType = "text"
)
