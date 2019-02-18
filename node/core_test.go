package node_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ajnavarro/genprog/node"
	"github.com/ajnavarro/genprog/node/function"
	"github.com/ajnavarro/genprog/node/terminal"
)

const (
	X terminal.Variable = iota
	Y
	Z
)

var fixtures = []struct {
	Title  string
	Node   node.Node
	Result interface{}
}{
	{
		"multiply 2 variables",
		function.NewMultiply(
			terminal.NewNumberVariable(X),
			terminal.NewNumberVariable(Y),
		),
		4.00, //All numbers are floats
	},
	{
		"multiply and sum 3 variables",
		function.NewSum(
			function.NewMultiply(
				terminal.NewNumberVariable(X),
				terminal.NewNumberVariable(Y),
			),
			terminal.NewNumberVariable(Z),
		),
		7.00, //All numbers are floats
	},
}

func TestFixtures(t *testing.T) {
	ctx :=
		context.WithValue(
			context.WithValue(
				context.WithValue(
					context.TODO(),
					X, 2),
				Y, 2),
			Z, 3)
	for _, f := range fixtures {
		t.Run(f.Title, func(subt *testing.T) {
			fmt.Println(f.Node)
			r := f.Node.Eval(ctx)
			if f.Result != r {
				subt.Log(f.Result, "!=", r)
				subt.Fail()
			}
		})
	}
}
