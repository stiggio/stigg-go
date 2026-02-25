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

// Archives a product, preventing new subscriptions. All plans and addons are
// archived.
func (r *V1ProductService) ArchiveProduct(ctx context.Context, id string, opts ...option.RequestOption) (res *V1ProductArchiveProductResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/products/%s/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Creates a new product.
func (r *V1ProductService) NewProduct(ctx context.Context, body V1ProductNewProductParams, opts ...option.RequestOption) (res *V1ProductNewProductResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Duplicates an existing product, including its plans, addons, and configuration.
func (r *V1ProductService) DuplicateProduct(ctx context.Context, id string, body V1ProductDuplicateProductParams, opts ...option.RequestOption) (res *V1ProductDuplicateProductResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/products/%s/duplicate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

// Restores an archived product, allowing new subscriptions to be created.
func (r *V1ProductService) UnarchiveProduct(ctx context.Context, id string, opts ...option.RequestOption) (res *V1ProductUnarchiveProductResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/products/%s/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Updates an existing product's properties such as display name, description, and
// metadata.
func (r *V1ProductService) UpdateProduct(ctx context.Context, id string, body V1ProductUpdateProductParams, opts ...option.RequestOption) (res *V1ProductUpdateProductResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Response object
type V1ProductArchiveProductResponse struct {
	// Product configuration object
	Data V1ProductArchiveProductResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ProductArchiveProductResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ProductArchiveProductResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product configuration object
type V1ProductArchiveProductResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description of the product
	Description string `json:"description" api:"required"`
	// Display name of the product
	DisplayName string `json:"displayName" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// Indicates if multiple subscriptions to this product are allowed
	MultipleSubscriptions bool `json:"multipleSubscriptions" api:"required"`
	// The status of the product
	//
	// Any of "PUBLISHED", "ARCHIVED".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Product behavior settings for subscription lifecycle management.
	ProductSettings V1ProductArchiveProductResponseDataProductSettings `json:"productSettings"`
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
func (r V1ProductArchiveProductResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ProductArchiveProductResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product behavior settings for subscription lifecycle management.
type V1ProductArchiveProductResponseDataProductSettings struct {
	// Time when the subscription will be cancelled
	//
	// Any of "END_OF_BILLING_PERIOD", "IMMEDIATE", "SPECIFIC_DATE".
	SubscriptionCancellationTime string `json:"subscriptionCancellationTime" api:"required"`
	// Setup for the end of the subscription
	//
	// Any of "DOWNGRADE_TO_FREE", "CANCEL_SUBSCRIPTION".
	SubscriptionEndSetup string `json:"subscriptionEndSetup" api:"required"`
	// Setup for the start of the subscription
	//
	// Any of "PLAN_SELECTION", "TRIAL_PERIOD", "FREE_PLAN".
	SubscriptionStartSetup string `json:"subscriptionStartSetup" api:"required"`
	// ID of the plan to downgrade to at the end of the billing period
	DowngradePlanID string `json:"downgradePlanId" api:"nullable"`
	// Indicates if the subscription should be prorated at the end of the billing
	// period
	ProrateAtEndOfBillingPeriod bool `json:"prorateAtEndOfBillingPeriod" api:"nullable"`
	// ID of the plan to start the subscription with
	SubscriptionStartPlanID string `json:"subscriptionStartPlanId" api:"nullable"`
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
func (r V1ProductArchiveProductResponseDataProductSettings) RawJSON() string { return r.JSON.raw }
func (r *V1ProductArchiveProductResponseDataProductSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1ProductNewProductResponse struct {
	// Product configuration object
	Data V1ProductNewProductResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ProductNewProductResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ProductNewProductResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product configuration object
type V1ProductNewProductResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description of the product
	Description string `json:"description" api:"required"`
	// Display name of the product
	DisplayName string `json:"displayName" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// Indicates if multiple subscriptions to this product are allowed
	MultipleSubscriptions bool `json:"multipleSubscriptions" api:"required"`
	// The status of the product
	//
	// Any of "PUBLISHED", "ARCHIVED".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Product behavior settings for subscription lifecycle management.
	ProductSettings V1ProductNewProductResponseDataProductSettings `json:"productSettings"`
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
func (r V1ProductNewProductResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ProductNewProductResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product behavior settings for subscription lifecycle management.
type V1ProductNewProductResponseDataProductSettings struct {
	// Time when the subscription will be cancelled
	//
	// Any of "END_OF_BILLING_PERIOD", "IMMEDIATE", "SPECIFIC_DATE".
	SubscriptionCancellationTime string `json:"subscriptionCancellationTime" api:"required"`
	// Setup for the end of the subscription
	//
	// Any of "DOWNGRADE_TO_FREE", "CANCEL_SUBSCRIPTION".
	SubscriptionEndSetup string `json:"subscriptionEndSetup" api:"required"`
	// Setup for the start of the subscription
	//
	// Any of "PLAN_SELECTION", "TRIAL_PERIOD", "FREE_PLAN".
	SubscriptionStartSetup string `json:"subscriptionStartSetup" api:"required"`
	// ID of the plan to downgrade to at the end of the billing period
	DowngradePlanID string `json:"downgradePlanId" api:"nullable"`
	// Indicates if the subscription should be prorated at the end of the billing
	// period
	ProrateAtEndOfBillingPeriod bool `json:"prorateAtEndOfBillingPeriod" api:"nullable"`
	// ID of the plan to start the subscription with
	SubscriptionStartPlanID string `json:"subscriptionStartPlanId" api:"nullable"`
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
func (r V1ProductNewProductResponseDataProductSettings) RawJSON() string { return r.JSON.raw }
func (r *V1ProductNewProductResponseDataProductSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1ProductDuplicateProductResponse struct {
	// Product configuration object
	Data V1ProductDuplicateProductResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ProductDuplicateProductResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ProductDuplicateProductResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product configuration object
type V1ProductDuplicateProductResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description of the product
	Description string `json:"description" api:"required"`
	// Display name of the product
	DisplayName string `json:"displayName" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// Indicates if multiple subscriptions to this product are allowed
	MultipleSubscriptions bool `json:"multipleSubscriptions" api:"required"`
	// The status of the product
	//
	// Any of "PUBLISHED", "ARCHIVED".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Product behavior settings for subscription lifecycle management.
	ProductSettings V1ProductDuplicateProductResponseDataProductSettings `json:"productSettings"`
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
func (r V1ProductDuplicateProductResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ProductDuplicateProductResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product behavior settings for subscription lifecycle management.
type V1ProductDuplicateProductResponseDataProductSettings struct {
	// Time when the subscription will be cancelled
	//
	// Any of "END_OF_BILLING_PERIOD", "IMMEDIATE", "SPECIFIC_DATE".
	SubscriptionCancellationTime string `json:"subscriptionCancellationTime" api:"required"`
	// Setup for the end of the subscription
	//
	// Any of "DOWNGRADE_TO_FREE", "CANCEL_SUBSCRIPTION".
	SubscriptionEndSetup string `json:"subscriptionEndSetup" api:"required"`
	// Setup for the start of the subscription
	//
	// Any of "PLAN_SELECTION", "TRIAL_PERIOD", "FREE_PLAN".
	SubscriptionStartSetup string `json:"subscriptionStartSetup" api:"required"`
	// ID of the plan to downgrade to at the end of the billing period
	DowngradePlanID string `json:"downgradePlanId" api:"nullable"`
	// Indicates if the subscription should be prorated at the end of the billing
	// period
	ProrateAtEndOfBillingPeriod bool `json:"prorateAtEndOfBillingPeriod" api:"nullable"`
	// ID of the plan to start the subscription with
	SubscriptionStartPlanID string `json:"subscriptionStartPlanId" api:"nullable"`
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
func (r V1ProductDuplicateProductResponseDataProductSettings) RawJSON() string { return r.JSON.raw }
func (r *V1ProductDuplicateProductResponseDataProductSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product configuration object
type V1ProductListProductsResponse struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description of the product
	Description string `json:"description" api:"required"`
	// Display name of the product
	DisplayName string `json:"displayName" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// Indicates if multiple subscriptions to this product are allowed
	MultipleSubscriptions bool `json:"multipleSubscriptions" api:"required"`
	// The status of the product
	//
	// Any of "PUBLISHED", "ARCHIVED".
	Status V1ProductListProductsResponseStatus `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
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
	SubscriptionCancellationTime string `json:"subscriptionCancellationTime" api:"required"`
	// Setup for the end of the subscription
	//
	// Any of "DOWNGRADE_TO_FREE", "CANCEL_SUBSCRIPTION".
	SubscriptionEndSetup string `json:"subscriptionEndSetup" api:"required"`
	// Setup for the start of the subscription
	//
	// Any of "PLAN_SELECTION", "TRIAL_PERIOD", "FREE_PLAN".
	SubscriptionStartSetup string `json:"subscriptionStartSetup" api:"required"`
	// ID of the plan to downgrade to at the end of the billing period
	DowngradePlanID string `json:"downgradePlanId" api:"nullable"`
	// Indicates if the subscription should be prorated at the end of the billing
	// period
	ProrateAtEndOfBillingPeriod bool `json:"prorateAtEndOfBillingPeriod" api:"nullable"`
	// ID of the plan to start the subscription with
	SubscriptionStartPlanID string `json:"subscriptionStartPlanId" api:"nullable"`
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

// Response object
type V1ProductUnarchiveProductResponse struct {
	// Product configuration object
	Data V1ProductUnarchiveProductResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ProductUnarchiveProductResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ProductUnarchiveProductResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product configuration object
type V1ProductUnarchiveProductResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description of the product
	Description string `json:"description" api:"required"`
	// Display name of the product
	DisplayName string `json:"displayName" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// Indicates if multiple subscriptions to this product are allowed
	MultipleSubscriptions bool `json:"multipleSubscriptions" api:"required"`
	// The status of the product
	//
	// Any of "PUBLISHED", "ARCHIVED".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Product behavior settings for subscription lifecycle management.
	ProductSettings V1ProductUnarchiveProductResponseDataProductSettings `json:"productSettings"`
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
func (r V1ProductUnarchiveProductResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ProductUnarchiveProductResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product behavior settings for subscription lifecycle management.
type V1ProductUnarchiveProductResponseDataProductSettings struct {
	// Time when the subscription will be cancelled
	//
	// Any of "END_OF_BILLING_PERIOD", "IMMEDIATE", "SPECIFIC_DATE".
	SubscriptionCancellationTime string `json:"subscriptionCancellationTime" api:"required"`
	// Setup for the end of the subscription
	//
	// Any of "DOWNGRADE_TO_FREE", "CANCEL_SUBSCRIPTION".
	SubscriptionEndSetup string `json:"subscriptionEndSetup" api:"required"`
	// Setup for the start of the subscription
	//
	// Any of "PLAN_SELECTION", "TRIAL_PERIOD", "FREE_PLAN".
	SubscriptionStartSetup string `json:"subscriptionStartSetup" api:"required"`
	// ID of the plan to downgrade to at the end of the billing period
	DowngradePlanID string `json:"downgradePlanId" api:"nullable"`
	// Indicates if the subscription should be prorated at the end of the billing
	// period
	ProrateAtEndOfBillingPeriod bool `json:"prorateAtEndOfBillingPeriod" api:"nullable"`
	// ID of the plan to start the subscription with
	SubscriptionStartPlanID string `json:"subscriptionStartPlanId" api:"nullable"`
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
func (r V1ProductUnarchiveProductResponseDataProductSettings) RawJSON() string { return r.JSON.raw }
func (r *V1ProductUnarchiveProductResponseDataProductSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1ProductUpdateProductResponse struct {
	// Product configuration object
	Data V1ProductUpdateProductResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ProductUpdateProductResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ProductUpdateProductResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product configuration object
type V1ProductUpdateProductResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description of the product
	Description string `json:"description" api:"required"`
	// Display name of the product
	DisplayName string `json:"displayName" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// Indicates if multiple subscriptions to this product are allowed
	MultipleSubscriptions bool `json:"multipleSubscriptions" api:"required"`
	// The status of the product
	//
	// Any of "PUBLISHED", "ARCHIVED".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Product behavior settings for subscription lifecycle management.
	ProductSettings V1ProductUpdateProductResponseDataProductSettings `json:"productSettings"`
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
func (r V1ProductUpdateProductResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ProductUpdateProductResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product behavior settings for subscription lifecycle management.
type V1ProductUpdateProductResponseDataProductSettings struct {
	// Time when the subscription will be cancelled
	//
	// Any of "END_OF_BILLING_PERIOD", "IMMEDIATE", "SPECIFIC_DATE".
	SubscriptionCancellationTime string `json:"subscriptionCancellationTime" api:"required"`
	// Setup for the end of the subscription
	//
	// Any of "DOWNGRADE_TO_FREE", "CANCEL_SUBSCRIPTION".
	SubscriptionEndSetup string `json:"subscriptionEndSetup" api:"required"`
	// Setup for the start of the subscription
	//
	// Any of "PLAN_SELECTION", "TRIAL_PERIOD", "FREE_PLAN".
	SubscriptionStartSetup string `json:"subscriptionStartSetup" api:"required"`
	// ID of the plan to downgrade to at the end of the billing period
	DowngradePlanID string `json:"downgradePlanId" api:"nullable"`
	// Indicates if the subscription should be prorated at the end of the billing
	// period
	ProrateAtEndOfBillingPeriod bool `json:"prorateAtEndOfBillingPeriod" api:"nullable"`
	// ID of the plan to start the subscription with
	SubscriptionStartPlanID string `json:"subscriptionStartPlanId" api:"nullable"`
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
func (r V1ProductUpdateProductResponseDataProductSettings) RawJSON() string { return r.JSON.raw }
func (r *V1ProductUpdateProductResponseDataProductSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ProductNewProductParams struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Description of the product
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of the product
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Indicates if multiple subscriptions to this product are allowed
	MultipleSubscriptions param.Opt[bool] `json:"multipleSubscriptions,omitzero"`
	// Additional metadata for the product
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1ProductNewProductParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ProductNewProductParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ProductNewProductParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ProductDuplicateProductParams struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Description of the product
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of the product
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	paramObj
}

func (r V1ProductDuplicateProductParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ProductDuplicateProductParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ProductDuplicateProductParams) UnmarshalJSON(data []byte) error {
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

type V1ProductUpdateProductParams struct {
	// Description of the product
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of the product
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Indicates if multiple subscriptions to this product are allowed
	MultipleSubscriptions param.Opt[bool] `json:"multipleSubscriptions,omitzero"`
	// Additional metadata for the product
	Metadata        map[string]string                           `json:"metadata,omitzero"`
	ProductSettings V1ProductUpdateProductParamsProductSettings `json:"productSettings,omitzero"`
	// Rule defining when usage resets upon subscription update.
	UsageResetCutoffRule V1ProductUpdateProductParamsUsageResetCutoffRule `json:"usageResetCutoffRule,omitzero"`
	paramObj
}

func (r V1ProductUpdateProductParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ProductUpdateProductParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ProductUpdateProductParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SubscriptionCancellationTime, SubscriptionEndSetup,
// SubscriptionStartSetup are required.
type V1ProductUpdateProductParamsProductSettings struct {
	// Time when the subscription will be cancelled
	//
	// Any of "END_OF_BILLING_PERIOD", "IMMEDIATE", "SPECIFIC_DATE".
	SubscriptionCancellationTime string `json:"subscriptionCancellationTime,omitzero" api:"required"`
	// Setup for the end of the subscription
	//
	// Any of "DOWNGRADE_TO_FREE", "CANCEL_SUBSCRIPTION".
	SubscriptionEndSetup string `json:"subscriptionEndSetup,omitzero" api:"required"`
	// Setup for the start of the subscription
	//
	// Any of "PLAN_SELECTION", "TRIAL_PERIOD", "FREE_PLAN".
	SubscriptionStartSetup string `json:"subscriptionStartSetup,omitzero" api:"required"`
	// ID of the plan to downgrade to at the end of the billing period
	DowngradePlanID             param.Opt[string] `json:"downgradePlanId,omitzero"`
	ProrateAtEndOfBillingPeriod param.Opt[bool]   `json:"prorateAtEndOfBillingPeriod,omitzero"`
	// ID of the plan to start the subscription with
	SubscriptionStartPlanID param.Opt[string] `json:"subscriptionStartPlanId,omitzero"`
	paramObj
}

func (r V1ProductUpdateProductParamsProductSettings) MarshalJSON() (data []byte, err error) {
	type shadow V1ProductUpdateProductParamsProductSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ProductUpdateProductParamsProductSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ProductUpdateProductParamsProductSettings](
		"subscriptionCancellationTime", "END_OF_BILLING_PERIOD", "IMMEDIATE", "SPECIFIC_DATE",
	)
	apijson.RegisterFieldValidator[V1ProductUpdateProductParamsProductSettings](
		"subscriptionEndSetup", "DOWNGRADE_TO_FREE", "CANCEL_SUBSCRIPTION",
	)
	apijson.RegisterFieldValidator[V1ProductUpdateProductParamsProductSettings](
		"subscriptionStartSetup", "PLAN_SELECTION", "TRIAL_PERIOD", "FREE_PLAN",
	)
}

// Rule defining when usage resets upon subscription update.
//
// The property Behavior is required.
type V1ProductUpdateProductParamsUsageResetCutoffRule struct {
	// Behavior of the usage reset cutoff rule
	//
	// Any of "NEVER_RESET", "ALWAYS_RESET", "BILLING_PERIOD_CHANGE".
	Behavior string `json:"behavior,omitzero" api:"required"`
	paramObj
}

func (r V1ProductUpdateProductParamsUsageResetCutoffRule) MarshalJSON() (data []byte, err error) {
	type shadow V1ProductUpdateProductParamsUsageResetCutoffRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ProductUpdateProductParamsUsageResetCutoffRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ProductUpdateProductParamsUsageResetCutoffRule](
		"behavior", "NEVER_RESET", "ALWAYS_RESET", "BILLING_PERIOD_CHANGE",
	)
}
