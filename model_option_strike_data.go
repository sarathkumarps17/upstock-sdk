/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

import (
	"time"
)

// Response data for option chain data
type OptionStrikeData struct {
	Expiry              time.Time               `json:"expiry,omitempty"`
	Pcr                 float64                 `json:"pcr,omitempty"`
	StrikePrice         float64                 `json:"strike_price,omitempty"`
	UnderlyingKey       string                  `json:"underlying_key,omitempty"`
	UnderlyingSpotPrice float64                 `json:"underlying_spot_price,omitempty"`
	CallOptions         *PutCallOptionChainData `json:"call_options,omitempty"`
	PutOptions          *PutCallOptionChainData `json:"put_options,omitempty"`
}