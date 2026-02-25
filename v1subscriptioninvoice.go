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

// V1SubscriptionInvoiceService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1SubscriptionInvoiceService] method instead.
type V1SubscriptionInvoiceService struct {
	Options []option.RequestOption
}

// NewV1SubscriptionInvoiceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1SubscriptionInvoiceService(opts ...option.RequestOption) (r V1SubscriptionInvoiceService) {
	r = V1SubscriptionInvoiceService{}
	r.Options = opts
	return
}

// Marks the latest invoice of a subscription as paid in the billing provider. The
// invoice must exist and have an OPEN status.
func (r *V1SubscriptionInvoiceService) MarkAsPaid(ctx context.Context, id string, opts ...option.RequestOption) (res *V1SubscriptionInvoiceMarkAsPaidResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s/invoice/paid", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Response object
type V1SubscriptionInvoiceMarkAsPaidResponse struct {
	// Result of marking a subscription invoice as paid.
	Data V1SubscriptionInvoiceMarkAsPaidResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionInvoiceMarkAsPaidResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionInvoiceMarkAsPaidResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of marking a subscription invoice as paid.
type V1SubscriptionInvoiceMarkAsPaidResponseData struct {
	// The subscription ID whose invoice was marked as paid
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SubscriptionInvoiceMarkAsPaidResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SubscriptionInvoiceMarkAsPaidResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
