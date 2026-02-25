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
	"github.com/stiggio/stigg-go/packages/pagination"
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

// Grants promotional entitlements to a customer, providing feature access outside
// their subscription. Entitlements can be time-limited or permanent.
func (r *V1CustomerPromotionalEntitlementService) New(ctx context.Context, id string, body V1CustomerPromotionalEntitlementNewParams, opts ...option.RequestOption) (res *V1CustomerPromotionalEntitlementNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/promotional-entitlements", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of a customer's promotional entitlements.
func (r *V1CustomerPromotionalEntitlementService) List(ctx context.Context, id string, query V1CustomerPromotionalEntitlementListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1CustomerPromotionalEntitlementListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/promotional-entitlements", id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieves a paginated list of a customer's promotional entitlements.
func (r *V1CustomerPromotionalEntitlementService) ListAutoPaging(ctx context.Context, id string, query V1CustomerPromotionalEntitlementListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1CustomerPromotionalEntitlementListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, id, query, opts...))
}

// Revokes a previously granted promotional entitlement from a customer for a
// specific feature.
func (r *V1CustomerPromotionalEntitlementService) Revoke(ctx context.Context, featureID string, body V1CustomerPromotionalEntitlementRevokeParams, opts ...option.RequestOption) (res *V1CustomerPromotionalEntitlementRevokeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if featureID == "" {
		err = errors.New("missing required featureId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/promotional-entitlements/%s", body.ID, featureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Response object
type V1CustomerPromotionalEntitlementNewResponse struct {
	Data []V1CustomerPromotionalEntitlementNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerPromotionalEntitlementNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPromotionalEntitlementNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granted feature entitlement
type V1CustomerPromotionalEntitlementNewResponseData struct {
	// Unique identifier for the entity
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The description of the entitlement
	Description string `json:"description" api:"required"`
	// The end date of the promotional entitlement
	EndDate time.Time `json:"endDate" api:"required" format:"date-time"`
	// The enum values of the entitlement
	EnumValues []string `json:"enumValues" api:"required"`
	// The unique identifier for the environment
	EnvironmentID string `json:"environmentId" api:"required" format:"uuid"`
	// Feature group IDs associated with this entitlement
	FeatureGroupIDs []string `json:"featureGroupIds" api:"required"`
	// The unique identifier of the entitlement feature
	FeatureID string `json:"featureId" api:"required" format:"uuid"`
	// Whether the entitlement has a soft limit
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// Whether the entitlement has an unlimited usage
	HasUnlimitedUsage bool `json:"hasUnlimitedUsage" api:"required"`
	// Whether the entitlement is visible
	IsVisible bool `json:"isVisible" api:"required"`
	// The grant period of the promotional entitlement
	//
	// Any of "1 week", "1 month", "6 month", "1 year", "lifetime", "custom".
	Period string `json:"period" api:"required"`
	// The reset period of the entitlement
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod" api:"required"`
	// The reset period configuration of the entitlement
	ResetPeriodConfiguration V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
	// The start date of the entitlement
	StartDate time.Time `json:"startDate" api:"required" format:"date-time"`
	// The status of the entitlement
	//
	// Any of "Active", "Expired", "Paused".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The usage limit of the entitlement
	UsageLimit float64 `json:"usageLimit" api:"required"`
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
func (r V1CustomerPromotionalEntitlementNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPromotionalEntitlementNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationUnion
// contains all possible properties and values from
// [V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementNewResponseDataResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granted feature entitlement
type V1CustomerPromotionalEntitlementListResponse struct {
	// Unique identifier for the entity
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The description of the entitlement
	Description string `json:"description" api:"required"`
	// The end date of the promotional entitlement
	EndDate time.Time `json:"endDate" api:"required" format:"date-time"`
	// The enum values of the entitlement
	EnumValues []string `json:"enumValues" api:"required"`
	// The unique identifier for the environment
	EnvironmentID string `json:"environmentId" api:"required" format:"uuid"`
	// Feature group IDs associated with this entitlement
	FeatureGroupIDs []string `json:"featureGroupIds" api:"required"`
	// The unique identifier of the entitlement feature
	FeatureID string `json:"featureId" api:"required" format:"uuid"`
	// Whether the entitlement has a soft limit
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// Whether the entitlement has an unlimited usage
	HasUnlimitedUsage bool `json:"hasUnlimitedUsage" api:"required"`
	// Whether the entitlement is visible
	IsVisible bool `json:"isVisible" api:"required"`
	// The grant period of the promotional entitlement
	//
	// Any of "1 week", "1 month", "6 month", "1 year", "lifetime", "custom".
	Period V1CustomerPromotionalEntitlementListResponsePeriod `json:"period" api:"required"`
	// The reset period of the entitlement
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod V1CustomerPromotionalEntitlementListResponseResetPeriod `json:"resetPeriod" api:"required"`
	// The reset period configuration of the entitlement
	ResetPeriodConfiguration V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
	// The start date of the entitlement
	StartDate time.Time `json:"startDate" api:"required" format:"date-time"`
	// The status of the entitlement
	//
	// Any of "Active", "Expired", "Paused".
	Status V1CustomerPromotionalEntitlementListResponseStatus `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The usage limit of the entitlement
	UsageLimit float64 `json:"usageLimit" api:"required"`
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
func (r V1CustomerPromotionalEntitlementListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerPromotionalEntitlementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The grant period of the promotional entitlement
type V1CustomerPromotionalEntitlementListResponsePeriod string

const (
	V1CustomerPromotionalEntitlementListResponsePeriod1Week    V1CustomerPromotionalEntitlementListResponsePeriod = "1 week"
	V1CustomerPromotionalEntitlementListResponsePeriod1Month   V1CustomerPromotionalEntitlementListResponsePeriod = "1 month"
	V1CustomerPromotionalEntitlementListResponsePeriod6Month   V1CustomerPromotionalEntitlementListResponsePeriod = "6 month"
	V1CustomerPromotionalEntitlementListResponsePeriod1Year    V1CustomerPromotionalEntitlementListResponsePeriod = "1 year"
	V1CustomerPromotionalEntitlementListResponsePeriodLifetime V1CustomerPromotionalEntitlementListResponsePeriod = "lifetime"
	V1CustomerPromotionalEntitlementListResponsePeriodCustom   V1CustomerPromotionalEntitlementListResponsePeriod = "custom"
)

// The reset period of the entitlement
type V1CustomerPromotionalEntitlementListResponseResetPeriod string

const (
	V1CustomerPromotionalEntitlementListResponseResetPeriodYear  V1CustomerPromotionalEntitlementListResponseResetPeriod = "YEAR"
	V1CustomerPromotionalEntitlementListResponseResetPeriodMonth V1CustomerPromotionalEntitlementListResponseResetPeriod = "MONTH"
	V1CustomerPromotionalEntitlementListResponseResetPeriodWeek  V1CustomerPromotionalEntitlementListResponseResetPeriod = "WEEK"
	V1CustomerPromotionalEntitlementListResponseResetPeriodDay   V1CustomerPromotionalEntitlementListResponseResetPeriod = "DAY"
	V1CustomerPromotionalEntitlementListResponseResetPeriodHour  V1CustomerPromotionalEntitlementListResponseResetPeriod = "HOUR"
)

// V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationUnion
// contains all possible properties and values from
// [V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationYearlyResetPeriodConfig],
// [V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationMonthlyResetPeriodConfig],
// [V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationWeeklyResetPeriodConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationUnion struct {
	AccordingTo string `json:"accordingTo"`
	JSON        struct {
		AccordingTo respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationUnion) AsYearlyResetPeriodConfig() (v V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationYearlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationUnion) AsMonthlyResetPeriodConfig() (v V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationMonthlyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationUnion) AsWeeklyResetPeriodConfig() (v V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationWeeklyResetPeriodConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Yearly reset configuration
type V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationYearlyResetPeriodConfig struct {
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
func (r V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationYearlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationYearlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly reset configuration
type V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationMonthlyResetPeriodConfig struct {
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
func (r V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationMonthlyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationMonthlyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weekly reset configuration
type V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationWeeklyResetPeriodConfig struct {
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
func (r V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationWeeklyResetPeriodConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerPromotionalEntitlementListResponseResetPeriodConfigurationWeeklyResetPeriodConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the entitlement
type V1CustomerPromotionalEntitlementListResponseStatus string

const (
	V1CustomerPromotionalEntitlementListResponseStatusActive  V1CustomerPromotionalEntitlementListResponseStatus = "Active"
	V1CustomerPromotionalEntitlementListResponseStatusExpired V1CustomerPromotionalEntitlementListResponseStatus = "Expired"
	V1CustomerPromotionalEntitlementListResponseStatusPaused  V1CustomerPromotionalEntitlementListResponseStatus = "Paused"
)

// Response object
type V1CustomerPromotionalEntitlementRevokeResponse struct {
	// Granted feature entitlement
	Data V1CustomerPromotionalEntitlementRevokeResponseData `json:"data" api:"required"`
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
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The description of the entitlement
	Description string `json:"description" api:"required"`
	// The end date of the promotional entitlement
	EndDate time.Time `json:"endDate" api:"required" format:"date-time"`
	// The enum values of the entitlement
	EnumValues []string `json:"enumValues" api:"required"`
	// The unique identifier for the environment
	EnvironmentID string `json:"environmentId" api:"required" format:"uuid"`
	// Feature group IDs associated with this entitlement
	FeatureGroupIDs []string `json:"featureGroupIds" api:"required"`
	// The unique identifier of the entitlement feature
	FeatureID string `json:"featureId" api:"required" format:"uuid"`
	// Whether the entitlement has a soft limit
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// Whether the entitlement has an unlimited usage
	HasUnlimitedUsage bool `json:"hasUnlimitedUsage" api:"required"`
	// Whether the entitlement is visible
	IsVisible bool `json:"isVisible" api:"required"`
	// The grant period of the promotional entitlement
	//
	// Any of "1 week", "1 month", "6 month", "1 year", "lifetime", "custom".
	Period string `json:"period" api:"required"`
	// The reset period of the entitlement
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod" api:"required"`
	// The reset period configuration of the entitlement
	ResetPeriodConfiguration V1CustomerPromotionalEntitlementRevokeResponseDataResetPeriodConfigurationUnion `json:"resetPeriodConfiguration" api:"required"`
	// The start date of the entitlement
	StartDate time.Time `json:"startDate" api:"required" format:"date-time"`
	// The status of the entitlement
	//
	// Any of "Active", "Expired", "Paused".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The usage limit of the entitlement
	UsageLimit float64 `json:"usageLimit" api:"required"`
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
	AccordingTo string `json:"accordingTo" api:"required"`
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
	AccordingTo string `json:"accordingTo" api:"required"`
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
	AccordingTo string `json:"accordingTo" api:"required"`
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

type V1CustomerPromotionalEntitlementNewParams struct {
	// Promotional entitlements to grant
	PromotionalEntitlements []V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlement `json:"promotionalEntitlements,omitzero" api:"required"`
	paramObj
}

func (r V1CustomerPromotionalEntitlementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalEntitlementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalEntitlementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single entitlement grant config
//
// The properties CustomEndDate, EnumValues, FeatureID, HasSoftLimit,
// HasUnlimitedUsage, IsVisible, MonthlyResetPeriodConfiguration, Period,
// ResetPeriod, UsageLimit, WeeklyResetPeriodConfiguration,
// YearlyResetPeriodConfiguration are required.
type V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlement struct {
	// The custom end date of the promotional entitlement
	CustomEndDate param.Opt[time.Time] `json:"customEndDate,omitzero" api:"required" format:"date-time"`
	// Whether the entitlement has a soft limit
	HasSoftLimit param.Opt[bool] `json:"hasSoftLimit,omitzero" api:"required"`
	// Whether the entitlement has an unlimited usage
	HasUnlimitedUsage param.Opt[bool] `json:"hasUnlimitedUsage,omitzero" api:"required"`
	// Whether the entitlement is visible
	IsVisible param.Opt[bool] `json:"isVisible,omitzero" api:"required"`
	// The usage limit of the entitlement
	UsageLimit param.Opt[int64] `json:"usageLimit,omitzero" api:"required"`
	// The enum values of the entitlement
	EnumValues []string `json:"enumValues,omitzero" api:"required"`
	// The monthly reset period configuration of the entitlement, defined when reset
	// period is monthly
	MonthlyResetPeriodConfiguration V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero" api:"required"`
	// The reset period of the entitlement
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,omitzero" api:"required"`
	// The weekly reset period configuration of the entitlement, defined when reset
	// period is weekly
	WeeklyResetPeriodConfiguration V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero" api:"required"`
	// The yearly reset period configuration of the entitlement, defined when reset
	// period is yearly
	YearlyResetPeriodConfiguration V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero" api:"required"`
	// The unique identifier of the entitlement feature
	FeatureID string `json:"featureId" api:"required"`
	// The grant period of the promotional entitlement
	//
	// Any of "1 week", "1 month", "6 month", "1 year", "lifetime", "custom".
	Period string `json:"period,omitzero" api:"required"`
	paramObj
}

func (r V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlement](
		"period", "1 week", "1 month", "6 month", "1 year", "lifetime", "custom",
	)
	apijson.RegisterFieldValidator[V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlement](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// The monthly reset period configuration of the entitlement, defined when reset
// period is monthly
//
// The property AccordingTo is required.
type V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// The weekly reset period configuration of the entitlement, defined when reset
// period is weekly
//
// The property AccordingTo is required.
type V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// The yearly reset period configuration of the entitlement, defined when reset
// period is yearly
//
// The property AccordingTo is required.
type V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1CustomerPromotionalEntitlementListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by promotional entitlement status. Supports comma-separated values for
	// multiple statuses
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter by creation date using range operators: gt, gte, lt, lte
	CreatedAt V1CustomerPromotionalEntitlementListParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerPromotionalEntitlementListParams]'s query
// parameters as `url.Values`.
func (r V1CustomerPromotionalEntitlementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by creation date using range operators: gt, gte, lt, lte
type V1CustomerPromotionalEntitlementListParamsCreatedAt struct {
	// Greater than the specified createdAt value
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// Greater than or equal to the specified createdAt value
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Less than the specified createdAt value
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	// Less than or equal to the specified createdAt value
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerPromotionalEntitlementListParamsCreatedAt]'s
// query parameters as `url.Values`.
func (r V1CustomerPromotionalEntitlementListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerPromotionalEntitlementRevokeParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
