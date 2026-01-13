// internal/analyzer/explainer.go
package analyzer

func Explain(timeC, spaceC string) (string, string) {
	explanation := ""
	impact := ""

	switch timeC {
	case "O(1)":
		explanation = "The algorithm runs in constant time regardless of input size."
		impact = "Very fast and scalable."
	case "O(n)":
		explanation = "Execution time grows linearly with input size."
		impact = "Acceptable for medium to large data."
	case "O(n^2)":
		explanation = "Execution time grows quadratically due to nested iterations."
		impact = "May become slow for large datasets."
	default:
		explanation = "Time complexity could not be determined precisely."
		impact = "Performance impact is unclear."
	}

	return explanation, impact
}
