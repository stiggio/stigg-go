// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stiggio/stigg-go/internal/apijson"
	"github.com/stiggio/stigg-go/internal/apiquery"
	"github.com/stiggio/stigg-go/internal/requestconfig"
	"github.com/stiggio/stigg-go/option"
	"github.com/stiggio/stigg-go/packages/pagination"
	"github.com/stiggio/stigg-go/packages/param"
	"github.com/stiggio/stigg-go/packages/respjson"
)

// V1BetaEntityTypeService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BetaEntityTypeService] method instead.
type V1BetaEntityTypeService struct {
	Options []option.RequestOption
}

// NewV1BetaEntityTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1BetaEntityTypeService(opts ...option.RequestOption) (r V1BetaEntityTypeService) {
	r = V1BetaEntityTypeService{}
	r.Options = opts
	return
}

// Returns a cursor-paginated list of entity types defined in the environment.
// Entity types are vendor-defined categories of resource that can be governed
// (e.g. Org, Team, User).
func (r *V1BetaEntityTypeService) List(ctx context.Context, params V1BetaEntityTypeListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1BetaEntityTypeListResponse], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1-beta/entity-types"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a cursor-paginated list of entity types defined in the environment.
// Entity types are vendor-defined categories of resource that can be governed
// (e.g. Org, Team, User).
func (r *V1BetaEntityTypeService) ListAutoPaging(ctx context.Context, params V1BetaEntityTypeListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1BetaEntityTypeListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, params, opts...))
}

// Batched create-or-update of entity types. Existing types matched by id are
// updated; new ids are created. Idempotent — re-submitting the same payload
// converges to the same state.
func (r *V1BetaEntityTypeService) Upsert(ctx context.Context, params V1BetaEntityTypeUpsertParams, opts ...option.RequestOption) (res *V1BetaEntityTypeUpsertResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1-beta/entity-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A vendor-defined category of resource that can be governed (e.g. Org, Team,
// User). Vendors define entity types once per environment; their customers create
// instances (entities) of these types and the governance engine tracks usage and
// enforces limits per instance.
type V1BetaEntityTypeListResponse struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Dimension keys used to attribute usage events to instances of this type (e.g.
	// ["orgId"]). Empty array means no attribution.
	AttributionKeys []string `json:"attributionKeys" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The display name for the entity type
	DisplayName string `json:"displayName" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AttributionKeys respjson.Field
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityTypeListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityTypeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entity types after upsert.
type V1BetaEntityTypeUpsertResponse struct {
	Data []V1BetaEntityTypeUpsertResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityTypeUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityTypeUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A vendor-defined category of resource that can be governed (e.g. Org, Team,
// User). Vendors define entity types once per environment; their customers create
// instances (entities) of these types and the governance engine tracks usage and
// enforces limits per instance.
type V1BetaEntityTypeUpsertResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Dimension keys used to attribute usage events to instances of this type (e.g.
	// ["orgId"]). Empty array means no attribution.
	AttributionKeys []string `json:"attributionKeys" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The display name for the entity type
	DisplayName string `json:"displayName" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AttributionKeys respjson.Field
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityTypeUpsertResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityTypeUpsertResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BetaEntityTypeListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit          param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BetaEntityTypeListParams]'s query parameters as
// `url.Values`.
func (r V1BetaEntityTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1BetaEntityTypeUpsertParams struct {
	// Entity types to upsert (1–100 per request)
	Types          []V1BetaEntityTypeUpsertParamsType `json:"types,omitzero" api:"required"`
	XAccountID     param.Opt[string]                  `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]                  `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1BetaEntityTypeUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaEntityTypeUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaEntityTypeUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entity type definition.
//
// The properties ID, AttributionKeys, DisplayName are required.
type V1BetaEntityTypeUpsertParamsType struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Dimension keys used to attribute usage events to instances of this type (e.g.
	// ["orgId"]). Empty array means no attribution.
	AttributionKeys []string `json:"attributionKeys,omitzero" api:"required"`
	// The display name for the entity type
	DisplayName string `json:"displayName" api:"required"`
	paramObj
}

func (r V1BetaEntityTypeUpsertParamsType) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaEntityTypeUpsertParamsType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaEntityTypeUpsertParamsType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
