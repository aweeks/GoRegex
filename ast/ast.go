package ast

import (
	"fmt"
)

type Regex interface {
}

/* Null Statement*/

type RegexNull struct {
}

func NewRegexNull() Regex {
	rn := new(RegexNull)
	return rn
}

func (rn RegexNull) String() string {
	return "<null>"
}

/* Primitive Statement */

type RegexPrimitive struct {
	r rune
}

func NewRegexPrimitive(r rune) Regex {
	rp := new(RegexPrimitive)
	rp.r = r
	return rp
}

func (rp RegexPrimitive) String() string {
	return fmt.Sprintf("<%q>", rp.r)
}

/* Sequence Statement */

type RegexConcat struct {
	first  Regex
	second Regex
}

func NewRegexConcat(first Regex, second Regex) Regex {
	rc := new(RegexConcat)
	rc.first = first
	rc.second = second
	return rc
}

func (rc RegexConcat) String() string {
	return fmt.Sprintf("<Concat %s %s>", rc.first, rc.second)
}

/* Or Statement */

type RegexOr struct {
	left  Regex
	right Regex
}

func NewRegexOr(left Regex, right Regex) Regex {
	ro := new(RegexOr)
	ro.left = left
	ro.right = right
	return ro
}

func (ro RegexOr) String() string {
	return fmt.Sprintf("<Or %s %s>", ro.left, ro.right)
}

/* Star Statement */

type RegexStar struct {
	regex Regex
}

func NewRegexStar(regex Regex) Regex {
	rs := new(RegexStar)
	rs.regex = regex
	return rs
}

func (rs RegexStar) String() string {
	return fmt.Sprintf("<Star %s>", rs.regex)
}
