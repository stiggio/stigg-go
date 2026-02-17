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

// V1CustomerService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerService] method instead.
type V1CustomerService struct {
	Options       []option.RequestOption
	PaymentMethod V1CustomerPaymentMethodService
}

// NewV1CustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomerService(opts ...option.RequestOption) (r V1CustomerService) {
	r = V1CustomerService{}
	r.Options = opts
	r.PaymentMethod = NewV1CustomerPaymentMethodService(opts...)
	return
}

// Retrieves a customer by their unique identifier, including billing information
// and subscription status.
func (r *V1CustomerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing customer's properties such as name, email, and billing
// information.
func (r *V1CustomerService) Update(ctx context.Context, id string, body V1CustomerUpdateParams, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
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
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Imports multiple customers in bulk. Used for migrating customer data from
// external systems.
func (r *V1CustomerService) Import(ctx context.Context, body V1CustomerImportParams, opts ...option.RequestOption) (res *V1CustomerImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/customers/import"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a list of customerresources
func (r *V1CustomerService) ListResources(ctx context.Context, id string, query V1CustomerListResourcesParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1CustomerListResourcesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
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

// Get a list of customerresources
func (r *V1CustomerService) ListResourcesAutoPaging(ctx context.Context, id string, query V1CustomerListResourcesParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1CustomerListResourcesResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListResources(ctx, id, query, opts...))
}

// Creates a new customer and optionally provisions an initial subscription in a
// single operation.
func (r *V1CustomerService) Provision(ctx context.Context, body V1CustomerProvisionParams, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Restores an archived customer, allowing them to create new subscriptions again.
func (r *V1CustomerService) Unarchive(ctx context.Context, id string, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Response object
type CustomerResponse struct {
	// A customer can be either an organization or an individual
	Data CustomerResponseData `json:"data,required"`
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
	ID string `json:"id,required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt,required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Customer level coupon
	CouponID string `json:"couponId,nullable"`
	// The default payment method details
	DefaultPaymentMethod CustomerResponseDataDefaultPaymentMethod `json:"defaultPaymentMethod,nullable"`
	// The email of the customer
	Email string `json:"email,nullable" format:"email"`
	// List of integrations
	Integrations []CustomerResponseDataIntegration `json:"integrations"`
	// Additional metadata
	Metadata map[string]string `json:"metadata"`
	// The name of the customer
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		ArchivedAt           respjson.Field
		CreatedAt            respjson.Field
		UpdatedAt            respjson.Field
		CouponID             respjson.Field
		DefaultPaymentMethod respjson.Field
		Email                respjson.Field
		Integrations         respjson.Field
		Metadata             respjson.Field
		Name                 respjson.Field
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
	BillingID string `json:"billingId,required"`
	// The expiration month of the default payment method
	CardExpiryMonth float64 `json:"cardExpiryMonth,required"`
	// The expiration year of the default payment method
	CardExpiryYear float64 `json:"cardExpiryYear,required"`
	// The last 4 digits of the default payment method
	CardLast4Digits string `json:"cardLast4Digits,required"`
	// The default payment method type
	//
	// Any of "CARD", "BANK", "CASH_APP".
	Type string `json:"type,required"`
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
	ID string `json:"id,required"`
	// Synced entity id
	SyncedEntityID string `json:"syncedEntityId,required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier,required"`
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

// A customer can be either an organization or an individual
type V1CustomerListResponse struct {
	// Customer slug
	ID string `json:"id,required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt,required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Customer level coupon
	CouponID string `json:"couponId,nullable"`
	// The default payment method details
	DefaultPaymentMethod V1CustomerListResponseDefaultPaymentMethod `json:"defaultPaymentMethod,nullable"`
	// The email of the customer
	Email string `json:"email,nullable" format:"email"`
	// List of integrations
	Integrations []V1CustomerListResponseIntegration `json:"integrations"`
	// Additional metadata
	Metadata map[string]string `json:"metadata"`
	// The name of the customer
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		ArchivedAt           respjson.Field
		CreatedAt            respjson.Field
		UpdatedAt            respjson.Field
		CouponID             respjson.Field
		DefaultPaymentMethod respjson.Field
		Email                respjson.Field
		Integrations         respjson.Field
		Metadata             respjson.Field
		Name                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The default payment method details
type V1CustomerListResponseDefaultPaymentMethod struct {
	// The default payment method id
	BillingID string `json:"billingId,required"`
	// The expiration month of the default payment method
	CardExpiryMonth float64 `json:"cardExpiryMonth,required"`
	// The expiration year of the default payment method
	CardExpiryYear float64 `json:"cardExpiryYear,required"`
	// The last 4 digits of the default payment method
	CardLast4Digits string `json:"cardLast4Digits,required"`
	// The default payment method type
	//
	// Any of "CARD", "BANK", "CASH_APP".
	Type string `json:"type,required"`
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
	ID string `json:"id,required"`
	// Synced entity id
	SyncedEntityID string `json:"syncedEntityId,required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier,required"`
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

// Response object
type V1CustomerImportResponse struct {
	// List of newly created customer IDs from the import operation.
	Data V1CustomerImportResponseData `json:"data,required"`
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
	NewCustomers []string `json:"newCustomers,required"`
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
	ID string `json:"id,required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
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

type V1CustomerUpdateParams struct {
	// Customer level coupon
	CouponID param.Opt[string] `json:"couponId,omitzero"`
	// The email of the customer
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// The name of the customer
	Name param.Opt[string] `json:"name,omitzero"`
	// List of integrations
	Integrations []V1CustomerUpdateParamsIntegration `json:"integrations,omitzero"`
	// Additional metadata
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External billing or CRM integration link
//
// The properties ID, SyncedEntityID, VendorIdentifier are required.
type V1CustomerUpdateParamsIntegration struct {
	// Synced entity id
	SyncedEntityID param.Opt[string] `json:"syncedEntityId,omitzero,required"`
	// Integration details
	ID string `json:"id,required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier,omitzero,required"`
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
	Customers []V1CustomerImportParamsCustomer `json:"customers,omitzero,required"`
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
	Email param.Opt[string] `json:"email,omitzero,required" format:"email"`
	// The name of the customer
	Name param.Opt[string] `json:"name,omitzero,required"`
	// Customer slug
	ID string `json:"id,required"`
	// Billing provider payment method id
	PaymentMethodID param.Opt[string] `json:"paymentMethodId,omitzero"`
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
	ID string `json:"id,required"`
	// Customer level coupon
	CouponID param.Opt[string] `json:"couponId,omitzero"`
	// The email of the customer
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// The name of the customer
	Name param.Opt[string] `json:"name,omitzero"`
	// The default payment method details
	DefaultPaymentMethod V1CustomerProvisionParamsDefaultPaymentMethod `json:"defaultPaymentMethod,omitzero"`
	// List of integrations
	Integrations []V1CustomerProvisionParamsIntegration `json:"integrations,omitzero"`
	// Additional metadata
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1CustomerProvisionParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerProvisionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerProvisionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The default payment method details
//
// The properties BillingID, CardExpiryMonth, CardExpiryYear, CardLast4Digits, Type
// are required.
type V1CustomerProvisionParamsDefaultPaymentMethod struct {
	// The default payment method id
	BillingID param.Opt[string] `json:"billingId,omitzero,required"`
	// The expiration month of the default payment method
	CardExpiryMonth param.Opt[float64] `json:"cardExpiryMonth,omitzero,required"`
	// The expiration year of the default payment method
	CardExpiryYear param.Opt[float64] `json:"cardExpiryYear,omitzero,required"`
	// The last 4 digits of the default payment method
	CardLast4Digits param.Opt[string] `json:"cardLast4Digits,omitzero,required"`
	// The default payment method type
	//
	// Any of "CARD", "BANK", "CASH_APP".
	Type string `json:"type,omitzero,required"`
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
	SyncedEntityID param.Opt[string] `json:"syncedEntityId,omitzero,required"`
	// Integration details
	ID string `json:"id,required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier,omitzero,required"`
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
