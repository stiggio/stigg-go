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
	"github.com/stiggio/stigg-go/packages/param"
	"github.com/stiggio/stigg-go/packages/respjson"
)

// InternalBetaEventQueueService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInternalBetaEventQueueService] method instead.
type InternalBetaEventQueueService struct {
	Options []option.RequestOption
}

// NewInternalBetaEventQueueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInternalBetaEventQueueService(opts ...option.RequestOption) (r InternalBetaEventQueueService) {
	r = InternalBetaEventQueueService{}
	r.Options = opts
	return
}

// Get event queue by queue name
func (r *InternalBetaEventQueueService) Get(ctx context.Context, queueName string, opts ...option.RequestOption) (res *EventQueueResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if queueName == "" {
		err = errors.New("missing required queueName parameter")
		return nil, err
	}
	path := fmt.Sprintf("internal/beta/event-queues/%s", queueName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update event queue configuration
func (r *InternalBetaEventQueueService) Update(ctx context.Context, queueName string, body InternalBetaEventQueueUpdateParams, opts ...option.RequestOption) (res *EventQueueResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if queueName == "" {
		err = errors.New("missing required queueName parameter")
		return nil, err
	}
	path := fmt.Sprintf("internal/beta/event-queues/%s", queueName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all event queues for the current environment
func (r *InternalBetaEventQueueService) List(ctx context.Context, opts ...option.RequestOption) (res *InternalBetaEventQueueListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "internal/beta/event-queues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete an event queue and tear down its infrastructure
func (r *InternalBetaEventQueueService) Delete(ctx context.Context, queueName string, opts ...option.RequestOption) (res *EventQueueResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if queueName == "" {
		err = errors.New("missing required queueName parameter")
		return nil, err
	}
	path := fmt.Sprintf("internal/beta/event-queues/%s", queueName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Provision SQS queue, SNS subscriptions, and IAM role for the current environment
func (r *InternalBetaEventQueueService) Provision(ctx context.Context, body InternalBetaEventQueueProvisionParams, opts ...option.RequestOption) (res *EventQueueResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "internal/beta/event-queues/provision"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Response object
type EventQueueResponse struct {
	// Event queue provisioning status and details
	Data EventQueueResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventQueueResponse) RawJSON() string { return r.JSON.raw }
func (r *EventQueueResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event queue provisioning status and details
type EventQueueResponseData struct {
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Unique queue identifier
	QueueName string `json:"queueName" api:"required"`
	// AWS region where the queue is deployed
	//
	// Any of "us-east-1", "us-east-2", "us-west-1", "us-west-2", "ca-central-1",
	// "eu-west-1", "eu-west-2", "eu-west-3", "eu-central-1", "eu-central-2",
	// "eu-north-1", "eu-south-1", "eu-south-2", "ap-southeast-1", "ap-southeast-2",
	// "ap-southeast-3", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3",
	// "ap-south-1", "ap-south-2", "ap-east-1", "sa-east-1", "af-south-1",
	// "me-south-1", "me-central-1", "il-central-1".
	Region string `json:"region" api:"required"`
	// Current provisioning status
	//
	// Any of "PROVISIONING", "ACTIVE", "FAILED", "DEPROVISIONING".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// SQS queue URL
	QueueURL string `json:"queueUrl" api:"nullable"`
	// IAM role ARN for queue access
	RoleArn string `json:"roleArn" api:"nullable"`
	// Queue suffix for disambiguation
	Suffix string `json:"suffix" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		QueueName   respjson.Field
		Region      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		QueueURL    respjson.Field
		RoleArn     respjson.Field
		Suffix      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventQueueResponseData) RawJSON() string { return r.JSON.raw }
func (r *EventQueueResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response list object
type InternalBetaEventQueueListResponse struct {
	Data []InternalBetaEventQueueListResponseData `json:"data" api:"required"`
	// Pagination metadata including cursors for navigating through results
	Pagination InternalBetaEventQueueListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternalBetaEventQueueListResponse) RawJSON() string { return r.JSON.raw }
func (r *InternalBetaEventQueueListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event queue provisioning status and details
type InternalBetaEventQueueListResponseData struct {
	// Timestamp of when the record was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Unique queue identifier
	QueueName string `json:"queueName" api:"required"`
	// AWS region where the queue is deployed
	//
	// Any of "us-east-1", "us-east-2", "us-west-1", "us-west-2", "ca-central-1",
	// "eu-west-1", "eu-west-2", "eu-west-3", "eu-central-1", "eu-central-2",
	// "eu-north-1", "eu-south-1", "eu-south-2", "ap-southeast-1", "ap-southeast-2",
	// "ap-southeast-3", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3",
	// "ap-south-1", "ap-south-2", "ap-east-1", "sa-east-1", "af-south-1",
	// "me-south-1", "me-central-1", "il-central-1".
	Region string `json:"region" api:"required"`
	// Current provisioning status
	//
	// Any of "PROVISIONING", "ACTIVE", "FAILED", "DEPROVISIONING".
	Status string `json:"status" api:"required"`
	// Timestamp of when the record was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// SQS queue URL
	QueueURL string `json:"queueUrl" api:"nullable"`
	// IAM role ARN for queue access
	RoleArn string `json:"roleArn" api:"nullable"`
	// Queue suffix for disambiguation
	Suffix string `json:"suffix" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		QueueName   respjson.Field
		Region      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		QueueURL    respjson.Field
		RoleArn     respjson.Field
		Suffix      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternalBetaEventQueueListResponseData) RawJSON() string { return r.JSON.raw }
func (r *InternalBetaEventQueueListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata including cursors for navigating through results
type InternalBetaEventQueueListResponsePagination struct {
	// Cursor for fetching the next page of results, or null if no additional pages
	// exist
	Next string `json:"next" api:"required" format:"uuid"`
	// Cursor for fetching the previous page of results, or null if at the beginning
	Prev string `json:"prev" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		Prev        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternalBetaEventQueueListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *InternalBetaEventQueueListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternalBetaEventQueueUpdateParams struct {
	// Whether to create separate low-priority queues for standard topic events
	CreateLowPriorityQueues param.Opt[bool] `json:"createLowPriorityQueues,omitzero"`
	AllowedAssumeRoleArns   []string        `json:"allowedAssumeRoleArns,omitzero"`
	// Any of "MEMBER_INVITED", "SYNC_SUBSCRIPTION", "SYNC_CREDIT_GRANT",
	// "CUSTOMER_CREATED", "CUSTOMER_UPDATED", "CUSTOMER_DELETED", "SYNC_CUSTOMER",
	// "SUBSCRIPTION_CREATED", "SUBSCRIPTION_CANCELED", "SUBSCRIPTION_EXPIRED",
	// "SUBSCRIPTION_UPDATED", "SUBSCRIPTION_TRIAL_STARTED",
	// "SUBSCRIPTION_TRIAL_EXPIRED", "SUBSCRIPTION_TRIAL_CONVERTED",
	// "SUBSCRIPTION_TRIAL_ENDS_SOON", "SYNC_SUBSCRIPTION_USAGE",
	// "SUBSCRIPTION_USAGE_UPDATED", "SUBSCRIPTION_SPENT_LIMIT_EXCEEDED",
	// "CREATE_SUBSCRIPTION_FAILED", "PLAN_CREATED", "PLAN_UPDATED", "PLAN_DELETED",
	// "ADDON_CREATED", "ADDON_UPDATED", "ADDON_DELETED", "SYNC_PACKAGE",
	// "FEATURE_CREATED", "FEATURE_UPDATED", "FEATURE_DELETED", "FEATURE_ARCHIVED",
	// "API_KEY_CREATED", "API_KEY_UPDATED", "API_KEY_ROTATED", "API_KEY_REVOKED",
	// "ENTITLEMENT_REQUESTED", "ENTITLEMENT_GRANTED", "ENTITLEMENT_DENIED",
	// "MEASUREMENT_REPORTED", "USAGE_THRESHOLD_EXCEEDED",
	// "PROMOTIONAL_ENTITLEMENT_GRANTED", "PROMOTIONAL_ENTITLEMENT_REVOKED",
	// "PROMOTIONAL_ENTITLEMENT_UPDATED", "PROMOTIONAL_ENTITLEMENT_EXPIRED",
	// "PROMOTIONAL_ENTITLEMENT_ENDS_SOON", "PACKAGE_PUBLISHED",
	// "MIGRATE_SUBSCRIPTIONS", "RECALCULATE_MIGRATED_ENTITLEMENTS_BATCH",
	// "MIGRATE_SUBSCRIPTIONS_SCHEDULED_UPDATES", "ENTITLEMENTS_UPDATED",
	// "RESYNC_INTEGRATION_TRIGGERED", "COUPON_CREATED", "COUPON_UPDATED",
	// "IMPORT_INTEGRATION_CATALOG_TRIGGERED",
	// "IMPORT_INTEGRATION_CUSTOMERS_TRIGGERED", "INCOMING_STRIPE_WEBHOOK",
	// "INCOMING_AWS_MARKETPLACE_WEBHOOK", "INCOMING_ZUORA_WEBHOOK",
	// "INCOMING_DOGGO_WEBHOOK", "INCOMING_APP_STORE_WEBHOOK", "RESYNC_INTEGRATION",
	// "SYNC_COUPON", "IMPORT_INTEGRATION_CATALOG", "IMPORT_INTEGRATION_CUSTOMERS",
	// "SYNC_FAILED", "CUSTOMER_PAYMENT_FAILED", "PRODUCT_CREATED", "PRODUCT_UPDATED",
	// "PRODUCT_DELETED", "PRODUCT_UNARCHIVED", "PACKAGE_GROUP_CREATED",
	// "PACKAGE_GROUP_UPDATED", "ENVIRONMENT_DELETED", "WIDGET_CONFIGURATION_UPDATED",
	// "EDGE_API_DATA_RESYNC", "EDGE_API_DOGGO_RESYNC",
	// "EDGE_API_CLIENT_CONFIGURATION_DATA_RESYNC",
	// "PURGE_CUSTOMER_PERSISTENT_CACHE_REQUESTED",
	// "CUSTOMER_RESOURCE_ENTITLEMENT_CALCULATION_TRIGGERED",
	// "RECALCULATE_RESOURCE_ENTITLEMENTS",
	// "CUSTOMER_ENTITLEMENT_CALCULATION_TRIGGERED[",
	// "RECALCULATE_ENTITLEMENTS_TRIGGERED", "IMPORT_SUBSCRIPTIONS_BULK_TRIGGERED",
	// "EDGE_API_CUSTOMER_DATA_RESYNC", "EDGE_API_SUBSCRIPTIONS_DATA_RESYNC",
	// "EDGE_API_PACKAGE_ENTITLEMENTS_DATA_RESYNC",
	// "EDGE_API_PRODUCT_CACHE_DATA_RESYNC", "EDGE_API_PLAN_CACHE_DATA_RESYNC",
	// "EDGE_API_CUSTOM_CURRENCY_CACHE_DATA_RESYNC", "REPLAY_WEBHOOK_EVENT",
	// "SUBSCRIPTIONS_MIGRATED", "SUBSCRIPTIONS_MIGRATION_TRIGGERED",
	// "SUBSCRIPTION_BILLING_MONTH_ENDS_SOON", "SUBSCRIPTION_USAGE_CHARGE_TRIGGERED",
	// "SCHEDULER_BATCH", "EVENT_LOG_CREATED", "CREDIT_GRANT_CREATED",
	// "CREDIT_GRANT_EXPIRED", "CREDIT_GRANT_VOIDED", "CREDIT_GRANT_UPDATED",
	// "CREDIT_GRANT_DEPLETED", "CREDIT_GRANT_BALANCE_LOW", "CREDIT_BALANCE_UPDATED",
	// "CREDIT_BALANCE_DEPLETED", "CREDIT_BALANCE_LOW",
	// "CREDIT_GRANT_PROCESS_COMPLETED", "AUTOMATIC_RECHARGE_THRESHOLD_BREACH",
	// "AUTOMATIC_RECHARGE_OPERATION_ATTEMPTED",
	// "CREDITS_AUTOMATIC_RECHARGE_LIMIT_EXCEEDED",
	// "AUTOMATIC_RECHARGE_CONFIGURATION_CHANGED", "FEATURE_GROUP_CREATED",
	// "FEATURE_GROUP_UPDATED", "FEATURE_GROUP_ARCHIVED", "FEATURE_GROUP_UN_ARCHIVED",
	// "CUSTOM_CURRENCY_CREATED", "CUSTOM_CURRENCY_UPDATED",
	// "CUSTOM_CURRENCY_ARCHIVED", "CUSTOM_CURRENCY_UNARCHIVED",
	// "STRIPE_APP_DRAWER_VIEWED", "EVENT_QUEUE_PROVISIONING_REQUESTED",
	// "EVENT_QUEUE_DEPROVISIONING_REQUESTED".
	EventTypes []string `json:"eventTypes,omitzero"`
	paramObj
}

func (r InternalBetaEventQueueUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InternalBetaEventQueueUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InternalBetaEventQueueUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternalBetaEventQueueProvisionParams struct {
	// AWS region for the SQS queue (e.g., us-east-1, eu-west-1)
	//
	// Any of "us-east-1", "us-east-2", "us-west-1", "us-west-2", "ca-central-1",
	// "eu-west-1", "eu-west-2", "eu-west-3", "eu-central-1", "eu-central-2",
	// "eu-north-1", "eu-south-1", "eu-south-2", "ap-southeast-1", "ap-southeast-2",
	// "ap-southeast-3", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3",
	// "ap-south-1", "ap-south-2", "ap-east-1", "sa-east-1", "af-south-1",
	// "me-south-1", "me-central-1", "il-central-1".
	Region InternalBetaEventQueueProvisionParamsRegion `json:"region,omitzero" api:"required"`
	// Whether to create separate low-priority queues for standard topic events
	CreateLowPriorityQueues param.Opt[bool] `json:"createLowPriorityQueues,omitzero"`
	// Optional suffix to allow multiple queues for the same environment and region
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// Additional IAM role ARNs allowed to assume the external role for queue access
	AllowedAssumeRoleArns []string `json:"allowedAssumeRoleArns,omitzero"`
	// Event types to subscribe to. Defaults to entitlements, measurements, and
	// migrations.
	//
	// Any of "MEMBER_INVITED", "SYNC_SUBSCRIPTION", "SYNC_CREDIT_GRANT",
	// "CUSTOMER_CREATED", "CUSTOMER_UPDATED", "CUSTOMER_DELETED", "SYNC_CUSTOMER",
	// "SUBSCRIPTION_CREATED", "SUBSCRIPTION_CANCELED", "SUBSCRIPTION_EXPIRED",
	// "SUBSCRIPTION_UPDATED", "SUBSCRIPTION_TRIAL_STARTED",
	// "SUBSCRIPTION_TRIAL_EXPIRED", "SUBSCRIPTION_TRIAL_CONVERTED",
	// "SUBSCRIPTION_TRIAL_ENDS_SOON", "SYNC_SUBSCRIPTION_USAGE",
	// "SUBSCRIPTION_USAGE_UPDATED", "SUBSCRIPTION_SPENT_LIMIT_EXCEEDED",
	// "CREATE_SUBSCRIPTION_FAILED", "PLAN_CREATED", "PLAN_UPDATED", "PLAN_DELETED",
	// "ADDON_CREATED", "ADDON_UPDATED", "ADDON_DELETED", "SYNC_PACKAGE",
	// "FEATURE_CREATED", "FEATURE_UPDATED", "FEATURE_DELETED", "FEATURE_ARCHIVED",
	// "API_KEY_CREATED", "API_KEY_UPDATED", "API_KEY_ROTATED", "API_KEY_REVOKED",
	// "ENTITLEMENT_REQUESTED", "ENTITLEMENT_GRANTED", "ENTITLEMENT_DENIED",
	// "MEASUREMENT_REPORTED", "USAGE_THRESHOLD_EXCEEDED",
	// "PROMOTIONAL_ENTITLEMENT_GRANTED", "PROMOTIONAL_ENTITLEMENT_REVOKED",
	// "PROMOTIONAL_ENTITLEMENT_UPDATED", "PROMOTIONAL_ENTITLEMENT_EXPIRED",
	// "PROMOTIONAL_ENTITLEMENT_ENDS_SOON", "PACKAGE_PUBLISHED",
	// "MIGRATE_SUBSCRIPTIONS", "RECALCULATE_MIGRATED_ENTITLEMENTS_BATCH",
	// "MIGRATE_SUBSCRIPTIONS_SCHEDULED_UPDATES", "ENTITLEMENTS_UPDATED",
	// "RESYNC_INTEGRATION_TRIGGERED", "COUPON_CREATED", "COUPON_UPDATED",
	// "IMPORT_INTEGRATION_CATALOG_TRIGGERED",
	// "IMPORT_INTEGRATION_CUSTOMERS_TRIGGERED", "INCOMING_STRIPE_WEBHOOK",
	// "INCOMING_AWS_MARKETPLACE_WEBHOOK", "INCOMING_ZUORA_WEBHOOK",
	// "INCOMING_DOGGO_WEBHOOK", "INCOMING_APP_STORE_WEBHOOK", "RESYNC_INTEGRATION",
	// "SYNC_COUPON", "IMPORT_INTEGRATION_CATALOG", "IMPORT_INTEGRATION_CUSTOMERS",
	// "SYNC_FAILED", "CUSTOMER_PAYMENT_FAILED", "PRODUCT_CREATED", "PRODUCT_UPDATED",
	// "PRODUCT_DELETED", "PRODUCT_UNARCHIVED", "PACKAGE_GROUP_CREATED",
	// "PACKAGE_GROUP_UPDATED", "ENVIRONMENT_DELETED", "WIDGET_CONFIGURATION_UPDATED",
	// "EDGE_API_DATA_RESYNC", "EDGE_API_DOGGO_RESYNC",
	// "EDGE_API_CLIENT_CONFIGURATION_DATA_RESYNC",
	// "PURGE_CUSTOMER_PERSISTENT_CACHE_REQUESTED",
	// "CUSTOMER_RESOURCE_ENTITLEMENT_CALCULATION_TRIGGERED",
	// "RECALCULATE_RESOURCE_ENTITLEMENTS",
	// "CUSTOMER_ENTITLEMENT_CALCULATION_TRIGGERED[",
	// "RECALCULATE_ENTITLEMENTS_TRIGGERED", "IMPORT_SUBSCRIPTIONS_BULK_TRIGGERED",
	// "EDGE_API_CUSTOMER_DATA_RESYNC", "EDGE_API_SUBSCRIPTIONS_DATA_RESYNC",
	// "EDGE_API_PACKAGE_ENTITLEMENTS_DATA_RESYNC",
	// "EDGE_API_PRODUCT_CACHE_DATA_RESYNC", "EDGE_API_PLAN_CACHE_DATA_RESYNC",
	// "EDGE_API_CUSTOM_CURRENCY_CACHE_DATA_RESYNC", "REPLAY_WEBHOOK_EVENT",
	// "SUBSCRIPTIONS_MIGRATED", "SUBSCRIPTIONS_MIGRATION_TRIGGERED",
	// "SUBSCRIPTION_BILLING_MONTH_ENDS_SOON", "SUBSCRIPTION_USAGE_CHARGE_TRIGGERED",
	// "SCHEDULER_BATCH", "EVENT_LOG_CREATED", "CREDIT_GRANT_CREATED",
	// "CREDIT_GRANT_EXPIRED", "CREDIT_GRANT_VOIDED", "CREDIT_GRANT_UPDATED",
	// "CREDIT_GRANT_DEPLETED", "CREDIT_GRANT_BALANCE_LOW", "CREDIT_BALANCE_UPDATED",
	// "CREDIT_BALANCE_DEPLETED", "CREDIT_BALANCE_LOW",
	// "CREDIT_GRANT_PROCESS_COMPLETED", "AUTOMATIC_RECHARGE_THRESHOLD_BREACH",
	// "AUTOMATIC_RECHARGE_OPERATION_ATTEMPTED",
	// "CREDITS_AUTOMATIC_RECHARGE_LIMIT_EXCEEDED",
	// "AUTOMATIC_RECHARGE_CONFIGURATION_CHANGED", "FEATURE_GROUP_CREATED",
	// "FEATURE_GROUP_UPDATED", "FEATURE_GROUP_ARCHIVED", "FEATURE_GROUP_UN_ARCHIVED",
	// "CUSTOM_CURRENCY_CREATED", "CUSTOM_CURRENCY_UPDATED",
	// "CUSTOM_CURRENCY_ARCHIVED", "CUSTOM_CURRENCY_UNARCHIVED",
	// "STRIPE_APP_DRAWER_VIEWED", "EVENT_QUEUE_PROVISIONING_REQUESTED",
	// "EVENT_QUEUE_DEPROVISIONING_REQUESTED".
	EventTypes []string `json:"eventTypes,omitzero"`
	paramObj
}

func (r InternalBetaEventQueueProvisionParams) MarshalJSON() (data []byte, err error) {
	type shadow InternalBetaEventQueueProvisionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InternalBetaEventQueueProvisionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AWS region for the SQS queue (e.g., us-east-1, eu-west-1)
type InternalBetaEventQueueProvisionParamsRegion string

const (
	InternalBetaEventQueueProvisionParamsRegionUsEast1      InternalBetaEventQueueProvisionParamsRegion = "us-east-1"
	InternalBetaEventQueueProvisionParamsRegionUsEast2      InternalBetaEventQueueProvisionParamsRegion = "us-east-2"
	InternalBetaEventQueueProvisionParamsRegionUsWest1      InternalBetaEventQueueProvisionParamsRegion = "us-west-1"
	InternalBetaEventQueueProvisionParamsRegionUsWest2      InternalBetaEventQueueProvisionParamsRegion = "us-west-2"
	InternalBetaEventQueueProvisionParamsRegionCaCentral1   InternalBetaEventQueueProvisionParamsRegion = "ca-central-1"
	InternalBetaEventQueueProvisionParamsRegionEuWest1      InternalBetaEventQueueProvisionParamsRegion = "eu-west-1"
	InternalBetaEventQueueProvisionParamsRegionEuWest2      InternalBetaEventQueueProvisionParamsRegion = "eu-west-2"
	InternalBetaEventQueueProvisionParamsRegionEuWest3      InternalBetaEventQueueProvisionParamsRegion = "eu-west-3"
	InternalBetaEventQueueProvisionParamsRegionEuCentral1   InternalBetaEventQueueProvisionParamsRegion = "eu-central-1"
	InternalBetaEventQueueProvisionParamsRegionEuCentral2   InternalBetaEventQueueProvisionParamsRegion = "eu-central-2"
	InternalBetaEventQueueProvisionParamsRegionEuNorth1     InternalBetaEventQueueProvisionParamsRegion = "eu-north-1"
	InternalBetaEventQueueProvisionParamsRegionEuSouth1     InternalBetaEventQueueProvisionParamsRegion = "eu-south-1"
	InternalBetaEventQueueProvisionParamsRegionEuSouth2     InternalBetaEventQueueProvisionParamsRegion = "eu-south-2"
	InternalBetaEventQueueProvisionParamsRegionApSoutheast1 InternalBetaEventQueueProvisionParamsRegion = "ap-southeast-1"
	InternalBetaEventQueueProvisionParamsRegionApSoutheast2 InternalBetaEventQueueProvisionParamsRegion = "ap-southeast-2"
	InternalBetaEventQueueProvisionParamsRegionApSoutheast3 InternalBetaEventQueueProvisionParamsRegion = "ap-southeast-3"
	InternalBetaEventQueueProvisionParamsRegionApNortheast1 InternalBetaEventQueueProvisionParamsRegion = "ap-northeast-1"
	InternalBetaEventQueueProvisionParamsRegionApNortheast2 InternalBetaEventQueueProvisionParamsRegion = "ap-northeast-2"
	InternalBetaEventQueueProvisionParamsRegionApNortheast3 InternalBetaEventQueueProvisionParamsRegion = "ap-northeast-3"
	InternalBetaEventQueueProvisionParamsRegionApSouth1     InternalBetaEventQueueProvisionParamsRegion = "ap-south-1"
	InternalBetaEventQueueProvisionParamsRegionApSouth2     InternalBetaEventQueueProvisionParamsRegion = "ap-south-2"
	InternalBetaEventQueueProvisionParamsRegionApEast1      InternalBetaEventQueueProvisionParamsRegion = "ap-east-1"
	InternalBetaEventQueueProvisionParamsRegionSaEast1      InternalBetaEventQueueProvisionParamsRegion = "sa-east-1"
	InternalBetaEventQueueProvisionParamsRegionAfSouth1     InternalBetaEventQueueProvisionParamsRegion = "af-south-1"
	InternalBetaEventQueueProvisionParamsRegionMeSouth1     InternalBetaEventQueueProvisionParamsRegion = "me-south-1"
	InternalBetaEventQueueProvisionParamsRegionMeCentral1   InternalBetaEventQueueProvisionParamsRegion = "me-central-1"
	InternalBetaEventQueueProvisionParamsRegionIlCentral1   InternalBetaEventQueueProvisionParamsRegion = "il-central-1"
)
