package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_out "github.com/input-api/mcp-server/tools/out"
	tools_route "github.com/input-api/mcp-server/tools/route"
	tools_flights "github.com/input-api/mcp-server/tools/flights"
	tools_in "github.com/input-api/mcp-server/tools/in"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_out.CreateGet_out_airport_airportidTool(cfg),
		tools_route.CreateGet_route_source_sourceairportid_dest_destinationairportidTool(cfg),
		tools_flights.CreateGet_flightsTool(cfg),
		tools_in.CreateGet_in_airport_airportidTool(cfg),
	}
}
