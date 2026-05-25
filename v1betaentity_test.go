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

func TestV1BetaEntityGet(t *testing.T) {
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
	_, err := client.V1Beta.Entities.Get(
		context.TODO(),
		"x",
		stigg.V1BetaEntityGetParams{
			ID: "id",
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

func TestV1BetaEntityListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1Beta.Entities.List(
		context.TODO(),
		"id",
		stigg.V1BetaEntityListParams{
			After:           stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Before:          stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			IncludeArchived: stigg.V1BetaEntityListParamsIncludeArchivedTrue,
			Limit:           stigg.Int(1),
			TypeRefID:       stigg.String("typeRefId"),
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

func TestV1BetaEntityArchive(t *testing.T) {
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
	_, err := client.V1Beta.Entities.Archive(
		context.TODO(),
		"id",
		stigg.V1BetaEntityArchiveParams{
			IDs: []string{"user-7f3a0c1d", "user-c4d1b2e9"},
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

func TestV1BetaEntityUnarchive(t *testing.T) {
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
	_, err := client.V1Beta.Entities.Unarchive(
		context.TODO(),
		"id",
		stigg.V1BetaEntityUnarchiveParams{
			IDs: []string{"user-7f3a0c1d", "user-c4d1b2e9"},
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

func TestV1BetaEntityUpsert(t *testing.T) {
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
	_, err := client.V1Beta.Entities.Upsert(
		context.TODO(),
		"id",
		stigg.V1BetaEntityUpsertParams{
			Entities: []stigg.V1BetaEntityUpsertParamsEntity{{
				ID: "user-7f3a0c1d",
				Metadata: map[string]string{
					"email": "jane@acme.com",
					"role":  "admin",
				},
				TypeRefID: stigg.String("user"),
			}, {
				ID: "user-c4d1b2e9",
				Metadata: map[string]string{
					"email": "john@acme.com",
				},
				TypeRefID: stigg.String("user"),
			}},
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
