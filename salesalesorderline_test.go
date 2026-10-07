// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/open-mrp/openmrp-go"
	"github.com/open-mrp/openmrp-go/internal/testutil"
	"github.com/open-mrp/openmrp-go/option"
)

func TestSaleSalesOrderLineNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Lines.New(
		context.TODO(),
		"or_9lqo07quiwyb",
		openmrp.SaleSalesOrderLineNewParams{
			CreateSalesOrderLineRequest: openmrp.CreateSalesOrderLineRequestParam{
				Metadata: map[string]string{
					"foo": "string",
				},
				ProductID:  "pd_07oe0r7adh2w",
				ProductSKU: "WIDGET-001",
				Quantity: openmrp.QuantityInputParam{
					UnitID: "un_82bd37dae5po",
					Value:  "10",
				},
				ProductDescription: openmrp.String("product_description"),
				UnitPrice: openmrp.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
			},
			Include: []string{"product"},
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

func TestSaleSalesOrderLineUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Lines.Update(
		context.TODO(),
		"example",
		openmrp.SaleSalesOrderLineUpdateParams{
			ID:      "or_9lqo07quiwyb",
			Include: []string{"product"},
			UpdateSalesOrderLineRequest: openmrp.UpdateSalesOrderLineRequestParam{
				Metadata: map[string]string{
					"foo": "string",
				},
				ProductDescription: openmrp.String("product_description"),
				ProductSKU:         openmrp.String("product_sku"),
				Quantity: openmrp.QuantityInputParam{
					UnitID: "un_82bd37dae5po",
					Value:  "20",
				},
				UnitCost: openmrp.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
				UnitPrice: openmrp.RateInputParam{
					DenominatorUnitID: "un_82bd37dae5po",
					NumeratorUnitID:   "un_82bd37dae5po",
					Value:             "30.00",
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

func TestSaleSalesOrderLineDelete(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Lines.Delete(
		context.TODO(),
		"example",
		openmrp.SaleSalesOrderLineDeleteParams{
			ID: "or_9lqo07quiwyb",
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
