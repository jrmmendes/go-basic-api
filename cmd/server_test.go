package main

import (
	"api/pkg/core"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	Message string `json:"message"`
}

func setupTestEnv() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func TestGetHelloMessage(t *testing.T) {
	setupTestEnv()

	app := core.Bootstrap()

	request, _ := http.NewRequest("GET", "/", nil)

	response, error := app.Test(request)

	assert.Nil(t, error)
	assert.Equal(t, 200, response.StatusCode)

	body, error := io.ReadAll(response.Body)

	var data Response

	json.Unmarshal(body, &data)

	assert.Equal(t, "your first json response", data.Message)
}
