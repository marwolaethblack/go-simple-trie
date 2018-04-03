package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marwolaethblack/webdev/express"
)

type Node struct {
	Value    rune
	Children map[rune]*Node
	End      bool
	Depth    int
}

type Tree struct {
	Root *Node
	Size int
}

func NewTree() *Tree {
	tree := &Tree{Root: &Node{0, map[rune]*Node{}, false, 0}, Size: 0}
	return tree
}

func (t *Tree) AddNode(node *Node, Value rune) *Node {

	fmt.Println(node.Children[Value])

	if node.Children[Value] == nil {
		newnode := &Node{Value, map[rune]*Node{}, true, node.Depth + 1}
		node.Children[Value] = newnode
		t.Size++
		node.End = false
		return newnode
	}

	return node.Children[Value]

}

func (t *Tree) AddWord(word string) {
	currentNode := t.Root
	for _, letter := range word {
		currentNode = t.AddNode(currentNode, letter)
	}
}

func main() {
	tree := NewTree()
	word := "hello"

	tree.AddWord(word)

	j, _ := json.Marshal(tree)
	fmt.Println(string(j))

	var app express.App

	h := func(w http.ResponseWriter, req *http.Request, stop func(message string)) {
		express.GzipJSON(w, tree)
	}

	app.Get("/", h)

	app.Run(":8080")
}
