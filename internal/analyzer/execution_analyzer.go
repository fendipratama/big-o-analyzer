package analyzer

import (
	"fmt"

	"github.com/fendipratama/big-o-analyzer/internal/domain"
)

// EstimateExecutions mengubah Big-O menjadi estimasi waktu eksekusi nyata
func EstimateExecutions(timeC string, inputs []int) []domain.EstimatedExecution {
	results := []domain.EstimatedExecution{}

	for _, n := range inputs {
		var operations float64

		switch timeC {
		case "O(1)":
			operations = 1
		case "O(n)":
			operations = float64(n)
		case "O(n^2)":
			operations = float64(n * n)
		case "O(n^3)":
			operations = float64(n * n * n)
		default:
			continue
		}

		// asumsi 50 nanosecond per operasi
		timeMs := operations * 50 / 1_000_000

		results = append(results, domain.EstimatedExecution{
			InputSize:  fmt.Sprintf("n=%d", n),
			Operations: fmt.Sprintf("%.0f", operations),
			TimeMs:     timeMs,
		})
	}

	return results
}
