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
	"github.com/linq-team/linq-go/packages/respjson"
)

// Request a contact's location, retrieve location for contacts sharing with you,
// and subscribe to webhooks when someone starts or stops sharing.
//
// **Coordinates** are returned in
// [GeoJSON](https://datatracker.ietf.org/doc/html/rfc7946) format:
// `[longitude, latitude]`.
//
// ### Reading location is poll-based
//
// Poll `GET /v3/chats/{chatId}/location` whenever you need the latest position.
// **There is no webhook that pushes updated coordinates** — the
// `location.sharing.started` / `location.sharing.stopped` webhooks fire only when
// a contact begins or ends sharing, not on each position update. To track a moving
// contact, poll the `GET` endpoint.
//
// ### Freshness
//
// Each feature's `properties.updated_at` tells you when that participant's
// location was last updated — use it to judge freshness.
//
// ### Polling guidance
//
// Locations refresh on Apple's cadence, not per request — polling faster than a
// participant's location actually updates just returns the same position. Poll at
// a modest interval (for example, once every few minutes per chat) rather than
// continuously.
//
// ### Why is location empty after `location.sharing.started` fired?
//
// If the contact started sharing from the **standalone Find My app** instead of
// the Messages conversation, the share may be tied to their **Apple ID email**
// rather than their phone number — the webhook's `shared_by` field shows the email
// in that case. Location is readable only through a chat with the handle that
// shared, so `GET /v3/chats/{chatId}/location` on the phone-number chat stays
// empty.
//
// The fix: have the contact stop sharing and re-share from **Find My inside the
// Messages conversation** with your number.
//
// ChatLocationService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatLocationService] method instead.
type ChatLocationService struct {
	Options []option.RequestOption
}

// NewChatLocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatLocationService(opts ...option.RequestOption) (r ChatLocationService) {
	r = ChatLocationService{}
	r.Options = opts
	return
}

// Retrieve the current location for contacts sharing with you in a chat.
//
// The response is wrapped in the standard `{ "success": true, "data": ... }`
// envelope — the body is **not** a bare GeoJSON document. `data` is a
// [GeoJSON](https://datatracker.ietf.org/doc/html/rfc7946) `FeatureCollection`
// with a `Feature` for each participant actively sharing their location.
//
// Works for both 1:1 and group chats. In group chats, `data.features` contains a
// separate feature for each participant who is sharing. Each feature's
// `properties.handle` identifies the user.
//
// A participant appears as soon as their first position arrives, typically within
// a second or two of sharing starting.
//
// Returns an empty `data.features` array if no one is sharing or no location data
// is available yet. If sharing started but this stays empty, see the **Location
// Sharing** overview.
//
// Poll this endpoint to track a moving contact. `properties.updated_at` reflects
// when each participant's location was last updated. There is no coordinate-update
// webhook. See the **Location Sharing** overview for polling guidance.
func (r *ChatLocationService) Get(ctx context.Context, chatID string, opts ...option.RequestOption) (res *GetChatLocationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/chats/%s/location", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Request a contact in a chat to share their location. They receive an iMessage
// prompt and must accept before any location is available; once they do, read
// their location coordinates with `GET /v3/chats/{chatId}/location`.
//
// The request is delivered asynchronously. The endpoint returns immediately with
// `{ "success": true, "message": "Location request sent" }` and does not return
// coordinates.
//
// Rejected with `409` if the recipient is already sharing — read their location
// with `GET /v3/chats/{chatId}/location` instead of re-requesting.
//
// Rate limited per chat, since each request prompts the recipient's device.
// Exceeding it returns `429` with a `Retry-After` header.
//
// Location requests only work in **1:1 iMessage chats** (Apple limitation):
//
//   - Group chats (any service) return `409` with code `2016`
//     (`GroupChatNotSupported`).
//   - 1:1 SMS and RCS chats return `409` with code `2017`
//     (`ChatServiceNotSupported`).
func (r *ChatLocationService) Request(ctx context.Context, chatID string, opts ...option.RequestOption) (res *LocationRequestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/chats/%s/location/request", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type GetChatLocationResponse struct {
	Data    GetChatLocationResponseData `json:"data" api:"required"`
	Success bool                        `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetChatLocationResponse) RawJSON() string { return r.JSON.raw }
func (r *GetChatLocationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetChatLocationResponseData struct {
	Features []GetChatLocationResponseDataFeature `json:"features" api:"required"`
	// Any of "FeatureCollection".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Features    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetChatLocationResponseData) RawJSON() string { return r.JSON.raw }
func (r *GetChatLocationResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetChatLocationResponseDataFeature struct {
	Geometry   GetChatLocationResponseDataFeatureGeometry   `json:"geometry" api:"required"`
	Properties GetChatLocationResponseDataFeatureProperties `json:"properties" api:"required"`
	// Any of "Feature".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geometry    respjson.Field
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetChatLocationResponseDataFeature) RawJSON() string { return r.JSON.raw }
func (r *GetChatLocationResponseDataFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetChatLocationResponseDataFeatureGeometry struct {
	// [longitude, latitude]
	Coordinates []float64 `json:"coordinates" api:"required"`
	// Any of "Point".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetChatLocationResponseDataFeatureGeometry) RawJSON() string { return r.JSON.raw }
func (r *GetChatLocationResponseDataFeatureGeometry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetChatLocationResponseDataFeatureProperties struct {
	// Phone number or email of the person sharing their location
	Handle string `json:"handle" api:"required"`
	// Full street address
	Address string `json:"address"`
	// City or locality name
	Locality string `json:"locality"`
	// When the location was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handle      respjson.Field
		Address     respjson.Field
		Locality    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetChatLocationResponseDataFeatureProperties) RawJSON() string { return r.JSON.raw }
func (r *GetChatLocationResponseDataFeatureProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationRequestResponse struct {
	Message string `json:"message" api:"required"`
	Success bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationRequestResponse) RawJSON() string { return r.JSON.raw }
func (r *LocationRequestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
