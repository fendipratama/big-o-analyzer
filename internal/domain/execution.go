// internal/domain/execution.go
package domain

type EstimatedExecution struct {
	InputSize string `json:"input_size"`
	Operations string `json:"operations"`
	TimeMs float64 `json:"time_ms"`
}
