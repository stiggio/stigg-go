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
	"github.com/stiggio/stigg-go/shared/constant"
)

// V1ContractService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContractService] method instead.
type V1ContractService struct {
	Options []option.RequestOption
}

// NewV1ContractService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ContractService(opts ...option.RequestOption) (r V1ContractService) {
	r = V1ContractService{}
	r.Options = opts
	return
}

// Creates a contract for a customer together with all of its (custom)
// subscriptions in a single atomic operation. Every new subscription is created
// inside one transaction — any validation or creation failure rolls the whole
// contract back. Each subscription entry is either a new subscription to create or
// a reference to an existing custom subscription. Returns the created contract.
func (r *V1ContractService) New(ctx context.Context, params V1ContractNewParams, opts ...option.RequestOption) (res *V1ContractNewResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/contracts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a single contract by its ID, enriched with a preview of its upcoming
// (next) invoice when one is available. Returns 404 when no contract with that ID
// exists in the environment.
func (r *V1ContractService) Get(ctx context.Context, id string, query V1ContractGetParams, opts ...option.RequestOption) (res *V1ContractGetResponse, err error) {
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
	path := fmt.Sprintf("api/v1/contracts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a contract's metadata (name, PO number, activation dates) and optionally
// re-links its subscriptions. Best-effort re-syncs the change to the connected
// billing provider.
func (r *V1ContractService) Update(ctx context.Context, id string, params V1ContractUpdateParams, opts ...option.RequestOption) (res *V1ContractUpdateResponse, err error) {
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
	path := fmt.Sprintf("api/v1/contracts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieves a cursor-paginated list of contracts in the environment, fetched live
// from the connected billing provider. Each contract is enriched with a preview of
// its upcoming (next) invoice when one is available. Returns an empty list when no
// billing provider is connected. Supports filtering by customer external ID,
// state, and name.
func (r *V1ContractService) List(ctx context.Context, params V1ContractListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1ContractListResponse], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/contracts"
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

// Retrieves a cursor-paginated list of contracts in the environment, fetched live
// from the connected billing provider. Each contract is enriched with a preview of
// its upcoming (next) invoice when one is available. Returns an empty list when no
// billing provider is connected. Supports filtering by customer external ID,
// state, and name.
func (r *V1ContractService) ListAutoPaging(ctx context.Context, params V1ContractListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1ContractListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a contract: cancels the contract in the connected billing provider and
// cancels every subscription linked to it.
func (r *V1ContractService) Delete(ctx context.Context, id string, body V1ContractDeleteParams, opts ...option.RequestOption) (res *V1ContractDeleteResponse, err error) {
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
	path := fmt.Sprintf("api/v1/contracts/%s/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Response object
type V1ContractNewResponse struct {
	// A billing contract as reported by the connected billing provider.
	Data V1ContractNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A billing contract as reported by the connected billing provider.
type V1ContractNewResponseData struct {
	// The persisted Stigg contract id (matches a subscription’s contractId; present
	// for Stigg-managed contracts)
	ID string `json:"id" api:"required"`
	// The date the contract activation ends
	ActivationEndDate time.Time `json:"activationEndDate" api:"required" format:"date-time"`
	// The date the contract becomes active
	ActivationStartDate time.Time `json:"activationStartDate" api:"required" format:"date-time"`
	// The billing provider (Received) contract ID; null until the contract has synced
	// to the billing provider
	BillingID string `json:"billingId" api:"required"`
	// The current state of the contract
	//
	// Any of "DRAFT", "ACTIVE", "CANCELED", "END_BILLING".
	BillingState string `json:"billingState" api:"required"`
	// The Stigg contract ref ID (the key used to fetch/update/delete this contract)
	ContractID string `json:"contractId" api:"required"`
	// The date the contract was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The external identifier of the customer the contract belongs to
	CustomerExternalID string `json:"customerExternalId" api:"required"`
	// The external identifier of the contract
	ExternalID string `json:"externalId" api:"required"`
	// The most recent non-draft invoice for this contract (open, paid, or canceled),
	// or null when none exists
	LatestInvoice V1ContractNewResponseDataLatestInvoice `json:"latestInvoice" api:"required"`
	// The contract name (the purchase-order number when set, otherwise the
	// contract/customer name)
	Name string `json:"name" api:"required"`
	// A preview of the contract's upcoming invoice, or null when none is available
	NextInvoice V1ContractNewResponseDataNextInvoice `json:"nextInvoice" api:"required"`
	// Purchase-order number, when set on the contract
	PoNumber string `json:"poNumber" api:"required"`
	// The Stigg contract ref ID (present for Stigg-managed contracts; the key used to
	// update/delete)
	RefID string `json:"refId" api:"required"`
	// The current state of the contract
	//
	// Any of "DRAFT", "ACTIVE", "CANCELED", "END_BILLING".
	State string `json:"state" api:"required"`
	// The custom subscriptions attached to this contract (empty when none)
	Subscriptions []V1ContractNewResponseDataSubscription `json:"subscriptions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ActivationEndDate   respjson.Field
		ActivationStartDate respjson.Field
		BillingID           respjson.Field
		BillingState        respjson.Field
		ContractID          respjson.Field
		CreatedAt           respjson.Field
		CustomerExternalID  respjson.Field
		ExternalID          respjson.Field
		LatestInvoice       respjson.Field
		Name                respjson.Field
		NextInvoice         respjson.Field
		PoNumber            respjson.Field
		RefID               respjson.Field
		State               respjson.Field
		Subscriptions       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ContractNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The most recent non-draft invoice for this contract (open, paid, or canceled),
// or null when none exists
type V1ContractNewResponseDataLatestInvoice struct {
	// Invoice billing ID
	BillingID string `json:"billingId" api:"required"`
	// Invoice creation date
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether payment requires action
	RequiresAction bool `json:"requiresAction" api:"required"`
	// Invoice status
	//
	// Any of "OPEN", "CANCELED", "PAID".
	Status string `json:"status" api:"required"`
	// Amount due
	AmountDue float64 `json:"amountDue" api:"nullable"`
	// Billing reason
	//
	// Any of "BILLING_CYCLE", "SUBSCRIPTION_CREATION", "SUBSCRIPTION_UPDATE",
	// "MANUAL", "MINIMUM_INVOICE_AMOUNT_EXCEEDED", "OTHER".
	BillingReason string `json:"billingReason" api:"nullable"`
	// Invoice currency
	Currency string `json:"currency" api:"nullable"`
	// Invoice PDF URL
	PdfURL string `json:"pdfUrl" api:"nullable"`
	// Total amount
	Total float64 `json:"total" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		CreatedAt      respjson.Field
		RequiresAction respjson.Field
		Status         respjson.Field
		AmountDue      respjson.Field
		BillingReason  respjson.Field
		Currency       respjson.Field
		PdfURL         respjson.Field
		Total          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractNewResponseDataLatestInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1ContractNewResponseDataLatestInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A preview of the contract's upcoming invoice, or null when none is available
type V1ContractNewResponseDataNextInvoice struct {
	// The total amount of the upcoming invoice
	Amount V1ContractNewResponseDataNextInvoiceAmount `json:"amount" api:"required"`
	// The date the upcoming invoice is due
	DueDate time.Time `json:"dueDate" api:"required" format:"date-time"`
	// The billing provider ID of the draft invoice this preview describes
	InvoiceID string `json:"invoiceId" api:"required"`
	// The end of the billing period the upcoming invoice covers
	PeriodEnd time.Time `json:"periodEnd" api:"required" format:"date-time"`
	// The start of the billing period the upcoming invoice covers
	PeriodStart time.Time `json:"periodStart" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		DueDate     respjson.Field
		InvoiceID   respjson.Field
		PeriodEnd   respjson.Field
		PeriodStart respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractNewResponseDataNextInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1ContractNewResponseDataNextInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The total amount of the upcoming invoice
type V1ContractNewResponseDataNextInvoiceAmount struct {
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
func (r V1ContractNewResponseDataNextInvoiceAmount) RawJSON() string { return r.JSON.raw }
func (r *V1ContractNewResponseDataNextInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A custom subscription attached to a contract.
type V1ContractNewResponseDataSubscription struct {
	// Display name of the subscription plan
	PlanDisplayName string `json:"planDisplayName" api:"required"`
	// Display name of the product the subscription plan belongs to
	ProductDisplayName string `json:"productDisplayName" api:"required"`
	// The subscription ref ID (use it to deep-link to the subscription)
	SubscriptionID string `json:"subscriptionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlanDisplayName    respjson.Field
		ProductDisplayName respjson.Field
		SubscriptionID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractNewResponseDataSubscription) RawJSON() string { return r.JSON.raw }
func (r *V1ContractNewResponseDataSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1ContractGetResponse struct {
	// A billing contract as reported by the connected billing provider.
	Data V1ContractGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A billing contract as reported by the connected billing provider.
type V1ContractGetResponseData struct {
	// The persisted Stigg contract id (matches a subscription’s contractId; present
	// for Stigg-managed contracts)
	ID string `json:"id" api:"required"`
	// The date the contract activation ends
	ActivationEndDate time.Time `json:"activationEndDate" api:"required" format:"date-time"`
	// The date the contract becomes active
	ActivationStartDate time.Time `json:"activationStartDate" api:"required" format:"date-time"`
	// The billing provider (Received) contract ID; null until the contract has synced
	// to the billing provider
	BillingID string `json:"billingId" api:"required"`
	// The current state of the contract
	//
	// Any of "DRAFT", "ACTIVE", "CANCELED", "END_BILLING".
	BillingState string `json:"billingState" api:"required"`
	// The Stigg contract ref ID (the key used to fetch/update/delete this contract)
	ContractID string `json:"contractId" api:"required"`
	// The date the contract was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The external identifier of the customer the contract belongs to
	CustomerExternalID string `json:"customerExternalId" api:"required"`
	// The external identifier of the contract
	ExternalID string `json:"externalId" api:"required"`
	// The most recent non-draft invoice for this contract (open, paid, or canceled),
	// or null when none exists
	LatestInvoice V1ContractGetResponseDataLatestInvoice `json:"latestInvoice" api:"required"`
	// The contract name (the purchase-order number when set, otherwise the
	// contract/customer name)
	Name string `json:"name" api:"required"`
	// A preview of the contract's upcoming invoice, or null when none is available
	NextInvoice V1ContractGetResponseDataNextInvoice `json:"nextInvoice" api:"required"`
	// Purchase-order number, when set on the contract
	PoNumber string `json:"poNumber" api:"required"`
	// The Stigg contract ref ID (present for Stigg-managed contracts; the key used to
	// update/delete)
	RefID string `json:"refId" api:"required"`
	// The current state of the contract
	//
	// Any of "DRAFT", "ACTIVE", "CANCELED", "END_BILLING".
	State string `json:"state" api:"required"`
	// The custom subscriptions attached to this contract (empty when none)
	Subscriptions []V1ContractGetResponseDataSubscription `json:"subscriptions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ActivationEndDate   respjson.Field
		ActivationStartDate respjson.Field
		BillingID           respjson.Field
		BillingState        respjson.Field
		ContractID          respjson.Field
		CreatedAt           respjson.Field
		CustomerExternalID  respjson.Field
		ExternalID          respjson.Field
		LatestInvoice       respjson.Field
		Name                respjson.Field
		NextInvoice         respjson.Field
		PoNumber            respjson.Field
		RefID               respjson.Field
		State               respjson.Field
		Subscriptions       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ContractGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The most recent non-draft invoice for this contract (open, paid, or canceled),
// or null when none exists
type V1ContractGetResponseDataLatestInvoice struct {
	// Invoice billing ID
	BillingID string `json:"billingId" api:"required"`
	// Invoice creation date
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether payment requires action
	RequiresAction bool `json:"requiresAction" api:"required"`
	// Invoice status
	//
	// Any of "OPEN", "CANCELED", "PAID".
	Status string `json:"status" api:"required"`
	// Amount due
	AmountDue float64 `json:"amountDue" api:"nullable"`
	// Billing reason
	//
	// Any of "BILLING_CYCLE", "SUBSCRIPTION_CREATION", "SUBSCRIPTION_UPDATE",
	// "MANUAL", "MINIMUM_INVOICE_AMOUNT_EXCEEDED", "OTHER".
	BillingReason string `json:"billingReason" api:"nullable"`
	// Invoice currency
	Currency string `json:"currency" api:"nullable"`
	// Invoice PDF URL
	PdfURL string `json:"pdfUrl" api:"nullable"`
	// Total amount
	Total float64 `json:"total" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		CreatedAt      respjson.Field
		RequiresAction respjson.Field
		Status         respjson.Field
		AmountDue      respjson.Field
		BillingReason  respjson.Field
		Currency       respjson.Field
		PdfURL         respjson.Field
		Total          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetResponseDataLatestInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1ContractGetResponseDataLatestInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A preview of the contract's upcoming invoice, or null when none is available
type V1ContractGetResponseDataNextInvoice struct {
	// The total amount of the upcoming invoice
	Amount V1ContractGetResponseDataNextInvoiceAmount `json:"amount" api:"required"`
	// The date the upcoming invoice is due
	DueDate time.Time `json:"dueDate" api:"required" format:"date-time"`
	// The billing provider ID of the draft invoice this preview describes
	InvoiceID string `json:"invoiceId" api:"required"`
	// The end of the billing period the upcoming invoice covers
	PeriodEnd time.Time `json:"periodEnd" api:"required" format:"date-time"`
	// The start of the billing period the upcoming invoice covers
	PeriodStart time.Time `json:"periodStart" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		DueDate     respjson.Field
		InvoiceID   respjson.Field
		PeriodEnd   respjson.Field
		PeriodStart respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetResponseDataNextInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1ContractGetResponseDataNextInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The total amount of the upcoming invoice
type V1ContractGetResponseDataNextInvoiceAmount struct {
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
func (r V1ContractGetResponseDataNextInvoiceAmount) RawJSON() string { return r.JSON.raw }
func (r *V1ContractGetResponseDataNextInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A custom subscription attached to a contract.
type V1ContractGetResponseDataSubscription struct {
	// Display name of the subscription plan
	PlanDisplayName string `json:"planDisplayName" api:"required"`
	// Display name of the product the subscription plan belongs to
	ProductDisplayName string `json:"productDisplayName" api:"required"`
	// The subscription ref ID (use it to deep-link to the subscription)
	SubscriptionID string `json:"subscriptionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlanDisplayName    respjson.Field
		ProductDisplayName respjson.Field
		SubscriptionID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractGetResponseDataSubscription) RawJSON() string { return r.JSON.raw }
func (r *V1ContractGetResponseDataSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1ContractUpdateResponse struct {
	// A billing contract as reported by the connected billing provider.
	Data V1ContractUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A billing contract as reported by the connected billing provider.
type V1ContractUpdateResponseData struct {
	// The persisted Stigg contract id (matches a subscription’s contractId; present
	// for Stigg-managed contracts)
	ID string `json:"id" api:"required"`
	// The date the contract activation ends
	ActivationEndDate time.Time `json:"activationEndDate" api:"required" format:"date-time"`
	// The date the contract becomes active
	ActivationStartDate time.Time `json:"activationStartDate" api:"required" format:"date-time"`
	// The billing provider (Received) contract ID; null until the contract has synced
	// to the billing provider
	BillingID string `json:"billingId" api:"required"`
	// The current state of the contract
	//
	// Any of "DRAFT", "ACTIVE", "CANCELED", "END_BILLING".
	BillingState string `json:"billingState" api:"required"`
	// The Stigg contract ref ID (the key used to fetch/update/delete this contract)
	ContractID string `json:"contractId" api:"required"`
	// The date the contract was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The external identifier of the customer the contract belongs to
	CustomerExternalID string `json:"customerExternalId" api:"required"`
	// The external identifier of the contract
	ExternalID string `json:"externalId" api:"required"`
	// The most recent non-draft invoice for this contract (open, paid, or canceled),
	// or null when none exists
	LatestInvoice V1ContractUpdateResponseDataLatestInvoice `json:"latestInvoice" api:"required"`
	// The contract name (the purchase-order number when set, otherwise the
	// contract/customer name)
	Name string `json:"name" api:"required"`
	// A preview of the contract's upcoming invoice, or null when none is available
	NextInvoice V1ContractUpdateResponseDataNextInvoice `json:"nextInvoice" api:"required"`
	// Purchase-order number, when set on the contract
	PoNumber string `json:"poNumber" api:"required"`
	// The Stigg contract ref ID (present for Stigg-managed contracts; the key used to
	// update/delete)
	RefID string `json:"refId" api:"required"`
	// The current state of the contract
	//
	// Any of "DRAFT", "ACTIVE", "CANCELED", "END_BILLING".
	State string `json:"state" api:"required"`
	// The custom subscriptions attached to this contract (empty when none)
	Subscriptions []V1ContractUpdateResponseDataSubscription `json:"subscriptions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ActivationEndDate   respjson.Field
		ActivationStartDate respjson.Field
		BillingID           respjson.Field
		BillingState        respjson.Field
		ContractID          respjson.Field
		CreatedAt           respjson.Field
		CustomerExternalID  respjson.Field
		ExternalID          respjson.Field
		LatestInvoice       respjson.Field
		Name                respjson.Field
		NextInvoice         respjson.Field
		PoNumber            respjson.Field
		RefID               respjson.Field
		State               respjson.Field
		Subscriptions       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ContractUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The most recent non-draft invoice for this contract (open, paid, or canceled),
// or null when none exists
type V1ContractUpdateResponseDataLatestInvoice struct {
	// Invoice billing ID
	BillingID string `json:"billingId" api:"required"`
	// Invoice creation date
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether payment requires action
	RequiresAction bool `json:"requiresAction" api:"required"`
	// Invoice status
	//
	// Any of "OPEN", "CANCELED", "PAID".
	Status string `json:"status" api:"required"`
	// Amount due
	AmountDue float64 `json:"amountDue" api:"nullable"`
	// Billing reason
	//
	// Any of "BILLING_CYCLE", "SUBSCRIPTION_CREATION", "SUBSCRIPTION_UPDATE",
	// "MANUAL", "MINIMUM_INVOICE_AMOUNT_EXCEEDED", "OTHER".
	BillingReason string `json:"billingReason" api:"nullable"`
	// Invoice currency
	Currency string `json:"currency" api:"nullable"`
	// Invoice PDF URL
	PdfURL string `json:"pdfUrl" api:"nullable"`
	// Total amount
	Total float64 `json:"total" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		CreatedAt      respjson.Field
		RequiresAction respjson.Field
		Status         respjson.Field
		AmountDue      respjson.Field
		BillingReason  respjson.Field
		Currency       respjson.Field
		PdfURL         respjson.Field
		Total          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractUpdateResponseDataLatestInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1ContractUpdateResponseDataLatestInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A preview of the contract's upcoming invoice, or null when none is available
type V1ContractUpdateResponseDataNextInvoice struct {
	// The total amount of the upcoming invoice
	Amount V1ContractUpdateResponseDataNextInvoiceAmount `json:"amount" api:"required"`
	// The date the upcoming invoice is due
	DueDate time.Time `json:"dueDate" api:"required" format:"date-time"`
	// The billing provider ID of the draft invoice this preview describes
	InvoiceID string `json:"invoiceId" api:"required"`
	// The end of the billing period the upcoming invoice covers
	PeriodEnd time.Time `json:"periodEnd" api:"required" format:"date-time"`
	// The start of the billing period the upcoming invoice covers
	PeriodStart time.Time `json:"periodStart" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		DueDate     respjson.Field
		InvoiceID   respjson.Field
		PeriodEnd   respjson.Field
		PeriodStart respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractUpdateResponseDataNextInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1ContractUpdateResponseDataNextInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The total amount of the upcoming invoice
type V1ContractUpdateResponseDataNextInvoiceAmount struct {
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
func (r V1ContractUpdateResponseDataNextInvoiceAmount) RawJSON() string { return r.JSON.raw }
func (r *V1ContractUpdateResponseDataNextInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A custom subscription attached to a contract.
type V1ContractUpdateResponseDataSubscription struct {
	// Display name of the subscription plan
	PlanDisplayName string `json:"planDisplayName" api:"required"`
	// Display name of the product the subscription plan belongs to
	ProductDisplayName string `json:"productDisplayName" api:"required"`
	// The subscription ref ID (use it to deep-link to the subscription)
	SubscriptionID string `json:"subscriptionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlanDisplayName    respjson.Field
		ProductDisplayName respjson.Field
		SubscriptionID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractUpdateResponseDataSubscription) RawJSON() string { return r.JSON.raw }
func (r *V1ContractUpdateResponseDataSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A billing contract as reported by the connected billing provider.
type V1ContractListResponse struct {
	// The persisted Stigg contract id (matches a subscription’s contractId; present
	// for Stigg-managed contracts)
	ID string `json:"id" api:"required"`
	// The date the contract activation ends
	ActivationEndDate time.Time `json:"activationEndDate" api:"required" format:"date-time"`
	// The date the contract becomes active
	ActivationStartDate time.Time `json:"activationStartDate" api:"required" format:"date-time"`
	// The billing provider (Received) contract ID; null until the contract has synced
	// to the billing provider
	BillingID string `json:"billingId" api:"required"`
	// The current state of the contract
	//
	// Any of "DRAFT", "ACTIVE", "CANCELED", "END_BILLING".
	BillingState V1ContractListResponseBillingState `json:"billingState" api:"required"`
	// The Stigg contract ref ID (the key used to fetch/update/delete this contract)
	ContractID string `json:"contractId" api:"required"`
	// The date the contract was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The external identifier of the customer the contract belongs to
	CustomerExternalID string `json:"customerExternalId" api:"required"`
	// The external identifier of the contract
	ExternalID string `json:"externalId" api:"required"`
	// The most recent non-draft invoice for this contract (open, paid, or canceled),
	// or null when none exists
	LatestInvoice V1ContractListResponseLatestInvoice `json:"latestInvoice" api:"required"`
	// The contract name (the purchase-order number when set, otherwise the
	// contract/customer name)
	Name string `json:"name" api:"required"`
	// A preview of the contract's upcoming invoice, or null when none is available
	NextInvoice V1ContractListResponseNextInvoice `json:"nextInvoice" api:"required"`
	// Purchase-order number, when set on the contract
	PoNumber string `json:"poNumber" api:"required"`
	// The Stigg contract ref ID (present for Stigg-managed contracts; the key used to
	// update/delete)
	RefID string `json:"refId" api:"required"`
	// The current state of the contract
	//
	// Any of "DRAFT", "ACTIVE", "CANCELED", "END_BILLING".
	State V1ContractListResponseState `json:"state" api:"required"`
	// The custom subscriptions attached to this contract (empty when none)
	Subscriptions []V1ContractListResponseSubscription `json:"subscriptions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ActivationEndDate   respjson.Field
		ActivationStartDate respjson.Field
		BillingID           respjson.Field
		BillingState        respjson.Field
		ContractID          respjson.Field
		CreatedAt           respjson.Field
		CustomerExternalID  respjson.Field
		ExternalID          respjson.Field
		LatestInvoice       respjson.Field
		Name                respjson.Field
		NextInvoice         respjson.Field
		PoNumber            respjson.Field
		RefID               respjson.Field
		State               respjson.Field
		Subscriptions       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current state of the contract
type V1ContractListResponseBillingState string

const (
	V1ContractListResponseBillingStateDraft      V1ContractListResponseBillingState = "DRAFT"
	V1ContractListResponseBillingStateActive     V1ContractListResponseBillingState = "ACTIVE"
	V1ContractListResponseBillingStateCanceled   V1ContractListResponseBillingState = "CANCELED"
	V1ContractListResponseBillingStateEndBilling V1ContractListResponseBillingState = "END_BILLING"
)

// The most recent non-draft invoice for this contract (open, paid, or canceled),
// or null when none exists
type V1ContractListResponseLatestInvoice struct {
	// Invoice billing ID
	BillingID string `json:"billingId" api:"required"`
	// Invoice creation date
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether payment requires action
	RequiresAction bool `json:"requiresAction" api:"required"`
	// Invoice status
	//
	// Any of "OPEN", "CANCELED", "PAID".
	Status string `json:"status" api:"required"`
	// Amount due
	AmountDue float64 `json:"amountDue" api:"nullable"`
	// Billing reason
	//
	// Any of "BILLING_CYCLE", "SUBSCRIPTION_CREATION", "SUBSCRIPTION_UPDATE",
	// "MANUAL", "MINIMUM_INVOICE_AMOUNT_EXCEEDED", "OTHER".
	BillingReason string `json:"billingReason" api:"nullable"`
	// Invoice currency
	Currency string `json:"currency" api:"nullable"`
	// Invoice PDF URL
	PdfURL string `json:"pdfUrl" api:"nullable"`
	// Total amount
	Total float64 `json:"total" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		CreatedAt      respjson.Field
		RequiresAction respjson.Field
		Status         respjson.Field
		AmountDue      respjson.Field
		BillingReason  respjson.Field
		Currency       respjson.Field
		PdfURL         respjson.Field
		Total          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractListResponseLatestInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1ContractListResponseLatestInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A preview of the contract's upcoming invoice, or null when none is available
type V1ContractListResponseNextInvoice struct {
	// The total amount of the upcoming invoice
	Amount V1ContractListResponseNextInvoiceAmount `json:"amount" api:"required"`
	// The date the upcoming invoice is due
	DueDate time.Time `json:"dueDate" api:"required" format:"date-time"`
	// The billing provider ID of the draft invoice this preview describes
	InvoiceID string `json:"invoiceId" api:"required"`
	// The end of the billing period the upcoming invoice covers
	PeriodEnd time.Time `json:"periodEnd" api:"required" format:"date-time"`
	// The start of the billing period the upcoming invoice covers
	PeriodStart time.Time `json:"periodStart" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		DueDate     respjson.Field
		InvoiceID   respjson.Field
		PeriodEnd   respjson.Field
		PeriodStart respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractListResponseNextInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1ContractListResponseNextInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The total amount of the upcoming invoice
type V1ContractListResponseNextInvoiceAmount struct {
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
func (r V1ContractListResponseNextInvoiceAmount) RawJSON() string { return r.JSON.raw }
func (r *V1ContractListResponseNextInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current state of the contract
type V1ContractListResponseState string

const (
	V1ContractListResponseStateDraft      V1ContractListResponseState = "DRAFT"
	V1ContractListResponseStateActive     V1ContractListResponseState = "ACTIVE"
	V1ContractListResponseStateCanceled   V1ContractListResponseState = "CANCELED"
	V1ContractListResponseStateEndBilling V1ContractListResponseState = "END_BILLING"
)

// A custom subscription attached to a contract.
type V1ContractListResponseSubscription struct {
	// Display name of the subscription plan
	PlanDisplayName string `json:"planDisplayName" api:"required"`
	// Display name of the product the subscription plan belongs to
	ProductDisplayName string `json:"productDisplayName" api:"required"`
	// The subscription ref ID (use it to deep-link to the subscription)
	SubscriptionID string `json:"subscriptionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlanDisplayName    respjson.Field
		ProductDisplayName respjson.Field
		SubscriptionID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractListResponseSubscription) RawJSON() string { return r.JSON.raw }
func (r *V1ContractListResponseSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1ContractDeleteResponse struct {
	// A billing contract as reported by the connected billing provider.
	Data V1ContractDeleteResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContractDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A billing contract as reported by the connected billing provider.
type V1ContractDeleteResponseData struct {
	// The persisted Stigg contract id (matches a subscription’s contractId; present
	// for Stigg-managed contracts)
	ID string `json:"id" api:"required"`
	// The date the contract activation ends
	ActivationEndDate time.Time `json:"activationEndDate" api:"required" format:"date-time"`
	// The date the contract becomes active
	ActivationStartDate time.Time `json:"activationStartDate" api:"required" format:"date-time"`
	// The billing provider (Received) contract ID; null until the contract has synced
	// to the billing provider
	BillingID string `json:"billingId" api:"required"`
	// The current state of the contract
	//
	// Any of "DRAFT", "ACTIVE", "CANCELED", "END_BILLING".
	BillingState string `json:"billingState" api:"required"`
	// The Stigg contract ref ID (the key used to fetch/update/delete this contract)
	ContractID string `json:"contractId" api:"required"`
	// The date the contract was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The external identifier of the customer the contract belongs to
	CustomerExternalID string `json:"customerExternalId" api:"required"`
	// The external identifier of the contract
	ExternalID string `json:"externalId" api:"required"`
	// The most recent non-draft invoice for this contract (open, paid, or canceled),
	// or null when none exists
	LatestInvoice V1ContractDeleteResponseDataLatestInvoice `json:"latestInvoice" api:"required"`
	// The contract name (the purchase-order number when set, otherwise the
	// contract/customer name)
	Name string `json:"name" api:"required"`
	// A preview of the contract's upcoming invoice, or null when none is available
	NextInvoice V1ContractDeleteResponseDataNextInvoice `json:"nextInvoice" api:"required"`
	// Purchase-order number, when set on the contract
	PoNumber string `json:"poNumber" api:"required"`
	// The Stigg contract ref ID (present for Stigg-managed contracts; the key used to
	// update/delete)
	RefID string `json:"refId" api:"required"`
	// The current state of the contract
	//
	// Any of "DRAFT", "ACTIVE", "CANCELED", "END_BILLING".
	State string `json:"state" api:"required"`
	// The custom subscriptions attached to this contract (empty when none)
	Subscriptions []V1ContractDeleteResponseDataSubscription `json:"subscriptions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ActivationEndDate   respjson.Field
		ActivationStartDate respjson.Field
		BillingID           respjson.Field
		BillingState        respjson.Field
		ContractID          respjson.Field
		CreatedAt           respjson.Field
		CustomerExternalID  respjson.Field
		ExternalID          respjson.Field
		LatestInvoice       respjson.Field
		Name                respjson.Field
		NextInvoice         respjson.Field
		PoNumber            respjson.Field
		RefID               respjson.Field
		State               respjson.Field
		Subscriptions       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ContractDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The most recent non-draft invoice for this contract (open, paid, or canceled),
// or null when none exists
type V1ContractDeleteResponseDataLatestInvoice struct {
	// Invoice billing ID
	BillingID string `json:"billingId" api:"required"`
	// Invoice creation date
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether payment requires action
	RequiresAction bool `json:"requiresAction" api:"required"`
	// Invoice status
	//
	// Any of "OPEN", "CANCELED", "PAID".
	Status string `json:"status" api:"required"`
	// Amount due
	AmountDue float64 `json:"amountDue" api:"nullable"`
	// Billing reason
	//
	// Any of "BILLING_CYCLE", "SUBSCRIPTION_CREATION", "SUBSCRIPTION_UPDATE",
	// "MANUAL", "MINIMUM_INVOICE_AMOUNT_EXCEEDED", "OTHER".
	BillingReason string `json:"billingReason" api:"nullable"`
	// Invoice currency
	Currency string `json:"currency" api:"nullable"`
	// Invoice PDF URL
	PdfURL string `json:"pdfUrl" api:"nullable"`
	// Total amount
	Total float64 `json:"total" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		CreatedAt      respjson.Field
		RequiresAction respjson.Field
		Status         respjson.Field
		AmountDue      respjson.Field
		BillingReason  respjson.Field
		Currency       respjson.Field
		PdfURL         respjson.Field
		Total          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractDeleteResponseDataLatestInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1ContractDeleteResponseDataLatestInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A preview of the contract's upcoming invoice, or null when none is available
type V1ContractDeleteResponseDataNextInvoice struct {
	// The total amount of the upcoming invoice
	Amount V1ContractDeleteResponseDataNextInvoiceAmount `json:"amount" api:"required"`
	// The date the upcoming invoice is due
	DueDate time.Time `json:"dueDate" api:"required" format:"date-time"`
	// The billing provider ID of the draft invoice this preview describes
	InvoiceID string `json:"invoiceId" api:"required"`
	// The end of the billing period the upcoming invoice covers
	PeriodEnd time.Time `json:"periodEnd" api:"required" format:"date-time"`
	// The start of the billing period the upcoming invoice covers
	PeriodStart time.Time `json:"periodStart" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		DueDate     respjson.Field
		InvoiceID   respjson.Field
		PeriodEnd   respjson.Field
		PeriodStart respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractDeleteResponseDataNextInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1ContractDeleteResponseDataNextInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The total amount of the upcoming invoice
type V1ContractDeleteResponseDataNextInvoiceAmount struct {
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
func (r V1ContractDeleteResponseDataNextInvoiceAmount) RawJSON() string { return r.JSON.raw }
func (r *V1ContractDeleteResponseDataNextInvoiceAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A custom subscription attached to a contract.
type V1ContractDeleteResponseDataSubscription struct {
	// Display name of the subscription plan
	PlanDisplayName string `json:"planDisplayName" api:"required"`
	// Display name of the product the subscription plan belongs to
	ProductDisplayName string `json:"productDisplayName" api:"required"`
	// The subscription ref ID (use it to deep-link to the subscription)
	SubscriptionID string `json:"subscriptionId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlanDisplayName    respjson.Field
		ProductDisplayName respjson.Field
		SubscriptionID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContractDeleteResponseDataSubscription) RawJSON() string { return r.JSON.raw }
func (r *V1ContractDeleteResponseDataSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractNewParams struct {
	// The customer ref ID the contract belongs to
	CustomerID string `json:"customerId" api:"required"`
	// The subscriptions to attach to the contract (must be non-empty). Each entry is
	// either a new subscription to create or a reference to an existing custom
	// subscription.
	Subscriptions []V1ContractNewParamsSubscription `json:"subscriptions,omitzero" api:"required"`
	// Optional contract name
	Name param.Opt[string] `json:"name,omitzero"`
	// Optional purchase-order number
	PoNumber param.Opt[string] `json:"poNumber,omitzero"`
	// Optional contract activation end date
	ActivationEndDate param.Opt[time.Time] `json:"activationEndDate,omitzero" format:"date-time"`
	// Optional contract activation start date
	ActivationStartDate param.Opt[time.Time] `json:"activationStartDate,omitzero" format:"date-time"`
	// Whether to set up billing for the contract by creating a billing contract in the
	// connected billing provider. When false, the contract only provisions access
	// (grants entitlements) and no billing contract is created. Defaults to true.
	SetupBilling   param.Opt[bool]   `json:"setupBilling,omitzero"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1ContractNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single subscription on a contract: exactly one of newSubscription or
// existingSubscriptionId must be set.
type V1ContractNewParamsSubscription struct {
	// The subscription ref ID of an already-created custom subscription to link
	ExistingSubscriptionID param.Opt[string] `json:"existingSubscriptionId,omitzero"`
	// A new subscription to create, using the same body the provision-subscription
	// endpoint accepts
	NewSubscription V1ContractNewParamsSubscriptionNewSubscription `json:"newSubscription,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscription) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A new subscription to create, using the same body the provision-subscription
// endpoint accepts
//
// The properties CustomerID, PlanID are required.
type V1ContractNewParamsSubscriptionNewSubscription struct {
	// Customer ID to provision the subscription for
	CustomerID string `json:"customerId" api:"required"`
	// Plan ID to provision
	PlanID string `json:"planId" api:"required"`
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
	// Subscription cancellation date
	CancellationDate param.Opt[time.Time] `json:"cancellationDate,omitzero" format:"date-time"`
	// Subscription start date
	StartDate param.Opt[time.Time] `json:"startDate,omitzero" format:"date-time"`
	// Unit quantity for per-unit pricing. Minimum is 0 (zero is allowed).
	UnitQuantity param.Opt[int64]                                     `json:"unitQuantity,omitzero"`
	Budget       V1ContractNewParamsSubscriptionNewSubscriptionBudget `json:"budget,omitzero"`
	// Minimum spend amount
	MinimumSpend V1ContractNewParamsSubscriptionNewSubscriptionMinimumSpend `json:"minimumSpend,omitzero"`
	Addons       []V1ContractNewParamsSubscriptionNewSubscriptionAddon      `json:"addons,omitzero"`
	// Coupon configuration
	AppliedCoupon V1ContractNewParamsSubscriptionNewSubscriptionAppliedCoupon `json:"appliedCoupon,omitzero"`
	// Billing cycle anchor behavior for the subscription
	//
	// Any of "UNCHANGED", "NOW".
	BillingCycleAnchor string                                                           `json:"billingCycleAnchor,omitzero"`
	BillingInformation V1ContractNewParamsSubscriptionNewSubscriptionBillingInformation `json:"billingInformation,omitzero"`
	// Billing period (MONTHLY or ANNUALLY)
	//
	// Any of "MONTHLY", "ANNUALLY".
	BillingPeriod string                                                 `json:"billingPeriod,omitzero"`
	Charges       []V1ContractNewParamsSubscriptionNewSubscriptionCharge `json:"charges,omitzero"`
	// Checkout page configuration for payment collection
	CheckoutOptions V1ContractNewParamsSubscriptionNewSubscriptionCheckoutOptions    `json:"checkoutOptions,omitzero"`
	Entitlements    []V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion `json:"entitlements,omitzero"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata,omitzero"`
	// How payments should be collected for this subscription
	//
	// Any of "CHARGE", "INVOICE", "NONE".
	PaymentCollectionMethod string                                                        `json:"paymentCollectionMethod,omitzero"`
	PriceOverrides          []V1ContractNewParamsSubscriptionNewSubscriptionPriceOverride `json:"priceOverrides,omitzero"`
	// Strategy for scheduling subscription changes
	//
	// Any of "END_OF_BILLING_PERIOD", "END_OF_BILLING_MONTH", "IMMEDIATE".
	ScheduleStrategy string `json:"scheduleStrategy,omitzero"`
	// Trial period override settings
	TrialOverrideConfiguration V1ContractNewParamsSubscriptionNewSubscriptionTrialOverrideConfiguration `json:"trialOverrideConfiguration,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscription) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscription](
		"billingCycleAnchor", "UNCHANGED", "NOW",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscription](
		"billingPeriod", "MONTHLY", "ANNUALLY",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscription](
		"paymentCollectionMethod", "CHARGE", "INVOICE", "NONE",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscription](
		"scheduleStrategy", "END_OF_BILLING_PERIOD", "END_OF_BILLING_MONTH", "IMMEDIATE",
	)
}

// Addon configuration
//
// The properties ID, Quantity are required.
type V1ContractNewParamsSubscriptionNewSubscriptionAddon struct {
	// Addon ID
	ID string `json:"id" api:"required"`
	// Number of addon instances
	Quantity int64 `json:"quantity" api:"required"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionAddon) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionAddon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionAddon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coupon configuration
type V1ContractNewParamsSubscriptionNewSubscriptionAppliedCoupon struct {
	// Billing provider coupon ID
	BillingCouponID param.Opt[string] `json:"billingCouponId,omitzero"`
	// Stigg coupon ID
	CouponID param.Opt[string] `json:"couponId,omitzero"`
	// Promotion code to apply
	PromotionCode param.Opt[string] `json:"promotionCode,omitzero"`
	// Coupon timing configuration
	Configuration V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponConfiguration `json:"configuration,omitzero"`
	// Ad-hoc discount configuration
	Discount V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscount `json:"discount,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionAppliedCoupon) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionAppliedCoupon
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionAppliedCoupon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coupon timing configuration
type V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponConfiguration struct {
	// Coupon start date
	StartDate param.Opt[time.Time] `json:"startDate,omitzero" format:"date-time"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ad-hoc discount configuration
type V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscount struct {
	// Ad-hoc discount
	Description param.Opt[string] `json:"description,omitzero"`
	// Duration in months
	DurationInMonths param.Opt[float64] `json:"durationInMonths,omitzero"`
	// Discount name
	Name param.Opt[string] `json:"name,omitzero"`
	// Percentage discount
	PercentOff param.Opt[float64] `json:"percentOff,omitzero"`
	// Fixed amounts off by currency
	AmountsOff []V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscountAmountsOff `json:"amountsOff,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscount) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Currency are required.
type V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscountAmountsOff struct {
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
	Currency string `json:"currency,omitzero" api:"required"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscountAmountsOff) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscountAmountsOff
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscountAmountsOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscountAmountsOff](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

type V1ContractNewParamsSubscriptionNewSubscriptionBillingInformation struct {
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
	BillingAddress V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationBillingAddress `json:"billingAddress,omitzero"`
	// Additional metadata for the subscription
	Metadata map[string]string `json:"metadata,omitzero"`
	// How to handle proration for billing changes
	//
	// Any of "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE".
	ProrationBehavior string `json:"prorationBehavior,omitzero"`
	// Customer tax identification numbers
	TaxIDs []V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationTaxID `json:"taxIds,omitzero"`
	// Tax rate identifiers to apply
	TaxRateIDs []string `json:"taxRateIds,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionBillingInformation) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionBillingInformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionBillingInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionBillingInformation](
		"prorationBehavior", "INVOICE_IMMEDIATELY", "CREATE_PRORATIONS", "NONE",
	)
}

// Billing address for the subscription
type V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationBillingAddress struct {
	City       param.Opt[string] `json:"city,omitzero"`
	Country    param.Opt[string] `json:"country,omitzero"`
	Line1      param.Opt[string] `json:"line1,omitzero"`
	Line2      param.Opt[string] `json:"line2,omitzero"`
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	State      param.Opt[string] `json:"state,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tax identifier with type and value for customer tax exemptions.
//
// The properties Type, Value are required.
type V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationTaxID struct {
	// The type of tax exemption identifier, such as VAT.
	Type string `json:"type" api:"required"`
	// The actual tax identifier value
	Value string `json:"value" api:"required"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationTaxID) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationTaxID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationTaxID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties HasSoftLimit, Limit are required.
type V1ContractNewParamsSubscriptionNewSubscriptionBudget struct {
	// Whether the budget is a soft limit
	HasSoftLimit bool `json:"hasSoftLimit" api:"required"`
	// Maximum spending limit
	Limit float64 `json:"limit" api:"required"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionBudget) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionBudget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A charge selection for a subscription (references a catalog charge with a
// quantity).
//
// The properties ID, Quantity, Type are required.
type V1ContractNewParamsSubscriptionNewSubscriptionCharge struct {
	// Charge ID
	ID string `json:"id" api:"required"`
	// Charge quantity. Minimum is 0 (zero is allowed).
	Quantity float64 `json:"quantity" api:"required"`
	// Charge type
	//
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionCharge) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionCharge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionCharge](
		"type", "FEATURE", "CREDIT",
	)
}

// Checkout page configuration for payment collection
//
// The properties CancelURL, SuccessURL are required.
type V1ContractNewParamsSubscriptionNewSubscriptionCheckoutOptions struct {
	// URL to redirect to if checkout is canceled
	CancelURL string `json:"cancelUrl" api:"required" format:"uri"`
	// URL to redirect to after successful checkout
	SuccessURL string `json:"successUrl" api:"required" format:"uri"`
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

func (r V1ContractNewParamsSubscriptionNewSubscriptionCheckoutOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionCheckoutOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionCheckoutOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion struct {
	OfFeature *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeature `json:",omitzero,inline"`
	OfCredit  *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementCredit  `json:",omitzero,inline"`
	paramUnion
}

func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFeature, u.OfCredit)
}
func (u *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) asAny() any {
	if !param.IsOmitted(u.OfFeature) {
		return u.OfFeature
	} else if !param.IsOmitted(u.OfCredit) {
		return u.OfCredit
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) GetHasSoftLimit() *bool {
	if vt := u.OfFeature; vt != nil && vt.HasSoftLimit.Valid() {
		return &vt.HasSoftLimit.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) GetHasUnlimitedUsage() *bool {
	if vt := u.OfFeature; vt != nil && vt.HasUnlimitedUsage.Valid() {
		return &vt.HasUnlimitedUsage.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) GetMonthlyResetPeriodConfiguration() *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureMonthlyResetPeriodConfiguration {
	if vt := u.OfFeature; vt != nil {
		return &vt.MonthlyResetPeriodConfiguration
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) GetResetPeriod() *string {
	if vt := u.OfFeature; vt != nil {
		return &vt.ResetPeriod
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) GetUsageLimit() *int64 {
	if vt := u.OfFeature; vt != nil && vt.UsageLimit.Valid() {
		return &vt.UsageLimit.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) GetWeeklyResetPeriodConfiguration() *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureWeeklyResetPeriodConfiguration {
	if vt := u.OfFeature; vt != nil {
		return &vt.WeeklyResetPeriodConfiguration
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) GetYearlyResetPeriodConfiguration() *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureYearlyResetPeriodConfiguration {
	if vt := u.OfFeature; vt != nil {
		return &vt.YearlyResetPeriodConfiguration
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) GetAmount() *float64 {
	if vt := u.OfCredit; vt != nil {
		return &vt.Amount
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) GetCadence() *string {
	if vt := u.OfCredit; vt != nil {
		return &vt.Cadence
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) GetID() *string {
	if vt := u.OfFeature; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfCredit; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion) GetType() *string {
	if vt := u.OfFeature; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCredit; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion](
		"type",
		apijson.Discriminator[V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeature]("FEATURE"),
		apijson.Discriminator[V1ContractNewParamsSubscriptionNewSubscriptionEntitlementCredit]("CREDIT"),
	)
}

// Feature entitlement configuration for a subscription
//
// The properties ID, Type are required.
type V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeature struct {
	// The feature ID to attach the entitlement to
	ID string `json:"id" api:"required"`
	// Whether the usage limit is a soft limit
	HasSoftLimit param.Opt[bool] `json:"hasSoftLimit,omitzero"`
	// Whether usage is unlimited
	HasUnlimitedUsage param.Opt[bool] `json:"hasUnlimitedUsage,omitzero"`
	// Maximum allowed usage for the feature
	UsageLimit param.Opt[int64] `json:"usageLimit,omitzero"`
	// Configuration for monthly reset period
	MonthlyResetPeriodConfiguration V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureMonthlyResetPeriodConfiguration `json:"monthlyResetPeriodConfiguration,omitzero"`
	// Configuration for weekly reset period
	WeeklyResetPeriodConfiguration V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureWeeklyResetPeriodConfiguration `json:"weeklyResetPeriodConfiguration,omitzero"`
	// Configuration for yearly reset period
	YearlyResetPeriodConfiguration V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureYearlyResetPeriodConfiguration `json:"yearlyResetPeriodConfiguration,omitzero"`
	// Period at which usage resets
	//
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string `json:"resetPeriod,omitzero"`
	// SubscriptionFeatureEntitlementRequest
	//
	// This field can be elided, and will marshal its zero value as "FEATURE".
	Type constant.Feature `json:"type" default:"FEATURE"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeature) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeature
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeature](
		"resetPeriod", "YEAR", "MONTH", "WEEK", "DAY", "HOUR",
	)
}

// Configuration for monthly reset period
//
// The property AccordingTo is required.
type V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureMonthlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or StartOfTheMonth)
	//
	// Any of "SubscriptionStart", "StartOfTheMonth".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureMonthlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureMonthlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureMonthlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureMonthlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "StartOfTheMonth",
	)
}

// Configuration for weekly reset period
//
// The property AccordingTo is required.
type V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureWeeklyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart or specific day)
	//
	// Any of "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday",
	// "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureWeeklyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureWeeklyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureWeeklyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureWeeklyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart", "EverySunday", "EveryMonday", "EveryTuesday", "EveryWednesday", "EveryThursday", "EveryFriday", "EverySaturday",
	)
}

// Configuration for yearly reset period
//
// The property AccordingTo is required.
type V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureYearlyResetPeriodConfiguration struct {
	// Reset anchor (SubscriptionStart)
	//
	// Any of "SubscriptionStart".
	AccordingTo string `json:"accordingTo,omitzero" api:"required"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureYearlyResetPeriodConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureYearlyResetPeriodConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureYearlyResetPeriodConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureYearlyResetPeriodConfiguration](
		"accordingTo", "SubscriptionStart",
	)
}

// Credit entitlement configuration for a subscription
//
// The properties ID, Amount, Cadence, Type are required.
type V1ContractNewParamsSubscriptionNewSubscriptionEntitlementCredit struct {
	// The custom currency ID for the credit entitlement
	ID string `json:"id" api:"required"`
	// Credit grant amount
	Amount float64 `json:"amount" api:"required"`
	// Credit grant cadence (MONTH or YEAR)
	//
	// Any of "MONTH", "YEAR".
	Cadence string `json:"cadence,omitzero" api:"required"`
	// SubscriptionCreditEntitlementRequest
	//
	// This field can be elided, and will marshal its zero value as "CREDIT".
	Type constant.Credit `json:"type" default:"CREDIT"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionEntitlementCredit) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionEntitlementCredit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionEntitlementCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionEntitlementCredit](
		"cadence", "MONTH", "YEAR",
	)
}

// Minimum spend amount
type V1ContractNewParamsSubscriptionNewSubscriptionMinimumSpend struct {
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

func (r V1ContractNewParamsSubscriptionNewSubscriptionMinimumSpend) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionMinimumSpend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionMinimumSpend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionMinimumSpend](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

type V1ContractNewParamsSubscriptionNewSubscriptionPriceOverride struct {
	// Addon identifier for the price override
	AddonID param.Opt[string] `json:"addonId,omitzero"`
	// Feature identifier for the price override
	FeatureID param.Opt[string] `json:"featureId,omitzero"`
	// The price amount
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Whether this is a base charge override
	BaseCharge param.Opt[bool] `json:"baseCharge,omitzero"`
	// The billing country code of the price
	BillingCountryCode param.Opt[string] `json:"billingCountryCode,omitzero"`
	// Block size for pricing
	BlockSize param.Opt[float64] `json:"blockSize,omitzero"`
	// Any of "BEGINNING_OF_BILLING_PERIOD", "MONTHLY".
	CreditGrantCadence string                                                                `json:"creditGrantCadence,omitzero"`
	CreditRate         V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideCreditRate `json:"creditRate,omitzero"`
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
	// Pricing tiers configuration
	Tiers []V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTier `json:"tiers,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionPriceOverride) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionPriceOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionPriceOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionPriceOverride](
		"creditGrantCadence", "BEGINNING_OF_BILLING_PERIOD", "MONTHLY",
	)
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionPriceOverride](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// The properties Amount, CurrencyID are required.
type V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideCreditRate struct {
	// The credit rate amount
	Amount float64 `json:"amount" api:"required"`
	// The custom currency refId for the credit rate
	CurrencyID string `json:"currencyId" api:"required"`
	// A custom formula for calculating cost based on single event dimensions
	CostFormula param.Opt[string] `json:"costFormula,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideCreditRate) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideCreditRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideCreditRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTier struct {
	// The up to quantity of the price tier
	UpTo param.Opt[float64] `json:"upTo,omitzero"`
	// The flat fee price of the price tier
	FlatPrice V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierFlatPrice `json:"flatPrice,omitzero"`
	// The unit price of the price tier
	UnitPrice V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierUnitPrice `json:"unitPrice,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTier) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The flat fee price of the price tier
//
// The properties Amount, Currency are required.
type V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierFlatPrice struct {
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
	Currency string `json:"currency,omitzero" api:"required"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierFlatPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierFlatPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierFlatPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierFlatPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// The unit price of the price tier
//
// The properties Amount, Currency are required.
type V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierUnitPrice struct {
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
	Currency string `json:"currency,omitzero" api:"required"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierUnitPrice) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierUnitPrice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierUnitPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierUnitPrice](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Trial period override settings
//
// The property IsTrial is required.
type V1ContractNewParamsSubscriptionNewSubscriptionTrialOverrideConfiguration struct {
	// Whether the subscription should start with a trial period
	IsTrial bool `json:"isTrial" api:"required"`
	// Custom trial end date
	TrialEndDate param.Opt[time.Time] `json:"trialEndDate,omitzero" format:"date-time"`
	// Behavior when trial ends: CONVERT_TO_PAID or CANCEL_SUBSCRIPTION
	//
	// Any of "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION".
	TrialEndBehavior string `json:"trialEndBehavior,omitzero"`
	paramObj
}

func (r V1ContractNewParamsSubscriptionNewSubscriptionTrialOverrideConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractNewParamsSubscriptionNewSubscriptionTrialOverrideConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractNewParamsSubscriptionNewSubscriptionTrialOverrideConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ContractNewParamsSubscriptionNewSubscriptionTrialOverrideConfiguration](
		"trialEndBehavior", "CONVERT_TO_PAID", "CANCEL_SUBSCRIPTION",
	)
}

type V1ContractGetParams struct {
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

type V1ContractUpdateParams struct {
	// New contract name
	Name param.Opt[string] `json:"name,omitzero"`
	// New purchase-order number
	PoNumber param.Opt[string] `json:"poNumber,omitzero"`
	// New activation end date
	ActivationEndDate param.Opt[time.Time] `json:"activationEndDate,omitzero" format:"date-time"`
	// New activation start date
	ActivationStartDate param.Opt[time.Time] `json:"activationStartDate,omitzero" format:"date-time"`
	// Enable billing on a provision-access-only contract by creating a billing
	// contract in the connected billing provider. Only takes effect when true and the
	// contract has no billing yet; omitting it leaves billing unchanged. Billing is
	// never removed by an update.
	SetupBilling   param.Opt[bool]   `json:"setupBilling,omitzero"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	// When provided, replaces the set of subscriptions linked to the contract
	// (subscription ref IDs)
	SubscriptionIDs []string `json:"subscriptionIds,omitzero"`
	paramObj
}

func (r V1ContractUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContractUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContractUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContractListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter by the exact external ID of the customer the contract belongs to
	CustomerExternalID param.Opt[string] `query:"customerExternalId,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by exact contract name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter by contract state. Supports comma-separated values for multiple states
	State          param.Opt[string] `query:"state,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1ContractListParams]'s query parameters as `url.Values`.
func (r V1ContractListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ContractDeleteParams struct {
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}
