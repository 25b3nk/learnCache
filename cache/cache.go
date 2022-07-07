package cache

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jellydator/ttlcache/v2"
)

var Cache ttlcache.SimpleCache = ttlcache.NewCache()

func VerifyCache(c *fiber.Ctx) error {
	id := c.Params("id")
	val, err := Cache.Get(id)
	if err != ttlcache.ErrNotFound {
		// If we have the data in cache we return it
		return c.JSON(fiber.Map{"cached": val})
	}
	// If we do not have data in cache we go to the next function in the route
	return c.Next()
}
