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

// V1Service contains methods and other services that help with interacting with
// the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	Options       []option.RequestOption
	Customers     V1CustomerService
	Subscriptions V1SubscriptionService
	Coupons       V1CouponService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r V1Service) {
	r = V1Service{}
	r.Options = opts
	r.Customers = NewV1CustomerService(opts...)
	r.Subscriptions = NewV1SubscriptionService(opts...)
	r.Coupons = NewV1CouponService(opts...)
	return
}

// Report usage events
func (r *V1Service) NewEvent(ctx context.Context, body V1NewEventParams, opts ...option.RequestOption) (res *V1NewEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Report usage measurements
func (r *V1Service) NewUsage(ctx context.Context, body V1NewUsageParams, opts ...option.RequestOption) (res *V1NewUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response object
type V1NewEventResponse struct {
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
func (r V1NewEventResponse) RawJSON() string { return r.JSON.raw }
func (r *V1NewEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing reported usage measurements with current usage values,
// period information, and reset dates for each measurement.
type V1NewUsageResponse struct {
	// Array of usage measurements with current values and period info
	Data []V1NewUsageResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1NewUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *V1NewUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recorded usage with period info
type V1NewUsageResponseData struct {
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
func (r V1NewUsageResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1NewUsageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1NewEventParams struct {
	// A list of usage events to report
	Events []V1NewEventParamsEvent `json:"events,omitzero,required"`
	paramObj
}

func (r V1NewEventParams) MarshalJSON() (data []byte, err error) {
	type shadow V1NewEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1NewEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Raw usage event
//
// The properties CustomerID, EventName, IdempotencyKey are required.
type V1NewEventParamsEvent struct {
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
	Dimensions map[string]V1NewEventParamsEventDimensionUnion `json:"dimensions,omitzero"`
	paramObj
}

func (r V1NewEventParamsEvent) MarshalJSON() (data []byte, err error) {
	type shadow V1NewEventParamsEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1NewEventParamsEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1NewEventParamsEventDimensionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u V1NewEventParamsEventDimensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *V1NewEventParamsEventDimensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1NewEventParamsEventDimensionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type V1NewUsageParams struct {
	// A list of usage reports to be submitted in bulk
	Usages []V1NewUsageParamsUsage `json:"usages,omitzero,required"`
	paramObj
}

func (r V1NewUsageParams) MarshalJSON() (data []byte, err error) {
	type shadow V1NewUsageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1NewUsageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single usage measurement
//
// The properties CustomerID, FeatureID, Value are required.
type V1NewUsageParamsUsage struct {
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
	Dimensions map[string]V1NewUsageParamsUsageDimensionUnion `json:"dimensions,omitzero"`
	// The method by which the usage value should be updated
	//
	// Any of "DELTA", "SET".
	UpdateBehavior string `json:"updateBehavior,omitzero"`
	paramObj
}

func (r V1NewUsageParamsUsage) MarshalJSON() (data []byte, err error) {
	type shadow V1NewUsageParamsUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1NewUsageParamsUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1NewUsageParamsUsage](
		"updateBehavior", "DELTA", "SET",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1NewUsageParamsUsageDimensionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u V1NewUsageParamsUsageDimensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *V1NewUsageParamsUsageDimensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1NewUsageParamsUsageDimensionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}
