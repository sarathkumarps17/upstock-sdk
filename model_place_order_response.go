/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type PlaceOrderResponse struct {
	Status string          `json:"status,omitempty"`
	Data   *PlaceOrderData `json:"data,omitempty"`
}
