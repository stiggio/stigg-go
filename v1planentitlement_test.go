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

func TestV1PlanEntitlementNew(t *testing.T) {
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
	_, err := client.V1.Plans.Entitlements.New(
		context.TODO(),
		"planId",
		stigg.V1PlanEntitlementNewParams{
			Entitlements: []stigg.V1PlanEntitlementNewParamsEntitlement{{
				Credit: stigg.V1PlanEntitlementNewParamsEntitlementCredit{
					Amount:              stigg.Float(1),
					Cadence:             "MONTH",
					CustomCurrencyID:    "customCurrencyId",
					Behavior:            "Increment",
					Description:         stigg.String("description"),
					DisplayNameOverride: stigg.String("displayNameOverride"),
					HiddenFromWidgets:   []string{"PAYWALL"},
					IsCustom:            stigg.Bool(true),
					IsGranted:           stigg.Bool(true),
					Order:               stigg.Float(0),
				},
				Feature: stigg.V1PlanEntitlementNewParamsEntitlementFeature{
					FeatureID:           "featureId",
					Behavior:            "Increment",
					Description:         stigg.String("description"),
					DisplayNameOverride: stigg.String("displayNameOverride"),
					EnumValues:          []string{"string"},
					HasSoftLimit:        stigg.Bool(true),
					HasUnlimitedUsage:   stigg.Bool(true),
					HiddenFromWidgets:   []string{"PAYWALL"},
					IsCustom:            stigg.Bool(true),
					IsGranted:           stigg.Bool(true),
					MonthlyResetPeriodConfiguration: stigg.V1PlanEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration{
						AccordingTo: "SubscriptionStart",
					},
					Order:       stigg.Float(0),
					ResetPeriod: "YEAR",
					UsageLimit:  stigg.Int(0),
					WeeklyResetPeriodConfiguration: stigg.V1PlanEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration{
						AccordingTo: "SubscriptionStart",
					},
					YearlyResetPeriodConfiguration: stigg.V1PlanEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration{
						AccordingTo: "SubscriptionStart",
					},
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

func TestV1PlanEntitlementUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Plans.Entitlements.Update(
		context.TODO(),
		"id",
		stigg.V1PlanEntitlementUpdateParams{
			PlanID: "planId",
			Credit: stigg.V1PlanEntitlementUpdateParamsCredit{
				Amount:              stigg.Float(1),
				Behavior:            "Increment",
				Cadence:             "MONTH",
				Description:         stigg.String("description"),
				DisplayNameOverride: stigg.String("displayNameOverride"),
				HiddenFromWidgets:   []string{"PAYWALL"},
				IsCustom:            stigg.Bool(true),
				IsGranted:           stigg.Bool(true),
				Order:               stigg.Float(0),
			},
			Feature: stigg.V1PlanEntitlementUpdateParamsFeature{
				Behavior:            "Increment",
				Description:         stigg.String("description"),
				DisplayNameOverride: stigg.String("displayNameOverride"),
				EnumValues:          []string{"string"},
				HasSoftLimit:        stigg.Bool(true),
				HasUnlimitedUsage:   stigg.Bool(true),
				HiddenFromWidgets:   []string{"PAYWALL"},
				IsCustom:            stigg.Bool(true),
				IsGranted:           stigg.Bool(true),
				MonthlyResetPeriodConfiguration: stigg.V1PlanEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				Order:       stigg.Float(0),
				ResetPeriod: "YEAR",
				UsageLimit:  stigg.Int(0),
				WeeklyResetPeriodConfiguration: stigg.V1PlanEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				YearlyResetPeriodConfiguration: stigg.V1PlanEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
			},
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

func TestV1PlanEntitlementList(t *testing.T) {
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
	_, err := client.V1.Plans.Entitlements.List(context.TODO(), "planId")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PlanEntitlementDelete(t *testing.T) {
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
	_, err := client.V1.Plans.Entitlements.Delete(
		context.TODO(),
		"id",
		stigg.V1PlanEntitlementDeleteParams{
			PlanID: "planId",
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
