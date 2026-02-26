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

// V1EventFeatureService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventFeatureService] method instead.
type V1EventFeatureService struct {
	Options []option.RequestOption
}

// NewV1EventFeatureService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1EventFeatureService(opts ...option.RequestOption) (r V1EventFeatureService) {
	r = V1EventFeatureService{}
	r.Options = opts
	return
}

// Archives a feature, preventing it from being used in new entitlements.
func (r *V1EventFeatureService) ArchiveFeature(ctx context.Context, id string, opts ...option.RequestOption) (res *Feature, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/features/%s/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Creates a new feature with the specified type, metering, and configuration.
func (r *V1EventFeatureService) NewFeature(ctx context.Context, body V1EventFeatureNewFeatureParams, opts ...option.RequestOption) (res *Feature, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/features"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of features in the environment.
func (r *V1EventFeatureService) ListFeatures(ctx context.Context, query V1EventFeatureListFeaturesParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1EventFeatureListFeaturesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/features"
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

// Retrieves a paginated list of features in the environment.
func (r *V1EventFeatureService) ListFeaturesAutoPaging(ctx context.Context, query V1EventFeatureListFeaturesParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1EventFeatureListFeaturesResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListFeatures(ctx, query, opts...))
}

// Retrieves a feature by its unique identifier.
func (r *V1EventFeatureService) GetFeature(ctx context.Context, id string, opts ...option.RequestOption) (res *Feature, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/features/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Restores an archived feature, allowing it to be used in entitlements again.
func (r *V1EventFeatureService) UnarchiveFeature(ctx context.Context, id string, opts ...option.RequestOption) (res *Feature, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/features/%s/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Updates an existing feature's properties such as display name, description, and
// configuration.
func (r *V1EventFeatureService) UpdateFeature(ctx context.Context, id string, body V1EventFeatureUpdateFeatureParams, opts ...option.RequestOption) (res *Feature, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/features/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Response object
type Feature struct {
	// Feature configuration object
	Data FeatureData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Feature) RawJSON() string { return r.JSON.raw }
func (r *Feature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature configuration object
type FeatureData struct {
	// The unique identifier for the feature
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The description for the feature
	Description string `json:"description" api:"required"`
	// The display name for the feature
	DisplayName string `json:"displayName" api:"required"`
	// The configuration data for the feature
	EnumConfiguration []FeatureDataEnumConfiguration `json:"enumConfiguration" api:"required"`
	// The status of the feature
	//
	// Any of "NEW", "SUSPENDED", "ACTIVE".
	FeatureStatus string `json:"featureStatus" api:"required"`
	// The type of the feature
	//
	// Any of "BOOLEAN", "NUMBER", "ENUM".
	FeatureType string `json:"featureType" api:"required"`
	// The units for the feature
	FeatureUnits string `json:"featureUnits" api:"required"`
	// The plural units for the feature
	FeatureUnitsPlural string `json:"featureUnitsPlural" api:"required"`
	// The additional metadata for the feature
	Metadata map[string]string `json:"metadata" api:"required"`
	// The meter type for the feature
	//
	// Any of "None", "FLUCTUATING", "INCREMENTAL".
	MeterType string `json:"meterType" api:"required"`
	// Unit transformation to be applied to the reported usage
	UnitTransformation FeatureDataUnitTransformation `json:"unitTransformation" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		DisplayName        respjson.Field
		EnumConfiguration  respjson.Field
		FeatureStatus      respjson.Field
		FeatureType        respjson.Field
		FeatureUnits       respjson.Field
		FeatureUnitsPlural respjson.Field
		Metadata           respjson.Field
		MeterType          respjson.Field
		UnitTransformation respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureData) RawJSON() string { return r.JSON.raw }
func (r *FeatureData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureDataEnumConfiguration struct {
	// The display name for the enum configuration entity
	DisplayName string `json:"displayName" api:"required"`
	// The unique value identifier for the enum configuration entity
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureDataEnumConfiguration) RawJSON() string { return r.JSON.raw }
func (r *FeatureDataEnumConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unit transformation to be applied to the reported usage
type FeatureDataUnitTransformation struct {
	// Divide usage by this number
	Divide float64 `json:"divide" api:"required"`
	// Singular feature units after the transformation
	FeatureUnits string `json:"featureUnits" api:"required"`
	// Plural feature units after the transformation
	FeatureUnitsPlural string `json:"featureUnitsPlural" api:"required"`
	// After division, either round the result up or down
	//
	// Any of "UP", "DOWN".
	Round string `json:"round" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Divide             respjson.Field
		FeatureUnits       respjson.Field
		FeatureUnitsPlural respjson.Field
		Round              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureDataUnitTransformation) RawJSON() string { return r.JSON.raw }
func (r *FeatureDataUnitTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature configuration object
type V1EventFeatureListFeaturesResponse struct {
	// The unique identifier for the feature
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The description for the feature
	Description string `json:"description" api:"required"`
	// The display name for the feature
	DisplayName string `json:"displayName" api:"required"`
	// The configuration data for the feature
	EnumConfiguration []V1EventFeatureListFeaturesResponseEnumConfiguration `json:"enumConfiguration" api:"required"`
	// The status of the feature
	//
	// Any of "NEW", "SUSPENDED", "ACTIVE".
	FeatureStatus V1EventFeatureListFeaturesResponseFeatureStatus `json:"featureStatus" api:"required"`
	// The type of the feature
	//
	// Any of "BOOLEAN", "NUMBER", "ENUM".
	FeatureType V1EventFeatureListFeaturesResponseFeatureType `json:"featureType" api:"required"`
	// The units for the feature
	FeatureUnits string `json:"featureUnits" api:"required"`
	// The plural units for the feature
	FeatureUnitsPlural string `json:"featureUnitsPlural" api:"required"`
	// The additional metadata for the feature
	Metadata map[string]string `json:"metadata" api:"required"`
	// The meter type for the feature
	//
	// Any of "None", "FLUCTUATING", "INCREMENTAL".
	MeterType V1EventFeatureListFeaturesResponseMeterType `json:"meterType" api:"required"`
	// Unit transformation to be applied to the reported usage
	UnitTransformation V1EventFeatureListFeaturesResponseUnitTransformation `json:"unitTransformation" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		DisplayName        respjson.Field
		EnumConfiguration  respjson.Field
		FeatureStatus      respjson.Field
		FeatureType        respjson.Field
		FeatureUnits       respjson.Field
		FeatureUnitsPlural respjson.Field
		Metadata           respjson.Field
		MeterType          respjson.Field
		UnitTransformation respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventFeatureListFeaturesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1EventFeatureListFeaturesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventFeatureListFeaturesResponseEnumConfiguration struct {
	// The display name for the enum configuration entity
	DisplayName string `json:"displayName" api:"required"`
	// The unique value identifier for the enum configuration entity
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventFeatureListFeaturesResponseEnumConfiguration) RawJSON() string { return r.JSON.raw }
func (r *V1EventFeatureListFeaturesResponseEnumConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the feature
type V1EventFeatureListFeaturesResponseFeatureStatus string

const (
	V1EventFeatureListFeaturesResponseFeatureStatusNew       V1EventFeatureListFeaturesResponseFeatureStatus = "NEW"
	V1EventFeatureListFeaturesResponseFeatureStatusSuspended V1EventFeatureListFeaturesResponseFeatureStatus = "SUSPENDED"
	V1EventFeatureListFeaturesResponseFeatureStatusActive    V1EventFeatureListFeaturesResponseFeatureStatus = "ACTIVE"
)

// The type of the feature
type V1EventFeatureListFeaturesResponseFeatureType string

const (
	V1EventFeatureListFeaturesResponseFeatureTypeBoolean V1EventFeatureListFeaturesResponseFeatureType = "BOOLEAN"
	V1EventFeatureListFeaturesResponseFeatureTypeNumber  V1EventFeatureListFeaturesResponseFeatureType = "NUMBER"
	V1EventFeatureListFeaturesResponseFeatureTypeEnum    V1EventFeatureListFeaturesResponseFeatureType = "ENUM"
)

// The meter type for the feature
type V1EventFeatureListFeaturesResponseMeterType string

const (
	V1EventFeatureListFeaturesResponseMeterTypeNone        V1EventFeatureListFeaturesResponseMeterType = "None"
	V1EventFeatureListFeaturesResponseMeterTypeFluctuating V1EventFeatureListFeaturesResponseMeterType = "FLUCTUATING"
	V1EventFeatureListFeaturesResponseMeterTypeIncremental V1EventFeatureListFeaturesResponseMeterType = "INCREMENTAL"
)

// Unit transformation to be applied to the reported usage
type V1EventFeatureListFeaturesResponseUnitTransformation struct {
	// Divide usage by this number
	Divide float64 `json:"divide" api:"required"`
	// Singular feature units after the transformation
	FeatureUnits string `json:"featureUnits" api:"required"`
	// Plural feature units after the transformation
	FeatureUnitsPlural string `json:"featureUnitsPlural" api:"required"`
	// After division, either round the result up or down
	//
	// Any of "UP", "DOWN".
	Round string `json:"round" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Divide             respjson.Field
		FeatureUnits       respjson.Field
		FeatureUnitsPlural respjson.Field
		Round              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1EventFeatureListFeaturesResponseUnitTransformation) RawJSON() string { return r.JSON.raw }
func (r *V1EventFeatureListFeaturesResponseUnitTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1EventFeatureNewFeatureParams struct {
	// The unique identifier for the feature
	ID string `json:"id" api:"required"`
	// The display name for the feature
	DisplayName string `json:"displayName" api:"required"`
	// The type of the feature
	//
	// Any of "BOOLEAN", "NUMBER", "ENUM".
	FeatureType V1EventFeatureNewFeatureParamsFeatureType `json:"featureType,omitzero" api:"required"`
	// The description for the feature
	Description param.Opt[string] `json:"description,omitzero"`
	// The units for the feature
	FeatureUnits param.Opt[string] `json:"featureUnits,omitzero"`
	// The plural units for the feature
	FeatureUnitsPlural param.Opt[string] `json:"featureUnitsPlural,omitzero"`
	// Unit transformation to be applied to the reported usage
	UnitTransformation V1EventFeatureNewFeatureParamsUnitTransformation `json:"unitTransformation,omitzero"`
	// The configuration data for the feature
	EnumConfiguration []V1EventFeatureNewFeatureParamsEnumConfiguration `json:"enumConfiguration,omitzero"`
	// The status of the feature
	//
	// Any of "NEW", "SUSPENDED", "ACTIVE".
	FeatureStatus V1EventFeatureNewFeatureParamsFeatureStatus `json:"featureStatus,omitzero"`
	// The additional metadata for the feature
	Metadata map[string]string `json:"metadata,omitzero"`
	// The meter type for the feature
	//
	// Any of "None", "FLUCTUATING", "INCREMENTAL".
	MeterType V1EventFeatureNewFeatureParamsMeterType `json:"meterType,omitzero"`
	paramObj
}

func (r V1EventFeatureNewFeatureParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventFeatureNewFeatureParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventFeatureNewFeatureParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the feature
type V1EventFeatureNewFeatureParamsFeatureType string

const (
	V1EventFeatureNewFeatureParamsFeatureTypeBoolean V1EventFeatureNewFeatureParamsFeatureType = "BOOLEAN"
	V1EventFeatureNewFeatureParamsFeatureTypeNumber  V1EventFeatureNewFeatureParamsFeatureType = "NUMBER"
	V1EventFeatureNewFeatureParamsFeatureTypeEnum    V1EventFeatureNewFeatureParamsFeatureType = "ENUM"
)

// The properties DisplayName, Value are required.
type V1EventFeatureNewFeatureParamsEnumConfiguration struct {
	// The display name for the enum configuration entity
	DisplayName string `json:"displayName" api:"required"`
	// The unique value identifier for the enum configuration entity
	Value string `json:"value" api:"required"`
	paramObj
}

func (r V1EventFeatureNewFeatureParamsEnumConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventFeatureNewFeatureParamsEnumConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventFeatureNewFeatureParamsEnumConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the feature
type V1EventFeatureNewFeatureParamsFeatureStatus string

const (
	V1EventFeatureNewFeatureParamsFeatureStatusNew       V1EventFeatureNewFeatureParamsFeatureStatus = "NEW"
	V1EventFeatureNewFeatureParamsFeatureStatusSuspended V1EventFeatureNewFeatureParamsFeatureStatus = "SUSPENDED"
	V1EventFeatureNewFeatureParamsFeatureStatusActive    V1EventFeatureNewFeatureParamsFeatureStatus = "ACTIVE"
)

// The meter type for the feature
type V1EventFeatureNewFeatureParamsMeterType string

const (
	V1EventFeatureNewFeatureParamsMeterTypeNone        V1EventFeatureNewFeatureParamsMeterType = "None"
	V1EventFeatureNewFeatureParamsMeterTypeFluctuating V1EventFeatureNewFeatureParamsMeterType = "FLUCTUATING"
	V1EventFeatureNewFeatureParamsMeterTypeIncremental V1EventFeatureNewFeatureParamsMeterType = "INCREMENTAL"
)

// Unit transformation to be applied to the reported usage
//
// The property Divide is required.
type V1EventFeatureNewFeatureParamsUnitTransformation struct {
	// Divide usage by this number
	Divide int64 `json:"divide" api:"required"`
	// Singular feature units after the transformation
	FeatureUnits param.Opt[string] `json:"featureUnits,omitzero"`
	// Plural feature units after the transformation
	FeatureUnitsPlural param.Opt[string] `json:"featureUnitsPlural,omitzero"`
	// After division, either round the result up or down
	//
	// Any of "UP", "DOWN".
	Round string `json:"round,omitzero"`
	paramObj
}

func (r V1EventFeatureNewFeatureParamsUnitTransformation) MarshalJSON() (data []byte, err error) {
	type shadow V1EventFeatureNewFeatureParamsUnitTransformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventFeatureNewFeatureParamsUnitTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventFeatureNewFeatureParamsUnitTransformation](
		"round", "UP", "DOWN",
	)
}

type V1EventFeatureListFeaturesParams struct {
	// Filter by entity ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Return items that come after this cursor
	After param.Opt[string] `query:"after,omitzero" format:"uuid" json:"-"`
	// Return items that come before this cursor
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	// Filter by feature type. Supports comma-separated values for multiple types
	FeatureType param.Opt[string] `query:"featureType,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by meter type. Supports comma-separated values for multiple types
	MeterType param.Opt[string] `query:"meterType,omitzero" json:"-"`
	// Filter by feature status. Supports comma-separated values for multiple statuses
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter by creation date using range operators: gt, gte, lt, lte
	CreatedAt V1EventFeatureListFeaturesParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1EventFeatureListFeaturesParams]'s query parameters as
// `url.Values`.
func (r V1EventFeatureListFeaturesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by creation date using range operators: gt, gte, lt, lte
type V1EventFeatureListFeaturesParamsCreatedAt struct {
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

// URLQuery serializes [V1EventFeatureListFeaturesParamsCreatedAt]'s query
// parameters as `url.Values`.
func (r V1EventFeatureListFeaturesParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1EventFeatureUpdateFeatureParams struct {
	// The description for the feature
	Description param.Opt[string] `json:"description,omitzero"`
	// The display name for the feature
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// The units for the feature
	FeatureUnits param.Opt[string] `json:"featureUnits,omitzero"`
	// The plural units for the feature
	FeatureUnitsPlural param.Opt[string] `json:"featureUnitsPlural,omitzero"`
	// Unit transformation to be applied to the reported usage
	UnitTransformation V1EventFeatureUpdateFeatureParamsUnitTransformation `json:"unitTransformation,omitzero"`
	// The configuration data for the feature
	EnumConfiguration []V1EventFeatureUpdateFeatureParamsEnumConfiguration `json:"enumConfiguration,omitzero"`
	// The additional metadata for the feature
	Metadata map[string]string                      `json:"metadata,omitzero"`
	Meter    V1EventFeatureUpdateFeatureParamsMeter `json:"meter,omitzero"`
	paramObj
}

func (r V1EventFeatureUpdateFeatureParams) MarshalJSON() (data []byte, err error) {
	type shadow V1EventFeatureUpdateFeatureParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventFeatureUpdateFeatureParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DisplayName, Value are required.
type V1EventFeatureUpdateFeatureParamsEnumConfiguration struct {
	// The display name for the enum configuration entity
	DisplayName string `json:"displayName" api:"required"`
	// The unique value identifier for the enum configuration entity
	Value string `json:"value" api:"required"`
	paramObj
}

func (r V1EventFeatureUpdateFeatureParamsEnumConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1EventFeatureUpdateFeatureParamsEnumConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventFeatureUpdateFeatureParamsEnumConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Aggregation, Filters are required.
type V1EventFeatureUpdateFeatureParamsMeter struct {
	Aggregation V1EventFeatureUpdateFeatureParamsMeterAggregation `json:"aggregation,omitzero" api:"required"`
	Filters     []V1EventFeatureUpdateFeatureParamsMeterFilter    `json:"filters,omitzero" api:"required"`
	paramObj
}

func (r V1EventFeatureUpdateFeatureParamsMeter) MarshalJSON() (data []byte, err error) {
	type shadow V1EventFeatureUpdateFeatureParamsMeter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventFeatureUpdateFeatureParamsMeter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Function is required.
type V1EventFeatureUpdateFeatureParamsMeterAggregation struct {
	// Any of "SUM", "MAX", "MIN", "AVG", "COUNT", "UNIQUE".
	Function string            `json:"function,omitzero" api:"required"`
	Field    param.Opt[string] `json:"field,omitzero"`
	paramObj
}

func (r V1EventFeatureUpdateFeatureParamsMeterAggregation) MarshalJSON() (data []byte, err error) {
	type shadow V1EventFeatureUpdateFeatureParamsMeterAggregation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventFeatureUpdateFeatureParamsMeterAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventFeatureUpdateFeatureParamsMeterAggregation](
		"function", "SUM", "MAX", "MIN", "AVG", "COUNT", "UNIQUE",
	)
}

// The property Conditions is required.
type V1EventFeatureUpdateFeatureParamsMeterFilter struct {
	Conditions []V1EventFeatureUpdateFeatureParamsMeterFilterCondition `json:"conditions,omitzero" api:"required"`
	paramObj
}

func (r V1EventFeatureUpdateFeatureParamsMeterFilter) MarshalJSON() (data []byte, err error) {
	type shadow V1EventFeatureUpdateFeatureParamsMeterFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventFeatureUpdateFeatureParamsMeterFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Field, Operation are required.
type V1EventFeatureUpdateFeatureParamsMeterFilterCondition struct {
	Field string `json:"field" api:"required"`
	// Any of "EQUALS", "NOT_EQUALS", "GREATER_THAN", "GREATER_THAN_OR_EQUAL",
	// "LESS_THAN", "LESS_THAN_OR_EQUAL", "IS_NULL", "IS_NOT_NULL", "CONTAINS",
	// "STARTS_WITH", "ENDS_WITH", "IN".
	Operation string            `json:"operation,omitzero" api:"required"`
	Value     param.Opt[string] `json:"value,omitzero"`
	Values    []string          `json:"values,omitzero"`
	paramObj
}

func (r V1EventFeatureUpdateFeatureParamsMeterFilterCondition) MarshalJSON() (data []byte, err error) {
	type shadow V1EventFeatureUpdateFeatureParamsMeterFilterCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventFeatureUpdateFeatureParamsMeterFilterCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventFeatureUpdateFeatureParamsMeterFilterCondition](
		"operation", "EQUALS", "NOT_EQUALS", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "IS_NULL", "IS_NOT_NULL", "CONTAINS", "STARTS_WITH", "ENDS_WITH", "IN",
	)
}

// Unit transformation to be applied to the reported usage
//
// The property Divide is required.
type V1EventFeatureUpdateFeatureParamsUnitTransformation struct {
	// Divide usage by this number
	Divide int64 `json:"divide" api:"required"`
	// Singular feature units after the transformation
	FeatureUnits param.Opt[string] `json:"featureUnits,omitzero"`
	// Plural feature units after the transformation
	FeatureUnitsPlural param.Opt[string] `json:"featureUnitsPlural,omitzero"`
	// After division, either round the result up or down
	//
	// Any of "UP", "DOWN".
	Round string `json:"round,omitzero"`
	paramObj
}

func (r V1EventFeatureUpdateFeatureParamsUnitTransformation) MarshalJSON() (data []byte, err error) {
	type shadow V1EventFeatureUpdateFeatureParamsUnitTransformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1EventFeatureUpdateFeatureParamsUnitTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1EventFeatureUpdateFeatureParamsUnitTransformation](
		"round", "UP", "DOWN",
	)
}
