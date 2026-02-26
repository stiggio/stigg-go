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

// V1FeatureService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1FeatureService] method instead.
type V1FeatureService struct {
	Options []option.RequestOption
}

// NewV1FeatureService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1FeatureService(opts ...option.RequestOption) (r V1FeatureService) {
	r = V1FeatureService{}
	r.Options = opts
	return
}

// Archives a feature, preventing it from being used in new entitlements.
func (r *V1FeatureService) ArchiveFeature(ctx context.Context, id string, opts ...option.RequestOption) (res *Feature, err error) {
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
func (r *V1FeatureService) NewFeature(ctx context.Context, body V1FeatureNewFeatureParams, opts ...option.RequestOption) (res *Feature, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/features"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of features in the environment.
func (r *V1FeatureService) ListFeatures(ctx context.Context, query V1FeatureListFeaturesParams, opts ...option.RequestOption) (res *pagination.MyCursorIDPage[V1FeatureListFeaturesResponse], err error) {
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
func (r *V1FeatureService) ListFeaturesAutoPaging(ctx context.Context, query V1FeatureListFeaturesParams, opts ...option.RequestOption) *pagination.MyCursorIDPageAutoPager[V1FeatureListFeaturesResponse] {
	return pagination.NewMyCursorIDPageAutoPager(r.ListFeatures(ctx, query, opts...))
}

// Retrieves a feature by its unique identifier.
func (r *V1FeatureService) GetFeature(ctx context.Context, id string, opts ...option.RequestOption) (res *Feature, err error) {
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
func (r *V1FeatureService) UnarchiveFeature(ctx context.Context, id string, opts ...option.RequestOption) (res *Feature, err error) {
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
func (r *V1FeatureService) UpdateFeature(ctx context.Context, id string, body V1FeatureUpdateFeatureParams, opts ...option.RequestOption) (res *Feature, err error) {
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
type V1FeatureListFeaturesResponse struct {
	// The unique identifier for the feature
	ID string `json:"id" api:"required"`
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The description for the feature
	Description string `json:"description" api:"required"`
	// The display name for the feature
	DisplayName string `json:"displayName" api:"required"`
	// The configuration data for the feature
	EnumConfiguration []V1FeatureListFeaturesResponseEnumConfiguration `json:"enumConfiguration" api:"required"`
	// The status of the feature
	//
	// Any of "NEW", "SUSPENDED", "ACTIVE".
	FeatureStatus V1FeatureListFeaturesResponseFeatureStatus `json:"featureStatus" api:"required"`
	// The type of the feature
	//
	// Any of "BOOLEAN", "NUMBER", "ENUM".
	FeatureType V1FeatureListFeaturesResponseFeatureType `json:"featureType" api:"required"`
	// The units for the feature
	FeatureUnits string `json:"featureUnits" api:"required"`
	// The plural units for the feature
	FeatureUnitsPlural string `json:"featureUnitsPlural" api:"required"`
	// The additional metadata for the feature
	Metadata map[string]string `json:"metadata" api:"required"`
	// The meter type for the feature
	//
	// Any of "None", "FLUCTUATING", "INCREMENTAL".
	MeterType V1FeatureListFeaturesResponseMeterType `json:"meterType" api:"required"`
	// Unit transformation to be applied to the reported usage
	UnitTransformation V1FeatureListFeaturesResponseUnitTransformation `json:"unitTransformation" api:"required"`
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
func (r V1FeatureListFeaturesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1FeatureListFeaturesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1FeatureListFeaturesResponseEnumConfiguration struct {
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
func (r V1FeatureListFeaturesResponseEnumConfiguration) RawJSON() string { return r.JSON.raw }
func (r *V1FeatureListFeaturesResponseEnumConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the feature
type V1FeatureListFeaturesResponseFeatureStatus string

const (
	V1FeatureListFeaturesResponseFeatureStatusNew       V1FeatureListFeaturesResponseFeatureStatus = "NEW"
	V1FeatureListFeaturesResponseFeatureStatusSuspended V1FeatureListFeaturesResponseFeatureStatus = "SUSPENDED"
	V1FeatureListFeaturesResponseFeatureStatusActive    V1FeatureListFeaturesResponseFeatureStatus = "ACTIVE"
)

// The type of the feature
type V1FeatureListFeaturesResponseFeatureType string

const (
	V1FeatureListFeaturesResponseFeatureTypeBoolean V1FeatureListFeaturesResponseFeatureType = "BOOLEAN"
	V1FeatureListFeaturesResponseFeatureTypeNumber  V1FeatureListFeaturesResponseFeatureType = "NUMBER"
	V1FeatureListFeaturesResponseFeatureTypeEnum    V1FeatureListFeaturesResponseFeatureType = "ENUM"
)

// The meter type for the feature
type V1FeatureListFeaturesResponseMeterType string

const (
	V1FeatureListFeaturesResponseMeterTypeNone        V1FeatureListFeaturesResponseMeterType = "None"
	V1FeatureListFeaturesResponseMeterTypeFluctuating V1FeatureListFeaturesResponseMeterType = "FLUCTUATING"
	V1FeatureListFeaturesResponseMeterTypeIncremental V1FeatureListFeaturesResponseMeterType = "INCREMENTAL"
)

// Unit transformation to be applied to the reported usage
type V1FeatureListFeaturesResponseUnitTransformation struct {
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
func (r V1FeatureListFeaturesResponseUnitTransformation) RawJSON() string { return r.JSON.raw }
func (r *V1FeatureListFeaturesResponseUnitTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1FeatureNewFeatureParams struct {
	// The unique identifier for the feature
	ID string `json:"id" api:"required"`
	// The display name for the feature
	DisplayName string `json:"displayName" api:"required"`
	// The type of the feature
	//
	// Any of "BOOLEAN", "NUMBER", "ENUM".
	FeatureType V1FeatureNewFeatureParamsFeatureType `json:"featureType,omitzero" api:"required"`
	// The description for the feature
	Description param.Opt[string] `json:"description,omitzero"`
	// The units for the feature
	FeatureUnits param.Opt[string] `json:"featureUnits,omitzero"`
	// The plural units for the feature
	FeatureUnitsPlural param.Opt[string] `json:"featureUnitsPlural,omitzero"`
	// Unit transformation to be applied to the reported usage
	UnitTransformation V1FeatureNewFeatureParamsUnitTransformation `json:"unitTransformation,omitzero"`
	// The configuration data for the feature
	EnumConfiguration []V1FeatureNewFeatureParamsEnumConfiguration `json:"enumConfiguration,omitzero"`
	// The status of the feature
	//
	// Any of "NEW", "SUSPENDED", "ACTIVE".
	FeatureStatus V1FeatureNewFeatureParamsFeatureStatus `json:"featureStatus,omitzero"`
	// The additional metadata for the feature
	Metadata map[string]string `json:"metadata,omitzero"`
	// The meter type for the feature
	//
	// Any of "None", "FLUCTUATING", "INCREMENTAL".
	MeterType V1FeatureNewFeatureParamsMeterType `json:"meterType,omitzero"`
	paramObj
}

func (r V1FeatureNewFeatureParams) MarshalJSON() (data []byte, err error) {
	type shadow V1FeatureNewFeatureParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1FeatureNewFeatureParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the feature
type V1FeatureNewFeatureParamsFeatureType string

const (
	V1FeatureNewFeatureParamsFeatureTypeBoolean V1FeatureNewFeatureParamsFeatureType = "BOOLEAN"
	V1FeatureNewFeatureParamsFeatureTypeNumber  V1FeatureNewFeatureParamsFeatureType = "NUMBER"
	V1FeatureNewFeatureParamsFeatureTypeEnum    V1FeatureNewFeatureParamsFeatureType = "ENUM"
)

// The properties DisplayName, Value are required.
type V1FeatureNewFeatureParamsEnumConfiguration struct {
	// The display name for the enum configuration entity
	DisplayName string `json:"displayName" api:"required"`
	// The unique value identifier for the enum configuration entity
	Value string `json:"value" api:"required"`
	paramObj
}

func (r V1FeatureNewFeatureParamsEnumConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1FeatureNewFeatureParamsEnumConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1FeatureNewFeatureParamsEnumConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the feature
type V1FeatureNewFeatureParamsFeatureStatus string

const (
	V1FeatureNewFeatureParamsFeatureStatusNew       V1FeatureNewFeatureParamsFeatureStatus = "NEW"
	V1FeatureNewFeatureParamsFeatureStatusSuspended V1FeatureNewFeatureParamsFeatureStatus = "SUSPENDED"
	V1FeatureNewFeatureParamsFeatureStatusActive    V1FeatureNewFeatureParamsFeatureStatus = "ACTIVE"
)

// The meter type for the feature
type V1FeatureNewFeatureParamsMeterType string

const (
	V1FeatureNewFeatureParamsMeterTypeNone        V1FeatureNewFeatureParamsMeterType = "None"
	V1FeatureNewFeatureParamsMeterTypeFluctuating V1FeatureNewFeatureParamsMeterType = "FLUCTUATING"
	V1FeatureNewFeatureParamsMeterTypeIncremental V1FeatureNewFeatureParamsMeterType = "INCREMENTAL"
)

// Unit transformation to be applied to the reported usage
//
// The property Divide is required.
type V1FeatureNewFeatureParamsUnitTransformation struct {
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

func (r V1FeatureNewFeatureParamsUnitTransformation) MarshalJSON() (data []byte, err error) {
	type shadow V1FeatureNewFeatureParamsUnitTransformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1FeatureNewFeatureParamsUnitTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1FeatureNewFeatureParamsUnitTransformation](
		"round", "UP", "DOWN",
	)
}

type V1FeatureListFeaturesParams struct {
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
	CreatedAt V1FeatureListFeaturesParamsCreatedAt `query:"createdAt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1FeatureListFeaturesParams]'s query parameters as
// `url.Values`.
func (r V1FeatureListFeaturesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by creation date using range operators: gt, gte, lt, lte
type V1FeatureListFeaturesParamsCreatedAt struct {
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

// URLQuery serializes [V1FeatureListFeaturesParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r V1FeatureListFeaturesParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1FeatureUpdateFeatureParams struct {
	// The description for the feature
	Description param.Opt[string] `json:"description,omitzero"`
	// The display name for the feature
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// The units for the feature
	FeatureUnits param.Opt[string] `json:"featureUnits,omitzero"`
	// The plural units for the feature
	FeatureUnitsPlural param.Opt[string] `json:"featureUnitsPlural,omitzero"`
	// Unit transformation to be applied to the reported usage
	UnitTransformation V1FeatureUpdateFeatureParamsUnitTransformation `json:"unitTransformation,omitzero"`
	// The configuration data for the feature
	EnumConfiguration []V1FeatureUpdateFeatureParamsEnumConfiguration `json:"enumConfiguration,omitzero"`
	// The additional metadata for the feature
	Metadata map[string]string                 `json:"metadata,omitzero"`
	Meter    V1FeatureUpdateFeatureParamsMeter `json:"meter,omitzero"`
	paramObj
}

func (r V1FeatureUpdateFeatureParams) MarshalJSON() (data []byte, err error) {
	type shadow V1FeatureUpdateFeatureParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1FeatureUpdateFeatureParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DisplayName, Value are required.
type V1FeatureUpdateFeatureParamsEnumConfiguration struct {
	// The display name for the enum configuration entity
	DisplayName string `json:"displayName" api:"required"`
	// The unique value identifier for the enum configuration entity
	Value string `json:"value" api:"required"`
	paramObj
}

func (r V1FeatureUpdateFeatureParamsEnumConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow V1FeatureUpdateFeatureParamsEnumConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1FeatureUpdateFeatureParamsEnumConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Aggregation, Filters are required.
type V1FeatureUpdateFeatureParamsMeter struct {
	Aggregation V1FeatureUpdateFeatureParamsMeterAggregation `json:"aggregation,omitzero" api:"required"`
	Filters     []V1FeatureUpdateFeatureParamsMeterFilter    `json:"filters,omitzero" api:"required"`
	paramObj
}

func (r V1FeatureUpdateFeatureParamsMeter) MarshalJSON() (data []byte, err error) {
	type shadow V1FeatureUpdateFeatureParamsMeter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1FeatureUpdateFeatureParamsMeter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Function is required.
type V1FeatureUpdateFeatureParamsMeterAggregation struct {
	// Any of "SUM", "MAX", "MIN", "AVG", "COUNT", "UNIQUE".
	Function string            `json:"function,omitzero" api:"required"`
	Field    param.Opt[string] `json:"field,omitzero"`
	paramObj
}

func (r V1FeatureUpdateFeatureParamsMeterAggregation) MarshalJSON() (data []byte, err error) {
	type shadow V1FeatureUpdateFeatureParamsMeterAggregation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1FeatureUpdateFeatureParamsMeterAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1FeatureUpdateFeatureParamsMeterAggregation](
		"function", "SUM", "MAX", "MIN", "AVG", "COUNT", "UNIQUE",
	)
}

// The property Conditions is required.
type V1FeatureUpdateFeatureParamsMeterFilter struct {
	Conditions []V1FeatureUpdateFeatureParamsMeterFilterCondition `json:"conditions,omitzero" api:"required"`
	paramObj
}

func (r V1FeatureUpdateFeatureParamsMeterFilter) MarshalJSON() (data []byte, err error) {
	type shadow V1FeatureUpdateFeatureParamsMeterFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1FeatureUpdateFeatureParamsMeterFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Field, Operation are required.
type V1FeatureUpdateFeatureParamsMeterFilterCondition struct {
	Field string `json:"field" api:"required"`
	// Any of "EQUALS", "NOT_EQUALS", "GREATER_THAN", "GREATER_THAN_OR_EQUAL",
	// "LESS_THAN", "LESS_THAN_OR_EQUAL", "IS_NULL", "IS_NOT_NULL", "CONTAINS",
	// "STARTS_WITH", "ENDS_WITH", "IN".
	Operation string            `json:"operation,omitzero" api:"required"`
	Value     param.Opt[string] `json:"value,omitzero"`
	Values    []string          `json:"values,omitzero"`
	paramObj
}

func (r V1FeatureUpdateFeatureParamsMeterFilterCondition) MarshalJSON() (data []byte, err error) {
	type shadow V1FeatureUpdateFeatureParamsMeterFilterCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1FeatureUpdateFeatureParamsMeterFilterCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1FeatureUpdateFeatureParamsMeterFilterCondition](
		"operation", "EQUALS", "NOT_EQUALS", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "IS_NULL", "IS_NOT_NULL", "CONTAINS", "STARTS_WITH", "ENDS_WITH", "IN",
	)
}

// Unit transformation to be applied to the reported usage
//
// The property Divide is required.
type V1FeatureUpdateFeatureParamsUnitTransformation struct {
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

func (r V1FeatureUpdateFeatureParamsUnitTransformation) MarshalJSON() (data []byte, err error) {
	type shadow V1FeatureUpdateFeatureParamsUnitTransformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1FeatureUpdateFeatureParamsUnitTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1FeatureUpdateFeatureParamsUnitTransformation](
		"round", "UP", "DOWN",
	)
}
