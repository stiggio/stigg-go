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

func TestV1EventBetaCustomerAssignmentListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Events.Beta.Customers.Assignments.List(
		context.TODO(),
		"id",
		stigg.V1EventBetaCustomerAssignmentListParams{
			After:          stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Before:         stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			CapabilityID:   stigg.String("capabilityId"),
			EntityID:       stigg.String("entityId"),
			Limit:          stigg.Int(1),
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

func TestV1EventBetaCustomerAssignmentUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Events.Beta.Customers.Assignments.Upsert(
		context.TODO(),
		"id",
		stigg.V1EventBetaCustomerAssignmentUpsertParams{
			Assignments: []stigg.V1EventBetaCustomerAssignmentUpsertParamsAssignment{{
				CapabilityID: "compute-minutes",
				EntityID:     "workspace-001",
				Cadence:      "MONTH",
				UsageLimit:   stigg.Float(1000),
			}, {
				CapabilityID: "compute-minutes",
				EntityID:     "workspace-002",
				Cadence:      "MONTH",
				UsageLimit:   stigg.Float(2000),
			}},
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
