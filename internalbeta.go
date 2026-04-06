// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"github.com/stiggio/stigg-go/option"
)

// InternalBetaService contains methods and other services that help with
// interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInternalBetaService] method instead.
type InternalBetaService struct {
	Options     []option.RequestOption
	EventQueues InternalBetaEventQueueService
}

// NewInternalBetaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInternalBetaService(opts ...option.RequestOption) (r InternalBetaService) {
	r = InternalBetaService{}
	r.Options = opts
	r.EventQueues = NewInternalBetaEventQueueService(opts...)
	return
}
