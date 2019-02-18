package utils_test

import (
	"fmt"
	"testing"

	"github.com/ajnavarro/genprog/utils"
)

func TestBranchingFactor(t *testing.T) {
	tree := utils.GenerateRandomTree(ff, tf, 2, utils.GROW)
	fmt.Println(tree)

	fmt.Println(utils.BranchingFactor(tree))
}

func TestCrossover(t *testing.T) {
	p1 := utils.GenerateRandomTree(ff, tf, 3, utils.FULL)
	p2 := utils.GenerateRandomTree(ff, tf, 3, utils.FULL)
	fmt.Println(p1)
	fmt.Println(":::::::::::::::::")
	fmt.Println(p2)
	cross, err := utils.Crossover(3, p1, p2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(":::::::::::::::::")
	fmt.Println(cross)
}
