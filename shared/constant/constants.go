// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/linq-team/linq-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type HTTPSDocsLinqappComGuidesWebhooksEvents string // Always "https://docs.linqapp.com/guides/webhooks/events"
type IMessageApp string                             // Always "imessage_app"
type Link string                                    // Always "link"

func (c HTTPSDocsLinqappComGuidesWebhooksEvents) Default() HTTPSDocsLinqappComGuidesWebhooksEvents {
	return "https://docs.linqapp.com/guides/webhooks/events"
}
func (c IMessageApp) Default() IMessageApp { return "imessage_app" }
func (c Link) Default() Link               { return "link" }

func (c HTTPSDocsLinqappComGuidesWebhooksEvents) MarshalJSON() ([]byte, error) {
	return marshalString(c)
}
func (c IMessageApp) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Link) MarshalJSON() ([]byte, error)        { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
