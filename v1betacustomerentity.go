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

// V1BetaCustomerEntityService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BetaCustomerEntityService] method instead.
type V1BetaCustomerEntityService struct {
	Options []option.RequestOption
}

// NewV1BetaCustomerEntityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1BetaCustomerEntityService(opts ...option.RequestOption) (r V1BetaCustomerEntityService) {
	r = V1BetaCustomerEntityService{}
	r.Options = opts
	return
}

// Retrieves a single entity for the given customer by its identifier.
func (r *V1BetaCustomerEntityService) Get(ctx context.Context, entityID string, params V1BetaCustomerEntityGetParams, opts ...option.RequestOption) (res *V1BetaCustomerEntityGetResponse, err error) {
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
	path := fmt.Sprintf("api/v1-beta/customers/%s/entities/%s", params.ID, entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves a paginated list of entities for the given customer.
func (r *V1BetaCustomerEntityService) List(ctx context.Context, id string, params V1BetaCustomerEntityListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1BetaCustomerEntityListResponse], err error) {
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
	path := fmt.Sprintf("api/v1-beta/customers/%s/entities", id)
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
func (r *V1BetaCustomerEntityService) ListAutoPaging(ctx context.Context, id string, params V1BetaCustomerEntityListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1BetaCustomerEntityListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, id, params, opts...))
}

// Archives entities in bulk for the given customer by id.
func (r *V1BetaCustomerEntityService) Archive(ctx context.Context, id string, params V1BetaCustomerEntityArchiveParams, opts ...option.RequestOption) (res *V1BetaCustomerEntityArchiveResponse, err error) {
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
	path := fmt.Sprintf("api/v1-beta/customers/%s/entities/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restores previously archived entities in bulk for the given customer by id.
func (r *V1BetaCustomerEntityService) Unarchive(ctx context.Context, id string, params V1BetaCustomerEntityUnarchiveParams, opts ...option.RequestOption) (res *V1BetaCustomerEntityUnarchiveResponse, err error) {
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
	path := fmt.Sprintf("api/v1-beta/customers/%s/entities/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Creates or updates entities in bulk for the given customer. Existing entities
// matched by id are updated; new ids are created.
func (r *V1BetaCustomerEntityService) Upsert(ctx context.Context, id string, params V1BetaCustomerEntityUpsertParams, opts ...option.RequestOption) (res *V1BetaCustomerEntityUpsertResponse, err error) {
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
	path := fmt.Sprintf("api/v1-beta/customers/%s/entities", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Response object
type V1BetaCustomerEntityGetResponse struct {
	// A stored entity instance tracked by the governance service for a given customer
	Data V1BetaCustomerEntityGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaCustomerEntityGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerEntityGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1BetaCustomerEntityGetResponseData struct {
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
func (r V1BetaCustomerEntityGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerEntityGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1BetaCustomerEntityListResponse struct {
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
func (r V1BetaCustomerEntityListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerEntityListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapped response echoing the ids that were acted on by an archive/unarchive call
type V1BetaCustomerEntityArchiveResponse struct {
	// List of entity identifiers that were acted on
	Data V1BetaCustomerEntityArchiveResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaCustomerEntityArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerEntityArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entity identifiers that were acted on
type V1BetaCustomerEntityArchiveResponseData struct {
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
func (r V1BetaCustomerEntityArchiveResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerEntityArchiveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapped response echoing the ids that were acted on by an archive/unarchive call
type V1BetaCustomerEntityUnarchiveResponse struct {
	// List of entity identifiers that were acted on
	Data V1BetaCustomerEntityUnarchiveResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaCustomerEntityUnarchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerEntityUnarchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entity identifiers that were acted on
type V1BetaCustomerEntityUnarchiveResponseData struct {
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
func (r V1BetaCustomerEntityUnarchiveResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerEntityUnarchiveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entities created or updated by an upsert request
type V1BetaCustomerEntityUpsertResponse struct {
	Data []V1BetaCustomerEntityUpsertResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaCustomerEntityUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerEntityUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored entity instance tracked by the governance service for a given customer
type V1BetaCustomerEntityUpsertResponseData struct {
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
func (r V1BetaCustomerEntityUpsertResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerEntityUpsertResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BetaCustomerEntityGetParams struct {
	ID             string            `path:"id" api:"required" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

type V1BetaCustomerEntityListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter results to entities of a specific entity type, by the type's refId
	TypeRefID      param.Opt[string] `query:"typeRefId,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	// Whether to include archived entities. One of: true, false
	//
	// Any of "true", "false".
	IncludeArchived V1BetaCustomerEntityListParamsIncludeArchived `query:"includeArchived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BetaCustomerEntityListParams]'s query parameters as
// `url.Values`.
func (r V1BetaCustomerEntityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to include archived entities. One of: true, false
type V1BetaCustomerEntityListParamsIncludeArchived string

const (
	V1BetaCustomerEntityListParamsIncludeArchivedTrue  V1BetaCustomerEntityListParamsIncludeArchived = "true"
	V1BetaCustomerEntityListParamsIncludeArchivedFalse V1BetaCustomerEntityListParamsIncludeArchived = "false"
)

type V1BetaCustomerEntityArchiveParams struct {
	// Entity identifiers to act on
	IDs            []string          `json:"ids,omitzero" api:"required"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1BetaCustomerEntityArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaCustomerEntityArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaCustomerEntityArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BetaCustomerEntityUnarchiveParams struct {
	// Entity identifiers to act on
	IDs            []string          `json:"ids,omitzero" api:"required"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1BetaCustomerEntityUnarchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaCustomerEntityUnarchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaCustomerEntityUnarchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BetaCustomerEntityUpsertParams struct {
	// List of entities to create or update (1-100 entries)
	Entities       []V1BetaCustomerEntityUpsertParamsEntity `json:"entities,omitzero" api:"required"`
	XAccountID     param.Opt[string]                        `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]                        `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1BetaCustomerEntityUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaCustomerEntityUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaCustomerEntityUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entity to create or update.
//
// The property ID is required.
type V1BetaCustomerEntityUpsertParamsEntity struct {
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

func (r V1BetaCustomerEntityUpsertParamsEntity) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaCustomerEntityUpsertParamsEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaCustomerEntityUpsertParamsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
