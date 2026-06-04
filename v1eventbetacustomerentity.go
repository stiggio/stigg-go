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

// V1EventBetaCustomerEntityService contains methods and other services that help
// with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventBetaCustomerEntityService] method instead.
type V1EventBetaCustomerEntityService struct {
	Options []option.RequestOption
}

// NewV1EventBetaCustomerEntityService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1EventBetaCustomerEntityService(opts ...option.RequestOption) (r V1EventBetaCustomerEntityService) {
	r = V1EventBetaCustomerEntityService{}
	r.Options = opts
	return
}

// Retrieves a single entity for the given customer by its identifier.
func (r *V1EventBetaCustomerEntityService) Get(ctx context.Context, entityID string, query V1EventBetaCustomerEntityGetParams, opts ...option.RequestOption) (res *V1EventBetaCustomerEntityGetResponse, err error) {
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
func (r *V1EventBetaCustomerEntityService) List(ctx context.Context, id string, query V1EventBetaCustomerEntityListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1EventBetaCustomerEntityListResponse], err error) {
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
func (r *V1EventBetaCustomerEntityService) ListAutoPaging(ctx context.Context, id string, query V1EventBetaCustomerEntityListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1EventBetaCustomerEntityListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, id, query, opts...))
}

// Archives entities in bulk for the given customer by id.
func (r *V1EventBetaCustomerEntityService) Archive(ctx context.Context, id string, body V1EventBetaCustomerEntityArchiveParams, opts ...option.RequestOption) (res *V1EventBetaCustomerEntityArchiveResponse, err error) {
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
func (r *V1EventBetaCustomerEntityService) Unarchive(ctx context.Context, id string, body V1EventBetaCustomerEntityUnarchiveParams, opts ...option.RequestOption) (res *V1EventBetaCustomerEntityUnarchiveResponse, err error) {
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
func (r *V1EventBetaCustomerEntityService) Upsert(ctx context.Context, id string, body V1EventBetaCustomerEntityUpsertParams, opts ...option.RequestOption) (res *V1EventBetaCustomerEntityUpsertResponse, err error) {
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
type V1EventBetaCustomerEntityGetResponse struct {
	// A stored entity instance tracked by the governance service for a given customer
	Data V1EventBetaCustomerEntityGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerEntityGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntityGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1EventBetaCustomerEntityGetResponseData struct {
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
func (r V1EventBetaCustomerEntityGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntityGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1EventBetaCustomerEntityListResponse struct {
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
func (r V1EventBetaCustomerEntityListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntityListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapped response echoing the ids that were acted on by an archive/unarchive call
type V1EventBetaCustomerEntityArchiveResponse struct {
	// List of entity identifiers that were acted on
	Data V1EventBetaCustomerEntityArchiveResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerEntityArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntityArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entity identifiers that were acted on
type V1EventBetaCustomerEntityArchiveResponseData struct {
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
func (r V1EventBetaCustomerEntityArchiveResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntityArchiveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapped response echoing the ids that were acted on by an archive/unarchive call
type V1EventBetaCustomerEntityUnarchiveResponse struct {
	// List of entity identifiers that were acted on
	Data V1EventBetaCustomerEntityUnarchiveResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerEntityUnarchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntityUnarchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entity identifiers that were acted on
type V1EventBetaCustomerEntityUnarchiveResponseData struct {
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
func (r V1EventBetaCustomerEntityUnarchiveResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntityUnarchiveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entities created or updated by an upsert request
type V1EventBetaCustomerEntityUpsertResponse struct {
	Data []V1EventBetaCustomerEntityUpsertResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerEntityUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntityUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1EventBetaCustomerEntityUpsertResponseData struct {
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
func (r V1EventBetaCustomerEntityUpsertResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntityUpsertResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventBetaCustomerEntityGetParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type V1EventBetaCustomerEntityListParams struct {
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
	IncludeArchived V1EventBetaCustomerEntityListParamsIncludeArchived `query:"includeArchived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventBetaCustomerEntityListParams]'s query parameters as
// `url.Values`.
func (r V1EventBetaCustomerEntityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to include archived entities. One of: true, false
type V1EventBetaCustomerEntityListParamsIncludeArchived string

const (
	V1EventBetaCustomerEntityListParamsIncludeArchivedTrue  V1EventBetaCustomerEntityListParamsIncludeArchived = "true"
	V1EventBetaCustomerEntityListParamsIncludeArchivedFalse V1EventBetaCustomerEntityListParamsIncludeArchived = "false"
)

type V1EventBetaCustomerEntityArchiveParams struct {
	// Entity identifiers to act on
	IDs []string `json:"ids,omitzero" api:"required"`
	paramObj
}

func (r V1EventBetaCustomerEntityArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventBetaCustomerEntityArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventBetaCustomerEntityArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventBetaCustomerEntityUnarchiveParams struct {
	// Entity identifiers to act on
	IDs []string `json:"ids,omitzero" api:"required"`
	paramObj
}

func (r V1EventBetaCustomerEntityUnarchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventBetaCustomerEntityUnarchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventBetaCustomerEntityUnarchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventBetaCustomerEntityUpsertParams struct {
	// List of entities to create or update (1-100 entries)
	Entities []V1EventBetaCustomerEntityUpsertParamsEntity `json:"entities,omitzero" api:"required"`
	paramObj
}

func (r V1EventBetaCustomerEntityUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventBetaCustomerEntityUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventBetaCustomerEntityUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entity to create or update.
//
// The property ID is required.
type V1EventBetaCustomerEntityUpsertParamsEntity struct {
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

func (r V1EventBetaCustomerEntityUpsertParamsEntity) MarshalJSON() (data []byte, err error) {
	type shadow V1EventBetaCustomerEntityUpsertParamsEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventBetaCustomerEntityUpsertParamsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
