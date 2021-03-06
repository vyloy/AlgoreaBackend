package service

import (
	"net/http"

	"github.com/go-chi/render"
)

// Response is used for generating non-data responses, i.e. on error or on POST/PUT/PATCH/DELETE request
type Response struct {
	HTTPStatusCode int         `json:"-"`
	Success        bool        `json:"success"`
	Message        string      `json:"message"`
	Data           interface{} `json:"data,omitempty"`
}

// Render generates the HTTP response from Response
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if resp.Success && resp.Message == "" {
		resp.Message = "success"
	}
	render.Status(r, resp.HTTPStatusCode)
	return nil
}

// CreationSuccess generated a success response for a POST creation
func CreationSuccess(data interface{}) render.Renderer {
	return &Response{
		HTTPStatusCode: http.StatusCreated,
		Success:        true,
		Data:           data,
	}
}
