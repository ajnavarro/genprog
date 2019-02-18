package node

import (
	"context"
)

type FactoryNode struct {
	Arity   int
	Factory func([]Node) Node
}

type Node interface {
	Type() Type
	Eval(ctx context.Context) interface{}
	String() string
	Depth() int
	Children() []Node
	SetChildren([]Node)
}

type Function interface {
	Node
}

type BinaryFunction struct {
	Function
	L, R Node
}

func (f *BinaryFunction) Children() []Node {
	return []Node{f.L, f.R}
}

func (f *BinaryFunction) SetChildren(ch []Node) {
	if len(ch) != 2 {
		panic("BinaryFunction only accepts len==2 children arrays")
	}

	f.L = ch[0]
	f.R = ch[1]
}

func (f *BinaryFunction) Depth() int {
	var cd = 0
	if f.L.Depth() > f.R.Depth() {
		cd = f.L.Depth()
	} else {
		cd = f.R.Depth()
	}

	return cd + 1
}

type Terminal interface {
	Node
}

type BaseTerminal struct {
	Terminal
}

func (BaseTerminal) Children() []Node {
	return nil
}

func (BaseTerminal) SetChildren(ch []Node) {
}

type Type int

const (
	Invalid Type = iota
	Number
	String
	Boolean
)
