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
	"github.com/stiggio/stigg-go/shared/constant"
)

// V1CustomerService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerService] method instead.
type V1CustomerService struct {
	Options []option.RequestOption
	// Operations related to customers
	PaymentMethod V1CustomerPaymentMethodService
	// Operations related to promotional entitlements
	PromotionalEntitlements V1CustomerPromotionalEntitlementService
}

// NewV1CustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomerService(opts ...option.RequestOption) (r V1CustomerService) {
	r = V1CustomerService{}
	r.Options = opts
	r.PaymentMethod = NewV1CustomerPaymentMethodService(opts...)
	r.PromotionalEntitlements = NewV1CustomerPromotionalEntitlementService(opts...)
	return
}

// Retrieves a customer by their unique identifier, including billing information
// and subscription status.
func (r *V1CustomerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing customer's properties such as name, email, and billing
// information.
func (r *V1CustomerService) Update(ctx context.Context, id string, body V1CustomerUpdateParams, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieves a paginated list of customers in the environment.
func (r *V1CustomerService) List(ctx context.Context, query V1CustomerListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1CustomerListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/customers"
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

// Retrieves a paginated list of customers in the environment.
func (r *V1CustomerService) ListAutoPaging(ctx context.Context, query V1CustomerListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1CustomerListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, query, opts...))
}

// Archives a customer, preventing new subscriptions. Optionally cancels existing
// subscriptions.
func (r *V1CustomerService) Archive(ctx context.Context, id string, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Imports multiple customers in bulk. Used for migrating customer data from
// external systems.
func (r *V1CustomerService) Import(ctx context.Context, body V1CustomerImportParams, opts ...option.RequestOption) (res *V1CustomerImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/customers/import"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a paginated list of resources within the same customer.
func (r *V1CustomerService) ListResources(ctx context.Context, id string, query V1CustomerListResourcesParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1CustomerListResourcesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/resources", id)
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

// Retrieves a paginated list of resources within the same customer.
func (r *V1CustomerService) ListResourcesAutoPaging(ctx context.Context, id string, query V1CustomerListResourcesParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1CustomerListResourcesResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListResources(ctx, id, query, opts...))
}

// Creates a new customer and optionally provisions an initial subscription in a
// single operation.
func (r *V1CustomerService) Provision(ctx context.Context, body V1CustomerProvisionParams, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the effective entitlements for a customer or resource, including
// feature and credit entitlements.
func (r *V1CustomerService) GetEntitlements(ctx context.Context, id string, query V1CustomerGetEntitlementsParams, opts ...option.RequestOption) (res *V1CustomerGetEntitlementsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/entitlements", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Restores an archived customer, allowing them to create new subscriptions again.
func (r *V1CustomerService) Unarchive(ctx context.Context, id string, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Response object
type CustomerResponse struct {
	// A customer can be either an organization or an individual
	Data CustomerResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A customer can be either an organization or an individual
type CustomerResponseData struct {
	// Customer slug
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The billing currency of the customer
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
	BillingCurrency string `json:"billingCurrency" api:"nullable"`
	// The unique identifier for the entity in the billing provider
	BillingID string `json:"billingId" api:"nullable"`
	// Customer level coupon
	CouponID string `json:"couponId" api:"nullable"`
	// The default payment method details
	DefaultPaymentMethod CustomerResponseDataDefaultPaymentMethod `json:"defaultPaymentMethod" api:"nullable"`
	// The email of the customer
	Email string `json:"email" api:"nullable" format:"email"`
	// List of integrations
	Integrations []CustomerResponseDataIntegration `json:"integrations"`
	// Language to use for this customer
	Language string `json:"language" api:"nullable"`
	// Additional metadata
	Metadata map[string]string `json:"metadata"`
	// The name of the customer
	Name string `json:"name" api:"nullable"`
	// Vendor-specific billing passthrough fields.
	Passthrough CustomerResponseDataPassthrough `json:"passthrough"`
	// Timezone to use for this customer
	Timezone string `json:"timezone" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		ArchivedAt           respjson.Field
		CreatedAt            respjson.Field
		UpdatedAt            respjson.Field
		BillingCurrency      respjson.Field
		BillingID            respjson.Field
		CouponID             respjson.Field
		DefaultPaymentMethod respjson.Field
		Email                respjson.Field
		Integrations         respjson.Field
		Language             respjson.Field
		Metadata             respjson.Field
		Name                 respjson.Field
		Passthrough          respjson.Field
		Timezone             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponseData) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The default payment method details
type CustomerResponseDataDefaultPaymentMethod struct {
	// The default payment method id
	BillingID string `json:"billingId" api:"required"`
	// The expiration month of the default payment method
	CardExpiryMonth float64 `json:"cardExpiryMonth" api:"required"`
	// The expiration year of the default payment method
	CardExpiryYear float64 `json:"cardExpiryYear" api:"required"`
	// The last 4 digits of the default payment method
	CardLast4Digits string `json:"cardLast4Digits" api:"required"`
	// The default payment method type
	//
	// Any of "CARD", "BANK", "CASH_APP".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID       respjson.Field
		CardExpiryMonth respjson.Field
		CardExpiryYear  respjson.Field
		CardLast4Digits respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponseDataDefaultPaymentMethod) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponseDataDefaultPaymentMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External billing or CRM integration link
type CustomerResponseDataIntegration struct {
	// Integration details
	ID string `json:"id" api:"required"`
	// Synced entity id
	SyncedEntityID string `json:"syncedEntityId" api:"required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		SyncedEntityID   respjson.Field
		VendorIdentifier respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponseDataIntegration) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponseDataIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vendor-specific billing passthrough fields.
type CustomerResponseDataPassthrough struct {
	// Stripe-specific billing fields for the customer.
	Stripe CustomerResponseDataPassthroughStripe `json:"stripe"`
	// Zuora-specific billing fields for the customer.
	Zuora CustomerResponseDataPassthroughZuora `json:"zuora"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Stripe      respjson.Field
		Zuora       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponseDataPassthrough) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponseDataPassthrough) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stripe-specific billing fields for the customer.
type CustomerResponseDataPassthroughStripe struct {
	// Physical address
	BillingAddress CustomerResponseDataPassthroughStripeBillingAddress `json:"billingAddress"`
	// Customer name
	CustomerName string `json:"customerName"`
	// Invoice custom fields
	InvoiceCustomFields map[string]string `json:"invoiceCustomFields"`
	// Additional metadata
	Metadata map[string]string `json:"metadata"`
	// Billing provider payment method id, attached to this customer
	PaymentMethodID string `json:"paymentMethodId"`
	// Physical address
	ShippingAddress CustomerResponseDataPassthroughStripeShippingAddress `json:"shippingAddress"`
	// Tax IDs
	TaxIDs []CustomerResponseDataPassthroughStripeTaxID `json:"taxIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingAddress      respjson.Field
		CustomerName        respjson.Field
		InvoiceCustomFields respjson.Field
		Metadata            respjson.Field
		PaymentMethodID     respjson.Field
		ShippingAddress     respjson.Field
		TaxIDs              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponseDataPassthroughStripe) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponseDataPassthroughStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address
type CustomerResponseDataPassthroughStripeBillingAddress struct {
	// City name
	City string `json:"city"`
	// Country code or name
	Country string `json:"country"`
	// Street address line 1
	Line1 string `json:"line1"`
	// Street address line 2
	Line2 string `json:"line2"`
	// Postal or ZIP code
	PostalCode string `json:"postalCode"`
	// State or province
	State string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		Line2       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponseDataPassthroughStripeBillingAddress) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponseDataPassthroughStripeBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address
type CustomerResponseDataPassthroughStripeShippingAddress struct {
	// City name
	City string `json:"city"`
	// Country code or name
	Country string `json:"country"`
	// Street address line 1
	Line1 string `json:"line1"`
	// Street address line 2
	Line2 string `json:"line2"`
	// Postal or ZIP code
	PostalCode string `json:"postalCode"`
	// State or province
	State string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		Line2       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponseDataPassthroughStripeShippingAddress) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponseDataPassthroughStripeShippingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tax identifier with type and value for customer tax exemptions.
type CustomerResponseDataPassthroughStripeTaxID struct {
	// The type of tax exemption identifier, such as VAT.
	Type string `json:"type" api:"required"`
	// The actual tax identifier value
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponseDataPassthroughStripeTaxID) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponseDataPassthroughStripeTaxID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Zuora-specific billing fields for the customer.
type CustomerResponseDataPassthroughZuora struct {
	// Physical address
	BillingAddress CustomerResponseDataPassthroughZuoraBillingAddress `json:"billingAddress"`
	// Customers selected currency
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
	// Additional metadata
	Metadata map[string]string `json:"metadata"`
	// Billing provider payment method id, attached to this customer
	PaymentMethodID string `json:"paymentMethodId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingAddress  respjson.Field
		Currency        respjson.Field
		Metadata        respjson.Field
		PaymentMethodID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponseDataPassthroughZuora) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponseDataPassthroughZuora) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address
type CustomerResponseDataPassthroughZuoraBillingAddress struct {
	// City name
	City string `json:"city"`
	// Country code or name
	Country string `json:"country"`
	// Street address line 1
	Line1 string `json:"line1"`
	// Street address line 2
	Line2 string `json:"line2"`
	// Postal or ZIP code
	PostalCode string `json:"postalCode"`
	// State or province
	State string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		Line2       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponseDataPassthroughZuoraBillingAddress) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponseDataPassthroughZuoraBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A customer can be either an organization or an individual
type V1CustomerListResponse struct {
	// Customer slug
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The billing currency of the customer
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
	BillingCurrency V1CustomerListResponseBillingCurrency `json:"billingCurrency" api:"nullable"`
	// The unique identifier for the entity in the billing provider
	BillingID string `json:"billingId" api:"nullable"`
	// Customer level coupon
	CouponID V1CustomerListResponseCouponID `json:"couponId" api:"nullable"`
	// The default payment method details
	DefaultPaymentMethod V1CustomerListResponseDefaultPaymentMethod `json:"defaultPaymentMethod" api:"nullable"`
	// The email of the customer
	Email string `json:"email" api:"nullable" format:"email"`
	// List of integrations
	Integrations []V1CustomerListResponseIntegration `json:"integrations"`
	// Language to use for this customer
	Language string `json:"language" api:"nullable"`
	// Additional metadata
	Metadata map[string]string `json:"metadata"`
	// The name of the customer
	Name string `json:"name" api:"nullable"`
	// Vendor-specific billing passthrough fields.
	Passthrough V1CustomerListResponsePassthrough `json:"passthrough"`
	// Timezone to use for this customer
	Timezone string `json:"timezone" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		ArchivedAt           respjson.Field
		CreatedAt            respjson.Field
		UpdatedAt            respjson.Field
		BillingCurrency      respjson.Field
		BillingID            respjson.Field
		CouponID             respjson.Field
		DefaultPaymentMethod respjson.Field
		Email                respjson.Field
		Integrations         respjson.Field
		Language             respjson.Field
		Metadata             respjson.Field
		Name                 respjson.Field
		Passthrough          respjson.Field
		Timezone             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The billing currency of the customer
type V1CustomerListResponseBillingCurrency string

const (
	V1CustomerListResponseBillingCurrencyUsd V1CustomerListResponseBillingCurrency = "usd"
	V1CustomerListResponseBillingCurrencyAed V1CustomerListResponseBillingCurrency = "aed"
	V1CustomerListResponseBillingCurrencyAll V1CustomerListResponseBillingCurrency = "all"
	V1CustomerListResponseBillingCurrencyAmd V1CustomerListResponseBillingCurrency = "amd"
	V1CustomerListResponseBillingCurrencyAng V1CustomerListResponseBillingCurrency = "ang"
	V1CustomerListResponseBillingCurrencyAud V1CustomerListResponseBillingCurrency = "aud"
	V1CustomerListResponseBillingCurrencyAwg V1CustomerListResponseBillingCurrency = "awg"
	V1CustomerListResponseBillingCurrencyAzn V1CustomerListResponseBillingCurrency = "azn"
	V1CustomerListResponseBillingCurrencyBam V1CustomerListResponseBillingCurrency = "bam"
	V1CustomerListResponseBillingCurrencyBbd V1CustomerListResponseBillingCurrency = "bbd"
	V1CustomerListResponseBillingCurrencyBdt V1CustomerListResponseBillingCurrency = "bdt"
	V1CustomerListResponseBillingCurrencyBgn V1CustomerListResponseBillingCurrency = "bgn"
	V1CustomerListResponseBillingCurrencyBif V1CustomerListResponseBillingCurrency = "bif"
	V1CustomerListResponseBillingCurrencyBmd V1CustomerListResponseBillingCurrency = "bmd"
	V1CustomerListResponseBillingCurrencyBnd V1CustomerListResponseBillingCurrency = "bnd"
	V1CustomerListResponseBillingCurrencyBsd V1CustomerListResponseBillingCurrency = "bsd"
	V1CustomerListResponseBillingCurrencyBwp V1CustomerListResponseBillingCurrency = "bwp"
	V1CustomerListResponseBillingCurrencyByn V1CustomerListResponseBillingCurrency = "byn"
	V1CustomerListResponseBillingCurrencyBzd V1CustomerListResponseBillingCurrency = "bzd"
	V1CustomerListResponseBillingCurrencyBrl V1CustomerListResponseBillingCurrency = "brl"
	V1CustomerListResponseBillingCurrencyCad V1CustomerListResponseBillingCurrency = "cad"
	V1CustomerListResponseBillingCurrencyCdf V1CustomerListResponseBillingCurrency = "cdf"
	V1CustomerListResponseBillingCurrencyChf V1CustomerListResponseBillingCurrency = "chf"
	V1CustomerListResponseBillingCurrencyCny V1CustomerListResponseBillingCurrency = "cny"
	V1CustomerListResponseBillingCurrencyCzk V1CustomerListResponseBillingCurrency = "czk"
	V1CustomerListResponseBillingCurrencyDkk V1CustomerListResponseBillingCurrency = "dkk"
	V1CustomerListResponseBillingCurrencyDop V1CustomerListResponseBillingCurrency = "dop"
	V1CustomerListResponseBillingCurrencyDzd V1CustomerListResponseBillingCurrency = "dzd"
	V1CustomerListResponseBillingCurrencyEgp V1CustomerListResponseBillingCurrency = "egp"
	V1CustomerListResponseBillingCurrencyEtb V1CustomerListResponseBillingCurrency = "etb"
	V1CustomerListResponseBillingCurrencyEur V1CustomerListResponseBillingCurrency = "eur"
	V1CustomerListResponseBillingCurrencyFjd V1CustomerListResponseBillingCurrency = "fjd"
	V1CustomerListResponseBillingCurrencyGbp V1CustomerListResponseBillingCurrency = "gbp"
	V1CustomerListResponseBillingCurrencyGel V1CustomerListResponseBillingCurrency = "gel"
	V1CustomerListResponseBillingCurrencyGip V1CustomerListResponseBillingCurrency = "gip"
	V1CustomerListResponseBillingCurrencyGmd V1CustomerListResponseBillingCurrency = "gmd"
	V1CustomerListResponseBillingCurrencyGyd V1CustomerListResponseBillingCurrency = "gyd"
	V1CustomerListResponseBillingCurrencyHkd V1CustomerListResponseBillingCurrency = "hkd"
	V1CustomerListResponseBillingCurrencyHrk V1CustomerListResponseBillingCurrency = "hrk"
	V1CustomerListResponseBillingCurrencyHtg V1CustomerListResponseBillingCurrency = "htg"
	V1CustomerListResponseBillingCurrencyIdr V1CustomerListResponseBillingCurrency = "idr"
	V1CustomerListResponseBillingCurrencyIls V1CustomerListResponseBillingCurrency = "ils"
	V1CustomerListResponseBillingCurrencyInr V1CustomerListResponseBillingCurrency = "inr"
	V1CustomerListResponseBillingCurrencyIsk V1CustomerListResponseBillingCurrency = "isk"
	V1CustomerListResponseBillingCurrencyJmd V1CustomerListResponseBillingCurrency = "jmd"
	V1CustomerListResponseBillingCurrencyJpy V1CustomerListResponseBillingCurrency = "jpy"
	V1CustomerListResponseBillingCurrencyKes V1CustomerListResponseBillingCurrency = "kes"
	V1CustomerListResponseBillingCurrencyKgs V1CustomerListResponseBillingCurrency = "kgs"
	V1CustomerListResponseBillingCurrencyKhr V1CustomerListResponseBillingCurrency = "khr"
	V1CustomerListResponseBillingCurrencyKmf V1CustomerListResponseBillingCurrency = "kmf"
	V1CustomerListResponseBillingCurrencyKrw V1CustomerListResponseBillingCurrency = "krw"
	V1CustomerListResponseBillingCurrencyKyd V1CustomerListResponseBillingCurrency = "kyd"
	V1CustomerListResponseBillingCurrencyKzt V1CustomerListResponseBillingCurrency = "kzt"
	V1CustomerListResponseBillingCurrencyLbp V1CustomerListResponseBillingCurrency = "lbp"
	V1CustomerListResponseBillingCurrencyLkr V1CustomerListResponseBillingCurrency = "lkr"
	V1CustomerListResponseBillingCurrencyLrd V1CustomerListResponseBillingCurrency = "lrd"
	V1CustomerListResponseBillingCurrencyLsl V1CustomerListResponseBillingCurrency = "lsl"
	V1CustomerListResponseBillingCurrencyMad V1CustomerListResponseBillingCurrency = "mad"
	V1CustomerListResponseBillingCurrencyMdl V1CustomerListResponseBillingCurrency = "mdl"
	V1CustomerListResponseBillingCurrencyMga V1CustomerListResponseBillingCurrency = "mga"
	V1CustomerListResponseBillingCurrencyMkd V1CustomerListResponseBillingCurrency = "mkd"
	V1CustomerListResponseBillingCurrencyMmk V1CustomerListResponseBillingCurrency = "mmk"
	V1CustomerListResponseBillingCurrencyMnt V1CustomerListResponseBillingCurrency = "mnt"
	V1CustomerListResponseBillingCurrencyMop V1CustomerListResponseBillingCurrency = "mop"
	V1CustomerListResponseBillingCurrencyMro V1CustomerListResponseBillingCurrency = "mro"
	V1CustomerListResponseBillingCurrencyMvr V1CustomerListResponseBillingCurrency = "mvr"
	V1CustomerListResponseBillingCurrencyMwk V1CustomerListResponseBillingCurrency = "mwk"
	V1CustomerListResponseBillingCurrencyMxn V1CustomerListResponseBillingCurrency = "mxn"
	V1CustomerListResponseBillingCurrencyMyr V1CustomerListResponseBillingCurrency = "myr"
	V1CustomerListResponseBillingCurrencyMzn V1CustomerListResponseBillingCurrency = "mzn"
	V1CustomerListResponseBillingCurrencyNad V1CustomerListResponseBillingCurrency = "nad"
	V1CustomerListResponseBillingCurrencyNgn V1CustomerListResponseBillingCurrency = "ngn"
	V1CustomerListResponseBillingCurrencyNok V1CustomerListResponseBillingCurrency = "nok"
	V1CustomerListResponseBillingCurrencyNpr V1CustomerListResponseBillingCurrency = "npr"
	V1CustomerListResponseBillingCurrencyNzd V1CustomerListResponseBillingCurrency = "nzd"
	V1CustomerListResponseBillingCurrencyPgk V1CustomerListResponseBillingCurrency = "pgk"
	V1CustomerListResponseBillingCurrencyPhp V1CustomerListResponseBillingCurrency = "php"
	V1CustomerListResponseBillingCurrencyPkr V1CustomerListResponseBillingCurrency = "pkr"
	V1CustomerListResponseBillingCurrencyPln V1CustomerListResponseBillingCurrency = "pln"
	V1CustomerListResponseBillingCurrencyQar V1CustomerListResponseBillingCurrency = "qar"
	V1CustomerListResponseBillingCurrencyRon V1CustomerListResponseBillingCurrency = "ron"
	V1CustomerListResponseBillingCurrencyRsd V1CustomerListResponseBillingCurrency = "rsd"
	V1CustomerListResponseBillingCurrencyRub V1CustomerListResponseBillingCurrency = "rub"
	V1CustomerListResponseBillingCurrencyRwf V1CustomerListResponseBillingCurrency = "rwf"
	V1CustomerListResponseBillingCurrencySar V1CustomerListResponseBillingCurrency = "sar"
	V1CustomerListResponseBillingCurrencySbd V1CustomerListResponseBillingCurrency = "sbd"
	V1CustomerListResponseBillingCurrencyScr V1CustomerListResponseBillingCurrency = "scr"
	V1CustomerListResponseBillingCurrencySek V1CustomerListResponseBillingCurrency = "sek"
	V1CustomerListResponseBillingCurrencySgd V1CustomerListResponseBillingCurrency = "sgd"
	V1CustomerListResponseBillingCurrencySle V1CustomerListResponseBillingCurrency = "sle"
	V1CustomerListResponseBillingCurrencySll V1CustomerListResponseBillingCurrency = "sll"
	V1CustomerListResponseBillingCurrencySos V1CustomerListResponseBillingCurrency = "sos"
	V1CustomerListResponseBillingCurrencySzl V1CustomerListResponseBillingCurrency = "szl"
	V1CustomerListResponseBillingCurrencyThb V1CustomerListResponseBillingCurrency = "thb"
	V1CustomerListResponseBillingCurrencyTjs V1CustomerListResponseBillingCurrency = "tjs"
	V1CustomerListResponseBillingCurrencyTop V1CustomerListResponseBillingCurrency = "top"
	V1CustomerListResponseBillingCurrencyTry V1CustomerListResponseBillingCurrency = "try"
	V1CustomerListResponseBillingCurrencyTtd V1CustomerListResponseBillingCurrency = "ttd"
	V1CustomerListResponseBillingCurrencyTzs V1CustomerListResponseBillingCurrency = "tzs"
	V1CustomerListResponseBillingCurrencyUah V1CustomerListResponseBillingCurrency = "uah"
	V1CustomerListResponseBillingCurrencyUzs V1CustomerListResponseBillingCurrency = "uzs"
	V1CustomerListResponseBillingCurrencyVnd V1CustomerListResponseBillingCurrency = "vnd"
	V1CustomerListResponseBillingCurrencyVuv V1CustomerListResponseBillingCurrency = "vuv"
	V1CustomerListResponseBillingCurrencyWst V1CustomerListResponseBillingCurrency = "wst"
	V1CustomerListResponseBillingCurrencyXaf V1CustomerListResponseBillingCurrency = "xaf"
	V1CustomerListResponseBillingCurrencyXcd V1CustomerListResponseBillingCurrency = "xcd"
	V1CustomerListResponseBillingCurrencyYer V1CustomerListResponseBillingCurrency = "yer"
	V1CustomerListResponseBillingCurrencyZar V1CustomerListResponseBillingCurrency = "zar"
	V1CustomerListResponseBillingCurrencyZmw V1CustomerListResponseBillingCurrency = "zmw"
	V1CustomerListResponseBillingCurrencyClp V1CustomerListResponseBillingCurrency = "clp"
	V1CustomerListResponseBillingCurrencyDjf V1CustomerListResponseBillingCurrency = "djf"
	V1CustomerListResponseBillingCurrencyGnf V1CustomerListResponseBillingCurrency = "gnf"
	V1CustomerListResponseBillingCurrencyUgx V1CustomerListResponseBillingCurrency = "ugx"
	V1CustomerListResponseBillingCurrencyPyg V1CustomerListResponseBillingCurrency = "pyg"
	V1CustomerListResponseBillingCurrencyXof V1CustomerListResponseBillingCurrency = "xof"
	V1CustomerListResponseBillingCurrencyXpf V1CustomerListResponseBillingCurrency = "xpf"
)

// Customer level coupon
type V1CustomerListResponseCouponID string

const (
	V1CustomerListResponseCouponIDEmpty V1CustomerListResponseCouponID = ""
)

// The default payment method details
type V1CustomerListResponseDefaultPaymentMethod struct {
	// The default payment method id
	BillingID string `json:"billingId" api:"required"`
	// The expiration month of the default payment method
	CardExpiryMonth float64 `json:"cardExpiryMonth" api:"required"`
	// The expiration year of the default payment method
	CardExpiryYear float64 `json:"cardExpiryYear" api:"required"`
	// The last 4 digits of the default payment method
	CardLast4Digits string `json:"cardLast4Digits" api:"required"`
	// The default payment method type
	//
	// Any of "CARD", "BANK", "CASH_APP".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID       respjson.Field
		CardExpiryMonth respjson.Field
		CardExpiryYear  respjson.Field
		CardLast4Digits respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponseDefaultPaymentMethod) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponseDefaultPaymentMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External billing or CRM integration link
type V1CustomerListResponseIntegration struct {
	// Integration details
	ID string `json:"id" api:"required"`
	// Synced entity id
	SyncedEntityID string `json:"syncedEntityId" api:"required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		SyncedEntityID   respjson.Field
		VendorIdentifier respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponseIntegration) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponseIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vendor-specific billing passthrough fields.
type V1CustomerListResponsePassthrough struct {
	// Stripe-specific billing fields for the customer.
	Stripe V1CustomerListResponsePassthroughStripe `json:"stripe"`
	// Zuora-specific billing fields for the customer.
	Zuora V1CustomerListResponsePassthroughZuora `json:"zuora"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Stripe      respjson.Field
		Zuora       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponsePassthrough) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponsePassthrough) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stripe-specific billing fields for the customer.
type V1CustomerListResponsePassthroughStripe struct {
	// Physical address
	BillingAddress V1CustomerListResponsePassthroughStripeBillingAddress `json:"billingAddress"`
	// Customer name
	CustomerName string `json:"customerName"`
	// Invoice custom fields
	InvoiceCustomFields map[string]string `json:"invoiceCustomFields"`
	// Additional metadata
	Metadata map[string]string `json:"metadata"`
	// Billing provider payment method id, attached to this customer
	PaymentMethodID string `json:"paymentMethodId"`
	// Physical address
	ShippingAddress V1CustomerListResponsePassthroughStripeShippingAddress `json:"shippingAddress"`
	// Tax IDs
	TaxIDs []V1CustomerListResponsePassthroughStripeTaxID `json:"taxIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingAddress      respjson.Field
		CustomerName        respjson.Field
		InvoiceCustomFields respjson.Field
		Metadata            respjson.Field
		PaymentMethodID     respjson.Field
		ShippingAddress     respjson.Field
		TaxIDs              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponsePassthroughStripe) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponsePassthroughStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address
type V1CustomerListResponsePassthroughStripeBillingAddress struct {
	// City name
	City string `json:"city"`
	// Country code or name
	Country string `json:"country"`
	// Street address line 1
	Line1 string `json:"line1"`
	// Street address line 2
	Line2 string `json:"line2"`
	// Postal or ZIP code
	PostalCode string `json:"postalCode"`
	// State or province
	State string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		Line2       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponsePassthroughStripeBillingAddress) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponsePassthroughStripeBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address
type V1CustomerListResponsePassthroughStripeShippingAddress struct {
	// City name
	City string `json:"city"`
	// Country code or name
	Country string `json:"country"`
	// Street address line 1
	Line1 string `json:"line1"`
	// Street address line 2
	Line2 string `json:"line2"`
	// Postal or ZIP code
	PostalCode string `json:"postalCode"`
	// State or province
	State string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		Line2       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponsePassthroughStripeShippingAddress) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponsePassthroughStripeShippingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tax identifier with type and value for customer tax exemptions.
type V1CustomerListResponsePassthroughStripeTaxID struct {
	// The type of tax exemption identifier, such as VAT.
	Type string `json:"type" api:"required"`
	// The actual tax identifier value
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponsePassthroughStripeTaxID) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponsePassthroughStripeTaxID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Zuora-specific billing fields for the customer.
type V1CustomerListResponsePassthroughZuora struct {
	// Physical address
	BillingAddress V1CustomerListResponsePassthroughZuoraBillingAddress `json:"billingAddress"`
	// Customers selected currency
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
	// Additional metadata
	Metadata map[string]string `json:"metadata"`
	// Billing provider payment method id, attached to this customer
	PaymentMethodID string `json:"paymentMethodId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingAddress  respjson.Field
		Currency        respjson.Field
		Metadata        respjson.Field
		PaymentMethodID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponsePassthroughZuora) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponsePassthroughZuora) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address
type V1CustomerListResponsePassthroughZuoraBillingAddress struct {
	// City name
	City string `json:"city"`
	// Country code or name
	Country string `json:"country"`
	// Street address line 1
	Line1 string `json:"line1"`
	// Street address line 2
	Line2 string `json:"line2"`
	// Postal or ZIP code
	PostalCode string `json:"postalCode"`
	// State or province
	State string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		Line2       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponsePassthroughZuoraBillingAddress) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponsePassthroughZuoraBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1CustomerImportResponse struct {
	// List of newly created customer IDs from the import operation.
	Data V1CustomerImportResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerImportResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerImportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of newly created customer IDs from the import operation.
type V1CustomerImportResponseData struct {
	// Customer IDs created during import
	NewCustomers []string `json:"newCustomers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NewCustomers respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerImportResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerImportResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource object that belongs to a customer
type V1CustomerListResourcesResponse struct {
	// Resource slug
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
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
func (r V1CustomerListResourcesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResourcesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1CustomerGetEntitlementsResponse struct {
	// The effective entitlements state for a customer or resource.
	Data V1CustomerGetEntitlementsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerGetEntitlementsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerGetEntitlementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The effective entitlements state for a customer or resource.
type V1CustomerGetEntitlementsResponseData struct {
	// Reason why entitlements access was denied, if applicable
	//
	// Any of "CustomerNotFound", "NoActiveSubscription", "CustomerIsArchived".
	AccessDeniedReason string `json:"accessDeniedReason" api:"required"`
	// List of effective feature and credit entitlements
	Entitlements []V1CustomerGetEntitlementsResponseDataEntitlementUnion `json:"entitlements" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessDeniedReason respjson.Field
		Entitlements       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerGetEntitlementsResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerGetEntitlementsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1CustomerGetEntitlementsResponseDataEntitlementUnion contains all possible
// properties and values from
// [V1CustomerGetEntitlementsResponseDataEntitlementFeature],
// [V1CustomerGetEntitlementsResponseDataEntitlementCredit].
//
// Use the [V1CustomerGetEntitlementsResponseDataEntitlementUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerGetEntitlementsResponseDataEntitlementUnion struct {
	AccessDeniedReason string `json:"accessDeniedReason"`
	IsGranted          bool   `json:"isGranted"`
	// Any of "FEATURE", "CREDIT".
	Type                 string    `json:"type"`
	CurrentUsage         float64   `json:"currentUsage"`
	EntitlementUpdatedAt time.Time `json:"entitlementUpdatedAt"`
	// This field is from variant
	// [V1CustomerGetEntitlementsResponseDataEntitlementFeature].
	Feature V1CustomerGetEntitlementsResponseDataEntitlementFeatureFeature `json:"feature"`
	// This field is from variant
	// [V1CustomerGetEntitlementsResponseDataEntitlementFeature].
	HasUnlimitedUsage bool `json:"hasUnlimitedUsage"`
	// This field is from variant
	// [V1CustomerGetEntitlementsResponseDataEntitlementFeature].
	ResetPeriod string  `json:"resetPeriod"`
	UsageLimit  float64 `json:"usageLimit"`
	// This field is from variant
	// [V1CustomerGetEntitlementsResponseDataEntitlementFeature].
	UsagePeriodAnchor time.Time `json:"usagePeriodAnchor"`
	UsagePeriodEnd    time.Time `json:"usagePeriodEnd"`
	// This field is from variant
	// [V1CustomerGetEntitlementsResponseDataEntitlementFeature].
	UsagePeriodStart time.Time `json:"usagePeriodStart"`
	ValidUntil       time.Time `json:"validUntil"`
	// This field is from variant
	// [V1CustomerGetEntitlementsResponseDataEntitlementCredit].
	Currency V1CustomerGetEntitlementsResponseDataEntitlementCreditCurrency `json:"currency"`
	// This field is from variant
	// [V1CustomerGetEntitlementsResponseDataEntitlementCredit].
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

// anyV1CustomerGetEntitlementsResponseDataEntitlement is implemented by each
// variant of [V1CustomerGetEntitlementsResponseDataEntitlementUnion] to add type
// safety for the return type of
// [V1CustomerGetEntitlementsResponseDataEntitlementUnion.AsAny]
type anyV1CustomerGetEntitlementsResponseDataEntitlement interface {
	implV1CustomerGetEntitlementsResponseDataEntitlementUnion()
}

func (V1CustomerGetEntitlementsResponseDataEntitlementFeature) implV1CustomerGetEntitlementsResponseDataEntitlementUnion() {
}
func (V1CustomerGetEntitlementsResponseDataEntitlementCredit) implV1CustomerGetEntitlementsResponseDataEntitlementUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := V1CustomerGetEntitlementsResponseDataEntitlementUnion.AsAny().(type) {
//	case stigg.V1CustomerGetEntitlementsResponseDataEntitlementFeature:
//	case stigg.V1CustomerGetEntitlementsResponseDataEntitlementCredit:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u V1CustomerGetEntitlementsResponseDataEntitlementUnion) AsAny() anyV1CustomerGetEntitlementsResponseDataEntitlement {
	switch u.Type {
	case "FEATURE":
		return u.AsFeature()
	case "CREDIT":
		return u.AsCredit()
	}
	return nil
}

func (u V1CustomerGetEntitlementsResponseDataEntitlementUnion) AsFeature() (v V1CustomerGetEntitlementsResponseDataEntitlementFeature) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerGetEntitlementsResponseDataEntitlementUnion) AsCredit() (v V1CustomerGetEntitlementsResponseDataEntitlementCredit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerGetEntitlementsResponseDataEntitlementUnion) RawJSON() string { return u.JSON.raw }

func (r *V1CustomerGetEntitlementsResponseDataEntitlementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerGetEntitlementsResponseDataEntitlementFeature struct {
	// Any of "FeatureNotFound", "CustomerNotFound", "CustomerIsArchived",
	// "CustomerResourceNotFound", "NoActiveSubscription",
	// "NoFeatureEntitlementInSubscription", "RequestedUsageExceedingLimit",
	// "RequestedValuesMismatch", "BudgetExceeded", "Unknown", "FeatureTypeMismatch",
	// "Revoked", "InsufficientCredits", "EntitlementNotFound".
	AccessDeniedReason string           `json:"accessDeniedReason" api:"required"`
	IsGranted          bool             `json:"isGranted" api:"required"`
	Type               constant.Feature `json:"type" api:"required"`
	CurrentUsage       float64          `json:"currentUsage"`
	// Timestamp of the last update to the entitlement grant or configuration.
	EntitlementUpdatedAt time.Time                                                      `json:"entitlementUpdatedAt" format:"date-time"`
	Feature              V1CustomerGetEntitlementsResponseDataEntitlementFeatureFeature `json:"feature"`
	HasUnlimitedUsage    bool                                                           `json:"hasUnlimitedUsage"`
	// Any of "YEAR", "MONTH", "WEEK", "DAY", "HOUR".
	ResetPeriod string  `json:"resetPeriod" api:"nullable"`
	UsageLimit  float64 `json:"usageLimit" api:"nullable"`
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
func (r V1CustomerGetEntitlementsResponseDataEntitlementFeature) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerGetEntitlementsResponseDataEntitlementFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerGetEntitlementsResponseDataEntitlementFeatureFeature struct {
	// The unique reference ID of the entitlement.
	ID string `json:"id" api:"required"`
	// The human-readable name of the entitlement, shown in UI elements.
	DisplayName string `json:"displayName" api:"required"`
	// The current status of the feature.
	//
	// Any of "NEW", "SUSPENDED", "ACTIVE".
	FeatureStatus string `json:"featureStatus" api:"required"`
	// The type of feature associated with the entitlement.
	//
	// Any of "BOOLEAN", "NUMBER", "ENUM".
	FeatureType string `json:"featureType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		DisplayName   respjson.Field
		FeatureStatus respjson.Field
		FeatureType   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerGetEntitlementsResponseDataEntitlementFeatureFeature) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerGetEntitlementsResponseDataEntitlementFeatureFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerGetEntitlementsResponseDataEntitlementCredit struct {
	// Any of "FeatureNotFound", "CustomerNotFound", "CustomerIsArchived",
	// "CustomerResourceNotFound", "NoActiveSubscription",
	// "NoFeatureEntitlementInSubscription", "RequestedUsageExceedingLimit",
	// "RequestedValuesMismatch", "BudgetExceeded", "Unknown", "FeatureTypeMismatch",
	// "Revoked", "InsufficientCredits", "EntitlementNotFound".
	AccessDeniedReason string `json:"accessDeniedReason" api:"required"`
	// The currency associated with a credit entitlement.
	Currency     V1CustomerGetEntitlementsResponseDataEntitlementCreditCurrency `json:"currency" api:"required"`
	CurrentUsage float64                                                        `json:"currentUsage" api:"required"`
	IsGranted    bool                                                           `json:"isGranted" api:"required"`
	Type         constant.Credit                                                `json:"type" api:"required"`
	UsageLimit   float64                                                        `json:"usageLimit" api:"required"`
	// Timestamp of the last update to the credit usage.
	UsageUpdatedAt time.Time `json:"usageUpdatedAt" api:"required" format:"date-time"`
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
func (r V1CustomerGetEntitlementsResponseDataEntitlementCredit) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerGetEntitlementsResponseDataEntitlementCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The currency associated with a credit entitlement.
type V1CustomerGetEntitlementsResponseDataEntitlementCreditCurrency struct {
	// The unique identifier of the custom currency.
	CurrencyID string `json:"currencyId" api:"required"`
	// The display name of the currency.
	DisplayName string `json:"displayName" api:"required"`
	// A description of the currency.
	Description string `json:"description" api:"nullable"`
	// Additional metadata associated with the currency.
	Metadata map[string]string `json:"metadata" api:"nullable"`
	// The plural form of the currency unit.
	UnitPlural string `json:"unitPlural" api:"nullable"`
	// The singular form of the currency unit.
	UnitSingular string `json:"unitSingular" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyID   respjson.Field
		DisplayName  respjson.Field
		Description  respjson.Field
		Metadata     respjson.Field
		UnitPlural   respjson.Field
		UnitSingular respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerGetEntitlementsResponseDataEntitlementCreditCurrency) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerGetEntitlementsResponseDataEntitlementCreditCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerUpdateParams struct {
	// The unique identifier for the entity in the billing provider
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// The email of the customer
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Language to use for this customer
	Language param.Opt[string] `json:"language,omitzero"`
	// The name of the customer
	Name param.Opt[string] `json:"name,omitzero"`
	// Timezone to use for this customer
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// The billing currency of the customer
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
	BillingCurrency V1CustomerUpdateParamsBillingCurrency `json:"billingCurrency,omitzero"`
	// Customer level coupon
	CouponID V1CustomerUpdateParamsCouponID `json:"couponId,omitzero"`
	// List of integrations
	Integrations []V1CustomerUpdateParamsIntegration `json:"integrations,omitzero"`
	// Additional metadata
	Metadata map[string]string `json:"metadata,omitzero"`
	// Vendor-specific billing passthrough fields.
	Passthrough V1CustomerUpdateParamsPassthrough `json:"passthrough,omitzero"`
	paramObj
}

func (r V1CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The billing currency of the customer
type V1CustomerUpdateParamsBillingCurrency string

const (
	V1CustomerUpdateParamsBillingCurrencyUsd V1CustomerUpdateParamsBillingCurrency = "usd"
	V1CustomerUpdateParamsBillingCurrencyAed V1CustomerUpdateParamsBillingCurrency = "aed"
	V1CustomerUpdateParamsBillingCurrencyAll V1CustomerUpdateParamsBillingCurrency = "all"
	V1CustomerUpdateParamsBillingCurrencyAmd V1CustomerUpdateParamsBillingCurrency = "amd"
	V1CustomerUpdateParamsBillingCurrencyAng V1CustomerUpdateParamsBillingCurrency = "ang"
	V1CustomerUpdateParamsBillingCurrencyAud V1CustomerUpdateParamsBillingCurrency = "aud"
	V1CustomerUpdateParamsBillingCurrencyAwg V1CustomerUpdateParamsBillingCurrency = "awg"
	V1CustomerUpdateParamsBillingCurrencyAzn V1CustomerUpdateParamsBillingCurrency = "azn"
	V1CustomerUpdateParamsBillingCurrencyBam V1CustomerUpdateParamsBillingCurrency = "bam"
	V1CustomerUpdateParamsBillingCurrencyBbd V1CustomerUpdateParamsBillingCurrency = "bbd"
	V1CustomerUpdateParamsBillingCurrencyBdt V1CustomerUpdateParamsBillingCurrency = "bdt"
	V1CustomerUpdateParamsBillingCurrencyBgn V1CustomerUpdateParamsBillingCurrency = "bgn"
	V1CustomerUpdateParamsBillingCurrencyBif V1CustomerUpdateParamsBillingCurrency = "bif"
	V1CustomerUpdateParamsBillingCurrencyBmd V1CustomerUpdateParamsBillingCurrency = "bmd"
	V1CustomerUpdateParamsBillingCurrencyBnd V1CustomerUpdateParamsBillingCurrency = "bnd"
	V1CustomerUpdateParamsBillingCurrencyBsd V1CustomerUpdateParamsBillingCurrency = "bsd"
	V1CustomerUpdateParamsBillingCurrencyBwp V1CustomerUpdateParamsBillingCurrency = "bwp"
	V1CustomerUpdateParamsBillingCurrencyByn V1CustomerUpdateParamsBillingCurrency = "byn"
	V1CustomerUpdateParamsBillingCurrencyBzd V1CustomerUpdateParamsBillingCurrency = "bzd"
	V1CustomerUpdateParamsBillingCurrencyBrl V1CustomerUpdateParamsBillingCurrency = "brl"
	V1CustomerUpdateParamsBillingCurrencyCad V1CustomerUpdateParamsBillingCurrency = "cad"
	V1CustomerUpdateParamsBillingCurrencyCdf V1CustomerUpdateParamsBillingCurrency = "cdf"
	V1CustomerUpdateParamsBillingCurrencyChf V1CustomerUpdateParamsBillingCurrency = "chf"
	V1CustomerUpdateParamsBillingCurrencyCny V1CustomerUpdateParamsBillingCurrency = "cny"
	V1CustomerUpdateParamsBillingCurrencyCzk V1CustomerUpdateParamsBillingCurrency = "czk"
	V1CustomerUpdateParamsBillingCurrencyDkk V1CustomerUpdateParamsBillingCurrency = "dkk"
	V1CustomerUpdateParamsBillingCurrencyDop V1CustomerUpdateParamsBillingCurrency = "dop"
	V1CustomerUpdateParamsBillingCurrencyDzd V1CustomerUpdateParamsBillingCurrency = "dzd"
	V1CustomerUpdateParamsBillingCurrencyEgp V1CustomerUpdateParamsBillingCurrency = "egp"
	V1CustomerUpdateParamsBillingCurrencyEtb V1CustomerUpdateParamsBillingCurrency = "etb"
	V1CustomerUpdateParamsBillingCurrencyEur V1CustomerUpdateParamsBillingCurrency = "eur"
	V1CustomerUpdateParamsBillingCurrencyFjd V1CustomerUpdateParamsBillingCurrency = "fjd"
	V1CustomerUpdateParamsBillingCurrencyGbp V1CustomerUpdateParamsBillingCurrency = "gbp"
	V1CustomerUpdateParamsBillingCurrencyGel V1CustomerUpdateParamsBillingCurrency = "gel"
	V1CustomerUpdateParamsBillingCurrencyGip V1CustomerUpdateParamsBillingCurrency = "gip"
	V1CustomerUpdateParamsBillingCurrencyGmd V1CustomerUpdateParamsBillingCurrency = "gmd"
	V1CustomerUpdateParamsBillingCurrencyGyd V1CustomerUpdateParamsBillingCurrency = "gyd"
	V1CustomerUpdateParamsBillingCurrencyHkd V1CustomerUpdateParamsBillingCurrency = "hkd"
	V1CustomerUpdateParamsBillingCurrencyHrk V1CustomerUpdateParamsBillingCurrency = "hrk"
	V1CustomerUpdateParamsBillingCurrencyHtg V1CustomerUpdateParamsBillingCurrency = "htg"
	V1CustomerUpdateParamsBillingCurrencyIdr V1CustomerUpdateParamsBillingCurrency = "idr"
	V1CustomerUpdateParamsBillingCurrencyIls V1CustomerUpdateParamsBillingCurrency = "ils"
	V1CustomerUpdateParamsBillingCurrencyInr V1CustomerUpdateParamsBillingCurrency = "inr"
	V1CustomerUpdateParamsBillingCurrencyIsk V1CustomerUpdateParamsBillingCurrency = "isk"
	V1CustomerUpdateParamsBillingCurrencyJmd V1CustomerUpdateParamsBillingCurrency = "jmd"
	V1CustomerUpdateParamsBillingCurrencyJpy V1CustomerUpdateParamsBillingCurrency = "jpy"
	V1CustomerUpdateParamsBillingCurrencyKes V1CustomerUpdateParamsBillingCurrency = "kes"
	V1CustomerUpdateParamsBillingCurrencyKgs V1CustomerUpdateParamsBillingCurrency = "kgs"
	V1CustomerUpdateParamsBillingCurrencyKhr V1CustomerUpdateParamsBillingCurrency = "khr"
	V1CustomerUpdateParamsBillingCurrencyKmf V1CustomerUpdateParamsBillingCurrency = "kmf"
	V1CustomerUpdateParamsBillingCurrencyKrw V1CustomerUpdateParamsBillingCurrency = "krw"
	V1CustomerUpdateParamsBillingCurrencyKyd V1CustomerUpdateParamsBillingCurrency = "kyd"
	V1CustomerUpdateParamsBillingCurrencyKzt V1CustomerUpdateParamsBillingCurrency = "kzt"
	V1CustomerUpdateParamsBillingCurrencyLbp V1CustomerUpdateParamsBillingCurrency = "lbp"
	V1CustomerUpdateParamsBillingCurrencyLkr V1CustomerUpdateParamsBillingCurrency = "lkr"
	V1CustomerUpdateParamsBillingCurrencyLrd V1CustomerUpdateParamsBillingCurrency = "lrd"
	V1CustomerUpdateParamsBillingCurrencyLsl V1CustomerUpdateParamsBillingCurrency = "lsl"
	V1CustomerUpdateParamsBillingCurrencyMad V1CustomerUpdateParamsBillingCurrency = "mad"
	V1CustomerUpdateParamsBillingCurrencyMdl V1CustomerUpdateParamsBillingCurrency = "mdl"
	V1CustomerUpdateParamsBillingCurrencyMga V1CustomerUpdateParamsBillingCurrency = "mga"
	V1CustomerUpdateParamsBillingCurrencyMkd V1CustomerUpdateParamsBillingCurrency = "mkd"
	V1CustomerUpdateParamsBillingCurrencyMmk V1CustomerUpdateParamsBillingCurrency = "mmk"
	V1CustomerUpdateParamsBillingCurrencyMnt V1CustomerUpdateParamsBillingCurrency = "mnt"
	V1CustomerUpdateParamsBillingCurrencyMop V1CustomerUpdateParamsBillingCurrency = "mop"
	V1CustomerUpdateParamsBillingCurrencyMro V1CustomerUpdateParamsBillingCurrency = "mro"
	V1CustomerUpdateParamsBillingCurrencyMvr V1CustomerUpdateParamsBillingCurrency = "mvr"
	V1CustomerUpdateParamsBillingCurrencyMwk V1CustomerUpdateParamsBillingCurrency = "mwk"
	V1CustomerUpdateParamsBillingCurrencyMxn V1CustomerUpdateParamsBillingCurrency = "mxn"
	V1CustomerUpdateParamsBillingCurrencyMyr V1CustomerUpdateParamsBillingCurrency = "myr"
	V1CustomerUpdateParamsBillingCurrencyMzn V1CustomerUpdateParamsBillingCurrency = "mzn"
	V1CustomerUpdateParamsBillingCurrencyNad V1CustomerUpdateParamsBillingCurrency = "nad"
	V1CustomerUpdateParamsBillingCurrencyNgn V1CustomerUpdateParamsBillingCurrency = "ngn"
	V1CustomerUpdateParamsBillingCurrencyNok V1CustomerUpdateParamsBillingCurrency = "nok"
	V1CustomerUpdateParamsBillingCurrencyNpr V1CustomerUpdateParamsBillingCurrency = "npr"
	V1CustomerUpdateParamsBillingCurrencyNzd V1CustomerUpdateParamsBillingCurrency = "nzd"
	V1CustomerUpdateParamsBillingCurrencyPgk V1CustomerUpdateParamsBillingCurrency = "pgk"
	V1CustomerUpdateParamsBillingCurrencyPhp V1CustomerUpdateParamsBillingCurrency = "php"
	V1CustomerUpdateParamsBillingCurrencyPkr V1CustomerUpdateParamsBillingCurrency = "pkr"
	V1CustomerUpdateParamsBillingCurrencyPln V1CustomerUpdateParamsBillingCurrency = "pln"
	V1CustomerUpdateParamsBillingCurrencyQar V1CustomerUpdateParamsBillingCurrency = "qar"
	V1CustomerUpdateParamsBillingCurrencyRon V1CustomerUpdateParamsBillingCurrency = "ron"
	V1CustomerUpdateParamsBillingCurrencyRsd V1CustomerUpdateParamsBillingCurrency = "rsd"
	V1CustomerUpdateParamsBillingCurrencyRub V1CustomerUpdateParamsBillingCurrency = "rub"
	V1CustomerUpdateParamsBillingCurrencyRwf V1CustomerUpdateParamsBillingCurrency = "rwf"
	V1CustomerUpdateParamsBillingCurrencySar V1CustomerUpdateParamsBillingCurrency = "sar"
	V1CustomerUpdateParamsBillingCurrencySbd V1CustomerUpdateParamsBillingCurrency = "sbd"
	V1CustomerUpdateParamsBillingCurrencyScr V1CustomerUpdateParamsBillingCurrency = "scr"
	V1CustomerUpdateParamsBillingCurrencySek V1CustomerUpdateParamsBillingCurrency = "sek"
	V1CustomerUpdateParamsBillingCurrencySgd V1CustomerUpdateParamsBillingCurrency = "sgd"
	V1CustomerUpdateParamsBillingCurrencySle V1CustomerUpdateParamsBillingCurrency = "sle"
	V1CustomerUpdateParamsBillingCurrencySll V1CustomerUpdateParamsBillingCurrency = "sll"
	V1CustomerUpdateParamsBillingCurrencySos V1CustomerUpdateParamsBillingCurrency = "sos"
	V1CustomerUpdateParamsBillingCurrencySzl V1CustomerUpdateParamsBillingCurrency = "szl"
	V1CustomerUpdateParamsBillingCurrencyThb V1CustomerUpdateParamsBillingCurrency = "thb"
	V1CustomerUpdateParamsBillingCurrencyTjs V1CustomerUpdateParamsBillingCurrency = "tjs"
	V1CustomerUpdateParamsBillingCurrencyTop V1CustomerUpdateParamsBillingCurrency = "top"
	V1CustomerUpdateParamsBillingCurrencyTry V1CustomerUpdateParamsBillingCurrency = "try"
	V1CustomerUpdateParamsBillingCurrencyTtd V1CustomerUpdateParamsBillingCurrency = "ttd"
	V1CustomerUpdateParamsBillingCurrencyTzs V1CustomerUpdateParamsBillingCurrency = "tzs"
	V1CustomerUpdateParamsBillingCurrencyUah V1CustomerUpdateParamsBillingCurrency = "uah"
	V1CustomerUpdateParamsBillingCurrencyUzs V1CustomerUpdateParamsBillingCurrency = "uzs"
	V1CustomerUpdateParamsBillingCurrencyVnd V1CustomerUpdateParamsBillingCurrency = "vnd"
	V1CustomerUpdateParamsBillingCurrencyVuv V1CustomerUpdateParamsBillingCurrency = "vuv"
	V1CustomerUpdateParamsBillingCurrencyWst V1CustomerUpdateParamsBillingCurrency = "wst"
	V1CustomerUpdateParamsBillingCurrencyXaf V1CustomerUpdateParamsBillingCurrency = "xaf"
	V1CustomerUpdateParamsBillingCurrencyXcd V1CustomerUpdateParamsBillingCurrency = "xcd"
	V1CustomerUpdateParamsBillingCurrencyYer V1CustomerUpdateParamsBillingCurrency = "yer"
	V1CustomerUpdateParamsBillingCurrencyZar V1CustomerUpdateParamsBillingCurrency = "zar"
	V1CustomerUpdateParamsBillingCurrencyZmw V1CustomerUpdateParamsBillingCurrency = "zmw"
	V1CustomerUpdateParamsBillingCurrencyClp V1CustomerUpdateParamsBillingCurrency = "clp"
	V1CustomerUpdateParamsBillingCurrencyDjf V1CustomerUpdateParamsBillingCurrency = "djf"
	V1CustomerUpdateParamsBillingCurrencyGnf V1CustomerUpdateParamsBillingCurrency = "gnf"
	V1CustomerUpdateParamsBillingCurrencyUgx V1CustomerUpdateParamsBillingCurrency = "ugx"
	V1CustomerUpdateParamsBillingCurrencyPyg V1CustomerUpdateParamsBillingCurrency = "pyg"
	V1CustomerUpdateParamsBillingCurrencyXof V1CustomerUpdateParamsBillingCurrency = "xof"
	V1CustomerUpdateParamsBillingCurrencyXpf V1CustomerUpdateParamsBillingCurrency = "xpf"
)

// Customer level coupon
type V1CustomerUpdateParamsCouponID string

const (
	V1CustomerUpdateParamsCouponIDEmpty V1CustomerUpdateParamsCouponID = ""
)

// External billing or CRM integration link
//
// The properties ID, SyncedEntityID, VendorIdentifier are required.
type V1CustomerUpdateParamsIntegration struct {
	// Synced entity id
	SyncedEntityID param.Opt[string] `json:"syncedEntityId,omitzero" api:"required"`
	// Integration details
	ID string `json:"id" api:"required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier,omitzero" api:"required"`
	paramObj
}

func (r V1CustomerUpdateParamsIntegration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerUpdateParamsIntegration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerUpdateParamsIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerUpdateParamsIntegration](
		"vendorIdentifier", "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE", "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE",
	)
}

// Vendor-specific billing passthrough fields.
type V1CustomerUpdateParamsPassthrough struct {
	// Stripe-specific billing fields for the customer.
	Stripe V1CustomerUpdateParamsPassthroughStripe `json:"stripe,omitzero"`
	// Zuora-specific billing fields for the customer.
	Zuora V1CustomerUpdateParamsPassthroughZuora `json:"zuora,omitzero"`
	paramObj
}

func (r V1CustomerUpdateParamsPassthrough) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerUpdateParamsPassthrough
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerUpdateParamsPassthrough) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stripe-specific billing fields for the customer.
type V1CustomerUpdateParamsPassthroughStripe struct {
	// Customer name
	CustomerName param.Opt[string] `json:"customerName,omitzero"`
	// Billing provider payment method id, attached to this customer
	PaymentMethodID param.Opt[string] `json:"paymentMethodId,omitzero"`
	// Physical address
	BillingAddress V1CustomerUpdateParamsPassthroughStripeBillingAddress `json:"billingAddress,omitzero"`
	// Invoice custom fields
	InvoiceCustomFields map[string]string `json:"invoiceCustomFields,omitzero"`
	// Additional metadata
	Metadata map[string]string `json:"metadata,omitzero"`
	// Physical address
	ShippingAddress V1CustomerUpdateParamsPassthroughStripeShippingAddress `json:"shippingAddress,omitzero"`
	// Tax IDs
	TaxIDs []V1CustomerUpdateParamsPassthroughStripeTaxID `json:"taxIds,omitzero"`
	paramObj
}

func (r V1CustomerUpdateParamsPassthroughStripe) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerUpdateParamsPassthroughStripe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerUpdateParamsPassthroughStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address
type V1CustomerUpdateParamsPassthroughStripeBillingAddress struct {
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

func (r V1CustomerUpdateParamsPassthroughStripeBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerUpdateParamsPassthroughStripeBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerUpdateParamsPassthroughStripeBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address
type V1CustomerUpdateParamsPassthroughStripeShippingAddress struct {
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

func (r V1CustomerUpdateParamsPassthroughStripeShippingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerUpdateParamsPassthroughStripeShippingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerUpdateParamsPassthroughStripeShippingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tax identifier with type and value for customer tax exemptions.
//
// The properties Type, Value are required.
type V1CustomerUpdateParamsPassthroughStripeTaxID struct {
	// The type of tax exemption identifier, such as VAT.
	Type string `json:"type" api:"required"`
	// The actual tax identifier value
	Value string `json:"value" api:"required"`
	paramObj
}

func (r V1CustomerUpdateParamsPassthroughStripeTaxID) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerUpdateParamsPassthroughStripeTaxID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerUpdateParamsPassthroughStripeTaxID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Zuora-specific billing fields for the customer.
type V1CustomerUpdateParamsPassthroughZuora struct {
	// Billing provider payment method id, attached to this customer
	PaymentMethodID param.Opt[string] `json:"paymentMethodId,omitzero"`
	// Physical address
	BillingAddress V1CustomerUpdateParamsPassthroughZuoraBillingAddress `json:"billingAddress,omitzero"`
	// Customers selected currency
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
	// Additional metadata
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1CustomerUpdateParamsPassthroughZuora) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerUpdateParamsPassthroughZuora
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerUpdateParamsPassthroughZuora) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerUpdateParamsPassthroughZuora](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Physical address
type V1CustomerUpdateParamsPassthroughZuoraBillingAddress struct {
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

func (r V1CustomerUpdateParamsPassthroughZuoraBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerUpdateParamsPassthroughZuoraBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerUpdateParamsPassthroughZuoraBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Filter by exact customer email address
	Email param.Opt[string] `query:"email,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by exact customer name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter by creation date using range operators: gt, gte, lt, lte
	CreatedAt V1CustomerListParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerListParams]'s query parameters as `url.Values`.
func (r V1CustomerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by creation date using range operators: gt, gte, lt, lte
type V1CustomerListParamsCreatedAt struct {
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

// URLQuery serializes [V1CustomerListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r V1CustomerListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerImportParams struct {
	// List of customer objects to import
	Customers []V1CustomerImportParamsCustomer `json:"customers,omitzero" api:"required"`
	// Integration details
	IntegrationID param.Opt[string] `json:"integrationId,omitzero"`
	paramObj
}

func (r V1CustomerImportParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerImportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerImportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Email, Name are required.
type V1CustomerImportParamsCustomer struct {
	// The email of the customer
	Email param.Opt[string] `json:"email,omitzero" api:"required" format:"email"`
	// The name of the customer
	Name param.Opt[string] `json:"name,omitzero" api:"required"`
	// Customer slug
	ID string `json:"id" api:"required"`
	// Id in the billing provider
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// Billing provider payment method id
	PaymentMethodID param.Opt[string] `json:"paymentMethodId,omitzero"`
	// The unique identifier for the customer in Salesforce integration
	SalesforceID param.Opt[string] `json:"salesforceId,omitzero"`
	// Timestamp of when the record was last updated
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Additional metadata
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1CustomerImportParamsCustomer) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerImportParamsCustomer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerImportParamsCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerListResourcesParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerListResourcesParams]'s query parameters as
// `url.Values`.
func (r V1CustomerListResourcesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerProvisionParams struct {
	// Customer slug
	ID string `json:"id" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// The email of the customer
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Language to use for this customer
	Language param.Opt[string] `json:"language,omitzero"`
	// The name of the customer
	Name param.Opt[string] `json:"name,omitzero"`
	// Timezone to use for this customer
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// The billing currency of the customer
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
	BillingCurrency V1CustomerProvisionParamsBillingCurrency `json:"billingCurrency,omitzero"`
	// Customer level coupon
	CouponID V1CustomerProvisionParamsCouponID `json:"couponId,omitzero"`
	// The default payment method details
	DefaultPaymentMethod V1CustomerProvisionParamsDefaultPaymentMethod `json:"defaultPaymentMethod,omitzero"`
	// List of integrations
	Integrations []V1CustomerProvisionParamsIntegration `json:"integrations,omitzero"`
	// Additional metadata
	Metadata map[string]string `json:"metadata,omitzero"`
	// Vendor-specific billing passthrough fields.
	Passthrough V1CustomerProvisionParamsPassthrough `json:"passthrough,omitzero"`
	paramObj
}

func (r V1CustomerProvisionParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerProvisionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerProvisionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The billing currency of the customer
type V1CustomerProvisionParamsBillingCurrency string

const (
	V1CustomerProvisionParamsBillingCurrencyUsd V1CustomerProvisionParamsBillingCurrency = "usd"
	V1CustomerProvisionParamsBillingCurrencyAed V1CustomerProvisionParamsBillingCurrency = "aed"
	V1CustomerProvisionParamsBillingCurrencyAll V1CustomerProvisionParamsBillingCurrency = "all"
	V1CustomerProvisionParamsBillingCurrencyAmd V1CustomerProvisionParamsBillingCurrency = "amd"
	V1CustomerProvisionParamsBillingCurrencyAng V1CustomerProvisionParamsBillingCurrency = "ang"
	V1CustomerProvisionParamsBillingCurrencyAud V1CustomerProvisionParamsBillingCurrency = "aud"
	V1CustomerProvisionParamsBillingCurrencyAwg V1CustomerProvisionParamsBillingCurrency = "awg"
	V1CustomerProvisionParamsBillingCurrencyAzn V1CustomerProvisionParamsBillingCurrency = "azn"
	V1CustomerProvisionParamsBillingCurrencyBam V1CustomerProvisionParamsBillingCurrency = "bam"
	V1CustomerProvisionParamsBillingCurrencyBbd V1CustomerProvisionParamsBillingCurrency = "bbd"
	V1CustomerProvisionParamsBillingCurrencyBdt V1CustomerProvisionParamsBillingCurrency = "bdt"
	V1CustomerProvisionParamsBillingCurrencyBgn V1CustomerProvisionParamsBillingCurrency = "bgn"
	V1CustomerProvisionParamsBillingCurrencyBif V1CustomerProvisionParamsBillingCurrency = "bif"
	V1CustomerProvisionParamsBillingCurrencyBmd V1CustomerProvisionParamsBillingCurrency = "bmd"
	V1CustomerProvisionParamsBillingCurrencyBnd V1CustomerProvisionParamsBillingCurrency = "bnd"
	V1CustomerProvisionParamsBillingCurrencyBsd V1CustomerProvisionParamsBillingCurrency = "bsd"
	V1CustomerProvisionParamsBillingCurrencyBwp V1CustomerProvisionParamsBillingCurrency = "bwp"
	V1CustomerProvisionParamsBillingCurrencyByn V1CustomerProvisionParamsBillingCurrency = "byn"
	V1CustomerProvisionParamsBillingCurrencyBzd V1CustomerProvisionParamsBillingCurrency = "bzd"
	V1CustomerProvisionParamsBillingCurrencyBrl V1CustomerProvisionParamsBillingCurrency = "brl"
	V1CustomerProvisionParamsBillingCurrencyCad V1CustomerProvisionParamsBillingCurrency = "cad"
	V1CustomerProvisionParamsBillingCurrencyCdf V1CustomerProvisionParamsBillingCurrency = "cdf"
	V1CustomerProvisionParamsBillingCurrencyChf V1CustomerProvisionParamsBillingCurrency = "chf"
	V1CustomerProvisionParamsBillingCurrencyCny V1CustomerProvisionParamsBillingCurrency = "cny"
	V1CustomerProvisionParamsBillingCurrencyCzk V1CustomerProvisionParamsBillingCurrency = "czk"
	V1CustomerProvisionParamsBillingCurrencyDkk V1CustomerProvisionParamsBillingCurrency = "dkk"
	V1CustomerProvisionParamsBillingCurrencyDop V1CustomerProvisionParamsBillingCurrency = "dop"
	V1CustomerProvisionParamsBillingCurrencyDzd V1CustomerProvisionParamsBillingCurrency = "dzd"
	V1CustomerProvisionParamsBillingCurrencyEgp V1CustomerProvisionParamsBillingCurrency = "egp"
	V1CustomerProvisionParamsBillingCurrencyEtb V1CustomerProvisionParamsBillingCurrency = "etb"
	V1CustomerProvisionParamsBillingCurrencyEur V1CustomerProvisionParamsBillingCurrency = "eur"
	V1CustomerProvisionParamsBillingCurrencyFjd V1CustomerProvisionParamsBillingCurrency = "fjd"
	V1CustomerProvisionParamsBillingCurrencyGbp V1CustomerProvisionParamsBillingCurrency = "gbp"
	V1CustomerProvisionParamsBillingCurrencyGel V1CustomerProvisionParamsBillingCurrency = "gel"
	V1CustomerProvisionParamsBillingCurrencyGip V1CustomerProvisionParamsBillingCurrency = "gip"
	V1CustomerProvisionParamsBillingCurrencyGmd V1CustomerProvisionParamsBillingCurrency = "gmd"
	V1CustomerProvisionParamsBillingCurrencyGyd V1CustomerProvisionParamsBillingCurrency = "gyd"
	V1CustomerProvisionParamsBillingCurrencyHkd V1CustomerProvisionParamsBillingCurrency = "hkd"
	V1CustomerProvisionParamsBillingCurrencyHrk V1CustomerProvisionParamsBillingCurrency = "hrk"
	V1CustomerProvisionParamsBillingCurrencyHtg V1CustomerProvisionParamsBillingCurrency = "htg"
	V1CustomerProvisionParamsBillingCurrencyIdr V1CustomerProvisionParamsBillingCurrency = "idr"
	V1CustomerProvisionParamsBillingCurrencyIls V1CustomerProvisionParamsBillingCurrency = "ils"
	V1CustomerProvisionParamsBillingCurrencyInr V1CustomerProvisionParamsBillingCurrency = "inr"
	V1CustomerProvisionParamsBillingCurrencyIsk V1CustomerProvisionParamsBillingCurrency = "isk"
	V1CustomerProvisionParamsBillingCurrencyJmd V1CustomerProvisionParamsBillingCurrency = "jmd"
	V1CustomerProvisionParamsBillingCurrencyJpy V1CustomerProvisionParamsBillingCurrency = "jpy"
	V1CustomerProvisionParamsBillingCurrencyKes V1CustomerProvisionParamsBillingCurrency = "kes"
	V1CustomerProvisionParamsBillingCurrencyKgs V1CustomerProvisionParamsBillingCurrency = "kgs"
	V1CustomerProvisionParamsBillingCurrencyKhr V1CustomerProvisionParamsBillingCurrency = "khr"
	V1CustomerProvisionParamsBillingCurrencyKmf V1CustomerProvisionParamsBillingCurrency = "kmf"
	V1CustomerProvisionParamsBillingCurrencyKrw V1CustomerProvisionParamsBillingCurrency = "krw"
	V1CustomerProvisionParamsBillingCurrencyKyd V1CustomerProvisionParamsBillingCurrency = "kyd"
	V1CustomerProvisionParamsBillingCurrencyKzt V1CustomerProvisionParamsBillingCurrency = "kzt"
	V1CustomerProvisionParamsBillingCurrencyLbp V1CustomerProvisionParamsBillingCurrency = "lbp"
	V1CustomerProvisionParamsBillingCurrencyLkr V1CustomerProvisionParamsBillingCurrency = "lkr"
	V1CustomerProvisionParamsBillingCurrencyLrd V1CustomerProvisionParamsBillingCurrency = "lrd"
	V1CustomerProvisionParamsBillingCurrencyLsl V1CustomerProvisionParamsBillingCurrency = "lsl"
	V1CustomerProvisionParamsBillingCurrencyMad V1CustomerProvisionParamsBillingCurrency = "mad"
	V1CustomerProvisionParamsBillingCurrencyMdl V1CustomerProvisionParamsBillingCurrency = "mdl"
	V1CustomerProvisionParamsBillingCurrencyMga V1CustomerProvisionParamsBillingCurrency = "mga"
	V1CustomerProvisionParamsBillingCurrencyMkd V1CustomerProvisionParamsBillingCurrency = "mkd"
	V1CustomerProvisionParamsBillingCurrencyMmk V1CustomerProvisionParamsBillingCurrency = "mmk"
	V1CustomerProvisionParamsBillingCurrencyMnt V1CustomerProvisionParamsBillingCurrency = "mnt"
	V1CustomerProvisionParamsBillingCurrencyMop V1CustomerProvisionParamsBillingCurrency = "mop"
	V1CustomerProvisionParamsBillingCurrencyMro V1CustomerProvisionParamsBillingCurrency = "mro"
	V1CustomerProvisionParamsBillingCurrencyMvr V1CustomerProvisionParamsBillingCurrency = "mvr"
	V1CustomerProvisionParamsBillingCurrencyMwk V1CustomerProvisionParamsBillingCurrency = "mwk"
	V1CustomerProvisionParamsBillingCurrencyMxn V1CustomerProvisionParamsBillingCurrency = "mxn"
	V1CustomerProvisionParamsBillingCurrencyMyr V1CustomerProvisionParamsBillingCurrency = "myr"
	V1CustomerProvisionParamsBillingCurrencyMzn V1CustomerProvisionParamsBillingCurrency = "mzn"
	V1CustomerProvisionParamsBillingCurrencyNad V1CustomerProvisionParamsBillingCurrency = "nad"
	V1CustomerProvisionParamsBillingCurrencyNgn V1CustomerProvisionParamsBillingCurrency = "ngn"
	V1CustomerProvisionParamsBillingCurrencyNok V1CustomerProvisionParamsBillingCurrency = "nok"
	V1CustomerProvisionParamsBillingCurrencyNpr V1CustomerProvisionParamsBillingCurrency = "npr"
	V1CustomerProvisionParamsBillingCurrencyNzd V1CustomerProvisionParamsBillingCurrency = "nzd"
	V1CustomerProvisionParamsBillingCurrencyPgk V1CustomerProvisionParamsBillingCurrency = "pgk"
	V1CustomerProvisionParamsBillingCurrencyPhp V1CustomerProvisionParamsBillingCurrency = "php"
	V1CustomerProvisionParamsBillingCurrencyPkr V1CustomerProvisionParamsBillingCurrency = "pkr"
	V1CustomerProvisionParamsBillingCurrencyPln V1CustomerProvisionParamsBillingCurrency = "pln"
	V1CustomerProvisionParamsBillingCurrencyQar V1CustomerProvisionParamsBillingCurrency = "qar"
	V1CustomerProvisionParamsBillingCurrencyRon V1CustomerProvisionParamsBillingCurrency = "ron"
	V1CustomerProvisionParamsBillingCurrencyRsd V1CustomerProvisionParamsBillingCurrency = "rsd"
	V1CustomerProvisionParamsBillingCurrencyRub V1CustomerProvisionParamsBillingCurrency = "rub"
	V1CustomerProvisionParamsBillingCurrencyRwf V1CustomerProvisionParamsBillingCurrency = "rwf"
	V1CustomerProvisionParamsBillingCurrencySar V1CustomerProvisionParamsBillingCurrency = "sar"
	V1CustomerProvisionParamsBillingCurrencySbd V1CustomerProvisionParamsBillingCurrency = "sbd"
	V1CustomerProvisionParamsBillingCurrencyScr V1CustomerProvisionParamsBillingCurrency = "scr"
	V1CustomerProvisionParamsBillingCurrencySek V1CustomerProvisionParamsBillingCurrency = "sek"
	V1CustomerProvisionParamsBillingCurrencySgd V1CustomerProvisionParamsBillingCurrency = "sgd"
	V1CustomerProvisionParamsBillingCurrencySle V1CustomerProvisionParamsBillingCurrency = "sle"
	V1CustomerProvisionParamsBillingCurrencySll V1CustomerProvisionParamsBillingCurrency = "sll"
	V1CustomerProvisionParamsBillingCurrencySos V1CustomerProvisionParamsBillingCurrency = "sos"
	V1CustomerProvisionParamsBillingCurrencySzl V1CustomerProvisionParamsBillingCurrency = "szl"
	V1CustomerProvisionParamsBillingCurrencyThb V1CustomerProvisionParamsBillingCurrency = "thb"
	V1CustomerProvisionParamsBillingCurrencyTjs V1CustomerProvisionParamsBillingCurrency = "tjs"
	V1CustomerProvisionParamsBillingCurrencyTop V1CustomerProvisionParamsBillingCurrency = "top"
	V1CustomerProvisionParamsBillingCurrencyTry V1CustomerProvisionParamsBillingCurrency = "try"
	V1CustomerProvisionParamsBillingCurrencyTtd V1CustomerProvisionParamsBillingCurrency = "ttd"
	V1CustomerProvisionParamsBillingCurrencyTzs V1CustomerProvisionParamsBillingCurrency = "tzs"
	V1CustomerProvisionParamsBillingCurrencyUah V1CustomerProvisionParamsBillingCurrency = "uah"
	V1CustomerProvisionParamsBillingCurrencyUzs V1CustomerProvisionParamsBillingCurrency = "uzs"
	V1CustomerProvisionParamsBillingCurrencyVnd V1CustomerProvisionParamsBillingCurrency = "vnd"
	V1CustomerProvisionParamsBillingCurrencyVuv V1CustomerProvisionParamsBillingCurrency = "vuv"
	V1CustomerProvisionParamsBillingCurrencyWst V1CustomerProvisionParamsBillingCurrency = "wst"
	V1CustomerProvisionParamsBillingCurrencyXaf V1CustomerProvisionParamsBillingCurrency = "xaf"
	V1CustomerProvisionParamsBillingCurrencyXcd V1CustomerProvisionParamsBillingCurrency = "xcd"
	V1CustomerProvisionParamsBillingCurrencyYer V1CustomerProvisionParamsBillingCurrency = "yer"
	V1CustomerProvisionParamsBillingCurrencyZar V1CustomerProvisionParamsBillingCurrency = "zar"
	V1CustomerProvisionParamsBillingCurrencyZmw V1CustomerProvisionParamsBillingCurrency = "zmw"
	V1CustomerProvisionParamsBillingCurrencyClp V1CustomerProvisionParamsBillingCurrency = "clp"
	V1CustomerProvisionParamsBillingCurrencyDjf V1CustomerProvisionParamsBillingCurrency = "djf"
	V1CustomerProvisionParamsBillingCurrencyGnf V1CustomerProvisionParamsBillingCurrency = "gnf"
	V1CustomerProvisionParamsBillingCurrencyUgx V1CustomerProvisionParamsBillingCurrency = "ugx"
	V1CustomerProvisionParamsBillingCurrencyPyg V1CustomerProvisionParamsBillingCurrency = "pyg"
	V1CustomerProvisionParamsBillingCurrencyXof V1CustomerProvisionParamsBillingCurrency = "xof"
	V1CustomerProvisionParamsBillingCurrencyXpf V1CustomerProvisionParamsBillingCurrency = "xpf"
)

// Customer level coupon
type V1CustomerProvisionParamsCouponID string

const (
	V1CustomerProvisionParamsCouponIDEmpty V1CustomerProvisionParamsCouponID = ""
)

// The default payment method details
//
// The properties BillingID, CardExpiryMonth, CardExpiryYear, CardLast4Digits, Type
// are required.
type V1CustomerProvisionParamsDefaultPaymentMethod struct {
	// The default payment method id
	BillingID param.Opt[string] `json:"billingId,omitzero" api:"required"`
	// The expiration month of the default payment method
	CardExpiryMonth param.Opt[float64] `json:"cardExpiryMonth,omitzero" api:"required"`
	// The expiration year of the default payment method
	CardExpiryYear param.Opt[float64] `json:"cardExpiryYear,omitzero" api:"required"`
	// The last 4 digits of the default payment method
	CardLast4Digits param.Opt[string] `json:"cardLast4Digits,omitzero" api:"required"`
	// The default payment method type
	//
	// Any of "CARD", "BANK", "CASH_APP".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r V1CustomerProvisionParamsDefaultPaymentMethod) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerProvisionParamsDefaultPaymentMethod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerProvisionParamsDefaultPaymentMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerProvisionParamsDefaultPaymentMethod](
		"type", "CARD", "BANK", "CASH_APP",
	)
}

// External billing or CRM integration link
//
// The properties ID, SyncedEntityID, VendorIdentifier are required.
type V1CustomerProvisionParamsIntegration struct {
	// Synced entity id
	SyncedEntityID param.Opt[string] `json:"syncedEntityId,omitzero" api:"required"`
	// Integration details
	ID string `json:"id" api:"required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier,omitzero" api:"required"`
	paramObj
}

func (r V1CustomerProvisionParamsIntegration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerProvisionParamsIntegration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerProvisionParamsIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerProvisionParamsIntegration](
		"vendorIdentifier", "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE", "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE",
	)
}

// Vendor-specific billing passthrough fields.
type V1CustomerProvisionParamsPassthrough struct {
	// Stripe-specific billing fields for the customer.
	Stripe V1CustomerProvisionParamsPassthroughStripe `json:"stripe,omitzero"`
	// Zuora-specific billing fields for the customer.
	Zuora V1CustomerProvisionParamsPassthroughZuora `json:"zuora,omitzero"`
	paramObj
}

func (r V1CustomerProvisionParamsPassthrough) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerProvisionParamsPassthrough
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerProvisionParamsPassthrough) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stripe-specific billing fields for the customer.
type V1CustomerProvisionParamsPassthroughStripe struct {
	// Customer name
	CustomerName param.Opt[string] `json:"customerName,omitzero"`
	// Billing provider payment method id, attached to this customer
	PaymentMethodID param.Opt[string] `json:"paymentMethodId,omitzero"`
	// Physical address
	BillingAddress V1CustomerProvisionParamsPassthroughStripeBillingAddress `json:"billingAddress,omitzero"`
	// Invoice custom fields
	InvoiceCustomFields map[string]string `json:"invoiceCustomFields,omitzero"`
	// Additional metadata
	Metadata map[string]string `json:"metadata,omitzero"`
	// Physical address
	ShippingAddress V1CustomerProvisionParamsPassthroughStripeShippingAddress `json:"shippingAddress,omitzero"`
	// Tax IDs
	TaxIDs []V1CustomerProvisionParamsPassthroughStripeTaxID `json:"taxIds,omitzero"`
	paramObj
}

func (r V1CustomerProvisionParamsPassthroughStripe) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerProvisionParamsPassthroughStripe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerProvisionParamsPassthroughStripe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address
type V1CustomerProvisionParamsPassthroughStripeBillingAddress struct {
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

func (r V1CustomerProvisionParamsPassthroughStripeBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerProvisionParamsPassthroughStripeBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerProvisionParamsPassthroughStripeBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Physical address
type V1CustomerProvisionParamsPassthroughStripeShippingAddress struct {
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

func (r V1CustomerProvisionParamsPassthroughStripeShippingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerProvisionParamsPassthroughStripeShippingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerProvisionParamsPassthroughStripeShippingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tax identifier with type and value for customer tax exemptions.
//
// The properties Type, Value are required.
type V1CustomerProvisionParamsPassthroughStripeTaxID struct {
	// The type of tax exemption identifier, such as VAT.
	Type string `json:"type" api:"required"`
	// The actual tax identifier value
	Value string `json:"value" api:"required"`
	paramObj
}

func (r V1CustomerProvisionParamsPassthroughStripeTaxID) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerProvisionParamsPassthroughStripeTaxID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerProvisionParamsPassthroughStripeTaxID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Zuora-specific billing fields for the customer.
type V1CustomerProvisionParamsPassthroughZuora struct {
	// Billing provider payment method id, attached to this customer
	PaymentMethodID param.Opt[string] `json:"paymentMethodId,omitzero"`
	// Physical address
	BillingAddress V1CustomerProvisionParamsPassthroughZuoraBillingAddress `json:"billingAddress,omitzero"`
	// Customers selected currency
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
	// Additional metadata
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1CustomerProvisionParamsPassthroughZuora) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerProvisionParamsPassthroughZuora
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerProvisionParamsPassthroughZuora) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerProvisionParamsPassthroughZuora](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

// Physical address
type V1CustomerProvisionParamsPassthroughZuoraBillingAddress struct {
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

func (r V1CustomerProvisionParamsPassthroughZuoraBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerProvisionParamsPassthroughZuoraBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerProvisionParamsPassthroughZuoraBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerGetEntitlementsParams struct {
	// Resource ID to scope entitlements to a specific resource
	ResourceID param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerGetEntitlementsParams]'s query parameters as
// `url.Values`.
func (r V1CustomerGetEntitlementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
