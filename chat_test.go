// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqapiv3_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/linq-api-v3-go"
	"github.com/stainless-sdks/linq-api-v3-go/internal/testutil"
	"github.com/stainless-sdks/linq-api-v3-go/option"
)

func TestChatNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.New(context.TODO(), linqapiv3.ChatNewParams{
		From: "+12052535597",
		Message: linqapiv3.MessageContentParam{
			Parts: []linqapiv3.MessageContentPartUnionParam{{
				OfText: &linqapiv3.MessageContentPartTextParam{
					Value:          "Hello! How can I help you today?",
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
		To: []string{"+12052532136"},
	})
	if err != nil {
		var apierr *linqapiv3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatGet(t *testing.T) {
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
	_, err := client.Chats.Get(context.TODO(), "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		var apierr *linqapiv3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Update(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		linqapiv3.ChatUpdateParams{
			DisplayName:   linqapiv3.String("Team Discussion"),
			GroupChatIcon: linqapiv3.String("https://example.com/icon.png"),
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

func TestChatListWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.List(context.TODO(), linqapiv3.ChatListParams{
		From:   "+13343284472",
		Cursor: linqapiv3.String("20"),
		Limit:  linqapiv3.Int(20),
	})
	if err != nil {
		var apierr *linqapiv3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatMarkAsRead(t *testing.T) {
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
	err := client.Chats.MarkAsRead(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *linqapiv3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatSendVoicememo(t *testing.T) {
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
	_, err := client.Chats.SendVoicememo(
		context.TODO(),
		"f19ee7b8-8533-4c5c-83ec-4ef8d6d1ddbd",
		linqapiv3.ChatSendVoicememoParams{
			From:         "+12052535597",
			VoiceMemoURL: "https://example.com/voice-memo.m4a",
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

func TestChatShareContactCard(t *testing.T) {
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
	err := client.Chats.ShareContactCard(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *linqapiv3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
