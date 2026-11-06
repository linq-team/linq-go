// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/apiquery"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// Read the carrier (BYOC) voice calls received on your phone numbers. Each record
// is the durable form of the `call_voice.*` webhooks for that call and shares the
// same `id`, so you can look a call up by the id a webhook delivered. See SIP
// Profiles for routing inbound calls to your endpoint.
//
// CallVoiceService contains methods and other services that help with interacting
// with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallVoiceService] method instead.
type CallVoiceService struct {
	Options []option.RequestOption
}

// NewCallVoiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCallVoiceService(opts ...option.RequestOption) (r CallVoiceService) {
	r = CallVoiceService{}
	r.Options = opts
	return
}

// Retrieve a single carrier voice call by its `id` — the same id delivered in the
// `call_voice.*` webhooks and returned by the list endpoint. The call must belong
// to the authenticated partner.
func (r *CallVoiceService) Get(ctx context.Context, callVoiceID string, opts ...option.RequestOption) (res *CallVoiceRecord, err error) {
	opts = slices.Concat(r.Options, opts)
	if callVoiceID == "" {
		err = errors.New("missing required callVoiceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/call-voice/%s", callVoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the carrier voice calls received on one of your phone numbers, newest
// first, with cursor pagination. Each record is the durable form of the
// `call_voice.*` webhooks for that call; the `id` is the same id those webhooks
// carry. Results are scoped to the authenticated partner — a number that isn't
// yours (or simply has no calls) returns `200` with an empty `call_voice` array,
// not an error.
func (r *CallVoiceService) List(ctx context.Context, phoneNumber string, query CallVoiceListParams, opts ...option.RequestOption) (res *CallVoiceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/phone-numbers/%s/call-voice", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A carrier (BYOC) inbound voice call. Mirrors the call_voice.\* webhook payload.
type CallVoiceRecord struct {
	// Stable call id, identical to the id in the call_voice.\* webhooks.
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "inbound".
	Direction CallVoiceRecordDirection `json:"direction" api:"required"`
	// External caller, E.164.
	From string `json:"from" api:"required"`
	// Terminal call outcome. Null while the call is still in progress; set once it
	// ends.
	//
	// Any of "completed", "busy", "no_answer", "failed", "canceled".
	Status CallVoiceRecordStatus `json:"status" api:"required"`
	// Your line that was called, E.164.
	To string `json:"to" api:"required"`
	// When the call ended. Null while still in progress.
	EndedAt time.Time `json:"ended_at" api:"nullable" format:"date-time"`
	// When the call was answered. Null if it never connected.
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Direction   respjson.Field
		From        respjson.Field
		Status      respjson.Field
		To          respjson.Field
		EndedAt     respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallVoiceRecord) RawJSON() string { return r.JSON.raw }
func (r *CallVoiceRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallVoiceRecordDirection string

const (
	CallVoiceRecordDirectionInbound CallVoiceRecordDirection = "inbound"
)

// Terminal call outcome. Null while the call is still in progress; set once it
// ends.
type CallVoiceRecordStatus string

const (
	CallVoiceRecordStatusCompleted CallVoiceRecordStatus = "completed"
	CallVoiceRecordStatusBusy      CallVoiceRecordStatus = "busy"
	CallVoiceRecordStatusNoAnswer  CallVoiceRecordStatus = "no_answer"
	CallVoiceRecordStatusFailed    CallVoiceRecordStatus = "failed"
	CallVoiceRecordStatusCanceled  CallVoiceRecordStatus = "canceled"
)

type CallVoiceListResponse struct {
	// Calls for the phone number, newest first.
	CallVoice []CallVoiceRecord `json:"call_voice" api:"required"`
	// Cursor for the next page of results. Null when there are no more. Pass this
	// value as the `cursor` parameter in the next request.
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallVoice   respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallVoiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *CallVoiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallVoiceListParams struct {
	// Pagination cursor from a previous next_cursor response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of calls to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CallVoiceListParams]'s query parameters as `url.Values`.
func (r CallVoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
