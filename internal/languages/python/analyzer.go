// internal/languages/python/analyzer.go
package python

import (
	"strings"

	"github.com/fendipratama/big-o-analyzer/internal/domain"
)

type PythonAnalyzer struct{}

func (PythonAnalyzer) Name() string { return "Python" }

func (PythonAnalyzer) Detect(code string) bool {
	return strings.Contains(code, "def ")
}

func (PythonAnalyzer) Analyze(code string) domain.Result {
	loops := strings.Count(code, "for ") + strings.Count(code, "while ")

	time := "O(1)"
	if loops == 1 {
		time = "O(n)"
	} else if loops > 1 {
		time = "O(n^2)"
	}

	return domain.Result{
		Language:        "Python",
		TimeComplexity:  time,
		SpaceComplexity: "O(1)",
		Confidence:      0.75,
		Reasoning:       []string{"Loop-based heuristic"},
	}
}
