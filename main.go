package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"

	"github.com/jinzhu/copier"

	"github.com/ajnavarro/genprog/node"
	"github.com/ajnavarro/genprog/node/function"
	"github.com/ajnavarro/genprog/node/terminal"
	"github.com/ajnavarro/genprog/utils"

	"github.com/MaxHalford/eaopt"
)

const maxDepth = 6

const (
	X terminal.Variable = iota
	Y
)

var vars = []terminal.Variable{X, Y}

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

var sumDataset = &utils.DataSet{
	From:     -5,
	To:       5,
	NSamples: 5,
	NVars:    2,
	Evaluator: func(ins ...interface{}) interface{} {
		v1 := ins[0].(int)
		v2 := ins[1].(int)
		out := v1 + v2

		return out
	},
}

var funcDataset = &utils.DataSet{
	From:     -5,
	To:       5,
	NSamples: 10,
	NVars:    3,
	Evaluator: func(ins ...interface{}) interface{} {
		v1 := ins[0].(int)
		v2 := ins[1].(int)
		return v1*v1 + v2 - 3
	},
}

var func2Dataset = &utils.DataSet{
	From:     -5,
	To:       5,
	NSamples: 10,
	NVars:    3,
	Evaluator: func(ins ...interface{}) interface{} {
		v1 := ins[0].(int)
		v2 := ins[1].(int)
		return v1*v1 + v2 - 3 + v1*v2
	},
}

var data = funcDataset.Generate()

type Gnome struct {
	n node.Node
	m utils.GenerationMethod
}

func (g *Gnome) Evaluate() (float64, error) {
	var absError float64
	for _, ds := range data {
		ctx := context.TODO()
		for i, v := range vars {
			ctx = context.WithValue(ctx, v, ds.Vars[i])
		}

		result := g.n.Eval(ctx)

		switch g.n.Type() {
		case node.String:
			panic("string type not implemented")
		case node.Boolean:
			panic("boolean type not implemented")
		case node.Number:
			dsResult := ds.Result.(int)
			gnomeResult := result.(int)
			abs := math.Abs(float64(dsResult - gnomeResult))
			absError = absError + abs
		}
	}
	return float64(absError), nil
}

func (g *Gnome) Mutate(rng *rand.Rand) {
	r := rng.Float64()
	random := utils.GenerateRandomTree(rng, ff, tf, maxDepth, g.m)
	branchingFactor := 1
	if r > 0.89 {
		// functions with branching factor
		branchingFactor = rng.Intn(maxDepth) + 1
	}
	node, err := utils.Crossover(rng, branchingFactor, g.n, random)
	if err != nil {
		panic(fmt.Errorf("Error on mutation", err))
	}

	g.n = node
}

func (g *Gnome) Crossover(genome eaopt.Genome, rng *rand.Rand) {
	gen := genome.(*Gnome)

	r := rng.Float64()
	branchingFactor := 1
	if r > 0.89 {
		// functions with branching factor
		branchingFactor = rng.Intn(maxDepth) + 1
	}
	node, err := utils.Crossover(rng, branchingFactor, g.n, gen.n)
	if err != nil {
		panic(fmt.Errorf("Error on crossover", err))
	}
	g.n = node
}

func (g *Gnome) Clone() eaopt.Genome {
	out := &Gnome{}
	copier.Copy(&out, &g)
	return out
}

func GenomeFactory(rng *rand.Rand) eaopt.Genome {
	genMethodChoice := rand.Float64()
	m := utils.GROW
	if genMethodChoice > 0.5 {
		m = utils.FULL
	}

	return &Gnome{
		m: m,
		n: utils.GenerateRandomTree(rng, ff, tf, maxDepth, m),
	}
}

func main() {
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	ga.NPops = 100
	ga.NGenerations = 10000
	ga.PopSize = 50
	ga.HofSize = 5
	ga.ParallelEval = true
	ga.Callback = func(ga *eaopt.GA) {
		var totalFitness float64 = 0
		for _, f := range ga.HallOfFame {
			totalFitness = totalFitness + f.Fitness
		}
		fmt.Printf("fitness at generation %d: %f\n", ga.Generations, totalFitness)
	}

	ga.EarlyStop = func(ga *eaopt.GA) bool {
		return ga.HallOfFame[0].Fitness == 0
	}

	err = ga.Minimize(GenomeFactory)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ga.HallOfFame[0].Genome.(*Gnome).n)

}
