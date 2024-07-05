package linkedlist_test

import (
	"strings"
	"testing"

	linkedlist "github.com/ryicoh/linkedlist-go"
)

type key string

func (k key) Less(than linkedlist.Orderable) bool {
	return strings.Compare(string(k), string(than.(key))) < 0
}

func (k key) Equal(other linkedlist.Orderable) bool {
	return string(k) == string(other.(key))
}

func TestLinkedList(t *testing.T) {
	ll := linkedlist.NewLinkedList[key, string]()
	ll.Set("k1", "v1")

	expectGet(t, ll, "k1", "v1")

	ll.Set("k3", "v3")
	ll.Set("k2", "v2")

	expectGet(t, ll, "k1", "v1")
	expectGet(t, ll, "k3", "v3")
	expectGet(t, ll, "k2", "v2")
}

func expectGet(t *testing.T, ll *linkedlist.LinkedList[key, string], key key, expected string) {
	t.Helper()

	real, ok := ll.Get(key)
	if !ok {
		t.Fatalf("expected true, got false")
	}
	if *real != expected {
		t.Fatalf("expected %s, got %s", expected, *real)
	}
}
