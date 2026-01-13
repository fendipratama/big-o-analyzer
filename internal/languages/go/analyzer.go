// internal/languages/go/analyzer.go
package golang

import (
	"strings"

	"github.com/fendipratama/big-o-analyzer/internal/domain"
)

type GoAnalyzer struct{}

func (GoAnalyzer) Name() string { return "Go" }

func (GoAnalyzer) Detect(code string) bool {
	return strings.Contains(code, "package ") && strings.Contains(code, "func ")
}

func (GoAnalyzer) Analyze(code string) domain.Result {
	loops := strings.Count(code, "for ")

	time := "O(1)"
	if loops == 1 {
		time = "O(n)"
	} else if loops > 1 {
		time = "O(n^2)"
	}

	return domain.Result{
		Language:        "Go",
		TimeComplexity:  time,
		SpaceComplexity: "O(1)",
		Confidence:      0.85,
		Reasoning:       []string{"for-loop detected"},
	}
}
