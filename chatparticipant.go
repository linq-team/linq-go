// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqapiv3

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/linq-api-v3-go/internal/apijson"
	"github.com/stainless-sdks/linq-api-v3-go/internal/requestconfig"
	"github.com/stainless-sdks/linq-api-v3-go/option"
	"github.com/stainless-sdks/linq-api-v3-go/packages/param"
	"github.com/stainless-sdks/linq-api-v3-go/packages/respjson"
)

// ChatParticipantService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatParticipantService] method instead.
type ChatParticipantService struct {
	Options []option.RequestOption
}

// NewChatParticipantService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatParticipantService(opts ...option.RequestOption) (r ChatParticipantService) {
	r = ChatParticipantService{}
	r.Options = opts
	return
}

// Add a new participant to an existing group chat.
//
// **Requirements:**
//
//   - Group chats only (3+ existing participants)
//   - New participant must support the same messaging service as the group
//   - Cross-service additions not allowed (e.g., can't add RCS-only user to iMessage
//     group)
//   - For cross-service scenarios, create a new chat instead
func (r *ChatParticipantService) Add(ctx context.Context, chatID string, body ChatParticipantAddParams, opts ...option.RequestOption) (res *ChatParticipantAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return
	}
	path := fmt.Sprintf("v3/chats/%s/participants", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove a participant from an existing group chat.
//
// **Requirements:**
//
// - Group chats only
// - Must have 3+ participants after removal
func (r *ChatParticipantService) Remove(ctx context.Context, chatID string, body ChatParticipantRemoveParams, opts ...option.RequestOption) (res *ChatParticipantRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return
	}
	path := fmt.Sprintf("v3/chats/%s/participants", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type ChatParticipantAddResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	TraceID string `json:"trace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		TraceID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatParticipantAddResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatParticipantAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatParticipantRemoveResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	TraceID string `json:"trace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		TraceID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatParticipantRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatParticipantRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatParticipantAddParams struct {
	// Phone number (E.164 format) or email address of the participant to add
	Handle string `json:"handle,required"`
	paramObj
}

func (r ChatParticipantAddParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatParticipantAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatParticipantAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatParticipantRemoveParams struct {
	// Phone number (E.164 format) or email address of the participant to remove
	Handle string `json:"handle,required"`
	paramObj
}

func (r ChatParticipantRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatParticipantRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatParticipantRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
