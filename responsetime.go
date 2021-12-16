package responsetime

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		diff := time.Since(start)
		c.Append("X-Response-Time-Mili", fmt.Sprintf("%dms", diff.Milliseconds()))
		// c.Append("X-Response-Time-Micro", fmt.Sprintf("%vus", diff.Microseconds()))
		// c.Append("X-Response-Time-Nano", fmt.Sprintf("%vns", diff.Nanoseconds()))
		return err
	}
}
