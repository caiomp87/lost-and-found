package main

import (
	"log"

	"github.com/caiomp87/lost-and-found/src/handlers/user"
	"github.com/caiomp87/lost-and-found/src/modules/user/repositories"
	"github.com/caiomp87/lost-and-found/src/modules/user/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	userRepository := repositories.NewUserRepository(nil)
	userService := services.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	appv1 := app.Group("/v1/api")

	appUser := appv1.Group("/users")
	appUser.Post("/", userHandler.Create)

	log.Println("API run on port 3333!")

	log.Fatal(app.Listen(":3333"))
}
