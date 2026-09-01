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
// covered phone numbers is given a **retention window configured for your
// account** — after that window the platform permanently deletes the message from
// Linq storage. There is no per-message flag; ephemerality is applied
// automatically based on your configuration.
//
// The window can be set anywhere from **60 minutes to 24 hours**, and defaults to
// **24 hours**. Ask your Linq support contact to configure a shorter window; it
// cannot be changed through the API.
//
// You can request it at two scopes:
//
// | Scope                | Effect                                                                                                                            |
// | -------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
// | **Partner-wide**     | Every outbound and inbound message on every phone number under your account is retained for your configured window, then deleted. |
// | **Per phone number** | Only the specified phone numbers have their messages auto-deleted. The rest follow the standard message-retention policy.         |
//
// **Behavioral differences vs the standard default:**
//
// | Aspect                  | Standard                                           | Ephemeral                                                                                                                                                                                                                                                                                                                                   |
// | ----------------------- | -------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
// | Retention               | Retained per the standard message-retention policy | **Hard backstop: your configured window** (60 minutes – 24 hours, default 24 hours) from when the message is created                                                                                                                                                                                                                        |
// | After expiry            | Message stays retrievable                          | Message is permanently deleted — `GET /v3/messages/{messageId}` returns `404` and it no longer appears in `GET /v3/chats/{chatId}/messages`                                                                                                                                                                                                 |
// | Content on expiry       | N/A                                                | Text, formatting, and attachment references are scrubbed; the message is gone, not blanked out                                                                                                                                                                                                                                              |
// | Attachments             | Retained                                           | Media sent on the **ephemeral attachments tier** is removed on its own storage backstop — within roughly 24–48 hours of upload — independently of the message window, so it can outlast a window shorter than a day. Attachments on the persistent tier (including pre-uploads via `POST /v3/attachments`) are kept until you `DELETE` them |
// | Cross-partner isolation | Enforced                                           | Enforced                                                                                                                                                                                                                                                                                                                                    |
//
// **How the retention window works:**
//
//   - The window runs from **message creation** (`created_at`). It is configured for
//     your account (60 minutes – 24 hours, default 24 hours) and cannot be set per
//     message.
//   - Attachment media follows its own storage backstop rather than the message
//     window — see the Attachments row above.
//   - Expiry is delivery-independent — the clock starts when the message is created,
//     not when it is delivered or read.
//   - **Deletion happens shortly _after_ the window, not exactly at it.** A
//     background sweep runs every ~5 minutes, so a message typically stops being
//     retrievable within about 5 minutes of its expiry, and longer while a backlog
//     is being worked through. Treat the window as the guaranteed _minimum_
//     retention, never as an exact deletion time or an upper bound.
//
// **What you observe:**
//
//   - **No expiry timestamp is exposed.** API responses and webhook payloads do not
//     include the deletion time, and they do not report your configured window
//     either — so if you are on a window shorter than 24 hours you cannot derive a
//     message's expiry from the API today. Track the window you agreed with your
//     Linq support contact and compute `created_at + window` yourself.
//   - **No deletion webhook is sent.** There is no `message.deleted` event — a
//     message simply stops being retrievable once its window passes.
//   - **The backstop governs Linq storage.** API retrievability (the `404` behavior
//     above) ends at your configured window. Ephemeral-tier media objects are
//     removed on their own storage backstop — within roughly 24–48 hours of upload —
//     which is independent of the message window and can outlast a window shorter
//     than a day. Removal of the corresponding entries from the sending device
//     happens asynchronously and can complete after the backstop.
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
// no longer return the message once its window passes, persist anything you need
// to keep from the webhook payload at the time it is delivered.
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
