// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo

import (
	"github.com/linq-team/linq-go/internal/apierror"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Messaging service type
//
// This is an alias to an internal type.
type ServiceType = shared.ServiceType

// Equals "iMessage"
const ServiceTypeIMessage = shared.ServiceTypeIMessage

// Equals "SMS"
const ServiceTypeSMS = shared.ServiceTypeSMS

// Equals "RCS"
const ServiceTypeRcs = shared.ServiceTypeRcs
