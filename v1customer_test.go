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

func TestV1CustomerGet(t *testing.T) {
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
	_, err := client.V1.Customers.Get(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Update(
		context.TODO(),
		"x",
		stigg.V1CustomerUpdateParams{
			BillingCurrency: stigg.V1CustomerUpdateParamsBillingCurrencyUsd,
			BillingID:       stigg.String("billingId"),
			CouponID:        stigg.V1CustomerUpdateParamsCouponIDEmpty,
			Email:           stigg.String("dev@stainless.com"),
			Integrations: []stigg.V1CustomerUpdateParamsIntegration{{
				ID:               "id",
				SyncedEntityID:   stigg.String("syncedEntityId"),
				VendorIdentifier: "AUTH0",
			}},
			Language: stigg.String("language"),
			Metadata: map[string]string{
				"foo": "string",
			},
			Name: stigg.String("name"),
			Passthrough: stigg.V1CustomerUpdateParamsPassthrough{
				Stripe: stigg.V1CustomerUpdateParamsPassthroughStripe{
					BillingAddress: stigg.V1CustomerUpdateParamsPassthroughStripeBillingAddress{
						City:       stigg.String("city"),
						Country:    stigg.String("country"),
						Line1:      stigg.String("line1"),
						Line2:      stigg.String("line2"),
						PostalCode: stigg.String("postalCode"),
						State:      stigg.String("state"),
					},
					CustomerName: stigg.String("customerName"),
					InvoiceCustomFields: map[string]string{
						"foo": "string",
					},
					Metadata: map[string]string{
						"foo": "string",
					},
					PaymentMethodID: stigg.String("paymentMethodId"),
					ShippingAddress: stigg.V1CustomerUpdateParamsPassthroughStripeShippingAddress{
						City:       stigg.String("city"),
						Country:    stigg.String("country"),
						Line1:      stigg.String("line1"),
						Line2:      stigg.String("line2"),
						PostalCode: stigg.String("postalCode"),
						State:      stigg.String("state"),
					},
					TaxIDs: []stigg.V1CustomerUpdateParamsPassthroughStripeTaxID{{
						Type:  "type",
						Value: "value",
					}},
				},
				Zuora: stigg.V1CustomerUpdateParamsPassthroughZuora{
					BillingAddress: stigg.V1CustomerUpdateParamsPassthroughZuoraBillingAddress{
						City:       stigg.String("city"),
						Country:    stigg.String("country"),
						Line1:      stigg.String("line1"),
						Line2:      stigg.String("line2"),
						PostalCode: stigg.String("postalCode"),
						State:      stigg.String("state"),
					},
					Currency: "usd",
					Metadata: map[string]string{
						"foo": "string",
					},
					PaymentMethodID: stigg.String("paymentMethodId"),
				},
			},
			Timezone: stigg.String("timezone"),
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

func TestV1CustomerListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.List(context.TODO(), stigg.V1CustomerListParams{
		After:  stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Before: stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CreatedAt: stigg.V1CustomerListParamsCreatedAt{
			Gt:  stigg.Time(time.Now()),
			Gte: stigg.Time(time.Now()),
			Lt:  stigg.Time(time.Now()),
			Lte: stigg.Time(time.Now()),
		},
		Email: stigg.String("email"),
		Limit: stigg.Int(1),
		Name:  stigg.String("name"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerArchive(t *testing.T) {
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
	_, err := client.V1.Customers.Archive(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerImportWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Import(context.TODO(), stigg.V1CustomerImportParams{
		Customers: []stigg.V1CustomerImportParamsCustomer{{
			ID:        "id",
			Email:     stigg.String("dev@stainless.com"),
			Name:      stigg.String("name"),
			BillingID: stigg.String("billingId"),
			Metadata: map[string]string{
				"foo": "string",
			},
			PaymentMethodID: stigg.String("paymentMethodId"),
			SalesforceID:    stigg.String("salesforceId"),
			UpdatedAt:       stigg.Time(time.Now()),
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

func TestV1CustomerListResourcesWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.ListResources(
		context.TODO(),
		"x",
		stigg.V1CustomerListResourcesParams{
			After:  stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Before: stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Limit:  stigg.Int(1),
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

func TestV1CustomerProvisionWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Customers.Provision(context.TODO(), stigg.V1CustomerProvisionParams{
		ID:              "id",
		BillingCurrency: stigg.V1CustomerProvisionParamsBillingCurrencyUsd,
		BillingID:       stigg.String("billingId"),
		CouponID:        stigg.V1CustomerProvisionParamsCouponIDEmpty,
		DefaultPaymentMethod: stigg.V1CustomerProvisionParamsDefaultPaymentMethod{
			BillingID:       stigg.String("billingId"),
			CardExpiryMonth: stigg.Float(0),
			CardExpiryYear:  stigg.Float(0),
			CardLast4Digits: stigg.String("cardLast4Digits"),
			Type:            "CARD",
		},
		Email: stigg.String("dev@stainless.com"),
		Integrations: []stigg.V1CustomerProvisionParamsIntegration{{
			ID:               "id",
			SyncedEntityID:   stigg.String("syncedEntityId"),
			VendorIdentifier: "AUTH0",
		}},
		Language: stigg.String("language"),
		Metadata: map[string]string{
			"foo": "string",
		},
		Name: stigg.String("name"),
		Passthrough: stigg.V1CustomerProvisionParamsPassthrough{
			Stripe: stigg.V1CustomerProvisionParamsPassthroughStripe{
				BillingAddress: stigg.V1CustomerProvisionParamsPassthroughStripeBillingAddress{
					City:       stigg.String("city"),
					Country:    stigg.String("country"),
					Line1:      stigg.String("line1"),
					Line2:      stigg.String("line2"),
					PostalCode: stigg.String("postalCode"),
					State:      stigg.String("state"),
				},
				CustomerName: stigg.String("customerName"),
				InvoiceCustomFields: map[string]string{
					"foo": "string",
				},
				Metadata: map[string]string{
					"foo": "string",
				},
				PaymentMethodID: stigg.String("paymentMethodId"),
				ShippingAddress: stigg.V1CustomerProvisionParamsPassthroughStripeShippingAddress{
					City:       stigg.String("city"),
					Country:    stigg.String("country"),
					Line1:      stigg.String("line1"),
					Line2:      stigg.String("line2"),
					PostalCode: stigg.String("postalCode"),
					State:      stigg.String("state"),
				},
				TaxIDs: []stigg.V1CustomerProvisionParamsPassthroughStripeTaxID{{
					Type:  "type",
					Value: "value",
				}},
			},
			Zuora: stigg.V1CustomerProvisionParamsPassthroughZuora{
				BillingAddress: stigg.V1CustomerProvisionParamsPassthroughZuoraBillingAddress{
					City:       stigg.String("city"),
					Country:    stigg.String("country"),
					Line1:      stigg.String("line1"),
					Line2:      stigg.String("line2"),
					PostalCode: stigg.String("postalCode"),
					State:      stigg.String("state"),
				},
				Currency: "usd",
				Metadata: map[string]string{
					"foo": "string",
				},
				PaymentMethodID: stigg.String("paymentMethodId"),
			},
		},
		Timezone: stigg.String("timezone"),
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CustomerUnarchive(t *testing.T) {
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
	_, err := client.V1.Customers.Unarchive(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
