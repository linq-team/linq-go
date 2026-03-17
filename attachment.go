// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// Send files (images, videos, documents, audio) with messages by providing a URL
// in a media part. Pre-uploading via `POST /v3/attachments` is **optional** and
// only needed for specific optimization scenarios.
//
// ## Sending Media via URL (up to 10MB)
//
// Provide a publicly accessible HTTPS URL with a
// [supported media type](#supported-file-types) in the `url` field of a media
// part.
//
// ```json
//
//	{
//	  "parts": [{ "type": "media", "url": "https://your-cdn.com/images/photo.jpg" }]
//	}
//
// ```
//
// This works with any URL you already host — no pre-upload step required.
// **Maximum file size: 10MB.**
//
// ## Pre-Upload (required for files over 10MB)
//
// Use `POST /v3/attachments` when you want to:
//
//   - **Send files larger than 10MB** (up to 100MB) — URL-based downloads are
//     limited to 10MB
//   - **Send the same file to many recipients** — upload once, reuse the
//     `attachment_id` without re-downloading each time
//   - **Reduce message send latency** — the file is already stored, so sending is
//     faster
//
// **How it works:**
//
//  1. `POST /v3/attachments` with file metadata → returns a presigned `upload_url`
//     (valid for **15 minutes**) and a permanent `attachment_id`
//  2. PUT the raw file bytes to the `upload_url` with the `required_headers` (no
//     JSON or multipart — just the binary content)
//  3. Reference the `attachment_id` in your media part when sending messages (no
//     expiration)
//
// **Key difference:** When you provide an external `url`, we download and process
// the file on every send. When you use a pre-uploaded `attachment_id`, the file is
// already stored — so repeated sends skip the download step entirely.
//
// ## Domain Allowlisting
//
// Attachment URLs in API responses are served from `cdn.linqapp.com`. This
// includes:
//
// - `url` fields in media and voice memo message parts
// - `download_url` fields in attachment and upload response objects
//
// If your application enforces domain allowlists (e.g., for SSRF protection), add:
//
// ```
// cdn.linqapp.com
// ```
//
// ## Supported File Types
//
// - **Images:** JPEG, PNG, GIF, HEIC, HEIF, TIFF, BMP
// - **Videos:** MP4, MOV, M4V
// - **Audio:** M4A, AAC, MP3, WAV, AIFF, CAF, AMR
// - **Documents:** PDF, TXT, RTF, CSV, Office formats, ZIP
// - **Contact & Calendar:** VCF, ICS
//
// ## Audio: Attachment vs Voice Memo
//
// Audio files sent as media parts appear as **downloadable file attachments** in
// iMessage. To send audio as an **iMessage voice memo bubble** (with native inline
// playback UI), use the dedicated `POST /v3/chats/{chatId}/voicememo` endpoint
// instead.
//
// ## File Size Limits
//
// - **URL-based (`url` field):** 10MB maximum
// - **Pre-upload (`attachment_id`):** 100MB maximum
//
// AttachmentService contains methods and other services that help with interacting
// with the linq-api-v3 API.
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
// POST /v3/chats
//
//	{
//	  "from": "+15559876543",
//	  "to": ["+15551234567"],
//	  "message": {
//	    "parts": [
//	      { "type": "media", "attachment_id": "<attachment_id from step 1>" }
//	    ]
//	  }
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
	return res, err
}

// Retrieve metadata for a specific attachment including its status, file
// information, and URLs for downloading.
func (r *AttachmentService) Get(ctx context.Context, attachmentID string, opts ...option.RequestOption) (res *AttachmentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if attachmentID == "" {
		err = errors.New("missing required attachmentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/attachments/%s", attachmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
type SupportedContentType string

const (
	SupportedContentTypeImageJpeg                                                            SupportedContentType = "image/jpeg"
	SupportedContentTypeImageJpg                                                             SupportedContentType = "image/jpg"
	SupportedContentTypeImagePng                                                             SupportedContentType = "image/png"
	SupportedContentTypeImageGif                                                             SupportedContentType = "image/gif"
	SupportedContentTypeImageHeic                                                            SupportedContentType = "image/heic"
	SupportedContentTypeImageHeif                                                            SupportedContentType = "image/heif"
	SupportedContentTypeImageTiff                                                            SupportedContentType = "image/tiff"
	SupportedContentTypeImageBmp                                                             SupportedContentType = "image/bmp"
	SupportedContentTypeImageXMsBmp                                                          SupportedContentType = "image/x-ms-bmp"
	SupportedContentTypeVideoMP4                                                             SupportedContentType = "video/mp4"
	SupportedContentTypeVideoQuicktime                                                       SupportedContentType = "video/quicktime"
	SupportedContentTypeVideoMpeg                                                            SupportedContentType = "video/mpeg"
	SupportedContentTypeVideoXM4v                                                            SupportedContentType = "video/x-m4v"
	SupportedContentTypeVideo3gpp                                                            SupportedContentType = "video/3gpp"
	SupportedContentTypeAudioMpeg                                                            SupportedContentType = "audio/mpeg"
	SupportedContentTypeAudioMP3                                                             SupportedContentType = "audio/mp3"
	SupportedContentTypeAudioMP4                                                             SupportedContentType = "audio/mp4"
	SupportedContentTypeAudioXM4a                                                            SupportedContentType = "audio/x-m4a"
	SupportedContentTypeAudioM4a                                                             SupportedContentType = "audio/m4a"
	SupportedContentTypeAudioXCaf                                                            SupportedContentType = "audio/x-caf"
	SupportedContentTypeAudioWav                                                             SupportedContentType = "audio/wav"
	SupportedContentTypeAudioXWav                                                            SupportedContentType = "audio/x-wav"
	SupportedContentTypeAudioAiff                                                            SupportedContentType = "audio/aiff"
	SupportedContentTypeAudioXAiff                                                           SupportedContentType = "audio/x-aiff"
	SupportedContentTypeAudioAac                                                             SupportedContentType = "audio/aac"
	SupportedContentTypeAudioXAac                                                            SupportedContentType = "audio/x-aac"
	SupportedContentTypeAudioAmr                                                             SupportedContentType = "audio/amr"
	SupportedContentTypeApplicationPdf                                                       SupportedContentType = "application/pdf"
	SupportedContentTypeTextPlain                                                            SupportedContentType = "text/plain"
	SupportedContentTypeTextMarkdown                                                         SupportedContentType = "text/markdown"
	SupportedContentTypeTextVcard                                                            SupportedContentType = "text/vcard"
	SupportedContentTypeTextXVcard                                                           SupportedContentType = "text/x-vcard"
	SupportedContentTypeTextRtf                                                              SupportedContentType = "text/rtf"
	SupportedContentTypeApplicationRtf                                                       SupportedContentType = "application/rtf"
	SupportedContentTypeTextCsv                                                              SupportedContentType = "text/csv"
	SupportedContentTypeTextHTML                                                             SupportedContentType = "text/html"
	SupportedContentTypeTextCalendar                                                         SupportedContentType = "text/calendar"
	SupportedContentTypeApplicationMsword                                                    SupportedContentType = "application/msword"
	SupportedContentTypeApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocument   SupportedContentType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	SupportedContentTypeApplicationVndMsExcel                                                SupportedContentType = "application/vnd.ms-excel"
	SupportedContentTypeApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheet         SupportedContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	SupportedContentTypeApplicationVndMsPowerpoint                                           SupportedContentType = "application/vnd.ms-powerpoint"
	SupportedContentTypeApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresentation SupportedContentType = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	SupportedContentTypeApplicationVndApplePages                                             SupportedContentType = "application/vnd.apple.pages"
	SupportedContentTypeApplicationXIworkPagesSffpages                                       SupportedContentType = "application/x-iwork-pages-sffpages"
	SupportedContentTypeApplicationVndAppleNumbers                                           SupportedContentType = "application/vnd.apple.numbers"
	SupportedContentTypeApplicationXIworkNumbersSffnumbers                                   SupportedContentType = "application/x-iwork-numbers-sffnumbers"
	SupportedContentTypeApplicationVndAppleKeynote                                           SupportedContentType = "application/vnd.apple.keynote"
	SupportedContentTypeApplicationXIworkKeynoteSffkey                                       SupportedContentType = "application/x-iwork-keynote-sffkey"
	SupportedContentTypeApplicationEpubZip                                                   SupportedContentType = "application/epub+zip"
	SupportedContentTypeApplicationZip                                                       SupportedContentType = "application/zip"
	SupportedContentTypeApplicationXZipCompressed                                            SupportedContentType = "application/x-zip-compressed"
)

type AttachmentNewResponse struct {
	// Unique identifier for the attachment (for status checks via GET
	// /v3/attachments/{id})
	AttachmentID string `json:"attachment_id" api:"required" format:"uuid"`
	// Permanent CDN URL for the file. Does not expire. Use the `attachment_id` to
	// reference this file in media parts when sending messages.
	DownloadURL string `json:"download_url" api:"required" format:"uri"`
	// When the upload URL expires (15 minutes from now)
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// HTTP method to use for upload (always PUT)
	//
	// Any of "PUT".
	HTTPMethod AttachmentNewResponseHTTPMethod `json:"http_method" api:"required"`
	// HTTP headers required for the upload request
	RequiredHeaders map[string]string `json:"required_headers" api:"required"`
	// Presigned URL for uploading the file. PUT the raw binary file content to this
	// URL with the `required_headers`. Do not JSON-encode or multipart-wrap the body.
	// Expires after 15 minutes.
	UploadURL string `json:"upload_url" api:"required" format:"uri"`
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
	ID string `json:"id" api:"required"`
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
	ContentType SupportedContentType `json:"content_type" api:"required"`
	// When the attachment was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Original filename of the attachment
	Filename string `json:"filename" api:"required"`
	// Size of the attachment in bytes
	SizeBytes int64 `json:"size_bytes" api:"required"`
	// Current upload/processing status
	//
	// Any of "pending", "complete", "failed".
	Status AttachmentGetResponseStatus `json:"status" api:"required"`
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
	ContentType SupportedContentType `json:"content_type,omitzero" api:"required"`
	// Name of the file to upload
	Filename string `json:"filename" api:"required"`
	// Size of the file in bytes (max 100MB)
	SizeBytes int64 `json:"size_bytes" api:"required"`
	paramObj
}

func (r AttachmentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AttachmentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttachmentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
