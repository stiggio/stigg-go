// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stiggio/stigg-go/internal/apijson"
	"github.com/stiggio/stigg-go/internal/requestconfig"
	"github.com/stiggio/stigg-go/option"
	"github.com/stiggio/stigg-go/packages/respjson"
)

// V1EventAddonDraftService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventAddonDraftService] method instead.
type V1EventAddonDraftService struct {
	Options []option.RequestOption
}

// NewV1EventAddonDraftService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1EventAddonDraftService(opts ...option.RequestOption) (r V1EventAddonDraftService) {
	r = V1EventAddonDraftService{}
	r.Options = opts
	return
}

// Creates a draft version of an existing addon for modification before publishing.
func (r *V1EventAddonDraftService) NewAddonDraft(ctx context.Context, id string, opts ...option.RequestOption) (res *V1EventAddonDraftNewAddonDraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s/draft", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Removes a draft version of an addon.
func (r *V1EventAddonDraftService) RemoveAddonDraft(ctx context.Context, id string, opts ...option.RequestOption) (res *V1EventAddonDraftRemoveAddonDraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s/draft", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Response object
type V1EventAddonDraftNewAddonDraftResponse struct {
	// Addon configuration object
	Data V1EventAddonDraftNewAddonDraftResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonDraftNewAddonDraftResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonDraftNewAddonDraftResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Addon configuration object
type V1EventAddonDraftNewAddonDraftResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id,required"`
	// The unique identifier for the entity in the billing provider
	BillingID string `json:"billingId,required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// List of addons the addon is dependant on
	Dependencies []string `json:"dependencies,required"`
	// The description of the package
	Description string `json:"description,required"`
	// The display name of the package
	DisplayName string `json:"displayName,required"`
	// List of entitlements for the addon
	Entitlements []V1EventAddonDraftNewAddonDraftResponseDataEntitlement `json:"entitlements,required"`
	// Indicates if the package is the latest version
	IsLatest bool `json:"isLatest,required"`
	// The maximum quantity of this addon that can be added to a subscription
	MaxQuantity int64 `json:"maxQuantity,required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,required"`
	// The pricing type of the package
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType string `json:"pricingType,required"`
	// The status of the package
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status string `json:"status,required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The version number of the package
	VersionNumber int64 `json:"versionNumber,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BillingID     respjson.Field
		CreatedAt     respjson.Field
		Dependencies  respjson.Field
		Description   respjson.Field
		DisplayName   respjson.Field
		Entitlements  respjson.Field
		IsLatest      respjson.Field
		MaxQuantity   respjson.Field
		Metadata      respjson.Field
		PricingType   respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		VersionNumber respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonDraftNewAddonDraftResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonDraftNewAddonDraftResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type V1EventAddonDraftNewAddonDraftResponseDataEntitlement struct {
	// The unique identifier for the entity
	ID string `json:"id,required"`
	// Any of "FEATURE", "CREDIT".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonDraftNewAddonDraftResponseDataEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonDraftNewAddonDraftResponseDataEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response confirming the addon draft was removed.
type V1EventAddonDraftRemoveAddonDraftResponse struct {
	Data V1EventAddonDraftRemoveAddonDraftResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonDraftRemoveAddonDraftResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonDraftRemoveAddonDraftResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventAddonDraftRemoveAddonDraftResponseData struct {
	// The unique identifier for the entity
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonDraftRemoveAddonDraftResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonDraftRemoveAddonDraftResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
