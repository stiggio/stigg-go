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

// V1BetaCustomerGovernanceService contains methods and other services that help
// with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BetaCustomerGovernanceService] method instead.
type V1BetaCustomerGovernanceService struct {
	Options []option.RequestOption
}

// NewV1BetaCustomerGovernanceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1BetaCustomerGovernanceService(opts ...option.RequestOption) (r V1BetaCustomerGovernanceService) {
	r = V1BetaCustomerGovernanceService{}
	r.Options = opts
	return
}

// Queries the customer's governance hierarchy tree, returning a cursor-paginated
// list of nodes with their usage configuration (limit, cadence, scope) and current
// usage, sortable and filterable by usage. Each node carries `parentId` so the
// tree can be rebuilt client-side. Usage is read from a periodically-refreshed
// read model and never gates access.
func (r *V1BetaCustomerGovernanceService) List(ctx context.Context, id string, params V1BetaCustomerGovernanceListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1BetaCustomerGovernanceListResponse], err error) {
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
	path := fmt.Sprintf("api/v1-beta/customers/%s/governance", id)
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

// Queries the customer's governance hierarchy tree, returning a cursor-paginated
// list of nodes with their usage configuration (limit, cadence, scope) and current
// usage, sortable and filterable by usage. Each node carries `parentId` so the
// tree can be rebuilt client-side. Usage is read from a periodically-refreshed
// read model and never gates access.
func (r *V1BetaCustomerGovernanceService) ListAutoPaging(ctx context.Context, id string, params V1BetaCustomerGovernanceListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1BetaCustomerGovernanceListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, id, params, opts...))
}

// A node of the governance hierarchy tree with its usage configuration (limit,
// cadence, scope) and current usage. Usage is read from a periodically-refreshed
// read model and may lag the live counter by a short interval; it never gates
// access.
type V1BetaCustomerGovernanceListResponse struct {
	// Usage-reset cadence as an ISO-8601 single-unit duration, e.g. `P1M`, `P30D`,
	// `PT1M`; `null` when the node has no usage configuration.
	Cadence string `json:"cadence" api:"required"`
	// Usage consumed in the current cadence period (may lag the live counter by a
	// short interval).
	CurrentUsage float64 `json:"currentUsage" api:"required"`
	// Human-readable name of the entity, or null when none is set (display the entity
	// id instead).
	DisplayName string `json:"displayName" api:"required"`
	// External id of the entity at this node.
	EntityID string `json:"entityId" api:"required"`
	// External id of the entity type (e.g. `team`, `user`).
	EntityTypeID string `json:"entityTypeId" api:"required"`
	// External id of the parent entity in the tree. `null` means the entity is either
	// a root or not yet placed in the hierarchy — placement rides on an assignment, so
	// an entity with no limits set has no parent yet. Both render at the top level;
	// use it to rebuild the tree.
	ParentID string `json:"parentId" api:"required"`
	// The configuration scope (entity ids). Empty is the node-wide configuration; a
	// non-empty set is a dimension-scoped sub-configuration.
	ScopeEntityIDs []string `json:"scopeEntityIds" api:"required"`
	// Hard usage limit for this node per cadence period.
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// Exclusive end of the cadence period in progress now — when usage resets. `null`
	// when the node has no usage configuration, or when a stored cadence cannot be
	// parsed.
	UsagePeriodEnd time.Time `json:"usagePeriodEnd" api:"required" format:"date-time"`
	// Start of the cadence period in progress now, derived from the cadence and the
	// assignment anchor — it stays correct across a rollover. `null` when the node has
	// no usage configuration, or when a stored cadence cannot be parsed.
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
		DisplayName      respjson.Field
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
func (r V1BetaCustomerGovernanceListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerGovernanceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BetaCustomerGovernanceListParams struct {
	// Opaque cursor from a page’s `next`; returns the page after it. Mutually
	// exclusive with `before`.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Opaque cursor from a page’s `prev`; returns the page before it. Mutually
	// exclusive with `after`.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Case-insensitive substring match on the entity id or its display name (`%`/`_`
	// matched literally).
	EntityIDSearch param.Opt[string] `query:"entityIdSearch,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only nodes with utilization ≥ this value (e.g. 0.8 for ≥80%, 1 for at/over
	// limit).
	MinUtilization param.Opt[float64] `query:"minUtilization,omitzero" json:"-"`
	XAccountID     param.Opt[string]  `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]  `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	// Currency ids to include, comma-separated (e.g. `?currencyIds=credits`), up
	// to 100. Omit both featureIds and currencyIds for tree mode.
	CurrencyIDs []string `query:"currencyIds,omitzero" json:"-"`
	// Filter to one or more entity types, comma-separated (e.g.
	// `?entityTypeIds=team,user`), up to 100.
	EntityTypeIDs []string `query:"entityTypeIds,omitzero" json:"-"`
	// Feature ids to include, comma-separated (e.g. `?featureIds=ai-tokens,seats`), up
	// to 100. Omit both featureIds and currencyIds for tree mode — every node in the
	// hierarchy with no usage configuration attached.
	FeatureIDs []string `query:"featureIds,omitzero" json:"-"`
	// Sort direction: `asc` or `desc` (default `desc`).
	//
	// Any of "asc", "desc".
	Order V1BetaCustomerGovernanceListParamsOrder `query:"order,omitzero" json:"-"`
	// Filter by configuration scope: `all` (default), `nodeWide` (`[]` only), or
	// `scoped` (non-empty only).
	//
	// Any of "all", "nodeWide", "scoped".
	Scope V1BetaCustomerGovernanceListParamsScope `query:"scope,omitzero" json:"-"`
	// Sort key: `utilization` (default, cross-capability-safe), `currentUsage`,
	// `usageLimit`, `scopeSize`, `id`, or `createdAt`.
	//
	// Any of "utilization", "currentUsage", "usageLimit", "scopeSize", "id",
	// "createdAt".
	SortBy V1BetaCustomerGovernanceListParamsSortBy `query:"sortBy,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BetaCustomerGovernanceListParams]'s query parameters as
// `url.Values`.
func (r V1BetaCustomerGovernanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction: `asc` or `desc` (default `desc`).
type V1BetaCustomerGovernanceListParamsOrder string

const (
	V1BetaCustomerGovernanceListParamsOrderAsc  V1BetaCustomerGovernanceListParamsOrder = "asc"
	V1BetaCustomerGovernanceListParamsOrderDesc V1BetaCustomerGovernanceListParamsOrder = "desc"
)

// Filter by configuration scope: `all` (default), `nodeWide` (`[]` only), or
// `scoped` (non-empty only).
type V1BetaCustomerGovernanceListParamsScope string

const (
	V1BetaCustomerGovernanceListParamsScopeAll      V1BetaCustomerGovernanceListParamsScope = "all"
	V1BetaCustomerGovernanceListParamsScopeNodeWide V1BetaCustomerGovernanceListParamsScope = "nodeWide"
	V1BetaCustomerGovernanceListParamsScopeScoped   V1BetaCustomerGovernanceListParamsScope = "scoped"
)

// Sort key: `utilization` (default, cross-capability-safe), `currentUsage`,
// `usageLimit`, `scopeSize`, `id`, or `createdAt`.
type V1BetaCustomerGovernanceListParamsSortBy string

const (
	V1BetaCustomerGovernanceListParamsSortByUtilization  V1BetaCustomerGovernanceListParamsSortBy = "utilization"
	V1BetaCustomerGovernanceListParamsSortByCurrentUsage V1BetaCustomerGovernanceListParamsSortBy = "currentUsage"
	V1BetaCustomerGovernanceListParamsSortByUsageLimit   V1BetaCustomerGovernanceListParamsSortBy = "usageLimit"
	V1BetaCustomerGovernanceListParamsSortByScopeSize    V1BetaCustomerGovernanceListParamsSortBy = "scopeSize"
	V1BetaCustomerGovernanceListParamsSortByID           V1BetaCustomerGovernanceListParamsSortBy = "id"
	V1BetaCustomerGovernanceListParamsSortByCreatedAt    V1BetaCustomerGovernanceListParamsSortBy = "createdAt"
)
