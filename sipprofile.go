// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// ## SIP configuration
//
// Your Linq representative provides the signaling and media IP ranges and ports to
// allow through your firewall during onboarding.
//
// Set `primary_uri` to your endpoint as scheme + host — `sips:` for TLS or `sip:`
// for UDP, with no user part (we add it). Each call arrives at your endpoint as:
//
// ```
// sip:<called-number>@<your-host>
// ```
//
// so you can route on the dialed number. The caller's number is in the `From`
// header in E.164.
//
// SipProfileService contains methods and other services that help with interacting
// with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSipProfileService] method instead.
type SipProfileService struct {
	Options []option.RequestOption
}

// NewSipProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSipProfileService(opts ...option.RequestOption) (r SipProfileService) {
	r = SipProfileService{}
	r.Options = opts
	return
}

// ## SIP configuration
//
// Your Linq representative provides the signaling and media IP ranges and ports to
// allow through your firewall during onboarding.
//
// Set `primary_uri` to your endpoint as scheme + host — `sips:` for TLS or `sip:`
// for UDP, with no user part (we add it). Each call arrives at your endpoint as:
//
// ```
// sip:<called-number>@<your-host>
// ```
//
// so you can route on the dialed number. The caller's number is in the `From`
// header in E.164.
//
// ---
//
// Create a SIP routing profile for your inbound calls, then assign it to a phone
// number to start routing calls to its SIP URI.
func (r *SipProfileService) New(ctx context.Context, body SipProfileNewParams, opts ...option.RequestOption) (res *SipProfile, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/sip-profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve details for a specific SIP profile.
func (r *SipProfileService) Get(ctx context.Context, sipProfileID string, opts ...option.RequestOption) (res *SipProfile, err error) {
	opts = slices.Concat(r.Options, opts)
	if sipProfileID == "" {
		err = errors.New("missing required sipProfileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/sip-profiles/%s", url.PathEscape(sipProfileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a SIP profile. Only the fields you send are changed; omitted fields keep
// their current value. Set `backup_uri` to an empty string to clear it.
func (r *SipProfileService) Update(ctx context.Context, sipProfileID string, body SipProfileUpdateParams, opts ...option.RequestOption) (res *SipProfile, err error) {
	opts = slices.Concat(r.Options, opts)
	if sipProfileID == "" {
		err = errors.New("missing required sipProfileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/sip-profiles/%s", url.PathEscape(sipProfileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Retrieve all SIP profiles for the authenticated partner.
func (r *SipProfileService) List(ctx context.Context, opts ...option.RequestOption) (res *SipProfileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/sip-profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a SIP profile. The profile must not be assigned to any phone number;
// unassign it first or the request is rejected.
func (r *SipProfileService) Delete(ctx context.Context, sipProfileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if sipProfileID == "" {
		err = errors.New("missing required sipProfileId parameter")
		return err
	}
	path := fmt.Sprintf("v3/sip-profiles/%s", url.PathEscape(sipProfileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// ## SIP configuration
//
// Your Linq representative provides the signaling and media IP ranges and ports to
// allow through your firewall during onboarding.
//
// Set `primary_uri` to your endpoint as scheme + host — `sips:` for TLS or `sip:`
// for UDP, with no user part (we add it). Each call arrives at your endpoint as:
//
// ```
// sip:<called-number>@<your-host>
// ```
//
// so you can route on the dialed number. The caller's number is in the `From`
// header in E.164.
//
// ---
//
// Assign a SIP profile to route inbound calls to its SIP URI. This takes
// precedence over an existing `forwarding_number` without clearing it. Send
// `sip_profile_id: null` to remove the SIP assignment and resume existing
// forwarding, if configured. Allow about 5 minutes for routing changes to
// propagate. The phone number must belong to the authenticated partner.
func (r *SipProfileService) Assign(ctx context.Context, phoneNumber string, body SipProfileAssignParams, opts ...option.RequestOption) (res *SipProfileAssignResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/phone-numbers/%s/sip-profile", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type SipProfile struct {
	// Unique identifier for the SIP profile
	ID string `json:"id" api:"required"`
	// How many phone numbers reference this profile
	AssignedNumberCount int64     `json:"assigned_number_count" api:"required"`
	CreatedAt           time.Time `json:"created_at" api:"required" format:"date-time"`
	// Human-readable name
	Name string `json:"name" api:"required"`
	// Your SIP endpoint as scheme + host (`sip:` UDP / `sips:` TLS); no user part — we
	// set it to the called number on each call.
	PrimaryUri string    `json:"primary_uri" api:"required"`
	UpdatedAt  time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Monotonic version, incremented on each update
	Version int64 `json:"version" api:"required"`
	// Optional failover SIP URI. Null when unset.
	BackupUri string `json:"backup_uri" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AssignedNumberCount respjson.Field
		CreatedAt           respjson.Field
		Name                respjson.Field
		PrimaryUri          respjson.Field
		UpdatedAt           respjson.Field
		Version             respjson.Field
		BackupUri           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SipProfile) RawJSON() string { return r.JSON.raw }
func (r *SipProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SipProfileListResponse struct {
	// The partner's SIP profiles
	SipProfiles []SipProfile `json:"sip_profiles" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SipProfiles respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SipProfileListResponse) RawJSON() string { return r.JSON.raw }
func (r *SipProfileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SipProfileAssignResponse struct {
	// The phone number whose assignment changed
	PhoneNumber string `json:"phone_number" api:"required"`
	// The assigned SIP profile id, or null when cleared
	SipProfileID string `json:"sip_profile_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber  respjson.Field
		SipProfileID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SipProfileAssignResponse) RawJSON() string { return r.JSON.raw }
func (r *SipProfileAssignResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SipProfileNewParams struct {
	Name string `json:"name" api:"required"`
	// Your SIP endpoint as scheme + host, e.g. `sips:sip.partner.example`. Must start
	// with `sip:` (UDP) or `sips:` (TLS) — the scheme sets the transport. Do not
	// include a user part: we set it to the called number on each call
	// (`sip:<called-number>@<your-host>`). A bare host with no scheme is rejected.
	PrimaryUri string `json:"primary_uri" api:"required"`
	// Optional failover SIP URI
	BackupUri param.Opt[string] `json:"backup_uri,omitzero"`
	paramObj
}

func (r SipProfileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SipProfileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SipProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SipProfileUpdateParams struct {
	// Set to "" to clear the backup URI.
	BackupUri  param.Opt[string] `json:"backup_uri,omitzero"`
	Name       param.Opt[string] `json:"name,omitzero"`
	PrimaryUri param.Opt[string] `json:"primary_uri,omitzero"`
	paramObj
}

func (r SipProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SipProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SipProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SipProfileAssignParams struct {
	// The SIP profile to assign, or null to clear the number's current assignment.
	SipProfileID param.Opt[string] `json:"sip_profile_id,omitzero" api:"required"`
	paramObj
}

func (r SipProfileAssignParams) MarshalJSON() (data []byte, err error) {
	type shadow SipProfileAssignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SipProfileAssignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
