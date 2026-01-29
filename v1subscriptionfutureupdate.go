// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stiggio/stigg-go/internal/apijson"
	"github.com/stiggio/stigg-go/internal/requestconfig"
	"github.com/stiggio/stigg-go/option"
	"github.com/stiggio/stigg-go/packages/respjson"
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

// Cancel pending payment update
func (r *V1SubscriptionFutureUpdateService) CancelPendingPayment(ctx context.Context, id string, opts ...option.RequestOption) (res *CancelSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/future-update/pending-payment", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Cancel scheduled update
func (r *V1SubscriptionFutureUpdateService) CancelSchedule(ctx context.Context, id string, opts ...option.RequestOption) (res *CancelSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/future-update/schedule", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Response object
type CancelSubscription struct {
	Data CancelSubscriptionData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CancelSubscription) RawJSON() string { return r.JSON.raw }
func (r *CancelSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CancelSubscriptionData struct {
	// Subscription ID
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CancelSubscriptionData) RawJSON() string { return r.JSON.raw }
func (r *CancelSubscriptionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
