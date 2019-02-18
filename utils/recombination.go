package utils

import (
	"github.com/ajnavarro/genprog/node"
	"github.com/jinzhu/copier"
)

func Crossover(branchingFactor int, p1, p2 node.Node) (node.Node, error) {
	var p1Copy node.Node
	if err := copier.Copy(&p1Copy, &p1); err != nil {
		return nil, err
	}

	var p2Copy node.Node
	if err := copier.Copy(&p2Copy, &p2); err != nil {
		return nil, err
	}

	return walk(branchingFactor, p1Copy, p2Copy), nil
}

func walk(bf int, p1, p2 node.Node) node.Node {
	bfp1 := BranchingFactor(p1)
	if bfp1 <= bf {
		return getCrossoverNode(bf, p2)
	}
	children := p1.Children()
	shuffle := getRandom().Perm(len(children))
	for _, idx := range shuffle {
		result := walk(bf, children[idx], p2)
		if result == nil {
			continue
		}

		children[idx] = result
		p1.SetChildren(children)
		return p1
	}

	return nil
}

func getCrossoverNode(bf int, p node.Node) node.Node {
	bfp := BranchingFactor(p)
	if bfp <= bf {
		return p
	}

	return getCrossoverNode(bf, pickOneNode(p.Children()))
}

func BranchingFactor(n node.Node) int {
	return branchingFactorRec(n, -1)
}

func branchingFactorRec(n node.Node, sum int) int {
	sum = sum + 1

	for _, node := range n.Children() {
		sum = branchingFactorRec(node, sum)
	}

	return sum
}
