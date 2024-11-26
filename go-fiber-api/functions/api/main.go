package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// Handler function for the serverless environment
func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	app := fiber.New()

	// Define your routes
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, Netlify!"})
	})

	// Use Fiber as an AWS Lambda handler
	return adaptor.HTTPHandlerFunc(app.Handler())(req)
}

func main() {
	lambda.Start(handler)
}