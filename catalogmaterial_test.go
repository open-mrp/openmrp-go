// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/open-mrp/openmrp-go"
	"github.com/open-mrp/openmrp-go/internal/testutil"
	"github.com/open-mrp/openmrp-go/option"
)

func TestCatalogMaterialNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Catalog.Materials.New(context.TODO(), openmrp.CatalogMaterialNewParams{
		CreateMaterialRequest: openmrp.CreateMaterialRequestParam{
			CategoryID:   "ic_d06g9c6yc9ck",
			SKU:          "MAT-001",
			AttributeIDs: []string{"at_rf1w295jt5ia"},
			Description:  openmrp.String("Cold-rolled 304 stainless steel sheet, 1.5mm"),
			LeadTime: openmrp.QuantityInputRequestParam{
				UnitID: "un_82bd37dae5po",
				Value:  "7.00",
			},
			Notes: openmrp.String("Store flat in a dry area to avoid surface oxidation."),
			OrderPoint: openmrp.QuantityInputRequestParam{
				UnitID: "un_82bd37dae5po",
				Value:  "100.00",
			},
			UnitCost: openmrp.RateInputParam{
				DenominatorUnitID: "un_82bd37dae5po",
				NumeratorUnitID:   "un_82bd37dae5po",
				Value:             "8.25",
			},
			UnitPrice: openmrp.RateInputParam{
				DenominatorUnitID: "un_82bd37dae5po",
				NumeratorUnitID:   "un_82bd37dae5po",
				Value:             "12.50",
			},
		},
		Include: []string{"item"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogMaterialGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Catalog.Materials.Get(
		context.TODO(),
		"ml_ow202v78slbl",
		openmrp.CatalogMaterialGetParams{
			Include: []string{"item"},
		},
	)
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogMaterialUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Catalog.Materials.Update(
		context.TODO(),
		"ml_ow202v78slbl",
		openmrp.CatalogMaterialUpdateParams{
			Include: []string{"item"},
			UpdateMaterialRequest: openmrp.UpdateMaterialRequestParam{
				CategoryID:  openmrp.String("ic_d06g9c6yc9ck"),
				Description: openmrp.String("Cold-rolled 304 stainless steel sheet, 2.0mm"),
				LeadTime: openmrp.QuantityInputRequestParam{
					UnitID: "un_82bd37dae5po",
					Value:  "10.00",
				},
				Notes: openmrp.String("Reorder point raised after Q2 demand spike."),
				OrderPoint: openmrp.QuantityInputRequestParam{
					UnitID: "un_82bd37dae5po",
					Value:  "150.00",
				},
				SKU: openmrp.String("MAT-001-UPDATED"),
				UnitCost: openmrp.RateInputParam{
					DenominatorUnitID: "un_82bd37dae5po",
					NumeratorUnitID:   "un_82bd37dae5po",
					Value:             "9.10",
				},
			},
		},
	)
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogMaterialListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Catalog.Materials.List(context.TODO(), openmrp.CatalogMaterialListParams{
		AttributeIDs: []string{"string"},
		CategoryIDs:  []string{"string"},
		Cursor:       openmrp.String("cursor"),
		EndsAt:       openmrp.Time(time.Now()),
		Include:      []string{"item"},
		Limit:        openmrp.Int(0),
		Q:            openmrp.String("q"),
		StartsAt:     openmrp.Time(time.Now()),
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogMaterialDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Catalog.Materials.Delete(context.TODO(), "ml_ow202v78slbl")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
