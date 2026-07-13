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
	"github.com/stiggio/stigg-go/packages/param"
	"github.com/stiggio/stigg-go/packages/respjson"
)

// V1EventDataExportDestinationService contains methods and other services that
// help with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventDataExportDestinationService] method instead.
type V1EventDataExportDestinationService struct {
	Options []option.RequestOption
}

// NewV1EventDataExportDestinationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1EventDataExportDestinationService(opts ...option.RequestOption) (r V1EventDataExportDestinationService) {
	r = V1EventDataExportDestinationService{}
	r.Options = opts
	return
}

// Register a destination on the environment's DATA_EXPORT integration.
// Lazy-creates the integration row + provider recipient on first call. Idempotent
// on destinationId.
func (r *V1EventDataExportDestinationService) New(ctx context.Context, params V1EventDataExportDestinationNewParams, opts ...option.RequestOption) (res *V1EventDataExportDestinationNewResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/data-export/destinations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a destination's entity selection. Pushes the new enabled_models to the
// provider first, then persists the selection. Applies on the next scheduled
// transfer.
func (r *V1EventDataExportDestinationService) Update(ctx context.Context, destinationID string, params V1EventDataExportDestinationUpdateParams, opts ...option.RequestOption) (res *V1EventDataExportDestinationUpdateResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if destinationID == "" {
		err = errors.New("missing required destinationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/data-export/destinations/%s", destinationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Disconnect a destination: stops the provider sync (deletes the provider
// destination) and removes it from the DATA_EXPORT integration. Non-destructive —
// the warehouse table is left intact. Idempotent.
func (r *V1EventDataExportDestinationService) Delete(ctx context.Context, destinationID string, body V1EventDataExportDestinationDeleteParams, opts ...option.RequestOption) (res *V1EventDataExportDestinationDeleteResponse, err error) {
	if !param.IsOmitted(body.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", body.XAccountID.Value)))
	}
	if !param.IsOmitted(body.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", body.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if destinationID == "" {
		err = errors.New("missing required destinationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/data-export/destinations/%s", destinationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Response object
type V1EventDataExportDestinationNewResponse struct {
	// Current destinations under the DATA_EXPORT integration.
	Data V1EventDataExportDestinationNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportDestinationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current destinations under the DATA_EXPORT integration.
type V1EventDataExportDestinationNewResponseData struct {
	// Current destinations under the DATA_EXPORT integration
	Destinations []V1EventDataExportDestinationNewResponseDataDestination `json:"destinations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Destinations respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportDestinationNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single destination entry under the DATA_EXPORT integration.
type V1EventDataExportDestinationNewResponseDataDestination struct {
	// ISO8601 timestamp of when the destination was connected
	ConnectedAt string `json:"connectedAt" api:"required"`
	// Provider destination ID
	DestinationID string `json:"destinationId" api:"required"`
	// Destination type (snowflake, bigquery, ...)
	Type string `json:"type" api:"required"`
	// Connection status of the destination (connected, failed)
	ConnectionStatus string   `json:"connectionStatus"`
	EnabledModels    []string `json:"enabledModels"`
	// Latest sync snapshot for the destination, refreshed by the provider webhook
	LastSyncStatus V1EventDataExportDestinationNewResponseDataDestinationLastSyncStatus `json:"lastSyncStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectedAt      respjson.Field
		DestinationID    respjson.Field
		Type             respjson.Field
		ConnectionStatus respjson.Field
		EnabledModels    respjson.Field
		LastSyncStatus   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationNewResponseDataDestination) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportDestinationNewResponseDataDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Latest sync snapshot for the destination, refreshed by the provider webhook
type V1EventDataExportDestinationNewResponseDataDestinationLastSyncStatus struct {
	// ISO8601 timestamp of when the latest sync finished
	FinishedAt string `json:"finishedAt" api:"required"`
	// Sync status (PENDING, RUNNING, INCOMPLETE, FAILED, SUCCEEDED, CANCELLED)
	Status string `json:"status" api:"required"`
	// Provider transfer ID of the latest sync
	TransferID string `json:"transferId" api:"required"`
	// Party responsible for a failed sync, as reported by the data-export provider
	BlamedParty string `json:"blamedParty"`
	// Customer-friendly failure message, when the latest sync failed
	FailureMessage string `json:"failureMessage"`
	// Number of rows transferred in the latest sync
	RowsTransferred float64 `json:"rowsTransferred"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishedAt      respjson.Field
		Status          respjson.Field
		TransferID      respjson.Field
		BlamedParty     respjson.Field
		FailureMessage  respjson.Field
		RowsTransferred respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationNewResponseDataDestinationLastSyncStatus) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventDataExportDestinationNewResponseDataDestinationLastSyncStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1EventDataExportDestinationUpdateResponse struct {
	// Current destinations under the DATA_EXPORT integration.
	Data V1EventDataExportDestinationUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportDestinationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current destinations under the DATA_EXPORT integration.
type V1EventDataExportDestinationUpdateResponseData struct {
	// Current destinations under the DATA_EXPORT integration
	Destinations []V1EventDataExportDestinationUpdateResponseDataDestination `json:"destinations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Destinations respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportDestinationUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single destination entry under the DATA_EXPORT integration.
type V1EventDataExportDestinationUpdateResponseDataDestination struct {
	// ISO8601 timestamp of when the destination was connected
	ConnectedAt string `json:"connectedAt" api:"required"`
	// Provider destination ID
	DestinationID string `json:"destinationId" api:"required"`
	// Destination type (snowflake, bigquery, ...)
	Type string `json:"type" api:"required"`
	// Connection status of the destination (connected, failed)
	ConnectionStatus string   `json:"connectionStatus"`
	EnabledModels    []string `json:"enabledModels"`
	// Latest sync snapshot for the destination, refreshed by the provider webhook
	LastSyncStatus V1EventDataExportDestinationUpdateResponseDataDestinationLastSyncStatus `json:"lastSyncStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectedAt      respjson.Field
		DestinationID    respjson.Field
		Type             respjson.Field
		ConnectionStatus respjson.Field
		EnabledModels    respjson.Field
		LastSyncStatus   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationUpdateResponseDataDestination) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventDataExportDestinationUpdateResponseDataDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Latest sync snapshot for the destination, refreshed by the provider webhook
type V1EventDataExportDestinationUpdateResponseDataDestinationLastSyncStatus struct {
	// ISO8601 timestamp of when the latest sync finished
	FinishedAt string `json:"finishedAt" api:"required"`
	// Sync status (PENDING, RUNNING, INCOMPLETE, FAILED, SUCCEEDED, CANCELLED)
	Status string `json:"status" api:"required"`
	// Provider transfer ID of the latest sync
	TransferID string `json:"transferId" api:"required"`
	// Party responsible for a failed sync, as reported by the data-export provider
	BlamedParty string `json:"blamedParty"`
	// Customer-friendly failure message, when the latest sync failed
	FailureMessage string `json:"failureMessage"`
	// Number of rows transferred in the latest sync
	RowsTransferred float64 `json:"rowsTransferred"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishedAt      respjson.Field
		Status          respjson.Field
		TransferID      respjson.Field
		BlamedParty     respjson.Field
		FailureMessage  respjson.Field
		RowsTransferred respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationUpdateResponseDataDestinationLastSyncStatus) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventDataExportDestinationUpdateResponseDataDestinationLastSyncStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1EventDataExportDestinationDeleteResponse struct {
	// Current destinations under the DATA_EXPORT integration.
	Data V1EventDataExportDestinationDeleteResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportDestinationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current destinations under the DATA_EXPORT integration.
type V1EventDataExportDestinationDeleteResponseData struct {
	// Current destinations under the DATA_EXPORT integration
	Destinations []V1EventDataExportDestinationDeleteResponseDataDestination `json:"destinations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Destinations respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportDestinationDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single destination entry under the DATA_EXPORT integration.
type V1EventDataExportDestinationDeleteResponseDataDestination struct {
	// ISO8601 timestamp of when the destination was connected
	ConnectedAt string `json:"connectedAt" api:"required"`
	// Provider destination ID
	DestinationID string `json:"destinationId" api:"required"`
	// Destination type (snowflake, bigquery, ...)
	Type string `json:"type" api:"required"`
	// Connection status of the destination (connected, failed)
	ConnectionStatus string   `json:"connectionStatus"`
	EnabledModels    []string `json:"enabledModels"`
	// Latest sync snapshot for the destination, refreshed by the provider webhook
	LastSyncStatus V1EventDataExportDestinationDeleteResponseDataDestinationLastSyncStatus `json:"lastSyncStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectedAt      respjson.Field
		DestinationID    respjson.Field
		Type             respjson.Field
		ConnectionStatus respjson.Field
		EnabledModels    respjson.Field
		LastSyncStatus   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationDeleteResponseDataDestination) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventDataExportDestinationDeleteResponseDataDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Latest sync snapshot for the destination, refreshed by the provider webhook
type V1EventDataExportDestinationDeleteResponseDataDestinationLastSyncStatus struct {
	// ISO8601 timestamp of when the latest sync finished
	FinishedAt string `json:"finishedAt" api:"required"`
	// Sync status (PENDING, RUNNING, INCOMPLETE, FAILED, SUCCEEDED, CANCELLED)
	Status string `json:"status" api:"required"`
	// Provider transfer ID of the latest sync
	TransferID string `json:"transferId" api:"required"`
	// Party responsible for a failed sync, as reported by the data-export provider
	BlamedParty string `json:"blamedParty"`
	// Customer-friendly failure message, when the latest sync failed
	FailureMessage string `json:"failureMessage"`
	// Number of rows transferred in the latest sync
	RowsTransferred float64 `json:"rowsTransferred"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishedAt      respjson.Field
		Status          respjson.Field
		TransferID      respjson.Field
		BlamedParty     respjson.Field
		FailureMessage  respjson.Field
		RowsTransferred respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportDestinationDeleteResponseDataDestinationLastSyncStatus) RawJSON() string {
	return r.JSON.raw
}
func (r *V1EventDataExportDestinationDeleteResponseDataDestinationLastSyncStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventDataExportDestinationNewParams struct {
	// The provider destination ID returned by the embedded SDK on connect
	DestinationID string `json:"destinationId" api:"required"`
	// The destination type (e.g. snowflake, bigquery)
	DestinationType string            `json:"destinationType" api:"required"`
	XAccountID      param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID  param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	EnabledModels   []string          `json:"enabledModels,omitzero"`
	paramObj
}

func (r V1EventDataExportDestinationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventDataExportDestinationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventDataExportDestinationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventDataExportDestinationUpdateParams struct {
	EnabledModels []string `json:"enabledModels,omitzero" api:"required"`
	// Target integration row hosting the destination
	IntegrationID  string            `json:"integrationId" api:"required"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1EventDataExportDestinationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventDataExportDestinationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventDataExportDestinationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventDataExportDestinationDeleteParams struct {
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}
