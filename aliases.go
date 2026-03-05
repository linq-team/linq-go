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

// This is an alias to an internal type.
type Reaction = shared.Reaction

// Sticker attachment details when reaction_type is "sticker". Null for non-sticker
// reactions.
//
// This is an alias to an internal type.
type ReactionSticker = shared.ReactionSticker

// Type of reaction. Standard iMessage tapbacks are love, like, dislike, laugh,
// emphasize, question. Custom emoji reactions have type "custom" with the actual
// emoji in the custom_emoji field. Sticker reactions have type "sticker" with
// sticker attachment details in the sticker field.
//
// This is an alias to an internal type.
type ReactionType = shared.ReactionType

// Equals "love"
const ReactionTypeLove = shared.ReactionTypeLove

// Equals "like"
const ReactionTypeLike = shared.ReactionTypeLike

// Equals "dislike"
const ReactionTypeDislike = shared.ReactionTypeDislike

// Equals "laugh"
const ReactionTypeLaugh = shared.ReactionTypeLaugh

// Equals "emphasize"
const ReactionTypeEmphasize = shared.ReactionTypeEmphasize

// Equals "question"
const ReactionTypeQuestion = shared.ReactionTypeQuestion

// Equals "custom"
const ReactionTypeCustom = shared.ReactionTypeCustom

// Equals "sticker"
const ReactionTypeSticker = shared.ReactionTypeSticker

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
