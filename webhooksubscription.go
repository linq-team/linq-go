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
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// WebhookSubscriptionService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookSubscriptionService] method instead.
type WebhookSubscriptionService struct {
	Options []option.RequestOption
}

// NewWebhookSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWebhookSubscriptionService(opts ...option.RequestOption) (r WebhookSubscriptionService) {
	r = WebhookSubscriptionService{}
	r.Options = opts
	return
}

// Create a new webhook subscription to receive events at a target URL. Upon
// creation, a signing secret is generated for verifying webhook authenticity.
// **Store this secret securely — it cannot be retrieved later.**
//
// **Webhook Delivery:**
//
//   - Events are sent via HTTP POST to the target URL
//   - Each request includes `X-Webhook-Signature` and `X-Webhook-Timestamp` headers
//   - Signature is HMAC-SHA256 over `{timestamp}.{payload}` — see
//     [Webhook Events](/docs/webhook-events) for verification details
//   - Failed deliveries (5xx, 429, network errors) are retried up to 10 times over
//     ~2 hours with exponential backoff
//   - Client errors (4xx except 429) are not retried
func (r *WebhookSubscriptionService) New(ctx context.Context, body WebhookSubscriptionNewParams, opts ...option.RequestOption) (res *WebhookSubscriptionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/webhook-subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details for a specific webhook subscription including its target URL,
// subscribed events, and current status.
func (r *WebhookSubscriptionService) Get(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *WebhookSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return
	}
	path := fmt.Sprintf("v3/webhook-subscriptions/%s", url.PathEscape(subscriptionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing webhook subscription. You can modify the target URL,
// subscribed events, or activate/deactivate the subscription.
//
// **Note:** The signing secret cannot be changed via this endpoint.
func (r *WebhookSubscriptionService) Update(ctx context.Context, subscriptionID string, body WebhookSubscriptionUpdateParams, opts ...option.RequestOption) (res *WebhookSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return
	}
	path := fmt.Sprintf("v3/webhook-subscriptions/%s", url.PathEscape(subscriptionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve all webhook subscriptions for the authenticated partner. Returns a list
// of active and inactive subscriptions with their configuration and status.
func (r *WebhookSubscriptionService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookSubscriptionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/webhook-subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a webhook subscription.
func (r *WebhookSubscriptionService) Delete(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return
	}
	path := fmt.Sprintf("v3/webhook-subscriptions/%s", url.PathEscape(subscriptionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type WebhookSubscription struct {
	// Unique identifier for the webhook subscription
	ID string `json:"id,required"`
	// When the subscription was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether this subscription is currently active
	IsActive bool `json:"is_active,required"`
	// List of event types this subscription receives
	SubscribedEvents []WebhookEventType `json:"subscribed_events,required"`
	// URL where webhook events will be sent
	TargetURL string `json:"target_url,required" format:"uri"`
	// When the subscription was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		IsActive         respjson.Field
		SubscribedEvents respjson.Field
		TargetURL        respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookSubscription) RawJSON() string { return r.JSON.raw }
func (r *WebhookSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response returned when creating a webhook subscription. Includes the signing
// secret which is only shown once.
type WebhookSubscriptionNewResponse struct {
	// Unique identifier for the webhook subscription
	ID string `json:"id,required"`
	// When the subscription was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether this subscription is currently active
	IsActive bool `json:"is_active,required"`
	// Secret for verifying webhook signatures. Store this securely - it cannot be
	// retrieved again.
	SigningSecret string `json:"signing_secret,required"`
	// List of event types this subscription receives
	SubscribedEvents []WebhookEventType `json:"subscribed_events,required"`
	// URL where webhook events will be sent
	TargetURL string `json:"target_url,required" format:"uri"`
	// When the subscription was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		IsActive         respjson.Field
		SigningSecret    respjson.Field
		SubscribedEvents respjson.Field
		TargetURL        respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookSubscriptionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookSubscriptionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookSubscriptionListResponse struct {
	// List of webhook subscriptions
	Subscriptions []WebhookSubscription `json:"subscriptions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Subscriptions respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookSubscriptionListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookSubscriptionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookSubscriptionNewParams struct {
	// List of event types to subscribe to
	SubscribedEvents []WebhookEventType `json:"subscribed_events,omitzero,required"`
	// URL where webhook events will be sent. Must be HTTPS.
	TargetURL string `json:"target_url,required" format:"uri"`
	paramObj
}

func (r WebhookSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookSubscriptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookSubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookSubscriptionUpdateParams struct {
	// Activate or deactivate the subscription
	IsActive param.Opt[bool] `json:"is_active,omitzero"`
	// New target URL for webhook events
	TargetURL param.Opt[string] `json:"target_url,omitzero" format:"uri"`
	// Updated list of event types to subscribe to
	SubscribedEvents []WebhookEventType `json:"subscribed_events,omitzero"`
	paramObj
}

func (r WebhookSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookSubscriptionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookSubscriptionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
