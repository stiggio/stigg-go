// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"encoding/json"
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
	"github.com/stiggio/stigg-go/shared/constant"
)

// V1EventBetaCustomerEntitlementService contains methods and other services that
// help with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventBetaCustomerEntitlementService] method instead.
type V1EventBetaCustomerEntitlementService struct {
	Options []option.RequestOption
}

// NewV1EventBetaCustomerEntitlementService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1EventBetaCustomerEntitlementService(opts ...option.RequestOption) (r V1EventBetaCustomerEntitlementService) {
	r = V1EventBetaCustomerEntitlementService{}
	r.Options = opts
	return
}

// Experimental — request and response shapes may change without notice. Same
// semantics as `Check entitlement`, plus an optional `dimensions` query param that
// resolves to per-entity governance limits surfaced as `chains` on the response.
func (r *V1EventBetaCustomerEntitlementService) Check(ctx context.Context, id string, params V1EventBetaCustomerEntitlementCheckParams, opts ...option.RequestOption) (res *V1EventBetaCustomerEntitlementCheckResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1-beta/customers/%s/entitlements/check", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Response object
type V1EventBetaCustomerEntitlementCheckResponse struct {
	// Feature entitlement with optional governance chains attached.
	Data V1EventBetaCustomerEntitlementCheckResponseDataUnion `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerEntitlementCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntitlementCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1EventBetaCustomerEntitlementCheckResponseDataUnion contains all possible
// properties and values from
// [V1EventBetaCustomerEntitlementCheckResponseDataFeature],
// [V1EventBetaCustomerEntitlementCheckResponseDataCredit].
//
// Use the [V1EventBetaCustomerEntitlementCheckResponseDataUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1EventBetaCustomerEntitlementCheckResponseDataUnion struct {
	AccessDeniedReason string `json:"accessDeniedReason"`
	IsGranted          bool   `json:"isGranted"`
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type"`
	// This field is a union of
	// [[][]V1EventBetaCustomerEntitlementCheckResponseDataFeatureChain],
	// [[][]V1EventBetaCustomerEntitlementCheckResponseDataCreditChain]
	Chains               V1EventBetaCustomerEntitlementCheckResponseDataUnionChains `json:"chains"`
	CurrentUsage         float64                                                    `json:"currentUsage"`
	EntitlementUpdatedAt time.Time                                                  `json:"entitlementUpdatedAt"`
	// This field is from variant
	// [V1EventBetaCustomerEntitlementCheckResponseDataFeature].
	Feature V1EventBetaCustomerEntitlementCheckResponseDataFeatureFeature `json:"feature"`
	// This field is from variant
	// [V1EventBetaCustomerEntitlementCheckResponseDataFeature].
	HasUnlimitedUsage bool `json:"hasUnlimitedUsage"`
	// This field is from variant
	// [V1EventBetaCustomerEntitlementCheckResponseDataFeature].
	ResetPeriod string  `json:"resetPeriod"`
	UsageLimit  float64 `json:"usageLimit"`
	// This field is from variant
	// [V1EventBetaCustomerEntitlementCheckResponseDataFeature].
	UsagePeriodAnchor time.Time `json:"usagePeriodAnchor"`
	UsagePeriodEnd    time.Time `json:"usagePeriodEnd"`
	// This field is from variant
	// [V1EventBetaCustomerEntitlementCheckResponseDataFeature].
	UsagePeriodStart time.Time `json:"usagePeriodStart"`
	ValidUntil       time.Time `json:"validUntil"`
	// This field is from variant
	// [V1EventBetaCustomerEntitlementCheckResponseDataCredit].
	Currency V1EventBetaCustomerEntitlementCheckResponseDataCreditCurrency `json:"currency"`
	// This field is from variant
	// [V1EventBetaCustomerEntitlementCheckResponseDataCredit].
	UsageUpdatedAt time.Time `json:"usageUpdatedAt"`
	JSON           struct {
		AccessDeniedReason   respjson.Field
		IsGranted            respjson.Field
		Type                 respjson.Field
		Chains               respjson.Field
		CurrentUsage         respjson.Field
		EntitlementUpdatedAt respjson.Field
		Feature              respjson.Field
		HasUnlimitedUsage    respjson.Field
		ResetPeriod          respjson.Field
		UsageLimit           respjson.Field
		UsagePeriodAnchor    respjson.Field
		UsagePeriodEnd       respjson.Field
		UsagePeriodStart     respjson.Field
		ValidUntil           respjson.Field
		Currency             respjson.Field
		UsageUpdatedAt       respjson.Field
		raw                  string
	} `json:"-"`
}

// anyV1EventBetaCustomerEntitlementCheckResponseData is implemented by each
// variant of [V1EventBetaCustomerEntitlementCheckResponseDataUnion] to add type
// safety for the return type of
// [V1EventBetaCustomerEntitlementCheckResponseDataUnion.AsAny]
type anyV1EventBetaCustomerEntitlementCheckResponseData interface {
	implV1EventBetaCustomerEntitlementCheckResponseDataUnion()
}

func (V1EventBetaCustomerEntitlementCheckResponseDataFeature) implV1EventBetaCustomerEntitlementCheckResponseDataUnion() {
}
func (V1EventBetaCustomerEntitlementCheckResponseDataCredit) implV1EventBetaCustomerEntitlementCheckResponseDataUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := V1EventBetaCustomerEntitlementCheckResponseDataUnion.AsAny().(type) {
//	case stigg.V1EventBetaCustomerEntitlementCheckResponseDataFeature:
//	case stigg.V1EventBetaCustomerEntitlementCheckResponseDataCredit:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u V1EventBetaCustomerEntitlementCheckResponseDataUnion) AsAny() anyV1EventBetaCustomerEntitlementCheckResponseData {
	switch u.Type {
	case "FEATURE":
		return u.AsFeature()
	case "CREDIT":
		return u.AsCredit()
	}
	return nil
}

func (u V1EventBetaCustomerEntitlementCheckResponseDataUnion) AsFeature() (v V1EventBetaCustomerEntitlementCheckResponseDataFeature) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1EventBetaCustomerEntitlementCheckResponseDataUnion) AsCredit() (v V1EventBetaCustomerEntitlementCheckResponseDataCredit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1EventBetaCustomerEntitlementCheckResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1EventBetaCustomerEntitlementCheckResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1EventBetaCustomerEntitlementCheckResponseDataUnionChains is an implicit
// subunion of [V1EventBetaCustomerEntitlementCheckResponseDataUnion].
// V1EventBetaCustomerEntitlementCheckResponseDataUnionChains provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1EventBetaCustomerEntitlementCheckResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBetaChains]
type V1EventBetaCustomerEntitlementCheckResponseDataUnionChains struct {
	// This field will be present if the value is a
	// [[][]V1EventBetaCustomerEntitlementCheckResponseDataFeatureChain] instead of an
	// object.
	OfBetaChains [][]V1EventBetaCustomerEntitlementCheckResponseDataFeatureChain `json:",inline"`
	JSON         struct {
		OfBetaChains respjson.Field
		raw          string
	} `json:"-"`
}

func (r *V1EventBetaCustomerEntitlementCheckResponseDataUnionChains) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature entitlement with optional governance chains attached.
type V1EventBetaCustomerEntitlementCheckResponseDataFeature struct {
	// Any of "FeatureNotFound", "CustomerNotFound", "CustomerIsArchived",
	// "CustomerResourceNotFound", "NoActiveSubscription",
	// "NoFeatureEntitlementInSubscription", "RequestedUsageExceedingLimit",
	// "RequestedValuesMismatch", "BudgetExceeded", "Unknown", "FeatureTypeMismatch",
	// "Revoked", "InsufficientCredits", "EntitlementNotFound".
	AccessDeniedReason string           `json:"accessDeniedReason" api:"required"`
	IsGranted          bool             `json:"isGranted" api:"required"`
	Type               constant.Feature `json:"type" default:"FEATURE"`
	// Per-entity rollups, one chain per resolved dimension. Omitted when dimensions
	// was not provided.
	Chains       [][]V1EventBetaCustomerEntitlementCheckResponseDataFeatureChain `json:"chains"`
	CurrentUsage float64                                                         `json:"currentUsage"`
	// Timestamp of the last update to the entitlement grant or configuration.
	EntitlementUpdatedAt time.Time                                                     `json:"entitlementUpdatedAt" format:"date-time"`
	Feature              V1EventBetaCustomerEntitlementCheckResponseDataFeatureFeature `json:"feature"`
	HasUnlimitedUsage    bool                                                          `json:"hasUnlimitedUsage"`
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string  `json:"resetPeriod" api:"nullable"`
	UsageLimit  float64 `json:"usageLimit" api:"nullable"`
	// The anchor for calculating the usage period for metered entitlements with a
	// reset period configured
	UsagePeriodAnchor time.Time `json:"usagePeriodAnchor" format:"date-time"`
	// The end date of the usage period for metered entitlements with a reset period
	// configured
	UsagePeriodEnd time.Time `json:"usagePeriodEnd" format:"date-time"`
	// The start date of the usage period for metered entitlements with a reset period
	// configured
	UsagePeriodStart time.Time `json:"usagePeriodStart" format:"date-time"`
	// The next time the entitlement should be recalculated
	ValidUntil time.Time `json:"validUntil" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessDeniedReason   respjson.Field
		IsGranted            respjson.Field
		Type                 respjson.Field
		Chains               respjson.Field
		CurrentUsage         respjson.Field
		EntitlementUpdatedAt respjson.Field
		Feature              respjson.Field
		HasUnlimitedUsage    respjson.Field
		ResetPeriod          respjson.Field
		UsageLimit           respjson.Field
		UsagePeriodAnchor    respjson.Field
		UsagePeriodEnd       respjson.Field
		UsagePeriodStart     respjson.Field
		ValidUntil           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerEntitlementCheckResponseDataFeature) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntitlementCheckResponseDataFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-entity governance node — limit and current usage for a single resolved
// entity.
type V1EventBetaCustomerEntitlementCheckResponseDataFeatureChain struct {
	// Amount consumed by this entity in the current cadence period.
	CurrentUsage float64 `json:"currentUsage" api:"required"`
	// External id of the entity within the customer.
	EntityID string `json:"entityId" api:"required"`
	// Whether this node alone permits the requested usage.
	IsGranted bool `json:"isGranted" api:"required"`
	// Hard usage limit for this node; null when no assignment is configured.
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentUsage respjson.Field
		EntityID     respjson.Field
		IsGranted    respjson.Field
		UsageLimit   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerEntitlementCheckResponseDataFeatureChain) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventBetaCustomerEntitlementCheckResponseDataFeatureChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventBetaCustomerEntitlementCheckResponseDataFeatureFeature struct {
	// The unique reference ID of the entitlement.
	ID string `json:"id" api:"required"`
	// The human-readable name of the entitlement, shown in UI elements.
	DisplayName string `json:"displayName" api:"required"`
	// The current status of the feature.
	//
	// Any of "NEW", "SUSPENDED", "ACTIVE".
	FeatureStatus string `json:"featureStatus" api:"required"`
	// The type of feature associated with the entitlement.
	//
	// Any of "BOOLEAN", "NUMBER", "ENUM".
	FeatureType string `json:"featureType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		DisplayName   respjson.Field
		FeatureStatus respjson.Field
		FeatureType   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerEntitlementCheckResponseDataFeatureFeature) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventBetaCustomerEntitlementCheckResponseDataFeatureFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement with optional governance chains attached.
type V1EventBetaCustomerEntitlementCheckResponseDataCredit struct {
	// Any of "FeatureNotFound", "CustomerNotFound", "CustomerIsArchived",
	// "CustomerResourceNotFound", "NoActiveSubscription",
	// "NoFeatureEntitlementInSubscription", "RequestedUsageExceedingLimit",
	// "RequestedValuesMismatch", "BudgetExceeded", "Unknown", "FeatureTypeMismatch",
	// "Revoked", "InsufficientCredits", "EntitlementNotFound".
	AccessDeniedReason string `json:"accessDeniedReason" api:"required"`
	// The currency associated with a credit entitlement.
	Currency     V1EventBetaCustomerEntitlementCheckResponseDataCreditCurrency `json:"currency" api:"required"`
	CurrentUsage float64                                                       `json:"currentUsage" api:"required"`
	IsGranted    bool                                                          `json:"isGranted" api:"required"`
	Type         constant.Credit                                               `json:"type" default:"CREDIT"`
	UsageLimit   float64                                                       `json:"usageLimit" api:"required"`
	// Timestamp of the last update to the credit usage.
	UsageUpdatedAt time.Time `json:"usageUpdatedAt" api:"required" format:"date-time"`
	// Per-entity rollups, one chain per resolved dimension. Omitted when dimensions
	// was not provided.
	Chains [][]V1EventBetaCustomerEntitlementCheckResponseDataCreditChain `json:"chains"`
	// Timestamp of the last update to the entitlement grant or configuration.
	EntitlementUpdatedAt time.Time `json:"entitlementUpdatedAt" format:"date-time"`
	// The end date of the current billing period for recurring credit grants.
	UsagePeriodEnd time.Time `json:"usagePeriodEnd" format:"date-time"`
	// The next time the entitlement should be recalculated
	ValidUntil time.Time `json:"validUntil" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessDeniedReason   respjson.Field
		Currency             respjson.Field
		CurrentUsage         respjson.Field
		IsGranted            respjson.Field
		Type                 respjson.Field
		UsageLimit           respjson.Field
		UsageUpdatedAt       respjson.Field
		Chains               respjson.Field
		EntitlementUpdatedAt respjson.Field
		UsagePeriodEnd       respjson.Field
		ValidUntil           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerEntitlementCheckResponseDataCredit) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerEntitlementCheckResponseDataCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The currency associated with a credit entitlement.
type V1EventBetaCustomerEntitlementCheckResponseDataCreditCurrency struct {
	// The unique identifier of the custom currency.
	CurrencyID string `json:"currencyId" api:"required"`
	// The display name of the currency.
	DisplayName string `json:"displayName" api:"required"`
	// A description of the currency.
	Description string `json:"description" api:"nullable"`
	// Additional metadata associated with the currency.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// The plural form of the currency unit.
	UnitPlural string `json:"unitPlural" api:"nullable"`
	// The singular form of the currency unit.
	UnitSingular string `json:"unitSingular" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyID   respjson.Field
		DisplayName  respjson.Field
		Description  respjson.Field
		Metadata     respjson.Field
		UnitPlural   respjson.Field
		UnitSingular respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerEntitlementCheckResponseDataCreditCurrency) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventBetaCustomerEntitlementCheckResponseDataCreditCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-entity governance node — limit and current usage for a single resolved
// entity.
type V1EventBetaCustomerEntitlementCheckResponseDataCreditChain struct {
	// Amount consumed by this entity in the current cadence period.
	CurrentUsage float64 `json:"currentUsage" api:"required"`
	// External id of the entity within the customer.
	EntityID string `json:"entityId" api:"required"`
	// Whether this node alone permits the requested usage.
	IsGranted bool `json:"isGranted" api:"required"`
	// Hard usage limit for this node; null when no assignment is configured.
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentUsage respjson.Field
		EntityID     respjson.Field
		IsGranted    respjson.Field
		UsageLimit   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerEntitlementCheckResponseDataCreditChain) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventBetaCustomerEntitlementCheckResponseDataCreditChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventBetaCustomerEntitlementCheckParams struct {
	// Currency ID (refId) to check for credit entitlements. Mutually exclusive with
	// `featureId`.
	CurrencyID param.Opt[string] `query:"currencyId,omitzero" json:"-"`
	// Feature ID (refId) to check. Mutually exclusive with `currencyId`.
	FeatureID param.Opt[string] `query:"featureId,omitzero" json:"-"`
	// Requested usage amount to evaluate against the entitlement limit (numeric
	// features only)
	RequestedUsage param.Opt[int64] `query:"requestedUsage,omitzero" json:"-"`
	// Resource ID to scope the entitlement check to a specific resource
	ResourceID     param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	// Optional attribution map (e.g. `dimensions[userId]=u1`). When provided, the
	// response includes a `chains` array with per-entity governance limits.
	Dimensions map[string]string `query:"dimensions,omitzero" json:"-"`
	// Requested values to evaluate against allowed values (enum features only)
	RequestedValues []string `query:"requestedValues,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventBetaCustomerEntitlementCheckParams]'s query
// parameters as `url.Values`.
func (r V1EventBetaCustomerEntitlementCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
