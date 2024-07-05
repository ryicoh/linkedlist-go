package linkedlist

import (
	"fmt"
	"strings"
	"testing"
)

type key string

func (k key) Less(than Orderable) bool {
	return strings.Compare(string(k), string(than.(key))) < 0
}

func (k key) Equal(other Orderable) bool {
	return string(k) == string(other.(key))
}

func TestGetNodeLessOrEqual(t *testing.T) {
	testCases := []struct {
		head     *node[key, string]
		key      key
		expected key
		found    bool
	}{
		{&node[key, string]{}, "k1", "", false},
		{&node[key, string]{key: "k1"}, "k1", "k1", true},
		{&node[key, string]{key: "k1"}, "k0", "", false},
		{&node[key, string]{key: "k1"}, "k2", "k1", false},
		{&node[key, string]{key: "k1", next: &node[key, string]{key: "k3"}}, "k1", "k1", true},
		{&node[key, string]{key: "k1", next: &node[key, string]{key: "k3"}}, "k3", "k3", true},
		{&node[key, string]{key: "k1", next: &node[key, string]{key: "k3"}}, "k0", "", false},
		{&node[key, string]{key: "k1", next: &node[key, string]{key: "k3", prev: &node[key, string]{key: "k1"}}}, "k2", "k1", false},
		{&node[key, string]{key: "k1", next: &node[key, string]{key: "k3", prev: &node[key, string]{key: "k1"}}}, "k4", "k3", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("head=%v,key=%s,expected=%s,found=%t", tc.head, tc.key, tc.expected, tc.found), func(t *testing.T) {
			ll := NewLinkedList[key, string]()
			ll.head = tc.head

			actual, found := ll.getNodeLessOrEqual(tc.key)
			if found != tc.found {
				t.Errorf("expected %t, got %t", tc.found, found)
			}

			expectKey := key("")
			if actual != nil {
				expectKey = actual.key
			}
			if !expectKey.Equal(tc.expected) {
				t.Errorf("expected %s, got %s", tc.expected, expectKey)
			}
		})
	}
}
