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

// V1EventPlanDraftService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventPlanDraftService] method instead.
type V1EventPlanDraftService struct {
	Options []option.RequestOption
}

// NewV1EventPlanDraftService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1EventPlanDraftService(opts ...option.RequestOption) (r V1EventPlanDraftService) {
	r = V1EventPlanDraftService{}
	r.Options = opts
	return
}

// Creates a draft version of an existing plan for modification before publishing.
func (r *V1EventPlanDraftService) New(ctx context.Context, id string, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s/draft", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Removes a draft version of a plan.
func (r *V1EventPlanDraftService) Remove(ctx context.Context, id string, opts ...option.RequestOption) (res *V1EventPlanDraftRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/plans/%s/draft", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Response confirming the plan draft was removed.
type V1EventPlanDraftRemoveResponse struct {
	Data V1EventPlanDraftRemoveResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanDraftRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanDraftRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventPlanDraftRemoveResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventPlanDraftRemoveResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventPlanDraftRemoveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
