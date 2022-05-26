package cache

import (
	"sort"
	"testing"
	"time"
)

func TestGetExist(t *testing.T) {
	cache := NewCache()
	cache.Put("test_key", "test_value")
	want_value, want_ok := "test_value", true
	if got_value, got_ok := cache.Get("test_key"); got_value != want_value || got_ok != want_ok {
		t.Errorf("got %s (%t), want %s (%t)", got_value, got_ok, want_value, want_ok)
	}
}

func TestGetNotExist(t *testing.T) {
	cache := NewCache()
	cache.Put("test_key", "test_value")
	want_value, want_ok := "", false
	if got_value, got_ok := cache.Get("not_test_key"); got_value != want_value || got_ok != want_ok {
		t.Errorf("got %s (%t), want %s (%t)", got_value, got_ok, want_value, want_ok)
	}
}

func TestKeys(t *testing.T) {
	cache := NewCache()
	cache.Put("test_key1", "test_value1")
	cache.Put("test_key2", "test_value2")
	cache.Put("test_key3", "test_value3")
	want := []string{"test_key1", "test_key2", "test_key3"}
	if got := cache.Keys(); !Equal(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}

func Equal(a, b []string) bool {
	sort.Strings(a)
	sort.Strings(b)
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestEmpty(t *testing.T) {
	cache := NewCache()
	want_value, want_ok := "", false
	if got_value, got_ok := cache.Get("not_test_key"); got_value != want_value || got_ok != want_ok {
		t.Errorf("got %s (%t), want %s (%t)", got_value, got_ok, want_value, want_ok)
	}
}

func TestOverride(t *testing.T) {
	cache := NewCache()
	cache.Put("test_key1", "test_value1")
	cache.Put("test_key1", "test_value3")
	want_value, want_ok := "test_value3", true
	if got_value, got_ok := cache.Get("test_key1"); got_value != want_value || got_ok != want_ok {
		t.Errorf("got %s (%t), want %s (%t)", got_value, got_ok, want_value, want_ok)
	}
}

func TestDeadline(t *testing.T) {
	cache := NewCache()
	cache.Put("test_key1", "test_value1")
	cache.PutTill("test_key2", "test_value2", time.Now().Add(time.Second*2))
	cache.PutTill("test_key3", "test_value2", time.Now().Add(time.Second*4))
	time.Sleep(time.Second * 1)
	want := []string{"test_key1", "test_key2", "test_key3"}
	if got := cache.Keys(); !Equal(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
	time.Sleep(time.Second * 2)
	want = []string{"test_key1", "test_key3"}
	if got := cache.Keys(); !Equal(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}
