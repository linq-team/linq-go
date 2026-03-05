// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/linq-team/linq-go/internal/apijson"
	"github.com/linq-team/linq-go/internal/requestconfig"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/packages/param"
	"github.com/linq-team/linq-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type ListChatsPagination[T any] struct {
	Chats      []T    `json:"chats"`
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chats       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ListChatsPagination[T]) RawJSON() string { return r.JSON.raw }
func (r *ListChatsPagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ListChatsPagination[T]) GetNextPage() (res *ListChatsPagination[T], err error) {
	if len(r.Chats) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ListChatsPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ListChatsPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ListChatsPaginationAutoPager[T any] struct {
	page *ListChatsPagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewListChatsPaginationAutoPager[T any](page *ListChatsPagination[T], err error) *ListChatsPaginationAutoPager[T] {
	return &ListChatsPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ListChatsPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Chats) == 0 {
		return false
	}
	if r.idx >= len(r.page.Chats) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Chats) == 0 {
			return false
		}
	}
	r.cur = r.page.Chats[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ListChatsPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *ListChatsPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *ListChatsPaginationAutoPager[T]) Index() int {
	return r.run
}

type ListMessagesPagination[T any] struct {
	Messages   []T    `json:"messages"`
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages    respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ListMessagesPagination[T]) RawJSON() string { return r.JSON.raw }
func (r *ListMessagesPagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ListMessagesPagination[T]) GetNextPage() (res *ListMessagesPagination[T], err error) {
	if len(r.Messages) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ListMessagesPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ListMessagesPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ListMessagesPaginationAutoPager[T any] struct {
	page *ListMessagesPagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewListMessagesPaginationAutoPager[T any](page *ListMessagesPagination[T], err error) *ListMessagesPaginationAutoPager[T] {
	return &ListMessagesPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ListMessagesPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Messages) == 0 {
		return false
	}
	if r.idx >= len(r.page.Messages) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Messages) == 0 {
			return false
		}
	}
	r.cur = r.page.Messages[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ListMessagesPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *ListMessagesPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *ListMessagesPaginationAutoPager[T]) Index() int {
	return r.run
}
