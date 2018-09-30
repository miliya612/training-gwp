package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.ID != 1 {
		t.Error("wrong id, was expecting 1 but got", post.ID)
	}
	if post.Content != "Hello World!" {
		t.Error("wrong content, was expecting 'Hello World!' but got", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("skipping encoding for now")
}

func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second)
}
