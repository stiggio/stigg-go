// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
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

// V1EventBetaEntityTypeService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventBetaEntityTypeService] method instead.
type V1EventBetaEntityTypeService struct {
	Options []option.RequestOption
}

// NewV1EventBetaEntityTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1EventBetaEntityTypeService(opts ...option.RequestOption) (r V1EventBetaEntityTypeService) {
	r = V1EventBetaEntityTypeService{}
	r.Options = opts
	return
}

// Returns a cursor-paginated list of entity types defined in the environment.
// Entity types are vendor-defined categories of resource that can be governed
// (e.g. Org, Team, User).
func (r *V1EventBetaEntityTypeService) List(ctx context.Context, query V1EventBetaEntityTypeListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1EventBetaEntityTypeListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1-beta/entity-types"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *V1EventBetaEntityTypeService) ListAutoPaging(ctx context.Context, query V1EventBetaEntityTypeListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1EventBetaEntityTypeListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, query, opts...))
}

// Batched create-or-update of entity types. Existing types matched by id are
// updated; new ids are created. Idempotent — re-submitting the same payload
// converges to the same state.
func (r *V1EventBetaEntityTypeService) Upsert(ctx context.Context, body V1EventBetaEntityTypeUpsertParams, opts ...option.RequestOption) (res *V1EventBetaEntityTypeUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1-beta/entity-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// A vendor-defined category of resource that can be governed (e.g. Org, Team,
// User). Vendors define entity types once per environment; their customers create
// instances (entities) of these types and the governance engine tracks usage and
// enforces limits per instance.
type V1EventBetaEntityTypeListResponse struct {
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
func (r V1EventBetaEntityTypeListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaEntityTypeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entity types after upsert.
type V1EventBetaEntityTypeUpsertResponse struct {
	Data []V1EventBetaEntityTypeUpsertResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaEntityTypeUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaEntityTypeUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A vendor-defined category of resource that can be governed (e.g. Org, Team,
// User). Vendors define entity types once per environment; their customers create
// instances (entities) of these types and the governance engine tracks usage and
// enforces limits per instance.
type V1EventBetaEntityTypeUpsertResponseData struct {
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
func (r V1EventBetaEntityTypeUpsertResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaEntityTypeUpsertResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventBetaEntityTypeListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventBetaEntityTypeListParams]'s query parameters as
// `url.Values`.
func (r V1EventBetaEntityTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1EventBetaEntityTypeUpsertParams struct {
	// Entity types to upsert (1–100 per request)
	Types []V1EventBetaEntityTypeUpsertParamsType `json:"types,omitzero" api:"required"`
	paramObj
}

func (r V1EventBetaEntityTypeUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventBetaEntityTypeUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventBetaEntityTypeUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entity type definition.
//
// The properties ID, AttributionKeys, DisplayName are required.
type V1EventBetaEntityTypeUpsertParamsType struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Dimension keys used to attribute usage events to instances of this type (e.g.
	// ["orgId"]). Empty array means no attribution.
	AttributionKeys []string `json:"attributionKeys,omitzero" api:"required"`
	// The display name for the entity type
	DisplayName string `json:"displayName" api:"required"`
	paramObj
}

func (r V1EventBetaEntityTypeUpsertParamsType) MarshalJSON() (data []byte, err error) {
	type shadow V1EventBetaEntityTypeUpsertParamsType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventBetaEntityTypeUpsertParamsType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
