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

// Operations related to custom currencies
//
// V1CreditCustomCurrencyService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CreditCustomCurrencyService] method instead.
type V1CreditCustomCurrencyService struct {
	Options []option.RequestOption
}

// NewV1CreditCustomCurrencyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CreditCustomCurrencyService(opts ...option.RequestOption) (r V1CreditCustomCurrencyService) {
	r = V1CreditCustomCurrencyService{}
	r.Options = opts
	return
}

// Creates a new custom currency in the environment.
func (r *V1CreditCustomCurrencyService) New(ctx context.Context, body V1CreditCustomCurrencyNewParams, opts ...option.RequestOption) (res *CustomCurrency, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/credits/custom-currencies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates an existing custom currency. Only the supplied fields are modified.
func (r *V1CreditCustomCurrencyService) Update(ctx context.Context, currencyID string, body V1CreditCustomCurrencyUpdateParams, opts ...option.RequestOption) (res *CustomCurrency, err error) {
	opts = slices.Concat(r.Options, opts)
	if currencyID == "" {
		err = errors.New("missing required currencyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/credits/custom-currencies/%s", currencyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieves a paginated list of custom currencies in the environment. Archived
// currencies are excluded by default; pass `status=ARCHIVED` (or
// `status=ACTIVE,ARCHIVED`) to include them.
func (r *V1CreditCustomCurrencyService) List(ctx context.Context, query V1CreditCustomCurrencyListParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1CreditCustomCurrencyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/credits/custom-currencies"
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

// Retrieves a paginated list of custom currencies in the environment. Archived
// currencies are excluded by default; pass `status=ARCHIVED` (or
// `status=ACTIVE,ARCHIVED`) to include them.
func (r *V1CreditCustomCurrencyService) ListAutoPaging(ctx context.Context, query V1CreditCustomCurrencyListParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1CreditCustomCurrencyListResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.List(ctx, query, opts...))
}

// Archives a custom currency. Fails if the currency is still associated with any
// active plan or addon — use the associated-entities endpoint first to inspect
// dependencies.
func (r *V1CreditCustomCurrencyService) Archive(ctx context.Context, currencyID string, opts ...option.RequestOption) (res *CustomCurrency, err error) {
	opts = slices.Concat(r.Options, opts)
	if currencyID == "" {
		err = errors.New("missing required currencyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/credits/custom-currencies/%s/archive", currencyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Lists the active plans and addons that reference a custom currency. Useful
// before archiving to inspect dependencies.
func (r *V1CreditCustomCurrencyService) ListAssociatedEntities(ctx context.Context, currencyID string, opts ...option.RequestOption) (res *V1CreditCustomCurrencyListAssociatedEntitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if currencyID == "" {
		err = errors.New("missing required currencyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/credits/custom-currencies/%s/associated-entities", currencyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Restores a previously archived custom currency. Fails if another active currency
// with the same ID already exists.
func (r *V1CreditCustomCurrencyService) Unarchive(ctx context.Context, currencyID string, opts ...option.RequestOption) (res *CustomCurrency, err error) {
	opts = slices.Concat(r.Options, opts)
	if currencyID == "" {
		err = errors.New("missing required currencyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/credits/custom-currencies/%s/unarchive", currencyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Response object
type CustomCurrency struct {
	// A custom currency used to denominate credit-based entitlements and pricing
	Data CustomCurrencyData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomCurrency) RawJSON() string { return r.JSON.raw }
func (r *CustomCurrency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A custom currency used to denominate credit-based entitlements and pricing
type CustomCurrencyData struct {
	// The unique identifier for the custom currency
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description of the currency
	Description string `json:"description" api:"required"`
	// The display name of the custom currency
	DisplayName string `json:"displayName" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The symbol used to represent the custom currency
	Symbol string `json:"symbol" api:"required"`
	// Singular and plural unit labels for a custom currency
	Units CustomCurrencyDataUnits `json:"units" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ArchivedAt  respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		DisplayName respjson.Field
		Metadata    respjson.Field
		Symbol      respjson.Field
		Units       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomCurrencyData) RawJSON() string { return r.JSON.raw }
func (r *CustomCurrencyData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Singular and plural unit labels for a custom currency
type CustomCurrencyDataUnits struct {
	// Plural form of the unit label
	Plural string `json:"plural" api:"required"`
	// Singular form of the unit label
	Singular string `json:"singular" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Plural      respjson.Field
		Singular    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomCurrencyDataUnits) RawJSON() string { return r.JSON.raw }
func (r *CustomCurrencyDataUnits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A custom currency used to denominate credit-based entitlements and pricing
type V1CreditCustomCurrencyListResponse struct {
	// The unique identifier for the custom currency
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was deleted
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description of the currency
	Description string `json:"description" api:"required"`
	// The display name of the custom currency
	DisplayName string `json:"displayName" api:"required"`
	// Metadata associated with the entity
	Metadata map[string]string `json:"metadata" api:"required"`
	// The symbol used to represent the custom currency
	Symbol string `json:"symbol" api:"required"`
	// Singular and plural unit labels for a custom currency
	Units V1CreditCustomCurrencyListResponseUnits `json:"units" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ArchivedAt  respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		DisplayName respjson.Field
		Metadata    respjson.Field
		Symbol      respjson.Field
		Units       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditCustomCurrencyListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditCustomCurrencyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Singular and plural unit labels for a custom currency
type V1CreditCustomCurrencyListResponseUnits struct {
	// Plural form of the unit label
	Plural string `json:"plural" api:"required"`
	// Singular form of the unit label
	Singular string `json:"singular" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Plural      respjson.Field
		Singular    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditCustomCurrencyListResponseUnits) RawJSON() string { return r.JSON.raw }
func (r *V1CreditCustomCurrencyListResponseUnits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of entities (plans or addons) that reference a custom currency
type V1CreditCustomCurrencyListAssociatedEntitiesResponse struct {
	Data []V1CreditCustomCurrencyListAssociatedEntitiesResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditCustomCurrencyListAssociatedEntitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CreditCustomCurrencyListAssociatedEntitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An entity (plan or addon) that references a custom currency
type V1CreditCustomCurrencyListAssociatedEntitiesResponseData struct {
	// The reference ID of the associated entity
	ID string `json:"id" api:"required"`
	// The display name of the associated entity
	DisplayName string `json:"displayName" api:"required"`
	// The kind of entity referencing the currency (e.g., Plan)
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DisplayName respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CreditCustomCurrencyListAssociatedEntitiesResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CreditCustomCurrencyListAssociatedEntitiesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditCustomCurrencyNewParams struct {
	// The unique identifier for the new custom currency
	ID string `json:"id" api:"required"`
	// The display name of the custom currency
	DisplayName string `json:"displayName" api:"required"`
	// Description of the currency
	Description param.Opt[string] `json:"description,omitzero"`
	// The symbol used to represent the custom currency
	Symbol param.Opt[string] `json:"symbol,omitzero"`
	// Additional metadata to attach to the custom currency
	Metadata map[string]string `json:"metadata,omitzero"`
	// Singular and plural unit labels for a custom currency. Both fields are required
	// when supplied.
	Units V1CreditCustomCurrencyNewParamsUnits `json:"units,omitzero"`
	paramObj
}

func (r V1CreditCustomCurrencyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditCustomCurrencyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditCustomCurrencyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Singular and plural unit labels for a custom currency. Both fields are required
// when supplied.
//
// The properties Plural, Singular are required.
type V1CreditCustomCurrencyNewParamsUnits struct {
	// Plural form of the unit label
	Plural string `json:"plural" api:"required"`
	// Singular form of the unit label
	Singular string `json:"singular" api:"required"`
	paramObj
}

func (r V1CreditCustomCurrencyNewParamsUnits) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditCustomCurrencyNewParamsUnits
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditCustomCurrencyNewParamsUnits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditCustomCurrencyUpdateParams struct {
	// A human-readable description of the custom currency. Send an empty string to
	// clear.
	Description param.Opt[string] `json:"description,omitzero"`
	// The symbol used to represent the custom currency. Send an empty string to clear.
	Symbol param.Opt[string] `json:"symbol,omitzero"`
	// The display name of the custom currency
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Additional metadata to attach to the custom currency
	Metadata map[string]string `json:"metadata,omitzero"`
	// Singular and plural unit labels for a custom currency. Both fields are required
	// when supplied.
	Units V1CreditCustomCurrencyUpdateParamsUnits `json:"units,omitzero"`
	paramObj
}

func (r V1CreditCustomCurrencyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditCustomCurrencyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditCustomCurrencyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Singular and plural unit labels for a custom currency. Both fields are required
// when supplied.
//
// The properties Plural, Singular are required.
type V1CreditCustomCurrencyUpdateParamsUnits struct {
	// Plural form of the unit label
	Plural string `json:"plural" api:"required"`
	// Singular form of the unit label
	Singular string `json:"singular" api:"required"`
	paramObj
}

func (r V1CreditCustomCurrencyUpdateParamsUnits) MarshalJSON() (data []byte, err error) {
	type shadow V1CreditCustomCurrencyUpdateParamsUnits
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CreditCustomCurrencyUpdateParamsUnits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CreditCustomCurrencyListParams struct {
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by custom currency status. Supports comma-separated values (e.g.,
	// `ACTIVE,ARCHIVED`). Defaults to `ACTIVE`.
	//
	// Any of "ACTIVE", "ARCHIVED".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CreditCustomCurrencyListParams]'s query parameters as
// `url.Values`.
func (r V1CreditCustomCurrencyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
