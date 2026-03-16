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
			BillingID: stigg.String("billingId"),
			Charges: stigg.V1PlanUpdateParamsCharges{
				PricingType: "FREE",
				BillingID:   stigg.String("billingId"),
				MinimumSpend: []stigg.V1PlanUpdateParamsChargesMinimumSpend{{
					BillingPeriod: "MONTHLY",
					Minimum: stigg.V1PlanUpdateParamsChargesMinimumSpendMinimum{
						Amount:   0,
						Currency: "usd",
					},
				}},
				OverageBillingPeriod: "ON_SUBSCRIPTION_RENEWAL",
				OveragePricingModels: []stigg.V1PlanUpdateParamsChargesOveragePricingModel{{
					BillingModel: "FLAT_FEE",
					PricePeriods: []stigg.V1PlanUpdateParamsChargesOveragePricingModelPricePeriod{{
						BillingPeriod:      "MONTHLY",
						BillingCountryCode: stigg.String("billingCountryCode"),
						BlockSize:          stigg.Float(0),
						CreditGrantCadence: "BEGINNING_OF_BILLING_PERIOD",
						CreditRate: stigg.V1PlanUpdateParamsChargesOveragePricingModelPricePeriodCreditRate{
							Amount:      1,
							CurrencyID:  "currencyId",
							CostFormula: stigg.String("costFormula"),
						},
						Price: stigg.V1PlanUpdateParamsChargesOveragePricingModelPricePeriodPrice{
							Amount:   0,
							Currency: "usd",
						},
						Tiers: []stigg.V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTier{{
							FlatPrice: stigg.V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierFlatPrice{
								Amount:   0,
								Currency: "usd",
							},
							UnitPrice: stigg.V1PlanUpdateParamsChargesOveragePricingModelPricePeriodTierUnitPrice{
								Amount:   0,
								Currency: "usd",
							},
							UpTo: stigg.Float(0),
						}},
					}},
					BillingCadence: "RECURRING",
					Entitlement: stigg.V1PlanUpdateParamsChargesOveragePricingModelEntitlement{
						FeatureID:         "featureId",
						HasSoftLimit:      stigg.Bool(true),
						HasUnlimitedUsage: stigg.Bool(true),
						MonthlyResetPeriodConfiguration: stigg.V1PlanUpdateParamsChargesOveragePricingModelEntitlementMonthlyResetPeriodConfiguration{
							AccordingTo: "SubscriptionStart",
						},
						ResetPeriod: "YEAR",
						UsageLimit:  stigg.Float(0),
						WeeklyResetPeriodConfiguration: stigg.V1PlanUpdateParamsChargesOveragePricingModelEntitlementWeeklyResetPeriodConfiguration{
							AccordingTo: "SubscriptionStart",
						},
						YearlyResetPeriodConfiguration: stigg.V1PlanUpdateParamsChargesOveragePricingModelEntitlementYearlyResetPeriodConfiguration{
							AccordingTo: "SubscriptionStart",
						},
					},
					FeatureID:             stigg.String("featureId"),
					TopUpCustomCurrencyID: stigg.String("topUpCustomCurrencyId"),
				}},
				PricingModels: []stigg.V1PlanUpdateParamsChargesPricingModel{{
					BillingModel: "FLAT_FEE",
					PricePeriods: []stigg.V1PlanUpdateParamsChargesPricingModelPricePeriod{{
						BillingPeriod:      "MONTHLY",
						BillingCountryCode: stigg.String("billingCountryCode"),
						BlockSize:          stigg.Float(0),
						CreditGrantCadence: "BEGINNING_OF_BILLING_PERIOD",
						CreditRate: stigg.V1PlanUpdateParamsChargesPricingModelPricePeriodCreditRate{
							Amount:      1,
							CurrencyID:  "currencyId",
							CostFormula: stigg.String("costFormula"),
						},
						Price: stigg.V1PlanUpdateParamsChargesPricingModelPricePeriodPrice{
							Amount:   0,
							Currency: "usd",
						},
						Tiers: []stigg.V1PlanUpdateParamsChargesPricingModelPricePeriodTier{{
							FlatPrice: stigg.V1PlanUpdateParamsChargesPricingModelPricePeriodTierFlatPrice{
								Amount:   0,
								Currency: "usd",
							},
							UnitPrice: stigg.V1PlanUpdateParamsChargesPricingModelPricePeriodTierUnitPrice{
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
					MonthlyResetPeriodConfiguration: stigg.V1PlanUpdateParamsChargesPricingModelMonthlyResetPeriodConfiguration{
						AccordingTo: "SubscriptionStart",
					},
					ResetPeriod:           "YEAR",
					TiersMode:             "VOLUME",
					TopUpCustomCurrencyID: stigg.String("topUpCustomCurrencyId"),
					WeeklyResetPeriodConfiguration: stigg.V1PlanUpdateParamsChargesPricingModelWeeklyResetPeriodConfiguration{
						AccordingTo: "SubscriptionStart",
					},
					YearlyResetPeriodConfiguration: stigg.V1PlanUpdateParamsChargesPricingModelYearlyResetPeriodConfiguration{
						AccordingTo: "SubscriptionStart",
					},
				}},
			},
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
