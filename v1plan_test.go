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

func TestV1PlanNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Plans.New(context.TODO(), stigg.V1PlanNewParams{
		ID:          "id",
		DisplayName: "displayName",
		ProductID:   "productId",
		BillingID:   stigg.String("billingId"),
		DefaultTrialConfig: stigg.V1PlanNewParamsDefaultTrialConfig{
			Duration: 0,
			Units:    "DAY",
			Budget: stigg.V1PlanNewParamsDefaultTrialConfigBudget{
				HasSoftLimit: true,
				Limit:        0,
			},
			TrialEndBehavior: "CONVERT_TO_PAID",
		},
		Description: stigg.String("description"),
		Metadata: map[string]string{
			"foo": "string",
		},
		ParentPlanID: stigg.String("parentPlanId"),
		PricingType:  stigg.V1PlanNewParamsPricingTypeFree,
		Status:       stigg.V1PlanNewParamsStatusDraft,
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PlanGet(t *testing.T) {
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
	_, err := client.V1.Plans.Get(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PlanUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Plans.Update(
		context.TODO(),
		"x",
		stigg.V1PlanUpdateParams{
			BillingID:          stigg.String("billingId"),
			CompatibleAddonIDs: []string{"string"},
			DefaultTrialConfig: stigg.V1PlanUpdateParamsDefaultTrialConfig{
				Duration: 0,
				Units:    "DAY",
				Budget: stigg.V1PlanUpdateParamsDefaultTrialConfigBudget{
					HasSoftLimit: true,
					Limit:        0,
				},
				TrialEndBehavior: "CONVERT_TO_PAID",
			},
			Description: stigg.String("description"),
			DisplayName: stigg.String("displayName"),
			Metadata: map[string]string{
				"foo": "string",
			},
			ParentPlanID: stigg.String("parentPlanId"),
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

func TestV1PlanListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Plans.List(context.TODO(), stigg.V1PlanListParams{
		After:  stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Before: stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CreatedAt: stigg.V1PlanListParamsCreatedAt{
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

func TestV1PlanArchive(t *testing.T) {
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
	_, err := client.V1.Plans.Archive(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PlanNewDraft(t *testing.T) {
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
	_, err := client.V1.Plans.NewDraft(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PlanPublish(t *testing.T) {
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
	_, err := client.V1.Plans.Publish(
		context.TODO(),
		"x",
		stigg.V1PlanPublishParams{
			MigrationType: stigg.V1PlanPublishParamsMigrationTypeNewCustomers,
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

func TestV1PlanRemoveDraft(t *testing.T) {
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
	_, err := client.V1.Plans.RemoveDraft(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PlanSetPricingWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Plans.SetPricing(
		context.TODO(),
		"x",
		stigg.V1PlanSetPricingParams{
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
