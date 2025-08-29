package utils

// Create linked list
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (ll *LinkedList) Append(data interface{}) {
	newNode := &Node{Data: data}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		newNode.Prev = ll.Tail
		ll.Tail = newNode
	}
	ll.Length++
}

func (ll *LinkedList) Prepend(data interface{}) {
	newNode := &Node{Data: data}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		newNode.Next = ll.Head
		ll.Head.Prev = newNode
		ll.Head = newNode
	}
	ll.Length++
}

func (ll *LinkedList) Remove(data interface{}) {
	current := ll.Head

	for current != nil {
		if current.Data == data {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				ll.Head = current.Next
			}

			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				ll.Tail = current.Prev
			}

			ll.Length--
			return
		}
		current = current.Next
	}
}

func (ll *LinkedList) Find(data interface{}) *Node {
	current := ll.Head

	for current != nil {
		if current.Data == data {
			return current
		}
		current = current.Next
	}
	return nil
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Length == 0
}

func (ll *LinkedList) Display() []interface{} {
	var elements []interface{}
	current := ll.Head

	for current != nil {
		elements = append(elements, current.Data)
		current = current.Next
	}
	return elements
}

