// internal/languages/language.go
package languages

import "github.com/fendipratama/big-o-analyzer/internal/analyzer"

type LanguageAnalyzer interface {
	Name() string
	Detect(code string) bool
	Analyze(code string) analyzer.Result
}
