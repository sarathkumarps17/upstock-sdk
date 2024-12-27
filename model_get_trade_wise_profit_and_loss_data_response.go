/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type GetTradeWiseProfitAndLossDataResponse struct {
	Status string `json:"status,omitempty"`
	// Response data for trade wise data details
	Data     []TradeWiseProfitAndLossData  `json:"data,omitempty"`
	Metadata *ProfitAndLossMetaDataWrapper `json:"metadata,omitempty"`
}
