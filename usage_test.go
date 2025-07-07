// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/stigg-go"
	"github.com/stainless-sdks/stigg-go/internal/testutil"
	"github.com/stainless-sdks/stigg-go/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.V1.Permissions.Check(context.TODO(), stigg.V1PermissionCheckParams{
		UserID: "REPLACE_ME",
		ResourcesAndActions: []stigg.V1PermissionCheckParamsResourcesAndAction{{
			Action:   "read",
			Resource: "product",
		}},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.PermittedList)
}
