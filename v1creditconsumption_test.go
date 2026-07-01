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

func TestV1CreditConsumptionConsumeWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Credits.Consumption.Consume(context.TODO(), stigg.V1CreditConsumptionConsumeParams{
		Amount:         1,
		CurrencyID:     "currencyId",
		CustomerID:     "customerId",
		IdempotencyKey: "x",
		CreatedAt:      stigg.Time(time.Now()),
		Dimensions: map[string]stigg.V1CreditConsumptionConsumeParamsDimensionUnion{
			"foo": {
				OfString: stigg.String("string"),
			},
		},
		ResourceID:     stigg.String("resourceId"),
		XAccountID:     stigg.String("X-ACCOUNT-ID"),
		XEnvironmentID: stigg.String("X-ENVIRONMENT-ID"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CreditConsumptionConsumeAsyncWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Credits.Consumption.ConsumeAsync(context.TODO(), stigg.V1CreditConsumptionConsumeAsyncParams{
		Consumptions: []stigg.V1CreditConsumptionConsumeAsyncParamsConsumption{{
			Amount:         1,
			CurrencyID:     "currencyId",
			CustomerID:     "customerId",
			IdempotencyKey: "x",
			CreatedAt:      stigg.Time(time.Now()),
			Dimensions: map[string]stigg.V1CreditConsumptionConsumeAsyncParamsConsumptionDimensionUnion{
				"foo": {
					OfString: stigg.String("string"),
				},
			},
			ResourceID: stigg.String("resourceId"),
		}},
		XAccountID:     stigg.String("X-ACCOUNT-ID"),
		XEnvironmentID: stigg.String("X-ENVIRONMENT-ID"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
