// / node has both linked list and graph support. should handle correct
package utils

type Node struct {
	Value string
	Data  interface{}
	Left  *Node
	Right *Node
	Next  *Node
	Prev  *Node
}
