// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// Check whether a recipient address supports iMessage or RCS before sending a
// message.
//
// CapabilityService contains methods and other services that help with interacting
// with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCapabilityService] method instead.
type CapabilityService struct {
	Options []option.RequestOption
}

// NewCapabilityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCapabilityService(opts ...option.RequestOption) (r CapabilityService) {
	r = CapabilityService{}
	r.Options = opts
	return
}

// Check whether a recipient address (phone number or email) is reachable via
// iMessage.
func (r *CapabilityService) CheckiMessage(ctx context.Context, body CapabilityCheckiMessageParams, opts ...option.RequestOption) (res *CapabilityCheckiMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/capability/check_imessage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Check whether a recipient address (phone number) supports RCS messaging.
func (r *CapabilityService) CheckRCS(ctx context.Context, body CapabilityCheckRCSParams, opts ...option.RequestOption) (res *CapabilityCheckRCSResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/capability/check_rcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CapabilityCheckiMessageResponse struct {
	// The recipient address that was checked
	Address string `json:"address" api:"required"`
	// Whether the recipient supports the checked messaging service
	Available bool `json:"available" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Available   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilityCheckiMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *CapabilityCheckiMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilityCheckRCSResponse struct {
	// The recipient address that was checked
	Address string `json:"address" api:"required"`
	// Whether the recipient supports the checked messaging service
	Available bool `json:"available" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Available   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilityCheckRCSResponse) RawJSON() string { return r.JSON.raw }
func (r *CapabilityCheckRCSResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilityCheckiMessageParams struct {
	// The recipient phone number or email address to check
	Address string `json:"address" api:"required"`
	// Optional sender phone number. If omitted, an available phone from your pool is
	// used automatically.
	From param.Opt[string] `json:"from,omitzero"`
	paramObj
}

func (r CapabilityCheckiMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow CapabilityCheckiMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CapabilityCheckiMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilityCheckRCSParams struct {
	// The recipient phone number or email address to check
	Address string `json:"address" api:"required"`
	// Optional sender phone number. If omitted, an available phone from your pool is
	// used automatically.
	From param.Opt[string] `json:"from,omitzero"`
	paramObj
}

func (r CapabilityCheckRCSParams) MarshalJSON() (data []byte, err error) {
	type shadow CapabilityCheckRCSParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CapabilityCheckRCSParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
