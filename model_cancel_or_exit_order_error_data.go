/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

// Error data for cancel or exit order request
type CancelOrExitOrderErrorData struct {
	// Unique code for the error state

	// Verbose message for the error state
	Message string `json:"message,omitempty"`
	// Path to property failing validation
	CPropertyPath string `json:"propertyPath,omitempty"`
	// Invalid value for the property failing validation

	ErrorCode    string       `json:"error_code,omitempty"`
	PropertyPath string       `json:"property_path,omitempty"`
	InvalidValue *interface{} `json:"invalid_value,omitempty"`
	// Key of instrument
	InstrumentKey string `json:"instrument_key,omitempty"`
	// Reference order ID
	OrderId string `json:"order_id,omitempty"`
}
