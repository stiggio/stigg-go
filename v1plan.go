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
	shimjson "github.com/stiggio/stigg-go/internal/encoding/json"
	"github.com/stiggio/stigg-go/internal/requestconfig"
	"github.com/stiggio/stigg-go/option"
	"github.com/stiggio/stigg-go/packages/pagination"
	"github.com/stiggio/stigg-go/packages/param"
	"github.com/stiggio/stigg-go/packages/respjson"
)

// V1PlanService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PlanService] method instead.
type V1PlanService struct {
	Options      []option.RequestOption
	Entitlements V1PlanEntitlementService
}

// NewV1PlanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1PlanService(opts ...option.RequestOption) (r V1PlanService) {
	r = V1PlanService{}
	r.Options = opts
	r.Entitlements = NewV1PlanEntitlementService(opts...)
	return
}

// Creates a new plan in draft status.
func (r *V1PlanService) New(ctx context.Context, body V1PlanNewParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/plans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a plan by its unique identifier, including entitlements and pricing
// details.
func (r *V1PlanService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing plan's properties such as display name, description, and
// metadata.
func (r *V1PlanService) Update(ctx context.Context, id string, body V1PlanUpdateParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of plans in the environment.
func (r *V1PlanService) List(ctx context.Context, query V1PlanListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1PlanListResponse], err error) {
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
func (r *V1PlanService) ListAutoPaging(ctx context.Context, query V1PlanListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1PlanListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, query, opts...))
}

// Archives a plan, preventing it from being used in new subscriptions.
func (r *V1PlanService) Archive(ctx context.Context, id string, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Creates a draft version of an existing plan for modification before publishing.
func (r *V1PlanService) NewDraft(ctx context.Context, id string, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s/draft", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Publishes a draft plan, making it available for use in subscriptions.
func (r *V1PlanService) Publish(ctx context.Context, id string, body V1PlanPublishParams, opts ...option.RequestOption) (res *V1PlanPublishResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s/publish", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes a draft version of a plan.
func (r *V1PlanService) RemoveDraft(ctx context.Context, id string, opts ...option.RequestOption) (res *V1PlanRemoveDraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s/draft", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Sets the pricing configuration for a plan, including pricing models, overage
// pricing, and minimum spend.
func (r *V1PlanService) SetPricing(ctx context.Context, id string, body V1PlanSetPricingParams, opts ...option.RequestOption) (res *SetPackagePricingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s/charges", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Response object
type Plan struct {
	// Plan configuration object
	Data PlanData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Plan) RawJSON() string { return r.JSON.raw }
func (r *Plan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Plan configuration object
type PlanData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID          string   `json:"billingId" api:"required"`
	CompatibleAddonIDs []string `json:"compatibleAddonIds" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Default trial configuration for the plan
	DefaultTrialConfig PlanDataDefaultTrialConfig `json:"defaultTrialConfig" api:"required"`
	// The description of the package
	Description string `json:"description" api:"required"`
	// The display name of the package
	DisplayName string `json:"displayName" api:"required"`
	// List of entitlements of the package
	Entitlements []PlanDataEntitlement `json:"entitlements" api:"required"`
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
func (r PlanData) RawJSON() string { return r.JSON.raw }
func (r *PlanData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default trial configuration for the plan
type PlanDataDefaultTrialConfig struct {
	// The duration of the trial in the specified units
	Duration float64 `json:"duration" api:"required"`
	// The time unit for the trial duration (DAY or MONTH)
	//
	// Any of "DAY", "MONTH".
	Units string `json:"units" api:"required"`
	// Budget configuration for the trial
	Budget PlanDataDefaultTrialConfigBudget `json:"budget" api:"nullable"`
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
func (r PlanDataDefaultTrialConfig) RawJSON() string { return r.JSON.raw }
func (r *PlanDataDefaultTrialConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Budget configuration for the trial
type PlanDataDefaultTrialConfigBudget struct {
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
func (r PlanDataDefaultTrialConfigBudget) RawJSON() string { return r.JSON.raw }
func (r *PlanDataDefaultTrialConfigBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type PlanDataEntitlement struct {
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
func (r PlanDataEntitlement) RawJSON() string { return r.JSON.raw }
func (r *PlanDataEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Plan configuration object
type V1PlanListResponse struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID          string   `json:"billingId" api:"required"`
	CompatibleAddonIDs []string `json:"compatibleAddonIds" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Default trial configuration for the plan
	DefaultTrialConfig V1PlanListResponseDefaultTrialConfig `json:"defaultTrialConfig" api:"required"`
	// The description of the package
	Description string `json:"description" api:"required"`
	// The display name of the package
	DisplayName string `json:"displayName" api:"required"`
	// List of entitlements of the package
	Entitlements []V1PlanListResponseEntitlement `json:"entitlements" api:"required"`
	// Indicates if the package is the latest version
	IsLatest bool `json:"isLatest" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The ID of the parent plan, if applicable
	ParentPlanID string `json:"parentPlanId" api:"required"`
	// The pricing type of the package
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType V1PlanListResponsePricingType `json:"pricingType" api:"required"`
	// The product id of the package
	ProductID string `json:"productId" api:"required"`
	// The status of the package
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status V1PlanListResponseStatus `json:"status" api:"required"`
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
func (r V1PlanListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default trial configuration for the plan
type V1PlanListResponseDefaultTrialConfig struct {
	// The duration of the trial in the specified units
	Duration float64 `json:"duration" api:"required"`
	// The time unit for the trial duration (DAY or MONTH)
	//
	// Any of "DAY", "MONTH".
	Units string `json:"units" api:"required"`
	// Budget configuration for the trial
	Budget V1PlanListResponseDefaultTrialConfigBudget `json:"budget" api:"nullable"`
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
func (r V1PlanListResponseDefaultTrialConfig) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListResponseDefaultTrialConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Budget configuration for the trial
type V1PlanListResponseDefaultTrialConfigBudget struct {
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
func (r V1PlanListResponseDefaultTrialConfigBudget) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListResponseDefaultTrialConfigBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type V1PlanListResponseEntitlement struct {
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
func (r V1PlanListResponseEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListResponseEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pricing type of the package
type V1PlanListResponsePricingType string

const (
	V1PlanListResponsePricingTypeFree   V1PlanListResponsePricingType = "FREE"
	V1PlanListResponsePricingTypePaid   V1PlanListResponsePricingType = "PAID"
	V1PlanListResponsePricingTypeCustom V1PlanListResponsePricingType = "CUSTOM"
)

// The status of the package
type V1PlanListResponseStatus string

const (
	V1PlanListResponseStatusDraft     V1PlanListResponseStatus = "DRAFT"
	V1PlanListResponseStatusPublished V1PlanListResponseStatus = "PUBLISHED"
	V1PlanListResponseStatusArchived  V1PlanListResponseStatus = "ARCHIVED"
)

// Response containing task ID for publish operation
type V1PlanPublishResponse struct {
	Data V1PlanPublishResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanPublishResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PlanPublishResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanPublishResponseData struct {
	// Task ID for tracking the async publish operation
	TaskID string `json:"taskId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TaskID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanPublishResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1PlanPublishResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response confirming the plan draft was removed.
type V1PlanRemoveDraftResponse struct {
	Data V1PlanRemoveDraftResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanRemoveDraftResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PlanRemoveDraftResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanRemoveDraftResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanRemoveDraftResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1PlanRemoveDraftResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanNewParams struct {
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
	DefaultTrialConfig V1PlanNewParamsDefaultTrialConfig `json:"defaultTrialConfig,omitzero"`
	// The pricing type of the package
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType V1PlanNewParamsPricingType `json:"pricingType,omitzero"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,omitzero"`
	// The status of the package
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status V1PlanNewParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r V1PlanNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default trial configuration for the plan
//
// The properties Duration, Units are required.
type V1PlanNewParamsDefaultTrialConfig struct {
	// The duration of the trial in the specified units
	Duration float64 `json:"duration" api:"required"`
	// The time unit for the trial duration (DAY or MONTH)
	//
	// Any of "DAY", "MONTH".
	Units string `json:"units,omitzero" api:"required"`
	// Budget configuration for the trial
	Budget V1PlanNewParamsDefaultTrialConfigBudget `json:"budget,omitzero"`
	// Behavior when the trial ends (CONVERT_TO_PAID or CANCEL_SUBSCRIPTION)
	//
	// Any of "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION".
	TrialEndBehavior string `json:"trialEndBehavior,omitzero"`
	paramObj
}

func (r V1PlanNewParamsDefaultTrialConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanNewParamsDefaultTrialConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanNewParamsDefaultTrialConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanNewParamsDefaultTrialConfig](
		"units", "DAY", "MONTH",
	)
	apijson.RegisterFieldValidator[V1PlanNewParamsDefaultTrialConfig](
		"trialEndBehavior", "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION",
	)
}

// Budget configuration for the trial
//
// The properties HasSoftLimit, Limit are required.
type V1PlanNewParamsDefaultTrialConfigBudget struct {
	// Whether the budget limit is a soft limit (allows overage) or hard limit
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// The budget limit amount
	Limit float64 `json:"limit" api:"required"`
	paramObj
}

func (r V1PlanNewParamsDefaultTrialConfigBudget) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanNewParamsDefaultTrialConfigBudget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanNewParamsDefaultTrialConfigBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pricing type of the package
type V1PlanNewParamsPricingType string

const (
	V1PlanNewParamsPricingTypeFree   V1PlanNewParamsPricingType = "FREE"
	V1PlanNewParamsPricingTypePaid   V1PlanNewParamsPricingType = "PAID"
	V1PlanNewParamsPricingTypeCustom V1PlanNewParamsPricingType = "CUSTOM"
)

// The status of the package
type V1PlanNewParamsStatus string

const (
	V1PlanNewParamsStatusDraft     V1PlanNewParamsStatus = "DRAFT"
	V1PlanNewParamsStatusPublished V1PlanNewParamsStatus = "PUBLISHED"
	V1PlanNewParamsStatusArchived  V1PlanNewParamsStatus = "ARCHIVED"
)

type V1PlanUpdateParams struct {
	// The unique identifier for the entity in the billing provider
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// The description of the package
	Description param.Opt[string] `json:"description,omitzero"`
	// The ID of the parent plan, if applicable
	ParentPlanID param.Opt[string] `json:"parentPlanId,omitzero"`
	// The display name of the package
	DisplayName        param.Opt[string] `json:"displayName,omitzero"`
	CompatibleAddonIDs []string          `json:"compatibleAddonIds,omitzero"`
	// Default trial configuration for the plan
	DefaultTrialConfig V1PlanUpdateParamsDefaultTrialConfig `json:"defaultTrialConfig,omitzero"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1PlanUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default trial configuration for the plan
//
// The properties Duration, Units are required.
type V1PlanUpdateParamsDefaultTrialConfig struct {
	// The duration of the trial in the specified units
	Duration float64 `json:"duration" api:"required"`
	// The time unit for the trial duration (DAY or MONTH)
	//
	// Any of "DAY", "MONTH".
	Units string `json:"units,omitzero" api:"required"`
	// Budget configuration for the trial
	Budget V1PlanUpdateParamsDefaultTrialConfigBudget `json:"budget,omitzero"`
	// Behavior when the trial ends (CONVERT_TO_PAID or CANCEL_SUBSCRIPTION)
	//
	// Any of "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION".
	TrialEndBehavior string `json:"trialEndBehavior,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsDefaultTrialConfig) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsDefaultTrialConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsDefaultTrialConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsDefaultTrialConfig](
		"units", "DAY", "MONTH",
	)
	apijson.RegisterFieldValidator[V1PlanUpdateParamsDefaultTrialConfig](
		"trialEndBehavior", "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION",
	)
}

// Budget configuration for the trial
//
// The properties HasSoftLimit, Limit are required.
type V1PlanUpdateParamsDefaultTrialConfigBudget struct {
	// Whether the budget limit is a soft limit (allows overage) or hard limit
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// The budget limit amount
	Limit float64 `json:"limit" api:"required"`
	paramObj
}

func (r V1PlanUpdateParamsDefaultTrialConfigBudget) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsDefaultTrialConfigBudget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsDefaultTrialConfigBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PlanListParams struct {
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
	CreatedAt V1PlanListParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1PlanListParams]'s query parameters as `url.Values`.
func (r V1PlanListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by creation date using range operators: gt, gte, lt, lte
type V1PlanListParamsCreatedAt struct {
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

// URLQuery serializes [V1PlanListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r V1PlanListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PlanPublishParams struct {
	// The migration type of the package
	//
	// Any of "NEW_CUSTOMERS", "ALL_CUSTOMERS".
	MigrationType V1PlanPublishParamsMigrationType `json:"migrationType,omitzero" api:"required"`
	paramObj
}

func (r V1PlanPublishParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanPublishParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanPublishParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The migration type of the package
type V1PlanPublishParamsMigrationType string

const (
	V1PlanPublishParamsMigrationTypeNewCustomers V1PlanPublishParamsMigrationType = "NEW_CUSTOMERS"
	V1PlanPublishParamsMigrationTypeAllCustomers V1PlanPublishParamsMigrationType = "ALL_CUSTOMERS"
)

type V1PlanSetPricingParams struct {
	// Request to set the pricing configuration for a plan or addon.
	SetPackagePricing SetPackagePricingParam
	paramObj
}

func (r V1PlanSetPricingParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetPackagePricing)
}
func (r *V1PlanSetPricingParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SetPackagePricing)
}
