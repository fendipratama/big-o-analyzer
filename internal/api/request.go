// internal/api/request.go
package api

type AnalyzeRequest struct {
	Code     string `json:"code"`
	Language string `json:"language,omitempty"`
}
