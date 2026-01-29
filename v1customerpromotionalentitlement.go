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

// V1CustomerPromotionalEntitlementService contains methods and other services that
// help with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerPromotionalEntitlementService] method instead.
type V1CustomerPromotionalEntitlementService struct {
	Options []option.RequestOption
}

// NewV1CustomerPromotionalEntitlementService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1CustomerPromotionalEntitlementService(opts ...option.RequestOption) (r V1CustomerPromotionalEntitlementService) {
	r = V1CustomerPromotionalEntitlementService{}
	r.Options = opts
	return
}

// Create a promotional entitlements
func (r *V1CustomerPromotionalEntitlementService) Grant(ctx context.Context, customerID string, body V1CustomerPromotionalEntitlementGrantParams, opts ...option.RequestOption) (res *V1CustomerPromotionalEntitlementGrantResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/promotional", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Revoke promotional entitlement
func (r *V1CustomerPromotionalEntitlementService) Revoke(ctx context.Context, featureID string, body V1CustomerPromotionalEntitlementRevokeParams, opts ...option.RequestOption) (res *V1CustomerPromotionalEntitlementRevokeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.CustomerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	if featureID == "" {
		err = errors.New("missing required featureId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/promotional/%s", body.CustomerID, featureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Response object
type V1CustomerPromotionalEntitlementGrantResponse struct {
	Data []V1CustomerPromotionalEntitlementGrantResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalEntitlementGrantResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPromotionalEntitlementGrantResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granted feature entitlement
type V1CustomerPromotionalEntitlementGrantResponseData struct {
	// Unique identifier for the entity
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// The description of the entitlement
	Description string `json:"description,required"`
	// The end date of the promotional entitlement
	EndDate time.Time `json:"endDate,required" format:"date-time"`
	// The enum values of the entitlement
	EnumValues []string `json:"enumValues,required"`
	// The unique identifier for the environment
	EnvironmentID string `json:"environmentId,required" format:"uuid"`
	// Feature group IDs associated with this entitlement
	FeatureGroupIDs []string `json:"featureGroupIds,required"`
	// The unique identifier of the entitlement feature
	FeatureID string `json:"featureId,required" format:"uuid"`
	// Whether the entitlement has a soft limit
	HasSoftLimit bool `json:"hasSoftLimit,required"`
	// Whether the entitlement has an unlimited usage
	HasUnlimitedUsage bool `json:"hasUnlimitedUsage,required"`
	// Whether the entitlement is visible
	IsVisible bool `json:"isVisible,required"`
	// The grant period of the promotional entitlement
	//
	// Any of "1 week", "1 month", "6 month", "1 year", "lifetime", "custom".
	Period string `json:"period,required"`
	// The reset period of the entitlement
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,required"`
	// The reset period configuration of the entitlement
	ResetPeriodConfiguration V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration,required"`
	// The start date of the entitlement
	StartDate time.Time `json:"startDate,required" format:"date-time"`
	// The status of the entitlement
	//
	// Any of "Active", "Expired", "Paused".
	Status string `json:"status,required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The usage limit of the entitlement
	UsageLimit float64 `json:"usageLimit,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		Description              respjson.Field
		EndDate                  respjson.Field
		EnumValues               respjson.Field
		EnvironmentID            respjson.Field
		FeatureGroupIDs          respjson.Field
		FeatureID                respjson.Field
		HasSoftLimit             respjson.Field
		HasUnlimitedUsage        respjson.Field
		IsVisible                respjson.Field
		Period                   respjson.Field
		ResetPeriod              respjson.Field
		ResetPeriodConfiguration respjson.Field
		StartDate                respjson.Field
		Status                   respjson.Field
		UpdatedAt                respjson.Field
		UsageLimit               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalEntitlementGrantResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPromotionalEntitlementGrantResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationUnion
// contains all possible properties and values from
// [V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementGrantResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1CustomerPromotionalEntitlementRevokeResponse struct {
	// Granted feature entitlement
	Data V1CustomerPromotionalEntitlementRevokeResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalEntitlementRevokeResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPromotionalEntitlementRevokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granted feature entitlement
type V1CustomerPromotionalEntitlementRevokeResponseData struct {
	// Unique identifier for the entity
	ID string `json:"id,required" format:"uuid"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// The description of the entitlement
	Description string `json:"description,required"`
	// The end date of the promotional entitlement
	EndDate time.Time `json:"endDate,required" format:"date-time"`
	// The enum values of the entitlement
	EnumValues []string `json:"enumValues,required"`
	// The unique identifier for the environment
	EnvironmentID string `json:"environmentId,required" format:"uuid"`
	// Feature group IDs associated with this entitlement
	FeatureGroupIDs []string `json:"featureGroupIds,required"`
	// The unique identifier of the entitlement feature
	FeatureID string `json:"featureId,required" format:"uuid"`
	// Whether the entitlement has a soft limit
	HasSoftLimit bool `json:"hasSoftLimit,required"`
	// Whether the entitlement has an unlimited usage
	HasUnlimitedUsage bool `json:"hasUnlimitedUsage,required"`
	// Whether the entitlement is visible
	IsVisible bool `json:"isVisible,required"`
	// The grant period of the promotional entitlement
	//
	// Any of "1 week", "1 month", "6 month", "1 year", "lifetime", "custom".
	Period string `json:"period,required"`
	// The reset period of the entitlement
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,required"`
	// The reset period configuration of the entitlement
	ResetPeriodConfiguration V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration,required"`
	// The start date of the entitlement
	StartDate time.Time `json:"startDate,required" format:"date-time"`
	// The status of the entitlement
	//
	// Any of "Active", "Expired", "Paused".
	Status string `json:"status,required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The usage limit of the entitlement
	UsageLimit float64 `json:"usageLimit,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		Description              respjson.Field
		EndDate                  respjson.Field
		EnumValues               respjson.Field
		EnvironmentID            respjson.Field
		FeatureGroupIDs          respjson.Field
		FeatureID                respjson.Field
		HasSoftLimit             respjson.Field
		HasUnlimitedUsage        respjson.Field
		IsVisible                respjson.Field
		Period                   respjson.Field
		ResetPeriod              respjson.Field
		ResetPeriodConfiguration respjson.Field
		StartDate                respjson.Field
		Status                   respjson.Field
		UpdatedAt                respjson.Field
		UsageLimit               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalEntitlementRevokeResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPromotionalEntitlementRevokeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationUnion
// contains all possible properties and values from
// [V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccordingTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPromotionalEntitlementGrantParams struct {
	// Promotional entitlements to grant
	PromotionalEntitlements []V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlement `json:"promotionalEntitlements,omitzero,required"`
	paramObj
}

func (r V1CustomerPromotionalEntitlementGrantParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalEntitlementGrantParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalEntitlementGrantParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single entitlement grant config
//
// The properties CustomEndDate, EnumValues, FeatureID, HasSoftLimit,
// HasUnlimitedUsage, IsVisible, MonthlyResetPeriodConfiguration, Period,
// ResetPeriod, UsageLimit, WeeklyResetPeriodConfiguration,
// YearlyResetPeriodConfiguration are required.
type V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlement struct {
	// The custom end date of the promotional entitlement
	CustomEndDate param.Opt[time.Time] `json:"customEndDate,omitzero,required" format:"date-time"`
	// Whether the entitlement has a soft limit
	HasSoftLimit param.Opt[bool] `json:"hasSoftLimit,omitzero,required"`
	// Whether the entitlement has an unlimited usage
	HasUnlimitedUsage param.Opt[bool] `json:"hasUnlimitedUsage,omitzero,required"`
	// Whether the entitlement is visible
	IsVisible param.Opt[bool] `json:"isVisible,omitzero,required"`
	// The usage limit of the entitlement
	UsageLimit param.Opt[int64] `json:"usageLimit,omitzero,required"`
	// The enum values of the entitlement
	EnumValues []string `json:"enumValues,omitzero,required"`
	// The monthly reset period configuration of the entitlement, defined when reset
	// period is monthly
	MonthlyResetPeriodConfiguration V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero,required"`
	// The reset period of the entitlement
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,omitzero,required"`
	// The weekly reset period configuration of the entitlement, defined when reset
	// period is weekly
	WeeklyResetPeriodConfiguration V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero,required"`
	// The yearly reset period configuration of the entitlement, defined when reset
	// period is yearly
	YearlyResetPeriodConfiguration V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero,required"`
	// The unique identifier of the entitlement feature
	FeatureID string `json:"featureId,required"`
	// The grant period of the promotional entitlement
	//
	// Any of "1 week", "1 month", "6 month", "1 year", "lifetime", "custom".
	Period string `json:"period,omitzero,required"`
	paramObj
}

func (r V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlement](
		"period", "1 week", "1 month", "6 month", "1 year", "lifetime", "custom",
	)
	apijson.RegisterFieldValidator[V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlement](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// The monthly reset period configuration of the entitlement, defined when reset
// period is monthly
//
// The property AccordingTo is required.
type V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero,required"`
	paramObj
}

func (r V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// The weekly reset period configuration of the entitlement, defined when reset
// period is weekly
//
// The property AccordingTo is required.
type V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero,required"`
	paramObj
}

func (r V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// The yearly reset period configuration of the entitlement, defined when reset
// period is yearly
//
// The property AccordingTo is required.
type V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero,required"`
	paramObj
}

func (r V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1CustomerPromotionalEntitlementRevokeParams struct {
	CustomerID string `path:"customerId,required" json:"-"`
	paramObj
}
