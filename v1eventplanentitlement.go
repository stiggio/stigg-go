// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"encoding/json"
	"errors"
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

// V1EventPlanEntitlementService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventPlanEntitlementService] method instead.
type V1EventPlanEntitlementService struct {
	Options []option.RequestOption
}

// NewV1EventPlanEntitlementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1EventPlanEntitlementService(opts ...option.RequestOption) (r V1EventPlanEntitlementService) {
	r = V1EventPlanEntitlementService{}
	r.Options = opts
	return
}

// Creates one or more entitlements (feature or credit) on a draft plan.
func (r *V1EventPlanEntitlementService) New(ctx context.Context, planID string, body V1EventPlanEntitlementNewParams, opts ...option.RequestOption) (res *V1EventPlanEntitlementNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required planId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s/entitlements", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing entitlement on a draft plan.
func (r *V1EventPlanEntitlementService) Update(ctx context.Context, id string, params V1EventPlanEntitlementUpdateParams, opts ...option.RequestOption) (res *PlanEntitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.PlanID == "" {
		err = errors.New("missing required planId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s/entitlements/%s", params.PlanID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieves a list of entitlements for a plan.
func (r *V1EventPlanEntitlementService) List(ctx context.Context, planID string, opts ...option.RequestOption) (res *V1EventPlanEntitlementListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required planId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s/entitlements", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an entitlement from a draft plan.
func (r *V1EventPlanEntitlementService) Delete(ctx context.Context, id string, body V1EventPlanEntitlementDeleteParams, opts ...option.RequestOption) (res *PlanEntitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.PlanID == "" {
		err = errors.New("missing required planId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s/entitlements/%s", body.PlanID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Response object
type PlanEntitlement struct {
	// Feature or credit entitlement on a plan
	Data PlanEntitlementData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanEntitlement) RawJSON() string { return r.JSON.raw }
func (r *PlanEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature or credit entitlement on a plan
type PlanEntitlementData struct {
	// Unique identifier of the entitlement
	ID string `json:"id" api:"required"`
	// Credit amount (for credit entitlements)
	Amount float64 `json:"amount" api:"required"`
	// Entitlement behavior (Increment or Override)
	//
	// Any of "Increment", "Override".
	Behavior string `json:"behavior" api:"required"`
	// Credit grant cadence (for credit entitlements)
	//
	// Any of "MONTH", "YEAR".
	Cadence string `json:"cadence" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Custom currency ID (for credit entitlements)
	CustomCurrencyID string `json:"customCurrencyId" api:"required"`
	// Optional description of the entitlement
	Description string `json:"description" api:"required"`
	// Override display name for the entitlement
	DisplayNameOverride string `json:"displayNameOverride" api:"required"`
	// Allowed enum values (for feature entitlements)
	EnumValues []string `json:"enumValues" api:"required"`
	// Feature ID (for feature entitlements)
	FeatureID string `json:"featureId" api:"required"`
	// Whether the usage limit is a soft limit (for feature entitlements)
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// Whether usage is unlimited (for feature entitlements)
	HasUnlimitedUsage bool `json:"hasUnlimitedUsage" api:"required"`
	// Widget types where this entitlement is hidden
	//
	// Any of "PAYWALL", "CUSTOMER_PORTAL", "CHECKOUT".
	HiddenFromWidgets []string `json:"hiddenFromWidgets" api:"required"`
	// Whether this is a custom entitlement
	IsCustom bool `json:"isCustom" api:"required"`
	// Whether the entitlement is granted
	IsGranted bool `json:"isGranted" api:"required"`
	// Display order of the entitlement
	Order float64 `json:"order" api:"required"`
	// Usage reset period (for feature entitlements)
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod" api:"required"`
	// Reset period configuration (for feature entitlements)
	ResetPeriodConfiguration PlanEntitlementDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
	// Entitlement type (FEATURE or CREDIT)
	//
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Usage limit (for feature entitlements)
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Amount                   respjson.Field
		Behavior                 respjson.Field
		Cadence                  respjson.Field
		CreatedAt                respjson.Field
		CustomCurrencyID         respjson.Field
		Description              respjson.Field
		DisplayNameOverride      respjson.Field
		EnumValues               respjson.Field
		FeatureID                respjson.Field
		HasSoftLimit             respjson.Field
		HasUnlimitedUsage        respjson.Field
		HiddenFromWidgets        respjson.Field
		IsCustom                 respjson.Field
		IsGranted                respjson.Field
		Order                    respjson.Field
		ResetPeriod              respjson.Field
		ResetPeriodConfiguration respjson.Field
		Type                     respjson.Field
		UpdatedAt                respjson.Field
		UsageLimit               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanEntitlementData) RawJSON() string { return r.JSON.raw }
func (r *PlanEntitlementData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PlanEntitlementDataResetPeriodConfigurationUnion contains all possible
// properties and values from
// [PlanEntitlementDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [PlanEntitlementDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [PlanEntitlementDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PlanEntitlementDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u PlanEntitlementDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v PlanEntitlementDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlanEntitlementDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v PlanEntitlementDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlanEntitlementDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v PlanEntitlementDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlanEntitlementDataResetPeriodConfigurationUnion) RawJSON() string { return u.JSON.raw }

func (r *PlanEntitlementDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type PlanEntitlementDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanEntitlementDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *PlanEntitlementDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type PlanEntitlementDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanEntitlementDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *PlanEntitlementDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type PlanEntitlementDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanEntitlementDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *PlanEntitlementDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1EventPlanEntitlementNewResponse struct {
	Data []V1EventPlanEntitlementNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanEntitlementNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanEntitlementNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature or credit entitlement on a plan
type V1EventPlanEntitlementNewResponseData struct {
	// Unique identifier of the entitlement
	ID string `json:"id" api:"required"`
	// Credit amount (for credit entitlements)
	Amount float64 `json:"amount" api:"required"`
	// Entitlement behavior (Increment or Override)
	//
	// Any of "Increment", "Override".
	Behavior string `json:"behavior" api:"required"`
	// Credit grant cadence (for credit entitlements)
	//
	// Any of "MONTH", "YEAR".
	Cadence string `json:"cadence" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Custom currency ID (for credit entitlements)
	CustomCurrencyID string `json:"customCurrencyId" api:"required"`
	// Optional description of the entitlement
	Description string `json:"description" api:"required"`
	// Override display name for the entitlement
	DisplayNameOverride string `json:"displayNameOverride" api:"required"`
	// Allowed enum values (for feature entitlements)
	EnumValues []string `json:"enumValues" api:"required"`
	// Feature ID (for feature entitlements)
	FeatureID string `json:"featureId" api:"required"`
	// Whether the usage limit is a soft limit (for feature entitlements)
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// Whether usage is unlimited (for feature entitlements)
	HasUnlimitedUsage bool `json:"hasUnlimitedUsage" api:"required"`
	// Widget types where this entitlement is hidden
	//
	// Any of "PAYWALL", "CUSTOMER_PORTAL", "CHECKOUT".
	HiddenFromWidgets []string `json:"hiddenFromWidgets" api:"required"`
	// Whether this is a custom entitlement
	IsCustom bool `json:"isCustom" api:"required"`
	// Whether the entitlement is granted
	IsGranted bool `json:"isGranted" api:"required"`
	// Display order of the entitlement
	Order float64 `json:"order" api:"required"`
	// Usage reset period (for feature entitlements)
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod" api:"required"`
	// Reset period configuration (for feature entitlements)
	ResetPeriodConfiguration V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
	// Entitlement type (FEATURE or CREDIT)
	//
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Usage limit (for feature entitlements)
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Amount                   respjson.Field
		Behavior                 respjson.Field
		Cadence                  respjson.Field
		CreatedAt                respjson.Field
		CustomCurrencyID         respjson.Field
		Description              respjson.Field
		DisplayNameOverride      respjson.Field
		EnumValues               respjson.Field
		FeatureID                respjson.Field
		HasSoftLimit             respjson.Field
		HasUnlimitedUsage        respjson.Field
		HiddenFromWidgets        respjson.Field
		IsCustom                 respjson.Field
		IsGranted                respjson.Field
		Order                    respjson.Field
		ResetPeriod              respjson.Field
		ResetPeriodConfiguration respjson.Field
		Type                     respjson.Field
		UpdatedAt                respjson.Field
		UsageLimit               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanEntitlementNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanEntitlementNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationUnion contains all
// possible properties and values from
// [V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventPlanEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response list object
type V1EventPlanEntitlementListResponse struct {
	Data []V1EventPlanEntitlementListResponseData `json:"data" api:"required"`
	// Pagination metadata including cursors for navigating through results
	Pagination V1EventPlanEntitlementListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanEntitlementListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanEntitlementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature or credit entitlement on a plan
type V1EventPlanEntitlementListResponseData struct {
	// Unique identifier of the entitlement
	ID string `json:"id" api:"required"`
	// Credit amount (for credit entitlements)
	Amount float64 `json:"amount" api:"required"`
	// Entitlement behavior (Increment or Override)
	//
	// Any of "Increment", "Override".
	Behavior string `json:"behavior" api:"required"`
	// Credit grant cadence (for credit entitlements)
	//
	// Any of "MONTH", "YEAR".
	Cadence string `json:"cadence" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Custom currency ID (for credit entitlements)
	CustomCurrencyID string `json:"customCurrencyId" api:"required"`
	// Optional description of the entitlement
	Description string `json:"description" api:"required"`
	// Override display name for the entitlement
	DisplayNameOverride string `json:"displayNameOverride" api:"required"`
	// Allowed enum values (for feature entitlements)
	EnumValues []string `json:"enumValues" api:"required"`
	// Feature ID (for feature entitlements)
	FeatureID string `json:"featureId" api:"required"`
	// Whether the usage limit is a soft limit (for feature entitlements)
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// Whether usage is unlimited (for feature entitlements)
	HasUnlimitedUsage bool `json:"hasUnlimitedUsage" api:"required"`
	// Widget types where this entitlement is hidden
	//
	// Any of "PAYWALL", "CUSTOMER_PORTAL", "CHECKOUT".
	HiddenFromWidgets []string `json:"hiddenFromWidgets" api:"required"`
	// Whether this is a custom entitlement
	IsCustom bool `json:"isCustom" api:"required"`
	// Whether the entitlement is granted
	IsGranted bool `json:"isGranted" api:"required"`
	// Display order of the entitlement
	Order float64 `json:"order" api:"required"`
	// Usage reset period (for feature entitlements)
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod" api:"required"`
	// Reset period configuration (for feature entitlements)
	ResetPeriodConfiguration V1EventPlanEntitlementListResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
	// Entitlement type (FEATURE or CREDIT)
	//
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Usage limit (for feature entitlements)
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Amount                   respjson.Field
		Behavior                 respjson.Field
		Cadence                  respjson.Field
		CreatedAt                respjson.Field
		CustomCurrencyID         respjson.Field
		Description              respjson.Field
		DisplayNameOverride      respjson.Field
		EnumValues               respjson.Field
		FeatureID                respjson.Field
		HasSoftLimit             respjson.Field
		HasUnlimitedUsage        respjson.Field
		HiddenFromWidgets        respjson.Field
		IsCustom                 respjson.Field
		IsGranted                respjson.Field
		Order                    respjson.Field
		ResetPeriod              respjson.Field
		ResetPeriodConfiguration respjson.Field
		Type                     respjson.Field
		UpdatedAt                respjson.Field
		UsageLimit               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanEntitlementListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanEntitlementListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1EventPlanEntitlementListResponseDataResetPeriodConfigurationUnion contains all
// possible properties and values from
// [V1EventPlanEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1EventPlanEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1EventPlanEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1EventPlanEntitlementListResponseDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1EventPlanEntitlementListResponseDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1EventPlanEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1EventPlanEntitlementListResponseDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1EventPlanEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1EventPlanEntitlementListResponseDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1EventPlanEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1EventPlanEntitlementListResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1EventPlanEntitlementListResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1EventPlanEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventPlanEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1EventPlanEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventPlanEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1EventPlanEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventPlanEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata including cursors for navigating through results
type V1EventPlanEntitlementListResponsePagination struct {
	// Cursor for fetching the next page of results, or null if no additional pages
	// exist
	Next string `json:"next" api:"required" format:"uuid"`
	// Cursor for fetching the previous page of results, or null if at the beginning
	Prev string `json:"prev" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		Prev        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanEntitlementListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanEntitlementListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventPlanEntitlementNewParams struct {
	// Entitlements to create
	Entitlements []V1EventPlanEntitlementNewParamsEntitlement `json:"entitlements,omitzero" api:"required"`
	paramObj
}

func (r V1EventPlanEntitlementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entitlement to create. Provide exactly one of feature or credit.
type V1EventPlanEntitlementNewParamsEntitlement struct {
	// Credit entitlement to create
	Credit V1EventPlanEntitlementNewParamsEntitlementCredit `json:"credit,omitzero"`
	// Feature entitlement to create
	Feature V1EventPlanEntitlementNewParamsEntitlementFeature `json:"feature,omitzero"`
	paramObj
}

func (r V1EventPlanEntitlementNewParamsEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementNewParamsEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementNewParamsEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement to create
//
// The properties Amount, Cadence, CustomCurrencyID are required.
type V1EventPlanEntitlementNewParamsEntitlementCredit struct {
	// Credit grant amount
	Amount param.Opt[float64] `json:"amount,omitzero" api:"required"`
	// Credit grant cadence (MONTH or YEAR)
	//
	// Any of "MONTH", "YEAR".
	Cadence string `json:"cadence,omitzero" api:"required"`
	// The custom currency ID for the credit entitlement
	CustomCurrencyID string `json:"customCurrencyId" api:"required"`
	// Description of the entitlement
	Description param.Opt[string] `json:"description,omitzero"`
	// Override display name for the entitlement
	DisplayNameOverride param.Opt[string] `json:"displayNameOverride,omitzero"`
	// Whether this is a custom entitlement
	IsCustom param.Opt[bool] `json:"isCustom,omitzero"`
	// Whether the entitlement is granted
	IsGranted param.Opt[bool] `json:"isGranted,omitzero"`
	// Display order of the entitlement
	Order param.Opt[float64] `json:"order,omitzero"`
	// Entitlement behavior (Increment or Override)
	//
	// Any of "Increment", "Override".
	Behavior string `json:"behavior,omitzero"`
	// Widget types where this entitlement is hidden
	//
	// Any of "PAYWALL", "CUSTOMER_PORTAL", "CHECKOUT".
	HiddenFromWidgets []string `json:"hiddenFromWidgets,omitzero"`
	paramObj
}

func (r V1EventPlanEntitlementNewParamsEntitlementCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementNewParamsEntitlementCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementNewParamsEntitlementCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventPlanEntitlementNewParamsEntitlementCredit](
		"cadence", "MONTH", "YEAR",
	)
	apijson.RegisterFieldValidator[V1EventPlanEntitlementNewParamsEntitlementCredit](
		"behavior", "Increment", "Override",
	)
}

// Feature entitlement to create
//
// The property FeatureID is required.
type V1EventPlanEntitlementNewParamsEntitlementFeature struct {
	// The feature ID to attach the entitlement to
	FeatureID string `json:"featureId" api:"required"`
	// Maximum allowed usage for the feature
	UsageLimit param.Opt[int64] `json:"usageLimit,omitzero"`
	// Description of the entitlement
	Description param.Opt[string] `json:"description,omitzero"`
	// Override display name for the entitlement
	DisplayNameOverride param.Opt[string] `json:"displayNameOverride,omitzero"`
	// Whether the usage limit is a soft limit
	HasSoftLimit param.Opt[bool] `json:"hasSoftLimit,omitzero"`
	// Whether usage is unlimited
	HasUnlimitedUsage param.Opt[bool] `json:"hasUnlimitedUsage,omitzero"`
	// Whether this is a custom entitlement
	IsCustom param.Opt[bool] `json:"isCustom,omitzero"`
	// Whether the entitlement is granted
	IsGranted param.Opt[bool] `json:"isGranted,omitzero"`
	// Display order of the entitlement
	Order param.Opt[float64] `json:"order,omitzero"`
	// Configuration for monthly reset period
	MonthlyResetPeriodConfiguration V1EventPlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Configuration for weekly reset period
	WeeklyResetPeriodConfiguration V1EventPlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Configuration for yearly reset period
	YearlyResetPeriodConfiguration V1EventPlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
	// Entitlement behavior (Increment or Override)
	//
	// Any of "Increment", "Override".
	Behavior string `json:"behavior,omitzero"`
	// Allowed enum values for the feature entitlement
	EnumValues []string `json:"enumValues,omitzero"`
	// Widget types where this entitlement is hidden
	//
	// Any of "PAYWALL", "CUSTOMER_PORTAL", "CHECKOUT".
	HiddenFromWidgets []string `json:"hiddenFromWidgets,omitzero"`
	// Period at which usage resets
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,omitzero"`
	paramObj
}

func (r V1EventPlanEntitlementNewParamsEntitlementFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementNewParamsEntitlementFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementNewParamsEntitlementFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventPlanEntitlementNewParamsEntitlementFeature](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1EventPlanEntitlementNewParamsEntitlementFeature](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Configuration for monthly reset period
//
// The property AccordingTo is required.
type V1EventPlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventPlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventPlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Configuration for weekly reset period
//
// The property AccordingTo is required.
type V1EventPlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventPlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventPlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Configuration for yearly reset period
//
// The property AccordingTo is required.
type V1EventPlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventPlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventPlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1EventPlanEntitlementUpdateParams struct {
	PlanID string `path:"planId" api:"required" json:"-"`
	// Credit entitlement fields to update
	Credit V1EventPlanEntitlementUpdateParamsCredit `json:"credit,omitzero"`
	// Feature entitlement fields to update
	Feature V1EventPlanEntitlementUpdateParamsFeature `json:"feature,omitzero"`
	paramObj
}

func (r V1EventPlanEntitlementUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement fields to update
type V1EventPlanEntitlementUpdateParamsCredit struct {
	// Credit grant amount
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Description of the entitlement
	Description param.Opt[string] `json:"description,omitzero"`
	// Override display name for the entitlement
	DisplayNameOverride param.Opt[string] `json:"displayNameOverride,omitzero"`
	// Whether this is a custom entitlement
	IsCustom param.Opt[bool] `json:"isCustom,omitzero"`
	// Whether the entitlement is granted
	IsGranted param.Opt[bool] `json:"isGranted,omitzero"`
	// Display order of the entitlement
	Order param.Opt[float64] `json:"order,omitzero"`
	// Entitlement behavior (Increment or Override)
	//
	// Any of "Increment", "Override".
	Behavior string `json:"behavior,omitzero"`
	// Credit grant cadence (MONTH or YEAR)
	//
	// Any of "MONTH", "YEAR".
	Cadence string `json:"cadence,omitzero"`
	// Widget types where this entitlement is hidden
	//
	// Any of "PAYWALL", "CUSTOMER_PORTAL", "CHECKOUT".
	HiddenFromWidgets []string `json:"hiddenFromWidgets,omitzero"`
	paramObj
}

func (r V1EventPlanEntitlementUpdateParamsCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementUpdateParamsCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementUpdateParamsCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventPlanEntitlementUpdateParamsCredit](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1EventPlanEntitlementUpdateParamsCredit](
		"cadence", "MONTH", "YEAR",
	)
}

// Feature entitlement fields to update
type V1EventPlanEntitlementUpdateParamsFeature struct {
	// Maximum allowed usage for the feature
	UsageLimit param.Opt[int64] `json:"usageLimit,omitzero"`
	// Description of the entitlement
	Description param.Opt[string] `json:"description,omitzero"`
	// Override display name for the entitlement
	DisplayNameOverride param.Opt[string] `json:"displayNameOverride,omitzero"`
	// Whether the usage limit is a soft limit
	HasSoftLimit param.Opt[bool] `json:"hasSoftLimit,omitzero"`
	// Whether usage is unlimited
	HasUnlimitedUsage param.Opt[bool] `json:"hasUnlimitedUsage,omitzero"`
	// Whether this is a custom entitlement
	IsCustom param.Opt[bool] `json:"isCustom,omitzero"`
	// Whether the entitlement is granted
	IsGranted param.Opt[bool] `json:"isGranted,omitzero"`
	// Display order of the entitlement
	Order param.Opt[float64] `json:"order,omitzero"`
	// Configuration for monthly reset period
	MonthlyResetPeriodConfiguration V1EventPlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Configuration for weekly reset period
	WeeklyResetPeriodConfiguration V1EventPlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Configuration for yearly reset period
	YearlyResetPeriodConfiguration V1EventPlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
	// Entitlement behavior (Increment or Override)
	//
	// Any of "Increment", "Override".
	Behavior string `json:"behavior,omitzero"`
	// Allowed enum values for the feature entitlement
	EnumValues []string `json:"enumValues,omitzero"`
	// Widget types where this entitlement is hidden
	//
	// Any of "PAYWALL", "CUSTOMER_PORTAL", "CHECKOUT".
	HiddenFromWidgets []string `json:"hiddenFromWidgets,omitzero"`
	// Period at which usage resets
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,omitzero"`
	paramObj
}

func (r V1EventPlanEntitlementUpdateParamsFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementUpdateParamsFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementUpdateParamsFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventPlanEntitlementUpdateParamsFeature](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1EventPlanEntitlementUpdateParamsFeature](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Configuration for monthly reset period
//
// The property AccordingTo is required.
type V1EventPlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventPlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventPlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Configuration for weekly reset period
//
// The property AccordingTo is required.
type V1EventPlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventPlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventPlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Configuration for yearly reset period
//
// The property AccordingTo is required.
type V1EventPlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventPlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventPlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1EventPlanEntitlementDeleteParams struct {
	PlanID string `path:"planId" api:"required" json:"-"`
	paramObj
}
