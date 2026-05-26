// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"github.com/stiggio/stigg-go/option"
)

// V1BetaService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BetaService] method instead.
type V1BetaService struct {
	Options     []option.RequestOption
	Customers   V1BetaCustomerService
	EntityTypes V1BetaEntityTypeService
}

// NewV1BetaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1BetaService(opts ...option.RequestOption) (r V1BetaService) {
	r = V1BetaService{}
	r.Options = opts
	r.Customers = NewV1BetaCustomerService(opts...)
	r.EntityTypes = NewV1BetaEntityTypeService(opts...)
	return
}
