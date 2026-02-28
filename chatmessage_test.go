// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/linq-team/linq-go"
	"github.com/linq-team/linq-go/internal/testutil"
	"github.com/linq-team/linq-go/option"
	"github.com/linq-team/linq-go/shared"
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
	client := linqgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chats.Messages.List(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		linqgo.ChatMessageListParams{
			Cursor: linqgo.String("cursor"),
			Limit:  linqgo.Int(1),
		},
	)
	if err != nil {
		var apierr *linqgo.Error
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
	client := linqgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chats.Messages.Send(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		linqgo.ChatMessageSendParams{
			Message: linqgo.MessageContentParam{
				Parts: []linqgo.MessageContentPartUnionParam{{
					OfText: &linqgo.MessageContentPartTextParam{
						Value: "Hello, world!",
					},
				}},
				Effect: linqgo.MessageEffectParam{
					Name: linqgo.String("confetti"),
					Type: linqgo.MessageEffectTypeScreen,
				},
				IdempotencyKey:   linqgo.String("msg-abc123xyz"),
				PreferredService: shared.ServiceTypeIMessage,
				ReplyTo: linqgo.ReplyToParam{
					MessageID: "550e8400-e29b-41d4-a716-446655440000",
					PartIndex: linqgo.Int(0),
				},
			},
		},
	)
	if err != nil {
		var apierr *linqgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
