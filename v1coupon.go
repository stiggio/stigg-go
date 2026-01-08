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

	"github.com/stainless-sdks/stigg-go/internal/apijson"
	"github.com/stainless-sdks/stigg-go/internal/apiquery"
	"github.com/stainless-sdks/stigg-go/internal/requestconfig"
	"github.com/stainless-sdks/stigg-go/option"
	"github.com/stainless-sdks/stigg-go/packages/param"
	"github.com/stainless-sdks/stigg-go/packages/respjson"
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

// Create a new Coupon
func (r *V1CouponService) New(ctx context.Context, body V1CouponNewParams, opts ...option.RequestOption) (res *V1CouponNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/coupons"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single Coupon by id
func (r *V1CouponService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *V1CouponGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/coupons/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of Coupons
func (r *V1CouponService) List(ctx context.Context, query V1CouponListParams, opts ...option.RequestOption) (res *V1CouponListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/coupons"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type V1CouponNewResponse struct {
	Data V1CouponNewResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CouponNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CouponNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CouponNewResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id,required"`
	// Fixed amount discounts in different currencies
	AmountsOff []V1CouponNewResponseDataAmountsOff `json:"amountsOff,required"`
	// The unique identifier for the entity in the billing provider
	BillingID string `json:"billingId,required"`
	// The URL to the entity in the billing provider
	BillingLinkURL string `json:"billingLinkUrl,required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Description of the coupon
	Description string `json:"description,required"`
	// Duration of the coupon validity in months
	DurationInMonths float64 `json:"durationInMonths,required"`
	// Name of the coupon
	Name string `json:"name,required"`
	// Percentage discount off the original price
	PercentOff float64 `json:"percentOff,required"`
	// The source of the coupon
	//
	// Any of "STIGG", "STIGG_ADHOC", "STRIPE".
	Source string `json:"source,required"`
	// Current status of the coupon
	//
	// Any of "ACTIVE", "ARCHIVED".
	Status string `json:"status,required"`
	// Type of the coupon (percentage or fixed amount)
	//
	// Any of "FIXED", "PERCENTAGE".
	Type string `json:"type,required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
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
func (r V1CouponNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CouponNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CouponNewResponseDataAmountsOff struct {
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
	Currency string `json:"currency,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CouponNewResponseDataAmountsOff) RawJSON() string { return r.JSON.raw }
func (r *V1CouponNewResponseDataAmountsOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CouponGetResponse struct {
	Data V1CouponGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CouponGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CouponGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CouponGetResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id,required"`
	// Fixed amount discounts in different currencies
	AmountsOff []V1CouponGetResponseDataAmountsOff `json:"amountsOff,required"`
	// The unique identifier for the entity in the billing provider
	BillingID string `json:"billingId,required"`
	// The URL to the entity in the billing provider
	BillingLinkURL string `json:"billingLinkUrl,required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Description of the coupon
	Description string `json:"description,required"`
	// Duration of the coupon validity in months
	DurationInMonths float64 `json:"durationInMonths,required"`
	// Name of the coupon
	Name string `json:"name,required"`
	// Percentage discount off the original price
	PercentOff float64 `json:"percentOff,required"`
	// The source of the coupon
	//
	// Any of "STIGG", "STIGG_ADHOC", "STRIPE".
	Source string `json:"source,required"`
	// Current status of the coupon
	//
	// Any of "ACTIVE", "ARCHIVED".
	Status string `json:"status,required"`
	// Type of the coupon (percentage or fixed amount)
	//
	// Any of "FIXED", "PERCENTAGE".
	Type string `json:"type,required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
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
func (r V1CouponGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CouponGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CouponGetResponseDataAmountsOff struct {
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
	Currency string `json:"currency,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CouponGetResponseDataAmountsOff) RawJSON() string { return r.JSON.raw }
func (r *V1CouponGetResponseDataAmountsOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CouponListResponse struct {
	Data []V1CouponListResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CouponListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CouponListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CouponListResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id,required"`
	// Fixed amount discounts in different currencies
	AmountsOff []V1CouponListResponseDataAmountsOff `json:"amountsOff,required"`
	// The unique identifier for the entity in the billing provider
	BillingID string `json:"billingId,required"`
	// The URL to the entity in the billing provider
	BillingLinkURL string `json:"billingLinkUrl,required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Cursor ID for query pagination
	CursorID string `json:"cursorId,required" format:"uuid"`
	// Description of the coupon
	Description string `json:"description,required"`
	// Duration of the coupon validity in months
	DurationInMonths float64 `json:"durationInMonths,required"`
	// Name of the coupon
	Name string `json:"name,required"`
	// Percentage discount off the original price
	PercentOff float64 `json:"percentOff,required"`
	// The source of the coupon
	//
	// Any of "STIGG", "STIGG_ADHOC", "STRIPE".
	Source string `json:"source,required"`
	// Current status of the coupon
	//
	// Any of "ACTIVE", "ARCHIVED".
	Status string `json:"status,required"`
	// Type of the coupon (percentage or fixed amount)
	//
	// Any of "FIXED", "PERCENTAGE".
	Type string `json:"type,required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AmountsOff       respjson.Field
		BillingID        respjson.Field
		BillingLinkURL   respjson.Field
		CreatedAt        respjson.Field
		CursorID         respjson.Field
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
func (r V1CouponListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CouponListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CouponListResponseDataAmountsOff struct {
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
	Currency string `json:"currency,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CouponListResponseDataAmountsOff) RawJSON() string { return r.JSON.raw }
func (r *V1CouponListResponseDataAmountsOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CouponNewParams struct {
	// Description of the coupon
	Description param.Opt[string] `json:"description,omitzero,required"`
	// Duration of the coupon validity in months
	DurationInMonths param.Opt[int64] `json:"durationInMonths,omitzero,required"`
	// Percentage discount off the original price
	PercentOff param.Opt[float64] `json:"percentOff,omitzero,required"`
	// Fixed amount discounts in different currencies
	AmountsOff []V1CouponNewParamsAmountsOff `json:"amountsOff,omitzero,required"`
	// The unique identifier for the entity
	ID string `json:"id,required"`
	// Name of the coupon
	Name string `json:"name,required"`
	// Metadata associated with the entity
	AdditionalMetaData any `json:"additionalMetaData,omitzero"`
	paramObj
}

func (r V1CouponNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CouponNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CouponNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Currency are required.
type V1CouponNewParamsAmountsOff struct {
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
	// Ending before this UUID for pagination
	EndingBefore param.Opt[string] `query:"endingBefore,omitzero" format:"uuid" json:"-"`
	// Items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Starting after this UUID for pagination
	StartingAfter param.Opt[string] `query:"startingAfter,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [V1CouponListParams]'s query parameters as `url.Values`.
func (r V1CouponListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
