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

// V1SubscriptionService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1SubscriptionService] method instead.
type V1SubscriptionService struct {
	Options      []option.RequestOption
	FutureUpdate V1SubscriptionFutureUpdateService
}

// NewV1SubscriptionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1SubscriptionService(opts ...option.RequestOption) (r V1SubscriptionService) {
	r = V1SubscriptionService{}
	r.Options = opts
	r.FutureUpdate = NewV1SubscriptionFutureUpdateService(opts...)
	return
}

// Create a new Subscription
func (r *V1SubscriptionService) New(ctx context.Context, body V1SubscriptionNewParams, opts ...option.RequestOption) (res *V1SubscriptionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single Subscription by id
func (r *V1SubscriptionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *V1SubscriptionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of Subscriptions
func (r *V1SubscriptionService) List(ctx context.Context, query V1SubscriptionListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1SubscriptionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/subscriptions"
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

// Get a list of Subscriptions
func (r *V1SubscriptionService) ListAutoPaging(ctx context.Context, query V1SubscriptionListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1SubscriptionListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, query, opts...))
}

// Perform delegate on a Subscription
func (r *V1SubscriptionService) Delegate(ctx context.Context, id string, body V1SubscriptionDelegateParams, opts ...option.RequestOption) (res *V1SubscriptionDelegateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/delegate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform migrate to latest plan version on a Subscription
func (r *V1SubscriptionService) Migrate(ctx context.Context, id string, body V1SubscriptionMigrateParams, opts ...option.RequestOption) (res *V1SubscriptionMigrateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/migrate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new Subscription Preview
func (r *V1SubscriptionService) Preview(ctx context.Context, body V1SubscriptionPreviewParams, opts ...option.RequestOption) (res *V1SubscriptionPreviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/subscriptions/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform transfer to resource on a Subscription
func (r *V1SubscriptionService) Transfer(ctx context.Context, id string, body V1SubscriptionTransferParams, opts ...option.RequestOption) (res *V1SubscriptionTransferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/transfer", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1SubscriptionNewResponse struct {
	Data V1SubscriptionNewResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewResponseData struct {
	// Unique identifier for the provisioned subscription
	ID           string                                     `json:"id,required"`
	Entitlements []V1SubscriptionNewResponseDataEntitlement `json:"entitlements,required"`
	// Provision status: SUCCESS or PAYMENT_REQUIRED
	//
	// Any of "SUCCESS", "PAYMENT_REQUIRED".
	Status string `json:"status,required"`
	// Checkout billing ID when payment is required
	CheckoutBillingID string `json:"checkoutBillingId,nullable"`
	// URL to complete payment when PAYMENT_REQUIRED
	CheckoutURL string `json:"checkoutUrl,nullable"`
	// Whether the subscription is scheduled for future activation
	IsScheduled  bool                                      `json:"isScheduled"`
	Subscription V1SubscriptionNewResponseDataSubscription `json:"subscription"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Entitlements      respjson.Field
		Status            respjson.Field
		CheckoutBillingID respjson.Field
		CheckoutURL       respjson.Field
		IsScheduled       respjson.Field
		Subscription      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewResponseDataEntitlement struct {
	AccessDeniedReason string  `json:"accessDeniedReason,nullable"`
	CurrentUsage       float64 `json:"currentUsage"`
	// entitlement updated at
	EntitlementUpdatedAt time.Time                                       `json:"entitlementUpdatedAt,nullable" format:"date-time"`
	Feature              V1SubscriptionNewResponseDataEntitlementFeature `json:"feature,nullable"`
	HasUnlimitedUsage    bool                                            `json:"hasUnlimitedUsage,nullable"`
	IsGranted            bool                                            `json:"isGranted"`
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string  `json:"resetPeriod,nullable"`
	UsageLimit  float64 `json:"usageLimit,nullable"`
	// usage period anchor
	UsagePeriodAnchor time.Time `json:"usagePeriodAnchor,nullable" format:"date-time"`
	// usage period end
	UsagePeriodEnd time.Time `json:"usagePeriodEnd,nullable" format:"date-time"`
	// usage period start
	UsagePeriodStart time.Time `json:"usagePeriodStart,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessDeniedReason   respjson.Field
		CurrentUsage         respjson.Field
		EntitlementUpdatedAt respjson.Field
		Feature              respjson.Field
		HasUnlimitedUsage    respjson.Field
		IsGranted            respjson.Field
		ResetPeriod          respjson.Field
		UsageLimit           respjson.Field
		UsagePeriodAnchor    respjson.Field
		UsagePeriodEnd       respjson.Field
		UsagePeriodStart     respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionNewResponseDataEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionNewResponseDataEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewResponseDataEntitlementFeature struct {
	RefID string `json:"refId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RefID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionNewResponseDataEntitlementFeature) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionNewResponseDataEntitlementFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewResponseDataSubscription struct {
	// Subscription ID
	ID string `json:"id,required"`
	// Billing ID
	BillingID string `json:"billingId,required"`
	// Created at
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Customer ID
	CustomerID string `json:"customerId,required"`
	// Payment collection
	//
	// Any of "NOT_REQUIRED", "PROCESSING", "FAILED", "ACTION_REQUIRED".
	PaymentCollection string `json:"paymentCollection,required"`
	// Plan ID
	PlanID string `json:"planId,required"`
	// Pricing type
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType string `json:"pricingType,required"`
	// Subscription start date
	StartDate time.Time `json:"startDate,required" format:"date-time"`
	// Subscription status
	//
	// Any of "PAYMENT_PENDING", "ACTIVE", "EXPIRED", "IN_TRIAL", "CANCELED",
	// "NOT_STARTED".
	Status string `json:"status,required"`
	// Subscription cancellation date
	CancellationDate time.Time `json:"cancellationDate,nullable" format:"date-time"`
	// Subscription cancel reason
	//
	// Any of "UPGRADE_OR_DOWNGRADE", "CANCELLED_BY_BILLING", "EXPIRED",
	// "DETACH_BILLING", "TRIAL_ENDED", "Immediate", "TRIAL_CONVERTED",
	// "PENDING_PAYMENT_EXPIRED", "ScheduledCancellation", "CustomerArchived",
	// "AutoCancellationRule".
	CancelReason string `json:"cancelReason,nullable"`
	// End of the current billing period
	CurrentBillingPeriodEnd time.Time `json:"currentBillingPeriodEnd,nullable" format:"date-time"`
	// Start of the current billing period
	CurrentBillingPeriodStart time.Time `json:"currentBillingPeriodStart,nullable" format:"date-time"`
	// Subscription effective end date
	EffectiveEndDate time.Time `json:"effectiveEndDate,nullable" format:"date-time"`
	// Subscription end date
	EndDate time.Time `json:"endDate,nullable" format:"date-time"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata"`
	// Paying customer ID for delegated billing
	PayingCustomerID string `json:"payingCustomerId,nullable"`
	// The method used to collect payments for a subscription
	//
	// Any of "CHARGE", "INVOICE", "NONE".
	PaymentCollectionMethod string                                           `json:"paymentCollectionMethod,nullable"`
	Prices                  []V1SubscriptionNewResponseDataSubscriptionPrice `json:"prices"`
	// Resource ID
	ResourceID string `json:"resourceId,nullable"`
	// Subscription trial end date
	TrialEndDate time.Time `json:"trialEndDate,nullable" format:"date-time"`
	UnitQuantity float64   `json:"unitQuantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		BillingID                 respjson.Field
		CreatedAt                 respjson.Field
		CustomerID                respjson.Field
		PaymentCollection         respjson.Field
		PlanID                    respjson.Field
		PricingType               respjson.Field
		StartDate                 respjson.Field
		Status                    respjson.Field
		CancellationDate          respjson.Field
		CancelReason              respjson.Field
		CurrentBillingPeriodEnd   respjson.Field
		CurrentBillingPeriodStart respjson.Field
		EffectiveEndDate          respjson.Field
		EndDate                   respjson.Field
		Metadata                  respjson.Field
		PayingCustomerID          respjson.Field
		PaymentCollectionMethod   respjson.Field
		Prices                    respjson.Field
		ResourceID                respjson.Field
		TrialEndDate              respjson.Field
		UnitQuantity              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionNewResponseDataSubscription) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionNewResponseDataSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewResponseDataSubscriptionPrice struct {
	// Addon identifier for the price override
	AddonID string `json:"addonId,nullable"`
	// Whether this is a base charge override
	BaseCharge bool `json:"baseCharge"`
	// Block size for pricing
	BlockSize float64 `json:"blockSize"`
	// Feature identifier for the price override
	FeatureID string `json:"featureId,nullable"`
	// Override price amount
	Price V1SubscriptionNewResponseDataSubscriptionPricePrice `json:"price"`
	// Pricing tiers configuration
	Tiers []V1SubscriptionNewResponseDataSubscriptionPriceTier `json:"tiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddonID     respjson.Field
		BaseCharge  respjson.Field
		BlockSize   respjson.Field
		FeatureID   respjson.Field
		Price       respjson.Field
		Tiers       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionNewResponseDataSubscriptionPrice) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionNewResponseDataSubscriptionPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Override price amount
type V1SubscriptionNewResponseDataSubscriptionPricePrice struct {
	// The price amount
	Amount float64 `json:"amount"`
	// The billing country code of the price
	BillingCountryCode string `json:"billingCountryCode,nullable"`
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
	Currency string `json:"currency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount             respjson.Field
		BillingCountryCode respjson.Field
		Currency           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionNewResponseDataSubscriptionPricePrice) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionNewResponseDataSubscriptionPricePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewResponseDataSubscriptionPriceTier struct {
	// The flat fee price of the price tier
	FlatPrice V1SubscriptionNewResponseDataSubscriptionPriceTierFlatPrice `json:"flatPrice"`
	// The unit price of the price tier
	UnitPrice V1SubscriptionNewResponseDataSubscriptionPriceTierUnitPrice `json:"unitPrice"`
	// The up to quantity of the price tier
	UpTo float64 `json:"upTo"`
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
func (r V1SubscriptionNewResponseDataSubscriptionPriceTier) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionNewResponseDataSubscriptionPriceTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The flat fee price of the price tier
type V1SubscriptionNewResponseDataSubscriptionPriceTierFlatPrice struct {
	// The price amount
	Amount float64 `json:"amount"`
	// The billing country code of the price
	BillingCountryCode string `json:"billingCountryCode,nullable"`
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
	Currency string `json:"currency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount             respjson.Field
		BillingCountryCode respjson.Field
		Currency           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionNewResponseDataSubscriptionPriceTierFlatPrice) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionNewResponseDataSubscriptionPriceTierFlatPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The unit price of the price tier
type V1SubscriptionNewResponseDataSubscriptionPriceTierUnitPrice struct {
	// The price amount
	Amount float64 `json:"amount"`
	// The billing country code of the price
	BillingCountryCode string `json:"billingCountryCode,nullable"`
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
	Currency string `json:"currency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount             respjson.Field
		BillingCountryCode respjson.Field
		Currency           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionNewResponseDataSubscriptionPriceTierUnitPrice) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionNewResponseDataSubscriptionPriceTierUnitPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionGetResponse struct {
	Data V1SubscriptionGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionGetResponseData struct {
	// Subscription ID
	ID string `json:"id,required"`
	// Billing ID
	BillingID string `json:"billingId,required"`
	// Created at
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Customer ID
	CustomerID string `json:"customerId,required"`
	// Payment collection
	//
	// Any of "NOT_REQUIRED", "PROCESSING", "FAILED", "ACTION_REQUIRED".
	PaymentCollection string `json:"paymentCollection,required"`
	// Plan ID
	PlanID string `json:"planId,required"`
	// Pricing type
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType string `json:"pricingType,required"`
	// Subscription start date
	StartDate time.Time `json:"startDate,required" format:"date-time"`
	// Subscription status
	//
	// Any of "PAYMENT_PENDING", "ACTIVE", "EXPIRED", "IN_TRIAL", "CANCELED",
	// "NOT_STARTED".
	Status string `json:"status,required"`
	// Subscription cancellation date
	CancellationDate time.Time `json:"cancellationDate,nullable" format:"date-time"`
	// Subscription cancel reason
	//
	// Any of "UPGRADE_OR_DOWNGRADE", "CANCELLED_BY_BILLING", "EXPIRED",
	// "DETACH_BILLING", "TRIAL_ENDED", "Immediate", "TRIAL_CONVERTED",
	// "PENDING_PAYMENT_EXPIRED", "ScheduledCancellation", "CustomerArchived",
	// "AutoCancellationRule".
	CancelReason string `json:"cancelReason,nullable"`
	// End of the current billing period
	CurrentBillingPeriodEnd time.Time `json:"currentBillingPeriodEnd,nullable" format:"date-time"`
	// Start of the current billing period
	CurrentBillingPeriodStart time.Time `json:"currentBillingPeriodStart,nullable" format:"date-time"`
	// Subscription effective end date
	EffectiveEndDate time.Time `json:"effectiveEndDate,nullable" format:"date-time"`
	// Subscription end date
	EndDate time.Time `json:"endDate,nullable" format:"date-time"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata"`
	// Paying customer ID for delegated billing
	PayingCustomerID string `json:"payingCustomerId,nullable"`
	// The method used to collect payments for a subscription
	//
	// Any of "CHARGE", "INVOICE", "NONE".
	PaymentCollectionMethod string                               `json:"paymentCollectionMethod,nullable"`
	Prices                  []V1SubscriptionGetResponseDataPrice `json:"prices"`
	// Resource ID
	ResourceID string `json:"resourceId,nullable"`
	// Subscription trial end date
	TrialEndDate time.Time `json:"trialEndDate,nullable" format:"date-time"`
	UnitQuantity float64   `json:"unitQuantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		BillingID                 respjson.Field
		CreatedAt                 respjson.Field
		CustomerID                respjson.Field
		PaymentCollection         respjson.Field
		PlanID                    respjson.Field
		PricingType               respjson.Field
		StartDate                 respjson.Field
		Status                    respjson.Field
		CancellationDate          respjson.Field
		CancelReason              respjson.Field
		CurrentBillingPeriodEnd   respjson.Field
		CurrentBillingPeriodStart respjson.Field
		EffectiveEndDate          respjson.Field
		EndDate                   respjson.Field
		Metadata                  respjson.Field
		PayingCustomerID          respjson.Field
		PaymentCollectionMethod   respjson.Field
		Prices                    respjson.Field
		ResourceID                respjson.Field
		TrialEndDate              respjson.Field
		UnitQuantity              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionGetResponseDataPrice struct {
	// Price ID
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt string `json:"createdAt,required"`
	// Last update timestamp
	UpdatedAt   string         `json:"updatedAt,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionGetResponseDataPrice) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionGetResponseDataPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionListResponse struct {
	// Subscription ID
	ID string `json:"id,required"`
	// Billing ID
	BillingID string `json:"billingId,required"`
	// Created at
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Customer ID
	CustomerID string `json:"customerId,required"`
	// Payment collection
	//
	// Any of "NOT_REQUIRED", "PROCESSING", "FAILED", "ACTION_REQUIRED".
	PaymentCollection V1SubscriptionListResponsePaymentCollection `json:"paymentCollection,required"`
	// Plan ID
	PlanID string `json:"planId,required"`
	// Pricing type
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType V1SubscriptionListResponsePricingType `json:"pricingType,required"`
	// Subscription start date
	StartDate time.Time `json:"startDate,required" format:"date-time"`
	// Subscription status
	//
	// Any of "PAYMENT_PENDING", "ACTIVE", "EXPIRED", "IN_TRIAL", "CANCELED",
	// "NOT_STARTED".
	Status V1SubscriptionListResponseStatus `json:"status,required"`
	// Subscription cancellation date
	CancellationDate time.Time `json:"cancellationDate,nullable" format:"date-time"`
	// Subscription cancel reason
	//
	// Any of "UPGRADE_OR_DOWNGRADE", "CANCELLED_BY_BILLING", "EXPIRED",
	// "DETACH_BILLING", "TRIAL_ENDED", "Immediate", "TRIAL_CONVERTED",
	// "PENDING_PAYMENT_EXPIRED", "ScheduledCancellation", "CustomerArchived",
	// "AutoCancellationRule".
	CancelReason V1SubscriptionListResponseCancelReason `json:"cancelReason,nullable"`
	// End of the current billing period
	CurrentBillingPeriodEnd time.Time `json:"currentBillingPeriodEnd,nullable" format:"date-time"`
	// Start of the current billing period
	CurrentBillingPeriodStart time.Time `json:"currentBillingPeriodStart,nullable" format:"date-time"`
	// Subscription effective end date
	EffectiveEndDate time.Time `json:"effectiveEndDate,nullable" format:"date-time"`
	// Subscription end date
	EndDate time.Time `json:"endDate,nullable" format:"date-time"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata"`
	// Paying customer ID for delegated billing
	PayingCustomerID string `json:"payingCustomerId,nullable"`
	// The method used to collect payments for a subscription
	//
	// Any of "CHARGE", "INVOICE", "NONE".
	PaymentCollectionMethod V1SubscriptionListResponsePaymentCollectionMethod `json:"paymentCollectionMethod,nullable"`
	Prices                  []V1SubscriptionListResponsePrice                 `json:"prices"`
	// Resource ID
	ResourceID string `json:"resourceId,nullable"`
	// Subscription trial end date
	TrialEndDate time.Time `json:"trialEndDate,nullable" format:"date-time"`
	UnitQuantity float64   `json:"unitQuantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		BillingID                 respjson.Field
		CreatedAt                 respjson.Field
		CustomerID                respjson.Field
		PaymentCollection         respjson.Field
		PlanID                    respjson.Field
		PricingType               respjson.Field
		StartDate                 respjson.Field
		Status                    respjson.Field
		CancellationDate          respjson.Field
		CancelReason              respjson.Field
		CurrentBillingPeriodEnd   respjson.Field
		CurrentBillingPeriodStart respjson.Field
		EffectiveEndDate          respjson.Field
		EndDate                   respjson.Field
		Metadata                  respjson.Field
		PayingCustomerID          respjson.Field
		PaymentCollectionMethod   respjson.Field
		Prices                    respjson.Field
		ResourceID                respjson.Field
		TrialEndDate              respjson.Field
		UnitQuantity              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payment collection
type V1SubscriptionListResponsePaymentCollection string

const (
	V1SubscriptionListResponsePaymentCollectionNotRequired    V1SubscriptionListResponsePaymentCollection = "NOT_REQUIRED"
	V1SubscriptionListResponsePaymentCollectionProcessing     V1SubscriptionListResponsePaymentCollection = "PROCESSING"
	V1SubscriptionListResponsePaymentCollectionFailed         V1SubscriptionListResponsePaymentCollection = "FAILED"
	V1SubscriptionListResponsePaymentCollectionActionRequired V1SubscriptionListResponsePaymentCollection = "ACTION_REQUIRED"
)

// Pricing type
type V1SubscriptionListResponsePricingType string

const (
	V1SubscriptionListResponsePricingTypeFree   V1SubscriptionListResponsePricingType = "FREE"
	V1SubscriptionListResponsePricingTypePaid   V1SubscriptionListResponsePricingType = "PAID"
	V1SubscriptionListResponsePricingTypeCustom V1SubscriptionListResponsePricingType = "CUSTOM"
)

// Subscription status
type V1SubscriptionListResponseStatus string

const (
	V1SubscriptionListResponseStatusPaymentPending V1SubscriptionListResponseStatus = "PAYMENT_PENDING"
	V1SubscriptionListResponseStatusActive         V1SubscriptionListResponseStatus = "ACTIVE"
	V1SubscriptionListResponseStatusExpired        V1SubscriptionListResponseStatus = "EXPIRED"
	V1SubscriptionListResponseStatusInTrial        V1SubscriptionListResponseStatus = "IN_TRIAL"
	V1SubscriptionListResponseStatusCanceled       V1SubscriptionListResponseStatus = "CANCELED"
	V1SubscriptionListResponseStatusNotStarted     V1SubscriptionListResponseStatus = "NOT_STARTED"
)

// Subscription cancel reason
type V1SubscriptionListResponseCancelReason string

const (
	V1SubscriptionListResponseCancelReasonUpgradeOrDowngrade    V1SubscriptionListResponseCancelReason = "UPGRADE_OR_DOWNGRADE"
	V1SubscriptionListResponseCancelReasonCancelledByBilling    V1SubscriptionListResponseCancelReason = "CANCELLED_BY_BILLING"
	V1SubscriptionListResponseCancelReasonExpired               V1SubscriptionListResponseCancelReason = "EXPIRED"
	V1SubscriptionListResponseCancelReasonDetachBilling         V1SubscriptionListResponseCancelReason = "DETACH_BILLING"
	V1SubscriptionListResponseCancelReasonTrialEnded            V1SubscriptionListResponseCancelReason = "TRIAL_ENDED"
	V1SubscriptionListResponseCancelReasonImmediate             V1SubscriptionListResponseCancelReason = "Immediate"
	V1SubscriptionListResponseCancelReasonTrialConverted        V1SubscriptionListResponseCancelReason = "TRIAL_CONVERTED"
	V1SubscriptionListResponseCancelReasonPendingPaymentExpired V1SubscriptionListResponseCancelReason = "PENDING_PAYMENT_EXPIRED"
	V1SubscriptionListResponseCancelReasonScheduledCancellation V1SubscriptionListResponseCancelReason = "ScheduledCancellation"
	V1SubscriptionListResponseCancelReasonCustomerArchived      V1SubscriptionListResponseCancelReason = "CustomerArchived"
	V1SubscriptionListResponseCancelReasonAutoCancellationRule  V1SubscriptionListResponseCancelReason = "AutoCancellationRule"
)

// The method used to collect payments for a subscription
type V1SubscriptionListResponsePaymentCollectionMethod string

const (
	V1SubscriptionListResponsePaymentCollectionMethodCharge  V1SubscriptionListResponsePaymentCollectionMethod = "CHARGE"
	V1SubscriptionListResponsePaymentCollectionMethodInvoice V1SubscriptionListResponsePaymentCollectionMethod = "INVOICE"
	V1SubscriptionListResponsePaymentCollectionMethodNone    V1SubscriptionListResponsePaymentCollectionMethod = "NONE"
)

type V1SubscriptionListResponsePrice struct {
	// Price ID
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt string `json:"createdAt,required"`
	// Last update timestamp
	UpdatedAt   string         `json:"updatedAt,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionListResponsePrice) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionListResponsePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionDelegateResponse struct {
	Data V1SubscriptionDelegateResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionDelegateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionDelegateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionDelegateResponseData struct {
	// Subscription ID
	ID string `json:"id,required"`
	// Billing ID
	BillingID string `json:"billingId,required"`
	// Created at
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Customer ID
	CustomerID string `json:"customerId,required"`
	// Payment collection
	//
	// Any of "NOT_REQUIRED", "PROCESSING", "FAILED", "ACTION_REQUIRED".
	PaymentCollection string `json:"paymentCollection,required"`
	// Plan ID
	PlanID string `json:"planId,required"`
	// Pricing type
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType string `json:"pricingType,required"`
	// Subscription start date
	StartDate time.Time `json:"startDate,required" format:"date-time"`
	// Subscription status
	//
	// Any of "PAYMENT_PENDING", "ACTIVE", "EXPIRED", "IN_TRIAL", "CANCELED",
	// "NOT_STARTED".
	Status string `json:"status,required"`
	// Subscription cancellation date
	CancellationDate time.Time `json:"cancellationDate,nullable" format:"date-time"`
	// Subscription cancel reason
	//
	// Any of "UPGRADE_OR_DOWNGRADE", "CANCELLED_BY_BILLING", "EXPIRED",
	// "DETACH_BILLING", "TRIAL_ENDED", "Immediate", "TRIAL_CONVERTED",
	// "PENDING_PAYMENT_EXPIRED", "ScheduledCancellation", "CustomerArchived",
	// "AutoCancellationRule".
	CancelReason string `json:"cancelReason,nullable"`
	// End of the current billing period
	CurrentBillingPeriodEnd time.Time `json:"currentBillingPeriodEnd,nullable" format:"date-time"`
	// Start of the current billing period
	CurrentBillingPeriodStart time.Time `json:"currentBillingPeriodStart,nullable" format:"date-time"`
	// Subscription effective end date
	EffectiveEndDate time.Time `json:"effectiveEndDate,nullable" format:"date-time"`
	// Subscription end date
	EndDate time.Time `json:"endDate,nullable" format:"date-time"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata"`
	// Paying customer ID for delegated billing
	PayingCustomerID string `json:"payingCustomerId,nullable"`
	// The method used to collect payments for a subscription
	//
	// Any of "CHARGE", "INVOICE", "NONE".
	PaymentCollectionMethod string                                    `json:"paymentCollectionMethod,nullable"`
	Prices                  []V1SubscriptionDelegateResponseDataPrice `json:"prices"`
	// Resource ID
	ResourceID string `json:"resourceId,nullable"`
	// Subscription trial end date
	TrialEndDate time.Time `json:"trialEndDate,nullable" format:"date-time"`
	UnitQuantity float64   `json:"unitQuantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		BillingID                 respjson.Field
		CreatedAt                 respjson.Field
		CustomerID                respjson.Field
		PaymentCollection         respjson.Field
		PlanID                    respjson.Field
		PricingType               respjson.Field
		StartDate                 respjson.Field
		Status                    respjson.Field
		CancellationDate          respjson.Field
		CancelReason              respjson.Field
		CurrentBillingPeriodEnd   respjson.Field
		CurrentBillingPeriodStart respjson.Field
		EffectiveEndDate          respjson.Field
		EndDate                   respjson.Field
		Metadata                  respjson.Field
		PayingCustomerID          respjson.Field
		PaymentCollectionMethod   respjson.Field
		Prices                    respjson.Field
		ResourceID                respjson.Field
		TrialEndDate              respjson.Field
		UnitQuantity              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionDelegateResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionDelegateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionDelegateResponseDataPrice struct {
	// Price ID
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt string `json:"createdAt,required"`
	// Last update timestamp
	UpdatedAt   string         `json:"updatedAt,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionDelegateResponseDataPrice) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionDelegateResponseDataPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionMigrateResponse struct {
	Data V1SubscriptionMigrateResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionMigrateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionMigrateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionMigrateResponseData struct {
	// Subscription ID
	ID string `json:"id,required"`
	// Billing ID
	BillingID string `json:"billingId,required"`
	// Created at
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Customer ID
	CustomerID string `json:"customerId,required"`
	// Payment collection
	//
	// Any of "NOT_REQUIRED", "PROCESSING", "FAILED", "ACTION_REQUIRED".
	PaymentCollection string `json:"paymentCollection,required"`
	// Plan ID
	PlanID string `json:"planId,required"`
	// Pricing type
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType string `json:"pricingType,required"`
	// Subscription start date
	StartDate time.Time `json:"startDate,required" format:"date-time"`
	// Subscription status
	//
	// Any of "PAYMENT_PENDING", "ACTIVE", "EXPIRED", "IN_TRIAL", "CANCELED",
	// "NOT_STARTED".
	Status string `json:"status,required"`
	// Subscription cancellation date
	CancellationDate time.Time `json:"cancellationDate,nullable" format:"date-time"`
	// Subscription cancel reason
	//
	// Any of "UPGRADE_OR_DOWNGRADE", "CANCELLED_BY_BILLING", "EXPIRED",
	// "DETACH_BILLING", "TRIAL_ENDED", "Immediate", "TRIAL_CONVERTED",
	// "PENDING_PAYMENT_EXPIRED", "ScheduledCancellation", "CustomerArchived",
	// "AutoCancellationRule".
	CancelReason string `json:"cancelReason,nullable"`
	// End of the current billing period
	CurrentBillingPeriodEnd time.Time `json:"currentBillingPeriodEnd,nullable" format:"date-time"`
	// Start of the current billing period
	CurrentBillingPeriodStart time.Time `json:"currentBillingPeriodStart,nullable" format:"date-time"`
	// Subscription effective end date
	EffectiveEndDate time.Time `json:"effectiveEndDate,nullable" format:"date-time"`
	// Subscription end date
	EndDate time.Time `json:"endDate,nullable" format:"date-time"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata"`
	// Paying customer ID for delegated billing
	PayingCustomerID string `json:"payingCustomerId,nullable"`
	// The method used to collect payments for a subscription
	//
	// Any of "CHARGE", "INVOICE", "NONE".
	PaymentCollectionMethod string                                   `json:"paymentCollectionMethod,nullable"`
	Prices                  []V1SubscriptionMigrateResponseDataPrice `json:"prices"`
	// Resource ID
	ResourceID string `json:"resourceId,nullable"`
	// Subscription trial end date
	TrialEndDate time.Time `json:"trialEndDate,nullable" format:"date-time"`
	UnitQuantity float64   `json:"unitQuantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		BillingID                 respjson.Field
		CreatedAt                 respjson.Field
		CustomerID                respjson.Field
		PaymentCollection         respjson.Field
		PlanID                    respjson.Field
		PricingType               respjson.Field
		StartDate                 respjson.Field
		Status                    respjson.Field
		CancellationDate          respjson.Field
		CancelReason              respjson.Field
		CurrentBillingPeriodEnd   respjson.Field
		CurrentBillingPeriodStart respjson.Field
		EffectiveEndDate          respjson.Field
		EndDate                   respjson.Field
		Metadata                  respjson.Field
		PayingCustomerID          respjson.Field
		PaymentCollectionMethod   respjson.Field
		Prices                    respjson.Field
		ResourceID                respjson.Field
		TrialEndDate              respjson.Field
		UnitQuantity              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionMigrateResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionMigrateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionMigrateResponseDataPrice struct {
	// Price ID
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt string `json:"createdAt,required"`
	// Last update timestamp
	UpdatedAt   string         `json:"updatedAt,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionMigrateResponseDataPrice) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionMigrateResponseDataPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponse struct {
	Data V1SubscriptionPreviewResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionPreviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseData struct {
	ImmediateInvoice    V1SubscriptionPreviewResponseDataImmediateInvoice   `json:"immediateInvoice,required"`
	BillingPeriodRange  V1SubscriptionPreviewResponseDataBillingPeriodRange `json:"billingPeriodRange"`
	FreeItems           []V1SubscriptionPreviewResponseDataFreeItem         `json:"freeItems"`
	HasScheduledUpdates bool                                                `json:"hasScheduledUpdates"`
	IsPlanDowngrade     bool                                                `json:"isPlanDowngrade"`
	RecurringInvoice    V1SubscriptionPreviewResponseDataRecurringInvoice   `json:"recurringInvoice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImmediateInvoice    respjson.Field
		BillingPeriodRange  respjson.Field
		FreeItems           respjson.Field
		HasScheduledUpdates respjson.Field
		IsPlanDowngrade     respjson.Field
		RecurringInvoice    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionPreviewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataImmediateInvoice struct {
	SubTotal           float64                                                             `json:"subTotal,required"`
	Total              float64                                                             `json:"total,required"`
	BillingPeriodRange V1SubscriptionPreviewResponseDataImmediateInvoiceBillingPeriodRange `json:"billingPeriodRange"`
	Currency           string                                                              `json:"currency,nullable"`
	Discount           float64                                                             `json:"discount"`
	DiscountDetails    V1SubscriptionPreviewResponseDataImmediateInvoiceDiscountDetails    `json:"discountDetails"`
	Discounts          []V1SubscriptionPreviewResponseDataImmediateInvoiceDiscount         `json:"discounts"`
	Lines              []V1SubscriptionPreviewResponseDataImmediateInvoiceLine             `json:"lines"`
	Tax                float64                                                             `json:"tax"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubTotal           respjson.Field
		Total              respjson.Field
		BillingPeriodRange respjson.Field
		Currency           respjson.Field
		Discount           respjson.Field
		DiscountDetails    respjson.Field
		Discounts          respjson.Field
		Lines              respjson.Field
		Tax                respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataImmediateInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionPreviewResponseDataImmediateInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataImmediateInvoiceBillingPeriodRange struct {
	// Billing period end date
	End time.Time `json:"end,required" format:"date-time"`
	// Billing period start date
	Start time.Time `json:"start,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataImmediateInvoiceBillingPeriodRange) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionPreviewResponseDataImmediateInvoiceBillingPeriodRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataImmediateInvoiceDiscountDetails struct {
	Code        string  `json:"code"`
	FixedAmount float64 `json:"fixedAmount"`
	Percentage  float64 `json:"percentage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		FixedAmount respjson.Field
		Percentage  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataImmediateInvoiceDiscountDetails) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionPreviewResponseDataImmediateInvoiceDiscountDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataImmediateInvoiceDiscount struct {
	Amount      float64 `json:"amount,required"`
	Currency    string  `json:"currency,required"`
	Description string  `json:"description,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataImmediateInvoiceDiscount) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionPreviewResponseDataImmediateInvoiceDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataImmediateInvoiceLine struct {
	Currency    string  `json:"currency,required"`
	Description string  `json:"description,required"`
	SubTotal    float64 `json:"subTotal,required"`
	UnitPrice   float64 `json:"unitPrice,required"`
	Quantity    float64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Description respjson.Field
		SubTotal    respjson.Field
		UnitPrice   respjson.Field
		Quantity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataImmediateInvoiceLine) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionPreviewResponseDataImmediateInvoiceLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataBillingPeriodRange struct {
	// Billing period end date
	End time.Time `json:"end" format:"date-time"`
	// Billing period start date
	Start time.Time `json:"start" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataBillingPeriodRange) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionPreviewResponseDataBillingPeriodRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataFreeItem struct {
	AddonID  string  `json:"addonId,required"`
	Quantity float64 `json:"quantity,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddonID     respjson.Field
		Quantity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataFreeItem) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionPreviewResponseDataFreeItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataRecurringInvoice struct {
	SubTotal           float64                                                             `json:"subTotal,required"`
	Total              float64                                                             `json:"total,required"`
	BillingPeriodRange V1SubscriptionPreviewResponseDataRecurringInvoiceBillingPeriodRange `json:"billingPeriodRange"`
	Currency           string                                                              `json:"currency,nullable"`
	Discount           float64                                                             `json:"discount"`
	DiscountDetails    V1SubscriptionPreviewResponseDataRecurringInvoiceDiscountDetails    `json:"discountDetails"`
	Discounts          []V1SubscriptionPreviewResponseDataRecurringInvoiceDiscount         `json:"discounts"`
	Lines              []V1SubscriptionPreviewResponseDataRecurringInvoiceLine             `json:"lines"`
	Tax                float64                                                             `json:"tax"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubTotal           respjson.Field
		Total              respjson.Field
		BillingPeriodRange respjson.Field
		Currency           respjson.Field
		Discount           respjson.Field
		DiscountDetails    respjson.Field
		Discounts          respjson.Field
		Lines              respjson.Field
		Tax                respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataRecurringInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionPreviewResponseDataRecurringInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataRecurringInvoiceBillingPeriodRange struct {
	// Billing period end date
	End time.Time `json:"end,required" format:"date-time"`
	// Billing period start date
	Start time.Time `json:"start,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataRecurringInvoiceBillingPeriodRange) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionPreviewResponseDataRecurringInvoiceBillingPeriodRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataRecurringInvoiceDiscountDetails struct {
	Code        string  `json:"code"`
	FixedAmount float64 `json:"fixedAmount"`
	Percentage  float64 `json:"percentage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		FixedAmount respjson.Field
		Percentage  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataRecurringInvoiceDiscountDetails) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionPreviewResponseDataRecurringInvoiceDiscountDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataRecurringInvoiceDiscount struct {
	Amount      float64 `json:"amount,required"`
	Currency    string  `json:"currency,required"`
	Description string  `json:"description,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataRecurringInvoiceDiscount) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionPreviewResponseDataRecurringInvoiceDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewResponseDataRecurringInvoiceLine struct {
	Currency    string  `json:"currency,required"`
	Description string  `json:"description,required"`
	SubTotal    float64 `json:"subTotal,required"`
	UnitPrice   float64 `json:"unitPrice,required"`
	Quantity    float64 `json:"quantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Description respjson.Field
		SubTotal    respjson.Field
		UnitPrice   respjson.Field
		Quantity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionPreviewResponseDataRecurringInvoiceLine) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionPreviewResponseDataRecurringInvoiceLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionTransferResponse struct {
	Data V1SubscriptionTransferResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionTransferResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionTransferResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionTransferResponseData struct {
	// Subscription ID
	ID string `json:"id,required"`
	// Billing ID
	BillingID string `json:"billingId,required"`
	// Created at
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Customer ID
	CustomerID string `json:"customerId,required"`
	// Payment collection
	//
	// Any of "NOT_REQUIRED", "PROCESSING", "FAILED", "ACTION_REQUIRED".
	PaymentCollection string `json:"paymentCollection,required"`
	// Plan ID
	PlanID string `json:"planId,required"`
	// Pricing type
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType string `json:"pricingType,required"`
	// Subscription start date
	StartDate time.Time `json:"startDate,required" format:"date-time"`
	// Subscription status
	//
	// Any of "PAYMENT_PENDING", "ACTIVE", "EXPIRED", "IN_TRIAL", "CANCELED",
	// "NOT_STARTED".
	Status string `json:"status,required"`
	// Subscription cancellation date
	CancellationDate time.Time `json:"cancellationDate,nullable" format:"date-time"`
	// Subscription cancel reason
	//
	// Any of "UPGRADE_OR_DOWNGRADE", "CANCELLED_BY_BILLING", "EXPIRED",
	// "DETACH_BILLING", "TRIAL_ENDED", "Immediate", "TRIAL_CONVERTED",
	// "PENDING_PAYMENT_EXPIRED", "ScheduledCancellation", "CustomerArchived",
	// "AutoCancellationRule".
	CancelReason string `json:"cancelReason,nullable"`
	// End of the current billing period
	CurrentBillingPeriodEnd time.Time `json:"currentBillingPeriodEnd,nullable" format:"date-time"`
	// Start of the current billing period
	CurrentBillingPeriodStart time.Time `json:"currentBillingPeriodStart,nullable" format:"date-time"`
	// Subscription effective end date
	EffectiveEndDate time.Time `json:"effectiveEndDate,nullable" format:"date-time"`
	// Subscription end date
	EndDate time.Time `json:"endDate,nullable" format:"date-time"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata"`
	// Paying customer ID for delegated billing
	PayingCustomerID string `json:"payingCustomerId,nullable"`
	// The method used to collect payments for a subscription
	//
	// Any of "CHARGE", "INVOICE", "NONE".
	PaymentCollectionMethod string                                    `json:"paymentCollectionMethod,nullable"`
	Prices                  []V1SubscriptionTransferResponseDataPrice `json:"prices"`
	// Resource ID
	ResourceID string `json:"resourceId,nullable"`
	// Subscription trial end date
	TrialEndDate time.Time `json:"trialEndDate,nullable" format:"date-time"`
	UnitQuantity float64   `json:"unitQuantity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		BillingID                 respjson.Field
		CreatedAt                 respjson.Field
		CustomerID                respjson.Field
		PaymentCollection         respjson.Field
		PlanID                    respjson.Field
		PricingType               respjson.Field
		StartDate                 respjson.Field
		Status                    respjson.Field
		CancellationDate          respjson.Field
		CancelReason              respjson.Field
		CurrentBillingPeriodEnd   respjson.Field
		CurrentBillingPeriodStart respjson.Field
		EffectiveEndDate          respjson.Field
		EndDate                   respjson.Field
		Metadata                  respjson.Field
		PayingCustomerID          respjson.Field
		PaymentCollectionMethod   respjson.Field
		Prices                    respjson.Field
		ResourceID                respjson.Field
		TrialEndDate              respjson.Field
		UnitQuantity              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionTransferResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionTransferResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionTransferResponseDataPrice struct {
	// Price ID
	ID string `json:"id,required"`
	// Creation timestamp
	CreatedAt string `json:"createdAt,required"`
	// Last update timestamp
	UpdatedAt   string         `json:"updatedAt,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionTransferResponseDataPrice) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionTransferResponseDataPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewParams struct {
	// Customer ID to provision the subscription for
	CustomerID string `json:"customerId,required"`
	// Plan ID to provision
	PlanID string `json:"planId,required"`
	// Unique identifier for the subscription
	ID param.Opt[string] `json:"id,omitzero"`
	// The ISO 3166-1 alpha-2 country code for billing
	BillingCountryCode param.Opt[string] `json:"billingCountryCode,omitzero"`
	// External billing system identifier
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// Optional paying customer ID for split billing scenarios
	PayingCustomerID param.Opt[string] `json:"payingCustomerId,omitzero"`
	// Optional resource ID for multi-instance subscriptions
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Salesforce ID
	SalesforceID param.Opt[string] `json:"salesforceId,omitzero"`
	// Whether to wait for payment confirmation before returning the subscription
	AwaitPaymentConfirmation param.Opt[bool] `json:"awaitPaymentConfirmation,omitzero"`
	// Subscription start date
	StartDate          param.Opt[time.Time]                      `json:"startDate,omitzero" format:"date-time"`
	UnitQuantity       param.Opt[float64]                        `json:"unitQuantity,omitzero"`
	Budget             V1SubscriptionNewParamsBudget             `json:"budget,omitzero"`
	MinimumSpend       V1SubscriptionNewParamsMinimumSpend       `json:"minimumSpend,omitzero"`
	Addons             []V1SubscriptionNewParamsAddon            `json:"addons,omitzero"`
	AppliedCoupon      V1SubscriptionNewParamsAppliedCoupon      `json:"appliedCoupon,omitzero"`
	BillingInformation V1SubscriptionNewParamsBillingInformation `json:"billingInformation,omitzero"`
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod   V1SubscriptionNewParamsBillingPeriod   `json:"billingPeriod,omitzero"`
	Charges         []V1SubscriptionNewParamsCharge        `json:"charges,omitzero"`
	CheckoutOptions V1SubscriptionNewParamsCheckoutOptions `json:"checkoutOptions,omitzero"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata,omitzero"`
	// How payments should be collected for this subscription
	//
	// Any of "CHARGE", "INVOICE", "NONE".
	PaymentCollectionMethod V1SubscriptionNewParamsPaymentCollectionMethod `json:"paymentCollectionMethod,omitzero"`
	PriceOverrides          []V1SubscriptionNewParamsPriceOverride         `json:"priceOverrides,omitzero"`
	// Strategy for scheduling subscription changes
	//
	// Any of "END_OF_BILLING_PERIOD", "END_OF_BILLING_MONTH", "IMMEDIATE".
	ScheduleStrategy           V1SubscriptionNewParamsScheduleStrategy           `json:"scheduleStrategy,omitzero"`
	SubscriptionEntitlements   []V1SubscriptionNewParamsSubscriptionEntitlement  `json:"subscriptionEntitlements,omitzero"`
	TrialOverrideConfiguration V1SubscriptionNewParamsTrialOverrideConfiguration `json:"trialOverrideConfiguration,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property AddonID is required.
type V1SubscriptionNewParamsAddon struct {
	// Addon identifier
	AddonID string `json:"addonId,required"`
	// Number of addon units
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsAddon) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsAddon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsAddon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewParamsAppliedCoupon struct {
	BillingCouponID param.Opt[string]                                 `json:"billingCouponId,omitzero"`
	CouponID        param.Opt[string]                                 `json:"couponId,omitzero"`
	PromotionCode   param.Opt[string]                                 `json:"promotionCode,omitzero"`
	Configuration   V1SubscriptionNewParamsAppliedCouponConfiguration `json:"configuration,omitzero"`
	Discount        V1SubscriptionNewParamsAppliedCouponDiscount      `json:"discount,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsAppliedCoupon) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsAppliedCoupon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsAppliedCoupon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewParamsAppliedCouponConfiguration struct {
	// Coupon start date
	StartDate param.Opt[time.Time] `json:"startDate,omitzero" format:"date-time"`
	paramObj
}

func (r V1SubscriptionNewParamsAppliedCouponConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsAppliedCouponConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsAppliedCouponConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewParamsAppliedCouponDiscount struct {
	Description      param.Opt[string]                                        `json:"description,omitzero"`
	DurationInMonths param.Opt[float64]                                       `json:"durationInMonths,omitzero"`
	Name             param.Opt[string]                                        `json:"name,omitzero"`
	PercentOff       param.Opt[float64]                                       `json:"percentOff,omitzero"`
	AmountsOff       []V1SubscriptionNewParamsAppliedCouponDiscountAmountsOff `json:"amountsOff,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsAppliedCouponDiscount) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsAppliedCouponDiscount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsAppliedCouponDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Amount is required.
type V1SubscriptionNewParamsAppliedCouponDiscountAmountsOff struct {
	Amount float64 `json:"amount,required"`
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

func (r V1SubscriptionNewParamsAppliedCouponDiscountAmountsOff) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsAppliedCouponDiscountAmountsOff
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsAppliedCouponDiscountAmountsOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionNewParamsAppliedCouponDiscountAmountsOff](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

type V1SubscriptionNewParamsBillingInformation struct {
	// Stripe Connect account to charge on behalf of
	ChargeOnBehalfOfAccount param.Opt[string] `json:"chargeOnBehalfOfAccount,omitzero"`
	// Billing integration identifier
	IntegrationID param.Opt[string] `json:"integrationId,omitzero"`
	// Number of days until invoice is due
	InvoiceDaysUntilDue param.Opt[float64] `json:"invoiceDaysUntilDue,omitzero"`
	// Whether the subscription is backdated
	IsBackdated param.Opt[bool] `json:"isBackdated,omitzero"`
	// Whether the invoice is marked as paid
	IsInvoicePaid param.Opt[bool] `json:"isInvoicePaid,omitzero"`
	// Tax percentage (0-100)
	TaxPercentage param.Opt[float64] `json:"taxPercentage,omitzero"`
	// Billing address for the subscription
	BillingAddress V1SubscriptionNewParamsBillingInformationBillingAddress `json:"billingAddress,omitzero"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata,omitzero"`
	// How to handle proration for billing changes
	//
	// Any of "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE".
	ProrationBehavior string `json:"prorationBehavior,omitzero"`
	// Customer tax identification numbers
	TaxIDs []V1SubscriptionNewParamsBillingInformationTaxID `json:"taxIds,omitzero"`
	// Tax rate identifiers to apply
	TaxRateIDs []string `json:"taxRateIds,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsBillingInformation) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsBillingInformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsBillingInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionNewParamsBillingInformation](
		"prorationBehavior", "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE",
	)
}

// Billing address for the subscription
type V1SubscriptionNewParamsBillingInformationBillingAddress struct {
	City       param.Opt[string] `json:"city,omitzero"`
	Country    param.Opt[string] `json:"country,omitzero"`
	Line1      param.Opt[string] `json:"line1,omitzero"`
	Line2      param.Opt[string] `json:"line2,omitzero"`
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	State      param.Opt[string] `json:"state,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsBillingInformationBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsBillingInformationBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsBillingInformationBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, Value are required.
type V1SubscriptionNewParamsBillingInformationTaxID struct {
	// The type of tax exemption identifier, such as VAT.
	Type string `json:"type,required"`
	// The actual tax identifier value
	Value string `json:"value,required"`
	paramObj
}

func (r V1SubscriptionNewParamsBillingInformationTaxID) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsBillingInformationTaxID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsBillingInformationTaxID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewParamsBillingPeriod string

const (
	V1SubscriptionNewParamsBillingPeriodMonthly  V1SubscriptionNewParamsBillingPeriod = "MONTHLY"
	V1SubscriptionNewParamsBillingPeriodAnnually V1SubscriptionNewParamsBillingPeriod = "ANNUALLY"
)

// The properties HasSoftLimit, Limit are required.
type V1SubscriptionNewParamsBudget struct {
	// Whether the budget is a soft limit
	HasSoftLimit bool `json:"hasSoftLimit,required"`
	// Maximum spending limit
	Limit float64 `json:"limit,required"`
	paramObj
}

func (r V1SubscriptionNewParamsBudget) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsBudget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Quantity, Type are required.
type V1SubscriptionNewParamsCharge struct {
	// Charge ID
	ID string `json:"id,required"`
	// Charge quantity
	Quantity float64 `json:"quantity,required"`
	// Charge type
	//
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r V1SubscriptionNewParamsCharge) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionNewParamsCharge](
		"type", "FEATURE", "CREDIT",
	)
}

// The properties CancelURL, SuccessURL are required.
type V1SubscriptionNewParamsCheckoutOptions struct {
	// URL to redirect to if checkout is canceled
	CancelURL string `json:"cancelUrl,required" format:"uri"`
	// URL to redirect to after successful checkout
	SuccessURL string `json:"successUrl,required" format:"uri"`
	// Optional reference ID for the checkout session
	ReferenceID param.Opt[string] `json:"referenceId,omitzero"`
	// Allow promotional codes during checkout
	AllowPromoCodes param.Opt[bool] `json:"allowPromoCodes,omitzero"`
	// Allow tax ID collection during checkout
	AllowTaxIDCollection param.Opt[bool] `json:"allowTaxIdCollection,omitzero"`
	// Collect billing address during checkout
	CollectBillingAddress param.Opt[bool] `json:"collectBillingAddress,omitzero"`
	// Collect phone number during checkout
	CollectPhoneNumber param.Opt[bool] `json:"collectPhoneNumber,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsCheckoutOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsCheckoutOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsCheckoutOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionNewParamsMinimumSpend struct {
	// Minimum spend amount
	Minimum V1SubscriptionNewParamsMinimumSpendMinimum `json:"minimum,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsMinimumSpend) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsMinimumSpend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsMinimumSpend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Minimum spend amount
type V1SubscriptionNewParamsMinimumSpendMinimum struct {
	// The billing country code of the price
	BillingCountryCode param.Opt[string] `json:"billingCountryCode,omitzero"`
	// The price amount
	Amount param.Opt[float64] `json:"amount,omitzero"`
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

func (r V1SubscriptionNewParamsMinimumSpendMinimum) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsMinimumSpendMinimum
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsMinimumSpendMinimum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionNewParamsMinimumSpendMinimum](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// How payments should be collected for this subscription
type V1SubscriptionNewParamsPaymentCollectionMethod string

const (
	V1SubscriptionNewParamsPaymentCollectionMethodCharge  V1SubscriptionNewParamsPaymentCollectionMethod = "CHARGE"
	V1SubscriptionNewParamsPaymentCollectionMethodInvoice V1SubscriptionNewParamsPaymentCollectionMethod = "INVOICE"
	V1SubscriptionNewParamsPaymentCollectionMethodNone    V1SubscriptionNewParamsPaymentCollectionMethod = "NONE"
)

type V1SubscriptionNewParamsPriceOverride struct {
	// Addon identifier for the price override
	AddonID param.Opt[string] `json:"addonId,omitzero"`
	// Feature identifier for the price override
	FeatureID param.Opt[string] `json:"featureId,omitzero"`
	// Whether this is a base charge override
	BaseCharge param.Opt[bool] `json:"baseCharge,omitzero"`
	// Block size for pricing
	BlockSize param.Opt[float64] `json:"blockSize,omitzero"`
	// Any of "BEGINNING_OF_BILLING_PERIOD", "MONTHLY".
	CreditGrantCadence string                                         `json:"creditGrantCadence,omitzero"`
	CreditRate         V1SubscriptionNewParamsPriceOverrideCreditRate `json:"creditRate,omitzero"`
	// Override price amount
	Price V1SubscriptionNewParamsPriceOverridePrice `json:"price,omitzero"`
	// Pricing tiers configuration
	Tiers []V1SubscriptionNewParamsPriceOverrideTier `json:"tiers,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsPriceOverride) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsPriceOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsPriceOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionNewParamsPriceOverride](
		"creditGrantCadence", "BEGINNING_OF_BILLING_PERIOD", "MONTHLY",
	)
}

// The properties Amount, CurrencyID are required.
type V1SubscriptionNewParamsPriceOverrideCreditRate struct {
	// The credit rate amount
	Amount float64 `json:"amount,required"`
	// The custom currency refId for the credit rate
	CurrencyID string `json:"currencyId,required"`
	// A custom formula for calculating cost based on single event dimensions
	CostFormula param.Opt[string] `json:"costFormula,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsPriceOverrideCreditRate) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsPriceOverrideCreditRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsPriceOverrideCreditRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Override price amount
type V1SubscriptionNewParamsPriceOverridePrice struct {
	// The billing country code of the price
	BillingCountryCode param.Opt[string] `json:"billingCountryCode,omitzero"`
	// The price amount
	Amount param.Opt[float64] `json:"amount,omitzero"`
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

func (r V1SubscriptionNewParamsPriceOverridePrice) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsPriceOverridePrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsPriceOverridePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionNewParamsPriceOverridePrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

type V1SubscriptionNewParamsPriceOverrideTier struct {
	// The up to quantity of the price tier
	UpTo param.Opt[float64] `json:"upTo,omitzero"`
	// The flat fee price of the price tier
	FlatPrice V1SubscriptionNewParamsPriceOverrideTierFlatPrice `json:"flatPrice,omitzero"`
	// The unit price of the price tier
	UnitPrice V1SubscriptionNewParamsPriceOverrideTierUnitPrice `json:"unitPrice,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsPriceOverrideTier) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsPriceOverrideTier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsPriceOverrideTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The flat fee price of the price tier
type V1SubscriptionNewParamsPriceOverrideTierFlatPrice struct {
	// The billing country code of the price
	BillingCountryCode param.Opt[string] `json:"billingCountryCode,omitzero"`
	// The price amount
	Amount param.Opt[float64] `json:"amount,omitzero"`
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

func (r V1SubscriptionNewParamsPriceOverrideTierFlatPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsPriceOverrideTierFlatPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsPriceOverrideTierFlatPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionNewParamsPriceOverrideTierFlatPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// The unit price of the price tier
type V1SubscriptionNewParamsPriceOverrideTierUnitPrice struct {
	// The billing country code of the price
	BillingCountryCode param.Opt[string] `json:"billingCountryCode,omitzero"`
	// The price amount
	Amount param.Opt[float64] `json:"amount,omitzero"`
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

func (r V1SubscriptionNewParamsPriceOverrideTierUnitPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsPriceOverrideTierUnitPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsPriceOverrideTierUnitPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionNewParamsPriceOverrideTierUnitPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Strategy for scheduling subscription changes
type V1SubscriptionNewParamsScheduleStrategy string

const (
	V1SubscriptionNewParamsScheduleStrategyEndOfBillingPeriod V1SubscriptionNewParamsScheduleStrategy = "END_OF_BILLING_PERIOD"
	V1SubscriptionNewParamsScheduleStrategyEndOfBillingMonth  V1SubscriptionNewParamsScheduleStrategy = "END_OF_BILLING_MONTH"
	V1SubscriptionNewParamsScheduleStrategyImmediate          V1SubscriptionNewParamsScheduleStrategy = "IMMEDIATE"
)

// The properties FeatureID, UsageLimit are required.
type V1SubscriptionNewParamsSubscriptionEntitlement struct {
	FeatureID  string          `json:"featureId,required"`
	UsageLimit float64         `json:"usageLimit,required"`
	IsGranted  param.Opt[bool] `json:"isGranted,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsSubscriptionEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsSubscriptionEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsSubscriptionEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property IsTrial is required.
type V1SubscriptionNewParamsTrialOverrideConfiguration struct {
	// Whether the subscription should start with a trial period
	IsTrial bool `json:"isTrial,required"`
	// Custom trial end date
	TrialEndDate param.Opt[time.Time] `json:"trialEndDate,omitzero" format:"date-time"`
	// Behavior when trial ends: CONVERT_TO_PAID or CANCEL_SUBSCRIPTION
	//
	// Any of "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION".
	TrialEndBehavior string `json:"trialEndBehavior,omitzero"`
	paramObj
}

func (r V1SubscriptionNewParamsTrialOverrideConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionNewParamsTrialOverrideConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionNewParamsTrialOverrideConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionNewParamsTrialOverrideConfiguration](
		"trialEndBehavior", "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION",
	)
}

type V1SubscriptionListParams struct {
	// Starting after this UUID for pagination
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Ending before this UUID for pagination
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Filter by customer ID
	CustomerID param.Opt[string] `query:"customerId,omitzero" json:"-"`
	// Items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by subscription status (comma-separated for multiple statuses, e.g.,
	// ACTIVE,IN_TRIAL)
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1SubscriptionListParams]'s query parameters as
// `url.Values`.
func (r V1SubscriptionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1SubscriptionDelegateParams struct {
	// The customer ID to delegate the subscription to
	TargetCustomerID string `json:"targetCustomerId,required"`
	paramObj
}

func (r V1SubscriptionDelegateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionDelegateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionDelegateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionMigrateParams struct {
	// When to migrate the subscription: IMMEDIATE or END_OF_BILLING_PERIOD
	//
	// Any of "END_OF_BILLING_PERIOD", "IMMEDIATE".
	SubscriptionMigrationTime V1SubscriptionMigrateParamsSubscriptionMigrationTime `json:"subscriptionMigrationTime,omitzero"`
	paramObj
}

func (r V1SubscriptionMigrateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionMigrateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionMigrateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When to migrate the subscription: IMMEDIATE or END_OF_BILLING_PERIOD
type V1SubscriptionMigrateParamsSubscriptionMigrationTime string

const (
	V1SubscriptionMigrateParamsSubscriptionMigrationTimeEndOfBillingPeriod V1SubscriptionMigrateParamsSubscriptionMigrationTime = "END_OF_BILLING_PERIOD"
	V1SubscriptionMigrateParamsSubscriptionMigrationTimeImmediate          V1SubscriptionMigrateParamsSubscriptionMigrationTime = "IMMEDIATE"
)

type V1SubscriptionPreviewParams struct {
	// Customer ID
	CustomerID string `json:"customerId,required"`
	// Plan ID
	PlanID             string            `json:"planId,required"`
	BillingCountryCode param.Opt[string] `json:"billingCountryCode,omitzero"`
	PayingCustomerID   param.Opt[string] `json:"payingCustomerId,omitzero"`
	ResourceID         param.Opt[string] `json:"resourceId,omitzero"`
	// Subscription start date
	StartDate          param.Opt[time.Time]                          `json:"startDate,omitzero" format:"date-time"`
	UnitQuantity       param.Opt[float64]                            `json:"unitQuantity,omitzero"`
	Addons             []V1SubscriptionPreviewParamsAddon            `json:"addons,omitzero"`
	AppliedCoupon      V1SubscriptionPreviewParamsAppliedCoupon      `json:"appliedCoupon,omitzero"`
	BillableFeatures   []V1SubscriptionPreviewParamsBillableFeature  `json:"billableFeatures,omitzero"`
	BillingInformation V1SubscriptionPreviewParamsBillingInformation `json:"billingInformation,omitzero"`
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod V1SubscriptionPreviewParamsBillingPeriod `json:"billingPeriod,omitzero"`
	Charges       []V1SubscriptionPreviewParamsCharge      `json:"charges,omitzero"`
	// Any of "END_OF_BILLING_PERIOD", "END_OF_BILLING_MONTH", "IMMEDIATE".
	ScheduleStrategy           V1SubscriptionPreviewParamsScheduleStrategy           `json:"scheduleStrategy,omitzero"`
	TrialOverrideConfiguration V1SubscriptionPreviewParamsTrialOverrideConfiguration `json:"trialOverrideConfiguration,omitzero"`
	paramObj
}

func (r V1SubscriptionPreviewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property AddonID is required.
type V1SubscriptionPreviewParamsAddon struct {
	// Addon ID
	AddonID  string           `json:"addonId,required"`
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	paramObj
}

func (r V1SubscriptionPreviewParamsAddon) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsAddon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsAddon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewParamsAppliedCoupon struct {
	BillingCouponID param.Opt[string]                                     `json:"billingCouponId,omitzero"`
	CouponID        param.Opt[string]                                     `json:"couponId,omitzero"`
	PromotionCode   param.Opt[string]                                     `json:"promotionCode,omitzero"`
	Configuration   V1SubscriptionPreviewParamsAppliedCouponConfiguration `json:"configuration,omitzero"`
	Discount        V1SubscriptionPreviewParamsAppliedCouponDiscount      `json:"discount,omitzero"`
	paramObj
}

func (r V1SubscriptionPreviewParamsAppliedCoupon) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsAppliedCoupon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsAppliedCoupon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewParamsAppliedCouponConfiguration struct {
	// Coupon start date
	StartDate param.Opt[time.Time] `json:"startDate,omitzero" format:"date-time"`
	paramObj
}

func (r V1SubscriptionPreviewParamsAppliedCouponConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsAppliedCouponConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsAppliedCouponConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewParamsAppliedCouponDiscount struct {
	Description      param.Opt[string]                                            `json:"description,omitzero"`
	DurationInMonths param.Opt[float64]                                           `json:"durationInMonths,omitzero"`
	Name             param.Opt[string]                                            `json:"name,omitzero"`
	PercentOff       param.Opt[float64]                                           `json:"percentOff,omitzero"`
	AmountsOff       []V1SubscriptionPreviewParamsAppliedCouponDiscountAmountsOff `json:"amountsOff,omitzero"`
	paramObj
}

func (r V1SubscriptionPreviewParamsAppliedCouponDiscount) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsAppliedCouponDiscount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsAppliedCouponDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Amount is required.
type V1SubscriptionPreviewParamsAppliedCouponDiscountAmountsOff struct {
	Amount float64 `json:"amount,required"`
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

func (r V1SubscriptionPreviewParamsAppliedCouponDiscountAmountsOff) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsAppliedCouponDiscountAmountsOff
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsAppliedCouponDiscountAmountsOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionPreviewParamsAppliedCouponDiscountAmountsOff](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// The properties FeatureID, Quantity are required.
type V1SubscriptionPreviewParamsBillableFeature struct {
	// Feature ID
	FeatureID string  `json:"featureId,required"`
	Quantity  float64 `json:"quantity,required"`
	paramObj
}

func (r V1SubscriptionPreviewParamsBillableFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsBillableFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsBillableFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewParamsBillingInformation struct {
	ChargeOnBehalfOfAccount param.Opt[string]                                           `json:"chargeOnBehalfOfAccount,omitzero"`
	IntegrationID           param.Opt[string]                                           `json:"integrationId,omitzero"`
	InvoiceDaysUntilDue     param.Opt[float64]                                          `json:"invoiceDaysUntilDue,omitzero"`
	IsBackdated             param.Opt[bool]                                             `json:"isBackdated,omitzero"`
	IsInvoicePaid           param.Opt[bool]                                             `json:"isInvoicePaid,omitzero"`
	TaxPercentage           param.Opt[float64]                                          `json:"taxPercentage,omitzero"`
	BillingAddress          V1SubscriptionPreviewParamsBillingInformationBillingAddress `json:"billingAddress,omitzero"`
	Metadata                any                                                         `json:"metadata,omitzero"`
	// Any of "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE".
	ProrationBehavior string                                               `json:"prorationBehavior,omitzero"`
	TaxIDs            []V1SubscriptionPreviewParamsBillingInformationTaxID `json:"taxIds,omitzero"`
	TaxRateIDs        []string                                             `json:"taxRateIds,omitzero"`
	paramObj
}

func (r V1SubscriptionPreviewParamsBillingInformation) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsBillingInformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsBillingInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionPreviewParamsBillingInformation](
		"prorationBehavior", "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE",
	)
}

type V1SubscriptionPreviewParamsBillingInformationBillingAddress struct {
	City       param.Opt[string] `json:"city,omitzero"`
	Country    param.Opt[string] `json:"country,omitzero"`
	Line1      param.Opt[string] `json:"line1,omitzero"`
	Line2      param.Opt[string] `json:"line2,omitzero"`
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	State      param.Opt[string] `json:"state,omitzero"`
	paramObj
}

func (r V1SubscriptionPreviewParamsBillingInformationBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsBillingInformationBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsBillingInformationBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, Value are required.
type V1SubscriptionPreviewParamsBillingInformationTaxID struct {
	Type  string `json:"type,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r V1SubscriptionPreviewParamsBillingInformationTaxID) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsBillingInformationTaxID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsBillingInformationTaxID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionPreviewParamsBillingPeriod string

const (
	V1SubscriptionPreviewParamsBillingPeriodMonthly  V1SubscriptionPreviewParamsBillingPeriod = "MONTHLY"
	V1SubscriptionPreviewParamsBillingPeriodAnnually V1SubscriptionPreviewParamsBillingPeriod = "ANNUALLY"
)

// The properties ID, Quantity, Type are required.
type V1SubscriptionPreviewParamsCharge struct {
	// Charge ID
	ID string `json:"id,required"`
	// Charge quantity
	Quantity float64 `json:"quantity,required"`
	// Charge type
	//
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r V1SubscriptionPreviewParamsCharge) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionPreviewParamsCharge](
		"type", "FEATURE", "CREDIT",
	)
}

type V1SubscriptionPreviewParamsScheduleStrategy string

const (
	V1SubscriptionPreviewParamsScheduleStrategyEndOfBillingPeriod V1SubscriptionPreviewParamsScheduleStrategy = "END_OF_BILLING_PERIOD"
	V1SubscriptionPreviewParamsScheduleStrategyEndOfBillingMonth  V1SubscriptionPreviewParamsScheduleStrategy = "END_OF_BILLING_MONTH"
	V1SubscriptionPreviewParamsScheduleStrategyImmediate          V1SubscriptionPreviewParamsScheduleStrategy = "IMMEDIATE"
)

// The property IsTrial is required.
type V1SubscriptionPreviewParamsTrialOverrideConfiguration struct {
	IsTrial bool `json:"isTrial,required"`
	// Trial end date
	TrialEndDate param.Opt[time.Time] `json:"trialEndDate,omitzero" format:"date-time"`
	// Any of "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION".
	TrialEndBehavior string `json:"trialEndBehavior,omitzero"`
	paramObj
}

func (r V1SubscriptionPreviewParamsTrialOverrideConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsTrialOverrideConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsTrialOverrideConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionPreviewParamsTrialOverrideConfiguration](
		"trialEndBehavior", "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION",
	)
}

type V1SubscriptionTransferParams struct {
	// The resource ID to transfer the subscription to. The destination resource must
	// belong to the same customer.
	DestinationResourceID string `json:"destinationResourceId,required"`
	paramObj
}

func (r V1SubscriptionTransferParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionTransferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionTransferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
