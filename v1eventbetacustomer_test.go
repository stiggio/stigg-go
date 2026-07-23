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

func TestV1EventBetaCustomerGetGovernanceWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Events.Beta.Customers.GetGovernance(
		context.TODO(),
		"id",
		stigg.V1EventBetaCustomerGetGovernanceParams{
			After:          stigg.String("after"),
			CurrencyIDs:    []string{"string"},
			EntityIDSearch: stigg.String("x"),
			EntityTypeIDs:  []string{"string"},
			FeatureIDs:     []string{"string"},
			Limit:          stigg.Int(1),
			MinUtilization: stigg.Float(0),
			Order:          stigg.V1EventBetaCustomerGetGovernanceParamsOrderAsc,
			Scope:          stigg.V1EventBetaCustomerGetGovernanceParamsScopeAll,
			SortBy:         stigg.V1EventBetaCustomerGetGovernanceParamsSortByUtilization,
			XAccountID:     stigg.String("X-ACCOUNT-ID"),
			XEnvironmentID: stigg.String("X-ENVIRONMENT-ID"),
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
