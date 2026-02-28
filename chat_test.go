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

func TestChatNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.New(context.TODO(), linqgo.ChatNewParams{
		From: "+12052535597",
		Message: linqgo.MessageContentParam{
			Parts: []linqgo.MessageContentPartUnionParam{{
				OfText: &linqgo.MessageContentPartTextParam{
					Value: "Hello! How can I help you today?",
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
		To: []string{"+12052532136"},
	})
	if err != nil {
		var apierr *linqgo.Error
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
	client := linqgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chats.Get(context.TODO(), "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		var apierr *linqgo.Error
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
	client := linqgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chats.Update(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		linqgo.ChatUpdateParams{
			DisplayName:   linqgo.String("Team Discussion"),
			GroupChatIcon: linqgo.String("https://example.com/icon.png"),
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

func TestChatListWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.List(context.TODO(), linqgo.ChatListParams{
		From:   "+13343284472",
		Cursor: linqgo.String("20"),
		Limit:  linqgo.Int(20),
	})
	if err != nil {
		var apierr *linqgo.Error
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
	client := linqgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Chats.MarkAsRead(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *linqgo.Error
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
	client := linqgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chats.SendVoicememo(
		context.TODO(),
		"f19ee7b8-8533-4c5c-83ec-4ef8d6d1ddbd",
		linqgo.ChatSendVoicememoParams{
			From:         "+12052535597",
			VoiceMemoURL: "https://example.com/voice-memo.m4a",
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

func TestChatShareContactCard(t *testing.T) {
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
	err := client.Chats.ShareContactCard(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *linqgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
