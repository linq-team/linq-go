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

func TestCapabilityCheckImessageWithOptionalParams(t *testing.T) {
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
	_, err := client.Capability.CheckImessage(context.TODO(), linqapiv3.CapabilityCheckImessageParams{
		Address: "+15551234567",
		From:    linqapiv3.String("+15559876543"),
	})
	if err != nil {
		var apierr *linqapiv3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCapabilityCheckRcsWithOptionalParams(t *testing.T) {
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
	_, err := client.Capability.CheckRcs(context.TODO(), linqapiv3.CapabilityCheckRcsParams{
		Address: "+15551234567",
		From:    linqapiv3.String("+15559876543"),
	})
	if err != nil {
		var apierr *linqapiv3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
