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

func TestV1SubscriptionGet(t *testing.T) {
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
	_, err := client.V1.Subscriptions.Get(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1SubscriptionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Subscriptions.Update(
		context.TODO(),
		"x",
		stigg.V1SubscriptionUpdateParams{
			Addons: []stigg.V1SubscriptionUpdateParamsAddon{{
				AddonID:  "addonId",
				Quantity: 0,
			}},
			AppliedCoupon: stigg.V1SubscriptionUpdateParamsAppliedCoupon{
				BillingCouponID: stigg.String("billingCouponId"),
				Configuration: stigg.V1SubscriptionUpdateParamsAppliedCouponConfiguration{
					StartDate: stigg.Time(time.Now()),
				},
				CouponID: stigg.String("couponId"),
				Discount: stigg.V1SubscriptionUpdateParamsAppliedCouponDiscount{
					AmountsOff: []stigg.V1SubscriptionUpdateParamsAppliedCouponDiscountAmountsOff{{
						Amount:   0,
						Currency: "usd",
					}},
					Description:      stigg.String("description"),
					DurationInMonths: stigg.Float(1),
					Name:             stigg.String("name"),
					PercentOff:       stigg.Float(1),
				},
				PromotionCode: stigg.String("promotionCode"),
			},
			AwaitPaymentConfirmation: stigg.Bool(true),
			BillingInformation: stigg.V1SubscriptionUpdateParamsBillingInformation{
				BillingAddress: stigg.V1SubscriptionUpdateParamsBillingInformationBillingAddress{
					City:       stigg.String("city"),
					Country:    stigg.String("country"),
					Line1:      stigg.String("line1"),
					Line2:      stigg.String("line2"),
					PostalCode: stigg.String("postalCode"),
					State:      stigg.String("state"),
				},
				ChargeOnBehalfOfAccount: stigg.String("chargeOnBehalfOfAccount"),
				CouponID:                stigg.String("couponId"),
				IntegrationID:           stigg.String("integrationId"),
				InvoiceDaysUntilDue:     stigg.Float(0),
				IsBackdated:             stigg.Bool(true),
				IsInvoicePaid:           stigg.Bool(true),
				Metadata: map[string]any{
					"foo": "bar",
				},
				ProrationBehavior: "INVOICE_IMMEDIATELY",
				TaxIDs: []stigg.V1SubscriptionUpdateParamsBillingInformationTaxID{{
					Type:  "type",
					Value: "value",
				}},
				TaxPercentage: stigg.Float(0),
				TaxRateIDs:    []string{"string"},
			},
			BillingPeriod: stigg.V1SubscriptionUpdateParamsBillingPeriodMonthly,
			Budget: stigg.V1SubscriptionUpdateParamsBudget{
				HasSoftLimit: true,
				Limit:        0,
			},
			Charges: []stigg.V1SubscriptionUpdateParamsCharge{{
				ID:       "id",
				Quantity: 1,
				Type:     "FEATURE",
			}},
			Metadata: map[string]string{
				"foo": "string",
			},
			MinimumSpend: stigg.V1SubscriptionUpdateParamsMinimumSpend{
				Minimum: stigg.V1SubscriptionUpdateParamsMinimumSpendMinimum{
					Amount:   0,
					Currency: "usd",
				},
			},
			PriceOverrides: []stigg.V1SubscriptionUpdateParamsPriceOverride{{
				AddonID:    stigg.String("addonId"),
				BaseCharge: stigg.Bool(true),
				CurrencyID: stigg.String("currencyId"),
				FeatureID:  stigg.String("featureId"),
				Price: stigg.V1SubscriptionUpdateParamsPriceOverridePrice{
					Amount:   0,
					Currency: "usd",
				},
			}},
			PromotionCode:    stigg.String("promotionCode"),
			ScheduleStrategy: stigg.V1SubscriptionUpdateParamsScheduleStrategyEndOfBillingPeriod,
			SubscriptionEntitlements: []stigg.V1SubscriptionUpdateParamsSubscriptionEntitlement{{
				ID:                stigg.String("id"),
				FeatureID:         stigg.String("featureId"),
				HasSoftLimit:      stigg.Bool(true),
				HasUnlimitedUsage: stigg.Bool(true),
				MonthlyResetPeriodConfiguration: stigg.V1SubscriptionUpdateParamsSubscriptionEntitlementMonthlyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				ResetPeriod: "YEAR",
				UsageLimit:  stigg.Float(0),
				WeeklyResetPeriodConfiguration: stigg.V1SubscriptionUpdateParamsSubscriptionEntitlementWeeklyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
				YearlyResetPeriodConfiguration: stigg.V1SubscriptionUpdateParamsSubscriptionEntitlementYearlyResetPeriodConfiguration{
					AccordingTo: "SubscriptionStart",
				},
			}},
			TrialEndDate: stigg.Time(time.Now()),
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

func TestV1SubscriptionListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Subscriptions.List(context.TODO(), stigg.V1SubscriptionListParams{
		After:  stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Before: stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CreatedAt: stigg.V1SubscriptionListParamsCreatedAt{
			Gt:  stigg.Time(time.Now()),
			Gte: stigg.Time(time.Now()),
			Lt:  stigg.Time(time.Now()),
			Lte: stigg.Time(time.Now()),
		},
		CustomerID:  stigg.String("customerId"),
		Limit:       stigg.Int(1),
		PlanID:      stigg.String("planId"),
		PricingType: stigg.String("pricingType"),
		ResourceID:  stigg.String("resourceId"),
		Status:      stigg.String("status"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1SubscriptionCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Subscriptions.Cancel(
		context.TODO(),
		"x",
		stigg.V1SubscriptionCancelParams{
			CancellationAction: stigg.V1SubscriptionCancelParamsCancellationActionDefault,
			CancellationTime:   stigg.V1SubscriptionCancelParamsCancellationTimeEndOfBillingPeriod,
			EndDate:            stigg.Time(time.Now()),
			Prorate:            stigg.Bool(true),
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

func TestV1SubscriptionDelegate(t *testing.T) {
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
	_, err := client.V1.Subscriptions.Delegate(
		context.TODO(),
		"x",
		stigg.V1SubscriptionDelegateParams{
			TargetCustomerID: "targetCustomerId",
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

func TestV1SubscriptionImportWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Subscriptions.Import(context.TODO(), stigg.V1SubscriptionImportParams{
		Subscriptions: []stigg.V1SubscriptionImportParamsSubscription{{
			ID:         "id",
			CustomerID: "customerId",
			PlanID:     "planId",
			BillingID:  stigg.String("billingId"),
			EndDate:    stigg.Time(time.Now()),
			Metadata: map[string]string{
				"foo": "string",
			},
			ResourceID: stigg.String("resourceId"),
			StartDate:  stigg.Time(time.Now()),
		}},
		IntegrationID: stigg.String("integrationId"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1SubscriptionMigrateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Subscriptions.Migrate(
		context.TODO(),
		"x",
		stigg.V1SubscriptionMigrateParams{
			SubscriptionMigrationTime: stigg.V1SubscriptionMigrateParamsSubscriptionMigrationTimeEndOfBillingPeriod,
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

func TestV1SubscriptionPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Subscriptions.Preview(context.TODO(), stigg.V1SubscriptionPreviewParams{
		CustomerID: "customerId",
		PlanID:     "planId",
		Addons: []stigg.V1SubscriptionPreviewParamsAddon{{
			AddonID:  "addonId",
			Quantity: stigg.Int(1),
		}},
		AppliedCoupon: stigg.V1SubscriptionPreviewParamsAppliedCoupon{
			BillingCouponID: stigg.String("billingCouponId"),
			Configuration: stigg.V1SubscriptionPreviewParamsAppliedCouponConfiguration{
				StartDate: stigg.Time(time.Now()),
			},
			CouponID: stigg.String("couponId"),
			Discount: stigg.V1SubscriptionPreviewParamsAppliedCouponDiscount{
				AmountsOff: []stigg.V1SubscriptionPreviewParamsAppliedCouponDiscountAmountsOff{{
					Amount:   0,
					Currency: "usd",
				}},
				Description:      stigg.String("description"),
				DurationInMonths: stigg.Float(1),
				Name:             stigg.String("name"),
				PercentOff:       stigg.Float(1),
			},
			PromotionCode: stigg.String("promotionCode"),
		},
		BillableFeatures: []stigg.V1SubscriptionPreviewParamsBillableFeature{{
			FeatureID: "featureId",
			Quantity:  1,
		}},
		BillingCountryCode: stigg.String("billingCountryCode"),
		BillingInformation: stigg.V1SubscriptionPreviewParamsBillingInformation{
			BillingAddress: stigg.V1SubscriptionPreviewParamsBillingInformationBillingAddress{
				City:       stigg.String("city"),
				Country:    stigg.String("country"),
				Line1:      stigg.String("line1"),
				Line2:      stigg.String("line2"),
				PostalCode: stigg.String("postalCode"),
				State:      stigg.String("state"),
			},
			ChargeOnBehalfOfAccount: stigg.String("chargeOnBehalfOfAccount"),
			IntegrationID:           stigg.String("integrationId"),
			InvoiceDaysUntilDue:     stigg.Float(0),
			IsBackdated:             stigg.Bool(true),
			IsInvoicePaid:           stigg.Bool(true),
			Metadata:                map[string]any{},
			ProrationBehavior:       "INVOICE_IMMEDIATELY",
			TaxIDs: []stigg.V1SubscriptionPreviewParamsBillingInformationTaxID{{
				Type:  "type",
				Value: "value",
			}},
			TaxPercentage: stigg.Float(0),
			TaxRateIDs:    []string{"string"},
		},
		BillingPeriod: stigg.V1SubscriptionPreviewParamsBillingPeriodMonthly,
		Charges: []stigg.V1SubscriptionPreviewParamsCharge{{
			ID:       "id",
			Quantity: 1,
			Type:     "FEATURE",
		}},
		PayingCustomerID: stigg.String("payingCustomerId"),
		ResourceID:       stigg.String("resourceId"),
		ScheduleStrategy: stigg.V1SubscriptionPreviewParamsScheduleStrategyEndOfBillingPeriod,
		StartDate:        stigg.Time(time.Now()),
		TrialOverrideConfiguration: stigg.V1SubscriptionPreviewParamsTrialOverrideConfiguration{
			IsTrial:          true,
			TrialEndBehavior: "CONVERT_TO_PAID",
			TrialEndDate:     stigg.Time(time.Now()),
		},
		UnitQuantity: stigg.Float(1),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1SubscriptionProvisionWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Subscriptions.Provision(context.TODO(), stigg.V1SubscriptionProvisionParams{
		CustomerID: "customerId",
		PlanID:     "planId",
		ID:         stigg.String("id"),
		Addons: []stigg.V1SubscriptionProvisionParamsAddon{{
			AddonID:  "addonId",
			Quantity: stigg.Int(1),
		}},
		AppliedCoupon: stigg.V1SubscriptionProvisionParamsAppliedCoupon{
			BillingCouponID: stigg.String("billingCouponId"),
			Configuration: stigg.V1SubscriptionProvisionParamsAppliedCouponConfiguration{
				StartDate: stigg.Time(time.Now()),
			},
			CouponID: stigg.String("couponId"),
			Discount: stigg.V1SubscriptionProvisionParamsAppliedCouponDiscount{
				AmountsOff: []stigg.V1SubscriptionProvisionParamsAppliedCouponDiscountAmountsOff{{
					Amount:   0,
					Currency: "usd",
				}},
				Description:      stigg.String("description"),
				DurationInMonths: stigg.Float(1),
				Name:             stigg.String("name"),
				PercentOff:       stigg.Float(1),
			},
			PromotionCode: stigg.String("promotionCode"),
		},
		AwaitPaymentConfirmation: stigg.Bool(true),
		BillingCountryCode:       stigg.String("billingCountryCode"),
		BillingID:                stigg.String("billingId"),
		BillingInformation: stigg.V1SubscriptionProvisionParamsBillingInformation{
			BillingAddress: stigg.V1SubscriptionProvisionParamsBillingInformationBillingAddress{
				City:       stigg.String("city"),
				Country:    stigg.String("country"),
				Line1:      stigg.String("line1"),
				Line2:      stigg.String("line2"),
				PostalCode: stigg.String("postalCode"),
				State:      stigg.String("state"),
			},
			ChargeOnBehalfOfAccount: stigg.String("chargeOnBehalfOfAccount"),
			IntegrationID:           stigg.String("integrationId"),
			InvoiceDaysUntilDue:     stigg.Float(0),
			IsBackdated:             stigg.Bool(true),
			IsInvoicePaid:           stigg.Bool(true),
			Metadata: map[string]string{
				"foo": "string",
			},
			ProrationBehavior: "INVOICE_IMMEDIATELY",
			TaxIDs: []stigg.V1SubscriptionProvisionParamsBillingInformationTaxID{{
				Type:  "type",
				Value: "value",
			}},
			TaxPercentage: stigg.Float(0),
			TaxRateIDs:    []string{"string"},
		},
		BillingPeriod: stigg.V1SubscriptionProvisionParamsBillingPeriodMonthly,
		Budget: stigg.V1SubscriptionProvisionParamsBudget{
			HasSoftLimit: true,
			Limit:        0,
		},
		Charges: []stigg.V1SubscriptionProvisionParamsCharge{{
			ID:       "id",
			Quantity: 1,
			Type:     "FEATURE",
		}},
		CheckoutOptions: stigg.V1SubscriptionProvisionParamsCheckoutOptions{
			CancelURL:             "https://example.com",
			SuccessURL:            "https://example.com",
			AllowPromoCodes:       stigg.Bool(true),
			AllowTaxIDCollection:  stigg.Bool(true),
			CollectBillingAddress: stigg.Bool(true),
			CollectPhoneNumber:    stigg.Bool(true),
			ReferenceID:           stigg.String("referenceId"),
		},
		Metadata: map[string]string{
			"foo": "string",
		},
		MinimumSpend: stigg.V1SubscriptionProvisionParamsMinimumSpend{
			Minimum: stigg.V1SubscriptionProvisionParamsMinimumSpendMinimum{
				Amount:             stigg.Float(0),
				BillingCountryCode: stigg.String("billingCountryCode"),
				Currency:           "usd",
			},
		},
		PayingCustomerID:        stigg.String("payingCustomerId"),
		PaymentCollectionMethod: stigg.V1SubscriptionProvisionParamsPaymentCollectionMethodCharge,
		PriceOverrides: []stigg.V1SubscriptionProvisionParamsPriceOverride{{
			AddonID:            stigg.String("addonId"),
			BaseCharge:         stigg.Bool(true),
			BlockSize:          stigg.Float(0),
			CreditGrantCadence: "BEGINNING_OF_BILLING_PERIOD",
			CreditRate: stigg.V1SubscriptionProvisionParamsPriceOverrideCreditRate{
				Amount:      1,
				CurrencyID:  "currencyId",
				CostFormula: stigg.String("costFormula"),
			},
			FeatureID: stigg.String("featureId"),
			Price: stigg.V1SubscriptionProvisionParamsPriceOverridePrice{
				Amount:             stigg.Float(0),
				BillingCountryCode: stigg.String("billingCountryCode"),
				Currency:           "usd",
			},
			Tiers: []stigg.V1SubscriptionProvisionParamsPriceOverrideTier{{
				FlatPrice: stigg.V1SubscriptionProvisionParamsPriceOverrideTierFlatPrice{
					Amount:             stigg.Float(0),
					BillingCountryCode: stigg.String("billingCountryCode"),
					Currency:           "usd",
				},
				UnitPrice: stigg.V1SubscriptionProvisionParamsPriceOverrideTierUnitPrice{
					Amount:             stigg.Float(0),
					BillingCountryCode: stigg.String("billingCountryCode"),
					Currency:           "usd",
				},
				UpTo: stigg.Float(0),
			}},
		}},
		ResourceID:       stigg.String("resourceId"),
		SalesforceID:     stigg.String("salesforceId"),
		ScheduleStrategy: stigg.V1SubscriptionProvisionParamsScheduleStrategyEndOfBillingPeriod,
		StartDate:        stigg.Time(time.Now()),
		SubscriptionEntitlements: []stigg.V1SubscriptionProvisionParamsSubscriptionEntitlement{{
			FeatureID:  "featureId",
			UsageLimit: 0,
			IsGranted:  stigg.Bool(true),
		}},
		TrialOverrideConfiguration: stigg.V1SubscriptionProvisionParamsTrialOverrideConfiguration{
			IsTrial:          true,
			TrialEndBehavior: "CONVERT_TO_PAID",
			TrialEndDate:     stigg.Time(time.Now()),
		},
		UnitQuantity: stigg.Float(1),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1SubscriptionTransfer(t *testing.T) {
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
	_, err := client.V1.Subscriptions.Transfer(
		context.TODO(),
		"x",
		stigg.V1SubscriptionTransferParams{
			DestinationResourceID: "destinationResourceId",
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
