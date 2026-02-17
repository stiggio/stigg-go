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

// V1EventAddonService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventAddonService] method instead.
type V1EventAddonService struct {
	Options []option.RequestOption
	Draft   V1EventAddonDraftService
}

// NewV1EventAddonService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1EventAddonService(opts ...option.RequestOption) (r V1EventAddonService) {
	r = V1EventAddonService{}
	r.Options = opts
	r.Draft = NewV1EventAddonDraftService(opts...)
	return
}

// Archives an addon, preventing it from being used in new subscriptions.
func (r *V1EventAddonService) ArchiveAddon(ctx context.Context, id string, opts ...option.RequestOption) (res *V1EventAddonArchiveAddonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Creates a new addon in draft status, associated with a specific product.
func (r *V1EventAddonService) NewAddon(ctx context.Context, body V1EventAddonNewAddonParams, opts ...option.RequestOption) (res *V1EventAddonNewAddonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/addons"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of addons in the environment.
func (r *V1EventAddonService) ListAddons(ctx context.Context, query V1EventAddonListAddonsParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1EventAddonListAddonsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/addons"
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

// Retrieves a paginated list of addons in the environment.
func (r *V1EventAddonService) ListAddonsAutoPaging(ctx context.Context, query V1EventAddonListAddonsParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1EventAddonListAddonsResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListAddons(ctx, query, opts...))
}

// Publishes a draft addon, making it available for use in subscriptions.
func (r *V1EventAddonService) PublishAddon(ctx context.Context, id string, body V1EventAddonPublishAddonParams, opts ...option.RequestOption) (res *V1EventAddonPublishAddonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s/publish", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves an addon by its unique identifier, including entitlements and pricing
// details.
func (r *V1EventAddonService) GetAddon(ctx context.Context, id string, opts ...option.RequestOption) (res *V1EventAddonGetAddonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing addon's properties such as display name, description, and
// metadata.
func (r *V1EventAddonService) UpdateAddon(ctx context.Context, id string, body V1EventAddonUpdateAddonParams, opts ...option.RequestOption) (res *V1EventAddonUpdateAddonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/addons/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Response object
type V1EventAddonArchiveAddonResponse struct {
	// Addon configuration object
	Data V1EventAddonArchiveAddonResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonArchiveAddonResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonArchiveAddonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Addon configuration object
type V1EventAddonArchiveAddonResponseData struct {
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
	Entitlements []V1EventAddonArchiveAddonResponseDataEntitlement `json:"entitlements,required"`
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
func (r V1EventAddonArchiveAddonResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonArchiveAddonResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type V1EventAddonArchiveAddonResponseDataEntitlement struct {
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
func (r V1EventAddonArchiveAddonResponseDataEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonArchiveAddonResponseDataEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1EventAddonNewAddonResponse struct {
	// Addon configuration object
	Data V1EventAddonNewAddonResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonNewAddonResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonNewAddonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Addon configuration object
type V1EventAddonNewAddonResponseData struct {
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
	Entitlements []V1EventAddonNewAddonResponseDataEntitlement `json:"entitlements,required"`
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
func (r V1EventAddonNewAddonResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonNewAddonResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type V1EventAddonNewAddonResponseDataEntitlement struct {
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
func (r V1EventAddonNewAddonResponseDataEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonNewAddonResponseDataEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Addon configuration object
type V1EventAddonListAddonsResponse struct {
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
	Entitlements []V1EventAddonListAddonsResponseEntitlement `json:"entitlements,required"`
	// Indicates if the package is the latest version
	IsLatest bool `json:"isLatest,required"`
	// The maximum quantity of this addon that can be added to a subscription
	MaxQuantity int64 `json:"maxQuantity,required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,required"`
	// The pricing type of the package
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType V1EventAddonListAddonsResponsePricingType `json:"pricingType,required"`
	// The status of the package
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status V1EventAddonListAddonsResponseStatus `json:"status,required"`
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
func (r V1EventAddonListAddonsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonListAddonsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type V1EventAddonListAddonsResponseEntitlement struct {
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
func (r V1EventAddonListAddonsResponseEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonListAddonsResponseEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pricing type of the package
type V1EventAddonListAddonsResponsePricingType string

const (
	V1EventAddonListAddonsResponsePricingTypeFree   V1EventAddonListAddonsResponsePricingType = "FREE"
	V1EventAddonListAddonsResponsePricingTypePaid   V1EventAddonListAddonsResponsePricingType = "PAID"
	V1EventAddonListAddonsResponsePricingTypeCustom V1EventAddonListAddonsResponsePricingType = "CUSTOM"
)

// The status of the package
type V1EventAddonListAddonsResponseStatus string

const (
	V1EventAddonListAddonsResponseStatusDraft     V1EventAddonListAddonsResponseStatus = "DRAFT"
	V1EventAddonListAddonsResponseStatusPublished V1EventAddonListAddonsResponseStatus = "PUBLISHED"
	V1EventAddonListAddonsResponseStatusArchived  V1EventAddonListAddonsResponseStatus = "ARCHIVED"
)

// Response containing task ID for publish operation
type V1EventAddonPublishAddonResponse struct {
	Data V1EventAddonPublishAddonResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonPublishAddonResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonPublishAddonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventAddonPublishAddonResponseData struct {
	// Task ID for tracking the async publish operation
	TaskID string `json:"taskId,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TaskID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonPublishAddonResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonPublishAddonResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1EventAddonGetAddonResponse struct {
	// Addon configuration object
	Data V1EventAddonGetAddonResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonGetAddonResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonGetAddonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Addon configuration object
type V1EventAddonGetAddonResponseData struct {
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
	Entitlements []V1EventAddonGetAddonResponseDataEntitlement `json:"entitlements,required"`
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
func (r V1EventAddonGetAddonResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonGetAddonResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type V1EventAddonGetAddonResponseDataEntitlement struct {
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
func (r V1EventAddonGetAddonResponseDataEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonGetAddonResponseDataEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1EventAddonUpdateAddonResponse struct {
	// Addon configuration object
	Data V1EventAddonUpdateAddonResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventAddonUpdateAddonResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonUpdateAddonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Addon configuration object
type V1EventAddonUpdateAddonResponseData struct {
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
	Entitlements []V1EventAddonUpdateAddonResponseDataEntitlement `json:"entitlements,required"`
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
func (r V1EventAddonUpdateAddonResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonUpdateAddonResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entitlement reference with type and identifier
type V1EventAddonUpdateAddonResponseDataEntitlement struct {
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
func (r V1EventAddonUpdateAddonResponseDataEntitlement) RawJSON() string { return r.JSON.raw }
func (r *V1EventAddonUpdateAddonResponseDataEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventAddonNewAddonParams struct {
	// The unique identifier for the entity
	ID string `json:"id,required"`
	// The display name of the package
	DisplayName string `json:"displayName,required"`
	// The product ID to associate the addon with
	ProductID string `json:"productId,required"`
	// The unique identifier for the entity in the billing provider
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// The description of the package
	Description param.Opt[string] `json:"description,omitzero"`
	// The maximum quantity of this addon that can be added to a subscription
	MaxQuantity param.Opt[int64] `json:"maxQuantity,omitzero"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,omitzero"`
	// The pricing type of the package
	//
	// Any of "FREE", "PAID", "CUSTOM".
	PricingType V1EventAddonNewAddonParamsPricingType `json:"pricingType,omitzero"`
	// The status of the package
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status V1EventAddonNewAddonParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r V1EventAddonNewAddonParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonNewAddonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonNewAddonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pricing type of the package
type V1EventAddonNewAddonParamsPricingType string

const (
	V1EventAddonNewAddonParamsPricingTypeFree   V1EventAddonNewAddonParamsPricingType = "FREE"
	V1EventAddonNewAddonParamsPricingTypePaid   V1EventAddonNewAddonParamsPricingType = "PAID"
	V1EventAddonNewAddonParamsPricingTypeCustom V1EventAddonNewAddonParamsPricingType = "CUSTOM"
)

// The status of the package
type V1EventAddonNewAddonParamsStatus string

const (
	V1EventAddonNewAddonParamsStatusDraft     V1EventAddonNewAddonParamsStatus = "DRAFT"
	V1EventAddonNewAddonParamsStatusPublished V1EventAddonNewAddonParamsStatus = "PUBLISHED"
	V1EventAddonNewAddonParamsStatusArchived  V1EventAddonNewAddonParamsStatus = "ARCHIVED"
)

type V1EventAddonListAddonsParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by product ID
	ProductID param.Opt[string] `query:"productId,omitzero" json:"-"`
	// Filter by creation date using range operators: gt, gte, lt, lte
	CreatedAt V1EventAddonListAddonsParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	// Filter by addon status. Supports comma-separated values for multiple statuses
	//
	// Any of "DRAFT", "PUBLISHED", "ARCHIVED".
	Status V1EventAddonListAddonsParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventAddonListAddonsParams]'s query parameters as
// `url.Values`.
func (r V1EventAddonListAddonsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by creation date using range operators: gt, gte, lt, lte
type V1EventAddonListAddonsParamsCreatedAt struct {
	// Greater than the specified createdAt value
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// Greater than or equal to the specified createdAt value
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Less than the specified createdAt value
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	// Less than or equal to the specified createdAt value
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventAddonListAddonsParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r V1EventAddonListAddonsParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by addon status. Supports comma-separated values for multiple statuses
type V1EventAddonListAddonsParamsStatus string

const (
	V1EventAddonListAddonsParamsStatusDraft     V1EventAddonListAddonsParamsStatus = "DRAFT"
	V1EventAddonListAddonsParamsStatusPublished V1EventAddonListAddonsParamsStatus = "PUBLISHED"
	V1EventAddonListAddonsParamsStatusArchived  V1EventAddonListAddonsParamsStatus = "ARCHIVED"
)

type V1EventAddonPublishAddonParams struct {
	// The migration type of the package
	//
	// Any of "NEW_CUSTOMERS", "ALL_CUSTOMERS".
	MigrationType V1EventAddonPublishAddonParamsMigrationType `json:"migrationType,omitzero,required"`
	paramObj
}

func (r V1EventAddonPublishAddonParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonPublishAddonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonPublishAddonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The migration type of the package
type V1EventAddonPublishAddonParamsMigrationType string

const (
	V1EventAddonPublishAddonParamsMigrationTypeNewCustomers V1EventAddonPublishAddonParamsMigrationType = "NEW_CUSTOMERS"
	V1EventAddonPublishAddonParamsMigrationTypeAllCustomers V1EventAddonPublishAddonParamsMigrationType = "ALL_CUSTOMERS"
)

type V1EventAddonUpdateAddonParams struct {
	// The unique identifier for the entity in the billing provider
	BillingID param.Opt[string] `json:"billingId,omitzero"`
	// The description of the package
	Description param.Opt[string] `json:"description,omitzero"`
	// The maximum quantity of this addon that can be added to a subscription
	MaxQuantity param.Opt[int64] `json:"maxQuantity,omitzero"`
	// The display name of the package
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// List of addons the addon is dependant on
	Dependencies []string `json:"dependencies,omitzero"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r V1EventAddonUpdateAddonParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventAddonUpdateAddonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventAddonUpdateAddonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
