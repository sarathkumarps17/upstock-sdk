/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type PlaceOrderRequest struct {
	// Quantity with which the order is to be placed
	Quantity int32 `json:"quantity"`
	// Signifies if the order was either Intraday, Delivery, CO or OCO
	Product string `json:"product"`
	// It can be one of the following - DAY(default), IOC
	Validity string `json:"validity"`
	// Price at which the order will be placed
	Price float32 `json:"price"`
	// Tag for a particular order
	Tag string `json:"tag,omitempty"`
	// Key of the instrument
	InstrumentToken string `json:"instrument_token"`
	// Type of order. It can be one of the following MARKET refers to market order LIMIT refers to Limit Order SL refers to Stop Loss Limit SL-M refers to Stop Loss Market
	OrderType string `json:"order_type"`
	// Indicates whether its a buy or sell order
	TransactionType string `json:"transaction_type"`
	// The quantity that should be disclosed in the market depth
	DisclosedQuantity int32 `json:"disclosed_quantity"`
	// If the order is a stop loss order then the trigger price to be set is mentioned here
	TriggerPrice float32 `json:"trigger_price"`
	// Signifies if the order is an After Market Order
	IsAmo bool `json:"is_amo"`
}
