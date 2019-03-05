package utils

import (
	"math/rand"

	"github.com/ajnavarro/genprog/node"
)

func pickOneNode(rnd *rand.Rand, fn []node.Node) node.Node {
	idx := rnd.Intn(len(fn))
	return fn[idx]
}

func pickOneFactory(rnd *rand.Rand, fn []*node.FactoryNode) *node.FactoryNode {
	return fn[rnd.Intn(len(fn))]
}
