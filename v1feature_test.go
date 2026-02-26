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

func TestV1FeatureArchiveFeature(t *testing.T) {
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
	_, err := client.V1.Features.ArchiveFeature(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1FeatureNewFeatureWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Features.NewFeature(context.TODO(), stigg.V1FeatureNewFeatureParams{
		ID:          "id",
		DisplayName: "displayName",
		FeatureType: stigg.V1FeatureNewFeatureParamsFeatureTypeBoolean,
		Description: stigg.String("description"),
		EnumConfiguration: []stigg.V1FeatureNewFeatureParamsEnumConfiguration{{
			DisplayName: "displayName",
			Value:       "value",
		}},
		FeatureStatus:      stigg.V1FeatureNewFeatureParamsFeatureStatusNew,
		FeatureUnits:       stigg.String("featureUnits"),
		FeatureUnitsPlural: stigg.String("featureUnitsPlural"),
		Metadata: map[string]string{
			"foo": "string",
		},
		MeterType: stigg.V1FeatureNewFeatureParamsMeterTypeNone,
		UnitTransformation: stigg.V1FeatureNewFeatureParamsUnitTransformation{
			Divide:             0,
			FeatureUnits:       stigg.String("featureUnits"),
			FeatureUnitsPlural: stigg.String("featureUnitsPlural"),
			Round:              "UP",
		},
	})
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1FeatureListFeaturesWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Features.ListFeatures(context.TODO(), stigg.V1FeatureListFeaturesParams{
		ID:     stigg.String("id"),
		After:  stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Before: stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CreatedAt: stigg.V1FeatureListFeaturesParamsCreatedAt{
			Gt:  stigg.Time(time.Now()),
			Gte: stigg.Time(time.Now()),
			Lt:  stigg.Time(time.Now()),
			Lte: stigg.Time(time.Now()),
		},
		FeatureType: stigg.String("featureType"),
		Limit:       stigg.Int(1),
		MeterType:   stigg.String("meterType"),
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

func TestV1FeatureGetFeature(t *testing.T) {
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
	_, err := client.V1.Features.GetFeature(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1FeatureUnarchiveFeature(t *testing.T) {
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
	_, err := client.V1.Features.UnarchiveFeature(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1FeatureUpdateFeatureWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Features.UpdateFeature(
		context.TODO(),
		"x",
		stigg.V1FeatureUpdateFeatureParams{
			Description: stigg.String("description"),
			DisplayName: stigg.String("displayName"),
			EnumConfiguration: []stigg.V1FeatureUpdateFeatureParamsEnumConfiguration{{
				DisplayName: "displayName",
				Value:       "value",
			}},
			FeatureUnits:       stigg.String("featureUnits"),
			FeatureUnitsPlural: stigg.String("featureUnitsPlural"),
			Metadata: map[string]string{
				"foo": "string",
			},
			Meter: stigg.V1FeatureUpdateFeatureParamsMeter{
				Aggregation: stigg.V1FeatureUpdateFeatureParamsMeterAggregation{
					Function: "SUM",
					Field:    stigg.String("field"),
				},
				Filters: []stigg.V1FeatureUpdateFeatureParamsMeterFilter{{
					Conditions: []stigg.V1FeatureUpdateFeatureParamsMeterFilterCondition{{
						Field:     "field",
						Operation: "EQUALS",
						Value:     stigg.String("value"),
						Values:    []string{"string"},
					}},
				}},
			},
			UnitTransformation: stigg.V1FeatureUpdateFeatureParamsUnitTransformation{
				Divide:             0,
				FeatureUnits:       stigg.String("featureUnits"),
				FeatureUnitsPlural: stigg.String("featureUnitsPlural"),
				Round:              "UP",
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
