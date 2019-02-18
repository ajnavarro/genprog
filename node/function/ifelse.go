package function

import (
	"context"

	"github.com/ajnavarro/genprog/utils"

	"github.com/ajnavarro/genprog/node"
	"github.com/spf13/cast"
)

var _ node.Function = &ifelse{}

var IfElseFactory = &node.FactoryNode{
	Arity: 4,
	Factory: func(ns []node.Node) node.Node {
		return NewIfElse(ns[0], ns[1], ns[2], ns[3])
	},
}

type ifelse struct {
	Test, Yes, No, Child node.Node
}

func NewIfElse(test, yes, no, child node.Node) *ifelse {
	return &ifelse{test, yes, no, child}
}

func (f *ifelse) Type() node.Type {
	return node.Boolean
}

func (f *ifelse) Eval(ctx context.Context) interface{} {
	if cast.ToBool(f.Test.Eval(ctx)) {
		f.Yes.Eval(ctx)
	} else {
		if f.No != nil {
			f.No.Eval(ctx)
		}
	}

	return f.Child.Eval(ctx)
}

func (f *ifelse) Children() []node.Node {
	return []node.Node{f.Test, f.Yes, f.No, f.Child}
}

func (f *ifelse) SetChildren(ch []node.Node) {
	if len(ch) != 4 {
		panic("ifelse node panic on SetChildren")
	}

	f.Test = ch[0]
	f.Yes = ch[1]
	f.No = ch[2]
	f.Child = ch[3]
}

func (f *ifelse) Depth() int {
	var max = 0
	for _, v := range []node.Node{f.Test, f.Yes, f.No, f.Child} {
		if max < v.Depth() {
			max = v.Depth()
		}
	}

	return max + 1
}

func (f *ifelse) String() string {
	t := utils.NewTreePrinter()
	t.WriteNode("IF(%s)", f.Test.String())
	t.WriteChildren(f.Yes.String(), f.No.String(), f.Child.String())
	return t.String()
}
