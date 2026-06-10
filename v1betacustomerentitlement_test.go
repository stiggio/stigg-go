// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stiggio/stigg-go"
	"github.com/stiggio/stigg-go/internal/testutil"
	"github.com/stiggio/stigg-go/option"
)

func TestV1BetaCustomerEntitlementCheckWithOptionalParams(t *testing.T) {
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
	_, err := client.V1Beta.Customers.Entitlements.Check(
		context.TODO(),
		"x",
		stigg.V1BetaCustomerEntitlementCheckParams{
			CurrencyID: stigg.String("x"),
			Dimensions: map[string]string{
				"foo": "string",
			},
			FeatureID:       stigg.String("x"),
			RequestedUsage:  stigg.Int(0),
			RequestedValues: []string{"string"},
			ResourceID:      stigg.String("x"),
			XAccountID:      stigg.String("X-ACCOUNT-ID"),
			XEnvironmentID:  stigg.String("X-ENVIRONMENT-ID"),
		},
	)
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
