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
	chat, err := client.Chats.New(context.TODO(), linqapiv3.ChatNewParams{
		From: "+12025551234",
		Message: linqapiv3.ChatNewParamsMessage{
			Parts: []linqapiv3.ChatNewParamsMessagePartUnion{{
				OfText: &linqapiv3.ChatNewParamsMessagePartText{
					Value: "Hello from Linq SDK!",
				},
			}},
		},
		To: []string{"+19876543210"},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", chat.Chat)
}
