package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFibonacci(t *testing.T) {
	router := gin.New()
	router.GET("/fibonacci", Fibonacci)

	t.Run("EmptyQuery", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/fibonacci", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		if resp.Code != 400 {
			t.Fatalf("should get %d but got %d", 400, resp.Code)
		}
	})
	t.Run("IncorrectQuery", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/fibonacci", nil)
		resp := httptest.NewRecorder()
		qs := req.URL.Query()
		qs.Set("number", "abc")
		req.URL.RawQuery = qs.Encode()
		router.ServeHTTP(resp, req)
		if resp.Code != 400 {
			t.Fatalf("should get %d but got %d", 400, resp.Code)
		}
	})
	t.Run("OK", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/fibonacci", nil)
		resp := httptest.NewRecorder()
		qs := req.URL.Query()
		qs.Set("number", "5")
		req.URL.RawQuery = qs.Encode()
		router.ServeHTTP(resp, req)
		if resp.Code != 200 {
			t.Fatalf("should get %d but got %d", 200, resp.Code)
		}
	})
}
