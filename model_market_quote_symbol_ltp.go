/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type MarketQuoteSymbolLtp struct {
	// The last traded price of symbol
	LastPrice       float64 `json:"last_price,omitempty"`
	InstrumentToken string  `json:"instrument_token,omitempty"`
}
