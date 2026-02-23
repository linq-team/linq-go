// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqapiv3

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/linq-api-v3-go/internal/apijson"
	"github.com/stainless-sdks/linq-api-v3-go/internal/requestconfig"
	"github.com/stainless-sdks/linq-api-v3-go/option"
	"github.com/stainless-sdks/linq-api-v3-go/packages/respjson"
)

// PhoneNumberService contains methods and other services that help with
// interacting with the linq API.
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

// **Deprecated.** Use `GET /v3/phone_numbers` instead.
//
// Deprecated: Use `list` instead, which calls `GET /v3/phone_numbers`.
func (r *PhoneNumberService) ListDeprecated(ctx context.Context, opts ...option.RequestOption) (res *PhoneNumberListDeprecatedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/phonenumbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PhoneNumberListResponse struct {
	// List of phone numbers assigned to the partner
	PhoneNumbers []PhoneNumberListResponsePhoneNumber `json:"phone_numbers,required"`
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
	ID string `json:"id,required" format:"uuid"`
	// Phone number in E.164 format
	PhoneNumber string `json:"phone_number,required"`
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

type PhoneNumberListDeprecatedResponse struct {
	// List of phone numbers assigned to the partner
	PhoneNumbers []PhoneNumberListDeprecatedResponsePhoneNumber `json:"phone_numbers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumbers respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberListDeprecatedResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberListDeprecatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberListDeprecatedResponsePhoneNumber struct {
	// Unique identifier for the phone number
	ID string `json:"id,required" format:"uuid"`
	// Phone number in E.164 format
	PhoneNumber  string                                                   `json:"phone_number,required"`
	Capabilities PhoneNumberListDeprecatedResponsePhoneNumberCapabilities `json:"capabilities"`
	// Deprecated. Always null.
	CountryCode string `json:"country_code"`
	// Deprecated. Always null.
	//
	// Any of "TWILIO", "APPLE_ID".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		PhoneNumber  respjson.Field
		Capabilities respjson.Field
		CountryCode  respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberListDeprecatedResponsePhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberListDeprecatedResponsePhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberListDeprecatedResponsePhoneNumberCapabilities struct {
	// Whether MMS messaging is supported
	Mms bool `json:"mms,required"`
	// Whether SMS messaging is supported
	SMS bool `json:"sms,required"`
	// Whether voice calls are supported
	Voice bool `json:"voice,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mms         respjson.Field
		SMS         respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberListDeprecatedResponsePhoneNumberCapabilities) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberListDeprecatedResponsePhoneNumberCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
