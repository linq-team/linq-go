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

// Request a payment from a recipient over iMessage. You create a payment request,
// send its `checkout_url` to the recipient, and they pay with Apple Pay or card.
// Funds settle **directly to your own Stripe account** — Linq never holds the
// money.
//
// ## How it works
//
//  1. **Create** a payment request with an amount and currency. You get back a
//     `checkout_url` and a `status` of `requested`.
//  2. **Send** the `checkout_url` to the recipient as a `link` message part so it
//     arrives as a tappable card (see _Sending the link_ below).
//  3. The recipient **pays** on the hosted checkout (Apple Pay App Clip on a
//     supported iPhone, web checkout everywhere else).
//  4. You receive a **`payment.succeeded`** webhook and the request's `status`
//     becomes `succeeded`. Requests you don't collect eventually `expire`.
//
// ## Connected accounts (Stripe Standard, direct charges)
//
// Payments run on **Stripe Connect Standard accounts** using **direct charges**:
// the charge is created on _your_ connected account and **you are the merchant of
// record**. That means the money, the payout schedule, the customer relationship,
// and the compliance surface are all yours — Linq orchestrates the request and the
// checkout but is never in the funds flow.
//
// **Refunds, disputes, and chargebacks are handled by you, in your own Stripe
// Dashboard.** Because charges settle directly to your account, Linq has no
// custody of the funds and cannot issue refunds or contest disputes on your behalf
// — and there is no refund/dispute endpoint in this API by design. Use the Stripe
// Dashboard (or the Stripe API on your own account) for the money lifecycle after
// a payment succeeds.
//
// ## Getting set up
//
// Open **Agent Pay** in your Linq dashboard
// (`https://zero.linqapp.com/organization/payments`), click **Connect Stripe**,
// and complete Stripe's onboarding (business details + a bank account). When your
// account reaches `charges_enabled`, request creation unlocks; until you connect
// Stripe, `POST /v3/payment_requests` returns `403`. You can keep collecting even
// while Stripe finishes background verification.
//
// ## Subscriptions
//
// Set `mode: subscription` on `POST /v3/payment_requests` to start an
// **auto-renewing subscription** instead of a one-time charge. Instead of an
// amount, you pass a `price_id` — an active **recurring Price** on your connected
// Stripe account (create one in your Stripe Dashboard under Product catalog; if
// you sell through Stripe Payment Links today, reuse the price your link is built
// from). The recipient pays the first invoice at the same checkout, and their
// payment method is saved to the subscription for automatic renewals.
//
// The division of labor is deliberate: **Linq handles the first payment, your
// Stripe account handles the rest.** The request reaches `succeeded` when the
// first invoice is paid; from then on the subscription lives entirely on your
// connected account. The response's `stripe` object gives you the join keys —
// `customer_id` and `subscription_id` — so renewals, plan changes, dunning, and
// cancellation are managed with your own Stripe Dashboard/API and your own Stripe
// webhooks. Your `metadata` is stamped on the Customer and Subscription, so
// correlating in either direction is trivial. There are no renewal webhooks from
// Linq by design.
//
// ### Discounts
//
// Pass a `discount` with a **coupon** or **promotion code** from your connected
// Stripe account to apply it to the subscription. Create either in your Stripe
// Dashboard under Product catalog → Coupons; Linq only forwards the id.
//
// ```json
//
//	{
//	  "mode": "subscription",
//	  "price_id": "price_1QAbCdEfGhIjKlMn",
//	  "discount": {
//	    "coupon": "7fKCMvBh",
//	    "label": "50% OFF FIRST MONTH"
//	  }
//	}
//
// ```
//
// Stripe applies the coupon and prices the first invoice; the `amount` we return
// is that invoice's amount due, so a `$50.00/month` price with a
// 50%-off-first-month coupon comes back as `2500` and the recipient is charged
// **$25.00** at checkout. A coupon that covers the whole first invoice returns
// `amount: 0`; checkout shows $0.00 and collects the card for the renewal rather
// than charging now. Renewals bill at the full price automatically — how long a
// discount lasts is the coupon's `duration`, enforced by Stripe on your account,
// and Linq never re-prices anything.
//
// Use `promotion_code` instead of `coupon` to apply a promotion code by id
// (`promo_...`, not the customer-facing code string); pass one or the other, never
// both.
//
// `label` is the customer-facing promotion name displayed at checkout instead of
// the coupon or promotion code ID. The label is displayed exactly as provided, so
// include important terms such as "FIRST MONTH" or "FIRST 3 MONTHS" when
// applicable. These terms are not displayed elsewhere on the checkout screen.
//
// If omitted, Stripe uses the coupon's name as the promotion label.
//
// ### Free trials
//
// Add `trial_period_days` (or a fixed `trial_end` timestamp) to start the
// subscription with a free trial. The checkout still collects the recipient's
// payment method — the pay sheet shows "$0 due today" with the first charge date —
// and saves it to the subscription; Stripe bills it automatically when the trial
// ends. The request reaches `succeeded` when the card is collected, and the
// response carries `trial_end`. If the trial would end without a payment method on
// file, the subscription cancels rather than generating unpayable invoices. Trial
// lifecycle after checkout (extending, ending early) is managed in your own Stripe
// account via `stripe.subscription_id`.
//
// A subscription request you cancel (or that expires unpaid) cancels the
// incomplete Stripe subscription — nothing lingers on your account.
//
// ## Pre-created customers
//
// By default each request stands alone: payment mode attaches no Customer, and
// subscription mode creates a fresh one. If you already manage Customers on your
// connected account, pass their id as `customer_id` (`cus_...`) on create — in
// payment mode the charge lands on that customer's payment history, and in
// subscription mode the subscription is created on them instead of on a new
// Customer. The id must reference an existing, non-deleted customer on your
// connected account or the request fails with `400`. We never modify a customer
// you pass — no metadata is stamped on it.
//
// ## Sending the link
//
// Deliver the `checkout_url` as a **`link` message part** via
// `POST /v3/chats/{chatId}/messages` — it renders as a rich card with your
// branding (title, amount, image) instead of a bare URL, which converts far
// better. A `link` part must be the only part in the message. See
// [Rich Link Previews](/channel/imessage/guides/messaging/sending-messages).
//
// On a supported iPhone the link opens an **Apple Pay App Clip** — a native,
// no-install checkout sheet. Everywhere else (Android, desktop, iPhones without
// the App Clip yet) the same URL opens the web checkout, so the link always works.
// The App Clip experience for your payment links is registered automatically by
// Linq and refreshed whenever you update your payments branding; a newly
// registered experience can take up to ~24 hours to activate on Apple's side,
// during which links open the web checkout.
//
// ## Sending it as a card instead
//
// A `link` part is one way to deliver a request. The other is the **`agentpay`
// experience**, which sends the same request as a native card in Linq's iMessage
// app — the amount and reason are drawn in the bubble, and it turns itself into
// "Paid" in place once the payment succeeds, without a second message.
//
// Send it to `POST /v3/chats/{chatId}/messages`:
//
// ```json
//
//	{
//	  "message": {
//	    "experience": {
//	      "name": "agentpay",
//	      "action": "request_payment",
//	      "params": {
//	        "checkout_url": "https://zero.linqapp.com/pay/acme?session=tok_..."
//	      }
//	    }
//	  }
//	}
//
// ```
//
// `checkout_url` is the only required field — pass back exactly what
// `POST /v3/payment_requests` returned. **The amount and reason are read from that
// request, never from you**, so the card can never claim a different figure than
// the checkout will charge. Optional `title` and `note` override the copy only.
// The link must be one of your own payment requests; another partner's is
// rejected.
//
// The trade-off against a `link` part: a card is an app card, so it is
// iMessage-only, and recipients without the app see a static version of it. A link
// works everywhere and is what opens the Apple Pay App Clip. Send whichever suits
// the conversation — both settle the same payment request and fire the same
// webhooks.
//
// ## Webhooks
//
// Subscribe to payment lifecycle events to reconcile server-side rather than
// polling: `payment.succeeded`, `payment.canceled`, and `payment.expired`. Each
// event carries the payment request id, amount, currency, and your `metadata`. See
// [Webhooks](/channel/imessage/guides/webhooks).
//
// PaymentRequestService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentRequestService] method instead.
type PaymentRequestService struct {
	Options []option.RequestOption
}

// NewPaymentRequestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentRequestService(opts ...option.RequestOption) (r PaymentRequestService) {
	r = PaymentRequestService{}
	r.Options = opts
	return
}

// Creates a payment request and returns a `checkout_url` the recipient opens to
// pay with Apple Pay or card. Funds settle directly to your connected Stripe
// account. A payment request is independent of any chat; to associate one with a
// chat for your records, store the chat id in `metadata`. Requires your connected
// account to be `charges_enabled` (returns `403` otherwise).
//
// Set `mode: subscription` with a recurring `price_id` from your connected Stripe
// account to start an **auto-renewing subscription** instead of a one-time charge
// — the recipient pays the first invoice at checkout and the response's `stripe`
// object carries the customer and subscription ids for the ongoing lifecycle in
// your own Stripe account. See the _Subscriptions_ section of the tag overview.
//
// In either mode, pass `customer_id` to attach the request to an **existing
// Customer** on your connected account instead of creating a new one — see
// _Pre-created customers_ in the tag overview.
func (r *PaymentRequestService) New(ctx context.Context, params PaymentRequestNewParams, opts ...option.RequestOption) (res *PaymentRequest, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/payment_requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a payment request's status and details.
func (r *PaymentRequestService) Get(ctx context.Context, paymentRequestID string, opts ...option.RequestOption) (res *PaymentRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentRequestID == "" {
		err = errors.New("missing required paymentRequestId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/payment_requests/%s", paymentRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists your payment requests, newest first, for reconciliation. Paginate with
// `limit` + `offset`; `has_more` indicates whether another page exists.
func (r *PaymentRequestService) List(ctx context.Context, query PaymentRequestListParams, opts ...option.RequestOption) (res *PaymentRequestListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/payment_requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Cancels an unpaid payment request: the underlying payment intent is canceled and
// the request moves to `canceled`. A request that is already paid, canceled, or
// expired returns 409.
func (r *PaymentRequestService) Cancel(ctx context.Context, paymentRequestID string, opts ...option.RequestOption) (res *PaymentRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentRequestID == "" {
		err = errors.New("missing required paymentRequestId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/payment_requests/%s/cancel", paymentRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type PaymentRequest struct {
	// Unique identifier of the payment request.
	ID string `json:"id" api:"required" format:"uuid"`
	// What the recipient is charged at checkout, in the currency's minor units. In
	// `subscription` mode this is the first invoice's amount due — all items after any
	// discounts are applied — so a discount that covers the whole invoice returns `0`
	// and checkout shows $0.00.
	Amount int64 `json:"amount" api:"required"`
	// URL the recipient opens to pay:
	// `https://zero.linqapp.com/pay/{slug}?session=...`, where `{slug}` is your
	// partner checkout slug.
	CheckoutURL string    `json:"checkout_url" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Currency    string    `json:"currency" api:"required"`
	// Whether this request collects a one-time charge or starts a subscription.
	//
	// Any of "payment", "subscription".
	Mode   PaymentRequestMode `json:"mode" api:"required"`
	Object string             `json:"object" api:"required"`
	// Lifecycle status of the payment request.
	//
	// Any of "requested", "succeeded", "canceled", "expired".
	Status      PaymentRequestStatus `json:"status" api:"required"`
	Description string               `json:"description"`
	// Subscription mode — the discount applied, as Stripe applied it.
	Discount PaymentRequestDiscount `json:"discount"`
	// When an unpaid request auto-expires.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Subscription mode — how often the subscription renews.
	//
	// Any of "day", "week", "month", "year".
	Interval PaymentRequestInterval `json:"interval"`
	// Subscription mode — intervals per renewal (e.g. `3` + `month` = quarterly).
	IntervalCount int64             `json:"interval_count"`
	Metadata      map[string]string `json:"metadata"`
	// Natural-rail join keys, present when `rail: natural`.
	Natural PaymentRequestNatural `json:"natural"`
	// When the request was paid. Absent until it succeeds.
	PaidAt time.Time `json:"paid_at" format:"date-time"`
	// Subscription mode — the recurring price this request subscribes to.
	PriceID string `json:"price_id"`
	// Subscription mode — units of the price subscribed to.
	Quantity int64 `json:"quantity"`
	// The rail this request settled on.
	//
	// Any of "stripe", "natural".
	Rail PaymentRequestRail `json:"rail"`
	// Ids of the Stripe objects created **on your connected account** — your join keys
	// into your own Stripe Dashboard, webhooks, and API. After a subscription's first
	// payment succeeds, its ongoing lifecycle (renewals, plan changes, cancellation)
	// is managed in your Stripe account using `subscription_id`.
	Stripe PaymentRequestStripe `json:"stripe"`
	// Subscription mode — when the free trial ends and the first charge happens.
	// Present only on trial requests; `paid_at`/`succeeded` mean the payment method
	// was collected (no funds move until this time).
	TrialEnd  time.Time `json:"trial_end" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		CheckoutURL   respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Mode          respjson.Field
		Object        respjson.Field
		Status        respjson.Field
		Description   respjson.Field
		Discount      respjson.Field
		ExpiresAt     respjson.Field
		Interval      respjson.Field
		IntervalCount respjson.Field
		Metadata      respjson.Field
		Natural       respjson.Field
		PaidAt        respjson.Field
		PriceID       respjson.Field
		Quantity      respjson.Field
		Rail          respjson.Field
		Stripe        respjson.Field
		TrialEnd      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentRequest) RawJSON() string { return r.JSON.raw }
func (r *PaymentRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this request collects a one-time charge or starts a subscription.
type PaymentRequestMode string

const (
	PaymentRequestModePayment      PaymentRequestMode = "payment"
	PaymentRequestModeSubscription PaymentRequestMode = "subscription"
)

// Lifecycle status of the payment request.
type PaymentRequestStatus string

const (
	PaymentRequestStatusRequested PaymentRequestStatus = "requested"
	PaymentRequestStatusSucceeded PaymentRequestStatus = "succeeded"
	PaymentRequestStatusCanceled  PaymentRequestStatus = "canceled"
	PaymentRequestStatusExpired   PaymentRequestStatus = "expired"
)

// Subscription mode — the discount applied, as Stripe applied it.
type PaymentRequestDiscount struct {
	// The ID of the coupon applied.
	Coupon string `json:"coupon"`
	// The customer-facing discount description shown at checkout.
	Label string `json:"label"`
	// The ID of the promotion code applied, if you passed one.
	PromotionCode string `json:"promotion_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coupon        respjson.Field
		Label         respjson.Field
		PromotionCode respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentRequestDiscount) RawJSON() string { return r.JSON.raw }
func (r *PaymentRequestDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription mode — how often the subscription renews.
type PaymentRequestInterval string

const (
	PaymentRequestIntervalDay   PaymentRequestInterval = "day"
	PaymentRequestIntervalWeek  PaymentRequestInterval = "week"
	PaymentRequestIntervalMonth PaymentRequestInterval = "month"
	PaymentRequestIntervalYear  PaymentRequestInterval = "year"
)

// Natural-rail join keys, present when `rail: natural`.
type PaymentRequestNatural struct {
	// The Natural payment request (`prq_...`).
	PaymentRequestID string `json:"payment_request_id"`
	// The settled transaction (`txn_...`).
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentRequestID respjson.Field
		TransactionID    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentRequestNatural) RawJSON() string { return r.JSON.raw }
func (r *PaymentRequestNatural) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The rail this request settled on.
type PaymentRequestRail string

const (
	PaymentRequestRailStripe  PaymentRequestRail = "stripe"
	PaymentRequestRailNatural PaymentRequestRail = "natural"
)

// Ids of the Stripe objects created **on your connected account** — your join keys
// into your own Stripe Dashboard, webhooks, and API. After a subscription's first
// payment succeeds, its ongoing lifecycle (renewals, plan changes, cancellation)
// is managed in your Stripe account using `subscription_id`.
type PaymentRequestStripe struct {
	// The Customer this request is attached to (`cus_...`). Always set in subscription
	// mode (created for you unless you passed `customer_id`); set in payment mode only
	// when you passed one.
	CustomerID string `json:"customer_id"`
	// The PaymentIntent collected at checkout (`pi_...`).
	PaymentIntentID string `json:"payment_intent_id"`
	// Subscription mode — the Subscription (`sub_...`).
	SubscriptionID string `json:"subscription_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID      respjson.Field
		PaymentIntentID respjson.Field
		SubscriptionID  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentRequestStripe) RawJSON() string { return r.JSON.raw }
func (r *PaymentRequestStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentRequestListResponse struct {
	Data []PaymentRequest `json:"data" api:"required"`
	// Whether more results exist beyond this page.
	HasMore bool `json:"has_more" api:"required"`
	// Any of "list".
	Object PaymentRequestListResponseObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentRequestListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentRequestListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentRequestListResponseObject string

const (
	PaymentRequestListResponseObjectList PaymentRequestListResponseObject = "list"
)

type PaymentRequestNewParams struct {
	// Amount to charge, in the currency's minor units (e.g. cents). Must be at least
	// the payment provider's minimum (50 for `usd`). Required in `payment` mode; must
	// be omitted in `subscription` mode (the amount comes from the price).
	Amount param.Opt[int64] `json:"amount,omitzero"`
	// Three-letter ISO 4217 currency code. Only `usd` is currently supported. Required
	// in `payment` mode; must be omitted in `subscription` mode (the currency comes
	// from the price).
	Currency param.Opt[string] `json:"currency,omitzero"`
	// Optional id of an **existing Customer** on your connected Stripe account
	// (`cus_...`) to attach this request to, instead of a new Customer being created.
	// In `payment` mode the charge lands on that customer's payment history; in
	// `subscription` mode the subscription is created on them. The customer must exist
	// (and not be deleted) on your connected account.
	CustomerID param.Opt[string] `json:"customer_id,omitzero"`
	// Optional description shown to the recipient at checkout.
	Description param.Opt[string] `json:"description,omitzero"`
	// Required for `rail: natural`. The line the request is sent from, in E.164
	// format. Must be a phone number your organization owns.
	From param.Opt[string] `json:"from,omitzero"`
	// Required for `rail: natural`. The payer to bill, in E.164 format.
	PayerHandle param.Opt[string] `json:"payer_handle,omitzero"`
	// Subscription mode only (required there): id of an **active recurring Price** on
	// your connected Stripe account (`price_...`). If you sell through Stripe Payment
	// Links today, pass the same price the link was built from to get the native
	// iMessage checkout for it.
	PriceID param.Opt[string] `json:"price_id,omitzero"`
	// Subscription mode only — units of the price to subscribe to.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// Subscription mode only — end the free trial at a fixed timestamp (must be in the
	// future) instead of a day count. Mutually exclusive with `trial_period_days`.
	TrialEnd param.Opt[time.Time] `json:"trial_end,omitzero" format:"date-time"`
	// Subscription mode only — start with a free trial of this many days. The
	// recipient's card is still collected at checkout (Apple Pay or card), saved to
	// the subscription, and first charged when the trial ends. Mutually exclusive with
	// `trial_end`.
	TrialPeriodDays param.Opt[int64]  `json:"trial_period_days,omitzero"`
	IdempotencyKey  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	// Subscription mode only. The coupon or promotion code to apply to this
	// subscription payment. Currently, only accept one coupon or one promo code.
	Discount PaymentRequestNewParamsDiscount `json:"discount,omitzero"`
	// Optional key/value metadata (up to 49 keys) echoed back on retrieval and on
	// `payment.*` webhooks, and stamped on the Stripe objects we create on your
	// connected account (the PaymentIntent, and in subscription mode the Subscription
	// and any Customer created for you — a customer you pass via `customer_id` is
	// never modified) — use it to correlate a request with your own records (e.g. a
	// chat id). Keys starting with `linq_` are reserved.
	Metadata map[string]string `json:"metadata,omitzero"`
	// `payment` (default) collects a one-time charge for `amount` + `currency`.
	// `subscription` starts an auto-renewing subscription from a recurring `price_id`
	// on your connected Stripe account: the recipient pays the first invoice at
	// checkout and Stripe renews it automatically from then on.
	//
	// Any of "payment", "subscription".
	Mode PaymentRequestNewParamsMode `json:"mode,omitzero"`
	// Payment rail. `stripe` (default) is the direct-charge flow that settles to your
	// connected Stripe account. `natural` collects through the Natural custodial
	// wallet; it requires `from` + `payer_handle` and that your organization has
	// completed Natural merchant onboarding.
	//
	// Any of "stripe", "natural".
	Rail PaymentRequestNewParamsRail `json:"rail,omitzero"`
	paramObj
}

func (r PaymentRequestNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentRequestNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentRequestNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription mode only. The coupon or promotion code to apply to this
// subscription payment. Currently, only accept one coupon or one promo code.
type PaymentRequestNewParamsDiscount struct {
	// The ID of the coupon to apply to this subscription.
	Coupon param.Opt[string] `json:"coupon,omitzero"`
	// Name of the coupon/promo code displayed to customers.
	Label param.Opt[string] `json:"label,omitzero"`
	// The ID of a promotion code to apply to this subscription.
	PromotionCode param.Opt[string] `json:"promotion_code,omitzero"`
	paramObj
}

func (r PaymentRequestNewParamsDiscount) MarshalJSON() (data []byte, err error) {
	type shadow PaymentRequestNewParamsDiscount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentRequestNewParamsDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `payment` (default) collects a one-time charge for `amount` + `currency`.
// `subscription` starts an auto-renewing subscription from a recurring `price_id`
// on your connected Stripe account: the recipient pays the first invoice at
// checkout and Stripe renews it automatically from then on.
type PaymentRequestNewParamsMode string

const (
	PaymentRequestNewParamsModePayment      PaymentRequestNewParamsMode = "payment"
	PaymentRequestNewParamsModeSubscription PaymentRequestNewParamsMode = "subscription"
)

// Payment rail. `stripe` (default) is the direct-charge flow that settles to your
// connected Stripe account. `natural` collects through the Natural custodial
// wallet; it requires `from` + `payer_handle` and that your organization has
// completed Natural merchant onboarding.
type PaymentRequestNewParamsRail string

const (
	PaymentRequestNewParamsRailStripe  PaymentRequestNewParamsRail = "stripe"
	PaymentRequestNewParamsRailNatural PaymentRequestNewParamsRail = "natural"
)

type PaymentRequestListParams struct {
	// Max results to return (default 20, max 100).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of results to skip.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by lifecycle status.
	//
	// Any of "requested", "authorized", "succeeded", "canceled", "expired",
	// "declined".
	Status PaymentRequestListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaymentRequestListParams]'s query parameters as
// `url.Values`.
func (r PaymentRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by lifecycle status.
type PaymentRequestListParamsStatus string

const (
	PaymentRequestListParamsStatusRequested  PaymentRequestListParamsStatus = "requested"
	PaymentRequestListParamsStatusAuthorized PaymentRequestListParamsStatus = "authorized"
	PaymentRequestListParamsStatusSucceeded  PaymentRequestListParamsStatus = "succeeded"
	PaymentRequestListParamsStatusCanceled   PaymentRequestListParamsStatus = "canceled"
	PaymentRequestListParamsStatusExpired    PaymentRequestListParamsStatus = "expired"
	PaymentRequestListParamsStatusDeclined   PaymentRequestListParamsStatus = "declined"
)
