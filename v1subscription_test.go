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

func TestV1SubscriptionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Subscriptions.New(context.TODO(), stigg.V1SubscriptionNewParams{
		CustomerID: "customerId",
		PlanID:     "planId",
		ID:         stigg.String("id"),
		Addons: []stigg.V1SubscriptionNewParamsAddon{{
			AddonID:  "addonId",
			Quantity: stigg.Int(1),
		}},
		AppliedCoupon: stigg.V1SubscriptionNewParamsAppliedCoupon{
			BillingCouponID: stigg.String("billingCouponId"),
			Configuration: stigg.V1SubscriptionNewParamsAppliedCouponConfiguration{
				StartDate: stigg.Time(time.Now()),
			},
			CouponID: stigg.String("couponId"),
			Discount: stigg.V1SubscriptionNewParamsAppliedCouponDiscount{
				AmountsOff: []stigg.V1SubscriptionNewParamsAppliedCouponDiscountAmountsOff{{
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
		BillingInformation: stigg.V1SubscriptionNewParamsBillingInformation{
			BillingAddress: stigg.V1SubscriptionNewParamsBillingInformationBillingAddress{
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
			TaxIDs: []stigg.V1SubscriptionNewParamsBillingInformationTaxID{{
				Type:  "type",
				Value: "value",
			}},
			TaxPercentage: stigg.Float(0),
			TaxRateIDs:    []string{"string"},
		},
		BillingPeriod: stigg.V1SubscriptionNewParamsBillingPeriodMonthly,
		Budget: stigg.V1SubscriptionNewParamsBudget{
			HasSoftLimit: true,
			Limit:        0,
		},
		Charges: []stigg.V1SubscriptionNewParamsCharge{{
			ID:       "id",
			Quantity: 1,
			Type:     "FEATURE",
		}},
		CheckoutOptions: stigg.V1SubscriptionNewParamsCheckoutOptions{
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
		MinimumSpend: stigg.V1SubscriptionNewParamsMinimumSpend{
			Minimum: stigg.V1SubscriptionNewParamsMinimumSpendMinimum{
				Amount:             stigg.Float(0),
				BillingCountryCode: stigg.String("billingCountryCode"),
				Currency:           "usd",
			},
		},
		PayingCustomerID:        stigg.String("payingCustomerId"),
		PaymentCollectionMethod: stigg.V1SubscriptionNewParamsPaymentCollectionMethodCharge,
		PriceOverrides: []stigg.V1SubscriptionNewParamsPriceOverride{{
			AddonID:            stigg.String("addonId"),
			BaseCharge:         stigg.Bool(true),
			BlockSize:          stigg.Float(0),
			CreditGrantCadence: "BEGINNING_OF_BILLING_PERIOD",
			CreditRate: stigg.V1SubscriptionNewParamsPriceOverrideCreditRate{
				Amount:      1,
				CurrencyID:  "currencyId",
				CostFormula: stigg.String("costFormula"),
			},
			FeatureID: stigg.String("featureId"),
			Price: stigg.V1SubscriptionNewParamsPriceOverridePrice{
				Amount:             stigg.Float(0),
				BillingCountryCode: stigg.String("billingCountryCode"),
				Currency:           "usd",
			},
			Tiers: []stigg.V1SubscriptionNewParamsPriceOverrideTier{{
				FlatPrice: stigg.V1SubscriptionNewParamsPriceOverrideTierFlatPrice{
					Amount:             stigg.Float(0),
					BillingCountryCode: stigg.String("billingCountryCode"),
					Currency:           "usd",
				},
				UnitPrice: stigg.V1SubscriptionNewParamsPriceOverrideTierUnitPrice{
					Amount:             stigg.Float(0),
					BillingCountryCode: stigg.String("billingCountryCode"),
					Currency:           "usd",
				},
				UpTo: stigg.Float(0),
			}},
		}},
		ResourceID:       stigg.String("resourceId"),
		SalesforceID:     stigg.String("salesforceId"),
		ScheduleStrategy: stigg.V1SubscriptionNewParamsScheduleStrategyEndOfBillingPeriod,
		StartDate:        stigg.Time(time.Now()),
		SubscriptionEntitlements: []stigg.V1SubscriptionNewParamsSubscriptionEntitlement{{
			FeatureID:  "featureId",
			UsageLimit: 0,
			IsGranted:  stigg.Bool(true),
		}},
		TrialOverrideConfiguration: stigg.V1SubscriptionNewParamsTrialOverrideConfiguration{
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

func TestV1SubscriptionGet(t *testing.T) {
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
	_, err := client.V1.Subscriptions.Get(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1SubscriptionListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Subscriptions.List(context.TODO(), stigg.V1SubscriptionListParams{
		After:      stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Before:     stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CustomerID: stigg.String("customerId"),
		Limit:      stigg.Int(1),
		Status:     stigg.String("status"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1SubscriptionDelegate(t *testing.T) {
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

func TestV1SubscriptionMigrateWithOptionalParams(t *testing.T) {
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

func TestV1SubscriptionTransfer(t *testing.T) {
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
