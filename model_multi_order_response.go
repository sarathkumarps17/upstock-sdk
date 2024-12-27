/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type MultiOrderResponse struct {
	Status string `json:"status,omitempty"`
	// Response data for multi order request
	Data []MultiOrderData `json:"data,omitempty"`
	// Error details for multi order request
	Errors  []MultiOrderError  `json:"errors,omitempty"`
	Summary *MultiOrderSummary `json:"summary,omitempty"`
}
