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

// V1AddonEntitlementService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AddonEntitlementService] method instead.
type V1AddonEntitlementService struct {
	Options []option.RequestOption
}

// NewV1AddonEntitlementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1AddonEntitlementService(opts ...option.RequestOption) (r V1AddonEntitlementService) {
	r = V1AddonEntitlementService{}
	r.Options = opts
	return
}

// Creates one or more entitlements (feature or credit) on a draft addon.
func (r *V1AddonEntitlementService) New(ctx context.Context, addonID string, body V1AddonEntitlementNewParams, opts ...option.RequestOption) (res *V1AddonEntitlementNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if addonID == "" {
		err = errors.New("missing required addonId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s/entitlements", addonID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing entitlement on a draft addon.
func (r *V1AddonEntitlementService) Update(ctx context.Context, id string, params V1AddonEntitlementUpdateParams, opts ...option.RequestOption) (res *AddonPackageEntitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AddonID == "" {
		err = errors.New("missing required addonId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s/entitlements/%s", params.AddonID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieves a list of entitlements for an addon.
func (r *V1AddonEntitlementService) List(ctx context.Context, addonID string, opts ...option.RequestOption) (res *V1AddonEntitlementListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if addonID == "" {
		err = errors.New("missing required addonId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s/entitlements", addonID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an entitlement from a draft addon.
func (r *V1AddonEntitlementService) Delete(ctx context.Context, id string, body V1AddonEntitlementDeleteParams, opts ...option.RequestOption) (res *AddonPackageEntitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AddonID == "" {
		err = errors.New("missing required addonId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s/entitlements/%s", body.AddonID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Response object
type AddonPackageEntitlement struct {
	// Feature or credit entitlement on an addon
	Data AddonPackageEntitlementData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddonPackageEntitlement) RawJSON() string { return r.JSON.raw }
func (r *AddonPackageEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature or credit entitlement on an addon
type AddonPackageEntitlementData struct {
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
	ResetPeriodConfiguration AddonPackageEntitlementDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
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
func (r AddonPackageEntitlementData) RawJSON() string { return r.JSON.raw }
func (r *AddonPackageEntitlementData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AddonPackageEntitlementDataResetPeriodConfigurationUnion contains all possible
// properties and values from
// [AddonPackageEntitlementDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [AddonPackageEntitlementDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [AddonPackageEntitlementDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AddonPackageEntitlementDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u AddonPackageEntitlementDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v AddonPackageEntitlementDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AddonPackageEntitlementDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v AddonPackageEntitlementDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AddonPackageEntitlementDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v AddonPackageEntitlementDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AddonPackageEntitlementDataResetPeriodConfigurationUnion) RawJSON() string { return u.JSON.raw }

func (r *AddonPackageEntitlementDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type AddonPackageEntitlementDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r AddonPackageEntitlementDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AddonPackageEntitlementDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type AddonPackageEntitlementDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r AddonPackageEntitlementDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AddonPackageEntitlementDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type AddonPackageEntitlementDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r AddonPackageEntitlementDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AddonPackageEntitlementDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1AddonEntitlementNewResponse struct {
	Data []V1AddonEntitlementNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AddonEntitlementNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AddonEntitlementNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature or credit entitlement on an addon
type V1AddonEntitlementNewResponseData struct {
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
	ResetPeriodConfiguration V1AddonEntitlementNewResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
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
func (r V1AddonEntitlementNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1AddonEntitlementNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1AddonEntitlementNewResponseDataResetPeriodConfigurationUnion contains all
// possible properties and values from
// [V1AddonEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1AddonEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1AddonEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1AddonEntitlementNewResponseDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1AddonEntitlementNewResponseDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1AddonEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1AddonEntitlementNewResponseDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1AddonEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1AddonEntitlementNewResponseDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1AddonEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1AddonEntitlementNewResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1AddonEntitlementNewResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1AddonEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1AddonEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1AddonEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1AddonEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1AddonEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1AddonEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response list object
type V1AddonEntitlementListResponse struct {
	Data []V1AddonEntitlementListResponseData `json:"data" api:"required"`
	// Pagination metadata including cursors for navigating through results
	Pagination V1AddonEntitlementListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AddonEntitlementListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AddonEntitlementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature or credit entitlement on an addon
type V1AddonEntitlementListResponseData struct {
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
	ResetPeriodConfiguration V1AddonEntitlementListResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
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
func (r V1AddonEntitlementListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1AddonEntitlementListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1AddonEntitlementListResponseDataResetPeriodConfigurationUnion contains all
// possible properties and values from
// [V1AddonEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1AddonEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1AddonEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1AddonEntitlementListResponseDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1AddonEntitlementListResponseDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1AddonEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1AddonEntitlementListResponseDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1AddonEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1AddonEntitlementListResponseDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1AddonEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1AddonEntitlementListResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1AddonEntitlementListResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1AddonEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1AddonEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1AddonEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1AddonEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1AddonEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1AddonEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata including cursors for navigating through results
type V1AddonEntitlementListResponsePagination struct {
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
func (r V1AddonEntitlementListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *V1AddonEntitlementListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AddonEntitlementNewParams struct {
	// Entitlements to create
	Entitlements []V1AddonEntitlementNewParamsEntitlement `json:"entitlements,omitzero" api:"required"`
	paramObj
}

func (r V1AddonEntitlementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entitlement to create. Provide exactly one of feature or credit.
type V1AddonEntitlementNewParamsEntitlement struct {
	// Credit entitlement to create
	Credit V1AddonEntitlementNewParamsEntitlementCredit `json:"credit,omitzero"`
	// Feature entitlement to create
	Feature V1AddonEntitlementNewParamsEntitlementFeature `json:"feature,omitzero"`
	paramObj
}

func (r V1AddonEntitlementNewParamsEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementNewParamsEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementNewParamsEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement to create
//
// The properties Amount, Cadence, CustomCurrencyID are required.
type V1AddonEntitlementNewParamsEntitlementCredit struct {
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

func (r V1AddonEntitlementNewParamsEntitlementCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementNewParamsEntitlementCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementNewParamsEntitlementCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementNewParamsEntitlementCredit](
		"cadence", "MONTH", "YEAR",
	)
	apijson.RegisterFieldValidator[V1AddonEntitlementNewParamsEntitlementCredit](
		"behavior", "Increment", "Override",
	)
}

// Feature entitlement to create
//
// The property FeatureID is required.
type V1AddonEntitlementNewParamsEntitlementFeature struct {
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
	MonthlyResetPeriodConfiguration V1AddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Configuration for weekly reset period
	WeeklyResetPeriodConfiguration V1AddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Configuration for yearly reset period
	YearlyResetPeriodConfiguration V1AddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
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

func (r V1AddonEntitlementNewParamsEntitlementFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementNewParamsEntitlementFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementNewParamsEntitlementFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementNewParamsEntitlementFeature](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1AddonEntitlementNewParamsEntitlementFeature](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Configuration for monthly reset period
//
// The property AccordingTo is required.
type V1AddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1AddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Configuration for weekly reset period
//
// The property AccordingTo is required.
type V1AddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1AddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Configuration for yearly reset period
//
// The property AccordingTo is required.
type V1AddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1AddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1AddonEntitlementUpdateParams struct {
	AddonID string `path:"addonId" api:"required" json:"-"`
	// Credit entitlement fields to update
	Credit V1AddonEntitlementUpdateParamsCredit `json:"credit,omitzero"`
	// Feature entitlement fields to update
	Feature V1AddonEntitlementUpdateParamsFeature `json:"feature,omitzero"`
	paramObj
}

func (r V1AddonEntitlementUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement fields to update
type V1AddonEntitlementUpdateParamsCredit struct {
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

func (r V1AddonEntitlementUpdateParamsCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementUpdateParamsCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementUpdateParamsCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsCredit](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsCredit](
		"cadence", "MONTH", "YEAR",
	)
}

// Feature entitlement fields to update
type V1AddonEntitlementUpdateParamsFeature struct {
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
	MonthlyResetPeriodConfiguration V1AddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Configuration for weekly reset period
	WeeklyResetPeriodConfiguration V1AddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Configuration for yearly reset period
	YearlyResetPeriodConfiguration V1AddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
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

func (r V1AddonEntitlementUpdateParamsFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementUpdateParamsFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementUpdateParamsFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsFeature](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsFeature](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Configuration for monthly reset period
//
// The property AccordingTo is required.
type V1AddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1AddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Configuration for weekly reset period
//
// The property AccordingTo is required.
type V1AddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1AddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Configuration for yearly reset period
//
// The property AccordingTo is required.
type V1AddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1AddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1AddonEntitlementDeleteParams struct {
	AddonID string `path:"addonId" api:"required" json:"-"`
	paramObj
}
