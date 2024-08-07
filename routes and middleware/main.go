package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//middlware
	/*
		We can store the messages returned by the Logger function in a database.
		Handling Common Middleware Tasks
		Middleware in Fiber can be used to handle a wide range of common tasks. Let’s explore some of the tasks that are commonly handled using middleware:

		Authentication: Middleware can be used to authenticate users before allowing them access to certain routes. You can check user credentials, verify tokens, or implement any authentication logic.

		Logging: Middleware functions are ideal for logging requests, responses, and application events. Logging helps in debugging, monitoring, and analyzing the application’s behavior.

		Request Parsing: Middleware can preprocess and parse incoming requests, such as extracting data from request bodies or headers.

		Authorization: Similar to authentication, authorization middleware can determine whether a user has the necessary permissions to access a specific route.

		CORS (Cross-Origin Resource Sharing): Middleware can handle CORS headers and ensure secure cross-origin requests.

		Compression: Middleware can compress responses to reduce bandwidth and improve application performance.

		Error Handling: Middleware can catch and handle errors that occur during the request-response cycle, providing consistent error responses to clients.

		Rate Limiting: Middleware can implement rate limiting to control the number of requests a client can make within a certain time frame.

		By using middleware, you can modularize and structure your application’s code effectively, making it more maintainable and readable.
	*/
	app.Use(Logger)

	//basic root directory
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	//about directory
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("About Fiber!")
	})
	/*
		I created two directories here: one represents the root directory /, and the other represents the about directory.
	*/

	//dynamic directory
	/*
		Dynamic routing is frequently preferred in areas such as user operations and database transactions."
	*/
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		userID := c.Params("id")
		return c.SendString("User ID: " + userID)
	})

	app.Get("/protected", Logger, func(c *fiber.Ctx) error {
		return c.SendString("This route is protected by the logger func.")
	})

	app.Listen(":3000")

}

func Logger(c *fiber.Ctx) error {
	println("Middleware: Request Received.")
	return c.Next()
}
