// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/respjson"
)

// Let an agent pay on a customer's behalf with a single-use virtual card. Connect
// a customer once, then create a payment — a virtual card is minted scoped to that
// purchase and the card details are handed back for checkout.
//
// ExperienceService contains methods and other services that help with interacting
// with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperienceService] method instead.
type ExperienceService struct {
	Options []option.RequestOption
}

// NewExperienceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExperienceService(opts ...option.RequestOption) (r ExperienceService) {
	r = ExperienceService{}
	r.Options = opts
	return
}

// Get one experience
func (r *ExperienceService) Get(ctx context.Context, experience string, opts ...option.RequestOption) (res *ExperienceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if experience == "" {
		err = errors.New("missing required experience parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/experiences/%s", url.PathEscape(experience))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The experiences enabled for your account, with the actions you may invoke on
// each and the fields each action accepts. This is the authoritative list — an
// action missing here cannot be sent.
func (r *ExperienceService) List(ctx context.Context, opts ...option.RequestOption) (res *ExperienceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/experiences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// What an experience offers you. Deliberately a projection: where its templates
// live and how they are built is not yours to depend on, so it is not here.
type ExperienceGetResponse struct {
	Actions     []ExperienceGetResponseAction `json:"actions"`
	DisplayName string                        `json:"display_name"`
	Experience  string                        `json:"experience"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions     respjson.Field
		DisplayName respjson.Field
		Experience  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperienceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperienceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperienceGetResponseAction struct {
	// Fields you may send in `params`, keyed by the exact name to use.
	Fields  map[string]ExperienceGetResponseActionField `json:"fields"`
	Name    string                                      `json:"name"`
	Summary string                                      `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields      respjson.Field
		Name        respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperienceGetResponseAction) RawJSON() string { return r.JSON.raw }
func (r *ExperienceGetResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperienceGetResponseActionField struct {
	// Maximum length, for strings.
	Max      int64 `json:"max"`
	Required bool  `json:"required"`
	// Any of "string", "cents", "int", "url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperienceGetResponseActionField) RawJSON() string { return r.JSON.raw }
func (r *ExperienceGetResponseActionField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperienceListResponse struct {
	Experiences []ExperienceListResponseExperience `json:"experiences"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Experiences respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperienceListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperienceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What an experience offers you. Deliberately a projection: where its templates
// live and how they are built is not yours to depend on, so it is not here.
type ExperienceListResponseExperience struct {
	Actions     []ExperienceListResponseExperienceAction `json:"actions"`
	DisplayName string                                   `json:"display_name"`
	Experience  string                                   `json:"experience"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions     respjson.Field
		DisplayName respjson.Field
		Experience  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperienceListResponseExperience) RawJSON() string { return r.JSON.raw }
func (r *ExperienceListResponseExperience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperienceListResponseExperienceAction struct {
	// Fields you may send in `params`, keyed by the exact name to use.
	Fields  map[string]ExperienceListResponseExperienceActionField `json:"fields"`
	Name    string                                                 `json:"name"`
	Summary string                                                 `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields      respjson.Field
		Name        respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperienceListResponseExperienceAction) RawJSON() string { return r.JSON.raw }
func (r *ExperienceListResponseExperienceAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperienceListResponseExperienceActionField struct {
	// Maximum length, for strings.
	Max      int64 `json:"max"`
	Required bool  `json:"required"`
	// Any of "string", "cents", "int", "url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperienceListResponseExperienceActionField) RawJSON() string { return r.JSON.raw }
func (r *ExperienceListResponseExperienceActionField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
