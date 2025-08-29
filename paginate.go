package main

import (
	"fmt"

	"architecture.com/m/utils"
)

type Display struct {
	node  *utils.Node
	stack Stack
}

func (d *Display) PushAll(node *utils.Node) {
	d.stack.Push(node)

	if node.Left == nil {
		return
	} else {
		d.PushAll(node.Left)
	}
}

func (d *Display) NextItem() *utils.Node {
	_, item := d.stack.Pop()

	if item.Right != nil {
		d.PushAll(item.Right)
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
			data = append(data, item.Value)
		}
	}

	return data
}

func new(node *utils.Node) *Display {
	return &Display{node: node}
}

func Paginate() {
	var bst utils.BinarySearchTree

	names := []string{"Jeanette", "Latasha", "Elvira", "Caryl", "Antoinette", "Cassie", "Charity", "Lyn", "Elia", "Anya", "Albert", "Cherlyn", "Lala", "Kandice", "Iliana"}

	for _, name := range names {
		bst.Insert(name)
	}

	display := new(bst.Root)

	display.PushAll(display.node)
	for _, item := range display.NextPage() {
		fmt.Println(item)
	}
}
