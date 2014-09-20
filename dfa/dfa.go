package dfa

import (
	"fmt"
	"strings"
)

type Node struct {
	Terminal   bool
	Out_edges  map[rune]*Node
	Out_uncond *Node
}

func (n *Node) Next(r rune) (*Node, bool) {
	next, ok := n.Out_edges[r]
	if ok {
		return next, true
	}
	return n.Out_uncond, n.Out_uncond != nil
}

func (n *Node) MatchReader(reader *strings.Reader) bool {
	if reader.Len() == 0 && n.Terminal {
		return true
	}

	r, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println("Bad rune")
	}

	next, ok := n.Next(r)

	if !ok {
		return false
	}

	return next.MatchReader(reader)
}

func (n *Node) MatchString(s string) bool {
	return n.MatchReader(strings.NewReader(s))
}

func NewNode() *Node {
	n := new(Node)

	n.Out_edges = make(map[rune]*Node)
	return n
}
