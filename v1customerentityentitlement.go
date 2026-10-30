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

// V1CustomerEntityEntitlementService contains methods and other services that help
// with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerEntityEntitlementService] method instead.
type V1CustomerEntityEntitlementService struct {
	Options []option.RequestOption
}

// NewV1CustomerEntityEntitlementService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1CustomerEntityEntitlementService(opts ...option.RequestOption) (r V1CustomerEntityEntitlementService) {
	r = V1CustomerEntityEntitlementService{}
	r.Options = opts
	return
}

// Returns a cursor-paginated list of entity entitlements for the given customer.
// An entity entitlement ties an entity to a feature or currency with a usage limit
// and reset cadence.
func (r *V1CustomerEntityEntitlementService) List(ctx context.Context, id string, params V1CustomerEntityEntitlementListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1CustomerEntityEntitlementListResponse], err error) {
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
	path := fmt.Sprintf("api/v1/customers/%s/entities/entitlements", id)
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

// Returns a cursor-paginated list of entity entitlements for the given customer.
// An entity entitlement ties an entity to a feature or currency with a usage limit
// and reset cadence.
func (r *V1CustomerEntityEntitlementService) ListAutoPaging(ctx context.Context, id string, params V1CustomerEntityEntitlementListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1CustomerEntityEntitlementListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, id, params, opts...))
}

// Batched create-or-update of entity entitlements. Existing entity entitlements
// matched by (entityId, featureId or currencyId, scopeEntityIds) are updated; new
// ones are created. On update, omitted fields (usageLimit, cadence) are preserved;
// on create both are required.
func (r *V1CustomerEntityEntitlementService) Upsert(ctx context.Context, id string, params V1CustomerEntityEntitlementUpsertParams, opts ...option.RequestOption) (res *V1CustomerEntityEntitlementUpsertResponse, err error) {
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
	path := fmt.Sprintf("api/v1/customers/%s/entities/entitlements", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// An entity entitlement for an entity belonging to a customer. Defines how much of
// the feature or currency the entity may consume (`usageLimit`) and how often the
// counter resets (`cadence`).
type V1CustomerEntityEntitlementListResponse struct {
	// Synthetic UUID identifier — also the cursor anchor for paginated lists
	ID string `json:"id" api:"required" format:"uuid"`
	// Usage-reset cadence as an ISO-8601 single-unit duration, e.g. `P1M`, `P30D`,
	// `PT1M`.
	Cadence string `json:"cadence" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The entity ID this entity entitlement is attached to
	EntityID string `json:"entityId" api:"required"`
	// Parent entity ID in the hierarchy, or `null` for a root.
	ParentID string `json:"parentId" api:"required"`
	// Dimension-scoped sub-limit key: the set of entity IDs this limit applies to.
	// Empty is the node-wide limit that always matches; a non-empty set only applies
	// when every listed entity is present in the resolved set (order-insensitive).
	ScopeEntityIDs []string `json:"scopeEntityIds" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Maximum usage allowed within one cadence window
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// Currency ID this entity entitlement grants (present when it grants a credit
	// currency).
	CurrencyID string `json:"currencyId"`
	// Feature ID this entity entitlement grants (present when it grants a feature).
	FeatureID string `json:"featureId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Cadence        respjson.Field
		CreatedAt      respjson.Field
		EntityID       respjson.Field
		ParentID       respjson.Field
		ScopeEntityIDs respjson.Field
		UpdatedAt      respjson.Field
		UsageLimit     respjson.Field
		CurrencyID     respjson.Field
		FeatureID      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerEntityEntitlementListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerEntityEntitlementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entity entitlements after upsert.
type V1CustomerEntityEntitlementUpsertResponse struct {
	Data []V1CustomerEntityEntitlementUpsertResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerEntityEntitlementUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerEntityEntitlementUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An entity entitlement for an entity belonging to a customer. Defines how much of
// the feature or currency the entity may consume (`usageLimit`) and how often the
// counter resets (`cadence`).
type V1CustomerEntityEntitlementUpsertResponseData struct {
	// Synthetic UUID identifier — also the cursor anchor for paginated lists
	ID string `json:"id" api:"required" format:"uuid"`
	// Usage-reset cadence as an ISO-8601 single-unit duration, e.g. `P1M`, `P30D`,
	// `PT1M`.
	Cadence string `json:"cadence" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The entity ID this entity entitlement is attached to
	EntityID string `json:"entityId" api:"required"`
	// Parent entity ID in the hierarchy, or `null` for a root.
	ParentID string `json:"parentId" api:"required"`
	// Dimension-scoped sub-limit key: the set of entity IDs this limit applies to.
	// Empty is the node-wide limit that always matches; a non-empty set only applies
	// when every listed entity is present in the resolved set (order-insensitive).
	ScopeEntityIDs []string `json:"scopeEntityIds" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Maximum usage allowed within one cadence window
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// Currency ID this entity entitlement grants (present when it grants a credit
	// currency).
	CurrencyID string `json:"currencyId"`
	// Feature ID this entity entitlement grants (present when it grants a feature).
	FeatureID string `json:"featureId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Cadence        respjson.Field
		CreatedAt      respjson.Field
		EntityID       respjson.Field
		ParentID       respjson.Field
		ScopeEntityIDs respjson.Field
		UpdatedAt      respjson.Field
		UsageLimit     respjson.Field
		CurrencyID     respjson.Field
		FeatureID      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerEntityEntitlementUpsertResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerEntityEntitlementUpsertResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerEntityEntitlementListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Filter entity entitlements to a specific currency, by its ID. Mutually exclusive
	// with `featureId`.
	CurrencyID param.Opt[string] `query:"currencyId,omitzero" json:"-"`
	// Filter entity entitlements to a specific entity ID
	EntityID param.Opt[string] `query:"entityId,omitzero" json:"-"`
	// Filter entity entitlements to a specific feature, by its ID. Mutually exclusive
	// with `currencyId`.
	FeatureID param.Opt[string] `query:"featureId,omitzero" json:"-"`
	// Maximum number of items to return
	Limit          param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerEntityEntitlementListParams]'s query parameters
// as `url.Values`.
func (r V1CustomerEntityEntitlementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerEntityEntitlementUpsertParams struct {
	// Entity entitlements to upsert (1–100 per request)
	EntityEntitlements []V1CustomerEntityEntitlementUpsertParamsEntityEntitlement `json:"entityEntitlements,omitzero" api:"required"`
	XAccountID         param.Opt[string]                                          `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID     param.Opt[string]                                          `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1CustomerEntityEntitlementUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerEntityEntitlementUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerEntityEntitlementUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entity entitlement to create or update. Identify the feature or
// currency with exactly one of `featureId` or `currencyId`. The natural key is the
// `(entityId, featureId or currencyId, scopeEntityIds)` triple. On create both
// `usageLimit` and `cadence` are required; on update they may be omitted
// individually to preserve the existing value.
//
// The property EntityID is required.
type V1CustomerEntityEntitlementUpsertParamsEntityEntitlement struct {
	// The entity ID this entity entitlement is attached to
	EntityID string `json:"entityId" api:"required"`
	// Parent entity ID in the hierarchy. Omit to leave the current parent untouched (a
	// new node defaults to a root); `null` detaches to a root; an ID sets or changes
	// the parent. Reparenting or detaching an existing node is leaf-only; place a
	// parent before its children, or in the same batch.
	ParentID param.Opt[string] `json:"parentId,omitzero"`
	// Maximum usage allowed within one cadence window (required on create)
	UsageLimit param.Opt[float64] `json:"usageLimit,omitzero"`
	// Usage-reset cadence (required on create) as an ISO-8601 single-unit duration,
	// e.g. `P1M`, `P30D`, `PT1M`.
	Cadence param.Opt[string] `json:"cadence,omitzero"`
	// Currency ID this entity entitlement grants (credit limits). Mutually exclusive
	// with `featureId`.
	CurrencyID param.Opt[string] `json:"currencyId,omitzero"`
	// Feature ID this entity entitlement grants. Mutually exclusive with `currencyId`.
	FeatureID param.Opt[string] `json:"featureId,omitzero"`
	// Dimension-scoped sub-limit: the set of entity IDs this limit applies to
	// (order-insensitive, unique, at most 5). Omit or leave empty for the node-wide
	// limit.
	ScopeEntityIDs []string `json:"scopeEntityIds,omitzero"`
	paramObj
}

func (r V1CustomerEntityEntitlementUpsertParamsEntityEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerEntityEntitlementUpsertParamsEntityEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerEntityEntitlementUpsertParamsEntityEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
