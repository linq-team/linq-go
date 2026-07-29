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
// PaymentProviderService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentProviderService] method instead.
type PaymentProviderService struct {
	Options []option.RequestOption
}

// NewPaymentProviderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentProviderService(opts ...option.RequestOption) (r PaymentProviderService) {
	r = PaymentProviderService{}
	r.Options = opts
	return
}

// Returns your organization's onboarding status for a payment provider.
func (r *PaymentProviderService) Get(ctx context.Context, provider string, opts ...option.RequestOption) (res *PaymentProvider, err error) {
	opts = slices.Concat(r.Options, opts)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/payments/providers/%s", url.PathEscape(provider))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Begins connecting your organization to a payment provider (e.g. `agentcard`).
// Returns a hosted URL where an admin authorizes the connection; on completion the
// provider redirects back and Linq stores your connected credentials.
func (r *PaymentProviderService) Connect(ctx context.Context, provider string, body PaymentProviderConnectParams, opts ...option.RequestOption) (res *PaymentProviderConnectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/payments/providers/%s/connect", url.PathEscape(provider))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PaymentProvider struct {
	Provider string `json:"provider"`
	// Any of "onboarding", "ready", "disabled".
	Status PaymentProviderStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Provider    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentProvider) RawJSON() string { return r.JSON.raw }
func (r *PaymentProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentProviderStatus string

const (
	PaymentProviderStatusOnboarding PaymentProviderStatus = "onboarding"
	PaymentProviderStatusReady      PaymentProviderStatus = "ready"
	PaymentProviderStatusDisabled   PaymentProviderStatus = "disabled"
)

type PaymentProviderConnectResponse struct {
	// Send the admin here to authorize the connection.
	HostedURL string `json:"hosted_url"`
	SessionID string `json:"session_id"`
	Status    string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HostedURL   respjson.Field
		SessionID   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentProviderConnectResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentProviderConnectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentProviderConnectParams struct {
	// Where to send the admin after they authorize the connection.
	ReturnURL string `json:"return_url" api:"required"`
	paramObj
}

func (r PaymentProviderConnectParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentProviderConnectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentProviderConnectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
