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

func TestMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.New(context.TODO(), linqgo.MessageNewParams{
		Message: linqgo.MessageContentParam{
			Parts: []linqgo.MessageContentPartUnionParam{{
				OfText: &linqgo.TextPartParam{
					Type:  linqgo.TextPartTypeText,
					Value: "Hi! Thanks for reaching out — how can we help?",
					TextDecorations: []shared.TextDecorationParam{{
						Range:     []int64{0, 5},
						Animation: shared.TextDecorationAnimationShake,
						Style:     shared.TextDecorationStyleBold,
					}, {
						Range:     []int64{6, 11},
						Animation: shared.TextDecorationAnimationShake,
						Style:     shared.TextDecorationStyleBold,
					}},
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
		To: []string{"+14155559876"},
		ContinuationMessage: linqgo.MessageNewParamsContinuationMessage{
			Text: "Hi, it's Acme Support reaching you from a new number.",
		},
		IdempotencyKey: linqgo.String("send-abc123xyz"),
	})
	if err != nil {
		var apierr *linqgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageGet(t *testing.T) {
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
	_, err := client.Messages.Get(context.TODO(), "69a37c7d-af4f-4b5e-af42-e28e98ce873a")
	if err != nil {
		var apierr *linqgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.Update(
		context.TODO(),
		"69a37c7d-af4f-4b5e-af42-e28e98ce873a",
		linqgo.MessageUpdateParams{
			Text:      "This is the edited message content",
			PartIndex: linqgo.Int(0),
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

func TestMessageDelete(t *testing.T) {
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
	err := client.Messages.Delete(context.TODO(), "69a37c7d-af4f-4b5e-af42-e28e98ce873a")
	if err != nil {
		var apierr *linqgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageAddReactionWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.AddReaction(
		context.TODO(),
		"69a37c7d-af4f-4b5e-af42-e28e98ce873a",
		linqgo.MessageAddReactionParams{
			Operation:   linqgo.MessageAddReactionParamsOperationAdd,
			Type:        shared.ReactionTypeLove,
			CustomEmoji: linqgo.String("custom_emoji"),
			PartIndex:   linqgo.Int(1),
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

func TestMessageListMessagesThreadWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.ListMessagesThread(
		context.TODO(),
		"69a37c7d-af4f-4b5e-af42-e28e98ce873a",
		linqgo.MessageListMessagesThreadParams{
			Cursor: linqgo.String("cursor"),
			Limit:  linqgo.Int(1),
			Order:  linqgo.MessageListMessagesThreadParamsOrderAsc,
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

func TestMessageUpdateAppCardWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.UpdateAppCard(
		context.TODO(),
		"69a37c7d-af4f-4b5e-af42-e28e98ce873a",
		linqgo.MessageUpdateAppCardParams{
			Layout: linqgo.MessageUpdateAppCardParamsLayout{
				Caption:            linqgo.String("Score: 2 – 1"),
				Subcaption:         linqgo.String("You said: hello"),
				TrailingCaption:    linqgo.String("2 min"),
				TrailingSubcaption: linqgo.String("expires"),
			},
			FallbackText: linqgo.String("Score update"),
			Interactive:  linqgo.Bool(true),
			URL:          linqgo.String("https://app.example.com/card?game=7f3a&move=2"),
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
