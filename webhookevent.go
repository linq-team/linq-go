// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqapiv3

import (
	"context"
	"net/http"
	"slices"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/respjson"
	"github.com/linq-team/linq-go/shared/constant"
)

// WebhookEventService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookEventService] method instead.
type WebhookEventService struct {
	Options []option.RequestOption
}

// NewWebhookEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookEventService(opts ...option.RequestOption) (r WebhookEventService) {
	r = WebhookEventService{}
	r.Options = opts
	return
}

// Returns all available webhook event types that can be subscribed to. Use this
// endpoint to discover valid values for the `subscribed_events` field when
// creating or updating webhook subscriptions.
func (r *WebhookEventService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/webhook-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Valid webhook event types that can be subscribed to
type WebhookEventType string

const (
	WebhookEventTypeMessageSent                WebhookEventType = "message.sent"
	WebhookEventTypeMessageReceived            WebhookEventType = "message.received"
	WebhookEventTypeMessageRead                WebhookEventType = "message.read"
	WebhookEventTypeMessageDelivered           WebhookEventType = "message.delivered"
	WebhookEventTypeMessageFailed              WebhookEventType = "message.failed"
	WebhookEventTypeReactionAdded              WebhookEventType = "reaction.added"
	WebhookEventTypeReactionRemoved            WebhookEventType = "reaction.removed"
	WebhookEventTypeParticipantAdded           WebhookEventType = "participant.added"
	WebhookEventTypeParticipantRemoved         WebhookEventType = "participant.removed"
	WebhookEventTypeChatCreated                WebhookEventType = "chat.created"
	WebhookEventTypeChatGroupNameUpdated       WebhookEventType = "chat.group_name_updated"
	WebhookEventTypeChatGroupIconUpdated       WebhookEventType = "chat.group_icon_updated"
	WebhookEventTypeChatGroupNameUpdateFailed  WebhookEventType = "chat.group_name_update_failed"
	WebhookEventTypeChatGroupIconUpdateFailed  WebhookEventType = "chat.group_icon_update_failed"
	WebhookEventTypeChatTypingIndicatorStarted WebhookEventType = "chat.typing_indicator.started"
	WebhookEventTypeChatTypingIndicatorStopped WebhookEventType = "chat.typing_indicator.stopped"
	WebhookEventTypePhoneNumberStatusUpdated   WebhookEventType = "phone_number.status_updated"
)

type WebhookEventListResponse struct {
	// URL to the webhook events documentation
	DocURL constant.HTTPSApidocsLinqappComDocumentationWebhookEvents `json:"doc_url,required"`
	// List of all available webhook event types
	Events []WebhookEventType `json:"events,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocURL      respjson.Field
		Events      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
