// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
)

// A Chat is a conversation thread with one or more participants.
//
// To begin a chat, you must create a Chat with at least one recipient handle.
// Including multiple handles creates a group chat.
//
// When creating a chat, the `from` field specifies which of your authorized phone
// numbers the message originates from. Your authentication token grants access to
// one or more phone numbers, but the `from` field determines the actual sender.
//
// **Handle Format:**
//
//   - Handles can be phone numbers or email addresses
//   - Phone numbers MUST be in E.164 format (starting with +)
//   - Phone format: `+[country code][subscriber number]`
//   - Example phone: `+12223334444` (US), `+442071234567` (UK), `+81312345678`
//     (Japan)
//   - Example email: `user@example.com`
//   - No spaces, dashes, or parentheses in phone numbers
//
// ChatTypingService contains methods and other services that help with interacting
// with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatTypingService] method instead.
type ChatTypingService struct {
	Options []option.RequestOption
}

// NewChatTypingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatTypingService(opts ...option.RequestOption) (r ChatTypingService) {
	r = ChatTypingService{}
	r.Options = opts
	return
}

// Send a typing indicator to show that someone is typing in the chat.
//
// ## Behavior
//
// Typing indicators are best-effort signals that behave as follows:
//
//   - **iMessage chats only:** Typing indicators are only supported for iMessage
//     chats. Requests for RCS or SMS chats are accepted (`204`) but no indicator is
//     delivered.
//
//   - **Send a message first for reliable delivery:** Typing indicators are
//     best-effort. If you have not sent a message in this chat recently (roughly the
//     **last 5 minutes**), a typing indicator may not reach the recipient — the
//     request is still accepted (`204`), but delivery is not deterministic. Once you
//     have sent a message in the chat, typing indicators reliably reach the
//     recipient.
//
//   - **No delivery guarantee:** Even for active chats, a `204` response only
//     indicates the request was accepted for processing.
//
//   - **Direct and group chats:** Typing indicators work in both direct and group
//     chats.
//
// ## Duration & keeping it visible
//
//   - A single call shows the indicator for about **85–90 seconds**, then it clears
//     automatically.
//
//   - To keep it visible longer, call this endpoint again every **60 seconds**. Each
//     call refreshes the indicator so it stays visible continuously.
//
// - Sending a message clears the indicator.
//
// - To resume typing after sending a message, call this endpoint again.
//
// - Incoming messages do not affect the indicator.
//
// ## Recipient re-opening the chat
//
// If the recipient brings their messaging app to the foreground while the chat has
// an unread message, their device clears any showing typing indicator. Calling
// this endpoint again on its own may not bring it back. To make it reappear,
// either send a message, or call `DELETE /v3/chats/{chatId}/typing` (stop) and
// then call start typing again.
//
// ## Recommended usage
//
// Call this endpoint when composing begins, call it again every 60 seconds while
// composing, and send the message to clear the indicator. To clear the indicator
// without sending a message, call `DELETE /v3/chats/{chatId}/typing`.
func (r *ChatTypingService) Start(ctx context.Context, chatID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return err
	}
	path := fmt.Sprintf("v3/chats/%s/typing", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Immediately clears the typing indicator for the chat, without sending a message.
//
// The typing indicator also clears automatically when you send a message, or about
// 85–90 seconds after the last `POST /v3/chats/{chatId}/typing` (start typing)
// request.
//
// See the start typing endpoint (`POST /v3/chats/{chatId}/typing`) above for
// behavior details.
//
// **Note:** Works in both direct and group chats.
func (r *ChatTypingService) Stop(ctx context.Context, chatID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return err
	}
	path := fmt.Sprintf("v3/chats/%s/typing", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}
