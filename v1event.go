// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stiggio/stigg-go/internal/apijson"
	"github.com/stiggio/stigg-go/internal/requestconfig"
	"github.com/stiggio/stigg-go/option"
	"github.com/stiggio/stigg-go/packages/param"
	"github.com/stiggio/stigg-go/packages/respjson"
)

// V1EventService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventService] method instead.
type V1EventService struct {
	Options  []option.RequestOption
	Features V1EventFeatureService
	Addons   V1EventAddonService
}

// NewV1EventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1EventService(opts ...option.RequestOption) (r V1EventService) {
	r = V1EventService{}
	r.Options = opts
	r.Features = NewV1EventFeatureService(opts...)
	r.Addons = NewV1EventAddonService(opts...)
	return
}

// Reports raw usage events for event-based metering. Events are ingested
// asynchronously and aggregated into usage totals.
func (r *V1EventService) Report(ctx context.Context, body V1EventReportParams, opts ...option.RequestOption) (res *V1EventReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response object
type V1EventReportResponse struct {
	// Empty success response confirming that events were successfully ingested and
	// queued for processing by Stigg's metering system.
	Data any `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventReportResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventReportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventReportParams struct {
	// A list of usage events to report
	Events []V1EventReportParamsEvent `json:"events,omitzero,required"`
	paramObj
}

func (r V1EventReportParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventReportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventReportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Raw usage event
//
// The properties CustomerID, EventName, IdempotencyKey are required.
type V1EventReportParamsEvent struct {
	// Customer id
	CustomerID string `json:"customerId,required"`
	// The name of the usage event
	EventName string `json:"eventName,required"`
	// Idempotency key
	IdempotencyKey string `json:"idempotencyKey,required"`
	// Resource id
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Timestamp
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	// Dimensions associated with the usage event
	Dimensions map[string]V1EventReportParamsEventDimensionUnion `json:"dimensions,omitzero"`
	paramObj
}

func (r V1EventReportParamsEvent) MarshalJSON() (data []byte, err error) {
	type shadow V1EventReportParamsEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventReportParamsEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1EventReportParamsEventDimensionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u V1EventReportParamsEventDimensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *V1EventReportParamsEventDimensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1EventReportParamsEventDimensionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}
