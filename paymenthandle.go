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
// PaymentHandleService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentHandleService] method instead.
type PaymentHandleService struct {
	Options []option.RequestOption
}

// NewPaymentHandleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentHandleService(opts ...option.RequestOption) (r PaymentHandleService) {
	r = PaymentHandleService{}
	r.Options = opts
	return
}

// Starts connecting a customer (by phone/email) so an agent can pay on their
// behalf. Linq drives the OTP + consent ceremony through the messaging channel;
// this returns `pending` and a `connection.created` webhook fires once the
// customer completes it.
func (r *PaymentHandleService) Connect(ctx context.Context, handle string, opts ...option.RequestOption) (res *PaymentHandleConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if handle == "" {
		err = errors.New("missing required handle parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/payments/handles/%s/connect", url.PathEscape(handle))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get a handle's connection status
func (r *PaymentHandleService) Connection(ctx context.Context, handle string, opts ...option.RequestOption) (res *PaymentHandleConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if handle == "" {
		err = errors.New("missing required handle parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/payments/handles/%s/connection", url.PathEscape(handle))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revokes this partner's grant for the customer. Only your grant is removed; the
// customer's wallet at the provider is untouched.
func (r *PaymentHandleService) Revoke(ctx context.Context, handle string, opts ...option.RequestOption) (res *PaymentHandleConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if handle == "" {
		err = errors.New("missing required handle parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/payments/handles/%s/connection", url.PathEscape(handle))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Completes the ceremony `connect` started: verifies the code, records the
// customer's consent, and stores the connection. Returns `connected` on success,
// after which payments for this handle no longer need the customer present.
//
// The code reaches you however your channel works — typically the customer replies
// with it in the thread. Codes are single-use and short-lived; if one has expired,
// call `connect` again for a fresh `connect_id`.
func (r *PaymentHandleService) Verify(ctx context.Context, handle string, body PaymentHandleVerifyParams, opts ...option.RequestOption) (res *PaymentHandleConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if handle == "" {
		err = errors.New("missing required handle parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/payments/handles/%s/verify", url.PathEscape(handle))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PaymentHandleConnection struct {
	// Returned only by `connect`, and only while the ceremony is pending. Nothing on
	// our side persists it — it comes back from the provider and is required again to
	// verify — so hold it until you submit the code.
	ConnectID string `json:"connect_id"`
	Handle    string `json:"handle"`
	// Any of "not_connected", "pending", "connected", "revoked".
	Status PaymentHandleConnectionStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectID   respjson.Field
		Handle      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentHandleConnection) RawJSON() string { return r.JSON.raw }
func (r *PaymentHandleConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentHandleConnectionStatus string

const (
	PaymentHandleConnectionStatusNotConnected PaymentHandleConnectionStatus = "not_connected"
	PaymentHandleConnectionStatusPending      PaymentHandleConnectionStatus = "pending"
	PaymentHandleConnectionStatusConnected    PaymentHandleConnectionStatus = "connected"
	PaymentHandleConnectionStatusRevoked      PaymentHandleConnectionStatus = "revoked"
)

type PaymentHandleVerifyParams struct {
	// The one-time code the customer received.
	Code string `json:"code" api:"required"`
	// The id returned by `connect`.
	ConnectID string `json:"connect_id" api:"required"`
	paramObj
}

func (r PaymentHandleVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentHandleVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentHandleVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
