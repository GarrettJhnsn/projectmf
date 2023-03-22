package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"favicon", "/favicon", "GET", []postData{}, http.StatusOK},
	{"login", "/login", "GET", []postData{}, http.StatusOK},
	{"packages", "/packages", "GET", []postData{}, http.StatusOK},
	{"consultation", "/consultation", "GET", []postData{}, http.StatusOK},
	{"confirmation", "/confirmation", "GET", []postData{}, http.StatusOK},
	{"tos", "/tos", "GET", []postData{}, http.StatusOK},
	{"consultation-request", "/confirmation", "POST", []postData{
		{key: "first_name", value: "Garrett"},
		{key: "last_name", value: "Johnson"},
		{key: "email", value: "garrettjhns0n8@gmail.com"},
		{key: "phone_number", value: "8312566826"},
		{key: "appointment_date", value: "2023-04-01"},
		{key: "appointment_time", value: "12:00PM"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, i := range theTests {
		if i.method == "GET" {
			response, err := ts.Client().Get(ts.URL + i.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if response.StatusCode != i.expectedStatusCode {
				t.Errorf("for %s, expeced %d but got %d", i.name, i.expectedStatusCode, response.StatusCode)
			}
		} else {
			values := url.Values{}

			for _, j := range i.params {
				values.Add(j.key, j.value)
			}
			response, err := ts.Client().Get(ts.URL + i.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if response.StatusCode != i.expectedStatusCode {
				t.Errorf("for %s, expeced %d but got %d", i.name, i.expectedStatusCode, response.StatusCode)
			}
		}
	}
}
