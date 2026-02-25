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

// V1CouponService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CouponService] method instead.
type V1CouponService struct {
	Options []option.RequestOption
}

// NewV1CouponService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CouponService(opts ...option.RequestOption) (r V1CouponService) {
	r = V1CouponService{}
	r.Options = opts
	return
}

// Creates a new discount coupon with percentage or fixed amount off, applicable to
// customer subscriptions.
func (r *V1CouponService) New(ctx context.Context, body V1CouponNewParams, opts ...option.RequestOption) (res *Coupon, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/coupons"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a coupon by its unique identifier.
func (r *V1CouponService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Coupon, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/coupons/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a paginated list of coupons in the environment.
func (r *V1CouponService) List(ctx context.Context, query V1CouponListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1CouponListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/coupons"
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

// Retrieves a paginated list of coupons in the environment.
func (r *V1CouponService) ListAutoPaging(ctx context.Context, query V1CouponListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1CouponListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, query, opts...))
}

// Archives a coupon, preventing it from being applied to new subscriptions.
func (r *V1CouponService) ArchiveCoupon(ctx context.Context, id string, opts ...option.RequestOption) (res *Coupon, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/coupons/%s/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Updates an existing coupon's properties such as name, description, and metadata.
func (r *V1CouponService) UpdateCoupon(ctx context.Context, id string, body V1CouponUpdateCouponParams, opts ...option.RequestOption) (res *Coupon, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/coupons/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Response object
type Coupon struct {
	// Discount instrument with percentage or fixed amount
	Data CouponData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Coupon) RawJSON() string { return r.JSON.raw }
func (r *Coupon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discount instrument with percentage or fixed amount
type CouponData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Fixed amount discounts in different currencies
	AmountsOff []CouponDataAmountsOff `json:"amountsOff" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID string `json:"billingId" api:"required"`
	// The URL to the entity in the billing provider
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description of the coupon
	Description string `json:"description" api:"required"`
	// Duration of the coupon validity in months
	DurationInMonths float64 `json:"durationInMonths" api:"required"`
	// Name of the coupon
	Name string `json:"name" api:"required"`
	// Percentage discount off the original price
	PercentOff float64 `json:"percentOff" api:"required"`
	// The source of the coupon
	//
	// Any of "STIGG", "STIGG_ADHOC", "STRIPE".
	Source string `json:"source" api:"required"`
	// Current status of the coupon
	//
	// Any of "ACTIVE", "ARCHIVED".
	Status string `json:"status" api:"required"`
	// Type of the coupon (percentage or fixed amount)
	//
	// Any of "FIXED", "PERCENTAGE".
	Type string `json:"type" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AmountsOff       respjson.Field
		BillingID        respjson.Field
		BillingLinkURL   respjson.Field
		CreatedAt        respjson.Field
		Description      respjson.Field
		DurationInMonths respjson.Field
		Name             respjson.Field
		PercentOff       respjson.Field
		Source           respjson.Field
		Status           respjson.Field
		Type             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CouponData) RawJSON() string { return r.JSON.raw }
func (r *CouponData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monetary amount with currency
type CouponDataAmountsOff struct {
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
func (r CouponDataAmountsOff) RawJSON() string { return r.JSON.raw }
func (r *CouponDataAmountsOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discount instrument with percentage or fixed amount
type V1CouponListResponse struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Fixed amount discounts in different currencies
	AmountsOff []V1CouponListResponseAmountsOff `json:"amountsOff" api:"required"`
	// The unique identifier for the entity in the billing provider
	BillingID string `json:"billingId" api:"required"`
	// The URL to the entity in the billing provider
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description of the coupon
	Description string `json:"description" api:"required"`
	// Duration of the coupon validity in months
	DurationInMonths float64 `json:"durationInMonths" api:"required"`
	// Name of the coupon
	Name string `json:"name" api:"required"`
	// Percentage discount off the original price
	PercentOff float64 `json:"percentOff" api:"required"`
	// The source of the coupon
	//
	// Any of "STIGG", "STIGG_ADHOC", "STRIPE".
	Source V1CouponListResponseSource `json:"source" api:"required"`
	// Current status of the coupon
	//
	// Any of "ACTIVE", "ARCHIVED".
	Status V1CouponListResponseStatus `json:"status" api:"required"`
	// Type of the coupon (percentage or fixed amount)
	//
	// Any of "FIXED", "PERCENTAGE".
	Type V1CouponListResponseType `json:"type" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AmountsOff       respjson.Field
		BillingID        respjson.Field
		BillingLinkURL   respjson.Field
		CreatedAt        respjson.Field
		Description      respjson.Field
		DurationInMonths respjson.Field
		Name             respjson.Field
		PercentOff       respjson.Field
		Source           respjson.Field
		Status           respjson.Field
		Type             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CouponListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CouponListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monetary amount with currency
type V1CouponListResponseAmountsOff struct {
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
func (r V1CouponListResponseAmountsOff) RawJSON() string { return r.JSON.raw }
func (r *V1CouponListResponseAmountsOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The source of the coupon
type V1CouponListResponseSource string

const (
	V1CouponListResponseSourceStigg      V1CouponListResponseSource = "STIGG"
	V1CouponListResponseSourceStiggAdhoc V1CouponListResponseSource = "STIGG_ADHOC"
	V1CouponListResponseSourceStripe     V1CouponListResponseSource = "STRIPE"
)

// Current status of the coupon
type V1CouponListResponseStatus string

const (
	V1CouponListResponseStatusActive   V1CouponListResponseStatus = "ACTIVE"
	V1CouponListResponseStatusArchived V1CouponListResponseStatus = "ARCHIVED"
)

// Type of the coupon (percentage or fixed amount)
type V1CouponListResponseType string

const (
	V1CouponListResponseTypeFixed      V1CouponListResponseType = "FIXED"
	V1CouponListResponseTypePercentage V1CouponListResponseType = "PERCENTAGE"
)

type V1CouponNewParams struct {
	// Description of the coupon
	Description param.Opt[string] `json:"description,omitzero" api:"required"`
	// Duration of the coupon validity in months
	DurationInMonths param.Opt[int64] `json:"durationInMonths,omitzero" api:"required"`
	// Percentage discount off the original price
	PercentOff param.Opt[float64] `json:"percentOff,omitzero" api:"required"`
	// Fixed amount discounts in different currencies
	AmountsOff []V1CouponNewParamsAmountsOff `json:"amountsOff,omitzero" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,omitzero" api:"required"`
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// Name of the coupon
	Name string `json:"name" api:"required"`
	paramObj
}

func (r V1CouponNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CouponNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CouponNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monetary amount with currency
//
// The properties Amount, Currency are required.
type V1CouponNewParamsAmountsOff struct {
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
	Currency string `json:"currency,omitzero" api:"required"`
	paramObj
}

func (r V1CouponNewParamsAmountsOff) MarshalJSON() (data []byte, err error) {
	type shadow V1CouponNewParamsAmountsOff
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CouponNewParamsAmountsOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1CouponNewParamsAmountsOff](
		"currency", "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd", "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad", "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd", "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr", "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp", "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro", "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk", "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr", "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd", "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw", "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf",
	)
}

type V1CouponListParams struct {
	// Filter by entity ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by coupon status. Supports comma-separated values for multiple statuses
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter by creation date using range operators: gt, gte, lt, lte
	CreatedAt V1CouponListParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	// Filter by coupon type (FIXED or PERCENTAGE)
	//
	// Any of "FIXED", "PERCENTAGE".
	Type V1CouponListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CouponListParams]'s query parameters as `url.Values`.
func (r V1CouponListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by creation date using range operators: gt, gte, lt, lte
type V1CouponListParamsCreatedAt struct {
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

// URLQuery serializes [V1CouponListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r V1CouponListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by coupon type (FIXED or PERCENTAGE)
type V1CouponListParamsType string

const (
	V1CouponListParamsTypeFixed      V1CouponListParamsType = "FIXED"
	V1CouponListParamsTypePercentage V1CouponListParamsType = "PERCENTAGE"
)

type V1CouponUpdateCouponParams struct {
	// Description of the coupon
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the coupon
	Name param.Opt[string] `json:"name,omitzero"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1CouponUpdateCouponParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CouponUpdateCouponParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CouponUpdateCouponParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
