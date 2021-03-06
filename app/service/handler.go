package service

import (
	"net/http"

	"github.com/go-chi/render"
)

// AppHandler is a type that implements http.Handler and makes handling
// errors easier. When its method returns an error, it prints it to the logs
// and shows a JSON formatted error to the user.
type AppHandler func(http.ResponseWriter, *http.Request) APIError

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiErr := fn(w, r)
	if apiErr != NoError { // err is APIError, not os.Error
		if err := render.Render(w, r, apiErr.httpResponse()); err != nil {
			panic(err) // if unable to render errors, panic
		}
	}
}
