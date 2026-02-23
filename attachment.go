// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqapiv3

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/linq-api-v3-go/internal/apijson"
	"github.com/stainless-sdks/linq-api-v3-go/internal/requestconfig"
	"github.com/stainless-sdks/linq-api-v3-go/option"
	"github.com/stainless-sdks/linq-api-v3-go/packages/param"
	"github.com/stainless-sdks/linq-api-v3-go/packages/respjson"
)

// AttachmentService contains methods and other services that help with interacting
// with the linq API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttachmentService] method instead.
type AttachmentService struct {
	Options []option.RequestOption
}

// NewAttachmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttachmentService(opts ...option.RequestOption) (r AttachmentService) {
	r = AttachmentService{}
	r.Options = opts
	return
}

// **This endpoint is optional.** You can send media by simply providing a URL in
// your message's media part — no pre-upload required. Use this endpoint only when
// you want to upload a file ahead of time for reuse or latency optimization.
//
// Returns a presigned upload URL and a permanent `attachment_id` you can reference
// in future messages.
//
// ## Step 1: Request an upload URL
//
// Call this endpoint with file metadata:
//
// ```json
// POST /v3/attachments
//
//	{
//	  "filename": "photo.jpg",
//	  "content_type": "image/jpeg",
//	  "size_bytes": 1024000
//	}
//
// ```
//
// The response includes an `upload_url` (valid for 15 minutes) and a permanent
// `attachment_id`.
//
// ## Step 2: Upload the file
//
// Make a PUT request to the `upload_url` with the raw file bytes as the request
// body.. Include the headers from `required_headers`. The request body is the
// binary file content — **not** JSON, **not** multipart form data.
//
// ```bash
//
//	curl -X PUT "<upload_url from step 1>" \
//	  -H "Content-Type: image/jpeg" \
//	  --data-binary @filebytes
//
// ```
//
// ## Step 3: Send a message with the attachment
//
// Reference the `attachment_id` in a media part. The ID never expires — use it in
// as many messages as you want.
//
// ```json
// POST /v3/messages
//
//	{
//	  "to": ["+15551234567"],
//	  "from": "+15559876543",
//	  "parts": [
//	    { "type": "media", "attachment_id": "<attachment_id from step 1>" }
//	  ]
//	}
//
// ```
//
// ## When to use this instead of a URL in the media part
//
// - Sending the same file to multiple recipients (avoids re-downloading each time)
// - Large files where you want to separate upload from message send
// - Latency-sensitive sends where the file should already be stored
//
// If you just need to send a file once, skip all of this and pass a `url` directly
// in the media part instead.
//
// **File Size Limit:** 100MB
//
// **Unsupported Types:** WebP, SVG, FLAC, OGG, and executable files are explicitly
// rejected.
func (r *AttachmentService) New(ctx context.Context, body AttachmentNewParams, opts ...option.RequestOption) (res *AttachmentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/attachments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve metadata for a specific attachment including its status, file
// information, and URLs for downloading.
func (r *AttachmentService) Get(ctx context.Context, attachmentID string, opts ...option.RequestOption) (res *AttachmentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if attachmentID == "" {
		err = errors.New("missing required attachmentId parameter")
		return
	}
	path := fmt.Sprintf("v3/attachments/%s", attachmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AttachmentNewResponse struct {
	// Unique identifier for the attachment (for status checks via GET
	// /v3/attachments/{id})
	AttachmentID string `json:"attachment_id,required" format:"uuid"`
	// Permanent CDN URL for the file. Does not expire. Use the `attachment_id` to
	// reference this file in media parts when sending messages.
	DownloadURL string `json:"download_url,required" format:"uri"`
	// When the upload URL expires (15 minutes from now)
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// HTTP method to use for upload (always PUT)
	//
	// Any of "PUT".
	HTTPMethod AttachmentNewResponseHTTPMethod `json:"http_method,required"`
	// HTTP headers required for the upload request
	RequiredHeaders map[string]string `json:"required_headers,required"`
	// Presigned URL for uploading the file. PUT the raw binary file content to this
	// URL with the `required_headers`. Do not JSON-encode or multipart-wrap the body.
	// Expires after 15 minutes.
	UploadURL string `json:"upload_url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachmentID    respjson.Field
		DownloadURL     respjson.Field
		ExpiresAt       respjson.Field
		HTTPMethod      respjson.Field
		RequiredHeaders respjson.Field
		UploadURL       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttachmentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AttachmentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP method to use for upload (always PUT)
type AttachmentNewResponseHTTPMethod string

const (
	AttachmentNewResponseHTTPMethodPut AttachmentNewResponseHTTPMethod = "PUT"
)

type AttachmentGetResponse struct {
	// Unique identifier for the attachment (UUID)
	ID string `json:"id,required"`
	// Supported MIME types for file attachments and media URLs.
	//
	// **Images:** image/jpeg, image/png, image/gif, image/heic, image/heif,
	// image/tiff, image/bmp
	//
	// **Videos:** video/mp4, video/quicktime, video/mpeg, video/3gpp
	//
	// **Audio:** audio/mpeg, audio/mp4, audio/x-m4a, audio/x-caf, audio/wav,
	// audio/aiff, audio/aac, audio/amr
	//
	// **Documents:** application/pdf, text/plain, text/markdown, text/vcard, text/rtf,
	// text/csv, text/html, text/calendar, application/msword,
	// application/vnd.openxmlformats-officedocument.wordprocessingml.document,
	// application/vnd.ms-excel,
	// application/vnd.openxmlformats-officedocument.spreadsheetml.sheet,
	// application/vnd.ms-powerpoint,
	// application/vnd.openxmlformats-officedocument.presentationml.presentation,
	// application/vnd.apple.pages, application/vnd.apple.numbers,
	// application/vnd.apple.keynote, application/epub+zip, application/zip
	//
	// **Unsupported:** WebP, SVG, FLAC, OGG, and executable files are explicitly
	// rejected.
	//
	// Any of "image/jpeg", "image/jpg", "image/png", "image/gif", "image/heic",
	// "image/heif", "image/tiff", "image/bmp", "image/x-ms-bmp", "video/mp4",
	// "video/quicktime", "video/mpeg", "video/x-m4v", "video/3gpp", "audio/mpeg",
	// "audio/mp3", "audio/mp4", "audio/x-m4a", "audio/m4a", "audio/x-caf",
	// "audio/wav", "audio/x-wav", "audio/aiff", "audio/x-aiff", "audio/aac",
	// "audio/x-aac", "audio/amr", "application/pdf", "text/plain", "text/markdown",
	// "text/vcard", "text/x-vcard", "text/rtf", "application/rtf", "text/csv",
	// "text/html", "text/calendar", "application/msword",
	// "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	// "application/vnd.ms-excel",
	// "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	// "application/vnd.ms-powerpoint",
	// "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	// "application/vnd.apple.pages", "application/x-iwork-pages-sffpages",
	// "application/vnd.apple.numbers", "application/x-iwork-numbers-sffnumbers",
	// "application/vnd.apple.keynote", "application/x-iwork-keynote-sffkey",
	// "application/epub+zip", "application/zip", "application/x-zip-compressed".
	ContentType AttachmentGetResponseContentType `json:"content_type,required"`
	// When the attachment was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Original filename of the attachment
	Filename string `json:"filename,required"`
	// Size of the attachment in bytes
	SizeBytes int64 `json:"size_bytes,required"`
	// Current upload/processing status
	//
	// Any of "pending", "complete", "failed".
	Status AttachmentGetResponseStatus `json:"status,required"`
	// URL to download the attachment
	DownloadURL string `json:"download_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ContentType respjson.Field
		CreatedAt   respjson.Field
		Filename    respjson.Field
		SizeBytes   respjson.Field
		Status      respjson.Field
		DownloadURL respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttachmentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AttachmentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Supported MIME types for file attachments and media URLs.
//
// **Images:** image/jpeg, image/png, image/gif, image/heic, image/heif,
// image/tiff, image/bmp
//
// **Videos:** video/mp4, video/quicktime, video/mpeg, video/3gpp
//
// **Audio:** audio/mpeg, audio/mp4, audio/x-m4a, audio/x-caf, audio/wav,
// audio/aiff, audio/aac, audio/amr
//
// **Documents:** application/pdf, text/plain, text/markdown, text/vcard, text/rtf,
// text/csv, text/html, text/calendar, application/msword,
// application/vnd.openxmlformats-officedocument.wordprocessingml.document,
// application/vnd.ms-excel,
// application/vnd.openxmlformats-officedocument.spreadsheetml.sheet,
// application/vnd.ms-powerpoint,
// application/vnd.openxmlformats-officedocument.presentationml.presentation,
// application/vnd.apple.pages, application/vnd.apple.numbers,
// application/vnd.apple.keynote, application/epub+zip, application/zip
//
// **Unsupported:** WebP, SVG, FLAC, OGG, and executable files are explicitly
// rejected.
type AttachmentGetResponseContentType string

const (
	AttachmentGetResponseContentTypeImageJpeg                                                            AttachmentGetResponseContentType = "image/jpeg"
	AttachmentGetResponseContentTypeImageJpg                                                             AttachmentGetResponseContentType = "image/jpg"
	AttachmentGetResponseContentTypeImagePng                                                             AttachmentGetResponseContentType = "image/png"
	AttachmentGetResponseContentTypeImageGif                                                             AttachmentGetResponseContentType = "image/gif"
	AttachmentGetResponseContentTypeImageHeic                                                            AttachmentGetResponseContentType = "image/heic"
	AttachmentGetResponseContentTypeImageHeif                                                            AttachmentGetResponseContentType = "image/heif"
	AttachmentGetResponseContentTypeImageTiff                                                            AttachmentGetResponseContentType = "image/tiff"
	AttachmentGetResponseContentTypeImageBmp                                                             AttachmentGetResponseContentType = "image/bmp"
	AttachmentGetResponseContentTypeImageXMsBmp                                                          AttachmentGetResponseContentType = "image/x-ms-bmp"
	AttachmentGetResponseContentTypeVideoMP4                                                             AttachmentGetResponseContentType = "video/mp4"
	AttachmentGetResponseContentTypeVideoQuicktime                                                       AttachmentGetResponseContentType = "video/quicktime"
	AttachmentGetResponseContentTypeVideoMpeg                                                            AttachmentGetResponseContentType = "video/mpeg"
	AttachmentGetResponseContentTypeVideoXM4v                                                            AttachmentGetResponseContentType = "video/x-m4v"
	AttachmentGetResponseContentTypeVideo3gpp                                                            AttachmentGetResponseContentType = "video/3gpp"
	AttachmentGetResponseContentTypeAudioMpeg                                                            AttachmentGetResponseContentType = "audio/mpeg"
	AttachmentGetResponseContentTypeAudioMP3                                                             AttachmentGetResponseContentType = "audio/mp3"
	AttachmentGetResponseContentTypeAudioMP4                                                             AttachmentGetResponseContentType = "audio/mp4"
	AttachmentGetResponseContentTypeAudioXM4a                                                            AttachmentGetResponseContentType = "audio/x-m4a"
	AttachmentGetResponseContentTypeAudioM4a                                                             AttachmentGetResponseContentType = "audio/m4a"
	AttachmentGetResponseContentTypeAudioXCaf                                                            AttachmentGetResponseContentType = "audio/x-caf"
	AttachmentGetResponseContentTypeAudioWav                                                             AttachmentGetResponseContentType = "audio/wav"
	AttachmentGetResponseContentTypeAudioXWav                                                            AttachmentGetResponseContentType = "audio/x-wav"
	AttachmentGetResponseContentTypeAudioAiff                                                            AttachmentGetResponseContentType = "audio/aiff"
	AttachmentGetResponseContentTypeAudioXAiff                                                           AttachmentGetResponseContentType = "audio/x-aiff"
	AttachmentGetResponseContentTypeAudioAac                                                             AttachmentGetResponseContentType = "audio/aac"
	AttachmentGetResponseContentTypeAudioXAac                                                            AttachmentGetResponseContentType = "audio/x-aac"
	AttachmentGetResponseContentTypeAudioAmr                                                             AttachmentGetResponseContentType = "audio/amr"
	AttachmentGetResponseContentTypeApplicationPdf                                                       AttachmentGetResponseContentType = "application/pdf"
	AttachmentGetResponseContentTypeTextPlain                                                            AttachmentGetResponseContentType = "text/plain"
	AttachmentGetResponseContentTypeTextMarkdown                                                         AttachmentGetResponseContentType = "text/markdown"
	AttachmentGetResponseContentTypeTextVcard                                                            AttachmentGetResponseContentType = "text/vcard"
	AttachmentGetResponseContentTypeTextXVcard                                                           AttachmentGetResponseContentType = "text/x-vcard"
	AttachmentGetResponseContentTypeTextRtf                                                              AttachmentGetResponseContentType = "text/rtf"
	AttachmentGetResponseContentTypeApplicationRtf                                                       AttachmentGetResponseContentType = "application/rtf"
	AttachmentGetResponseContentTypeTextCsv                                                              AttachmentGetResponseContentType = "text/csv"
	AttachmentGetResponseContentTypeTextHTML                                                             AttachmentGetResponseContentType = "text/html"
	AttachmentGetResponseContentTypeTextCalendar                                                         AttachmentGetResponseContentType = "text/calendar"
	AttachmentGetResponseContentTypeApplicationMsword                                                    AttachmentGetResponseContentType = "application/msword"
	AttachmentGetResponseContentTypeApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocument   AttachmentGetResponseContentType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	AttachmentGetResponseContentTypeApplicationVndMsExcel                                                AttachmentGetResponseContentType = "application/vnd.ms-excel"
	AttachmentGetResponseContentTypeApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheet         AttachmentGetResponseContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	AttachmentGetResponseContentTypeApplicationVndMsPowerpoint                                           AttachmentGetResponseContentType = "application/vnd.ms-powerpoint"
	AttachmentGetResponseContentTypeApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresentation AttachmentGetResponseContentType = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	AttachmentGetResponseContentTypeApplicationVndApplePages                                             AttachmentGetResponseContentType = "application/vnd.apple.pages"
	AttachmentGetResponseContentTypeApplicationXIworkPagesSffpages                                       AttachmentGetResponseContentType = "application/x-iwork-pages-sffpages"
	AttachmentGetResponseContentTypeApplicationVndAppleNumbers                                           AttachmentGetResponseContentType = "application/vnd.apple.numbers"
	AttachmentGetResponseContentTypeApplicationXIworkNumbersSffnumbers                                   AttachmentGetResponseContentType = "application/x-iwork-numbers-sffnumbers"
	AttachmentGetResponseContentTypeApplicationVndAppleKeynote                                           AttachmentGetResponseContentType = "application/vnd.apple.keynote"
	AttachmentGetResponseContentTypeApplicationXIworkKeynoteSffkey                                       AttachmentGetResponseContentType = "application/x-iwork-keynote-sffkey"
	AttachmentGetResponseContentTypeApplicationEpubZip                                                   AttachmentGetResponseContentType = "application/epub+zip"
	AttachmentGetResponseContentTypeApplicationZip                                                       AttachmentGetResponseContentType = "application/zip"
	AttachmentGetResponseContentTypeApplicationXZipCompressed                                            AttachmentGetResponseContentType = "application/x-zip-compressed"
)

// Current upload/processing status
type AttachmentGetResponseStatus string

const (
	AttachmentGetResponseStatusPending  AttachmentGetResponseStatus = "pending"
	AttachmentGetResponseStatusComplete AttachmentGetResponseStatus = "complete"
	AttachmentGetResponseStatusFailed   AttachmentGetResponseStatus = "failed"
)

type AttachmentNewParams struct {
	// Supported MIME types for file attachments and media URLs.
	//
	// **Images:** image/jpeg, image/png, image/gif, image/heic, image/heif,
	// image/tiff, image/bmp
	//
	// **Videos:** video/mp4, video/quicktime, video/mpeg, video/3gpp
	//
	// **Audio:** audio/mpeg, audio/mp4, audio/x-m4a, audio/x-caf, audio/wav,
	// audio/aiff, audio/aac, audio/amr
	//
	// **Documents:** application/pdf, text/plain, text/markdown, text/vcard, text/rtf,
	// text/csv, text/html, text/calendar, application/msword,
	// application/vnd.openxmlformats-officedocument.wordprocessingml.document,
	// application/vnd.ms-excel,
	// application/vnd.openxmlformats-officedocument.spreadsheetml.sheet,
	// application/vnd.ms-powerpoint,
	// application/vnd.openxmlformats-officedocument.presentationml.presentation,
	// application/vnd.apple.pages, application/vnd.apple.numbers,
	// application/vnd.apple.keynote, application/epub+zip, application/zip
	//
	// **Unsupported:** WebP, SVG, FLAC, OGG, and executable files are explicitly
	// rejected.
	//
	// Any of "image/jpeg", "image/jpg", "image/png", "image/gif", "image/heic",
	// "image/heif", "image/tiff", "image/bmp", "image/x-ms-bmp", "video/mp4",
	// "video/quicktime", "video/mpeg", "video/x-m4v", "video/3gpp", "audio/mpeg",
	// "audio/mp3", "audio/mp4", "audio/x-m4a", "audio/m4a", "audio/x-caf",
	// "audio/wav", "audio/x-wav", "audio/aiff", "audio/x-aiff", "audio/aac",
	// "audio/x-aac", "audio/amr", "application/pdf", "text/plain", "text/markdown",
	// "text/vcard", "text/x-vcard", "text/rtf", "application/rtf", "text/csv",
	// "text/html", "text/calendar", "application/msword",
	// "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	// "application/vnd.ms-excel",
	// "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	// "application/vnd.ms-powerpoint",
	// "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	// "application/vnd.apple.pages", "application/x-iwork-pages-sffpages",
	// "application/vnd.apple.numbers", "application/x-iwork-numbers-sffnumbers",
	// "application/vnd.apple.keynote", "application/x-iwork-keynote-sffkey",
	// "application/epub+zip", "application/zip", "application/x-zip-compressed".
	ContentType AttachmentNewParamsContentType `json:"content_type,omitzero,required"`
	// Name of the file to upload
	Filename string `json:"filename,required"`
	// Size of the file in bytes (max 100MB)
	SizeBytes int64 `json:"size_bytes,required"`
	paramObj
}

func (r AttachmentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AttachmentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttachmentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Supported MIME types for file attachments and media URLs.
//
// **Images:** image/jpeg, image/png, image/gif, image/heic, image/heif,
// image/tiff, image/bmp
//
// **Videos:** video/mp4, video/quicktime, video/mpeg, video/3gpp
//
// **Audio:** audio/mpeg, audio/mp4, audio/x-m4a, audio/x-caf, audio/wav,
// audio/aiff, audio/aac, audio/amr
//
// **Documents:** application/pdf, text/plain, text/markdown, text/vcard, text/rtf,
// text/csv, text/html, text/calendar, application/msword,
// application/vnd.openxmlformats-officedocument.wordprocessingml.document,
// application/vnd.ms-excel,
// application/vnd.openxmlformats-officedocument.spreadsheetml.sheet,
// application/vnd.ms-powerpoint,
// application/vnd.openxmlformats-officedocument.presentationml.presentation,
// application/vnd.apple.pages, application/vnd.apple.numbers,
// application/vnd.apple.keynote, application/epub+zip, application/zip
//
// **Unsupported:** WebP, SVG, FLAC, OGG, and executable files are explicitly
// rejected.
type AttachmentNewParamsContentType string

const (
	AttachmentNewParamsContentTypeImageJpeg                                                            AttachmentNewParamsContentType = "image/jpeg"
	AttachmentNewParamsContentTypeImageJpg                                                             AttachmentNewParamsContentType = "image/jpg"
	AttachmentNewParamsContentTypeImagePng                                                             AttachmentNewParamsContentType = "image/png"
	AttachmentNewParamsContentTypeImageGif                                                             AttachmentNewParamsContentType = "image/gif"
	AttachmentNewParamsContentTypeImageHeic                                                            AttachmentNewParamsContentType = "image/heic"
	AttachmentNewParamsContentTypeImageHeif                                                            AttachmentNewParamsContentType = "image/heif"
	AttachmentNewParamsContentTypeImageTiff                                                            AttachmentNewParamsContentType = "image/tiff"
	AttachmentNewParamsContentTypeImageBmp                                                             AttachmentNewParamsContentType = "image/bmp"
	AttachmentNewParamsContentTypeImageXMsBmp                                                          AttachmentNewParamsContentType = "image/x-ms-bmp"
	AttachmentNewParamsContentTypeVideoMP4                                                             AttachmentNewParamsContentType = "video/mp4"
	AttachmentNewParamsContentTypeVideoQuicktime                                                       AttachmentNewParamsContentType = "video/quicktime"
	AttachmentNewParamsContentTypeVideoMpeg                                                            AttachmentNewParamsContentType = "video/mpeg"
	AttachmentNewParamsContentTypeVideoXM4v                                                            AttachmentNewParamsContentType = "video/x-m4v"
	AttachmentNewParamsContentTypeVideo3gpp                                                            AttachmentNewParamsContentType = "video/3gpp"
	AttachmentNewParamsContentTypeAudioMpeg                                                            AttachmentNewParamsContentType = "audio/mpeg"
	AttachmentNewParamsContentTypeAudioMP3                                                             AttachmentNewParamsContentType = "audio/mp3"
	AttachmentNewParamsContentTypeAudioMP4                                                             AttachmentNewParamsContentType = "audio/mp4"
	AttachmentNewParamsContentTypeAudioXM4a                                                            AttachmentNewParamsContentType = "audio/x-m4a"
	AttachmentNewParamsContentTypeAudioM4a                                                             AttachmentNewParamsContentType = "audio/m4a"
	AttachmentNewParamsContentTypeAudioXCaf                                                            AttachmentNewParamsContentType = "audio/x-caf"
	AttachmentNewParamsContentTypeAudioWav                                                             AttachmentNewParamsContentType = "audio/wav"
	AttachmentNewParamsContentTypeAudioXWav                                                            AttachmentNewParamsContentType = "audio/x-wav"
	AttachmentNewParamsContentTypeAudioAiff                                                            AttachmentNewParamsContentType = "audio/aiff"
	AttachmentNewParamsContentTypeAudioXAiff                                                           AttachmentNewParamsContentType = "audio/x-aiff"
	AttachmentNewParamsContentTypeAudioAac                                                             AttachmentNewParamsContentType = "audio/aac"
	AttachmentNewParamsContentTypeAudioXAac                                                            AttachmentNewParamsContentType = "audio/x-aac"
	AttachmentNewParamsContentTypeAudioAmr                                                             AttachmentNewParamsContentType = "audio/amr"
	AttachmentNewParamsContentTypeApplicationPdf                                                       AttachmentNewParamsContentType = "application/pdf"
	AttachmentNewParamsContentTypeTextPlain                                                            AttachmentNewParamsContentType = "text/plain"
	AttachmentNewParamsContentTypeTextMarkdown                                                         AttachmentNewParamsContentType = "text/markdown"
	AttachmentNewParamsContentTypeTextVcard                                                            AttachmentNewParamsContentType = "text/vcard"
	AttachmentNewParamsContentTypeTextXVcard                                                           AttachmentNewParamsContentType = "text/x-vcard"
	AttachmentNewParamsContentTypeTextRtf                                                              AttachmentNewParamsContentType = "text/rtf"
	AttachmentNewParamsContentTypeApplicationRtf                                                       AttachmentNewParamsContentType = "application/rtf"
	AttachmentNewParamsContentTypeTextCsv                                                              AttachmentNewParamsContentType = "text/csv"
	AttachmentNewParamsContentTypeTextHTML                                                             AttachmentNewParamsContentType = "text/html"
	AttachmentNewParamsContentTypeTextCalendar                                                         AttachmentNewParamsContentType = "text/calendar"
	AttachmentNewParamsContentTypeApplicationMsword                                                    AttachmentNewParamsContentType = "application/msword"
	AttachmentNewParamsContentTypeApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocument   AttachmentNewParamsContentType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	AttachmentNewParamsContentTypeApplicationVndMsExcel                                                AttachmentNewParamsContentType = "application/vnd.ms-excel"
	AttachmentNewParamsContentTypeApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheet         AttachmentNewParamsContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	AttachmentNewParamsContentTypeApplicationVndMsPowerpoint                                           AttachmentNewParamsContentType = "application/vnd.ms-powerpoint"
	AttachmentNewParamsContentTypeApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresentation AttachmentNewParamsContentType = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	AttachmentNewParamsContentTypeApplicationVndApplePages                                             AttachmentNewParamsContentType = "application/vnd.apple.pages"
	AttachmentNewParamsContentTypeApplicationXIworkPagesSffpages                                       AttachmentNewParamsContentType = "application/x-iwork-pages-sffpages"
	AttachmentNewParamsContentTypeApplicationVndAppleNumbers                                           AttachmentNewParamsContentType = "application/vnd.apple.numbers"
	AttachmentNewParamsContentTypeApplicationXIworkNumbersSffnumbers                                   AttachmentNewParamsContentType = "application/x-iwork-numbers-sffnumbers"
	AttachmentNewParamsContentTypeApplicationVndAppleKeynote                                           AttachmentNewParamsContentType = "application/vnd.apple.keynote"
	AttachmentNewParamsContentTypeApplicationXIworkKeynoteSffkey                                       AttachmentNewParamsContentType = "application/x-iwork-keynote-sffkey"
	AttachmentNewParamsContentTypeApplicationEpubZip                                                   AttachmentNewParamsContentType = "application/epub+zip"
	AttachmentNewParamsContentTypeApplicationZip                                                       AttachmentNewParamsContentType = "application/zip"
	AttachmentNewParamsContentTypeApplicationXZipCompressed                                            AttachmentNewParamsContentType = "application/x-zip-compressed"
)
