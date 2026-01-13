package analyzer

import "github.com/fendipratama/big-o-analyzer/internal/domain"

type Engine struct {
	Analyzers []domain.LanguageAnalyzer
}

func (e *Engine) Analyze(code string) domain.Result {
	var result domain.Result

	// 1. Detect & analyze language
	for _, a := range e.Analyzers {
		if a.Detect(code) {
			result = a.Analyze(code)
			break
		}
	}

	// 2. Fallback jika tidak terdeteksi
	if result.Language == "" {
		result = domain.Result{
			Language:        "Unknown",
			TimeComplexity:  "Unknown",
			SpaceComplexity: "Unknown",
			Confidence:      0,
			Reasoning:       []string{"Language not recognized"},
		}
	}

	// 3. ✅ ENRICHMENT LAYER (INI KUNCI)
	result.Explanation, result.Impact =
		Explain(result.TimeComplexity, result.SpaceComplexity)

	result.Optimizations =
		Recommend(result.TimeComplexity, result.Reasoning)

	result.EstimatedExecution =
		EstimateExecutions(result.TimeComplexity, []int{1000})

	return result
}
