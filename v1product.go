// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
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

// V1ProductService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ProductService] method instead.
type V1ProductService struct {
	Options []option.RequestOption
}

// NewV1ProductService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ProductService(opts ...option.RequestOption) (r V1ProductService) {
	r = V1ProductService{}
	r.Options = opts
	return
}

// Retrieves a paginated list of products in the environment.
func (r *V1ProductService) ListProducts(ctx context.Context, query V1ProductListProductsParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1ProductListProductsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/products"
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

// Retrieves a paginated list of products in the environment.
func (r *V1ProductService) ListProductsAutoPaging(ctx context.Context, query V1ProductListProductsParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1ProductListProductsResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListProducts(ctx, query, opts...))
}

// Product configuration object
type V1ProductListProductsResponse struct {
	// The unique identifier for the entity
	ID string `json:"id,required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Description of the product
	Description string `json:"description,required"`
	// Display name of the product
	DisplayName string `json:"displayName,required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata,required"`
	// Indicates if multiple subscriptions to this product are allowed
	MultipleSubscriptions bool `json:"multipleSubscriptions,required"`
	// The status of the product
	//
	// Any of "PUBLISHED", "ARCHIVED".
	Status V1ProductListProductsResponseStatus `json:"status,required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		Description           respjson.Field
		DisplayName           respjson.Field
		Metadata              respjson.Field
		MultipleSubscriptions respjson.Field
		Status                respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ProductListProductsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ProductListProductsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the product
type V1ProductListProductsResponseStatus string

const (
	V1ProductListProductsResponseStatusPublished V1ProductListProductsResponseStatus = "PUBLISHED"
	V1ProductListProductsResponseStatusArchived  V1ProductListProductsResponseStatus = "ARCHIVED"
)

type V1ProductListProductsParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1ProductListProductsParams]'s query parameters as
// `url.Values`.
func (r V1ProductListProductsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
