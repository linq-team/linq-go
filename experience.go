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

// An **experience** renders inside Linq's iMessage app as a native card, instead
// of as text or a link. You invoke one by name; Linq resolves the recipient, mints
// any session it needs, composes the card and sends it.
//
// Send it to `POST /v3/chats/{chatId}/messages`:
//
// ```json
//
//	{
//	  "message": {
//	    "experience": {
//	      "name": "agentpay",
//	      "action": "request_payment",
//	      "params": {
//	        "checkout_url": "https://zero.linqapp.com/pay/acme?session=tok_..."
//	      }
//	    }
//	  }
//	}
//
// ```
//
// The key is `experience` — what you're invoking. Nested under it is its `name`,
// the action you're invoking on it, and that action's params. A card **is** the
// whole message on Apple's side, so a message carries either `experience` or
// `parts`, never both, and an action goes to exactly one recipient.
//
// ## What you can invoke
//
// | Experience  | Action            | What the customer sees                                                                        |
// | ----------- | ----------------- | --------------------------------------------------------------------------------------------- |
// | `agentpay`  | `request_payment` | A payment request they can pay in the app. Turns itself into "Paid" in place once it settles. |
// | `agentcard` | `attach_card`     | A prompt to add a card to their wallet.                                                       |
// | `agentcard` | `approve_card`    | A passkey approval for a virtual card.                                                        |
// | `link`      | `open`            | A card that opens a URL you supply.                                                           |
//
// `GET /v3/experiences` is the authoritative list for your account, with every
// action and the fields each accepts — an action missing there cannot be sent.
// Fields are display copy unless documented otherwise.
//
// ## Params are checked before the card is sent
//
// Unknown fields are **rejected rather than ignored**, so copy that would never
// have rendered fails for you now instead of arriving wrong on somebody's phone.
// Some fields are read rather than sent: `agentpay`'s `request_payment` takes only
// a `checkout_url` and resolves the amount and reason from that payment request,
// so a card can never claim a figure the checkout will not charge.
//
// Cards are **iMessage-only**. Recipients without the app see a static version
// built from the same copy; SMS and RCS recipients cannot receive one at all
// (error codes 2018 and 4005).
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
