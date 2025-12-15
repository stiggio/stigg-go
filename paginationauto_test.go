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

func TestAutoPagination(t *testing.T) {
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
	iter := client.V1.Customers.ListAutoPaging(context.TODO(), stigg.V1CustomerListParams{
		Limit: stigg.Int(30),
	})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		customer := iter.Current()
		t.Logf("%+v\n", customer.CursorID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
