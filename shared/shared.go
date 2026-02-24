// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/linq-team/linq-go/packages/param"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Messaging service type
type ServiceType string

const (
	ServiceTypeIMessage ServiceType = "iMessage"
	ServiceTypeSMS      ServiceType = "SMS"
	ServiceTypeRcs      ServiceType = "RCS"
)
