package utils_test

import (
	"fmt"
	"testing"

	"github.com/ajnavarro/genprog/node"
	"github.com/ajnavarro/genprog/node/function"
	"github.com/ajnavarro/genprog/node/terminal"
	"github.com/ajnavarro/genprog/utils"
)

const (
	X terminal.Variable = iota
	Y
)

var ff = []*node.FactoryNode{
	function.SumFactory,
	function.SubFactory,
	function.DivideFactory,
	function.MultiplyFactory,
	//TODO function.IfElseFactory,
}

var tf = []*node.FactoryNode{
	terminal.NewNumberLiteralFactory(-5, 5),
	terminal.NewNumberVariableFactory(X, Y),
}

func TestGenerateRandomTreeGrow(t *testing.T) {
	tree := utils.GenerateRandomTree(ff, tf, 3, utils.GROW)
	fmt.Println(tree)
}

func TestGenerateRandomTreeFull(t *testing.T) {
	tree := utils.GenerateRandomTree(ff, tf, 5, utils.FULL)
	fmt.Println(tree)
}
