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

// V1PlanEntitlementService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PlanEntitlementService] method instead.
type V1PlanEntitlementService struct {
	Options []option.RequestOption
}

// NewV1PlanEntitlementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1PlanEntitlementService(opts ...option.RequestOption) (r V1PlanEntitlementService) {
	r = V1PlanEntitlementService{}
	r.Options = opts
	return
}

// Creates one or more entitlements (feature or credit) on a draft plan.
func (r *V1PlanEntitlementService) New(ctx context.Context, planID string, body V1PlanEntitlementNewParams, opts ...option.RequestOption) (res *V1PlanEntitlementNewResponse, err error) {
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
func (r *V1PlanEntitlementService) Update(ctx context.Context, id string, params V1PlanEntitlementUpdateParams, opts ...option.RequestOption) (res *PlanEntitlement, err error) {
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
func (r *V1PlanEntitlementService) List(ctx context.Context, planID string, opts ...option.RequestOption) (res *V1PlanEntitlementListResponse, err error) {
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
func (r *V1PlanEntitlementService) Delete(ctx context.Context, id string, body V1PlanEntitlementDeleteParams, opts ...option.RequestOption) (res *PlanEntitlement, err error) {
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
type V1PlanEntitlementNewResponse struct {
	Data []V1PlanEntitlementNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanEntitlementNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PlanEntitlementNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature or credit entitlement on a plan
type V1PlanEntitlementNewResponseData struct {
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
	ResetPeriodConfiguration V1PlanEntitlementNewResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
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
func (r V1PlanEntitlementNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1PlanEntitlementNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1PlanEntitlementNewResponseDataResetPeriodConfigurationUnion contains all
// possible properties and values from
// [V1PlanEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1PlanEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1PlanEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1PlanEntitlementNewResponseDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1PlanEntitlementNewResponseDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1PlanEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1PlanEntitlementNewResponseDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1PlanEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1PlanEntitlementNewResponseDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1PlanEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1PlanEntitlementNewResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1PlanEntitlementNewResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1PlanEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1PlanEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1PlanEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1PlanEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1PlanEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1PlanEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response list object
type V1PlanEntitlementListResponse struct {
	Data []V1PlanEntitlementListResponseData `json:"data" api:"required"`
	// Pagination metadata including cursors for navigating through results
	Pagination V1PlanEntitlementListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanEntitlementListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PlanEntitlementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature or credit entitlement on a plan
type V1PlanEntitlementListResponseData struct {
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
	ResetPeriodConfiguration V1PlanEntitlementListResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
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
func (r V1PlanEntitlementListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1PlanEntitlementListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1PlanEntitlementListResponseDataResetPeriodConfigurationUnion contains all
// possible properties and values from
// [V1PlanEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1PlanEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1PlanEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1PlanEntitlementListResponseDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1PlanEntitlementListResponseDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1PlanEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1PlanEntitlementListResponseDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1PlanEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1PlanEntitlementListResponseDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1PlanEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1PlanEntitlementListResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1PlanEntitlementListResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1PlanEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1PlanEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1PlanEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1PlanEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1PlanEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1PlanEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata including cursors for navigating through results
type V1PlanEntitlementListResponsePagination struct {
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
func (r V1PlanEntitlementListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *V1PlanEntitlementListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanEntitlementNewParams struct {
	// Entitlements to create
	Entitlements []V1PlanEntitlementNewParamsEntitlement `json:"entitlements,omitzero" api:"required"`
	paramObj
}

func (r V1PlanEntitlementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entitlement to create. Provide exactly one of feature or credit.
type V1PlanEntitlementNewParamsEntitlement struct {
	// Credit entitlement to create
	Credit V1PlanEntitlementNewParamsEntitlementCredit `json:"credit,omitzero"`
	// Feature entitlement to create
	Feature V1PlanEntitlementNewParamsEntitlementFeature `json:"feature,omitzero"`
	paramObj
}

func (r V1PlanEntitlementNewParamsEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementNewParamsEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementNewParamsEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement to create
//
// The properties Amount, Cadence, CustomCurrencyID are required.
type V1PlanEntitlementNewParamsEntitlementCredit struct {
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

func (r V1PlanEntitlementNewParamsEntitlementCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementNewParamsEntitlementCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementNewParamsEntitlementCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementNewParamsEntitlementCredit](
		"cadence", "MONTH", "YEAR",
	)
	apijson.RegisterFieldValidator[V1PlanEntitlementNewParamsEntitlementCredit](
		"behavior", "Increment", "Override",
	)
}

// Feature entitlement to create
//
// The property FeatureID is required.
type V1PlanEntitlementNewParamsEntitlementFeature struct {
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
	MonthlyResetPeriodConfiguration V1PlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Configuration for weekly reset period
	WeeklyResetPeriodConfiguration V1PlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Configuration for yearly reset period
	YearlyResetPeriodConfiguration V1PlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
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

func (r V1PlanEntitlementNewParamsEntitlementFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementNewParamsEntitlementFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementNewParamsEntitlementFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementNewParamsEntitlementFeature](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1PlanEntitlementNewParamsEntitlementFeature](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Configuration for monthly reset period
//
// The property AccordingTo is required.
type V1PlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Configuration for weekly reset period
//
// The property AccordingTo is required.
type V1PlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Configuration for yearly reset period
//
// The property AccordingTo is required.
type V1PlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1PlanEntitlementUpdateParams struct {
	PlanID string `path:"planId" api:"required" json:"-"`
	// Credit entitlement fields to update
	Credit V1PlanEntitlementUpdateParamsCredit `json:"credit,omitzero"`
	// Feature entitlement fields to update
	Feature V1PlanEntitlementUpdateParamsFeature `json:"feature,omitzero"`
	paramObj
}

func (r V1PlanEntitlementUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement fields to update
type V1PlanEntitlementUpdateParamsCredit struct {
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

func (r V1PlanEntitlementUpdateParamsCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementUpdateParamsCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementUpdateParamsCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsCredit](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsCredit](
		"cadence", "MONTH", "YEAR",
	)
}

// Feature entitlement fields to update
type V1PlanEntitlementUpdateParamsFeature struct {
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
	MonthlyResetPeriodConfiguration V1PlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Configuration for weekly reset period
	WeeklyResetPeriodConfiguration V1PlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Configuration for yearly reset period
	YearlyResetPeriodConfiguration V1PlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
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

func (r V1PlanEntitlementUpdateParamsFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementUpdateParamsFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementUpdateParamsFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsFeature](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsFeature](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Configuration for monthly reset period
//
// The property AccordingTo is required.
type V1PlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Configuration for weekly reset period
//
// The property AccordingTo is required.
type V1PlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Configuration for yearly reset period
//
// The property AccordingTo is required.
type V1PlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1PlanEntitlementDeleteParams struct {
	PlanID string `path:"planId" api:"required" json:"-"`
	paramObj
}
