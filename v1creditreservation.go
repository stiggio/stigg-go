// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"errors"
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

// V1CreditReservationService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CreditReservationService] method instead.
type V1CreditReservationService struct {
	Options []option.RequestOption
}

// NewV1CreditReservationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CreditReservationService(opts ...option.RequestOption) (r V1CreditReservationService) {
	r = V1CreditReservationService{}
	r.Options = opts
	return
}

// Releases a credit reservation, consuming the given amount and returning the rest
// of the held credits to the balance. Retry a failed release with the same
// reservation ID; it is applied once.
func (r *V1CreditReservationService) Release(ctx context.Context, reservationID string, params V1CreditReservationReleaseParams, opts ...option.RequestOption) (res *V1CreditReservationReleaseResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if reservationID == "" {
		err = errors.New("missing required reservationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/credits/reservations/%s/release", reservationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Response object
type V1CreditReservationReleaseResponse struct {
	// Result of releasing a credit reservation
	Data V1CreditReservationReleaseResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditReservationReleaseResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditReservationReleaseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of releasing a credit reservation
type V1CreditReservationReleaseResponseData struct {
	// The optimistic credit balance after the release
	Credit V1CreditReservationReleaseResponseDataCredit `json:"credit" api:"required"`
	// The released reservation
	Reservation V1CreditReservationReleaseResponseDataReservation `json:"reservation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Credit      respjson.Field
		Reservation respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditReservationReleaseResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CreditReservationReleaseResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The optimistic credit balance after the release
type V1CreditReservationReleaseResponseDataCredit struct {
	// The credit currency identifier
	CurrencyID string `json:"currencyId" api:"required"`
	// The wallet's total consumed credits for this currency (optimistic — includes
	// not-yet-reconciled usage), shared across every feature that draws on the
	// currency. This is the running balance, not this call's deduction — see
	// `consumed` for that.
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
func (r V1CreditReservationReleaseResponseDataCredit) RawJSON() string { return r.JSON.raw }
func (r *V1CreditReservationReleaseResponseDataCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The released reservation
type V1CreditReservationReleaseResponseDataReservation struct {
	// The client-provided reservation ID.
	ID string `json:"id" api:"required"`
	// The amount of credits held.
	Amount float64 `json:"amount" api:"required"`
	// The amount consumed at release
	ConsumedAmount float64 `json:"consumedAmount" api:"required"`
	// When the reservation expires. Unreleased credits return to the balance at this
	// time.
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// The status of the reservation.
	//
	// Any of "ACTIVE", "RELEASED", "EXPIRED".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Amount         respjson.Field
		ConsumedAmount respjson.Field
		ExpiresAt      respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditReservationReleaseResponseDataReservation) RawJSON() string { return r.JSON.raw }
func (r *V1CreditReservationReleaseResponseDataReservation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditReservationReleaseParams struct {
	// The amount of credits to consume at release. Up to the reserved amount is
	// charged to the credits that backed the reservation; the rest returns to the
	// balance, and any amount above the reservation is consumed normally.
	ConsumedAmount float64 `json:"consumedAmount" api:"required"`
	// The credit currency of the reservation
	CurrencyID string `json:"currencyId" api:"required"`
	// The customer the reservation was made for
	CustomerID string `json:"customerId" api:"required"`
	// The resource the reservation was made for, if any
	ResourceID     param.Opt[string] `json:"resourceId,omitzero"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1CreditReservationReleaseParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditReservationReleaseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditReservationReleaseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
