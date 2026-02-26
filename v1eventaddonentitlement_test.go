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

func TestV1EventAddonEntitlementNew(t *testing.T) {
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
	_, err := client.V1.Events.Addons.Entitlements.New(
		context.TODO(),
		"addonId",
		stigg.V1EventAddonEntitlementNewParams{
			Entitlements: []stigg.V1EventAddonEntitlementNewParamsEntitlement{{
				Credit: stigg.V1EventAddonEntitlementNewParamsEntitlementCredit{
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
				Feature: stigg.V1EventAddonEntitlementNewParamsEntitlementFeature{
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
					MonthlyResetPeriodConfiguration: stigg.V1EventAddonEntitlementNewParamsEntitlementFeatureMonthlyResetPeriodConfiguration{
						AccordingTo: "SubscriptionStart",
					},
					Order:       stigg.Float(0),
					ResetPeriod: "YEAR",
					UsageLimit:  stigg.Int(0),
					WeeklyResetPeriodConfiguration: stigg.V1EventAddonEntitlementNewParamsEntitlementFeatureWeeklyResetPeriodConfiguration{
						AccordingTo: "SubscriptionStart",
					},
					YearlyResetPeriodConfiguration: stigg.V1EventAddonEntitlementNewParamsEntitlementFeatureYearlyResetPeriodConfiguration{
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

func TestV1EventAddonEntitlementUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Events.Addons.Entitlements.Update(
		context.TODO(),
		"id",
		stigg.V1EventAddonEntitlementUpdateParams{
			AddonID: "addonId",
			Credit: stigg.V1EventAddonEntitlementUpdateParamsCredit{
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
			Feature: stigg.V1EventAddonEntitlementUpdateParamsFeature{
				Behavior:            "Increment",
				Description:         stigg.String("description"),
				DisplayNameOverride: stigg.String("displayNameOverride"),
				EnumValues:          []string{"string"},
				HasSoftLimit:        stigg.Bool(true),
				HasUnlimitedUsage:   stigg.Bool(true),
				HiddenFromWidgets:   []string{"PAYWALL"},
				IsCustom:            stigg.Bool(true),
				IsGranted:           stigg.Bool(true),
				MonthlyResetPeriodConfiguration: stigg.V1EventAddonEntitlementUpdateParamsFeatureMonthlyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				Order:       stigg.Float(0),
				ResetPeriod: "YEAR",
				UsageLimit:  stigg.Int(0),
				WeeklyResetPeriodConfiguration: stigg.V1EventAddonEntitlementUpdateParamsFeatureWeeklyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				YearlyResetPeriodConfiguration: stigg.V1EventAddonEntitlementUpdateParamsFeatureYearlyResetPeriodConfiguration{
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

func TestV1EventAddonEntitlementList(t *testing.T) {
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
	_, err := client.V1.Events.Addons.Entitlements.List(context.TODO(), "addonId")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1EventAddonEntitlementDelete(t *testing.T) {
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
	_, err := client.V1.Events.Addons.Entitlements.Delete(
		context.TODO(),
		"id",
		stigg.V1EventAddonEntitlementDeleteParams{
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
