package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.ID != 1 {
		t.Errorf("cannot retrieve json post")
	}
}

func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	jsonBody := strings.NewReader(`{"content": "Updated post", "author": "Sau Sheong"}`)
	request, _ := http.NewRequest("POST", "/post/1", jsonBody)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusCreated {
		t.Errorf("response code is %v", writer.Code)
	}
}