package terminal

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ajnavarro/genprog/node"
)

type Variable int

var _ node.Terminal = &variable{}

type variable struct {
	node.BaseTerminal
	k Variable
	t node.Type
}

func NewNumberVariableFactory(vars ...Variable) *node.FactoryNode {
	return &node.FactoryNode{
		Arity: 0,
		Factory: func(ns []node.Node) node.Node {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			max := len(vars)
			return NewNumberVariable(vars[r.Intn(max)])
		},
	}
}

func NewStringVariable(k Variable) *variable {
	return &variable{k: k, t: node.String}
}

func NewNumberVariable(k Variable) *variable {
	return &variable{k: k, t: node.Number}
}

func NewBooleanVariable(k Variable) *variable {
	return &variable{k: k, t: node.Boolean}
}

func (l *variable) Type() node.Type {
	return l.t
}

func (l *variable) Eval(ctx context.Context) interface{} {
	return ctx.Value(l.k)
}

func (l *variable) String() string {
	return fmt.Sprintf("var_%v", l.k)
}
