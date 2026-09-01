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

// Phone Numbers represent the phone numbers assigned to your partner account.
//
// Use the list phone numbers endpoint to discover which phone numbers are
// available for sending messages.
//
// When creating chats, listing chats, or sending a voice memo, use one of your
// assigned phone numbers in the `from` field.
//
// **Ineligible numbers.** A number can temporarily lose the ability to deliver
// messages. While it is in that state, requests that would produce new activity on
// it — sending a message, creating a chat, reacting, typing, group actions — are
// rejected with `403` (error code `2027`) before anything is created. Reads keep
// working, so your existing chats, messages, and history stay available. Omit
// `from` on `POST /v3/messages` and we pick an eligible number for you, skipping
// ineligible ones; if none of your assigned numbers are eligible, you get `409`
// (no `from` number was ever chosen, so there's no specific number to blame with a
// `403`).
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

// Returns the audit's status and, once complete, the report. Audits are scoped to
// the line in the URL — an `auditId` started on a different line returns `404`.
func (r *PhoneNumberService) GetReputationAudit(ctx context.Context, auditID string, query PhoneNumberGetReputationAuditParams, opts ...option.RequestOption) (res *ReputationAudit, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.PhoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return nil, err
	}
	if auditID == "" {
		err = errors.New("missing required auditId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/phone_numbers/%s/reputation_audit/%s", url.PathEscape(query.PhoneNumber), url.PathEscape(auditID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Starts an asynchronous reputation audit for a line and returns an `audit_id`.
// Poll the GET endpoint for the result.
//
// Rate limited per line: only one audit may run at a time. Starting one while
// another is still running returns `202` with the running audit's `audit_id`
// rather than an error, so a retried start picks that audit back up instead of
// losing it — poll the id you were given.
//
// Once an audit finishes, a new one can't be started for the same line until a
// cooldown elapses (`429`, with `Retry-After` carrying the wait). Keep the
// `audit_id` from the original `202`: it stays readable on the GET endpoint for 24
// hours, and the cooldown response does not repeat it.
func (r *PhoneNumberService) StartReputationAudit(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *ReputationAuditStarted, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/phone_numbers/%s/reputation_audit", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ReputationActionItem struct {
	Detail string `json:"detail"`
	// Any of "high", "medium", "low".
	ExpectedImpact ReputationActionItemExpectedImpact `json:"expected_impact"`
	// 1 = do first
	Priority int64  `json:"priority"`
	Title    string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail         respjson.Field
		ExpectedImpact respjson.Field
		Priority       respjson.Field
		Title          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationActionItem) RawJSON() string { return r.JSON.raw }
func (r *ReputationActionItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReputationActionItemExpectedImpact string

const (
	ReputationActionItemExpectedImpactHigh   ReputationActionItemExpectedImpact = "high"
	ReputationActionItemExpectedImpactMedium ReputationActionItemExpectedImpact = "medium"
	ReputationActionItemExpectedImpactLow    ReputationActionItemExpectedImpact = "low"
)

type ReputationAudit struct {
	AuditID string `json:"audit_id" api:"required"`
	// `pending` until the report is ready — poll until `complete` or `error`.
	//
	// Any of "pending", "complete", "error".
	Status ReputationAuditStatus `json:"status" api:"required"`
	// Present only when `status` is `error`. Short, generic reason safe to display.
	Error string `json:"error"`
	// When the report was generated; signals reflect the line at this moment.
	GeneratedAt time.Time `json:"generated_at" format:"date-time"`
	// The line audited, E.164.
	Phone string `json:"phone"`
	// Present only when `status` is `complete`.
	Report ReputationReport `json:"report"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuditID     respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		GeneratedAt respjson.Field
		Phone       respjson.Field
		Report      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationAudit) RawJSON() string { return r.JSON.raw }
func (r *ReputationAudit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `pending` until the report is ready — poll until `complete` or `error`.
type ReputationAuditStatus string

const (
	ReputationAuditStatusPending  ReputationAuditStatus = "pending"
	ReputationAuditStatusComplete ReputationAuditStatus = "complete"
	ReputationAuditStatusError    ReputationAuditStatus = "error"
)

type ReputationAuditStarted struct {
	// Identifier for this audit. Poll
	// `GET /v3/phone_numbers/{phoneNumber}/reputation_audit/{auditId}` until `status`
	// is `complete` or `error`.
	AuditID string `json:"audit_id" api:"required"`
	// A newly started audit is `pending`.
	//
	// Any of "pending", "complete", "error".
	Status ReputationAuditStartedStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuditID     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationAuditStarted) RawJSON() string { return r.JSON.raw }
func (r *ReputationAuditStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A newly started audit is `pending`.
type ReputationAuditStartedStatus string

const (
	ReputationAuditStartedStatusPending  ReputationAuditStartedStatus = "pending"
	ReputationAuditStartedStatusComplete ReputationAuditStartedStatus = "complete"
	ReputationAuditStartedStatusError    ReputationAuditStartedStatus = "error"
)

type ReputationDriver struct {
	// Stable driver-category identifier — what is dragging the line, or one of its
	// conversations, down.
	//
	//   - `low_engagement` — The conversation is one-sided: several messages sent, few
	//     or no replies back. Pause or rework outreach where recipients are not
	//     replying, and lead with messages that invite a response. Conversation-level:
	//     it appears on `evidence.unhealthy_chats[].driver_keys`, never in `drivers`.
	//   - `overall_conversation_health` — A large share of the line's active
	//     conversations are trending unhealthy. Fix the unhealthy conversations first —
	//     review their content and timing, and whether recipients are engaging.
	//   - `volume_spike` — The line's daily sending volume jumped far above its own
	//     normal level while few recipients were replying, or exceeded the recommended
	//     daily volume for a single line. Ramp volume gradually instead of spiking,
	//     prioritize people who have already engaged with you, and spread sustained high
	//     volume across additional lines.
	//   - `new_conversation_rate` — The line is starting too many brand-new
	//     conversations in a single day. Spread new conversations out over time instead
	//     of starting many at once.
	//   - `opt_out_handling` — Recipients asked this line to stop. Honor every stop
	//     request immediately: send nothing further to that recipient unless they opt
	//     back in. Every send to them is rejected with `403` (error code `2024`),
	//     including a final courtesy message — to send one telling them they can reply
	//     to resume, set `override_optout: true` on that single request.
	//   - `flagged` — The line is currently restricted and its messages may not be
	//     reaching recipients. Move active traffic to a healthy line now, and let this
	//     one recover before sending more.
	//   - `other` — Fallback for a signal without dedicated partner copy.
	//
	// Any of "low_engagement", "overall_conversation_health", "volume_spike",
	// "new_conversation_rate", "opt_out_handling", "flagged", "other".
	Key ReputationDriverKey `json:"key"`
	// A specific observed figure when available; otherwise a short qualitative note.
	Metric string `json:"metric"`
	// One plain-English sentence.
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Metric      respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationDriver) RawJSON() string { return r.JSON.raw }
func (r *ReputationDriver) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable driver-category identifier — what is dragging the line, or one of its
// conversations, down.
//
//   - `low_engagement` — The conversation is one-sided: several messages sent, few
//     or no replies back. Pause or rework outreach where recipients are not
//     replying, and lead with messages that invite a response. Conversation-level:
//     it appears on `evidence.unhealthy_chats[].driver_keys`, never in `drivers`.
//   - `overall_conversation_health` — A large share of the line's active
//     conversations are trending unhealthy. Fix the unhealthy conversations first —
//     review their content and timing, and whether recipients are engaging.
//   - `volume_spike` — The line's daily sending volume jumped far above its own
//     normal level while few recipients were replying, or exceeded the recommended
//     daily volume for a single line. Ramp volume gradually instead of spiking,
//     prioritize people who have already engaged with you, and spread sustained high
//     volume across additional lines.
//   - `new_conversation_rate` — The line is starting too many brand-new
//     conversations in a single day. Spread new conversations out over time instead
//     of starting many at once.
//   - `opt_out_handling` — Recipients asked this line to stop. Honor every stop
//     request immediately: send nothing further to that recipient unless they opt
//     back in. Every send to them is rejected with `403` (error code `2024`),
//     including a final courtesy message — to send one telling them they can reply
//     to resume, set `override_optout: true` on that single request.
//   - `flagged` — The line is currently restricted and its messages may not be
//     reaching recipients. Move active traffic to a healthy line now, and let this
//     one recover before sending more.
//   - `other` — Fallback for a signal without dedicated partner copy.
type ReputationDriverKey string

const (
	ReputationDriverKeyLowEngagement             ReputationDriverKey = "low_engagement"
	ReputationDriverKeyOverallConversationHealth ReputationDriverKey = "overall_conversation_health"
	ReputationDriverKeyVolumeSpike               ReputationDriverKey = "volume_spike"
	ReputationDriverKeyNewConversationRate       ReputationDriverKey = "new_conversation_rate"
	ReputationDriverKeyOptOutHandling            ReputationDriverKey = "opt_out_handling"
	ReputationDriverKeyFlagged                   ReputationDriverKey = "flagged"
	ReputationDriverKeyOther                     ReputationDriverKey = "other"
)

// The specific conversations behind the drivers, so partners can verify every
// claim against their own send logs. Each `chat_id` can be fetched via
// `GET /v3/chats/{chatId}` — its current health appears there.
type ReputationEvidence struct {
	// Worst first — most messages sent after the stop request; honor these
	// immediately.
	OptOutChats []ReputationOptOutChat `json:"opt_out_chats"`
	// Up to 15, worst first.
	UnhealthyChats []ReputationUnhealthyChat `json:"unhealthy_chats"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OptOutChats    respjson.Field
		UnhealthyChats respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationEvidence) RawJSON() string { return r.JSON.raw }
func (r *ReputationEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReputationOptOutChat struct {
	ChatID string `json:"chat_id"`
	// Outbound messages sent after the recipient asked to stop.
	MessagesAfterStop int64 `json:"messages_after_stop"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID            respjson.Field
		MessagesAfterStop respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationOptOutChat) RawJSON() string { return r.JSON.raw }
func (r *ReputationOptOutChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReputationReport struct {
	// Ordered by `priority`; 1 = do first.
	ActionItems []ReputationActionItem `json:"action_items"`
	// Ranked, highest impact first.
	Drivers []ReputationDriver `json:"drivers"`
	// The specific conversations behind the drivers, so partners can verify every
	// claim against their own send logs. Each `chat_id` can be fetched via
	// `GET /v3/chats/{chatId}` — its current health appears there.
	Evidence ReputationEvidence `json:"evidence"`
	// The `key` of the most important driver. Empty string when the line has nothing
	// to act on — the report then carries a single reassurance action item. Its values
	// are the `ReputationDriverKey` vocabulary — see that schema for what each means
	// and what to do about it.
	PrimaryDriver string `json:"primary_driver"`
	// Current reputation of this phone line.
	//
	//   - `HEALTHY` — The line is in good standing. Send normally.
	//   - `AT_RISK` — Warning signs on the line: engagement is low across many of its
	//     conversations, or it's starting too many brand-new conversations in a single
	//     day — and a spike in send volume can add to either. Slow the line's send pace,
	//     avoid opening many new conversations at once, and review your messaging
	//     patterns.
	//   - `CRITICAL` — Strong signals that messages from this line aren't landing well.
	//     Pause outbound on the line until it recovers.
	//
	// Defaults to `HEALTHY` for lines that have not yet been scored.
	//
	// Any of "HEALTHY", "AT_RISK", "CRITICAL".
	Severity ReputationReportSeverity `json:"severity"`
	// Deterministic markdown rendering of this report, suitable for feeding directly
	// to automated systems and AI agents as investigation context. Rendered from the
	// structured fields above, which remain the source of truth.
	SummaryMarkdown string `json:"summary_markdown"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionItems     respjson.Field
		Drivers         respjson.Field
		Evidence        respjson.Field
		PrimaryDriver   respjson.Field
		Severity        respjson.Field
		SummaryMarkdown respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationReport) RawJSON() string { return r.JSON.raw }
func (r *ReputationReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current reputation of this phone line.
//
//   - `HEALTHY` — The line is in good standing. Send normally.
//   - `AT_RISK` — Warning signs on the line: engagement is low across many of its
//     conversations, or it's starting too many brand-new conversations in a single
//     day — and a spike in send volume can add to either. Slow the line's send pace,
//     avoid opening many new conversations at once, and review your messaging
//     patterns.
//   - `CRITICAL` — Strong signals that messages from this line aren't landing well.
//     Pause outbound on the line until it recovers.
//
// Defaults to `HEALTHY` for lines that have not yet been scored.
type ReputationReportSeverity string

const (
	ReputationReportSeverityHealthy  ReputationReportSeverity = "HEALTHY"
	ReputationReportSeverityAtRisk   ReputationReportSeverity = "AT_RISK"
	ReputationReportSeverityCritical ReputationReportSeverity = "CRITICAL"
)

type ReputationUnhealthyChat struct {
	ChatID string `json:"chat_id"`
	// What is dragging this conversation down, in the same vocabulary as the report's
	// drivers. Each key's meaning and the fix for it are documented on
	// `ReputationDriverKey`.
	DriverKeys []ReputationDriverKey `json:"driver_keys"`
	// The conversation's current health — the same value `GET /v3/chats/{chatId}`
	// reports for it.
	//
	// Any of "AT_RISK", "CRITICAL", "OPTED_OUT".
	Status ReputationUnhealthyChatStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		DriverKeys  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationUnhealthyChat) RawJSON() string { return r.JSON.raw }
func (r *ReputationUnhealthyChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The conversation's current health — the same value `GET /v3/chats/{chatId}`
// reports for it.
type ReputationUnhealthyChatStatus string

const (
	ReputationUnhealthyChatStatusAtRisk   ReputationUnhealthyChatStatus = "AT_RISK"
	ReputationUnhealthyChatStatusCritical ReputationUnhealthyChatStatus = "CRITICAL"
	ReputationUnhealthyChatStatusOptedOut ReputationUnhealthyChatStatus = "OPTED_OUT"
)

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
	// Phone number in E.164 format
	PhoneNumber string `json:"phone_number" api:"required"`
	// **[BETA]** Current reputation for a phone line. Always present — lines start at
	// `HEALTHY` and may shift based on aggregate engagement and delivery signals
	// across all conversations on the line.
	//
	// Unlike chat health, line reputation does not include `opted_out` — opt-out
	// applies to individual recipients, not the whole line.
	//
	// See the
	// [Phone Reputation guide](/channel/imessage/guides/phone-numbers/phone-reputation)
	// for what each status means and how to react.
	Reputation PhoneNumberListResponsePhoneNumberReputation `json:"reputation" api:"required"`
	// The forwarding number associated with this phone number, in E.164 format. Null
	// when no forwarding number is configured.
	ForwardingNumber string `json:"forwarding_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
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
// See the
// [Phone Reputation guide](/channel/imessage/guides/phone-numbers/phone-reputation)
// for what each status means and how to react.
type PhoneNumberListResponsePhoneNumberReputation struct {
	// Deep-link to the relevant section of the Phone Reputation guide for this status.
	DocURL string `json:"doc_url" api:"required" format:"uri"`
	// Current reputation of this phone line.
	//
	//   - `HEALTHY` — The line is in good standing. Send normally.
	//   - `AT_RISK` — Warning signs on the line: engagement is low across many of its
	//     conversations, or it's starting too many brand-new conversations in a single
	//     day — and a spike in send volume can add to either. Slow the line's send pace,
	//     avoid opening many new conversations at once, and review your messaging
	//     patterns.
	//   - `CRITICAL` — Strong signals that messages from this line aren't landing well.
	//     Pause outbound on the line until it recovers.
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

type PhoneNumberGetReputationAuditParams struct {
	PhoneNumber string `path:"phoneNumber" api:"required" json:"-"`
	paramObj
}
