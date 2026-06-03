// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"github.com/stiggio/stigg-go/option"
)

// V1BetaCustomerService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BetaCustomerService] method instead.
type V1BetaCustomerService struct {
	Options      []option.RequestOption
	Entitlements V1BetaCustomerEntitlementService
	Entities     V1BetaCustomerEntityService
	Assignments  V1BetaCustomerAssignmentService
}

// NewV1BetaCustomerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1BetaCustomerService(opts ...option.RequestOption) (r V1BetaCustomerService) {
	r = V1BetaCustomerService{}
	r.Options = opts
	r.Entitlements = NewV1BetaCustomerEntitlementService(opts...)
	r.Entities = NewV1BetaCustomerEntityService(opts...)
	r.Assignments = NewV1BetaCustomerAssignmentService(opts...)
	return
}
