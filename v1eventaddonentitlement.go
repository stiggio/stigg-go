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

// V1EventAddonEntitlementService contains methods and other services that help
// with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventAddonEntitlementService] method instead.
type V1EventAddonEntitlementService struct {
	Options []option.RequestOption
}

// NewV1EventAddonEntitlementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1EventAddonEntitlementService(opts ...option.RequestOption) (r V1EventAddonEntitlementService) {
	r = V1EventAddonEntitlementService{}
	r.Options = opts
	return
}

// Creates one or more entitlements (feature or credit) on a draft addon.
func (r *V1EventAddonEntitlementService) New(ctx context.Context, addonID string, body V1EventAddonEntitlementNewParams, opts ...option.RequestOption) (res *V1EventAddonEntitlementNewResponse, err error) {
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
func (r *V1EventAddonEntitlementService) Update(ctx context.Context, id string, params V1EventAddonEntitlementUpdateParams, opts ...option.RequestOption) (res *AddonPackageEntitlement, err error) {
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
func (r *V1EventAddonEntitlementService) List(ctx context.Context, addonID string, opts ...option.RequestOption) (res *V1EventAddonEntitlementListResponse, err error) {
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
func (r *V1EventAddonEntitlementService) Delete(ctx context.Context, id string, body V1EventAddonEntitlementDeleteParams, opts ...option.RequestOption) (res *AddonPackageEntitlement, err error) {
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
type V1EventAddonEntitlementNewResponse struct {
	Data []V1EventAddonEntitlementNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonEntitlementNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonEntitlementNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature or credit entitlement on an addon
type V1EventAddonEntitlementNewResponseData struct {
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
	ResetPeriodConfiguration V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
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
func (r V1EventAddonEntitlementNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonEntitlementNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationUnion contains all
// possible properties and values from
// [V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventAddonEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response list object
type V1EventAddonEntitlementListResponse struct {
	Data []V1EventAddonEntitlementListResponseData `json:"data" api:"required"`
	// Pagination metadata including cursors for navigating through results
	Pagination V1EventAddonEntitlementListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonEntitlementListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonEntitlementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature or credit entitlement on an addon
type V1EventAddonEntitlementListResponseData struct {
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
	ResetPeriodConfiguration V1EventAddonEntitlementListResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
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
func (r V1EventAddonEntitlementListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonEntitlementListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1EventAddonEntitlementListResponseDataResetPeriodConfigurationUnion contains
// all possible properties and values from
// [V1EventAddonEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1EventAddonEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1EventAddonEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1EventAddonEntitlementListResponseDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1EventAddonEntitlementListResponseDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1EventAddonEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1EventAddonEntitlementListResponseDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1EventAddonEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1EventAddonEntitlementListResponseDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1EventAddonEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1EventAddonEntitlementListResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1EventAddonEntitlementListResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1EventAddonEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1EventAddonEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventAddonEntitlementListResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1EventAddonEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1EventAddonEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventAddonEntitlementListResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1EventAddonEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1EventAddonEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventAddonEntitlementListResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata including cursors for navigating through results
type V1EventAddonEntitlementListResponsePagination struct {
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
func (r V1EventAddonEntitlementListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonEntitlementListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventAddonEntitlementNewParams struct {
	// Entitlements to create
	Entitlements []V1EventAddonEntitlementNewParamsEntitlement `json:"entitlements,omitzero" api:"required"`
	paramObj
}

func (r V1EventAddonEntitlementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single entitlement to create. Provide exactly one of feature or credit.
type V1EventAddonEntitlementNewParamsEntitlement struct {
	// Credit entitlement to create
	Credit V1EventAddonEntitlementNewParamsEntitlementCredit `json:"credit,omitzero"`
	// Feature entitlement to create
	Feature V1EventAddonEntitlementNewParamsEntitlementFeature `json:"feature,omitzero"`
	paramObj
}

func (r V1EventAddonEntitlementNewParamsEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementNewParamsEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementNewParamsEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement to create
//
// The properties Amount, Cadence, CustomCurrencyID are required.
type V1EventAddonEntitlementNewParamsEntitlementCredit struct {
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

func (r V1EventAddonEntitlementNewParamsEntitlementCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementNewParamsEntitlementCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementNewParamsEntitlementCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventAddonEntitlementNewParamsEntitlementCredit](
		"cadence", "MONTH", "YEAR",
	)
	apijson.RegisterFieldValidator[V1EventAddonEntitlementNewParamsEntitlementCredit](
		"behavior", "Increment", "Override",
	)
}

// Feature entitlement to create
//
// The property FeatureID is required.
type V1EventAddonEntitlementNewParamsEntitlementFeature struct {
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
	MonthlyResetPeriodConfiguration V1EventAddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Configuration for weekly reset period
	WeeklyResetPeriodConfiguration V1EventAddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Configuration for yearly reset period
	YearlyResetPeriodConfiguration V1EventAddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
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

func (r V1EventAddonEntitlementNewParamsEntitlementFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementNewParamsEntitlementFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementNewParamsEntitlementFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventAddonEntitlementNewParamsEntitlementFeature](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1EventAddonEntitlementNewParamsEntitlementFeature](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Configuration for monthly reset period
//
// The property AccordingTo is required.
type V1EventAddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventAddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventAddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Configuration for weekly reset period
//
// The property AccordingTo is required.
type V1EventAddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventAddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventAddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Configuration for yearly reset period
//
// The property AccordingTo is required.
type V1EventAddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventAddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventAddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1EventAddonEntitlementUpdateParams struct {
	AddonID string `path:"addonId" api:"required" json:"-"`
	// Credit entitlement fields to update
	Credit V1EventAddonEntitlementUpdateParamsCredit `json:"credit,omitzero"`
	// Feature entitlement fields to update
	Feature V1EventAddonEntitlementUpdateParamsFeature `json:"feature,omitzero"`
	paramObj
}

func (r V1EventAddonEntitlementUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement fields to update
type V1EventAddonEntitlementUpdateParamsCredit struct {
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

func (r V1EventAddonEntitlementUpdateParamsCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementUpdateParamsCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementUpdateParamsCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventAddonEntitlementUpdateParamsCredit](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1EventAddonEntitlementUpdateParamsCredit](
		"cadence", "MONTH", "YEAR",
	)
}

// Feature entitlement fields to update
type V1EventAddonEntitlementUpdateParamsFeature struct {
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
	MonthlyResetPeriodConfiguration V1EventAddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Configuration for weekly reset period
	WeeklyResetPeriodConfiguration V1EventAddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Configuration for yearly reset period
	YearlyResetPeriodConfiguration V1EventAddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
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

func (r V1EventAddonEntitlementUpdateParamsFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementUpdateParamsFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementUpdateParamsFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventAddonEntitlementUpdateParamsFeature](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1EventAddonEntitlementUpdateParamsFeature](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Configuration for monthly reset period
//
// The property AccordingTo is required.
type V1EventAddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventAddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventAddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Configuration for weekly reset period
//
// The property AccordingTo is required.
type V1EventAddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventAddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventAddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Configuration for yearly reset period
//
// The property AccordingTo is required.
type V1EventAddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1EventAddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventAddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1EventAddonEntitlementDeleteParams struct {
	AddonID string `path:"addonId" api:"required" json:"-"`
	paramObj
}
