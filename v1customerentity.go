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

// V1CustomerEntityService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerEntityService] method instead.
type V1CustomerEntityService struct {
	Options      []option.RequestOption
	Entitlements V1CustomerEntityEntitlementService
}

// NewV1CustomerEntityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerEntityService(opts ...option.RequestOption) (r V1CustomerEntityService) {
	r = V1CustomerEntityService{}
	r.Options = opts
	r.Entitlements = NewV1CustomerEntityEntitlementService(opts...)
	return
}

// Retrieves a single entity for the given customer by its identifier.
func (r *V1CustomerEntityService) Get(ctx context.Context, entityID string, params V1CustomerEntityGetParams, opts ...option.RequestOption) (res *V1CustomerEntityGetResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if entityID == "" {
		err = errors.New("missing required entityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/entities/%s", params.ID, entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves a paginated list of entities for the given customer.
func (r *V1CustomerEntityService) List(ctx context.Context, id string, params V1CustomerEntityListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1CustomerEntityListResponse], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/entities", id)
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

// Retrieves a paginated list of entities for the given customer.
func (r *V1CustomerEntityService) ListAutoPaging(ctx context.Context, id string, params V1CustomerEntityListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1CustomerEntityListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, id, params, opts...))
}

// Archives entities in bulk for the given customer by id. Upsert an archived
// entity to restore it.
func (r *V1CustomerEntityService) Archive(ctx context.Context, id string, params V1CustomerEntityArchiveParams, opts ...option.RequestOption) (res *V1CustomerEntityArchiveResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/entities/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Creates or updates entities in bulk for the given customer. Existing entities
// matched by id are updated; new ids are created. An upsert of an archived entity
// restores it. An entity on an archived entity type fails with 400.
func (r *V1CustomerEntityService) Upsert(ctx context.Context, id string, params V1CustomerEntityUpsertParams, opts ...option.RequestOption) (res *V1CustomerEntityUpsertResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/entities", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Response object
type V1CustomerEntityGetResponse struct {
	// A stored entity instance tracked by the governance service for a given customer
	Data V1CustomerEntityGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerEntityGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerEntityGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1CustomerEntityGetResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Human-readable name for the entity, or null when none is set — in which case
	// clients display the entity ID
	DisplayName string `json:"displayName" api:"required"`
	// The entity type identifier this entity instantiates
	EntityTypeID string `json:"entityTypeId" api:"required"`
	// Free-form key/value metadata attached to the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ArchivedAt   respjson.Field
		CreatedAt    respjson.Field
		DisplayName  respjson.Field
		EntityTypeID respjson.Field
		Metadata     respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerEntityGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerEntityGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1CustomerEntityListResponse struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Human-readable name for the entity, or null when none is set — in which case
	// clients display the entity ID
	DisplayName string `json:"displayName" api:"required"`
	// The entity type identifier this entity instantiates
	EntityTypeID string `json:"entityTypeId" api:"required"`
	// Free-form key/value metadata attached to the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ArchivedAt   respjson.Field
		CreatedAt    respjson.Field
		DisplayName  respjson.Field
		EntityTypeID respjson.Field
		Metadata     respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerEntityListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerEntityListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapped response echoing the ids that were acted on by an archive/unarchive call
type V1CustomerEntityArchiveResponse struct {
	// List of entity identifiers that were acted on
	Data V1CustomerEntityArchiveResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerEntityArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerEntityArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entity identifiers that were acted on
type V1CustomerEntityArchiveResponseData struct {
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
func (r V1CustomerEntityArchiveResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerEntityArchiveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entities created or updated by an upsert request
type V1CustomerEntityUpsertResponse struct {
	Data []V1CustomerEntityUpsertResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerEntityUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerEntityUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1CustomerEntityUpsertResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Human-readable name for the entity, or null when none is set — in which case
	// clients display the entity ID
	DisplayName string `json:"displayName" api:"required"`
	// The entity type identifier this entity instantiates
	EntityTypeID string `json:"entityTypeId" api:"required"`
	// Free-form key/value metadata attached to the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ArchivedAt   respjson.Field
		CreatedAt    respjson.Field
		DisplayName  respjson.Field
		EntityTypeID respjson.Field
		Metadata     respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerEntityUpsertResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerEntityUpsertResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerEntityGetParams struct {
	ID             string            `path:"id" api:"required" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

type V1CustomerEntityListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Filter results to entities of a specific entity type, by the type's ID
	EntityTypeID param.Opt[string] `query:"entityTypeId,omitzero" json:"-"`
	// Maximum number of items to return
	Limit          param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	// Whether to include archived entities. One of: true, false
	//
	// Any of "true", "false".
	IncludeArchived V1CustomerEntityListParamsIncludeArchived `query:"includeArchived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerEntityListParams]'s query parameters as
// `url.Values`.
func (r V1CustomerEntityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to include archived entities. One of: true, false
type V1CustomerEntityListParamsIncludeArchived string

const (
	V1CustomerEntityListParamsIncludeArchivedTrue  V1CustomerEntityListParamsIncludeArchived = "true"
	V1CustomerEntityListParamsIncludeArchivedFalse V1CustomerEntityListParamsIncludeArchived = "false"
)

type V1CustomerEntityArchiveParams struct {
	// Entity identifiers to act on
	IDs            []string          `json:"ids,omitzero" api:"required"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1CustomerEntityArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerEntityArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerEntityArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerEntityUpsertParams struct {
	// List of entities to create or update (1-100 entries)
	Entities       []V1CustomerEntityUpsertParamsEntity `json:"entities,omitzero" api:"required"`
	XAccountID     param.Opt[string]                    `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]                    `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1CustomerEntityUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerEntityUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerEntityUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entity to create or update.
//
// The property ID is required.
type V1CustomerEntityUpsertParamsEntity struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Human-readable name for the entity. Omit to preserve the stored value, or send
	// an empty string or null to clear it.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// The entity type ID this entity instantiates. Required when creating a new
	// entity; on a re-upsert may be omitted to preserve the existing type. Governance
	// returns 400 if missing on create.
	EntityTypeID param.Opt[string] `json:"entityTypeId,omitzero"`
	// Free-form key/value metadata. Patch semantics: empty-string value removes a key,
	// omitted keys are preserved.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1CustomerEntityUpsertParamsEntity) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerEntityUpsertParamsEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerEntityUpsertParamsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
