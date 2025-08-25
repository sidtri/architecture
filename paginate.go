package main

import "fmt"

type Display struct {
	node  *Node
	stack Stack
}

func (d *Display) PushAll(node *Node) {
	d.stack.Push(node)

	if node.left == nil {
		return
	} else {
		d.PushAll(node.left)
	}
}

func (d *Display) NextItem() *Node {
	_, item := d.stack.Pop()

	if item.right != nil {
		d.PushAll(item.right)
	}

	return item
}

func (d *Display) HasNext() bool {
	return len(d.stack.items) > 0
}

func (d *Display) NextPage() []string {
	var data []string

	for i := 0; i < 10; i++ {
		if d.HasNext() {
			item := d.NextItem()
			data = append(data, item.value)
		}
	}

	return data
}

func new(node *Node) *Display {
	return &Display{node: node}
}

func Paginate() {
	var bst BinarySearchTree

	names := []string{"Jeanette", "Latasha", "Elvira", "Caryl", "Antoinette", "Cassie", "Charity", "Lyn", "Elia", "Anya", "Albert", "Cherlyn", "Lala", "Kandice", "Iliana"}

	for _, name := range names {
		bst.Insert(name)
	}

	display := new(bst.root)

	display.PushAll(display.node)
	for _, item := range display.NextPage() {
		fmt.Println(item)
	}
}
