// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/linq-team/linq-go"
	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// A media attachment part
type MediaPartResponse struct {
	// Unique attachment identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Original filename
	Filename string `json:"filename" api:"required"`
	// MIME type of the file
	MimeType string `json:"mime_type" api:"required"`
	// Reactions on this message part
	Reactions []linqgo.Reaction `json:"reactions" api:"required"`
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

// Messaging service type
type ServiceType string

const (
	ServiceTypeiMessage ServiceType = "iMessage"
	ServiceTypeSMS      ServiceType = "SMS"
	ServiceTypeRCS      ServiceType = "RCS"
)

// A text message part
type TextPartResponse struct {
	// Reactions on this message part
	Reactions []linqgo.Reaction `json:"reactions" api:"required"`
	// Indicates this is a text message part
	//
	// Any of "text".
	Type TextPartResponseType `json:"type" api:"required"`
	// The text content
	Value string `json:"value" api:"required"`
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
func (r TextPartResponse) RawJSON() string { return r.JSON.raw }
func (r *TextPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates this is a text message part
type TextPartResponseType string

const (
	TextPartResponseTypeText TextPartResponseType = "text"
)
