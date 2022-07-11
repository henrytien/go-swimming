package main

import (
	"net/http"
	"net/http/httptest"
    "log"
	"testing"
)

func TestIndexHander(t *testing.T) {
	req, err := http.NewRequest("Get", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req) // https://stackoverflow.com/questions/49668070/how-does-servehttp-work

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

	expected := "Hello, World!"
    log.Printf(rr.Body.String())
	if rr.Body.String() != expected {
		t.Errorf(
			"unexpected body: got (%v) want (%v)",
			rr.Body.String(),
			"Hello, World!",
		)
	}
}

func TestIndexHanderNotFound(t *testing.T) {
    req, err := http.NewRequest("Get", "/404", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(indexHandler)
    log.Printf("req %v", req)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusNotFound {
        t.Errorf(
            "unexpected status: got (%v) want (%v)",
            status,
            http.StatusNotFound,
        )
    }
}
