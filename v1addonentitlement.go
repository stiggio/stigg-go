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
	"github.com/stiggio/stigg-go/shared/constant"
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
		return nil, err
	}
	path := fmt.Sprintf("api/v1/addons/%s/entitlements", addonID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates an existing entitlement on a draft addon.
func (r *V1AddonEntitlementService) Update(ctx context.Context, id string, params V1AddonEntitlementUpdateParams, opts ...option.RequestOption) (res *AddonPackageEntitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AddonID == "" {
		err = errors.New("missing required addonId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/addons/%s/entitlements/%s", params.AddonID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieves a list of entitlements for an addon.
func (r *V1AddonEntitlementService) List(ctx context.Context, addonID string, opts ...option.RequestOption) (res *V1AddonEntitlementListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if addonID == "" {
		err = errors.New("missing required addonId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/addons/%s/entitlements", addonID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes an entitlement from a draft addon.
func (r *V1AddonEntitlementService) Delete(ctx context.Context, id string, body V1AddonEntitlementDeleteParams, opts ...option.RequestOption) (res *AddonPackageEntitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AddonID == "" {
		err = errors.New("missing required addonId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/addons/%s/entitlements/%s", body.AddonID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Response object
type AddonPackageEntitlement struct {
	// Feature entitlement response
	Data AddonPackageEntitlementDataUnion `json:"data" api:"required"`
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

// AddonPackageEntitlementDataUnion contains all possible properties and values
// from [AddonPackageEntitlementDataFeature], [AddonPackageEntitlementDataCredit].
//
// Use the [AddonPackageEntitlementDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AddonPackageEntitlementDataUnion struct {
	ID                  string    `json:"id"`
	Behavior            string    `json:"behavior"`
	CreatedAt           time.Time `json:"createdAt"`
	Description         string    `json:"description"`
	DisplayNameOverride string    `json:"displayNameOverride"`
	// This field is from variant [AddonPackageEntitlementDataFeature].
	EnumValues []string `json:"enumValues"`
	// This field is from variant [AddonPackageEntitlementDataFeature].
	HasSoftLimit bool `json:"hasSoftLimit"`
	// This field is from variant [AddonPackageEntitlementDataFeature].
	HasUnlimitedUsage bool     `json:"hasUnlimitedUsage"`
	HiddenFromWidgets []string `json:"hiddenFromWidgets"`
	IsCustom          bool     `json:"isCustom"`
	IsGranted         bool     `json:"isGranted"`
	Order             float64  `json:"order"`
	// This field is from variant [AddonPackageEntitlementDataFeature].
	ResetPeriod string `json:"resetPeriod"`
	// This field is from variant [AddonPackageEntitlementDataFeature].
	ResetPeriodConfiguration AddonPackageEntitlementDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration"`
	// Any of "FEATURE", "CREDIT".
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updatedAt"`
	// This field is from variant [AddonPackageEntitlementDataFeature].
	UsageLimit float64 `json:"usageLimit"`
	// This field is from variant [AddonPackageEntitlementDataCredit].
	Amount float64 `json:"amount"`
	// This field is from variant [AddonPackageEntitlementDataCredit].
	Cadence string `json:"cadence"`
	// This field is from variant [AddonPackageEntitlementDataCredit].
	DependencyFeatureID string `json:"dependencyFeatureId"`
	JSON                struct {
		ID                       respjson.Field
		Behavior                 respjson.Field
		CreatedAt                respjson.Field
		Description              respjson.Field
		DisplayNameOverride      respjson.Field
		EnumValues               respjson.Field
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
		Amount                   respjson.Field
		Cadence                  respjson.Field
		DependencyFeatureID      respjson.Field
		raw                      string
	} `json:"-"`
}

// anyAddonPackageEntitlementData is implemented by each variant of
// [AddonPackageEntitlementDataUnion] to add type safety for the return type of
// [AddonPackageEntitlementDataUnion.AsAny]
type anyAddonPackageEntitlementData interface {
	implAddonPackageEntitlementDataUnion()
}

func (AddonPackageEntitlementDataFeature) implAddonPackageEntitlementDataUnion() {}
func (AddonPackageEntitlementDataCredit) implAddonPackageEntitlementDataUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AddonPackageEntitlementDataUnion.AsAny().(type) {
//	case stigg.AddonPackageEntitlementDataFeature:
//	case stigg.AddonPackageEntitlementDataCredit:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AddonPackageEntitlementDataUnion) AsAny() anyAddonPackageEntitlementData {
	switch u.Type {
	case "FEATURE":
		return u.AsFeature()
	case "CREDIT":
		return u.AsCredit()
	}
	return nil
}

func (u AddonPackageEntitlementDataUnion) AsFeature() (v AddonPackageEntitlementDataFeature) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AddonPackageEntitlementDataUnion) AsCredit() (v AddonPackageEntitlementDataCredit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AddonPackageEntitlementDataUnion) RawJSON() string { return u.JSON.raw }

func (r *AddonPackageEntitlementDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature entitlement response
type AddonPackageEntitlementDataFeature struct {
	// Unique identifier of the entitlement
	ID string `json:"id" api:"required"`
	// Entitlement behavior (Increment or Override)
	//
	// Any of "Increment", "Override".
	Behavior string `json:"behavior" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Optional description of the entitlement
	Description string `json:"description" api:"required"`
	// Override display name for the entitlement
	DisplayNameOverride string `json:"displayNameOverride" api:"required"`
	// Allowed enum values (for feature entitlements)
	EnumValues []string `json:"enumValues" api:"required"`
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
	ResetPeriodConfiguration AddonPackageEntitlementDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
	// Entitlement type (FEATURE or CREDIT)
	Type constant.Feature `json:"type" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Usage limit (for feature entitlements)
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Behavior                 respjson.Field
		CreatedAt                respjson.Field
		Description              respjson.Field
		DisplayNameOverride      respjson.Field
		EnumValues               respjson.Field
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
func (r AddonPackageEntitlementDataFeature) RawJSON() string { return r.JSON.raw }
func (r *AddonPackageEntitlementDataFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AddonPackageEntitlementDataFeatureResetPeriodConfigurationUnion contains all
// possible properties and values from
// [AddonPackageEntitlementDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig],
// [AddonPackageEntitlementDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig],
// [AddonPackageEntitlementDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AddonPackageEntitlementDataFeatureResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u AddonPackageEntitlementDataFeatureResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v AddonPackageEntitlementDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AddonPackageEntitlementDataFeatureResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v AddonPackageEntitlementDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AddonPackageEntitlementDataFeatureResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v AddonPackageEntitlementDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AddonPackageEntitlementDataFeatureResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AddonPackageEntitlementDataFeatureResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type AddonPackageEntitlementDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r AddonPackageEntitlementDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AddonPackageEntitlementDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type AddonPackageEntitlementDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r AddonPackageEntitlementDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AddonPackageEntitlementDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type AddonPackageEntitlementDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r AddonPackageEntitlementDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AddonPackageEntitlementDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement response
type AddonPackageEntitlementDataCredit struct {
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
	// Optional description of the entitlement
	Description string `json:"description" api:"required"`
	// Override display name for the entitlement
	DisplayNameOverride string `json:"displayNameOverride" api:"required"`
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
	// Entitlement type (FEATURE or CREDIT)
	Type constant.Credit `json:"type" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The feature ID this entitlement depends on (for credit entitlements). The
	// entitlement value will be calculated as: base amount × dependency feature usage
	// limit
	DependencyFeatureID string `json:"dependencyFeatureId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Amount              respjson.Field
		Behavior            respjson.Field
		Cadence             respjson.Field
		CreatedAt           respjson.Field
		Description         respjson.Field
		DisplayNameOverride respjson.Field
		HiddenFromWidgets   respjson.Field
		IsCustom            respjson.Field
		IsGranted           respjson.Field
		Order               respjson.Field
		Type                respjson.Field
		UpdatedAt           respjson.Field
		DependencyFeatureID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddonPackageEntitlementDataCredit) RawJSON() string { return r.JSON.raw }
func (r *AddonPackageEntitlementDataCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1AddonEntitlementNewResponse struct {
	Data []V1AddonEntitlementNewResponseDataUnion `json:"data" api:"required"`
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

// V1AddonEntitlementNewResponseDataUnion contains all possible properties and
// values from [V1AddonEntitlementNewResponseDataFeature],
// [V1AddonEntitlementNewResponseDataCredit].
//
// Use the [V1AddonEntitlementNewResponseDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1AddonEntitlementNewResponseDataUnion struct {
	ID                  string    `json:"id"`
	Behavior            string    `json:"behavior"`
	CreatedAt           time.Time `json:"createdAt"`
	Description         string    `json:"description"`
	DisplayNameOverride string    `json:"displayNameOverride"`
	// This field is from variant [V1AddonEntitlementNewResponseDataFeature].
	EnumValues []string `json:"enumValues"`
	// This field is from variant [V1AddonEntitlementNewResponseDataFeature].
	HasSoftLimit bool `json:"hasSoftLimit"`
	// This field is from variant [V1AddonEntitlementNewResponseDataFeature].
	HasUnlimitedUsage bool     `json:"hasUnlimitedUsage"`
	HiddenFromWidgets []string `json:"hiddenFromWidgets"`
	IsCustom          bool     `json:"isCustom"`
	IsGranted         bool     `json:"isGranted"`
	Order             float64  `json:"order"`
	// This field is from variant [V1AddonEntitlementNewResponseDataFeature].
	ResetPeriod string `json:"resetPeriod"`
	// This field is from variant [V1AddonEntitlementNewResponseDataFeature].
	ResetPeriodConfiguration V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration"`
	// Any of "FEATURE", "CREDIT".
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updatedAt"`
	// This field is from variant [V1AddonEntitlementNewResponseDataFeature].
	UsageLimit float64 `json:"usageLimit"`
	// This field is from variant [V1AddonEntitlementNewResponseDataCredit].
	Amount float64 `json:"amount"`
	// This field is from variant [V1AddonEntitlementNewResponseDataCredit].
	Cadence string `json:"cadence"`
	// This field is from variant [V1AddonEntitlementNewResponseDataCredit].
	DependencyFeatureID string `json:"dependencyFeatureId"`
	JSON                struct {
		ID                       respjson.Field
		Behavior                 respjson.Field
		CreatedAt                respjson.Field
		Description              respjson.Field
		DisplayNameOverride      respjson.Field
		EnumValues               respjson.Field
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
		Amount                   respjson.Field
		Cadence                  respjson.Field
		DependencyFeatureID      respjson.Field
		raw                      string
	} `json:"-"`
}

// anyV1AddonEntitlementNewResponseData is implemented by each variant of
// [V1AddonEntitlementNewResponseDataUnion] to add type safety for the return type
// of [V1AddonEntitlementNewResponseDataUnion.AsAny]
type anyV1AddonEntitlementNewResponseData interface {
	implV1AddonEntitlementNewResponseDataUnion()
}

func (V1AddonEntitlementNewResponseDataFeature) implV1AddonEntitlementNewResponseDataUnion() {}
func (V1AddonEntitlementNewResponseDataCredit) implV1AddonEntitlementNewResponseDataUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := V1AddonEntitlementNewResponseDataUnion.AsAny().(type) {
//	case stigg.V1AddonEntitlementNewResponseDataFeature:
//	case stigg.V1AddonEntitlementNewResponseDataCredit:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u V1AddonEntitlementNewResponseDataUnion) AsAny() anyV1AddonEntitlementNewResponseData {
	switch u.Type {
	case "FEATURE":
		return u.AsFeature()
	case "CREDIT":
		return u.AsCredit()
	}
	return nil
}

func (u V1AddonEntitlementNewResponseDataUnion) AsFeature() (v V1AddonEntitlementNewResponseDataFeature) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1AddonEntitlementNewResponseDataUnion) AsCredit() (v V1AddonEntitlementNewResponseDataCredit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1AddonEntitlementNewResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1AddonEntitlementNewResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature entitlement response
type V1AddonEntitlementNewResponseDataFeature struct {
	// Unique identifier of the entitlement
	ID string `json:"id" api:"required"`
	// Entitlement behavior (Increment or Override)
	//
	// Any of "Increment", "Override".
	Behavior string `json:"behavior" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Optional description of the entitlement
	Description string `json:"description" api:"required"`
	// Override display name for the entitlement
	DisplayNameOverride string `json:"displayNameOverride" api:"required"`
	// Allowed enum values (for feature entitlements)
	EnumValues []string `json:"enumValues" api:"required"`
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
	ResetPeriodConfiguration V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
	// Entitlement type (FEATURE or CREDIT)
	Type constant.Feature `json:"type" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Usage limit (for feature entitlements)
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Behavior                 respjson.Field
		CreatedAt                respjson.Field
		Description              respjson.Field
		DisplayNameOverride      respjson.Field
		EnumValues               respjson.Field
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
func (r V1AddonEntitlementNewResponseDataFeature) RawJSON() string { return r.JSON.raw }
func (r *V1AddonEntitlementNewResponseDataFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion contains
// all possible properties and values from
// [V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementNewResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement response
type V1AddonEntitlementNewResponseDataCredit struct {
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
	// Optional description of the entitlement
	Description string `json:"description" api:"required"`
	// Override display name for the entitlement
	DisplayNameOverride string `json:"displayNameOverride" api:"required"`
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
	// Entitlement type (FEATURE or CREDIT)
	Type constant.Credit `json:"type" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The feature ID this entitlement depends on (for credit entitlements). The
	// entitlement value will be calculated as: base amount × dependency feature usage
	// limit
	DependencyFeatureID string `json:"dependencyFeatureId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Amount              respjson.Field
		Behavior            respjson.Field
		Cadence             respjson.Field
		CreatedAt           respjson.Field
		Description         respjson.Field
		DisplayNameOverride respjson.Field
		HiddenFromWidgets   respjson.Field
		IsCustom            respjson.Field
		IsGranted           respjson.Field
		Order               respjson.Field
		Type                respjson.Field
		UpdatedAt           respjson.Field
		DependencyFeatureID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AddonEntitlementNewResponseDataCredit) RawJSON() string { return r.JSON.raw }
func (r *V1AddonEntitlementNewResponseDataCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response list object
type V1AddonEntitlementListResponse struct {
	Data []V1AddonEntitlementListResponseDataUnion `json:"data" api:"required"`
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

// V1AddonEntitlementListResponseDataUnion contains all possible properties and
// values from [V1AddonEntitlementListResponseDataFeature],
// [V1AddonEntitlementListResponseDataCredit].
//
// Use the [V1AddonEntitlementListResponseDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1AddonEntitlementListResponseDataUnion struct {
	ID                  string    `json:"id"`
	Behavior            string    `json:"behavior"`
	CreatedAt           time.Time `json:"createdAt"`
	Description         string    `json:"description"`
	DisplayNameOverride string    `json:"displayNameOverride"`
	// This field is from variant [V1AddonEntitlementListResponseDataFeature].
	EnumValues []string `json:"enumValues"`
	// This field is from variant [V1AddonEntitlementListResponseDataFeature].
	HasSoftLimit bool `json:"hasSoftLimit"`
	// This field is from variant [V1AddonEntitlementListResponseDataFeature].
	HasUnlimitedUsage bool     `json:"hasUnlimitedUsage"`
	HiddenFromWidgets []string `json:"hiddenFromWidgets"`
	IsCustom          bool     `json:"isCustom"`
	IsGranted         bool     `json:"isGranted"`
	Order             float64  `json:"order"`
	// This field is from variant [V1AddonEntitlementListResponseDataFeature].
	ResetPeriod string `json:"resetPeriod"`
	// This field is from variant [V1AddonEntitlementListResponseDataFeature].
	ResetPeriodConfiguration V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration"`
	// Any of "FEATURE", "CREDIT".
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updatedAt"`
	// This field is from variant [V1AddonEntitlementListResponseDataFeature].
	UsageLimit float64 `json:"usageLimit"`
	// This field is from variant [V1AddonEntitlementListResponseDataCredit].
	Amount float64 `json:"amount"`
	// This field is from variant [V1AddonEntitlementListResponseDataCredit].
	Cadence string `json:"cadence"`
	// This field is from variant [V1AddonEntitlementListResponseDataCredit].
	DependencyFeatureID string `json:"dependencyFeatureId"`
	JSON                struct {
		ID                       respjson.Field
		Behavior                 respjson.Field
		CreatedAt                respjson.Field
		Description              respjson.Field
		DisplayNameOverride      respjson.Field
		EnumValues               respjson.Field
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
		Amount                   respjson.Field
		Cadence                  respjson.Field
		DependencyFeatureID      respjson.Field
		raw                      string
	} `json:"-"`
}

// anyV1AddonEntitlementListResponseData is implemented by each variant of
// [V1AddonEntitlementListResponseDataUnion] to add type safety for the return type
// of [V1AddonEntitlementListResponseDataUnion.AsAny]
type anyV1AddonEntitlementListResponseData interface {
	implV1AddonEntitlementListResponseDataUnion()
}

func (V1AddonEntitlementListResponseDataFeature) implV1AddonEntitlementListResponseDataUnion() {}
func (V1AddonEntitlementListResponseDataCredit) implV1AddonEntitlementListResponseDataUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := V1AddonEntitlementListResponseDataUnion.AsAny().(type) {
//	case stigg.V1AddonEntitlementListResponseDataFeature:
//	case stigg.V1AddonEntitlementListResponseDataCredit:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u V1AddonEntitlementListResponseDataUnion) AsAny() anyV1AddonEntitlementListResponseData {
	switch u.Type {
	case "FEATURE":
		return u.AsFeature()
	case "CREDIT":
		return u.AsCredit()
	}
	return nil
}

func (u V1AddonEntitlementListResponseDataUnion) AsFeature() (v V1AddonEntitlementListResponseDataFeature) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1AddonEntitlementListResponseDataUnion) AsCredit() (v V1AddonEntitlementListResponseDataCredit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1AddonEntitlementListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1AddonEntitlementListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature entitlement response
type V1AddonEntitlementListResponseDataFeature struct {
	// Unique identifier of the entitlement
	ID string `json:"id" api:"required"`
	// Entitlement behavior (Increment or Override)
	//
	// Any of "Increment", "Override".
	Behavior string `json:"behavior" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Optional description of the entitlement
	Description string `json:"description" api:"required"`
	// Override display name for the entitlement
	DisplayNameOverride string `json:"displayNameOverride" api:"required"`
	// Allowed enum values (for feature entitlements)
	EnumValues []string `json:"enumValues" api:"required"`
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
	ResetPeriodConfiguration V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
	// Entitlement type (FEATURE or CREDIT)
	Type constant.Feature `json:"type" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Usage limit (for feature entitlements)
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Behavior                 respjson.Field
		CreatedAt                respjson.Field
		Description              respjson.Field
		DisplayNameOverride      respjson.Field
		EnumValues               respjson.Field
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
func (r V1AddonEntitlementListResponseDataFeature) RawJSON() string { return r.JSON.raw }
func (r *V1AddonEntitlementListResponseDataFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationUnion contains
// all possible properties and values from
// [V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AddonEntitlementListResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement response
type V1AddonEntitlementListResponseDataCredit struct {
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
	// Optional description of the entitlement
	Description string `json:"description" api:"required"`
	// Override display name for the entitlement
	DisplayNameOverride string `json:"displayNameOverride" api:"required"`
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
	// Entitlement type (FEATURE or CREDIT)
	Type constant.Credit `json:"type" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The feature ID this entitlement depends on (for credit entitlements). The
	// entitlement value will be calculated as: base amount × dependency feature usage
	// limit
	DependencyFeatureID string `json:"dependencyFeatureId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Amount              respjson.Field
		Behavior            respjson.Field
		Cadence             respjson.Field
		CreatedAt           respjson.Field
		Description         respjson.Field
		DisplayNameOverride respjson.Field
		HiddenFromWidgets   respjson.Field
		IsCustom            respjson.Field
		IsGranted           respjson.Field
		Order               respjson.Field
		Type                respjson.Field
		UpdatedAt           respjson.Field
		DependencyFeatureID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AddonEntitlementListResponseDataCredit) RawJSON() string { return r.JSON.raw }
func (r *V1AddonEntitlementListResponseDataCredit) UnmarshalJSON(data []byte) error {
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
	Entitlements []V1AddonEntitlementNewParamsEntitlementUnion `json:"entitlements,omitzero" api:"required"`
	paramObj
}

func (r V1AddonEntitlementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1AddonEntitlementNewParamsEntitlementUnion struct {
	OfFeature *V1AddonEntitlementNewParamsEntitlementFeature `json:",omitzero,inline"`
	OfCredit  *V1AddonEntitlementNewParamsEntitlementCredit  `json:",omitzero,inline"`
	paramUnion
}

func (u V1AddonEntitlementNewParamsEntitlementUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFeature, u.OfCredit)
}
func (u *V1AddonEntitlementNewParamsEntitlementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1AddonEntitlementNewParamsEntitlementUnion) asAny() any {
	if !param.IsOmitted(u.OfFeature) {
		return u.OfFeature
	} else if !param.IsOmitted(u.OfCredit) {
		return u.OfCredit
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetEnumValues() []string {
	if vt := u.OfFeature; vt != nil {
		return vt.EnumValues
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetHasSoftLimit() *bool {
	if vt := u.OfFeature; vt != nil && vt.HasSoftLimit.Valid() {
		return &vt.HasSoftLimit.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetHasUnlimitedUsage() *bool {
	if vt := u.OfFeature; vt != nil && vt.HasUnlimitedUsage.Valid() {
		return &vt.HasUnlimitedUsage.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetMonthlyResetPeriodConfiguration() *V1AddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration {
	if vt := u.OfFeature; vt != nil {
		return &vt.MonthlyResetPeriodConfiguration
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetResetPeriod() *string {
	if vt := u.OfFeature; vt != nil {
		return &vt.ResetPeriod
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetUsageLimit() *int64 {
	if vt := u.OfFeature; vt != nil && vt.UsageLimit.Valid() {
		return &vt.UsageLimit.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetWeeklyResetPeriodConfiguration() *V1AddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration {
	if vt := u.OfFeature; vt != nil {
		return &vt.WeeklyResetPeriodConfiguration
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetYearlyResetPeriodConfiguration() *V1AddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration {
	if vt := u.OfFeature; vt != nil {
		return &vt.YearlyResetPeriodConfiguration
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetAmount() *float64 {
	if vt := u.OfCredit; vt != nil && vt.Amount.Valid() {
		return &vt.Amount.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetCadence() *string {
	if vt := u.OfCredit; vt != nil {
		return &vt.Cadence
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetDependencyFeatureID() *string {
	if vt := u.OfCredit; vt != nil && vt.DependencyFeatureID.Valid() {
		return &vt.DependencyFeatureID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetID() *string {
	if vt := u.OfFeature; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfCredit; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetType() *string {
	if vt := u.OfFeature; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCredit; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetBehavior() *string {
	if vt := u.OfFeature; vt != nil {
		return (*string)(&vt.Behavior)
	} else if vt := u.OfCredit; vt != nil {
		return (*string)(&vt.Behavior)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetDescription() *string {
	if vt := u.OfFeature; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfCredit; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetDisplayNameOverride() *string {
	if vt := u.OfFeature; vt != nil && vt.DisplayNameOverride.Valid() {
		return &vt.DisplayNameOverride.Value
	} else if vt := u.OfCredit; vt != nil && vt.DisplayNameOverride.Valid() {
		return &vt.DisplayNameOverride.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetIsCustom() *bool {
	if vt := u.OfFeature; vt != nil && vt.IsCustom.Valid() {
		return &vt.IsCustom.Value
	} else if vt := u.OfCredit; vt != nil && vt.IsCustom.Valid() {
		return &vt.IsCustom.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetIsGranted() *bool {
	if vt := u.OfFeature; vt != nil && vt.IsGranted.Valid() {
		return &vt.IsGranted.Value
	} else if vt := u.OfCredit; vt != nil && vt.IsGranted.Valid() {
		return &vt.IsGranted.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetOrder() *float64 {
	if vt := u.OfFeature; vt != nil && vt.Order.Valid() {
		return &vt.Order.Value
	} else if vt := u.OfCredit; vt != nil && vt.Order.Valid() {
		return &vt.Order.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's HiddenFromWidgets property, if
// present.
func (u V1AddonEntitlementNewParamsEntitlementUnion) GetHiddenFromWidgets() []string {
	if vt := u.OfFeature; vt != nil {
		return vt.HiddenFromWidgets
	} else if vt := u.OfCredit; vt != nil {
		return vt.HiddenFromWidgets
	}
	return nil
}

func init() {
	apijson.RegisterUnion[V1AddonEntitlementNewParamsEntitlementUnion](
		"type",
		apijson.Discriminator[V1AddonEntitlementNewParamsEntitlementFeature]("FEATURE"),
		apijson.Discriminator[V1AddonEntitlementNewParamsEntitlementCredit]("CREDIT"),
	)
}

// Request to create a feature entitlement
//
// The properties ID, Type are required.
type V1AddonEntitlementNewParamsEntitlementFeature struct {
	// The feature ID to attach the entitlement to
	ID string `json:"id" api:"required"`
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
	// CreateFeatureEntitlementRequest
	//
	// This field can be elided, and will marshal its zero value as "FEATURE".
	Type constant.Feature `json:"type" api:"required"`
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

// Request to create a credit entitlement
//
// The properties ID, Amount, Cadence, Type are required.
type V1AddonEntitlementNewParamsEntitlementCredit struct {
	// Credit grant amount
	Amount param.Opt[float64] `json:"amount,omitzero" api:"required"`
	// The custom currency ID for the credit entitlement
	ID string `json:"id" api:"required"`
	// Credit grant cadence (MONTH or YEAR)
	//
	// Any of "MONTH", "YEAR".
	Cadence string `json:"cadence,omitzero" api:"required"`
	// The feature ID this entitlement depends on. The entitlement value will be
	// calculated as: base amount × dependency feature usage limit
	DependencyFeatureID param.Opt[string] `json:"dependencyFeatureId,omitzero"`
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
	// CreateCreditEntitlementRequest
	//
	// This field can be elided, and will marshal its zero value as "CREDIT".
	Type constant.Credit `json:"type" api:"required"`
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

type V1AddonEntitlementUpdateParams struct {
	AddonID string `path:"addonId" api:"required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Fields
	// to update on a feature entitlement
	OfFeature *V1AddonEntitlementUpdateParamsBodyFeature `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Fields
	// to update on a credit entitlement
	OfCredit *V1AddonEntitlementUpdateParamsBodyCredit `json:",inline"`

	paramObj
}

func (u V1AddonEntitlementUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFeature, u.OfCredit)
}
func (r *V1AddonEntitlementUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fields to update on a feature entitlement
//
// The property Type is required.
type V1AddonEntitlementUpdateParamsBodyFeature struct {
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
	MonthlyResetPeriodConfiguration V1AddonEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Configuration for weekly reset period
	WeeklyResetPeriodConfiguration V1AddonEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Configuration for yearly reset period
	YearlyResetPeriodConfiguration V1AddonEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
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
	// UpdateFeatureEntitlementRequest
	//
	// This field can be elided, and will marshal its zero value as "FEATURE".
	Type constant.Feature `json:"type" api:"required"`
	paramObj
}

func (r V1AddonEntitlementUpdateParamsBodyFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementUpdateParamsBodyFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementUpdateParamsBodyFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsBodyFeature](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsBodyFeature](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Configuration for monthly reset period
//
// The property AccordingTo is required.
type V1AddonEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1AddonEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Configuration for weekly reset period
//
// The property AccordingTo is required.
type V1AddonEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1AddonEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Configuration for yearly reset period
//
// The property AccordingTo is required.
type V1AddonEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1AddonEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

// Fields to update on a credit entitlement
//
// The property Type is required.
type V1AddonEntitlementUpdateParamsBodyCredit struct {
	// The feature ID this entitlement depends on. The entitlement value will be
	// calculated as: base amount × dependency feature usage limit
	DependencyFeatureID param.Opt[string] `json:"dependencyFeatureId,omitzero"`
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
	// UpdateCreditEntitlementRequest
	//
	// This field can be elided, and will marshal its zero value as "CREDIT".
	Type constant.Credit `json:"type" api:"required"`
	paramObj
}

func (r V1AddonEntitlementUpdateParamsBodyCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1AddonEntitlementUpdateParamsBodyCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AddonEntitlementUpdateParamsBodyCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsBodyCredit](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1AddonEntitlementUpdateParamsBodyCredit](
		"cadence", "MONTH", "YEAR",
	)
}

type V1AddonEntitlementDeleteParams struct {
	AddonID string `path:"addonId" api:"required" json:"-"`
	paramObj
}
