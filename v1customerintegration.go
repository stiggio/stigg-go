// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stiggio/stigg-go/internal/apijson"
	"github.com/stiggio/stigg-go/internal/apiquery"
	"github.com/stiggio/stigg-go/internal/requestconfig"
	"github.com/stiggio/stigg-go/option"
	"github.com/stiggio/stigg-go/packages/pagination"
	"github.com/stiggio/stigg-go/packages/param"
	"github.com/stiggio/stigg-go/packages/respjson"
)

// V1CustomerIntegrationService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerIntegrationService] method instead.
type V1CustomerIntegrationService struct {
	Options []option.RequestOption
}

// NewV1CustomerIntegrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerIntegrationService(opts ...option.RequestOption) (r V1CustomerIntegrationService) {
	r = V1CustomerIntegrationService{}
	r.Options = opts
	return
}

// Retrieves a specific integration for a customer by integration ID.
func (r *V1CustomerIntegrationService) Get(ctx context.Context, integrationID string, query V1CustomerIntegrationGetParams, opts ...option.RequestOption) (res *V1CustomerIntegrationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if integrationID == "" {
		err = errors.New("missing required integrationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/integrations/%s", query.ID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a customer's integration link, such as changing the synced external
// entity ID.
func (r *V1CustomerIntegrationService) Update(ctx context.Context, integrationID string, params V1CustomerIntegrationUpdateParams, opts ...option.RequestOption) (res *V1CustomerIntegrationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if integrationID == "" {
		err = errors.New("missing required integrationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/integrations/%s", params.ID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieves a paginated list of a customer's external integrations (billing, CRM,
// etc.).
func (r *V1CustomerIntegrationService) List(ctx context.Context, id string, query V1CustomerIntegrationListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1CustomerIntegrationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/integrations", id)
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

// Retrieves a paginated list of a customer's external integrations (billing, CRM,
// etc.).
func (r *V1CustomerIntegrationService) ListAutoPaging(ctx context.Context, id string, query V1CustomerIntegrationListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1CustomerIntegrationListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, id, query, opts...))
}

// Links a customer to an external integration by specifying the vendor and
// external entity ID.
func (r *V1CustomerIntegrationService) Link(ctx context.Context, id string, body V1CustomerIntegrationLinkParams, opts ...option.RequestOption) (res *V1CustomerIntegrationLinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/integrations", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Removes the link between a customer and an external integration.
func (r *V1CustomerIntegrationService) Unlink(ctx context.Context, integrationID string, body V1CustomerIntegrationUnlinkParams, opts ...option.RequestOption) (res *V1CustomerIntegrationUnlinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if integrationID == "" {
		err = errors.New("missing required integrationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/customers/%s/integrations/%s", body.ID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Response object
type V1CustomerIntegrationGetResponse struct {
	// External billing or CRM integration link
	Data V1CustomerIntegrationGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerIntegrationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External billing or CRM integration link
type V1CustomerIntegrationGetResponseData struct {
	// Integration details
	ID string `json:"id" api:"required"`
	// Synced entity id
	SyncedEntityID string `json:"syncedEntityId" api:"required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier" api:"required"`
	// Price billing sync revision data containing billing ID, link URL, and price
	// group package billing ID
	SyncData V1CustomerIntegrationGetResponseDataSyncDataUnion `json:"syncData" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		SyncedEntityID   respjson.Field
		VendorIdentifier respjson.Field
		SyncData         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerIntegrationGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1CustomerIntegrationGetResponseDataSyncDataUnion contains all possible
// properties and values from
// [V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionPriceBillingData],
// [V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionBillingData],
// [V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionMarketplaceData].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerIntegrationGetResponseDataSyncDataUnion struct {
	BillingID      string `json:"billingId"`
	BillingLinkURL string `json:"billingLinkUrl"`
	// This field is from variant
	// [V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionPriceBillingData].
	PriceGroupPackageBillingID string `json:"priceGroupPackageBillingId"`
	// This field is from variant
	// [V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionMarketplaceData].
	Dimensions string `json:"dimensions"`
	JSON       struct {
		BillingID                  respjson.Field
		BillingLinkURL             respjson.Field
		PriceGroupPackageBillingID respjson.Field
		Dimensions                 respjson.Field
		raw                        string
	} `json:"-"`
}

func (u V1CustomerIntegrationGetResponseDataSyncDataUnion) AsSyncRevisionPriceBillingData() (v V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionPriceBillingData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerIntegrationGetResponseDataSyncDataUnion) AsSyncRevisionBillingData() (v V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionBillingData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerIntegrationGetResponseDataSyncDataUnion) AsSyncRevisionMarketplaceData() (v V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionMarketplaceData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerIntegrationGetResponseDataSyncDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1CustomerIntegrationGetResponseDataSyncDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price billing sync revision data containing billing ID, link URL, and price
// group package billing ID
type V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionPriceBillingData struct {
	// Billing integration id
	BillingID string `json:"billingId" api:"required"`
	// Billing integration url
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// Price group package billing id
	PriceGroupPackageBillingID string `json:"priceGroupPackageBillingId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID                  respjson.Field
		BillingLinkURL             respjson.Field
		PriceGroupPackageBillingID respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionPriceBillingData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionPriceBillingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing sync revision data containing billing ID and link URL
type V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionBillingData struct {
	// Billing integration id
	BillingID string `json:"billingId" api:"required"`
	// Billing integration url
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		BillingLinkURL respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionBillingData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionBillingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Marketplace sync revision data containing dimensions
type V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionMarketplaceData struct {
	// Dimensions of the marketplace sync revision
	Dimensions string `json:"dimensions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimensions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionMarketplaceData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationGetResponseDataSyncDataSyncRevisionMarketplaceData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1CustomerIntegrationUpdateResponse struct {
	// External billing or CRM integration link
	Data V1CustomerIntegrationUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerIntegrationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External billing or CRM integration link
type V1CustomerIntegrationUpdateResponseData struct {
	// Integration details
	ID string `json:"id" api:"required"`
	// Synced entity id
	SyncedEntityID string `json:"syncedEntityId" api:"required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier" api:"required"`
	// Price billing sync revision data containing billing ID, link URL, and price
	// group package billing ID
	SyncData V1CustomerIntegrationUpdateResponseDataSyncDataUnion `json:"syncData" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		SyncedEntityID   respjson.Field
		VendorIdentifier respjson.Field
		SyncData         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerIntegrationUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1CustomerIntegrationUpdateResponseDataSyncDataUnion contains all possible
// properties and values from
// [V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionPriceBillingData],
// [V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionBillingData],
// [V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionMarketplaceData].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerIntegrationUpdateResponseDataSyncDataUnion struct {
	BillingID      string `json:"billingId"`
	BillingLinkURL string `json:"billingLinkUrl"`
	// This field is from variant
	// [V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionPriceBillingData].
	PriceGroupPackageBillingID string `json:"priceGroupPackageBillingId"`
	// This field is from variant
	// [V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionMarketplaceData].
	Dimensions string `json:"dimensions"`
	JSON       struct {
		BillingID                  respjson.Field
		BillingLinkURL             respjson.Field
		PriceGroupPackageBillingID respjson.Field
		Dimensions                 respjson.Field
		raw                        string
	} `json:"-"`
}

func (u V1CustomerIntegrationUpdateResponseDataSyncDataUnion) AsSyncRevisionPriceBillingData() (v V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionPriceBillingData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerIntegrationUpdateResponseDataSyncDataUnion) AsSyncRevisionBillingData() (v V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionBillingData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerIntegrationUpdateResponseDataSyncDataUnion) AsSyncRevisionMarketplaceData() (v V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionMarketplaceData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerIntegrationUpdateResponseDataSyncDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1CustomerIntegrationUpdateResponseDataSyncDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price billing sync revision data containing billing ID, link URL, and price
// group package billing ID
type V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionPriceBillingData struct {
	// Billing integration id
	BillingID string `json:"billingId" api:"required"`
	// Billing integration url
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// Price group package billing id
	PriceGroupPackageBillingID string `json:"priceGroupPackageBillingId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID                  respjson.Field
		BillingLinkURL             respjson.Field
		PriceGroupPackageBillingID respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionPriceBillingData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionPriceBillingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing sync revision data containing billing ID and link URL
type V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionBillingData struct {
	// Billing integration id
	BillingID string `json:"billingId" api:"required"`
	// Billing integration url
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		BillingLinkURL respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionBillingData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionBillingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Marketplace sync revision data containing dimensions
type V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionMarketplaceData struct {
	// Dimensions of the marketplace sync revision
	Dimensions string `json:"dimensions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimensions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionMarketplaceData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationUpdateResponseDataSyncDataSyncRevisionMarketplaceData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External billing or CRM integration link
type V1CustomerIntegrationListResponse struct {
	// Integration details
	ID string `json:"id" api:"required"`
	// Synced entity id
	SyncedEntityID string `json:"syncedEntityId" api:"required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier V1CustomerIntegrationListResponseVendorIdentifier `json:"vendorIdentifier" api:"required"`
	// Price billing sync revision data containing billing ID, link URL, and price
	// group package billing ID
	SyncData V1CustomerIntegrationListResponseSyncDataUnion `json:"syncData" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		SyncedEntityID   respjson.Field
		VendorIdentifier respjson.Field
		SyncData         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerIntegrationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The vendor identifier of integration
type V1CustomerIntegrationListResponseVendorIdentifier string

const (
	V1CustomerIntegrationListResponseVendorIdentifierAuth0          V1CustomerIntegrationListResponseVendorIdentifier = "AUTH0"
	V1CustomerIntegrationListResponseVendorIdentifierZuora          V1CustomerIntegrationListResponseVendorIdentifier = "ZUORA"
	V1CustomerIntegrationListResponseVendorIdentifierStripe         V1CustomerIntegrationListResponseVendorIdentifier = "STRIPE"
	V1CustomerIntegrationListResponseVendorIdentifierHubspot        V1CustomerIntegrationListResponseVendorIdentifier = "HUBSPOT"
	V1CustomerIntegrationListResponseVendorIdentifierAwsMarketplace V1CustomerIntegrationListResponseVendorIdentifier = "AWS_MARKETPLACE"
	V1CustomerIntegrationListResponseVendorIdentifierSnowflake      V1CustomerIntegrationListResponseVendorIdentifier = "SNOWFLAKE"
	V1CustomerIntegrationListResponseVendorIdentifierSalesforce     V1CustomerIntegrationListResponseVendorIdentifier = "SALESFORCE"
	V1CustomerIntegrationListResponseVendorIdentifierBigQuery       V1CustomerIntegrationListResponseVendorIdentifier = "BIG_QUERY"
	V1CustomerIntegrationListResponseVendorIdentifierOpenFga        V1CustomerIntegrationListResponseVendorIdentifier = "OPEN_FGA"
	V1CustomerIntegrationListResponseVendorIdentifierAppStore       V1CustomerIntegrationListResponseVendorIdentifier = "APP_STORE"
)

// V1CustomerIntegrationListResponseSyncDataUnion contains all possible properties
// and values from
// [V1CustomerIntegrationListResponseSyncDataSyncRevisionPriceBillingData],
// [V1CustomerIntegrationListResponseSyncDataSyncRevisionBillingData],
// [V1CustomerIntegrationListResponseSyncDataSyncRevisionMarketplaceData].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerIntegrationListResponseSyncDataUnion struct {
	BillingID      string `json:"billingId"`
	BillingLinkURL string `json:"billingLinkUrl"`
	// This field is from variant
	// [V1CustomerIntegrationListResponseSyncDataSyncRevisionPriceBillingData].
	PriceGroupPackageBillingID string `json:"priceGroupPackageBillingId"`
	// This field is from variant
	// [V1CustomerIntegrationListResponseSyncDataSyncRevisionMarketplaceData].
	Dimensions string `json:"dimensions"`
	JSON       struct {
		BillingID                  respjson.Field
		BillingLinkURL             respjson.Field
		PriceGroupPackageBillingID respjson.Field
		Dimensions                 respjson.Field
		raw                        string
	} `json:"-"`
}

func (u V1CustomerIntegrationListResponseSyncDataUnion) AsSyncRevisionPriceBillingData() (v V1CustomerIntegrationListResponseSyncDataSyncRevisionPriceBillingData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerIntegrationListResponseSyncDataUnion) AsSyncRevisionBillingData() (v V1CustomerIntegrationListResponseSyncDataSyncRevisionBillingData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerIntegrationListResponseSyncDataUnion) AsSyncRevisionMarketplaceData() (v V1CustomerIntegrationListResponseSyncDataSyncRevisionMarketplaceData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerIntegrationListResponseSyncDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1CustomerIntegrationListResponseSyncDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price billing sync revision data containing billing ID, link URL, and price
// group package billing ID
type V1CustomerIntegrationListResponseSyncDataSyncRevisionPriceBillingData struct {
	// Billing integration id
	BillingID string `json:"billingId" api:"required"`
	// Billing integration url
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// Price group package billing id
	PriceGroupPackageBillingID string `json:"priceGroupPackageBillingId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID                  respjson.Field
		BillingLinkURL             respjson.Field
		PriceGroupPackageBillingID respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationListResponseSyncDataSyncRevisionPriceBillingData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationListResponseSyncDataSyncRevisionPriceBillingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing sync revision data containing billing ID and link URL
type V1CustomerIntegrationListResponseSyncDataSyncRevisionBillingData struct {
	// Billing integration id
	BillingID string `json:"billingId" api:"required"`
	// Billing integration url
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		BillingLinkURL respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationListResponseSyncDataSyncRevisionBillingData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationListResponseSyncDataSyncRevisionBillingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Marketplace sync revision data containing dimensions
type V1CustomerIntegrationListResponseSyncDataSyncRevisionMarketplaceData struct {
	// Dimensions of the marketplace sync revision
	Dimensions string `json:"dimensions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimensions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationListResponseSyncDataSyncRevisionMarketplaceData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationListResponseSyncDataSyncRevisionMarketplaceData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1CustomerIntegrationLinkResponse struct {
	// External billing or CRM integration link
	Data V1CustomerIntegrationLinkResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationLinkResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerIntegrationLinkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External billing or CRM integration link
type V1CustomerIntegrationLinkResponseData struct {
	// Integration details
	ID string `json:"id" api:"required"`
	// Synced entity id
	SyncedEntityID string `json:"syncedEntityId" api:"required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier" api:"required"`
	// Price billing sync revision data containing billing ID, link URL, and price
	// group package billing ID
	SyncData V1CustomerIntegrationLinkResponseDataSyncDataUnion `json:"syncData" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		SyncedEntityID   respjson.Field
		VendorIdentifier respjson.Field
		SyncData         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationLinkResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerIntegrationLinkResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1CustomerIntegrationLinkResponseDataSyncDataUnion contains all possible
// properties and values from
// [V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionPriceBillingData],
// [V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionBillingData],
// [V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionMarketplaceData].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerIntegrationLinkResponseDataSyncDataUnion struct {
	BillingID      string `json:"billingId"`
	BillingLinkURL string `json:"billingLinkUrl"`
	// This field is from variant
	// [V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionPriceBillingData].
	PriceGroupPackageBillingID string `json:"priceGroupPackageBillingId"`
	// This field is from variant
	// [V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionMarketplaceData].
	Dimensions string `json:"dimensions"`
	JSON       struct {
		BillingID                  respjson.Field
		BillingLinkURL             respjson.Field
		PriceGroupPackageBillingID respjson.Field
		Dimensions                 respjson.Field
		raw                        string
	} `json:"-"`
}

func (u V1CustomerIntegrationLinkResponseDataSyncDataUnion) AsSyncRevisionPriceBillingData() (v V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionPriceBillingData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerIntegrationLinkResponseDataSyncDataUnion) AsSyncRevisionBillingData() (v V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionBillingData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerIntegrationLinkResponseDataSyncDataUnion) AsSyncRevisionMarketplaceData() (v V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionMarketplaceData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerIntegrationLinkResponseDataSyncDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1CustomerIntegrationLinkResponseDataSyncDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price billing sync revision data containing billing ID, link URL, and price
// group package billing ID
type V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionPriceBillingData struct {
	// Billing integration id
	BillingID string `json:"billingId" api:"required"`
	// Billing integration url
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// Price group package billing id
	PriceGroupPackageBillingID string `json:"priceGroupPackageBillingId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID                  respjson.Field
		BillingLinkURL             respjson.Field
		PriceGroupPackageBillingID respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionPriceBillingData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionPriceBillingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing sync revision data containing billing ID and link URL
type V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionBillingData struct {
	// Billing integration id
	BillingID string `json:"billingId" api:"required"`
	// Billing integration url
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		BillingLinkURL respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionBillingData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionBillingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Marketplace sync revision data containing dimensions
type V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionMarketplaceData struct {
	// Dimensions of the marketplace sync revision
	Dimensions string `json:"dimensions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimensions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionMarketplaceData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationLinkResponseDataSyncDataSyncRevisionMarketplaceData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response object
type V1CustomerIntegrationUnlinkResponse struct {
	// External billing or CRM integration link
	Data V1CustomerIntegrationUnlinkResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationUnlinkResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerIntegrationUnlinkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External billing or CRM integration link
type V1CustomerIntegrationUnlinkResponseData struct {
	// Integration details
	ID string `json:"id" api:"required"`
	// Synced entity id
	SyncedEntityID string `json:"syncedEntityId" api:"required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier string `json:"vendorIdentifier" api:"required"`
	// Price billing sync revision data containing billing ID, link URL, and price
	// group package billing ID
	SyncData V1CustomerIntegrationUnlinkResponseDataSyncDataUnion `json:"syncData" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		SyncedEntityID   respjson.Field
		VendorIdentifier respjson.Field
		SyncData         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationUnlinkResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CustomerIntegrationUnlinkResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1CustomerIntegrationUnlinkResponseDataSyncDataUnion contains all possible
// properties and values from
// [V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionPriceBillingData],
// [V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionBillingData],
// [V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionMarketplaceData].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1CustomerIntegrationUnlinkResponseDataSyncDataUnion struct {
	BillingID      string `json:"billingId"`
	BillingLinkURL string `json:"billingLinkUrl"`
	// This field is from variant
	// [V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionPriceBillingData].
	PriceGroupPackageBillingID string `json:"priceGroupPackageBillingId"`
	// This field is from variant
	// [V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionMarketplaceData].
	Dimensions string `json:"dimensions"`
	JSON       struct {
		BillingID                  respjson.Field
		BillingLinkURL             respjson.Field
		PriceGroupPackageBillingID respjson.Field
		Dimensions                 respjson.Field
		raw                        string
	} `json:"-"`
}

func (u V1CustomerIntegrationUnlinkResponseDataSyncDataUnion) AsSyncRevisionPriceBillingData() (v V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionPriceBillingData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerIntegrationUnlinkResponseDataSyncDataUnion) AsSyncRevisionBillingData() (v V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionBillingData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1CustomerIntegrationUnlinkResponseDataSyncDataUnion) AsSyncRevisionMarketplaceData() (v V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionMarketplaceData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1CustomerIntegrationUnlinkResponseDataSyncDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1CustomerIntegrationUnlinkResponseDataSyncDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price billing sync revision data containing billing ID, link URL, and price
// group package billing ID
type V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionPriceBillingData struct {
	// Billing integration id
	BillingID string `json:"billingId" api:"required"`
	// Billing integration url
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// Price group package billing id
	PriceGroupPackageBillingID string `json:"priceGroupPackageBillingId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID                  respjson.Field
		BillingLinkURL             respjson.Field
		PriceGroupPackageBillingID respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionPriceBillingData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionPriceBillingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing sync revision data containing billing ID and link URL
type V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionBillingData struct {
	// Billing integration id
	BillingID string `json:"billingId" api:"required"`
	// Billing integration url
	BillingLinkURL string `json:"billingLinkUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingID      respjson.Field
		BillingLinkURL respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionBillingData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionBillingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Marketplace sync revision data containing dimensions
type V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionMarketplaceData struct {
	// Dimensions of the marketplace sync revision
	Dimensions string `json:"dimensions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimensions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionMarketplaceData) RawJSON() string {
	return r.JSON.raw
}
func (r *V1CustomerIntegrationUnlinkResponseDataSyncDataSyncRevisionMarketplaceData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerIntegrationGetParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type V1CustomerIntegrationUpdateParams struct {
	// Synced entity id
	SyncedEntityID param.Opt[string] `json:"syncedEntityId,omitzero" api:"required"`
	ID             string            `path:"id" api:"required" json:"-"`
	paramObj
}

func (r V1CustomerIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerIntegrationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerIntegrationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CustomerIntegrationListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by vendor identifier. Supports comma-separated values for multiple
	// vendors (e.g., STRIPE,HUBSPOT)
	VendorIdentifier param.Opt[string] `query:"vendorIdentifier,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CustomerIntegrationListParams]'s query parameters as
// `url.Values`.
func (r V1CustomerIntegrationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CustomerIntegrationLinkParams struct {
	// Integration details
	ID string `json:"id" api:"required"`
	// Synced entity id
	SyncedEntityID string `json:"syncedEntityId" api:"required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier V1CustomerIntegrationLinkParamsVendorIdentifier `json:"vendorIdentifier,omitzero" api:"required"`
	paramObj
}

func (r V1CustomerIntegrationLinkParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerIntegrationLinkParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerIntegrationLinkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The vendor identifier of integration
type V1CustomerIntegrationLinkParamsVendorIdentifier string

const (
	V1CustomerIntegrationLinkParamsVendorIdentifierAuth0          V1CustomerIntegrationLinkParamsVendorIdentifier = "AUTH0"
	V1CustomerIntegrationLinkParamsVendorIdentifierZuora          V1CustomerIntegrationLinkParamsVendorIdentifier = "ZUORA"
	V1CustomerIntegrationLinkParamsVendorIdentifierStripe         V1CustomerIntegrationLinkParamsVendorIdentifier = "STRIPE"
	V1CustomerIntegrationLinkParamsVendorIdentifierHubspot        V1CustomerIntegrationLinkParamsVendorIdentifier = "HUBSPOT"
	V1CustomerIntegrationLinkParamsVendorIdentifierAwsMarketplace V1CustomerIntegrationLinkParamsVendorIdentifier = "AWS_MARKETPLACE"
	V1CustomerIntegrationLinkParamsVendorIdentifierSnowflake      V1CustomerIntegrationLinkParamsVendorIdentifier = "SNOWFLAKE"
	V1CustomerIntegrationLinkParamsVendorIdentifierSalesforce     V1CustomerIntegrationLinkParamsVendorIdentifier = "SALESFORCE"
	V1CustomerIntegrationLinkParamsVendorIdentifierBigQuery       V1CustomerIntegrationLinkParamsVendorIdentifier = "BIG_QUERY"
	V1CustomerIntegrationLinkParamsVendorIdentifierOpenFga        V1CustomerIntegrationLinkParamsVendorIdentifier = "OPEN_FGA"
	V1CustomerIntegrationLinkParamsVendorIdentifierAppStore       V1CustomerIntegrationLinkParamsVendorIdentifier = "APP_STORE"
)

type V1CustomerIntegrationUnlinkParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
