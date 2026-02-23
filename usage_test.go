// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqapiv3_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/linq-api-v3-go"
	"github.com/stainless-sdks/linq-api-v3-go/internal/testutil"
	"github.com/stainless-sdks/linq-api-v3-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := linqapiv3.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Mock server tests are disabled")
	chat, err := client.Chats.New(context.TODO(), linqapiv3.ChatNewParams{
		From: "+12052535597",
		Message: linqapiv3.MessageContentParam{
			Parts: []linqapiv3.MessageContentPartUnionParam{{
				OfText: &linqapiv3.MessageContentPartTextParam{
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
