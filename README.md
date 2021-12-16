# fiber-responsetime

X-Response-Time middleware for fiber/v2

```go
package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	responsetime "github.com/husanu/fiber-responsetime"
)

func main() {
	app := fiber.New()

	app.Use(responsetime.New())

	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(666 * time.Millisecond)
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
```

```bash
curl -i http://localhost:3000/
HTTP/1.1 200 OK
Date: Thu, 16 Dec 2021 10:04:13 GMT
Content-Type: text/plain; charset=utf-8
Content-Length: 13
X-Response-Time: 670ms

Hello, World!
```
