/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type AnalyticsData struct {
	Vega  float64 `json:"vega,omitempty"`
	Theta float64 `json:"theta,omitempty"`
	Gamma float64 `json:"gamma,omitempty"`
	Delta float64 `json:"delta,omitempty"`
	Iv    float64 `json:"iv,omitempty"`
	Pop   float64 `json:"pop,omitempty"`
}
