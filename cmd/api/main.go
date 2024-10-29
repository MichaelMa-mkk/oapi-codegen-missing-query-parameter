// Package main handles calls to the team app API
package main

import (
	"log"
	"net/http"
	api "oapi-codegen-sample/cmd/api/gen"

	"github.com/gofiber/fiber/v2"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=gen/types.cfg.yaml ./api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=gen/server.cfg.yaml ./api.yaml

type Server struct{}

// (GET /ping)
func (Server) GetPing(ctx *fiber.Ctx, _ api.GetPingParams) error {
	resp := api.Pong{
		Ping: "pong",
	}

	return ctx.
		Status(http.StatusOK).
		JSON(resp)
}

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := Server{}

	app := fiber.New()

	api.RegisterHandlers(app, server)

	// And we serve HTTP until the world ends.
	log.Fatal(app.Listen("0.0.0.0:8080"))
}
