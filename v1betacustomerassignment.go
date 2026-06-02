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

// V1BetaCustomerAssignmentService contains methods and other services that help
// with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BetaCustomerAssignmentService] method instead.
type V1BetaCustomerAssignmentService struct {
	Options []option.RequestOption
}

// NewV1BetaCustomerAssignmentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1BetaCustomerAssignmentService(opts ...option.RequestOption) (r V1BetaCustomerAssignmentService) {
	r = V1BetaCustomerAssignmentService{}
	r.Options = opts
	return
}

// Returns a cursor-paginated list of capability assignments for the given
// customer. An assignment ties an entity to a capability with a usage limit and
// reset cadence.
func (r *V1BetaCustomerAssignmentService) List(ctx context.Context, id string, query V1BetaCustomerAssignmentListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1BetaCustomerAssignmentListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1-beta/customers/%s/assignments", id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *V1BetaCustomerAssignmentService) ListAutoPaging(ctx context.Context, id string, query V1BetaCustomerAssignmentListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1BetaCustomerAssignmentListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, id, query, opts...))
}

// Batched create-or-update of capability assignments. Existing assignments matched
// by (entityId, capabilityId) are updated; new pairs are created. On update,
// omitted fields (usageLimit, cadence) are preserved; on create both are required
// by the governance service.
func (r *V1BetaCustomerAssignmentService) Upsert(ctx context.Context, id string, body V1BetaCustomerAssignmentUpsertParams, opts ...option.RequestOption) (res *V1BetaCustomerAssignmentUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1-beta/customers/%s/assignments", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// A capability assignment for an entity belonging to a customer. Defines how much
// of the capability the entity may consume (`usageLimit`) and how often the
// counter resets (`cadence`).
type V1BetaCustomerAssignmentListResponse struct {
	// Synthetic UUID identifier — also the cursor anchor for paginated lists
	ID string `json:"id" api:"required" format:"uuid"`
	// Usage-reset cadence. Currently only `MONTH` is supported
	//
	// Any of "MONTH".
	Cadence V1BetaCustomerAssignmentListResponseCadence `json:"cadence" api:"required"`
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
func (r V1BetaCustomerAssignmentListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerAssignmentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage-reset cadence. Currently only `MONTH` is supported
type V1BetaCustomerAssignmentListResponseCadence string

const (
	V1BetaCustomerAssignmentListResponseCadenceMonth V1BetaCustomerAssignmentListResponseCadence = "MONTH"
)

// Assignments after upsert.
type V1BetaCustomerAssignmentUpsertResponse struct {
	Data []V1BetaCustomerAssignmentUpsertResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BetaCustomerAssignmentUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerAssignmentUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A capability assignment for an entity belonging to a customer. Defines how much
// of the capability the entity may consume (`usageLimit`) and how often the
// counter resets (`cadence`).
type V1BetaCustomerAssignmentUpsertResponseData struct {
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
func (r V1BetaCustomerAssignmentUpsertResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BetaCustomerAssignmentUpsertResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BetaCustomerAssignmentListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Filter assignments to a specific capability refId
	CapabilityID param.Opt[string] `query:"capabilityId,omitzero" json:"-"`
	// Filter assignments to a specific entity refId
	EntityID param.Opt[string] `query:"entityId,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BetaCustomerAssignmentListParams]'s query parameters as
// `url.Values`.
func (r V1BetaCustomerAssignmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1BetaCustomerAssignmentUpsertParams struct {
	// Assignments to upsert (1–100 per request)
	Assignments []V1BetaCustomerAssignmentUpsertParamsAssignment `json:"assignments,omitzero" api:"required"`
	paramObj
}

func (r V1BetaCustomerAssignmentUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaCustomerAssignmentUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaCustomerAssignmentUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single assignment to create or update. The natural key is the
// `(entityId, capabilityId)` pair. On create both `usageLimit` and `cadence` are
// required; on update they may be omitted individually to preserve the existing
// value.
//
// The properties CapabilityID, EntityID are required.
type V1BetaCustomerAssignmentUpsertParamsAssignment struct {
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

func (r V1BetaCustomerAssignmentUpsertParamsAssignment) MarshalJSON() (data []byte, err error) {
	type shadow V1BetaCustomerAssignmentUpsertParamsAssignment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BetaCustomerAssignmentUpsertParamsAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BetaCustomerAssignmentUpsertParamsAssignment](
		"cadence", "MONTH",
	)
}
