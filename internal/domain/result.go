package domain

type Result struct {
	Language        string   `json:"language"`
	TimeComplexity  string   `json:"time_complexity"`
	SpaceComplexity string   `json:"space_complexity"`
	Confidence      float64  `json:"confidence"`

	Explanation   string   `json:"explanation,omitempty"`
	Impact        string   `json:"impact,omitempty"`
	Reasoning     []string `json:"reasoning,omitempty"`
	Optimizations []string `json:"optimizations,omitempty"`

	EstimatedExecution []EstimatedExecution `json:"estimated_execution,omitempty"`
}
