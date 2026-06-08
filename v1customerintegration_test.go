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

func TestV1CustomerIntegrationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Integrations.Get(
		context.TODO(),
		"integrationId",
		stigg.V1CustomerIntegrationGetParams{
			ID:             "id",
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

func TestV1CustomerIntegrationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Integrations.Update(
		context.TODO(),
		"integrationId",
		stigg.V1CustomerIntegrationUpdateParams{
			ID:             "id",
			SyncedEntityID: stigg.String("syncedEntityId"),
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

func TestV1CustomerIntegrationListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Integrations.List(
		context.TODO(),
		"x",
		stigg.V1CustomerIntegrationListParams{
			After:            stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Before:           stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Limit:            stigg.Int(1),
			VendorIdentifier: []string{"AUTH0"},
			XAccountID:       stigg.String("X-ACCOUNT-ID"),
			XEnvironmentID:   stigg.String("X-ENVIRONMENT-ID"),
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

func TestV1CustomerIntegrationLinkWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Integrations.Link(
		context.TODO(),
		"x",
		stigg.V1CustomerIntegrationLinkParams{
			ID:               "id",
			SyncedEntityID:   "syncedEntityId",
			VendorIdentifier: stigg.V1CustomerIntegrationLinkParamsVendorIdentifierAuth0,
			XAccountID:       stigg.String("X-ACCOUNT-ID"),
			XEnvironmentID:   stigg.String("X-ENVIRONMENT-ID"),
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

func TestV1CustomerIntegrationUnlinkWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Integrations.Unlink(
		context.TODO(),
		"integrationId",
		stigg.V1CustomerIntegrationUnlinkParams{
			ID:             "id",
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
