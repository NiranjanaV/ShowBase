package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

/*
	This function is responsible for turning the test mode
	before executing the tests.
	Parameters: m testing.M
	Returns:
*/
func ChangeEnvironmentMode(m *testing.M) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Proceed with the testing process
	os.Exit(m.Run())
}

/*
	This function is responsible for creating a new Gin Engine
	object.
	Parameters:
	Returns: r *gin.Engine
*/
func getRouter() *gin.Engine {
	r := gin.Default()

	return r
}

/*
	This function acts as a Helper function to process a request and
	record its response to test the routes.
	Parameters: t testing.T, r gin.Engine, request http.Request
	Returns:
*/
func testHttpRequest(t *testing.T, r *gin.Engine, request *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create a new ResponseRecorder object
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, request)

	if !f(rec) {
		t.Fail()
	}
}
