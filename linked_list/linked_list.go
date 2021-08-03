package linkedlist

type LinkedList struct {
	count int
	head  *Node
	tail  *Node
}

type Node struct {
	next *Node
	val  interface{}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		count: 0,
		head:  nil,
		tail:  nil,
	}
}

func NewNode(next *Node, value interface{}) *Node {
	return &Node{
		next: next,
		val:  value,
	}
}

func (ll *LinkedList) Add(elem interface{}) {
	node := NewNode(nil, elem)
	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		ll.tail = node
	}

	ll.count++
}

func (ll *LinkedList) Remove(el interface{}) (int, bool) {
	current := ll.head
	var previous *Node
	index := 0
	for current != nil {
		if current.val == el {
			if previous != nil {
				previous.next = current.next

				if current.next == nil {
					ll.tail = previous
				}
			} else {
				ll.head = current.next
				if ll.head == nil {
					ll.tail = nil
				}
			}
			ll.count--
			return index, true
		}
		index++
		previous = current
		current = current.next
	}
	return index, false
}
