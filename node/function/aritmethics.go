package function

import (
	"context"

	"github.com/ajnavarro/genprog/node"
	"github.com/ajnavarro/genprog/utils"
	"github.com/spf13/cast"
)

var _ node.Function = &sum{}

var SumFactory = &node.FactoryNode{
	Arity: 2,
	Factory: func(ns []node.Node) node.Node {
		return NewSum(ns[0], ns[1])
	},
}

func NewSum(l, r node.Node) *sum {
	return &sum{node.BinaryFunction{L: l, R: r}}
}

type sum struct {
	node.BinaryFunction
}

func (f *sum) Type() node.Type {
	return node.Number
}

func (f *sum) Eval(ctx context.Context) interface{} {
	return cast.ToFloat64(f.L.Eval(ctx)) + cast.ToFloat64(f.R.Eval(ctx))
}

func (f *sum) String() string {
	t := utils.NewTreePrinter()
	t.WriteNode("ADD")
	t.WriteChildren(f.L.String(), f.R.String())
	return t.String()
}

var _ node.Function = &sub{}

var SubFactory = &node.FactoryNode{
	Arity: 2,
	Factory: func(ns []node.Node) node.Node {
		return NewSub(ns[0], ns[1])
	},
}

func NewSub(l, r node.Node) *sub {
	return &sub{node.BinaryFunction{L: l, R: r}}
}

type sub struct {
	node.BinaryFunction
}

func (f *sub) Type() node.Type {
	return node.Number
}

func (f *sub) Eval(ctx context.Context) interface{} {
	return cast.ToFloat64(f.L.Eval(ctx)) - cast.ToFloat64(f.R.Eval(ctx))
}

func (f *sub) String() string {
	t := utils.NewTreePrinter()
	t.WriteNode("SUBSTRACT")
	t.WriteChildren(f.L.String(), f.R.String())
	return t.String()

}

var _ node.Function = &multiply{}

var MultiplyFactory = &node.FactoryNode{
	Arity: 2,
	Factory: func(ns []node.Node) node.Node {
		return NewMultiply(ns[0], ns[1])
	},
}

func NewMultiply(l, r node.Node) *multiply {
	return &multiply{node.BinaryFunction{L: l, R: r}}
}

type multiply struct {
	node.BinaryFunction
}

func (f *multiply) Type() node.Type {
	return node.Number
}

func (f *multiply) Eval(ctx context.Context) interface{} {
	return cast.ToFloat64(f.L.Eval(ctx)) * cast.ToFloat64(f.R.Eval(ctx))
}

func (f *multiply) String() string {
	t := utils.NewTreePrinter()
	t.WriteNode("MULTIPLY")
	t.WriteChildren(f.L.String(), f.R.String())
	return t.String()

}

var _ node.Function = &divide{}

var DivideFactory = &node.FactoryNode{
	Arity: 2,
	Factory: func(ns []node.Node) node.Node {
		return NewDivide(ns[0], ns[1])
	},
}

func NewDivide(l, r node.Node) *divide {
	return &divide{node.BinaryFunction{L: l, R: r}}
}

type divide struct {
	node.BinaryFunction
}

func (f *divide) Type() node.Type {
	return node.Number
}

func (f *divide) Eval(ctx context.Context) interface{} {
	return cast.ToFloat64(f.L.Eval(ctx)) / cast.ToFloat64(f.R.Eval(ctx))
}

func (f *divide) String() string {
	t := utils.NewTreePrinter()
	t.WriteNode("DIVIDE")
	t.WriteChildren(f.L.String(), f.R.String())
	return t.String()

}
