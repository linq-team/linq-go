// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"github.com/linq-team/linq-go/option"
)

// LinkPaymentService contains methods and other services that help with
// interacting with the linq-api-v3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLinkPaymentService] method instead.
type LinkPaymentService struct {
	Options []option.RequestOption
}

// NewLinkPaymentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLinkPaymentService(opts ...option.RequestOption) (r LinkPaymentService) {
	r = LinkPaymentService{}
	r.Options = opts
	return
}
