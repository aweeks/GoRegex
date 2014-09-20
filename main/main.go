package main

import (
	"dfa"
	"fmt"
	"parser"
)

func main() {

	n1 := dfa.NewNode()
	n2 := dfa.NewNode()
	n3 := dfa.NewNode()

	fmt.Println(n1)
	n1.Out_edges['a'] = n2
	n2.Out_edges['b'] = n2
	n2.Out_edges['a'] = n3
	n2.Out_uncond = n1
	n3.Terminal = true

	fmt.Println(n1.MatchString("abbbcba"))

	fmt.Println(parser.NewParser("ab*").ParseRegex())
	fmt.Println(parser.NewParser("ab*c").ParseRegex())
	fmt.Println(parser.NewParser("ab*|c").ParseRegex())
	fmt.Println(parser.NewParser("(ab)*").ParseRegex())
	fmt.Println(parser.NewParser("(ab)*|cd").ParseRegex())

}
