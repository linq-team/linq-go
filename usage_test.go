// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo_test

import (
	"context"
	"os"
	"testing"

	"github.com/linq-team/linq-go"
	"github.com/linq-team/linq-go/internal/testutil"
	"github.com/linq-team/linq-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := linqgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Mock server tests are disabled")
	chat, err := client.Chats.New(context.TODO(), linqgo.ChatNewParams{
		From: "+12052535597",
		Message: linqgo.MessageContentParam{
			Parts: []linqgo.MessageContentPartUnionParam{{
				OfText: &linqgo.MessageContentPartTextParam{
					Value: "Hello! How can I help you today?",
				},
			}},
		},
		To: []string{"+12052532136"},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", chat.Chat)
}
