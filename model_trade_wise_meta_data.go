/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

// Response data for brokerage
type TradeWiseMetaData struct {
	// Total count of trades in the trade wise P and L report
	TradesCount int32 `json:"trades_count,omitempty"`
	// Maximum number of trades in a page of the trade wise P and L report API
	PageSizeLimit int32 `json:"page_size_limit,omitempty"`
}
