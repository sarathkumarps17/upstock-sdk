/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type OtherTaxes struct {
	// Transaction charges
	Transaction float32 `json:"transaction,omitempty"`
	// Clearing charges
	Clearing float32 `json:"clearing,omitempty"`
	// IPF charges
	Ipft float32 `json:"ipft,omitempty"`
	// SEBI turnover charges
	SebiTurnover float32 `json:"sebi_turnover,omitempty"`
}
