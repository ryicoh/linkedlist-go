package linkedlist

type Orderable interface {
	Less(than Orderable) bool
	Equal(other Orderable) bool
}

type LinkedList[K Orderable, V any] struct {
	size int
	head *node[K, V]
}

type node[K Orderable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

func NewLinkedList[K Orderable, V any]() *LinkedList[K, V] {
	return &LinkedList[K, V]{
		size: 0,
		head: nil,
	}
}

func (ll *LinkedList[K, V]) Get(key K) (*V, bool) {
	n, eq := ll.getNodeLessOrEqual(key)
	if eq {
		return &n.value, true
	}

	return nil, false
}

func (ll *LinkedList[K, V]) getNodeLessOrEqual(key K) (n *node[K, V], eq bool) {
	cur := ll.head
	if cur == nil {
		return nil, false
	}

	for cur.key.Less(key) {
		if cur.next == nil {
			return cur, false
		}
		cur = cur.next
	}

	if cur.key.Equal(key) {
		return cur, true
	}

	return cur.prev, false
}

func (ll *LinkedList[K, V]) Set(key K, value V) {
	n, eq := ll.getNodeLessOrEqual(key)
	if eq {
		n.value = value
		return
	}

	if n == nil {
		ll.head = &node[K, V]{
			key:   key,
			value: value,
			prev:  nil,
			next:  ll.head,
		}
		ll.size++
		return
	}

	newNode := &node[K, V]{
		key:   key,
		value: value,
		prev:  n,
		next:  n.next,
	}

	if n.next != nil {
		n.next.prev = newNode
	}
	n.next = newNode

	ll.size++
}
