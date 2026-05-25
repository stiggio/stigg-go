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

// V1BetaEntityService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BetaEntityService] method instead.
type V1BetaEntityService struct {
	Options []option.RequestOption
}

// NewV1BetaEntityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1BetaEntityService(opts ...option.RequestOption) (r V1BetaEntityService) {
	r = V1BetaEntityService{}
	r.Options = opts
	return
}

// Retrieves a single entity for the given customer by its identifier.
func (r *V1BetaEntityService) Get(ctx context.Context, entityID string, query V1BetaEntityGetParams, opts ...option.RequestOption) (res *V1BetaEntityGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if entityID == "" {
		err = errors.New("missing required entityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1-beta/customers/%s/entities/%s", query.ID, entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves a paginated list of entities for the given customer.
func (r *V1BetaEntityService) List(ctx context.Context, id string, query V1BetaEntityListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1BetaEntityListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1-beta/customers/%s/entities", id)
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

// Retrieves a paginated list of entities for the given customer.
func (r *V1BetaEntityService) ListAutoPaging(ctx context.Context, id string, query V1BetaEntityListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1BetaEntityListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, id, query, opts...))
}

// Archives entities in bulk for the given customer by id.
func (r *V1BetaEntityService) Archive(ctx context.Context, id string, body V1BetaEntityArchiveParams, opts ...option.RequestOption) (res *V1BetaEntityArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1-beta/customers/%s/entities/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Restores previously archived entities in bulk for the given customer by id.
func (r *V1BetaEntityService) Unarchive(ctx context.Context, id string, body V1BetaEntityUnarchiveParams, opts ...option.RequestOption) (res *V1BetaEntityUnarchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1-beta/customers/%s/entities/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Creates or updates entities in bulk for the given customer. Existing entities
// matched by id are updated; new ids are created.
func (r *V1BetaEntityService) Upsert(ctx context.Context, id string, body V1BetaEntityUpsertParams, opts ...option.RequestOption) (res *V1BetaEntityUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1-beta/customers/%s/entities", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Response object
type V1BetaEntityGetResponse struct {
	// A stored entity instance tracked by the governance service for a given customer
	Data V1BetaEntityGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1BetaEntityGetResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Free-form key/value metadata attached to the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The entity type identifier this entity instantiates
	TypeID string `json:"typeId" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ArchivedAt  respjson.Field
		CreatedAt   respjson.Field
		Metadata    respjson.Field
		TypeID      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1BetaEntityListResponse struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Free-form key/value metadata attached to the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The entity type identifier this entity instantiates
	TypeID string `json:"typeId" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ArchivedAt  respjson.Field
		CreatedAt   respjson.Field
		Metadata    respjson.Field
		TypeID      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapped response echoing the ids that were acted on by an archive/unarchive call
type V1BetaEntityArchiveResponse struct {
	// List of entity identifiers that were acted on
	Data V1BetaEntityArchiveResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entity identifiers that were acted on
type V1BetaEntityArchiveResponseData struct {
	// Entity identifiers to act on
	IDs []string `json:"ids" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IDs         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityArchiveResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityArchiveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapped response echoing the ids that were acted on by an archive/unarchive call
type V1BetaEntityUnarchiveResponse struct {
	// List of entity identifiers that were acted on
	Data V1BetaEntityUnarchiveResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityUnarchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityUnarchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entity identifiers that were acted on
type V1BetaEntityUnarchiveResponseData struct {
	// Entity identifiers to act on
	IDs []string `json:"ids" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IDs         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityUnarchiveResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityUnarchiveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entities created or updated by an upsert request
type V1BetaEntityUpsertResponse struct {
	Data []V1BetaEntityUpsertResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1BetaEntityUpsertResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Free-form key/value metadata attached to the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The entity type identifier this entity instantiates
	TypeID string `json:"typeId" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ArchivedAt  respjson.Field
		CreatedAt   respjson.Field
		Metadata    respjson.Field
		TypeID      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaEntityUpsertResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BetaEntityUpsertResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BetaEntityGetParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type V1BetaEntityListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter results to entities of a specific entity type, by the type's refId
	TypeRefID param.Opt[string] `query:"typeRefId,omitzero" json:"-"`
	// Whether to include archived entities. One of: true, false
	//
	// Any of "true", "false".
	IncludeArchived V1BetaEntityListParamsIncludeArchived `query:"includeArchived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BetaEntityListParams]'s query parameters as `url.Values`.
func (r V1BetaEntityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to include archived entities. One of: true, false
type V1BetaEntityListParamsIncludeArchived string

const (
	V1BetaEntityListParamsIncludeArchivedTrue  V1BetaEntityListParamsIncludeArchived = "true"
	V1BetaEntityListParamsIncludeArchivedFalse V1BetaEntityListParamsIncludeArchived = "false"
)

type V1BetaEntityArchiveParams struct {
	// Entity identifiers to act on
	IDs []string `json:"ids,omitzero" api:"required"`
	paramObj
}

func (r V1BetaEntityArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaEntityArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaEntityArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BetaEntityUnarchiveParams struct {
	// Entity identifiers to act on
	IDs []string `json:"ids,omitzero" api:"required"`
	paramObj
}

func (r V1BetaEntityUnarchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaEntityUnarchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaEntityUnarchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BetaEntityUpsertParams struct {
	// List of entities to create or update (1-100 entries)
	Entities []V1BetaEntityUpsertParamsEntity `json:"entities,omitzero" api:"required"`
	paramObj
}

func (r V1BetaEntityUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaEntityUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaEntityUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entity to create or update.
//
// The property ID is required.
type V1BetaEntityUpsertParamsEntity struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// The entity type refId this entity instantiates. Required when creating a new
	// entity; on a re-upsert may be omitted to preserve the existing type. Governance
	// returns 400 if missing on create.
	TypeRefID param.Opt[string] `json:"typeRefId,omitzero"`
	// Free-form key/value metadata. Patch semantics: empty-string value removes a key,
	// omitted keys are preserved.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1BetaEntityUpsertParamsEntity) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaEntityUpsertParamsEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaEntityUpsertParamsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
