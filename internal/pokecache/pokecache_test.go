package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(time.Second)
	cache.Add("test", []byte("test"))

	data, ok := cache.Get("test")
	if !ok {
		t.Errorf("expected to get data, but got nothing")
	}

	if string(data) != "test" {
		t.Errorf("expected to get 'test', but got '%s'", string(data))
	}
}
