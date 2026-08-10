// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// Block handles — phone numbers, email addresses, SMS short codes, or sender IDs.
// Inbound messages from a blocked handle are dropped before they reach your
// webhooks, and direct sends to a blocked handle are rejected with `403` (error
// code `2026`). Group sends that include unblocked members are not restricted.
//
// BlockedHandleService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockedHandleService] method instead.
type BlockedHandleService struct {
	Options []option.RequestOption
}

// NewBlockedHandleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockedHandleService(opts ...option.RequestOption) (r BlockedHandleService) {
	r = BlockedHandleService{}
	r.Options = opts
	return
}

// Returns all handles you have blocked. Inbound messages from a blocked handle are
// dropped and produce no webhooks, and direct sends to a blocked handle are
// rejected with `403` (error code `2026`). Group sends that include unblocked
// members are not restricted.
func (r *BlockedHandleService) List(ctx context.Context, opts ...option.RequestOption) (res *BlockedHandleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/blocked_handles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Blocks a handle — an E.164 phone number, an email address (iMessage sender), an
// SMS short code (e.g. `262966`), or an alphanumeric sender ID. Inbound messages
// from it are dropped and produce no webhooks, and direct sends to it are rejected
// with `403` (error code `2026`); group sends that include unblocked members are
// not restricted. Blocking is idempotent — re-blocking an already blocked handle
// returns the existing entry.
func (r *BlockedHandleService) Block(ctx context.Context, body BlockedHandleBlockParams, opts ...option.RequestOption) (res *BlockedHandleBlockResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/blocked_handles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Removes a handle from your blocklist. Inbound messages from it will be delivered
// again and sends to it are allowed again. The handle goes in the request body,
// mirroring block — no URL encoding needed.
func (r *BlockedHandleService) Unblock(ctx context.Context, body BlockedHandleUnblockParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v3/blocked_handles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

type BlockedHandleEntry struct {
	// When the handle was blocked
	BlockedAt time.Time `json:"blocked_at" api:"required" format:"date-time"`
	// The blocked handle, normalized (E.164 phone, lowercased email, short code, or
	// sender ID)
	Handle string `json:"handle" api:"required"`
	// Optional note recorded when the handle was blocked
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockedAt   respjson.Field
		Handle      respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BlockedHandleEntry) RawJSON() string { return r.JSON.raw }
func (r *BlockedHandleEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlockedHandleListResponse struct {
	// All handles blocked by the partner, newest first
	BlockedHandles []BlockedHandleEntry `json:"blocked_handles" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockedHandles respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BlockedHandleListResponse) RawJSON() string { return r.JSON.raw }
func (r *BlockedHandleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlockedHandleBlockResponse struct {
	BlockedHandle BlockedHandleEntry `json:"blocked_handle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockedHandle respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BlockedHandleBlockResponse) RawJSON() string { return r.JSON.raw }
func (r *BlockedHandleBlockResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlockedHandleBlockParams struct {
	// The handle to block: an E.164 phone number, an email address, an SMS short code
	// (3-8 digits), or an alphanumeric sender ID.
	Handle string `json:"handle" api:"required"`
	// Optional free-text note on why the handle was blocked
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r BlockedHandleBlockParams) MarshalJSON() (data []byte, err error) {
	type shadow BlockedHandleBlockParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlockedHandleBlockParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlockedHandleUnblockParams struct {
	// The handle to unblock
	Handle string `json:"handle" api:"required"`
	paramObj
}

func (r BlockedHandleUnblockParams) MarshalJSON() (data []byte, err error) {
	type shadow BlockedHandleUnblockParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlockedHandleUnblockParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
