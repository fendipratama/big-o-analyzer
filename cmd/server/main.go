// cmd/server/main.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/fendipratama/big-o-analyzer/internal/analyzer"
	"github.com/fendipratama/big-o-analyzer/internal/api"
	"github.com/fendipratama/big-o-analyzer/internal/domain"
	golang "github.com/fendipratama/big-o-analyzer/internal/languages/go"
	"github.com/fendipratama/big-o-analyzer/internal/languages/javascript"
	"github.com/fendipratama/big-o-analyzer/internal/languages/php"
	"github.com/fendipratama/big-o-analyzer/internal/languages/python"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:5173",
        AllowHeaders: "Origin, Content-Type, Accept",
        AllowMethods: "GET,POST,OPTIONS",
    }))

	engine := &analyzer.Engine{
		Analyzers: []domain.LanguageAnalyzer{
			javascript.JSAnalyzer{},
			php.PHPAnalyzer{},
			python.PythonAnalyzer{},
			golang.GoAnalyzer{},
		},
	}

	app.Post("/analyze", api.AnalyzeHandler(engine))
	app.Listen(":3000")
}
