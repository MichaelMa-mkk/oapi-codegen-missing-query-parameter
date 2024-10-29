// Package main handles calls to the team app API
package main

import (
	"context"
	"errors"
	api "oapi-codegen-sample/cmd/api/gen"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=gen/types.cfg.yaml ./api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=gen/server.cfg.yaml ./api.yaml

type Server struct{}

// (GET /ping)
func (Server) GetPing(_ context.Context, request api.GetPingRequestObject) (api.GetPingResponseObject, error) {
	if request.Params.Option == "error" {
		return nil, errors.New("customised internal server error")
	}
	return api.GetPing200JSONResponse{Ping: "pong"}, nil
}

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := Server{}

	app := fiber.New()

	api.RegisterHandlers(app, api.NewStrictHandler(server, nil))

	// And we serve HTTP until the world ends.
	log.Fatal().Err(app.Listen("0.0.0.0:8080")).Msg("Failed to start server")
}
