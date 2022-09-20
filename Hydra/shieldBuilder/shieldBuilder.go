package shieldBuilder

import "strings"

// This module implements a Builder pattern in Go
type shield struct {
	front bool
	back  bool
	right bool
	left  bool
}

type shieldbuilder struct {
	code string
}

// constructor for a new shield builder
func NewShieldBuilder() *shieldbuilder {
	return new(shieldbuilder)
}

func (sb *shieldbuilder) RaiseFront() *shieldbuilder {
	sb.code += "F"
	return sb
}

func (sb *shieldbuilder) RaiseBack() *shieldbuilder {
	sb.code += "B"
	return sb
}

func (sb *shieldbuilder) RaiseRight() *shieldbuilder {
	sb.code += "R"
	return sb
}

func (sb *shieldbuilder) RaiseLeft() *shieldbuilder {
	sb.code += "L"
	return sb
}

func (sb *shieldbuilder) Build() *shield {
	code := sb.code
	return &shield{
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
		right: strings.Contains(code, "R"),
		left:  strings.Contains(code, "L"),
	}
}
