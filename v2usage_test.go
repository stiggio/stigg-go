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

func TestV2UsageReportWithOptionalParams(t *testing.T) {
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
	_, err := client.V2.Usage.Report(context.TODO(), stigg.V2UsageReportParams{
		Usages: []stigg.V2UsageReportParamsUsage{{
			CustomerID: "customerId",
			FeatureID:  "featureId",
			Value:      -9007199254740991,
			CreatedAt:  stigg.Time(time.Now()),
			Dimensions: map[string]stigg.V2UsageReportParamsUsageDimensionUnion{
				"foo": {
					OfString: stigg.String("string"),
				},
			},
			IdempotencyKey: stigg.String("x"),
			ResourceID:     stigg.String("resourceId"),
			UpdateBehavior: "DELTA",
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
