// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"github.com/stiggio/stigg-go/option"
)

// V1EventBetaCustomerService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1EventBetaCustomerService] method instead.
type V1EventBetaCustomerService struct {
	Options      []option.RequestOption
	Entitlements V1EventBetaCustomerEntitlementService
}

// NewV1EventBetaCustomerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1EventBetaCustomerService(opts ...option.RequestOption) (r V1EventBetaCustomerService) {
	r = V1EventBetaCustomerService{}
	r.Options = opts
	r.Entitlements = NewV1EventBetaCustomerEntitlementService(opts...)
	return
}
