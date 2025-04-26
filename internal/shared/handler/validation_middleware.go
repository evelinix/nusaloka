package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BindAndValidate adalah helper middleware untuk binding + validasi JSON input
func BindAndValidate[T any](c *gin.Context) (*T, bool) {
	var req T

	if err := c.ShouldBindJSON(&req); err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError

		switch {
		case errors.Is(err, io.EOF):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request body tidak boleh kosong"})
		case errors.As(err, &syntaxErr):
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON tidak valid"})
		case errors.As(err, &unmarshalTypeErr):
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Tipe field '%s' salah", unmarshalTypeErr.Field)})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return nil, false
	}

	return &req, true
}
