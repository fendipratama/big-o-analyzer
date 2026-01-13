// internal/domain/language.go
package domain

type LanguageAnalyzer interface {
	Name() string
	Detect(code string) bool
	Analyze(code string) Result
}
