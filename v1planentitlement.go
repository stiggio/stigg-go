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
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s/entitlements", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates an existing entitlement on a draft plan.
func (r *V1PlanEntitlementService) Update(ctx context.Context, id string, params V1PlanEntitlementUpdateParams, opts ...option.RequestOption) (res *PlanEntitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.PlanID == "" {
		err = errors.New("missing required planId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s/entitlements/%s", params.PlanID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieves a list of entitlements for a plan.
func (r *V1PlanEntitlementService) List(ctx context.Context, planID string, opts ...option.RequestOption) (res *V1PlanEntitlementListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required planId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s/entitlements", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes an entitlement from a draft plan.
func (r *V1PlanEntitlementService) Delete(ctx context.Context, id string, body V1PlanEntitlementDeleteParams, opts ...option.RequestOption) (res *PlanEntitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.PlanID == "" {
		err = errors.New("missing required planId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s/entitlements/%s", body.PlanID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Response object
type PlanEntitlement struct {
	// Feature entitlement response
	Data PlanEntitlementDataUnion `json:"data" api:"required"`
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

// PlanEntitlementDataUnion contains all possible properties and values from
// [PlanEntitlementDataFeature], [PlanEntitlementDataCredit].
//
// Use the [PlanEntitlementDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PlanEntitlementDataUnion struct {
	ID                  string    `json:"id"`
	Behavior            string    `json:"behavior"`
	CreatedAt           time.Time `json:"createdAt"`
	Description         string    `json:"description"`
	DisplayNameOverride string    `json:"displayNameOverride"`
	// This field is from variant [PlanEntitlementDataFeature].
	EnumValues []string `json:"enumValues"`
	// This field is from variant [PlanEntitlementDataFeature].
	HasSoftLimit bool `json:"hasSoftLimit"`
	// This field is from variant [PlanEntitlementDataFeature].
	HasUnlimitedUsage bool     `json:"hasUnlimitedUsage"`
	HiddenFromWidgets []string `json:"hiddenFromWidgets"`
	IsCustom          bool     `json:"isCustom"`
	IsGranted         bool     `json:"isGranted"`
	Order             float64  `json:"order"`
	// This field is from variant [PlanEntitlementDataFeature].
	ResetPeriod string `json:"resetPeriod"`
	// This field is from variant [PlanEntitlementDataFeature].
	ResetPeriodConfiguration PlanEntitlementDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration"`
	// Any of "FEATURE", "CREDIT".
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updatedAt"`
	// This field is from variant [PlanEntitlementDataFeature].
	UsageLimit float64 `json:"usageLimit"`
	// This field is from variant [PlanEntitlementDataCredit].
	Amount float64 `json:"amount"`
	// This field is from variant [PlanEntitlementDataCredit].
	Cadence string `json:"cadence"`
	JSON    struct {
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
		raw                      string
	} `json:"-"`
}

// anyPlanEntitlementData is implemented by each variant of
// [PlanEntitlementDataUnion] to add type safety for the return type of
// [PlanEntitlementDataUnion.AsAny]
type anyPlanEntitlementData interface {
	implPlanEntitlementDataUnion()
}

func (PlanEntitlementDataFeature) implPlanEntitlementDataUnion() {}
func (PlanEntitlementDataCredit) implPlanEntitlementDataUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PlanEntitlementDataUnion.AsAny().(type) {
//	case stigg.PlanEntitlementDataFeature:
//	case stigg.PlanEntitlementDataCredit:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PlanEntitlementDataUnion) AsAny() anyPlanEntitlementData {
	switch u.Type {
	case "FEATURE":
		return u.AsFeature()
	case "CREDIT":
		return u.AsCredit()
	}
	return nil
}

func (u PlanEntitlementDataUnion) AsFeature() (v PlanEntitlementDataFeature) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlanEntitlementDataUnion) AsCredit() (v PlanEntitlementDataCredit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlanEntitlementDataUnion) RawJSON() string { return u.JSON.raw }

func (r *PlanEntitlementDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature entitlement response
type PlanEntitlementDataFeature struct {
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
	ResetPeriodConfiguration PlanEntitlementDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
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
func (r PlanEntitlementDataFeature) RawJSON() string { return r.JSON.raw }
func (r *PlanEntitlementDataFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PlanEntitlementDataFeatureResetPeriodConfigurationUnion contains all possible
// properties and values from
// [PlanEntitlementDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig],
// [PlanEntitlementDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig],
// [PlanEntitlementDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PlanEntitlementDataFeatureResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u PlanEntitlementDataFeatureResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v PlanEntitlementDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlanEntitlementDataFeatureResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v PlanEntitlementDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlanEntitlementDataFeatureResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v PlanEntitlementDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlanEntitlementDataFeatureResetPeriodConfigurationUnion) RawJSON() string { return u.JSON.raw }

func (r *PlanEntitlementDataFeatureResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type PlanEntitlementDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r PlanEntitlementDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *PlanEntitlementDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type PlanEntitlementDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r PlanEntitlementDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *PlanEntitlementDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type PlanEntitlementDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r PlanEntitlementDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *PlanEntitlementDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement response
type PlanEntitlementDataCredit struct {
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
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanEntitlementDataCredit) RawJSON() string { return r.JSON.raw }
func (r *PlanEntitlementDataCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1PlanEntitlementNewResponse struct {
	Data []V1PlanEntitlementNewResponseDataUnion `json:"data" api:"required"`
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

// V1PlanEntitlementNewResponseDataUnion contains all possible properties and
// values from [V1PlanEntitlementNewResponseDataFeature],
// [V1PlanEntitlementNewResponseDataCredit].
//
// Use the [V1PlanEntitlementNewResponseDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1PlanEntitlementNewResponseDataUnion struct {
	ID                  string    `json:"id"`
	Behavior            string    `json:"behavior"`
	CreatedAt           time.Time `json:"createdAt"`
	Description         string    `json:"description"`
	DisplayNameOverride string    `json:"displayNameOverride"`
	// This field is from variant [V1PlanEntitlementNewResponseDataFeature].
	EnumValues []string `json:"enumValues"`
	// This field is from variant [V1PlanEntitlementNewResponseDataFeature].
	HasSoftLimit bool `json:"hasSoftLimit"`
	// This field is from variant [V1PlanEntitlementNewResponseDataFeature].
	HasUnlimitedUsage bool     `json:"hasUnlimitedUsage"`
	HiddenFromWidgets []string `json:"hiddenFromWidgets"`
	IsCustom          bool     `json:"isCustom"`
	IsGranted         bool     `json:"isGranted"`
	Order             float64  `json:"order"`
	// This field is from variant [V1PlanEntitlementNewResponseDataFeature].
	ResetPeriod string `json:"resetPeriod"`
	// This field is from variant [V1PlanEntitlementNewResponseDataFeature].
	ResetPeriodConfiguration V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration"`
	// Any of "FEATURE", "CREDIT".
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updatedAt"`
	// This field is from variant [V1PlanEntitlementNewResponseDataFeature].
	UsageLimit float64 `json:"usageLimit"`
	// This field is from variant [V1PlanEntitlementNewResponseDataCredit].
	Amount float64 `json:"amount"`
	// This field is from variant [V1PlanEntitlementNewResponseDataCredit].
	Cadence string `json:"cadence"`
	JSON    struct {
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
		raw                      string
	} `json:"-"`
}

// anyV1PlanEntitlementNewResponseData is implemented by each variant of
// [V1PlanEntitlementNewResponseDataUnion] to add type safety for the return type
// of [V1PlanEntitlementNewResponseDataUnion.AsAny]
type anyV1PlanEntitlementNewResponseData interface {
	implV1PlanEntitlementNewResponseDataUnion()
}

func (V1PlanEntitlementNewResponseDataFeature) implV1PlanEntitlementNewResponseDataUnion() {}
func (V1PlanEntitlementNewResponseDataCredit) implV1PlanEntitlementNewResponseDataUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := V1PlanEntitlementNewResponseDataUnion.AsAny().(type) {
//	case stigg.V1PlanEntitlementNewResponseDataFeature:
//	case stigg.V1PlanEntitlementNewResponseDataCredit:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u V1PlanEntitlementNewResponseDataUnion) AsAny() anyV1PlanEntitlementNewResponseData {
	switch u.Type {
	case "FEATURE":
		return u.AsFeature()
	case "CREDIT":
		return u.AsCredit()
	}
	return nil
}

func (u V1PlanEntitlementNewResponseDataUnion) AsFeature() (v V1PlanEntitlementNewResponseDataFeature) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1PlanEntitlementNewResponseDataUnion) AsCredit() (v V1PlanEntitlementNewResponseDataCredit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1PlanEntitlementNewResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1PlanEntitlementNewResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature entitlement response
type V1PlanEntitlementNewResponseDataFeature struct {
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
	ResetPeriodConfiguration V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
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
func (r V1PlanEntitlementNewResponseDataFeature) RawJSON() string { return r.JSON.raw }
func (r *V1PlanEntitlementNewResponseDataFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion contains
// all possible properties and values from
// [V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementNewResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement response
type V1PlanEntitlementNewResponseDataCredit struct {
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
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanEntitlementNewResponseDataCredit) RawJSON() string { return r.JSON.raw }
func (r *V1PlanEntitlementNewResponseDataCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response list object
type V1PlanEntitlementListResponse struct {
	Data []V1PlanEntitlementListResponseDataUnion `json:"data" api:"required"`
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

// V1PlanEntitlementListResponseDataUnion contains all possible properties and
// values from [V1PlanEntitlementListResponseDataFeature],
// [V1PlanEntitlementListResponseDataCredit].
//
// Use the [V1PlanEntitlementListResponseDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1PlanEntitlementListResponseDataUnion struct {
	ID                  string    `json:"id"`
	Behavior            string    `json:"behavior"`
	CreatedAt           time.Time `json:"createdAt"`
	Description         string    `json:"description"`
	DisplayNameOverride string    `json:"displayNameOverride"`
	// This field is from variant [V1PlanEntitlementListResponseDataFeature].
	EnumValues []string `json:"enumValues"`
	// This field is from variant [V1PlanEntitlementListResponseDataFeature].
	HasSoftLimit bool `json:"hasSoftLimit"`
	// This field is from variant [V1PlanEntitlementListResponseDataFeature].
	HasUnlimitedUsage bool     `json:"hasUnlimitedUsage"`
	HiddenFromWidgets []string `json:"hiddenFromWidgets"`
	IsCustom          bool     `json:"isCustom"`
	IsGranted         bool     `json:"isGranted"`
	Order             float64  `json:"order"`
	// This field is from variant [V1PlanEntitlementListResponseDataFeature].
	ResetPeriod string `json:"resetPeriod"`
	// This field is from variant [V1PlanEntitlementListResponseDataFeature].
	ResetPeriodConfiguration V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration"`
	// Any of "FEATURE", "CREDIT".
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updatedAt"`
	// This field is from variant [V1PlanEntitlementListResponseDataFeature].
	UsageLimit float64 `json:"usageLimit"`
	// This field is from variant [V1PlanEntitlementListResponseDataCredit].
	Amount float64 `json:"amount"`
	// This field is from variant [V1PlanEntitlementListResponseDataCredit].
	Cadence string `json:"cadence"`
	JSON    struct {
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
		raw                      string
	} `json:"-"`
}

// anyV1PlanEntitlementListResponseData is implemented by each variant of
// [V1PlanEntitlementListResponseDataUnion] to add type safety for the return type
// of [V1PlanEntitlementListResponseDataUnion.AsAny]
type anyV1PlanEntitlementListResponseData interface {
	implV1PlanEntitlementListResponseDataUnion()
}

func (V1PlanEntitlementListResponseDataFeature) implV1PlanEntitlementListResponseDataUnion() {}
func (V1PlanEntitlementListResponseDataCredit) implV1PlanEntitlementListResponseDataUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := V1PlanEntitlementListResponseDataUnion.AsAny().(type) {
//	case stigg.V1PlanEntitlementListResponseDataFeature:
//	case stigg.V1PlanEntitlementListResponseDataCredit:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u V1PlanEntitlementListResponseDataUnion) AsAny() anyV1PlanEntitlementListResponseData {
	switch u.Type {
	case "FEATURE":
		return u.AsFeature()
	case "CREDIT":
		return u.AsCredit()
	}
	return nil
}

func (u V1PlanEntitlementListResponseDataUnion) AsFeature() (v V1PlanEntitlementListResponseDataFeature) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1PlanEntitlementListResponseDataUnion) AsCredit() (v V1PlanEntitlementListResponseDataCredit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1PlanEntitlementListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1PlanEntitlementListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature entitlement response
type V1PlanEntitlementListResponseDataFeature struct {
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
	ResetPeriodConfiguration V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
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
func (r V1PlanEntitlementListResponseDataFeature) RawJSON() string { return r.JSON.raw }
func (r *V1PlanEntitlementListResponseDataFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationUnion contains
// all possible properties and values from
// [V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1PlanEntitlementListResponseDataFeatureResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit entitlement response
type V1PlanEntitlementListResponseDataCredit struct {
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
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanEntitlementListResponseDataCredit) RawJSON() string { return r.JSON.raw }
func (r *V1PlanEntitlementListResponseDataCredit) UnmarshalJSON(data []byte) error {
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
	Entitlements []V1PlanEntitlementNewParamsEntitlementUnion `json:"entitlements,omitzero" api:"required"`
	paramObj
}

func (r V1PlanEntitlementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1PlanEntitlementNewParamsEntitlementUnion struct {
	OfFeature *V1PlanEntitlementNewParamsEntitlementFeature `json:",omitzero,inline"`
	OfCredit  *V1PlanEntitlementNewParamsEntitlementCredit  `json:",omitzero,inline"`
	paramUnion
}

func (u V1PlanEntitlementNewParamsEntitlementUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFeature, u.OfCredit)
}
func (u *V1PlanEntitlementNewParamsEntitlementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1PlanEntitlementNewParamsEntitlementUnion) asAny() any {
	if !param.IsOmitted(u.OfFeature) {
		return u.OfFeature
	} else if !param.IsOmitted(u.OfCredit) {
		return u.OfCredit
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetEnumValues() []string {
	if vt := u.OfFeature; vt != nil {
		return vt.EnumValues
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetHasSoftLimit() *bool {
	if vt := u.OfFeature; vt != nil && vt.HasSoftLimit.Valid() {
		return &vt.HasSoftLimit.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetHasUnlimitedUsage() *bool {
	if vt := u.OfFeature; vt != nil && vt.HasUnlimitedUsage.Valid() {
		return &vt.HasUnlimitedUsage.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetMonthlyResetPeriodConfiguration() *V1PlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration {
	if vt := u.OfFeature; vt != nil {
		return &vt.MonthlyResetPeriodConfiguration
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetResetPeriod() *string {
	if vt := u.OfFeature; vt != nil {
		return &vt.ResetPeriod
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetUsageLimit() *int64 {
	if vt := u.OfFeature; vt != nil && vt.UsageLimit.Valid() {
		return &vt.UsageLimit.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetWeeklyResetPeriodConfiguration() *V1PlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration {
	if vt := u.OfFeature; vt != nil {
		return &vt.WeeklyResetPeriodConfiguration
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetYearlyResetPeriodConfiguration() *V1PlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration {
	if vt := u.OfFeature; vt != nil {
		return &vt.YearlyResetPeriodConfiguration
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetAmount() *float64 {
	if vt := u.OfCredit; vt != nil && vt.Amount.Valid() {
		return &vt.Amount.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetCadence() *string {
	if vt := u.OfCredit; vt != nil {
		return &vt.Cadence
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetID() *string {
	if vt := u.OfFeature; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfCredit; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetType() *string {
	if vt := u.OfFeature; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCredit; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetBehavior() *string {
	if vt := u.OfFeature; vt != nil {
		return (*string)(&vt.Behavior)
	} else if vt := u.OfCredit; vt != nil {
		return (*string)(&vt.Behavior)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetDescription() *string {
	if vt := u.OfFeature; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfCredit; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetDisplayNameOverride() *string {
	if vt := u.OfFeature; vt != nil && vt.DisplayNameOverride.Valid() {
		return &vt.DisplayNameOverride.Value
	} else if vt := u.OfCredit; vt != nil && vt.DisplayNameOverride.Valid() {
		return &vt.DisplayNameOverride.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetIsCustom() *bool {
	if vt := u.OfFeature; vt != nil && vt.IsCustom.Valid() {
		return &vt.IsCustom.Value
	} else if vt := u.OfCredit; vt != nil && vt.IsCustom.Valid() {
		return &vt.IsCustom.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetIsGranted() *bool {
	if vt := u.OfFeature; vt != nil && vt.IsGranted.Valid() {
		return &vt.IsGranted.Value
	} else if vt := u.OfCredit; vt != nil && vt.IsGranted.Valid() {
		return &vt.IsGranted.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetOrder() *float64 {
	if vt := u.OfFeature; vt != nil && vt.Order.Valid() {
		return &vt.Order.Value
	} else if vt := u.OfCredit; vt != nil && vt.Order.Valid() {
		return &vt.Order.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's HiddenFromWidgets property, if
// present.
func (u V1PlanEntitlementNewParamsEntitlementUnion) GetHiddenFromWidgets() []string {
	if vt := u.OfFeature; vt != nil {
		return vt.HiddenFromWidgets
	} else if vt := u.OfCredit; vt != nil {
		return vt.HiddenFromWidgets
	}
	return nil
}

func init() {
	apijson.RegisterUnion[V1PlanEntitlementNewParamsEntitlementUnion](
		"type",
		apijson.Discriminator[V1PlanEntitlementNewParamsEntitlementFeature]("FEATURE"),
		apijson.Discriminator[V1PlanEntitlementNewParamsEntitlementCredit]("CREDIT"),
	)
}

// Request to create a feature entitlement
//
// The properties ID, Type are required.
type V1PlanEntitlementNewParamsEntitlementFeature struct {
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
	// CreateFeatureEntitlementRequest
	//
	// This field can be elided, and will marshal its zero value as "FEATURE".
	Type constant.Feature `json:"type" api:"required"`
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

// Request to create a credit entitlement
//
// The properties ID, Amount, Cadence, Type are required.
type V1PlanEntitlementNewParamsEntitlementCredit struct {
	// Credit grant amount
	Amount param.Opt[float64] `json:"amount,omitzero" api:"required"`
	// The custom currency ID for the credit entitlement
	ID string `json:"id" api:"required"`
	// Credit grant cadence (MONTH or YEAR)
	//
	// Any of "MONTH", "YEAR".
	Cadence string `json:"cadence,omitzero" api:"required"`
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

type V1PlanEntitlementUpdateParams struct {
	PlanID string `path:"planId" api:"required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Fields
	// to update on a feature entitlement
	OfFeature *V1PlanEntitlementUpdateParamsBodyFeature `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Fields
	// to update on a credit entitlement
	OfCredit *V1PlanEntitlementUpdateParamsBodyCredit `json:",inline"`

	paramObj
}

func (u V1PlanEntitlementUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFeature, u.OfCredit)
}
func (r *V1PlanEntitlementUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fields to update on a feature entitlement
//
// The property Type is required.
type V1PlanEntitlementUpdateParamsBodyFeature struct {
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
	MonthlyResetPeriodConfiguration V1PlanEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Configuration for weekly reset period
	WeeklyResetPeriodConfiguration V1PlanEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Configuration for yearly reset period
	YearlyResetPeriodConfiguration V1PlanEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
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

func (r V1PlanEntitlementUpdateParamsBodyFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementUpdateParamsBodyFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementUpdateParamsBodyFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsBodyFeature](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsBodyFeature](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Configuration for monthly reset period
//
// The property AccordingTo is required.
type V1PlanEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsBodyFeatureMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Configuration for weekly reset period
//
// The property AccordingTo is required.
type V1PlanEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsBodyFeatureWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Configuration for yearly reset period
//
// The property AccordingTo is required.
type V1PlanEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsBodyFeatureYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

// Fields to update on a credit entitlement
//
// The property Type is required.
type V1PlanEntitlementUpdateParamsBodyCredit struct {
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

func (r V1PlanEntitlementUpdateParamsBodyCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanEntitlementUpdateParamsBodyCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanEntitlementUpdateParamsBodyCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsBodyCredit](
		"behavior", "Increment", "Override",
	)
	apijson.RegisterFieldValidator[V1PlanEntitlementUpdateParamsBodyCredit](
		"cadence", "MONTH", "YEAR",
	)
}

type V1PlanEntitlementDeleteParams struct {
	PlanID string `path:"planId" api:"required" json:"-"`
	paramObj
}
