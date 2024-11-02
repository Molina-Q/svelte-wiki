package api

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	// the response code
	Code int

	// the response message
	Status string
}

type Error struct {
	// the error code
	Code int

	// the error message
	Message string
}