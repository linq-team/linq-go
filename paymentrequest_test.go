// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package linqgo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/linq-team/linq-go"
	"github.com/linq-team/linq-go/internal/testutil"
	"github.com/linq-team/linq-go/option"
)

func TestPaymentRequestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentRequests.New(context.TODO(), linqgo.PaymentRequestNewParams{
		Amount:      linqgo.Int(497),
		Currency:    linqgo.String("usd"),
		CustomerID:  linqgo.String("cus_QAbCdEfGhIjKlMn"),
		Description: linqgo.String("Coffee with Ava"),
		Discount: linqgo.PaymentRequestNewParamsDiscount{
			Coupon:        linqgo.String("7fKCMvBh"),
			Label:         linqgo.String("15% OFF FIRST 3 MONTHS"),
			PromotionCode: linqgo.String("promo_1QAbCdEfGhIjKlMn"),
		},
		From: linqgo.String("+12025550123"),
		Metadata: map[string]string{
			"order_id": "order_8675309",
		},
		Mode:            linqgo.PaymentRequestNewParamsModePayment,
		PayerHandle:     linqgo.String("+12025550199"),
		PriceID:         linqgo.String("price_1QAbCdEfGhIjKlMn"),
		Quantity:        linqgo.Int(1),
		Rail:            linqgo.PaymentRequestNewParamsRailStripe,
		TrialEnd:        linqgo.Time(time.Now()),
		TrialPeriodDays: linqgo.Int(14),
		IdempotencyKey:  linqgo.String("pr-abc123xyz"),
	})
	if err != nil {
		var apierr *linqgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentRequestGet(t *testing.T) {
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
	_, err := client.PaymentRequests.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *linqgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentRequests.List(context.TODO(), linqgo.PaymentRequestListParams{
		Limit:  linqgo.Int(1),
		Offset: linqgo.Int(0),
		Status: linqgo.PaymentRequestListParamsStatusRequested,
	})
	if err != nil {
		var apierr *linqgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentRequestCancel(t *testing.T) {
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
	_, err := client.PaymentRequests.Cancel(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *linqgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
