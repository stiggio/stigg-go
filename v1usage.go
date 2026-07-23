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

// Operations related to usage & metering
//
// V1UsageService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1UsageService] method instead.
type V1UsageService struct {
	Options []option.RequestOption
}

// NewV1UsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1UsageService(opts ...option.RequestOption) (r V1UsageService) {
	r = V1UsageService{}
	r.Options = opts
	return
}

// Estimates the credit cost of a usage report without recording it. Returns the
// estimated cost per credit currency, the current balance, and the balance after
// the estimated consumption.
func (r *V1UsageService) Estimate(ctx context.Context, params V1UsageEstimateParams, opts ...option.RequestOption) (res *V1UsageEstimateResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/usage/estimate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves historical usage data for a customer's metered feature over time.
func (r *V1UsageService) History(ctx context.Context, featureID string, params V1UsageHistoryParams, opts ...option.RequestOption) (res *V1UsageHistoryResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.CustomerID == "" {
		err = errors.New("missing required customerId parameter")
		return nil, err
	}
	if featureID == "" {
		err = errors.New("missing required featureId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/usage/%s/history/%s", params.CustomerID, featureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Reports usage measurements for metered features. The reported usage is used to
// track, limit, and bill customer consumption.
func (r *V1UsageService) Report(ctx context.Context, params V1UsageReportParams, opts ...option.RequestOption) (res *V1UsageReportResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Response object
type V1UsageEstimateResponse struct {
	// Estimated credit cost, current balance and balance after
	Data V1UsageEstimateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageEstimateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1UsageEstimateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Estimated credit cost, current balance and balance after
type V1UsageEstimateResponseData struct {
	// Per-currency cost estimates
	Estimates []V1UsageEstimateResponseDataEstimate `json:"estimates" api:"required"`
	// Request-level warnings about the estimation context
	//
	// Any of "RESOURCE_SCOPED_SUBSCRIPTION_EXISTS", "FEATURE_NOT_FOUND",
	// "FEATURE_NOT_CREDIT_BASED".
	Warnings []string `json:"warnings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Estimates   respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageEstimateResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1UsageEstimateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1UsageEstimateResponseDataEstimate struct {
	// The credit balance after subtracting the estimated cost
	BalanceAfterEstimate float64 `json:"balanceAfterEstimate" api:"required"`
	// Estimated cost contribution per feature
	Breakdown []V1UsageEstimateResponseDataEstimateBreakdown `json:"breakdown" api:"required"`
	// The credit currency identifier
	CurrencyID string `json:"currencyId" api:"required"`
	// The current credit balance, including not-yet-reconciled consumption
	CurrentBalance float64 `json:"currentBalance" api:"required"`
	// The estimated credit cost of the reported event or usage
	EstimatedCost float64 `json:"estimatedCost" api:"required"`
	// Whether the estimated consumption would bring the balance below zero
	WouldOverdraft bool `json:"wouldOverdraft" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BalanceAfterEstimate respjson.Field
		Breakdown            respjson.Field
		CurrencyID           respjson.Field
		CurrentBalance       respjson.Field
		EstimatedCost        respjson.Field
		WouldOverdraft       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageEstimateResponseDataEstimate) RawJSON() string { return r.JSON.raw }
func (r *V1UsageEstimateResponseDataEstimate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1UsageEstimateResponseDataEstimateBreakdown struct {
	// The estimated credit cost contributed by this feature
	Cost float64 `json:"cost" api:"required"`
	// The feature whose meter contributed this cost
	FeatureID string `json:"featureId" api:"required"`
	// Warning explaining why this cost may be inaccurate, if any
	//
	// Any of "UNSUPPORTED_AGGREGATION".
	WarningCode string `json:"warningCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		FeatureID   respjson.Field
		WarningCode respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageEstimateResponseDataEstimateBreakdown) RawJSON() string { return r.JSON.raw }
func (r *V1UsageEstimateResponseDataEstimateBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1UsageHistoryResponse struct {
	// Historical usage time series
	Data V1UsageHistoryResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *V1UsageHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Historical usage time series
type V1UsageHistoryResponseData struct {
	// Markers for events that affecting feature usage
	Markers []V1UsageHistoryResponseDataMarker `json:"markers" api:"required"`
	// Series of usage history
	Series []V1UsageHistoryResponseDataSeries `json:"series" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markers     respjson.Field
		Series      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageHistoryResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1UsageHistoryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage reset or change marker
type V1UsageHistoryResponseDataMarker struct {
	// Timestamp of the marker
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Type of marker for a usage history point
	//
	// Any of "PERIODIC_RESET", "SUBSCRIPTION_CHANGE_RESET".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageHistoryResponseDataMarker) RawJSON() string { return r.JSON.raw }
func (r *V1UsageHistoryResponseDataMarker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage data points with tags
type V1UsageHistoryResponseDataSeries struct {
	// Points in the usage history series
	Points []V1UsageHistoryResponseDataSeriesPoint `json:"points" api:"required"`
	// Tags for the usage history series
	Tags []V1UsageHistoryResponseDataSeriesTag `json:"tags" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Points      respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageHistoryResponseDataSeries) RawJSON() string { return r.JSON.raw }
func (r *V1UsageHistoryResponseDataSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single usage data point
type V1UsageHistoryResponseDataSeriesPoint struct {
	// Indicates whether there was usage reset in this point, see `markers` for details
	IsResetPoint bool `json:"isResetPoint" api:"required"`
	// Timestamp of the usage history point
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Value of the usage history point
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsResetPoint respjson.Field
		Timestamp    respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageHistoryResponseDataSeriesPoint) RawJSON() string { return r.JSON.raw }
func (r *V1UsageHistoryResponseDataSeriesPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Grouping tag key-value
type V1UsageHistoryResponseDataSeriesTag struct {
	// Key of the tag
	Key string `json:"key" api:"required"`
	// Value of the tag
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
func (r V1UsageHistoryResponseDataSeriesTag) RawJSON() string { return r.JSON.raw }
func (r *V1UsageHistoryResponseDataSeriesTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing reported usage measurements with current usage values,
// period information, and reset dates for each measurement.
type V1UsageReportResponse struct {
	// Array of usage measurements with current values and period info
	Data []V1UsageReportResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageReportResponse) RawJSON() string { return r.JSON.raw }
func (r *V1UsageReportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recorded usage with period info
type V1UsageReportResponseData struct {
	// Unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Customer id
	CustomerID string `json:"customerId" api:"required"`
	// Feature id
	FeatureID string `json:"featureId" api:"required"`
	// Timestamp
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The usage measurement record
	Value int64 `json:"value" api:"required"`
	// Optimistic credit balance for a credit-backed feature
	Credit V1UsageReportResponseDataCredit `json:"credit" api:"nullable"`
	// The current measured usage value
	CurrentUsage float64 `json:"currentUsage" api:"nullable"`
	// The date when the next usage reset will occur
	NextResetDate time.Time `json:"nextResetDate" api:"nullable" format:"date-time"`
	// Resource id
	ResourceID string `json:"resourceId" api:"nullable"`
	// The end date of the usage period in which this measurement resides (for
	// entitlements with a reset period)
	UsagePeriodEnd time.Time `json:"usagePeriodEnd" api:"nullable" format:"date-time"`
	// The start date of the usage period in which this measurement resides (for
	// entitlements with a reset period)
	UsagePeriodStart time.Time `json:"usagePeriodStart" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		CustomerID       respjson.Field
		FeatureID        respjson.Field
		Timestamp        respjson.Field
		Value            respjson.Field
		Credit           respjson.Field
		CurrentUsage     respjson.Field
		NextResetDate    respjson.Field
		ResourceID       respjson.Field
		UsagePeriodEnd   respjson.Field
		UsagePeriodStart respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageReportResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1UsageReportResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optimistic credit balance for a credit-backed feature
type V1UsageReportResponseDataCredit struct {
	// The credits this single reportUsage call deducted, in credit units — scoped to
	// this one measurement (0 for idempotency duplicates). Contrast `currentUsage`,
	// which is the wallet-wide running total shared across all features on this
	// currency. Use it to reconcile expected per-call deductions.
	Consumed float64 `json:"consumed" api:"required"`
	// The credit currency identifier
	CurrencyID string `json:"currencyId" api:"required"`
	// The wallet's total consumed credits for this currency (optimistic — includes
	// not-yet-reconciled usage), shared across every feature that draws on the
	// currency. This is the running balance, not this call's deduction — see
	// `consumed` for that.
	CurrentUsage float64 `json:"currentUsage" api:"required"`
	// The grant-version timestamp of this balance, used by the SDK for last-write-wins
	// reconciliation
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The total credits granted
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// End of the current credit grant period (when recurring credits reset), if
	// applicable
	UsagePeriodEnd time.Time `json:"usagePeriodEnd" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Consumed       respjson.Field
		CurrencyID     respjson.Field
		CurrentUsage   respjson.Field
		Timestamp      respjson.Field
		UsageLimit     respjson.Field
		UsagePeriodEnd respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1UsageReportResponseDataCredit) RawJSON() string { return r.JSON.raw }
func (r *V1UsageReportResponseDataCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1UsageEstimateParams struct {
	// Customer id
	CustomerID string `json:"customerId" api:"required"`
	// Feature id
	FeatureID string `json:"featureId" api:"required"`
	// The value to report for usage
	Value int64 `json:"value" api:"required"`
	// Resource id
	ResourceID     param.Opt[string] `json:"resourceId,omitzero"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	// Additional dimensions for the usage report
	Dimensions map[string]V1UsageEstimateParamsDimensionUnion `json:"dimensions,omitzero"`
	// The method by which the usage value should be updated
	//
	// Any of "DELTA", "SET".
	UpdateBehavior V1UsageEstimateParamsUpdateBehavior `json:"updateBehavior,omitzero"`
	paramObj
}

func (r V1UsageEstimateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1UsageEstimateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1UsageEstimateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1UsageEstimateParamsDimensionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u V1UsageEstimateParamsDimensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *V1UsageEstimateParamsDimensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1UsageEstimateParamsDimensionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The method by which the usage value should be updated
type V1UsageEstimateParamsUpdateBehavior string

const (
	V1UsageEstimateParamsUpdateBehaviorDelta V1UsageEstimateParamsUpdateBehavior = "DELTA"
	V1UsageEstimateParamsUpdateBehaviorSet   V1UsageEstimateParamsUpdateBehavior = "SET"
)

type V1UsageHistoryParams struct {
	CustomerID string `path:"customerId" api:"required" json:"-"`
	// The start date of the range
	StartDate time.Time `query:"startDate" api:"required" format:"date-time" json:"-"`
	// Resource id
	ResourceID param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	// The end date of the range
	EndDate param.Opt[time.Time] `query:"endDate,omitzero" format:"date-time" json:"-"`
	// Criteria by which to group the usage history
	GroupBy        param.Opt[string] `query:"groupBy,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1UsageHistoryParams]'s query parameters as `url.Values`.
func (r V1UsageHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1UsageReportParams struct {
	// A list of usage reports to be submitted in bulk
	Usages         []V1UsageReportParamsUsage `json:"usages,omitzero" api:"required"`
	XAccountID     param.Opt[string]          `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]          `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1UsageReportParams) MarshalJSON() (data []byte, err error) {
	type shadow V1UsageReportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1UsageReportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single usage measurement
//
// The properties CustomerID, FeatureID, Value are required.
type V1UsageReportParamsUsage struct {
	// Customer id
	CustomerID string `json:"customerId" api:"required"`
	// Feature id
	FeatureID string `json:"featureId" api:"required"`
	// The value to report for usage
	Value int64 `json:"value" api:"required"`
	// Resource id
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Timestamp of when the record was created
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Idempotency key
	IdempotencyKey param.Opt[string] `json:"idempotencyKey,omitzero"`
	// Additional dimensions for the usage report
	Dimensions map[string]V1UsageReportParamsUsageDimensionUnion `json:"dimensions,omitzero"`
	// The method by which the usage value should be updated
	//
	// Any of "DELTA", "SET".
	UpdateBehavior string `json:"updateBehavior,omitzero"`
	paramObj
}

func (r V1UsageReportParamsUsage) MarshalJSON() (data []byte, err error) {
	type shadow V1UsageReportParamsUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1UsageReportParamsUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1UsageReportParamsUsage](
		"updateBehavior", "DELTA", "SET",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1UsageReportParamsUsageDimensionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u V1UsageReportParamsUsageDimensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *V1UsageReportParamsUsageDimensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1UsageReportParamsUsageDimensionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}
