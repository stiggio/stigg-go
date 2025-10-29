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
	"github.com/stainless-sdks/stigg-go/packages/param"
	"github.com/stainless-sdks/stigg-go/packages/respjson"
)

// V1CustomerService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerService] method instead.
type V1CustomerService struct {
	Options     []option.RequestOption
	SubCustomer V1CustomerSubCustomerService
}

// NewV1CustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomerService(opts ...option.RequestOption) (r V1CustomerService) {
	r = V1CustomerService{}
	r.Options = opts
	r.SubCustomer = NewV1CustomerSubCustomerService(opts...)
	return
}

// Get a single customer by id
func (r *V1CustomerService) Get(ctx context.Context, refID string, query V1CustomerGetParams, opts ...option.RequestOption) (res *V1CustomerGetResponse, err error) {
	if !param.IsOmitted(query.XAPIKey) {
		opts = append(opts, option.WithHeader("X-API-KEY", fmt.Sprintf("%s", query.XAPIKey)))
	}
	if !param.IsOmitted(query.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%s", query.XEnvironmentID)))
	}
	opts = slices.Concat(r.Options, opts)
	if refID == "" {
		err = errors.New("missing required refId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s", refID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type V1CustomerGetResponse struct {
	// Unique identifier for the entity
	ID string `json:"id,required"`
	// The email of the customer
	Email string `json:"email,required"`
	// The name of the customer
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerGetParams struct {
	XAPIKey        string `header:"X-API-KEY,required" json:"-"`
	XEnvironmentID string `header:"X-ENVIRONMENT-ID,required" json:"-"`
	paramObj
}
