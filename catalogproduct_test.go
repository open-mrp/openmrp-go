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

func TestCatalogProductNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.New(context.TODO(), openmrp.CatalogProductNewParams{
		CreateProductRequest: openmrp.CreateProductRequestParam{
			CategoryID:       "ic_d06g9c6yc9ck",
			SKU:              "ALM-2024-1001",
			Type:             openmrp.CreateProductRequestTypeSale,
			AttributeIDs:     []string{"at_rf1w295jt5ia"},
			Description:      openmrp.String("Wireless barcode scanner with charging cradle"),
			Notes:            openmrp.String("Ships with a 2-year warranty; register for extended coverage."),
			PortalVisibility: openmrp.CreateProductRequestPortalVisibilityVisible,
			ProductLineID:    openmrp.String("pdln_k9bnlgvxhxjh"),
			UnitCost: openmrp.RateInputParam{
				DenominatorUnitID: "un_82bd37dae5po",
				NumeratorUnitID:   "un_82bd37dae5po",
				Value:             "112.00",
			},
			UnitPrice: openmrp.RateInputParam{
				DenominatorUnitID: "un_82bd37dae5po",
				NumeratorUnitID:   "un_82bd37dae5po",
				Value:             "199.00",
			},
		},
		Include: []string{"product_line"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogProductGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.Get(
		context.TODO(),
		"pd_07oe0r7adh2w",
		openmrp.CatalogProductGetParams{
			Include: []string{"product_line"},
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

func TestCatalogProductUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.Update(
		context.TODO(),
		"pd_07oe0r7adh2w",
		openmrp.CatalogProductUpdateParams{
			Include: []string{"product_line"},
			UpdateProductRequest: openmrp.UpdateProductRequestParam{
				CategoryID:       openmrp.String("ic_d06g9c6yc9ck"),
				Description:      openmrp.String("Wireless barcode scanner with charging cradle (v2)"),
				Notes:            openmrp.String("Firmware 2.1 improves Bluetooth pairing reliability."),
				PortalVisibility: openmrp.UpdateProductRequestPortalVisibilityVisible,
				SKU:              openmrp.String("SKU-002"),
				UnitPrice: openmrp.RateInputParam{
					DenominatorUnitID: "un_82bd37dae5po",
					NumeratorUnitID:   "un_82bd37dae5po",
					Value:             "219.00",
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

func TestCatalogProductListWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.List(context.TODO(), openmrp.CatalogProductListParams{
		AttributeIDs:     []string{"string"},
		CategoryIDs:      []string{"string"},
		Cursor:           openmrp.String("cursor"),
		CustomerIDs:      []string{"string"},
		EndsAt:           openmrp.Time(time.Now()),
		Include:          []string{"product_line"},
		Limit:            openmrp.Int(0),
		PortalVisibility: openmrp.CatalogProductListParamsPortalVisibilityVisible,
		ProductLineIDs:   []string{"string"},
		Q:                openmrp.String("q"),
		StartsAt:         openmrp.Time(time.Now()),
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogProductDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.Delete(
		context.TODO(),
		"pd_07oe0r7adh2w",
		openmrp.CatalogProductDeleteParams{
			Include: []string{"product_line"},
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

func TestCatalogProductChangeProductLineWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.ChangeProductLine(
		context.TODO(),
		"pdln_k9bnlgvxhxjh",
		openmrp.CatalogProductChangeProductLineParams{
			ID:      "pd_07oe0r7adh2w",
			Include: []string{"product_line"},
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
