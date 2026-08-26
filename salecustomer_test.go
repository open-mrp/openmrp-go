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

func TestSaleCustomerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Customers.New(context.TODO(), openmrp.SaleCustomerNewParams{
		CreateCustomerRequest: openmrp.CreateCustomerRequestParam{
			BillToAddress: openmrp.AddressInputParam{
				Country:           "US",
				Name:              "Acme Inc.",
				Email:             openmrp.String("warehouse@acme.com"),
				Locality:          openmrp.String("New York"),
				Phone:             openmrp.String("555-123-4567"),
				PostalCode:        openmrp.String("10001"),
				ReceiveCalendarID: openmrp.String("receive_calendar_id"),
				State:             openmrp.String("NY"),
				StreetLine1:       openmrp.String("123 Main St"),
				StreetLine2:       openmrp.String("Suite 400"),
				Type:              openmrp.AddressInputTypeStandard,
			},
			CustomerTypeGroupID:   "acgp_6p4z57e9alaf",
			DefaultCarrierID:      "cr_tv5vfjtgu1n3",
			DefaultPaymentTermID:  "pytm_skssmsy21lem",
			DefaultShippingTermID: "shtm_c5gxy05whw6r",
			Name:                  "Acme Inc.",
			ShipToAddress: openmrp.AddressInputParam{
				Country:           "US",
				Name:              "Acme Inc.",
				Email:             openmrp.String("warehouse@acme.com"),
				Locality:          openmrp.String("New York"),
				Phone:             openmrp.String("555-123-4567"),
				PostalCode:        openmrp.String("10001"),
				ReceiveCalendarID: openmrp.String("receive_calendar_id"),
				State:             openmrp.String("NY"),
				StreetLine1:       openmrp.String("123 Main St"),
				StreetLine2:       openmrp.String("Suite 400"),
				Type:              openmrp.AddressInputTypeStandard,
			},
			CarrierBillingAccount: openmrp.String("123456789"),
			CarrierBillingType:    openmrp.CreateCustomerRequestCarrierBillingTypeSender,
			CommissionPolicy:      openmrp.CreateCustomerRequestCommissionPolicyCommissionApplied,
			CreditLimit: openmrp.QuantityInputParam{
				UnitID: "un_82bd37dae5po",
				Value:  "10000.00",
			},
			CustomerPriceGroupIDs: []string{"acgp_6p4z57e9alaf"},
			DefaultPriority:       openmrp.CreateCustomerRequestDefaultPriorityNormal,
			DefaultSalesRepID:     openmrp.String("acus_e5zu8bde0z3h"),
			DefaultServiceLevelID: openmrp.String("crop_4ilk9p6gccrx"),
			EdiStatus:             openmrp.CreateCustomerRequestEdiStatusDisabled,
			Email:                 openmrp.String("orders@acme.com"),
			FreightPolicy:         openmrp.CreateCustomerRequestFreightPolicyBilledFreight,
			FulfillmentPolicy:     openmrp.CreateCustomerRequestFulfillmentPolicyMakeToOrder,
			LeadTimeDays:          openmrp.Int(0),
			Note:                  openmrp.String("Key enterprise account"),
			Number:                openmrp.String("100042"),
			Phone:                 openmrp.String("555-123-4567"),
			ReceiveCalendarID:     openmrp.String("receive_calendar_id"),
			Status:                openmrp.CreateCustomerRequestStatusNormal,
			URL:                   openmrp.String("https://acme.com"),
		},
		Include: []string{"bill_to_address"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleCustomerGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Customers.Get(
		context.TODO(),
		"ac_opnlh43ymyee",
		openmrp.SaleCustomerGetParams{
			Include: []string{"bill_to_address"},
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

func TestSaleCustomerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Customers.Update(
		context.TODO(),
		"ac_opnlh43ymyee",
		openmrp.SaleCustomerUpdateParams{
			Include: []string{"bill_to_address"},
			UpdateCustomerRequest: openmrp.UpdateCustomerRequestParam{
				BillToAddressID:       openmrp.String("ad_npqa5y43q26z"),
				CarrierBillingAccount: openmrp.String("123456789"),
				CarrierBillingType:    openmrp.UpdateCustomerRequestCarrierBillingTypeSender,
				CommissionPolicy:      openmrp.UpdateCustomerRequestCommissionPolicyCommissionApplied,
				CreditLimit: openmrp.QuantityInputParam{
					UnitID: "un_82bd37dae5po",
					Value:  "10000.00",
				},
				CustomerPriceGroupIDs: []string{"acgp_6p4z57e9alaf"},
				CustomerTypeGroupID:   openmrp.String("acgp_6p4z57e9alaf"),
				DefaultCarrierID:      openmrp.String("cr_tv5vfjtgu1n3"),
				DefaultPaymentTermID:  openmrp.String("pytm_skssmsy21lem"),
				DefaultPriority:       openmrp.UpdateCustomerRequestDefaultPriorityNormal,
				DefaultSalesRepID:     openmrp.String("acus_e5zu8bde0z3h"),
				DefaultServiceLevelID: openmrp.String("crop_4ilk9p6gccrx"),
				DefaultShippingTermID: openmrp.String("shtm_c5gxy05whw6r"),
				EdiStatus:             openmrp.UpdateCustomerRequestEdiStatusDisabled,
				Email:                 openmrp.String("orders@acme.com"),
				FreightPolicy:         openmrp.UpdateCustomerRequestFreightPolicyBilledFreight,
				FulfillmentPolicy:     openmrp.UpdateCustomerRequestFulfillmentPolicyMakeToOrder,
				LeadTimeDays:          openmrp.Int(0),
				Name:                  openmrp.String("Acme Corp Updated"),
				Note:                  openmrp.String("Updated account notes"),
				Number:                openmrp.String("100042"),
				Phone:                 openmrp.String("555-123-4567"),
				ReceiveCalendarID:     openmrp.String("receive_calendar_id"),
				ShipToAddressID:       openmrp.String("ad_npqa5y43q26z"),
				Status:                openmrp.UpdateCustomerRequestStatusNormal,
				URL:                   openmrp.String("https://acme.com"),
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

func TestSaleCustomerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Customers.List(context.TODO(), openmrp.SaleCustomerListParams{
		CarrierIDs:            []string{"string"},
		City:                  openmrp.String("city"),
		CommissionStatusCodes: []string{"commission_applied"},
		Cursor:                openmrp.String("cursor"),
		CustomerGroupIDs:      []string{"string"},
		EndsAt:                openmrp.Time(time.Now()),
		FreightStatusCodes:    []string{"free_freight"},
		Include:               []string{"bill_to_address"},
		Limit:                 openmrp.Int(0),
		ParentAccountStatus:   openmrp.SaleCustomerListParamsParentAccountStatusParent,
		PaymentTermIDs:        []string{"string"},
		PostalCode:            openmrp.String("postal_code"),
		PricingGroupIDs:       []string{"string"},
		Q:                     openmrp.String("q"),
		SalesRepIDs:           []string{"string"},
		ServiceLevelIDs:       []string{"string"},
		ShippingTermIDs:       []string{"string"},
		StartsAt:              openmrp.Time(time.Now()),
		State:                 openmrp.String("state"),
		StatusCodes:           []string{"normal"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleCustomerDelete(t *testing.T) {
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
	_, err := client.Sales.Customers.Delete(context.TODO(), "ac_opnlh43ymyee")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleCustomerGetLeadTimeWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Customers.GetLeadTime(
		context.TODO(),
		"ac_opnlh43ymyee",
		openmrp.SaleCustomerGetLeadTimeParams{
			Include: []string{"account_group"},
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
