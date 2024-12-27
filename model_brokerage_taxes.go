/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

// Taxes levied on order
type BrokerageTaxes struct {
	// GST charges
	Gst float32 `json:"gst,omitempty"`
	// STT charges
	Stt float32 `json:"stt,omitempty"`
	// Stamp duty charges
	StampDuty float32 `json:"stamp_duty,omitempty"`
}