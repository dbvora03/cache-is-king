package cachinghandler

import (
	"context"
	"net/http"
	"time"

	unitmodel "cacheisking.com/models"
	redisconnection "cacheisking.com/utils/redis"
	fiber "github.com/gofiber/fiber/v2"
	redis "github.com/go-redis/redis/v8"
)
var ctx = context.Background()


func SetCache(c *fiber.Ctx) error {

	unit_instance := new(unitmodel.Unit)

	if err := c.BodyParser(unit_instance) ; err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return c.Status(400).JSON(err)
	}

	var key string = unit_instance.Project + "/" + unit_instance.UnitID

	err := redisconnection.RDB.Set(ctx, key, unit_instance.Value, time.Duration(unit_instance.Duration)).Err()
	if err != nil {
		return c.Status(400).JSON(err)
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"unit": unit_instance,
		"message": "Added to cache successfully",
	})

	return nil
}

func SetPermCache(c *fiber.Ctx) error {

	unit_instance := new(unitmodel.UnitNoTime)

	if err := c.BodyParser(unit_instance) ; err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return c.Status(400).JSON(err)
	}

	var key string = unit_instance.Project + "/" + unit_instance.UnitID

	err := redisconnection.RDB.Set(ctx, key, unit_instance.Value, 0).Err()
	if err != nil {
		return c.Status(400).JSON(err)
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"unit": unit_instance,
		"message": "Added to cache successfully",
	})

	return nil
}

func GetCache(c *fiber.Ctx) error {

	unit_instance := new(unitmodel.UnitReq)

	if err := c.BodyParser(unit_instance) ; err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return c.Status(400).JSON(err)
	}

	var key string = unit_instance.Project + "/" + unit_instance.UnitID

	val, err := redisconnection.RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		return c.Status(400).JSON(err)
	} else if err != nil {
		return c.Status(400).JSON(err)
	} else {
		c.Status(200).JSON(&fiber.Map{
			"success": true,
			"value": val,
			"message": "Added to cache successfully",
		})
	}

	return nil
}