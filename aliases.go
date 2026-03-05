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

// A media attachment part
//
// This is an alias to an internal type.
type MediaPartResponse = shared.MediaPartResponse

// Indicates this is a media attachment part
//
// This is an alias to an internal type.
type MediaPartResponseType = shared.MediaPartResponseType

// Equals "media"
const MediaPartResponseTypeMedia = shared.MediaPartResponseTypeMedia

// Messaging service type
//
// This is an alias to an internal type.
type ServiceType = shared.ServiceType

// Equals "iMessage"
const ServiceTypeiMessage = shared.ServiceTypeiMessage

// Equals "SMS"
const ServiceTypeSMS = shared.ServiceTypeSMS

// Equals "RCS"
const ServiceTypeRCS = shared.ServiceTypeRCS

// A text message part
//
// This is an alias to an internal type.
type TextPartResponse = shared.TextPartResponse

// Indicates this is a text message part
//
// This is an alias to an internal type.
type TextPartResponseType = shared.TextPartResponseType

// Equals "text"
const TextPartResponseTypeText = shared.TextPartResponseTypeText
