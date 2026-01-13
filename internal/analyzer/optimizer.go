// internal/analyzer/optimizer.go
package analyzer

func Recommend(timeC string, reasoning []string) []string {
	recs := []string{}

	if timeC == "O(n^2)" {
		recs = append(recs,
			"Consider reducing nested loops by using a hash map or set",
			"Check if pre-processing data can avoid repeated comparisons",
		)
	}

	if timeC == "O(n)" {
		recs = append(recs,
			"Ensure loop body runs in constant time",
		)
	}

	return recs
}
