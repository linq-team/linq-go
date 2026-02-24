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

// PhonenumberService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhonenumberService] method instead.
type PhonenumberService struct {
	Options []option.RequestOption
}

// NewPhonenumberService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPhonenumberService(opts ...option.RequestOption) (r PhonenumberService) {
	r = PhonenumberService{}
	r.Options = opts
	return
}

// **Deprecated.** Use `GET /v3/phone_numbers` instead.
//
// Deprecated: deprecated
func (r *PhonenumberService) List(ctx context.Context, opts ...option.RequestOption) (res *PhonenumberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/phonenumbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PhonenumberListResponse struct {
	// List of phone numbers assigned to the partner
	PhoneNumbers []PhonenumberListResponsePhoneNumber `json:"phone_numbers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumbers respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhonenumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *PhonenumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhonenumberListResponsePhoneNumber struct {
	// Unique identifier for the phone number
	ID string `json:"id,required" format:"uuid"`
	// Phone number in E.164 format
	PhoneNumber  string                                         `json:"phone_number,required"`
	Capabilities PhonenumberListResponsePhoneNumberCapabilities `json:"capabilities"`
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
func (r PhonenumberListResponsePhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *PhonenumberListResponsePhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhonenumberListResponsePhoneNumberCapabilities struct {
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
func (r PhonenumberListResponsePhoneNumberCapabilities) RawJSON() string { return r.JSON.raw }
func (r *PhonenumberListResponsePhoneNumberCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
