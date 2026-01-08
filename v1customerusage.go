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

	"github.com/stainless-sdks/stigg-go/internal/apijson"
	"github.com/stainless-sdks/stigg-go/internal/apiquery"
	"github.com/stainless-sdks/stigg-go/internal/requestconfig"
	"github.com/stainless-sdks/stigg-go/option"
	"github.com/stainless-sdks/stigg-go/packages/param"
	"github.com/stainless-sdks/stigg-go/packages/respjson"
)

// V1CustomerUsageService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerUsageService] method instead.
type V1CustomerUsageService struct {
	Options []option.RequestOption
}

// NewV1CustomerUsageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomerUsageService(opts ...option.RequestOption) (r V1CustomerUsageService) {
	r = V1CustomerUsageService{}
	r.Options = opts
	return
}

// Perform retrieval on a Usage history
func (r *V1CustomerUsageService) Get(ctx context.Context, featureID string, params V1CustomerUsageGetParams, opts ...option.RequestOption) (res *V1CustomerUsageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CustomerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	if featureID == "" {
		err = errors.New("missing required featureId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/usage/features/%s", params.CustomerID, featureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type V1CustomerUsageGetResponse struct {
	Data V1CustomerUsageGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerUsageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerUsageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerUsageGetResponseData struct {
	// Markers for events that affecting feature usage
	Markers []V1CustomerUsageGetResponseDataMarker `json:"markers,required"`
	// Series of usage history
	Series []V1CustomerUsageGetResponseDataSeries `json:"series,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markers     respjson.Field
		Series      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerUsageGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerUsageGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerUsageGetResponseDataMarker struct {
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
func (r V1CustomerUsageGetResponseDataMarker) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerUsageGetResponseDataMarker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerUsageGetResponseDataSeries struct {
	// Points in the usage history series
	Points []V1CustomerUsageGetResponseDataSeriesPoint `json:"points,required"`
	// Tags for the usage history series
	Tags []V1CustomerUsageGetResponseDataSeriesTag `json:"tags,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Points      respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerUsageGetResponseDataSeries) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerUsageGetResponseDataSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerUsageGetResponseDataSeriesPoint struct {
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
func (r V1CustomerUsageGetResponseDataSeriesPoint) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerUsageGetResponseDataSeriesPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerUsageGetResponseDataSeriesTag struct {
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
func (r V1CustomerUsageGetResponseDataSeriesTag) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerUsageGetResponseDataSeriesTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerUsageGetParams struct {
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

// URLQuery serializes [V1CustomerUsageGetParams]'s query parameters as
// `url.Values`.
func (r V1CustomerUsageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
