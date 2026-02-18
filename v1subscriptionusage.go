// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stiggio/stigg-go/internal/apijson"
	"github.com/stiggio/stigg-go/internal/requestconfig"
	"github.com/stiggio/stigg-go/option"
	"github.com/stiggio/stigg-go/packages/param"
	"github.com/stiggio/stigg-go/packages/respjson"
)

// V1SubscriptionUsageService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1SubscriptionUsageService] method instead.
type V1SubscriptionUsageService struct {
	Options []option.RequestOption
}

// NewV1SubscriptionUsageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1SubscriptionUsageService(opts ...option.RequestOption) (r V1SubscriptionUsageService) {
	r = V1SubscriptionUsageService{}
	r.Options = opts
	return
}

// Immediately charges usage for a subscription via the billing integration.
// Calculates usage since the last charge and creates an invoice.
func (r *V1SubscriptionUsageService) ChargeUsage(ctx context.Context, id string, body V1SubscriptionUsageChargeUsageParams, opts ...option.RequestOption) (res *V1SubscriptionUsageChargeUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/usage/charge", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response object
type V1SubscriptionUsageChargeUsageResponse struct {
	// Result of charging subscription usage including the billing period and charged
	// items.
	Data V1SubscriptionUsageChargeUsageResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionUsageChargeUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionUsageChargeUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of charging subscription usage including the billing period and charged
// items.
type V1SubscriptionUsageChargeUsageResponseData struct {
	// The invoice ID in the billing integration
	InvoiceBillingID string `json:"invoiceBillingId,required"`
	// End of the usage billing period
	PeriodEnd time.Time `json:"periodEnd,required" format:"date-time"`
	// Start of the usage billing period
	PeriodStart time.Time `json:"periodStart,required" format:"date-time"`
	// The subscription ID for which usage was charged
	SubscriptionID string `json:"subscriptionId,required"`
	// Usage items that were charged
	UsageCharged []V1SubscriptionUsageChargeUsageResponseDataUsageCharged `json:"usageCharged,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InvoiceBillingID respjson.Field
		PeriodEnd        respjson.Field
		PeriodStart      respjson.Field
		SubscriptionID   respjson.Field
		UsageCharged     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionUsageChargeUsageResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionUsageChargeUsageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single usage item that was charged.
type V1SubscriptionUsageChargeUsageResponseDataUsageCharged struct {
	// The feature ID for which usage was charged
	FeatureID string `json:"featureId,required"`
	// The number of units charged
	UsageAmount float64 `json:"usageAmount,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FeatureID   respjson.Field
		UsageAmount respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionUsageChargeUsageResponseDataUsageCharged) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionUsageChargeUsageResponseDataUsageCharged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionUsageChargeUsageParams struct {
	// Cutoff date for usage calculation. If not provided, the current time is used.
	UntilDate param.Opt[time.Time] `json:"untilDate,omitzero" format:"date-time"`
	paramObj
}

func (r V1SubscriptionUsageChargeUsageParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUsageChargeUsageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUsageChargeUsageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
