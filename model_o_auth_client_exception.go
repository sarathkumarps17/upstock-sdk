/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/upstock_sdk-api/upstock_sdk-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/upstock_sdk-api/upstock_sdk-codegen.git)
 */
package upstock_sdk

type OAuthClientException struct {
	Cause            *OAuthClientExceptionCause            `json:"cause,omitempty"`
	StackTrace       []OAuthClientExceptionCauseStackTrace `json:"stackTrace,omitempty"`
	Message          string                                `json:"message,omitempty"`
	Suppressed       []OAuthClientExceptionCauseSuppressed `json:"suppressed,omitempty"`
	LocalizedMessage string                                `json:"localizedMessage,omitempty"`
}
