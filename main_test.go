package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kenlomaxhybris/goworkshopII/router"
)

/*
 200 – OK – Eyerything is working
201 – OK – New resource has been created
204 – OK – The resource was successfully deleted

304 – Not Modified – The client can use cached data

400 – Bad Request – The request was invalid or cannot be served. The exact error should be explained in the error payload. E.g. „The JSON is not valid“
401 – Unauthorized – The request requires an user authentication
403 – Forbidden – The server understood the request, but is refusing it or the access is not allowed.
404 – Not found – There is no resource behind the URI.
422 – Unprocessable Entity – Should be used if the server cannot process the enitity, e.g. if an image cannot be formatted or mandatory fields are missing in the payload.

500 – Internal Server Error – API developers should avoid this error. If an error occurs in the global catch blog, the stracktrace should be logged and not returned as response.
*/

func request(verb string, url string, payload string) (int, string /*, models.Workshop, []models.Workshop*/) {
	p := strings.NewReader(payload)
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(verb, url, p)
	r.ServeHTTP(rr, req)
	return rr.Code, strings.TrimSpace(rr.Body.String())
}

var r = router.InitRouter()

var tests = []struct {
	method  string
	url     string
	payload string
}{
	{"POST", "/", ""},
	{"POST", "/workshops", `{"Presenter":"P1","Title":"T1"}`},
	{"POST", "/workshops", `{"Presenter":"P2","Title":"T2"}`},
	{"POST", "/workshops", `{}`},
	{"POST", "/workshops", `{"Presenter":"This is too long, and won't fit as we have a maximum lenght of 100 characters, and this comes in at over 100","Title":"T2"}`},
	{"POST", "/workshops", `{"Presenter":"Some presenter","Title":"This is too long, and wonÄt fit as we have a maximum lenght of 100 characters, and this comes in at over 100"}`},
	{"GET", "/workshops", ``},
	{"GET", "/workshops/0", ``},
	{"GET", "/workshops/7", ``},
	{"PUT", "/workshops/1", `{"Presenter":"P11","Title":"T11"}`},
	{"PUT", "/workshops/7", `{"Presenter":"P11","Title":"T11"}`},
	{"PUT", "/workshops/1", `{}`},
	{"DELETE", "/workshops/0", ``},
	{"DELETE", "/workshops/7", ``},
}

func Example() {
	for _, t := range tests {
		code, body := request(t.method, t.url, t.payload)
		fmt.Printf("%s %s %s -> HTTP Status: %s(%d), Body: %s\n", t.method, t.url, t.payload, http.StatusText(code), code, body)
	}

	// Output:
	// POST /  -> HTTP Status: Not Found(404), Body: 404 page not found
	// POST /workshops {"Presenter":"P1","Title":"T1"} -> HTTP Status: OK(200), Body: {"ID":0,"Presenter":"P1","Title":"T1"}
	// POST /workshops {"Presenter":"P2","Title":"T2"} -> HTTP Status: OK(200), Body: {"ID":1,"Presenter":"P2","Title":"T2"}
	// POST /workshops {} -> HTTP Status: Unprocessable Entity(422), Body: {"error":"Missing/too much Data"}
	// POST /workshops {"Presenter":"This is too long, and won't fit as we have a maximum lenght of 100 characters, and this comes in at over 100","Title":"T2"} -> HTTP Status: Unprocessable Entity(422), Body: {"error":"Missing/too much Data"}
	// POST /workshops {"Presenter":"Some presenter","Title":"This is too long, and wonÄt fit as we have a maximum lenght of 100 characters, and this comes in at over 100"} -> HTTP Status: Unprocessable Entity(422), Body: {"error":"Missing/too much Data"}
	// GET /workshops  -> HTTP Status: OK(200), Body: [{"ID":0,"Presenter":"P1","Title":"T1"},{"ID":1,"Presenter":"P2","Title":"T2"}]
	// GET /workshops/0  -> HTTP Status: OK(200), Body: {"ID":0,"Presenter":"P1","Title":"T1"}
	// GET /workshops/7  -> HTTP Status: Not Found(404), Body: {"error":"ID 7 not found"}
	// PUT /workshops/1 {"Presenter":"P11","Title":"T11"} -> HTTP Status: OK(200), Body: {"ID":1,"Presenter":"P2","Title":"T2"}
	// PUT /workshops/7 {"Presenter":"P11","Title":"T11"} -> HTTP Status: Not Found(404), Body: {"error":"ID 7 not found"}
	// PUT /workshops/1 {} -> HTTP Status: Unprocessable Entity(422), Body: {"error":"Missing/too much Data"}
	// DELETE /workshops/0  -> HTTP Status: OK(200), Body: ""
	// DELETE /workshops/7  -> HTTP Status: Not Found(404), Body: {"error":"ID 7 not found"}
}

func TestConcurrency(t *testing.T) {
	for i := 0; i < 500; i++ {
		for _, t := range tests {
			go request(t.method, t.url, t.payload)
		}
	}
}
func Benchmark(b *testing.B) {
	var tests = []struct {
		method  string
		url     string
		payload string
	}{
		{"POST", "/workshops", `{"Presenter":"P1","Title":"T1"}`},
		{"POST", "/workshops", `{"Presenter":"P2","Title":"T2"}`},
		{"POST", "/workshops", `{"Presenter":"This is too long, and won't fit as we have a maximum lenght of 100 characters, and this comes in at over 100","Title":"T2"}`},
		{"POST", "/workshops", `{"Presenter":"Some presenter","Title":"This is too long, and wonÄt fit as we have a maximum lenght of 100 characters, and this comes in at over 100"}`},
		{"GET", "/workshops", ``},
		{"GET", "/workshops/0", ``},
		{"GET", "/workshops/7", ``},
		{"PUT", "/workshops/1", `{"Presenter":"P11","Title":"T11"}`},
		{"PUT", "/workshops/7", `{"Presenter":"P11","Title":"T11"}`},
		{"DELETE", "/workshops/0", ``},
		{"DELETE", "/workshops/7", ``},
	}

	for i := 0; i < b.N; i++ {
		t := tests[rand.Intn(9)]
		request(t.method, t.url, t.payload)
		//log.Printf("%s %s %s -> \n\tHTTP Status: %s(%d), Body: %s\n", t.method, t.url, t.payload, http.StatusText(code), code, body)
	}
}
