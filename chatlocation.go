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

// Request and retrieve real-time location data via iMessage.
//
// Use these endpoints to request a contact's location, retrieve location data for
// contacts who are sharing with you, and subscribe to webhooks when someone starts
// or stops sharing their location.
//
// **Coordinates** are returned in
// [GeoJSON](https://datatracker.ietf.org/doc/html/rfc7946) format:
// `[longitude, latitude]` or `[longitude, latitude, altitude]` if altitude is
// available.
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
// Returns a [GeoJSON](https://datatracker.ietf.org/doc/html/rfc7946)
// `FeatureCollection` with a `Feature` for each participant actively sharing their
// location.
//
// Works for both 1:1 and group chats. In group chats, returns a separate feature
// for each participant who is sharing. Each feature's `properties.handle`
// identifies the user.
//
// Returns an empty `features` array if no one is sharing or no location data is
// available yet.
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

// Send a location sharing request to a contact. They will receive an iMessage
// prompt asking them to share their location.
//
// Location requests only work in **1:1 iMessage chats** (Apple limitation).
// Attempting to request location in a group chat, or in an SMS or RCS chat,
// returns `409` (Operation not supported on this chat's service type).
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
	// [longitude, latitude] or [longitude, latitude, altitude]
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
