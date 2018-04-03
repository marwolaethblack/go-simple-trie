package simpletrie

import "testing"

func TestAddNode(t *testing.T) {
	tree := NewTree()

	if tree == nil {
		t.Error()
	}

	newNode := tree.AddNode(tree.Root, 'v')

	if newNode == nil {
		t.Error()
	}
}

func TestAddWord(t *testing.T) {
	tree := NewTree()

	tree.AddWord("hello")

	if tree.Root.Children['h'].Children['e'].Children['l'].Children['l'].Children['o'] == nil {
		t.Error()
	}

}

func TestClosestMatch(t *testing.T) {
	tree := NewTree()

	tree.AddWord("Unilateral")

	if tree.ClosestMatch("Unil") != "Unil" {
		t.Error()
	}

	if tree.ClosestMatch("Unilateral") != "Unilateral" {
		t.Error()
	}
}

func BenchmarkClosestMatch(b *testing.B) {

	tree := NewTree()
	tree.AddWord("Hello")
	tree.AddWord("Help")
	tree.AddWord("Hell")
	tree.AddWord("Helicopter")

	for index := 0; index < b.N; index++ {
		tree.ClosestMatch("Hell")
	}
}
