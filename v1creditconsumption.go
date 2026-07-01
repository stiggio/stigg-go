// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stiggio/stigg-go/internal/apijson"
	"github.com/stiggio/stigg-go/internal/requestconfig"
	"github.com/stiggio/stigg-go/option"
	"github.com/stiggio/stigg-go/packages/param"
	"github.com/stiggio/stigg-go/packages/respjson"
)

// V1CreditConsumptionService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CreditConsumptionService] method instead.
type V1CreditConsumptionService struct {
	Options []option.RequestOption
}

// NewV1CreditConsumptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CreditConsumptionService(opts ...option.RequestOption) (r V1CreditConsumptionService) {
	r = V1CreditConsumptionService{}
	r.Options = opts
	return
}

// Consumes a specified amount of credits directly from a customer wallet, with no
// feature mapping. Returns the optimistic balance.
func (r *V1CreditConsumptionService) Consume(ctx context.Context, params V1CreditConsumptionConsumeParams, opts ...option.RequestOption) (res *V1CreditConsumptionConsumeResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/credits/consumption"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Consumes credits directly from customer wallets asynchronously. Consumptions are
// reconciled asynchronously into the credit balances.
func (r *V1CreditConsumptionService) ConsumeAsync(ctx context.Context, params V1CreditConsumptionConsumeAsyncParams, opts ...option.RequestOption) (res *V1CreditConsumptionConsumeAsyncResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/credits/consumption/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Response object
type V1CreditConsumptionConsumeResponse struct {
	// Result of a synchronous direct credit consumption
	Data V1CreditConsumptionConsumeResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditConsumptionConsumeResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditConsumptionConsumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of a synchronous direct credit consumption
type V1CreditConsumptionConsumeResponseData struct {
	// The amount of credits consumed
	Amount float64 `json:"amount" api:"required"`
	// The credit currency the credits were consumed from
	CurrencyID string `json:"currencyId" api:"required"`
	// The customer the credits were consumed from
	CustomerID string `json:"customerId" api:"required"`
	// The timestamp the consumption was attributed to
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The optimistic credit balance after consumption (when sync credit consumption is
	// enabled)
	Credit V1CreditConsumptionConsumeResponseDataCredit `json:"credit" api:"nullable"`
	// The resource the consumption was attributed to
	ResourceID string `json:"resourceId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		CurrencyID  respjson.Field
		CustomerID  respjson.Field
		Timestamp   respjson.Field
		Credit      respjson.Field
		ResourceID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditConsumptionConsumeResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CreditConsumptionConsumeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The optimistic credit balance after consumption (when sync credit consumption is
// enabled)
type V1CreditConsumptionConsumeResponseDataCredit struct {
	// The credit currency identifier
	CurrencyID string `json:"currencyId" api:"required"`
	// The credits consumed (optimistic — includes not-yet-reconciled usage)
	CurrentUsage float64 `json:"currentUsage" api:"required"`
	// The grant-version timestamp of this balance, used by the SDK for last-write-wins
	// reconciliation
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The total credits granted
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// End of the current credit grant period (when recurring credits reset), if
	// applicable
	UsagePeriodEnd time.Time `json:"usagePeriodEnd" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyID     respjson.Field
		CurrentUsage   respjson.Field
		Timestamp      respjson.Field
		UsageLimit     respjson.Field
		UsagePeriodEnd respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditConsumptionConsumeResponseDataCredit) RawJSON() string { return r.JSON.raw }
func (r *V1CreditConsumptionConsumeResponseDataCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1CreditConsumptionConsumeAsyncResponse struct {
	// Confirmation that the credit consumptions were accepted for processing
	Data any `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditConsumptionConsumeAsyncResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditConsumptionConsumeAsyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditConsumptionConsumeParams struct {
	// The amount of credits to consume
	Amount float64 `json:"amount" api:"required"`
	// The credit currency to consume from (required)
	CurrencyID string `json:"currencyId" api:"required"`
	// The customer to consume credits from (required)
	CustomerID string `json:"customerId" api:"required"`
	// A unique key used to deduplicate the consumption (required)
	IdempotencyKey string `json:"idempotencyKey" api:"required"`
	// Optional timestamp the consumption is attributed to
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Optional resource the consumption is attributed to
	ResourceID     param.Opt[string] `json:"resourceId,omitzero"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	// Optional dimensions describing the consumption
	Dimensions map[string]V1CreditConsumptionConsumeParamsDimensionUnion `json:"dimensions,omitzero"`
	paramObj
}

func (r V1CreditConsumptionConsumeParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditConsumptionConsumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditConsumptionConsumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1CreditConsumptionConsumeParamsDimensionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u V1CreditConsumptionConsumeParamsDimensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *V1CreditConsumptionConsumeParamsDimensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1CreditConsumptionConsumeParamsDimensionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type V1CreditConsumptionConsumeAsyncParams struct {
	// The credit consumptions to report (up to 1000)
	Consumptions   []V1CreditConsumptionConsumeAsyncParamsConsumption `json:"consumptions,omitzero" api:"required"`
	XAccountID     param.Opt[string]                                  `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]                                  `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1CreditConsumptionConsumeAsyncParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditConsumptionConsumeAsyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditConsumptionConsumeAsyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for consuming credits directly from a wallet
//
// The properties Amount, CurrencyID, CustomerID, IdempotencyKey are required.
type V1CreditConsumptionConsumeAsyncParamsConsumption struct {
	// The amount of credits to consume
	Amount float64 `json:"amount" api:"required"`
	// The credit currency to consume from (required)
	CurrencyID string `json:"currencyId" api:"required"`
	// The customer to consume credits from (required)
	CustomerID string `json:"customerId" api:"required"`
	// A unique key used to deduplicate the consumption (required)
	IdempotencyKey string `json:"idempotencyKey" api:"required"`
	// Optional timestamp the consumption is attributed to
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Optional resource the consumption is attributed to
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Optional dimensions describing the consumption
	Dimensions map[string]V1CreditConsumptionConsumeAsyncParamsConsumptionDimensionUnion `json:"dimensions,omitzero"`
	paramObj
}

func (r V1CreditConsumptionConsumeAsyncParamsConsumption) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditConsumptionConsumeAsyncParamsConsumption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditConsumptionConsumeAsyncParamsConsumption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1CreditConsumptionConsumeAsyncParamsConsumptionDimensionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u V1CreditConsumptionConsumeAsyncParamsConsumptionDimensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *V1CreditConsumptionConsumeAsyncParamsConsumptionDimensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1CreditConsumptionConsumeAsyncParamsConsumptionDimensionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}
