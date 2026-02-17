// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"github.com/stiggio/stigg-go/option"
)

// V1CustomerPromotionalEntitlementService contains methods and other services that
// help with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerPromotionalEntitlementService] method instead.
type V1CustomerPromotionalEntitlementService struct {
	Options []option.RequestOption
}

// NewV1CustomerPromotionalEntitlementService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1CustomerPromotionalEntitlementService(opts ...option.RequestOption) (r V1CustomerPromotionalEntitlementService) {
	r = V1CustomerPromotionalEntitlementService{}
	r.Options = opts
	return
}
