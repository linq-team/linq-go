// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// Let an agent pay on a customer's behalf with a single-use virtual card. Connect
// a customer once, then create a payment — a virtual card is minted scoped to that
// purchase and the card details are handed back for checkout.
//
// PaymentService contains methods and other services that help with interacting
// with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	Options []option.RequestOption
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r PaymentService) {
	r = PaymentService{}
	r.Options = opts
	return
}

// Advances the pay flow for a connected customer handle and returns a `status`
// describing where it is (`needs_connection`, `awaiting_user_action`, `ready`,
// ...). A payment `id` appears once a card is minted. Idempotent on the
// `Idempotency-Key` header.
func (r *PaymentService) New(ctx context.Context, body PaymentNewParams, opts ...option.RequestOption) (res *Payment, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a payment
func (r *PaymentService) Get(ctx context.Context, paymentID string, opts ...option.RequestOption) (res *Payment, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentID == "" {
		err = errors.New("missing required paymentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/payments/%s", url.PathEscape(paymentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Closes the virtual card and cancels the payment.
func (r *PaymentService) Cancel(ctx context.Context, paymentID string, opts ...option.RequestOption) (res *Payment, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentID == "" {
		err = errors.New("missing required paymentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/payments/%s/cancel", url.PathEscape(paymentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns a short-lived handoff for a `ready` payment. Fetch the card credentials
// **directly from the provider** with the returned `user_token` at `fetch_url` —
// the card number never passes through Linq. Do not persist PAN/CVC.
func (r *PaymentService) Credentials(ctx context.Context, paymentID string, opts ...option.RequestOption) (res *PaymentCredentialsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentID == "" {
		err = errors.New("missing required paymentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/payments/%s/credentials", url.PathEscape(paymentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Payment struct {
	ID          string `json:"id"`
	AmountCents int64  `json:"amount_cents"`
	// Present when the customer must approve with a passkey.
	ApprovalURL string `json:"approval_url"`
	// Present when the customer must attach a card.
	AttachURL   string `json:"attach_url"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Handle      string `json:"handle"`
	// Any of "needs_connection", "connecting", "awaiting_user_action", "ready",
	// "authorized", "succeeded", "declined", "canceled", "expired".
	Status PaymentStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AmountCents respjson.Field
		ApprovalURL respjson.Field
		AttachURL   respjson.Field
		Currency    respjson.Field
		Description respjson.Field
		Handle      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Payment) RawJSON() string { return r.JSON.raw }
func (r *Payment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentStatus string

const (
	PaymentStatusNeedsConnection    PaymentStatus = "needs_connection"
	PaymentStatusConnecting         PaymentStatus = "connecting"
	PaymentStatusAwaitingUserAction PaymentStatus = "awaiting_user_action"
	PaymentStatusReady              PaymentStatus = "ready"
	PaymentStatusAuthorized         PaymentStatus = "authorized"
	PaymentStatusSucceeded          PaymentStatus = "succeeded"
	PaymentStatusDeclined           PaymentStatus = "declined"
	PaymentStatusCanceled           PaymentStatus = "canceled"
	PaymentStatusExpired            PaymentStatus = "expired"
)

type PaymentCredentialsResponse struct {
	// Fetch the card directly from the provider with these — never through Linq.
	Handoff PaymentCredentialsResponseHandoff `json:"handoff"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handoff     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentCredentialsResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentCredentialsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fetch the card directly from the provider with these — never through Linq.
type PaymentCredentialsResponseHandoff struct {
	CardRef  string `json:"card_ref"`
	FetchURL string `json:"fetch_url"`
	Provider string `json:"provider"`
	// Short-lived bearer to fetch the card from the provider.
	UserToken string `json:"user_token"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CardRef     respjson.Field
		FetchURL    respjson.Field
		Provider    respjson.Field
		UserToken   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentCredentialsResponseHandoff) RawJSON() string { return r.JSON.raw }
func (r *PaymentCredentialsResponseHandoff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentNewParams struct {
	AmountCents int64  `json:"amount_cents" api:"required"`
	Currency    string `json:"currency" api:"required"`
	// Customer phone (E.164) or email.
	Handle      string                   `json:"handle" api:"required"`
	Description param.Opt[string]        `json:"description,omitzero"`
	Merchant    PaymentNewParamsMerchant `json:"merchant,omitzero"`
	Metadata    map[string]string        `json:"metadata,omitzero"`
	paramObj
}

func (r PaymentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentNewParamsMerchant struct {
	Name param.Opt[string] `json:"name,omitzero"`
	URL  param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r PaymentNewParamsMerchant) MarshalJSON() (data []byte, err error) {
	type shadow PaymentNewParamsMerchant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentNewParamsMerchant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
