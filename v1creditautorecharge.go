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
	"github.com/stiggio/stigg-go/packages/respjson"
)

// V1CreditAutoRechargeService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CreditAutoRechargeService] method instead.
type V1CreditAutoRechargeService struct {
	Options []option.RequestOption
}

// NewV1CreditAutoRechargeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CreditAutoRechargeService(opts ...option.RequestOption) (r V1CreditAutoRechargeService) {
	r = V1CreditAutoRechargeService{}
	r.Options = opts
	return
}

// Retrieves the automatic recharge configuration for a customer and currency.
// Returns default settings if no configuration exists.
func (r *V1CreditAutoRechargeService) GetAutoRecharge(ctx context.Context, query V1CreditAutoRechargeGetAutoRechargeParams, opts ...option.RequestOption) (res *V1CreditAutoRechargeGetAutoRechargeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/credits/auto-recharge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Response object
type V1CreditAutoRechargeGetAutoRechargeResponse struct {
	// Automatic recharge configuration for a customer and currency
	Data V1CreditAutoRechargeGetAutoRechargeResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditAutoRechargeGetAutoRechargeResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditAutoRechargeGetAutoRechargeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Automatic recharge configuration for a customer and currency
type V1CreditAutoRechargeGetAutoRechargeResponseData struct {
	// The unique configuration ID
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The currency ID for this configuration
	CurrencyID string `json:"currencyId" api:"required"`
	// The customer ID this configuration belongs to
	CustomerID string `json:"customerId" api:"required"`
	// Expiration period for auto-recharge grants (1_MONTH or 1_YEAR)
	//
	// Any of "1_MONTH", "1_YEAR".
	GrantExpirationPeriod string `json:"grantExpirationPeriod" api:"required"`
	// Whether automatic recharge is enabled
	IsEnabled bool `json:"isEnabled" api:"required"`
	// Maximum monthly spend limit for automatic recharges
	MaxSpendLimit float64 `json:"maxSpendLimit" api:"required"`
	// The target credit balance to recharge to
	TargetBalance float64 `json:"targetBalance" api:"required"`
	// The threshold type (CREDIT_AMOUNT or DOLLAR_AMOUNT)
	//
	// Any of "CREDIT_AMOUNT", "DOLLAR_AMOUNT".
	ThresholdType string `json:"thresholdType" api:"required"`
	// The threshold value that triggers a recharge
	ThresholdValue float64 `json:"thresholdValue" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		CurrencyID            respjson.Field
		CustomerID            respjson.Field
		GrantExpirationPeriod respjson.Field
		IsEnabled             respjson.Field
		MaxSpendLimit         respjson.Field
		TargetBalance         respjson.Field
		ThresholdType         respjson.Field
		ThresholdValue        respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditAutoRechargeGetAutoRechargeResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CreditAutoRechargeGetAutoRechargeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditAutoRechargeGetAutoRechargeParams struct {
	// Filter by currency ID (required)
	CurrencyID string `query:"currencyId" api:"required" json:"-"`
	// Filter by customer ID (required)
	CustomerID string `query:"customerId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1CreditAutoRechargeGetAutoRechargeParams]'s query
// parameters as `url.Values`.
func (r V1CreditAutoRechargeGetAutoRechargeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
