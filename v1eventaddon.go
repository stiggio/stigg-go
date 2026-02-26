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

// V1EventAddonService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventAddonService] method instead.
type V1EventAddonService struct {
	Options      []option.RequestOption
	Draft        V1EventAddonDraftService
	Entitlements V1EventAddonEntitlementService
}

// NewV1EventAddonService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1EventAddonService(opts ...option.RequestOption) (r V1EventAddonService) {
	r = V1EventAddonService{}
	r.Options = opts
	r.Draft = NewV1EventAddonDraftService(opts...)
	r.Entitlements = NewV1EventAddonEntitlementService(opts...)
	return
}

// Archives an addon, preventing it from being used in new subscriptions.
func (r *V1EventAddonService) ArchiveAddon(ctx context.Context, id string, opts ...option.RequestOption) (res *Addon, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Creates a new addon in draft status, associated with a specific product.
func (r *V1EventAddonService) NewAddon(ctx context.Context, body V1EventAddonNewAddonParams, opts ...option.RequestOption) (res *Addon, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/addons"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of addons in the environment.
func (r *V1EventAddonService) ListAddons(ctx context.Context, query V1EventAddonListAddonsParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1EventAddonListAddonsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/addons"
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

// Retrieves a paginated list of addons in the environment.
func (r *V1EventAddonService) ListAddonsAutoPaging(ctx context.Context, query V1EventAddonListAddonsParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1EventAddonListAddonsResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListAddons(ctx, query, opts...))
}

// Publishes a draft addon, making it available for use in subscriptions.
func (r *V1EventAddonService) PublishAddon(ctx context.Context, id string, body V1EventAddonPublishAddonParams, opts ...option.RequestOption) (res *V1EventAddonPublishAddonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s/publish", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves an addon by its unique identifier, including entitlements and pricing
// details.
func (r *V1EventAddonService) GetAddon(ctx context.Context, id string, opts ...option.RequestOption) (res *Addon, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets the pricing configuration for an addon.
func (r *V1EventAddonService) SetPricing(ctx context.Context, id string, body V1EventAddonSetPricingParams, opts ...option.RequestOption) (res *SetPackagePricingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s/charges", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Updates an existing addon's properties such as display name, description, and
// metadata.
func (r *V1EventAddonService) UpdateAddon(ctx context.Context, id string, body V1EventAddonUpdateAddonParams, opts ...option.RequestOption) (res *Addon, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Response object
type Addon struct {
	// Addon configuration object
	Data AddonData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Addon) RawJSON() string { return r.JSON.raw }
func (r *Addon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Addon configuration object
type AddonData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID string `json:"billingId" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// List of addons the addon is dependant on
	Dependencies []string `json:"dependencies" api:"required"`
	// The description of the package
	Description string `json:"description" api:"required"`
	// The display name of the package
	DisplayName string `json:"displayName" api:"required"`
	// List of entitlements of the package
	Entitlements []AddonDataEntitlement `json:"entitlements" api:"required"`
	// Indicates if the package is the latest version
	IsLatest bool `json:"isLatest" api:"required"`
	// The maximum quantity of this addon that can be added to a subscription
	MaxQuantity int64 `json:"maxQuantity" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
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
		ID            respjson.Field
		BillingID     respjson.Field
		CreatedAt     respjson.Field
		Dependencies  respjson.Field
		Description   respjson.Field
		DisplayName   respjson.Field
		Entitlements  respjson.Field
		IsLatest      respjson.Field
		MaxQuantity   respjson.Field
		Metadata      respjson.Field
		PricingType   respjson.Field
		ProductID     respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		VersionNumber respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddonData) RawJSON() string { return r.JSON.raw }
func (r *AddonData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type AddonDataEntitlement struct {
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
func (r AddonDataEntitlement) RawJSON() string { return r.JSON.raw }
func (r *AddonDataEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to set the pricing configuration for a plan or addon.
//
// The property PricingType is required.
type SetPackagePricingParam struct {
	// The pricing type (FREE, PAID, or CUSTOM)
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType SetPackagePricingPricingType `json:"pricingType,omitzero" api:"required"`
	// Deprecated: billing integration ID
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// Minimum spend configuration per billing period
	MinimumSpend []SetPackagePricingMinimumSpendParam `json:"minimumSpend,omitzero"`
	// When overage charges are billed
	//
	// Any of "ON_SUBSCRIPTION_RENEWAL", "MONTHLY".
	OverageBillingPeriod SetPackagePricingOverageBillingPeriod `json:"overageBillingPeriod,omitzero"`
	// Array of overage pricing model configurations
	OveragePricingModels []SetPackagePricingOveragePricingModelParam `json:"overagePricingModels,omitzero"`
	// Array of pricing model configurations
	PricingModels []SetPackagePricingPricingModelParam `json:"pricingModels,omitzero"`
	paramObj
}

func (r SetPackagePricingParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pricing type (FREE, PAID, or CUSTOM)
type SetPackagePricingPricingType string

const (
	SetPackagePricingPricingTypeFree   SetPackagePricingPricingType = "FREE"
	SetPackagePricingPricingTypePaid   SetPackagePricingPricingType = "PAID"
	SetPackagePricingPricingTypeCustom SetPackagePricingPricingType = "CUSTOM"
)

// Minimum spend configuration for a billing period.
//
// The properties BillingPeriod, Minimum are required.
type SetPackagePricingMinimumSpendParam struct {
	// The billing period
	//
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod string `json:"billingPeriod,omitzero" api:"required"`
	// The minimum spend amount
	Minimum SetPackagePricingMinimumSpendMinimumParam `json:"minimum,omitzero" api:"required"`
	paramObj
}

func (r SetPackagePricingMinimumSpendParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingMinimumSpendParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingMinimumSpendParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingMinimumSpendParam](
		"billingPeriod", "MONTHLY", "ANNUALLY",
	)
}

// The minimum spend amount
//
// The property Amount is required.
type SetPackagePricingMinimumSpendMinimumParam struct {
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

func (r SetPackagePricingMinimumSpendMinimumParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingMinimumSpendMinimumParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingMinimumSpendMinimumParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingMinimumSpendMinimumParam](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// When overage charges are billed
type SetPackagePricingOverageBillingPeriod string

const (
	SetPackagePricingOverageBillingPeriodOnSubscriptionRenewal SetPackagePricingOverageBillingPeriod = "ON_SUBSCRIPTION_RENEWAL"
	SetPackagePricingOverageBillingPeriodMonthly               SetPackagePricingOverageBillingPeriod = "MONTHLY"
)

// Overage pricing model configuration.
//
// The properties BillingModel, PricePeriods are required.
type SetPackagePricingOveragePricingModelParam struct {
	// The billing model for overages
	//
	// Any of "FLAT_FEE", "MINIMUM_SPEND", "PER_UNIT", "USAGE_BASED", "CREDIT_BASED".
	BillingModel string `json:"billingModel,omitzero" api:"required"`
	// Price periods for overage pricing
	PricePeriods []SetPackagePricingOveragePricingModelPricePeriodParam `json:"pricePeriods,omitzero" api:"required"`
	// The feature ID for overage pricing
	FeatureID param.Opt[string] `json:"featureId,omitzero"`
	// Custom currency ID for overage top-up
	TopUpCustomCurrencyID param.Opt[string] `json:"topUpCustomCurrencyId,omitzero"`
	// The billing cadence for overages
	//
	// Any of "RECURRING", "ONE_OFF".
	BillingCadence string `json:"billingCadence,omitzero"`
	// Entitlement configuration for the overage feature
	Entitlement SetPackagePricingOveragePricingModelEntitlementParam `json:"entitlement,omitzero"`
	paramObj
}

func (r SetPackagePricingOveragePricingModelParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingOveragePricingModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingOveragePricingModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingOveragePricingModelParam](
		"billingModel", "FLAT_FEE", "MINIMUM_SPEND", "PER_UNIT", "USAGE_BASED", "CREDIT_BASED",
	)
	apijson.RegisterFieldValidator[SetPackagePricingOveragePricingModelParam](
		"billingCadence", "RECURRING", "ONE_OFF",
	)
}

// Price configuration for a specific billing period.
//
// The property BillingPeriod is required.
type SetPackagePricingOveragePricingModelPricePeriodParam struct {
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
	CreditRate SetPackagePricingOveragePricingModelPricePeriodCreditRateParam `json:"creditRate,omitzero"`
	// The price amount and currency
	Price SetPackagePricingOveragePricingModelPricePeriodPriceParam `json:"price,omitzero"`
	// Tiered pricing configuration
	Tiers []SetPackagePricingOveragePricingModelPricePeriodTierParam `json:"tiers,omitzero"`
	paramObj
}

func (r SetPackagePricingOveragePricingModelPricePeriodParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingOveragePricingModelPricePeriodParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingOveragePricingModelPricePeriodParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingOveragePricingModelPricePeriodParam](
		"billingPeriod", "MONTHLY", "ANNUALLY",
	)
	apijson.RegisterFieldValidator[SetPackagePricingOveragePricingModelPricePeriodParam](
		"creditGrantCadence", "BEGINNING_OF_BILLING_PERIOD", "MONTHLY",
	)
}

// Credit rate configuration for credit-based pricing
//
// The properties Amount, CurrencyID are required.
type SetPackagePricingOveragePricingModelPricePeriodCreditRateParam struct {
	// The credit rate amount
	Amount float64 `json:"amount" api:"required"`
	// The custom currency ID
	CurrencyID string `json:"currencyId" api:"required"`
	// Optional cost formula expression
	CostFormula param.Opt[string] `json:"costFormula,omitzero"`
	paramObj
}

func (r SetPackagePricingOveragePricingModelPricePeriodCreditRateParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingOveragePricingModelPricePeriodCreditRateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingOveragePricingModelPricePeriodCreditRateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The price amount and currency
//
// The property Amount is required.
type SetPackagePricingOveragePricingModelPricePeriodPriceParam struct {
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

func (r SetPackagePricingOveragePricingModelPricePeriodPriceParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingOveragePricingModelPricePeriodPriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingOveragePricingModelPricePeriodPriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingOveragePricingModelPricePeriodPriceParam](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// A tier in tiered pricing.
type SetPackagePricingOveragePricingModelPricePeriodTierParam struct {
	// Upper bound of this tier (null for unlimited)
	UpTo param.Opt[float64] `json:"upTo,omitzero"`
	// Flat price for this tier
	FlatPrice SetPackagePricingOveragePricingModelPricePeriodTierFlatPriceParam `json:"flatPrice,omitzero"`
	// Per-unit price in this tier
	UnitPrice SetPackagePricingOveragePricingModelPricePeriodTierUnitPriceParam `json:"unitPrice,omitzero"`
	paramObj
}

func (r SetPackagePricingOveragePricingModelPricePeriodTierParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingOveragePricingModelPricePeriodTierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingOveragePricingModelPricePeriodTierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flat price for this tier
//
// The property Amount is required.
type SetPackagePricingOveragePricingModelPricePeriodTierFlatPriceParam struct {
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

func (r SetPackagePricingOveragePricingModelPricePeriodTierFlatPriceParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingOveragePricingModelPricePeriodTierFlatPriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingOveragePricingModelPricePeriodTierFlatPriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingOveragePricingModelPricePeriodTierFlatPriceParam](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Per-unit price in this tier
//
// The property Amount is required.
type SetPackagePricingOveragePricingModelPricePeriodTierUnitPriceParam struct {
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

func (r SetPackagePricingOveragePricingModelPricePeriodTierUnitPriceParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingOveragePricingModelPricePeriodTierUnitPriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingOveragePricingModelPricePeriodTierUnitPriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingOveragePricingModelPricePeriodTierUnitPriceParam](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Entitlement configuration for the overage feature
//
// The property FeatureID is required.
type SetPackagePricingOveragePricingModelEntitlementParam struct {
	// The feature ID for the entitlement
	FeatureID string `json:"featureId" api:"required"`
	// Whether the limit is soft (allows overage)
	HasSoftLimit param.Opt[bool] `json:"hasSoftLimit,omitzero"`
	// Whether usage is unlimited
	HasUnlimitedUsage param.Opt[bool] `json:"hasUnlimitedUsage,omitzero"`
	// The usage limit before overage kicks in
	UsageLimit param.Opt[float64] `json:"usageLimit,omitzero"`
	// Monthly reset configuration
	MonthlyResetPeriodConfiguration SetPackagePricingOveragePricingModelEntitlementMonthlyResetPeriodConfigurationParam `json:"monthlyResetPeriodConfiguration,omitzero"`
	// The usage reset period
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,omitzero"`
	// Weekly reset configuration
	WeeklyResetPeriodConfiguration SetPackagePricingOveragePricingModelEntitlementWeeklyResetPeriodConfigurationParam `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Yearly reset configuration
	YearlyResetPeriodConfiguration SetPackagePricingOveragePricingModelEntitlementYearlyResetPeriodConfigurationParam `json:"yearlyResetPeriodConfiguration,omitzero"`
	paramObj
}

func (r SetPackagePricingOveragePricingModelEntitlementParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingOveragePricingModelEntitlementParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingOveragePricingModelEntitlementParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingOveragePricingModelEntitlementParam](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Monthly reset configuration
//
// The property AccordingTo is required.
type SetPackagePricingOveragePricingModelEntitlementMonthlyResetPeriodConfigurationParam struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r SetPackagePricingOveragePricingModelEntitlementMonthlyResetPeriodConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingOveragePricingModelEntitlementMonthlyResetPeriodConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingOveragePricingModelEntitlementMonthlyResetPeriodConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingOveragePricingModelEntitlementMonthlyResetPeriodConfigurationParam](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Weekly reset configuration
//
// The property AccordingTo is required.
type SetPackagePricingOveragePricingModelEntitlementWeeklyResetPeriodConfigurationParam struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r SetPackagePricingOveragePricingModelEntitlementWeeklyResetPeriodConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingOveragePricingModelEntitlementWeeklyResetPeriodConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingOveragePricingModelEntitlementWeeklyResetPeriodConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingOveragePricingModelEntitlementWeeklyResetPeriodConfigurationParam](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Yearly reset configuration
//
// The property AccordingTo is required.
type SetPackagePricingOveragePricingModelEntitlementYearlyResetPeriodConfigurationParam struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r SetPackagePricingOveragePricingModelEntitlementYearlyResetPeriodConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingOveragePricingModelEntitlementYearlyResetPeriodConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingOveragePricingModelEntitlementYearlyResetPeriodConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingOveragePricingModelEntitlementYearlyResetPeriodConfigurationParam](
		"accordingTo", "SubscriptionStart",
	)
}

// A pricing model configuration with billing details and price periods.
//
// The properties BillingModel, PricePeriods are required.
type SetPackagePricingPricingModelParam struct {
	// The billing model (FLAT_FEE, PER_UNIT, USAGE_BASED, CREDIT_BASED)
	//
	// Any of "FLAT_FEE", "MINIMUM_SPEND", "PER_UNIT", "USAGE_BASED", "CREDIT_BASED".
	BillingModel string `json:"billingModel,omitzero" api:"required"`
	// Array of price period configurations (at least one required)
	PricePeriods []SetPackagePricingPricingModelPricePeriodParam `json:"pricePeriods,omitzero" api:"required"`
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
	MonthlyResetPeriodConfiguration SetPackagePricingPricingModelMonthlyResetPeriodConfigurationParam `json:"monthlyResetPeriodConfiguration,omitzero"`
	// The usage reset period
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,omitzero"`
	// The tiered pricing mode (VOLUME or GRADUATED)
	//
	// Any of "VOLUME", "GRADUATED".
	TiersMode string `json:"tiersMode,omitzero"`
	// Weekly reset period configuration
	WeeklyResetPeriodConfiguration SetPackagePricingPricingModelWeeklyResetPeriodConfigurationParam `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Yearly reset period configuration
	YearlyResetPeriodConfiguration SetPackagePricingPricingModelYearlyResetPeriodConfigurationParam `json:"yearlyResetPeriodConfiguration,omitzero"`
	paramObj
}

func (r SetPackagePricingPricingModelParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingPricingModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingPricingModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelParam](
		"billingModel", "FLAT_FEE", "MINIMUM_SPEND", "PER_UNIT", "USAGE_BASED", "CREDIT_BASED",
	)
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelParam](
		"billingCadence", "RECURRING", "ONE_OFF",
	)
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelParam](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelParam](
		"tiersMode", "VOLUME", "GRADUATED",
	)
}

// Price configuration for a specific billing period.
//
// The property BillingPeriod is required.
type SetPackagePricingPricingModelPricePeriodParam struct {
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
	CreditRate SetPackagePricingPricingModelPricePeriodCreditRateParam `json:"creditRate,omitzero"`
	// The price amount and currency
	Price SetPackagePricingPricingModelPricePeriodPriceParam `json:"price,omitzero"`
	// Tiered pricing configuration
	Tiers []SetPackagePricingPricingModelPricePeriodTierParam `json:"tiers,omitzero"`
	paramObj
}

func (r SetPackagePricingPricingModelPricePeriodParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingPricingModelPricePeriodParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingPricingModelPricePeriodParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelPricePeriodParam](
		"billingPeriod", "MONTHLY", "ANNUALLY",
	)
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelPricePeriodParam](
		"creditGrantCadence", "BEGINNING_OF_BILLING_PERIOD", "MONTHLY",
	)
}

// Credit rate configuration for credit-based pricing
//
// The properties Amount, CurrencyID are required.
type SetPackagePricingPricingModelPricePeriodCreditRateParam struct {
	// The credit rate amount
	Amount float64 `json:"amount" api:"required"`
	// The custom currency ID
	CurrencyID string `json:"currencyId" api:"required"`
	// Optional cost formula expression
	CostFormula param.Opt[string] `json:"costFormula,omitzero"`
	paramObj
}

func (r SetPackagePricingPricingModelPricePeriodCreditRateParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingPricingModelPricePeriodCreditRateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingPricingModelPricePeriodCreditRateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The price amount and currency
//
// The property Amount is required.
type SetPackagePricingPricingModelPricePeriodPriceParam struct {
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

func (r SetPackagePricingPricingModelPricePeriodPriceParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingPricingModelPricePeriodPriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingPricingModelPricePeriodPriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelPricePeriodPriceParam](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// A tier in tiered pricing.
type SetPackagePricingPricingModelPricePeriodTierParam struct {
	// Upper bound of this tier (null for unlimited)
	UpTo param.Opt[float64] `json:"upTo,omitzero"`
	// Flat price for this tier
	FlatPrice SetPackagePricingPricingModelPricePeriodTierFlatPriceParam `json:"flatPrice,omitzero"`
	// Per-unit price in this tier
	UnitPrice SetPackagePricingPricingModelPricePeriodTierUnitPriceParam `json:"unitPrice,omitzero"`
	paramObj
}

func (r SetPackagePricingPricingModelPricePeriodTierParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingPricingModelPricePeriodTierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingPricingModelPricePeriodTierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flat price for this tier
//
// The property Amount is required.
type SetPackagePricingPricingModelPricePeriodTierFlatPriceParam struct {
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

func (r SetPackagePricingPricingModelPricePeriodTierFlatPriceParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingPricingModelPricePeriodTierFlatPriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingPricingModelPricePeriodTierFlatPriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelPricePeriodTierFlatPriceParam](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Per-unit price in this tier
//
// The property Amount is required.
type SetPackagePricingPricingModelPricePeriodTierUnitPriceParam struct {
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

func (r SetPackagePricingPricingModelPricePeriodTierUnitPriceParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingPricingModelPricePeriodTierUnitPriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingPricingModelPricePeriodTierUnitPriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelPricePeriodTierUnitPriceParam](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Monthly reset period configuration
//
// The property AccordingTo is required.
type SetPackagePricingPricingModelMonthlyResetPeriodConfigurationParam struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r SetPackagePricingPricingModelMonthlyResetPeriodConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingPricingModelMonthlyResetPeriodConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingPricingModelMonthlyResetPeriodConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelMonthlyResetPeriodConfigurationParam](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Weekly reset period configuration
//
// The property AccordingTo is required.
type SetPackagePricingPricingModelWeeklyResetPeriodConfigurationParam struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r SetPackagePricingPricingModelWeeklyResetPeriodConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingPricingModelWeeklyResetPeriodConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingPricingModelWeeklyResetPeriodConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelWeeklyResetPeriodConfigurationParam](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Yearly reset period configuration
//
// The property AccordingTo is required.
type SetPackagePricingPricingModelYearlyResetPeriodConfigurationParam struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r SetPackagePricingPricingModelYearlyResetPeriodConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow SetPackagePricingPricingModelYearlyResetPeriodConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetPackagePricingPricingModelYearlyResetPeriodConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SetPackagePricingPricingModelYearlyResetPeriodConfigurationParam](
		"accordingTo", "SubscriptionStart",
	)
}

// Response object
type SetPackagePricingResponse struct {
	// Result of setting package pricing.
	Data SetPackagePricingResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SetPackagePricingResponse) RawJSON() string { return r.JSON.raw }
func (r *SetPackagePricingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of setting package pricing.
type SetPackagePricingResponseData struct {
	// The package identifier (refId)
	ID string `json:"id" api:"required"`
	// The pricing type that was set
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType string `json:"pricingType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		PricingType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SetPackagePricingResponseData) RawJSON() string { return r.JSON.raw }
func (r *SetPackagePricingResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Addon configuration object
type V1EventAddonListAddonsResponse struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID string `json:"billingId" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// List of addons the addon is dependant on
	Dependencies []string `json:"dependencies" api:"required"`
	// The description of the package
	Description string `json:"description" api:"required"`
	// The display name of the package
	DisplayName string `json:"displayName" api:"required"`
	// List of entitlements of the package
	Entitlements []V1EventAddonListAddonsResponseEntitlement `json:"entitlements" api:"required"`
	// Indicates if the package is the latest version
	IsLatest bool `json:"isLatest" api:"required"`
	// The maximum quantity of this addon that can be added to a subscription
	MaxQuantity int64 `json:"maxQuantity" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The pricing type of the package
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType V1EventAddonListAddonsResponsePricingType `json:"pricingType" api:"required"`
	// The product id of the package
	ProductID string `json:"productId" api:"required"`
	// The status of the package
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status V1EventAddonListAddonsResponseStatus `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The version number of the package
	VersionNumber int64 `json:"versionNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BillingID     respjson.Field
		CreatedAt     respjson.Field
		Dependencies  respjson.Field
		Description   respjson.Field
		DisplayName   respjson.Field
		Entitlements  respjson.Field
		IsLatest      respjson.Field
		MaxQuantity   respjson.Field
		Metadata      respjson.Field
		PricingType   respjson.Field
		ProductID     respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		VersionNumber respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonListAddonsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonListAddonsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type V1EventAddonListAddonsResponseEntitlement struct {
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
func (r V1EventAddonListAddonsResponseEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonListAddonsResponseEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pricing type of the package
type V1EventAddonListAddonsResponsePricingType string

const (
	V1EventAddonListAddonsResponsePricingTypeFree   V1EventAddonListAddonsResponsePricingType = "FREE"
	V1EventAddonListAddonsResponsePricingTypePaid   V1EventAddonListAddonsResponsePricingType = "PAID"
	V1EventAddonListAddonsResponsePricingTypeCustom V1EventAddonListAddonsResponsePricingType = "CUSTOM"
)

// The status of the package
type V1EventAddonListAddonsResponseStatus string

const (
	V1EventAddonListAddonsResponseStatusDraft     V1EventAddonListAddonsResponseStatus = "DRAFT"
	V1EventAddonListAddonsResponseStatusPublished V1EventAddonListAddonsResponseStatus = "PUBLISHED"
	V1EventAddonListAddonsResponseStatusArchived  V1EventAddonListAddonsResponseStatus = "ARCHIVED"
)

// Response containing task ID for publish operation
type V1EventAddonPublishAddonResponse struct {
	Data V1EventAddonPublishAddonResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonPublishAddonResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonPublishAddonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventAddonPublishAddonResponseData struct {
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
func (r V1EventAddonPublishAddonResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonPublishAddonResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventAddonNewAddonParams struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// The display name of the package
	DisplayName string `json:"displayName" api:"required"`
	// The product id of the package
	ProductID string `json:"productId" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// The description of the package
	Description param.Opt[string] `json:"description,omitzero"`
	// The maximum quantity of this addon that can be added to a subscription
	MaxQuantity param.Opt[int64] `json:"maxQuantity,omitzero"`
	// The pricing type of the package
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType V1EventAddonNewAddonParamsPricingType `json:"pricingType,omitzero"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,omitzero"`
	// The status of the package
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status V1EventAddonNewAddonParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r V1EventAddonNewAddonParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonNewAddonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonNewAddonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pricing type of the package
type V1EventAddonNewAddonParamsPricingType string

const (
	V1EventAddonNewAddonParamsPricingTypeFree   V1EventAddonNewAddonParamsPricingType = "FREE"
	V1EventAddonNewAddonParamsPricingTypePaid   V1EventAddonNewAddonParamsPricingType = "PAID"
	V1EventAddonNewAddonParamsPricingTypeCustom V1EventAddonNewAddonParamsPricingType = "CUSTOM"
)

// The status of the package
type V1EventAddonNewAddonParamsStatus string

const (
	V1EventAddonNewAddonParamsStatusDraft     V1EventAddonNewAddonParamsStatus = "DRAFT"
	V1EventAddonNewAddonParamsStatusPublished V1EventAddonNewAddonParamsStatus = "PUBLISHED"
	V1EventAddonNewAddonParamsStatusArchived  V1EventAddonNewAddonParamsStatus = "ARCHIVED"
)

type V1EventAddonListAddonsParams struct {
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
	CreatedAt V1EventAddonListAddonsParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventAddonListAddonsParams]'s query parameters as
// `url.Values`.
func (r V1EventAddonListAddonsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by creation date using range operators: gt, gte, lt, lte
type V1EventAddonListAddonsParamsCreatedAt struct {
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

// URLQuery serializes [V1EventAddonListAddonsParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r V1EventAddonListAddonsParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1EventAddonPublishAddonParams struct {
	// The migration type of the package
	//
	// Any of "NEW_CUSTOMERS", "ALL_CUSTOMERS".
	MigrationType V1EventAddonPublishAddonParamsMigrationType `json:"migrationType,omitzero" api:"required"`
	paramObj
}

func (r V1EventAddonPublishAddonParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonPublishAddonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonPublishAddonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The migration type of the package
type V1EventAddonPublishAddonParamsMigrationType string

const (
	V1EventAddonPublishAddonParamsMigrationTypeNewCustomers V1EventAddonPublishAddonParamsMigrationType = "NEW_CUSTOMERS"
	V1EventAddonPublishAddonParamsMigrationTypeAllCustomers V1EventAddonPublishAddonParamsMigrationType = "ALL_CUSTOMERS"
)

type V1EventAddonSetPricingParams struct {
	// Request to set the pricing configuration for a plan or addon.
	SetPackagePricing SetPackagePricingParam
	paramObj
}

func (r V1EventAddonSetPricingParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetPackagePricing)
}
func (r *V1EventAddonSetPricingParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SetPackagePricing)
}

type V1EventAddonUpdateAddonParams struct {
	// The unique identifier for the entity in the billing provider
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// The description of the package
	Description param.Opt[string] `json:"description,omitzero"`
	// The maximum quantity of this addon that can be added to a subscription
	MaxQuantity param.Opt[int64] `json:"maxQuantity,omitzero"`
	// The display name of the package
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// List of addons the addon is dependant on
	Dependencies []string `json:"dependencies,omitzero"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1EventAddonUpdateAddonParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonUpdateAddonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonUpdateAddonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
