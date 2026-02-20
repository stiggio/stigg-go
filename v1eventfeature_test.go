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

func TestV1EventFeatureArchiveFeature(t *testing.T) {
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
	_, err := client.V1.Events.Features.ArchiveFeature(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1EventFeatureNewFeatureWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Events.Features.NewFeature(context.TODO(), stigg.V1EventFeatureNewFeatureParams{
		ID:          "id",
		DisplayName: "displayName",
		FeatureType: stigg.V1EventFeatureNewFeatureParamsFeatureTypeBoolean,
		Description: stigg.String("description"),
		EnumConfiguration: []stigg.V1EventFeatureNewFeatureParamsEnumConfiguration{{
			DisplayName: "displayName",
			Value:       "value",
		}},
		FeatureStatus:      stigg.V1EventFeatureNewFeatureParamsFeatureStatusNew,
		FeatureUnits:       stigg.String("featureUnits"),
		FeatureUnitsPlural: stigg.String("featureUnitsPlural"),
		Metadata: map[string]string{
			"foo": "string",
		},
		MeterType: stigg.V1EventFeatureNewFeatureParamsMeterTypeNone,
		UnitTransformation: stigg.V1EventFeatureNewFeatureParamsUnitTransformation{
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

func TestV1EventFeatureListFeaturesWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Events.Features.ListFeatures(context.TODO(), stigg.V1EventFeatureListFeaturesParams{
		ID:     stigg.String("id"),
		After:  stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Before: stigg.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CreatedAt: stigg.V1EventFeatureListFeaturesParamsCreatedAt{
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

func TestV1EventFeatureGetFeature(t *testing.T) {
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
	_, err := client.V1.Events.Features.GetFeature(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1EventFeatureUnarchiveFeature(t *testing.T) {
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
	_, err := client.V1.Events.Features.UnarchiveFeature(context.TODO(), "x")
	if err != nil {
		var apierr *stigg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1EventFeatureUpdateFeatureWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Events.Features.UpdateFeature(
		context.TODO(),
		"x",
		stigg.V1EventFeatureUpdateFeatureParams{
			Description: stigg.String("description"),
			DisplayName: stigg.String("displayName"),
			EnumConfiguration: []stigg.V1EventFeatureUpdateFeatureParamsEnumConfiguration{{
				DisplayName: "displayName",
				Value:       "value",
			}},
			FeatureUnits:       stigg.String("featureUnits"),
			FeatureUnitsPlural: stigg.String("featureUnitsPlural"),
			Metadata: map[string]string{
				"foo": "string",
			},
			Meter: stigg.V1EventFeatureUpdateFeatureParamsMeter{
				Aggregation: stigg.V1EventFeatureUpdateFeatureParamsMeterAggregation{
					Function: "SUM",
					Field:    stigg.String("field"),
				},
				Filters: []stigg.V1EventFeatureUpdateFeatureParamsMeterFilter{{
					Conditions: []stigg.V1EventFeatureUpdateFeatureParamsMeterFilterCondition{{
						Field:     "field",
						Operation: "EQUALS",
						Value:     stigg.String("value"),
						Values:    []string{"string"},
					}},
				}},
			},
			UnitTransformation: stigg.V1EventFeatureUpdateFeatureParamsUnitTransformation{
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
