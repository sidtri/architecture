package main

import (
	"fmt"
	"strings"

	"architecture.com/m/utils"
)

type Display struct {
	node  *utils.Node
	stack Stack
}

type Page struct {
	Items []string
}

func (page *Page) Serialize() string {
	return strings.Join(page.Items, ", ")
}

func (page *Page) Deserialize(str string) {
	page.Items = []string{}

	data := strings.Split(str, ",")
	page.Items = append(page.Items, data...)
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

func (d *Display) NextPage() Page {
	var page Page

	for i := 0; i < 10; i++ {
		if d.HasNext() {
			item := d.NextItem()
			page.Items = append(page.Items, item.Value)
		}
	}

	return page
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
	page := display.NextPage()

	fmt.Println(page.Serialize())
}
