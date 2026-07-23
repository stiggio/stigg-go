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
	"github.com/stiggio/stigg-go/packages/param"
	"github.com/stiggio/stigg-go/packages/respjson"
)

// V1EventBetaCustomerService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventBetaCustomerService] method instead.
type V1EventBetaCustomerService struct {
	Options []option.RequestOption
}

// NewV1EventBetaCustomerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1EventBetaCustomerService(opts ...option.RequestOption) (r V1EventBetaCustomerService) {
	r = V1EventBetaCustomerService{}
	r.Options = opts
	return
}

// Queries the customer's governance hierarchy tree, returning a cursor-paginated
// list of nodes with their usage configuration (limit, cadence, scope) and current
// usage, sortable and filterable by usage. Each node carries `parentId` so the
// tree can be rebuilt client-side. Usage is read from a periodically-refreshed
// read model and never gates access.
func (r *V1EventBetaCustomerService) GetGovernance(ctx context.Context, id string, params V1EventBetaCustomerGetGovernanceParams, opts ...option.RequestOption) (res *V1EventBetaCustomerGetGovernanceResponse, err error) {
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
	path := fmt.Sprintf("api/v1-beta/customers/%s/governance", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Paginated list of governance tree nodes, each with its usage configuration and
// current usage.
type V1EventBetaCustomerGetGovernanceResponse struct {
	Data       []V1EventBetaCustomerGetGovernanceResponseData     `json:"data" api:"required"`
	Pagination V1EventBetaCustomerGetGovernanceResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerGetGovernanceResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerGetGovernanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A node of the governance hierarchy tree with its usage configuration (limit,
// cadence, scope) and current usage. Usage is read from a periodically-refreshed
// read model and may lag the live counter by a short interval; it never gates
// access.
type V1EventBetaCustomerGetGovernanceResponseData struct {
	// Usage-reset cadence as an ISO-8601 single-unit duration, e.g. `P1M`, `P30D`,
	// `PT1M`; `null` when the node has no usage configuration.
	Cadence string `json:"cadence" api:"required"`
	// Usage consumed in the current cadence period (may lag the live counter by a
	// short interval).
	CurrentUsage float64 `json:"currentUsage" api:"required"`
	// External id of the entity at this node.
	EntityID string `json:"entityId" api:"required"`
	// External id of the entity type (e.g. `team`, `user`).
	EntityTypeID string `json:"entityTypeId" api:"required"`
	// External id of the parent entity in the tree; `null` for a root. Use it to
	// rebuild the tree.
	ParentID string `json:"parentId" api:"required"`
	// The configuration scope (entity ids). Empty is the node-wide configuration; a
	// non-empty set is a dimension-scoped sub-configuration.
	ScopeEntityIDs []string `json:"scopeEntityIds" api:"required"`
	// Hard usage limit for this node per cadence period.
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// Exclusive end of the cadence period — when usage resets; `null` once the period
	// has rolled over.
	UsagePeriodEnd time.Time `json:"usagePeriodEnd" api:"required" format:"date-time"`
	// Start of the cadence period the usage snapshot belongs to; `null` once the
	// period has rolled over.
	UsagePeriodStart time.Time `json:"usagePeriodStart" api:"required" format:"date-time"`
	// `currentUsage / usageLimit` (1 when usageLimit is 0 — always at limit). The
	// cross-capability-safe sort key.
	Utilization float64 `json:"utilization" api:"required"`
	// The metered currency ID (present when the configured capability is a credit
	// currency).
	CurrencyID string `json:"currencyId"`
	// The metered feature ID (present when the configured capability is a feature).
	FeatureID string `json:"featureId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cadence          respjson.Field
		CurrentUsage     respjson.Field
		EntityID         respjson.Field
		EntityTypeID     respjson.Field
		ParentID         respjson.Field
		ScopeEntityIDs   respjson.Field
		UsageLimit       respjson.Field
		UsagePeriodEnd   respjson.Field
		UsagePeriodStart respjson.Field
		Utilization      respjson.Field
		CurrencyID       respjson.Field
		FeatureID        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerGetGovernanceResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerGetGovernanceResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventBetaCustomerGetGovernanceResponsePagination struct {
	// Cursor for fetching the next page of results, or null if no additional pages
	// exist
	Next string `json:"next" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerGetGovernanceResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerGetGovernanceResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventBetaCustomerGetGovernanceParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Case-insensitive substring match on the entity id (`%`/`_` matched literally).
	EntityIDSearch param.Opt[string] `query:"entityIdSearch,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only nodes with utilization ≥ this value (e.g. 0.8 for ≥80%, 1 for at/over
	// limit).
	MinUtilization param.Opt[float64] `query:"minUtilization,omitzero" json:"-"`
	XAccountID     param.Opt[string]  `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]  `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	// Currency ids to include, repeated per value (e.g. `?currencyIds=credits`). Omit
	// both featureIds and currencyIds for tree mode.
	CurrencyIDs []string `query:"currencyIds,omitzero" json:"-"`
	// Filter to one or more entity types, repeated per value (e.g.
	// `?entityTypeIds=team&entityTypeIds=user`).
	EntityTypeIDs []string `query:"entityTypeIds,omitzero" json:"-"`
	// Feature ids to include, repeated per value (e.g.
	// `?featureIds=ai-tokens&featureIds=seats`). Omit both featureIds and currencyIds
	// for tree mode — every node in the hierarchy with no usage configuration
	// attached.
	FeatureIDs []string `query:"featureIds,omitzero" json:"-"`
	// Sort direction: `asc` or `desc` (default `desc`).
	//
	// Any of "asc", "desc".
	Order V1EventBetaCustomerGetGovernanceParamsOrder `query:"order,omitzero" json:"-"`
	// Filter by configuration scope: `all` (default), `nodeWide` (`[]` only), or
	// `scoped` (non-empty only).
	//
	// Any of "all", "nodeWide", "scoped".
	Scope V1EventBetaCustomerGetGovernanceParamsScope `query:"scope,omitzero" json:"-"`
	// Sort key: `utilization` (default, cross-capability-safe), `currentUsage`,
	// `usageLimit`, `scopeSize`, `id`, or `createdAt`.
	//
	// Any of "utilization", "currentUsage", "usageLimit", "scopeSize", "id",
	// "createdAt".
	SortBy V1EventBetaCustomerGetGovernanceParamsSortBy `query:"sortBy,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventBetaCustomerGetGovernanceParams]'s query parameters
// as `url.Values`.
func (r V1EventBetaCustomerGetGovernanceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction: `asc` or `desc` (default `desc`).
type V1EventBetaCustomerGetGovernanceParamsOrder string

const (
	V1EventBetaCustomerGetGovernanceParamsOrderAsc  V1EventBetaCustomerGetGovernanceParamsOrder = "asc"
	V1EventBetaCustomerGetGovernanceParamsOrderDesc V1EventBetaCustomerGetGovernanceParamsOrder = "desc"
)

// Filter by configuration scope: `all` (default), `nodeWide` (`[]` only), or
// `scoped` (non-empty only).
type V1EventBetaCustomerGetGovernanceParamsScope string

const (
	V1EventBetaCustomerGetGovernanceParamsScopeAll      V1EventBetaCustomerGetGovernanceParamsScope = "all"
	V1EventBetaCustomerGetGovernanceParamsScopeNodeWide V1EventBetaCustomerGetGovernanceParamsScope = "nodeWide"
	V1EventBetaCustomerGetGovernanceParamsScopeScoped   V1EventBetaCustomerGetGovernanceParamsScope = "scoped"
)

// Sort key: `utilization` (default, cross-capability-safe), `currentUsage`,
// `usageLimit`, `scopeSize`, `id`, or `createdAt`.
type V1EventBetaCustomerGetGovernanceParamsSortBy string

const (
	V1EventBetaCustomerGetGovernanceParamsSortByUtilization  V1EventBetaCustomerGetGovernanceParamsSortBy = "utilization"
	V1EventBetaCustomerGetGovernanceParamsSortByCurrentUsage V1EventBetaCustomerGetGovernanceParamsSortBy = "currentUsage"
	V1EventBetaCustomerGetGovernanceParamsSortByUsageLimit   V1EventBetaCustomerGetGovernanceParamsSortBy = "usageLimit"
	V1EventBetaCustomerGetGovernanceParamsSortByScopeSize    V1EventBetaCustomerGetGovernanceParamsSortBy = "scopeSize"
	V1EventBetaCustomerGetGovernanceParamsSortByID           V1EventBetaCustomerGetGovernanceParamsSortBy = "id"
	V1EventBetaCustomerGetGovernanceParamsSortByCreatedAt    V1EventBetaCustomerGetGovernanceParamsSortBy = "createdAt"
)
