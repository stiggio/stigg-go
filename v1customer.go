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
	Usage         V1CustomerUsageService
}

// NewV1CustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomerService(opts ...option.RequestOption) (r V1CustomerService) {
	r = V1CustomerService{}
	r.Options = opts
	r.PaymentMethod = NewV1CustomerPaymentMethodService(opts...)
	r.Usage = NewV1CustomerUsageService(opts...)
	return
}

// Create a new Customer
func (r *V1CustomerService) New(ctx context.Context, body V1CustomerNewParams, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single Customer by id
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

// Update an existing Customer
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

// Get a list of Customers
func (r *V1CustomerService) List(ctx context.Context, query V1CustomerListParams, opts ...option.RequestOption) (res *V1CustomerListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Perform archive on a Customer
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

// Perform unarchive on a Customer
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

type CustomerResponse struct {
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

type V1CustomerListResponse struct {
	Data []V1CustomerListResponseData `json:"data,required"`
	// Pagination information including cursors for navigation
	Pagination V1CustomerListResponsePagination `json:"pagination,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerListResponseData struct {
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
	DefaultPaymentMethod V1CustomerListResponseDataDefaultPaymentMethod `json:"defaultPaymentMethod,nullable"`
	// The email of the customer
	Email string `json:"email,nullable" format:"email"`
	// List of integrations
	Integrations []V1CustomerListResponseDataIntegration `json:"integrations"`
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
func (r V1CustomerListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The default payment method details
type V1CustomerListResponseDataDefaultPaymentMethod struct {
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
func (r V1CustomerListResponseDataDefaultPaymentMethod) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponseDataDefaultPaymentMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerListResponseDataIntegration struct {
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
func (r V1CustomerListResponseDataIntegration) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponseDataIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination information including cursors for navigation
type V1CustomerListResponsePagination struct {
	// Cursor to fetch the next page (use with after parameter), null if no more pages
	Next string `json:"next,required" format:"uuid"`
	// Cursor to fetch the previous page (use with before parameter), null if no
	// previous pages
	Prev string `json:"prev,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		Prev        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerNewParams struct {
	// Customer slug
	ID string `json:"id,required"`
	// Customer level coupon
	CouponID param.Opt[string] `json:"couponId,omitzero"`
	// The email of the customer
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// The name of the customer
	Name param.Opt[string] `json:"name,omitzero"`
	// The default payment method details
	DefaultPaymentMethod V1CustomerNewParamsDefaultPaymentMethod `json:"defaultPaymentMethod,omitzero"`
	// List of integrations
	Integrations []V1CustomerNewParamsIntegration `json:"integrations,omitzero"`
	// Additional metadata
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1CustomerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The default payment method details
//
// The properties BillingID, CardExpiryMonth, CardExpiryYear, CardLast4Digits, Type
// are required.
type V1CustomerNewParamsDefaultPaymentMethod struct {
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

func (r V1CustomerNewParamsDefaultPaymentMethod) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerNewParamsDefaultPaymentMethod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerNewParamsDefaultPaymentMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerNewParamsDefaultPaymentMethod](
		"type", "CARD", "BANK", "CASH_APP",
	)
}

// The properties ID, SyncedEntityID, VendorIdentifier are required.
type V1CustomerNewParamsIntegration struct {
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

func (r V1CustomerNewParamsIntegration) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerNewParamsIntegration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerNewParamsIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CustomerNewParamsIntegration](
		"vendorIdentifier", "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE", "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE",
	)
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
	// Starting after this UUID for pagination
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Ending before this UUID for pagination
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerListParams]'s query parameters as `url.Values`.
func (r V1CustomerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
