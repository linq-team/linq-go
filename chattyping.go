// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqapiv3

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/linq-api-v3-go/internal/requestconfig"
	"github.com/stainless-sdks/linq-api-v3-go/option"
)

// ChatTypingService contains methods and other services that help with interacting
// with the linq API.
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
// **Note:** Group chat typing indicators are not currently supported.
func (r *ChatTypingService) Start(ctx context.Context, chatID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return
	}
	path := fmt.Sprintf("v3/chats/%s/typing", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Stop the typing indicator for the chat.
//
// **Note:** Typing indicators are automatically stopped when a message is sent, so
// calling this endpoint after sending a message is unnecessary.
//
// **Note:** Group chat typing indicators are not currently supported.
func (r *ChatTypingService) Stop(ctx context.Context, chatID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return
	}
	path := fmt.Sprintf("v3/chats/%s/typing", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}
