// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/stigg-go/internal/apijson"
	"github.com/stainless-sdks/stigg-go/internal/requestconfig"
	"github.com/stainless-sdks/stigg-go/option"
	"github.com/stainless-sdks/stigg-go/packages/respjson"
)

// V1SubscriptionFutureUpdateService contains methods and other services that help
// with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1SubscriptionFutureUpdateService] method instead.
type V1SubscriptionFutureUpdateService struct {
	Options []option.RequestOption
}

// NewV1SubscriptionFutureUpdateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1SubscriptionFutureUpdateService(opts ...option.RequestOption) (r V1SubscriptionFutureUpdateService) {
	r = V1SubscriptionFutureUpdateService{}
	r.Options = opts
	return
}

// Perform cancel future update on a Subscription
func (r *V1SubscriptionFutureUpdateService) CancelPendingPayment(ctx context.Context, id string, opts ...option.RequestOption) (res *V1SubscriptionFutureUpdateCancelPendingPaymentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/future-update/pending-payment", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Perform cancel future update on a Subscription
func (r *V1SubscriptionFutureUpdateService) CancelSchedule(ctx context.Context, id string, opts ...option.RequestOption) (res *V1SubscriptionFutureUpdateCancelScheduleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/future-update/schedule", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type V1SubscriptionFutureUpdateCancelPendingPaymentResponse struct {
	Data V1SubscriptionFutureUpdateCancelPendingPaymentResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionFutureUpdateCancelPendingPaymentResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionFutureUpdateCancelPendingPaymentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionFutureUpdateCancelPendingPaymentResponseData struct {
	// external id of the canceled future update subscription
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionFutureUpdateCancelPendingPaymentResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1SubscriptionFutureUpdateCancelPendingPaymentResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionFutureUpdateCancelScheduleResponse struct {
	Data V1SubscriptionFutureUpdateCancelScheduleResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionFutureUpdateCancelScheduleResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionFutureUpdateCancelScheduleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SubscriptionFutureUpdateCancelScheduleResponseData struct {
	// external id of the canceled future update subscription
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionFutureUpdateCancelScheduleResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionFutureUpdateCancelScheduleResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
