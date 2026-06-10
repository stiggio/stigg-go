// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/stiggio/stigg-go/internal/apijson"
	"github.com/stiggio/stigg-go/internal/requestconfig"
	"github.com/stiggio/stigg-go/option"
	"github.com/stiggio/stigg-go/packages/param"
	"github.com/stiggio/stigg-go/packages/respjson"
)

// V1EventDataExportService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventDataExportService] method instead.
type V1EventDataExportService struct {
	Options      []option.RequestOption
	Destinations V1EventDataExportDestinationService
}

// NewV1EventDataExportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1EventDataExportService(opts ...option.RequestOption) (r V1EventDataExportService) {
	r = V1EventDataExportService{}
	r.Options = opts
	r.Destinations = NewV1EventDataExportDestinationService(opts...)
	return
}

// Mint a scoped JWT for the FE embedded SDK. Lazy-creates the DATA_EXPORT
// integration if needed.
func (r *V1EventDataExportService) MintScopedToken(ctx context.Context, params V1EventDataExportMintScopedTokenParams, opts ...option.RequestOption) (res *V1EventDataExportMintScopedTokenResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/data-export/scoped-token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Trigger a sync for one destination or all destinations under the provider
// entity.
func (r *V1EventDataExportService) TriggerSync(ctx context.Context, params V1EventDataExportTriggerSyncParams, opts ...option.RequestOption) (res *V1EventDataExportTriggerSyncResponse, err error) {
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("X-ACCOUNT-ID", fmt.Sprintf("%v", params.XAccountID.Value)))
	}
	if !param.IsOmitted(params.XEnvironmentID) {
		opts = append(opts, option.WithHeader("X-ENVIRONMENT-ID", fmt.Sprintf("%v", params.XEnvironmentID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/data-export/sync"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Response object
type V1EventDataExportMintScopedTokenResponse struct {
	// Scoped token + expiry + provider-specific metadata for the FE SDK.
	Data V1EventDataExportMintScopedTokenResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportMintScopedTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportMintScopedTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scoped token + expiry + provider-specific metadata for the FE SDK.
type V1EventDataExportMintScopedTokenResponseData struct {
	// Provider scoped JWT
	Token string `json:"token" api:"required"`
	// ISO8601 token expiry
	ExpiresAt string `json:"expiresAt" api:"required"`
	// Provider-specific extras the FE embedded SDK needs
	ProviderMetadata map[string]any `json:"providerMetadata" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token            respjson.Field
		ExpiresAt        respjson.Field
		ProviderMetadata respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportMintScopedTokenResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportMintScopedTokenResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1EventDataExportTriggerSyncResponse struct {
	// Per-destination trigger results across the batch.
	Data V1EventDataExportTriggerSyncResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportTriggerSyncResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportTriggerSyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-destination trigger results across the batch.
type V1EventDataExportTriggerSyncResponseData struct {
	// Per-destination trigger results
	Results []V1EventDataExportTriggerSyncResponseDataResult `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportTriggerSyncResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportTriggerSyncResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-destination trigger results.
type V1EventDataExportTriggerSyncResponseDataResult struct {
	// Provider destination ID
	DestinationID string `json:"destinationId" api:"required"`
	// True if a transfer was kicked
	Triggered bool `json:"triggered" api:"required"`
	// Error message if triggered=false on a hard failure
	ErrorMessage string `json:"errorMessage"`
	// Provider-side transfer ID
	TransferID string `json:"transferId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationID respjson.Field
		Triggered     respjson.Field
		ErrorMessage  respjson.Field
		TransferID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventDataExportTriggerSyncResponseDataResult) RawJSON() string { return r.JSON.raw }
func (r *V1EventDataExportTriggerSyncResponseDataResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventDataExportMintScopedTokenParams struct {
	// FE origin the resulting JWT is bound to (provider-side anti-fraud)
	ApplicationOrigin string `json:"applicationOrigin" api:"required"`
	// Pin the token to a specific warehouse connect flow
	DestinationType param.Opt[string] `json:"destinationType,omitzero"`
	XAccountID      param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID  param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1EventDataExportMintScopedTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventDataExportMintScopedTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventDataExportMintScopedTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventDataExportTriggerSyncParams struct {
	// Provider destination ID to sync. Omit to sync all destinations.
	DestinationID  param.Opt[string] `json:"destinationId,omitzero"`
	XAccountID     param.Opt[string] `header:"X-ACCOUNT-ID,omitzero" json:"-"`
	XEnvironmentID param.Opt[string] `header:"X-ENVIRONMENT-ID,omitzero" json:"-"`
	paramObj
}

func (r V1EventDataExportTriggerSyncParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventDataExportTriggerSyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventDataExportTriggerSyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
