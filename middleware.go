package microfiber

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger.Info(
			"request",
			zap.String("ip", c.IP()),
			zap.String("method", c.Method()),
		)
		return c.Next()
	}
}
