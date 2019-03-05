package terminal

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ajnavarro/genprog/node"
)

var _ node.Terminal = &literal{}

type literal struct {
	node.BaseTerminal
	v interface{}
	t node.Type
}

func NewNumberLiteralFactory(min, max int) *node.FactoryNode {
	return &node.FactoryNode{
		Arity: 0,
		Factory: func(ns []node.Node) node.Node {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			return NewNumberLiteral(min + r.Intn(max-min))
		},
	}
}

func NewStringLiteral(v string) *literal {
	return &literal{v: v, t: node.String}
}

func NewNumberLiteral(v int) *literal {
	return &literal{v: v, t: node.Number}
}

func NewBooleanLiteral(v bool) *literal {
	return &literal{v: v, t: node.Boolean}
}

func (l *literal) Type() node.Type {
	return l.t
}

func (l *literal) Eval(ctx context.Context) interface{} {
	return l.v
}

func (l *literal) String() string {
	return fmt.Sprintf("%v", l.v)
}
