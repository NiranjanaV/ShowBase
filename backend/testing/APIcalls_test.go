package tests

import (
	// "backend-v1/src/config"

	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	A "main/functions/APIcalls"
	DM "main/functions/dbAccess"
)

/*
	This function is responsible for checking whether
	HomeMovies API returns a success status code.
*/
func TestGetHomePageAPIStatusSuccess(t *testing.T) {
	r := getRouter()
	// config.CreateConn()

	r.GET("/top", A.HomeMovies)
	req, _ := http.NewRequest("GET", "/top", nil)

	testHttpRequest(t, r, req, func(w *httptest.ResponseRecorder) bool {
		s := w.Code == http.StatusOK
		return s
	})
}

/*
	This function is responsible for checking whether
	Search API returns all some questions that
	are present in database
*/
func TestGetSearchAPIResponse(t *testing.T) {
	r := getRouter()
	// config.CreateConn()

	r.GET("/search/:name", A.Search)
	req, _ := http.NewRequest("GET", "/search/bean", nil)

	testHttpRequest(t, r, req, func(w *httptest.ResponseRecorder) bool {
		s := w.Code == http.StatusOK
		return s
	})
}

/*
	This function is responsible for checking whether
	a wrong route returns a failure status code.
*/
func TestSearchWithoutParmAPIStatusFailure(t *testing.T) {
	r := getRouter()
	// config.CreateConn()

	r.GET("/search/:name", A.Search)
	req, _ := http.NewRequest("GET", "/search", nil)

	testHttpRequest(t, r, req, func(w *httptest.ResponseRecorder) bool {
		s := w.Code == http.StatusNotFound
		return s
	})
}

func TestGetSearchWithPagesAPIResponse(t *testing.T) {
	r := getRouter()
	// config.CreateConn()

	jsonData := []byte(`{
		"text": "bean",
		"page": "3"
	}`)

	r.GET("/searchPage", A.SearchWithPage)
	req, _ := http.NewRequest("GET", "/searchPage", bytes.NewBuffer(jsonData))

	testHttpRequest(t, r, req, func(w *httptest.ResponseRecorder) bool {
		s := w.Code == http.StatusOK
		return s
	})
}

func TestEnterDataIntoUser(t *testing.T) {
	r := getRouter()
	// config.CreateConn()
	//suppose to be text
	jsonData := []byte(`{
		"username": "Srikant",   
		"password": "123456"
	}`)

	r.PUT("/userReg", DM.InsertAuthTable)
	req, _ := http.NewRequest("PUT", "/userReg", bytes.NewBuffer(jsonData))

	testHttpRequest(t, r, req, func(w *httptest.ResponseRecorder) bool {
		fmt.Println(w.Body)
		s := w.Code == http.StatusBadRequest
		return s
	})
}

func TestGetSearchWithPageswrongJSONAPIResponse(t *testing.T) {
	r := getRouter()
	// config.CreateConn()
	//suppose to be text
	jsonData := []byte(`{
		"movie": "bean",   
		"page": "3"
	}`)

	r.GET("/searchPage", A.SearchWithPage)
	req, _ := http.NewRequest("GET", "/searchPage", bytes.NewBuffer(jsonData))

	testHttpRequest(t, r, req, func(w *httptest.ResponseRecorder) bool {
		s := w.Code == http.StatusBadRequest
		return s
	})
}
