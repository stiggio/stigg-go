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

// Get usage history
func (r *V1UsageService) History(ctx context.Context, featureID string, params V1UsageHistoryParams, opts ...option.RequestOption) (res *V1UsageHistoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CustomerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	if featureID == "" {
		err = errors.New("missing required featureId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/usage/%s/history/%s", params.CustomerID, featureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Report usage measurements
func (r *V1UsageService) Report(ctx context.Context, body V1UsageReportParams, opts ...option.RequestOption) (res *V1UsageReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response object
type V1UsageHistoryResponse struct {
	// Historical usage time series
	Data V1UsageHistoryResponseData `json:"data,required"`
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
	Markers []V1UsageHistoryResponseDataMarker `json:"markers,required"`
	// Series of usage history
	Series []V1UsageHistoryResponseDataSeries `json:"series,required"`
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
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Type of marker for a usage history point
	//
	// Any of "PERIODIC_RESET", "SUBSCRIPTION_CHANGE_RESET".
	Type string `json:"type,required"`
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
	Points []V1UsageHistoryResponseDataSeriesPoint `json:"points,required"`
	// Tags for the usage history series
	Tags []V1UsageHistoryResponseDataSeriesTag `json:"tags,required"`
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
	IsResetPoint bool `json:"isResetPoint,required"`
	// Timestamp of the usage history point
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Value of the usage history point
	Value float64 `json:"value,required"`
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
	Key string `json:"key,required"`
	// Value of the tag
	Value string `json:"value,required"`
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
	Data []V1UsageReportResponseData `json:"data,required"`
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
	ID string `json:"id,required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Customer id
	CustomerID string `json:"customerId,required"`
	// Feature id
	FeatureID string `json:"featureId,required"`
	// Timestamp
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The usage measurement record
	Value float64 `json:"value,required"`
	// The current measured usage value
	CurrentUsage float64 `json:"currentUsage,nullable"`
	// The date when the next usage reset will occur
	NextResetDate time.Time `json:"nextResetDate,nullable" format:"date-time"`
	// Resource id
	ResourceID string `json:"resourceId,nullable"`
	// The end date of the usage period in which this measurement resides (for
	// entitlements with a reset period)
	UsagePeriodEnd time.Time `json:"usagePeriodEnd,nullable" format:"date-time"`
	// The start date of the usage period in which this measurement resides (for
	// entitlements with a reset period)
	UsagePeriodStart time.Time `json:"usagePeriodStart,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		CustomerID       respjson.Field
		FeatureID        respjson.Field
		Timestamp        respjson.Field
		Value            respjson.Field
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

type V1UsageHistoryParams struct {
	CustomerID string `path:"customerId,required" json:"-"`
	// The start date of the range
	StartDate time.Time `query:"startDate,required" format:"date-time" json:"-"`
	// Resource id
	ResourceID param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	// The end date of the range
	EndDate param.Opt[time.Time] `query:"endDate,omitzero" format:"date-time" json:"-"`
	GroupBy param.Opt[string]    `query:"groupBy,omitzero" json:"-"`
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
	Usages []V1UsageReportParamsUsage `json:"usages,omitzero,required"`
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
	CustomerID string `json:"customerId,required"`
	// Feature id
	FeatureID string `json:"featureId,required"`
	// The value to report for usage
	Value int64 `json:"value,required"`
	// Resource id
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Timestamp of when the record was created
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
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
