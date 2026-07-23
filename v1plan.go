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

// Operations related to plans
//
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
func (r *V1PlanService) New(ctx context.Context, params V1PlanNewParams, opts ...option.RequestOption) (res *Plan, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/plans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a plan by its unique identifier, including entitlements and pricing
// details.
func (r *V1PlanService) Get(ctx context.Context, id string, query V1PlanGetParams, opts ...option.RequestOption) (res *Plan, err error) {
	if !param.IsOmitted(query.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", query.XAccountID.Value)))
	}
	if !param.IsOmitted(query.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", query.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing plan's properties such as display name, description, and
// metadata.
func (r *V1PlanService) Update(ctx context.Context, id string, params V1PlanUpdateParams, opts ...option.RequestOption) (res *Plan, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieves a paginated list of plans in the environment.
func (r *V1PlanService) List(ctx context.Context, params V1PlanListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1PlanListResponse], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/plans"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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
func (r *V1PlanService) ListAutoPaging(ctx context.Context, params V1PlanListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1PlanListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, params, opts...))
}

// Archives a plan, preventing it from being used in new subscriptions.
func (r *V1PlanService) Archive(ctx context.Context, id string, body V1PlanArchiveParams, opts ...option.RequestOption) (res *Plan, err error) {
	if !param.IsOmitted(body.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", body.XAccountID.Value)))
	}
	if !param.IsOmitted(body.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", body.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Creates a draft version of an existing plan for modification before publishing.
func (r *V1PlanService) NewDraft(ctx context.Context, id string, body V1PlanNewDraftParams, opts ...option.RequestOption) (res *Plan, err error) {
	if !param.IsOmitted(body.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", body.XAccountID.Value)))
	}
	if !param.IsOmitted(body.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", body.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s/draft", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Retrieves the list of charges configured on a plan.
func (r *V1PlanService) ListCharges(ctx context.Context, id string, params V1PlanListChargesParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1PlanListChargesResponse], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s/charges", id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Retrieves the list of charges configured on a plan.
func (r *V1PlanService) ListChargesAutoPaging(ctx context.Context, id string, params V1PlanListChargesParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1PlanListChargesResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListCharges(ctx, id, params, opts...))
}

// Retrieves the list of overage charges configured on a plan.
func (r *V1PlanService) ListOverageCharges(ctx context.Context, id string, params V1PlanListOverageChargesParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1PlanListOverageChargesResponse], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s/overage-charges", id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Retrieves the list of overage charges configured on a plan.
func (r *V1PlanService) ListOverageChargesAutoPaging(ctx context.Context, id string, params V1PlanListOverageChargesParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1PlanListOverageChargesResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListOverageCharges(ctx, id, params, opts...))
}

// Publishes a draft plan, making it available for use in subscriptions.
func (r *V1PlanService) Publish(ctx context.Context, id string, params V1PlanPublishParams, opts ...option.RequestOption) (res *V1PlanPublishResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s/publish", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Removes a draft version of a plan.
func (r *V1PlanService) RemoveDraft(ctx context.Context, id string, body V1PlanRemoveDraftParams, opts ...option.RequestOption) (res *V1PlanRemoveDraftResponse, err error) {
	if !param.IsOmitted(body.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", body.XAccountID.Value)))
	}
	if !param.IsOmitted(body.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", body.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/plans/%s/draft", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
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

// A single pricing row on a plan or addon. Each charge encodes one (billingPeriod,
// billingModel, billingCadence, billingCountryCode) combination. Plans and addons
// own many of these — one per currency / billing period / feature.
type V1PlanListChargesResponse struct {
	// Unique identifier of the charge
	ID string `json:"id" api:"required" format:"uuid"`
	// The billing cadence (RECURRING or ONE_OFF)
	//
	// Any of "RECURRING", "ONE_OFF".
	BillingCadence V1PlanListChargesResponseBillingCadence `json:"billingCadence" api:"required"`
	// The billing model (FLAT_FEE, PER_UNIT, USAGE_BASED, CREDIT_BASED, MINIMUM_SPEND)
	//
	// Any of "FLAT_FEE", "MINIMUM_SPEND", "PER_UNIT", "USAGE_BASED", "CREDIT_BASED".
	BillingModel V1PlanListChargesResponseBillingModel `json:"billingModel" api:"required"`
	// The billing period (MONTHLY or ANNUALLY)
	//
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod V1PlanListChargesResponseBillingPeriod `json:"billingPeriod" api:"required"`
	// Timestamp when the charge was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// ISO country code for localized pricing, if any
	BillingCountryCode string `json:"billingCountryCode" api:"nullable"`
	// Identifier in the external billing integration (e.g. Stripe price id), if any
	BillingID string `json:"billingId" api:"nullable"`
	// Block size for usage-based pricing
	BlockSize float64 `json:"blockSize" api:"nullable"`
	// When credits are granted (for credit-based pricing)
	//
	// Any of "BEGINNING_OF_BILLING_PERIOD", "MONTHLY".
	CreditGrantCadence V1PlanListChargesResponseCreditGrantCadence `json:"creditGrantCadence" api:"nullable"`
	// Credit rate configuration for credit-based pricing
	CreditRate V1PlanListChargesResponseCreditRate `json:"creditRate" api:"nullable"`
	// Identifier in the linked CRM, if any
	CRMID string `json:"crmId" api:"nullable"`
	// Deep link to the charge in the linked CRM, if any
	CRMLinkURL string `json:"crmLinkUrl" api:"nullable"`
	// The feature this charge meters, if metered
	FeatureID string `json:"featureId" api:"nullable"`
	// Maximum unit quantity that can be purchased
	MaxUnitQuantity float64 `json:"maxUnitQuantity" api:"nullable"`
	// Minimum unit quantity that can be purchased
	MinUnitQuantity float64 `json:"minUnitQuantity" api:"nullable"`
	// The flat price amount and currency, when applicable
	Price V1PlanListChargesResponsePrice `json:"price" api:"nullable"`
	// Tiered pricing rows when the charge is tiered
	Tiers []V1PlanListChargesResponseTier `json:"tiers" api:"nullable"`
	// Tiered pricing mode (VOLUME or GRADUATED) when the charge is tiered
	//
	// Any of "VOLUME", "GRADUATED".
	TiersMode V1PlanListChargesResponseTiersMode `json:"tiersMode" api:"nullable"`
	// Custom currency identifier for top-up pricing, if any
	TopUpCustomCurrencyID string `json:"topUpCustomCurrencyId" api:"nullable"`
	// True if this charge is referenced by at least one subscription
	UsedInSubscriptions bool `json:"usedInSubscriptions" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		BillingCadence        respjson.Field
		BillingModel          respjson.Field
		BillingPeriod         respjson.Field
		CreatedAt             respjson.Field
		BillingCountryCode    respjson.Field
		BillingID             respjson.Field
		BlockSize             respjson.Field
		CreditGrantCadence    respjson.Field
		CreditRate            respjson.Field
		CRMID                 respjson.Field
		CRMLinkURL            respjson.Field
		FeatureID             respjson.Field
		MaxUnitQuantity       respjson.Field
		MinUnitQuantity       respjson.Field
		Price                 respjson.Field
		Tiers                 respjson.Field
		TiersMode             respjson.Field
		TopUpCustomCurrencyID respjson.Field
		UsedInSubscriptions   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListChargesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListChargesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The billing cadence (RECURRING or ONE_OFF)
type V1PlanListChargesResponseBillingCadence string

const (
	V1PlanListChargesResponseBillingCadenceRecurring V1PlanListChargesResponseBillingCadence = "RECURRING"
	V1PlanListChargesResponseBillingCadenceOneOff    V1PlanListChargesResponseBillingCadence = "ONE_OFF"
)

// The billing model (FLAT_FEE, PER_UNIT, USAGE_BASED, CREDIT_BASED, MINIMUM_SPEND)
type V1PlanListChargesResponseBillingModel string

const (
	V1PlanListChargesResponseBillingModelFlatFee      V1PlanListChargesResponseBillingModel = "FLAT_FEE"
	V1PlanListChargesResponseBillingModelMinimumSpend V1PlanListChargesResponseBillingModel = "MINIMUM_SPEND"
	V1PlanListChargesResponseBillingModelPerUnit      V1PlanListChargesResponseBillingModel = "PER_UNIT"
	V1PlanListChargesResponseBillingModelUsageBased   V1PlanListChargesResponseBillingModel = "USAGE_BASED"
	V1PlanListChargesResponseBillingModelCreditBased  V1PlanListChargesResponseBillingModel = "CREDIT_BASED"
)

// The billing period (MONTHLY or ANNUALLY)
type V1PlanListChargesResponseBillingPeriod string

const (
	V1PlanListChargesResponseBillingPeriodMonthly  V1PlanListChargesResponseBillingPeriod = "MONTHLY"
	V1PlanListChargesResponseBillingPeriodAnnually V1PlanListChargesResponseBillingPeriod = "ANNUALLY"
)

// When credits are granted (for credit-based pricing)
type V1PlanListChargesResponseCreditGrantCadence string

const (
	V1PlanListChargesResponseCreditGrantCadenceBeginningOfBillingPeriod V1PlanListChargesResponseCreditGrantCadence = "BEGINNING_OF_BILLING_PERIOD"
	V1PlanListChargesResponseCreditGrantCadenceMonthly                  V1PlanListChargesResponseCreditGrantCadence = "MONTHLY"
)

// Credit rate configuration for credit-based pricing
type V1PlanListChargesResponseCreditRate struct {
	// Credit rate amount
	Amount float64 `json:"amount" api:"required"`
	// Custom currency identifier
	CurrencyID string `json:"currencyId" api:"required"`
	// Optional cost formula expression
	CostFormula string `json:"costFormula" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		CurrencyID  respjson.Field
		CostFormula respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListChargesResponseCreditRate) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListChargesResponseCreditRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The flat price amount and currency, when applicable
type V1PlanListChargesResponsePrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// ISO 4217 currency code
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListChargesResponsePrice) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListChargesResponsePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single tier within a tiered charge
type V1PlanListChargesResponseTier struct {
	// Flat price for this tier
	FlatPrice V1PlanListChargesResponseTierFlatPrice `json:"flatPrice" api:"nullable"`
	// Per-unit price in this tier
	UnitPrice V1PlanListChargesResponseTierUnitPrice `json:"unitPrice" api:"nullable"`
	// Upper bound of this tier (null for unlimited)
	UpTo float64 `json:"upTo" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlatPrice   respjson.Field
		UnitPrice   respjson.Field
		UpTo        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListChargesResponseTier) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListChargesResponseTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flat price for this tier
type V1PlanListChargesResponseTierFlatPrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// ISO 4217 currency code
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListChargesResponseTierFlatPrice) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListChargesResponseTierFlatPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-unit price in this tier
type V1PlanListChargesResponseTierUnitPrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// ISO 4217 currency code
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListChargesResponseTierUnitPrice) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListChargesResponseTierUnitPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tiered pricing mode (VOLUME or GRADUATED) when the charge is tiered
type V1PlanListChargesResponseTiersMode string

const (
	V1PlanListChargesResponseTiersModeVolume    V1PlanListChargesResponseTiersMode = "VOLUME"
	V1PlanListChargesResponseTiersModeGraduated V1PlanListChargesResponseTiersMode = "GRADUATED"
)

// A single pricing row on a plan or addon. Each charge encodes one (billingPeriod,
// billingModel, billingCadence, billingCountryCode) combination. Plans and addons
// own many of these — one per currency / billing period / feature.
type V1PlanListOverageChargesResponse struct {
	// Unique identifier of the charge
	ID string `json:"id" api:"required" format:"uuid"`
	// The billing cadence (RECURRING or ONE_OFF)
	//
	// Any of "RECURRING", "ONE_OFF".
	BillingCadence V1PlanListOverageChargesResponseBillingCadence `json:"billingCadence" api:"required"`
	// The billing model (FLAT_FEE, PER_UNIT, USAGE_BASED, CREDIT_BASED, MINIMUM_SPEND)
	//
	// Any of "FLAT_FEE", "MINIMUM_SPEND", "PER_UNIT", "USAGE_BASED", "CREDIT_BASED".
	BillingModel V1PlanListOverageChargesResponseBillingModel `json:"billingModel" api:"required"`
	// The billing period (MONTHLY or ANNUALLY)
	//
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod V1PlanListOverageChargesResponseBillingPeriod `json:"billingPeriod" api:"required"`
	// Timestamp when the charge was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// ISO country code for localized pricing, if any
	BillingCountryCode string `json:"billingCountryCode" api:"nullable"`
	// Identifier in the external billing integration (e.g. Stripe price id), if any
	BillingID string `json:"billingId" api:"nullable"`
	// Block size for usage-based pricing
	BlockSize float64 `json:"blockSize" api:"nullable"`
	// When credits are granted (for credit-based pricing)
	//
	// Any of "BEGINNING_OF_BILLING_PERIOD", "MONTHLY".
	CreditGrantCadence V1PlanListOverageChargesResponseCreditGrantCadence `json:"creditGrantCadence" api:"nullable"`
	// Credit rate configuration for credit-based pricing
	CreditRate V1PlanListOverageChargesResponseCreditRate `json:"creditRate" api:"nullable"`
	// Identifier in the linked CRM, if any
	CRMID string `json:"crmId" api:"nullable"`
	// Deep link to the charge in the linked CRM, if any
	CRMLinkURL string `json:"crmLinkUrl" api:"nullable"`
	// The feature this charge meters, if metered
	FeatureID string `json:"featureId" api:"nullable"`
	// Maximum unit quantity that can be purchased
	MaxUnitQuantity float64 `json:"maxUnitQuantity" api:"nullable"`
	// Minimum unit quantity that can be purchased
	MinUnitQuantity float64 `json:"minUnitQuantity" api:"nullable"`
	// The flat price amount and currency, when applicable
	Price V1PlanListOverageChargesResponsePrice `json:"price" api:"nullable"`
	// Tiered pricing rows when the charge is tiered
	Tiers []V1PlanListOverageChargesResponseTier `json:"tiers" api:"nullable"`
	// Tiered pricing mode (VOLUME or GRADUATED) when the charge is tiered
	//
	// Any of "VOLUME", "GRADUATED".
	TiersMode V1PlanListOverageChargesResponseTiersMode `json:"tiersMode" api:"nullable"`
	// Custom currency identifier for top-up pricing, if any
	TopUpCustomCurrencyID string `json:"topUpCustomCurrencyId" api:"nullable"`
	// True if this charge is referenced by at least one subscription
	UsedInSubscriptions bool `json:"usedInSubscriptions" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		BillingCadence        respjson.Field
		BillingModel          respjson.Field
		BillingPeriod         respjson.Field
		CreatedAt             respjson.Field
		BillingCountryCode    respjson.Field
		BillingID             respjson.Field
		BlockSize             respjson.Field
		CreditGrantCadence    respjson.Field
		CreditRate            respjson.Field
		CRMID                 respjson.Field
		CRMLinkURL            respjson.Field
		FeatureID             respjson.Field
		MaxUnitQuantity       respjson.Field
		MinUnitQuantity       respjson.Field
		Price                 respjson.Field
		Tiers                 respjson.Field
		TiersMode             respjson.Field
		TopUpCustomCurrencyID respjson.Field
		UsedInSubscriptions   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListOverageChargesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListOverageChargesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The billing cadence (RECURRING or ONE_OFF)
type V1PlanListOverageChargesResponseBillingCadence string

const (
	V1PlanListOverageChargesResponseBillingCadenceRecurring V1PlanListOverageChargesResponseBillingCadence = "RECURRING"
	V1PlanListOverageChargesResponseBillingCadenceOneOff    V1PlanListOverageChargesResponseBillingCadence = "ONE_OFF"
)

// The billing model (FLAT_FEE, PER_UNIT, USAGE_BASED, CREDIT_BASED, MINIMUM_SPEND)
type V1PlanListOverageChargesResponseBillingModel string

const (
	V1PlanListOverageChargesResponseBillingModelFlatFee      V1PlanListOverageChargesResponseBillingModel = "FLAT_FEE"
	V1PlanListOverageChargesResponseBillingModelMinimumSpend V1PlanListOverageChargesResponseBillingModel = "MINIMUM_SPEND"
	V1PlanListOverageChargesResponseBillingModelPerUnit      V1PlanListOverageChargesResponseBillingModel = "PER_UNIT"
	V1PlanListOverageChargesResponseBillingModelUsageBased   V1PlanListOverageChargesResponseBillingModel = "USAGE_BASED"
	V1PlanListOverageChargesResponseBillingModelCreditBased  V1PlanListOverageChargesResponseBillingModel = "CREDIT_BASED"
)

// The billing period (MONTHLY or ANNUALLY)
type V1PlanListOverageChargesResponseBillingPeriod string

const (
	V1PlanListOverageChargesResponseBillingPeriodMonthly  V1PlanListOverageChargesResponseBillingPeriod = "MONTHLY"
	V1PlanListOverageChargesResponseBillingPeriodAnnually V1PlanListOverageChargesResponseBillingPeriod = "ANNUALLY"
)

// When credits are granted (for credit-based pricing)
type V1PlanListOverageChargesResponseCreditGrantCadence string

const (
	V1PlanListOverageChargesResponseCreditGrantCadenceBeginningOfBillingPeriod V1PlanListOverageChargesResponseCreditGrantCadence = "BEGINNING_OF_BILLING_PERIOD"
	V1PlanListOverageChargesResponseCreditGrantCadenceMonthly                  V1PlanListOverageChargesResponseCreditGrantCadence = "MONTHLY"
)

// Credit rate configuration for credit-based pricing
type V1PlanListOverageChargesResponseCreditRate struct {
	// Credit rate amount
	Amount float64 `json:"amount" api:"required"`
	// Custom currency identifier
	CurrencyID string `json:"currencyId" api:"required"`
	// Optional cost formula expression
	CostFormula string `json:"costFormula" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		CurrencyID  respjson.Field
		CostFormula respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListOverageChargesResponseCreditRate) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListOverageChargesResponseCreditRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The flat price amount and currency, when applicable
type V1PlanListOverageChargesResponsePrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// ISO 4217 currency code
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListOverageChargesResponsePrice) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListOverageChargesResponsePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single tier within a tiered charge
type V1PlanListOverageChargesResponseTier struct {
	// Flat price for this tier
	FlatPrice V1PlanListOverageChargesResponseTierFlatPrice `json:"flatPrice" api:"nullable"`
	// Per-unit price in this tier
	UnitPrice V1PlanListOverageChargesResponseTierUnitPrice `json:"unitPrice" api:"nullable"`
	// Upper bound of this tier (null for unlimited)
	UpTo float64 `json:"upTo" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlatPrice   respjson.Field
		UnitPrice   respjson.Field
		UpTo        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListOverageChargesResponseTier) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListOverageChargesResponseTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flat price for this tier
type V1PlanListOverageChargesResponseTierFlatPrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// ISO 4217 currency code
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListOverageChargesResponseTierFlatPrice) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListOverageChargesResponseTierFlatPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-unit price in this tier
type V1PlanListOverageChargesResponseTierUnitPrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// ISO 4217 currency code
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1PlanListOverageChargesResponseTierUnitPrice) RawJSON() string { return r.JSON.raw }
func (r *V1PlanListOverageChargesResponseTierUnitPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tiered pricing mode (VOLUME or GRADUATED) when the charge is tiered
type V1PlanListOverageChargesResponseTiersMode string

const (
	V1PlanListOverageChargesResponseTiersModeVolume    V1PlanListOverageChargesResponseTiersMode = "VOLUME"
	V1PlanListOverageChargesResponseTiersModeGraduated V1PlanListOverageChargesResponseTiersMode = "GRADUATED"
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
	ParentPlanID   param.Opt[string] `json:"parentPlanId,omitzero"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
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

type V1PlanGetParams struct {
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

type V1PlanUpdateParams struct {
	// The unique identifier for the entity in the billing provider
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// The description of the package
	Description param.Opt[string] `json:"description,omitzero"`
	// The ID of the parent plan, if applicable
	ParentPlanID param.Opt[string] `json:"parentPlanId,omitzero"`
	// The display name of the package
	DisplayName        param.Opt[string] `json:"displayName,omitzero"`
	XAccountID         param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID     param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	CompatibleAddonIDs []string          `json:"compatibleAddonIds,omitzero"`
	// Default trial configuration for the plan
	DefaultTrialConfig V1PlanUpdateParamsDefaultTrialConfig `json:"defaultTrialConfig,omitzero"`
	// Pricing configuration to set on the plan draft
	Charges V1PlanUpdateParamsCharges `json:"charges,omitzero"`
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

// Pricing configuration to set on the plan draft
//
// The property PricingType is required.
type V1PlanUpdateParamsCharges struct {
	// The pricing type (FREE, PAID, or CUSTOM)
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType string `json:"pricingType,omitzero" api:"required"`
	// Deprecated: billing integration ID
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// Minimum spend configuration per billing period
	MinimumSpend []V1PlanUpdateParamsChargesMinimumSpend `json:"minimumSpend,omitzero"`
	// When overage charges are billed
	//
	// Any of "ON_SUBSCRIPTION_RENEWAL", "MONTHLY".
	OverageBillingPeriod string `json:"overageBillingPeriod,omitzero"`
	// Array of overage pricing model configurations
	OveragePricingModels []V1PlanUpdateParamsChargesOveragePricingModel `json:"overagePricingModels,omitzero"`
	// Array of pricing model configurations
	PricingModels []V1PlanUpdateParamsChargesPricingModel `json:"pricingModels,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsCharges) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsCharges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsCharges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsCharges](
		"pricingType", "FREE", "PAID", "CUSTOM",
	)
	apijson.RegisterFieldValidator[V1PlanUpdateParamsCharges](
		"overageBillingPeriod", "ON_SUBSCRIPTION_RENEWAL", "MONTHLY",
	)
}

// Minimum spend configuration for a billing period.
//
// The properties BillingPeriod, Minimum are required.
type V1PlanUpdateParamsChargesMinimumSpend struct {
	// The billing period
	//
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod string `json:"billingPeriod,omitzero" api:"required"`
	// The minimum spend amount
	Minimum V1PlanUpdateParamsChargesMinimumSpendMinimum `json:"minimum,omitzero" api:"required"`
	paramObj
}

func (r V1PlanUpdateParamsChargesMinimumSpend) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesMinimumSpend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesMinimumSpend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesMinimumSpend](
		"billingPeriod", "MONTHLY", "ANNUALLY",
	)
}

// The minimum spend amount
//
// The property Amount is required.
type V1PlanUpdateParamsChargesMinimumSpendMinimum struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// The price currency
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesMinimumSpendMinimum) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesMinimumSpendMinimum
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesMinimumSpendMinimum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesMinimumSpendMinimum](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Overage pricing model configuration.
//
// The properties BillingModel, PricePeriods are required.
type V1PlanUpdateParamsChargesOveragePricingModel struct {
	// The billing model for overages
	//
	// Any of "FLAT_FEE", "MINIMUM_SPEND", "PER_UNIT", "USAGE_BASED", "CREDIT_BASED".
	BillingModel string `json:"billingModel,omitzero" api:"required"`
	// Price periods for overage pricing
	PricePeriods []V1PlanUpdateParamsChargesOveragePricingModelPricePeriod `json:"pricePeriods,omitzero" api:"required"`
	// The refId of the custom currency this credit overage applies to
	CurrencyID param.Opt[string] `json:"currencyId,omitzero"`
	// The feature ID for overage pricing
	FeatureID param.Opt[string] `json:"featureId,omitzero"`
	// The billing cadence for overages
	//
	// Any of "RECURRING", "ONE_OFF".
	BillingCadence string `json:"billingCadence,omitzero"`
	// Entitlement configuration for the overage feature
	Entitlement V1PlanUpdateParamsChargesOveragePricingModelEntitlement `json:"entitlement,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesOveragePricingModel) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesOveragePricingModel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesOveragePricingModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesOveragePricingModel](
		"billingModel", "FLAT_FEE", "MINIMUM_SPEND", "PER_UNIT", "USAGE_BASED", "CREDIT_BASED",
	)
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesOveragePricingModel](
		"billingCadence", "RECURRING", "ONE_OFF",
	)
}

// Price configuration for a specific billing period.
//
// The property BillingPeriod is required.
type V1PlanUpdateParamsChargesOveragePricingModelPricePeriod struct {
	// The billing period (MONTHLY or ANNUALLY)
	//
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod string `json:"billingPeriod,omitzero" api:"required"`
	// ISO country code for localized pricing
	BillingCountryCode param.Opt[string] `json:"billingCountryCode,omitzero"`
	// Block size for usage-based pricing
	BlockSize param.Opt[float64] `json:"blockSize,omitzero"`
	// When credits are granted
	//
	// Any of "BEGINNING_OF_BILLING_PERIOD", "MONTHLY".
	CreditGrantCadence string `json:"creditGrantCadence,omitzero"`
	// Credit rate configuration for credit-based pricing
	CreditRate V1PlanUpdateParamsChargesOveragePricingModelPricePeriodCreditRate `json:"creditRate,omitzero"`
	// The price amount and currency
	Price V1PlanUpdateParamsChargesOveragePricingModelPricePeriodPrice `json:"price,omitzero"`
	// Tiered pricing configuration
	Tiers []V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTier `json:"tiers,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesOveragePricingModelPricePeriod) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesOveragePricingModelPricePeriod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesOveragePricingModelPricePeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesOveragePricingModelPricePeriod](
		"billingPeriod", "MONTHLY", "ANNUALLY",
	)
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesOveragePricingModelPricePeriod](
		"creditGrantCadence", "BEGINNING_OF_BILLING_PERIOD", "MONTHLY",
	)
}

// Credit rate configuration for credit-based pricing
//
// The properties Amount, CurrencyID are required.
type V1PlanUpdateParamsChargesOveragePricingModelPricePeriodCreditRate struct {
	// The credit rate amount
	Amount float64 `json:"amount" api:"required"`
	// The custom currency ID
	CurrencyID string `json:"currencyId" api:"required"`
	// Optional cost formula expression
	CostFormula param.Opt[string] `json:"costFormula,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesOveragePricingModelPricePeriodCreditRate) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesOveragePricingModelPricePeriodCreditRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesOveragePricingModelPricePeriodCreditRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The price amount and currency
//
// The property Amount is required.
type V1PlanUpdateParamsChargesOveragePricingModelPricePeriodPrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// The price currency
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesOveragePricingModelPricePeriodPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesOveragePricingModelPricePeriodPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesOveragePricingModelPricePeriodPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesOveragePricingModelPricePeriodPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// A tier in tiered pricing.
type V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTier struct {
	// Upper bound of this tier (null for unlimited)
	UpTo param.Opt[float64] `json:"upTo,omitzero"`
	// Flat price for this tier
	FlatPrice V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierFlatPrice `json:"flatPrice,omitzero"`
	// Per-unit price in this tier
	UnitPrice V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierUnitPrice `json:"unitPrice,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTier) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flat price for this tier
//
// The property Amount is required.
type V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierFlatPrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// The price currency
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierFlatPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierFlatPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierFlatPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierFlatPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Per-unit price in this tier
//
// The property Amount is required.
type V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierUnitPrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// The price currency
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierUnitPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierUnitPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierUnitPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierUnitPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Entitlement configuration for the overage feature
//
// The property FeatureID is required.
type V1PlanUpdateParamsChargesOveragePricingModelEntitlement struct {
	// The feature ID for the entitlement
	FeatureID string `json:"featureId" api:"required"`
	// Whether the limit is soft (allows overage)
	HasSoftLimit param.Opt[bool] `json:"hasSoftLimit,omitzero"`
	// Whether usage is unlimited
	HasUnlimitedUsage param.Opt[bool] `json:"hasUnlimitedUsage,omitzero"`
	// The usage limit before overage kicks in
	UsageLimit param.Opt[float64] `json:"usageLimit,omitzero"`
	// Monthly reset configuration
	MonthlyResetPeriodConfiguration V1PlanUpdateParamsChargesOveragePricingModelEntitlementMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// The usage reset period
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,omitzero"`
	// Weekly reset configuration
	WeeklyResetPeriodConfiguration V1PlanUpdateParamsChargesOveragePricingModelEntitlementWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Yearly reset configuration
	YearlyResetPeriodConfiguration V1PlanUpdateParamsChargesOveragePricingModelEntitlementYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesOveragePricingModelEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesOveragePricingModelEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesOveragePricingModelEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesOveragePricingModelEntitlement](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Monthly reset configuration
//
// The property AccordingTo is required.
type V1PlanUpdateParamsChargesOveragePricingModelEntitlementMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanUpdateParamsChargesOveragePricingModelEntitlementMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesOveragePricingModelEntitlementMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesOveragePricingModelEntitlementMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesOveragePricingModelEntitlementMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Weekly reset configuration
//
// The property AccordingTo is required.
type V1PlanUpdateParamsChargesOveragePricingModelEntitlementWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanUpdateParamsChargesOveragePricingModelEntitlementWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesOveragePricingModelEntitlementWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesOveragePricingModelEntitlementWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesOveragePricingModelEntitlementWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Yearly reset configuration
//
// The property AccordingTo is required.
type V1PlanUpdateParamsChargesOveragePricingModelEntitlementYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanUpdateParamsChargesOveragePricingModelEntitlementYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesOveragePricingModelEntitlementYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesOveragePricingModelEntitlementYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesOveragePricingModelEntitlementYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

// A pricing model configuration with billing details and price periods.
//
// The properties BillingModel, PricePeriods are required.
type V1PlanUpdateParamsChargesPricingModel struct {
	// The billing model (FLAT_FEE, PER_UNIT, USAGE_BASED, CREDIT_BASED)
	//
	// Any of "FLAT_FEE", "MINIMUM_SPEND", "PER_UNIT", "USAGE_BASED", "CREDIT_BASED".
	BillingModel string `json:"billingModel,omitzero" api:"required"`
	// Array of price period configurations (at least one required)
	PricePeriods []V1PlanUpdateParamsChargesPricingModelPricePeriod `json:"pricePeriods,omitzero" api:"required"`
	// The feature ID this pricing model is associated with
	FeatureID param.Opt[string] `json:"featureId,omitzero"`
	// Maximum number of units (max 999999)
	MaxUnitQuantity param.Opt[int64] `json:"maxUnitQuantity,omitzero"`
	// Minimum number of units
	MinUnitQuantity param.Opt[int64] `json:"minUnitQuantity,omitzero"`
	// The custom currency ID for top-up pricing
	TopUpCustomCurrencyID param.Opt[string] `json:"topUpCustomCurrencyId,omitzero"`
	// The billing cadence (RECURRING or ONE_OFF)
	//
	// Any of "RECURRING", "ONE_OFF".
	BillingCadence string `json:"billingCadence,omitzero"`
	// Monthly reset period configuration
	MonthlyResetPeriodConfiguration V1PlanUpdateParamsChargesPricingModelMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// The usage reset period
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,omitzero"`
	// The tiered pricing mode (VOLUME or GRADUATED)
	//
	// Any of "VOLUME", "GRADUATED".
	TiersMode string `json:"tiersMode,omitzero"`
	// Weekly reset period configuration
	WeeklyResetPeriodConfiguration V1PlanUpdateParamsChargesPricingModelWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Yearly reset period configuration
	YearlyResetPeriodConfiguration V1PlanUpdateParamsChargesPricingModelYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesPricingModel) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesPricingModel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesPricingModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModel](
		"billingModel", "FLAT_FEE", "MINIMUM_SPEND", "PER_UNIT", "USAGE_BASED", "CREDIT_BASED",
	)
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModel](
		"billingCadence", "RECURRING", "ONE_OFF",
	)
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModel](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModel](
		"tiersMode", "VOLUME", "GRADUATED",
	)
}

// Price configuration for a specific billing period.
//
// The property BillingPeriod is required.
type V1PlanUpdateParamsChargesPricingModelPricePeriod struct {
	// The billing period (MONTHLY or ANNUALLY)
	//
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod string `json:"billingPeriod,omitzero" api:"required"`
	// ISO country code for localized pricing
	BillingCountryCode param.Opt[string] `json:"billingCountryCode,omitzero"`
	// Block size for usage-based pricing
	BlockSize param.Opt[float64] `json:"blockSize,omitzero"`
	// When credits are granted
	//
	// Any of "BEGINNING_OF_BILLING_PERIOD", "MONTHLY".
	CreditGrantCadence string `json:"creditGrantCadence,omitzero"`
	// Credit rate configuration for credit-based pricing
	CreditRate V1PlanUpdateParamsChargesPricingModelPricePeriodCreditRate `json:"creditRate,omitzero"`
	// The price amount and currency
	Price V1PlanUpdateParamsChargesPricingModelPricePeriodPrice `json:"price,omitzero"`
	// Tiered pricing configuration
	Tiers []V1PlanUpdateParamsChargesPricingModelPricePeriodTier `json:"tiers,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesPricingModelPricePeriod) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesPricingModelPricePeriod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesPricingModelPricePeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModelPricePeriod](
		"billingPeriod", "MONTHLY", "ANNUALLY",
	)
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModelPricePeriod](
		"creditGrantCadence", "BEGINNING_OF_BILLING_PERIOD", "MONTHLY",
	)
}

// Credit rate configuration for credit-based pricing
//
// The properties Amount, CurrencyID are required.
type V1PlanUpdateParamsChargesPricingModelPricePeriodCreditRate struct {
	// The credit rate amount
	Amount float64 `json:"amount" api:"required"`
	// The custom currency ID
	CurrencyID string `json:"currencyId" api:"required"`
	// Optional cost formula expression
	CostFormula param.Opt[string] `json:"costFormula,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesPricingModelPricePeriodCreditRate) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesPricingModelPricePeriodCreditRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesPricingModelPricePeriodCreditRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The price amount and currency
//
// The property Amount is required.
type V1PlanUpdateParamsChargesPricingModelPricePeriodPrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// The price currency
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesPricingModelPricePeriodPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesPricingModelPricePeriodPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesPricingModelPricePeriodPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModelPricePeriodPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// A tier in tiered pricing.
type V1PlanUpdateParamsChargesPricingModelPricePeriodTier struct {
	// Upper bound of this tier (null for unlimited)
	UpTo param.Opt[float64] `json:"upTo,omitzero"`
	// Flat price for this tier
	FlatPrice V1PlanUpdateParamsChargesPricingModelPricePeriodTierFlatPrice `json:"flatPrice,omitzero"`
	// Per-unit price in this tier
	UnitPrice V1PlanUpdateParamsChargesPricingModelPricePeriodTierUnitPrice `json:"unitPrice,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesPricingModelPricePeriodTier) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesPricingModelPricePeriodTier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesPricingModelPricePeriodTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flat price for this tier
//
// The property Amount is required.
type V1PlanUpdateParamsChargesPricingModelPricePeriodTierFlatPrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// The price currency
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesPricingModelPricePeriodTierFlatPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesPricingModelPricePeriodTierFlatPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesPricingModelPricePeriodTierFlatPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModelPricePeriodTierFlatPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Per-unit price in this tier
//
// The property Amount is required.
type V1PlanUpdateParamsChargesPricingModelPricePeriodTierUnitPrice struct {
	// The price amount
	Amount float64 `json:"amount" api:"required"`
	// The price currency
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	Currency string `json:"currency,omitzero"`
	paramObj
}

func (r V1PlanUpdateParamsChargesPricingModelPricePeriodTierUnitPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesPricingModelPricePeriodTierUnitPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesPricingModelPricePeriodTierUnitPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModelPricePeriodTierUnitPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Monthly reset period configuration
//
// The property AccordingTo is required.
type V1PlanUpdateParamsChargesPricingModelMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanUpdateParamsChargesPricingModelMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesPricingModelMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesPricingModelMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModelMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Weekly reset period configuration
//
// The property AccordingTo is required.
type V1PlanUpdateParamsChargesPricingModelWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanUpdateParamsChargesPricingModelWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesPricingModelWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesPricingModelWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModelWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Yearly reset period configuration
//
// The property AccordingTo is required.
type V1PlanUpdateParamsChargesPricingModelYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1PlanUpdateParamsChargesPricingModelYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1PlanUpdateParamsChargesPricingModelYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PlanUpdateParamsChargesPricingModelYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1PlanUpdateParamsChargesPricingModelYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
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
	ProductID      param.Opt[string] `query:"productId,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	// Filter by creation date using range operators: gt, gte, lt, lte
	CreatedAt V1PlanListParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	// Filter by status. Supports comma-separated values for multiple statuses
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status []string `query:"status,omitzero" json:"-"`
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

type V1PlanArchiveParams struct {
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

type V1PlanNewDraftParams struct {
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

type V1PlanListChargesParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit          param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1PlanListChargesParams]'s query parameters as
// `url.Values`.
func (r V1PlanListChargesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PlanListOverageChargesParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit          param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1PlanListOverageChargesParams]'s query parameters as
// `url.Values`.
func (r V1PlanListOverageChargesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PlanPublishParams struct {
	// The migration type of the package
	//
	// Any of "NEW_CUSTOMERS", "ALL_CUSTOMERS".
	MigrationType  V1PlanPublishParamsMigrationType `json:"migrationType,omitzero" api:"required"`
	XAccountID     param.Opt[string]                `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]                `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
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

type V1PlanRemoveDraftParams struct {
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}
