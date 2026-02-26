// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"github.com/stiggio/stigg-go/option"
)

// V1Service contains methods and other services that help with interacting with
// the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	Options       []option.RequestOption
	Customers     V1CustomerService
	Subscriptions V1SubscriptionService
	Coupons       V1CouponService
	Events        V1EventService
	Features      V1FeatureService
	Addons        V1AddonService
	Plans         V1PlanService
	Usage         V1UsageService
	Products      V1ProductService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r V1Service) {
	r = V1Service{}
	r.Options = opts
	r.Customers = NewV1CustomerService(opts...)
	r.Subscriptions = NewV1SubscriptionService(opts...)
	r.Coupons = NewV1CouponService(opts...)
	r.Events = NewV1EventService(opts...)
	r.Features = NewV1FeatureService(opts...)
	r.Addons = NewV1AddonService(opts...)
	r.Plans = NewV1PlanService(opts...)
	r.Usage = NewV1UsageService(opts...)
	r.Products = NewV1ProductService(opts...)
	return
}
