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

// V1SubscriptionService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1SubscriptionService] method instead.
type V1SubscriptionService struct {
	Options      []option.RequestOption
	FutureUpdate V1SubscriptionFutureUpdateService
	Usage        V1SubscriptionUsageService
	Invoice      V1SubscriptionInvoiceService
}

// NewV1SubscriptionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1SubscriptionService(opts ...option.RequestOption) (r V1SubscriptionService) {
	r = V1SubscriptionService{}
	r.Options = opts
	r.FutureUpdate = NewV1SubscriptionFutureUpdateService(opts...)
	r.Usage = NewV1SubscriptionUsageService(opts...)
	r.Invoice = NewV1SubscriptionInvoiceService(opts...)
	return
}

// Retrieves a subscription by its unique identifier, including plan details,
// billing period, status, and add-ons.
func (r *V1SubscriptionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an active subscription's properties including billing period, add-ons,
// unit quantities, and discounts.
func (r *V1SubscriptionService) Update(ctx context.Context, id string, body V1SubscriptionUpdateParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of subscriptions, with optional filters for customer,
// status, and plan.
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

// Retrieves a paginated list of subscriptions, with optional filters for customer,
// status, and plan.
func (r *V1SubscriptionService) ListAutoPaging(ctx context.Context, query V1SubscriptionListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1SubscriptionListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, query, opts...))
}

// Cancels an active subscription, either immediately or at a specified time such
// as end of billing period.
func (r *V1SubscriptionService) Cancel(ctx context.Context, id string, body V1SubscriptionCancelParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delegates the payment responsibility of a subscription to a different customer.
// The delegated customer will be billed for this subscription.
func (r *V1SubscriptionService) Delegate(ctx context.Context, id string, body V1SubscriptionDelegateParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/delegate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Imports multiple subscriptions in bulk. Used for migrating subscription data
// from external systems.
func (r *V1SubscriptionService) Import(ctx context.Context, body V1SubscriptionImportParams, opts ...option.RequestOption) (res *V1SubscriptionImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/subscriptions/import"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Migrates a subscription to the latest published version of its plan or add-ons.
// Handles prorated charges or credits automatically.
func (r *V1SubscriptionService) Migrate(ctx context.Context, id string, body V1SubscriptionMigrateParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/migrate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Previews the pricing impact of creating or updating a subscription without
// making changes. Returns estimated costs, taxes, and proration details.
func (r *V1SubscriptionService) Preview(ctx context.Context, body V1SubscriptionPreviewParams, opts ...option.RequestOption) (res *V1SubscriptionPreviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/subscriptions/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a new subscription for an existing customer. When payment is required
// and no payment method exists, returns a checkout URL.
func (r *V1SubscriptionService) Provision(ctx context.Context, body V1SubscriptionProvisionParams, opts ...option.RequestOption) (res *V1SubscriptionProvisionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Transfers a subscription to a different resource ID. Used for multi-resource
// products where subscriptions apply to specific entities like websites or apps.
func (r *V1SubscriptionService) Transfer(ctx context.Context, id string, body V1SubscriptionTransferParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/transfer", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response object
type Subscription struct {
	// Customer subscription to a plan
	Data SubscriptionData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subscription) RawJSON() string { return r.JSON.raw }
func (r *Subscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer subscription to a plan
type SubscriptionData struct {
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
	PaymentCollectionMethod string                  `json:"paymentCollectionMethod,nullable"`
	Prices                  []SubscriptionDataPrice `json:"prices"`
	// Resource ID
	ResourceID string `json:"resourceId,nullable"`
	// Subscription trial end date
	TrialEndDate time.Time `json:"trialEndDate,nullable" format:"date-time"`
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
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionData) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionDataPrice struct {
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
func (r SubscriptionDataPrice) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionDataPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer subscription to a plan
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

// Response object
type V1SubscriptionImportResponse struct {
	Data V1SubscriptionImportResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionImportResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionImportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionImportResponseData struct {
	// Unique identifier for the import task
	TaskID string `json:"taskId,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TaskID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionImportResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionImportResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1SubscriptionPreviewResponse struct {
	// Pricing preview with invoices
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

// Pricing preview with invoices
type V1SubscriptionPreviewResponseData struct {
	// Invoice due immediately
	ImmediateInvoice V1SubscriptionPreviewResponseDataImmediateInvoice `json:"immediateInvoice,required"`
	// Billing period range
	BillingPeriodRange V1SubscriptionPreviewResponseDataBillingPeriodRange `json:"billingPeriodRange"`
	// Free items included
	FreeItems []V1SubscriptionPreviewResponseDataFreeItem `json:"freeItems"`
	// Whether updates are scheduled
	HasScheduledUpdates bool `json:"hasScheduledUpdates"`
	// Whether this is a downgrade
	IsPlanDowngrade bool `json:"isPlanDowngrade"`
	// Recurring invoice preview
	RecurringInvoice V1SubscriptionPreviewResponseDataRecurringInvoice `json:"recurringInvoice"`
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

// Invoice due immediately
type V1SubscriptionPreviewResponseDataImmediateInvoice struct {
	// Subtotal before discounts
	SubTotal float64 `json:"subTotal,required"`
	// Invoice total
	Total float64 `json:"total,required"`
	// Billing period covered
	BillingPeriodRange V1SubscriptionPreviewResponseDataImmediateInvoiceBillingPeriodRange `json:"billingPeriodRange"`
	// Currency code
	Currency string `json:"currency,nullable"`
	// Total discount amount
	Discount float64 `json:"discount"`
	// Discount breakdown
	DiscountDetails V1SubscriptionPreviewResponseDataImmediateInvoiceDiscountDetails `json:"discountDetails"`
	// Applied discounts
	Discounts []V1SubscriptionPreviewResponseDataImmediateInvoiceDiscount `json:"discounts"`
	// Line items
	Lines []V1SubscriptionPreviewResponseDataImmediateInvoiceLine `json:"lines"`
	// Tax amount
	Tax float64 `json:"tax"`
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

// Billing period covered
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

// Discount breakdown
type V1SubscriptionPreviewResponseDataImmediateInvoiceDiscountDetails struct {
	// Promo code used
	Code string `json:"code"`
	// Fixed discount amount
	FixedAmount float64 `json:"fixedAmount"`
	// Percentage discount
	Percentage float64 `json:"percentage"`
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

// Applied discount amount
type V1SubscriptionPreviewResponseDataImmediateInvoiceDiscount struct {
	// Discount amount
	Amount float64 `json:"amount,required"`
	// Currency code
	Currency string `json:"currency,required"`
	// Discount description
	Description string `json:"description,required"`
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

// Invoice line item
type V1SubscriptionPreviewResponseDataImmediateInvoiceLine struct {
	// Currency code
	Currency string `json:"currency,required"`
	// Line item description
	Description string `json:"description,required"`
	// Line subtotal
	SubTotal float64 `json:"subTotal,required"`
	// Price per unit
	UnitPrice float64 `json:"unitPrice,required"`
	// Quantity
	Quantity float64 `json:"quantity"`
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

// Billing period range
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

// Free item in subscription
type V1SubscriptionPreviewResponseDataFreeItem struct {
	// Addon ID
	AddonID string `json:"addonId,required"`
	// Quantity
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

// Recurring invoice preview
type V1SubscriptionPreviewResponseDataRecurringInvoice struct {
	// Subtotal before discounts
	SubTotal float64 `json:"subTotal,required"`
	// Invoice total
	Total float64 `json:"total,required"`
	// Billing period covered
	BillingPeriodRange V1SubscriptionPreviewResponseDataRecurringInvoiceBillingPeriodRange `json:"billingPeriodRange"`
	// Currency code
	Currency string `json:"currency,nullable"`
	// Total discount amount
	Discount float64 `json:"discount"`
	// Discount breakdown
	DiscountDetails V1SubscriptionPreviewResponseDataRecurringInvoiceDiscountDetails `json:"discountDetails"`
	// Applied discounts
	Discounts []V1SubscriptionPreviewResponseDataRecurringInvoiceDiscount `json:"discounts"`
	// Line items
	Lines []V1SubscriptionPreviewResponseDataRecurringInvoiceLine `json:"lines"`
	// Tax amount
	Tax float64 `json:"tax"`
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

// Billing period covered
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

// Discount breakdown
type V1SubscriptionPreviewResponseDataRecurringInvoiceDiscountDetails struct {
	// Promo code used
	Code string `json:"code"`
	// Fixed discount amount
	FixedAmount float64 `json:"fixedAmount"`
	// Percentage discount
	Percentage float64 `json:"percentage"`
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

// Applied discount amount
type V1SubscriptionPreviewResponseDataRecurringInvoiceDiscount struct {
	// Discount amount
	Amount float64 `json:"amount,required"`
	// Currency code
	Currency string `json:"currency,required"`
	// Discount description
	Description string `json:"description,required"`
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

// Invoice line item
type V1SubscriptionPreviewResponseDataRecurringInvoiceLine struct {
	// Currency code
	Currency string `json:"currency,required"`
	// Line item description
	Description string `json:"description,required"`
	// Line subtotal
	SubTotal float64 `json:"subTotal,required"`
	// Price per unit
	UnitPrice float64 `json:"unitPrice,required"`
	// Quantity
	Quantity float64 `json:"quantity"`
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

// Response object
type V1SubscriptionProvisionResponse struct {
	// Provisioning result with status and subscription or checkout URL.
	Data V1SubscriptionProvisionResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionProvisionResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionProvisionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning result with status and subscription or checkout URL.
type V1SubscriptionProvisionResponseData struct {
	// Unique identifier for the provisioned subscription
	ID           string                                                `json:"id,required"`
	Entitlements []V1SubscriptionProvisionResponseDataEntitlementUnion `json:"entitlements,required"`
	// Provision status: SUCCESS or PAYMENT_REQUIRED
	//
	// Any of "SUCCESS", "PAYMENT_REQUIRED".
	Status string `json:"status,required"`
	// Created subscription (when status is SUCCESS)
	Subscription V1SubscriptionProvisionResponseDataSubscription `json:"subscription,required"`
	// Checkout billing ID when payment is required
	CheckoutBillingID string `json:"checkoutBillingId"`
	// URL to complete payment when PAYMENT_REQUIRED
	CheckoutURL string `json:"checkoutUrl"`
	// Whether the subscription is scheduled for future activation
	IsScheduled bool `json:"isScheduled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Entitlements      respjson.Field
		Status            respjson.Field
		Subscription      respjson.Field
		CheckoutBillingID respjson.Field
		CheckoutURL       respjson.Field
		IsScheduled       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionProvisionResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionProvisionResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1SubscriptionProvisionResponseDataEntitlementUnion contains all possible
// properties and values from
// [V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0],
// [V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1SubscriptionProvisionResponseDataEntitlementUnion struct {
	AccessDeniedReason   string    `json:"accessDeniedReason"`
	IsGranted            bool      `json:"isGranted"`
	Type                 string    `json:"type"`
	CurrentUsage         float64   `json:"currentUsage"`
	EntitlementUpdatedAt time.Time `json:"entitlementUpdatedAt"`
	// This field is from variant
	// [V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0].
	Feature V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0Feature `json:"feature"`
	// This field is from variant
	// [V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0].
	HasUnlimitedUsage bool `json:"hasUnlimitedUsage"`
	// This field is from variant
	// [V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0].
	ResetPeriod string  `json:"resetPeriod"`
	UsageLimit  float64 `json:"usageLimit"`
	// This field is from variant
	// [V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0].
	UsagePeriodAnchor time.Time `json:"usagePeriodAnchor"`
	UsagePeriodEnd    time.Time `json:"usagePeriodEnd"`
	// This field is from variant
	// [V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0].
	UsagePeriodStart time.Time `json:"usagePeriodStart"`
	ValidUntil       time.Time `json:"validUntil"`
	// This field is from variant
	// [V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1].
	Currency V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1Currency `json:"currency"`
	// This field is from variant
	// [V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1].
	UsageUpdatedAt time.Time `json:"usageUpdatedAt"`
	JSON           struct {
		AccessDeniedReason   respjson.Field
		IsGranted            respjson.Field
		Type                 respjson.Field
		CurrentUsage         respjson.Field
		EntitlementUpdatedAt respjson.Field
		Feature              respjson.Field
		HasUnlimitedUsage    respjson.Field
		ResetPeriod          respjson.Field
		UsageLimit           respjson.Field
		UsagePeriodAnchor    respjson.Field
		UsagePeriodEnd       respjson.Field
		UsagePeriodStart     respjson.Field
		ValidUntil           respjson.Field
		Currency             respjson.Field
		UsageUpdatedAt       respjson.Field
		raw                  string
	} `json:"-"`
}

func (u V1SubscriptionProvisionResponseDataEntitlementUnion) AsV1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0() (v V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1SubscriptionProvisionResponseDataEntitlementUnion) AsV1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1() (v V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1SubscriptionProvisionResponseDataEntitlementUnion) RawJSON() string { return u.JSON.raw }

func (r *V1SubscriptionProvisionResponseDataEntitlementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0 struct {
	// Any of "FeatureNotFound", "CustomerNotFound", "CustomerIsArchived",
	// "CustomerResourceNotFound", "NoActiveSubscription",
	// "NoFeatureEntitlementInSubscription", "RequestedUsageExceedingLimit",
	// "RequestedValuesMismatch", "BudgetExceeded", "Unknown", "FeatureTypeMismatch",
	// "Revoked", "InsufficientCredits", "EntitlementNotFound".
	AccessDeniedReason string `json:"accessDeniedReason,required"`
	IsGranted          bool   `json:"isGranted,required"`
	// Any of "FEATURE".
	Type         string  `json:"type,required"`
	CurrentUsage float64 `json:"currentUsage"`
	// Timestamp of the last update to the entitlement grant or configuration.
	EntitlementUpdatedAt time.Time                                                                `json:"entitlementUpdatedAt" format:"date-time"`
	Feature              V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0Feature `json:"feature"`
	HasUnlimitedUsage    bool                                                                     `json:"hasUnlimitedUsage"`
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string  `json:"resetPeriod,nullable"`
	UsageLimit  float64 `json:"usageLimit,nullable"`
	// The anchor for calculating the usage period for metered entitlements with a
	// reset period configured
	UsagePeriodAnchor time.Time `json:"usagePeriodAnchor" format:"date-time"`
	// The end date of the usage period for metered entitlements with a reset period
	// configured
	UsagePeriodEnd time.Time `json:"usagePeriodEnd" format:"date-time"`
	// The start date of the usage period for metered entitlements with a reset period
	// configured
	UsagePeriodStart time.Time `json:"usagePeriodStart" format:"date-time"`
	// The next time the entitlement should be recalculated
	ValidUntil time.Time `json:"validUntil" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessDeniedReason   respjson.Field
		IsGranted            respjson.Field
		Type                 respjson.Field
		CurrentUsage         respjson.Field
		EntitlementUpdatedAt respjson.Field
		Feature              respjson.Field
		HasUnlimitedUsage    respjson.Field
		ResetPeriod          respjson.Field
		UsageLimit           respjson.Field
		UsagePeriodAnchor    respjson.Field
		UsagePeriodEnd       respjson.Field
		UsagePeriodStart     respjson.Field
		ValidUntil           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0Feature struct {
	// The human-readable name of the entitlement, shown in UI elements.
	DisplayName string `json:"displayName,required"`
	// The current status of the feature.
	//
	// Any of "NEW", "SUSPENDED", "ACTIVE".
	FeatureStatus string `json:"featureStatus,required"`
	// The type of feature associated with the entitlement.
	//
	// Any of "BOOLEAN", "NUMBER", "ENUM".
	FeatureType string `json:"featureType,required"`
	// The unique reference ID of the entitlement.
	RefID string `json:"refId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName   respjson.Field
		FeatureStatus respjson.Field
		FeatureType   respjson.Field
		RefID         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0Feature) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant0Feature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1 struct {
	// Any of "FeatureNotFound", "CustomerNotFound", "CustomerIsArchived",
	// "CustomerResourceNotFound", "NoActiveSubscription",
	// "NoFeatureEntitlementInSubscription", "RequestedUsageExceedingLimit",
	// "RequestedValuesMismatch", "BudgetExceeded", "Unknown", "FeatureTypeMismatch",
	// "Revoked", "InsufficientCredits", "EntitlementNotFound".
	AccessDeniedReason string `json:"accessDeniedReason,required"`
	// The currency associated with a credit entitlement.
	Currency     V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1Currency `json:"currency,required"`
	CurrentUsage float64                                                                   `json:"currentUsage,required"`
	IsGranted    bool                                                                      `json:"isGranted,required"`
	// Any of "CREDIT".
	Type       string  `json:"type,required"`
	UsageLimit float64 `json:"usageLimit,required"`
	// Timestamp of the last update to the credit usage.
	UsageUpdatedAt time.Time `json:"usageUpdatedAt,required" format:"date-time"`
	// Timestamp of the last update to the entitlement grant or configuration.
	EntitlementUpdatedAt time.Time `json:"entitlementUpdatedAt" format:"date-time"`
	// The end date of the current billing period for recurring credit grants.
	UsagePeriodEnd time.Time `json:"usagePeriodEnd" format:"date-time"`
	// The next time the entitlement should be recalculated
	ValidUntil time.Time `json:"validUntil" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessDeniedReason   respjson.Field
		Currency             respjson.Field
		CurrentUsage         respjson.Field
		IsGranted            respjson.Field
		Type                 respjson.Field
		UsageLimit           respjson.Field
		UsageUpdatedAt       respjson.Field
		EntitlementUpdatedAt respjson.Field
		UsagePeriodEnd       respjson.Field
		ValidUntil           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The currency associated with a credit entitlement.
type V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1Currency struct {
	// The unique identifier of the custom currency.
	CurrencyID string `json:"currencyId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1Currency) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionProvisionResponseDataEntitlementUnionObjectVariant1Currency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Created subscription (when status is SUCCESS)
type V1SubscriptionProvisionResponseDataSubscription struct {
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
	PaymentCollectionMethod string                                                 `json:"paymentCollectionMethod,nullable"`
	Prices                  []V1SubscriptionProvisionResponseDataSubscriptionPrice `json:"prices"`
	// Resource ID
	ResourceID string `json:"resourceId,nullable"`
	// Subscription trial end date
	TrialEndDate time.Time `json:"trialEndDate,nullable" format:"date-time"`
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
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionProvisionResponseDataSubscription) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionProvisionResponseDataSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionProvisionResponseDataSubscriptionPrice struct {
	// Addon identifier for the price override
	AddonID string `json:"addonId,nullable"`
	// Whether this is a base charge override
	BaseCharge bool `json:"baseCharge"`
	// Block size for pricing
	BlockSize float64 `json:"blockSize"`
	// Feature identifier for the price override
	FeatureID string `json:"featureId,nullable"`
	// Override price amount
	Price V1SubscriptionProvisionResponseDataSubscriptionPricePrice `json:"price"`
	// Pricing tiers configuration
	Tiers []V1SubscriptionProvisionResponseDataSubscriptionPriceTier `json:"tiers"`
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
func (r V1SubscriptionProvisionResponseDataSubscriptionPrice) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionProvisionResponseDataSubscriptionPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Override price amount
type V1SubscriptionProvisionResponseDataSubscriptionPricePrice struct {
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
func (r V1SubscriptionProvisionResponseDataSubscriptionPricePrice) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionProvisionResponseDataSubscriptionPricePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionProvisionResponseDataSubscriptionPriceTier struct {
	// The flat fee price of the price tier
	FlatPrice V1SubscriptionProvisionResponseDataSubscriptionPriceTierFlatPrice `json:"flatPrice"`
	// The unit price of the price tier
	UnitPrice V1SubscriptionProvisionResponseDataSubscriptionPriceTierUnitPrice `json:"unitPrice"`
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
func (r V1SubscriptionProvisionResponseDataSubscriptionPriceTier) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionProvisionResponseDataSubscriptionPriceTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The flat fee price of the price tier
type V1SubscriptionProvisionResponseDataSubscriptionPriceTierFlatPrice struct {
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
func (r V1SubscriptionProvisionResponseDataSubscriptionPriceTierFlatPrice) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionProvisionResponseDataSubscriptionPriceTierFlatPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The unit price of the price tier
type V1SubscriptionProvisionResponseDataSubscriptionPriceTierUnitPrice struct {
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
func (r V1SubscriptionProvisionResponseDataSubscriptionPriceTierUnitPrice) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionProvisionResponseDataSubscriptionPriceTierUnitPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionUpdateParams struct {
	AwaitPaymentConfirmation param.Opt[bool]   `json:"awaitPaymentConfirmation,omitzero"`
	PromotionCode            param.Opt[string] `json:"promotionCode,omitzero"`
	// Subscription trial end date
	TrialEndDate       param.Opt[time.Time]                         `json:"trialEndDate,omitzero" format:"date-time"`
	Budget             V1SubscriptionUpdateParamsBudget             `json:"budget,omitzero"`
	MinimumSpend       V1SubscriptionUpdateParamsMinimumSpend       `json:"minimumSpend,omitzero"`
	Addons             []V1SubscriptionUpdateParamsAddon            `json:"addons,omitzero"`
	AppliedCoupon      V1SubscriptionUpdateParamsAppliedCoupon      `json:"appliedCoupon,omitzero"`
	BillingInformation V1SubscriptionUpdateParamsBillingInformation `json:"billingInformation,omitzero"`
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod V1SubscriptionUpdateParamsBillingPeriod `json:"billingPeriod,omitzero"`
	Charges       []V1SubscriptionUpdateParamsCharge      `json:"charges,omitzero"`
	// Additional metadata for the subscription
	Metadata       map[string]string                         `json:"metadata,omitzero"`
	PriceOverrides []V1SubscriptionUpdateParamsPriceOverride `json:"priceOverrides,omitzero"`
	// Any of "END_OF_BILLING_PERIOD", "END_OF_BILLING_MONTH", "IMMEDIATE".
	ScheduleStrategy         V1SubscriptionUpdateParamsScheduleStrategy          `json:"scheduleStrategy,omitzero"`
	SubscriptionEntitlements []V1SubscriptionUpdateParamsSubscriptionEntitlement `json:"subscriptionEntitlements,omitzero"`
	paramObj
}

func (r V1SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AddonID, Quantity are required.
type V1SubscriptionUpdateParamsAddon struct {
	// Addon ID
	AddonID  string  `json:"addonId,required"`
	Quantity float64 `json:"quantity,required"`
	paramObj
}

func (r V1SubscriptionUpdateParamsAddon) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsAddon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsAddon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionUpdateParamsAppliedCoupon struct {
	PromotionCode   param.Opt[string]                                    `json:"promotionCode,omitzero"`
	BillingCouponID param.Opt[string]                                    `json:"billingCouponId,omitzero"`
	CouponID        param.Opt[string]                                    `json:"couponId,omitzero"`
	Configuration   V1SubscriptionUpdateParamsAppliedCouponConfiguration `json:"configuration,omitzero"`
	Discount        V1SubscriptionUpdateParamsAppliedCouponDiscount      `json:"discount,omitzero"`
	paramObj
}

func (r V1SubscriptionUpdateParamsAppliedCoupon) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsAppliedCoupon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsAppliedCoupon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionUpdateParamsAppliedCouponConfiguration struct {
	// Coupon start date
	StartDate param.Opt[time.Time] `json:"startDate,omitzero" format:"date-time"`
	paramObj
}

func (r V1SubscriptionUpdateParamsAppliedCouponConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsAppliedCouponConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsAppliedCouponConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionUpdateParamsAppliedCouponDiscount struct {
	Description      param.Opt[string]                                           `json:"description,omitzero"`
	DurationInMonths param.Opt[float64]                                          `json:"durationInMonths,omitzero"`
	Name             param.Opt[string]                                           `json:"name,omitzero"`
	PercentOff       param.Opt[float64]                                          `json:"percentOff,omitzero"`
	AmountsOff       []V1SubscriptionUpdateParamsAppliedCouponDiscountAmountsOff `json:"amountsOff,omitzero"`
	paramObj
}

func (r V1SubscriptionUpdateParamsAppliedCouponDiscount) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsAppliedCouponDiscount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsAppliedCouponDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Amount is required.
type V1SubscriptionUpdateParamsAppliedCouponDiscountAmountsOff struct {
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

func (r V1SubscriptionUpdateParamsAppliedCouponDiscountAmountsOff) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsAppliedCouponDiscountAmountsOff
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsAppliedCouponDiscountAmountsOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionUpdateParamsAppliedCouponDiscountAmountsOff](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

type V1SubscriptionUpdateParamsBillingInformation struct {
	ChargeOnBehalfOfAccount param.Opt[string]  `json:"chargeOnBehalfOfAccount,omitzero"`
	CouponID                param.Opt[string]  `json:"couponId,omitzero"`
	IntegrationID           param.Opt[string]  `json:"integrationId,omitzero"`
	InvoiceDaysUntilDue     param.Opt[float64] `json:"invoiceDaysUntilDue,omitzero"`
	IsBackdated             param.Opt[bool]    `json:"isBackdated,omitzero"`
	IsInvoicePaid           param.Opt[bool]    `json:"isInvoicePaid,omitzero"`
	TaxPercentage           param.Opt[float64] `json:"taxPercentage,omitzero"`
	// Physical address
	BillingAddress V1SubscriptionUpdateParamsBillingInformationBillingAddress `json:"billingAddress,omitzero"`
	// Additional metadata for the subscription
	Metadata map[string]any `json:"metadata,omitzero"`
	// Any of "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE".
	ProrationBehavior string                                              `json:"prorationBehavior,omitzero"`
	TaxIDs            []V1SubscriptionUpdateParamsBillingInformationTaxID `json:"taxIds,omitzero"`
	TaxRateIDs        []string                                            `json:"taxRateIds,omitzero"`
	paramObj
}

func (r V1SubscriptionUpdateParamsBillingInformation) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsBillingInformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsBillingInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionUpdateParamsBillingInformation](
		"prorationBehavior", "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE",
	)
}

// Physical address
type V1SubscriptionUpdateParamsBillingInformationBillingAddress struct {
	// City name
	City param.Opt[string] `json:"city,omitzero"`
	// Country code or name
	Country param.Opt[string] `json:"country,omitzero"`
	// Street address line 1
	Line1 param.Opt[string] `json:"line1,omitzero"`
	// Street address line 2
	Line2 param.Opt[string] `json:"line2,omitzero"`
	// Postal or ZIP code
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// State or province
	State param.Opt[string] `json:"state,omitzero"`
	paramObj
}

func (r V1SubscriptionUpdateParamsBillingInformationBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsBillingInformationBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsBillingInformationBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, Value are required.
type V1SubscriptionUpdateParamsBillingInformationTaxID struct {
	Type  string `json:"type,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r V1SubscriptionUpdateParamsBillingInformationTaxID) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsBillingInformationTaxID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsBillingInformationTaxID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionUpdateParamsBillingPeriod string

const (
	V1SubscriptionUpdateParamsBillingPeriodMonthly  V1SubscriptionUpdateParamsBillingPeriod = "MONTHLY"
	V1SubscriptionUpdateParamsBillingPeriodAnnually V1SubscriptionUpdateParamsBillingPeriod = "ANNUALLY"
)

// The properties HasSoftLimit, Limit are required.
type V1SubscriptionUpdateParamsBudget struct {
	HasSoftLimit bool    `json:"hasSoftLimit,required"`
	Limit        float64 `json:"limit,required"`
	paramObj
}

func (r V1SubscriptionUpdateParamsBudget) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsBudget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Quantity, Type are required.
type V1SubscriptionUpdateParamsCharge struct {
	// Charge ID
	ID       string  `json:"id,required"`
	Quantity float64 `json:"quantity,required"`
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r V1SubscriptionUpdateParamsCharge) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionUpdateParamsCharge](
		"type", "FEATURE", "CREDIT",
	)
}

type V1SubscriptionUpdateParamsMinimumSpend struct {
	Minimum V1SubscriptionUpdateParamsMinimumSpendMinimum `json:"minimum,omitzero"`
	paramObj
}

func (r V1SubscriptionUpdateParamsMinimumSpend) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsMinimumSpend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsMinimumSpend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Amount is required.
type V1SubscriptionUpdateParamsMinimumSpendMinimum struct {
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

func (r V1SubscriptionUpdateParamsMinimumSpendMinimum) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsMinimumSpendMinimum
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsMinimumSpendMinimum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionUpdateParamsMinimumSpendMinimum](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

type V1SubscriptionUpdateParamsPriceOverride struct {
	// Addon ID
	AddonID param.Opt[string] `json:"addonId,omitzero"`
	// Whether this is a base charge override
	BaseCharge param.Opt[bool] `json:"baseCharge,omitzero"`
	// The corresponding custom currency id of the recurring credits price
	CurrencyID param.Opt[string] `json:"currencyId,omitzero"`
	// Feature ID
	FeatureID param.Opt[string]                            `json:"featureId,omitzero"`
	Price     V1SubscriptionUpdateParamsPriceOverridePrice `json:"price,omitzero"`
	paramObj
}

func (r V1SubscriptionUpdateParamsPriceOverride) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsPriceOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsPriceOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Amount is required.
type V1SubscriptionUpdateParamsPriceOverridePrice struct {
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

func (r V1SubscriptionUpdateParamsPriceOverridePrice) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsPriceOverridePrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsPriceOverridePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionUpdateParamsPriceOverridePrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

type V1SubscriptionUpdateParamsScheduleStrategy string

const (
	V1SubscriptionUpdateParamsScheduleStrategyEndOfBillingPeriod V1SubscriptionUpdateParamsScheduleStrategy = "END_OF_BILLING_PERIOD"
	V1SubscriptionUpdateParamsScheduleStrategyEndOfBillingMonth  V1SubscriptionUpdateParamsScheduleStrategy = "END_OF_BILLING_MONTH"
	V1SubscriptionUpdateParamsScheduleStrategyImmediate          V1SubscriptionUpdateParamsScheduleStrategy = "IMMEDIATE"
)

type V1SubscriptionUpdateParamsSubscriptionEntitlement struct {
	ID                              param.Opt[string]                                                                `json:"id,omitzero"`
	FeatureID                       param.Opt[string]                                                                `json:"featureId,omitzero"`
	HasSoftLimit                    param.Opt[bool]                                                                  `json:"hasSoftLimit,omitzero"`
	HasUnlimitedUsage               param.Opt[bool]                                                                  `json:"hasUnlimitedUsage,omitzero"`
	UsageLimit                      param.Opt[float64]                                                               `json:"usageLimit,omitzero"`
	MonthlyResetPeriodConfiguration V1SubscriptionUpdateParamsSubscriptionEntitlementMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod                    string                                                                          `json:"resetPeriod,omitzero"`
	WeeklyResetPeriodConfiguration V1SubscriptionUpdateParamsSubscriptionEntitlementWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	YearlyResetPeriodConfiguration V1SubscriptionUpdateParamsSubscriptionEntitlementYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
	paramObj
}

func (r V1SubscriptionUpdateParamsSubscriptionEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsSubscriptionEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsSubscriptionEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionUpdateParamsSubscriptionEntitlement](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// The property AccordingTo is required.
type V1SubscriptionUpdateParamsSubscriptionEntitlementMonthlyResetPeriodConfiguration struct {
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero,required"`
	paramObj
}

func (r V1SubscriptionUpdateParamsSubscriptionEntitlementMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsSubscriptionEntitlementMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsSubscriptionEntitlementMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionUpdateParamsSubscriptionEntitlementMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// The property AccordingTo is required.
type V1SubscriptionUpdateParamsSubscriptionEntitlementWeeklyResetPeriodConfiguration struct {
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero,required"`
	paramObj
}

func (r V1SubscriptionUpdateParamsSubscriptionEntitlementWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsSubscriptionEntitlementWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsSubscriptionEntitlementWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionUpdateParamsSubscriptionEntitlementWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// The property AccordingTo is required.
type V1SubscriptionUpdateParamsSubscriptionEntitlementYearlyResetPeriodConfiguration struct {
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero,required"`
	paramObj
}

func (r V1SubscriptionUpdateParamsSubscriptionEntitlementYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionUpdateParamsSubscriptionEntitlementYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionUpdateParamsSubscriptionEntitlementYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionUpdateParamsSubscriptionEntitlementYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

type V1SubscriptionListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Filter by customer ID
	CustomerID param.Opt[string] `query:"customerId,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by plan ID
	PlanID param.Opt[string] `query:"planId,omitzero" json:"-"`
	// Filter by pricing type. Supports comma-separated values for multiple types
	PricingType param.Opt[string] `query:"pricingType,omitzero" json:"-"`
	// Filter by resource ID
	ResourceID param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	// Filter by subscription status. Supports comma-separated values for multiple
	// statuses
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter by creation date using range operators: gt, gte, lt, lte
	CreatedAt V1SubscriptionListParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
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

// Filter by creation date using range operators: gt, gte, lt, lte
type V1SubscriptionListParamsCreatedAt struct {
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

// URLQuery serializes [V1SubscriptionListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r V1SubscriptionListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1SubscriptionCancelParams struct {
	// Subscription end date
	EndDate param.Opt[time.Time] `json:"endDate,omitzero" format:"date-time"`
	// If set, enables or disables prorating of credits on subscription cancellation.
	Prorate param.Opt[bool] `json:"prorate,omitzero"`
	// Action on cancellation (downgrade or revoke)
	//
	// Any of "DEFAULT", "REVOKE_ENTITLEMENTS".
	CancellationAction V1SubscriptionCancelParamsCancellationAction `json:"cancellationAction,omitzero"`
	// When to cancel (immediate, period end, or date)
	//
	// Any of "END_OF_BILLING_PERIOD", "IMMEDIATE", "SPECIFIC_DATE".
	CancellationTime V1SubscriptionCancelParamsCancellationTime `json:"cancellationTime,omitzero"`
	paramObj
}

func (r V1SubscriptionCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action on cancellation (downgrade or revoke)
type V1SubscriptionCancelParamsCancellationAction string

const (
	V1SubscriptionCancelParamsCancellationActionDefault            V1SubscriptionCancelParamsCancellationAction = "DEFAULT"
	V1SubscriptionCancelParamsCancellationActionRevokeEntitlements V1SubscriptionCancelParamsCancellationAction = "REVOKE_ENTITLEMENTS"
)

// When to cancel (immediate, period end, or date)
type V1SubscriptionCancelParamsCancellationTime string

const (
	V1SubscriptionCancelParamsCancellationTimeEndOfBillingPeriod V1SubscriptionCancelParamsCancellationTime = "END_OF_BILLING_PERIOD"
	V1SubscriptionCancelParamsCancellationTimeImmediate          V1SubscriptionCancelParamsCancellationTime = "IMMEDIATE"
	V1SubscriptionCancelParamsCancellationTimeSpecificDate       V1SubscriptionCancelParamsCancellationTime = "SPECIFIC_DATE"
)

type V1SubscriptionDelegateParams struct {
	// The unique identifier of the customer who will assume payment responsibility for
	// this subscription. This customer must already exist in your Stigg account and
	// have a valid payment method if the subscription requires payment.
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

type V1SubscriptionImportParams struct {
	// List of subscription objects to import
	Subscriptions []V1SubscriptionImportParamsSubscription `json:"subscriptions,omitzero,required"`
	// Integration ID to use for importing subscriptions
	IntegrationID param.Opt[string] `json:"integrationId,omitzero"`
	paramObj
}

func (r V1SubscriptionImportParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionImportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionImportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, CustomerID, PlanID are required.
type V1SubscriptionImportParamsSubscription struct {
	// Subscription ID
	ID string `json:"id,required"`
	// Customer ID
	CustomerID string `json:"customerId,required"`
	// Plan ID
	PlanID string `json:"planId,required"`
	// Billing ID
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// Subscription end date
	EndDate param.Opt[time.Time] `json:"endDate,omitzero" format:"date-time"`
	// Resource ID
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Subscription start date
	StartDate param.Opt[time.Time] `json:"startDate,omitzero" format:"date-time"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1SubscriptionImportParamsSubscription) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionImportParamsSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionImportParamsSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionMigrateParams struct {
	// When to migrate (immediate or period end)
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

// When to migrate (immediate or period end)
type V1SubscriptionMigrateParamsSubscriptionMigrationTime string

const (
	V1SubscriptionMigrateParamsSubscriptionMigrationTimeEndOfBillingPeriod V1SubscriptionMigrateParamsSubscriptionMigrationTime = "END_OF_BILLING_PERIOD"
	V1SubscriptionMigrateParamsSubscriptionMigrationTimeImmediate          V1SubscriptionMigrateParamsSubscriptionMigrationTime = "IMMEDIATE"
)

type V1SubscriptionPreviewParams struct {
	// Customer ID
	CustomerID string `json:"customerId,required"`
	// Plan ID
	PlanID string `json:"planId,required"`
	// ISO 3166-1 country code for localization
	BillingCountryCode param.Opt[string] `json:"billingCountryCode,omitzero"`
	// Paying customer ID for delegated billing
	PayingCustomerID param.Opt[string] `json:"payingCustomerId,omitzero"`
	// Resource ID for multi-instance subscriptions
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Subscription start date
	StartDate param.Opt[time.Time] `json:"startDate,omitzero" format:"date-time"`
	// Unit quantity for per-unit pricing
	UnitQuantity param.Opt[float64] `json:"unitQuantity,omitzero"`
	// Addons to include
	Addons []V1SubscriptionPreviewParamsAddon `json:"addons,omitzero"`
	// Coupon or discount to apply
	AppliedCoupon V1SubscriptionPreviewParamsAppliedCoupon `json:"appliedCoupon,omitzero"`
	// Billable features with quantities
	BillableFeatures []V1SubscriptionPreviewParamsBillableFeature `json:"billableFeatures,omitzero"`
	// Billing and tax configuration
	BillingInformation V1SubscriptionPreviewParamsBillingInformation `json:"billingInformation,omitzero"`
	// Billing period (MONTHLY or ANNUALLY)
	//
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod V1SubscriptionPreviewParamsBillingPeriod `json:"billingPeriod,omitzero"`
	// One-time or recurring charges
	Charges []V1SubscriptionPreviewParamsCharge `json:"charges,omitzero"`
	// When to apply subscription changes
	//
	// Any of "END_OF_BILLING_PERIOD", "END_OF_BILLING_MONTH", "IMMEDIATE".
	ScheduleStrategy V1SubscriptionPreviewParamsScheduleStrategy `json:"scheduleStrategy,omitzero"`
	// Trial period override settings
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

// Addon configuration
//
// The property AddonID is required.
type V1SubscriptionPreviewParamsAddon struct {
	// Addon ID
	AddonID string `json:"addonId,required"`
	// Number of addon instances
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

// Coupon or discount to apply
type V1SubscriptionPreviewParamsAppliedCoupon struct {
	// Billing provider coupon ID
	BillingCouponID param.Opt[string] `json:"billingCouponId,omitzero"`
	// Stigg coupon ID
	CouponID param.Opt[string] `json:"couponId,omitzero"`
	// Promotion code to apply
	PromotionCode param.Opt[string] `json:"promotionCode,omitzero"`
	// Coupon timing configuration
	Configuration V1SubscriptionPreviewParamsAppliedCouponConfiguration `json:"configuration,omitzero"`
	// Ad-hoc discount configuration
	Discount V1SubscriptionPreviewParamsAppliedCouponDiscount `json:"discount,omitzero"`
	paramObj
}

func (r V1SubscriptionPreviewParamsAppliedCoupon) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsAppliedCoupon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsAppliedCoupon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coupon timing configuration
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

// Ad-hoc discount configuration
type V1SubscriptionPreviewParamsAppliedCouponDiscount struct {
	// Ad-hoc discount
	Description param.Opt[string] `json:"description,omitzero"`
	// Duration in months
	DurationInMonths param.Opt[float64] `json:"durationInMonths,omitzero"`
	// Discount name
	Name param.Opt[string] `json:"name,omitzero"`
	// Percentage discount
	PercentOff param.Opt[float64] `json:"percentOff,omitzero"`
	// Fixed amounts off by currency
	AmountsOff []V1SubscriptionPreviewParamsAppliedCouponDiscountAmountsOff `json:"amountsOff,omitzero"`
	paramObj
}

func (r V1SubscriptionPreviewParamsAppliedCouponDiscount) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsAppliedCouponDiscount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsAppliedCouponDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Currency are required.
type V1SubscriptionPreviewParamsAppliedCouponDiscountAmountsOff struct {
	// The price amount
	Amount float64 `json:"amount,required"`
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
	Currency string `json:"currency,omitzero,required"`
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

// Feature with quantity
//
// The properties FeatureID, Quantity are required.
type V1SubscriptionPreviewParamsBillableFeature struct {
	// Feature ID
	FeatureID string `json:"featureId,required"`
	// Quantity of feature units
	Quantity float64 `json:"quantity,required"`
	paramObj
}

func (r V1SubscriptionPreviewParamsBillableFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionPreviewParamsBillableFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionPreviewParamsBillableFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing and tax configuration
type V1SubscriptionPreviewParamsBillingInformation struct {
	// Connected account ID for platform billing
	ChargeOnBehalfOfAccount param.Opt[string] `json:"chargeOnBehalfOfAccount,omitzero"`
	// Billing integration ID
	IntegrationID param.Opt[string] `json:"integrationId,omitzero"`
	// Days until invoice is due
	InvoiceDaysUntilDue param.Opt[float64] `json:"invoiceDaysUntilDue,omitzero"`
	// Whether subscription is backdated
	IsBackdated param.Opt[bool] `json:"isBackdated,omitzero"`
	// Whether invoice is already paid
	IsInvoicePaid param.Opt[bool] `json:"isInvoicePaid,omitzero"`
	// Tax percentage to apply
	TaxPercentage param.Opt[float64] `json:"taxPercentage,omitzero"`
	// Billing address
	BillingAddress V1SubscriptionPreviewParamsBillingInformationBillingAddress `json:"billingAddress,omitzero"`
	// Additional billing metadata
	Metadata any `json:"metadata,omitzero"`
	// Proration behavior
	//
	// Any of "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE".
	ProrationBehavior string `json:"prorationBehavior,omitzero"`
	// Customer tax IDs
	TaxIDs []V1SubscriptionPreviewParamsBillingInformationTaxID `json:"taxIds,omitzero"`
	// Tax rate IDs from billing provider
	TaxRateIDs []string `json:"taxRateIds,omitzero"`
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

// Billing address
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

// Tax exemption identifier
//
// The properties Type, Value are required.
type V1SubscriptionPreviewParamsBillingInformationTaxID struct {
	// Tax exemption type (e.g., vat, gst)
	Type string `json:"type,required"`
	// Tax exemption identifier value
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

// Billing period (MONTHLY or ANNUALLY)
type V1SubscriptionPreviewParamsBillingPeriod string

const (
	V1SubscriptionPreviewParamsBillingPeriodMonthly  V1SubscriptionPreviewParamsBillingPeriod = "MONTHLY"
	V1SubscriptionPreviewParamsBillingPeriodAnnually V1SubscriptionPreviewParamsBillingPeriod = "ANNUALLY"
)

// Charge item
//
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

// When to apply subscription changes
type V1SubscriptionPreviewParamsScheduleStrategy string

const (
	V1SubscriptionPreviewParamsScheduleStrategyEndOfBillingPeriod V1SubscriptionPreviewParamsScheduleStrategy = "END_OF_BILLING_PERIOD"
	V1SubscriptionPreviewParamsScheduleStrategyEndOfBillingMonth  V1SubscriptionPreviewParamsScheduleStrategy = "END_OF_BILLING_MONTH"
	V1SubscriptionPreviewParamsScheduleStrategyImmediate          V1SubscriptionPreviewParamsScheduleStrategy = "IMMEDIATE"
)

// Trial period override settings
//
// The property IsTrial is required.
type V1SubscriptionPreviewParamsTrialOverrideConfiguration struct {
	// Whether to start as trial
	IsTrial bool `json:"isTrial,required"`
	// Trial end date
	TrialEndDate param.Opt[time.Time] `json:"trialEndDate,omitzero" format:"date-time"`
	// Behavior when trial ends
	//
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

type V1SubscriptionProvisionParams struct {
	// Customer ID to provision the subscription for
	CustomerID string `json:"customerId,required"`
	// Plan ID to provision
	PlanID string `json:"planId,required"`
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
	// Unique identifier for the subscription
	ID param.Opt[string] `json:"id,omitzero"`
	// Whether to wait for payment confirmation before returning the subscription
	AwaitPaymentConfirmation param.Opt[bool] `json:"awaitPaymentConfirmation,omitzero"`
	// Subscription start date
	StartDate    param.Opt[time.Time]                      `json:"startDate,omitzero" format:"date-time"`
	UnitQuantity param.Opt[float64]                        `json:"unitQuantity,omitzero"`
	Budget       V1SubscriptionProvisionParamsBudget       `json:"budget,omitzero"`
	MinimumSpend V1SubscriptionProvisionParamsMinimumSpend `json:"minimumSpend,omitzero"`
	Addons       []V1SubscriptionProvisionParamsAddon      `json:"addons,omitzero"`
	// Coupon configuration
	AppliedCoupon      V1SubscriptionProvisionParamsAppliedCoupon      `json:"appliedCoupon,omitzero"`
	BillingInformation V1SubscriptionProvisionParamsBillingInformation `json:"billingInformation,omitzero"`
	// Billing period (MONTHLY or ANNUALLY)
	//
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod V1SubscriptionProvisionParamsBillingPeriod `json:"billingPeriod,omitzero"`
	Charges       []V1SubscriptionProvisionParamsCharge      `json:"charges,omitzero"`
	// Checkout page configuration for payment collection
	CheckoutOptions V1SubscriptionProvisionParamsCheckoutOptions `json:"checkoutOptions,omitzero"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata,omitzero"`
	// How payments should be collected for this subscription
	//
	// Any of "CHARGE", "INVOICE", "NONE".
	PaymentCollectionMethod V1SubscriptionProvisionParamsPaymentCollectionMethod `json:"paymentCollectionMethod,omitzero"`
	PriceOverrides          []V1SubscriptionProvisionParamsPriceOverride         `json:"priceOverrides,omitzero"`
	// Strategy for scheduling subscription changes
	//
	// Any of "END_OF_BILLING_PERIOD", "END_OF_BILLING_MONTH", "IMMEDIATE".
	ScheduleStrategy         V1SubscriptionProvisionParamsScheduleStrategy          `json:"scheduleStrategy,omitzero"`
	SubscriptionEntitlements []V1SubscriptionProvisionParamsSubscriptionEntitlement `json:"subscriptionEntitlements,omitzero"`
	// Trial period override settings
	TrialOverrideConfiguration V1SubscriptionProvisionParamsTrialOverrideConfiguration `json:"trialOverrideConfiguration,omitzero"`
	paramObj
}

func (r V1SubscriptionProvisionParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property AddonID is required.
type V1SubscriptionProvisionParamsAddon struct {
	// Addon identifier
	AddonID string `json:"addonId,required"`
	// Number of addon units
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	paramObj
}

func (r V1SubscriptionProvisionParamsAddon) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsAddon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsAddon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coupon configuration
type V1SubscriptionProvisionParamsAppliedCoupon struct {
	// Billing provider coupon ID
	BillingCouponID param.Opt[string] `json:"billingCouponId,omitzero"`
	// Stigg coupon ID
	CouponID param.Opt[string] `json:"couponId,omitzero"`
	// Promotion code to apply
	PromotionCode param.Opt[string] `json:"promotionCode,omitzero"`
	// Coupon timing configuration
	Configuration V1SubscriptionProvisionParamsAppliedCouponConfiguration `json:"configuration,omitzero"`
	// Ad-hoc discount configuration
	Discount V1SubscriptionProvisionParamsAppliedCouponDiscount `json:"discount,omitzero"`
	paramObj
}

func (r V1SubscriptionProvisionParamsAppliedCoupon) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsAppliedCoupon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsAppliedCoupon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coupon timing configuration
type V1SubscriptionProvisionParamsAppliedCouponConfiguration struct {
	// Coupon start date
	StartDate param.Opt[time.Time] `json:"startDate,omitzero" format:"date-time"`
	paramObj
}

func (r V1SubscriptionProvisionParamsAppliedCouponConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsAppliedCouponConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsAppliedCouponConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ad-hoc discount configuration
type V1SubscriptionProvisionParamsAppliedCouponDiscount struct {
	// Ad-hoc discount
	Description param.Opt[string] `json:"description,omitzero"`
	// Duration in months
	DurationInMonths param.Opt[float64] `json:"durationInMonths,omitzero"`
	// Discount name
	Name param.Opt[string] `json:"name,omitzero"`
	// Percentage discount
	PercentOff param.Opt[float64] `json:"percentOff,omitzero"`
	// Fixed amounts off by currency
	AmountsOff []V1SubscriptionProvisionParamsAppliedCouponDiscountAmountsOff `json:"amountsOff,omitzero"`
	paramObj
}

func (r V1SubscriptionProvisionParamsAppliedCouponDiscount) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsAppliedCouponDiscount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsAppliedCouponDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Currency are required.
type V1SubscriptionProvisionParamsAppliedCouponDiscountAmountsOff struct {
	// The price amount
	Amount float64 `json:"amount,required"`
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
	Currency string `json:"currency,omitzero,required"`
	paramObj
}

func (r V1SubscriptionProvisionParamsAppliedCouponDiscountAmountsOff) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsAppliedCouponDiscountAmountsOff
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsAppliedCouponDiscountAmountsOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionProvisionParamsAppliedCouponDiscountAmountsOff](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

type V1SubscriptionProvisionParamsBillingInformation struct {
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
	BillingAddress V1SubscriptionProvisionParamsBillingInformationBillingAddress `json:"billingAddress,omitzero"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata,omitzero"`
	// How to handle proration for billing changes
	//
	// Any of "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE".
	ProrationBehavior string `json:"prorationBehavior,omitzero"`
	// Customer tax identification numbers
	TaxIDs []V1SubscriptionProvisionParamsBillingInformationTaxID `json:"taxIds,omitzero"`
	// Tax rate identifiers to apply
	TaxRateIDs []string `json:"taxRateIds,omitzero"`
	paramObj
}

func (r V1SubscriptionProvisionParamsBillingInformation) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsBillingInformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsBillingInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionProvisionParamsBillingInformation](
		"prorationBehavior", "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE",
	)
}

// Billing address for the subscription
type V1SubscriptionProvisionParamsBillingInformationBillingAddress struct {
	City       param.Opt[string] `json:"city,omitzero"`
	Country    param.Opt[string] `json:"country,omitzero"`
	Line1      param.Opt[string] `json:"line1,omitzero"`
	Line2      param.Opt[string] `json:"line2,omitzero"`
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	State      param.Opt[string] `json:"state,omitzero"`
	paramObj
}

func (r V1SubscriptionProvisionParamsBillingInformationBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsBillingInformationBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsBillingInformationBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, Value are required.
type V1SubscriptionProvisionParamsBillingInformationTaxID struct {
	// The type of tax exemption identifier, such as VAT.
	Type string `json:"type,required"`
	// The actual tax identifier value
	Value string `json:"value,required"`
	paramObj
}

func (r V1SubscriptionProvisionParamsBillingInformationTaxID) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsBillingInformationTaxID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsBillingInformationTaxID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing period (MONTHLY or ANNUALLY)
type V1SubscriptionProvisionParamsBillingPeriod string

const (
	V1SubscriptionProvisionParamsBillingPeriodMonthly  V1SubscriptionProvisionParamsBillingPeriod = "MONTHLY"
	V1SubscriptionProvisionParamsBillingPeriodAnnually V1SubscriptionProvisionParamsBillingPeriod = "ANNUALLY"
)

// The properties HasSoftLimit, Limit are required.
type V1SubscriptionProvisionParamsBudget struct {
	// Whether the budget is a soft limit
	HasSoftLimit bool `json:"hasSoftLimit,required"`
	// Maximum spending limit
	Limit float64 `json:"limit,required"`
	paramObj
}

func (r V1SubscriptionProvisionParamsBudget) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsBudget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Charge item
//
// The properties ID, Quantity, Type are required.
type V1SubscriptionProvisionParamsCharge struct {
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

func (r V1SubscriptionProvisionParamsCharge) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionProvisionParamsCharge](
		"type", "FEATURE", "CREDIT",
	)
}

// Checkout page configuration for payment collection
//
// The properties CancelURL, SuccessURL are required.
type V1SubscriptionProvisionParamsCheckoutOptions struct {
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

func (r V1SubscriptionProvisionParamsCheckoutOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsCheckoutOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsCheckoutOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionProvisionParamsMinimumSpend struct {
	// Minimum spend amount
	Minimum V1SubscriptionProvisionParamsMinimumSpendMinimum `json:"minimum,omitzero"`
	paramObj
}

func (r V1SubscriptionProvisionParamsMinimumSpend) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsMinimumSpend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsMinimumSpend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Minimum spend amount
type V1SubscriptionProvisionParamsMinimumSpendMinimum struct {
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

func (r V1SubscriptionProvisionParamsMinimumSpendMinimum) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsMinimumSpendMinimum
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsMinimumSpendMinimum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionProvisionParamsMinimumSpendMinimum](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// How payments should be collected for this subscription
type V1SubscriptionProvisionParamsPaymentCollectionMethod string

const (
	V1SubscriptionProvisionParamsPaymentCollectionMethodCharge  V1SubscriptionProvisionParamsPaymentCollectionMethod = "CHARGE"
	V1SubscriptionProvisionParamsPaymentCollectionMethodInvoice V1SubscriptionProvisionParamsPaymentCollectionMethod = "INVOICE"
	V1SubscriptionProvisionParamsPaymentCollectionMethodNone    V1SubscriptionProvisionParamsPaymentCollectionMethod = "NONE"
)

type V1SubscriptionProvisionParamsPriceOverride struct {
	// Addon identifier for the price override
	AddonID param.Opt[string] `json:"addonId,omitzero"`
	// Feature identifier for the price override
	FeatureID param.Opt[string] `json:"featureId,omitzero"`
	// Whether this is a base charge override
	BaseCharge param.Opt[bool] `json:"baseCharge,omitzero"`
	// Block size for pricing
	BlockSize param.Opt[float64] `json:"blockSize,omitzero"`
	// Any of "BEGINNING_OF_BILLING_PERIOD", "MONTHLY".
	CreditGrantCadence string                                               `json:"creditGrantCadence,omitzero"`
	CreditRate         V1SubscriptionProvisionParamsPriceOverrideCreditRate `json:"creditRate,omitzero"`
	// Override price amount
	Price V1SubscriptionProvisionParamsPriceOverridePrice `json:"price,omitzero"`
	// Pricing tiers configuration
	Tiers []V1SubscriptionProvisionParamsPriceOverrideTier `json:"tiers,omitzero"`
	paramObj
}

func (r V1SubscriptionProvisionParamsPriceOverride) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsPriceOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsPriceOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionProvisionParamsPriceOverride](
		"creditGrantCadence", "BEGINNING_OF_BILLING_PERIOD", "MONTHLY",
	)
}

// The properties Amount, CurrencyID are required.
type V1SubscriptionProvisionParamsPriceOverrideCreditRate struct {
	// The credit rate amount
	Amount float64 `json:"amount,required"`
	// The custom currency refId for the credit rate
	CurrencyID string `json:"currencyId,required"`
	// A custom formula for calculating cost based on single event dimensions
	CostFormula param.Opt[string] `json:"costFormula,omitzero"`
	paramObj
}

func (r V1SubscriptionProvisionParamsPriceOverrideCreditRate) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsPriceOverrideCreditRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsPriceOverrideCreditRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Override price amount
type V1SubscriptionProvisionParamsPriceOverridePrice struct {
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

func (r V1SubscriptionProvisionParamsPriceOverridePrice) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsPriceOverridePrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsPriceOverridePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionProvisionParamsPriceOverridePrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

type V1SubscriptionProvisionParamsPriceOverrideTier struct {
	// The up to quantity of the price tier
	UpTo param.Opt[float64] `json:"upTo,omitzero"`
	// The flat fee price of the price tier
	FlatPrice V1SubscriptionProvisionParamsPriceOverrideTierFlatPrice `json:"flatPrice,omitzero"`
	// The unit price of the price tier
	UnitPrice V1SubscriptionProvisionParamsPriceOverrideTierUnitPrice `json:"unitPrice,omitzero"`
	paramObj
}

func (r V1SubscriptionProvisionParamsPriceOverrideTier) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsPriceOverrideTier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsPriceOverrideTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The flat fee price of the price tier
type V1SubscriptionProvisionParamsPriceOverrideTierFlatPrice struct {
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

func (r V1SubscriptionProvisionParamsPriceOverrideTierFlatPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsPriceOverrideTierFlatPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsPriceOverrideTierFlatPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionProvisionParamsPriceOverrideTierFlatPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// The unit price of the price tier
type V1SubscriptionProvisionParamsPriceOverrideTierUnitPrice struct {
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

func (r V1SubscriptionProvisionParamsPriceOverrideTierUnitPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsPriceOverrideTierUnitPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsPriceOverrideTierUnitPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionProvisionParamsPriceOverrideTierUnitPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Strategy for scheduling subscription changes
type V1SubscriptionProvisionParamsScheduleStrategy string

const (
	V1SubscriptionProvisionParamsScheduleStrategyEndOfBillingPeriod V1SubscriptionProvisionParamsScheduleStrategy = "END_OF_BILLING_PERIOD"
	V1SubscriptionProvisionParamsScheduleStrategyEndOfBillingMonth  V1SubscriptionProvisionParamsScheduleStrategy = "END_OF_BILLING_MONTH"
	V1SubscriptionProvisionParamsScheduleStrategyImmediate          V1SubscriptionProvisionParamsScheduleStrategy = "IMMEDIATE"
)

// The properties FeatureID, UsageLimit are required.
type V1SubscriptionProvisionParamsSubscriptionEntitlement struct {
	// Feature ID
	FeatureID  string          `json:"featureId,required"`
	UsageLimit float64         `json:"usageLimit,required"`
	IsGranted  param.Opt[bool] `json:"isGranted,omitzero"`
	paramObj
}

func (r V1SubscriptionProvisionParamsSubscriptionEntitlement) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsSubscriptionEntitlement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsSubscriptionEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trial period override settings
//
// The property IsTrial is required.
type V1SubscriptionProvisionParamsTrialOverrideConfiguration struct {
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

func (r V1SubscriptionProvisionParamsTrialOverrideConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1SubscriptionProvisionParamsTrialOverrideConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SubscriptionProvisionParamsTrialOverrideConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1SubscriptionProvisionParamsTrialOverrideConfiguration](
		"trialEndBehavior", "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION",
	)
}

type V1SubscriptionTransferParams struct {
	// Resource ID to transfer the subscription to
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
