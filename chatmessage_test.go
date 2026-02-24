// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqapiv3_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/linq-team/linq-go"
	"github.com/linq-team/linq-go/internal/testutil"
	"github.com/linq-team/linq-go/option"
)

func TestChatMessageListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Chats.Messages.List(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		linqapiv3.ChatMessageListParams{
			Cursor: linqapiv3.String("cursor"),
			Limit:  linqapiv3.Int(1),
		},
	)
	if err != nil {
		var apierr *linqapiv3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatMessageSendWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Chats.Messages.Send(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		linqapiv3.ChatMessageSendParams{
			Message: linqapiv3.MessageContentParam{
				Parts: []linqapiv3.MessageContentPartUnionParam{{
					OfText: &linqapiv3.MessageContentPartTextParam{
						Value:          "Hello, world!",
						IdempotencyKey: linqapiv3.String("text-part-abc123"),
					},
				}},
				Effect: linqapiv3.MessageEffectParam{
					Name: linqapiv3.String("confetti"),
					Type: linqapiv3.MessageEffectTypeScreen,
				},
				IdempotencyKey:   linqapiv3.String("msg-abc123xyz"),
				PreferredService: linqapiv3.MessageContentPreferredServiceIMessage,
				ReplyTo: linqapiv3.ReplyToParam{
					MessageID: "550e8400-e29b-41d4-a716-446655440000",
					PartIndex: linqapiv3.Int(0),
				},
			},
		},
	)
	if err != nil {
		var apierr *linqapiv3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
