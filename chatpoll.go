// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
	"github.com/linq-team/linq-go/shared"
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
// ChatPollService contains methods and other services that help with interacting
// with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatPollService] method instead.
type ChatPollService struct {
	Options []option.RequestOption
}

// NewChatPollService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatPollService(opts ...option.RequestOption) (r ChatPollService) {
	r = ChatPollService{}
	r.Options = opts
	return
}

// Create an iMessage poll in an existing chat and send it. Polls are
// iMessage-only.
//
// The chat must already exist — **a poll cannot be the first message of a new
// chat** (use `POST /v3/chats` for that). Options are **add-only and immutable**:
// you can add options later via `POST /v3/messages/{messageId}/poll/options`, but
// never edit or remove them.
func (r *ChatPollService) New(ctx context.Context, chatID string, body ChatPollNewParams, opts ...option.RequestOption) (res *PollEnvelope, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/chats/%s/polls", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Poll content — options and the aggregate voter count.
type Poll struct {
	Options []PollOption `json:"options" api:"required"`
	// Distinct participants across the whole poll (a voter picking two options counts
	// once).
	TotalVoters int64 `json:"total_voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		TotalVoters respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Poll) RawJSON() string { return r.JSON.raw }
func (r *Poll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollOption struct {
	CanBeEdited bool `json:"can_be_edited" api:"required"`
	// The participant who added this option (poll creator for the initial options;
	// whoever added later ones).
	CreatorHandle shared.ChatHandle `json:"creator_handle" api:"required"`
	OptionID      string            `json:"option_id" api:"required" format:"uuid"`
	Text          string            `json:"text" api:"required"`
	// Participants who voted for this option (vote_count = voters.length).
	Voters []PollOptionVoter `json:"voters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanBeEdited   respjson.Field
		CreatorHandle respjson.Field
		OptionID      respjson.Field
		Text          respjson.Field
		Voters        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollOption) RawJSON() string { return r.JSON.raw }
func (r *PollOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PollOptionVoter struct {
	Handle  string    `json:"handle" api:"required"`
	VotedAt time.Time `json:"voted_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		VotedAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollOptionVoter) RawJSON() string { return r.JSON.raw }
func (r *PollOptionVoter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message-level envelope returned by every poll endpoint.
type PollEnvelope struct {
	ChatID    string    `json:"chat_id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The poll-definition message's ID — reference this poll by it.
	MessageID string `json:"message_id" api:"required" format:"uuid"`
	// Poll content — options and the aggregate voter count.
	Poll Poll `json:"poll" api:"required"`
	// Tapbacks/stickers on the whole poll (message part 0).
	Reactions []shared.Reaction `json:"reactions" api:"required"`
	UpdatedAt time.Time         `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		CreatedAt   respjson.Field
		MessageID   respjson.Field
		Poll        respjson.Field
		Reactions   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PollEnvelope) RawJSON() string { return r.JSON.raw }
func (r *PollEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatPollNewParams struct {
	// Poll content to create. A poll needs at least two options. Options are add-only
	// and immutable — there is no title/question (send that as a normal text message).
	Poll ChatPollNewParamsPoll `json:"poll,omitzero" api:"required"`
	paramObj
}

func (r ChatPollNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatPollNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatPollNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Poll content to create. A poll needs at least two options. Options are add-only
// and immutable — there is no title/question (send that as a normal text message).
//
// The property Options is required.
type ChatPollNewParamsPoll struct {
	Options []ChatPollNewParamsPollOption `json:"options,omitzero" api:"required"`
	// Optional key to deduplicate the poll creation.
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	paramObj
}

func (r ChatPollNewParamsPoll) MarshalJSON() (data []byte, err error) {
	type shadow ChatPollNewParamsPoll
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatPollNewParamsPoll) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Text is required.
type ChatPollNewParamsPollOption struct {
	Text string `json:"text" api:"required"`
	paramObj
}

func (r ChatPollNewParamsPollOption) MarshalJSON() (data []byte, err error) {
	type shadow ChatPollNewParamsPollOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatPollNewParamsPollOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
