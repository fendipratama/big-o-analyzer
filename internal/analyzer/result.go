package analyzer

type Result struct {
	Language       string   `json:"language"`
	TimeComplexity string   `json:"time_complexity"`
	SpaceComplexity string  `json:"space_complexity"`
	Confidence     float64  `json:"confidence"`
	Reasoning      []string `json:"reasoning"`
}
