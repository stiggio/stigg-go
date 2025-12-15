// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"github.com/stainless-sdks/stigg-go/option"
)

// V1Service contains methods and other services that help with interacting with
// the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	Options   []option.RequestOption
	Customers V1CustomerService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r V1Service) {
	r = V1Service{}
	r.Options = opts
	r.Customers = NewV1CustomerService(opts...)
	return
}
