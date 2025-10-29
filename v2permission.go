// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/stigg-go/internal/apijson"
	"github.com/stainless-sdks/stigg-go/internal/apiquery"
	"github.com/stainless-sdks/stigg-go/internal/requestconfig"
	"github.com/stainless-sdks/stigg-go/option"
	"github.com/stainless-sdks/stigg-go/packages/param"
	"github.com/stainless-sdks/stigg-go/packages/respjson"
)

// V2PermissionService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2PermissionService] method instead.
type V2PermissionService struct {
	Options []option.RequestOption
}

// NewV2PermissionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2PermissionService(opts ...option.RequestOption) (r V2PermissionService) {
	r = V2PermissionService{}
	r.Options = opts
	return
}

// Check multiple permissions for a user
func (r *V2PermissionService) Check(ctx context.Context, params V2PermissionCheckParams, opts ...option.RequestOption) (res *V2PermissionCheckResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/permissions/check"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Response for checking permissions
type V2PermissionCheckResponse struct {
	// List of booleans indicating whether the user has access to each resource and
	// action correspondingly
	PermittedList []bool `json:"permittedList,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PermittedList respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2PermissionCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *V2PermissionCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2PermissionCheckParams struct {
	// ID of the user to check permissions for
	UserID string `query:"userId,required" json:"-"`
	// List of resources and actions to check permissions for
	ResourcesAndActions []V2PermissionCheckParamsResourcesAndAction `json:"resourcesAndActions,omitzero,required"`
	paramObj
}

func (r V2PermissionCheckParams) MarshalJSON() (data []byte, err error) {
	type shadow V2PermissionCheckParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2PermissionCheckParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V2PermissionCheckParams]'s query parameters as
// `url.Values`.
func (r V2PermissionCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Data transfer object for resource and action pair
//
// The properties Action, Resource are required.
type V2PermissionCheckParamsResourcesAndAction struct {
	// Action to check permissions for
	Action any `json:"action,omitzero,required"`
	// Resource to check permissions for
	Resource string `json:"resource,required"`
	paramObj
}

func (r V2PermissionCheckParamsResourcesAndAction) MarshalJSON() (data []byte, err error) {
	type shadow V2PermissionCheckParamsResourcesAndAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2PermissionCheckParamsResourcesAndAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
