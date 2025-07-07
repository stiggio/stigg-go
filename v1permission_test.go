// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/stigg-go"
	"github.com/stainless-sdks/stigg-go/internal/testutil"
	"github.com/stainless-sdks/stigg-go/option"
)

func TestV1PermissionCheck(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.V1.Permissions.Check(context.TODO(), stigg.V1PermissionCheckParams{
		UserID: "userId",
		ResourcesAndActions: []stigg.V1PermissionCheckParamsResourcesAndAction{{
			Action:   "read",
			Resource: "product",
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
