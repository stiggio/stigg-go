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

func TestV1AddonEntitlementNew(t *testing.T) {
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
	_, err := client.V1.Addons.Entitlements.New(
		context.TODO(),
		"addonId",
		stigg.V1AddonEntitlementNewParams{
			Entitlements: []stigg.V1AddonEntitlementNewParamsEntitlement{{
				Credit: stigg.V1AddonEntitlementNewParamsEntitlementCredit{
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
				Feature: stigg.V1AddonEntitlementNewParamsEntitlementFeature{
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
					MonthlyResetPeriodConfiguration: stigg.V1AddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration{
						AccordingTo: "SubscriptionStart",
					},
					Order:       stigg.Float(0),
					ResetPeriod: "YEAR",
					UsageLimit:  stigg.Int(0),
					WeeklyResetPeriodConfiguration: stigg.V1AddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration{
						AccordingTo: "SubscriptionStart",
					},
					YearlyResetPeriodConfiguration: stigg.V1AddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration{
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

func TestV1AddonEntitlementUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Addons.Entitlements.Update(
		context.TODO(),
		"id",
		stigg.V1AddonEntitlementUpdateParams{
			AddonID: "addonId",
			Credit: stigg.V1AddonEntitlementUpdateParamsCredit{
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
			Feature: stigg.V1AddonEntitlementUpdateParamsFeature{
				Behavior:            "Increment",
				Description:         stigg.String("description"),
				DisplayNameOverride: stigg.String("displayNameOverride"),
				EnumValues:          []string{"string"},
				HasSoftLimit:        stigg.Bool(true),
				HasUnlimitedUsage:   stigg.Bool(true),
				HiddenFromWidgets:   []string{"PAYWALL"},
				IsCustom:            stigg.Bool(true),
				IsGranted:           stigg.Bool(true),
				MonthlyResetPeriodConfiguration: stigg.V1AddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				Order:       stigg.Float(0),
				ResetPeriod: "YEAR",
				UsageLimit:  stigg.Int(0),
				WeeklyResetPeriodConfiguration: stigg.V1AddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				YearlyResetPeriodConfiguration: stigg.V1AddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration{
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

func TestV1AddonEntitlementList(t *testing.T) {
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
	_, err := client.V1.Addons.Entitlements.List(context.TODO(), "addonId")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1AddonEntitlementDelete(t *testing.T) {
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
	_, err := client.V1.Addons.Entitlements.Delete(
		context.TODO(),
		"id",
		stigg.V1AddonEntitlementDeleteParams{
			AddonID: "addonId",
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
