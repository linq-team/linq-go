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
// AvailableNumberService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAvailableNumberService] method instead.
type AvailableNumberService struct {
	Options []option.RequestOption
}

// NewAvailableNumberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAvailableNumberService(opts ...option.RequestOption) (r AvailableNumberService) {
	r = AvailableNumberService{}
	r.Options = opts
	return
}

// Returns the best available line (E.164) to send from, applying smart number
// assignment. Optionally pass `to` recipients to make the choice "sticky" —
// reusing the line an existing chat with those recipients is already on. Without
// `to`, the best healthy line is chosen.
//
// This is advisory: it does not reserve the line or change selection state. Pass
// the returned `phone_number` as `from` when you create the chat to guarantee the
// same line.
//
// Also returns `vcf_url`: a time-limited link to a vCard (`.vcf`) for the chosen
// line, carrying its contact card (name/photo) with the chosen number as the
// primary `TEL` and the partner's other healthy lines as backups. Share it with
// recipients so they can save the line as a contact.
func (r *AvailableNumberService) Get(ctx context.Context, query AvailableNumberGetParams, opts ...option.RequestOption) (res *AvailableNumberGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/available_number"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// The line smart number assignment selected, plus a shareable vCard.
type AvailableNumberGetResponse struct {
	// The selected sending line in E.164 format.
	PhoneNumber string `json:"phone_number" api:"required"`
	// Time-limited link to a vCard (`.vcf`) for the selected line. The card carries
	// the line's contact details with the selected number as the primary `TEL` and the
	// partner's other healthy lines as backups. The link expires; re-call this
	// endpoint to mint a fresh one.
	VcfURL string `json:"vcf_url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		VcfURL      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailableNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AvailableNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableNumberGetParams struct {
	// Recipient handles (E.164 or email) the message is destined for. When provided,
	// an existing chat with these recipients makes the choice sticky. Repeat the
	// parameter for multiple recipients.
	To []string `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AvailableNumberGetParams]'s query parameters as
// `url.Values`.
func (r AvailableNumberGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
