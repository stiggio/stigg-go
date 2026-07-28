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

func TestV1ContractNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.New(context.TODO(), stigg.V1ContractNewParams{
		CustomerID: "customerId",
		Subscriptions: []stigg.V1ContractNewParamsSubscription{{
			ExistingSubscriptionID: stigg.String("existingSubscriptionId"),
			NewSubscription: stigg.V1ContractNewParamsSubscriptionNewSubscription{
				CustomerID: "customerId",
				PlanID:     "planId",
				ID:         stigg.String("id"),
				Addons: []stigg.V1ContractNewParamsSubscriptionNewSubscriptionAddon{{
					ID:       "id",
					Quantity: 0,
				}},
				AppliedCoupon: stigg.V1ContractNewParamsSubscriptionNewSubscriptionAppliedCoupon{
					BillingCouponID: stigg.String("billingCouponId"),
					Configuration: stigg.V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponConfiguration{
						StartDate: stigg.Time(time.Now()),
					},
					CouponID: stigg.String("couponId"),
					Discount: stigg.V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscount{
						AmountsOff: []stigg.V1ContractNewParamsSubscriptionNewSubscriptionAppliedCouponDiscountAmountsOff{{
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
				BillingCycleAnchor:       "UNCHANGED",
				BillingID:                stigg.String("billingId"),
				BillingInformation: stigg.V1ContractNewParamsSubscriptionNewSubscriptionBillingInformation{
					BillingAddress: stigg.V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationBillingAddress{
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
					TaxIDs: []stigg.V1ContractNewParamsSubscriptionNewSubscriptionBillingInformationTaxID{{
						Type:  "type",
						Value: "value",
					}},
					TaxPercentage: stigg.Float(0),
					TaxRateIDs:    []string{"string"},
				},
				BillingPeriod: "MONTHLY",
				Budget: stigg.V1ContractNewParamsSubscriptionNewSubscriptionBudget{
					HasSoftLimit: true,
					Limit:        0,
				},
				CancellationDate: stigg.Time(time.Now()),
				Charges: []stigg.V1ContractNewParamsSubscriptionNewSubscriptionCharge{{
					ID:       "id",
					Quantity: 0,
					Type:     "FEATURE",
				}},
				CheckoutOptions: stigg.V1ContractNewParamsSubscriptionNewSubscriptionCheckoutOptions{
					CancelURL:             "https://example.com",
					SuccessURL:            "https://example.com",
					AllowPromoCodes:       stigg.Bool(true),
					AllowTaxIDCollection:  stigg.Bool(true),
					CollectBillingAddress: stigg.Bool(true),
					CollectPhoneNumber:    stigg.Bool(true),
					ReferenceID:           stigg.String("referenceId"),
				},
				Entitlements: []stigg.V1ContractNewParamsSubscriptionNewSubscriptionEntitlementUnion{{
					OfFeature: &stigg.V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeature{
						ID:                "id",
						HasSoftLimit:      stigg.Bool(true),
						HasUnlimitedUsage: stigg.Bool(true),
						MonthlyResetPeriodConfiguration: stigg.V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureMonthlyResetPeriodConfiguration{
							AccordingTo: "SubscriptionStart",
						},
						ResetPeriod: "YEAR",
						UsageLimit:  stigg.Int(0),
						WeeklyResetPeriodConfiguration: stigg.V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureWeeklyResetPeriodConfiguration{
							AccordingTo: "SubscriptionStart",
						},
						YearlyResetPeriodConfiguration: stigg.V1ContractNewParamsSubscriptionNewSubscriptionEntitlementFeatureYearlyResetPeriodConfiguration{
							AccordingTo: "SubscriptionStart",
						},
					},
				}},
				Metadata: map[string]string{
					"foo": "string",
				},
				MinimumSpend: stigg.V1ContractNewParamsSubscriptionNewSubscriptionMinimumSpend{
					Amount:   stigg.Float(0),
					Currency: "usd",
				},
				PayingCustomerID:        stigg.String("payingCustomerId"),
				PaymentCollectionMethod: "CHARGE",
				PriceOverrides: []stigg.V1ContractNewParamsSubscriptionNewSubscriptionPriceOverride{{
					AddonID:            stigg.String("addonId"),
					Amount:             stigg.Float(0),
					BaseCharge:         stigg.Bool(true),
					BillingCountryCode: stigg.String("billingCountryCode"),
					BlockSize:          stigg.Float(0),
					CreditGrantCadence: "BEGINNING_OF_BILLING_PERIOD",
					CreditRate: stigg.V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideCreditRate{
						Amount:      1,
						CurrencyID:  "currencyId",
						CostFormula: stigg.String("costFormula"),
					},
					Currency:  "usd",
					FeatureID: stigg.String("featureId"),
					Tiers: []stigg.V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTier{{
						FlatPrice: stigg.V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierFlatPrice{
							Amount:   0,
							Currency: "usd",
						},
						UnitPrice: stigg.V1ContractNewParamsSubscriptionNewSubscriptionPriceOverrideTierUnitPrice{
							Amount:   0,
							Currency: "usd",
						},
						UpTo: stigg.Float(0),
					}},
				}},
				ResourceID:       stigg.String("resourceId"),
				SalesforceID:     stigg.String("salesforceId"),
				ScheduleStrategy: "END_OF_BILLING_PERIOD",
				StartDate:        stigg.Time(time.Now()),
				TrialOverrideConfiguration: stigg.V1ContractNewParamsSubscriptionNewSubscriptionTrialOverrideConfiguration{
					IsTrial:          true,
					TrialEndBehavior: "CONVERT_TO_PAID",
					TrialEndDate:     stigg.Time(time.Now()),
				},
				UnitQuantity: stigg.Int(0),
			},
		}},
		ActivationEndDate:   stigg.Time(time.Now()),
		ActivationStartDate: stigg.Time(time.Now()),
		Name:                stigg.String("name"),
		PoNumber:            stigg.String("poNumber"),
		SetupBilling:        stigg.Bool(true),
		XAccountID:          stigg.String("X-ACCOUNT-ID"),
		XEnvironmentID:      stigg.String("X-ENVIRONMENT-ID"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContractGetWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.Get(
		context.TODO(),
		"x",
		stigg.V1ContractGetParams{
			XAccountID:     stigg.String("X-ACCOUNT-ID"),
			XEnvironmentID: stigg.String("X-ENVIRONMENT-ID"),
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

func TestV1ContractUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.Update(
		context.TODO(),
		"x",
		stigg.V1ContractUpdateParams{
			ActivationEndDate:   stigg.Time(time.Now()),
			ActivationStartDate: stigg.Time(time.Now()),
			Name:                stigg.String("name"),
			PoNumber:            stigg.String("poNumber"),
			SetupBilling:        stigg.Bool(true),
			SubscriptionIDs:     []string{"NxI"},
			XAccountID:          stigg.String("X-ACCOUNT-ID"),
			XEnvironmentID:      stigg.String("X-ENVIRONMENT-ID"),
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

func TestV1ContractListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.List(context.TODO(), stigg.V1ContractListParams{
		After:              stigg.String("after"),
		Before:             stigg.String("before"),
		CustomerExternalID: stigg.String("customerExternalId"),
		Limit:              stigg.Int(1),
		Name:               stigg.String("name"),
		State:              stigg.String("state"),
		XAccountID:         stigg.String("X-ACCOUNT-ID"),
		XEnvironmentID:     stigg.String("X-ENVIRONMENT-ID"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContractDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Contracts.Delete(
		context.TODO(),
		"x",
		stigg.V1ContractDeleteParams{
			XAccountID:     stigg.String("X-ACCOUNT-ID"),
			XEnvironmentID: stigg.String("X-ENVIRONMENT-ID"),
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
