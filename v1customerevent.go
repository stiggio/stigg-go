// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"github.com/stiggio/stigg-go/option"
)

// V1CustomerEventService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerEventService] method instead.
type V1CustomerEventService struct {
	Options []option.RequestOption
}

// NewV1CustomerEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CustomerEventService(opts ...option.RequestOption) (r V1CustomerEventService) {
	r = V1CustomerEventService{}
	r.Options = opts
	return
}
