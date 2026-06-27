// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// Phone Numbers represent the phone numbers assigned to your partner account.
//
// Use the list phone numbers endpoint to discover which phone numbers are
// available for sending messages.
//
// When creating chats, listing chats, or sending a voice memo, use one of your
// assigned phone numbers in the `from` field.
//
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

// Updates the forwarding number for a phone number. The forwarding number is where
// inbound calls will be forwarded to.
//
// Pass an empty string to clear the forwarding number.
func (r *PhoneNumberService) Update(ctx context.Context, phoneNumberID string, body PhoneNumberUpdateParams, opts ...option.RequestOption) (res *PhoneNumberUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phoneNumberId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/phone_numbers/%s", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns all phone numbers assigned to the authenticated partner. Use this
// endpoint to discover which phone numbers are available for use as the `from`
// field when creating a chat, listing chats, or sending a voice memo.
func (r *PhoneNumberService) List(ctx context.Context, opts ...option.RequestOption) (res *PhoneNumberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/phone_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PhoneNumberUpdateResponse struct {
	// Unique identifier for the phone number
	ID string `json:"id" api:"required" format:"uuid"`
	// The forwarding number after the update. Null when cleared.
	ForwardingNumber string `json:"forwarding_number" api:"required"`
	// Phone number in E.164 format
	PhoneNumber string `json:"phone_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ForwardingNumber respjson.Field
		PhoneNumber      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// **[BETA]** Current reputation for a phone line. Always present — lines start at
	// `HEALTHY` and may shift based on aggregate engagement and delivery signals
	// across all conversations on the line.
	//
	// Unlike chat health, line reputation does not include `opted_out` — opt-out
	// applies to individual recipients, not the whole line.
	//
	// See the [Phone Reputation guide](/guides/phone-numbers/phone-reputation) for
	// what each status means and how to react.
	//
	// Deprecated: deprecated
	HealthStatus PhoneNumberListResponsePhoneNumberHealthStatus `json:"health_status" api:"required"`
	// Phone number in E.164 format
	PhoneNumber string `json:"phone_number" api:"required"`
	// **[BETA]** Current reputation for a phone line. Always present — lines start at
	// `HEALTHY` and may shift based on aggregate engagement and delivery signals
	// across all conversations on the line.
	//
	// Unlike chat health, line reputation does not include `opted_out` — opt-out
	// applies to individual recipients, not the whole line.
	//
	// See the [Phone Reputation guide](/guides/phone-numbers/phone-reputation) for
	// what each status means and how to react.
	Reputation PhoneNumberListResponsePhoneNumberReputation `json:"reputation" api:"required"`
	// The forwarding number associated with this phone number, in E.164 format. Null
	// when no forwarding number is configured.
	ForwardingNumber string `json:"forwarding_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		HealthStatus     respjson.Field
		PhoneNumber      respjson.Field
		Reputation       respjson.Field
		ForwardingNumber respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberListResponsePhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberListResponsePhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **[BETA]** Current reputation for a phone line. Always present — lines start at
// `HEALTHY` and may shift based on aggregate engagement and delivery signals
// across all conversations on the line.
//
// Unlike chat health, line reputation does not include `opted_out` — opt-out
// applies to individual recipients, not the whole line.
//
// See the [Phone Reputation guide](/guides/phone-numbers/phone-reputation) for
// what each status means and how to react.
//
// Deprecated: deprecated
type PhoneNumberListResponsePhoneNumberHealthStatus struct {
	// Deep-link to the relevant section of the Phone Reputation guide for this status.
	DocURL string `json:"doc_url" api:"required" format:"uri"`
	// Current reputation of this phone line as assessed by risk-service.
	//
	//   - `HEALTHY` — No elevated risk detected.
	//   - `AT_RISK` — Elevated risk indicators present; consider reducing send volume or
	//     reviewing messaging patterns.
	//   - `CRITICAL` — High risk; further sending may result in line flagging or
	//     restriction.
	//
	// Defaults to `HEALTHY` for lines that have not yet been scored.
	//
	// Any of "HEALTHY", "AT_RISK", "CRITICAL".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocURL      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberListResponsePhoneNumberHealthStatus) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberListResponsePhoneNumberHealthStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **[BETA]** Current reputation for a phone line. Always present — lines start at
// `HEALTHY` and may shift based on aggregate engagement and delivery signals
// across all conversations on the line.
//
// Unlike chat health, line reputation does not include `opted_out` — opt-out
// applies to individual recipients, not the whole line.
//
// See the [Phone Reputation guide](/guides/phone-numbers/phone-reputation) for
// what each status means and how to react.
type PhoneNumberListResponsePhoneNumberReputation struct {
	// Deep-link to the relevant section of the Phone Reputation guide for this status.
	DocURL string `json:"doc_url" api:"required" format:"uri"`
	// Current reputation of this phone line as assessed by risk-service.
	//
	//   - `HEALTHY` — No elevated risk detected.
	//   - `AT_RISK` — Elevated risk indicators present; consider reducing send volume or
	//     reviewing messaging patterns.
	//   - `CRITICAL` — High risk; further sending may result in line flagging or
	//     restriction.
	//
	// Defaults to `HEALTHY` for lines that have not yet been scored.
	//
	// Any of "HEALTHY", "AT_RISK", "CRITICAL".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocURL      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberListResponsePhoneNumberReputation) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberListResponsePhoneNumberReputation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberUpdateParams struct {
	// The forwarding number in E.164 format. Set to null or empty string to clear.
	ForwardingNumber param.Opt[string] `json:"forwarding_number,omitzero" api:"required"`
	paramObj
}

func (r PhoneNumberUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
