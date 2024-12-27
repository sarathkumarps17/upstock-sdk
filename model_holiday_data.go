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

// Response data for holiday list
type HolidayData struct {
	Date            time.Time            `json:"date,omitempty"`
	Description     string               `json:"description,omitempty"`
	HolidayType     string               `json:"holiday_type,omitempty"`
	ClosedExchanges []string             `json:"closed_exchanges,omitempty"`
	OpenExchanges   []ExchangeTimingData `json:"open_exchanges,omitempty"`
}