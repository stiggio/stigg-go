// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"github.com/stiggio/stigg-go/option"
)

// InternalService contains methods and other services that help with interacting
// with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInternalService] method instead.
type InternalService struct {
	Options []option.RequestOption
	Beta    InternalBetaService
}

// NewInternalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInternalService(opts ...option.RequestOption) (r InternalService) {
	r = InternalService{}
	r.Options = opts
	r.Beta = NewInternalBetaService(opts...)
	return
}
