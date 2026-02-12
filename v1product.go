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

// V1ProductService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ProductService] method instead.
type V1ProductService struct {
	Options []option.RequestOption
}

// NewV1ProductService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ProductService(opts ...option.RequestOption) (r V1ProductService) {
	r = V1ProductService{}
	r.Options = opts
	return
}

// Retrieves a paginated list of products in the environment.
func (r *V1ProductService) ListProducts(ctx context.Context, query V1ProductListProductsParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1ProductListProductsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/products"
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

// Retrieves a paginated list of products in the environment.
func (r *V1ProductService) ListProductsAutoPaging(ctx context.Context, query V1ProductListProductsParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1ProductListProductsResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListProducts(ctx, query, opts...))
}

// Product configuration object
type V1ProductListProductsResponse struct {
	// The unique identifier for the entity
	ID string `json:"id,required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Description of the product
	Description string `json:"description,required"`
	// Display name of the product
	DisplayName string `json:"displayName,required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,required"`
	// Indicates if multiple subscriptions to this product are allowed
	MultipleSubscriptions bool `json:"multipleSubscriptions,required"`
	// The status of the product
	//
	// Any of "PUBLISHED", "ARCHIVED".
	Status V1ProductListProductsResponseStatus `json:"status,required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Product behavior settings for subscription lifecycle management.
	ProductSettings V1ProductListProductsResponseProductSettings `json:"productSettings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		Description           respjson.Field
		DisplayName           respjson.Field
		Metadata              respjson.Field
		MultipleSubscriptions respjson.Field
		Status                respjson.Field
		UpdatedAt             respjson.Field
		ProductSettings       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ProductListProductsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ProductListProductsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the product
type V1ProductListProductsResponseStatus string

const (
	V1ProductListProductsResponseStatusPublished V1ProductListProductsResponseStatus = "PUBLISHED"
	V1ProductListProductsResponseStatusArchived  V1ProductListProductsResponseStatus = "ARCHIVED"
)

// Product behavior settings for subscription lifecycle management.
type V1ProductListProductsResponseProductSettings struct {
	// Time when the subscription will be cancelled
	//
	// Any of "END_OF_BILLING_PERIOD", "IMMEDIATE", "SPECIFIC_DATE".
	SubscriptionCancellationTime string `json:"subscriptionCancellationTime,required"`
	// Setup for the end of the subscription
	//
	// Any of "DOWNGRADE_TO_FREE", "CANCEL_SUBSCRIPTION".
	SubscriptionEndSetup string `json:"subscriptionEndSetup,required"`
	// Setup for the start of the subscription
	//
	// Any of "PLAN_SELECTION", "TRIAL_PERIOD", "FREE_PLAN".
	SubscriptionStartSetup string `json:"subscriptionStartSetup,required"`
	// ID of the plan to downgrade to at the end of the billing period
	DowngradePlanID string `json:"downgradePlanId,nullable"`
	// Indicates if the subscription should be prorated at the end of the billing
	// period
	ProrateAtEndOfBillingPeriod bool `json:"prorateAtEndOfBillingPeriod,nullable"`
	// ID of the plan to start the subscription with
	SubscriptionStartPlanID string `json:"subscriptionStartPlanId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubscriptionCancellationTime respjson.Field
		SubscriptionEndSetup         respjson.Field
		SubscriptionStartSetup       respjson.Field
		DowngradePlanID              respjson.Field
		ProrateAtEndOfBillingPeriod  respjson.Field
		SubscriptionStartPlanID      respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ProductListProductsResponseProductSettings) RawJSON() string { return r.JSON.raw }
func (r *V1ProductListProductsResponseProductSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ProductListProductsParams struct {
	// Filter by entity ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by product status. Supports comma-separated values for multiple statuses
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter by creation date using range operators: gt, gte, lt, lte
	CreatedAt V1ProductListProductsParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1ProductListProductsParams]'s query parameters as
// `url.Values`.
func (r V1ProductListProductsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by creation date using range operators: gt, gte, lt, lte
type V1ProductListProductsParamsCreatedAt struct {
	// Greater than the specified createdAt value
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// Greater than or equal to the specified createdAt value
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Less than the specified createdAt value
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	// Less than or equal to the specified createdAt value
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [V1ProductListProductsParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r V1ProductListProductsParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
