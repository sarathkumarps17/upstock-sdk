/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type ModifyOrderRequest struct {
	// Quantity with which the order was placed
	Quantity int32 `json:"quantity,omitempty"`
	// Order validity (DAY- Day and IOC- Immediate or Cancel (IOC) order)
	Validity string `json:"validity"`
	// Price at which the order was placed
	Price float32 `json:"price"`
	// The order ID for which the order must be modified
	OrderId string `json:"order_id"`
	// Type of order. It can be one of the following MARKET refers to market order LIMILT refers to Limit Order SL refers to Stop Loss Limit SL-M refers to Stop Loss Market
	OrderType string `json:"order_type"`
	// The quantity that should be disclosed in the market depth
	DisclosedQuantity int32 `json:"disclosed_quantity,omitempty"`
	// If the order is a stop loss order then the trigger price to be set is mentioned here
	TriggerPrice float32 `json:"trigger_price"`
}
