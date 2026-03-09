// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stiggio/stigg-go"
	"github.com/stiggio/stigg-go/internal/testutil"
	"github.com/stiggio/stigg-go/option"
)

func TestV1EventCreditGrantNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := stigg.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Events.Credits.Grants.New(context.TODO(), stigg.V1EventCreditGrantNewParams{
		Amount:                   0,
		CurrencyID:               "currencyId",
		CustomerID:               "customerId",
		DisplayName:              "displayName",
		GrantType:                stigg.V1EventCreditGrantNewParamsGrantTypePaid,
		AwaitPaymentConfirmation: stigg.Bool(true),
		BillingInformation: stigg.V1EventCreditGrantNewParamsBillingInformation{
			BillingAddress: stigg.V1EventCreditGrantNewParamsBillingInformationBillingAddress{
				City:       stigg.String("city"),
				Country:    stigg.String("country"),
				Line1:      stigg.String("line1"),
				Line2:      stigg.String("line2"),
				PostalCode: stigg.String("postalCode"),
				State:      stigg.String("state"),
			},
			InvoiceDaysUntilDue: stigg.Float(0),
			IsInvoicePaid:       stigg.Bool(true),
		},
		Comment: stigg.String("comment"),
		Cost: stigg.V1EventCreditGrantNewParamsCost{
			Amount:   0,
			Currency: "usd",
		},
		EffectiveAt: stigg.Time(time.Now()),
		ExpireAt:    stigg.Time(time.Now()),
		Metadata: map[string]string{
			"foo": "string",
		},
		PaymentCollectionMethod: stigg.V1EventCreditGrantNewParamsPaymentCollectionMethodCharge,
		Priority:                stigg.Int(0),
		ResourceID:              stigg.String("resourceId"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1EventCreditGrantListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := stigg.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Events.Credits.Grants.List(context.TODO(), stigg.V1EventCreditGrantListParams{
		CustomerID: "customerId",
		After:      stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Before:     stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CreatedAt: stigg.V1EventCreditGrantListParamsCreatedAt{
			Gt:  stigg.Time(time.Now()),
			Gte: stigg.Time(time.Now()),
			Lt:  stigg.Time(time.Now()),
			Lte: stigg.Time(time.Now()),
		},
		CurrencyID: stigg.String("currencyId"),
		Limit:      stigg.Int(1),
		ResourceID: stigg.String("resourceId"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1EventCreditGrantVoid(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := stigg.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Events.Credits.Grants.Void(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
