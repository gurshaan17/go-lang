package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateLimiting   int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "can't parse json"})
	}
	//implement rate limit & check if input is an actual URL

	if !govalidator.IsURL(body.URL) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid URL"})
	}
	// check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(503).JSON(fiber.Map{"error": "you can't access it"})
	}
	// enforce https, ssl
	body.URL = helpers.EnforceHTTP(body.URL)
}
