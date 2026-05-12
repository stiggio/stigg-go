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

// Operations related to credit grants
//
// V1CreditGrantService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CreditGrantService] method instead.
type V1CreditGrantService struct {
	Options []option.RequestOption
}

// NewV1CreditGrantService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CreditGrantService(opts ...option.RequestOption) (r V1CreditGrantService) {
	r = V1CreditGrantService{}
	r.Options = opts
	return
}

// Creates a new credit grant for a customer with specified amount, type, and
// optional billing configuration.
func (r *V1CreditGrantService) New(ctx context.Context, body V1CreditGrantNewParams, opts ...option.RequestOption) (res *CreditGrantResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/credits/grants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a paginated list of credit grants for a customer.
func (r *V1CreditGrantService) List(ctx context.Context, query V1CreditGrantListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1CreditGrantListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/credits/grants"
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

// Retrieves a paginated list of credit grants for a customer.
func (r *V1CreditGrantService) ListAutoPaging(ctx context.Context, query V1CreditGrantListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1CreditGrantListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, query, opts...))
}

// Voids an existing credit grant, preventing further consumption of the remaining
// credits.
func (r *V1CreditGrantService) Void(ctx context.Context, id string, opts ...option.RequestOption) (res *CreditGrantResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/credits/grants/%s/void", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Response object
type CreditGrantResponse struct {
	// Credit grant object representing allocated credits for a customer
	Data CreditGrantResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditGrantResponse) RawJSON() string { return r.JSON.raw }
func (r *CreditGrantResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit grant object representing allocated credits for a customer
type CreditGrantResponseData struct {
	// The unique readable identifier of the credit grant
	ID string `json:"id" api:"required"`
	// The total credits granted
	Amount float64 `json:"amount" api:"required"`
	// An optional comment on the credit grant
	Comment string `json:"comment" api:"required"`
	// The total credits consumed from this grant
	ConsumedAmount float64 `json:"consumedAmount" api:"required"`
	// The monetary cost of the credit grant
	Cost CreditGrantResponseDataCost `json:"cost" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The currency identifier for this grant
	CurrencyID string `json:"currencyId" api:"required"`
	// The customer ID this grant belongs to
	CustomerID string `json:"customerId" api:"required"`
	// The display name of the credit grant
	DisplayName string `json:"displayName" api:"required"`
	// The date when the credit grant becomes effective
	EffectiveAt time.Time `json:"effectiveAt" api:"required" format:"date-time"`
	// The date when the credit grant expires
	ExpireAt time.Time `json:"expireAt" api:"required" format:"date-time"`
	// The type of credit grant (PAID, PROMOTIONAL, RECURRING)
	//
	// Any of "PAID", "PROMOTIONAL", "RECURRING", "OVERDRAFT".
	GrantType string `json:"grantType" api:"required"`
	// The billing invoice ID associated with this grant
	InvoiceID string `json:"invoiceId" api:"required"`
	// The latest invoice details for this grant
	LatestInvoice CreditGrantResponseDataLatestInvoice `json:"latestInvoice" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The payment collection status
	//
	// Any of "NOT_REQUIRED", "PROCESSING", "FAILED", "ACTION_REQUIRED".
	PaymentCollection string `json:"paymentCollection" api:"required"`
	// The priority of the credit grant (lower number = higher priority)
	Priority float64 `json:"priority" api:"required"`
	// The resource ID this grant is scoped to
	ResourceID string `json:"resourceId" api:"required"`
	// The source type of the grant (PRICE, PLAN_ENTITLEMENT, ADDON_ENTITLEMENT)
	//
	// Any of "PRICE", "PLAN_ENTITLEMENT", "ADDON_ENTITLEMENT".
	SourceType string `json:"sourceType" api:"required"`
	// The effective status of the credit grant
	//
	// Any of "PAYMENT_PENDING", "ACTIVE", "EXPIRED", "VOIDED", "SCHEDULED".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The date when the credit grant was voided
	VoidedAt time.Time `json:"voidedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Amount            respjson.Field
		Comment           respjson.Field
		ConsumedAmount    respjson.Field
		Cost              respjson.Field
		CreatedAt         respjson.Field
		CurrencyID        respjson.Field
		CustomerID        respjson.Field
		DisplayName       respjson.Field
		EffectiveAt       respjson.Field
		ExpireAt          respjson.Field
		GrantType         respjson.Field
		InvoiceID         respjson.Field
		LatestInvoice     respjson.Field
		Metadata          respjson.Field
		PaymentCollection respjson.Field
		Priority          respjson.Field
		ResourceID        respjson.Field
		SourceType        respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		VoidedAt          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditGrantResponseData) RawJSON() string { return r.JSON.raw }
func (r *CreditGrantResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The monetary cost of the credit grant
type CreditGrantResponseDataCost struct {
	// The cost amount
	Amount float64 `json:"amount" api:"required"`
	// The currency code
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
func (r CreditGrantResponseDataCost) RawJSON() string { return r.JSON.raw }
func (r *CreditGrantResponseDataCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The latest invoice details for this grant
type CreditGrantResponseDataLatestInvoice struct {
	// The billing provider invoice ID
	BillingID string `json:"billingId" api:"required"`
	// The billing reason for the invoice
	//
	// Any of "MANUAL", "OTHER".
	BillingReason string `json:"billingReason" api:"required"`
	// The invoice creation date
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The invoice currency
	Currency string `json:"currency" api:"required"`
	// The invoice due date
	DueDate time.Time `json:"dueDate" api:"required" format:"date-time"`
	// Error message if payment failed
	ErrorMessage string `json:"errorMessage" api:"required"`
	// The payment URL for settling the invoice
	PaymentURL string `json:"paymentUrl" api:"required"`
	// The PDF URL of the invoice
	PdfURL string `json:"pdfUrl" api:"required"`
	// Whether the invoice requires user action
	RequiresAction bool `json:"requiresAction" api:"required"`
	// The invoice status
	//
	// Any of "OPEN", "PAID", "CANCELED".
	Status string `json:"status" api:"required"`
	// The subtotal amount before tax
	SubTotal float64 `json:"subTotal" api:"required"`
	// The tax amount
	Tax float64 `json:"tax" api:"required"`
	// The total amount including tax
	Total float64 `json:"total" api:"required"`
	// The invoice last update date
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		BillingReason  respjson.Field
		CreatedAt      respjson.Field
		Currency       respjson.Field
		DueDate        respjson.Field
		ErrorMessage   respjson.Field
		PaymentURL     respjson.Field
		PdfURL         respjson.Field
		RequiresAction respjson.Field
		Status         respjson.Field
		SubTotal       respjson.Field
		Tax            respjson.Field
		Total          respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditGrantResponseDataLatestInvoice) RawJSON() string { return r.JSON.raw }
func (r *CreditGrantResponseDataLatestInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit grant object representing allocated credits for a customer
type V1CreditGrantListResponse struct {
	// The unique readable identifier of the credit grant
	ID string `json:"id" api:"required"`
	// The total credits granted
	Amount float64 `json:"amount" api:"required"`
	// An optional comment on the credit grant
	Comment string `json:"comment" api:"required"`
	// The total credits consumed from this grant
	ConsumedAmount float64 `json:"consumedAmount" api:"required"`
	// The monetary cost of the credit grant
	Cost V1CreditGrantListResponseCost `json:"cost" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The currency identifier for this grant
	CurrencyID string `json:"currencyId" api:"required"`
	// The customer ID this grant belongs to
	CustomerID string `json:"customerId" api:"required"`
	// The display name of the credit grant
	DisplayName string `json:"displayName" api:"required"`
	// The date when the credit grant becomes effective
	EffectiveAt time.Time `json:"effectiveAt" api:"required" format:"date-time"`
	// The date when the credit grant expires
	ExpireAt time.Time `json:"expireAt" api:"required" format:"date-time"`
	// The type of credit grant (PAID, PROMOTIONAL, RECURRING)
	//
	// Any of "PAID", "PROMOTIONAL", "RECURRING", "OVERDRAFT".
	GrantType V1CreditGrantListResponseGrantType `json:"grantType" api:"required"`
	// The billing invoice ID associated with this grant
	InvoiceID string `json:"invoiceId" api:"required"`
	// The latest invoice details for this grant
	LatestInvoice V1CreditGrantListResponseLatestInvoice `json:"latestInvoice" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The payment collection status
	//
	// Any of "NOT_REQUIRED", "PROCESSING", "FAILED", "ACTION_REQUIRED".
	PaymentCollection V1CreditGrantListResponsePaymentCollection `json:"paymentCollection" api:"required"`
	// The priority of the credit grant (lower number = higher priority)
	Priority float64 `json:"priority" api:"required"`
	// The resource ID this grant is scoped to
	ResourceID string `json:"resourceId" api:"required"`
	// The source type of the grant (PRICE, PLAN_ENTITLEMENT, ADDON_ENTITLEMENT)
	//
	// Any of "PRICE", "PLAN_ENTITLEMENT", "ADDON_ENTITLEMENT".
	SourceType V1CreditGrantListResponseSourceType `json:"sourceType" api:"required"`
	// The effective status of the credit grant
	//
	// Any of "PAYMENT_PENDING", "ACTIVE", "EXPIRED", "VOIDED", "SCHEDULED".
	Status V1CreditGrantListResponseStatus `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The date when the credit grant was voided
	VoidedAt time.Time `json:"voidedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Amount            respjson.Field
		Comment           respjson.Field
		ConsumedAmount    respjson.Field
		Cost              respjson.Field
		CreatedAt         respjson.Field
		CurrencyID        respjson.Field
		CustomerID        respjson.Field
		DisplayName       respjson.Field
		EffectiveAt       respjson.Field
		ExpireAt          respjson.Field
		GrantType         respjson.Field
		InvoiceID         respjson.Field
		LatestInvoice     respjson.Field
		Metadata          respjson.Field
		PaymentCollection respjson.Field
		Priority          respjson.Field
		ResourceID        respjson.Field
		SourceType        respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		VoidedAt          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The monetary cost of the credit grant
type V1CreditGrantListResponseCost struct {
	// The cost amount
	Amount float64 `json:"amount" api:"required"`
	// The currency code
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
func (r V1CreditGrantListResponseCost) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListResponseCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of credit grant (PAID, PROMOTIONAL, RECURRING)
type V1CreditGrantListResponseGrantType string

const (
	V1CreditGrantListResponseGrantTypePaid        V1CreditGrantListResponseGrantType = "PAID"
	V1CreditGrantListResponseGrantTypePromotional V1CreditGrantListResponseGrantType = "PROMOTIONAL"
	V1CreditGrantListResponseGrantTypeRecurring   V1CreditGrantListResponseGrantType = "RECURRING"
	V1CreditGrantListResponseGrantTypeOverdraft   V1CreditGrantListResponseGrantType = "OVERDRAFT"
)

// The latest invoice details for this grant
type V1CreditGrantListResponseLatestInvoice struct {
	// The billing provider invoice ID
	BillingID string `json:"billingId" api:"required"`
	// The billing reason for the invoice
	//
	// Any of "MANUAL", "OTHER".
	BillingReason string `json:"billingReason" api:"required"`
	// The invoice creation date
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The invoice currency
	Currency string `json:"currency" api:"required"`
	// The invoice due date
	DueDate time.Time `json:"dueDate" api:"required" format:"date-time"`
	// Error message if payment failed
	ErrorMessage string `json:"errorMessage" api:"required"`
	// The payment URL for settling the invoice
	PaymentURL string `json:"paymentUrl" api:"required"`
	// The PDF URL of the invoice
	PdfURL string `json:"pdfUrl" api:"required"`
	// Whether the invoice requires user action
	RequiresAction bool `json:"requiresAction" api:"required"`
	// The invoice status
	//
	// Any of "OPEN", "PAID", "CANCELED".
	Status string `json:"status" api:"required"`
	// The subtotal amount before tax
	SubTotal float64 `json:"subTotal" api:"required"`
	// The tax amount
	Tax float64 `json:"tax" api:"required"`
	// The total amount including tax
	Total float64 `json:"total" api:"required"`
	// The invoice last update date
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		BillingReason  respjson.Field
		CreatedAt      respjson.Field
		Currency       respjson.Field
		DueDate        respjson.Field
		ErrorMessage   respjson.Field
		PaymentURL     respjson.Field
		PdfURL         respjson.Field
		RequiresAction respjson.Field
		Status         respjson.Field
		SubTotal       respjson.Field
		Tax            respjson.Field
		Total          respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditGrantListResponseLatestInvoice) RawJSON() string { return r.JSON.raw }
func (r *V1CreditGrantListResponseLatestInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payment collection status
type V1CreditGrantListResponsePaymentCollection string

const (
	V1CreditGrantListResponsePaymentCollectionNotRequired    V1CreditGrantListResponsePaymentCollection = "NOT_REQUIRED"
	V1CreditGrantListResponsePaymentCollectionProcessing     V1CreditGrantListResponsePaymentCollection = "PROCESSING"
	V1CreditGrantListResponsePaymentCollectionFailed         V1CreditGrantListResponsePaymentCollection = "FAILED"
	V1CreditGrantListResponsePaymentCollectionActionRequired V1CreditGrantListResponsePaymentCollection = "ACTION_REQUIRED"
)

// The source type of the grant (PRICE, PLAN_ENTITLEMENT, ADDON_ENTITLEMENT)
type V1CreditGrantListResponseSourceType string

const (
	V1CreditGrantListResponseSourceTypePrice            V1CreditGrantListResponseSourceType = "PRICE"
	V1CreditGrantListResponseSourceTypePlanEntitlement  V1CreditGrantListResponseSourceType = "PLAN_ENTITLEMENT"
	V1CreditGrantListResponseSourceTypeAddonEntitlement V1CreditGrantListResponseSourceType = "ADDON_ENTITLEMENT"
)

// The effective status of the credit grant
type V1CreditGrantListResponseStatus string

const (
	V1CreditGrantListResponseStatusPaymentPending V1CreditGrantListResponseStatus = "PAYMENT_PENDING"
	V1CreditGrantListResponseStatusActive         V1CreditGrantListResponseStatus = "ACTIVE"
	V1CreditGrantListResponseStatusExpired        V1CreditGrantListResponseStatus = "EXPIRED"
	V1CreditGrantListResponseStatusVoided         V1CreditGrantListResponseStatus = "VOIDED"
	V1CreditGrantListResponseStatusScheduled      V1CreditGrantListResponseStatus = "SCHEDULED"
)

type V1CreditGrantNewParams struct {
	// The credit amount to grant
	Amount float64 `json:"amount" api:"required"`
	// The credit currency ID (required)
	CurrencyID string `json:"currencyId" api:"required"`
	// The customer ID to grant credits to (required)
	CustomerID string `json:"customerId" api:"required"`
	// The display name for the credit grant
	DisplayName string `json:"displayName" api:"required"`
	// The type of credit grant (PAID, PROMOTIONAL)
	//
	// Any of "PAID", "PROMOTIONAL".
	GrantType V1CreditGrantNewParamsGrantType `json:"grantType,omitzero" api:"required"`
	// Whether to wait for payment confirmation before returning (default: true)
	AwaitPaymentConfirmation param.Opt[bool] `json:"awaitPaymentConfirmation,omitzero"`
	// An optional comment on the credit grant
	Comment param.Opt[string] `json:"comment,omitzero"`
	// The date when the credit grant becomes effective
	EffectiveAt param.Opt[time.Time] `json:"effectiveAt,omitzero" format:"date-time"`
	// The date when the credit grant expires
	ExpireAt param.Opt[time.Time] `json:"expireAt,omitzero" format:"date-time"`
	// The priority of the credit grant (lower number = higher priority)
	Priority param.Opt[int64] `json:"priority,omitzero"`
	// The resource ID to scope the grant to
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Billing information for the credit grant
	BillingInformation V1CreditGrantNewParamsBillingInformation `json:"billingInformation,omitzero"`
	// The monetary cost of the credit grant
	Cost V1CreditGrantNewParamsCost `json:"cost,omitzero"`
	// Additional metadata for the credit grant
	Metadata map[string]string `json:"metadata,omitzero"`
	// The payment collection method (CHARGE, INVOICE, NONE)
	//
	// Any of "CHARGE", "INVOICE", "NONE".
	PaymentCollectionMethod V1CreditGrantNewParamsPaymentCollectionMethod `json:"paymentCollectionMethod,omitzero"`
	paramObj
}

func (r V1CreditGrantNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of credit grant (PAID, PROMOTIONAL)
type V1CreditGrantNewParamsGrantType string

const (
	V1CreditGrantNewParamsGrantTypePaid        V1CreditGrantNewParamsGrantType = "PAID"
	V1CreditGrantNewParamsGrantTypePromotional V1CreditGrantNewParamsGrantType = "PROMOTIONAL"
)

// Billing information for the credit grant
type V1CreditGrantNewParamsBillingInformation struct {
	// Days until the invoice is due
	InvoiceDaysUntilDue param.Opt[float64] `json:"invoiceDaysUntilDue,omitzero"`
	// Whether the invoice is already paid
	IsInvoicePaid param.Opt[bool] `json:"isInvoicePaid,omitzero"`
	// The billing address
	BillingAddress V1CreditGrantNewParamsBillingInformationBillingAddress `json:"billingAddress,omitzero"`
	paramObj
}

func (r V1CreditGrantNewParamsBillingInformation) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantNewParamsBillingInformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantNewParamsBillingInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The billing address
type V1CreditGrantNewParamsBillingInformationBillingAddress struct {
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

func (r V1CreditGrantNewParamsBillingInformationBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantNewParamsBillingInformationBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantNewParamsBillingInformationBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The monetary cost of the credit grant
//
// The properties Amount, Currency are required.
type V1CreditGrantNewParamsCost struct {
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

func (r V1CreditGrantNewParamsCost) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditGrantNewParamsCost
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditGrantNewParamsCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CreditGrantNewParamsCost](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// The payment collection method (CHARGE, INVOICE, NONE)
type V1CreditGrantNewParamsPaymentCollectionMethod string

const (
	V1CreditGrantNewParamsPaymentCollectionMethodCharge  V1CreditGrantNewParamsPaymentCollectionMethod = "CHARGE"
	V1CreditGrantNewParamsPaymentCollectionMethodInvoice V1CreditGrantNewParamsPaymentCollectionMethod = "INVOICE"
	V1CreditGrantNewParamsPaymentCollectionMethodNone    V1CreditGrantNewParamsPaymentCollectionMethod = "NONE"
)

type V1CreditGrantListParams struct {
	// Filter by customer ID (required)
	CustomerID string `query:"customerId" api:"required" json:"-"`
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Filter by currency ID
	CurrencyID param.Opt[string] `query:"currencyId,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by resource ID. When omitted, only grants without a resource are returned
	ResourceID param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	// Filter by creation date using range operators: gt, gte, lt, lte
	CreatedAt V1CreditGrantListParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CreditGrantListParams]'s query parameters as
// `url.Values`.
func (r V1CreditGrantListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by creation date using range operators: gt, gte, lt, lte
type V1CreditGrantListParamsCreatedAt struct {
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

// URLQuery serializes [V1CreditGrantListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r V1CreditGrantListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
