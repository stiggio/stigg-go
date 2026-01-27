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

func TestV1NewEvent(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.V1.NewEvent(context.TODO(), stigg.V1NewEventParams{
		Events: []stigg.V1NewEventParamsEvent{{
			CustomerID:     "customerId",
			EventName:      "x",
			IdempotencyKey: "x",
			Dimensions: map[string]stigg.V1NewEventParamsEventDimensionUnion{
				"foo": {
					OfString: stigg.String("string"),
				},
			},
			ResourceID: stigg.String("resourceId"),
			Timestamp:  stigg.Time(time.Now()),
		}},
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1NewUsage(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.V1.NewUsage(context.TODO(), stigg.V1NewUsageParams{
		Usages: []stigg.V1NewUsageParamsUsage{{
			CustomerID: "customerId",
			FeatureID:  "featureId",
			Value:      -9007199254740991,
			CreatedAt:  stigg.Time(time.Now()),
			Dimensions: map[string]string{
				"foo": "string",
			},
			ResourceID:     stigg.String("resourceId"),
			UpdateBehavior: "DELTA",
		}},
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
