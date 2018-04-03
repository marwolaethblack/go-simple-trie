package simpletrie

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

//NewTree creates a new Radix Tree/Trie
func NewTree() *Tree {
	tree := &Tree{
		Root: &Node{0, map[rune]*Node{}, false, 0},
		Size: 0,
	}
	return tree
}

//AddNode adds a node to a parent node within the tree and returns a pointer to the added node
//If the parent node already has a child node with the given value it returns that child node instead
func (t *Tree) AddNode(parentNode *Node, Value rune) *Node {

	if parentNode.Children[Value] == nil {
		newChildNode := &Node{Value, map[rune]*Node{}, true, parentNode.Depth + 1}
		parentNode.Children[Value] = newChildNode
		t.Size++
		parentNode.End = false
		return newChildNode
	}

	return parentNode.Children[Value]

}

//AddWord adds a word to the tree by adding a node of each character if it doesnt exist
func (t *Tree) AddWord(word string) {
	currentNode := t.Root
	for _, letter := range word {
		currentNode = t.AddNode(currentNode, letter)
	}
}

//ClosestMatch traverses the tree and finds the closest match to the input within the tree
func (t *Tree) ClosestMatch(word string) string {

	match := ""

	currentNode := t.Root
	for _, letter := range word {
		currentNode = currentNode.Children[letter]
		if currentNode == nil {
			return match
		}

		match += string(currentNode.Value)

	}

	return match
}

// func main() {
// 	tree := NewTree()

// 	tree.AddWord("Hello")
// 	tree.AddWord("help")
// 	tree.AddWord("dolphin")
// 	tree.AddWord("hell")
// 	fmt.Println(tree.ClosestMatch("dolp"))

// 	j, _ := json.Marshal(tree)
// 	fmt.Println(string(j))

// 	var app express.App

// 	h := func(w http.ResponseWriter, req *http.Request, stop func(message string)) {
// 		express.GzipJSON(w, tree)
// 	}

// 	app.Get("/", h)

// 	app.Run(":8080")
// }
