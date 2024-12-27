/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type ConvertPositionRequest struct {
	// Key of the instrument
	InstrumentToken string `json:"instrument_token"`
	// Indicates the new product to use for the convert positions
	NewProduct string `json:"new_product"`
	// Indicates the old product to use for the convert positions
	OldProduct string `json:"old_product"`
	// Indicates whether its a buy(b) or sell(s) order
	TransactionType string `json:"transaction_type"`
	// Quantity with which the position to convert
	Quantity int32 `json:"quantity"`
}
