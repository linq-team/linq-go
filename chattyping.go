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
// **Note:** Group chat typing indicators are not currently supported. Attempting
// to start a typing indicator in a group chat will return a `403` error.
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

// Stop the typing indicator for the chat.
//
// **Note:** Typing indicators are automatically stopped when a message is sent, so
// calling this endpoint after sending a message is unnecessary.
//
// **Note:** Group chat typing indicators are not currently supported. Attempting
// to stop a typing indicator in a group chat will return a `403` error.
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
