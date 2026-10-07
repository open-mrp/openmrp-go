// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// The planning assumptions production schedules are solved against, and the
// per-resource overrides that mark which machines constrain the plan.
//
// OperationFulfillmentRecommendationService contains methods and other services
// that help with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationFulfillmentRecommendationService] method instead.
type OperationFulfillmentRecommendationService struct {
	options []option.RequestOption
	// The planning assumptions production schedules are solved against, and the
	// per-resource overrides that mark which machines constrain the plan.
	Actions OperationFulfillmentRecommendationActionService
}

// NewOperationFulfillmentRecommendationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewOperationFulfillmentRecommendationService(opts ...option.RequestOption) (r OperationFulfillmentRecommendationService) {
	r = OperationFulfillmentRecommendationService{}
	r.options = opts
	r.Actions = NewOperationFulfillmentRecommendationActionService(opts...)
	return
}

// Returns, for every sellable SKU, whether it should be built to stock or only
// against orders — and the measurement that decided.
//
// The rules are ordered and the first match wins. Lead-time feasibility is checked
// before anything else: if customers are promised less time than production needs,
// building to order is not possible rather than not preferred, and no amount of
// lumpy demand changes that. After that the engine looks for dead stock, a single
// contract customer, demand too erratic for a buffer to size, and slow-moving
// expensive units.
//
// Every verdict carries its numbers — demand interval, variability, customer
// concentration, promised lead time, annual cost of goods — so a planner can
// disagree with the rule rather than only with the answer. Thresholds are
// merchant-editable in the planning settings.
//
// Computed fresh on every call rather than stored. A recommendation is only
// meaningful next to current demand, and a saved one would go quietly stale; the
// durable artifact is the item setting written when someone agrees with it.
// Nothing here changes a plan on its own.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationFulfillmentRecommendationService) List(ctx context.Context, opts ...option.RequestOption) (res *ListFulfillmentRecommendation, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/fulfillment-recommendations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The engine's advice on how one SKU should be produced, with the measurements
// behind it.
type FulfillmentRecommendation struct {
	// Annual cost of goods for this item: demand times unit cost.
	//
	// Null unless the caller holds `costs:read`; customer and supplier portal users
	// never see it.
	AnnualCogs float64 `json:"annual_cogs" api:"required"`
	// Months observed divided by months with demand: 1 means it sells every month, 3
	// means once a quarter on average.
	//
	// Measured on monthly buckets, which cannot distinguish two orders in one month
	// from one.
	AverageDemandInterval float64 `json:"average_demand_interval" api:"required"`
	// Whether adopting the recommendation would change anything.
	Changes bool `json:"changes" api:"required"`
	// Squared coefficient of variation over the months that had demand, measuring how
	// uneven the quantities are.
	CoefficientOfVariation float64 `json:"coefficient_of_variation" api:"required"`
	// How the item is planned today.
	//
	// Any of "make_to_stock", "make_to_order".
	CurrentPolicy FulfillmentRecommendationCurrentPolicy `json:"current_policy" api:"required"`
	// Calendar days customers are promised on average, weighted by how much each buys.
	DemandWeightedLeadTimeDays float64 `json:"demand_weighted_lead_time_days" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Percentage of demand from customers whose own stated policy disagrees with the
	// recommendation.
	//
	// A policy is resolved per SKU, so an item sold to both a stocking distributor and
	// a contract customer gets one answer either way. A high share here is the signal
	// that the single answer is uncomfortable.
	MixedStreamSharePct float64 `json:"mixed_stream_share_pct" api:"required"`
	// Months since anything last sold, capped at the observation window.
	MonthsSinceLastSale int64 `json:"months_since_last_sale" api:"required"`
	// Resource type identifier.
	//
	// Any of "fulfillment_recommendation".
	Object FulfillmentRecommendationObject `json:"object" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductLine Entity `json:"product_line" api:"required"`
	// The rule that decided.
	//
	//   - `lead_time_infeasible`: customers are promised less time than production
	//     needs, so the stock has to exist before the order does. Checked first, because
	//     producing to order is not possible rather than not preferred.
	//   - `no_recent_demand`: nothing has sold for long enough that a buffer is dead
	//     stock.
	//   - `single_customer`: effectively one customer buys it, and that customer is
	//     served to order.
	//   - `lumpy_demand`: demand arrives rarely and in wildly different sizes, which is
	//     the shape a safety stock sizes worst.
	//   - `slow_moving_high_value`: expensive units, few sold — the buffer costs more
	//     than the service it buys.
	//   - `steady_demand`: regular enough to forecast, which is what stocking is for.
	//
	// Any of "lead_time_infeasible", "no_recent_demand", "single_customer",
	// "lumpy_demand", "slow_moving_high_value", "steady_demand".
	Reason FulfillmentRecommendationReason `json:"reason" api:"required"`
	// How the engine thinks it should be planned.
	//
	// Any of "make_to_stock", "make_to_order".
	RecommendedPolicy FulfillmentRecommendationRecommendedPolicy `json:"recommended_policy" api:"required"`
	// SKU of that item.
	SKU string `json:"sku" api:"required"`
	// Name of that customer.
	TopCustomerName string `json:"top_customer_name" api:"required"`
	// The largest customer's share of this item's demand, as a percentage.
	TopCustomerSharePct float64 `json:"top_customer_share_pct" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnnualCogs                 respjson.Field
		AverageDemandInterval      respjson.Field
		Changes                    respjson.Field
		CoefficientOfVariation     respjson.Field
		CurrentPolicy              respjson.Field
		DemandWeightedLeadTimeDays respjson.Field
		Item                       respjson.Field
		MixedStreamSharePct        respjson.Field
		MonthsSinceLastSale        respjson.Field
		Object                     respjson.Field
		ProductLine                respjson.Field
		Reason                     respjson.Field
		RecommendedPolicy          respjson.Field
		SKU                        respjson.Field
		TopCustomerName            respjson.Field
		TopCustomerSharePct        respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FulfillmentRecommendation) RawJSON() string { return r.JSON.raw }
func (r *FulfillmentRecommendation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the item is planned today.
type FulfillmentRecommendationCurrentPolicy string

const (
	FulfillmentRecommendationCurrentPolicyMakeToStock FulfillmentRecommendationCurrentPolicy = "make_to_stock"
	FulfillmentRecommendationCurrentPolicyMakeToOrder FulfillmentRecommendationCurrentPolicy = "make_to_order"
)

// Resource type identifier.
type FulfillmentRecommendationObject string

const (
	FulfillmentRecommendationObjectFulfillmentRecommendation FulfillmentRecommendationObject = "fulfillment_recommendation"
)

// The rule that decided.
//
//   - `lead_time_infeasible`: customers are promised less time than production
//     needs, so the stock has to exist before the order does. Checked first, because
//     producing to order is not possible rather than not preferred.
//   - `no_recent_demand`: nothing has sold for long enough that a buffer is dead
//     stock.
//   - `single_customer`: effectively one customer buys it, and that customer is
//     served to order.
//   - `lumpy_demand`: demand arrives rarely and in wildly different sizes, which is
//     the shape a safety stock sizes worst.
//   - `slow_moving_high_value`: expensive units, few sold — the buffer costs more
//     than the service it buys.
//   - `steady_demand`: regular enough to forecast, which is what stocking is for.
type FulfillmentRecommendationReason string

const (
	FulfillmentRecommendationReasonLeadTimeInfeasible  FulfillmentRecommendationReason = "lead_time_infeasible"
	FulfillmentRecommendationReasonNoRecentDemand      FulfillmentRecommendationReason = "no_recent_demand"
	FulfillmentRecommendationReasonSingleCustomer      FulfillmentRecommendationReason = "single_customer"
	FulfillmentRecommendationReasonLumpyDemand         FulfillmentRecommendationReason = "lumpy_demand"
	FulfillmentRecommendationReasonSlowMovingHighValue FulfillmentRecommendationReason = "slow_moving_high_value"
	FulfillmentRecommendationReasonSteadyDemand        FulfillmentRecommendationReason = "steady_demand"
)

// How the engine thinks it should be planned.
type FulfillmentRecommendationRecommendedPolicy string

const (
	FulfillmentRecommendationRecommendedPolicyMakeToStock FulfillmentRecommendationRecommendedPolicy = "make_to_stock"
	FulfillmentRecommendationRecommendedPolicyMakeToOrder FulfillmentRecommendationRecommendedPolicy = "make_to_order"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListFulfillmentRecommendation struct {
	// Resources in this page.
	Data []FulfillmentRecommendation `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListFulfillmentRecommendationObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListFulfillmentRecommendation) RawJSON() string { return r.JSON.raw }
func (r *ListFulfillmentRecommendation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListFulfillmentRecommendationObject string

const (
	ListFulfillmentRecommendationObjectList ListFulfillmentRecommendationObject = "list"
)
