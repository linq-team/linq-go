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
func (r *CapabilityService) CheckImessage(ctx context.Context, body CapabilityCheckImessageParams, opts ...option.RequestOption) (res *CapabilityCheckImessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/capability/check_imessage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check whether a recipient address (phone number) supports RCS messaging.
func (r *CapabilityService) CheckRcs(ctx context.Context, body CapabilityCheckRcsParams, opts ...option.RequestOption) (res *CapabilityCheckRcsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/capability/check_rcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CapabilityCheckImessageResponse struct {
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
func (r CapabilityCheckImessageResponse) RawJSON() string { return r.JSON.raw }
func (r *CapabilityCheckImessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilityCheckRcsResponse struct {
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
func (r CapabilityCheckRcsResponse) RawJSON() string { return r.JSON.raw }
func (r *CapabilityCheckRcsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilityCheckImessageParams struct {
	// The recipient phone number or email address to check
	Address string `json:"address" api:"required"`
	// Optional sender phone number. If omitted, an available phone from your pool is
	// used automatically.
	From param.Opt[string] `json:"from,omitzero"`
	paramObj
}

func (r CapabilityCheckImessageParams) MarshalJSON() (data []byte, err error) {
	type shadow CapabilityCheckImessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CapabilityCheckImessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilityCheckRcsParams struct {
	// The recipient phone number or email address to check
	Address string `json:"address" api:"required"`
	// Optional sender phone number. If omitted, an available phone from your pool is
	// used automatically.
	From param.Opt[string] `json:"from,omitzero"`
	paramObj
}

func (r CapabilityCheckRcsParams) MarshalJSON() (data []byte, err error) {
	type shadow CapabilityCheckRcsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CapabilityCheckRcsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
