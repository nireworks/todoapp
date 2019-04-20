package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleAddTodo(t *testing.T) {
	srv := server.server{
		db:    mockDatabase,
		email: mockEmailSender,
	}
	srv.routes()

	req, err := http.NewRequest("GET", "/about", nil)
	is.NoErr(err)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	is.Equal(w.StatusCode, http.StatusOK)
}
