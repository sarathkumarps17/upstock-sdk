/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type TradeHistoryResponsePageData struct {
	PageNumber   int32 `json:"page_number,omitempty"`
	PageSize     int32 `json:"page_size,omitempty"`
	TotalRecords int32 `json:"total_records,omitempty"`
	TotalPages   int32 `json:"total_pages,omitempty"`
}
