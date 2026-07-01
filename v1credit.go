// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
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

// V1CreditService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CreditService] method instead.
type V1CreditService struct {
	Options []option.RequestOption
	// Operations related to credit grants
	Grants V1CreditGrantService
	// Operations related to custom currencies
	CustomCurrencies V1CreditCustomCurrencyService
	Consumption      V1CreditConsumptionService
}

// NewV1CreditService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CreditService(opts ...option.RequestOption) (r V1CreditService) {
	r = V1CreditService{}
	r.Options = opts
	r.Grants = NewV1CreditGrantService(opts...)
	r.CustomCurrencies = NewV1CreditCustomCurrencyService(opts...)
	r.Consumption = NewV1CreditConsumptionService(opts...)
	return
}

// Retrieves the automatic recharge configuration for a customer and currency.
// Returns default settings if no configuration exists.
func (r *V1CreditService) GetAutoRecharge(ctx context.Context, params V1CreditGetAutoRechargeParams, opts ...option.RequestOption) (res *V1CreditGetAutoRechargeResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/credits/auto-recharge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Retrieves credit usage time-series data for a customer, grouped by feature, over
// a specified time range.
func (r *V1CreditService) GetUsage(ctx context.Context, params V1CreditGetUsageParams, opts ...option.RequestOption) (res *V1CreditGetUsageResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/credits/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Retrieves a paginated list of credit ledger events for a customer.
func (r *V1CreditService) ListLedger(ctx context.Context, params V1CreditListLedgerParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1CreditListLedgerResponse], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/credits/ledger"
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

// Retrieves a paginated list of credit ledger events for a customer.
func (r *V1CreditService) ListLedgerAutoPaging(ctx context.Context, params V1CreditListLedgerParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1CreditListLedgerResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListLedger(ctx, params, opts...))
}

// Response object
type V1CreditGetAutoRechargeResponse struct {
	// Automatic recharge configuration for a customer and currency
	Data V1CreditGetAutoRechargeResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGetAutoRechargeResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGetAutoRechargeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Automatic recharge configuration for a customer and currency
type V1CreditGetAutoRechargeResponseData struct {
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
func (r V1CreditGetAutoRechargeResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGetAutoRechargeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1CreditGetUsageResponse struct {
	// Credit usage data grouped by feature with time-series points
	Data V1CreditGetUsageResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGetUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGetUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage data grouped by feature with time-series points
type V1CreditGetUsageResponseData struct {
	// The custom currency used for credit measurement
	Currency V1CreditGetUsageResponseDataCurrency `json:"currency" api:"required"`
	// Cursor-based pagination for the returned series. `next`/`prev` are opaque
	// cursors; pass them back as `after`/`before` to traverse pages. The series axis
	// is `groupBy` when provided, otherwise `featureId`
	Pagination V1CreditGetUsageResponseDataPagination `json:"pagination" api:"required"`
	// Credit usage series grouped by feature
	Series []V1CreditGetUsageResponseDataSeries `json:"series" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Pagination  respjson.Field
		Series      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGetUsageResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGetUsageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The custom currency used for credit measurement
type V1CreditGetUsageResponseDataCurrency struct {
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
func (r V1CreditGetUsageResponseDataCurrency) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGetUsageResponseDataCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination for the returned series. `next`/`prev` are opaque
// cursors; pass them back as `after`/`before` to traverse pages. The series axis
// is `groupBy` when provided, otherwise `featureId`
type V1CreditGetUsageResponseDataPagination struct {
	// Cursor for fetching the next page of results, or null if no additional pages
	// exist
	Next string `json:"next" api:"required"`
	// Cursor for fetching the previous page of results, or null if at the beginning
	Prev string `json:"prev" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		Prev        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGetUsageResponseDataPagination) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGetUsageResponseDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit usage data for a single feature
type V1CreditGetUsageResponseDataSeries struct {
	// The feature ID; null when grouping by dimensions only
	FeatureID string `json:"featureId" api:"required"`
	// The display name of the feature; null when grouping by dimensions only
	FeatureName string `json:"featureName" api:"required"`
	// Time-series data points for this feature
	Points []V1CreditGetUsageResponseDataSeriesPoint `json:"points" api:"required"`
	// Total credits consumed by this feature in the time range
	TotalCredits float64 `json:"totalCredits" api:"required"`
	// Dimension key/value pairs identifying this series when groupBy is applied
	Tags []V1CreditGetUsageResponseDataSeriesTag `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FeatureID    respjson.Field
		FeatureName  respjson.Field
		Points       respjson.Field
		TotalCredits respjson.Field
		Tags         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGetUsageResponseDataSeries) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGetUsageResponseDataSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single data point in the credit usage time series
type V1CreditGetUsageResponseDataSeriesPoint struct {
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
func (r V1CreditGetUsageResponseDataSeriesPoint) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGetUsageResponseDataSeriesPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dimension key/value pair identifying a credit usage series
type V1CreditGetUsageResponseDataSeriesTag struct {
	// The dimension key
	Key string `json:"key" api:"required"`
	// The dimension value for this series
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGetUsageResponseDataSeriesTag) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGetUsageResponseDataSeriesTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A credit ledger event representing a change to credit balance
type V1CreditListLedgerResponse struct {
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
	EventType V1CreditListLedgerResponseEventType `json:"eventType" api:"required"`
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
func (r V1CreditListLedgerResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditListLedgerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of credit event
type V1CreditListLedgerResponseEventType string

const (
	V1CreditListLedgerResponseEventTypeCreditsGranted                   V1CreditListLedgerResponseEventType = "CREDITS_GRANTED"
	V1CreditListLedgerResponseEventTypeCreditsExpired                   V1CreditListLedgerResponseEventType = "CREDITS_EXPIRED"
	V1CreditListLedgerResponseEventTypeCreditsConsumed                  V1CreditListLedgerResponseEventType = "CREDITS_CONSUMED"
	V1CreditListLedgerResponseEventTypeCreditsVoided                    V1CreditListLedgerResponseEventType = "CREDITS_VOIDED"
	V1CreditListLedgerResponseEventTypeCreditsUpdated                   V1CreditListLedgerResponseEventType = "CREDITS_UPDATED"
	V1CreditListLedgerResponseEventTypeCreditsConsumptionTransferSource V1CreditListLedgerResponseEventType = "CREDITS_CONSUMPTION_TRANSFER_SOURCE"
	V1CreditListLedgerResponseEventTypeCreditsConsumptionTransferTarget V1CreditListLedgerResponseEventType = "CREDITS_CONSUMPTION_TRANSFER_TARGET"
)

type V1CreditGetAutoRechargeParams struct {
	// Filter by currency ID (required)
	CurrencyID string `query:"currencyId" api:"required" json:"-"`
	// Filter by customer ID (required)
	CustomerID     string            `query:"customerId" api:"required" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CreditGetAutoRechargeParams]'s query parameters as
// `url.Values`.
func (r V1CreditGetAutoRechargeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CreditGetUsageParams struct {
	// Filter by customer ID (required)
	CustomerID string `query:"customerId" api:"required" json:"-"`
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter by currency ID
	CurrencyID param.Opt[string] `query:"currencyId,omitzero" json:"-"`
	// End date for the credit usage time range (ISO 8601). Defaults to now when
	// startDate is provided
	EndDate param.Opt[time.Time] `query:"endDate,omitzero" format:"date-time" json:"-"`
	// Comma-separated list of feature dimension keys to group usage series by (up to
	// 3). Each key matches /^[a-zA-Z0-9_$-]+$/
	GroupBy param.Opt[string] `query:"groupBy,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by resource ID
	ResourceID param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	// Start date for the credit usage time range (ISO 8601). Takes precedence over
	// timeRange when provided
	StartDate      param.Opt[time.Time] `query:"startDate,omitzero" format:"date-time" json:"-"`
	XAccountID     param.Opt[string]    `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]    `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	// Time range for usage data (LAST_DAY, LAST_WEEK, LAST_MONTH, LAST_YEAR). Defaults
	// to LAST_MONTH
	//
	// Any of "LAST_DAY", "LAST_WEEK", "LAST_MONTH", "LAST_YEAR".
	TimeRange V1CreditGetUsageParamsTimeRange `query:"timeRange,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CreditGetUsageParams]'s query parameters as `url.Values`.
func (r V1CreditGetUsageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time range for usage data (LAST_DAY, LAST_WEEK, LAST_MONTH, LAST_YEAR). Defaults
// to LAST_MONTH
type V1CreditGetUsageParamsTimeRange string

const (
	V1CreditGetUsageParamsTimeRangeLastDay   V1CreditGetUsageParamsTimeRange = "LAST_DAY"
	V1CreditGetUsageParamsTimeRangeLastWeek  V1CreditGetUsageParamsTimeRange = "LAST_WEEK"
	V1CreditGetUsageParamsTimeRangeLastMonth V1CreditGetUsageParamsTimeRange = "LAST_MONTH"
	V1CreditGetUsageParamsTimeRangeLastYear  V1CreditGetUsageParamsTimeRange = "LAST_YEAR"
)

type V1CreditListLedgerParams struct {
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
	ResourceID     param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CreditListLedgerParams]'s query parameters as
// `url.Values`.
func (r V1CreditListLedgerParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
