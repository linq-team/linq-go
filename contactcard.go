// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/apiquery"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// Contact Card lets you set and share your contact information (name and profile
// photo) with chat participants via iMessage Name and Photo Sharing.
//
// Use `POST /v3/contact_card` to create or update a card for a phone number. Use
// `PATCH /v3/contact_card` to update an existing active card. Use
// `GET /v3/contact_card` to retrieve the active card(s) for your partner account.
//
// ContactCardService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContactCardService] method instead.
type ContactCardService struct {
	Options []option.RequestOption
}

// NewContactCardService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContactCardService(opts ...option.RequestOption) (r ContactCardService) {
	r = ContactCardService{}
	r.Options = opts
	return
}

// Creates a contact card for a phone number.
//
// The contact card is stored in an inactive state first. Once it's applied
// successfully, it is activated and `is_active` is returned as `true`. On failure,
// `is_active` is `false`.
func (r *ContactCardService) New(ctx context.Context, body ContactCardNewParams, opts ...option.RequestOption) (res *SetContactCard, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/contact_card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the contact card for a specific phone number, or all contact cards for
// the authenticated partner if no `phone_number` is provided.
func (r *ContactCardService) Get(ctx context.Context, query ContactCardGetParams, opts ...option.RequestOption) (res *ContactCardGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/contact_card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates an existing active contact card for a phone number.
//
// Fetches the current active contact card and merges the provided fields. Only
// fields present in the request body are updated; omitted fields retain their
// existing values.
//
// Requires an active contact card to exist for the phone number.
func (r *ContactCardService) Update(ctx context.Context, params ContactCardUpdateParams, opts ...option.RequestOption) (res *SetContactCard, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/contact_card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type SetContactCard struct {
	// First name on the contact card
	FirstName string `json:"first_name" api:"required"`
	// Whether the contact card was successfully applied to the device
	IsActive bool `json:"is_active" api:"required"`
	// The phone number the contact card is associated with
	PhoneNumber string `json:"phone_number" api:"required"`
	// Image URL on the contact card
	ImageURL string `json:"image_url"`
	// Last name on the contact card
	LastName string `json:"last_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		IsActive    respjson.Field
		PhoneNumber respjson.Field
		ImageURL    respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SetContactCard) RawJSON() string { return r.JSON.raw }
func (r *SetContactCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactCardGetResponse struct {
	ContactCards []ContactCardGetResponseContactCard `json:"contact_cards" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactCards respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactCardGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactCardGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactCardGetResponseContactCard struct {
	FirstName   string `json:"first_name" api:"required"`
	IsActive    bool   `json:"is_active" api:"required"`
	PhoneNumber string `json:"phone_number" api:"required"`
	ImageURL    string `json:"image_url"`
	LastName    string `json:"last_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		IsActive    respjson.Field
		PhoneNumber respjson.Field
		ImageURL    respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactCardGetResponseContactCard) RawJSON() string { return r.JSON.raw }
func (r *ContactCardGetResponseContactCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactCardNewParams struct {
	// First name for the contact card. Required.
	FirstName string `json:"first_name" api:"required"`
	// E.164 phone number to associate the contact card with
	PhoneNumber string `json:"phone_number" api:"required"`
	// URL of the profile image to rehost on the CDN. Only re-uploaded when a new value
	// is provided.
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Last name for the contact card. Optional.
	LastName param.Opt[string] `json:"last_name,omitzero"`
	paramObj
}

func (r ContactCardNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactCardNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactCardNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactCardGetParams struct {
	// E.164 phone number to filter by. If omitted, all my cards for the partner are
	// returned.
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ContactCardGetParams]'s query parameters as `url.Values`.
func (r ContactCardGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContactCardUpdateParams struct {
	// E.164 phone number of the contact card to update
	PhoneNumber string `query:"phone_number" api:"required" json:"-"`
	// Updated first name. If omitted, the existing value is kept.
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// Updated profile image URL. If omitted, the existing image is kept.
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Updated last name. If omitted, the existing value is kept.
	LastName param.Opt[string] `json:"last_name,omitzero"`
	paramObj
}

func (r ContactCardUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactCardUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactCardUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ContactCardUpdateParams]'s query parameters as
// `url.Values`.
func (r ContactCardUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
