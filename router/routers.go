package router

import (
	"KanbanBoardAPI/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router() *fiber.App {

	r := fiber.New()

	api := r.Group("/api") // /api

	users := api.Group("/users")    // /api/users

	users.Post("/register", controllers.Register)         // POST => /api/users/register
	//users.Post("/login")            // POST => /api/users/login
	//users.Put("/update-account")    // PUT => /api/users/update-account
	//users.Delete("/delete-account") // DELETE => /api/users/delete-account
	//
	//
	//categories := api.Group("/categories") // /api/categories
	//
	//categories.Post("/")                   // POST => /api/categories/
	//categories.Get("/")                    // GET => /api/categories/
	//categories.Patch("/:category_id")      // PATCH => /api/categories/:category_id
	//categories.Delete("/:category_id")     // DELETE => /api/categories/:category_id
	//
	//
	//tasks := api.Group("/tasks")              // /api/tasks
	//
	//tasks.Post("/")                           // POST => /api/tasks/
	//tasks.Get("/")                            // GET => /api/tasks/
	//tasks.Put("/:tasks_id")                   // PUT => /api/tasks/:tasks_id
	//tasks.Patch("/:tasks_id")                 // PATCH => /api/tasks/:tasks_id
	//tasks.Patch("/update-category/:tasks_id") // PATCH => /api/tasks/update-category/:tasks_id
	//tasks.Delete("/:tasks_id")                //DELETE => /api/tasks/:tasks_id

	return r

}
