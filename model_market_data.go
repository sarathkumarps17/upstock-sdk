/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type MarketData struct {
	Ltp        float64 `json:"ltp,omitempty"`
	Volume     int64   `json:"volume,omitempty"`
	Oi         float64 `json:"oi,omitempty"`
	ClosePrice float64 `json:"close_price,omitempty"`
	BidPrice   float64 `json:"bid_price,omitempty"`
	BidQty     int32   `json:"bid_qty,omitempty"`
	AskPrice   float64 `json:"ask_price,omitempty"`
	AskQty     int32   `json:"ask_qty,omitempty"`
	PrevOi     float64 `json:"prev_oi,omitempty"`
}
