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
// ChatBackgroundService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatBackgroundService] method instead.
type ChatBackgroundService struct {
	Options []option.RequestOption
}

// NewChatBackgroundService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatBackgroundService(opts ...option.RequestOption) (r ChatBackgroundService) {
	r = ChatBackgroundService{}
	r.Options = opts
	return
}

// Remove the transcript background from a chat, resetting it to the default.
func (r *ChatBackgroundService) Remove(ctx context.Context, chatID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return err
	}
	path := fmt.Sprintf("v3/chats/%s/background", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Set the transcript background for a chat.
//
// Provide one of: a **color** (a named preset or a custom 2-stop gradient), a
// **dynamic** animated style, or a **photo** (by URL). The request is accepted
// asynchronously; the terminal result arrives via the `chat.background_updated`
// webhook on success, or `chat.background_update_failed` on failure.
//
// **Group chats are supported.** Requests for RCS or SMS chats are accepted
// (`202`) but no background is applied and no `chat.background_updated` webhook
// fires.
func (r *ChatBackgroundService) Set(ctx context.Context, chatID string, body ChatBackgroundSetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return err
	}
	path := fmt.Sprintf("v3/chats/%s/background", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type ChatBackgroundSetParams struct {
	// The background family.
	//
	// Any of "color", "dynamic", "photo".
	Type ChatBackgroundSetParamsType `json:"type,omitzero" api:"required"`
	// Photo: the image URL to embed in the background. Must be an absolute `https` URL
	// pointing at an image (`.jpg`, `.png`, `.heic`, `.webp`), and the image is
	// fetched and re-hosted on our CDN before the request is accepted — the same way
	// `group_chat_icon` works. A URL we cannot fetch, or one that isn't an image, is
	// rejected with a `400` (`5007`/`5006`) rather than failing later on the device.
	ImageURL param.Opt[string] `json:"image_url,omitzero" format:"uri"`
	// Color: a named swatch — `mango`, `ice`, `plum`, `deep_sea`, `green_apple`,
	// `cherry`, `bubblegum`, `tangerine`, `magenta`, `lime`, `silver`, `carbon`,
	// `stone` — or `custom` (supply `shades`). Dynamic: the variant within the `style`
	// (e.g. `sunrise`).
	//
	// An unrecognized value still returns `202`, but no background is applied and no
	// `chat.background_updated` webhook fires. Send one of the values above.
	Variant param.Opt[string] `json:"variant,omitzero"`
	// Color with `variant: custom`: the two gradient stops as hex, top then bottom.
	// Ignored for named color variants (they carry their own two colors).
	Shades []string `json:"shades,omitzero"`
	// Dynamic: the animated style.
	//
	// Any of "sky", "water", "aurora", "glitter".
	Style ChatBackgroundSetParamsStyle `json:"style,omitzero"`
	paramObj
}

func (r ChatBackgroundSetParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatBackgroundSetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatBackgroundSetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The background family.
type ChatBackgroundSetParamsType string

const (
	ChatBackgroundSetParamsTypeColor   ChatBackgroundSetParamsType = "color"
	ChatBackgroundSetParamsTypeDynamic ChatBackgroundSetParamsType = "dynamic"
	ChatBackgroundSetParamsTypePhoto   ChatBackgroundSetParamsType = "photo"
)

// Dynamic: the animated style.
type ChatBackgroundSetParamsStyle string

const (
	ChatBackgroundSetParamsStyleSky     ChatBackgroundSetParamsStyle = "sky"
	ChatBackgroundSetParamsStyleWater   ChatBackgroundSetParamsStyle = "water"
	ChatBackgroundSetParamsStyleAurora  ChatBackgroundSetParamsStyle = "aurora"
	ChatBackgroundSetParamsStyleGlitter ChatBackgroundSetParamsStyle = "glitter"
)
