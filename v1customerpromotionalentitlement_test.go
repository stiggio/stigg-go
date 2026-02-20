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

func TestV1CustomerPromotionalEntitlementNew(t *testing.T) {
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
	_, err := client.V1.Customers.PromotionalEntitlements.New(
		context.TODO(),
		"x",
		stigg.V1CustomerPromotionalEntitlementNewParams{
			PromotionalEntitlements: []stigg.V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlement{{
				CustomEndDate:     stigg.Time(time.Now()),
				EnumValues:        []string{"string"},
				FeatureID:         "featureId",
				HasSoftLimit:      stigg.Bool(true),
				HasUnlimitedUsage: stigg.Bool(true),
				IsVisible:         stigg.Bool(true),
				MonthlyResetPeriodConfiguration: stigg.V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementMonthlyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				Period:      "1 week",
				ResetPeriod: "YEAR",
				UsageLimit:  stigg.Int(-9007199254740991),
				WeeklyResetPeriodConfiguration: stigg.V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementWeeklyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				YearlyResetPeriodConfiguration: stigg.V1CustomerPromotionalEntitlementNewParamsPromotionalEntitlementYearlyResetPeriodConfiguration{
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

func TestV1CustomerPromotionalEntitlementListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.PromotionalEntitlements.List(
		context.TODO(),
		"x",
		stigg.V1CustomerPromotionalEntitlementListParams{
			After:  stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Before: stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			CreatedAt: stigg.V1CustomerPromotionalEntitlementListParamsCreatedAt{
				Gt:  stigg.Time(time.Now()),
				Gte: stigg.Time(time.Now()),
				Lt:  stigg.Time(time.Now()),
				Lte: stigg.Time(time.Now()),
			},
			Limit:  stigg.Int(1),
			Status: stigg.String("status"),
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
	_, err := client.V1.Customers.PromotionalEntitlements.Revoke(
		context.TODO(),
		"featureId",
		stigg.V1CustomerPromotionalEntitlementRevokeParams{
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
