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

// V1EventBetaCustomerAssignmentService contains methods and other services that
// help with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventBetaCustomerAssignmentService] method instead.
type V1EventBetaCustomerAssignmentService struct {
	Options []option.RequestOption
}

// NewV1EventBetaCustomerAssignmentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1EventBetaCustomerAssignmentService(opts ...option.RequestOption) (r V1EventBetaCustomerAssignmentService) {
	r = V1EventBetaCustomerAssignmentService{}
	r.Options = opts
	return
}

// Returns a cursor-paginated list of capability assignments for the given
// customer. An assignment ties an entity to a capability with a usage limit and
// reset cadence.
func (r *V1EventBetaCustomerAssignmentService) List(ctx context.Context, id string, params V1EventBetaCustomerAssignmentListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1EventBetaCustomerAssignmentListResponse], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1-beta/customers/%s/assignments", id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Returns a cursor-paginated list of capability assignments for the given
// customer. An assignment ties an entity to a capability with a usage limit and
// reset cadence.
func (r *V1EventBetaCustomerAssignmentService) ListAutoPaging(ctx context.Context, id string, params V1EventBetaCustomerAssignmentListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1EventBetaCustomerAssignmentListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, id, params, opts...))
}

// Batched create-or-update of capability assignments. Existing assignments matched
// by (entityId, capabilityId) are updated; new pairs are created. On update,
// omitted fields (usageLimit, cadence) are preserved; on create both are required
// by the governance service.
func (r *V1EventBetaCustomerAssignmentService) Upsert(ctx context.Context, id string, params V1EventBetaCustomerAssignmentUpsertParams, opts ...option.RequestOption) (res *V1EventBetaCustomerAssignmentUpsertResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1-beta/customers/%s/assignments", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A capability assignment for an entity belonging to a customer. Defines how much
// of the capability the entity may consume (`usageLimit`) and how often the
// counter resets (`cadence`).
type V1EventBetaCustomerAssignmentListResponse struct {
	// Synthetic UUID identifier — also the cursor anchor for paginated lists
	ID string `json:"id" api:"required" format:"uuid"`
	// Usage-reset cadence. Currently only `MONTH` is supported
	//
	// Any of "MONTH".
	Cadence V1EventBetaCustomerAssignmentListResponseCadence `json:"cadence" api:"required"`
	// The capability refId this assignment grants
	CapabilityID string `json:"capabilityId" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The entity refId this assignment is attached to
	EntityID string `json:"entityId" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Maximum usage allowed within one cadence window
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Cadence      respjson.Field
		CapabilityID respjson.Field
		CreatedAt    respjson.Field
		EntityID     respjson.Field
		UpdatedAt    respjson.Field
		UsageLimit   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerAssignmentListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerAssignmentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage-reset cadence. Currently only `MONTH` is supported
type V1EventBetaCustomerAssignmentListResponseCadence string

const (
	V1EventBetaCustomerAssignmentListResponseCadenceMonth V1EventBetaCustomerAssignmentListResponseCadence = "MONTH"
)

// Assignments after upsert.
type V1EventBetaCustomerAssignmentUpsertResponse struct {
	Data []V1EventBetaCustomerAssignmentUpsertResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerAssignmentUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerAssignmentUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A capability assignment for an entity belonging to a customer. Defines how much
// of the capability the entity may consume (`usageLimit`) and how often the
// counter resets (`cadence`).
type V1EventBetaCustomerAssignmentUpsertResponseData struct {
	// Synthetic UUID identifier — also the cursor anchor for paginated lists
	ID string `json:"id" api:"required" format:"uuid"`
	// Usage-reset cadence. Currently only `MONTH` is supported
	//
	// Any of "MONTH".
	Cadence string `json:"cadence" api:"required"`
	// The capability refId this assignment grants
	CapabilityID string `json:"capabilityId" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The entity refId this assignment is attached to
	EntityID string `json:"entityId" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Maximum usage allowed within one cadence window
	UsageLimit float64 `json:"usageLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Cadence      respjson.Field
		CapabilityID respjson.Field
		CreatedAt    respjson.Field
		EntityID     respjson.Field
		UpdatedAt    respjson.Field
		UsageLimit   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventBetaCustomerAssignmentUpsertResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventBetaCustomerAssignmentUpsertResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventBetaCustomerAssignmentListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Filter assignments to a specific capability refId
	CapabilityID param.Opt[string] `query:"capabilityId,omitzero" json:"-"`
	// Filter assignments to a specific entity refId
	EntityID param.Opt[string] `query:"entityId,omitzero" json:"-"`
	// Maximum number of items to return
	Limit          param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventBetaCustomerAssignmentListParams]'s query parameters
// as `url.Values`.
func (r V1EventBetaCustomerAssignmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1EventBetaCustomerAssignmentUpsertParams struct {
	// Assignments to upsert (1–100 per request)
	Assignments    []V1EventBetaCustomerAssignmentUpsertParamsAssignment `json:"assignments,omitzero" api:"required"`
	XAccountID     param.Opt[string]                                     `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string]                                     `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1EventBetaCustomerAssignmentUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventBetaCustomerAssignmentUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventBetaCustomerAssignmentUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single assignment to create or update. The natural key is the
// `(entityId, capabilityId)` pair. On create both `usageLimit` and `cadence` are
// required; on update they may be omitted individually to preserve the existing
// value.
//
// The properties CapabilityID, EntityID are required.
type V1EventBetaCustomerAssignmentUpsertParamsAssignment struct {
	// The capability refId this assignment grants
	CapabilityID string `json:"capabilityId" api:"required"`
	// The entity refId this assignment is attached to
	EntityID string `json:"entityId" api:"required"`
	// Maximum usage allowed within one cadence window (required on create)
	UsageLimit param.Opt[float64] `json:"usageLimit,omitzero"`
	// Usage-reset cadence (required on create). Currently only `MONTH` is supported
	//
	// Any of "MONTH".
	Cadence string `json:"cadence,omitzero"`
	paramObj
}

func (r V1EventBetaCustomerAssignmentUpsertParamsAssignment) MarshalJSON() (data []byte, err error) {
	type shadow V1EventBetaCustomerAssignmentUpsertParamsAssignment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventBetaCustomerAssignmentUpsertParamsAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventBetaCustomerAssignmentUpsertParamsAssignment](
		"cadence", "MONTH",
	)
}
