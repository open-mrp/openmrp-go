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

func TestSaleSalesOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.New(context.TODO(), openmrp.SaleSalesOrderNewParams{
		CreateSalesOrderRequest: openmrp.CreateSalesOrderRequestParam{
			BillToAddressID: "ad_npqa5y43q26z",
			BuyerAccountID:  "ac_opnlh43ymyee",
			Lines: []openmrp.CreateSalesOrderLineInputParam{{
				ProductID: "pd_07oe0r7adh2w",
				Quantity: openmrp.QuantityInputParam{
					UnitID: "un_82bd37dae5po",
					Value:  "10",
				},
				ProductDescription: openmrp.String("product_description"),
				ProductSKU:         openmrp.String("product_sku"),
				UnitPrice: openmrp.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
			}},
			PriorityCode:    openmrp.CreateSalesOrderRequestPriorityCodeNormal,
			ShipToAddressID: "ad_npqa5y43q26z",
			AcknowledgementEmailContacts: []openmrp.SalesOrderEmailContactInputParam{{
				AccountUserID: "acus_e5zu8bde0z3h",
			}},
			CarrierBillingAccountNumber: openmrp.String("123456789"),
			CarrierBillingType:          openmrp.CreateSalesOrderRequestCarrierBillingTypeSender,
			CarrierID:                   openmrp.String("cr_tv5vfjtgu1n3"),
			CustomerPurchaseOrderNumber: openmrp.String("PO-88231"),
			InvoiceEmailContacts: []openmrp.SalesOrderEmailContactInputParam{{
				AccountUserID: "acus_e5zu8bde0z3h",
			}},
			LeadTimeOverrideDays: openmrp.Int(0),
			Note:                 openmrp.String("Rush order for trade show"),
			OrderDiscountID:      openmrp.String("ords_qnbrjvq5ih2q"),
			PaymentTermID:        openmrp.String("pytm_skssmsy21lem"),
			PromisedAt:           openmrp.Time(time.Now()),
			SalesRepID:           openmrp.String("acus_e5zu8bde0z3h"),
			ServiceLevelID:       openmrp.String("crop_4ilk9p6gccrx"),
			ShipByOverrideDate:   openmrp.Time(time.Now()),
			ShippingTermID:       openmrp.String("shtm_c5gxy05whw6r"),
		},
		Include: []string{"customer"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleSalesOrderGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Get(
		context.TODO(),
		"or_9lqo07quiwyb",
		openmrp.SaleSalesOrderGetParams{
			Include: []string{"customer"},
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

func TestSaleSalesOrderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Update(
		context.TODO(),
		"or_9lqo07quiwyb",
		openmrp.SaleSalesOrderUpdateParams{
			Include: []string{"customer"},
			UpdateSalesOrderRequest: openmrp.UpdateSalesOrderRequestParam{
				AcknowledgementEmailContacts: []openmrp.SalesOrderEmailContactInputParam{{
					AccountUserID: "account_user_id",
				}},
				AcknowledgmentStatus:        openmrp.UpdateSalesOrderRequestAcknowledgmentStatusNotSent,
				BillingAddressID:            openmrp.String("billing_address_id"),
				CarrierBillingAccountNumber: openmrp.String("carrier_billing_account_number"),
				CarrierBillingType:          openmrp.UpdateSalesOrderRequestCarrierBillingTypeSender,
				CarrierID:                   openmrp.String("cr_tv5vfjtgu1n3"),
				CustomerID:                  openmrp.String("customer_id"),
				CustomerPurchaseOrderNumber: openmrp.String("customer_purchase_order_number"),
				InvoiceEmailContacts: []openmrp.SalesOrderEmailContactInputParam{{
					AccountUserID: "account_user_id",
				}},
				LeadTimeOverrideDays: openmrp.Int(0),
				Note:                 openmrp.String("Updated shipping instructions"),
				OrderDiscountID:      openmrp.String("order_discount_id"),
				PaymentTermID:        openmrp.String("payment_term_id"),
				PriorityCode:         openmrp.UpdateSalesOrderRequestPriorityCodeNormal,
				PromisedAt:           openmrp.Time(time.Now()),
				SalesRepID:           openmrp.String("sales_rep_id"),
				ServiceLevelID:       openmrp.String("service_level_id"),
				ShipByOverrideDate:   openmrp.Time(time.Now()),
				ShippingAddressID:    openmrp.String("ad_npqa5y43q26z"),
				ShippingTermID:       openmrp.String("shipping_term_id"),
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

func TestSaleSalesOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.List(context.TODO(), openmrp.SaleSalesOrderListParams{
		Cursor:           openmrp.String("cursor"),
		CustomerGroupIDs: []string{"string"},
		CustomerIDs:      []string{"string"},
		EndsAt:           openmrp.String("ends_at"),
		Include:          []string{"customer"},
		ItemIDs:          []string{"string"},
		Limit:            openmrp.Int(0),
		PastDue:          openmrp.Bool(true),
		ProductLineIDs:   []string{"string"},
		Q:                openmrp.String("q"),
		SalesRepIDs:      []string{"string"},
		ShipByAfter:      openmrp.String("ship_by_after"),
		ShipByBefore:     openmrp.String("ship_by_before"),
		StartsAt:         openmrp.String("starts_at"),
		StatusCodes:      []string{"estimate"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleSalesOrderDelete(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Delete(context.TODO(), "or_9lqo07quiwyb")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleSalesOrderCheckout(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Checkout(
		context.TODO(),
		"or_9lqo07quiwyb",
		openmrp.SaleSalesOrderCheckoutParams{
			CheckoutSalesOrderRequest: openmrp.CheckoutSalesOrderRequestParam{
				Email: "operations@acme.example.com",
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

func TestSaleSalesOrderPriceQuote(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.PriceQuote(context.TODO(), openmrp.SaleSalesOrderPriceQuoteParams{
		QuoteSalesOrderPricesRequest: openmrp.QuoteSalesOrderPricesRequestParam{
			BuyerAccountID: "ac_opnlh43ymyee",
			Lines: []openmrp.QuoteSalesOrderLineInputParam{{
				ProductID: "pd_07oe0r7adh2w",
				Quantity: openmrp.QuantityInputParam{
					UnitID: "un_82bd37dae5po",
					Value:  "10",
				},
			}},
		},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleSalesOrderGetStatusesWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.GetStatuses(context.TODO(), openmrp.SaleSalesOrderGetStatusesParams{
		Cursor:  openmrp.String("cursor"),
		Include: []string{"owner"},
		Limit:   openmrp.Int(0),
		Q:       openmrp.String("q"),
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
