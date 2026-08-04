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
)

// Messages are individual communications within a chat thread.
//
// Messages can include text, media attachments, rich link previews, special
// effects (like confetti or fireworks), and reactions. All messages are associated
// with a specific chat and sent from a phone number you own.
//
// Messages support delivery status tracking, read receipts, and editing
// capabilities.
//
// ## Rich Link Previews
//
// Send a URL as a `link` part to deliver it with a rich preview card showing the
// page's title, description, and image (when available). A `link` part must be the
// **only** part in the message — it cannot be combined with text or media parts.
// To send a URL without a preview card, include it in a `text` part instead.
//
// **Limitations:**
//
// - A `link` part cannot be combined with other parts in the same message.
// - Maximum URL length: 2,048 characters.
//
// ## Ephemeral Messages (Privacy Tier)
//
// For regulated or sensitive conversations, opt in to the **ephemeral messages**
// tier by contacting your Linq support contact. When enabled, every message on the
// covered phone numbers is automatically given a fixed **24-hour retention
// window** — after that window the platform permanently deletes the message from
// Linq storage. There is no per-message flag; ephemerality is applied
// automatically based on your configuration.
//
// You can request it at two scopes:
//
// | Scope                | Effect                                                                                                                    |
// | -------------------- | ------------------------------------------------------------------------------------------------------------------------- |
// | **Partner-wide**     | Every outbound and inbound message on every phone number under your account is retained for 24 hours, then deleted.       |
// | **Per phone number** | Only the specified phone numbers have their messages auto-deleted. The rest follow the standard message-retention policy. |
//
// **Behavioral differences vs the standard default:**
//
// | Aspect                  | Standard                                           | Ephemeral                                                                                                                                   |
// | ----------------------- | -------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
// | Retention               | Retained per the standard message-retention policy | **Hard backstop: 24 hours** from when the message is created                                                                                |
// | After expiry            | Message stays retrievable                          | Message is permanently deleted — `GET /v3/messages/{messageId}` returns `404` and it no longer appears in `GET /v3/chats/{chatId}/messages` |
// | Content on expiry       | N/A                                                | Text, formatting, and attachment references are scrubbed; the message is gone, not blanked out                                              |
// | Cross-partner isolation | Enforced                                           | Enforced                                                                                                                                    |
//
// **How the 24-hour window works:**
//
//   - The window is fixed at **24 hours from message creation** (`created_at`) and
//     cannot be configured per message.
//   - It mirrors the ephemeral _attachments_ 1-day backstop, so a message and any
//     media it carries expire together.
//   - Expiry is delivery-independent — the clock starts when the message is created,
//     not when it is delivered or read.
//
// **What you observe:**
//
//   - **No expiry timestamp is exposed.** API responses and webhook payloads do not
//     include the deletion time. If you need it, compute `created_at + 24h`
//     yourself.
//   - **No deletion webhook is sent.** There is no `message.deleted` event — a
//     message simply stops being retrievable once its window passes.
//   - **Delivery is unaffected.** Ephemeral messages send, deliver, and fire the
//     usual `message.sent` / `message.received` and status webhooks exactly like
//     standard messages. Only retention changes.
//
// **When to choose ephemeral:**
//
//   - You have a compliance requirement that the platform must not retain message
//     content beyond a short window.
//   - The conversation is high-sensitivity (PHI, financial, identity verification)
//     and you do not want it sitting in storage long-term.
//   - Your application is the system of record — you capture what you need from the
//     delivery webhook in real time and do not rely on reading message history back
//     from Linq later.
//
// **Important:** ephemeral applies in _both directions_ — messages you send
// **and** messages received by the phone numbers in that scope. Because Linq can
// no longer return the message after 24 hours, persist anything you need to keep
// from the webhook payload at the time it is delivered.
//
// MessagePollService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagePollService] method instead.
type MessagePollService struct {
	Options []option.RequestOption
}

// NewMessagePollService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessagePollService(opts ...option.RequestOption) (r MessagePollService) {
	r = MessagePollService{}
	r.Options = opts
	return
}

// Return a poll's current results — its options, each option's voters, and the
// distinct total number of voters — by the poll-definition message's ID.
func (r *MessagePollService) Get(ctx context.Context, messageID string, opts ...option.RequestOption) (res *PollEnvelope, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/messages/%s/poll", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Add one or more options to an existing poll. Options are **add-only and
// immutable** — you can append options but never edit or remove them (Apple
// constraint). Returns the full poll.
func (r *MessagePollService) AddOptions(ctx context.Context, messageID string, body MessagePollAddOptionsParams, opts ...option.RequestOption) (res *PollEnvelope, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/messages/%s/poll/options", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Add or remove your line's vote on **one** poll option (per-option toggle —
// iMessage polls are toggled one option at a time). Returns the poll reflecting
// the toggle.
func (r *MessagePollService) Vote(ctx context.Context, messageID string, body MessagePollVoteParams, opts ...option.RequestOption) (res *PollEnvelope, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/messages/%s/poll/votes", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MessagePollAddOptionsParams struct {
	Options []MessagePollAddOptionsParamsOption `json:"options,omitzero" api:"required"`
	paramObj
}

func (r MessagePollAddOptionsParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagePollAddOptionsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagePollAddOptionsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Text is required.
type MessagePollAddOptionsParamsOption struct {
	Text string `json:"text" api:"required"`
	paramObj
}

func (r MessagePollAddOptionsParamsOption) MarshalJSON() (data []byte, err error) {
	type shadow MessagePollAddOptionsParamsOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagePollAddOptionsParamsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagePollVoteParams struct {
	// Add or remove your line's vote on the option.
	//
	// Any of "add", "remove".
	Operation MessagePollVoteParamsOperation `json:"operation,omitzero" api:"required"`
	// The option to toggle a vote on.
	OptionID string `json:"option_id" api:"required" format:"uuid"`
	paramObj
}

func (r MessagePollVoteParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagePollVoteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagePollVoteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Add or remove your line's vote on the option.
type MessagePollVoteParamsOperation string

const (
	MessagePollVoteParamsOperationAdd    MessagePollVoteParamsOperation = "add"
	MessagePollVoteParamsOperationRemove MessagePollVoteParamsOperation = "remove"
)
