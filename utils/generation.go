package utils

import (
	"math/rand"

	"github.com/ajnavarro/genprog/node"
)

type GenerationMethod int

const (
	GROW GenerationMethod = iota
	FULL
)

func GenerateRandomTree(
	rnd *rand.Rand,
	functions []*node.FactoryNode,
	terminals []*node.FactoryNode,
	maxDepth int,
	method GenerationMethod,
) node.Node {
	termLen := float64(len(terminals))
	funcLen := float64(len(functions))
	r := rnd.Float64()
	check := termLen / (termLen + funcLen)

	if maxDepth == 0 || (method == GROW && r < check) {
		return pickOneFactory(rnd, terminals).Factory(nil)
	}

	function := pickOneFactory(rnd, functions)
	var nodes []node.Node
	for i := 0; i < function.Arity; i++ {
		nodes = append(nodes, GenerateRandomTree(rnd, functions, terminals, maxDepth-1, method))
	}

	return function.Factory(nodes)
}
