// internal/languages/javascript/analyzer.go
package javascript

import (
	"strings"

	"github.com/fendipratama/big-o-analyzer/internal/domain"
)

type JSAnalyzer struct{}

func (JSAnalyzer) Name() string {
	return "JavaScript"
}

func (JSAnalyzer) Detect(code string) bool {
	return strings.Contains(code, "function") || strings.Contains(code, "=>")
}

func (JSAnalyzer) Analyze(code string) domain.Result {
	loops := strings.Count(code, "for") + strings.Count(code, "while")

	time := "O(1)"
	reason := []string{}

	if loops == 1 {
		time = "O(n)"
		reason = append(reason, "Single loop detected")
	} else if loops > 1 {
		time = "O(n^2)"
		reason = append(reason, "Nested loops detected")
	}

	if len(reason) == 0 {
		reason = append(reason, "No loop detected")
	}

	return domain.Result{
		Language:        "JavaScript",
		TimeComplexity:  time,
		SpaceComplexity: "O(1)",
		Confidence:      0.8,
		Reasoning:       reason,
	}
}
