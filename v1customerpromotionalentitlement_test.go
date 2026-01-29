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

func TestV1CustomerPromotionalEntitlementGrant(t *testing.T) {
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
	_, err := client.V1.Customers.PromotionalEntitlements.Grant(
		context.TODO(),
		"customerId",
		stigg.V1CustomerPromotionalEntitlementGrantParams{
			PromotionalEntitlements: []stigg.V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlement{{
				CustomEndDate:     stigg.Time(time.Now()),
				EnumValues:        []string{"string"},
				FeatureID:         "featureId",
				HasSoftLimit:      stigg.Bool(true),
				HasUnlimitedUsage: stigg.Bool(true),
				IsVisible:         stigg.Bool(true),
				MonthlyResetPeriodConfiguration: stigg.V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementMonthlyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				Period:      "1 week",
				ResetPeriod: "YEAR",
				UsageLimit:  stigg.Int(-9007199254740991),
				WeeklyResetPeriodConfiguration: stigg.V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementWeeklyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				YearlyResetPeriodConfiguration: stigg.V1CustomerPromotionalEntitlementGrantParamsPromotionalEntitlementYearlyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
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

func TestV1CustomerPromotionalEntitlementRevoke(t *testing.T) {
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
	_, err := client.V1.Customers.PromotionalEntitlements.Revoke(
		context.TODO(),
		"featureId",
		stigg.V1CustomerPromotionalEntitlementRevokeParams{
			CustomerID: "customerId",
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
