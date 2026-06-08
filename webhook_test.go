// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo_test

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/linq-team/linq-go"
	"github.com/linq-team/linq-go/option"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

func TestWebhookUnwrap(t *testing.T) {
	client := linqgo.NewClient(
		option.WithWebhookSecret("whsec_c2VjcmV0Cg=="),
		option.WithAPIKey("My API Key"),
	)
	payload := []byte(`{"api_version":"v3","created_at":"2025-11-23T17:30:00Z","data":{"id":"550e8400-e29b-41d4-a716-446655440001","chat":{"id":"550e8400-e29b-41d4-a716-446655440000","health_status":{"doc_url":"https://docs.linqapp.com/guides/chats/chat-health#at-risk","status":"AT_RISK","updated_at":"2026-05-01T18:28:25Z"},"is_group":true,"owner_handle":{"id":"550e8400-e29b-41d4-a716-446655440000","handle":"+15551234567","joined_at":"2025-05-21T15:30:00.000-05:00","service":"iMessage","is_me":false,"left_at":"2019-12-27T18:11:19.117Z","status":"active"}},"direction":"outbound","parts":[{"type":"text","value":"Hello!","text_decorations":[{"range":[0,5],"animation":"shake","style":"bold"}]}],"sender_handle":{"id":"550e8400-e29b-41d4-a716-446655440000","handle":"+15551234567","joined_at":"2025-05-21T15:30:00.000-05:00","service":"iMessage","is_me":false,"left_at":"2019-12-27T18:11:19.117Z","status":"active"},"service":"iMessage","delivered_at":"2026-01-30T20:49:20.352Z","effect":{"name":"gentle","type":"bubble"},"idempotency_key":"unique-key","preferred_service":"iMessage","read_at":null,"reply_to":{"message_id":"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e","part_index":0},"sent_at":"2026-01-30T20:49:19.704Z"},"event_id":"550e8400-e29b-41d4-a716-446655440000","event_type":"message.sent","partner_id":"partner_abc123","trace_id":"abc123def456","webhook_version":"2025-01-01"}`)
	wh, err := standardwebhooks.NewWebhook("whsec_c2VjcmV0Cg==")
	if err != nil {
		t.Fatal("Failed to sign test webhook message", err)
	}
	msgID := "1"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Fatal("Failed to sign test webhook message:", err)
	}
	headers := make(http.Header)
	headers.Set("webhook-signature", sig)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	_, err = client.Webhooks.Unwrap(payload, headers)
	if err != nil {
		t.Fatal("Failed to unwrap webhook:", err)
	}
}
