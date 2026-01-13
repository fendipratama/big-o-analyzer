// internal/api/handler.go
package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/fendipratama/big-o-analyzer/internal/analyzer"
)

type analyzeRequest struct {
	Code string `json:"code"`
}

func AnalyzeHandler(engine *analyzer.Engine) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req analyzeRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		result := engine.Analyze(req.Code)
		return c.JSON(result)
	}
}
