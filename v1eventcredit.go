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

// V1EventCreditService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventCreditService] method instead.
type V1EventCreditService struct {
	Options []option.RequestOption
	// Operations related to credit grants
	Grants V1EventCreditGrantService
}

// NewV1EventCreditService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1EventCreditService(opts ...option.RequestOption) (r V1EventCreditService) {
	r = V1EventCreditService{}
	r.Options = opts
	r.Grants = NewV1EventCreditGrantService(opts...)
	return
}

// Retrieves the automatic recharge configuration for a customer and currency.
// Returns default settings if no configuration exists.
func (r *V1EventCreditService) GetAutoRecharge(ctx context.Context, query V1EventCreditGetAutoRechargeParams, opts ...option.RequestOption) (res *V1EventCreditGetAutoRechargeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/credits/auto-recharge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves credit usage time-series data for a customer, grouped by feature, over
// a specified time range.
func (r *V1EventCreditService) GetUsage(ctx context.Context, query V1EventCreditGetUsageParams, opts ...option.RequestOption) (res *V1EventCreditGetUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/credits/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves a paginated list of credit ledger events for a customer.
func (r *V1EventCreditService) ListLedger(ctx context.Context, query V1EventCreditListLedgerParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1EventCreditListLedgerResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/credits/ledger"
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

// Retrieves a paginated list of credit ledger events for a customer.
func (r *V1EventCreditService) ListLedgerAutoPaging(ctx context.Context, query V1EventCreditListLedgerParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1EventCreditListLedgerResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListLedger(ctx, query, opts...))
}

// Response object
type V1EventCreditGetAutoRechargeResponse struct {
	// Automatic recharge configuration for a customer and currency
	Data V1EventCreditGetAutoRechargeResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventCreditGetAutoRechargeResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventCreditGetAutoRechargeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Automatic recharge configuration for a customer and currency
type V1EventCreditGetAutoRechargeResponseData struct {
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
func (r V1EventCreditGetAutoRechargeResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventCreditGetAutoRechargeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1EventCreditGetUsageResponse struct {
	// Credit usage data grouped by feature with time-series points
	Data V1EventCreditGetUsageResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventCreditGetUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventCreditGetUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage data grouped by feature with time-series points
type V1EventCreditGetUsageResponseData struct {
	// The custom currency used for credit measurement
	Currency V1EventCreditGetUsageResponseDataCurrency `json:"currency" api:"required"`
	// Credit usage series grouped by feature
	Series []V1EventCreditGetUsageResponseDataSeries `json:"series" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Series      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventCreditGetUsageResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventCreditGetUsageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The custom currency used for credit measurement
type V1EventCreditGetUsageResponseDataCurrency struct {
	// The currency identifier
	CurrencyID string `json:"currencyId" api:"required"`
	// The display name of the currency
	DisplayName string `json:"displayName" api:"required"`
	// Plural unit label
	Plural string `json:"plural" api:"required"`
	// Singular unit label
	Singular string `json:"singular" api:"required"`
	// The currency symbol
	Symbol string `json:"symbol" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyID  respjson.Field
		DisplayName respjson.Field
		Plural      respjson.Field
		Singular    respjson.Field
		Symbol      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventCreditGetUsageResponseDataCurrency) RawJSON() string { return r.JSON.raw }
func (r *V1EventCreditGetUsageResponseDataCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage data for a single feature
type V1EventCreditGetUsageResponseDataSeries struct {
	// The feature ID
	FeatureID string `json:"featureId" api:"required"`
	// The display name of the feature
	FeatureName string `json:"featureName" api:"required"`
	// Time-series data points for this feature
	Points []V1EventCreditGetUsageResponseDataSeriesPoint `json:"points" api:"required"`
	// Total credits consumed by this feature in the time range
	TotalCredits float64 `json:"totalCredits" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FeatureID    respjson.Field
		FeatureName  respjson.Field
		Points       respjson.Field
		TotalCredits respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventCreditGetUsageResponseDataSeries) RawJSON() string { return r.JSON.raw }
func (r *V1EventCreditGetUsageResponseDataSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single data point in the credit usage time series
type V1EventCreditGetUsageResponseDataSeriesPoint struct {
	// The timestamp of the data point
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The credit usage value at this point
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventCreditGetUsageResponseDataSeriesPoint) RawJSON() string { return r.JSON.raw }
func (r *V1EventCreditGetUsageResponseDataSeriesPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A credit ledger event representing a change to credit balance
type V1EventCreditListLedgerResponse struct {
	// The credit amount for this event
	Amount float64 `json:"amount" api:"required"`
	// The credit currency ID
	CreditCurrencyID string `json:"creditCurrencyId" api:"required"`
	// The credit grant ID associated with this event
	CreditGrantID string `json:"creditGrantId" api:"required"`
	// The customer ID this event belongs to
	CustomerID string `json:"customerId" api:"required"`
	// The unique event identifier
	EventID string `json:"eventId" api:"required"`
	// The type of credit event
	//
	// Any of "CREDITS_GRANTED", "CREDITS_EXPIRED", "CREDITS_CONSUMED",
	// "CREDITS_VOIDED", "CREDITS_UPDATED", "CREDITS_CONSUMPTION_TRANSFER_SOURCE",
	// "CREDITS_CONSUMPTION_TRANSFER_TARGET".
	EventType V1EventCreditListLedgerResponseEventType `json:"eventType" api:"required"`
	// The feature ID associated with this event
	FeatureID string `json:"featureId" api:"required"`
	// The resource ID this event is scoped to
	ResourceID string `json:"resourceId" api:"required"`
	// The timestamp when the event occurred
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount           respjson.Field
		CreditCurrencyID respjson.Field
		CreditGrantID    respjson.Field
		CustomerID       respjson.Field
		EventID          respjson.Field
		EventType        respjson.Field
		FeatureID        respjson.Field
		ResourceID       respjson.Field
		Timestamp        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventCreditListLedgerResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventCreditListLedgerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of credit event
type V1EventCreditListLedgerResponseEventType string

const (
	V1EventCreditListLedgerResponseEventTypeCreditsGranted                   V1EventCreditListLedgerResponseEventType = "CREDITS_GRANTED"
	V1EventCreditListLedgerResponseEventTypeCreditsExpired                   V1EventCreditListLedgerResponseEventType = "CREDITS_EXPIRED"
	V1EventCreditListLedgerResponseEventTypeCreditsConsumed                  V1EventCreditListLedgerResponseEventType = "CREDITS_CONSUMED"
	V1EventCreditListLedgerResponseEventTypeCreditsVoided                    V1EventCreditListLedgerResponseEventType = "CREDITS_VOIDED"
	V1EventCreditListLedgerResponseEventTypeCreditsUpdated                   V1EventCreditListLedgerResponseEventType = "CREDITS_UPDATED"
	V1EventCreditListLedgerResponseEventTypeCreditsConsumptionTransferSource V1EventCreditListLedgerResponseEventType = "CREDITS_CONSUMPTION_TRANSFER_SOURCE"
	V1EventCreditListLedgerResponseEventTypeCreditsConsumptionTransferTarget V1EventCreditListLedgerResponseEventType = "CREDITS_CONSUMPTION_TRANSFER_TARGET"
)

type V1EventCreditGetAutoRechargeParams struct {
	// Filter by currency ID (required)
	CurrencyID string `query:"currencyId" api:"required" json:"-"`
	// Filter by customer ID (required)
	CustomerID string `query:"customerId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventCreditGetAutoRechargeParams]'s query parameters as
// `url.Values`.
func (r V1EventCreditGetAutoRechargeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1EventCreditGetUsageParams struct {
	// Filter by customer ID (required)
	CustomerID string `query:"customerId" api:"required" json:"-"`
	// Filter by currency ID
	CurrencyID param.Opt[string] `query:"currencyId,omitzero" json:"-"`
	// Filter by resource ID
	ResourceID param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	// Time range for usage data (LAST_DAY, LAST_WEEK, LAST_MONTH, LAST_YEAR). Defaults
	// to LAST_MONTH
	//
	// Any of "LAST_DAY", "LAST_WEEK", "LAST_MONTH", "LAST_YEAR".
	TimeRange V1EventCreditGetUsageParamsTimeRange `query:"timeRange,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventCreditGetUsageParams]'s query parameters as
// `url.Values`.
func (r V1EventCreditGetUsageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time range for usage data (LAST_DAY, LAST_WEEK, LAST_MONTH, LAST_YEAR). Defaults
// to LAST_MONTH
type V1EventCreditGetUsageParamsTimeRange string

const (
	V1EventCreditGetUsageParamsTimeRangeLastDay   V1EventCreditGetUsageParamsTimeRange = "LAST_DAY"
	V1EventCreditGetUsageParamsTimeRangeLastWeek  V1EventCreditGetUsageParamsTimeRange = "LAST_WEEK"
	V1EventCreditGetUsageParamsTimeRangeLastMonth V1EventCreditGetUsageParamsTimeRange = "LAST_MONTH"
	V1EventCreditGetUsageParamsTimeRangeLastYear  V1EventCreditGetUsageParamsTimeRange = "LAST_YEAR"
)

type V1EventCreditListLedgerParams struct {
	// Filter by customer ID (required)
	CustomerID string `query:"customerId" api:"required" json:"-"`
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter by currency ID
	CurrencyID param.Opt[string] `query:"currencyId,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by resource ID
	ResourceID param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventCreditListLedgerParams]'s query parameters as
// `url.Values`.
func (r V1EventCreditListLedgerParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
