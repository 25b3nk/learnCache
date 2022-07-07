package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/25b3nk/learnCache/cache"
	"github.com/gofiber/fiber/v2"
)

// Response structure from jsonplaceholder
type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const url = "https://jsonplaceholder.typicode.com/todos/"

func main() {
	app := fiber.New()
	cache.Cache.SetTTL(time.Duration(10 * time.Second))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server running")
	})

	app.Get("/:id", cache.VerifyCache, func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := http.Get(url + id)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		todo := Todo{}
		// Unmarshal the json into the structure
		parseErr := json.Unmarshal(body, &todo)
		if parseErr != nil {
			return parseErr
		}
		cache.Cache.Set(id, todo)
		return c.JSON(fiber.Map{"Data": todo})
	})

	// Server hosted on localhost:3000
	app.Listen(":3000")
}
