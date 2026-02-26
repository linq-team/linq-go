// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/respjson"
)

// PhoneNumberService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberService] method instead.
type PhoneNumberService struct {
	Options []option.RequestOption
}

// NewPhoneNumberService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPhoneNumberService(opts ...option.RequestOption) (r PhoneNumberService) {
	r = PhoneNumberService{}
	r.Options = opts
	return
}

// Returns all phone numbers assigned to the authenticated partner. Use this
// endpoint to discover which phone numbers are available for use as the `from`
// field when creating a chat, listing chats, or sending a voice memo.
func (r *PhoneNumberService) List(ctx context.Context, opts ...option.RequestOption) (res *PhoneNumberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/phone_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PhoneNumberListResponse struct {
	// List of phone numbers assigned to the partner
	PhoneNumbers []PhoneNumberListResponsePhoneNumber `json:"phone_numbers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumbers respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberListResponsePhoneNumber struct {
	// Unique identifier for the phone number
	ID string `json:"id" api:"required" format:"uuid"`
	// Phone number in E.164 format
	PhoneNumber string `json:"phone_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberListResponsePhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberListResponsePhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
