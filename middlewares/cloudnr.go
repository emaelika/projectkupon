package middlewares

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FileMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request has a file
		if r.Method != "POST" || r.Header.Get("Content-Type") != "multipart/form-data" {
			next.ServeHTTP(w, r)
			return
		}

		// Parse the multipart form
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get the file from the form
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Read the file contents
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the file contents and filename to the response
		w.Header().Set("Content-Type", fileHeader.Header.Get("Content-Type"))
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileHeader.Filename))
		w.Write(fileBytes)
	})
}

func FileUpload(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check if the request has a file
		if c.Request().Method != "POST" {
			return next(c) // Pass the request to the next handler if not a file upload
		}

		// Parse the multipart form
		err := c.Request().ParseMultipartForm(32 << 20)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		// Get the file from the form
		file, fileHeader, err := c.Request().FormFile("image")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Process the uploaded file (read contents, perform validations, etc.)

		// Write the file contents and filename to the response (if applicable)
		// Read the file contents
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Create a map to store the file information
		fileInfo := map[string]interface{}{
			"filename": fileHeader.Filename,
			"contents": fileBytes,
		}

		// Set the file information in the context so it can be accessed by subsequent handlers
		c.Set("file", fileInfo)

		// Continue to the next handler
		return next(c)
	}
}
