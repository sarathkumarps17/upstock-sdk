/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

// Response data for charges details
type BrokerageData struct {
	// Total charges for the order
	Total float32 `json:"total,omitempty"`
	// Brokerage charges for the order
	Brokerage    float32         `json:"brokerage,omitempty"`
	Taxes        *BrokerageTaxes `json:"taxes,omitempty"`
	OtherTaxes   *OtherTaxes     `json:"otherTaxes,omitempty"`
	CDpPlan      *DpPlan         `json:"dpPlan,omitempty"`
	OtherCharges *OtherTaxes     `json:"other_charges,omitempty"`
	DpPlan       *DpPlan         `json:"dp_plan,omitempty"`
}
