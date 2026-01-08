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

	"github.com/stainless-sdks/stigg-go/internal/apijson"
	"github.com/stainless-sdks/stigg-go/internal/requestconfig"
	"github.com/stainless-sdks/stigg-go/option"
	"github.com/stainless-sdks/stigg-go/packages/param"
	"github.com/stainless-sdks/stigg-go/packages/respjson"
)

// V1CustomerPromotionalService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerPromotionalService] method instead.
type V1CustomerPromotionalService struct {
	Options []option.RequestOption
}

// NewV1CustomerPromotionalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerPromotionalService(opts ...option.RequestOption) (r V1CustomerPromotionalService) {
	r = V1CustomerPromotionalService{}
	r.Options = opts
	return
}

// Create a new Promotional Entitlements
func (r *V1CustomerPromotionalService) New(ctx context.Context, customerID string, body V1CustomerPromotionalNewParams, opts ...option.RequestOption) (res *V1CustomerPromotionalNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/promotional", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform revocation on a Promotional Entitlement
func (r *V1CustomerPromotionalService) Revoke(ctx context.Context, featureID string, body V1CustomerPromotionalRevokeParams, opts ...option.RequestOption) (res *V1CustomerPromotionalRevokeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.CustomerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	if featureID == "" {
		err = errors.New("missing required featureId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/promotional/featureId/%s", body.CustomerID, featureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type V1CustomerPromotionalNewResponse struct {
	Data []V1CustomerPromotionalNewResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPromotionalNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPromotionalNewResponseData struct {
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
	ResetPeriodConfiguration V1CustomerPromotionalNewResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration,required"`
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
func (r V1CustomerPromotionalNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPromotionalNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1CustomerPromotionalNewResponseDataResetPeriodConfigurationUnion contains all
// possible properties and values from
// [V1CustomerPromotionalNewResponseDataResetPeriodConfigurationAccordingTo],
// [V1CustomerPromotionalNewResponseDataResetPeriodConfigurationAccordingTo],
// [V1CustomerPromotionalNewResponseDataResetPeriodConfigurationAccordingTo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerPromotionalNewResponseDataResetPeriodConfigurationUnion struct {
	// This field is from variant
	// [V1CustomerPromotionalNewResponseDataResetPeriodConfigurationAccordingTo].
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1CustomerPromotionalNewResponseDataResetPeriodConfigurationUnion) AsV1CustomerPromotionalNewResponseDataResetPeriodConfigurationAccordingTo() (v V1CustomerPromotionalNewResponseDataResetPeriodConfigurationAccordingTo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalNewResponseDataResetPeriodConfigurationUnion) AsVariant2() (v V1CustomerPromotionalNewResponseDataResetPeriodConfigurationAccordingTo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalNewResponseDataResetPeriodConfigurationUnion) AsVariant3() (v V1CustomerPromotionalNewResponseDataResetPeriodConfigurationAccordingTo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerPromotionalNewResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1CustomerPromotionalNewResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPromotionalNewResponseDataResetPeriodConfigurationAccordingTo struct {
	// Yearly reset period according to
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
func (r V1CustomerPromotionalNewResponseDataResetPeriodConfigurationAccordingTo) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalNewResponseDataResetPeriodConfigurationAccordingTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPromotionalRevokeResponse struct {
	Data V1CustomerPromotionalRevokeResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalRevokeResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPromotionalRevokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPromotionalRevokeResponseData struct {
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
	ResetPeriodConfiguration V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration,required"`
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
func (r V1CustomerPromotionalRevokeResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPromotionalRevokeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationUnion contains
// all possible properties and values from
// [V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationAccordingTo],
// [V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationAccordingTo],
// [V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationAccordingTo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationUnion struct {
	// This field is from variant
	// [V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationAccordingTo].
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationUnion) AsV1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationAccordingTo() (v V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationAccordingTo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationUnion) AsVariant2() (v V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationAccordingTo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationUnion) AsVariant3() (v V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationAccordingTo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationAccordingTo struct {
	// Yearly reset period according to
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
func (r V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationAccordingTo) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalRevokeResponseDataResetPeriodConfigurationAccordingTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerPromotionalNewParams struct {
	// Promotional entitlements to grant
	PromotionalEntitlements []V1CustomerPromotionalNewParamsPromotionalEntitlement `json:"promotionalEntitlements,omitzero,required"`
	paramObj
}

func (r V1CustomerPromotionalNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CustomEndDate, EnumValues, FeatureID, HasSoftLimit,
// HasUnlimitedUsage, IsVisible, MonthlyResetPeriodConfiguration, Period,
// ResetPeriod, UsageLimit, WeeklyResetPeriodConfiguration,
// YearlyResetPeriodConfiguration are required.
type V1CustomerPromotionalNewParamsPromotionalEntitlement struct {
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
	MonthlyResetPeriodConfiguration V1CustomerPromotionalNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero,required"`
	// The reset period of the entitlement
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,omitzero,required"`
	// The weekly reset period configuration of the entitlement, defined when reset
	// period is weekly
	WeeklyResetPeriodConfiguration V1CustomerPromotionalNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero,required"`
	// The yearly reset period configuration of the entitlement, defined when reset
	// period is yearly
	YearlyResetPeriodConfiguration V1CustomerPromotionalNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero,required"`
	// The unique identifier of the entitlement feature
	FeatureID string `json:"featureId,required"`
	// The grant period of the promotional entitlement
	//
	// Any of "1 week", "1 month", "6 month", "1 year", "lifetime", "custom".
	Period string `json:"period,omitzero,required"`
	paramObj
}

func (r V1CustomerPromotionalNewParamsPromotionalEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalNewParamsPromotionalEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalNewParamsPromotionalEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalNewParamsPromotionalEntitlement](
		"period", "1 week", "1 month", "6 month", "1 year", "lifetime", "custom",
	)
	apijson.RegisterFieldValidator[V1CustomerPromotionalNewParamsPromotionalEntitlement](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// The monthly reset period configuration of the entitlement, defined when reset
// period is monthly
//
// The property AccordingTo is required.
type V1CustomerPromotionalNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration struct {
	// Monthly reset period according to
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero,required"`
	paramObj
}

func (r V1CustomerPromotionalNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// The weekly reset period configuration of the entitlement, defined when reset
// period is weekly
//
// The property AccordingTo is required.
type V1CustomerPromotionalNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration struct {
	// Weekly reset period according to
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero,required"`
	paramObj
}

func (r V1CustomerPromotionalNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// The yearly reset period configuration of the entitlement, defined when reset
// period is yearly
//
// The property AccordingTo is required.
type V1CustomerPromotionalNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration struct {
	// Yearly reset period according to
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero,required"`
	paramObj
}

func (r V1CustomerPromotionalNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1CustomerPromotionalRevokeParams struct {
	CustomerID string `path:"customerId,required" json:"-"`
	paramObj
}
