package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evelinix/nusaloka/internal/account/model"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
   // Setup router
   router := chi.NewRouter()

   // Sample input data
   newUser := model.User{
	   Password: "password123",
	   Email:    "testuser@example.com",
   }

   // Marshal user data into JSON
   requestBody, _ := json.Marshal(newUser)
   req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
   req.Header.Set("Content-Type", "application/json")

   // Recorder to capture the response
   recorder := httptest.NewRecorder()

   // Simulate request to the router
   router.ServeHTTP(recorder, req)

   // Assert status code and response body
   assert.Equal(t, http.StatusCreated, recorder.Code)
   var response map[string]string
   err := json.NewDecoder(recorder.Body).Decode(&response)
   assert.Nil(t, err)
   assert.Equal(t, "User registered successfully", response["message"])
}
