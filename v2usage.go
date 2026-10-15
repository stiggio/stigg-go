// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
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

// Operations related to usage & metering
//
// V2UsageService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2UsageService] method instead.
type V2UsageService struct {
	Options []option.RequestOption
}

// NewV2UsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2UsageService(opts ...option.RequestOption) (r V2UsageService) {
	r = V2UsageService{}
	r.Options = opts
	return
}

// Reports usage measurements for metered features. The reported usage is used to
// track, limit, and bill customer consumption. Credit balances are returned once
// per wallet in `credits`, not on each measurement.
func (r *V2UsageService) Report(ctx context.Context, params V2UsageReportParams, opts ...option.RequestOption) (res *V2UsageReportResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v2/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Response containing reported usage measurements with current usage values,
// period information and reset dates, plus the balance of each credit wallet they
// drew from.
type V2UsageReportResponse struct {
	// Reported usage measurements and the credit balances they drew from
	Data V2UsageReportResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2UsageReportResponse) RawJSON() string { return r.JSON.raw }
func (r *V2UsageReportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reported usage measurements and the credit balances they drew from
type V2UsageReportResponseData struct {
	// Usage measurements, in request order, with current values and period info
	Usages []V2UsageReportResponseDataUsage `json:"usages" api:"required"`
	// The balance of each credit wallet the request drew from. Omitted when no
	// reported feature consumes credits, or when credit balances aren't computed
	// synchronously for the environment; each usage's `credit` is then `null`.
	Credits []V2UsageReportResponseDataCredit `json:"credits"`
	// The current usage of each governance entitlement this request counted towards.
	// Absent when governance is not enabled or the usage is unknown. Entitlements
	// whose usage could not be read, and those of repeated idempotent requests, are
	// left out
	Entities []V2UsageReportResponseDataEntity `json:"entities"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Usages      respjson.Field
		Credits     respjson.Field
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2UsageReportResponseData) RawJSON() string { return r.JSON.raw }
func (r *V2UsageReportResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recorded usage with period info and the credits it deducted
type V2UsageReportResponseDataUsage struct {
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
	// The credits this measurement deducted. `null` when the measurement drew from no
	// wallet listed in `credits`. The balance of the wallet it drew from is in the
	// response `credits` list, matched by customer, resource and currency.
	Credit V2UsageReportResponseDataUsageCredit `json:"credit" api:"nullable"`
	// The current measured usage value
	CurrentUsage float64 `json:"currentUsage" api:"nullable"`
	// The date when the next usage reset will occur
	NextResetDate time.Time `json:"nextResetDate" api:"nullable" format:"date-time"`
	// The customer resource this usage applies to. Optional — only required if the
	// customer has multiple resources (for example, one subscription per workspace or
	// site) and usage needs to be tracked separately per resource; omit it to report
	// usage at the customer level.
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
func (r V2UsageReportResponseDataUsage) RawJSON() string { return r.JSON.raw }
func (r *V2UsageReportResponseDataUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The credits this measurement deducted. `null` when the measurement drew from no
// wallet listed in `credits`. The balance of the wallet it drew from is in the
// response `credits` list, matched by customer, resource and currency.
type V2UsageReportResponseDataUsageCredit struct {
	// The credits this single reportUsage call deducted, in credit units — scoped to
	// this one measurement (0 for idempotency duplicates). Contrast `currentUsage`,
	// which is the wallet-wide running total shared across all features on this
	// currency. Use it to reconcile expected per-call deductions.
	Consumed float64 `json:"consumed" api:"required"`
	// The credit currency identifier
	CurrencyID string `json:"currencyId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Consumed    respjson.Field
		CurrencyID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2UsageReportResponseDataUsageCredit) RawJSON() string { return r.JSON.raw }
func (r *V2UsageReportResponseDataUsageCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optimistic credit balance of one wallet (customer, resource and currency) after
// the request, listed once however many measurements drew from it
type V2UsageReportResponseDataCredit struct {
	// The credit currency identifier
	CurrencyID string `json:"currencyId" api:"required"`
	// The wallet's total consumed credits for this currency (optimistic — includes
	// not-yet-reconciled usage), shared across every feature that draws on the
	// currency. This is the running balance, not this call's deduction — see
	// `consumed` for that.
	CurrentUsage float64 `json:"currentUsage" api:"required"`
	// Customer id
	CustomerID string `json:"customerId" api:"required"`
	// The wallet's resource; `null` for the customer-level wallet
	ResourceID string `json:"resourceId" api:"required"`
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
		CurrencyID     respjson.Field
		CurrentUsage   respjson.Field
		CustomerID     respjson.Field
		ResourceID     respjson.Field
		Timestamp      respjson.Field
		UsageLimit     respjson.Field
		UsagePeriodEnd respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2UsageReportResponseDataCredit) RawJSON() string { return r.JSON.raw }
func (r *V2UsageReportResponseDataCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The usage counted against one governance entitlement
type V2UsageReportResponseDataEntity struct {
	// The absolute usage counted against the entitlement
	CurrentUsage float64 `json:"currentUsage" api:"required"`
	// The customer whose usage counted towards the entitlement
	CustomerID string `json:"customerId" api:"required"`
	// The governance entitlement the usage counts towards
	EntitlementID string `json:"entitlementId" api:"required"`
	// The governance entity the entitlement belongs to
	EntityID string `json:"entityId" api:"required"`
	// The customer resource the usage counted from; `null` for customer-level usage
	ResourceID string `json:"resourceId" api:"required"`
	// The entities the entitlement is scoped to (empty when it is not scoped)
	ScopeEntityIDs []string `json:"scopeEntityIds" api:"required"`
	// When the entitlement usage was last updated
	UsageUpdatedAt time.Time `json:"usageUpdatedAt" api:"required" format:"date-time"`
	// The credit currency the entitlement governs (exactly one of featureId and
	// currencyId is set)
	CurrencyID string `json:"currencyId"`
	// The feature the entitlement governs (exactly one of featureId and currencyId is
	// set)
	FeatureID string `json:"featureId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentUsage   respjson.Field
		CustomerID     respjson.Field
		EntitlementID  respjson.Field
		EntityID       respjson.Field
		ResourceID     respjson.Field
		ScopeEntityIDs respjson.Field
		UsageUpdatedAt respjson.Field
		CurrencyID     respjson.Field
		FeatureID      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2UsageReportResponseDataEntity) RawJSON() string { return r.JSON.raw }
func (r *V2UsageReportResponseDataEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2UsageReportParams struct {
	// A list of usage reports to be submitted in bulk
	Usages         []V2UsageReportParamsUsage `json:"usages,omitzero" api:"required"`
	XAccountID     param.Opt[string]          `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]          `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V2UsageReportParams) MarshalJSON() (data []byte, err error) {
	type shadow V2UsageReportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2UsageReportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single usage measurement
//
// The properties CustomerID, FeatureID, Value are required.
type V2UsageReportParamsUsage struct {
	// Customer id
	CustomerID string `json:"customerId" api:"required"`
	// Feature id
	FeatureID string `json:"featureId" api:"required"`
	// The value to report for usage. Must be a whole number — the REST API does not
	// accept fractional (float) usage values; scale up (e.g. report cents instead of
	// dollars, or milliseconds instead of seconds) if you need sub-unit precision.
	Value int64 `json:"value" api:"required"`
	// The customer resource this usage applies to. Optional — only required if the
	// customer has multiple resources (for example, one subscription per workspace or
	// site) and usage needs to be tracked separately per resource; omit it to report
	// usage at the customer level.
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Timestamp of when the record was created
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// A key you provide to safely retry the same usage report without double-counting
	// it. Reports with a previously-seen idempotency key are deduplicated for 7 days;
	// after that window a retry is treated as new usage.
	IdempotencyKey param.Opt[string] `json:"idempotencyKey,omitzero"`
	// Additional dimensions for the usage report
	Dimensions map[string]V2UsageReportParamsUsageDimensionUnion `json:"dimensions,omitzero"`
	// How the reported value is applied: DELTA (default) adds it to the feature's
	// current usage; SET treats it as the new absolute usage total, and Stigg computes
	// the delta internally.
	//
	// Any of "DELTA", "SET".
	UpdateBehavior string `json:"updateBehavior,omitzero"`
	paramObj
}

func (r V2UsageReportParamsUsage) MarshalJSON() (data []byte, err error) {
	type shadow V2UsageReportParamsUsage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2UsageReportParamsUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2UsageReportParamsUsage](
		"updateBehavior", "DELTA", "SET",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V2UsageReportParamsUsageDimensionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u V2UsageReportParamsUsageDimensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *V2UsageReportParamsUsageDimensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V2UsageReportParamsUsageDimensionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}
