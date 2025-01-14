/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type PutCallOptionChainData struct {
	InstrumentKey string         `json:"instrument_key,omitempty"`
	MarketData    *MarketData    `json:"market_data,omitempty"`
	OptionGreeks  *AnalyticsData `json:"option_greeks,omitempty"`
}
