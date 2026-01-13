// internal/languages/php/analyzer.go
package php

import (
	"strings"

	"github.com/fendipratama/big-o-analyzer/internal/domain"
)

type PHPAnalyzer struct{}

func (PHPAnalyzer) Name() string { return "PHP" }

func (PHPAnalyzer) Detect(code string) bool {
	return strings.Contains(code, "<?php")
}

func (PHPAnalyzer) Analyze(code string) domain.Result {
	loops := strings.Count(code, "for") +
		strings.Count(code, "foreach") +
		strings.Count(code, "while")

	time := "O(1)"
	if loops == 1 {
		time = "O(n)"
	} else if loops > 1 {
		time = "O(n^2)"
	}

	return domain.Result{
		Language:        "PHP",
		TimeComplexity:  time,
		SpaceComplexity: "O(1)",
		Confidence:      0.7,
		Reasoning:       []string{"Loop-based heuristic"},
	}
}
