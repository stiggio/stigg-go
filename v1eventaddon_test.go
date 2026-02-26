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

func TestV1EventAddonArchiveAddon(t *testing.T) {
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
	_, err := client.V1.Events.Addons.ArchiveAddon(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1EventAddonNewAddonWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Events.Addons.NewAddon(context.TODO(), stigg.V1EventAddonNewAddonParams{
		ID:          "id",
		DisplayName: "displayName",
		ProductID:   "productId",
		BillingID:   stigg.String("billingId"),
		Description: stigg.String("description"),
		MaxQuantity: stigg.Int(0),
		Metadata: map[string]string{
			"foo": "string",
		},
		PricingType: stigg.V1EventAddonNewAddonParamsPricingTypeFree,
		Status:      stigg.V1EventAddonNewAddonParamsStatusDraft,
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1EventAddonListAddonsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Events.Addons.ListAddons(context.TODO(), stigg.V1EventAddonListAddonsParams{
		After:  stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Before: stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CreatedAt: stigg.V1EventAddonListAddonsParamsCreatedAt{
			Gt:  stigg.Time(time.Now()),
			Gte: stigg.Time(time.Now()),
			Lt:  stigg.Time(time.Now()),
			Lte: stigg.Time(time.Now()),
		},
		Limit:     stigg.Int(1),
		ProductID: stigg.String("productId"),
		Status:    stigg.String("status"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1EventAddonPublishAddon(t *testing.T) {
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
	_, err := client.V1.Events.Addons.PublishAddon(
		context.TODO(),
		"x",
		stigg.V1EventAddonPublishAddonParams{
			MigrationType: stigg.V1EventAddonPublishAddonParamsMigrationTypeNewCustomers,
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

func TestV1EventAddonGetAddon(t *testing.T) {
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
	_, err := client.V1.Events.Addons.GetAddon(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1EventAddonSetPricingWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Events.Addons.SetPricing(
		context.TODO(),
		"x",
		stigg.V1EventAddonSetPricingParams{
			SetPackagePricing: stigg.SetPackagePricingParam{
				PricingType: stigg.SetPackagePricingPricingTypeFree,
				BillingID:   stigg.String("billingId"),
				MinimumSpend: []stigg.SetPackagePricingMinimumSpendParam{{
					BillingPeriod: "MONTHLY",
					Minimum: stigg.SetPackagePricingMinimumSpendMinimumParam{
						Amount:   0,
						Currency: "usd",
					},
				}},
				OverageBillingPeriod: stigg.SetPackagePricingOverageBillingPeriodOnSubscriptionRenewal,
				OveragePricingModels: []stigg.SetPackagePricingOveragePricingModelParam{{
					BillingModel: "FLAT_FEE",
					PricePeriods: []stigg.SetPackagePricingOveragePricingModelPricePeriodParam{{
						BillingPeriod:      "MONTHLY",
						BillingCountryCode: stigg.String("billingCountryCode"),
						BlockSize:          stigg.Float(0),
						CreditGrantCadence: "BEGINNING_OF_BILLING_PERIOD",
						CreditRate: stigg.SetPackagePricingOveragePricingModelPricePeriodCreditRateParam{
							Amount:      1,
							CurrencyID:  "currencyId",
							CostFormula: stigg.String("costFormula"),
						},
						Price: stigg.SetPackagePricingOveragePricingModelPricePeriodPriceParam{
							Amount:   0,
							Currency: "usd",
						},
						Tiers: []stigg.SetPackagePricingOveragePricingModelPricePeriodTierParam{{
							FlatPrice: stigg.SetPackagePricingOveragePricingModelPricePeriodTierFlatPriceParam{
								Amount:   0,
								Currency: "usd",
							},
							UnitPrice: stigg.SetPackagePricingOveragePricingModelPricePeriodTierUnitPriceParam{
								Amount:   0,
								Currency: "usd",
							},
							UpTo: stigg.Float(0),
						}},
					}},
					BillingCadence: "RECURRING",
					Entitlement: stigg.SetPackagePricingOveragePricingModelEntitlementParam{
						FeatureID:         "featureId",
						HasSoftLimit:      stigg.Bool(true),
						HasUnlimitedUsage: stigg.Bool(true),
						MonthlyResetPeriodConfiguration: stigg.SetPackagePricingOveragePricingModelEntitlementMonthlyResetPeriodConfigurationParam{
							AccordingTo: "SubscriptionStart",
						},
						ResetPeriod: "YEAR",
						UsageLimit:  stigg.Float(0),
						WeeklyResetPeriodConfiguration: stigg.SetPackagePricingOveragePricingModelEntitlementWeeklyResetPeriodConfigurationParam{
							AccordingTo: "SubscriptionStart",
						},
						YearlyResetPeriodConfiguration: stigg.SetPackagePricingOveragePricingModelEntitlementYearlyResetPeriodConfigurationParam{
							AccordingTo: "SubscriptionStart",
						},
					},
					FeatureID:             stigg.String("featureId"),
					TopUpCustomCurrencyID: stigg.String("topUpCustomCurrencyId"),
				}},
				PricingModels: []stigg.SetPackagePricingPricingModelParam{{
					BillingModel: "FLAT_FEE",
					PricePeriods: []stigg.SetPackagePricingPricingModelPricePeriodParam{{
						BillingPeriod:      "MONTHLY",
						BillingCountryCode: stigg.String("billingCountryCode"),
						BlockSize:          stigg.Float(0),
						CreditGrantCadence: "BEGINNING_OF_BILLING_PERIOD",
						CreditRate: stigg.SetPackagePricingPricingModelPricePeriodCreditRateParam{
							Amount:      1,
							CurrencyID:  "currencyId",
							CostFormula: stigg.String("costFormula"),
						},
						Price: stigg.SetPackagePricingPricingModelPricePeriodPriceParam{
							Amount:   0,
							Currency: "usd",
						},
						Tiers: []stigg.SetPackagePricingPricingModelPricePeriodTierParam{{
							FlatPrice: stigg.SetPackagePricingPricingModelPricePeriodTierFlatPriceParam{
								Amount:   0,
								Currency: "usd",
							},
							UnitPrice: stigg.SetPackagePricingPricingModelPricePeriodTierUnitPriceParam{
								Amount:   0,
								Currency: "usd",
							},
							UpTo: stigg.Float(0),
						}},
					}},
					BillingCadence:  "RECURRING",
					FeatureID:       stigg.String("featureId"),
					MaxUnitQuantity: stigg.Int(1),
					MinUnitQuantity: stigg.Int(1),
					MonthlyResetPeriodConfiguration: stigg.SetPackagePricingPricingModelMonthlyResetPeriodConfigurationParam{
						AccordingTo: "SubscriptionStart",
					},
					ResetPeriod:           "YEAR",
					TiersMode:             "VOLUME",
					TopUpCustomCurrencyID: stigg.String("topUpCustomCurrencyId"),
					WeeklyResetPeriodConfiguration: stigg.SetPackagePricingPricingModelWeeklyResetPeriodConfigurationParam{
						AccordingTo: "SubscriptionStart",
					},
					YearlyResetPeriodConfiguration: stigg.SetPackagePricingPricingModelYearlyResetPeriodConfigurationParam{
						AccordingTo: "SubscriptionStart",
					},
				}},
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

func TestV1EventAddonUpdateAddonWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Events.Addons.UpdateAddon(
		context.TODO(),
		"x",
		stigg.V1EventAddonUpdateAddonParams{
			BillingID:    stigg.String("billingId"),
			Dependencies: []string{"string"},
			Description:  stigg.String("description"),
			DisplayName:  stigg.String("displayName"),
			MaxQuantity:  stigg.Int(0),
			Metadata: map[string]string{
				"foo": "string",
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
