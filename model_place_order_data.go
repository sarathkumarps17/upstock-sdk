/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

// Response data for place order request
type PlaceOrderData struct {
	// An order ID for the order request placed
	OrderId string `json:"order_id,omitempty"`
}
