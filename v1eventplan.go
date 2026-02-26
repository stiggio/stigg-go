// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
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

// V1EventPlanService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventPlanService] method instead.
type V1EventPlanService struct {
	Options []option.RequestOption
}

// NewV1EventPlanService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1EventPlanService(opts ...option.RequestOption) (r V1EventPlanService) {
	r = V1EventPlanService{}
	r.Options = opts
	return
}

// Creates a new plan in draft status.
func (r *V1EventPlanService) New(ctx context.Context, body V1EventPlanNewParams, opts ...option.RequestOption) (res *V1EventPlanNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/plans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a plan by its unique identifier, including entitlements and pricing
// details.
func (r *V1EventPlanService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *V1EventPlanGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a paginated list of plans in the environment.
func (r *V1EventPlanService) List(ctx context.Context, query V1EventPlanListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1EventPlanListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/plans"
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

// Retrieves a paginated list of plans in the environment.
func (r *V1EventPlanService) ListAutoPaging(ctx context.Context, query V1EventPlanListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1EventPlanListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, query, opts...))
}

// Response object
type V1EventPlanNewResponse struct {
	// Plan configuration object
	Data V1EventPlanNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Plan configuration object
type V1EventPlanNewResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID          string   `json:"billingId" api:"required"`
	CompatibleAddonIDs []string `json:"compatibleAddonIds" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Default trial configuration for the plan
	DefaultTrialConfig V1EventPlanNewResponseDataDefaultTrialConfig `json:"defaultTrialConfig" api:"required"`
	// The description of the package
	Description string `json:"description" api:"required"`
	// The display name of the package
	DisplayName string `json:"displayName" api:"required"`
	// List of entitlements of the package
	Entitlements []V1EventPlanNewResponseDataEntitlement `json:"entitlements" api:"required"`
	// Indicates if the package is the latest version
	IsLatest bool `json:"isLatest" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The ID of the parent plan, if applicable
	ParentPlanID string `json:"parentPlanId" api:"required"`
	// The pricing type of the package
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType string `json:"pricingType" api:"required"`
	// The product id of the package
	ProductID string `json:"productId" api:"required"`
	// The status of the package
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The version number of the package
	VersionNumber int64 `json:"versionNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		BillingID          respjson.Field
		CompatibleAddonIDs respjson.Field
		CreatedAt          respjson.Field
		DefaultTrialConfig respjson.Field
		Description        respjson.Field
		DisplayName        respjson.Field
		Entitlements       respjson.Field
		IsLatest           respjson.Field
		Metadata           respjson.Field
		ParentPlanID       respjson.Field
		PricingType        respjson.Field
		ProductID          respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		VersionNumber      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default trial configuration for the plan
type V1EventPlanNewResponseDataDefaultTrialConfig struct {
	// The duration of the trial in the specified units
	Duration float64 `json:"duration" api:"required"`
	// The time unit for the trial duration (DAY or MONTH)
	//
	// Any of "DAY", "MONTH".
	Units string `json:"units" api:"required"`
	// Budget configuration for the trial
	Budget V1EventPlanNewResponseDataDefaultTrialConfigBudget `json:"budget" api:"nullable"`
	// Behavior when the trial ends (CONVERT_TO_PAID or CANCEL_SUBSCRIPTION)
	//
	// Any of "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION".
	TrialEndBehavior string `json:"trialEndBehavior" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration         respjson.Field
		Units            respjson.Field
		Budget           respjson.Field
		TrialEndBehavior respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanNewResponseDataDefaultTrialConfig) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanNewResponseDataDefaultTrialConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Budget configuration for the trial
type V1EventPlanNewResponseDataDefaultTrialConfigBudget struct {
	// Whether the budget limit is a soft limit (allows overage) or hard limit
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// The budget limit amount
	Limit float64 `json:"limit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasSoftLimit respjson.Field
		Limit        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanNewResponseDataDefaultTrialConfigBudget) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanNewResponseDataDefaultTrialConfigBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type V1EventPlanNewResponseDataEntitlement struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanNewResponseDataEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanNewResponseDataEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1EventPlanGetResponse struct {
	// Plan configuration object
	Data V1EventPlanGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Plan configuration object
type V1EventPlanGetResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID          string   `json:"billingId" api:"required"`
	CompatibleAddonIDs []string `json:"compatibleAddonIds" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Default trial configuration for the plan
	DefaultTrialConfig V1EventPlanGetResponseDataDefaultTrialConfig `json:"defaultTrialConfig" api:"required"`
	// The description of the package
	Description string `json:"description" api:"required"`
	// The display name of the package
	DisplayName string `json:"displayName" api:"required"`
	// List of entitlements of the package
	Entitlements []V1EventPlanGetResponseDataEntitlement `json:"entitlements" api:"required"`
	// Indicates if the package is the latest version
	IsLatest bool `json:"isLatest" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The ID of the parent plan, if applicable
	ParentPlanID string `json:"parentPlanId" api:"required"`
	// The pricing type of the package
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType string `json:"pricingType" api:"required"`
	// The product id of the package
	ProductID string `json:"productId" api:"required"`
	// The status of the package
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The version number of the package
	VersionNumber int64 `json:"versionNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		BillingID          respjson.Field
		CompatibleAddonIDs respjson.Field
		CreatedAt          respjson.Field
		DefaultTrialConfig respjson.Field
		Description        respjson.Field
		DisplayName        respjson.Field
		Entitlements       respjson.Field
		IsLatest           respjson.Field
		Metadata           respjson.Field
		ParentPlanID       respjson.Field
		PricingType        respjson.Field
		ProductID          respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		VersionNumber      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default trial configuration for the plan
type V1EventPlanGetResponseDataDefaultTrialConfig struct {
	// The duration of the trial in the specified units
	Duration float64 `json:"duration" api:"required"`
	// The time unit for the trial duration (DAY or MONTH)
	//
	// Any of "DAY", "MONTH".
	Units string `json:"units" api:"required"`
	// Budget configuration for the trial
	Budget V1EventPlanGetResponseDataDefaultTrialConfigBudget `json:"budget" api:"nullable"`
	// Behavior when the trial ends (CONVERT_TO_PAID or CANCEL_SUBSCRIPTION)
	//
	// Any of "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION".
	TrialEndBehavior string `json:"trialEndBehavior" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration         respjson.Field
		Units            respjson.Field
		Budget           respjson.Field
		TrialEndBehavior respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanGetResponseDataDefaultTrialConfig) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanGetResponseDataDefaultTrialConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Budget configuration for the trial
type V1EventPlanGetResponseDataDefaultTrialConfigBudget struct {
	// Whether the budget limit is a soft limit (allows overage) or hard limit
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// The budget limit amount
	Limit float64 `json:"limit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasSoftLimit respjson.Field
		Limit        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanGetResponseDataDefaultTrialConfigBudget) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanGetResponseDataDefaultTrialConfigBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type V1EventPlanGetResponseDataEntitlement struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanGetResponseDataEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanGetResponseDataEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Plan configuration object
type V1EventPlanListResponse struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID          string   `json:"billingId" api:"required"`
	CompatibleAddonIDs []string `json:"compatibleAddonIds" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Default trial configuration for the plan
	DefaultTrialConfig V1EventPlanListResponseDefaultTrialConfig `json:"defaultTrialConfig" api:"required"`
	// The description of the package
	Description string `json:"description" api:"required"`
	// The display name of the package
	DisplayName string `json:"displayName" api:"required"`
	// List of entitlements of the package
	Entitlements []V1EventPlanListResponseEntitlement `json:"entitlements" api:"required"`
	// Indicates if the package is the latest version
	IsLatest bool `json:"isLatest" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The ID of the parent plan, if applicable
	ParentPlanID string `json:"parentPlanId" api:"required"`
	// The pricing type of the package
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType V1EventPlanListResponsePricingType `json:"pricingType" api:"required"`
	// The product id of the package
	ProductID string `json:"productId" api:"required"`
	// The status of the package
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status V1EventPlanListResponseStatus `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The version number of the package
	VersionNumber int64 `json:"versionNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		BillingID          respjson.Field
		CompatibleAddonIDs respjson.Field
		CreatedAt          respjson.Field
		DefaultTrialConfig respjson.Field
		Description        respjson.Field
		DisplayName        respjson.Field
		Entitlements       respjson.Field
		IsLatest           respjson.Field
		Metadata           respjson.Field
		ParentPlanID       respjson.Field
		PricingType        respjson.Field
		ProductID          respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		VersionNumber      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default trial configuration for the plan
type V1EventPlanListResponseDefaultTrialConfig struct {
	// The duration of the trial in the specified units
	Duration float64 `json:"duration" api:"required"`
	// The time unit for the trial duration (DAY or MONTH)
	//
	// Any of "DAY", "MONTH".
	Units string `json:"units" api:"required"`
	// Budget configuration for the trial
	Budget V1EventPlanListResponseDefaultTrialConfigBudget `json:"budget" api:"nullable"`
	// Behavior when the trial ends (CONVERT_TO_PAID or CANCEL_SUBSCRIPTION)
	//
	// Any of "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION".
	TrialEndBehavior string `json:"trialEndBehavior" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration         respjson.Field
		Units            respjson.Field
		Budget           respjson.Field
		TrialEndBehavior respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanListResponseDefaultTrialConfig) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanListResponseDefaultTrialConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Budget configuration for the trial
type V1EventPlanListResponseDefaultTrialConfigBudget struct {
	// Whether the budget limit is a soft limit (allows overage) or hard limit
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// The budget limit amount
	Limit float64 `json:"limit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasSoftLimit respjson.Field
		Limit        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanListResponseDefaultTrialConfigBudget) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanListResponseDefaultTrialConfigBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type V1EventPlanListResponseEntitlement struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanListResponseEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanListResponseEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pricing type of the package
type V1EventPlanListResponsePricingType string

const (
	V1EventPlanListResponsePricingTypeFree   V1EventPlanListResponsePricingType = "FREE"
	V1EventPlanListResponsePricingTypePaid   V1EventPlanListResponsePricingType = "PAID"
	V1EventPlanListResponsePricingTypeCustom V1EventPlanListResponsePricingType = "CUSTOM"
)

// The status of the package
type V1EventPlanListResponseStatus string

const (
	V1EventPlanListResponseStatusDraft     V1EventPlanListResponseStatus = "DRAFT"
	V1EventPlanListResponseStatusPublished V1EventPlanListResponseStatus = "PUBLISHED"
	V1EventPlanListResponseStatusArchived  V1EventPlanListResponseStatus = "ARCHIVED"
)

type V1EventPlanNewParams struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// The display name of the package
	DisplayName string `json:"displayName" api:"required"`
	// The product ID to associate the plan with
	ProductID string `json:"productId" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// The description of the package
	Description param.Opt[string] `json:"description,omitzero"`
	// The ID of the parent plan, if applicable
	ParentPlanID param.Opt[string] `json:"parentPlanId,omitzero"`
	// Default trial configuration for the plan
	DefaultTrialConfig V1EventPlanNewParamsDefaultTrialConfig `json:"defaultTrialConfig,omitzero"`
	// The pricing type of the package
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType V1EventPlanNewParamsPricingType `json:"pricingType,omitzero"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,omitzero"`
	// The status of the package
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status V1EventPlanNewParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r V1EventPlanNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default trial configuration for the plan
//
// The properties Duration, Units are required.
type V1EventPlanNewParamsDefaultTrialConfig struct {
	// The duration of the trial in the specified units
	Duration float64 `json:"duration" api:"required"`
	// The time unit for the trial duration (DAY or MONTH)
	//
	// Any of "DAY", "MONTH".
	Units string `json:"units,omitzero" api:"required"`
	// Budget configuration for the trial
	Budget V1EventPlanNewParamsDefaultTrialConfigBudget `json:"budget,omitzero"`
	// Behavior when the trial ends (CONVERT_TO_PAID or CANCEL_SUBSCRIPTION)
	//
	// Any of "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION".
	TrialEndBehavior string `json:"trialEndBehavior,omitzero"`
	paramObj
}

func (r V1EventPlanNewParamsDefaultTrialConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanNewParamsDefaultTrialConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanNewParamsDefaultTrialConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventPlanNewParamsDefaultTrialConfig](
		"units", "DAY", "MONTH",
	)
	apijson.RegisterFieldValidator[V1EventPlanNewParamsDefaultTrialConfig](
		"trialEndBehavior", "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION",
	)
}

// Budget configuration for the trial
//
// The properties HasSoftLimit, Limit are required.
type V1EventPlanNewParamsDefaultTrialConfigBudget struct {
	// Whether the budget limit is a soft limit (allows overage) or hard limit
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// The budget limit amount
	Limit float64 `json:"limit" api:"required"`
	paramObj
}

func (r V1EventPlanNewParamsDefaultTrialConfigBudget) MarshalJSON() (data []byte, err error) {
	type shadow V1EventPlanNewParamsDefaultTrialConfigBudget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventPlanNewParamsDefaultTrialConfigBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pricing type of the package
type V1EventPlanNewParamsPricingType string

const (
	V1EventPlanNewParamsPricingTypeFree   V1EventPlanNewParamsPricingType = "FREE"
	V1EventPlanNewParamsPricingTypePaid   V1EventPlanNewParamsPricingType = "PAID"
	V1EventPlanNewParamsPricingTypeCustom V1EventPlanNewParamsPricingType = "CUSTOM"
)

// The status of the package
type V1EventPlanNewParamsStatus string

const (
	V1EventPlanNewParamsStatusDraft     V1EventPlanNewParamsStatus = "DRAFT"
	V1EventPlanNewParamsStatusPublished V1EventPlanNewParamsStatus = "PUBLISHED"
	V1EventPlanNewParamsStatusArchived  V1EventPlanNewParamsStatus = "ARCHIVED"
)

type V1EventPlanListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by product ID
	ProductID param.Opt[string] `query:"productId,omitzero" json:"-"`
	// Filter by status. Supports comma-separated values for multiple statuses
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter by creation date using range operators: gt, gte, lt, lte
	CreatedAt V1EventPlanListParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventPlanListParams]'s query parameters as `url.Values`.
func (r V1EventPlanListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by creation date using range operators: gt, gte, lt, lte
type V1EventPlanListParamsCreatedAt struct {
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

// URLQuery serializes [V1EventPlanListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r V1EventPlanListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
