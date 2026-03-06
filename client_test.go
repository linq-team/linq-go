// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/linq-team/linq-go"
	"github.com/linq-team/linq-go/internal"
	"github.com/linq-team/linq-go/option"
)

type closureTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *closureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

type closureDoer struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (d closureDoer) Do(req *http.Request) (*http.Response, error) {
	return d.fn(req)
}

func TestUserAgentHeader(t *testing.T) {
	var userAgent string
	client := linqgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					userAgent = req.Header.Get("User-Agent")
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)
	client.Chats.New(context.Background(), linqgo.ChatNewParams{
		From: "+12052535597",
		Message: linqgo.MessageContentParam{
			Parts: []linqgo.MessageContentPartUnionParam{{
				OfText: &linqgo.TextPartParam{
					Type:  linqgo.TextPartTypeText,
					Value: "Hello! How can I help you today?",
				},
			}},
		},
		To: []string{"+12052532136"},
	})
	if userAgent != fmt.Sprintf("LinqAPIV3/Go %s", internal.PackageVersion) {
		t.Errorf("Expected User-Agent to be correct, but got: %#v", userAgent)
	}
}

func TestRetryAfter(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := linqgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Chats.New(context.Background(), linqgo.ChatNewParams{
		From: "+12052535597",
		Message: linqgo.MessageContentParam{
			Parts: []linqgo.MessageContentPartUnionParam{{
				OfText: &linqgo.TextPartParam{
					Type:  linqgo.TextPartTypeText,
					Value: "Hello! How can I help you today?",
				},
			}},
		},
		To: []string{"+12052532136"},
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	attempts := len(retryCountHeaders)
	if attempts != 3 {
		t.Errorf("Expected %d attempts, got %d", 3, attempts)
	}

	expectedRetryCountHeaders := []string{"0", "1", "2"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestDeleteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := linqgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeaderDel("X-Stainless-Retry-Count"),
	)
	_, err := client.Chats.New(context.Background(), linqgo.ChatNewParams{
		From: "+12052535597",
		Message: linqgo.MessageContentParam{
			Parts: []linqgo.MessageContentPartUnionParam{{
				OfText: &linqgo.TextPartParam{
					Type:  linqgo.TextPartTypeText,
					Value: "Hello! How can I help you today?",
				},
			}},
		},
		To: []string{"+12052532136"},
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"", "", ""}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestOverwriteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := linqgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeader("X-Stainless-Retry-Count", "42"),
	)
	_, err := client.Chats.New(context.Background(), linqgo.ChatNewParams{
		From: "+12052535597",
		Message: linqgo.MessageContentParam{
			Parts: []linqgo.MessageContentPartUnionParam{{
				OfText: &linqgo.TextPartParam{
					Type:  linqgo.TextPartTypeText,
					Value: "Hello! How can I help you today?",
				},
			}},
		},
		To: []string{"+12052532136"},
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"42", "42", "42"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestRetryAfterMs(t *testing.T) {
	attempts := 0
	client := linqgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts++
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After-Ms"): []string{"100"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Chats.New(context.Background(), linqgo.ChatNewParams{
		From: "+12052535597",
		Message: linqgo.MessageContentParam{
			Parts: []linqgo.MessageContentPartUnionParam{{
				OfText: &linqgo.TextPartParam{
					Type:  linqgo.TextPartTypeText,
					Value: "Hello! How can I help you today?",
				},
			}},
		},
		To: []string{"+12052532136"},
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
	if want := 3; attempts != want {
		t.Errorf("Expected %d attempts, got %d", want, attempts)
	}
}

func TestContextCancel(t *testing.T) {
	client := linqgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := client.Chats.New(cancelCtx, linqgo.ChatNewParams{
		From: "+12052535597",
		Message: linqgo.MessageContentParam{
			Parts: []linqgo.MessageContentPartUnionParam{{
				OfText: &linqgo.TextPartParam{
					Type:  linqgo.TextPartTypeText,
					Value: "Hello! How can I help you today?",
				},
			}},
		},
		To: []string{"+12052532136"},
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
}

func TestContextCancelDelay(t *testing.T) {
	client := linqgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()
	_, err := client.Chats.New(cancelCtx, linqgo.ChatNewParams{
		From: "+12052535597",
		Message: linqgo.MessageContentParam{
			Parts: []linqgo.MessageContentPartUnionParam{{
				OfText: &linqgo.TextPartParam{
					Type:  linqgo.TextPartTypeText,
					Value: "Hello! How can I help you today?",
				},
			}},
		},
		To: []string{"+12052532136"},
	})
	if err == nil {
		t.Error("expected there to be a cancel error")
	}
}

func TestContextDeadline(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan struct{})

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := linqgo.NewClient(
			option.WithAPIKey("My API Key"),
			option.WithHTTPClient(&http.Client{
				Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						<-req.Context().Done()
						return nil, req.Context().Err()
					},
				},
			}),
		)
		_, err := client.Chats.New(deadlineCtx, linqgo.ChatNewParams{
			From: "+12052535597",
			Message: linqgo.MessageContentParam{
				Parts: []linqgo.MessageContentPartUnionParam{{
					OfText: &linqgo.TextPartParam{
						Type:  linqgo.TextPartTypeText,
						Value: "Hello! How can I help you today?",
					},
				}},
			},
			To: []string{"+12052532136"},
		})
		if err == nil {
			t.Error("expected there to be a deadline error")
		}
		close(testDone)
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		if diff := time.Since(deadline); diff < -30*time.Millisecond || 30*time.Millisecond < diff {
			t.Fatalf("client did not return within 30ms of context deadline, got %s", diff)
		}
	}
}

func TestPaginationPreservesCustomHTTPDoer(t *testing.T) {
	requests := make([]string, 0, 2)
	client := linqgo.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(closureDoer{
			fn: func(req *http.Request) (*http.Response, error) {
				requests = append(requests, req.URL.RawQuery)

				body := `{"chats":[{"id":"chat_123","created_at":"2026-01-01T00:00:00Z","display_name":"Test","handles":[],"is_archived":false,"is_group":false,"updated_at":"2026-01-01T00:00:00Z"}],"next_cursor":""}`
				if len(requests) == 1 {
					body = `{"chats":[{"id":"chat_123","created_at":"2026-01-01T00:00:00Z","display_name":"Test","handles":[],"is_archived":false,"is_group":false,"updated_at":"2026-01-01T00:00:00Z"}],"next_cursor":"cursor_2"}`
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     http.Header{"Content-Type": []string{"application/json"}},
					Body:       io.NopCloser(bytes.NewBufferString(body)),
				}, nil
			},
		}),
	)

	page, err := client.Chats.ListChats(context.Background(), linqgo.ChatListChatsParams{
		From: "+12052535597",
	})
	if err != nil {
		t.Fatalf("expected first page to succeed: %s", err)
	}

	nextPage, err := page.GetNextPage()
	if err != nil {
		t.Fatalf("expected second page to succeed: %s", err)
	}
	if nextPage == nil {
		t.Fatal("expected a second page")
	}

	expectedRequests := []string{"from=%2B12052535597", "cursor=cursor_2&from=%2B12052535597"}
	if !reflect.DeepEqual(requests, expectedRequests) {
		t.Fatalf("expected pagination requests %v, got %v", expectedRequests, requests)
	}
}
