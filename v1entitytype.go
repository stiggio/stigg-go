// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"errors"
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

// V1EntityTypeService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EntityTypeService] method instead.
type V1EntityTypeService struct {
	Options []option.RequestOption
}

// NewV1EntityTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1EntityTypeService(opts ...option.RequestOption) (r V1EntityTypeService) {
	r = V1EntityTypeService{}
	r.Options = opts
	return
}

// Returns a cursor-paginated list of entity types defined in the environment.
// Entity types are vendor-defined categories of resource that can be governed
// (e.g. Org, Team, User). Archived entity types are excluded unless
// includeArchived=true.
func (r *V1EntityTypeService) List(ctx context.Context, params V1EntityTypeListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1EntityTypeListResponse], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/entity-types"
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
// (e.g. Org, Team, User). Archived entity types are excluded unless
// includeArchived=true.
func (r *V1EntityTypeService) ListAutoPaging(ctx context.Context, params V1EntityTypeListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1EntityTypeListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, params, opts...))
}

// Archives entity types in bulk by id. An entity type with active entities cannot
// be archived: archive its entities first, or the whole request fails with 400.
func (r *V1EntityTypeService) Archive(ctx context.Context, params V1EntityTypeArchiveParams, opts ...option.RequestOption) (res *V1EntityTypeArchiveResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/entity-types/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns whether any active entity of this entity type exists, across every
// customer. An entity type can be archived only when this is false. An unknown
// entity type returns false.
func (r *V1EntityTypeService) CheckActiveEntities(ctx context.Context, id string, query V1EntityTypeCheckActiveEntitiesParams, opts ...option.RequestOption) (res *V1EntityTypeCheckActiveEntitiesResponse, err error) {
	if !param.IsOmitted(query.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", query.XAccountID.Value)))
	}
	if !param.IsOmitted(query.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", query.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/entity-types/%s/has-active-entities", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Batched create-or-update of entity types. Existing types matched by id are
// updated; new ids are created. Idempotent — re-submitting the same payload
// converges to the same state. An upsert of an archived entity type restores it.
func (r *V1EntityTypeService) Upsert(ctx context.Context, params V1EntityTypeUpsertParams, opts ...option.RequestOption) (res *V1EntityTypeUpsertResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/entity-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A vendor-defined category of resource that can be governed (e.g. Org, Team,
// User). Vendors define entity types once per environment; their customers create
// instances (entities) of these types and the governance engine tracks usage and
// enforces limits per instance.
type V1EntityTypeListResponse struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Dimension keys used to attribute usage events to instances of this type (e.g.
	// ["orgId"]). Empty array means no attribution.
	AttributionKeys []string `json:"attributionKeys" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// What this entity type represents and what it is for governing, or null when none
	// is set
	Description string `json:"description" api:"required"`
	// The display name for the entity type
	DisplayName string `json:"displayName" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ArchivedAt      respjson.Field
		AttributionKeys respjson.Field
		CreatedAt       respjson.Field
		Description     respjson.Field
		DisplayName     respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EntityTypeListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EntityTypeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapped response echoing the entity type ids that were acted on by an archive
// call
type V1EntityTypeArchiveResponse struct {
	// List of entity type identifiers that were acted on
	Data V1EntityTypeArchiveResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EntityTypeArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EntityTypeArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entity type identifiers that were acted on
type V1EntityTypeArchiveResponseData struct {
	// Entity type identifiers to act on
	IDs []string `json:"ids" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IDs         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EntityTypeArchiveResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EntityTypeArchiveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapped response with whether an entity type has active entities
type V1EntityTypeCheckActiveEntitiesResponse struct {
	// Whether an entity type has active entities
	Data V1EntityTypeCheckActiveEntitiesResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EntityTypeCheckActiveEntitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EntityTypeCheckActiveEntitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether an entity type has active entities
type V1EntityTypeCheckActiveEntitiesResponseData struct {
	// Whether at least one active (non-archived) entity of the entity type exists,
	// across every customer
	HasActiveEntities bool `json:"hasActiveEntities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasActiveEntities respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EntityTypeCheckActiveEntitiesResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EntityTypeCheckActiveEntitiesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entity types after upsert.
type V1EntityTypeUpsertResponse struct {
	Data []V1EntityTypeUpsertResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EntityTypeUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EntityTypeUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A vendor-defined category of resource that can be governed (e.g. Org, Team,
// User). Vendors define entity types once per environment; their customers create
// instances (entities) of these types and the governance engine tracks usage and
// enforces limits per instance.
type V1EntityTypeUpsertResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Dimension keys used to attribute usage events to instances of this type (e.g.
	// ["orgId"]). Empty array means no attribution.
	AttributionKeys []string `json:"attributionKeys" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// What this entity type represents and what it is for governing, or null when none
	// is set
	Description string `json:"description" api:"required"`
	// The display name for the entity type
	DisplayName string `json:"displayName" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ArchivedAt      respjson.Field
		AttributionKeys respjson.Field
		CreatedAt       respjson.Field
		Description     respjson.Field
		DisplayName     respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EntityTypeUpsertResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EntityTypeUpsertResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EntityTypeListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit          param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	// Whether to include archived entity types. One of: true, false
	//
	// Any of "true", "false".
	IncludeArchived V1EntityTypeListParamsIncludeArchived `query:"includeArchived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EntityTypeListParams]'s query parameters as `url.Values`.
func (r V1EntityTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to include archived entity types. One of: true, false
type V1EntityTypeListParamsIncludeArchived string

const (
	V1EntityTypeListParamsIncludeArchivedTrue  V1EntityTypeListParamsIncludeArchived = "true"
	V1EntityTypeListParamsIncludeArchivedFalse V1EntityTypeListParamsIncludeArchived = "false"
)

type V1EntityTypeArchiveParams struct {
	// Entity type identifiers to act on
	IDs            []string          `json:"ids,omitzero" api:"required"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1EntityTypeArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EntityTypeArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EntityTypeArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EntityTypeCheckActiveEntitiesParams struct {
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

type V1EntityTypeUpsertParams struct {
	// Entity types to upsert (1–100 per request)
	Types          []V1EntityTypeUpsertParamsType `json:"types,omitzero" api:"required"`
	XAccountID     param.Opt[string]              `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]              `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1EntityTypeUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EntityTypeUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EntityTypeUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entity type definition.
//
// The properties ID, AttributionKeys, DisplayName are required.
type V1EntityTypeUpsertParamsType struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Dimension keys used to attribute usage events to instances of this type (e.g.
	// ["orgId"]). Empty array means no attribution.
	AttributionKeys []string `json:"attributionKeys,omitzero" api:"required"`
	// The display name for the entity type
	DisplayName string `json:"displayName" api:"required"`
	// What this entity type represents and what it is for governing. Omit to preserve
	// the stored value, or send an empty string or null to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r V1EntityTypeUpsertParamsType) MarshalJSON() (data []byte, err error) {
	type shadow V1EntityTypeUpsertParamsType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EntityTypeUpsertParamsType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
